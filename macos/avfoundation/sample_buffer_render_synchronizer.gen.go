// Code generated by DarwinKit. DO NOT EDIT.

package avfoundation

import (
	"unsafe"

	"github.com/progrium/darwinkit/dispatch"
	"github.com/progrium/darwinkit/macos/coremedia"
	"github.com/progrium/darwinkit/macos/foundation"
	"github.com/progrium/darwinkit/objc"
)

// The class instance for the [SampleBufferRenderSynchronizer] class.
var SampleBufferRenderSynchronizerClass = _SampleBufferRenderSynchronizerClass{objc.GetClass("AVSampleBufferRenderSynchronizer")}

type _SampleBufferRenderSynchronizerClass struct {
	objc.Class
}

// An interface definition for the [SampleBufferRenderSynchronizer] class.
type ISampleBufferRenderSynchronizer interface {
	objc.IObject
	AddRenderer(renderer PQueuedSampleBufferRendering)
	AddRendererObject(rendererObject objc.IObject)
	SetRateTimeAtHostTime(rate float32, time coremedia.Time, hostTime coremedia.Time)
	RemoveRendererAtTimeCompletionHandler(renderer PQueuedSampleBufferRendering, time coremedia.Time, completionHandler func(didRemoveRenderer bool))
	RemoveRendererObjectAtTimeCompletionHandler(rendererObject objc.IObject, time coremedia.Time, completionHandler func(didRemoveRenderer bool))
	SetRateTime(rate float32, time coremedia.Time)
	AddBoundaryTimeObserverForTimesQueueUsingBlock(times []foundation.IValue, queue dispatch.Queue, block func()) objc.Object
	CurrentTime() coremedia.Time
	AddPeriodicTimeObserverForIntervalQueueUsingBlock(interval coremedia.Time, queue dispatch.Queue, block func(time coremedia.Time)) objc.Object
	RemoveTimeObserver(observer objc.IObject)
	Timebase() coremedia.TimebaseRef
	DelaysRateChangeUntilHasSufficientMediaData() bool
	SetDelaysRateChangeUntilHasSufficientMediaData(value bool)
	Rate() float32
	SetRate(value float32)
	Renderers() []QueuedSampleBufferRenderingObject
}

// An object used to synchronize multiple queued sample buffers to a single timeline. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avsamplebufferrendersynchronizer?language=objc
type SampleBufferRenderSynchronizer struct {
	objc.Object
}

func SampleBufferRenderSynchronizerFrom(ptr unsafe.Pointer) SampleBufferRenderSynchronizer {
	return SampleBufferRenderSynchronizer{
		Object: objc.ObjectFrom(ptr),
	}
}

func (sc _SampleBufferRenderSynchronizerClass) Alloc() SampleBufferRenderSynchronizer {
	rv := objc.Call[SampleBufferRenderSynchronizer](sc, objc.Sel("alloc"))
	return rv
}

