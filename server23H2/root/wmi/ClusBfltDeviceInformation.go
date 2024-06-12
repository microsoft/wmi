// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 6/12/2024 using wmigen
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

// ClusBfltDeviceInformation struct
type ClusBfltDeviceInformation struct {
	*cim.WmiInstance

	//
	Active bool

	//
	InstanceName string

	// Number of Paths.
	NumberOfPaths uint32

	// Path Info.
	PathInfo []ClusBfltPathInformation
}

func NewClusBfltDeviceInformationEx1(instance *cim.WmiInstance) (newInstance *ClusBfltDeviceInformation, err error) {
	tmp, err := instance, nil

	if err != nil {
		return
	}
	newInstance = &ClusBfltDeviceInformation{
		WmiInstance: tmp,
	}
	return
}

func NewClusBfltDeviceInformationEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *ClusBfltDeviceInformation, err error) {
	tmp, err := instance.GetWmiInstance(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &ClusBfltDeviceInformation{
		WmiInstance: tmp,
	}
	return
}

// SetActive sets the value of Active for the instance
func (instance *ClusBfltDeviceInformation) SetPropertyActive(value bool) (err error) {
	return instance.SetProperty("Active", (value))
}

// GetActive gets the value of Active for the instance
func (instance *ClusBfltDeviceInformation) GetPropertyActive() (value bool, err error) {
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

// SetInstanceName sets the value of InstanceName for the instance
func (instance *ClusBfltDeviceInformation) SetPropertyInstanceName(value string) (err error) {
	return instance.SetProperty("InstanceName", (value))
}

// GetInstanceName gets the value of InstanceName for the instance
func (instance *ClusBfltDeviceInformation) GetPropertyInstanceName() (value string, err error) {
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

// SetNumberOfPaths sets the value of NumberOfPaths for the instance
func (instance *ClusBfltDeviceInformation) SetPropertyNumberOfPaths(value uint32) (err error) {
	return instance.SetProperty("NumberOfPaths", (value))
}

// GetNumberOfPaths gets the value of NumberOfPaths for the instance
func (instance *ClusBfltDeviceInformation) GetPropertyNumberOfPaths() (value uint32, err error) {
	retValue, err := instance.GetProperty("NumberOfPaths")
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

// SetPathInfo sets the value of PathInfo for the instance
func (instance *ClusBfltDeviceInformation) SetPropertyPathInfo(value []ClusBfltPathInformation) (err error) {
	return instance.SetProperty("PathInfo", (value))
}

// GetPathInfo gets the value of PathInfo for the instance
func (instance *ClusBfltDeviceInformation) GetPropertyPathInfo() (value []ClusBfltPathInformation, err error) {
	retValue, err := instance.GetProperty("PathInfo")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	for _, interfaceValue := range retValue.([]interface{}) {
		valuetmp, ok := interfaceValue.(ClusBfltPathInformation)
		if !ok {
			err = errors.Wrapf(errors.InvalidType, " ClusBfltPathInformation is Invalid. Expected %s", reflect.TypeOf(interfaceValue))
			return
		}
		value = append(value, ClusBfltPathInformation(valuetmp))
	}

	return
}
