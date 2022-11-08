package main

import "fmt"

func main() {
	var name int
	var skillNames string
	fmt.Println("1.尝尝我的厉害吧！")
	fmt.Println("2.给你点颜色瞧瞧！")
	fmt.Println("3.我厉害吧！")
	fmt.Println("4.别眨眼哦！")
	fmt.Println("请选择您想使用的模板：")
	fmt.Scan(&name, &skillNames)
	ReleaseSkill(skillNames, func(skillName string) {
		if name == 1 {
			fmt.Println("尝尝我的厉害吧！", skillName)
		}
		if name == 2 {
			fmt.Println("给你点颜色瞧瞧！", skillName)
		}
		if name == 3 {
			fmt.Println("我厉害吧！", skillName)
		}
		if name == 4 {
			fmt.Println("别眨眼哦！", skillName)
		}
	})
}

func ReleaseSkill(skillNames string, releaseSkillFunc func(string)) {
	releaseSkillFunc(skillNames)
}