func (sc _SampleBufferRenderSynchronizerClass) New() SampleBufferRenderSynchronizer {
	rv := objc.Call[SampleBufferRenderSynchronizer](sc, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewSampleBufferRenderSynchronizer() SampleBufferRenderSynchronizer {
	return SampleBufferRenderSynchronizerClass.New()
}

func (s_ SampleBufferRenderSynchronizer) Init() SampleBufferRenderSynchronizer {
	rv := objc.Call[SampleBufferRenderSynchronizer](s_, objc.Sel("init"))
	return rv
}

// Adds a renderer to the list of renderers under the synchronizer's control. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avsamplebufferrendersynchronizer/2867828-addrenderer?language=objc
func (s_ SampleBufferRenderSynchronizer) AddRenderer(renderer PQueuedSampleBufferRendering) {
	po0 := objc.WrapAsProtocol("AVQueuedSampleBufferRendering", renderer)
	objc.Call[objc.Void](s_, objc.Sel("addRenderer:"), po0)
}

// Adds a renderer to the list of renderers under the synchronizer's control. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avsamplebufferrendersynchronizer/2867828-addrenderer?language=objc
func (s_ SampleBufferRenderSynchronizer) AddRendererObject(rendererObject objc.IObject) {
	objc.Call[objc.Void](s_, objc.Sel("addRenderer:"), rendererObject)
}

// Sets the playback rate and the relationship between the current time and host time. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avsamplebufferrendersynchronizer/3726157-setrate?language=objc
func (s_ SampleBufferRenderSynchronizer) SetRateTimeAtHostTime(rate float32, time coremedia.Time, hostTime coremedia.Time) {
	objc.Call[objc.Void](s_, objc.Sel("setRate:time:atHostTime:"), rate, time, hostTime)
}

// Removes a renderer from the synchronizer. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avsamplebufferrendersynchronizer/2867826-removerenderer?language=objc
func (s_ SampleBufferRenderSynchronizer) RemoveRendererAtTimeCompletionHandler(renderer PQueuedSampleBufferRendering, time coremedia.Time, completionHandler func(didRemoveRenderer bool)) {
	po0 := objc.WrapAsProtocol("AVQueuedSampleBufferRendering", renderer)
	objc.Call[objc.Void](s_, objc.Sel("removeRenderer:atTime:completionHandler:"), po0, time, completionHandler)
}

// Removes a renderer from the synchronizer. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avsamplebufferrendersynchronizer/2867826-removerenderer?language=objc
func (s_ SampleBufferRenderSynchronizer) RemoveRendererObjectAtTimeCompletionHandler(rendererObject objc.IObject, time coremedia.Time, completionHandler func(didRemoveRenderer bool)) {
	objc.Call[objc.Void](s_, objc.Sel("removeRenderer:atTime:completionHandler:"), rendererObject, time, completionHandler)
}

// Sets the renderer’s time and rate. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avsamplebufferrendersynchronizer/2867821-setrate?language=objc
func (s_ SampleBufferRenderSynchronizer) SetRateTime(rate float32, time coremedia.Time) {
	objc.Call[objc.Void](s_, objc.Sel("setRate:time:"), rate, time)
}

// Requests invocation of a block when specified times are traversed during normal rendering. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avsamplebufferrendersynchronizer/2867824-addboundarytimeobserverfortimes?language=objc
func (s_ SampleBufferRenderSynchronizer) AddBoundaryTimeObserverForTimesQueueUsingBlock(times []foundation.IValue, queue dispatch.Queue, block func()) objc.Object {
	rv := objc.Call[objc.Object](s_, objc.Sel("addBoundaryTimeObserverForTimes:queue:usingBlock:"), times, queue, block)
	return rv
}

// Returns the current time of the synchronizer. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avsamplebufferrendersynchronizer/3022467-currenttime?language=objc
func (s_ SampleBufferRenderSynchronizer) CurrentTime() coremedia.Time {
	rv := objc.Call[coremedia.Time](s_, objc.Sel("currentTime"))
	return rv
}

// Requests invocation of a block during rendering at specified time intervals. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avsamplebufferrendersynchronizer/2867825-addperiodictimeobserverforinterv?language=objc
func (s_ SampleBufferRenderSynchronizer) AddPeriodicTimeObserverForIntervalQueueUsingBlock(interval coremedia.Time, queue dispatch.Queue, block func(time coremedia.Time)) objc.Object {
	rv := objc.Call[objc.Object](s_, objc.Sel("addPeriodicTimeObserverForInterval:queue:usingBlock:"), interval, queue, block)
	return rv
}

// Cancels the specified time observer. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avsamplebufferrendersynchronizer/2867829-removetimeobserver?language=objc
func (s_ SampleBufferRenderSynchronizer) RemoveTimeObserver(observer objc.IObject) {
	objc.Call[objc.Void](s_, objc.Sel("removeTimeObserver:"), observer)
}

// The synchronizer’s rendering timebase which determines how it interprets timestamps. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avsamplebufferrendersynchronizer/2867822-timebase?language=objc
func (s_ SampleBufferRenderSynchronizer) Timebase() coremedia.TimebaseRef {
	rv := objc.Call[coremedia.TimebaseRef](s_, objc.Sel("timebase"))
	return rv
}

// A Boolean value that Indicates whether the playback should start immediately on rate change requests. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avsamplebufferrendersynchronizer/3726156-delaysratechangeuntilhassufficie?language=objc
func (s_ SampleBufferRenderSynchronizer) DelaysRateChangeUntilHasSufficientMediaData() bool {
	rv := objc.Call[bool](s_, objc.Sel("delaysRateChangeUntilHasSufficientMediaData"))
	return rv
}

// A Boolean value that Indicates whether the playback should start immediately on rate change requests. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avsamplebufferrendersynchronizer/3726156-delaysratechangeuntilhassufficie?language=objc
func (s_ SampleBufferRenderSynchronizer) SetDelaysRateChangeUntilHasSufficientMediaData(value bool) {
	objc.Call[objc.Void](s_, objc.Sel("setDelaysRateChangeUntilHasSufficientMediaData:"), value)
}

// The current playback rate. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avsamplebufferrendersynchronizer/2867823-rate?language=objc
func (s_ SampleBufferRenderSynchronizer) Rate() float32 {
	rv := objc.Call[float32](s_, objc.Sel("rate"))
	return rv
}

// The current playback rate. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avsamplebufferrendersynchronizer/2867823-rate?language=objc
func (s_ SampleBufferRenderSynchronizer) SetRate(value float32) {
	objc.Call[objc.Void](s_, objc.Sel("setRate:"), value)
}

// An array of queued sample buffer renderers currently attached to the synchronizer. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avsamplebufferrendersynchronizer/2867827-renderers?language=objc
func (s_ SampleBufferRenderSynchronizer) Renderers() []QueuedSampleBufferRenderingObject {
	rv := objc.Call[[]QueuedSampleBufferRenderingObject](s_, objc.Sel("renderers"))
	return rv
}
