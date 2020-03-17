// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 3/16/2020 using wmigen
//      Source root.Hardware
//////////////////////////////////////////////
package hardware

// CIM_Privilege struct
type CIM_Privilege struct {
	CIM_ManagedElement

	//
	Activities []uint16

	//
	ActivityQualifiers []string

	//
	InstanceID string

	//
	PrivilegeGranted bool

	//
	QualifierFormats []uint16
}

// SetActivities sets the value of Activities for the instance
func (instance *CIM_Privilege) SetPropertyActivities(value []uint16) (err error) {
	return instance.SetProperty("Activities", value)
}

// GetActivities gets the value of Activities for the instance
func (instance *CIM_Privilege) GetPropertyActivities() (value []uint16, err error) {
	retValue, err := instance.GetProperty("Activities")
	if err != nil {
		return
	}
	value, ok := retValue.([]uint16)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetActivityQualifiers sets the value of ActivityQualifiers for the instance
func (instance *CIM_Privilege) SetPropertyActivityQualifiers(value []string) (err error) {
	return instance.SetProperty("ActivityQualifiers", value)
}

// GetActivityQualifiers gets the value of ActivityQualifiers for the instance
func (instance *CIM_Privilege) GetPropertyActivityQualifiers() (value []string, err error) {
	retValue, err := instance.GetProperty("ActivityQualifiers")
	if err != nil {
		return
	}
	value, ok := retValue.([]string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetInstanceID sets the value of InstanceID for the instance
func (instance *CIM_Privilege) SetPropertyInstanceID(value string) (err error) {
	return instance.SetProperty("InstanceID", value)
}

// GetInstanceID gets the value of InstanceID for the instance
func (instance *CIM_Privilege) GetPropertyInstanceID() (value string, err error) {
	retValue, err := instance.GetProperty("InstanceID")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetPrivilegeGranted sets the value of PrivilegeGranted for the instance
func (instance *CIM_Privilege) SetPropertyPrivilegeGranted(value bool) (err error) {
	return instance.SetProperty("PrivilegeGranted", value)
}

// GetPrivilegeGranted gets the value of PrivilegeGranted for the instance
func (instance *CIM_Privilege) GetPropertyPrivilegeGranted() (value bool, err error) {
	retValue, err := instance.GetProperty("PrivilegeGranted")
	if err != nil {
		return
	}
	value, ok := retValue.(bool)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetQualifierFormats sets the value of QualifierFormats for the instance
func (instance *CIM_Privilege) SetPropertyQualifierFormats(value []uint16) (err error) {
	return instance.SetProperty("QualifierFormats", value)
}

// GetQualifierFormats gets the value of QualifierFormats for the instance
func (instance *CIM_Privilege) GetPropertyQualifierFormats() (value []uint16, err error) {
	retValue, err := instance.GetProperty("QualifierFormats")
	if err != nil {
		return
	}
	value, ok := retValue.([]uint16)
	if !ok {
		// TODO: Set an error
	}
	return
}
