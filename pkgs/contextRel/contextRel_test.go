package contextRel

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"testing"
	"time"
)

func dl2(ctx context.Context) {
	n := 1
	for {
		select {
		case <-ctx.Done():
			fmt.Println(ctx.Err())
			return
		default:
			fmt.Println("dl2 : ", n)
			n++
			time.Sleep(1 * time.Second)
		}
	}
}

func dl1(ctx context.Context) {
	n := 1
	for {
		select {
		case <-ctx.Done():
			fmt.Println(ctx.Err())
			return
		default:
			fmt.Println("dl1 : ", n)
			n++
			time.Sleep(2 * time.Second)
		}
	}
}
func Test01(t *testing.T) {
	// 设置deadline为当前时间之后的5秒那个时刻
	d := time.Now().Add(5 * time.Second)
	ctx, cancel := context.WithDeadline(context.Background(), d)
	defer cancel()
	go dl1(ctx)
	go dl2(ctx)
	for {
		select {
		case <-ctx.Done():
			fmt.Println("over", ctx.Err())
			return
		}
	}
}

func v3(ctx context.Context) {
	for {
		select {
		case <-ctx.Done():
			fmt.Println("v3 Done : ", ctx.Err())
			return
		default:
			fmt.Printf("in v3: %+v\n", ctx.Value("key"))
			time.Sleep(3 * time.Second)
		}
	}
}
func v2(ctx context.Context) {
	fmt.Printf("in v2: key=%+v\n", ctx.Value("key"))
	fmt.Printf("in v2: v1=%+v\n", ctx.Value("v1"))
	// 相同键,值覆盖
	ctx = context.WithValue(ctx, "key", "modify from v2")
	go v3(ctx)
}
func v1(ctx context.Context) {
	if v := ctx.Value("key"); v != nil {
		fmt.Printf("in v1: key = %+v\n", v)
	}
	ctx = context.WithValue(ctx, "v1", "value of v1 func")
	go v2(ctx)
	for {
		select {
		default:
			fmt.Println("print v1")
			time.Sleep(time.Second * 2)
		case <-ctx.Done():
			fmt.Println("v1 Done : ", ctx.Err())
			return
		}
	}
}
func Test02(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	// 向context中传递值
	ctx = context.WithValue(ctx, "key", "main")
	go v1(ctx)
	time.Sleep(10 * time.Second)
	cancel()
	time.Sleep(3 * time.Second)
}

func Test03(t *testing.T) {
	// 创建一个监听8000端口的服务器
	http.ListenAndServe(":8181", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()
		// 输出到STDOUT展示处理已经开始
		fmt.Fprint(os.Stdout, "processing request\n")
		// 通过select监听多个channel
		select {
		case <-time.After(2 * time.Second):
			// 如果两秒后接受到了一个消息后，意味请求已经处理完成
			// 我们写入"request processed"作为响应
			w.Write([]byte("request processed"))
		case <-ctx.Done():

			// 如果处理完成前取消了，在STDERR中记录请求被取消的消息
			fmt.Fprint(os.Stderr, "request cancelled\n")
		}
	}))
}
