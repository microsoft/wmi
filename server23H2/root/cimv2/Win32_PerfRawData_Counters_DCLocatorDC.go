// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 6/13/2024 using wmigen
//      Source root.CIMV2
//////////////////////////////////////////////
package cimv2

import (
	"github.com/microsoft/wmi/pkg/base/query"
	"github.com/microsoft/wmi/pkg/errors"
	cim "github.com/microsoft/wmi/pkg/wmiinstance"
	"reflect"
)

// Win32_PerfRawData_Counters_DCLocatorDC struct
type Win32_PerfRawData_Counters_DCLocatorDC struct {
	*Win32_PerfRawData

	//
	PingsActiveMailslotPings uint32

	//
	PingsActiveUDPLDAPPings uint32

	//
	PingsAverageMailslotPingLatencysecs uint32

	//
	PingsAverageMailslotPingLatencysecs_Base uint32

	//
	PingsAverageUDPLDAPPingLatencysecs uint32

	//
	PingsAverageUDPLDAPPingLatencysecs_Base uint32

	//
	PingsMailslotPingsReceivedPersec uint32

	//
	PingsUDPLDAPPingsReceivedPersec uint32
}

func NewWin32_PerfRawData_Counters_DCLocatorDCEx1(instance *cim.WmiInstance) (newInstance *Win32_PerfRawData_Counters_DCLocatorDC, err error) {
	tmp, err := NewWin32_PerfRawDataEx1(instance)

	if err != nil {
		return
	}
	newInstance = &Win32_PerfRawData_Counters_DCLocatorDC{
		Win32_PerfRawData: tmp,
	}
	return
}

func NewWin32_PerfRawData_Counters_DCLocatorDCEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *Win32_PerfRawData_Counters_DCLocatorDC, err error) {
	tmp, err := NewWin32_PerfRawDataEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &Win32_PerfRawData_Counters_DCLocatorDC{
		Win32_PerfRawData: tmp,
	}
	return
}

// SetPingsActiveMailslotPings sets the value of PingsActiveMailslotPings for the instance
func (instance *Win32_PerfRawData_Counters_DCLocatorDC) SetPropertyPingsActiveMailslotPings(value uint32) (err error) {
	return instance.SetProperty("PingsActiveMailslotPings", (value))
}

