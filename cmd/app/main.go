package main

import (
	"flag"
	"github.com/julienschmidt/httprouter"
	_ "github.com/maksgru/cryptogram-server/docs"
	"github.com/maksgru/cryptogram-server/internal/adapters/api"
	"github.com/maksgru/cryptogram-server/internal/config"
	httpSwagger "github.com/swaggo/http-swagger"
	"io"
	"log"
	"net/http"
)

var (
	configPath string
)

func init() {
	flag.StringVar(&configPath, "config-path", "configs/api.yml", "path to config file")
}

func main() {
	flag.Parse()
	cfg := config.GetConfig()
	router := httprouter.New()

	router.Handler(http.MethodGet, "/swagger", http.RedirectHandler("/swagger/index.html", http.StatusMovedPermanently))
	router.Handler(http.MethodGet, "/swagger/*any", httpSwagger.WrapHandler)

	api.NewApi(*router)

	router.GET("/hello", func(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
		io.WriteString(writer, "Hello from server")
	})

	err := http.ListenAndServe(cfg.APIServer.BindAddr, router)
	if err != nil {
		log.Fatal(err)
	}
}
