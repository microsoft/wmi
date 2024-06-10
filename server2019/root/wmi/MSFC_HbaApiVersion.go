// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

// Author:
//
//	Auto Generated on 6/6/2024 using wmigen
//	Source root.WMI
//
// ////////////////////////////////////////////
package wmi

import (
	"github.com/microsoft/wmi/pkg/base/instance"
	"github.com/microsoft/wmi/pkg/base/query"
	"github.com/microsoft/wmi/pkg/errors"
	cim "github.com/microsoft/wmi/pkg/wmiinstance"
	"reflect"
)

// MSFC_HbaApiVersion struct
type MSFC_HbaApiVersion struct {
	*cim.WmiInstance

	//
	Description string

	//
	HbaApiVersion uint32

	//
	WmiHbaApiVersion uint32
}

func NewMSFC_HbaApiVersionEx1(instance *cim.WmiInstance) (newInstance *MSFC_HbaApiVersion, err error) {
	tmp, err := instance, nil

	if err != nil {
		return
	}
	newInstance = &MSFC_HbaApiVersion{
		WmiInstance: tmp,
	}
	return
}

func NewMSFC_HbaApiVersionEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *MSFC_HbaApiVersion, err error) {
	tmp, err := instance.GetWmiInstance(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &MSFC_HbaApiVersion{
		WmiInstance: tmp,
	}
	return
}

// SetDescription sets the value of Description for the instance
func (instance *MSFC_HbaApiVersion) SetPropertyDescription(value string) (err error) {
	return instance.SetProperty("Description", (value))
}

// GetDescription gets the value of Description for the instance
func (instance *MSFC_HbaApiVersion) GetPropertyDescription() (value string, err error) {
	retValue, err := instance.GetProperty("Description")
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

// SetHbaApiVersion sets the value of HbaApiVersion for the instance
func (instance *MSFC_HbaApiVersion) SetPropertyHbaApiVersion(value uint32) (err error) {
	return instance.SetProperty("HbaApiVersion", (value))
}

// GetHbaApiVersion gets the value of HbaApiVersion for the instance
func (instance *MSFC_HbaApiVersion) GetPropertyHbaApiVersion() (value uint32, err error) {
	retValue, err := instance.GetProperty("HbaApiVersion")
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

// SetWmiHbaApiVersion sets the value of WmiHbaApiVersion for the instance
func (instance *MSFC_HbaApiVersion) SetPropertyWmiHbaApiVersion(value uint32) (err error) {
	return instance.SetProperty("WmiHbaApiVersion", (value))
}

// GetWmiHbaApiVersion gets the value of WmiHbaApiVersion for the instance
func (instance *MSFC_HbaApiVersion) GetPropertyWmiHbaApiVersion() (value uint32, err error) {
	retValue, err := instance.GetProperty("WmiHbaApiVersion")
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
