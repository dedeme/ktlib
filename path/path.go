// Copyright 26-May-2017 ºDeme
// GNU General Public License - V3 <http://www.gnu.org/licenses/>

// Functions to manage file paths.
package path

import (
	gpath "path"
)

// Returns the last component of path 'fpath'.
//
// If the path is empty, Base returns "."
//
// If the path consists entirely of slashes, Base returns "/".
func Base(fpath string) string {
	return gpath.Base(fpath)
}

// Removes duplicates '/' and redundant '..'.
//
// The returns ends in a slash only if it is the root "/".
//
// If the result were an empty strint, it returns ".".
func Canonical(fpath string) string {
	return gpath.Clean(fpath)
}

// Join elements of 'paths' with "/".
//
// The result is made canonical.
func Cat(paths ...string) string {
	return gpath.Join(paths...)
}

// Returns the extension of path 'fpath', including the dot, or an empty string
// if 'fpath' has not extension.
func Extension(fpath string) string {
	return gpath.Ext(fpath)
}

// Returns the parent directory or path 'fpath'.
//
// If there is no parent directory, it returns ".".
func Parent(fpath string) string {
	return gpath.Dir(fpath)
}
