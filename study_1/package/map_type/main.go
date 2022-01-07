package main

import "fmt"

func main() {
	//var countryCapitalMap map[string]string
	countryCapitalMap := make(map[string]string)
	countryCapitalMap["france"] = "巴黎"
	countryCapitalMap["Italy"] = "罗马"
	countryCapitalMap["Japan"] = "东京"
	countryCapitalMap["India"] = "新德里"

	for country := range countryCapitalMap {
		fmt.Println(country, "首都是", countryCapitalMap[country])
	}

	capital, ok := countryCapitalMap["American"]
	if ok {
		fmt.Println("American的首都是", capital)
	} else {
		fmt.Println("American的首都不存在")
	}
	m := make(map[string]int)
	m["Answer"] = 42
	fmt.Println("the value:", m["Answer"])

	m["Answer"] = 48
	fmt.Println("the value:", m["Answer"])
	delete(m, "Answer")
	fmt.Println("the value:", m["Answer"])
	v, ok := m["Answer"]
	fmt.Println("the value:", v, "Present?", ok)
}
