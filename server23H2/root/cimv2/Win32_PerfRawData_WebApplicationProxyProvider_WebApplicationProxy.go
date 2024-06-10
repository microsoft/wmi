// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

// Author:
//
//	Auto Generated on 6/6/2024 using wmigen
//	Source root.CIMV2
//
// ////////////////////////////////////////////
package cimv2

import (
	"github.com/microsoft/wmi/pkg/base/query"
	"github.com/microsoft/wmi/pkg/errors"
	cim "github.com/microsoft/wmi/pkg/wmiinstance"
	"reflect"
)

// Win32_PerfRawData_WebApplicationProxyProvider_WebApplicationProxy struct
type Win32_PerfRawData_WebApplicationProxyProvider_WebApplicationProxy struct {
	*Win32_PerfRawData

	//
	ActiveRequests uint32

	//
	ADFSforActiveEndpointCacheSize uint32

	//
	BackendRequestsPerSec uint32

	//
	ExternalRequestsPerSec uint32

	//
	IntegratedWindowsAuthenticationCacheSize uint32
}

func NewWin32_PerfRawData_WebApplicationProxyProvider_WebApplicationProxyEx1(instance *cim.WmiInstance) (newInstance *Win32_PerfRawData_WebApplicationProxyProvider_WebApplicationProxy, err error) {
	tmp, err := NewWin32_PerfRawDataEx1(instance)

	if err != nil {
		return
	}
	newInstance = &Win32_PerfRawData_WebApplicationProxyProvider_WebApplicationProxy{
		Win32_PerfRawData: tmp,
	}
	return
}

func NewWin32_PerfRawData_WebApplicationProxyProvider_WebApplicationProxyEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *Win32_PerfRawData_WebApplicationProxyProvider_WebApplicationProxy, err error) {
	tmp, err := NewWin32_PerfRawDataEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &Win32_PerfRawData_WebApplicationProxyProvider_WebApplicationProxy{
		Win32_PerfRawData: tmp,
	}
	return
}

// SetActiveRequests sets the value of ActiveRequests for the instance
func (instance *Win32_PerfRawData_WebApplicationProxyProvider_WebApplicationProxy) SetPropertyActiveRequests(value uint32) (err error) {
	return instance.SetProperty("ActiveRequests", (value))
}

// GetActiveRequests gets the value of ActiveRequests for the instance
func (instance *Win32_PerfRawData_WebApplicationProxyProvider_WebApplicationProxy) GetPropertyActiveRequests() (value uint32, err error) {
	retValue, err := instance.GetProperty("ActiveRequests")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(uint32)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " uint32 is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = uint32(valuetmp)

	return
}

// SetADFSforActiveEndpointCacheSize sets the value of ADFSforActiveEndpointCacheSize for the instance
func (instance *Win32_PerfRawData_WebApplicationProxyProvider_WebApplicationProxy) SetPropertyADFSforActiveEndpointCacheSize(value uint32) (err error) {
	return instance.SetProperty("ADFSforActiveEndpointCacheSize", (value))
}

// GetADFSforActiveEndpointCacheSize gets the value of ADFSforActiveEndpointCacheSize for the instance
func (instance *Win32_PerfRawData_WebApplicationProxyProvider_WebApplicationProxy) GetPropertyADFSforActiveEndpointCacheSize() (value uint32, err error) {
	retValue, err := instance.GetProperty("ADFSforActiveEndpointCacheSize")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(uint32)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " uint32 is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = uint32(valuetmp)

	return
}

// SetBackendRequestsPerSec sets the value of BackendRequestsPerSec for the instance
func (instance *Win32_PerfRawData_WebApplicationProxyProvider_WebApplicationProxy) SetPropertyBackendRequestsPerSec(value uint32) (err error) {
	return instance.SetProperty("BackendRequestsPerSec", (value))
}

// GetBackendRequestsPerSec gets the value of BackendRequestsPerSec for the instance
func (instance *Win32_PerfRawData_WebApplicationProxyProvider_WebApplicationProxy) GetPropertyBackendRequestsPerSec() (value uint32, err error) {
	retValue, err := instance.GetProperty("BackendRequestsPerSec")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(uint32)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " uint32 is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = uint32(valuetmp)

	return
}

// SetExternalRequestsPerSec sets the value of ExternalRequestsPerSec for the instance
func (instance *Win32_PerfRawData_WebApplicationProxyProvider_WebApplicationProxy) SetPropertyExternalRequestsPerSec(value uint32) (err error) {
	return instance.SetProperty("ExternalRequestsPerSec", (value))
}

// GetExternalRequestsPerSec gets the value of ExternalRequestsPerSec for the instance
func (instance *Win32_PerfRawData_WebApplicationProxyProvider_WebApplicationProxy) GetPropertyExternalRequestsPerSec() (value uint32, err error) {
	retValue, err := instance.GetProperty("ExternalRequestsPerSec")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(uint32)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " uint32 is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = uint32(valuetmp)

	return
}

// SetIntegratedWindowsAuthenticationCacheSize sets the value of IntegratedWindowsAuthenticationCacheSize for the instance
func (instance *Win32_PerfRawData_WebApplicationProxyProvider_WebApplicationProxy) SetPropertyIntegratedWindowsAuthenticationCacheSize(value uint32) (err error) {
	return instance.SetProperty("IntegratedWindowsAuthenticationCacheSize", (value))
}

// GetIntegratedWindowsAuthenticationCacheSize gets the value of IntegratedWindowsAuthenticationCacheSize for the instance
func (instance *Win32_PerfRawData_WebApplicationProxyProvider_WebApplicationProxy) GetPropertyIntegratedWindowsAuthenticationCacheSize() (value uint32, err error) {
	retValue, err := instance.GetProperty("IntegratedWindowsAuthenticationCacheSize")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(uint32)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " uint32 is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = uint32(valuetmp)

	return
}
