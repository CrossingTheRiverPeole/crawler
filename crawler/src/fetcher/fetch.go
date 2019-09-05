package fetcher

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

var rateLimiter = time.Tick(100 * time.Millisecond)

func Fetch(url string) ([]byte, error) {
	//<-rateLimiter
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	// 关闭resp
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("wrong status code: %d", resp.StatusCode)
	}

	// 获取内容
	all, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}
	return all, nil
}
