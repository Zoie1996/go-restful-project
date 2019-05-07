package util

import (
	"fmt"
	"strconv"
	"strings"
)

// StringToIntSlice 字符串转int切片
func StringToIntSlice(str string) []int {
	var int_arr []int
	if len(str) == 0 {
		return int_arr
	}
	str_arr := strings.Split(str, ",")
	for _, v := range str_arr {
		v1, _ := strconv.Atoi(v)
		int_arr = append(int_arr, v1)
	}
	return int_arr
}

// IntSliceToString int切片转字符串
func IntSliceToString(int_arr []int) string {
	if len(int_arr) == 0 {
		return ""
	}
	return strings.Replace(strings.Trim(fmt.Sprint(int_arr), "[]"), " ", ",", -1)
}

// StringSliceToString string切片转字符串
func StringSliceToString(int_arr []string) string {
	if len(int_arr) == 0 {
		return ""
	}
	return strings.Replace(strings.Trim(fmt.Sprint(int_arr), "[]"), " ", ",", -1)
}

// StringToStringSlice 字符串转String切片
func StringToStringSlice(str string) []string {
	var string_arr []string
	if len(str) == 0 {
		return string_arr
	}
	return strings.Split(str, ",")
}
