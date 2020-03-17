// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 3/16/2020 using wmigen
//      Source root.Hardware
//////////////////////////////////////////////
package hardware

// CIM_RegisteredProfile struct
type CIM_RegisteredProfile struct {
	CIM_ManagedElement

	//
	AdvertiseTypeDescriptions []string

	//
	AdvertiseTypes []uint16

	//
	InstanceID string

	//
	OtherRegisteredOrganization string

	//
	RegisteredName string

	//
	RegisteredOrganization uint16

	//
	RegisteredVersion string
}

// SetAdvertiseTypeDescriptions sets the value of AdvertiseTypeDescriptions for the instance
func (instance *CIM_RegisteredProfile) SetPropertyAdvertiseTypeDescriptions(value []string) (err error) {
	return instance.SetProperty("AdvertiseTypeDescriptions", value)
}

// GetAdvertiseTypeDescriptions gets the value of AdvertiseTypeDescriptions for the instance
func (instance *CIM_RegisteredProfile) GetPropertyAdvertiseTypeDescriptions() (value []string, err error) {
	retValue, err := instance.GetProperty("AdvertiseTypeDescriptions")
	if err != nil {
		return
	}
	value, ok := retValue.([]string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetAdvertiseTypes sets the value of AdvertiseTypes for the instance
func (instance *CIM_RegisteredProfile) SetPropertyAdvertiseTypes(value []uint16) (err error) {
	return instance.SetProperty("AdvertiseTypes", value)
}

// GetAdvertiseTypes gets the value of AdvertiseTypes for the instance
func (instance *CIM_RegisteredProfile) GetPropertyAdvertiseTypes() (value []uint16, err error) {
	retValue, err := instance.GetProperty("AdvertiseTypes")
	if err != nil {
		return
	}
	value, ok := retValue.([]uint16)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetInstanceID sets the value of InstanceID for the instance
func (instance *CIM_RegisteredProfile) SetPropertyInstanceID(value string) (err error) {
	return instance.SetProperty("InstanceID", value)
}

// GetInstanceID gets the value of InstanceID for the instance
func (instance *CIM_RegisteredProfile) GetPropertyInstanceID() (value string, err error) {
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

// SetOtherRegisteredOrganization sets the value of OtherRegisteredOrganization for the instance
func (instance *CIM_RegisteredProfile) SetPropertyOtherRegisteredOrganization(value string) (err error) {
	return instance.SetProperty("OtherRegisteredOrganization", value)
}

// GetOtherRegisteredOrganization gets the value of OtherRegisteredOrganization for the instance
func (instance *CIM_RegisteredProfile) GetPropertyOtherRegisteredOrganization() (value string, err error) {
	retValue, err := instance.GetProperty("OtherRegisteredOrganization")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetRegisteredName sets the value of RegisteredName for the instance
func (instance *CIM_RegisteredProfile) SetPropertyRegisteredName(value string) (err error) {
	return instance.SetProperty("RegisteredName", value)
}

// GetRegisteredName gets the value of RegisteredName for the instance
func (instance *CIM_RegisteredProfile) GetPropertyRegisteredName() (value string, err error) {
	retValue, err := instance.GetProperty("RegisteredName")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetRegisteredOrganization sets the value of RegisteredOrganization for the instance
func (instance *CIM_RegisteredProfile) SetPropertyRegisteredOrganization(value uint16) (err error) {
	return instance.SetProperty("RegisteredOrganization", value)
}

// GetRegisteredOrganization gets the value of RegisteredOrganization for the instance
func (instance *CIM_RegisteredProfile) GetPropertyRegisteredOrganization() (value uint16, err error) {
	retValue, err := instance.GetProperty("RegisteredOrganization")
	if err != nil {
		return
	}
	value, ok := retValue.(uint16)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetRegisteredVersion sets the value of RegisteredVersion for the instance
func (instance *CIM_RegisteredProfile) SetPropertyRegisteredVersion(value string) (err error) {
	return instance.SetProperty("RegisteredVersion", value)
}

// GetRegisteredVersion gets the value of RegisteredVersion for the instance
func (instance *CIM_RegisteredProfile) GetPropertyRegisteredVersion() (value string, err error) {
	retValue, err := instance.GetProperty("RegisteredVersion")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}
