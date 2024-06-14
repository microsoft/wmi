// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 6/13/2024 using wmigen
//      Source root.Microsoft.Windows.RemoteAccess.Server
//////////////////////////////////////////////
package server

import (
	"github.com/microsoft/wmi/pkg/base/query"
	"github.com/microsoft/wmi/pkg/errors"
	cim "github.com/microsoft/wmi/pkg/wmiinstance"
	"reflect"
)

// RemoteAccessActiveConnectionSummaryLocal struct
type RemoteAccessActiveConnectionSummaryLocal struct {
	*RemoteAccessConnectionSummaryLocal

	//
	MaxConcurrentConnections uint32

	//
	TotalBytesIn uint64

	//
	TotalBytesInOut uint64

	//
	TotalBytesOut uint64

	//
	TotalConnections uint32

	//
	TotalCumulativeConnections uint32

	//
	TotalDAConnections uint32

	//
	TotalVpnConnections uint32
}

func NewRemoteAccessActiveConnectionSummaryLocalEx1(instance *cim.WmiInstance) (newInstance *RemoteAccessActiveConnectionSummaryLocal, err error) {
	tmp, err := NewRemoteAccessConnectionSummaryLocalEx1(instance)

	if err != nil {
		return
	}
	newInstance = &RemoteAccessActiveConnectionSummaryLocal{
		RemoteAccessConnectionSummaryLocal: tmp,
	}
	return
}

func NewRemoteAccessActiveConnectionSummaryLocalEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *RemoteAccessActiveConnectionSummaryLocal, err error) {
	tmp, err := NewRemoteAccessConnectionSummaryLocalEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &RemoteAccessActiveConnectionSummaryLocal{
		RemoteAccessConnectionSummaryLocal: tmp,
	}
	return
}

// SetMaxConcurrentConnections sets the value of MaxConcurrentConnections for the instance
func (instance *RemoteAccessActiveConnectionSummaryLocal) SetPropertyMaxConcurrentConnections(value uint32) (err error) {
	return instance.SetProperty("MaxConcurrentConnections", (value))
}

// GetMaxConcurrentConnections gets the value of MaxConcurrentConnections for the instance
func (instance *RemoteAccessActiveConnectionSummaryLocal) GetPropertyMaxConcurrentConnections() (value uint32, err error) {
	retValue, err := instance.GetProperty("MaxConcurrentConnections")
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

// SetTotalBytesIn sets the value of TotalBytesIn for the instance
func (instance *RemoteAccessActiveConnectionSummaryLocal) SetPropertyTotalBytesIn(value uint64) (err error) {
	return instance.SetProperty("TotalBytesIn", (value))
}

// GetTotalBytesIn gets the value of TotalBytesIn for the instance
func (instance *RemoteAccessActiveConnectionSummaryLocal) GetPropertyTotalBytesIn() (value uint64, err error) {
	retValue, err := instance.GetProperty("TotalBytesIn")
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

// SetTotalBytesInOut sets the value of TotalBytesInOut for the instance
func (instance *RemoteAccessActiveConnectionSummaryLocal) SetPropertyTotalBytesInOut(value uint64) (err error) {
	return instance.SetProperty("TotalBytesInOut", (value))
}

// GetTotalBytesInOut gets the value of TotalBytesInOut for the instance
func (instance *RemoteAccessActiveConnectionSummaryLocal) GetPropertyTotalBytesInOut() (value uint64, err error) {
	retValue, err := instance.GetProperty("TotalBytesInOut")
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

// SetTotalBytesOut sets the value of TotalBytesOut for the instance
func (instance *RemoteAccessActiveConnectionSummaryLocal) SetPropertyTotalBytesOut(value uint64) (err error) {
	return instance.SetProperty("TotalBytesOut", (value))
}

// GetTotalBytesOut gets the value of TotalBytesOut for the instance
func (instance *RemoteAccessActiveConnectionSummaryLocal) GetPropertyTotalBytesOut() (value uint64, err error) {
	retValue, err := instance.GetProperty("TotalBytesOut")
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

// SetTotalConnections sets the value of TotalConnections for the instance
func (instance *RemoteAccessActiveConnectionSummaryLocal) SetPropertyTotalConnections(value uint32) (err error) {
	return instance.SetProperty("TotalConnections", (value))
}

// GetTotalConnections gets the value of TotalConnections for the instance
func (instance *RemoteAccessActiveConnectionSummaryLocal) GetPropertyTotalConnections() (value uint32, err error) {
	retValue, err := instance.GetProperty("TotalConnections")
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

// SetTotalCumulativeConnections sets the value of TotalCumulativeConnections for the instance
func (instance *RemoteAccessActiveConnectionSummaryLocal) SetPropertyTotalCumulativeConnections(value uint32) (err error) {
	return instance.SetProperty("TotalCumulativeConnections", (value))
}

// GetTotalCumulativeConnections gets the value of TotalCumulativeConnections for the instance
func (instance *RemoteAccessActiveConnectionSummaryLocal) GetPropertyTotalCumulativeConnections() (value uint32, err error) {
	retValue, err := instance.GetProperty("TotalCumulativeConnections")
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

// SetTotalDAConnections sets the value of TotalDAConnections for the instance
func (instance *RemoteAccessActiveConnectionSummaryLocal) SetPropertyTotalDAConnections(value uint32) (err error) {
	return instance.SetProperty("TotalDAConnections", (value))
}

// GetTotalDAConnections gets the value of TotalDAConnections for the instance
func (instance *RemoteAccessActiveConnectionSummaryLocal) GetPropertyTotalDAConnections() (value uint32, err error) {
	retValue, err := instance.GetProperty("TotalDAConnections")
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

// SetTotalVpnConnections sets the value of TotalVpnConnections for the instance
func (instance *RemoteAccessActiveConnectionSummaryLocal) SetPropertyTotalVpnConnections(value uint32) (err error) {
	return instance.SetProperty("TotalVpnConnections", (value))
}

// GetTotalVpnConnections gets the value of TotalVpnConnections for the instance
func (instance *RemoteAccessActiveConnectionSummaryLocal) GetPropertyTotalVpnConnections() (value uint32, err error) {
	retValue, err := instance.GetProperty("TotalVpnConnections")
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
