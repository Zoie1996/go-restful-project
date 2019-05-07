package app

import (
	"encoding/json"
	"go-restful-project/pkg/e"
	"log"
)

// type ErrResponse struct {
// 	ErrorCode string `json:"error_code"`
// 	Error     string `json:"error"`
// 	Message   string `json:"message"`
// }

// 响应设置gin.JSON
// func GetInfo(httpCode, errCode int, response *restful.Response) {
// func GetInfo(httpCode, errCode int) string {
func GetInfo(GetInfo interface{}) string {
	errResp := e.GetInfo(e.ERROR)
	if code, ok := GetInfo.(int); ok {
		errResp = e.GetInfo(code)
	}
	if msg, ok := GetInfo.(string); ok {
		errResp = e.ErrResponse{"INVALID", msg, "参数错误"}
	}
	data, err := json.Marshal(errResp)
	if err != nil {
		log.Printf("[json] json marshal error %s", err)
	}
	return string(data)
	// response.AddHeader("Content-Type", "application/json")
	// response.WriteHeaderAndEntity(httpCode, string(data))
}
