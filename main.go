package main

import (
	"demo/global"
	"demo/internal/routers"
	"demo/pkg/setting"
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

func init() {
	if err := setupSetting(); err != nil {
		log.Fatalf("init.setuSetting err: %v", err)
	}
}

func main() {
	gin.SetMode(global.ServerSetting.RunMode)
	router := routers.NewRouter()
	s := &http.Server{
		Addr:           ":" + global.ServerSetting.HttpPort,
		Handler:        router,
		ReadTimeout:    global.ServerSetting.ReadTimeout,
		WriteTimeout:   global.ServerSetting.WriteTimeout,
		MaxHeaderBytes: 1 << 20,
	}
	s.ListenAndServe()
}

func setupSetting() error {
	setting, err := setting.NewSetting()
	if err != nil {
		return err
	}

	if err := setting.ReadSection("Server", &global.ServerSetting); err != nil {
		return err
	}

	if err := setting.ReadSection("App", &global.AppSetting); err != nil {
		return err
	}

	if err := setting.ReadSection("Database", &global.DatabaseSetting); err != nil {
		return err
	}

	global.ServerSetting.ReadTimeout *= time.Second
	global.ServerSetting.WriteTimeout *= time.Second

	return nil
}
