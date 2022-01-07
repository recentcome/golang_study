package main

import "fmt"

func main() {
	//第一种方式
	var a map[string]string
	a = make(map[string]string)
	a["sdg"] = "sdgg"
	//第二种方式
	hero := make(map[string]string)
	hero["nihao"] = "ok"
	fmt.Println(hero)
	//第三种方式
	var heros = map[string]string{
		"hier": "sdg",
		"sdsg": "sdg1",
	}
	fmt.Println(heros)

	//map中包含map
	students := map[string]map[string]string{
		"stu01": {"name": "mary", "age": "10"},
		"stu02": {"name": "cc", "age": "11"},
	}
	fmt.Println(students)
	fmt.Println(students["stu01"])

	heros1 := map[string]string{
		"nihao": "",
	}

	//查找key
	//第一个返回值为值，第二个返回值是是否有值
	value, ok := heros1["nihao"]
	if ok {
		fmt.Println(value)
	} else {
		fmt.Println("没哟uKey")
	}

	//遍历map
	for k, v := range students {
		fmt.Println(k)
		for k2, v2 := range v {
			fmt.Printf("\tk=%v,v=%v\n", k2, v2)
		}
		fmt.Println()
	}
}
