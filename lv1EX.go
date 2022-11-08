package main

import (
	"fmt"
	"strings"
)

func main() {
	var name int
	var names []int
	var Names string
	skillNames := []string{"激光波", "龙卷风", "动感光波"}
	var judge bool
	var model string
	models := []string{"尝尝我的厉害吧！", "给你点颜色瞧瞧！", "我厉害吧！", "别眨眼哦！"}
	var skillname string
	fmt.Println("您可以选择以下模板：")
	fmt.Println("1.尝尝我的厉害吧！")
	fmt.Println("2.给你点颜色瞧瞧！")
	fmt.Println("3.我厉害吧！")
	fmt.Println("4.别眨眼哦！")
	fmt.Println("您是否想要自定义模板(想自定义则输入true,不想则输入false)")
	fmt.Println("提示：您最多可以自定义三个模板")
	fmt.Scan(&judge)
	if judge == true {
		fmt.Println("请输入您定义的模板：")
		fmt.Scan(&model)
		ok := strings.Contains(model, "sb")
		if ok == true {
			fmt.Println("模板中有敏感词汇，定义不成功")
		} else {
			fmt.Println("定义成功！")
			models = append(models, model)
		}
		for i := 0; i < 2; i++ {
			fmt.Println("您是否想继续定义新的模板(想自定义则输入true,不想则输入false)")
			fmt.Scan(&judge)
			if judge == false {
				break
			} else {
				fmt.Println("请输入您定义的模板：")
				fmt.Scan(&model)
				ok := strings.Contains(model, "sb")
				if ok == true {
					fmt.Println("模板中有敏感词汇，定义不成功")
				} else {
					fmt.Println("定义成功！")
					models = append(models, model)
				}
			}
		}
	}
	fmt.Println("您现在拥有以下模板：")
	for i := 0; i < len(models); i++ {
		fmt.Println(i+1, ".", models[i])
	}
	fmt.Println("您拥有以下技能：")
	fmt.Println("1.激光波")
	fmt.Println("请选择该技能您想使用的模板：")
	fmt.Scan(&name)
	names = append(names, name)
	fmt.Println("2.龙卷风")
	fmt.Println("请选择该技能您想使用的模板：")
	fmt.Scan(&name)
	names = append(names, name)
	fmt.Println("3.动感光波")
	fmt.Println("请选择该技能您想使用的模板：")
	fmt.Scan(&name)
	names = append(names, name)
	fmt.Println("您是否想自定义技能(想自定义则输入true,不想则输入false)")
	fmt.Println("提示：您最多可以自定义三个技能")
	fmt.Scan(&judge)
	if judge == true {
		fmt.Println("请输入您定义的技能名：")
		fmt.Scan(&Names)
		ok := strings.Contains(Names, "sb")
		if ok == true {
			fmt.Println("技能名中有敏感词汇，定义不成功")
		} else {
			fmt.Println("定义成功！")
			skillNames = append(skillNames, Names)
		}
		fmt.Println("请选择该技能您想使用的模板：")
		fmt.Scan(&name)
		names = append(names, name)
		for i := 0; i < 2; i++ {
			fmt.Println("您是否想继续定义新的技能(想自定义则输入true,不想则输入false)")
			fmt.Scan(&judge)
			if judge == false {
				break
			}
			fmt.Println("请输入您定义的技能名：")
			fmt.Scan(&Names)
			ok := strings.Contains(Names, "sb")
			if ok == true {
				fmt.Println("技能名中有敏感词汇，定义不成功")
			} else {
				fmt.Println("定义成功！")
				skillNames = append(skillNames, Names)
			}
			fmt.Println("请选择该技能您想使用的模板：")
			fmt.Scan(&name)
			names = append(names, name)
		}
	}
	fmt.Println("请输入您想释放的技能名：")
	fmt.Scan(&skillname)
	Release(skillname, func(skillName string) {
		for i := 0; i < len(names); i++ {
			if skillname == skillNames[i] {
				fmt.Println(models[names[i]-1], skillname)
			}
		}
	})
}
func Release(skill string, releaseSkillFunc func(string)) {
	releaseSkillFunc(skill)
}
