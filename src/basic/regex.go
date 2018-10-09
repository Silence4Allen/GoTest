package main

import (
	"regexp"
	"fmt"
)

const text = `My email is ccmouse@gmail.com
email1 is abc@def.org
email2 is hello@qq.com
email3 is silence4allen@gmail.com
email4 is weibo@abd.com.cn
`

func main() {
	re := regexp.MustCompile(`([a-zA-Z0-9]+)@([a-zA-Z0-9]+)(\.[a-zA-Z0-9].+)`)
	match := re.FindAllStringSubmatch(text, -1)
	for _, m := range match {
		fmt.Println(m)
	}
}
