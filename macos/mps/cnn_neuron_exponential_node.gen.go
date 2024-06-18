// Code generated by DarwinKit. DO NOT EDIT.

package mps

import (
	"unsafe"

	"github.com/progrium/darwinkit/objc"
)

// The class instance for the [CNNNeuronExponentialNode] class.
var CNNNeuronExponentialNodeClass = _CNNNeuronExponentialNodeClass{objc.GetClass("MPSCNNNeuronExponentialNode")}

type _CNNNeuronExponentialNodeClass struct {
	objc.Class
}

// An interface definition for the [CNNNeuronExponentialNode] class.
type ICNNNeuronExponentialNode interface {
	ICNNNeuronNode
}

// A representation of an exponential neuron filter. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnneuronexponentialnode?language=objc
type CNNNeuronExponentialNode struct {
	CNNNeuronNode
}

func CNNNeuronExponentialNodeFrom(ptr unsafe.Pointer) CNNNeuronExponentialNode {
	return CNNNeuronExponentialNode{
		CNNNeuronNode: CNNNeuronNodeFrom(ptr),
	}
}

func (cc _CNNNeuronExponentialNodeClass) NodeWithSourceABC(sourceNode INNImageNode, a float32, b float32, c float32) CNNNeuronExponentialNode {
	rv := objc.Call[CNNNeuronExponentialNode](cc, objc.Sel("nodeWithSource:a:b:c:"), sourceNode, a, b, c)
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnneuronexponentialnode/2951959-nodewithsource?language=objc
func CNNNeuronExponentialNode_NodeWithSourceABC(sourceNode INNImageNode, a float32, b float32, c float32) CNNNeuronExponentialNode {
	return CNNNeuronExponentialNodeClass.NodeWithSourceABC(sourceNode, a, b, c)
}

func (cc _CNNNeuronExponentialNodeClass) NodeWithSource(sourceNode INNImageNode) CNNNeuronExponentialNode {
	rv := objc.Call[CNNNeuronExponentialNode](cc, objc.Sel("nodeWithSource:"), sourceNode)
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnneuronexponentialnode/2951950-nodewithsource?language=objc
func CNNNeuronExponentialNode_NodeWithSource(sourceNode INNImageNode) CNNNeuronExponentialNode {
	return CNNNeuronExponentialNodeClass.NodeWithSource(sourceNode)
}

func (c_ CNNNeuronExponentialNode) InitWithSource(sourceNode INNImageNode) CNNNeuronExponentialNode {
	rv := objc.Call[CNNNeuronExponentialNode](c_, objc.Sel("initWithSource:"), sourceNode)
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnneuronexponentialnode/2951936-initwithsource?language=objc
func NewCNNNeuronExponentialNodeWithSource(sourceNode INNImageNode) CNNNeuronExponentialNode {
	instance := CNNNeuronExponentialNodeClass.Alloc().InitWithSource(sourceNode)
	instance.Autorelease()
	return instance
}

func (c_ CNNNeuronExponentialNode) InitWithSourceABC(sourceNode INNImageNode, a float32, b float32, c float32) CNNNeuronExponentialNode {
	rv := objc.Call[CNNNeuronExponentialNode](c_, objc.Sel("initWithSource:a:b:c:"), sourceNode, a, b, c)
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnneuronexponentialnode/2951933-initwithsource?language=objc
func NewCNNNeuronExponentialNodeWithSourceABC(sourceNode INNImageNode, a float32, b float32, c float32) CNNNeuronExponentialNode {
	instance := CNNNeuronExponentialNodeClass.Alloc().InitWithSourceABC(sourceNode, a, b, c)
	instance.Autorelease()
	return instance
}

func (cc _CNNNeuronExponentialNodeClass) Alloc() CNNNeuronExponentialNode {
	rv := objc.Call[CNNNeuronExponentialNode](cc, objc.Sel("alloc"))
	return rv
}

func (cc _CNNNeuronExponentialNodeClass) New() CNNNeuronExponentialNode {
	rv := objc.Call[CNNNeuronExponentialNode](cc, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewCNNNeuronExponentialNode() CNNNeuronExponentialNode {
	return CNNNeuronExponentialNodeClass.New()
}

func (c_ CNNNeuronExponentialNode) Init() CNNNeuronExponentialNode {
	rv := objc.Call[CNNNeuronExponentialNode](c_, objc.Sel("init"))
	return rv
}

func (cc _CNNNeuronExponentialNodeClass) NodeWithSourceDescriptor(sourceNode INNImageNode, descriptor INNNeuronDescriptor) CNNNeuronExponentialNode {
	rv := objc.Call[CNNNeuronExponentialNode](cc, objc.Sel("nodeWithSource:descriptor:"), sourceNode, descriptor)
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnneuronnode/3019333-nodewithsource?language=objc
func CNNNeuronExponentialNode_NodeWithSourceDescriptor(sourceNode INNImageNode, descriptor INNNeuronDescriptor) CNNNeuronExponentialNode {
	return CNNNeuronExponentialNodeClass.NodeWithSourceDescriptor(sourceNode, descriptor)
}
