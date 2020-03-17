// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 3/16/2020 using wmigen
//      Source root.StandardCimv2.mlnx
//////////////////////////////////////////////
package mlnx

// MLNX_NetAdapter struct
type MLNX_NetAdapter struct {
	CIM_EthernetPort

	//
	DroplessMode uint8

	//
	PciLocation string
}

// SetDroplessMode sets the value of DroplessMode for the instance
func (instance *MLNX_NetAdapter) SetPropertyDroplessMode(value uint8) (err error) {
	return instance.SetProperty("DroplessMode", value)
}

// GetDroplessMode gets the value of DroplessMode for the instance
func (instance *MLNX_NetAdapter) GetPropertyDroplessMode() (value uint8, err error) {
	retValue, err := instance.GetProperty("DroplessMode")
	if err != nil {
		return
	}
	value, ok := retValue.(uint8)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetPciLocation sets the value of PciLocation for the instance
func (instance *MLNX_NetAdapter) SetPropertyPciLocation(value string) (err error) {
	return instance.SetProperty("PciLocation", value)
}

// GetPciLocation gets the value of PciLocation for the instance
func (instance *MLNX_NetAdapter) GetPropertyPciLocation() (value string, err error) {
	retValue, err := instance.GetProperty("PciLocation")
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

// <param name="ReturnValue" type="uint8 "></param>
func (instance *MLNX_NetAdapter) Disable() (result uint8, err error) {
	retVal, err := instance.InvokeMethodWithReturn("Disable")
	if err != nil {
		return
	}
	result = uint8(retVal)
	return

}

//

// <param name="ReturnValue" type="uint8 "></param>
func (instance *MLNX_NetAdapter) Enable() (result uint8, err error) {
	retVal, err := instance.InvokeMethodWithReturn("Enable")
	if err != nil {
		return
	}
	result = uint8(retVal)
	return

}

//

// <param name="ReturnValue" type="uint8 "></param>
func (instance *MLNX_NetAdapter) Restart() (result uint8, err error) {
	retVal, err := instance.InvokeMethodWithReturn("Restart")
	if err != nil {
		return
	}
	result = uint8(retVal)
	return

}
