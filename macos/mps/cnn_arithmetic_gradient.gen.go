// Code generated by DarwinKit. DO NOT EDIT.

package mps

import (
	"unsafe"

	"github.com/progrium/darwinkit/macos/metal"
	"github.com/progrium/darwinkit/objc"
)

// The class instance for the [CNNArithmeticGradient] class.
var CNNArithmeticGradientClass = _CNNArithmeticGradientClass{objc.GetClass("MPSCNNArithmeticGradient")}

type _CNNArithmeticGradientClass struct {
	objc.Class
}

// An interface definition for the [CNNArithmeticGradient] class.
type ICNNArithmeticGradient interface {
	ICNNGradientKernel
	Bias() float32
	SetBias(value float32)
	SecondaryStrideInFeatureChannels() uint
	SetSecondaryStrideInFeatureChannels(value uint)
	MinimumValue() float32
	SetMinimumValue(value float32)
	IsSecondarySourceFilter() bool
	SecondaryScale() float32
	SetSecondaryScale(value float32)
	MaximumValue() float32
	SetMaximumValue(value float32)
	PrimaryScale() float32
	SetPrimaryScale(value float32)
}

// The base class for gradient arithmetic operators. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnarithmeticgradient?language=objc
type CNNArithmeticGradient struct {
	CNNGradientKernel
}

func CNNArithmeticGradientFrom(ptr unsafe.Pointer) CNNArithmeticGradient {
	return CNNArithmeticGradient{
		CNNGradientKernel: CNNGradientKernelFrom(ptr),
	}
}

func (cc _CNNArithmeticGradientClass) Alloc() CNNArithmeticGradient {
	rv := objc.Call[CNNArithmeticGradient](cc, objc.Sel("alloc"))
	return rv
}

func (cc _CNNArithmeticGradientClass) New() CNNArithmeticGradient {
	rv := objc.Call[CNNArithmeticGradient](cc, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewCNNArithmeticGradient() CNNArithmeticGradient {
	return CNNArithmeticGradientClass.New()
}

func (c_ CNNArithmeticGradient) Init() CNNArithmeticGradient {
	rv := objc.Call[CNNArithmeticGradient](c_, objc.Sel("init"))
	return rv
}

func (c_ CNNArithmeticGradient) InitWithDevice(device metal.PDevice) CNNArithmeticGradient {
	po0 := objc.WrapAsProtocol("MTLDevice", device)
	rv := objc.Call[CNNArithmeticGradient](c_, objc.Sel("initWithDevice:"), po0)
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnngradientkernel/2942657-initwithdevice?language=objc
func NewCNNArithmeticGradientWithDevice(device metal.PDevice) CNNArithmeticGradient {
	instance := CNNArithmeticGradientClass.Alloc().InitWithDevice(device)
	instance.Autorelease()
	return instance
}

func (c_ CNNArithmeticGradient) CopyWithZoneDevice(zone unsafe.Pointer, device metal.PDevice) CNNArithmeticGradient {
	po1 := objc.WrapAsProtocol("MTLDevice", device)
	rv := objc.Call[CNNArithmeticGradient](c_, objc.Sel("copyWithZone:device:"), zone, po1)
	return rv
}

// Makes a copy of this kernel object for a new device. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpskernel/1618912-copywithzone?language=objc
func CNNArithmeticGradient_CopyWithZoneDevice(zone unsafe.Pointer, device metal.PDevice) CNNArithmeticGradient {
	instance := CNNArithmeticGradientClass.Alloc().CopyWithZoneDevice(zone, device)
	instance.Autorelease()
	return instance
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnarithmeticgradient/2951855-bias?language=objc
func (c_ CNNArithmeticGradient) Bias() float32 {
	rv := objc.Call[float32](c_, objc.Sel("bias"))
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnarithmeticgradient/2951855-bias?language=objc
func (c_ CNNArithmeticGradient) SetBias(value float32) {
	objc.Call[objc.Void](c_, objc.Sel("setBias:"), value)
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnarithmeticgradient/2951853-secondarystrideinfeaturechannels?language=objc
func (c_ CNNArithmeticGradient) SecondaryStrideInFeatureChannels() uint {
	rv := objc.Call[uint](c_, objc.Sel("secondaryStrideInFeatureChannels"))
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnarithmeticgradient/2951853-secondarystrideinfeaturechannels?language=objc
func (c_ CNNArithmeticGradient) SetSecondaryStrideInFeatureChannels(value uint) {
	objc.Call[objc.Void](c_, objc.Sel("setSecondaryStrideInFeatureChannels:"), value)
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnarithmeticgradient/2951858-minimumvalue?language=objc
func (c_ CNNArithmeticGradient) MinimumValue() float32 {
	rv := objc.Call[float32](c_, objc.Sel("minimumValue"))
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnarithmeticgradient/2951858-minimumvalue?language=objc
func (c_ CNNArithmeticGradient) SetMinimumValue(value float32) {
	objc.Call[objc.Void](c_, objc.Sel("setMinimumValue:"), value)
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnarithmeticgradient/2951852-issecondarysourcefilter?language=objc
func (c_ CNNArithmeticGradient) IsSecondarySourceFilter() bool {
	rv := objc.Call[bool](c_, objc.Sel("isSecondarySourceFilter"))
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnarithmeticgradient/2951856-secondaryscale?language=objc
func (c_ CNNArithmeticGradient) SecondaryScale() float32 {
	rv := objc.Call[float32](c_, objc.Sel("secondaryScale"))
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnarithmeticgradient/2951856-secondaryscale?language=objc
func (c_ CNNArithmeticGradient) SetSecondaryScale(value float32) {
	objc.Call[objc.Void](c_, objc.Sel("setSecondaryScale:"), value)
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnarithmeticgradient/2951854-maximumvalue?language=objc
func (c_ CNNArithmeticGradient) MaximumValue() float32 {
	rv := objc.Call[float32](c_, objc.Sel("maximumValue"))
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnarithmeticgradient/2951854-maximumvalue?language=objc
func (c_ CNNArithmeticGradient) SetMaximumValue(value float32) {
	objc.Call[objc.Void](c_, objc.Sel("setMaximumValue:"), value)
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnarithmeticgradient/2951861-primaryscale?language=objc
func (c_ CNNArithmeticGradient) PrimaryScale() float32 {
	rv := objc.Call[float32](c_, objc.Sel("primaryScale"))
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnarithmeticgradient/2951861-primaryscale?language=objc
func (c_ CNNArithmeticGradient) SetPrimaryScale(value float32) {
	objc.Call[objc.Void](c_, objc.Sel("setPrimaryScale:"), value)
}
