// Code generated by DarwinKit. DO NOT EDIT.

package coreimage

import (
	"github.com/progrium/darwinkit/macos/coregraphics"
	"github.com/progrium/darwinkit/objc"
)

// The properties you use to configure a checkerboard generator filter. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cicheckerboardgenerator?language=objc
type PCheckerboardGenerator interface {
	// optional
	SetCenter(value coregraphics.Point)
	HasSetCenter() bool

	// optional
	Center() coregraphics.Point
	HasCenter() bool

	// optional
	SetColor0(value Color)
	HasSetColor0() bool

	// optional
	Color0() Color
	HasColor0() bool

	// optional
	SetWidth(value float32)
	HasSetWidth() bool

	// optional
	Width() float32
	HasWidth() bool

	// optional
	SetColor1(value Color)
	HasSetColor1() bool

	// optional
	Color1() Color
	HasColor1() bool

	// optional
	SetSharpness(value float32)
	HasSetSharpness() bool

	// optional
	Sharpness() float32
	HasSharpness() bool
}

// ensure impl type implements protocol interface
var _ PCheckerboardGenerator = (*CheckerboardGeneratorObject)(nil)

// A concrete type for the [PCheckerboardGenerator] protocol.
type CheckerboardGeneratorObject struct {
	objc.Object
}

func (c_ CheckerboardGeneratorObject) HasSetCenter() bool {
	return c_.RespondsToSelector(objc.Sel("setCenter:"))
}

// The center of the effect as x and y coordinates. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cicheckerboardgenerator/3228105-center?language=objc
func (c_ CheckerboardGeneratorObject) SetCenter(value coregraphics.Point) {
	objc.Call[objc.Void](c_, objc.Sel("setCenter:"), value)
}

func (c_ CheckerboardGeneratorObject) HasCenter() bool {
	return c_.RespondsToSelector(objc.Sel("center"))
}

// The center of the effect as x and y coordinates. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cicheckerboardgenerator/3228105-center?language=objc
func (c_ CheckerboardGeneratorObject) Center() coregraphics.Point {
	rv := objc.Call[coregraphics.Point](c_, objc.Sel("center"))
	return rv
}

func (c_ CheckerboardGeneratorObject) HasSetColor0() bool {
	return c_.RespondsToSelector(objc.Sel("setColor0:"))
}

// A color to use for the first set of squares. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cicheckerboardgenerator/3228106-color0?language=objc
func (c_ CheckerboardGeneratorObject) SetColor0(value Color) {
	objc.Call[objc.Void](c_, objc.Sel("setColor0:"), value)
}

func (c_ CheckerboardGeneratorObject) HasColor0() bool {
	return c_.RespondsToSelector(objc.Sel("color0"))
}

// A color to use for the first set of squares. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cicheckerboardgenerator/3228106-color0?language=objc
func (c_ CheckerboardGeneratorObject) Color0() Color {
	rv := objc.Call[Color](c_, objc.Sel("color0"))
	return rv
}

func (c_ CheckerboardGeneratorObject) HasSetWidth() bool {
	return c_.RespondsToSelector(objc.Sel("setWidth:"))
}

// The width of the squares in the pattern. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cicheckerboardgenerator/3228109-width?language=objc
func (c_ CheckerboardGeneratorObject) SetWidth(value float32) {
	objc.Call[objc.Void](c_, objc.Sel("setWidth:"), value)
}

func (c_ CheckerboardGeneratorObject) HasWidth() bool {
	return c_.RespondsToSelector(objc.Sel("width"))
}

// The width of the squares in the pattern. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cicheckerboardgenerator/3228109-width?language=objc
func (c_ CheckerboardGeneratorObject) Width() float32 {
	rv := objc.Call[float32](c_, objc.Sel("width"))
	return rv
}

func (c_ CheckerboardGeneratorObject) HasSetColor1() bool {
	return c_.RespondsToSelector(objc.Sel("setColor1:"))
}

// A color to use for the second set of squares. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cicheckerboardgenerator/3228107-color1?language=objc
func (c_ CheckerboardGeneratorObject) SetColor1(value Color) {
	objc.Call[objc.Void](c_, objc.Sel("setColor1:"), value)
}

func (c_ CheckerboardGeneratorObject) HasColor1() bool {
	return c_.RespondsToSelector(objc.Sel("color1"))
}

// A color to use for the second set of squares. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cicheckerboardgenerator/3228107-color1?language=objc
func (c_ CheckerboardGeneratorObject) Color1() Color {
	rv := objc.Call[Color](c_, objc.Sel("color1"))
	return rv
}

func (c_ CheckerboardGeneratorObject) HasSetSharpness() bool {
	return c_.RespondsToSelector(objc.Sel("setSharpness:"))
}

// The sharpness of the edges in the pattern. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cicheckerboardgenerator/3228108-sharpness?language=objc
func (c_ CheckerboardGeneratorObject) SetSharpness(value float32) {
	objc.Call[objc.Void](c_, objc.Sel("setSharpness:"), value)
}

func (c_ CheckerboardGeneratorObject) HasSharpness() bool {
	return c_.RespondsToSelector(objc.Sel("sharpness"))
}

// The sharpness of the edges in the pattern. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cicheckerboardgenerator/3228108-sharpness?language=objc
func (c_ CheckerboardGeneratorObject) Sharpness() float32 {
	rv := objc.Call[float32](c_, objc.Sel("sharpness"))
	return rv
}
