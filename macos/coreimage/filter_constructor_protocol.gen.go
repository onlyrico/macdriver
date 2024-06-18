// Code generated by DarwinKit. DO NOT EDIT.

package coreimage

import (
	"github.com/progrium/darwinkit/objc"
)

// A general interface for objects that produce CIFilter instances. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cifilterconstructor?language=objc
type PFilterConstructor interface {
	// optional
	FilterWithName(name string) Filter
	HasFilterWithName() bool
}

// ensure impl type implements protocol interface
var _ PFilterConstructor = (*FilterConstructorObject)(nil)

// A concrete type for the [PFilterConstructor] protocol.
type FilterConstructorObject struct {
	objc.Object
}

func (f_ FilterConstructorObject) HasFilterWithName() bool {
	return f_.RespondsToSelector(objc.Sel("filterWithName:"))
}

// Returns a filter object specified by name. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cifilterconstructor/1438018-filterwithname?language=objc
func (f_ FilterConstructorObject) FilterWithName(name string) Filter {
	rv := objc.Call[Filter](f_, objc.Sel("filterWithName:"), name)
	return rv
}
