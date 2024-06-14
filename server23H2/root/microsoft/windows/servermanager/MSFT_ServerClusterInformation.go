// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 6/14/2024 using wmigen
//      Source root.Microsoft.Windows.ServerManager
//////////////////////////////////////////////
package servermanager

import (
	"github.com/microsoft/wmi/pkg/base/instance"
	"github.com/microsoft/wmi/pkg/base/query"
	"github.com/microsoft/wmi/pkg/errors"
	cim "github.com/microsoft/wmi/pkg/wmiinstance"
	"reflect"
)

// MSFT_ServerClusterInformation struct
type MSFT_ServerClusterInformation struct {
	*cim.WmiInstance

	//
	GroupType int32

	//
	Name string

	//
	ObjectType uint8
}

func NewMSFT_ServerClusterInformationEx1(instance *cim.WmiInstance) (newInstance *MSFT_ServerClusterInformation, err error) {
	tmp, err := instance, nil

	if err != nil {
		return
	}
	newInstance = &MSFT_ServerClusterInformation{
		WmiInstance: tmp,
	}
	return
}

func NewMSFT_ServerClusterInformationEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *MSFT_ServerClusterInformation, err error) {
	tmp, err := instance.GetWmiInstance(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &MSFT_ServerClusterInformation{
		WmiInstance: tmp,
	}
	return
}

// SetGroupType sets the value of GroupType for the instance
func (instance *MSFT_ServerClusterInformation) SetPropertyGroupType(value int32) (err error) {
	return instance.SetProperty("GroupType", (value))
}

// GetGroupType gets the value of GroupType for the instance
func (instance *MSFT_ServerClusterInformation) GetPropertyGroupType() (value int32, err error) {
	retValue, err := instance.GetProperty("GroupType")
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

// SetName sets the value of Name for the instance
func (instance *MSFT_ServerClusterInformation) SetPropertyName(value string) (err error) {
	return instance.SetProperty("Name", (value))
}

// GetName gets the value of Name for the instance
func (instance *MSFT_ServerClusterInformation) GetPropertyName() (value string, err error) {
	retValue, err := instance.GetProperty("Name")
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

// SetObjectType sets the value of ObjectType for the instance
func (instance *MSFT_ServerClusterInformation) SetPropertyObjectType(value uint8) (err error) {
	return instance.SetProperty("ObjectType", (value))
}

// GetObjectType gets the value of ObjectType for the instance
func (instance *MSFT_ServerClusterInformation) GetPropertyObjectType() (value uint8, err error) {
	retValue, err := instance.GetProperty("ObjectType")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(uint8)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " uint8 is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = uint8(valuetmp)

	return
}
