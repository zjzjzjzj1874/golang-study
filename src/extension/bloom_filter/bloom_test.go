package bloomfilter

import (
	"fmt"
	"testing"
)

func TestNewBloomFilter(t *testing.T) {
	filter := NewBloomFilter()
	fmt.Println("seed = ",filter.Funcs[1].Seed)
	str1 := "hello,bloom filter!"
	filter.Add(str1)
	str2 := "你好,布隆过滤器"
	filter.Add(str2)
	str3 := "Greate wall"
	filter.Add(str3)
	str4 := "Greate walll"
	filter.Add(str4)
	str5 := "Greate willl"
	filter.Add(str5)

	fmt.Println("count:",filter.Set.Count())
	fmt.Println("str1 is contains in:",filter.Contains(str1))
	fmt.Println("str2 is contains in:",filter.Contains(str2))
	fmt.Println("str3 is contains in:",filter.Contains(str3))
	fmt.Println("str5 is contains in:",filter.Contains(str4))
	fmt.Println("str5 is contains in:",filter.Contains(str5))
	fmt.Println(filter.Contains("blockchain technology"))
	fmt.Println(filter.Contains("Greate willll"))
	fmt.Println(filter.Contains("Greate wallll"))
}
