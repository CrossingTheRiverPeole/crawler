package main

import (
	"fmt"
	"regexp"
)

const text = ` My email is ccmouse@gmail.com
email is ddd@dndn.com.cn
email is ddc@dsd.com
`

func main() {
	compile := regexp.MustCompile(`([a-zA-Z0-9]+)@([a-zA-Z0-9\.]+)\.([a-zA-Z0-9]+)`)
	match := compile.FindAllStringSubmatch(text, -1)

	for _, m := range match  {
		fmt.Println(m)
	}
}