// GetPingsActiveMailslotPings gets the value of PingsActiveMailslotPings for the instance
func (instance *Win32_PerfRawData_Counters_DCLocatorDC) GetPropertyPingsActiveMailslotPings() (value uint32, err error) {
	retValue, err := instance.GetProperty("PingsActiveMailslotPings")
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

// SetPingsActiveUDPLDAPPings sets the value of PingsActiveUDPLDAPPings for the instance
func (instance *Win32_PerfRawData_Counters_DCLocatorDC) SetPropertyPingsActiveUDPLDAPPings(value uint32) (err error) {
	return instance.SetProperty("PingsActiveUDPLDAPPings", (value))
}

// GetPingsActiveUDPLDAPPings gets the value of PingsActiveUDPLDAPPings for the instance
func (instance *Win32_PerfRawData_Counters_DCLocatorDC) GetPropertyPingsActiveUDPLDAPPings() (value uint32, err error) {
	retValue, err := instance.GetProperty("PingsActiveUDPLDAPPings")
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

// SetPingsAverageMailslotPingLatencysecs sets the value of PingsAverageMailslotPingLatencysecs for the instance
func (instance *Win32_PerfRawData_Counters_DCLocatorDC) SetPropertyPingsAverageMailslotPingLatencysecs(value uint32) (err error) {
	return instance.SetProperty("PingsAverageMailslotPingLatencysecs", (value))
}

// GetPingsAverageMailslotPingLatencysecs gets the value of PingsAverageMailslotPingLatencysecs for the instance
func (instance *Win32_PerfRawData_Counters_DCLocatorDC) GetPropertyPingsAverageMailslotPingLatencysecs() (value uint32, err error) {
	retValue, err := instance.GetProperty("PingsAverageMailslotPingLatencysecs")
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

// SetPingsAverageMailslotPingLatencysecs_Base sets the value of PingsAverageMailslotPingLatencysecs_Base for the instance
func (instance *Win32_PerfRawData_Counters_DCLocatorDC) SetPropertyPingsAverageMailslotPingLatencysecs_Base(value uint32) (err error) {
	return instance.SetProperty("PingsAverageMailslotPingLatencysecs_Base", (value))
}

// GetPingsAverageMailslotPingLatencysecs_Base gets the value of PingsAverageMailslotPingLatencysecs_Base for the instance
func (instance *Win32_PerfRawData_Counters_DCLocatorDC) GetPropertyPingsAverageMailslotPingLatencysecs_Base() (value uint32, err error) {
	retValue, err := instance.GetProperty("PingsAverageMailslotPingLatencysecs_Base")
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

// SetPingsAverageUDPLDAPPingLatencysecs sets the value of PingsAverageUDPLDAPPingLatencysecs for the instance
func (instance *Win32_PerfRawData_Counters_DCLocatorDC) SetPropertyPingsAverageUDPLDAPPingLatencysecs(value uint32) (err error) {
	return instance.SetProperty("PingsAverageUDPLDAPPingLatencysecs", (value))
}

// GetPingsAverageUDPLDAPPingLatencysecs gets the value of PingsAverageUDPLDAPPingLatencysecs for the instance
func (instance *Win32_PerfRawData_Counters_DCLocatorDC) GetPropertyPingsAverageUDPLDAPPingLatencysecs() (value uint32, err error) {
	retValue, err := instance.GetProperty("PingsAverageUDPLDAPPingLatencysecs")
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

// SetPingsAverageUDPLDAPPingLatencysecs_Base sets the value of PingsAverageUDPLDAPPingLatencysecs_Base for the instance
func (instance *Win32_PerfRawData_Counters_DCLocatorDC) SetPropertyPingsAverageUDPLDAPPingLatencysecs_Base(value uint32) (err error) {
	return instance.SetProperty("PingsAverageUDPLDAPPingLatencysecs_Base", (value))
}

// GetPingsAverageUDPLDAPPingLatencysecs_Base gets the value of PingsAverageUDPLDAPPingLatencysecs_Base for the instance
func (instance *Win32_PerfRawData_Counters_DCLocatorDC) GetPropertyPingsAverageUDPLDAPPingLatencysecs_Base() (value uint32, err error) {
	retValue, err := instance.GetProperty("PingsAverageUDPLDAPPingLatencysecs_Base")
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

// SetPingsMailslotPingsReceivedPersec sets the value of PingsMailslotPingsReceivedPersec for the instance
func (instance *Win32_PerfRawData_Counters_DCLocatorDC) SetPropertyPingsMailslotPingsReceivedPersec(value uint32) (err error) {
	return instance.SetProperty("PingsMailslotPingsReceivedPersec", (value))
}

// GetPingsMailslotPingsReceivedPersec gets the value of PingsMailslotPingsReceivedPersec for the instance
func (instance *Win32_PerfRawData_Counters_DCLocatorDC) GetPropertyPingsMailslotPingsReceivedPersec() (value uint32, err error) {
	retValue, err := instance.GetProperty("PingsMailslotPingsReceivedPersec")
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

// SetPingsUDPLDAPPingsReceivedPersec sets the value of PingsUDPLDAPPingsReceivedPersec for the instance
func (instance *Win32_PerfRawData_Counters_DCLocatorDC) SetPropertyPingsUDPLDAPPingsReceivedPersec(value uint32) (err error) {
	return instance.SetProperty("PingsUDPLDAPPingsReceivedPersec", (value))
}

// GetPingsUDPLDAPPingsReceivedPersec gets the value of PingsUDPLDAPPingsReceivedPersec for the instance
func (instance *Win32_PerfRawData_Counters_DCLocatorDC) GetPropertyPingsUDPLDAPPingsReceivedPersec() (value uint32, err error) {
	retValue, err := instance.GetProperty("PingsUDPLDAPPingsReceivedPersec")
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
