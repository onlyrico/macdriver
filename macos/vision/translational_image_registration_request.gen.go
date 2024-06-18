// Code generated by DarwinKit. DO NOT EDIT.

package vision

import (
	"unsafe"

	"github.com/progrium/darwinkit/macos/coregraphics"
	"github.com/progrium/darwinkit/macos/coreimage"
	"github.com/progrium/darwinkit/macos/coremedia"
	"github.com/progrium/darwinkit/macos/corevideo"
	"github.com/progrium/darwinkit/macos/foundation"
	"github.com/progrium/darwinkit/macos/imageio"
	"github.com/progrium/darwinkit/objc"
)

// The class instance for the [TranslationalImageRegistrationRequest] class.
var TranslationalImageRegistrationRequestClass = _TranslationalImageRegistrationRequestClass{objc.GetClass("VNTranslationalImageRegistrationRequest")}

type _TranslationalImageRegistrationRequestClass struct {
	objc.Class
}

// An interface definition for the [TranslationalImageRegistrationRequest] class.
type ITranslationalImageRegistrationRequest interface {
	IImageRegistrationRequest
}

// An image analysis request that determines the affine transform necessary to align the content of two images. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/vision/vntranslationalimageregistrationrequest?language=objc
type TranslationalImageRegistrationRequest struct {
	ImageRegistrationRequest
}

func TranslationalImageRegistrationRequestFrom(ptr unsafe.Pointer) TranslationalImageRegistrationRequest {
	return TranslationalImageRegistrationRequest{
		ImageRegistrationRequest: ImageRegistrationRequestFrom(ptr),
	}
}

func (tc _TranslationalImageRegistrationRequestClass) Alloc() TranslationalImageRegistrationRequest {
	rv := objc.Call[TranslationalImageRegistrationRequest](tc, objc.Sel("alloc"))
	return rv
}

