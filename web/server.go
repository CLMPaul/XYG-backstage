package web

import (
	"context"
	"errors"
	"fmt"
	"net"
	"net/http"
	"time"
	"xueyigou_demo/api"
	"xueyigou_demo/config"
	router2 "xueyigou_demo/router"
	"xueyigou_demo/static"

	"github.com/gin-contrib/cors"
	"github.com/gin-contrib/gzip"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"xueyigou_demo/internal/utils"
	"xueyigou_demo/internal/web/middlewares"
)

func NewServer(ctx context.Context) *http.Server {
	if logrus.StandardLogger().Level >= logrus.DebugLevel {
		gin.SetMode(gin.DebugMode)
	}
	gin.DefaultWriter = logrus.StandardLogger().WriterLevel(logrus.InfoLevel)
	gin.DefaultErrorWriter = logrus.StandardLogger().WriterLevel(logrus.ErrorLevel)
	router := gin.New()
	router.Use(gin.Logger())
	router.Use(middlewares.Recovery)
	router.Use(gzip.Gzip(gzip.DefaultCompression))

	//router.Use(middlewares.Tracing)
	if len(config.Config.Web.AllowOrigins) > 0 {
		router.Use(cors.New(cors.Config{
			AllowCredentials: true,
			AllowOrigins:     config.Config.Web.AllowOrigins,
			AllowHeaders: []string{
				"Origin", "Content-Length", "Content-Type", "Authentication",
				"X-WebAppID", "X-Page-Token",
			},
			AllowMethods: []string{"GET", "POST", "PUT", "PATCH", "DELETE", "HEAD"},
			MaxAge:       12 * time.Hour,
		}))
	}

	api.SetupRouter(router)
	// TODO 重构合并
	router2.InitRouter(router)
	router.NoRoute(static.NoRoute)

	return &http.Server{
		Addr:              fmt.Sprintf(":%d", config.Config.Web.HttpPort),
		Handler:           router,
		ReadHeaderTimeout: 20 * time.Second,
		ReadTimeout:       200 * time.Second,
		WriteTimeout:      90 * time.Second,
		BaseContext:       func(_ net.Listener) context.Context { return ctx },
	}
}

func RunServer(ctx context.Context) {
	server := NewServer(ctx)
	defer func() {
		_ = server.Shutdown(context.Background())
		logrus.Infoln("http server stopped")
	}()

	go func() {
		logrus.Infof("running http server on port %d", config.Config.Web.HttpPort)
		//err := server.ListenAndServe()
		err := server.ListenAndServeTLS("./config/ssl/8907110_xueyigou.cn.pem", "./config/ssl/8907110_xueyigou.cn.key") // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
		if !errors.Is(err, http.ErrServerClosed) {
			logrus.Errorf("failed starting http server: %v", err)
			utils.StopApp(ctx)
		}
	}()

	<-ctx.Done()
}

func init() {
	gin.SetMode(gin.ReleaseMode)
}
