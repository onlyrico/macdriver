// Code generated by DarwinKit. DO NOT EDIT.

package mps

import (
	"unsafe"

	"github.com/progrium/darwinkit/objc"
)

// The class instance for the [CNNLocalContrastNormalizationGradientNode] class.
var CNNLocalContrastNormalizationGradientNodeClass = _CNNLocalContrastNormalizationGradientNodeClass{objc.GetClass("MPSCNNLocalContrastNormalizationGradientNode")}

type _CNNLocalContrastNormalizationGradientNodeClass struct {
	objc.Class
}

// An interface definition for the [CNNLocalContrastNormalizationGradientNode] class.
type ICNNLocalContrastNormalizationGradientNode interface {
	INNGradientFilterNode
	Pm() float32
	SetPm(value float32)
	Beta() float32
	SetBeta(value float32)
	KernelWidth() uint
	Alpha() float32
	SetAlpha(value float32)
	P0() float32
	SetP0(value float32)
	Delta() float32
	SetDelta(value float32)
	KernelHeight() uint
	Ps() float32
	SetPs(value float32)
}

// A representation of a gradient local-contrast normalization kernel. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnlocalcontrastnormalizationgradientnode?language=objc
type CNNLocalContrastNormalizationGradientNode struct {
	NNGradientFilterNode
}

func CNNLocalContrastNormalizationGradientNodeFrom(ptr unsafe.Pointer) CNNLocalContrastNormalizationGradientNode {
	return CNNLocalContrastNormalizationGradientNode{
		NNGradientFilterNode: NNGradientFilterNodeFrom(ptr),
	}
}

