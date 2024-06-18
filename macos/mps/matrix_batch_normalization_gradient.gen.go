// Code generated by DarwinKit. DO NOT EDIT.

package mps

import (
	"unsafe"

	"github.com/progrium/darwinkit/macos/metal"
	"github.com/progrium/darwinkit/objc"
)

// The class instance for the [MatrixBatchNormalizationGradient] class.
var MatrixBatchNormalizationGradientClass = _MatrixBatchNormalizationGradientClass{objc.GetClass("MPSMatrixBatchNormalizationGradient")}

type _MatrixBatchNormalizationGradientClass struct {
	objc.Class
}

// An interface definition for the [MatrixBatchNormalizationGradient] class.
type IMatrixBatchNormalizationGradient interface {
	IMatrixBinaryKernel
	SetNeuronTypeParameterAParameterBParameterC(neuronType CNNNeuronType, parameterA float32, parameterB float32, parameterC float32)
	NeuronParameterB() float32
	NeuronParameterA() float32
	NeuronParameterC() float32
	NeuronType() CNNNeuronType
	EncodeToCommandBufferGradientMatrixInputMatrixMeanVectorVarianceVectorGammaVectorBetaVectorResultGradientForDataMatrixResultGradientForGammaVectorResultGradientForBetaVector(commandBuffer metal.PCommandBuffer, gradientMatrix IMatrix, inputMatrix IMatrix, meanVector IVector, varianceVector IVector, gammaVector IVector, betaVector IVector, resultGradientForDataMatrix IMatrix, resultGradientForGammaVector IVector, resultGradientForBetaVector IVector)
	EncodeToCommandBufferObjectGradientMatrixInputMatrixMeanVectorVarianceVectorGammaVectorBetaVectorResultGradientForDataMatrixResultGradientForGammaVectorResultGradientForBetaVector(commandBufferObject objc.IObject, gradientMatrix IMatrix, inputMatrix IMatrix, meanVector IVector, varianceVector IVector, gammaVector IVector, betaVector IVector, resultGradientForDataMatrix IMatrix, resultGradientForGammaVector IVector, resultGradientForBetaVector IVector)
	SourceNumberOfFeatureVectors() uint
	SetSourceNumberOfFeatureVectors(value uint)
	Epsilon() float32
	SetEpsilon(value float32)
	SourceInputFeatureChannels() uint
	SetSourceInputFeatureChannels(value uint)
}

// A batch normalization gradient kernel that operates on matrices. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsmatrixbatchnormalizationgradient?language=objc
type MatrixBatchNormalizationGradient struct {
	MatrixBinaryKernel
}

func MatrixBatchNormalizationGradientFrom(ptr unsafe.Pointer) MatrixBatchNormalizationGradient {
	return MatrixBatchNormalizationGradient{
		MatrixBinaryKernel: MatrixBinaryKernelFrom(ptr),
	}
}

