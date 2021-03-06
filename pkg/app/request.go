package app

import (
	"github.com/astaxie/beego/validation"

	"go-restful-project/pkg/logging"
)

// 错误日志
func MarkErrors(errors []*validation.Error) {
	for _, err := range errors {
		logging.Info(err.Key, err.Message)
	}

	return
}
