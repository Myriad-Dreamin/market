// Code generated by "stringer -type=GoodsStatus"; DO NOT EDIT.

package types

import "strconv"

func _() {
	// An "invalid array index" compiler error signifies that the constant values have changed.
	// Re-run the stringer command to generate them again.
	var x [1]struct{}
	_ = x[GoodsStatusUnknown-0]
	_ = x[GoodsStatusUnFinished-1]
	_ = x[GoodsStatusPending-2]
	_ = x[GoodsStatusFinished-3]
	_ = x[GoodsStatusCancelled-4]
}

const _GoodsStatus_name = "GoodsStatusUnknownGoodsStatusUnFinishedGoodsStatusPendingGoodsStatusFinishedGoodsStatusCancelled"

var _GoodsStatus_index = [...]uint8{0, 18, 39, 57, 76, 96}

func (i GoodsStatus) String() string {
	if i >= GoodsStatus(len(_GoodsStatus_index)-1) {
		return "GoodsStatus(" + strconv.FormatInt(int64(i), 10) + ")"
	}
	return _GoodsStatus_name[_GoodsStatus_index[i]:_GoodsStatus_index[i+1]]
}
