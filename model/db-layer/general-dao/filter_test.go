package general_dao

import (
	"encoding/json"
	"fmt"
	traits "github.com/Myriad-Dreamin/go-model-traits/example-traits"
	"testing"
)

func TestFilter(t *testing.T) {
	var x = GoodsFilter{
		Filter: traits.Filter{
			Order:    "123123",
			Page:     124124,
			PageSize: 444,
			BeforeID: 044,
		},
		BySeller:   0,
		ByBuyer:    0,
		HasType:    0,
		HasStatus:  0,
		WithName:   "",
		MinPriceL:  nil,
		MinPriceR:  nil,
		MaxPriceL:  nil,
		MaxPriceR:  nil,
		FixedTag:   nil,
		EndBefore:  nil,
		BeginAfter: nil,
	}
	b, err := json.Marshal(x)
	fmt.Println(string(b), err)
	x = GoodsFilter{}
	err = json.Unmarshal(b, &x)
	fmt.Println(x)
}
