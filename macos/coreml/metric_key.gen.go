// Code generated by DarwinKit. DO NOT EDIT.

package coreml

import (
	"unsafe"

	"github.com/progrium/darwinkit/objc"
)

// The class instance for the [MetricKey] class.
var MetricKeyClass = _MetricKeyClass{objc.GetClass("MLMetricKey")}

type _MetricKeyClass struct {
	objc.Class
}

// An interface definition for the [MetricKey] class.
type IMetricKey interface {
	IKey
}

// A key for the metrics dictionary in an update context. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreml/mlmetrickey?language=objc
type MetricKey struct {
	Key
}

func MetricKeyFrom(ptr unsafe.Pointer) MetricKey {
	return MetricKey{
		Key: KeyFrom(ptr),
	}
}

func (mc _MetricKeyClass) Alloc() MetricKey {
	rv := objc.Call[MetricKey](mc, objc.Sel("alloc"))
	return rv
}

func (mc _MetricKeyClass) New() MetricKey {
	rv := objc.Call[MetricKey](mc, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewMetricKey() MetricKey {
	return MetricKeyClass.New()
}

func (m_ MetricKey) Init() MetricKey {
	rv := objc.Call[MetricKey](m_, objc.Sel("init"))
	return rv
}

// The key you use to access the epoch index (an Int64 value). [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreml/mlmetrickey/3180090-epochindex?language=objc
func (mc _MetricKeyClass) EpochIndex() MetricKey {
	rv := objc.Call[MetricKey](mc, objc.Sel("epochIndex"))
	return rv
}

// The key you use to access the epoch index (an Int64 value). [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreml/mlmetrickey/3180090-epochindex?language=objc
func MetricKey_EpochIndex() MetricKey {
	return MetricKeyClass.EpochIndex()
}

// The key you use to access the current loss (a float value). [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreml/mlmetrickey/3180091-lossvalue?language=objc
func (mc _MetricKeyClass) LossValue() MetricKey {
	rv := objc.Call[MetricKey](mc, objc.Sel("lossValue"))
	return rv
}

// The key you use to access the current loss (a float value). [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreml/mlmetrickey/3180091-lossvalue?language=objc
func MetricKey_LossValue() MetricKey {
	return MetricKeyClass.LossValue()
}

// The key you use to access the mini-batch index (an Int64 value) within an epoch. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreml/mlmetrickey/3180092-minibatchindex?language=objc
func (mc _MetricKeyClass) MiniBatchIndex() MetricKey {
	rv := objc.Call[MetricKey](mc, objc.Sel("miniBatchIndex"))
	return rv
}

// The key you use to access the mini-batch index (an Int64 value) within an epoch. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreml/mlmetrickey/3180092-minibatchindex?language=objc
func MetricKey_MiniBatchIndex() MetricKey {
	return MetricKeyClass.MiniBatchIndex()
}
