package doc_gen

import (
	"fmt"
	"github.com/Myriad-Dreamin/market/lib/mock"
)

type Results = mock.ResultsI
type Records = mock.RecordsI

func FromResults(res []Results) {
	for i := range res {
		fmt.Println(res[i])
	}
}


