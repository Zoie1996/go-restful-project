package routers

import (
	"go-restful-project/pkg/setting"
	"go-restful-project/service/user_service"

	"github.com/emicklei/go-restful"
	swagger "github.com/emicklei/go-restful-swagger12"
)

func InitRouter() *restful.Container {

	wsContainer := restful.NewContainer()

	// 跨域过滤器
	cors := restful.CrossOriginResourceSharing{
		ExposeHeaders:  []string{"X-My-Header"},
		AllowedHeaders: []string{"Content-Type", "Accept"},
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
		WebServicesUrl: setting.AppSetting.PrefixUrl,
		ApiPath:        setting.SwaggerSetting.ApiPath,
		// ApiVersion:     "v2.0",
		// Optionally, specify where the UI is located
		SwaggerPath:     setting.SwaggerSetting.SwaggerPath,
		SwaggerFilePath: setting.SwaggerSetting.SwaggerFilePath}
	swagger.RegisterSwaggerService(config, wsContainer)
	//swagger.InstallSwaggerService(config)

	return wsContainer
}
