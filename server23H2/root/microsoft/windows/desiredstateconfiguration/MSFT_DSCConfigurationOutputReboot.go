// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 6/12/2024 using wmigen
//      Source root.Microsoft.Windows.DesiredStateConfiguration
//////////////////////////////////////////////
package desiredstateconfiguration

import (
	"github.com/microsoft/wmi/pkg/base/query"
	"github.com/microsoft/wmi/pkg/errors"
	cim "github.com/microsoft/wmi/pkg/wmiinstance"
	"reflect"
)

// MSFT_DSCConfigurationOutputReboot struct
type MSFT_DSCConfigurationOutputReboot struct {
	*MSFT_DSCConfigurationOutput

	//
	Automatic bool
}

func NewMSFT_DSCConfigurationOutputRebootEx1(instance *cim.WmiInstance) (newInstance *MSFT_DSCConfigurationOutputReboot, err error) {
	tmp, err := NewMSFT_DSCConfigurationOutputEx1(instance)

	if err != nil {
		return
	}
	newInstance = &MSFT_DSCConfigurationOutputReboot{
		MSFT_DSCConfigurationOutput: tmp,
	}
	return
}

func NewMSFT_DSCConfigurationOutputRebootEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *MSFT_DSCConfigurationOutputReboot, err error) {
	tmp, err := NewMSFT_DSCConfigurationOutputEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &MSFT_DSCConfigurationOutputReboot{
		MSFT_DSCConfigurationOutput: tmp,
	}
	return
}

// SetAutomatic sets the value of Automatic for the instance
func (instance *MSFT_DSCConfigurationOutputReboot) SetPropertyAutomatic(value bool) (err error) {
	return instance.SetProperty("Automatic", (value))
}

// GetAutomatic gets the value of Automatic for the instance
func (instance *MSFT_DSCConfigurationOutputReboot) GetPropertyAutomatic() (value bool, err error) {
	retValue, err := instance.GetProperty("Automatic")
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
