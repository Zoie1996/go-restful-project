package e

var MsgFlags = map[int]string{
	SUCCESS:                         "ok",
	ERROR:                           "fail",
	INVALID_PARAMS:                  "请求参数错误",
	ERROR_EXIST_TAG:                 "已存在该标签名称",
	ERROR_EXIST_TAG_FAIL:            "获取已存在标签失败",
	ERROR_NOT_EXIST_TAG:             "该标签不存在",
	ERROR_GET_TAGS_FAIL:             "获取所有标签失败",
	ERROR_COUNT_TAG_FAIL:            "统计标签失败",
	ERROR_ADD_TAG_FAIL:              "新增标签失败",
	ERROR_EDIT_TAG_FAIL:             "修改标签失败",
	ERROR_DELETE_TAG_FAIL:           "删除标签失败",
	ERROR_EXPORT_TAG_FAIL:           "导出标签失败",
	ERROR_IMPORT_TAG_FAIL:           "导入标签失败",
	ERROR_NOT_EXIST_ARTICLE:         "该文章不存在",
	ERROR_ADD_ARTICLE_FAIL:          "新增文章失败",
	ERROR_DELETE_ARTICLE_FAIL:       "删除文章失败",
	ERROR_CHECK_EXIST_ARTICLE_FAIL:  "检查文章是否存在失败",
	ERROR_EDIT_ARTICLE_FAIL:         "修改文章失败",
	ERROR_COUNT_ARTICLE_FAIL:        "统计文章失败",
	ERROR_GET_ARTICLES_FAIL:         "获取多个文章失败",
	ERROR_GET_ARTICLE_FAIL:          "获取单个文章失败",
	ERROR_GEN_ARTICLE_POSTER_FAIL:   "生成文章海报失败",
	ERROR_AUTH_CHECK_TOKEN_FAIL:     "Token鉴权失败",
	ERROR_AUTH_CHECK_TOKEN_TIMEOUT:  "Token已超时",
	ERROR_AUTH_TOKEN:                "Token生成失败",
	ERROR_AUTH:                      "Token错误",
	ERROR_UPLOAD_SAVE_IMAGE_FAIL:    "保存图片失败",
	ERROR_UPLOAD_CHECK_IMAGE_FAIL:   "检查图片失败",
	ERROR_UPLOAD_CHECK_IMAGE_FORMAT: "校验图片错误，图片格式或大小有问题",
}

// GetMsg get error information based on Code
func GetMsg(code int) string {
	msg, ok := MsgFlags[code]
	if ok {
		return msg
	}

	return MsgFlags[ERROR]
}

type ErrResponse struct {
	ErrorCode string      `json:"error_code"`
	Error     interface{} `json:"error"`
	Message   string      `json:"message"`
}

var MsgInfos = map[int]ErrResponse{
	ERROR:                    ErrResponse{"FAIL", "fail", "fail"},
	NOT_FOUND:                ErrResponse{"NOT_FOUND", "对象不存在", "对象不存在"},
	ERROR_EXIST_USER:         ErrResponse{"EXIST_USER", "用户已存在", "用户已存在"},
	ERROR_COUNT_USER_FAIL:    ErrResponse{"GET_COUNT_USER_FAIL", "获取用户总数失败", "获取用户总数失败"},
	ERROR_GET_USERS_FAIL:     ErrResponse{"GET_USERS_FAIL", "获取用户列表失败", "获取用户列表失败"},
	ERROR_ADD_USER_FAIL:      ErrResponse{"ADD_USER_FAIL", "添加用户失败", "添加用户失败"},
	INVALID_PARAMS:           ErrResponse{"INVALID_PARAMS", "请求参数错误", "请求参数错误"},
	Json_Marshal_Err:         ErrResponse{"Json_Marshal_Err", "JSON序列化失败", "JSON序列化失败"},
	ERROR_USER_NAME_INVALID:  ErrResponse{"ERROR_USER_NAME_INVALID", "仅允许3-36个字符的用户名", "仅允许3-36个字符的用户名"},
	ERROR_GET_ROLE_NAME_FAIL: ErrResponse{"ERROR_GET_ROLE_NAME", "获取角色名失败", "获取角色名失败"},
	ERROR_ROLE_NOT_FOUND:     ErrResponse{"ERROR_ROLE_NOT_FOUND", "获取角色失败", "获取角色失败"},
}

func GetInfo(code interface{}) ErrResponse {
	// msg, ok := MsgInfos[code]
	// if ok {
	// 	return msg
	// }
	// return MsgInfos[ERROR]

	if code, ok := code.(int); ok {
		msg, ok := MsgInfos[code]
		if ok {
			return msg
		}
	}
	if msg, ok := code.(error); ok {
		return ErrResponse{"INVALID", msg.Error(), "参数错误"}
	}
	if msg, ok := code.(string); ok {
		return ErrResponse{"INVALID", msg, "参数错误"}
	}
	return MsgInfos[ERROR]
}
