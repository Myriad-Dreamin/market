package traits

import (
	"fmt"
	"github.com/Myriad-Dreamin/market/model/db-layer"
	"testing"
)

func BenchmarkBaseTraits_objectFactory(b *testing.B) {
	var info = NewBaseTraits(dblayer.User{})
	b.ResetTimer()
	for i := 0; i < b.N; i ++ {
		_ = info.ObjectFactory()
	}
}

func cccccc() interface{} { return new(dblayer.User) }

func BenchmarkBaseTraits_objectFactoryPure(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i ++ {
		_ = cccccc()
	}
}

func BenchmarkBaseTraits_objectFactoryPureN(b *testing.B) {
	var New = cccccc
	b.ResetTimer()

	for i := 0; i < b.N; i ++ {
		_ = New()
	}
}

func TestTraits_objectFactory(t *testing.T) {
	var info = NewBaseTraits(dblayer.User{})
	if a, ok := info.ObjectFactory().(*dblayer.User); ok {
		fmt.Printf("%T %v\n", a, a)
	} else {
		t.Errorf("bad creation %T", a)
	}
}

func TestTraits_sliceFactory(t *testing.T) {
	var info = NewBaseTraits(dblayer.User{})
	if a, ok := info.SliceFactory().([]dblayer.User); ok {
		fmt.Printf("%T %v\n", a, a)
	} else {
		t.Errorf("bad creation %T", a)
	}
}