// Code generated by DarwinKit. DO NOT EDIT.

package coreimage

import (
	"github.com/progrium/darwinkit/macos/coregraphics"
	"github.com/progrium/darwinkit/objc"
)

// The properties you use to configure a crystalize filter. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cicrystallize?language=objc
type PCrystallize interface {
	// optional
	SetCenter(value coregraphics.Point)
	HasSetCenter() bool

	// optional
	Center() coregraphics.Point
	HasCenter() bool

	// optional
	SetInputImage(value Image)
	HasSetInputImage() bool

	// optional
	InputImage() Image
	HasInputImage() bool

	// optional
	SetRadius(value float32)
	HasSetRadius() bool

	// optional
	Radius() float32
	HasRadius() bool
}

// ensure impl type implements protocol interface
var _ PCrystallize = (*CrystallizeObject)(nil)

// A concrete type for the [PCrystallize] protocol.
type CrystallizeObject struct {
	objc.Object
}

func (c_ CrystallizeObject) HasSetCenter() bool {
	return c_.RespondsToSelector(objc.Sel("setCenter:"))
}

// The center of the effect as x and y coordinates. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cicrystallize/3228200-center?language=objc
func (c_ CrystallizeObject) SetCenter(value coregraphics.Point) {
	objc.Call[objc.Void](c_, objc.Sel("setCenter:"), value)
}

func (c_ CrystallizeObject) HasCenter() bool {
	return c_.RespondsToSelector(objc.Sel("center"))
}

// The center of the effect as x and y coordinates. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cicrystallize/3228200-center?language=objc
func (c_ CrystallizeObject) Center() coregraphics.Point {
	rv := objc.Call[coregraphics.Point](c_, objc.Sel("center"))
	return rv
}

func (c_ CrystallizeObject) HasSetInputImage() bool {
	return c_.RespondsToSelector(objc.Sel("setInputImage:"))
}

// The image to use as an input image. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cicrystallize/3228201-inputimage?language=objc
func (c_ CrystallizeObject) SetInputImage(value Image) {
	objc.Call[objc.Void](c_, objc.Sel("setInputImage:"), value)
}

func (c_ CrystallizeObject) HasInputImage() bool {
	return c_.RespondsToSelector(objc.Sel("inputImage"))
}

// The image to use as an input image. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cicrystallize/3228201-inputimage?language=objc
func (c_ CrystallizeObject) InputImage() Image {
	rv := objc.Call[Image](c_, objc.Sel("inputImage"))
	return rv
}

func (c_ CrystallizeObject) HasSetRadius() bool {
	return c_.RespondsToSelector(objc.Sel("setRadius:"))
}

// The radius, in pixels, of the effect. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cicrystallize/3228202-radius?language=objc
func (c_ CrystallizeObject) SetRadius(value float32) {
	objc.Call[objc.Void](c_, objc.Sel("setRadius:"), value)
}

func (c_ CrystallizeObject) HasRadius() bool {
	return c_.RespondsToSelector(objc.Sel("radius"))
}

// The radius, in pixels, of the effect. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cicrystallize/3228202-radius?language=objc
func (c_ CrystallizeObject) Radius() float32 {
	rv := objc.Call[float32](c_, objc.Sel("radius"))
	return rv
}
