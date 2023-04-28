/** @module infra-sk/modules/login */

import { LoginStatus } from './json';

/**
 * A Promise that will be resolved with the users current login status.
 *
 * The resolution object looks like:
 * <pre>
 *   {
 *     "Email": "fred@example.com",
 *     "LoginURL": "https://..."
 *   }
 * </pre>
 *
 * The Email will be the empty string if the user is not logged in.
 */
export const Login: Promise<LoginStatus> = fetch('/loginstatus/', {
  credentials: 'include',
}).then((res) => {
  if (res.ok) {
    return res.json();
  }
  throw new Error(`Problem reading /loginstatus/:${res.statusText}`);
});

/**
 * A function that returns a Promise that will be resolved with the users
 * current login status by accessing the specified loginStatusURL.
 *
 * The resolution object looks like:
 * <pre>
 *   {
 *     "Email": "fred@example.com",
 *     "LoginURL": "https://..."
 *   }
 * </pre>
 *
 * The Email will be the empty string if the user is not logged in.
 */
export const LoginTo = function (loginStatusURL: string): Promise<LoginStatus> {
  return fetch(loginStatusURL, {
    credentials: 'include',
  }).then((res) => {
    if (res.ok) {
      return res.json();
    }
    throw new Error(`Problem reading /loginstatus/:${res.statusText}`);
  });
};

// Add to the global sk namespace while we migrate away from Polymer.
if ((window as any).sk !== undefined) {
  (window as any).sk.Login = Login;
}

/**
 * Function that returns the root domain of a sub-domain.
 *
 * I.e. it will return "skia.org" if the current location is "perf.skia.org".
 *
 * In addition it will fallback to "skia.org" is case we are on corp.goog.
 */
export function rootDomain(): string {
  let ret = window.location.host.split('.').slice(-2).join('.');
  if (ret === 'corp.goog') {
    ret = 'skia.org';
  }
  return ret;
}
