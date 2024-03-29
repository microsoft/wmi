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

// MLNX_PCIDevice struct
type MLNX_PCIDevice struct {
	*CIM_PCIDevice

	//
	BusType string

	//
	DriverVersion string

	//
	FirmwareVersion string

	//
	PartNumber string
}

func NewMLNX_PCIDeviceEx1(instance *cim.WmiInstance) (newInstance *MLNX_PCIDevice, err error) {
	tmp, err := NewCIM_PCIDeviceEx1(instance)

	if err != nil {
		return
	}
	newInstance = &MLNX_PCIDevice{
		CIM_PCIDevice: tmp,
	}
	return
}

func NewMLNX_PCIDeviceEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *MLNX_PCIDevice, err error) {
	tmp, err := NewCIM_PCIDeviceEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &MLNX_PCIDevice{
		CIM_PCIDevice: tmp,
	}
	return
}

// SetBusType sets the value of BusType for the instance
func (instance *MLNX_PCIDevice) SetPropertyBusType(value string) (err error) {
	return instance.SetProperty("BusType", (value))
}

// GetBusType gets the value of BusType for the instance
func (instance *MLNX_PCIDevice) GetPropertyBusType() (value string, err error) {
	retValue, err := instance.GetProperty("BusType")
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

// SetDriverVersion sets the value of DriverVersion for the instance
func (instance *MLNX_PCIDevice) SetPropertyDriverVersion(value string) (err error) {
	return instance.SetProperty("DriverVersion", (value))
}

// GetDriverVersion gets the value of DriverVersion for the instance
func (instance *MLNX_PCIDevice) GetPropertyDriverVersion() (value string, err error) {
	retValue, err := instance.GetProperty("DriverVersion")
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

// SetFirmwareVersion sets the value of FirmwareVersion for the instance
func (instance *MLNX_PCIDevice) SetPropertyFirmwareVersion(value string) (err error) {
	return instance.SetProperty("FirmwareVersion", (value))
}

// GetFirmwareVersion gets the value of FirmwareVersion for the instance
func (instance *MLNX_PCIDevice) GetPropertyFirmwareVersion() (value string, err error) {
	retValue, err := instance.GetProperty("FirmwareVersion")
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

// SetPartNumber sets the value of PartNumber for the instance
func (instance *MLNX_PCIDevice) SetPropertyPartNumber(value string) (err error) {
	return instance.SetProperty("PartNumber", (value))
}

// GetPartNumber gets the value of PartNumber for the instance
func (instance *MLNX_PCIDevice) GetPropertyPartNumber() (value string, err error) {
	retValue, err := instance.GetProperty("PartNumber")
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

//

// <param name="ReturnValue" type="uint32 "></param>
func (instance *MLNX_PCIDevice) Disable() (result uint32, err error) {
	retVal, err := instance.InvokeMethodWithReturn("Disable")
	if err != nil {
		return
	}
	result = uint32(retVal)
	return

}

//

// <param name="ReturnValue" type="uint32 "></param>
func (instance *MLNX_PCIDevice) Enable() (result uint32, err error) {
	retVal, err := instance.InvokeMethodWithReturn("Enable")
	if err != nil {
		return
	}
	result = uint32(retVal)
	return

}

//

// <param name="ReturnValue" type="uint32 "></param>
func (instance *MLNX_PCIDevice) Restart() (result uint32, err error) {
	retVal, err := instance.InvokeMethodWithReturn("Restart")
	if err != nil {
		return
	}
	result = uint32(retVal)
	return

}
