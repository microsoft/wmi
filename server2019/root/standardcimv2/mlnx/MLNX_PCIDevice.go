// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 3/16/2020 using wmigen
//      Source root.StandardCimv2.mlnx
//////////////////////////////////////////////
package mlnx

// MLNX_PCIDevice struct
type MLNX_PCIDevice struct {
	CIM_PCIDevice

	//
	BusType string

	//
	DriverVersion string

	//
	FirmwareVersion string

	//
	PartNumber string
}

// SetBusType sets the value of BusType for the instance
func (instance *MLNX_PCIDevice) SetPropertyBusType(value string) (err error) {
	return instance.SetProperty("BusType", value)
}

// GetBusType gets the value of BusType for the instance
func (instance *MLNX_PCIDevice) GetPropertyBusType() (value string, err error) {
	retValue, err := instance.GetProperty("BusType")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetDriverVersion sets the value of DriverVersion for the instance
func (instance *MLNX_PCIDevice) SetPropertyDriverVersion(value string) (err error) {
	return instance.SetProperty("DriverVersion", value)
}

// GetDriverVersion gets the value of DriverVersion for the instance
func (instance *MLNX_PCIDevice) GetPropertyDriverVersion() (value string, err error) {
	retValue, err := instance.GetProperty("DriverVersion")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetFirmwareVersion sets the value of FirmwareVersion for the instance
func (instance *MLNX_PCIDevice) SetPropertyFirmwareVersion(value string) (err error) {
	return instance.SetProperty("FirmwareVersion", value)
}

// GetFirmwareVersion gets the value of FirmwareVersion for the instance
func (instance *MLNX_PCIDevice) GetPropertyFirmwareVersion() (value string, err error) {
	retValue, err := instance.GetProperty("FirmwareVersion")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetPartNumber sets the value of PartNumber for the instance
func (instance *MLNX_PCIDevice) SetPropertyPartNumber(value string) (err error) {
	return instance.SetProperty("PartNumber", value)
}

// GetPartNumber gets the value of PartNumber for the instance
func (instance *MLNX_PCIDevice) GetPropertyPartNumber() (value string, err error) {
	retValue, err := instance.GetProperty("PartNumber")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
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
