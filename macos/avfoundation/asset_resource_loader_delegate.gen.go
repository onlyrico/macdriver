// Code generated by DarwinKit. DO NOT EDIT.

package avfoundation

import (
	"github.com/progrium/darwinkit/macos/foundation"
	"github.com/progrium/darwinkit/objc"
)

// Methods you can implement to handle resource-loading requests coming from a URL asset. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avassetresourceloaderdelegate?language=objc
type PAssetResourceLoaderDelegate interface {
	// optional
	ResourceLoaderShouldWaitForRenewalOfRequestedResource(resourceLoader AssetResourceLoader, renewalRequest AssetResourceRenewalRequest) bool
	HasResourceLoaderShouldWaitForRenewalOfRequestedResource() bool

	// optional
	ResourceLoaderDidCancelLoadingRequest(resourceLoader AssetResourceLoader, loadingRequest AssetResourceLoadingRequest)
	HasResourceLoaderDidCancelLoadingRequest() bool

	// optional
	ResourceLoaderDidCancelAuthenticationChallenge(resourceLoader AssetResourceLoader, authenticationChallenge foundation.URLAuthenticationChallenge)
	HasResourceLoaderDidCancelAuthenticationChallenge() bool

	// optional
	ResourceLoaderShouldWaitForResponseToAuthenticationChallenge(resourceLoader AssetResourceLoader, authenticationChallenge foundation.URLAuthenticationChallenge) bool
	HasResourceLoaderShouldWaitForResponseToAuthenticationChallenge() bool

	// optional
	ResourceLoaderShouldWaitForLoadingOfRequestedResource(resourceLoader AssetResourceLoader, loadingRequest AssetResourceLoadingRequest) bool
	HasResourceLoaderShouldWaitForLoadingOfRequestedResource() bool
}

// A delegate implementation builder for the [PAssetResourceLoaderDelegate] protocol.
type AssetResourceLoaderDelegate struct {
	_ResourceLoaderShouldWaitForRenewalOfRequestedResource        func(resourceLoader AssetResourceLoader, renewalRequest AssetResourceRenewalRequest) bool
	_ResourceLoaderDidCancelLoadingRequest                        func(resourceLoader AssetResourceLoader, loadingRequest AssetResourceLoadingRequest)
	_ResourceLoaderDidCancelAuthenticationChallenge               func(resourceLoader AssetResourceLoader, authenticationChallenge foundation.URLAuthenticationChallenge)
	_ResourceLoaderShouldWaitForResponseToAuthenticationChallenge func(resourceLoader AssetResourceLoader, authenticationChallenge foundation.URLAuthenticationChallenge) bool
	_ResourceLoaderShouldWaitForLoadingOfRequestedResource        func(resourceLoader AssetResourceLoader, loadingRequest AssetResourceLoadingRequest) bool
}

func (di *AssetResourceLoaderDelegate) HasResourceLoaderShouldWaitForRenewalOfRequestedResource() bool {
	return di._ResourceLoaderShouldWaitForRenewalOfRequestedResource != nil
}

// Tells the delegate when assistance is required of the application to renew a resource. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avassetresourceloaderdelegate/1387058-resourceloader?language=objc
func (di *AssetResourceLoaderDelegate) SetResourceLoaderShouldWaitForRenewalOfRequestedResource(f func(resourceLoader AssetResourceLoader, renewalRequest AssetResourceRenewalRequest) bool) {
	di._ResourceLoaderShouldWaitForRenewalOfRequestedResource = f
}

// Tells the delegate when assistance is required of the application to renew a resource. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avassetresourceloaderdelegate/1387058-resourceloader?language=objc
func (di *AssetResourceLoaderDelegate) ResourceLoaderShouldWaitForRenewalOfRequestedResource(resourceLoader AssetResourceLoader, renewalRequest AssetResourceRenewalRequest) bool {
	return di._ResourceLoaderShouldWaitForRenewalOfRequestedResource(resourceLoader, renewalRequest)
}
func (di *AssetResourceLoaderDelegate) HasResourceLoaderDidCancelLoadingRequest() bool {
	return di._ResourceLoaderDidCancelLoadingRequest != nil
}

