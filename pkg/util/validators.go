package util

import (
	"fmt"
	"reflect"

	"github.com/smokezl/govalidators"
)

func uniqueMethod(params map[string]interface{}, val reflect.Value, args ...string) (bool, error) {
	fmt.Println("=====", "userMethod")
	return true, nil
}

func ValidatorMap() map[string]interface{} {

	msg := map[string]interface{}{
		"string": &govalidators.StringValidator{
			Range: govalidators.Range{
				RangeEMsg: map[string]string{
					"between": "[name] 长度必须在 [min] 和 [max] 之间",
				},
			},
		},
		"integer": &govalidators.IntegerValidator{
			Range: govalidators.Range{
				RangeEMsg: map[string]string{
					"between": "[name] 的值必须在 [min] 和 [max] 之间",
				},
			},
		},
		"in": &govalidators.InValidator{
			EMsg: "[name] 的值必须为 [args] 中的一个",
		},
		"email": &govalidators.EmailValidator{
			EMsg: "[name] 不是一个有效的email地址",
		},
		"url": &govalidators.UrlValidator{
			EMsg: "[name] 不是一个有效的url地址",
		},
		"datetime": &govalidators.DateTimeValidator{
			EMsg: "[name] 不是一个有效的日期",
		},
		// "unique": &govalidators.UniqueValidator{
		// 	EMsg: "[name] 不是唯一的",
		// },
		"unique": uniqueMethod,
	}
	return msg
}
