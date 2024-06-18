// Code generated by DarwinKit. DO NOT EDIT.

package mps

import (
	"github.com/progrium/darwinkit/objc"
)

// The protocol that provides resource identification. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpshandle?language=objc
type PHandle interface {
	// optional
	Label() string
	HasLabel() bool
}

// ensure impl type implements protocol interface
var _ PHandle = (*HandleObject)(nil)

// A concrete type for the [PHandle] protocol.
type HandleObject struct {
	objc.Object
}

func (h_ HandleObject) HasLabel() bool {
	return h_.RespondsToSelector(objc.Sel("label"))
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpshandle/2866414-label?language=objc
func (h_ HandleObject) Label() string {
	rv := objc.Call[string](h_, objc.Sel("label"))
	return rv
}
