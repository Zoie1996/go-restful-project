package util

import (
	"github.com/Unknwon/com"
	"github.com/emicklei/go-restful"

	"go-restful-project/pkg/setting"
)

// GetPage 获取page参数
func GetPage(request *restful.Request) int {
	result := 0
	page := com.StrTo(request.PathParameter("page")).MustInt()
	if page > 0 {
		result = (page - 1) * GetPageSize(request)
	}

	return result
}

// GetPageSize 获取pageSize参数
func GetPageSize(request *restful.Request) int {
	//设置默认为配置文件中的pageSize大小
	result := setting.AppSetting.PageSize
	pageSize := request.PathParameter("page_size")
	if pageSize != "" {
		pageSize := com.StrTo(request.PathParameter("page_size")).MustInt()
		//如果传入参数大于0小于配置文件pageSize大小则设置为传入参数
		if result-pageSize > 0 {
			result = pageSize
		}
	}

	return result
}
