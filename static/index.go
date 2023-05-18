package static

import (
	"embed"
	"github.com/gin-gonic/gin"
	"io"
	"io/fs"
	"net/http"
	"strings"
)

//go:embed all:dist
var staticFS embed.FS

func SetupStatic(router gin.IRouter) {
	staticSubFS, _ := fs.Sub(staticFS, "static")
	httpFS := http.FS(staticSubFS)
	router.StaticFS("/static", httpFS)
}

func NoRoute(c *gin.Context) {
	if c.IsAborted() {
		return
	}
	if strings.HasPrefix(c.Request.URL.Path, "/static") {
		c.AbortWithStatus(http.StatusOK)
		reader, _ := staticFS.Open("static/index.html")
		io.Copy(c.Writer, reader)
	}
}
