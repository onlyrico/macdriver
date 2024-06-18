// Code generated by DarwinKit. DO NOT EDIT.

package mpsgraph

import (
	"unsafe"

	"github.com/progrium/darwinkit/macos/foundation"
	"github.com/progrium/darwinkit/macos/metal"
	"github.com/progrium/darwinkit/macos/mps"
	"github.com/progrium/darwinkit/objc"
)

// The class instance for the [TensorData] class.
var TensorDataClass = _TensorDataClass{objc.GetClass("MPSGraphTensorData")}

type _TensorDataClass struct {
	objc.Class
}

// An interface definition for the [TensorData] class.
type ITensorData interface {
	objc.IObject
	Mpsndarray() mps.NDArray
	Shape() *foundation.Array
	Device() Device
	DataType() mps.DataType
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshadersgraph/mpsgraphtensordata?language=objc
type TensorData struct {
	objc.Object
}

func TensorDataFrom(ptr unsafe.Pointer) TensorData {
	return TensorData{
		Object: objc.ObjectFrom(ptr),
	}
}

func (t_ TensorData) InitWithMPSImageBatch(imageBatch *foundation.Array) TensorData {
	rv := objc.Call[TensorData](t_, objc.Sel("initWithMPSImageBatch:"), imageBatch)
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshadersgraph/mpsgraphtensordata/3564672-initwithmpsimagebatch?language=objc
func NewTensorDataWithMPSImageBatch(imageBatch *foundation.Array) TensorData {
	instance := TensorDataClass.Alloc().InitWithMPSImageBatch(imageBatch)
	instance.Autorelease()
	return instance
}

func (t_ TensorData) InitWithMPSVector(vector mps.IVector) TensorData {
	rv := objc.Call[TensorData](t_, objc.Sel("initWithMPSVector:"), vector)
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshadersgraph/mpsgraphtensordata/3564675-initwithmpsvector?language=objc
func NewTensorDataWithMPSVector(vector mps.IVector) TensorData {
	instance := TensorDataClass.Alloc().InitWithMPSVector(vector)
	instance.Autorelease()
	return instance
}

func (t_ TensorData) InitWithMPSMatrix(matrix mps.IMatrix) TensorData {
	rv := objc.Call[TensorData](t_, objc.Sel("initWithMPSMatrix:"), matrix)
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshadersgraph/mpsgraphtensordata/3564673-initwithmpsmatrix?language=objc
func NewTensorDataWithMPSMatrix(matrix mps.IMatrix) TensorData {
	instance := TensorDataClass.Alloc().InitWithMPSMatrix(matrix)
	instance.Autorelease()
	return instance
}

func (t_ TensorData) InitWithMTLBufferShapeDataType(buffer metal.PBuffer, shape *foundation.Array, dataType mps.DataType) TensorData {
	po0 := objc.WrapAsProtocol("MTLBuffer", buffer)
	rv := objc.Call[TensorData](t_, objc.Sel("initWithMTLBuffer:shape:dataType:"), po0, shape, dataType)
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshadersgraph/mpsgraphtensordata/3564676-initwithmtlbuffer?language=objc
func NewTensorDataWithMTLBufferShapeDataType(buffer metal.PBuffer, shape *foundation.Array, dataType mps.DataType) TensorData {
	instance := TensorDataClass.Alloc().InitWithMTLBufferShapeDataType(buffer, shape, dataType)
	instance.Autorelease()
	return instance
}

func (t_ TensorData) InitWithDeviceDataShapeDataType(device IDevice, data []byte, shape *foundation.Array, dataType mps.DataType) TensorData {
	rv := objc.Call[TensorData](t_, objc.Sel("initWithDevice:data:shape:dataType:"), device, data, shape, dataType)
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshadersgraph/mpsgraphtensordata/3564671-initwithdevice?language=objc
func NewTensorDataWithDeviceDataShapeDataType(device IDevice, data []byte, shape *foundation.Array, dataType mps.DataType) TensorData {
	instance := TensorDataClass.Alloc().InitWithDeviceDataShapeDataType(device, data, shape, dataType)
	instance.Autorelease()
	return instance
}

func (t_ TensorData) InitWithMTLBufferShapeDataTypeRowBytes(buffer metal.PBuffer, shape *foundation.Array, dataType mps.DataType, rowBytes uint) TensorData {
	po0 := objc.WrapAsProtocol("MTLBuffer", buffer)
	rv := objc.Call[TensorData](t_, objc.Sel("initWithMTLBuffer:shape:dataType:rowBytes:"), po0, shape, dataType, rowBytes)
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshadersgraph/mpsgraphtensordata/3925447-initwithmtlbuffer?language=objc
func NewTensorDataWithMTLBufferShapeDataTypeRowBytes(buffer metal.PBuffer, shape *foundation.Array, dataType mps.DataType, rowBytes uint) TensorData {
	instance := TensorDataClass.Alloc().InitWithMTLBufferShapeDataTypeRowBytes(buffer, shape, dataType, rowBytes)
	instance.Autorelease()
	return instance
}

func (t_ TensorData) InitWithMPSMatrixRank(matrix mps.IMatrix, rank uint) TensorData {
	rv := objc.Call[TensorData](t_, objc.Sel("initWithMPSMatrix:rank:"), matrix, rank)
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshadersgraph/mpsgraphtensordata/3667507-initwithmpsmatrix?language=objc
func NewTensorDataWithMPSMatrixRank(matrix mps.IMatrix, rank uint) TensorData {
	instance := TensorDataClass.Alloc().InitWithMPSMatrixRank(matrix, rank)
	instance.Autorelease()
	return instance
}

func (t_ TensorData) InitWithMPSNDArray(ndarray mps.INDArray) TensorData {
	rv := objc.Call[TensorData](t_, objc.Sel("initWithMPSNDArray:"), ndarray)
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshadersgraph/mpsgraphtensordata/3564674-initwithmpsndarray?language=objc
func NewTensorDataWithMPSNDArray(ndarray mps.INDArray) TensorData {
	instance := TensorDataClass.Alloc().InitWithMPSNDArray(ndarray)
	instance.Autorelease()
	return instance
}

func (t_ TensorData) InitWithMPSVectorRank(vector mps.IVector, rank uint) TensorData {
	rv := objc.Call[TensorData](t_, objc.Sel("initWithMPSVector:rank:"), vector, rank)
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshadersgraph/mpsgraphtensordata/3667508-initwithmpsvector?language=objc
func NewTensorDataWithMPSVectorRank(vector mps.IVector, rank uint) TensorData {
	instance := TensorDataClass.Alloc().InitWithMPSVectorRank(vector, rank)
	instance.Autorelease()
	return instance
}

func (tc _TensorDataClass) Alloc() TensorData {
	rv := objc.Call[TensorData](tc, objc.Sel("alloc"))
	return rv
}

func (tc _TensorDataClass) New() TensorData {
	rv := objc.Call[TensorData](tc, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewTensorData() TensorData {
	return TensorDataClass.New()
}

func (t_ TensorData) Init() TensorData {
	rv := objc.Call[TensorData](t_, objc.Sel("init"))
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshadersgraph/mpsgraphtensordata/3564677-mpsndarray?language=objc
func (t_ TensorData) Mpsndarray() mps.NDArray {
	rv := objc.Call[mps.NDArray](t_, objc.Sel("mpsndarray"))
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshadersgraph/mpsgraphtensordata/3564678-shape?language=objc
func (t_ TensorData) Shape() *foundation.Array {
	rv := objc.Call[*foundation.Array](t_, objc.Sel("shape"))
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshadersgraph/mpsgraphtensordata/3564670-device?language=objc
func (t_ TensorData) Device() Device {
	rv := objc.Call[Device](t_, objc.Sel("device"))
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshadersgraph/mpsgraphtensordata/3564669-datatype?language=objc
func (t_ TensorData) DataType() mps.DataType {
	rv := objc.Call[mps.DataType](t_, objc.Sel("dataType"))
	return rv
}
