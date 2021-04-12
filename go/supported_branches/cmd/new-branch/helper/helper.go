package helper

import (
	"bytes"
	"context"
	"fmt"
	"io/ioutil"
	"os"
	"regexp"
	"strings"

	cq_config "go.chromium.org/luci/cv/api/config/v2"

	"go.skia.org/infra/go/auth"
	"go.skia.org/infra/go/cq"
	"go.skia.org/infra/go/gerrit"
	"go.skia.org/infra/go/git"
	"go.skia.org/infra/go/gitiles"
	"go.skia.org/infra/go/httputils"
	"go.skia.org/infra/go/skerr"
	"go.skia.org/infra/go/supported_branches"
	"go.skia.org/infra/go/util"
)

// AddSupportedBranch adds a new supported branch, optionally deleting an old
// supported branch.
func AddSupportedBranch(repoUrl, branch, owner, deleteBranch string, excludeTrybots []string, submit bool) error {
	newRef := git.FullyQualifiedBranchName(branch)
	excludeTrybotRegexp := make([]*regexp.Regexp, 0, len(excludeTrybots))
	for _, excludeTrybot := range excludeTrybots {
		re, err := regexp.Compile(excludeTrybot)
		if err != nil {
			return skerr.Wrapf(err, "failed to compile regular expression from %q", excludeTrybot)
		}
		excludeTrybotRegexp = append(excludeTrybotRegexp, re)
	}

	// Setup.
	wd, err := ioutil.TempDir("", "new-branch")
	if err != nil {
		return skerr.Wrap(err)
	}
	defer util.RemoveAll(wd)

	ts, err := auth.NewDefaultTokenSource(true, auth.SCOPE_GERRIT)
	if err != nil {
		return skerr.Wrap(err)
	}
	client := httputils.DefaultClientConfig().WithTokenSource(ts).Client()
	gUrl := strings.Split(repoUrl, ".googlesource.com")[0] + "-review.googlesource.com"
	g, err := gerrit.NewGerrit(gUrl, client)
	if err != nil {
		return skerr.Wrap(err)
	}
	repo := gitiles.NewRepo(repoUrl, client)
	ctx := context.Background()
	baseCommitInfo, err := repo.Details(ctx, cq.CQ_CFG_REF)
	if err != nil {
		return skerr.Wrap(err)
	}
	baseCommit := baseCommitInfo.Hash

	// Download the CQ config file and modify it.
	cfgContents, err := repo.ReadFileAtRef(ctx, cq.CQ_CFG_FILE, baseCommit)
	if err != nil {
		return skerr.Wrap(err)
	}
	newCfgBytes, err := cq.WithUpdateCQConfig(cfgContents, func(cfg *cq_config.Config) error {
		cg, _, _, err := cq.MatchConfigGroup(cfg, newRef)
		if err != nil {
			return skerr.Wrap(err)
		}
		if cg != nil {
			_, _ = fmt.Fprintf(os.Stderr, "Already have %s in %s; not adding a duplicate.\n", newRef, cq.CQ_CFG_FILE)
		} else {
			if err := cq.CloneBranch(cfg, git.DefaultBranch, git.BranchBaseName(branch), false, false, excludeTrybotRegexp); err != nil {
				return skerr.Wrap(err)
			}
		}
		if deleteBranch != "" {
			if err := cq.DeleteBranch(cfg, deleteBranch); err != nil {
				return skerr.Wrap(err)
			}
		}
		return nil
	})

	// Download and modify the supported-branches.json file.
	branchesContents, err := repo.ReadFileAtRef(ctx, supported_branches.SUPPORTED_BRANCHES_FILE, baseCommit)
	if err != nil {
		return skerr.Wrap(err)
	}
	sbc, err := supported_branches.DecodeConfig(bytes.NewReader(branchesContents))
	if err != nil {
		return skerr.Wrap(err)
	}
	deleteRef := ""
	if deleteBranch != "" {
		deleteRef = git.FullyQualifiedBranchName(deleteBranch)
	}
	foundNewRef := false
	newBranches := make([]*supported_branches.SupportedBranch, 0, len(sbc.Branches)+1)
	for _, sb := range sbc.Branches {
		if deleteRef == "" || deleteRef != sb.Ref {
			newBranches = append(newBranches, sb)
		}
		if sb.Ref == newRef {
			foundNewRef = true
		}
	}
	if foundNewRef {
		_, _ = fmt.Fprintf(os.Stderr, "Already have %s in %s; not adding a duplicate.\n", newRef, supported_branches.SUPPORTED_BRANCHES_FILE)
	} else {
		newBranches = append(newBranches, &supported_branches.SupportedBranch{
			Ref:   newRef,
			Owner: owner,
		})
	}
	sbc.Branches = newBranches
	buf := bytes.Buffer{}
	if err := supported_branches.EncodeConfig(&buf, sbc); err != nil {
		return skerr.Wrap(err)
	}

	// Create the Gerrit CL.
	commitMsg := fmt.Sprintf("Add supported branch %s", branch)
	if deleteBranch != "" {
		commitMsg += fmt.Sprintf(", remove %s", deleteBranch)
	}
	repoSplit := strings.Split(repoUrl, "/")
	project := strings.TrimSuffix(repoSplit[len(repoSplit)-1], ".git")
	changes := map[string]string{
		cq.CQ_CFG_FILE: string(newCfgBytes),
		supported_branches.SUPPORTED_BRANCHES_FILE: string(buf.Bytes()),
	}
	ci, err := gerrit.CreateCLWithChanges(ctx, g, project, cq.CQ_CFG_REF, commitMsg, baseCommit, changes, submit)
	if ci != nil {
		fmt.Println(fmt.Sprintf("Uploaded change %s", g.Url(ci.Issue)))
	}
	return skerr.Wrap(err)
}