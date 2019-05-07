package app

// import (
// 	"github.com/gin-gonic/gin"

// 	"go-restful-project/pkg/e"
// )

// type ErrResponse struct {
// 	ErrorCode string `json:"error_code"`
// 	Error     string `json:"error"`
// 	Message   string `json:"message"`
// }

// // 响应设置gin.JSON
// func (g *Gin) ErrResponse(httpCode, errCode int) {
// 	msg := e.GetInfo(errCode)
// 	g.C.JSON(httpCode, ErrResponse{
// 		ErrorCode: msg[0],
// 		Error:     msg[1],
// 		Message:   msg[2],
// 	})
// 	return
// }

// type Response struct {
// 	Success bool        `json:"success"`
// 	Error   *Error      `json:"error,omitempty"`
// 	Data    interface{} `json:"data,omitempty"`
// }

// func WriteSuccess(resp *restful.Response) {
// 	NewResponse(true).WriteStatus(200, resp)
// }

// func WriteResponse(data interface{}, resp *restful.Response) {
// 	WriteResponseStatus(200, data, resp)
// }

// func WriteResponseStatus(status int, data interface{}, resp *restful.Response) {
// 	success := NewResponse(true)
// 	success.Data = data
// 	success.WriteStatus(status, resp)
// }

// func NewResponse(success bool) *Response {
// 	return &Response{Success: success}
// }

// func NewErrorResponse(err error) *Response {
// 	res := &Response{Success: false, Error: &Error{}}
// 	res.SetError(err)
// 	return res
// }

// func (r *Response) SetError(err error) {
// 	if err != nil {
// 		if r.Error == nil {
// 			r.Error = &Error{}
// 		}
// 		r.Error.Name = err.Error()
// 	}
// }

// func (r *Response) WriteStatus(status int, resp *restful.Response) {
// 	if r.Error != nil && r.Error.Code == 0 {
// 		r.Error.Code = status
// 	}
// 	resp.WriteHeaderAndEntity(status, r)
// }
