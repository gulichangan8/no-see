package main

import (
	"encoding/json"
	"fmt"
	"os"
	"strings"
)

func main() {
	var message = map[string]string{}
	var ok int
	var name string
	var password string
	file, _ := os.ReadFile("./users.data")
	json.Unmarshal(file, &message)
	fmt.Println("您是否想注册一个账号（想则输入1,否则输入2）")
	fmt.Scan(&ok)
	if ok == 1 {
		fmt.Println("请输入您新建的用户名：")
		fmt.Scan(&name)
		checkUser(name, message)
		fmt.Println("用户名设置成功！请设置密码：")
		fmt.Scan(&password)
		checkPass(password)
		fmt.Println("您是否想退出程序（想则输入1，否则输入2）")
		fmt.Scan(&ok)
		if ok == 1 {
			os.OpenFile("./users.data", os.O_APPEND, 0777)
			message[name] = password
			marshal, _ := json.Marshal(message)
			os.WriteFile("./users.data", marshal, 0777)
			fmt.Println("程序已退出")
		}
	}
}
func checkUser(name string, m map[string]string) {
	for i := 0; i < len([]rune(name)); i++ {
		if strings.Count(name, string(name[i])) >= 2 {
			fmt.Println("您设置的用户名格式错误！请重新设置：")
			fmt.Println("提示：用户名不能重复")
			fmt.Scan(&name)
			checkUser(name, m)
			break
		}

	}
	if strings.Contains(name, "\\") == true {
		fmt.Println("您设置的用户名格式错误！请重新设置：")
		fmt.Println("提示：用户名不能含有\\字符")
		fmt.Scan(&name)
		checkUser(name, m)
	}
}

func checkPass(password string) {
	switch {
	case len([]byte(password)) >= 6:
		fmt.Println("密码设置成功！")
	case strings.Contains(password, "\\") == true:
		fmt.Println("密码中不能含有\\，请重新设置：")
		fmt.Scan(&password)
		checkPass(password)
	case len(password) < 6:
		fmt.Println("密码长度小于6，请重新设置：")
		fmt.Scan(&password)
		checkPass(password)
	}
}