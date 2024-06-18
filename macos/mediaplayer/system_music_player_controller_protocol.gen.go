// Code generated by DarwinKit. DO NOT EDIT.

package mediaplayer

import (
	"github.com/progrium/darwinkit/objc"
)

// A protocol for playing videos in the Music app. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/mediaplayer/mpsystemmusicplayercontroller?language=objc
type PSystemMusicPlayerController interface {
}

// ensure impl type implements protocol interface
var _ PSystemMusicPlayerController = (*SystemMusicPlayerControllerObject)(nil)

// A concrete type for the [PSystemMusicPlayerController] protocol.
type SystemMusicPlayerControllerObject struct {
	objc.Object
}