func (m_ MatrixBatchNormalizationGradient) InitWithDevice(device metal.PDevice) MatrixBatchNormalizationGradient {
	po0 := objc.WrapAsProtocol("MTLDevice", device)
	rv := objc.Call[MatrixBatchNormalizationGradient](m_, objc.Sel("initWithDevice:"), po0)
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsmatrixbatchnormalizationgradient/2980748-initwithdevice?language=objc
func NewMatrixBatchNormalizationGradientWithDevice(device metal.PDevice) MatrixBatchNormalizationGradient {
	instance := MatrixBatchNormalizationGradientClass.Alloc().InitWithDevice(device)
	instance.Autorelease()
	return instance
}

func (m_ MatrixBatchNormalizationGradient) CopyWithZoneDevice(zone unsafe.Pointer, device metal.PDevice) MatrixBatchNormalizationGradient {
	po1 := objc.WrapAsProtocol("MTLDevice", device)
	rv := objc.Call[MatrixBatchNormalizationGradient](m_, objc.Sel("copyWithZone:device:"), zone, po1)
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsmatrixbatchnormalizationgradient/2980744-copywithzone?language=objc
func MatrixBatchNormalizationGradient_CopyWithZoneDevice(zone unsafe.Pointer, device metal.PDevice) MatrixBatchNormalizationGradient {
	instance := MatrixBatchNormalizationGradientClass.Alloc().CopyWithZoneDevice(zone, device)
	instance.Autorelease()
	return instance
}

func (mc _MatrixBatchNormalizationGradientClass) Alloc() MatrixBatchNormalizationGradient {
	rv := objc.Call[MatrixBatchNormalizationGradient](mc, objc.Sel("alloc"))
	return rv
}

func (mc _MatrixBatchNormalizationGradientClass) New() MatrixBatchNormalizationGradient {
	rv := objc.Call[MatrixBatchNormalizationGradient](mc, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewMatrixBatchNormalizationGradient() MatrixBatchNormalizationGradient {
	return MatrixBatchNormalizationGradientClass.New()
}

func (m_ MatrixBatchNormalizationGradient) Init() MatrixBatchNormalizationGradient {
	rv := objc.Call[MatrixBatchNormalizationGradient](m_, objc.Sel("init"))
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsmatrixbatchnormalizationgradient/2980753-setneurontype?language=objc
func (m_ MatrixBatchNormalizationGradient) SetNeuronTypeParameterAParameterBParameterC(neuronType CNNNeuronType, parameterA float32, parameterB float32, parameterC float32) {
	objc.Call[objc.Void](m_, objc.Sel("setNeuronType:parameterA:parameterB:parameterC:"), neuronType, parameterA, parameterB, parameterC)
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsmatrixbatchnormalizationgradient/2980750-neuronparameterb?language=objc
func (m_ MatrixBatchNormalizationGradient) NeuronParameterB() float32 {
	rv := objc.Call[float32](m_, objc.Sel("neuronParameterB"))
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsmatrixbatchnormalizationgradient/2980749-neuronparametera?language=objc
func (m_ MatrixBatchNormalizationGradient) NeuronParameterA() float32 {
	rv := objc.Call[float32](m_, objc.Sel("neuronParameterA"))
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsmatrixbatchnormalizationgradient/2980751-neuronparameterc?language=objc
func (m_ MatrixBatchNormalizationGradient) NeuronParameterC() float32 {
	rv := objc.Call[float32](m_, objc.Sel("neuronParameterC"))
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsmatrixbatchnormalizationgradient/2980752-neurontype?language=objc
func (m_ MatrixBatchNormalizationGradient) NeuronType() CNNNeuronType {
	rv := objc.Call[CNNNeuronType](m_, objc.Sel("neuronType"))
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsmatrixbatchnormalizationgradient/2980745-encodetocommandbuffer?language=objc
func (m_ MatrixBatchNormalizationGradient) EncodeToCommandBufferGradientMatrixInputMatrixMeanVectorVarianceVectorGammaVectorBetaVectorResultGradientForDataMatrixResultGradientForGammaVectorResultGradientForBetaVector(commandBuffer metal.PCommandBuffer, gradientMatrix IMatrix, inputMatrix IMatrix, meanVector IVector, varianceVector IVector, gammaVector IVector, betaVector IVector, resultGradientForDataMatrix IMatrix, resultGradientForGammaVector IVector, resultGradientForBetaVector IVector) {
	po0 := objc.WrapAsProtocol("MTLCommandBuffer", commandBuffer)
	objc.Call[objc.Void](m_, objc.Sel("encodeToCommandBuffer:gradientMatrix:inputMatrix:meanVector:varianceVector:gammaVector:betaVector:resultGradientForDataMatrix:resultGradientForGammaVector:resultGradientForBetaVector:"), po0, gradientMatrix, inputMatrix, meanVector, varianceVector, gammaVector, betaVector, resultGradientForDataMatrix, resultGradientForGammaVector, resultGradientForBetaVector)
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsmatrixbatchnormalizationgradient/2980745-encodetocommandbuffer?language=objc
func (m_ MatrixBatchNormalizationGradient) EncodeToCommandBufferObjectGradientMatrixInputMatrixMeanVectorVarianceVectorGammaVectorBetaVectorResultGradientForDataMatrixResultGradientForGammaVectorResultGradientForBetaVector(commandBufferObject objc.IObject, gradientMatrix IMatrix, inputMatrix IMatrix, meanVector IVector, varianceVector IVector, gammaVector IVector, betaVector IVector, resultGradientForDataMatrix IMatrix, resultGradientForGammaVector IVector, resultGradientForBetaVector IVector) {
	objc.Call[objc.Void](m_, objc.Sel("encodeToCommandBuffer:gradientMatrix:inputMatrix:meanVector:varianceVector:gammaVector:betaVector:resultGradientForDataMatrix:resultGradientForGammaVector:resultGradientForBetaVector:"), commandBufferObject, gradientMatrix, inputMatrix, meanVector, varianceVector, gammaVector, betaVector, resultGradientForDataMatrix, resultGradientForGammaVector, resultGradientForBetaVector)
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsmatrixbatchnormalizationgradient/2980755-sourcenumberoffeaturevectors?language=objc
func (m_ MatrixBatchNormalizationGradient) SourceNumberOfFeatureVectors() uint {
	rv := objc.Call[uint](m_, objc.Sel("sourceNumberOfFeatureVectors"))
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsmatrixbatchnormalizationgradient/2980755-sourcenumberoffeaturevectors?language=objc
func (m_ MatrixBatchNormalizationGradient) SetSourceNumberOfFeatureVectors(value uint) {
	objc.Call[objc.Void](m_, objc.Sel("setSourceNumberOfFeatureVectors:"), value)
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsmatrixbatchnormalizationgradient/2980746-epsilon?language=objc
func (m_ MatrixBatchNormalizationGradient) Epsilon() float32 {
	rv := objc.Call[float32](m_, objc.Sel("epsilon"))
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsmatrixbatchnormalizationgradient/2980746-epsilon?language=objc
func (m_ MatrixBatchNormalizationGradient) SetEpsilon(value float32) {
	objc.Call[objc.Void](m_, objc.Sel("setEpsilon:"), value)
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsmatrixbatchnormalizationgradient/2980754-sourceinputfeaturechannels?language=objc
func (m_ MatrixBatchNormalizationGradient) SourceInputFeatureChannels() uint {
	rv := objc.Call[uint](m_, objc.Sel("sourceInputFeatureChannels"))
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsmatrixbatchnormalizationgradient/2980754-sourceinputfeaturechannels?language=objc
func (m_ MatrixBatchNormalizationGradient) SetSourceInputFeatureChannels(value uint) {
	objc.Call[objc.Void](m_, objc.Sel("setSourceInputFeatureChannels:"), value)
}
