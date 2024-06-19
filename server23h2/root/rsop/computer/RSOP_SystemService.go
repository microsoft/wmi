// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 6/14/2024 using wmigen
//      Source root.RSOP.Computer
//////////////////////////////////////////////
package computer

import (
	"github.com/microsoft/wmi/pkg/base/query"
	"github.com/microsoft/wmi/pkg/errors"
	cim "github.com/microsoft/wmi/pkg/wmiinstance"
	"reflect"
)

// RSOP_SystemService struct
type RSOP_SystemService struct {
	*RSOP_SecuritySettings

	//
	SDDLString string

	//
	Service string

	//
	StartupMode SystemService_StartupMode
}

func NewRSOP_SystemServiceEx1(instance *cim.WmiInstance) (newInstance *RSOP_SystemService, err error) {
	tmp, err := NewRSOP_SecuritySettingsEx1(instance)

	if err != nil {
		return
	}
	newInstance = &RSOP_SystemService{
		RSOP_SecuritySettings: tmp,
	}
	return
}

func NewRSOP_SystemServiceEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *RSOP_SystemService, err error) {
	tmp, err := NewRSOP_SecuritySettingsEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &RSOP_SystemService{
		RSOP_SecuritySettings: tmp,
	}
	return
}

// SetSDDLString sets the value of SDDLString for the instance
func (instance *RSOP_SystemService) SetPropertySDDLString(value string) (err error) {
	return instance.SetProperty("SDDLString", (value))
}

// GetSDDLString gets the value of SDDLString for the instance
func (instance *RSOP_SystemService) GetPropertySDDLString() (value string, err error) {
	retValue, err := instance.GetProperty("SDDLString")
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

// SetService sets the value of Service for the instance
func (instance *RSOP_SystemService) SetPropertyService(value string) (err error) {
	return instance.SetProperty("Service", (value))
}

// GetService gets the value of Service for the instance
func (instance *RSOP_SystemService) GetPropertyService() (value string, err error) {
	retValue, err := instance.GetProperty("Service")
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

// SetStartupMode sets the value of StartupMode for the instance
func (instance *RSOP_SystemService) SetPropertyStartupMode(value SystemService_StartupMode) (err error) {
	return instance.SetProperty("StartupMode", (value))
}

// GetStartupMode gets the value of StartupMode for the instance
func (instance *RSOP_SystemService) GetPropertyStartupMode() (value SystemService_StartupMode, err error) {
	retValue, err := instance.GetProperty("StartupMode")
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

	value = SystemService_StartupMode(valuetmp)

	return
}