func (c_ CNNLocalContrastNormalizationGradientNode) InitWithSourceGradientSourceImageGradientStateKernelWidthKernelHeight(sourceGradient INNImageNode, sourceImage INNImageNode, gradientState INNGradientStateNode, kernelWidth uint, kernelHeight uint) CNNLocalContrastNormalizationGradientNode {
	rv := objc.Call[CNNLocalContrastNormalizationGradientNode](c_, objc.Sel("initWithSourceGradient:sourceImage:gradientState:kernelWidth:kernelHeight:"), sourceGradient, sourceImage, gradientState, kernelWidth, kernelHeight)
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnlocalcontrastnormalizationgradientnode/2948016-initwithsourcegradient?language=objc
func NewCNNLocalContrastNormalizationGradientNodeWithSourceGradientSourceImageGradientStateKernelWidthKernelHeight(sourceGradient INNImageNode, sourceImage INNImageNode, gradientState INNGradientStateNode, kernelWidth uint, kernelHeight uint) CNNLocalContrastNormalizationGradientNode {
	instance := CNNLocalContrastNormalizationGradientNodeClass.Alloc().InitWithSourceGradientSourceImageGradientStateKernelWidthKernelHeight(sourceGradient, sourceImage, gradientState, kernelWidth, kernelHeight)
	instance.Autorelease()
	return instance
}

func (cc _CNNLocalContrastNormalizationGradientNodeClass) NodeWithSourceGradientSourceImageGradientStateKernelWidthKernelHeight(sourceGradient INNImageNode, sourceImage INNImageNode, gradientState INNGradientStateNode, kernelWidth uint, kernelHeight uint) CNNLocalContrastNormalizationGradientNode {
	rv := objc.Call[CNNLocalContrastNormalizationGradientNode](cc, objc.Sel("nodeWithSourceGradient:sourceImage:gradientState:kernelWidth:kernelHeight:"), sourceGradient, sourceImage, gradientState, kernelWidth, kernelHeight)
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnlocalcontrastnormalizationgradientnode/2948047-nodewithsourcegradient?language=objc
func CNNLocalContrastNormalizationGradientNode_NodeWithSourceGradientSourceImageGradientStateKernelWidthKernelHeight(sourceGradient INNImageNode, sourceImage INNImageNode, gradientState INNGradientStateNode, kernelWidth uint, kernelHeight uint) CNNLocalContrastNormalizationGradientNode {
	return CNNLocalContrastNormalizationGradientNodeClass.NodeWithSourceGradientSourceImageGradientStateKernelWidthKernelHeight(sourceGradient, sourceImage, gradientState, kernelWidth, kernelHeight)
}

func (cc _CNNLocalContrastNormalizationGradientNodeClass) Alloc() CNNLocalContrastNormalizationGradientNode {
	rv := objc.Call[CNNLocalContrastNormalizationGradientNode](cc, objc.Sel("alloc"))
	return rv
}

func (cc _CNNLocalContrastNormalizationGradientNodeClass) New() CNNLocalContrastNormalizationGradientNode {
	rv := objc.Call[CNNLocalContrastNormalizationGradientNode](cc, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewCNNLocalContrastNormalizationGradientNode() CNNLocalContrastNormalizationGradientNode {
	return CNNLocalContrastNormalizationGradientNodeClass.New()
}

func (c_ CNNLocalContrastNormalizationGradientNode) Init() CNNLocalContrastNormalizationGradientNode {
	rv := objc.Call[CNNLocalContrastNormalizationGradientNode](c_, objc.Sel("init"))
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnlocalcontrastnormalizationgradientnode/2948053-pm?language=objc
func (c_ CNNLocalContrastNormalizationGradientNode) Pm() float32 {
	rv := objc.Call[float32](c_, objc.Sel("pm"))
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnlocalcontrastnormalizationgradientnode/2948053-pm?language=objc
func (c_ CNNLocalContrastNormalizationGradientNode) SetPm(value float32) {
	objc.Call[objc.Void](c_, objc.Sel("setPm:"), value)
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnlocalcontrastnormalizationgradientnode/2947977-beta?language=objc
func (c_ CNNLocalContrastNormalizationGradientNode) Beta() float32 {
	rv := objc.Call[float32](c_, objc.Sel("beta"))
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnlocalcontrastnormalizationgradientnode/2947977-beta?language=objc
func (c_ CNNLocalContrastNormalizationGradientNode) SetBeta(value float32) {
	objc.Call[objc.Void](c_, objc.Sel("setBeta:"), value)
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnlocalcontrastnormalizationgradientnode/2947965-kernelwidth?language=objc
func (c_ CNNLocalContrastNormalizationGradientNode) KernelWidth() uint {
	rv := objc.Call[uint](c_, objc.Sel("kernelWidth"))
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnlocalcontrastnormalizationgradientnode/2947973-alpha?language=objc
func (c_ CNNLocalContrastNormalizationGradientNode) Alpha() float32 {
	rv := objc.Call[float32](c_, objc.Sel("alpha"))
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnlocalcontrastnormalizationgradientnode/2947973-alpha?language=objc
func (c_ CNNLocalContrastNormalizationGradientNode) SetAlpha(value float32) {
	objc.Call[objc.Void](c_, objc.Sel("setAlpha:"), value)
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnlocalcontrastnormalizationgradientnode/2948017-p0?language=objc
func (c_ CNNLocalContrastNormalizationGradientNode) P0() float32 {
	rv := objc.Call[float32](c_, objc.Sel("p0"))
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnlocalcontrastnormalizationgradientnode/2948017-p0?language=objc
func (c_ CNNLocalContrastNormalizationGradientNode) SetP0(value float32) {
	objc.Call[objc.Void](c_, objc.Sel("setP0:"), value)
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnlocalcontrastnormalizationgradientnode/2948014-delta?language=objc
func (c_ CNNLocalContrastNormalizationGradientNode) Delta() float32 {
	rv := objc.Call[float32](c_, objc.Sel("delta"))
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnlocalcontrastnormalizationgradientnode/2948014-delta?language=objc
func (c_ CNNLocalContrastNormalizationGradientNode) SetDelta(value float32) {
	objc.Call[objc.Void](c_, objc.Sel("setDelta:"), value)
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnlocalcontrastnormalizationgradientnode/2948019-kernelheight?language=objc
func (c_ CNNLocalContrastNormalizationGradientNode) KernelHeight() uint {
	rv := objc.Call[uint](c_, objc.Sel("kernelHeight"))
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnlocalcontrastnormalizationgradientnode/2948008-ps?language=objc
func (c_ CNNLocalContrastNormalizationGradientNode) Ps() float32 {
	rv := objc.Call[float32](c_, objc.Sel("ps"))
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnlocalcontrastnormalizationgradientnode/2948008-ps?language=objc
func (c_ CNNLocalContrastNormalizationGradientNode) SetPs(value float32) {
	objc.Call[objc.Void](c_, objc.Sel("setPs:"), value)
}
