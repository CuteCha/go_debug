package cacheRel

import (
	"fmt"
	"github.com/patrickmn/go-cache"
	"testing"
	"time"
)

func Test01(t *testing.T) {
	// 创建一个cache对象，默认ttl 5分钟，每10分钟对过期数据进行一次清理
	//c := cache.New(5*time.Minute, 10*time.Minute)
	//c := cache.New(500*time.Millisecond, 500*time.Millisecond)
	c := cache.New(500000000, 500000000)

	// Set一个KV，key是"foo"，value是"bar"
	// TTL是默认值（上面创建对象的入参，也可以设置不同的值）5分钟
	c.Set("foo", "bar", cache.DefaultExpiration)

	// Set了一个没有TTL的KV，只有调用delete接口指定key时才会删除
	c.Set("baz", 42, cache.NoExpiration)

	// 从cache中获取key对应的value
	if foo, found := c.Get("foo"); found {
		fmt.Println(foo)
	} else {
		fmt.Println("get foo fail~")
	}

	if baz, found := c.Get("baz"); found {
		fmt.Println(baz)
	} else {
		fmt.Println("get baz fail~")
	}

	time.Sleep(500 * time.Millisecond)
	if foo, found := c.Get("foo"); found {
		fmt.Println(foo)
	} else {
		fmt.Println("get foo fail~")
	}
}
