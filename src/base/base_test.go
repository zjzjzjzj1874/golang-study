package base

import (
	"fmt"
	"testing"
	"time"
	"unicode/utf8"
)

func Test_StringOrRune(t *testing.T) {
	t.Run("#count", func(t *testing.T) {
		str := "hello,我是中国人"
		count := utf8.RuneCountInString(str)
		fmt.Println("RuneCountInString长度:", count) // 11

		c1 := utf8.RuneCount([]byte(str))
		fmt.Println("RuneCount长度:", c1) // 11

		fmt.Println("len长度:", len(str)) // 21:一个中文占3个字节
	})

}

func Test_timeTicker(t *testing.T) {
	t.Run("#time.Ticker test case", func(t *testing.T) {
		tc := time.NewTicker(time.Second) // second ticker => to execute every second task
		for v := range tc.C {
			fmt.Println(v)
		}
		time.Sleep(time.Minute)
	})

	t.Run("#time.Ticker test case", func(t *testing.T) {
		tc := time.NewTicker(time.Second) // second ticker => to execute every second task
		for range tc.C {
			fmt.Println(time.Now().Format(time.RFC3339))
		}
		time.Sleep(time.Minute)
	})
}

// 闭包函数测试
func Test_Closure(t *testing.T) {
	// 计算1-10之和
	t.Run("#closure闭包计算", func(t *testing.T) {
		fn := closure()
		sum := 0
		for i := 1; i <= 10; i++ {
			sum = fn(i) // sum也可以定义为局部变量,不过这样定义也好点,可以复用sum的地址,否则栈上内存创建又销毁
		}
		fmt.Printf("sum:%d\n", sum)
	})
	t.Run("#正常计算", func(t *testing.T) {
		// 使用闭包计算1-10之和
		sum := 0
		for i := 1; i <= 10; i++ {
			sum += i
		}
		fmt.Printf("sum:%d\n", sum)
	})

}
