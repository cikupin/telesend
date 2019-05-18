package telesend

import "fmt"

// ToString convert any data to string
func toString(v interface{}) string {
	str, ok := v.(string)
	if !ok {
		str = fmt.Sprintf("%#v", v)
	}
	return str
}
