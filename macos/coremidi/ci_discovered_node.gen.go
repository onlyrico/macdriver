// Code generated by DarwinKit. DO NOT EDIT.

package coremidi

import (
	"unsafe"

	"github.com/progrium/darwinkit/macos/foundation"
	"github.com/progrium/darwinkit/objc"
)

// The class instance for the [CIDiscoveredNode] class.
var CIDiscoveredNodeClass = _CIDiscoveredNodeClass{objc.GetClass("MIDICIDiscoveredNode")}

type _CIDiscoveredNodeClass struct {
	objc.Class
}

// An interface definition for the [CIDiscoveredNode] class.
type ICIDiscoveredNode interface {
	objc.IObject
	SupportsProfiles() bool
	Destination() EntityRef
	SupportsProperties() bool
	MaximumSysExSize() foundation.Number
	DeviceInfo() CIDeviceInfo
}

// A discovered MIDI-CI node that represents a MIDI source and destination that respond to capability inquiries. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coremidi/midicidiscoverednode?language=objc
type CIDiscoveredNode struct {
	objc.Object
}

func CIDiscoveredNodeFrom(ptr unsafe.Pointer) CIDiscoveredNode {
	return CIDiscoveredNode{
		Object: objc.ObjectFrom(ptr),
	}
}

func (cc _CIDiscoveredNodeClass) Alloc() CIDiscoveredNode {
	rv := objc.Call[CIDiscoveredNode](cc, objc.Sel("alloc"))
	return rv
}

func (cc _CIDiscoveredNodeClass) New() CIDiscoveredNode {
	rv := objc.Call[CIDiscoveredNode](cc, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewCIDiscoveredNode() CIDiscoveredNode {
	return CIDiscoveredNodeClass.New()
}

func (c_ CIDiscoveredNode) Init() CIDiscoveredNode {
	rv := objc.Call[CIDiscoveredNode](c_, objc.Sel("init"))
	return rv
}

// A Boolean value that indicates whether this node supports MIDI-CI profiles. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coremidi/midicidiscoverednode/3580323-supportsprofiles?language=objc
func (c_ CIDiscoveredNode) SupportsProfiles() bool {
	rv := objc.Call[bool](c_, objc.Sel("supportsProfiles"))
	return rv
}

// The node’s MIDI destination. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coremidi/midicidiscoverednode/3580320-destination?language=objc
func (c_ CIDiscoveredNode) Destination() EntityRef {
	rv := objc.Call[EntityRef](c_, objc.Sel("destination"))
	return rv
}

// A Boolean value that indicates whether this node supports MIDI-CI properties. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coremidi/midicidiscoverednode/3580324-supportsproperties?language=objc
func (c_ CIDiscoveredNode) SupportsProperties() bool {
	rv := objc.Call[bool](c_, objc.Sel("supportsProperties"))
	return rv
}

// The maximum size of a System Exclusive (SysEx) message this node supports. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coremidi/midicidiscoverednode/3580322-maximumsysexsize?language=objc
func (c_ CIDiscoveredNode) MaximumSysExSize() foundation.Number {
	rv := objc.Call[foundation.Number](c_, objc.Sel("maximumSysExSize"))
	return rv
}

// The available MIDI-CI device information. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coremidi/midicidiscoverednode/3580321-deviceinfo?language=objc
func (c_ CIDiscoveredNode) DeviceInfo() CIDeviceInfo {
	rv := objc.Call[CIDeviceInfo](c_, objc.Sel("deviceInfo"))
	return rv
}
