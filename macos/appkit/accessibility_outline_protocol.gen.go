// Code generated by DarwinKit. DO NOT EDIT.

package appkit

import (
	"github.com/progrium/darwinkit/objc"
)

// A role-based protocol that declares the minimum interface necessary for an accessibility element to act as an outline view. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsaccessibilityoutline?language=objc
type PAccessibilityOutline interface {
}

// ensure impl type implements protocol interface
var _ PAccessibilityOutline = (*AccessibilityOutlineObject)(nil)

// A concrete type for the [PAccessibilityOutline] protocol.
type AccessibilityOutlineObject struct {
	objc.Object
}
