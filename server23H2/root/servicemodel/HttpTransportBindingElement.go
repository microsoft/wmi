// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 6/14/2024 using wmigen
//      Source root.ServiceModel
//////////////////////////////////////////////
package servicemodel

import (
	"github.com/microsoft/wmi/pkg/base/query"
	"github.com/microsoft/wmi/pkg/errors"
	cim "github.com/microsoft/wmi/pkg/wmiinstance"
	"reflect"
)

// HttpTransportBindingElement struct
type HttpTransportBindingElement struct {
	*TransportBindingElement

	// A value that indicates whether the client accepts cookies and propagates them on future requests.
	AllowCookies bool

	// The authentication scheme used to authenticate client requests being processed by an HTTP listener.
	AuthenticationScheme string

	// A value that indicates whether proxies are ignored for local addresses.
	BypassProxyOnLocal bool

	// A value that indicates whether HTTP decompression (gzip/deflate) is enabled on the client.
	DecompressionEnabled bool

	// The extended protection policy used by the server to validate incoming client connections.
	ExtendedProtectionPolicy ExtendedProtectionPolicy

	// A value that indicates whether the hostname is used to reach the service when matching on the URI.
	HostNameComparisonMode string

	// When enabled http connections are kept alive regardless of activity level.
	KeepAliveEnabled bool

	// The maximum size of the buffer pool
	MaxBufferSize int32

	// A URI that contains the address of the proxy to use for HTTP requests.
	ProxyAddress string

	// The authentication scheme used to authenticate client requests being processed by an HTTP proxy.
	ProxyAuthenticationScheme string

	// The authentication realm.
	Realm string

	// A value that specifies whether messages are buffered or streamed or a request or response.
	TransferMode string

	// A value that indicates whether Unsafe Connection Sharing is enabled on the server.
	UnsafeConnectionNtlmAuthentication bool

	// A value that indicates whether the machine-wide proxy settings are used rather than the user specific settings.
	UseDefaultWebProxy bool
}

func NewHttpTransportBindingElementEx1(instance *cim.WmiInstance) (newInstance *HttpTransportBindingElement, err error) {
	tmp, err := NewTransportBindingElementEx1(instance)

	if err != nil {
		return
	}
	newInstance = &HttpTransportBindingElement{
		TransportBindingElement: tmp,
	}
	return
}

func NewHttpTransportBindingElementEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *HttpTransportBindingElement, err error) {
	tmp, err := NewTransportBindingElementEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &HttpTransportBindingElement{
		TransportBindingElement: tmp,
	}
	return
}

// SetAllowCookies sets the value of AllowCookies for the instance
func (instance *HttpTransportBindingElement) SetPropertyAllowCookies(value bool) (err error) {
	return instance.SetProperty("AllowCookies", (value))
}

// GetAllowCookies gets the value of AllowCookies for the instance
func (instance *HttpTransportBindingElement) GetPropertyAllowCookies() (value bool, err error) {
	retValue, err := instance.GetProperty("AllowCookies")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(bool)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " bool is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = bool(valuetmp)

	return
}

// SetAuthenticationScheme sets the value of AuthenticationScheme for the instance
func (instance *HttpTransportBindingElement) SetPropertyAuthenticationScheme(value string) (err error) {
	return instance.SetProperty("AuthenticationScheme", (value))
}

// GetAuthenticationScheme gets the value of AuthenticationScheme for the instance
func (instance *HttpTransportBindingElement) GetPropertyAuthenticationScheme() (value string, err error) {
	retValue, err := instance.GetProperty("AuthenticationScheme")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(string)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " string is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = string(valuetmp)

	return
}

// SetBypassProxyOnLocal sets the value of BypassProxyOnLocal for the instance
func (instance *HttpTransportBindingElement) SetPropertyBypassProxyOnLocal(value bool) (err error) {
	return instance.SetProperty("BypassProxyOnLocal", (value))
}

// GetBypassProxyOnLocal gets the value of BypassProxyOnLocal for the instance
func (instance *HttpTransportBindingElement) GetPropertyBypassProxyOnLocal() (value bool, err error) {
	retValue, err := instance.GetProperty("BypassProxyOnLocal")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(bool)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " bool is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = bool(valuetmp)

	return
}

// SetDecompressionEnabled sets the value of DecompressionEnabled for the instance
func (instance *HttpTransportBindingElement) SetPropertyDecompressionEnabled(value bool) (err error) {
	return instance.SetProperty("DecompressionEnabled", (value))
}

// GetDecompressionEnabled gets the value of DecompressionEnabled for the instance
func (instance *HttpTransportBindingElement) GetPropertyDecompressionEnabled() (value bool, err error) {
	retValue, err := instance.GetProperty("DecompressionEnabled")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(bool)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " bool is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = bool(valuetmp)

	return
}

// SetExtendedProtectionPolicy sets the value of ExtendedProtectionPolicy for the instance
func (instance *HttpTransportBindingElement) SetPropertyExtendedProtectionPolicy(value ExtendedProtectionPolicy) (err error) {
	return instance.SetProperty("ExtendedProtectionPolicy", (value))
}

// GetExtendedProtectionPolicy gets the value of ExtendedProtectionPolicy for the instance
func (instance *HttpTransportBindingElement) GetPropertyExtendedProtectionPolicy() (value ExtendedProtectionPolicy, err error) {
	retValue, err := instance.GetProperty("ExtendedProtectionPolicy")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(ExtendedProtectionPolicy)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " ExtendedProtectionPolicy is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = ExtendedProtectionPolicy(valuetmp)

	return
}

