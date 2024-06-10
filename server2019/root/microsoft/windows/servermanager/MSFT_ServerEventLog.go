// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 9/18/2020 using wmigen
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

// MSFT_ServerEventLog struct
type MSFT_ServerEventLog struct {
	*cim.WmiInstance

	//
	LocalizedName string

	//
	Name string
}

func NewMSFT_ServerEventLogEx1(instance *cim.WmiInstance) (newInstance *MSFT_ServerEventLog, err error) {
	tmp, err := instance, nil

	if err != nil {
		return
	}
	newInstance = &MSFT_ServerEventLog{
		WmiInstance: tmp,
	}
	return
}

func NewMSFT_ServerEventLogEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *MSFT_ServerEventLog, err error) {
	tmp, err := instance.GetWmiInstance(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &MSFT_ServerEventLog{
		WmiInstance: tmp,
	}
	return
}

// SetLocalizedName sets the value of LocalizedName for the instance
func (instance *MSFT_ServerEventLog) SetPropertyLocalizedName(value string) (err error) {
	return instance.SetProperty("LocalizedName", (value))
}

// GetLocalizedName gets the value of LocalizedName for the instance
func (instance *MSFT_ServerEventLog) GetPropertyLocalizedName() (value string, err error) {
	retValue, err := instance.GetProperty("LocalizedName")
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

// SetName sets the value of Name for the instance
func (instance *MSFT_ServerEventLog) SetPropertyName(value string) (err error) {
	return instance.SetProperty("Name", (value))
}

// GetName gets the value of Name for the instance
func (instance *MSFT_ServerEventLog) GetPropertyName() (value string, err error) {
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
