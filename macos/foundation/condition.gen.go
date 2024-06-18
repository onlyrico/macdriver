// Code generated by DarwinKit. DO NOT EDIT.

package foundation

import (
	"unsafe"

	"github.com/progrium/darwinkit/objc"
)

// The class instance for the [Condition] class.
var ConditionClass = _ConditionClass{objc.GetClass("NSCondition")}

type _ConditionClass struct {
	objc.Class
}

// An interface definition for the [Condition] class.
type ICondition interface {
	objc.IObject
	Broadcast()
	Signal()
	Wait()
	WaitUntilDate(limit IDate) bool
	Name() string
	SetName(value string)
}

// A condition variable whose semantics follow those used for POSIX-style conditions. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nscondition?language=objc
type Condition struct {
	objc.Object
}

func ConditionFrom(ptr unsafe.Pointer) Condition {
	return Condition{
		Object: objc.ObjectFrom(ptr),
	}
}

func (cc _ConditionClass) Alloc() Condition {
	rv := objc.Call[Condition](cc, objc.Sel("alloc"))
	return rv
}

func (cc _ConditionClass) New() Condition {
	rv := objc.Call[Condition](cc, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewCondition() Condition {
	return ConditionClass.New()
}

func (c_ Condition) Init() Condition {
	rv := objc.Call[Condition](c_, objc.Sel("init"))
	return rv
}

// Signals the condition, waking up all threads waiting on it. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nscondition/1415094-broadcast?language=objc
func (c_ Condition) Broadcast() {
	objc.Call[objc.Void](c_, objc.Sel("broadcast"))
}

// Signals the condition, waking up one thread waiting on it. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nscondition/1415724-signal?language=objc
func (c_ Condition) Signal() {
	objc.Call[objc.Void](c_, objc.Sel("signal"))
}

// Blocks the current thread until the condition is signaled. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nscondition/1407950-wait?language=objc
func (c_ Condition) Wait() {
	objc.Call[objc.Void](c_, objc.Sel("wait"))
}

// Blocks the current thread until the condition is signaled or the specified time limit is reached. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nscondition/1418307-waituntildate?language=objc
func (c_ Condition) WaitUntilDate(limit IDate) bool {
	rv := objc.Call[bool](c_, objc.Sel("waitUntilDate:"), limit)
	return rv
}

// The name of the condition. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nscondition/1408091-name?language=objc
func (c_ Condition) Name() string {
	rv := objc.Call[string](c_, objc.Sel("name"))
	return rv
}

// The name of the condition. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nscondition/1408091-name?language=objc
func (c_ Condition) SetName(value string) {
	objc.Call[objc.Void](c_, objc.Sel("setName:"), value)
}
