// Code generated by DarwinKit. DO NOT EDIT.

package metal

import (
	"unsafe"

	"github.com/progrium/darwinkit/objc"
)

// The class instance for the [BufferLayoutDescriptorArray] class.
var BufferLayoutDescriptorArrayClass = _BufferLayoutDescriptorArrayClass{objc.GetClass("MTLBufferLayoutDescriptorArray")}

type _BufferLayoutDescriptorArrayClass struct {
	objc.Class
}

// An interface definition for the [BufferLayoutDescriptorArray] class.
type IBufferLayoutDescriptorArray interface {
	objc.IObject
	SetObjectAtIndexedSubscript(bufferDesc IBufferLayoutDescriptor, index uint)
	ObjectAtIndexedSubscript(index uint) BufferLayoutDescriptor
}

// An array of buffer layout descriptor objects. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlbufferlayoutdescriptorarray?language=objc
type BufferLayoutDescriptorArray struct {
	objc.Object
}

func BufferLayoutDescriptorArrayFrom(ptr unsafe.Pointer) BufferLayoutDescriptorArray {
	return BufferLayoutDescriptorArray{
		Object: objc.ObjectFrom(ptr),
	}
}

func (bc _BufferLayoutDescriptorArrayClass) Alloc() BufferLayoutDescriptorArray {
	rv := objc.Call[BufferLayoutDescriptorArray](bc, objc.Sel("alloc"))
	return rv
}

func (bc _BufferLayoutDescriptorArrayClass) New() BufferLayoutDescriptorArray {
	rv := objc.Call[BufferLayoutDescriptorArray](bc, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewBufferLayoutDescriptorArray() BufferLayoutDescriptorArray {
	return BufferLayoutDescriptorArrayClass.New()
}

func (b_ BufferLayoutDescriptorArray) Init() BufferLayoutDescriptorArray {
	rv := objc.Call[BufferLayoutDescriptorArray](b_, objc.Sel("init"))
	return rv
}

// Sets the state of the specified buffer layout. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlbufferlayoutdescriptorarray/2097295-setobject?language=objc
func (b_ BufferLayoutDescriptorArray) SetObjectAtIndexedSubscript(bufferDesc IBufferLayoutDescriptor, index uint) {
	objc.Call[objc.Void](b_, objc.Sel("setObject:atIndexedSubscript:"), bufferDesc, index)
}

// Returns the state of the specified buffer layout. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlbufferlayoutdescriptorarray/2097228-objectatindexedsubscript?language=objc
func (b_ BufferLayoutDescriptorArray) ObjectAtIndexedSubscript(index uint) BufferLayoutDescriptor {
	rv := objc.Call[BufferLayoutDescriptor](b_, objc.Sel("objectAtIndexedSubscript:"), index)
	return rv
}
