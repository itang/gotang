package gotang

import (
	"fmt"
	"log"
	"strings"
	"testing"

	"github.com/bmizerany/assert"
)

func TestHttpGetAsJSON(t *testing.T) {
	obj := make(map[string]interface{})
	err := HttpGetAsJSON("http://ip.taobao.com/service/getIpInfo.php?ip=8.8.8.8", &obj)

	if err != nil {
		log.Println(err)
		return
	}

	code, _ := obj["code"]
	log.Println("code", code)
	assert.Equal(t, "0", fmt.Sprintf("%v", code), "code should be 0")

	err1 := HttpGetAsJSON("http://somebadhost", &obj)
	log.Println(err1)
	assert.NotEqual(t, nil, err1, "bad host raise error")
}

func TestHttpGetAsString(t *testing.T) {
	content, err := HttpGetAsString("http://ip.taobao.com/service/getIpInfo.php?ip=8.8.8.8")
	if err != nil {
		log.Println(err)
		return
	}
	println(content)
	assert.Equal(t, true, strings.Contains(content, `"code":0`), "code should be 0")
}
