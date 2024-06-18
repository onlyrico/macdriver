// Code generated by DarwinKit. DO NOT EDIT.

package mps

import (
	"unsafe"

	"github.com/progrium/darwinkit/macos/metal"
	"github.com/progrium/darwinkit/objc"
)

// The class instance for the [ImageAreaMax] class.
var ImageAreaMaxClass = _ImageAreaMaxClass{objc.GetClass("MPSImageAreaMax")}

type _ImageAreaMaxClass struct {
	objc.Class
}

// An interface definition for the [ImageAreaMax] class.
type IImageAreaMax interface {
	IUnaryImageKernel
	KernelWidth() uint
	KernelHeight() uint
}

// A filter that finds the maximum pixel value in a rectangular region centered around each pixel in the source image. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsimageareamax?language=objc
type ImageAreaMax struct {
	UnaryImageKernel
}

func ImageAreaMaxFrom(ptr unsafe.Pointer) ImageAreaMax {
	return ImageAreaMax{
		UnaryImageKernel: UnaryImageKernelFrom(ptr),
	}
}

func (i_ ImageAreaMax) InitWithDeviceKernelWidthKernelHeight(device metal.PDevice, kernelWidth uint, kernelHeight uint) ImageAreaMax {
	po0 := objc.WrapAsProtocol("MTLDevice", device)
	rv := objc.Call[ImageAreaMax](i_, objc.Sel("initWithDevice:kernelWidth:kernelHeight:"), po0, kernelWidth, kernelHeight)
	return rv
}

// Initializes the kernel with a specified width and height. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsimageareamax/1618281-initwithdevice?language=objc
func NewImageAreaMaxWithDeviceKernelWidthKernelHeight(device metal.PDevice, kernelWidth uint, kernelHeight uint) ImageAreaMax {
	instance := ImageAreaMaxClass.Alloc().InitWithDeviceKernelWidthKernelHeight(device, kernelWidth, kernelHeight)
	instance.Autorelease()
	return instance
}

func (ic _ImageAreaMaxClass) Alloc() ImageAreaMax {
	rv := objc.Call[ImageAreaMax](ic, objc.Sel("alloc"))
	return rv
}

func (ic _ImageAreaMaxClass) New() ImageAreaMax {
	rv := objc.Call[ImageAreaMax](ic, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewImageAreaMax() ImageAreaMax {
	return ImageAreaMaxClass.New()
}

func (i_ ImageAreaMax) Init() ImageAreaMax {
	rv := objc.Call[ImageAreaMax](i_, objc.Sel("init"))
	return rv
}

func (i_ ImageAreaMax) InitWithDevice(device metal.PDevice) ImageAreaMax {
	po0 := objc.WrapAsProtocol("MTLDevice", device)
	rv := objc.Call[ImageAreaMax](i_, objc.Sel("initWithDevice:"), po0)
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsunaryimagekernel/2866332-initwithdevice?language=objc
func NewImageAreaMaxWithDevice(device metal.PDevice) ImageAreaMax {
	instance := ImageAreaMaxClass.Alloc().InitWithDevice(device)
	instance.Autorelease()
	return instance
}

func (i_ ImageAreaMax) CopyWithZoneDevice(zone unsafe.Pointer, device metal.PDevice) ImageAreaMax {
	po1 := objc.WrapAsProtocol("MTLDevice", device)
	rv := objc.Call[ImageAreaMax](i_, objc.Sel("copyWithZone:device:"), zone, po1)
	return rv
}

// Makes a copy of this kernel object for a new device. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpskernel/1618912-copywithzone?language=objc
func ImageAreaMax_CopyWithZoneDevice(zone unsafe.Pointer, device metal.PDevice) ImageAreaMax {
	instance := ImageAreaMaxClass.Alloc().CopyWithZoneDevice(zone, device)
	instance.Autorelease()
	return instance
}

// The width of the filter window. Must be an odd number. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsimageareamax/1618282-kernelwidth?language=objc
func (i_ ImageAreaMax) KernelWidth() uint {
	rv := objc.Call[uint](i_, objc.Sel("kernelWidth"))
	return rv
}

// The height of the filter window. Must be an odd number. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsimageareamax/1618277-kernelheight?language=objc
func (i_ ImageAreaMax) KernelHeight() uint {
	rv := objc.Call[uint](i_, objc.Sel("kernelHeight"))
	return rv
}
