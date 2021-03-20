package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
	"user_server/core"
	"user_server/global"
	"user_server/internal/cache"
	"user_server/internal/models"
	"user_server/internal/routers"
)

func main() {
	global.CFG = core.Viper()
	global.LOG = core.Zap()
	global.DB = models.NewDBEngine()
	fmt.Println(global.CONFIG.System.HttpPort)
	if global.DB != nil {
		db, _ := global.DB.DB()
		defer db.Close()
	}
	cache.Redis()
	gin.SetMode("debug")
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
