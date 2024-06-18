// Code generated by DarwinKit. DO NOT EDIT.

package mps

import (
	"unsafe"

	"github.com/progrium/darwinkit/macos/metal"
	"github.com/progrium/darwinkit/objc"
)

// The class instance for the [NNScaleNode] class.
var NNScaleNodeClass = _NNScaleNodeClass{objc.GetClass("MPSNNScaleNode")}

type _NNScaleNodeClass struct {
	objc.Class
}

// An interface definition for the [NNScaleNode] class.
type INNScaleNode interface {
	INNFilterNode
}

// Abstract node representing an image resampling filter. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsnnscalenode?language=objc
type NNScaleNode struct {
	NNFilterNode
}

func NNScaleNodeFrom(ptr unsafe.Pointer) NNScaleNode {
	return NNScaleNode{
		NNFilterNode: NNFilterNodeFrom(ptr),
	}
}

func (n_ NNScaleNode) InitWithSourceOutputSize(sourceNode INNImageNode, size metal.Size) NNScaleNode {
	rv := objc.Call[NNScaleNode](n_, objc.Sel("initWithSource:outputSize:"), sourceNode, size)
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsnnscalenode/2915285-initwithsource?language=objc
func NewNNScaleNodeWithSourceOutputSize(sourceNode INNImageNode, size metal.Size) NNScaleNode {
	instance := NNScaleNodeClass.Alloc().InitWithSourceOutputSize(sourceNode, size)
	instance.Autorelease()
	return instance
}

func (n_ NNScaleNode) InitWithSourceTransformProviderOutputSize(sourceNode INNImageNode, transformProvider PImageTransformProvider, size metal.Size) NNScaleNode {
	po1 := objc.WrapAsProtocol("MPSImageTransformProvider", transformProvider)
	rv := objc.Call[NNScaleNode](n_, objc.Sel("initWithSource:transformProvider:outputSize:"), sourceNode, po1, size)
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsnnscalenode/2915278-initwithsource?language=objc
func NewNNScaleNodeWithSourceTransformProviderOutputSize(sourceNode INNImageNode, transformProvider PImageTransformProvider, size metal.Size) NNScaleNode {
	instance := NNScaleNodeClass.Alloc().InitWithSourceTransformProviderOutputSize(sourceNode, transformProvider, size)
	instance.Autorelease()
	return instance
}

func (nc _NNScaleNodeClass) NodeWithSourceOutputSize(sourceNode INNImageNode, size metal.Size) NNScaleNode {
	rv := objc.Call[NNScaleNode](nc, objc.Sel("nodeWithSource:outputSize:"), sourceNode, size)
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsnnscalenode/2915280-nodewithsource?language=objc
func NNScaleNode_NodeWithSourceOutputSize(sourceNode INNImageNode, size metal.Size) NNScaleNode {
	return NNScaleNodeClass.NodeWithSourceOutputSize(sourceNode, size)
}

func (nc _NNScaleNodeClass) NodeWithSourceTransformProviderOutputSize(sourceNode INNImageNode, transformProvider PImageTransformProvider, size metal.Size) NNScaleNode {
	po1 := objc.WrapAsProtocol("MPSImageTransformProvider", transformProvider)
	rv := objc.Call[NNScaleNode](nc, objc.Sel("nodeWithSource:transformProvider:outputSize:"), sourceNode, po1, size)
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsnnscalenode/2915286-nodewithsource?language=objc
func NNScaleNode_NodeWithSourceTransformProviderOutputSize(sourceNode INNImageNode, transformProvider PImageTransformProvider, size metal.Size) NNScaleNode {
	return NNScaleNodeClass.NodeWithSourceTransformProviderOutputSize(sourceNode, transformProvider, size)
}

func (nc _NNScaleNodeClass) Alloc() NNScaleNode {
	rv := objc.Call[NNScaleNode](nc, objc.Sel("alloc"))
	return rv
}

func (nc _NNScaleNodeClass) New() NNScaleNode {
	rv := objc.Call[NNScaleNode](nc, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewNNScaleNode() NNScaleNode {
	return NNScaleNodeClass.New()
}

func (n_ NNScaleNode) Init() NNScaleNode {
	rv := objc.Call[NNScaleNode](n_, objc.Sel("init"))
	return rv
}
