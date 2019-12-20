//go:generate stringer -type=GoodsStatus
package types

type GoodsStatus uint8

const (
	GoodsStatusUnknown GoodsStatus = iota
	GoodsStatusUnFinished
	GoodsStatusPending
	GoodsStatusFinished
	GoodsStatusCancelled
)
