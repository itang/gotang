package gotang

import (
	"fmt"
	"log"
	"strings"
	"testing"

	"github.com/bmizerany/assert"
)

func TestGetJSON(t *testing.T) {
	obj := make(map[string]interface{})
	err := GetJSON("http://ip.taobao.com/service/getIpInfo.php?ip=8.8.8.8", &obj)

	if err != nil {
		log.Println(err)
		return
	}

	code, _ := obj["code"]
	log.Println("code", code)
	assert.Equal(t, "0", fmt.Sprintf("%v", code), "code should be 0")

	err1 := GetJSON("http://somebadhost", &obj)
	log.Println(err1)
	assert.NotEqual(t, nil, err1, "bad host raise error")
}

func TestGetString(t *testing.T) {
	content, err := GetString("http://ip.taobao.com/service/getIpInfo.php?ip=8.8.8.8")
	if err != nil {
		log.Println(err)
		return
	}
	println(content)
	assert.Equal(t, true, strings.Contains(content, `"code":0`), "code should be 0")
}