func (tc _TranslationalImageRegistrationRequestClass) New() TranslationalImageRegistrationRequest {
	rv := objc.Call[TranslationalImageRegistrationRequest](tc, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewTranslationalImageRegistrationRequest() TranslationalImageRegistrationRequest {
	return TranslationalImageRegistrationRequestClass.New()
}

func (t_ TranslationalImageRegistrationRequest) Init() TranslationalImageRegistrationRequest {
	rv := objc.Call[TranslationalImageRegistrationRequest](t_, objc.Sel("init"))
	return rv
}

func (t_ TranslationalImageRegistrationRequest) InitWithTargetedCGImageOrientationOptions(cgImage coregraphics.ImageRef, orientation imageio.ImagePropertyOrientation, options map[ImageOption]objc.IObject) TranslationalImageRegistrationRequest {
	rv := objc.Call[TranslationalImageRegistrationRequest](t_, objc.Sel("initWithTargetedCGImage:orientation:options:"), cgImage, orientation, options)
	return rv
}

// Creates a new request targeting a Core Graphics image of known orientation. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/vision/vntargetedimagerequest/2923444-initwithtargetedcgimage?language=objc
func NewTranslationalImageRegistrationRequestWithTargetedCGImageOrientationOptions(cgImage coregraphics.ImageRef, orientation imageio.ImagePropertyOrientation, options map[ImageOption]objc.IObject) TranslationalImageRegistrationRequest {
	instance := TranslationalImageRegistrationRequestClass.Alloc().InitWithTargetedCGImageOrientationOptions(cgImage, orientation, options)
	instance.Autorelease()
	return instance
}

func (t_ TranslationalImageRegistrationRequest) InitWithTargetedCGImageOptions(cgImage coregraphics.ImageRef, options map[ImageOption]objc.IObject) TranslationalImageRegistrationRequest {
	rv := objc.Call[TranslationalImageRegistrationRequest](t_, objc.Sel("initWithTargetedCGImage:options:"), cgImage, options)
	return rv
}

// Creates a new request targeting a Core Graphics image. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/vision/vntargetedimagerequest/2923452-initwithtargetedcgimage?language=objc
func NewTranslationalImageRegistrationRequestWithTargetedCGImageOptions(cgImage coregraphics.ImageRef, options map[ImageOption]objc.IObject) TranslationalImageRegistrationRequest {
	instance := TranslationalImageRegistrationRequestClass.Alloc().InitWithTargetedCGImageOptions(cgImage, options)
	instance.Autorelease()
	return instance
}

func (t_ TranslationalImageRegistrationRequest) InitWithTargetedImageURLOptionsCompletionHandler(imageURL foundation.IURL, options map[ImageOption]objc.IObject, completionHandler RequestCompletionHandler) TranslationalImageRegistrationRequest {
	rv := objc.Call[TranslationalImageRegistrationRequest](t_, objc.Sel("initWithTargetedImageURL:options:completionHandler:"), imageURL, options, completionHandler)
	return rv
}

// Creates a new request targeting an image at the specified URL, executing the completion handler when done. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/vision/vntargetedimagerequest/2923453-initwithtargetedimageurl?language=objc
func NewTranslationalImageRegistrationRequestWithTargetedImageURLOptionsCompletionHandler(imageURL foundation.IURL, options map[ImageOption]objc.IObject, completionHandler RequestCompletionHandler) TranslationalImageRegistrationRequest {
	instance := TranslationalImageRegistrationRequestClass.Alloc().InitWithTargetedImageURLOptionsCompletionHandler(imageURL, options, completionHandler)
	instance.Autorelease()
	return instance
}

func (t_ TranslationalImageRegistrationRequest) InitWithTargetedImageDataOptions(imageData []byte, options map[ImageOption]objc.IObject) TranslationalImageRegistrationRequest {
	rv := objc.Call[TranslationalImageRegistrationRequest](t_, objc.Sel("initWithTargetedImageData:options:"), imageData, options)
	return rv
}

// Creates a new request targeting an image as raw data. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/vision/vntargetedimagerequest/2923460-initwithtargetedimagedata?language=objc
func NewTranslationalImageRegistrationRequestWithTargetedImageDataOptions(imageData []byte, options map[ImageOption]objc.IObject) TranslationalImageRegistrationRequest {
	instance := TranslationalImageRegistrationRequestClass.Alloc().InitWithTargetedImageDataOptions(imageData, options)
	instance.Autorelease()
	return instance
}

func (t_ TranslationalImageRegistrationRequest) InitWithTargetedCVPixelBufferOptionsCompletionHandler(pixelBuffer corevideo.PixelBufferRef, options map[ImageOption]objc.IObject, completionHandler RequestCompletionHandler) TranslationalImageRegistrationRequest {
	rv := objc.Call[TranslationalImageRegistrationRequest](t_, objc.Sel("initWithTargetedCVPixelBuffer:options:completionHandler:"), pixelBuffer, options, completionHandler)
	return rv
}

// Creates a new request targeting an image in a CVPixelBufferRef, executing the completion handler when done. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/vision/vntargetedimagerequest/2923446-initwithtargetedcvpixelbuffer?language=objc
func NewTranslationalImageRegistrationRequestWithTargetedCVPixelBufferOptionsCompletionHandler(pixelBuffer corevideo.PixelBufferRef, options map[ImageOption]objc.IObject, completionHandler RequestCompletionHandler) TranslationalImageRegistrationRequest {
	instance := TranslationalImageRegistrationRequestClass.Alloc().InitWithTargetedCVPixelBufferOptionsCompletionHandler(pixelBuffer, options, completionHandler)
	instance.Autorelease()
	return instance
}

func (t_ TranslationalImageRegistrationRequest) InitWithTargetedCMSampleBufferOptionsCompletionHandler(sampleBuffer coremedia.SampleBufferRef, options map[ImageOption]objc.IObject, completionHandler RequestCompletionHandler) TranslationalImageRegistrationRequest {
	rv := objc.Call[TranslationalImageRegistrationRequest](t_, objc.Sel("initWithTargetedCMSampleBuffer:options:completionHandler:"), sampleBuffer, options, completionHandler)
	return rv
}

// Creates a new request with a completion handler that targets an image in a sample buffer. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/vision/vntargetedimagerequest/3571275-initwithtargetedcmsamplebuffer?language=objc
func NewTranslationalImageRegistrationRequestWithTargetedCMSampleBufferOptionsCompletionHandler(sampleBuffer coremedia.SampleBufferRef, options map[ImageOption]objc.IObject, completionHandler RequestCompletionHandler) TranslationalImageRegistrationRequest {
	instance := TranslationalImageRegistrationRequestClass.Alloc().InitWithTargetedCMSampleBufferOptionsCompletionHandler(sampleBuffer, options, completionHandler)
	instance.Autorelease()
	return instance
}

func (t_ TranslationalImageRegistrationRequest) InitWithTargetedImageDataOrientationOptionsCompletionHandler(imageData []byte, orientation imageio.ImagePropertyOrientation, options map[ImageOption]objc.IObject, completionHandler RequestCompletionHandler) TranslationalImageRegistrationRequest {
	rv := objc.Call[TranslationalImageRegistrationRequest](t_, objc.Sel("initWithTargetedImageData:orientation:options:completionHandler:"), imageData, orientation, options, completionHandler)
	return rv
}

// Creates a new request targeting a raw data image of known orientation, executing the completion handler when done. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/vision/vntargetedimagerequest/2923443-initwithtargetedimagedata?language=objc
func NewTranslationalImageRegistrationRequestWithTargetedImageDataOrientationOptionsCompletionHandler(imageData []byte, orientation imageio.ImagePropertyOrientation, options map[ImageOption]objc.IObject, completionHandler RequestCompletionHandler) TranslationalImageRegistrationRequest {
	instance := TranslationalImageRegistrationRequestClass.Alloc().InitWithTargetedImageDataOrientationOptionsCompletionHandler(imageData, orientation, options, completionHandler)
	instance.Autorelease()
	return instance
}

func (t_ TranslationalImageRegistrationRequest) InitWithTargetedCMSampleBufferOrientationOptions(sampleBuffer coremedia.SampleBufferRef, orientation imageio.ImagePropertyOrientation, options map[ImageOption]objc.IObject) TranslationalImageRegistrationRequest {
	rv := objc.Call[TranslationalImageRegistrationRequest](t_, objc.Sel("initWithTargetedCMSampleBuffer:orientation:options:"), sampleBuffer, orientation, options)
	return rv
}

// Creates a new request that targets an image of a known orientation in a sample buffer. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/vision/vntargetedimagerequest/3571276-initwithtargetedcmsamplebuffer?language=objc
func NewTranslationalImageRegistrationRequestWithTargetedCMSampleBufferOrientationOptions(sampleBuffer coremedia.SampleBufferRef, orientation imageio.ImagePropertyOrientation, options map[ImageOption]objc.IObject) TranslationalImageRegistrationRequest {
	instance := TranslationalImageRegistrationRequestClass.Alloc().InitWithTargetedCMSampleBufferOrientationOptions(sampleBuffer, orientation, options)
	instance.Autorelease()
	return instance
}

func (t_ TranslationalImageRegistrationRequest) InitWithTargetedCVPixelBufferOptions(pixelBuffer corevideo.PixelBufferRef, options map[ImageOption]objc.IObject) TranslationalImageRegistrationRequest {
	rv := objc.Call[TranslationalImageRegistrationRequest](t_, objc.Sel("initWithTargetedCVPixelBuffer:options:"), pixelBuffer, options)
	return rv
}

// Creates a new request targeting an image in a CVPixelBufferRef. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/vision/vntargetedimagerequest/2923445-initwithtargetedcvpixelbuffer?language=objc
func NewTranslationalImageRegistrationRequestWithTargetedCVPixelBufferOptions(pixelBuffer corevideo.PixelBufferRef, options map[ImageOption]objc.IObject) TranslationalImageRegistrationRequest {
	instance := TranslationalImageRegistrationRequestClass.Alloc().InitWithTargetedCVPixelBufferOptions(pixelBuffer, options)
	instance.Autorelease()
	return instance
}

func (t_ TranslationalImageRegistrationRequest) InitWithTargetedImageDataOptionsCompletionHandler(imageData []byte, options map[ImageOption]objc.IObject, completionHandler RequestCompletionHandler) TranslationalImageRegistrationRequest {
	rv := objc.Call[TranslationalImageRegistrationRequest](t_, objc.Sel("initWithTargetedImageData:options:completionHandler:"), imageData, options, completionHandler)
	return rv
}

// Creates a new request targeting an image as raw data, executing the completion handler when done. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/vision/vntargetedimagerequest/2923455-initwithtargetedimagedata?language=objc
func NewTranslationalImageRegistrationRequestWithTargetedImageDataOptionsCompletionHandler(imageData []byte, options map[ImageOption]objc.IObject, completionHandler RequestCompletionHandler) TranslationalImageRegistrationRequest {
	instance := TranslationalImageRegistrationRequestClass.Alloc().InitWithTargetedImageDataOptionsCompletionHandler(imageData, options, completionHandler)
	instance.Autorelease()
	return instance
}

func (t_ TranslationalImageRegistrationRequest) InitWithTargetedCIImageOptionsCompletionHandler(ciImage coreimage.IImage, options map[ImageOption]objc.IObject, completionHandler RequestCompletionHandler) TranslationalImageRegistrationRequest {
	rv := objc.Call[TranslationalImageRegistrationRequest](t_, objc.Sel("initWithTargetedCIImage:options:completionHandler:"), ciImage, options, completionHandler)
	return rv
}

// Creates a new request targeting a CIImage, executing the completion handler when done. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/vision/vntargetedimagerequest/2923454-initwithtargetedciimage?language=objc
func NewTranslationalImageRegistrationRequestWithTargetedCIImageOptionsCompletionHandler(ciImage coreimage.IImage, options map[ImageOption]objc.IObject, completionHandler RequestCompletionHandler) TranslationalImageRegistrationRequest {
	instance := TranslationalImageRegistrationRequestClass.Alloc().InitWithTargetedCIImageOptionsCompletionHandler(ciImage, options, completionHandler)
	instance.Autorelease()
	return instance
}

func (t_ TranslationalImageRegistrationRequest) InitWithTargetedCIImageOrientationOptions(ciImage coreimage.IImage, orientation imageio.ImagePropertyOrientation, options map[ImageOption]objc.IObject) TranslationalImageRegistrationRequest {
	rv := objc.Call[TranslationalImageRegistrationRequest](t_, objc.Sel("initWithTargetedCIImage:orientation:options:"), ciImage, orientation, options)
	return rv
}

// Creates a new request targeting a CIImage of known orientation. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/vision/vntargetedimagerequest/2923459-initwithtargetedciimage?language=objc
func NewTranslationalImageRegistrationRequestWithTargetedCIImageOrientationOptions(ciImage coreimage.IImage, orientation imageio.ImagePropertyOrientation, options map[ImageOption]objc.IObject) TranslationalImageRegistrationRequest {
	instance := TranslationalImageRegistrationRequestClass.Alloc().InitWithTargetedCIImageOrientationOptions(ciImage, orientation, options)
	instance.Autorelease()
	return instance
}

func (t_ TranslationalImageRegistrationRequest) InitWithTargetedImageURLOrientationOptions(imageURL foundation.IURL, orientation imageio.ImagePropertyOrientation, options map[ImageOption]objc.IObject) TranslationalImageRegistrationRequest {
	rv := objc.Call[TranslationalImageRegistrationRequest](t_, objc.Sel("initWithTargetedImageURL:orientation:options:"), imageURL, orientation, options)
	return rv
}

// Creates a new request targeting an image of known orientation, at the specified URL. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/vision/vntargetedimagerequest/2923456-initwithtargetedimageurl?language=objc
func NewTranslationalImageRegistrationRequestWithTargetedImageURLOrientationOptions(imageURL foundation.IURL, orientation imageio.ImagePropertyOrientation, options map[ImageOption]objc.IObject) TranslationalImageRegistrationRequest {
	instance := TranslationalImageRegistrationRequestClass.Alloc().InitWithTargetedImageURLOrientationOptions(imageURL, orientation, options)
	instance.Autorelease()
	return instance
}

func (t_ TranslationalImageRegistrationRequest) InitWithTargetedCIImageOrientationOptionsCompletionHandler(ciImage coreimage.IImage, orientation imageio.ImagePropertyOrientation, options map[ImageOption]objc.IObject, completionHandler RequestCompletionHandler) TranslationalImageRegistrationRequest {
	rv := objc.Call[TranslationalImageRegistrationRequest](t_, objc.Sel("initWithTargetedCIImage:orientation:options:completionHandler:"), ciImage, orientation, options, completionHandler)
	return rv
}

// Creates a new request targeting a CIImage of known orientation, executing the completion handler when done. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/vision/vntargetedimagerequest/2923451-initwithtargetedciimage?language=objc
func NewTranslationalImageRegistrationRequestWithTargetedCIImageOrientationOptionsCompletionHandler(ciImage coreimage.IImage, orientation imageio.ImagePropertyOrientation, options map[ImageOption]objc.IObject, completionHandler RequestCompletionHandler) TranslationalImageRegistrationRequest {
	instance := TranslationalImageRegistrationRequestClass.Alloc().InitWithTargetedCIImageOrientationOptionsCompletionHandler(ciImage, orientation, options, completionHandler)
	instance.Autorelease()
	return instance
}

func (t_ TranslationalImageRegistrationRequest) InitWithTargetedCGImageOptionsCompletionHandler(cgImage coregraphics.ImageRef, options map[ImageOption]objc.IObject, completionHandler RequestCompletionHandler) TranslationalImageRegistrationRequest {
	rv := objc.Call[TranslationalImageRegistrationRequest](t_, objc.Sel("initWithTargetedCGImage:options:completionHandler:"), cgImage, options, completionHandler)
	return rv
}

// Creates a new request targeting a Core Graphics image, executing the completion handler when done. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/vision/vntargetedimagerequest/2923448-initwithtargetedcgimage?language=objc
func NewTranslationalImageRegistrationRequestWithTargetedCGImageOptionsCompletionHandler(cgImage coregraphics.ImageRef, options map[ImageOption]objc.IObject, completionHandler RequestCompletionHandler) TranslationalImageRegistrationRequest {
	instance := TranslationalImageRegistrationRequestClass.Alloc().InitWithTargetedCGImageOptionsCompletionHandler(cgImage, options, completionHandler)
	instance.Autorelease()
	return instance
}

func (t_ TranslationalImageRegistrationRequest) InitWithTargetedImageDataOrientationOptions(imageData []byte, orientation imageio.ImagePropertyOrientation, options map[ImageOption]objc.IObject) TranslationalImageRegistrationRequest {
	rv := objc.Call[TranslationalImageRegistrationRequest](t_, objc.Sel("initWithTargetedImageData:orientation:options:"), imageData, orientation, options)
	return rv
}

// Creates a new request targeting a raw data image of known orientation. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/vision/vntargetedimagerequest/2923441-initwithtargetedimagedata?language=objc
func NewTranslationalImageRegistrationRequestWithTargetedImageDataOrientationOptions(imageData []byte, orientation imageio.ImagePropertyOrientation, options map[ImageOption]objc.IObject) TranslationalImageRegistrationRequest {
	instance := TranslationalImageRegistrationRequestClass.Alloc().InitWithTargetedImageDataOrientationOptions(imageData, orientation, options)
	instance.Autorelease()
	return instance
}

func (t_ TranslationalImageRegistrationRequest) InitWithTargetedCMSampleBufferOptions(sampleBuffer coremedia.SampleBufferRef, options map[ImageOption]objc.IObject) TranslationalImageRegistrationRequest {
	rv := objc.Call[TranslationalImageRegistrationRequest](t_, objc.Sel("initWithTargetedCMSampleBuffer:options:"), sampleBuffer, options)
	return rv
}

// Creates a new request that targets an image in a sample buffer. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/vision/vntargetedimagerequest/3571274-initwithtargetedcmsamplebuffer?language=objc
func NewTranslationalImageRegistrationRequestWithTargetedCMSampleBufferOptions(sampleBuffer coremedia.SampleBufferRef, options map[ImageOption]objc.IObject) TranslationalImageRegistrationRequest {
	instance := TranslationalImageRegistrationRequestClass.Alloc().InitWithTargetedCMSampleBufferOptions(sampleBuffer, options)
	instance.Autorelease()
	return instance
}

func (t_ TranslationalImageRegistrationRequest) InitWithTargetedImageURLOrientationOptionsCompletionHandler(imageURL foundation.IURL, orientation imageio.ImagePropertyOrientation, options map[ImageOption]objc.IObject, completionHandler RequestCompletionHandler) TranslationalImageRegistrationRequest {
	rv := objc.Call[TranslationalImageRegistrationRequest](t_, objc.Sel("initWithTargetedImageURL:orientation:options:completionHandler:"), imageURL, orientation, options, completionHandler)
	return rv
}

// Creates a new request targeting an image of known orientation, at the specified URL, executing the completion handler when done. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/vision/vntargetedimagerequest/2923457-initwithtargetedimageurl?language=objc
func NewTranslationalImageRegistrationRequestWithTargetedImageURLOrientationOptionsCompletionHandler(imageURL foundation.IURL, orientation imageio.ImagePropertyOrientation, options map[ImageOption]objc.IObject, completionHandler RequestCompletionHandler) TranslationalImageRegistrationRequest {
	instance := TranslationalImageRegistrationRequestClass.Alloc().InitWithTargetedImageURLOrientationOptionsCompletionHandler(imageURL, orientation, options, completionHandler)
	instance.Autorelease()
	return instance
}

func (t_ TranslationalImageRegistrationRequest) InitWithTargetedCMSampleBufferOrientationOptionsCompletionHandler(sampleBuffer coremedia.SampleBufferRef, orientation imageio.ImagePropertyOrientation, options map[ImageOption]objc.IObject, completionHandler RequestCompletionHandler) TranslationalImageRegistrationRequest {
	rv := objc.Call[TranslationalImageRegistrationRequest](t_, objc.Sel("initWithTargetedCMSampleBuffer:orientation:options:completionHandler:"), sampleBuffer, orientation, options, completionHandler)
	return rv
}

// Creates a new request with a completion handler that targets an image of a known orientation in a sample buffer. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/vision/vntargetedimagerequest/3571277-initwithtargetedcmsamplebuffer?language=objc
func NewTranslationalImageRegistrationRequestWithTargetedCMSampleBufferOrientationOptionsCompletionHandler(sampleBuffer coremedia.SampleBufferRef, orientation imageio.ImagePropertyOrientation, options map[ImageOption]objc.IObject, completionHandler RequestCompletionHandler) TranslationalImageRegistrationRequest {
	instance := TranslationalImageRegistrationRequestClass.Alloc().InitWithTargetedCMSampleBufferOrientationOptionsCompletionHandler(sampleBuffer, orientation, options, completionHandler)
	instance.Autorelease()
	return instance
}

func (t_ TranslationalImageRegistrationRequest) InitWithTargetedCIImageOptions(ciImage coreimage.IImage, options map[ImageOption]objc.IObject) TranslationalImageRegistrationRequest {
	rv := objc.Call[TranslationalImageRegistrationRequest](t_, objc.Sel("initWithTargetedCIImage:options:"), ciImage, options)
	return rv
}

// Creates a new request targeting a CIImage. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/vision/vntargetedimagerequest/2923447-initwithtargetedciimage?language=objc
func NewTranslationalImageRegistrationRequestWithTargetedCIImageOptions(ciImage coreimage.IImage, options map[ImageOption]objc.IObject) TranslationalImageRegistrationRequest {
	instance := TranslationalImageRegistrationRequestClass.Alloc().InitWithTargetedCIImageOptions(ciImage, options)
	instance.Autorelease()
	return instance
}

func (t_ TranslationalImageRegistrationRequest) InitWithTargetedCGImageOrientationOptionsCompletionHandler(cgImage coregraphics.ImageRef, orientation imageio.ImagePropertyOrientation, options map[ImageOption]objc.IObject, completionHandler RequestCompletionHandler) TranslationalImageRegistrationRequest {
	rv := objc.Call[TranslationalImageRegistrationRequest](t_, objc.Sel("initWithTargetedCGImage:orientation:options:completionHandler:"), cgImage, orientation, options, completionHandler)
	return rv
}

// Creates a new request targeting a Core Graphics image of known orientation, executing the completion handler when done. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/vision/vntargetedimagerequest/2923450-initwithtargetedcgimage?language=objc
func NewTranslationalImageRegistrationRequestWithTargetedCGImageOrientationOptionsCompletionHandler(cgImage coregraphics.ImageRef, orientation imageio.ImagePropertyOrientation, options map[ImageOption]objc.IObject, completionHandler RequestCompletionHandler) TranslationalImageRegistrationRequest {
	instance := TranslationalImageRegistrationRequestClass.Alloc().InitWithTargetedCGImageOrientationOptionsCompletionHandler(cgImage, orientation, options, completionHandler)
	instance.Autorelease()
	return instance
}

func (t_ TranslationalImageRegistrationRequest) InitWithTargetedCVPixelBufferOrientationOptions(pixelBuffer corevideo.PixelBufferRef, orientation imageio.ImagePropertyOrientation, options map[ImageOption]objc.IObject) TranslationalImageRegistrationRequest {
	rv := objc.Call[TranslationalImageRegistrationRequest](t_, objc.Sel("initWithTargetedCVPixelBuffer:orientation:options:"), pixelBuffer, orientation, options)
	return rv
}

// Creates a new request targeting an image in a CVPixelBufferRef of known orientation. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/vision/vntargetedimagerequest/2923442-initwithtargetedcvpixelbuffer?language=objc
func NewTranslationalImageRegistrationRequestWithTargetedCVPixelBufferOrientationOptions(pixelBuffer corevideo.PixelBufferRef, orientation imageio.ImagePropertyOrientation, options map[ImageOption]objc.IObject) TranslationalImageRegistrationRequest {
	instance := TranslationalImageRegistrationRequestClass.Alloc().InitWithTargetedCVPixelBufferOrientationOptions(pixelBuffer, orientation, options)
	instance.Autorelease()
	return instance
}

func (t_ TranslationalImageRegistrationRequest) InitWithTargetedCVPixelBufferOrientationOptionsCompletionHandler(pixelBuffer corevideo.PixelBufferRef, orientation imageio.ImagePropertyOrientation, options map[ImageOption]objc.IObject, completionHandler RequestCompletionHandler) TranslationalImageRegistrationRequest {
	rv := objc.Call[TranslationalImageRegistrationRequest](t_, objc.Sel("initWithTargetedCVPixelBuffer:orientation:options:completionHandler:"), pixelBuffer, orientation, options, completionHandler)
	return rv
}

// Creates a new request targeting an image in a CVPixelBufferRef of known orientation, executing the completion handler when done. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/vision/vntargetedimagerequest/2923449-initwithtargetedcvpixelbuffer?language=objc
func NewTranslationalImageRegistrationRequestWithTargetedCVPixelBufferOrientationOptionsCompletionHandler(pixelBuffer corevideo.PixelBufferRef, orientation imageio.ImagePropertyOrientation, options map[ImageOption]objc.IObject, completionHandler RequestCompletionHandler) TranslationalImageRegistrationRequest {
	instance := TranslationalImageRegistrationRequestClass.Alloc().InitWithTargetedCVPixelBufferOrientationOptionsCompletionHandler(pixelBuffer, orientation, options, completionHandler)
	instance.Autorelease()
	return instance
}

func (t_ TranslationalImageRegistrationRequest) InitWithTargetedImageURLOptions(imageURL foundation.IURL, options map[ImageOption]objc.IObject) TranslationalImageRegistrationRequest {
	rv := objc.Call[TranslationalImageRegistrationRequest](t_, objc.Sel("initWithTargetedImageURL:options:"), imageURL, options)
	return rv
}

// Creates a new request targeting an image at the specified URL. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/vision/vntargetedimagerequest/2923458-initwithtargetedimageurl?language=objc
func NewTranslationalImageRegistrationRequestWithTargetedImageURLOptions(imageURL foundation.IURL, options map[ImageOption]objc.IObject) TranslationalImageRegistrationRequest {
	instance := TranslationalImageRegistrationRequestClass.Alloc().InitWithTargetedImageURLOptions(imageURL, options)
	instance.Autorelease()
	return instance
}

func (t_ TranslationalImageRegistrationRequest) InitWithCompletionHandler(completionHandler RequestCompletionHandler) TranslationalImageRegistrationRequest {
	rv := objc.Call[TranslationalImageRegistrationRequest](t_, objc.Sel("initWithCompletionHandler:"), completionHandler)
	return rv
}

// Creates a new Vision request with an optional completion handler. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/vision/vnrequest/2875416-initwithcompletionhandler?language=objc
func NewTranslationalImageRegistrationRequestWithCompletionHandler(completionHandler RequestCompletionHandler) TranslationalImageRegistrationRequest {
	instance := TranslationalImageRegistrationRequestClass.Alloc().InitWithCompletionHandler(completionHandler)
	instance.Autorelease()
	return instance
}
