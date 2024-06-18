// Code generated by DarwinKit. DO NOT EDIT.

package vision

import (
	"unsafe"

	"github.com/progrium/darwinkit/objc"
)

// The class instance for the [TrackRectangleRequest] class.
var TrackRectangleRequestClass = _TrackRectangleRequestClass{objc.GetClass("VNTrackRectangleRequest")}

type _TrackRectangleRequestClass struct {
	objc.Class
}

// An interface definition for the [TrackRectangleRequest] class.
type ITrackRectangleRequest interface {
	ITrackingRequest
}

// An image analysis request that tracks movement of a previously identified rectangular object across multiple images or video frames. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/vision/vntrackrectanglerequest?language=objc
type TrackRectangleRequest struct {
	TrackingRequest
}

func TrackRectangleRequestFrom(ptr unsafe.Pointer) TrackRectangleRequest {
	return TrackRectangleRequest{
		TrackingRequest: TrackingRequestFrom(ptr),
	}
}

func (t_ TrackRectangleRequest) InitWithRectangleObservation(observation IRectangleObservation) TrackRectangleRequest {
	rv := objc.Call[TrackRectangleRequest](t_, objc.Sel("initWithRectangleObservation:"), observation)
	return rv
}

// Creates a new rectangle tracking request with a rectangle observation. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/vision/vntrackrectanglerequest/2887361-initwithrectangleobservation?language=objc
func NewTrackRectangleRequestWithRectangleObservation(observation IRectangleObservation) TrackRectangleRequest {
	instance := TrackRectangleRequestClass.Alloc().InitWithRectangleObservation(observation)
	instance.Autorelease()
	return instance
}

func (t_ TrackRectangleRequest) InitWithRectangleObservationCompletionHandler(observation IRectangleObservation, completionHandler RequestCompletionHandler) TrackRectangleRequest {
	rv := objc.Call[TrackRectangleRequest](t_, objc.Sel("initWithRectangleObservation:completionHandler:"), observation, completionHandler)
	return rv
}

// Creates a new rectangle tracking request with a rectangle observation. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/vision/vntrackrectanglerequest/2887357-initwithrectangleobservation?language=objc
func NewTrackRectangleRequestWithRectangleObservationCompletionHandler(observation IRectangleObservation, completionHandler RequestCompletionHandler) TrackRectangleRequest {
	instance := TrackRectangleRequestClass.Alloc().InitWithRectangleObservationCompletionHandler(observation, completionHandler)
	instance.Autorelease()
	return instance
}

func (tc _TrackRectangleRequestClass) Alloc() TrackRectangleRequest {
	rv := objc.Call[TrackRectangleRequest](tc, objc.Sel("alloc"))
	return rv
}

func (tc _TrackRectangleRequestClass) New() TrackRectangleRequest {
	rv := objc.Call[TrackRectangleRequest](tc, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewTrackRectangleRequest() TrackRectangleRequest {
	return TrackRectangleRequestClass.New()
}

func (t_ TrackRectangleRequest) Init() TrackRectangleRequest {
	rv := objc.Call[TrackRectangleRequest](t_, objc.Sel("init"))
	return rv
}

func (t_ TrackRectangleRequest) InitWithCompletionHandler(completionHandler RequestCompletionHandler) TrackRectangleRequest {
	rv := objc.Call[TrackRectangleRequest](t_, objc.Sel("initWithCompletionHandler:"), completionHandler)
	return rv
}

// Creates a new Vision request with an optional completion handler. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/vision/vnrequest/2875416-initwithcompletionhandler?language=objc
func NewTrackRectangleRequestWithCompletionHandler(completionHandler RequestCompletionHandler) TrackRectangleRequest {
	instance := TrackRectangleRequestClass.Alloc().InitWithCompletionHandler(completionHandler)
	instance.Autorelease()
	return instance
}
