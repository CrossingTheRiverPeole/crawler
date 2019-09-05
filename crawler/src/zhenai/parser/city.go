package parser

import (
	"crawler/src/engine"
	"fmt"
	"regexp"
)

var (
	profileRe = regexp.MustCompile(`<a href="(http://album.zhenai.com/u/\w+)"[^>]*>([^<]+)</a>`)

	cityUrlRe = regexp.MustCompile(`href="(http://www.zhenai.com/zhenghun/[^"]+)`)
)

func ParseCity(content []byte) engine.ParseResult {
	match := profileRe.FindAllSubmatch(content, -1)

	result := engine.ParseResult{}
	for _, m := range match {
		name := string(m[2])
		result.Items = append(result.Items, name)
		result.Requests = append(result.Requests, engine.Request{
			Url: string(m[1]),
			ParserFunc: func(c []byte) engine.ParseResult {
				return ParseProfile(c, name)
			},
		})
	}
	fmt.Printf("用户Url %v", result.Requests)
	return result
}
