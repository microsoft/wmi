// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 3/16/2020 using wmigen
//      Source root.virtualization.v2
//////////////////////////////////////////////
package v2

// Msvm_MountedStorageImage struct
type Msvm_MountedStorageImage struct {
	CIM_LogicalElement

	//
	Access uint16

	//
	Lun uint8

	//
	PathId uint8

	//
	PnpDevicePath string

	//
	PortNumber uint8

	//
	TargetId uint8

	//
	Type uint16
}

// SetAccess sets the value of Access for the instance
func (instance *Msvm_MountedStorageImage) SetPropertyAccess(value uint16) (err error) {
	return instance.SetProperty("Access", value)
}

// GetAccess gets the value of Access for the instance
func (instance *Msvm_MountedStorageImage) GetPropertyAccess() (value uint16, err error) {
	retValue, err := instance.GetProperty("Access")
	if err != nil {
		return
	}
	value, ok := retValue.(uint16)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetLun sets the value of Lun for the instance
func (instance *Msvm_MountedStorageImage) SetPropertyLun(value uint8) (err error) {
	return instance.SetProperty("Lun", value)
}

// GetLun gets the value of Lun for the instance
func (instance *Msvm_MountedStorageImage) GetPropertyLun() (value uint8, err error) {
	retValue, err := instance.GetProperty("Lun")
	if err != nil {
		return
	}
	value, ok := retValue.(uint8)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetPathId sets the value of PathId for the instance
func (instance *Msvm_MountedStorageImage) SetPropertyPathId(value uint8) (err error) {
	return instance.SetProperty("PathId", value)
}

// GetPathId gets the value of PathId for the instance
func (instance *Msvm_MountedStorageImage) GetPropertyPathId() (value uint8, err error) {
	retValue, err := instance.GetProperty("PathId")
	if err != nil {
		return
	}
	value, ok := retValue.(uint8)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetPnpDevicePath sets the value of PnpDevicePath for the instance
func (instance *Msvm_MountedStorageImage) SetPropertyPnpDevicePath(value string) (err error) {
	return instance.SetProperty("PnpDevicePath", value)
}

// GetPnpDevicePath gets the value of PnpDevicePath for the instance
func (instance *Msvm_MountedStorageImage) GetPropertyPnpDevicePath() (value string, err error) {
	retValue, err := instance.GetProperty("PnpDevicePath")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetPortNumber sets the value of PortNumber for the instance
func (instance *Msvm_MountedStorageImage) SetPropertyPortNumber(value uint8) (err error) {
	return instance.SetProperty("PortNumber", value)
}

// GetPortNumber gets the value of PortNumber for the instance
func (instance *Msvm_MountedStorageImage) GetPropertyPortNumber() (value uint8, err error) {
	retValue, err := instance.GetProperty("PortNumber")
	if err != nil {
		return
	}
	value, ok := retValue.(uint8)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetTargetId sets the value of TargetId for the instance
func (instance *Msvm_MountedStorageImage) SetPropertyTargetId(value uint8) (err error) {
	return instance.SetProperty("TargetId", value)
}

// GetTargetId gets the value of TargetId for the instance
func (instance *Msvm_MountedStorageImage) GetPropertyTargetId() (value uint8, err error) {
	retValue, err := instance.GetProperty("TargetId")
	if err != nil {
		return
	}
	value, ok := retValue.(uint8)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetType sets the value of Type for the instance
func (instance *Msvm_MountedStorageImage) SetPropertyType(value uint16) (err error) {
	return instance.SetProperty("Type", value)
}

// GetType gets the value of Type for the instance
func (instance *Msvm_MountedStorageImage) GetPropertyType() (value uint16, err error) {
	retValue, err := instance.GetProperty("Type")
	if err != nil {
		return
	}
	value, ok := retValue.(uint16)
	if !ok {
		// TODO: Set an error
	}
	return
}

//

// <param name="ReturnValue" type="uint32 "></param>
func (instance *Msvm_MountedStorageImage) DetachVirtualHardDisk() (result uint32, err error) {
	retVal, err := instance.InvokeMethodWithReturn("DetachVirtualHardDisk")
	if err != nil {
		return
	}
	result = uint32(retVal)
	return

}
