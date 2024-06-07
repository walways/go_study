package main

import (
	"context"
	"fmt"
	"testing"
	"time"

	"github.com/davecgh/go-spew/spew"
)

func TestChan(t *testing.T) {
	// 创建一个 chan，类型是 struct{}
	ch := make(chan struct{})

	go func() {
		select {
		// 这个 case 会在 chan 关闭或者收到值的时候执行，
		// 在这里的情况是关闭了 chan。
		case v, ok := <-ch:
			if !ok {
				// 输出 "chan ch is closed."
				fmt.Println("chan ch is closed.")
			}
			// 关闭 chan 之后得到的是 ch 的零值，也就是一个空结构体实例
			fmt.Println(v) // {}
		}
	}()

	// 关闭 chan，所有从 chan 读取的操作都会立即返回。
	// 关闭 chan 之后，<-ch 返回的第一个值是 chan 对应类型的零值，第二个参数是 false。
	// 如果不是关闭的 chan，第二个参数是 true，表示可以从 chan 获取到数据。
	close(ch)

	// 防止程序退出看不到效果
	time.Sleep(time.Second)

	// {}，chan struct{} 关闭后，从中获取值的时候会立即返回一个空结构体实例
	fmt.Println(<-ch)
	fmt.Println(<-ch)
}

func TestDeadLine(t *testing.T) {
	ctx, _ := context.WithTimeout(context.Background(), time.Second)
	// 输出 ctx 的 deadline，具体时间为 1 秒之后
	spew.Dump(ctx.Deadline())

	ctx1 := context.Background()
	// ctx1 的超时时间是一个零值
	spew.Dump(ctx1.Deadline())

	ctx, cancel := context.WithCancel(ctx1)
	cancel()
}
