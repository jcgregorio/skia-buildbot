package main

import (
	"path"
	"runtime"

	"go.skia.org/infra/go/auth"
	"go.skia.org/infra/go/gce"
	"go.skia.org/infra/go/gce/server"
)

func AutoRollBase(name, ipAddress string) *gce.Instance {
	vm := server.AddGitConfigs(server.Server20170518(name), name)
	vm.DataDisk.SizeGb = 64
	vm.DataDisk.Type = gce.DISK_TYPE_PERSISTENT_STANDARD
	vm.ExternalIpAddress = ipAddress
	vm.MachineType = gce.MACHINE_TYPE_STANDARD_2
	vm.Metadata["owner_primary"] = "borenet"
	vm.Metadata["owner_secondary"] = "rmistry"
	vm.Scopes = append(vm.Scopes,
		auth.SCOPE_GERRIT,
		auth.SCOPE_USERINFO_EMAIL,
		auth.SCOPE_USERINFO_PROFILE,
	)
	return vm
}

func Skia() *gce.Instance {
	return AutoRollBase("skia-autoroll", "104.154.112.12")
}

func SkiaInternal() *gce.Instance {
	return AutoRollBase("skia-internal-autoroll", "104.154.112.129")
}

func Catapult() *gce.Instance {
	return AutoRollBase("catapult-autoroll", "104.154.112.121")
}

func NaCl() *gce.Instance {
	return AutoRollBase("nacl-autoroll", "104.154.112.123")
}

func PDFium() *gce.Instance {
	return AutoRollBase("pdfium-autoroll", "104.154.123.210")
}

func AddAndroidConfigs(vm *gce.Instance) *gce.Instance {
	vm.DataDisk.SizeGb = 512
	vm.MachineType = gce.MACHINE_TYPE_HIGHMEM_16
	vm.Scopes = append(vm.Scopes, "https://www.googleapis.com/auth/androidbuild.internal")

	_, filename, _, _ := runtime.Caller(0)
	dir := path.Dir(filename)
	vm.SetupScript = path.Join(dir, "setup-script-android.sh")
	return vm
}

func AndroidMaster() *gce.Instance {
	return AddAndroidConfigs(AutoRollBase("android-master-autoroll", "104.154.123.206"))
}

func AndroidO() *gce.Instance {
	return AddAndroidConfigs(AutoRollBase("android-o-autoroll", "104.154.123.208"))
}

func Test() *gce.Instance {
	vm := AutoRollBase("borenet-instance-creation-test", "")
	server.AddGitConfigs(vm, "skia-autoroll") // Hack. Copy the skia-autoroll creds.
	return vm
}

func main() {
	server.Main(map[string]*gce.Instance{
		"skia":           Skia(),
		"skia-internal":  SkiaInternal(),
		"catapult":       Catapult(),
		"nacl":           NaCl(),
		"pdfium":         PDFium(),
		"android-master": AndroidMaster(),
		"android-o":      AndroidO(),
		"test":           Test(),
	})
}
