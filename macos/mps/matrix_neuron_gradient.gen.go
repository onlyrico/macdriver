// Code generated by DarwinKit. DO NOT EDIT.

package mps

import (
	"unsafe"

	"github.com/progrium/darwinkit/macos/metal"
	"github.com/progrium/darwinkit/objc"
)

// The class instance for the [MatrixNeuronGradient] class.
var MatrixNeuronGradientClass = _MatrixNeuronGradientClass{objc.GetClass("MPSMatrixNeuronGradient")}

type _MatrixNeuronGradientClass struct {
	objc.Class
}

// An interface definition for the [MatrixNeuronGradient] class.
type IMatrixNeuronGradient interface {
	IMatrixBinaryKernel
	EncodeToCommandBufferGradientMatrixInputMatrixBiasVectorResultGradientForDataMatrixResultGradientForBiasVector(commandBuffer metal.PCommandBuffer, gradientMatrix IMatrix, inputMatrix IMatrix, biasVector IVector, resultGradientForDataMatrix IMatrix, resultGradientForBiasVector IVector)
	EncodeToCommandBufferObjectGradientMatrixInputMatrixBiasVectorResultGradientForDataMatrixResultGradientForBiasVector(commandBufferObject objc.IObject, gradientMatrix IMatrix, inputMatrix IMatrix, biasVector IVector, resultGradientForDataMatrix IMatrix, resultGradientForBiasVector IVector)
	SetNeuronToPReLUWithParametersA(A []byte)
	NeuronParameterB() float32
	SetNeuronTypeParameterAParameterBParameterC(neuronType CNNNeuronType, parameterA float32, parameterB float32, parameterC float32)
	NeuronParameterA() float32
	NeuronType() CNNNeuronType
	NeuronParameterC() float32
	SourceNumberOfFeatureVectors() uint
	SetSourceNumberOfFeatureVectors(value uint)
	Alpha() float64
	SetAlpha(value float64)
	SourceInputFeatureChannels() uint
	SetSourceInputFeatureChannels(value uint)
}

// A gradient neuron activation kernel that operates on matrices. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsmatrixneurongradient?language=objc
type MatrixNeuronGradient struct {
	MatrixBinaryKernel
}

func MatrixNeuronGradientFrom(ptr unsafe.Pointer) MatrixNeuronGradient {
	return MatrixNeuronGradient{
		MatrixBinaryKernel: MatrixBinaryKernelFrom(ptr),
	}
}

