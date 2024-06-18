// Code generated by DarwinKit. DO NOT EDIT.

package webkit

import (
	"github.com/progrium/darwinkit/objc"
)

// WebDocumentText is an optional protocol for document view objects that display text. This protocol defines methods for accessing document content as strings, and methods for text selection. Classes that adopt this protocol should also adopt WebDocumentView and inherit from NSView. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/webkit/webdocumenttext?language=objc
type PWebDocumentText interface {
}

// ensure impl type implements protocol interface
var _ PWebDocumentText = (*WebDocumentTextObject)(nil)

// A concrete type for the [PWebDocumentText] protocol.
type WebDocumentTextObject struct {
	objc.Object
}
