package main

import (
	"api_server/core"
	"api_server/global"
	"api_server/internal/cache"
	"api_server/internal/models"
	"api_server/internal/routers"
	"fmt"
	"net/http"
	"time"
)

// @title API_Server
// @version 1.0
// @description 没啥好说的
// @termsOfService http://swagger.io/terms/

// @contact.name jinsoft
// @contact.url http://www.swagger.io/support
// @contact.email job@ainiok.com

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @host 127.0.0.1:8300
// @BasePath /
func main() {
	global.CFG = core.Viper()
	global.LOG = core.Zap()
	global.DB = models.NewDBEngine()
	if global.DB != nil {
		db, _ := global.DB.DB()
		defer db.Close()
	}
	cache.Redis()
	//gin.SetMode("debug")
	router := routers.NewRouter()
	s := &http.Server{
		Addr:           fmt.Sprintf(":%d", global.CONFIG.System.HttpPort),
		Handler:        router,
		ReadTimeout:    10 * time.Millisecond,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}

	if err := s.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		fmt.Printf("s.ListenAndServe err: %v", err)
	}
}