// Informs the delegate that a prior loading request has been cancelled. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avassetresourceloaderdelegate/1387722-resourceloader?language=objc
func (di *AssetResourceLoaderDelegate) SetResourceLoaderDidCancelLoadingRequest(f func(resourceLoader AssetResourceLoader, loadingRequest AssetResourceLoadingRequest)) {
	di._ResourceLoaderDidCancelLoadingRequest = f
}

// Informs the delegate that a prior loading request has been cancelled. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avassetresourceloaderdelegate/1387722-resourceloader?language=objc
func (di *AssetResourceLoaderDelegate) ResourceLoaderDidCancelLoadingRequest(resourceLoader AssetResourceLoader, loadingRequest AssetResourceLoadingRequest) {
	di._ResourceLoaderDidCancelLoadingRequest(resourceLoader, loadingRequest)
}
func (di *AssetResourceLoaderDelegate) HasResourceLoaderDidCancelAuthenticationChallenge() bool {
	return di._ResourceLoaderDidCancelAuthenticationChallenge != nil
}

// Informs the delegate that a prior authentication challenge has been cancelled. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avassetresourceloaderdelegate/1387929-resourceloader?language=objc
func (di *AssetResourceLoaderDelegate) SetResourceLoaderDidCancelAuthenticationChallenge(f func(resourceLoader AssetResourceLoader, authenticationChallenge foundation.URLAuthenticationChallenge)) {
	di._ResourceLoaderDidCancelAuthenticationChallenge = f
}

// Informs the delegate that a prior authentication challenge has been cancelled. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avassetresourceloaderdelegate/1387929-resourceloader?language=objc
func (di *AssetResourceLoaderDelegate) ResourceLoaderDidCancelAuthenticationChallenge(resourceLoader AssetResourceLoader, authenticationChallenge foundation.URLAuthenticationChallenge) {
	di._ResourceLoaderDidCancelAuthenticationChallenge(resourceLoader, authenticationChallenge)
}
func (di *AssetResourceLoaderDelegate) HasResourceLoaderShouldWaitForResponseToAuthenticationChallenge() bool {
	return di._ResourceLoaderShouldWaitForResponseToAuthenticationChallenge != nil
}

// Tells the delegate that assistance is required of the application to respond to an authentication challenge. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avassetresourceloaderdelegate/1388736-resourceloader?language=objc
func (di *AssetResourceLoaderDelegate) SetResourceLoaderShouldWaitForResponseToAuthenticationChallenge(f func(resourceLoader AssetResourceLoader, authenticationChallenge foundation.URLAuthenticationChallenge) bool) {
	di._ResourceLoaderShouldWaitForResponseToAuthenticationChallenge = f
}

// Tells the delegate that assistance is required of the application to respond to an authentication challenge. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avassetresourceloaderdelegate/1388736-resourceloader?language=objc
func (di *AssetResourceLoaderDelegate) ResourceLoaderShouldWaitForResponseToAuthenticationChallenge(resourceLoader AssetResourceLoader, authenticationChallenge foundation.URLAuthenticationChallenge) bool {
	return di._ResourceLoaderShouldWaitForResponseToAuthenticationChallenge(resourceLoader, authenticationChallenge)
}
func (di *AssetResourceLoaderDelegate) HasResourceLoaderShouldWaitForLoadingOfRequestedResource() bool {
	return di._ResourceLoaderShouldWaitForLoadingOfRequestedResource != nil
}

// Asks the delegate if it wants to load the requested resource. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avassetresourceloaderdelegate/1388121-resourceloader?language=objc
func (di *AssetResourceLoaderDelegate) SetResourceLoaderShouldWaitForLoadingOfRequestedResource(f func(resourceLoader AssetResourceLoader, loadingRequest AssetResourceLoadingRequest) bool) {
	di._ResourceLoaderShouldWaitForLoadingOfRequestedResource = f
}

// Asks the delegate if it wants to load the requested resource. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avassetresourceloaderdelegate/1388121-resourceloader?language=objc
func (di *AssetResourceLoaderDelegate) ResourceLoaderShouldWaitForLoadingOfRequestedResource(resourceLoader AssetResourceLoader, loadingRequest AssetResourceLoadingRequest) bool {
	return di._ResourceLoaderShouldWaitForLoadingOfRequestedResource(resourceLoader, loadingRequest)
}

