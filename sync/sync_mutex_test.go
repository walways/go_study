package sync

import (
	"fmt"
	"sync"
	"sync/atomic"
	"testing"
)

// 不要重复锁定互斥锁；
// 不要忘记解锁互斥锁，必要时使用defer语句；
// 不要对尚未锁定或者已解锁的互斥锁解锁；
// 不要在多个函数之间直接传递互斥锁。
func TestLock(t *testing.T) {

	mu := sync.RWMutex{}
	mu.Lock()
	defer mu.Unlock()
	lockInput(&mu)
}

func lockInput(mu *sync.RWMutex) {
	mu.Unlock()

}

func TestAtomic(t *testing.T) {
	var test1 int32 = 4
	addUint32 := atomic.AddInt32(&test1, int32(-3))
	fmt.Println(addUint32)
	for {
		if atomic.CompareAndSwapInt32(&test1, 1, 6) {
			fmt.Println("success")
			break
		}
	}
}

func TestAtomic1(t *testing.T) {

}

func TestOne(t *testing.T) {

	//one := sync.Once{}
	////testing.Coverage()
	//for {
	//	one.Do(func() {
	//		fmt.Println("hello")
	//	})
	//}
	//
	for {
		sync.OnceFunc(func() {
			fmt.Println("hello")
		})
	}

}
