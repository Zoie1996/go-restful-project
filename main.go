package main

import (
	"fmt"
	"log"
	"net/http"

	"go-restful-project/models"
	"go-restful-project/pkg/gredis"
	"go-restful-project/pkg/logging"
	"go-restful-project/pkg/setting"
	"go-restful-project/routers"
)

func init() {
	setting.Setup()
	models.Setup()
	logging.Setup()
	gredis.Setup()
}

func main() {
	// restful.DefaultRequestContentType(restful.MIME_JSON)
	// restful.DefaultResponseContentType(restful.MIME_JSON)
	wsContainer := routers.InitRouter()
	readTimeout := setting.ServerSetting.ReadTimeout
	writeTimeout := setting.ServerSetting.WriteTimeout
	endPoint := fmt.Sprintf(":%d", setting.ServerSetting.HttpPort)
	maxHeaderBytes := 1 << 20

	server := &http.Server{
		Addr:           endPoint,
		Handler:        wsContainer,
		ReadTimeout:    readTimeout,
		WriteTimeout:   writeTimeout,
		MaxHeaderBytes: maxHeaderBytes,
	}

	// log.Print("start listening on localhost:8080")
	log.Printf("[info] start http server listening %s", endPoint)
	// defer server.Close()
	log.Fatal(server.ListenAndServe())

}
