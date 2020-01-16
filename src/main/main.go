package main

import (
	Myconf "MyConf"
	"fmt"
)

func main() {
	conf := Myconf.MyConfig{}
	filename := "src/main/config.conf"
	conf.Init(filename)
	fmt.Println(conf.GetSectionList())
	conf.SetParamInteger("SectionA", "QQQ", 2323)
	fmt.Println(conf.GetSection("SectionA"))
}
