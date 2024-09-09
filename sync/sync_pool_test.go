package sync

import (
	"sync"
	"testing"
)

type Person struct {
	name string
	age  int
	data [1024]byte // 增加对象大小
}

var result int // 全局变量，防止编译器优化

func BenchmarkPoolWithConcurrency(b *testing.B) {
	pool := sync.Pool{
		New: func() any {
			p := &Person{name: "chou", age: 18}
			return p
		},
	}

	b.ResetTimer()
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			person := pool.Get().(*Person)
			person.age++
			result += person.age                            // 确保结果被使用
			person.data[person.age%1024] = byte(person.age) //过于简单的话，可能导致sync.pool 性能不如之前的

			pool.Put(person)
		}
	})
}

func BenchmarkNoPoolWithConcurrency(b *testing.B) {
	b.ResetTimer()
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			person := &Person{name: "chou", age: 18}
			person.age++
			person.data[person.age%1024] = byte(person.age)
			result += person.age // 确保结果被使用
		}
	})
}
