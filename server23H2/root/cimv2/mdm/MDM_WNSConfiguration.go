// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

// Author:
//
//	Auto Generated on 6/6/2024 using wmigen
//	Source root.CIMV2.mdm
//
// ////////////////////////////////////////////
package mdm

import (
	"github.com/microsoft/wmi/pkg/base/instance"
	"github.com/microsoft/wmi/pkg/base/query"
	"github.com/microsoft/wmi/pkg/errors"
	cim "github.com/microsoft/wmi/pkg/wmiinstance"
	"reflect"
)

// MDM_WNSConfiguration struct
type MDM_WNSConfiguration struct {
	*cim.WmiInstance

	//
	AppId string

	//
	ConfigurationStatus uint32
}

func NewMDM_WNSConfigurationEx1(instance *cim.WmiInstance) (newInstance *MDM_WNSConfiguration, err error) {
	tmp, err := instance, nil

	if err != nil {
		return
	}
	newInstance = &MDM_WNSConfiguration{
		WmiInstance: tmp,
	}
	return
}

func NewMDM_WNSConfigurationEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *MDM_WNSConfiguration, err error) {
	tmp, err := instance.GetWmiInstance(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &MDM_WNSConfiguration{
		WmiInstance: tmp,
	}
	return
}

// SetAppId sets the value of AppId for the instance
func (instance *MDM_WNSConfiguration) SetPropertyAppId(value string) (err error) {
	return instance.SetProperty("AppId", (value))
}

// GetAppId gets the value of AppId for the instance
func (instance *MDM_WNSConfiguration) GetPropertyAppId() (value string, err error) {
	retValue, err := instance.GetProperty("AppId")
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

// SetConfigurationStatus sets the value of ConfigurationStatus for the instance
func (instance *MDM_WNSConfiguration) SetPropertyConfigurationStatus(value uint32) (err error) {
	return instance.SetProperty("ConfigurationStatus", (value))
}

// GetConfigurationStatus gets the value of ConfigurationStatus for the instance
func (instance *MDM_WNSConfiguration) GetPropertyConfigurationStatus() (value uint32, err error) {
	retValue, err := instance.GetProperty("ConfigurationStatus")
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

//

// <param name="ConfigString" type="string "></param>

// <param name="ReturnValue" type="uint32 "></param>
func (instance *MDM_WNSConfiguration) UpdateConfiguration( /* IN */ ConfigString string) (result uint32, err error) {
	retVal, err := instance.InvokeMethodWithReturn("UpdateConfiguration", ConfigString)
	if err != nil {
		return
	}
	result = uint32(retVal)
	return

}
