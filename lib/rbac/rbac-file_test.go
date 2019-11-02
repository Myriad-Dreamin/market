package rbac

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

func TestAddPolicy(t *testing.T) {

	e := Accquire("./test.csv")
	e.AddPolicy("data2_admin", "data2", "write")
	if err := e.SavePolicy(); err != nil {
		t.Error(err)
	}
	Release("./test.csv")
}

func TestDeletePolicy(t *testing.T) {

	e := Accquire("./test.csv")
	e.RemovePolicy("data2_admin", "data2", "write")
	if err := e.SavePolicy(); err != nil {
		t.Error(err)
	}
	Release("./test.csv")
}

func TestAccquireRelease(t *testing.T) {
	e := Accquire("./test.csv")
	fmt.Println(e.GetPolicy())
	time.Sleep(time.Second * 5)
	e2 := Accquire("./test.csv")
	fmt.Printf("%p %p\n", e, e2)
	Release("./test.csv")
	time.Sleep(time.Second * 8)
	Release("./test.csv")
	time.Sleep(time.Second * 1)
}

func BenchmarkAccquire(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = Accquire("./test.csv")
	}
}

func BenchmarkAccquireReleaseMultiThread(b *testing.B) {
	for i := 0; i < b.N; i++ {
		var wg sync.WaitGroup
		for j := 0; j < 1000; j++ {
			wg.Add(1)
			go func() {
				e := Accquire("./test.csv")

				e.AddPolicy("data2_admin", "data2", "write")
				if err := e.SavePolicy(); err != nil {
					b.Error(err)
				}
				e.RemovePolicy("data2_admin", "data2", "write")
				if err := e.SavePolicy(); err != nil {
					b.Error(err)
				}
				Release("./test.csv")
				wg.Done()
			}()
		}
		wg.Wait()
	}
}

func BenchmarkAccquireRelease(b *testing.B) {
	for i := 0; i < b.N; i++ {
		e := Accquire("./test.csv")

		e.AddPolicy("data2_admin", "data2", "write")
		if err := e.SavePolicy(); err != nil {
			b.Error(err)
		}
		e.RemovePolicy("data2_admin", "data2", "write")
		if err := e.SavePolicy(); err != nil {
			b.Error(err)
		}
		Release("./test.csv")
	}
}
