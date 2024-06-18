// Code generated by DarwinKit. DO NOT EDIT.

package avfoundation

import (
	"unsafe"

	"github.com/progrium/darwinkit/macos/coremedia"
	"github.com/progrium/darwinkit/objc"
)

// The class instance for the [MutableAudioMixInputParameters] class.
var MutableAudioMixInputParametersClass = _MutableAudioMixInputParametersClass{objc.GetClass("AVMutableAudioMixInputParameters")}

type _MutableAudioMixInputParametersClass struct {
	objc.Class
}

// An interface definition for the [MutableAudioMixInputParameters] class.
type IMutableAudioMixInputParameters interface {
	IAudioMixInputParameters
	SetVolumeAtTime(volume float32, time coremedia.Time)
	SetVolumeRampFromStartVolumeToEndVolumeTimeRange(startVolume float32, endVolume float32, timeRange coremedia.TimeRange)
	SetAudioTimePitchAlgorithm(value AudioTimePitchAlgorithm)
	SetAudioTapProcessor(value objc.IObject)
	SetTrackID(value objc.IObject)
}

// The parameters you use when adding an audio track to a mix. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avmutableaudiomixinputparameters?language=objc
type MutableAudioMixInputParameters struct {
	AudioMixInputParameters
}

func MutableAudioMixInputParametersFrom(ptr unsafe.Pointer) MutableAudioMixInputParameters {
	return MutableAudioMixInputParameters{
		AudioMixInputParameters: AudioMixInputParametersFrom(ptr),
	}
}

func (mc _MutableAudioMixInputParametersClass) AudioMixInputParametersWithTrack(track IAssetTrack) MutableAudioMixInputParameters {
	rv := objc.Call[MutableAudioMixInputParameters](mc, objc.Sel("audioMixInputParametersWithTrack:"), track)
	return rv
}

// Creates a mutable input parameters object for a given track. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avmutableaudiomixinputparameters/1386858-audiomixinputparameterswithtrack?language=objc
func MutableAudioMixInputParameters_AudioMixInputParametersWithTrack(track IAssetTrack) MutableAudioMixInputParameters {
	return MutableAudioMixInputParametersClass.AudioMixInputParametersWithTrack(track)
}

func (mc _MutableAudioMixInputParametersClass) AudioMixInputParameters() MutableAudioMixInputParameters {
	rv := objc.Call[MutableAudioMixInputParameters](mc, objc.Sel("audioMixInputParameters"))
	return rv
}

// Creates a mutable input parameters object. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avmutableaudiomixinputparameters/1560974-audiomixinputparameters?language=objc
func MutableAudioMixInputParameters_AudioMixInputParameters() MutableAudioMixInputParameters {
	return MutableAudioMixInputParametersClass.AudioMixInputParameters()
}

func (mc _MutableAudioMixInputParametersClass) Alloc() MutableAudioMixInputParameters {
	rv := objc.Call[MutableAudioMixInputParameters](mc, objc.Sel("alloc"))
	return rv
}

func (mc _MutableAudioMixInputParametersClass) New() MutableAudioMixInputParameters {
	rv := objc.Call[MutableAudioMixInputParameters](mc, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewMutableAudioMixInputParameters() MutableAudioMixInputParameters {
	return MutableAudioMixInputParametersClass.New()
}

func (m_ MutableAudioMixInputParameters) Init() MutableAudioMixInputParameters {
	rv := objc.Call[MutableAudioMixInputParameters](m_, objc.Sel("init"))
	return rv
}

// Sets the value of the audio volume starting at the specified time. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avmutableaudiomixinputparameters/1389875-setvolume?language=objc
func (m_ MutableAudioMixInputParameters) SetVolumeAtTime(volume float32, time coremedia.Time) {
	objc.Call[objc.Void](m_, objc.Sel("setVolume:atTime:"), volume, time)
}

// Sets a volume ramp to apply during a specified time range. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avmutableaudiomixinputparameters/1386056-setvolumerampfromstartvolume?language=objc
func (m_ MutableAudioMixInputParameters) SetVolumeRampFromStartVolumeToEndVolumeTimeRange(startVolume float32, endVolume float32, timeRange coremedia.TimeRange) {
	objc.Call[objc.Void](m_, objc.Sel("setVolumeRampFromStartVolume:toEndVolume:timeRange:"), startVolume, endVolume, timeRange)
}

// The processing algorithm used to manage audio pitch for scaled audio edits. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avmutableaudiomixinputparameters/1388300-audiotimepitchalgorithm?language=objc
func (m_ MutableAudioMixInputParameters) SetAudioTimePitchAlgorithm(value AudioTimePitchAlgorithm) {
	objc.Call[objc.Void](m_, objc.Sel("setAudioTimePitchAlgorithm:"), value)
}

// The audio processing tap associated with the track. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avmutableaudiomixinputparameters/1389296-audiotapprocessor?language=objc
func (m_ MutableAudioMixInputParameters) SetAudioTapProcessor(value objc.IObject) {
	objc.Call[objc.Void](m_, objc.Sel("setAudioTapProcessor:"), value)
}

// The identifier of the audio track to which the parameters should be applied. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avmutableaudiomixinputparameters/1389209-trackid?language=objc
func (m_ MutableAudioMixInputParameters) SetTrackID(value objc.IObject) {
	objc.Call[objc.Void](m_, objc.Sel("setTrackID:"), value)
}
