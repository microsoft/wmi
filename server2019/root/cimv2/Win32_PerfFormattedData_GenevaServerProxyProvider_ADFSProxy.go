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

// Win32_PerfFormattedData_GenevaServerProxyProvider_ADFSProxy struct
type Win32_PerfFormattedData_GenevaServerProxyProvider_ADFSProxy struct {
	*Win32_PerfFormattedData

	//
	OutstandingRequests uint64

	//
	RejectedRequests uint64

	//
	RejectedRequestsPersec uint64

	//
	RequestLatency uint64

	//
	Requests uint64

	//
	RequestsPersec uint64
}

func NewWin32_PerfFormattedData_GenevaServerProxyProvider_ADFSProxyEx1(instance *cim.WmiInstance) (newInstance *Win32_PerfFormattedData_GenevaServerProxyProvider_ADFSProxy, err error) {
	tmp, err := NewWin32_PerfFormattedDataEx1(instance)

	if err != nil {
		return
	}
	newInstance = &Win32_PerfFormattedData_GenevaServerProxyProvider_ADFSProxy{
		Win32_PerfFormattedData: tmp,
	}
	return
}

func NewWin32_PerfFormattedData_GenevaServerProxyProvider_ADFSProxyEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *Win32_PerfFormattedData_GenevaServerProxyProvider_ADFSProxy, err error) {
	tmp, err := NewWin32_PerfFormattedDataEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &Win32_PerfFormattedData_GenevaServerProxyProvider_ADFSProxy{
		Win32_PerfFormattedData: tmp,
	}
	return
}

// SetOutstandingRequests sets the value of OutstandingRequests for the instance
func (instance *Win32_PerfFormattedData_GenevaServerProxyProvider_ADFSProxy) SetPropertyOutstandingRequests(value uint64) (err error) {
	return instance.SetProperty("OutstandingRequests", (value))
}

// GetOutstandingRequests gets the value of OutstandingRequests for the instance
func (instance *Win32_PerfFormattedData_GenevaServerProxyProvider_ADFSProxy) GetPropertyOutstandingRequests() (value uint64, err error) {
	retValue, err := instance.GetProperty("OutstandingRequests")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(uint64)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " uint64 is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = uint64(valuetmp)

	return
}

// SetRejectedRequests sets the value of RejectedRequests for the instance
func (instance *Win32_PerfFormattedData_GenevaServerProxyProvider_ADFSProxy) SetPropertyRejectedRequests(value uint64) (err error) {
	return instance.SetProperty("RejectedRequests", (value))
}

// GetRejectedRequests gets the value of RejectedRequests for the instance
func (instance *Win32_PerfFormattedData_GenevaServerProxyProvider_ADFSProxy) GetPropertyRejectedRequests() (value uint64, err error) {
	retValue, err := instance.GetProperty("RejectedRequests")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(uint64)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " uint64 is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = uint64(valuetmp)

	return
}

// SetRejectedRequestsPersec sets the value of RejectedRequestsPersec for the instance
func (instance *Win32_PerfFormattedData_GenevaServerProxyProvider_ADFSProxy) SetPropertyRejectedRequestsPersec(value uint64) (err error) {
	return instance.SetProperty("RejectedRequestsPersec", (value))
}

// GetRejectedRequestsPersec gets the value of RejectedRequestsPersec for the instance
func (instance *Win32_PerfFormattedData_GenevaServerProxyProvider_ADFSProxy) GetPropertyRejectedRequestsPersec() (value uint64, err error) {
	retValue, err := instance.GetProperty("RejectedRequestsPersec")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(uint64)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " uint64 is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = uint64(valuetmp)

	return
}

// SetRequestLatency sets the value of RequestLatency for the instance
func (instance *Win32_PerfFormattedData_GenevaServerProxyProvider_ADFSProxy) SetPropertyRequestLatency(value uint64) (err error) {
	return instance.SetProperty("RequestLatency", (value))
}

// GetRequestLatency gets the value of RequestLatency for the instance
func (instance *Win32_PerfFormattedData_GenevaServerProxyProvider_ADFSProxy) GetPropertyRequestLatency() (value uint64, err error) {
	retValue, err := instance.GetProperty("RequestLatency")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(uint64)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " uint64 is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = uint64(valuetmp)

	return
}

// SetRequests sets the value of Requests for the instance
func (instance *Win32_PerfFormattedData_GenevaServerProxyProvider_ADFSProxy) SetPropertyRequests(value uint64) (err error) {
	return instance.SetProperty("Requests", (value))
}

// GetRequests gets the value of Requests for the instance
func (instance *Win32_PerfFormattedData_GenevaServerProxyProvider_ADFSProxy) GetPropertyRequests() (value uint64, err error) {
	retValue, err := instance.GetProperty("Requests")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(uint64)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " uint64 is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = uint64(valuetmp)

	return
}

// SetRequestsPersec sets the value of RequestsPersec for the instance
func (instance *Win32_PerfFormattedData_GenevaServerProxyProvider_ADFSProxy) SetPropertyRequestsPersec(value uint64) (err error) {
	return instance.SetProperty("RequestsPersec", (value))
}

// GetRequestsPersec gets the value of RequestsPersec for the instance
func (instance *Win32_PerfFormattedData_GenevaServerProxyProvider_ADFSProxy) GetPropertyRequestsPersec() (value uint64, err error) {
	retValue, err := instance.GetProperty("RequestsPersec")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(uint64)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " uint64 is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = uint64(valuetmp)

	return
}
