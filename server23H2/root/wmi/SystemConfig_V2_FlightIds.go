// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 6/14/2024 using wmigen
//      Source root.WMI
//////////////////////////////////////////////
package wmi

import (
	"github.com/microsoft/wmi/pkg/base/query"
	"github.com/microsoft/wmi/pkg/errors"
	cim "github.com/microsoft/wmi/pkg/wmiinstance"
	"reflect"
)

// SystemConfig_V2_FlightIds struct
type SystemConfig_V2_FlightIds struct {
	*SystemConfig_V2

	//
	FlightIdList string

	//
	UpdateId string
}

func NewSystemConfig_V2_FlightIdsEx1(instance *cim.WmiInstance) (newInstance *SystemConfig_V2_FlightIds, err error) {
	tmp, err := NewSystemConfig_V2Ex1(instance)

	if err != nil {
		return
	}
	newInstance = &SystemConfig_V2_FlightIds{
		SystemConfig_V2: tmp,
	}
	return
}

func NewSystemConfig_V2_FlightIdsEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *SystemConfig_V2_FlightIds, err error) {
	tmp, err := NewSystemConfig_V2Ex6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &SystemConfig_V2_FlightIds{
		SystemConfig_V2: tmp,
	}
	return
}

// SetFlightIdList sets the value of FlightIdList for the instance
func (instance *SystemConfig_V2_FlightIds) SetPropertyFlightIdList(value string) (err error) {
	return instance.SetProperty("FlightIdList", (value))
}

// GetFlightIdList gets the value of FlightIdList for the instance
func (instance *SystemConfig_V2_FlightIds) GetPropertyFlightIdList() (value string, err error) {
	retValue, err := instance.GetProperty("FlightIdList")
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

// SetUpdateId sets the value of UpdateId for the instance
func (instance *SystemConfig_V2_FlightIds) SetPropertyUpdateId(value string) (err error) {
	return instance.SetProperty("UpdateId", (value))
}

// GetUpdateId gets the value of UpdateId for the instance
func (instance *SystemConfig_V2_FlightIds) GetPropertyUpdateId() (value string, err error) {
	retValue, err := instance.GetProperty("UpdateId")
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
