package parser

import (
	"crawler/src/engine"
	"crawler/src/model"
	"fmt"
	"reflect"
	"regexp"
)

var regexs = map[string]*regexp.Regexp{
	"Age":        regexp.MustCompile(`<td><span class="label">年龄：</span>(.+)</td>`),
	"Sex":        regexp.MustCompile(`<td><span class="label">性别：</span><span field="">(.+)</span></td>`),
	"Height":     regexp.MustCompile(`<td><span class="label">身高：</span>(.+)</td>`),
	"Marriage":   regexp.MustCompile(`<td><span class="label">婚况：</span>([^<]+)</td>`),
	"Edu":        regexp.MustCompile(`<td><span class="label">学历：</span>(.+)</td>`),
	"Job":        regexp.MustCompile(`<td><span class="label">职业：.*</span>([^<]+)</td>`),
	"JobAddress": regexp.MustCompile(`<td><span class="label">工作地：</span>([^<]+)</td>`),
	"HasChild":   regexp.MustCompile(`<td><span class="label">有无孩子：</span>(.+)</td>`),
	"Income":     regexp.MustCompile(`<td><span class="label">月收入：</span>([^<]+)</td>`),
}

func ParseProfile(content []byte, name string) engine.ParseResult {
	profile := model.Profile{Name: name}
	v := reflect.ValueOf(&profile).Elem()

	for k, r := range regexs {
		s := extractString(content, r)
		if s != "" {
			a := v.FieldByName(k)
			if a.IsValid() {
				a.Set(reflect.ValueOf(s))
			}
		} else {
			//log.Warn("未能解析的属性：%s", k)
		}
	}

	fmt.Printf("用户信息打印， %v",profile)

	rs := engine.ParseResult{
		Items: []interface{}{profile},
	}
	return rs

}
func extractString(c []byte, r *regexp.Regexp) string {

	match := r.FindSubmatch(c)
	if match != nil && len(match) >= 2 {
		return string(match[1])
	} else {
		return ""
	}

}
