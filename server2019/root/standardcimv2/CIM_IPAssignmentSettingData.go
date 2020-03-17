// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 3/16/2020 using wmigen
//      Source root.StandardCimv2
//////////////////////////////////////////////
package standardcimv2

// CIM_IPAssignmentSettingData struct
type CIM_IPAssignmentSettingData struct {
	CIM_SettingData

	// 689
	AddressOrigin IPAssignmentSettingData_AddressOrigin

	// 693
	ProtocolIFType IPAssignmentSettingData_ProtocolIFType
}

// SetAddressOrigin sets the value of AddressOrigin for the instance
func (instance *CIM_IPAssignmentSettingData) SetPropertyAddressOrigin(value IPAssignmentSettingData_AddressOrigin) (err error) {
	return instance.SetProperty("AddressOrigin", value)
}

// GetAddressOrigin gets the value of AddressOrigin for the instance
func (instance *CIM_IPAssignmentSettingData) GetPropertyAddressOrigin() (value IPAssignmentSettingData_AddressOrigin, err error) {
	retValue, err := instance.GetProperty("AddressOrigin")
	if err != nil {
		return
	}
	value, ok := retValue.(IPAssignmentSettingData_AddressOrigin)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetProtocolIFType sets the value of ProtocolIFType for the instance
func (instance *CIM_IPAssignmentSettingData) SetPropertyProtocolIFType(value IPAssignmentSettingData_ProtocolIFType) (err error) {
	return instance.SetProperty("ProtocolIFType", value)
}

// GetProtocolIFType gets the value of ProtocolIFType for the instance
func (instance *CIM_IPAssignmentSettingData) GetPropertyProtocolIFType() (value IPAssignmentSettingData_ProtocolIFType, err error) {
	retValue, err := instance.GetProperty("ProtocolIFType")
	if err != nil {
		return
	}
	value, ok := retValue.(IPAssignmentSettingData_ProtocolIFType)
	if !ok {
		// TODO: Set an error
	}
	return
}
