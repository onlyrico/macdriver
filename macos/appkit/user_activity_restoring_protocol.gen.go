// Code generated by DarwinKit. DO NOT EDIT.

package appkit

import (
	"github.com/progrium/darwinkit/macos/foundation"
	"github.com/progrium/darwinkit/objc"
)

// A protocol that marks classes to restore the state of your app to continue a user activity. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsuseractivityrestoring?language=objc
type PUserActivityRestoring interface {
	// optional
	RestoreUserActivityState(userActivity foundation.UserActivity)
	HasRestoreUserActivityState() bool
}

// ensure impl type implements protocol interface
var _ PUserActivityRestoring = (*UserActivityRestoringObject)(nil)

// A concrete type for the [PUserActivityRestoring] protocol.
type UserActivityRestoringObject struct {
	objc.Object
}

func (u_ UserActivityRestoringObject) HasRestoreUserActivityState() bool {
	return u_.RespondsToSelector(objc.Sel("restoreUserActivityState:"))
}

// Restores the state necessary to continue the specified user activity. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsuseractivityrestoring/3022485-restoreuseractivitystate?language=objc
func (u_ UserActivityRestoringObject) RestoreUserActivityState(userActivity foundation.UserActivity) {
	objc.Call[objc.Void](u_, objc.Sel("restoreUserActivityState:"), userActivity)
}
