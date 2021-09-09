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
