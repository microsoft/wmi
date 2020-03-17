// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 3/16/2020 using wmigen
//      Source root.StandardCimv2.mlnx
//////////////////////////////////////////////
package mlnx

// CIM_ElementSoftwareIdentity struct
type CIM_ElementSoftwareIdentity struct {
	CIM_Dependency

	//
	ElementSoftwareStatus []ElementSoftwareIdentity_ElementSoftwareStatus

	//
	OtherUpgradeCondition string

	//
	UpgradeCondition ElementSoftwareIdentity_UpgradeCondition
}

// SetElementSoftwareStatus sets the value of ElementSoftwareStatus for the instance
func (instance *CIM_ElementSoftwareIdentity) SetPropertyElementSoftwareStatus(value []ElementSoftwareIdentity_ElementSoftwareStatus) (err error) {
	return instance.SetProperty("ElementSoftwareStatus", value)
}

// GetElementSoftwareStatus gets the value of ElementSoftwareStatus for the instance
func (instance *CIM_ElementSoftwareIdentity) GetPropertyElementSoftwareStatus() (value []ElementSoftwareIdentity_ElementSoftwareStatus, err error) {
	retValue, err := instance.GetProperty("ElementSoftwareStatus")
	if err != nil {
		return
	}
	value, ok := retValue.([]ElementSoftwareIdentity_ElementSoftwareStatus)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetOtherUpgradeCondition sets the value of OtherUpgradeCondition for the instance
func (instance *CIM_ElementSoftwareIdentity) SetPropertyOtherUpgradeCondition(value string) (err error) {
	return instance.SetProperty("OtherUpgradeCondition", value)
}

// GetOtherUpgradeCondition gets the value of OtherUpgradeCondition for the instance
func (instance *CIM_ElementSoftwareIdentity) GetPropertyOtherUpgradeCondition() (value string, err error) {
	retValue, err := instance.GetProperty("OtherUpgradeCondition")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetUpgradeCondition sets the value of UpgradeCondition for the instance
func (instance *CIM_ElementSoftwareIdentity) SetPropertyUpgradeCondition(value ElementSoftwareIdentity_UpgradeCondition) (err error) {
	return instance.SetProperty("UpgradeCondition", value)
}

// GetUpgradeCondition gets the value of UpgradeCondition for the instance
func (instance *CIM_ElementSoftwareIdentity) GetPropertyUpgradeCondition() (value ElementSoftwareIdentity_UpgradeCondition, err error) {
	retValue, err := instance.GetProperty("UpgradeCondition")
	if err != nil {
		return
	}
	value, ok := retValue.(ElementSoftwareIdentity_UpgradeCondition)
	if !ok {
		// TODO: Set an error
	}
	return
}
