package template

import (
	"fmt"
	"html/template"
	"strings"
	"time"

	"github.com/itang/gotang"
)

var ExtTemplateFuncs = map[string]interface{}{
	"str": func(v interface{}) string {
		return fmt.Sprintf("%v", v)
	},
	"inc": func(i1, i2 int) int {
		return i1 + i2
	},
	"add": func(i1, i2 int) int {
		return i1 + i2
	},
	"sub": func(i1, i2 int) int {
		return i1 - i2
	},
	"emptyOr": func(value interface{}, other interface{}) interface{} {
		switch value.(type) {
		case string:
			{
				s, _ := value.(string)
				if s == "" {
					return other
				}
			}
		}
		if value == nil {
			return other
		}
		return value
	},
	"active": func(s1, s2 string) string {
		if strings.HasPrefix(s2, s1) {
			return "active"
		}
		return ""
	},
	"current": func(s1, s2 string) string {
		if strings.HasPrefix(s2, s1) {
			return "current"
		}
		return ""
	},
	"startsWith": strings.HasPrefix,
	"zeroAsEmpty": func(v interface{}) interface{} {
		switch v.(type) {
		case int, int32, int64:
			if v == 0 {
				return ""
			}
		case time.Time:
			if v.(time.Time).IsZero() {
				return ""
			}
		}
		return v
	},
	"boolStr": func(v bool) string {
		if v {
			return "true"
		}
		return "false"
	},
	"mod": func(i int, j int) int {
		return i % j
	},
	"rawjs": func(s string) template.JS {
		return template.JS(s)
	},
	"newline": func(index int, maxline int) bool {
		i := index + 1
		return i%maxline == 1 && i != 1
	},
	"truncStr": gotang.TruncStr,
}
