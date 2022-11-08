package main

import (
	"crypto/hmac"
	"crypto/md5"
	"encoding/hex"
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
	var p, u string
	file, _ := os.ReadFile("./users.data")
	json.Unmarshal(file, &message)
	fmt.Println("您是否想注册一个账号（想则输入1,否则输入2）")
	fmt.Scan(&ok)
	if ok == 1 {
		fmt.Println("请输入您新建的用户名：")
		fmt.Scan(&name)
		CheckUser(name, message)
		fmt.Println("用户名设置成功！请设置密码：")
		fmt.Scan(&password)
		CheckPass(password)
		fmt.Println("请登录您的账户")
		fmt.Println("请输入用户名:")
		fmt.Scan(&u)
		enter(u, name)
		fmt.Scan(&p)
		Enter(p, password)
		fmt.Println("您是否想修改密码（1.想 2.不想）")
		fmt.Scan(&ok)
		if ok == 1 {
			fmt.Println("请输入您的新密码:")
			fmt.Scan(&password)
			CheckPass(password)
		}
	}
	fmt.Println("您是否想退出程序（想则输入1，否则输入2）")
	fmt.Scan(&ok)
	if ok == 1 {
		os.OpenFile("./users.data", os.O_APPEND, 0777)
		n := Hmac("a", name)
		P := Hmac("a", password)
		message[n] = P
		marshal, _ := json.Marshal(message)
		os.WriteFile("./users.data", marshal, 0777)
		fmt.Println("程序已退出")
	}
}

func CheckUser(name string, m map[string]string) {
	for i := 0; i < len([]rune(name)); i++ {
		if strings.Count(name, string(name[i])) >= 2 {
			fmt.Println("您设置的用户名格式错误！请重新设置：")
			fmt.Println("提示：用户名不能重复")
			fmt.Scan(&name)
			CheckUser(name, m)
			break
		}
	}
	if strings.Contains(name, "\\") == true {
		fmt.Println("您设置的用户名格式错误！请重新设置：")
		fmt.Println("提示：用户名不能含有\\字符")
		fmt.Scan(&name)
		CheckUser(name, m)
	}
	for k := range m {
		if Hmac("a", name) == k {
			fmt.Println("用户名已存在！请重新设置：")
			fmt.Scan(&name)
			CheckUser(name, m)
		}
		break
	}

}

func CheckPass(password string) {
	switch {
	case len([]byte(password)) >= 6:
		fmt.Println("密码设置成功！")
	case strings.Contains(password, "\\") == true:
		fmt.Println("密码中不能含有\\，请重新设置：")
		fmt.Scan(&password)
		CheckPass(password)
	case len(password) < 6:
		fmt.Println("密码长度小于6，请重新设置：")
		fmt.Scan(&password)
		CheckPass(password)
	}
}

func enter(u, name string) {
	if u == name {
		fmt.Println("用户名存在，请输入对应的密码:")
	} else {
		fmt.Println("用户名不存在！请重新输入:")
		fmt.Scan(&u)
		enter(u, name)
	}
}

func Enter(p, password string) {
	if p == password {
		fmt.Println("成功登录账户！")
	} else {
		fmt.Println("密码错误！请重新输入:")
		fmt.Scan(&p)
		enter(p, password)
	}
}

func Hmac(key, date string) string {
	hash := hmac.New(md5.New, []byte(key))
	hash.Write([]byte(date))
	return hex.EncodeToString(hash.Sum(nil))
}
