package types

type Factory = func() interface{}
type PointerFactory = Factory
type ValueFactory = Factory
type SliceFactory = Factory
type PointerSliceFactory = Factory
