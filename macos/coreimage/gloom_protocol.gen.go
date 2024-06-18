// Code generated by DarwinKit. DO NOT EDIT.

package coreimage

import (
	"github.com/progrium/darwinkit/objc"
)

// The properties you use to configure a gloom filter. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cigloom?language=objc
type PGloom interface {
	// optional
	SetIntensity(value float32)
	HasSetIntensity() bool

	// optional
	Intensity() float32
	HasIntensity() bool

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
var _ PGloom = (*GloomObject)(nil)

// A concrete type for the [PGloom] protocol.
type GloomObject struct {
	objc.Object
}

func (g_ GloomObject) HasSetIntensity() bool {
	return g_.RespondsToSelector(objc.Sel("setIntensity:"))
}

// The intensity of the effect. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cigloom/3228478-intensity?language=objc
func (g_ GloomObject) SetIntensity(value float32) {
	objc.Call[objc.Void](g_, objc.Sel("setIntensity:"), value)
}

func (g_ GloomObject) HasIntensity() bool {
	return g_.RespondsToSelector(objc.Sel("intensity"))
}

// The intensity of the effect. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cigloom/3228478-intensity?language=objc
func (g_ GloomObject) Intensity() float32 {
	rv := objc.Call[float32](g_, objc.Sel("intensity"))
	return rv
}

func (g_ GloomObject) HasSetInputImage() bool {
	return g_.RespondsToSelector(objc.Sel("setInputImage:"))
}

// The image to use as an input image. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cigloom/3228477-inputimage?language=objc
func (g_ GloomObject) SetInputImage(value Image) {
	objc.Call[objc.Void](g_, objc.Sel("setInputImage:"), value)
}

func (g_ GloomObject) HasInputImage() bool {
	return g_.RespondsToSelector(objc.Sel("inputImage"))
}

// The image to use as an input image. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cigloom/3228477-inputimage?language=objc
func (g_ GloomObject) InputImage() Image {
	rv := objc.Call[Image](g_, objc.Sel("inputImage"))
	return rv
}

func (g_ GloomObject) HasSetRadius() bool {
	return g_.RespondsToSelector(objc.Sel("setRadius:"))
}

// The radius, in pixels, of the effect. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cigloom/3228479-radius?language=objc
func (g_ GloomObject) SetRadius(value float32) {
	objc.Call[objc.Void](g_, objc.Sel("setRadius:"), value)
}

func (g_ GloomObject) HasRadius() bool {
	return g_.RespondsToSelector(objc.Sel("radius"))
}

// The radius, in pixels, of the effect. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cigloom/3228479-radius?language=objc
func (g_ GloomObject) Radius() float32 {
	rv := objc.Call[float32](g_, objc.Sel("radius"))
	return rv
}
