package main

import (
	"context"
	"fmt"
	"testing"
)

//
//func TestChan(t *testing.T) {
//	// 创建一个 chan，类型是 struct{}
//	ch := make(chan struct{})
//
//	go func() {
//		select {
//		// 这个 case 会在 chan 关闭或者收到值的时候执行，
//		// 在这里的情况是关闭了 chan。
//		case v, ok := <-ch:
//			if !ok {
//				// 输出 "chan ch is closed."
//				fmt.Println("chan ch is closed.")
//				fmt.Println(v) // {}
//			}
//		}
//	}()
//
//	// 关闭 chan，所有从 chan 读取的操作都会立即返回。
//	// 关闭 chan 之后，<-ch 返回的第一个值是 chan 对应类型的零值，第二个参数是 false。
//	// 如果不是关闭的 chan，第二个参数是 true，表示可以从 chan 获取到数据。
//	close(ch)
//
//	// {}，chan struct{} 关闭后，从中获取值的时候会立即返回一个空结构体实例
//	fmt.Println(<-ch)
//	fmt.Println(<-ch)
//
//	// 防止程序退出看不到效果
//	time.Sleep(time.Second)
//
//}
//
//func TestDeadLine(t *testing.T) {
//	ctx, _ := context.WithTimeout(context.Background(), time.Second)
//	// 输出 ctx 的 deadline，具体时间为 1 秒之后
//	spew.Dump(ctx.Deadline())
//
//	ctx1 := context.Background()
//	// ctx1 的超时时间是一个零值
//	spew.Dump(ctx1.Deadline())
//
//	ctx, cancel := context.WithCancel(ctx1)
//	cancel()
//}

func TestCancel(t *testing.T) {

	ctx := context.Background()
	ctx11 := context.WithValue(ctx, "key", "value")
	ctx1, _ := context.WithCancel(ctx11)
	ctx2, cancel2 := context.WithCancel(ctx1)
	ctx22 := context.WithValue(ctx2, "key2", "value2")
	ctx3 := context.WithValue(ctx22, "key1", "value1")
	cancel2()
	fmt.Println(ctx3.Value("key"))
	fmt.Println(ctx11.Value("key"))
	fmt.Println(ctx3.Value("key1"))
	<-ctx11.Done()
	//go func() {
	//	time.Sleep(time.Second)
	//
	//	fmt.Println("撤销cancel1了")
	//}()
	for {
		select {
		case <-ctx1.Done():
			fmt.Println("ctx1 取消了")
			//case <-ctx2.Done():
			//	fmt.Println("ctx2 取消了")
			//case <-ctx3.Done():
			//	fmt.Println("ctx3 取消了")
			//	fmt.Println(ctx3.Value("key"))
			//	fmt.Println(ctx3.Value("key2"))
			//	fmt.Println(ctx3.Value("key1"))
			return
		}
	}
}
