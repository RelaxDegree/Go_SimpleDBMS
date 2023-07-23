package controller

import (
	"mypro/dao"
	"strings"
)

func isDBMsg(opt string) bool {
	if opt == "select" || opt == "SELECT" || opt == "insert" || opt == "INSERT" || opt == "update" || opt == "UPDATE" || opt == "delete" || opt == "DELETE" {
		return true
	} else if opt == "create" || opt == "CREATE" || opt == "drop" || opt == "DROP" || opt == "alter" || opt == "ALTER" {
		return true
	}
	return false
}
func decodeDBMsg(msg string, ipaddr string) string {
	if _, isOk := userpool[ipaddr]; !isOk {
		return "not login"
	}
	lst := strings.Split(msg, " ")
	if lst[0] == "select" || lst[0] == "SELECT" {
		return dao.DBQL(msg)
	} else if lst[0] == "insert" || lst[0] == "INSERT" || lst[0] == "update" || lst[0] == "UPDATE" || lst[0] == "delete" || lst[0] == "DELETE" {
		if dao.DBML(msg) {
			return "manipulate success"
		} else {
			return "manipulate failed"
		}
	} else if lst[0] == "create" || lst[0] == "CREATE" || lst[0] == "drop" || lst[0] == "DROP" || lst[0] == "alter" || lst[0] == "ALTER" {
		if dao.DBDL(msg) {
			return "manipulate success"
		} else {
			return "manipulate failed"
		}
	}
	return "unknown command"
}
