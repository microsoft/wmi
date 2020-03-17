// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 3/16/2020 using wmigen
//      Source root.InventoryLogging
//////////////////////////////////////////////
package inventorylogging

// MsftSil_UalAccess struct
type MsftSil_UalAccess struct {
	MsftSil_Data

	//
	ProductName string

	//
	RoleGuid string

	//
	RoleName string

	//
	SampleDate string

	//
	UniqueDeviceAccessCount uint32

	//
	UniqueUserAccessCount uint32
}

// SetProductName sets the value of ProductName for the instance
func (instance *MsftSil_UalAccess) SetPropertyProductName(value string) (err error) {
	return instance.SetProperty("ProductName", value)
}

// GetProductName gets the value of ProductName for the instance
func (instance *MsftSil_UalAccess) GetPropertyProductName() (value string, err error) {
	retValue, err := instance.GetProperty("ProductName")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetRoleGuid sets the value of RoleGuid for the instance
func (instance *MsftSil_UalAccess) SetPropertyRoleGuid(value string) (err error) {
	return instance.SetProperty("RoleGuid", value)
}

// GetRoleGuid gets the value of RoleGuid for the instance
func (instance *MsftSil_UalAccess) GetPropertyRoleGuid() (value string, err error) {
	retValue, err := instance.GetProperty("RoleGuid")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetRoleName sets the value of RoleName for the instance
func (instance *MsftSil_UalAccess) SetPropertyRoleName(value string) (err error) {
	return instance.SetProperty("RoleName", value)
}

// GetRoleName gets the value of RoleName for the instance
func (instance *MsftSil_UalAccess) GetPropertyRoleName() (value string, err error) {
	retValue, err := instance.GetProperty("RoleName")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetSampleDate sets the value of SampleDate for the instance
func (instance *MsftSil_UalAccess) SetPropertySampleDate(value string) (err error) {
	return instance.SetProperty("SampleDate", value)
}

// GetSampleDate gets the value of SampleDate for the instance
func (instance *MsftSil_UalAccess) GetPropertySampleDate() (value string, err error) {
	retValue, err := instance.GetProperty("SampleDate")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetUniqueDeviceAccessCount sets the value of UniqueDeviceAccessCount for the instance
func (instance *MsftSil_UalAccess) SetPropertyUniqueDeviceAccessCount(value uint32) (err error) {
	return instance.SetProperty("UniqueDeviceAccessCount", value)
}

// GetUniqueDeviceAccessCount gets the value of UniqueDeviceAccessCount for the instance
func (instance *MsftSil_UalAccess) GetPropertyUniqueDeviceAccessCount() (value uint32, err error) {
	retValue, err := instance.GetProperty("UniqueDeviceAccessCount")
	if err != nil {
		return
	}
	value, ok := retValue.(uint32)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetUniqueUserAccessCount sets the value of UniqueUserAccessCount for the instance
func (instance *MsftSil_UalAccess) SetPropertyUniqueUserAccessCount(value uint32) (err error) {
	return instance.SetProperty("UniqueUserAccessCount", value)
}

// GetUniqueUserAccessCount gets the value of UniqueUserAccessCount for the instance
func (instance *MsftSil_UalAccess) GetPropertyUniqueUserAccessCount() (value uint32, err error) {
	retValue, err := instance.GetProperty("UniqueUserAccessCount")
	if err != nil {
		return
	}
	value, ok := retValue.(uint32)
	if !ok {
		// TODO: Set an error
	}
	return
}
