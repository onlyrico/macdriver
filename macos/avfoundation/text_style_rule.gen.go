// Code generated by DarwinKit. DO NOT EDIT.

package avfoundation

import (
	"unsafe"

	"github.com/progrium/darwinkit/objc"
)

// The class instance for the [TextStyleRule] class.
var TextStyleRuleClass = _TextStyleRuleClass{objc.GetClass("AVTextStyleRule")}

type _TextStyleRuleClass struct {
	objc.Class
}

// An interface definition for the [TextStyleRule] class.
type ITextStyleRule interface {
	objc.IObject
	TextSelector() string
	TextMarkupAttributes() map[string]objc.Object
}

// An object that represents the text styling rules to apply to a media item’s textual content. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avtextstylerule?language=objc
type TextStyleRule struct {
	objc.Object
}

func TextStyleRuleFrom(ptr unsafe.Pointer) TextStyleRule {
	return TextStyleRule{
		Object: objc.ObjectFrom(ptr),
	}
}

func (t_ TextStyleRule) InitWithTextMarkupAttributes(textMarkupAttributes map[string]objc.IObject) TextStyleRule {
	rv := objc.Call[TextStyleRule](t_, objc.Sel("initWithTextMarkupAttributes:"), textMarkupAttributes)
	return rv
}

// Creates a text style rule object with the specified style attributes. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avtextstylerule/1385849-initwithtextmarkupattributes?language=objc
func NewTextStyleRuleWithTextMarkupAttributes(textMarkupAttributes map[string]objc.IObject) TextStyleRule {
	instance := TextStyleRuleClass.Alloc().InitWithTextMarkupAttributes(textMarkupAttributes)
	instance.Autorelease()
	return instance
}

func (t_ TextStyleRule) InitWithTextMarkupAttributesTextSelector(textMarkupAttributes map[string]objc.IObject, textSelector string) TextStyleRule {
	rv := objc.Call[TextStyleRule](t_, objc.Sel("initWithTextMarkupAttributes:textSelector:"), textMarkupAttributes, textSelector)
	return rv
}

// Creates a text style rule object with the specified style attributes and text range information. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avtextstylerule/1389854-initwithtextmarkupattributes?language=objc
func NewTextStyleRuleWithTextMarkupAttributesTextSelector(textMarkupAttributes map[string]objc.IObject, textSelector string) TextStyleRule {
	instance := TextStyleRuleClass.Alloc().InitWithTextMarkupAttributesTextSelector(textMarkupAttributes, textSelector)
	instance.Autorelease()
	return instance
}

func (tc _TextStyleRuleClass) Alloc() TextStyleRule {
	rv := objc.Call[TextStyleRule](tc, objc.Sel("alloc"))
	return rv
}

func (tc _TextStyleRuleClass) New() TextStyleRule {
	rv := objc.Call[TextStyleRule](tc, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewTextStyleRule() TextStyleRule {
	return TextStyleRuleClass.New()
}

func (t_ TextStyleRule) Init() TextStyleRule {
	rv := objc.Call[TextStyleRule](t_, objc.Sel("init"))
	return rv
}

// Creates an array of text style rule objects from the specified property-list object. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avtextstylerule/1387802-textstylerulesfrompropertylist?language=objc
func (tc _TextStyleRuleClass) TextStyleRulesFromPropertyList(plist objc.IObject) []TextStyleRule {
	rv := objc.Call[[]TextStyleRule](tc, objc.Sel("textStyleRulesFromPropertyList:"), plist)
	return rv
}

// Creates an array of text style rule objects from the specified property-list object. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avtextstylerule/1387802-textstylerulesfrompropertylist?language=objc
func TextStyleRule_TextStyleRulesFromPropertyList(plist objc.IObject) []TextStyleRule {
	return TextStyleRuleClass.TextStyleRulesFromPropertyList(plist)
}

// Creates a new text style rule object using the specified style attributes and text range information. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avtextstylerule/1584361-textstylerulewithtextmarkupattri?language=objc
func (tc _TextStyleRuleClass) TextStyleRuleWithTextMarkupAttributesTextSelector(textMarkupAttributes map[string]objc.IObject, textSelector string) TextStyleRule {
	rv := objc.Call[TextStyleRule](tc, objc.Sel("textStyleRuleWithTextMarkupAttributes:textSelector:"), textMarkupAttributes, textSelector)
	return rv
}

// Creates a new text style rule object using the specified style attributes and text range information. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avtextstylerule/1584361-textstylerulewithtextmarkupattri?language=objc
func TextStyleRule_TextStyleRuleWithTextMarkupAttributesTextSelector(textMarkupAttributes map[string]objc.IObject, textSelector string) TextStyleRule {
	return TextStyleRuleClass.TextStyleRuleWithTextMarkupAttributesTextSelector(textMarkupAttributes, textSelector)
}

// Creates a new text style rule object using the style attributes in the specified dictionary. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avtextstylerule/1584360-textstylerulewithtextmarkupattri?language=objc
func (tc _TextStyleRuleClass) TextStyleRuleWithTextMarkupAttributes(textMarkupAttributes map[string]objc.IObject) TextStyleRule {
	rv := objc.Call[TextStyleRule](tc, objc.Sel("textStyleRuleWithTextMarkupAttributes:"), textMarkupAttributes)
	return rv
}

// Creates a new text style rule object using the style attributes in the specified dictionary. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avtextstylerule/1584360-textstylerulewithtextmarkupattri?language=objc
func TextStyleRule_TextStyleRuleWithTextMarkupAttributes(textMarkupAttributes map[string]objc.IObject) TextStyleRule {
	return TextStyleRuleClass.TextStyleRuleWithTextMarkupAttributes(textMarkupAttributes)
}

// Converts one or more text style rules into a serializable property list object. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avtextstylerule/1387970-propertylistfortextstylerules?language=objc
func (tc _TextStyleRuleClass) PropertyListForTextStyleRules(textStyleRules []ITextStyleRule) objc.Object {
	rv := objc.Call[objc.Object](tc, objc.Sel("propertyListForTextStyleRules:"), textStyleRules)
	return rv
}

// Converts one or more text style rules into a serializable property list object. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avtextstylerule/1387970-propertylistfortextstylerules?language=objc
func TextStyleRule_PropertyListForTextStyleRules(textStyleRules []ITextStyleRule) objc.Object {
	return TextStyleRuleClass.PropertyListForTextStyleRules(textStyleRules)
}

// A string that identifies the text to which the attributes should apply. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avtextstylerule/1389451-textselector?language=objc
func (t_ TextStyleRule) TextSelector() string {
	rv := objc.Call[string](t_, objc.Sel("textSelector"))
	return rv
}

// A dictionary of text style attributes to apply to the text. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avtextstylerule/1387945-textmarkupattributes?language=objc
func (t_ TextStyleRule) TextMarkupAttributes() map[string]objc.Object {
	rv := objc.Call[map[string]objc.Object](t_, objc.Sel("textMarkupAttributes"))
	return rv
}
