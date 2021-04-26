package tipitaka

import (
	"path/filepath"
	"strings"
)

// ActionToUrlPath converts action string to url path.
func ActionToUrlPath(action string) string {
	noext := strings.TrimSuffix(action, filepath.Ext(action))
	return "/" + strings.Replace(noext, ".", "/", -1) + "/"
}
