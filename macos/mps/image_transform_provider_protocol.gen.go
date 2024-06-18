// Code generated by DarwinKit. DO NOT EDIT.

package mps

import (
	"github.com/progrium/darwinkit/objc"
)

// A general interface for objects that provide image resampling. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsimagetransformprovider?language=objc
type PImageTransformProvider interface {
	// optional
	TransformForSourceImageHandle(image Image, handle HandleObject) ScaleTransform
	HasTransformForSourceImageHandle() bool
}

// ensure impl type implements protocol interface
var _ PImageTransformProvider = (*ImageTransformProviderObject)(nil)

// A concrete type for the [PImageTransformProvider] protocol.
type ImageTransformProviderObject struct {
	objc.Object
}

func (i_ ImageTransformProviderObject) HasTransformForSourceImageHandle() bool {
	return i_.RespondsToSelector(objc.Sel("transformForSourceImage:handle:"))
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsimagetransformprovider/2915282-transformforsourceimage?language=objc
func (i_ ImageTransformProviderObject) TransformForSourceImageHandle(image Image, handle HandleObject) ScaleTransform {
	po1 := objc.WrapAsProtocol("MPSHandle", handle)
	rv := objc.Call[ScaleTransform](i_, objc.Sel("transformForSourceImage:handle:"), image, po1)
	return rv
}
