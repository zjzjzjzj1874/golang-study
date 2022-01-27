package captcha

import (
	"fmt"
	"testing"
)

func TestNewDriverCaptcha(t *testing.T) {
	captcha := NewDriverCaptcha()

	id, pic, err := captcha.Generate()
	fmt.Println(id, err, pic)

	val, _ := captcha.generated.Load(id)
	fmt.Println(val)
	fmt.Println(captcha.Verify(id, val.(string)))

	captcha.generated.Range(func(key, value interface{}) bool {
		fmt.Println(key, value)
		return false
	})
}
