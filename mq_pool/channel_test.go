package rabbitmq

import (
	"fmt"
	"testing"
	"time"
)

func TestChanClose(t *testing.T) {

	ch := make(chan int)
	//close(ch)
	for i := 0; i < 10; i++ {
		go func(i int) {
			select {
			case <-ch:
				fmt.Println(i)
			}
		}(i)
	}
	time.Sleep(1 * time.Second)
}

func TestChannel1(t *testing.T) {
	ch := make(chan int)
	go func() {
		time.Sleep(2 * time.Second)
		ch <- 1
		fmt.Println("close5")

	}()
	//<-ch

	select {
	case <-ch:
		fmt.Printf("select \n")
		break
	}
}

func TestChannel2(t *testing.T) {
	ch := make(chan int, 2)
	ch <- 1
	ch <- 1

	aa, ok := <-ch
	fmt.Printf("aa:%+d,ok:%+v \n", aa, ok)
	close(ch)
	aa1, ok := <-ch
	fmt.Printf("aa1:%+d,ok:%+v \n", aa1, ok)
	aa2, ok := <-ch
	fmt.Printf("aa2:%+d,ok:%+v \n", aa2, ok)
}

func TestChannel3(t *testing.T) {
	ch := make(chan int)
	//go func() {
	//	ch <- 1
	//}()
	//go func() {
	//	time.Sleep(2 * time.Second)
	//	close(ch)
	//	fmt.Println("close")
	//}()
	go func() {
		time.Sleep(3 * time.Second)
		ch = nil
		fmt.Println("nil")
	}()
	for {
		select {
		case <-ch:
			fmt.Println(111)
			break
		default:
			fmt.Println("default")
			time.Sleep(2 * time.Second)
		}
	}
}
