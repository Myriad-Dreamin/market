//go:generate stringer -type=GoodsType
package types

type GoodsType uint16

const (
	GoodsTypeUnknown GoodsType = iota
	GoodsTypeElectronic
	GoodsTypeDaily
	GoodsTypeBook
	GoodsTypeLength
)

var GoodsTypesMap map[string]GoodsType

func init() {
	GoodsTypesMap = make(map[string]GoodsType)
	for i := GoodsType(0); i < GoodsTypeLength; i++ {
		GoodsTypesMap[i.String()] = i
	}
}
