// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 3/16/2020 using wmigen
//      Source root.InventoryLogging
//////////////////////////////////////////////
package inventorylogging

// MsftSil_ComputerIdentity struct
type MsftSil_ComputerIdentity struct {
	MsftSil_Data

	//
	HypervisorHostName string

	//
	ID string

	//
	Name string

	//
	UUID string

	//
	VMGUID string
}

// SetHypervisorHostName sets the value of HypervisorHostName for the instance
func (instance *MsftSil_ComputerIdentity) SetPropertyHypervisorHostName(value string) (err error) {
	return instance.SetProperty("HypervisorHostName", value)
}

// GetHypervisorHostName gets the value of HypervisorHostName for the instance
func (instance *MsftSil_ComputerIdentity) GetPropertyHypervisorHostName() (value string, err error) {
	retValue, err := instance.GetProperty("HypervisorHostName")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetID sets the value of ID for the instance
func (instance *MsftSil_ComputerIdentity) SetPropertyID(value string) (err error) {
	return instance.SetProperty("ID", value)
}

// GetID gets the value of ID for the instance
func (instance *MsftSil_ComputerIdentity) GetPropertyID() (value string, err error) {
	retValue, err := instance.GetProperty("ID")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetName sets the value of Name for the instance
func (instance *MsftSil_ComputerIdentity) SetPropertyName(value string) (err error) {
	return instance.SetProperty("Name", value)
}

// GetName gets the value of Name for the instance
func (instance *MsftSil_ComputerIdentity) GetPropertyName() (value string, err error) {
	retValue, err := instance.GetProperty("Name")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetUUID sets the value of UUID for the instance
func (instance *MsftSil_ComputerIdentity) SetPropertyUUID(value string) (err error) {
	return instance.SetProperty("UUID", value)
}

// GetUUID gets the value of UUID for the instance
func (instance *MsftSil_ComputerIdentity) GetPropertyUUID() (value string, err error) {
	retValue, err := instance.GetProperty("UUID")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetVMGUID sets the value of VMGUID for the instance
func (instance *MsftSil_ComputerIdentity) SetPropertyVMGUID(value string) (err error) {
	return instance.SetProperty("VMGUID", value)
}

// GetVMGUID gets the value of VMGUID for the instance
func (instance *MsftSil_ComputerIdentity) GetPropertyVMGUID() (value string, err error) {
	retValue, err := instance.GetProperty("VMGUID")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}
