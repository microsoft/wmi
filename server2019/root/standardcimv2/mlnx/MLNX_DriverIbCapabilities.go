// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 9/18/2020 using wmigen
//      Source root.StandardCimv2.mlnx
//////////////////////////////////////////////
package mlnx

import (
	"github.com/microsoft/wmi/pkg/base/query"
	"github.com/microsoft/wmi/pkg/errors"
	cim "github.com/microsoft/wmi/pkg/wmiinstance"
	"reflect"
)

// MLNX_DriverIbCapabilities struct
type MLNX_DriverIbCapabilities struct {
	*MLNX_DriverCapabilities

	//
	DebugFlags_Max uint32

	//
	DebugFlags_Min uint32

	//
	IbalDebugFlags_Max uint32

	//
	IbalDebugFlags_Min uint32

	//
	IbalDebugLevel_Max uint32

	//
	IbalDebugLevel_Min uint32
}

func NewMLNX_DriverIbCapabilitiesEx1(instance *cim.WmiInstance) (newInstance *MLNX_DriverIbCapabilities, err error) {
	tmp, err := NewMLNX_DriverCapabilitiesEx1(instance)

	if err != nil {
		return
	}
	newInstance = &MLNX_DriverIbCapabilities{
		MLNX_DriverCapabilities: tmp,
	}
	return
}

func NewMLNX_DriverIbCapabilitiesEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *MLNX_DriverIbCapabilities, err error) {
	tmp, err := NewMLNX_DriverCapabilitiesEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &MLNX_DriverIbCapabilities{
		MLNX_DriverCapabilities: tmp,
	}
	return
}

// SetDebugFlags_Max sets the value of DebugFlags_Max for the instance
func (instance *MLNX_DriverIbCapabilities) SetPropertyDebugFlags_Max(value uint32) (err error) {
	return instance.SetProperty("DebugFlags_Max", (value))
}

// GetDebugFlags_Max gets the value of DebugFlags_Max for the instance
func (instance *MLNX_DriverIbCapabilities) GetPropertyDebugFlags_Max() (value uint32, err error) {
	retValue, err := instance.GetProperty("DebugFlags_Max")
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

// SetDebugFlags_Min sets the value of DebugFlags_Min for the instance
func (instance *MLNX_DriverIbCapabilities) SetPropertyDebugFlags_Min(value uint32) (err error) {
	return instance.SetProperty("DebugFlags_Min", (value))
}

// GetDebugFlags_Min gets the value of DebugFlags_Min for the instance
func (instance *MLNX_DriverIbCapabilities) GetPropertyDebugFlags_Min() (value uint32, err error) {
	retValue, err := instance.GetProperty("DebugFlags_Min")
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

// SetIbalDebugFlags_Max sets the value of IbalDebugFlags_Max for the instance
func (instance *MLNX_DriverIbCapabilities) SetPropertyIbalDebugFlags_Max(value uint32) (err error) {
	return instance.SetProperty("IbalDebugFlags_Max", (value))
}

// GetIbalDebugFlags_Max gets the value of IbalDebugFlags_Max for the instance
func (instance *MLNX_DriverIbCapabilities) GetPropertyIbalDebugFlags_Max() (value uint32, err error) {
	retValue, err := instance.GetProperty("IbalDebugFlags_Max")
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

// SetIbalDebugFlags_Min sets the value of IbalDebugFlags_Min for the instance
func (instance *MLNX_DriverIbCapabilities) SetPropertyIbalDebugFlags_Min(value uint32) (err error) {
	return instance.SetProperty("IbalDebugFlags_Min", (value))
}

// GetIbalDebugFlags_Min gets the value of IbalDebugFlags_Min for the instance
func (instance *MLNX_DriverIbCapabilities) GetPropertyIbalDebugFlags_Min() (value uint32, err error) {
	retValue, err := instance.GetProperty("IbalDebugFlags_Min")
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

// SetIbalDebugLevel_Max sets the value of IbalDebugLevel_Max for the instance
func (instance *MLNX_DriverIbCapabilities) SetPropertyIbalDebugLevel_Max(value uint32) (err error) {
	return instance.SetProperty("IbalDebugLevel_Max", (value))
}

// GetIbalDebugLevel_Max gets the value of IbalDebugLevel_Max for the instance
func (instance *MLNX_DriverIbCapabilities) GetPropertyIbalDebugLevel_Max() (value uint32, err error) {
	retValue, err := instance.GetProperty("IbalDebugLevel_Max")
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

// SetIbalDebugLevel_Min sets the value of IbalDebugLevel_Min for the instance
func (instance *MLNX_DriverIbCapabilities) SetPropertyIbalDebugLevel_Min(value uint32) (err error) {
	return instance.SetProperty("IbalDebugLevel_Min", (value))
}

// GetIbalDebugLevel_Min gets the value of IbalDebugLevel_Min for the instance
func (instance *MLNX_DriverIbCapabilities) GetPropertyIbalDebugLevel_Min() (value uint32, err error) {
	retValue, err := instance.GetProperty("IbalDebugLevel_Min")
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
