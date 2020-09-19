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

// MLNX_DriverIpOverIbCapabilities struct
type MLNX_DriverIpOverIbCapabilities struct {
	*MLNX_DriverCapabilities

	//
	ModeFlags_Max uint32

	//
	ModeFlags_Min uint32

	//
	NdkDebugFlags_Max uint32

	//
	NdkDebugFlags_Min uint32

	//
	NdkDebugLevel_Max uint32

	//
	NdkDebugLevel_Min uint32
}

func NewMLNX_DriverIpOverIbCapabilitiesEx1(instance *cim.WmiInstance) (newInstance *MLNX_DriverIpOverIbCapabilities, err error) {
	tmp, err := NewMLNX_DriverCapabilitiesEx1(instance)

	if err != nil {
		return
	}
	newInstance = &MLNX_DriverIpOverIbCapabilities{
		MLNX_DriverCapabilities: tmp,
	}
	return
}

func NewMLNX_DriverIpOverIbCapabilitiesEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *MLNX_DriverIpOverIbCapabilities, err error) {
	tmp, err := NewMLNX_DriverCapabilitiesEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &MLNX_DriverIpOverIbCapabilities{
		MLNX_DriverCapabilities: tmp,
	}
	return
}

// SetModeFlags_Max sets the value of ModeFlags_Max for the instance
func (instance *MLNX_DriverIpOverIbCapabilities) SetPropertyModeFlags_Max(value uint32) (err error) {
	return instance.SetProperty("ModeFlags_Max", (value))
}

// GetModeFlags_Max gets the value of ModeFlags_Max for the instance
func (instance *MLNX_DriverIpOverIbCapabilities) GetPropertyModeFlags_Max() (value uint32, err error) {
	retValue, err := instance.GetProperty("ModeFlags_Max")
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

// SetModeFlags_Min sets the value of ModeFlags_Min for the instance
func (instance *MLNX_DriverIpOverIbCapabilities) SetPropertyModeFlags_Min(value uint32) (err error) {
	return instance.SetProperty("ModeFlags_Min", (value))
}

// GetModeFlags_Min gets the value of ModeFlags_Min for the instance
func (instance *MLNX_DriverIpOverIbCapabilities) GetPropertyModeFlags_Min() (value uint32, err error) {
	retValue, err := instance.GetProperty("ModeFlags_Min")
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

// SetNdkDebugFlags_Max sets the value of NdkDebugFlags_Max for the instance
func (instance *MLNX_DriverIpOverIbCapabilities) SetPropertyNdkDebugFlags_Max(value uint32) (err error) {
	return instance.SetProperty("NdkDebugFlags_Max", (value))
}

// GetNdkDebugFlags_Max gets the value of NdkDebugFlags_Max for the instance
func (instance *MLNX_DriverIpOverIbCapabilities) GetPropertyNdkDebugFlags_Max() (value uint32, err error) {
	retValue, err := instance.GetProperty("NdkDebugFlags_Max")
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

// SetNdkDebugFlags_Min sets the value of NdkDebugFlags_Min for the instance
func (instance *MLNX_DriverIpOverIbCapabilities) SetPropertyNdkDebugFlags_Min(value uint32) (err error) {
	return instance.SetProperty("NdkDebugFlags_Min", (value))
}

// GetNdkDebugFlags_Min gets the value of NdkDebugFlags_Min for the instance
func (instance *MLNX_DriverIpOverIbCapabilities) GetPropertyNdkDebugFlags_Min() (value uint32, err error) {
	retValue, err := instance.GetProperty("NdkDebugFlags_Min")
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

// SetNdkDebugLevel_Max sets the value of NdkDebugLevel_Max for the instance
func (instance *MLNX_DriverIpOverIbCapabilities) SetPropertyNdkDebugLevel_Max(value uint32) (err error) {
	return instance.SetProperty("NdkDebugLevel_Max", (value))
}

// GetNdkDebugLevel_Max gets the value of NdkDebugLevel_Max for the instance
func (instance *MLNX_DriverIpOverIbCapabilities) GetPropertyNdkDebugLevel_Max() (value uint32, err error) {
	retValue, err := instance.GetProperty("NdkDebugLevel_Max")
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

// SetNdkDebugLevel_Min sets the value of NdkDebugLevel_Min for the instance
func (instance *MLNX_DriverIpOverIbCapabilities) SetPropertyNdkDebugLevel_Min(value uint32) (err error) {
	return instance.SetProperty("NdkDebugLevel_Min", (value))
}

// GetNdkDebugLevel_Min gets the value of NdkDebugLevel_Min for the instance
func (instance *MLNX_DriverIpOverIbCapabilities) GetPropertyNdkDebugLevel_Min() (value uint32, err error) {
	retValue, err := instance.GetProperty("NdkDebugLevel_Min")
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
