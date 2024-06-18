// Code generated by DarwinKit. DO NOT EDIT.

package avfoundation

import (
	"unsafe"

	"github.com/progrium/darwinkit/macos/coremedia"
	"github.com/progrium/darwinkit/macos/foundation"
	"github.com/progrium/darwinkit/objc"
)

// The class instance for the [CompositionTrackSegment] class.
var CompositionTrackSegmentClass = _CompositionTrackSegmentClass{objc.GetClass("AVCompositionTrackSegment")}

type _CompositionTrackSegmentClass struct {
	objc.Class
}

// An interface definition for the [CompositionTrackSegment] class.
type ICompositionTrackSegment interface {
	IAssetTrackSegment
	SourceURL() foundation.URL
	SourceTrackID() objc.Object
}

// A track segment that maps a time from the source media track to the composition track. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcompositiontracksegment?language=objc
type CompositionTrackSegment struct {
	AssetTrackSegment
}

func CompositionTrackSegmentFrom(ptr unsafe.Pointer) CompositionTrackSegment {
	return CompositionTrackSegment{
		AssetTrackSegment: AssetTrackSegmentFrom(ptr),
	}
}

func (c_ CompositionTrackSegment) InitWithURLTrackIDSourceTimeRangeTargetTimeRange(URL foundation.IURL, trackID objc.IObject, sourceTimeRange coremedia.TimeRange, targetTimeRange coremedia.TimeRange) CompositionTrackSegment {
	rv := objc.Call[CompositionTrackSegment](c_, objc.Sel("initWithURL:trackID:sourceTimeRange:targetTimeRange:"), URL, trackID, sourceTimeRange, targetTimeRange)
	return rv
}

// Creates an object that presents a segment of a media file that the specified URL references. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcompositiontracksegment/1390282-initwithurl?language=objc
func NewCompositionTrackSegmentWithURLTrackIDSourceTimeRangeTargetTimeRange(URL foundation.IURL, trackID objc.IObject, sourceTimeRange coremedia.TimeRange, targetTimeRange coremedia.TimeRange) CompositionTrackSegment {
	instance := CompositionTrackSegmentClass.Alloc().InitWithURLTrackIDSourceTimeRangeTargetTimeRange(URL, trackID, sourceTimeRange, targetTimeRange)
	instance.Autorelease()
	return instance
}

func (cc _CompositionTrackSegmentClass) CompositionTrackSegmentWithTimeRange(timeRange coremedia.TimeRange) CompositionTrackSegment {
	rv := objc.Call[CompositionTrackSegment](cc, objc.Sel("compositionTrackSegmentWithTimeRange:"), timeRange)
	return rv
}

// Returns a new object that presents an empty composition track segment. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcompositiontracksegment/1400556-compositiontracksegmentwithtimer?language=objc
func CompositionTrackSegment_CompositionTrackSegmentWithTimeRange(timeRange coremedia.TimeRange) CompositionTrackSegment {
	return CompositionTrackSegmentClass.CompositionTrackSegmentWithTimeRange(timeRange)
}

func (cc _CompositionTrackSegmentClass) CompositionTrackSegmentWithURLTrackIDSourceTimeRangeTargetTimeRange(URL foundation.IURL, trackID objc.IObject, sourceTimeRange coremedia.TimeRange, targetTimeRange coremedia.TimeRange) CompositionTrackSegment {
	rv := objc.Call[CompositionTrackSegment](cc, objc.Sel("compositionTrackSegmentWithURL:trackID:sourceTimeRange:targetTimeRange:"), URL, trackID, sourceTimeRange, targetTimeRange)
	return rv
}

// Returns a new an object that presents a segment of a media file that the specified URL references. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcompositiontracksegment/1400552-compositiontracksegmentwithurl?language=objc
func CompositionTrackSegment_CompositionTrackSegmentWithURLTrackIDSourceTimeRangeTargetTimeRange(URL foundation.IURL, trackID objc.IObject, sourceTimeRange coremedia.TimeRange, targetTimeRange coremedia.TimeRange) CompositionTrackSegment {
	return CompositionTrackSegmentClass.CompositionTrackSegmentWithURLTrackIDSourceTimeRangeTargetTimeRange(URL, trackID, sourceTimeRange, targetTimeRange)
}

func (c_ CompositionTrackSegment) InitWithTimeRange(timeRange coremedia.TimeRange) CompositionTrackSegment {
	rv := objc.Call[CompositionTrackSegment](c_, objc.Sel("initWithTimeRange:"), timeRange)
	return rv
}

// Creates an object that presents an empty composition track segment. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcompositiontracksegment/1386841-initwithtimerange?language=objc
func NewCompositionTrackSegmentWithTimeRange(timeRange coremedia.TimeRange) CompositionTrackSegment {
	instance := CompositionTrackSegmentClass.Alloc().InitWithTimeRange(timeRange)
	instance.Autorelease()
	return instance
}

func (cc _CompositionTrackSegmentClass) Alloc() CompositionTrackSegment {
	rv := objc.Call[CompositionTrackSegment](cc, objc.Sel("alloc"))
	return rv
}

func (cc _CompositionTrackSegmentClass) New() CompositionTrackSegment {
	rv := objc.Call[CompositionTrackSegment](cc, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewCompositionTrackSegment() CompositionTrackSegment {
	return CompositionTrackSegmentClass.New()
}

func (c_ CompositionTrackSegment) Init() CompositionTrackSegment {
	rv := objc.Call[CompositionTrackSegment](c_, objc.Sel("init"))
	return rv
}

// A URL of the container file whose media this track segment presents. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcompositiontracksegment/1386814-sourceurl?language=objc
func (c_ CompositionTrackSegment) SourceURL() foundation.URL {
	rv := objc.Call[foundation.URL](c_, objc.Sel("sourceURL"))
	return rv
}

// An identifier of a track in the container file whose media this track segment presents. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcompositiontracksegment/1388326-sourcetrackid?language=objc
func (c_ CompositionTrackSegment) SourceTrackID() objc.Object {
	rv := objc.Call[objc.Object](c_, objc.Sel("sourceTrackID"))
	return rv
}
