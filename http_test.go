package gotang

import (
	"fmt"
	"log"
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
}
