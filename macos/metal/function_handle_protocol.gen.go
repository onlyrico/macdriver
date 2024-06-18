// Code generated by DarwinKit. DO NOT EDIT.

package metal

import (
	"github.com/progrium/darwinkit/objc"
)

// An object representing a function that you can add to a visible function table. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlfunctionhandle?language=objc
type PFunctionHandle interface {
	// optional
	FunctionType() FunctionType
	HasFunctionType() bool

	// optional
	Device() DeviceObject
	HasDevice() bool

	// optional
	Name() string
	HasName() bool
}

// ensure impl type implements protocol interface
var _ PFunctionHandle = (*FunctionHandleObject)(nil)

// A concrete type for the [PFunctionHandle] protocol.
type FunctionHandleObject struct {
	objc.Object
}

func (f_ FunctionHandleObject) HasFunctionType() bool {
	return f_.RespondsToSelector(objc.Sel("functionType"))
}

// The shader function’s type. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlfunctionhandle/3554003-functiontype?language=objc
func (f_ FunctionHandleObject) FunctionType() FunctionType {
	rv := objc.Call[FunctionType](f_, objc.Sel("functionType"))
	return rv
}

func (f_ FunctionHandleObject) HasDevice() bool {
	return f_.RespondsToSelector(objc.Sel("device"))
}

// The device object that created the shader function. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlfunctionhandle/3554002-device?language=objc
func (f_ FunctionHandleObject) Device() DeviceObject {
	rv := objc.Call[DeviceObject](f_, objc.Sel("device"))
	return rv
}

func (f_ FunctionHandleObject) HasName() bool {
	return f_.RespondsToSelector(objc.Sel("name"))
}

// The function’s name. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlfunctionhandle/3554004-name?language=objc
func (f_ FunctionHandleObject) Name() string {
	rv := objc.Call[string](f_, objc.Sel("name"))
	return rv
}
