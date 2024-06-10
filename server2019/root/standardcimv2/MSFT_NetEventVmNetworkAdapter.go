// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 9/18/2020 using wmigen
//      Source root.StandardCimv2
//////////////////////////////////////////////
package standardcimv2

import (
	"github.com/microsoft/wmi/pkg/base/query"
	"github.com/microsoft/wmi/pkg/errors"
	cim "github.com/microsoft/wmi/pkg/wmiinstance"
	"reflect"
)

// MSFT_NetEventVmNetworkAdapter struct
type MSFT_NetEventVmNetworkAdapter struct {
	*MSFT_NetEventPacketCaptureTarget

	//
	MacAddress string

	//
	PortName string

	//
	SwitchName string

	//
	VMId string

	//
	VMName string
}

func NewMSFT_NetEventVmNetworkAdapterEx1(instance *cim.WmiInstance) (newInstance *MSFT_NetEventVmNetworkAdapter, err error) {
	tmp, err := NewMSFT_NetEventPacketCaptureTargetEx1(instance)

	if err != nil {
		return
	}
	newInstance = &MSFT_NetEventVmNetworkAdapter{
		MSFT_NetEventPacketCaptureTarget: tmp,
	}
	return
}

func NewMSFT_NetEventVmNetworkAdapterEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *MSFT_NetEventVmNetworkAdapter, err error) {
	tmp, err := NewMSFT_NetEventPacketCaptureTargetEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &MSFT_NetEventVmNetworkAdapter{
		MSFT_NetEventPacketCaptureTarget: tmp,
	}
	return
}

// SetMacAddress sets the value of MacAddress for the instance
func (instance *MSFT_NetEventVmNetworkAdapter) SetPropertyMacAddress(value string) (err error) {
	return instance.SetProperty("MacAddress", (value))
}

// GetMacAddress gets the value of MacAddress for the instance
func (instance *MSFT_NetEventVmNetworkAdapter) GetPropertyMacAddress() (value string, err error) {
	retValue, err := instance.GetProperty("MacAddress")
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

// SetPortName sets the value of PortName for the instance
func (instance *MSFT_NetEventVmNetworkAdapter) SetPropertyPortName(value string) (err error) {
	return instance.SetProperty("PortName", (value))
}

// GetPortName gets the value of PortName for the instance
func (instance *MSFT_NetEventVmNetworkAdapter) GetPropertyPortName() (value string, err error) {
	retValue, err := instance.GetProperty("PortName")
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

// SetSwitchName sets the value of SwitchName for the instance
func (instance *MSFT_NetEventVmNetworkAdapter) SetPropertySwitchName(value string) (err error) {
	return instance.SetProperty("SwitchName", (value))
}

// GetSwitchName gets the value of SwitchName for the instance
func (instance *MSFT_NetEventVmNetworkAdapter) GetPropertySwitchName() (value string, err error) {
	retValue, err := instance.GetProperty("SwitchName")
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

// SetVMId sets the value of VMId for the instance
func (instance *MSFT_NetEventVmNetworkAdapter) SetPropertyVMId(value string) (err error) {
	return instance.SetProperty("VMId", (value))
}

// GetVMId gets the value of VMId for the instance
func (instance *MSFT_NetEventVmNetworkAdapter) GetPropertyVMId() (value string, err error) {
	retValue, err := instance.GetProperty("VMId")
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

// SetVMName sets the value of VMName for the instance
func (instance *MSFT_NetEventVmNetworkAdapter) SetPropertyVMName(value string) (err error) {
	return instance.SetProperty("VMName", (value))
}

// GetVMName gets the value of VMName for the instance
func (instance *MSFT_NetEventVmNetworkAdapter) GetPropertyVMName() (value string, err error) {
	retValue, err := instance.GetProperty("VMName")
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