// ensure impl type implements protocol interface
var _ PAssetResourceLoaderDelegate = (*AssetResourceLoaderDelegateObject)(nil)

// A concrete type for the [PAssetResourceLoaderDelegate] protocol.
type AssetResourceLoaderDelegateObject struct {
	objc.Object
}

func (a_ AssetResourceLoaderDelegateObject) HasResourceLoaderShouldWaitForRenewalOfRequestedResource() bool {
	return a_.RespondsToSelector(objc.Sel("resourceLoader:shouldWaitForRenewalOfRequestedResource:"))
}

// Tells the delegate when assistance is required of the application to renew a resource. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avassetresourceloaderdelegate/1387058-resourceloader?language=objc
func (a_ AssetResourceLoaderDelegateObject) ResourceLoaderShouldWaitForRenewalOfRequestedResource(resourceLoader AssetResourceLoader, renewalRequest AssetResourceRenewalRequest) bool {
	rv := objc.Call[bool](a_, objc.Sel("resourceLoader:shouldWaitForRenewalOfRequestedResource:"), resourceLoader, renewalRequest)
	return rv
}

func (a_ AssetResourceLoaderDelegateObject) HasResourceLoaderDidCancelLoadingRequest() bool {
	return a_.RespondsToSelector(objc.Sel("resourceLoader:didCancelLoadingRequest:"))
}

// Informs the delegate that a prior loading request has been cancelled. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avassetresourceloaderdelegate/1387722-resourceloader?language=objc
func (a_ AssetResourceLoaderDelegateObject) ResourceLoaderDidCancelLoadingRequest(resourceLoader AssetResourceLoader, loadingRequest AssetResourceLoadingRequest) {
	objc.Call[objc.Void](a_, objc.Sel("resourceLoader:didCancelLoadingRequest:"), resourceLoader, loadingRequest)
}

func (a_ AssetResourceLoaderDelegateObject) HasResourceLoaderDidCancelAuthenticationChallenge() bool {
	return a_.RespondsToSelector(objc.Sel("resourceLoader:didCancelAuthenticationChallenge:"))
}

// Informs the delegate that a prior authentication challenge has been cancelled. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avassetresourceloaderdelegate/1387929-resourceloader?language=objc
func (a_ AssetResourceLoaderDelegateObject) ResourceLoaderDidCancelAuthenticationChallenge(resourceLoader AssetResourceLoader, authenticationChallenge foundation.URLAuthenticationChallenge) {
	objc.Call[objc.Void](a_, objc.Sel("resourceLoader:didCancelAuthenticationChallenge:"), resourceLoader, authenticationChallenge)
}

func (a_ AssetResourceLoaderDelegateObject) HasResourceLoaderShouldWaitForResponseToAuthenticationChallenge() bool {
	return a_.RespondsToSelector(objc.Sel("resourceLoader:shouldWaitForResponseToAuthenticationChallenge:"))
}

// Tells the delegate that assistance is required of the application to respond to an authentication challenge. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avassetresourceloaderdelegate/1388736-resourceloader?language=objc
func (a_ AssetResourceLoaderDelegateObject) ResourceLoaderShouldWaitForResponseToAuthenticationChallenge(resourceLoader AssetResourceLoader, authenticationChallenge foundation.URLAuthenticationChallenge) bool {
	rv := objc.Call[bool](a_, objc.Sel("resourceLoader:shouldWaitForResponseToAuthenticationChallenge:"), resourceLoader, authenticationChallenge)
	return rv
}

func (a_ AssetResourceLoaderDelegateObject) HasResourceLoaderShouldWaitForLoadingOfRequestedResource() bool {
	return a_.RespondsToSelector(objc.Sel("resourceLoader:shouldWaitForLoadingOfRequestedResource:"))
}

// Asks the delegate if it wants to load the requested resource. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avassetresourceloaderdelegate/1388121-resourceloader?language=objc
func (a_ AssetResourceLoaderDelegateObject) ResourceLoaderShouldWaitForLoadingOfRequestedResource(resourceLoader AssetResourceLoader, loadingRequest AssetResourceLoadingRequest) bool {
	rv := objc.Call[bool](a_, objc.Sel("resourceLoader:shouldWaitForLoadingOfRequestedResource:"), resourceLoader, loadingRequest)
	return rv
}