func (m_ MatrixNeuronGradient) CopyWithZoneDevice(zone unsafe.Pointer, device metal.PDevice) MatrixNeuronGradient {
	po1 := objc.WrapAsProtocol("MTLDevice", device)
	rv := objc.Call[MatrixNeuronGradient](m_, objc.Sel("copyWithZone:device:"), zone, po1)
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsmatrixneurongradient/2966674-copywithzone?language=objc
func MatrixNeuronGradient_CopyWithZoneDevice(zone unsafe.Pointer, device metal.PDevice) MatrixNeuronGradient {
	instance := MatrixNeuronGradientClass.Alloc().CopyWithZoneDevice(zone, device)
	instance.Autorelease()
	return instance
}

func (m_ MatrixNeuronGradient) InitWithDevice(device metal.PDevice) MatrixNeuronGradient {
	po0 := objc.WrapAsProtocol("MTLDevice", device)
	rv := objc.Call[MatrixNeuronGradient](m_, objc.Sel("initWithDevice:"), po0)
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsmatrixneurongradient/2966677-initwithdevice?language=objc
func NewMatrixNeuronGradientWithDevice(device metal.PDevice) MatrixNeuronGradient {
	instance := MatrixNeuronGradientClass.Alloc().InitWithDevice(device)
	instance.Autorelease()
	return instance
}

func (mc _MatrixNeuronGradientClass) Alloc() MatrixNeuronGradient {
	rv := objc.Call[MatrixNeuronGradient](mc, objc.Sel("alloc"))
	return rv
}

func (mc _MatrixNeuronGradientClass) New() MatrixNeuronGradient {
	rv := objc.Call[MatrixNeuronGradient](mc, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewMatrixNeuronGradient() MatrixNeuronGradient {
	return MatrixNeuronGradientClass.New()
}

func (m_ MatrixNeuronGradient) Init() MatrixNeuronGradient {
	rv := objc.Call[MatrixNeuronGradient](m_, objc.Sel("init"))
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsmatrixneurongradient/2966675-encodetocommandbuffer?language=objc
func (m_ MatrixNeuronGradient) EncodeToCommandBufferGradientMatrixInputMatrixBiasVectorResultGradientForDataMatrixResultGradientForBiasVector(commandBuffer metal.PCommandBuffer, gradientMatrix IMatrix, inputMatrix IMatrix, biasVector IVector, resultGradientForDataMatrix IMatrix, resultGradientForBiasVector IVector) {
	po0 := objc.WrapAsProtocol("MTLCommandBuffer", commandBuffer)
	objc.Call[objc.Void](m_, objc.Sel("encodeToCommandBuffer:gradientMatrix:inputMatrix:biasVector:resultGradientForDataMatrix:resultGradientForBiasVector:"), po0, gradientMatrix, inputMatrix, biasVector, resultGradientForDataMatrix, resultGradientForBiasVector)
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsmatrixneurongradient/2966675-encodetocommandbuffer?language=objc
func (m_ MatrixNeuronGradient) EncodeToCommandBufferObjectGradientMatrixInputMatrixBiasVectorResultGradientForDataMatrixResultGradientForBiasVector(commandBufferObject objc.IObject, gradientMatrix IMatrix, inputMatrix IMatrix, biasVector IVector, resultGradientForDataMatrix IMatrix, resultGradientForBiasVector IVector) {
	objc.Call[objc.Void](m_, objc.Sel("encodeToCommandBuffer:gradientMatrix:inputMatrix:biasVector:resultGradientForDataMatrix:resultGradientForBiasVector:"), commandBufferObject, gradientMatrix, inputMatrix, biasVector, resultGradientForDataMatrix, resultGradientForBiasVector)
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsmatrixneurongradient/2966682-setneurontopreluwithparametersa?language=objc
func (m_ MatrixNeuronGradient) SetNeuronToPReLUWithParametersA(A []byte) {
	objc.Call[objc.Void](m_, objc.Sel("setNeuronToPReLUWithParametersA:"), A)
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsmatrixneurongradient/2966679-neuronparameterb?language=objc
func (m_ MatrixNeuronGradient) NeuronParameterB() float32 {
	rv := objc.Call[float32](m_, objc.Sel("neuronParameterB"))
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsmatrixneurongradient/2966683-setneurontype?language=objc
func (m_ MatrixNeuronGradient) SetNeuronTypeParameterAParameterBParameterC(neuronType CNNNeuronType, parameterA float32, parameterB float32, parameterC float32) {
	objc.Call[objc.Void](m_, objc.Sel("setNeuronType:parameterA:parameterB:parameterC:"), neuronType, parameterA, parameterB, parameterC)
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsmatrixneurongradient/2966678-neuronparametera?language=objc
func (m_ MatrixNeuronGradient) NeuronParameterA() float32 {
	rv := objc.Call[float32](m_, objc.Sel("neuronParameterA"))
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsmatrixneurongradient/2966681-neurontype?language=objc
func (m_ MatrixNeuronGradient) NeuronType() CNNNeuronType {
	rv := objc.Call[CNNNeuronType](m_, objc.Sel("neuronType"))
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsmatrixneurongradient/2966680-neuronparameterc?language=objc
func (m_ MatrixNeuronGradient) NeuronParameterC() float32 {
	rv := objc.Call[float32](m_, objc.Sel("neuronParameterC"))
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsmatrixneurongradient/2966685-sourcenumberoffeaturevectors?language=objc
func (m_ MatrixNeuronGradient) SourceNumberOfFeatureVectors() uint {
	rv := objc.Call[uint](m_, objc.Sel("sourceNumberOfFeatureVectors"))
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsmatrixneurongradient/2966685-sourcenumberoffeaturevectors?language=objc
func (m_ MatrixNeuronGradient) SetSourceNumberOfFeatureVectors(value uint) {
	objc.Call[objc.Void](m_, objc.Sel("setSourceNumberOfFeatureVectors:"), value)
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsmatrixneurongradient/2966673-alpha?language=objc
func (m_ MatrixNeuronGradient) Alpha() float64 {
	rv := objc.Call[float64](m_, objc.Sel("alpha"))
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsmatrixneurongradient/2966673-alpha?language=objc
func (m_ MatrixNeuronGradient) SetAlpha(value float64) {
	objc.Call[objc.Void](m_, objc.Sel("setAlpha:"), value)
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsmatrixneurongradient/2966684-sourceinputfeaturechannels?language=objc
func (m_ MatrixNeuronGradient) SourceInputFeatureChannels() uint {
	rv := objc.Call[uint](m_, objc.Sel("sourceInputFeatureChannels"))
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsmatrixneurongradient/2966684-sourceinputfeaturechannels?language=objc
func (m_ MatrixNeuronGradient) SetSourceInputFeatureChannels(value uint) {
	objc.Call[objc.Void](m_, objc.Sel("setSourceInputFeatureChannels:"), value)
}
