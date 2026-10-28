// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 10/28/2026 using wmigen
//      Source root.WMI
//////////////////////////////////////////////
package wmi

import (
	"github.com/microsoft/wmi/pkg/base/instance"
	"github.com/microsoft/wmi/pkg/base/query"
	"github.com/microsoft/wmi/pkg/errors"
	cim "github.com/microsoft/wmi/pkg/wmiinstance"
	"reflect"
)

// TeeDriverMeVersionWmiInfo struct
type TeeDriverMeVersionWmiInfo struct {
	*cim.WmiInstance

	//
	Active bool

	// Current FW Image Type, Corporate or Consumer Image
	FWType uint32

	// Current ME\TXE FW Version
	FWVer []WmiFWVersion

	//
	InstanceName string
}

func NewTeeDriverMeVersionWmiInfoEx1(instance *cim.WmiInstance) (newInstance *TeeDriverMeVersionWmiInfo, err error) {
	tmp, err := instance, nil

	if err != nil {
		return
	}
	newInstance = &TeeDriverMeVersionWmiInfo{
		WmiInstance: tmp,
	}
	return
}

func NewTeeDriverMeVersionWmiInfoEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *TeeDriverMeVersionWmiInfo, err error) {
	tmp, err := instance.GetWmiInstance(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &TeeDriverMeVersionWmiInfo{
		WmiInstance: tmp,
	}
	return
}

// SetActive sets the value of Active for the instance
func (instance *TeeDriverMeVersionWmiInfo) SetPropertyActive(value bool) (err error) {
	return instance.SetProperty("Active", (value))
}

// GetActive gets the value of Active for the instance
func (instance *TeeDriverMeVersionWmiInfo) GetPropertyActive() (value bool, err error) {
	retValue, err := instance.GetProperty("Active")
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

// SetFWType sets the value of FWType for the instance
func (instance *TeeDriverMeVersionWmiInfo) SetPropertyFWType(value uint32) (err error) {
	return instance.SetProperty("FWType", (value))
}

// GetFWType gets the value of FWType for the instance
func (instance *TeeDriverMeVersionWmiInfo) GetPropertyFWType() (value uint32, err error) {
	retValue, err := instance.GetProperty("FWType")
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

// SetFWVer sets the value of FWVer for the instance
func (instance *TeeDriverMeVersionWmiInfo) SetPropertyFWVer(value []WmiFWVersion) (err error) {
	return instance.SetProperty("FWVer", (value))
}

// GetFWVer gets the value of FWVer for the instance
func (instance *TeeDriverMeVersionWmiInfo) GetPropertyFWVer() (value []WmiFWVersion, err error) {
	retValue, err := instance.GetProperty("FWVer")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	for _, interfaceValue := range retValue.([]interface{}) {
		valuetmp, ok := interfaceValue.(WmiFWVersion)
		if !ok {
			err = errors.Wrapf(errors.InvalidType, " WmiFWVersion is Invalid. Expected %s", reflect.TypeOf(interfaceValue))
			return
		}
		value = append(value, WmiFWVersion(valuetmp))
	}

	return
}

// SetInstanceName sets the value of InstanceName for the instance
func (instance *TeeDriverMeVersionWmiInfo) SetPropertyInstanceName(value string) (err error) {
	return instance.SetProperty("InstanceName", (value))
}

// GetInstanceName gets the value of InstanceName for the instance
func (instance *TeeDriverMeVersionWmiInfo) GetPropertyInstanceName() (value string, err error) {
	retValue, err := instance.GetProperty("InstanceName")
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
