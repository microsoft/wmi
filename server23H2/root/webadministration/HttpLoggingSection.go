// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

// Author:
//
//	Auto Generated on 6/6/2024 using wmigen
//	Source root.WebAdministration
//
// ////////////////////////////////////////////
package webadministration

import (
	"github.com/microsoft/wmi/pkg/base/query"
	"github.com/microsoft/wmi/pkg/errors"
	cim "github.com/microsoft/wmi/pkg/wmiinstance"
	"reflect"
)

// HttpLoggingSection struct
type HttpLoggingSection struct {
	*ConfigurationSection

	//
	DontLog bool

	//
	SelectiveLogging int32
}

func NewHttpLoggingSectionEx1(instance *cim.WmiInstance) (newInstance *HttpLoggingSection, err error) {
	tmp, err := NewConfigurationSectionEx1(instance)

	if err != nil {
		return
	}
	newInstance = &HttpLoggingSection{
		ConfigurationSection: tmp,
	}
	return
}

func NewHttpLoggingSectionEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *HttpLoggingSection, err error) {
	tmp, err := NewConfigurationSectionEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &HttpLoggingSection{
		ConfigurationSection: tmp,
	}
	return
}

// SetDontLog sets the value of DontLog for the instance
func (instance *HttpLoggingSection) SetPropertyDontLog(value bool) (err error) {
	return instance.SetProperty("DontLog", (value))
}

// GetDontLog gets the value of DontLog for the instance
func (instance *HttpLoggingSection) GetPropertyDontLog() (value bool, err error) {
	retValue, err := instance.GetProperty("DontLog")
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

// SetSelectiveLogging sets the value of SelectiveLogging for the instance
func (instance *HttpLoggingSection) SetPropertySelectiveLogging(value int32) (err error) {
	return instance.SetProperty("SelectiveLogging", (value))
}

// GetSelectiveLogging gets the value of SelectiveLogging for the instance
func (instance *HttpLoggingSection) GetPropertySelectiveLogging() (value int32, err error) {
	retValue, err := instance.GetProperty("SelectiveLogging")
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
