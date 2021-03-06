package fs

import (
	"net/http"

	"github.com/7sDream/rikka/common/util"
	"github.com/7sDream/rikka/plugins"
)

const (
	fileURLPath = "/files/"
)

// ExtraHandlers return value will be add to http handle list.
// In fs plugin, we start a static file server to serve image file we accepted in /files/taskID path.
func (fsp fsPlugin) ExtraHandlers() (handlers []plugins.HandlerWithPattern) {
	// only accept GET method
	requestFilterFileServer := util.RequestFilter(
		"", "GET", l,
		// disable list dir
		util.DisableListDir(
			l,
			// Strip prefix path
			http.StripPrefix(
				// reserve last /
				fileURLPath[:len(fileURLPath)-1],
				// get a base file server
				http.FileServer(http.Dir(imageDir)),
			).ServeHTTP,
		),
	)

	handlers = []plugins.HandlerWithPattern{
		{
			Pattern: fileURLPath, Handler: requestFilterFileServer,
		},
	}

	return handlers
}
