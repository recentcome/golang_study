package main

import "fmt"

//如果某个用户名存在，则修改其密码“999999”，如果不存在就增加这个用户名的昵称和密码
//编写一个函数modifyUser(users map[string]map[string]string, name string)
func modifyUser(users map[string]map[string]string, name string) {
	if users[name] != nil {
		users[name]["pwd"] = "999999"
	} else {
		users[name] = make(map[string]string)
		users[name]["nickname"] = name + "-nick"
		users[name]["pwd"] = "888888"

	}
}

func main() {
	users := make(map[string]map[string]string)
	users["simth"] = make(map[string]string)
	users["simth"]["pwd"] = "777777"
	users["simth"]["nickname"] = "小花猫"

	modifyUser(users, "tom")
	modifyUser(users, "mary")
	modifyUser(users, "simth")
	fmt.Println(users)
}
