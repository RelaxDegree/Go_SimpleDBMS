package controller

import (
	"strings"
)

func decodeDBMsg(msg string, ipaddr string) string {
	if _, isOk := userpool[ipaddr]; !isOk {
		return "not login"
	}
	lst := strings.Split(msg, " ")
	if lst[0] == "select" || lst[0] == "SELECT" {
		return msg
	}
	return "unknown command"
}
