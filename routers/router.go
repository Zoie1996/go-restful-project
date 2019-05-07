package routers

import (
	"flag"
	"go-restful-project/service/user_service"
	"log"

	"github.com/emicklei/go-restful"
	swagger "github.com/emicklei/go-restful-swagger12"
	"github.com/magiconair/properties"
)

var (
	props          *properties.Properties
	propertiesFile = flag.String("config", "config.properties", "the configuration file")

	SwaggerPath string
)

func InitRouter() *restful.Container {
	flag.Parse()

	// Load configurations from a file
	log.Printf("loading configuration from [%s]", *propertiesFile)
	var err error
	if props, err = properties.LoadFile(*propertiesFile, properties.UTF8); err != nil {
		log.Fatalf("[error] Unable to read properties:%v\n", err)
	}

	addr := props.MustGet("http.server.host") + ":" + props.MustGet("http.server.port")
	basePath := "http://" + addr

	wsContainer := restful.NewContainer()

	// 跨域过滤器
	cors := restful.CrossOriginResourceSharing{
		ExposeHeaders:  []string{"X-My-Header"},
		AllowedHeaders: []string{"Content-Type","Accept"},
		AllowedMethods: []string{"GET", "POST"},
		CookiesAllowed: false,
		Container:      wsContainer}
	wsContainer.Filter(cors.Filter)

	// Add container filter to respond to OPTIONS
	wsContainer.Filter(wsContainer.OPTIONSFilter)

	// wsContainer.Router(restful.CurlyRouter{})

	user_service.UsersRegister(wsContainer)

	config := swagger.Config{
		WebServices:    wsContainer.RegisteredWebServices(), // you control what services are visible
		WebServicesUrl: basePath,
		ApiPath:        props.GetString("api.path", ""),
		// ApiVersion:     "v2.0",
		// Optionally, specify where the UI is located
		SwaggerPath:     props.GetString("swagger.path", ""),
		SwaggerFilePath: props.GetString("swagger.file.path", "")}
	swagger.RegisterSwaggerService(config, wsContainer)
	//swagger.InstallSwaggerService(config)

	return wsContainer
}
