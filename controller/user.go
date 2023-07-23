package controller

import (
	"fmt"
	"mypro/dao"
	"strings"
)

var userpool map[string]string = make(map[string]string) //ipaddr->username

func DecodeUserMsg(msg string, ipaddr string) string {
	lst := strings.Split(msg, " ")
	ret := ""
	if lst[0] == "login" {
		fmt.Println(lst[1], lst[2])
		if userpool[ipaddr] != "" {
			ret = "already login"
		}
		if dao.UserLogin(lst[1], lst[2]) {
			userpool[ipaddr] = lst[1]
			ret = "login success"
		} else {
			ret = "login failed"
		}
	} else if lst[0] == "register" {
		if dao.UserRegister(lst[1], lst[2]) {
			userpool[ipaddr] = lst[1]
			ret = "register success "
		} else {
			ret = "register failed"
		}
	} else if lst[0] == "changepswd" {
		if dao.UserChangePswd(userpool[ipaddr], lst[1], lst[2]) {
			ret = "change password success"
		} else {
			ret = "change password failed"
		}
	} else if lst[0] == "quit" {
		if userpool[ipaddr] == "" {
			ret = "not login"
		} else {
			delete(userpool, ipaddr)
			ret = "quit success"
		}
	} else if lst[0] == "help" {
		return "login [username] [password]\nregister [username] [password]\nchangepswd [username] [password]\nexit\n"
	} else {
		ret = "unknown command"
	}
	dao.Record(userpool[ipaddr], msg, ipaddr, ret)
	return ret
}
