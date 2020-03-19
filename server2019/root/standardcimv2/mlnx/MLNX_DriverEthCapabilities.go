// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 3/19/2020 using wmigen
//      Source root.StandardCimv2.mlnx
//////////////////////////////////////////////
package mlnx

import (
	"github.com/microsoft/wmi/pkg/base/query"
	cim "github.com/microsoft/wmi/pkg/wmiinstance"
)

// MLNX_DriverEthCapabilities struct
type MLNX_DriverEthCapabilities struct {
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

func NewMLNX_DriverEthCapabilitiesEx1(instance *cim.WmiInstance) (newInstance *MLNX_DriverEthCapabilities, err error) {
	tmp, err := NewMLNX_DriverCapabilitiesEx1(instance)

	if err != nil {
		return
	}
	newInstance = &MLNX_DriverEthCapabilities{
		MLNX_DriverCapabilities: tmp,
	}
	return
}

func NewMLNX_DriverEthCapabilitiesEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *MLNX_DriverEthCapabilities, err error) {
	tmp, err := NewMLNX_DriverCapabilitiesEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &MLNX_DriverEthCapabilities{
		MLNX_DriverCapabilities: tmp,
	}
	return
}

// SetModeFlags_Max sets the value of ModeFlags_Max for the instance
func (instance *MLNX_DriverEthCapabilities) SetPropertyModeFlags_Max(value uint32) (err error) {
	return instance.SetProperty("ModeFlags_Max", value)
}

// GetModeFlags_Max gets the value of ModeFlags_Max for the instance
func (instance *MLNX_DriverEthCapabilities) GetPropertyModeFlags_Max() (value uint32, err error) {
	retValue, err := instance.GetProperty("ModeFlags_Max")
	if err != nil {
		return
	}
	value, ok := retValue.(uint32)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetModeFlags_Min sets the value of ModeFlags_Min for the instance
func (instance *MLNX_DriverEthCapabilities) SetPropertyModeFlags_Min(value uint32) (err error) {
	return instance.SetProperty("ModeFlags_Min", value)
}

// GetModeFlags_Min gets the value of ModeFlags_Min for the instance
func (instance *MLNX_DriverEthCapabilities) GetPropertyModeFlags_Min() (value uint32, err error) {
	retValue, err := instance.GetProperty("ModeFlags_Min")
	if err != nil {
		return
	}
	value, ok := retValue.(uint32)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetNdkDebugFlags_Max sets the value of NdkDebugFlags_Max for the instance
func (instance *MLNX_DriverEthCapabilities) SetPropertyNdkDebugFlags_Max(value uint32) (err error) {
	return instance.SetProperty("NdkDebugFlags_Max", value)
}

// GetNdkDebugFlags_Max gets the value of NdkDebugFlags_Max for the instance
func (instance *MLNX_DriverEthCapabilities) GetPropertyNdkDebugFlags_Max() (value uint32, err error) {
	retValue, err := instance.GetProperty("NdkDebugFlags_Max")
	if err != nil {
		return
	}
	value, ok := retValue.(uint32)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetNdkDebugFlags_Min sets the value of NdkDebugFlags_Min for the instance
func (instance *MLNX_DriverEthCapabilities) SetPropertyNdkDebugFlags_Min(value uint32) (err error) {
	return instance.SetProperty("NdkDebugFlags_Min", value)
}

// GetNdkDebugFlags_Min gets the value of NdkDebugFlags_Min for the instance
func (instance *MLNX_DriverEthCapabilities) GetPropertyNdkDebugFlags_Min() (value uint32, err error) {
	retValue, err := instance.GetProperty("NdkDebugFlags_Min")
	if err != nil {
		return
	}
	value, ok := retValue.(uint32)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetNdkDebugLevel_Max sets the value of NdkDebugLevel_Max for the instance
func (instance *MLNX_DriverEthCapabilities) SetPropertyNdkDebugLevel_Max(value uint32) (err error) {
	return instance.SetProperty("NdkDebugLevel_Max", value)
}

// GetNdkDebugLevel_Max gets the value of NdkDebugLevel_Max for the instance
func (instance *MLNX_DriverEthCapabilities) GetPropertyNdkDebugLevel_Max() (value uint32, err error) {
	retValue, err := instance.GetProperty("NdkDebugLevel_Max")
	if err != nil {
		return
	}
	value, ok := retValue.(uint32)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetNdkDebugLevel_Min sets the value of NdkDebugLevel_Min for the instance
func (instance *MLNX_DriverEthCapabilities) SetPropertyNdkDebugLevel_Min(value uint32) (err error) {
	return instance.SetProperty("NdkDebugLevel_Min", value)
}

// GetNdkDebugLevel_Min gets the value of NdkDebugLevel_Min for the instance
func (instance *MLNX_DriverEthCapabilities) GetPropertyNdkDebugLevel_Min() (value uint32, err error) {
	retValue, err := instance.GetProperty("NdkDebugLevel_Min")
	if err != nil {
		return
	}
	value, ok := retValue.(uint32)
	if !ok {
		// TODO: Set an error
	}
	return
}
