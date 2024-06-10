// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

// Author:
//
//	Auto Generated on 6/6/2024 using wmigen
//	Source root.Microsoft.Windows.RemoteAccess
//
// ////////////////////////////////////////////
package remoteaccess

import (
	"github.com/microsoft/wmi/pkg/base/instance"
	"github.com/microsoft/wmi/pkg/base/query"
	"github.com/microsoft/wmi/pkg/errors"
	cim "github.com/microsoft/wmi/pkg/wmiinstance"
	"reflect"
)

// DAClient struct
type DAClient struct {
	*cim.WmiInstance

	//
	ClientEntrypoint DAClientEntrypoint

	//
	ClientSettings DAClientSettings

	//
	GpoName []string

	//
	SecurityGroupNameList []string
}

func NewDAClientEx1(instance *cim.WmiInstance) (newInstance *DAClient, err error) {
	tmp, err := instance, nil

	if err != nil {
		return
	}
	newInstance = &DAClient{
		WmiInstance: tmp,
	}
	return
}

func NewDAClientEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *DAClient, err error) {
	tmp, err := instance.GetWmiInstance(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &DAClient{
		WmiInstance: tmp,
	}
	return
}

// SetClientEntrypoint sets the value of ClientEntrypoint for the instance
func (instance *DAClient) SetPropertyClientEntrypoint(value DAClientEntrypoint) (err error) {
	return instance.SetProperty("ClientEntrypoint", (value))
}

// GetClientEntrypoint gets the value of ClientEntrypoint for the instance
func (instance *DAClient) GetPropertyClientEntrypoint() (value DAClientEntrypoint, err error) {
	retValue, err := instance.GetProperty("ClientEntrypoint")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(DAClientEntrypoint)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " DAClientEntrypoint is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = DAClientEntrypoint(valuetmp)

	return
}

// SetClientSettings sets the value of ClientSettings for the instance
func (instance *DAClient) SetPropertyClientSettings(value DAClientSettings) (err error) {
	return instance.SetProperty("ClientSettings", (value))
}

// GetClientSettings gets the value of ClientSettings for the instance
func (instance *DAClient) GetPropertyClientSettings() (value DAClientSettings, err error) {
	retValue, err := instance.GetProperty("ClientSettings")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(DAClientSettings)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " DAClientSettings is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = DAClientSettings(valuetmp)

	return
}

// SetGpoName sets the value of GpoName for the instance
func (instance *DAClient) SetPropertyGpoName(value []string) (err error) {
	return instance.SetProperty("GpoName", (value))
}

// GetGpoName gets the value of GpoName for the instance
func (instance *DAClient) GetPropertyGpoName() (value []string, err error) {
	retValue, err := instance.GetProperty("GpoName")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	for _, interfaceValue := range retValue.([]interface{}) {
		valuetmp, ok := interfaceValue.(string)
		if !ok {
			err = errors.Wrapf(errors.InvalidType, " string is Invalid. Expected %s", reflect.TypeOf(interfaceValue))
			return
		}
		value = append(value, string(valuetmp))
	}

	return
}

// SetSecurityGroupNameList sets the value of SecurityGroupNameList for the instance
func (instance *DAClient) SetPropertySecurityGroupNameList(value []string) (err error) {
	return instance.SetProperty("SecurityGroupNameList", (value))
}

// GetSecurityGroupNameList gets the value of SecurityGroupNameList for the instance
func (instance *DAClient) GetPropertySecurityGroupNameList() (value []string, err error) {
	retValue, err := instance.GetProperty("SecurityGroupNameList")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	for _, interfaceValue := range retValue.([]interface{}) {
		valuetmp, ok := interfaceValue.(string)
		if !ok {
			err = errors.Wrapf(errors.InvalidType, " string is Invalid. Expected %s", reflect.TypeOf(interfaceValue))
			return
		}
		value = append(value, string(valuetmp))
	}

	return
}
