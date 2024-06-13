// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 6/13/2024 using wmigen
//      Source root.Microsoft.Windows.Wdac
//////////////////////////////////////////////
package wdac

import (
	"github.com/microsoft/wmi/pkg/base/instance"
	"github.com/microsoft/wmi/pkg/base/query"
	"github.com/microsoft/wmi/pkg/errors"
	cim "github.com/microsoft/wmi/pkg/wmiinstance"
	"reflect"
)

// MSFT_OdbcPerfCounter struct
type MSFT_OdbcPerfCounter struct {
	*cim.WmiInstance

	//
	IsEnabled bool

	//
	Platform string
}

func NewMSFT_OdbcPerfCounterEx1(instance *cim.WmiInstance) (newInstance *MSFT_OdbcPerfCounter, err error) {
	tmp, err := instance, nil

	if err != nil {
		return
	}
	newInstance = &MSFT_OdbcPerfCounter{
		WmiInstance: tmp,
	}
	return
}

func NewMSFT_OdbcPerfCounterEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *MSFT_OdbcPerfCounter, err error) {
	tmp, err := instance.GetWmiInstance(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &MSFT_OdbcPerfCounter{
		WmiInstance: tmp,
	}
	return
}

// SetIsEnabled sets the value of IsEnabled for the instance
func (instance *MSFT_OdbcPerfCounter) SetPropertyIsEnabled(value bool) (err error) {
	return instance.SetProperty("IsEnabled", (value))
}

// GetIsEnabled gets the value of IsEnabled for the instance
func (instance *MSFT_OdbcPerfCounter) GetPropertyIsEnabled() (value bool, err error) {
	retValue, err := instance.GetProperty("IsEnabled")
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

// SetPlatform sets the value of Platform for the instance
func (instance *MSFT_OdbcPerfCounter) SetPropertyPlatform(value string) (err error) {
	return instance.SetProperty("Platform", (value))
}

// GetPlatform gets the value of Platform for the instance
func (instance *MSFT_OdbcPerfCounter) GetPropertyPlatform() (value string, err error) {
	retValue, err := instance.GetProperty("Platform")
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
