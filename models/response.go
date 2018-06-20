package models

import (
	"strconv"
)

type ResponseMessage struct {
	Code    string
	Message string
	Data    interface{}
}

var (
	//注册
	A_REGISTER_SUCCESS = ResponseMessage{strconv.Itoa(0), "注册成功！", nil}
	A_REGISTER_EXIST   = ResponseMessage{strconv.Itoa(1), "用户名已存在！", nil}

	//登录
	A_LOGIN_SUCCESS   = ResponseMessage{strconv.Itoa(0), "登陆成功！", nil}
	A_LOGIN_NOT_EXIST = ResponseMessage{strconv.Itoa(1), "用户名不存在！", nil}
	A_LOGIN_NOT_MATCH = ResponseMessage{strconv.Itoa(2), "用户名和密码不匹配", nil}

	//修改密码
	A_MODITY_PASSWD_SUCCESS   = ResponseMessage{strconv.Itoa(0), "密码修改成功！", nil}
	A_MODIFY_PASSWD_NOT_MATCH = ResponseMessage{strconv.Itoa(1), "新密码输入不一致！", nil}
	A_MODIFY_PASSWD_WRONG     = ResponseMessage{strconv.Itoa(2), "原密码不正确！", nil}

	//忘记密码
	A_GORGET_PASSWD_SUCCESS   = ResponseMessage{strconv.Itoa(0), "密码修改成功", nil}
	A_FORGET_PASSWD_NOT_MATCH = ResponseMessage{strconv.Itoa(1), "密码输入不一致!!", nil}

	//组合
	A_COMBINATION_SUCCESS                = ResponseMessage{strconv.Itoa(0), "组合创建成功!!!", nil}
	A_COMBINATION_CREATE_FAILD           = ResponseMessage{strconv.Itoa(1), "组合创建失败!!!", nil}
	A_COMBINATION_MODIFY_FAILD           = ResponseMessage{strconv.Itoa(2), "修改组合失败!!!", nil}
	A_COMBINATION_CREATE_ADDCOIN_FAILD   = ResponseMessage{strconv.Itoa(3), "组合添加币种失败!!!", nil}
	A_COMBINATION_CREATE_ADDCOIN_SUCCESS = ResponseMessage{strconv.Itoa(4), "组合添加币种成功!!!", nil}
	A_COMBINATION_DELETE_FAILD           = ResponseMessage{strconv.Itoa(5), "组合删除失败!!!", nil}
	A_COMBINATION_DELETE_SUCCESS         = ResponseMessage{strconv.Itoa(4), "组合删除成功!!!", nil}

	//币
	A_COIN_DELETE_SUCCESS = ResponseMessage{strconv.Itoa(0), "币种删除成功!!!", nil}
	A_COIN_DELETE_FAILD   = ResponseMessage{strconv.Itoa(1), "币种删除失败!!!", nil}
	A_COIN_FIND_FAILD     = ResponseMessage{strconv.Itoa(2), "查找失败!!!", nil}
)