// SetHostNameComparisonMode sets the value of HostNameComparisonMode for the instance
func (instance *HttpTransportBindingElement) SetPropertyHostNameComparisonMode(value string) (err error) {
	return instance.SetProperty("HostNameComparisonMode", (value))
}

// GetHostNameComparisonMode gets the value of HostNameComparisonMode for the instance
func (instance *HttpTransportBindingElement) GetPropertyHostNameComparisonMode() (value string, err error) {
	retValue, err := instance.GetProperty("HostNameComparisonMode")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(string)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " string is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = string(valuetmp)

	return
}

// SetKeepAliveEnabled sets the value of KeepAliveEnabled for the instance
func (instance *HttpTransportBindingElement) SetPropertyKeepAliveEnabled(value bool) (err error) {
	return instance.SetProperty("KeepAliveEnabled", (value))
}

// GetKeepAliveEnabled gets the value of KeepAliveEnabled for the instance
func (instance *HttpTransportBindingElement) GetPropertyKeepAliveEnabled() (value bool, err error) {
	retValue, err := instance.GetProperty("KeepAliveEnabled")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(bool)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " bool is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = bool(valuetmp)

	return
}

// SetMaxBufferSize sets the value of MaxBufferSize for the instance
func (instance *HttpTransportBindingElement) SetPropertyMaxBufferSize(value int32) (err error) {
	return instance.SetProperty("MaxBufferSize", (value))
}

// GetMaxBufferSize gets the value of MaxBufferSize for the instance
func (instance *HttpTransportBindingElement) GetPropertyMaxBufferSize() (value int32, err error) {
	retValue, err := instance.GetProperty("MaxBufferSize")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(int32)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " int32 is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = int32(valuetmp)

	return
}

// SetProxyAddress sets the value of ProxyAddress for the instance
func (instance *HttpTransportBindingElement) SetPropertyProxyAddress(value string) (err error) {
	return instance.SetProperty("ProxyAddress", (value))
}

// GetProxyAddress gets the value of ProxyAddress for the instance
func (instance *HttpTransportBindingElement) GetPropertyProxyAddress() (value string, err error) {
	retValue, err := instance.GetProperty("ProxyAddress")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(string)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " string is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = string(valuetmp)

	return
}

// SetProxyAuthenticationScheme sets the value of ProxyAuthenticationScheme for the instance
func (instance *HttpTransportBindingElement) SetPropertyProxyAuthenticationScheme(value string) (err error) {
	return instance.SetProperty("ProxyAuthenticationScheme", (value))
}

// GetProxyAuthenticationScheme gets the value of ProxyAuthenticationScheme for the instance
func (instance *HttpTransportBindingElement) GetPropertyProxyAuthenticationScheme() (value string, err error) {
	retValue, err := instance.GetProperty("ProxyAuthenticationScheme")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(string)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " string is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = string(valuetmp)

	return
}

// SetRealm sets the value of Realm for the instance
func (instance *HttpTransportBindingElement) SetPropertyRealm(value string) (err error) {
	return instance.SetProperty("Realm", (value))
}

// GetRealm gets the value of Realm for the instance
func (instance *HttpTransportBindingElement) GetPropertyRealm() (value string, err error) {
	retValue, err := instance.GetProperty("Realm")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(string)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " string is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = string(valuetmp)

	return
}

// SetTransferMode sets the value of TransferMode for the instance
func (instance *HttpTransportBindingElement) SetPropertyTransferMode(value string) (err error) {
	return instance.SetProperty("TransferMode", (value))
}

// GetTransferMode gets the value of TransferMode for the instance
func (instance *HttpTransportBindingElement) GetPropertyTransferMode() (value string, err error) {
	retValue, err := instance.GetProperty("TransferMode")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(string)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " string is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = string(valuetmp)

	return
}

// SetUnsafeConnectionNtlmAuthentication sets the value of UnsafeConnectionNtlmAuthentication for the instance
func (instance *HttpTransportBindingElement) SetPropertyUnsafeConnectionNtlmAuthentication(value bool) (err error) {
	return instance.SetProperty("UnsafeConnectionNtlmAuthentication", (value))
}

// GetUnsafeConnectionNtlmAuthentication gets the value of UnsafeConnectionNtlmAuthentication for the instance
func (instance *HttpTransportBindingElement) GetPropertyUnsafeConnectionNtlmAuthentication() (value bool, err error) {
	retValue, err := instance.GetProperty("UnsafeConnectionNtlmAuthentication")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(bool)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " bool is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = bool(valuetmp)

	return
}

// SetUseDefaultWebProxy sets the value of UseDefaultWebProxy for the instance
func (instance *HttpTransportBindingElement) SetPropertyUseDefaultWebProxy(value bool) (err error) {
	return instance.SetProperty("UseDefaultWebProxy", (value))
}

// GetUseDefaultWebProxy gets the value of UseDefaultWebProxy for the instance
func (instance *HttpTransportBindingElement) GetPropertyUseDefaultWebProxy() (value bool, err error) {
	retValue, err := instance.GetProperty("UseDefaultWebProxy")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(bool)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " bool is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = bool(valuetmp)

	return
}
