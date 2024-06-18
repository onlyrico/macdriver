// Code generated by DarwinKit. DO NOT EDIT.

package appkit

import (
	"github.com/progrium/darwinkit/objc"
)

// A role-based protocol that declares the minimum interface necessary for an accessibility element to act as a slider. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsaccessibilityslider?language=objc
type PAccessibilitySlider interface {
	// optional
	AccessibilityPerformIncrement() bool
	HasAccessibilityPerformIncrement() bool

	// optional
	AccessibilityValue() objc.Object
	HasAccessibilityValue() bool

	// optional
	AccessibilityLabel() string
	HasAccessibilityLabel() bool

	// optional
	AccessibilityPerformDecrement() bool
	HasAccessibilityPerformDecrement() bool
}

// ensure impl type implements protocol interface
var _ PAccessibilitySlider = (*AccessibilitySliderObject)(nil)

// A concrete type for the [PAccessibilitySlider] protocol.
type AccessibilitySliderObject struct {
	objc.Object
}

func (a_ AccessibilitySliderObject) HasAccessibilityPerformIncrement() bool {
	return a_.RespondsToSelector(objc.Sel("accessibilityPerformIncrement"))
}

// Increments the slider’s value. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsaccessibilityslider/1528478-accessibilityperformincrement?language=objc
func (a_ AccessibilitySliderObject) AccessibilityPerformIncrement() bool {
	rv := objc.Call[bool](a_, objc.Sel("accessibilityPerformIncrement"))
	return rv
}

func (a_ AccessibilitySliderObject) HasAccessibilityValue() bool {
	return a_.RespondsToSelector(objc.Sel("accessibilityValue"))
}

// Returns the slider’s value. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsaccessibilityslider/1530335-accessibilityvalue?language=objc
func (a_ AccessibilitySliderObject) AccessibilityValue() objc.Object {
	rv := objc.Call[objc.Object](a_, objc.Sel("accessibilityValue"))
	return rv
}

func (a_ AccessibilitySliderObject) HasAccessibilityLabel() bool {
	return a_.RespondsToSelector(objc.Sel("accessibilityLabel"))
}

// Returns a short description of the slider. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsaccessibilityslider/1530176-accessibilitylabel?language=objc
func (a_ AccessibilitySliderObject) AccessibilityLabel() string {
	rv := objc.Call[string](a_, objc.Sel("accessibilityLabel"))
	return rv
}

func (a_ AccessibilitySliderObject) HasAccessibilityPerformDecrement() bool {
	return a_.RespondsToSelector(objc.Sel("accessibilityPerformDecrement"))
}

// Decrements the slider’s value. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsaccessibilityslider/1534967-accessibilityperformdecrement?language=objc
func (a_ AccessibilitySliderObject) AccessibilityPerformDecrement() bool {
	rv := objc.Call[bool](a_, objc.Sel("accessibilityPerformDecrement"))
	return rv
}
