// Package captcha 图形验证码工具
package captcha

import (
	captcha "github.com/mojocn/base64Captcha"
	"math/rand"
	"sync"
	"time"
)

func init() {
	//init rand seed
	rand.Seed(time.Now().UnixNano())
}

type DriverCaptcha struct {
	driver    captcha.Driver
	generated *sync.Map
}

// NewDriverCaptcha maybe later can init DriverDigit outside && manage every captcha's life cycle
func NewDriverCaptcha() *DriverCaptcha {
	return &DriverCaptcha{
		driver:    captcha.NewDriverDigit(40, 100, 4, 0.01, 1),
		generated: &sync.Map{},
	}
}

//Generate generates a random id, base64 image string or an error if any
func (c *DriverCaptcha) Generate() (id, b64s string, err error) {
	id, content, answer := c.driver.GenerateIdQuestionAnswer()
	item, err := c.driver.DrawCaptcha(content)
	if err != nil {
		return "", "", err
	}
	c.generated.Store(id, answer)
	b64s = item.EncodeB64string()
	return
}

//Verify by a given id key and remove the captcha value in store,
//return boolean value.
//if you has multiple captcha instances which share a same store.
//You may want to call `store.Verify` method instead.
func (c *DriverCaptcha) Verify(id, answer string) (match bool) {
	if val, ok := c.generated.Load(id); ok && val == answer {
		match = true
	}
	return
}

// Clear todo store in cache,and manage it's life cycle
func (c *DriverCaptcha) Clear(id string) {
	c.generated.Delete(id)
}
