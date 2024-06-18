// Code generated by DarwinKit. DO NOT EDIT.

package fileprovider

import (
	"github.com/progrium/darwinkit/macos/coregraphics"
	"github.com/progrium/darwinkit/macos/foundation"
	"github.com/progrium/darwinkit/objc"
)

// Support for item thumbnails. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/fileprovider/nsfileproviderthumbnailing?language=objc
type PFileProviderThumbnailing interface {
	// optional
	FetchThumbnailsForItemIdentifiersRequestedSizePerThumbnailCompletionHandlerCompletionHandler(itemIdentifiers []FileProviderItemIdentifier, size coregraphics.Size, perThumbnailCompletionHandler func(identifier FileProviderItemIdentifier, imageData []byte, error foundation.Error), completionHandler func(error foundation.Error)) foundation.Progress
	HasFetchThumbnailsForItemIdentifiersRequestedSizePerThumbnailCompletionHandlerCompletionHandler() bool
}

// ensure impl type implements protocol interface
var _ PFileProviderThumbnailing = (*FileProviderThumbnailingObject)(nil)

// A concrete type for the [PFileProviderThumbnailing] protocol.
type FileProviderThumbnailingObject struct {
	objc.Object
}

func (f_ FileProviderThumbnailingObject) HasFetchThumbnailsForItemIdentifiersRequestedSizePerThumbnailCompletionHandlerCompletionHandler() bool {
	return f_.RespondsToSelector(objc.Sel("fetchThumbnailsForItemIdentifiers:requestedSize:perThumbnailCompletionHandler:completionHandler:"))
}

// Asks the file provider for a thumbnail of the specified items. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/fileprovider/nsfileproviderthumbnailing/3553313-fetchthumbnailsforitemidentifier?language=objc
func (f_ FileProviderThumbnailingObject) FetchThumbnailsForItemIdentifiersRequestedSizePerThumbnailCompletionHandlerCompletionHandler(itemIdentifiers []FileProviderItemIdentifier, size coregraphics.Size, perThumbnailCompletionHandler func(identifier FileProviderItemIdentifier, imageData []byte, error foundation.Error), completionHandler func(error foundation.Error)) foundation.Progress {
	rv := objc.Call[foundation.Progress](f_, objc.Sel("fetchThumbnailsForItemIdentifiers:requestedSize:perThumbnailCompletionHandler:completionHandler:"), itemIdentifiers, size, perThumbnailCompletionHandler, completionHandler)
	return rv
}
