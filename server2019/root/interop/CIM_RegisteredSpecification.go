// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 3/16/2020 using wmigen
//      Source root.Interop
//////////////////////////////////////////////
package interop

// CIM_RegisteredSpecification struct
type CIM_RegisteredSpecification struct {
	CIM_ManagedElement

	//
	AdvertiseTypeDescriptions []string

	//
	AdvertiseTypes []uint16

	//
	OtherRegisteredOrganization string

	//
	OtherSpecificationType string

	//
	RegisteredName string

	//
	RegisteredOrganization uint16

	//
	RegisteredVersion string

	//
	SpecificationType uint16
}

// SetAdvertiseTypeDescriptions sets the value of AdvertiseTypeDescriptions for the instance
func (instance *CIM_RegisteredSpecification) SetPropertyAdvertiseTypeDescriptions(value []string) (err error) {
	return instance.SetProperty("AdvertiseTypeDescriptions", value)
}

// GetAdvertiseTypeDescriptions gets the value of AdvertiseTypeDescriptions for the instance
func (instance *CIM_RegisteredSpecification) GetPropertyAdvertiseTypeDescriptions() (value []string, err error) {
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
func (instance *CIM_RegisteredSpecification) SetPropertyAdvertiseTypes(value []uint16) (err error) {
	return instance.SetProperty("AdvertiseTypes", value)
}

// GetAdvertiseTypes gets the value of AdvertiseTypes for the instance
func (instance *CIM_RegisteredSpecification) GetPropertyAdvertiseTypes() (value []uint16, err error) {
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

// SetOtherRegisteredOrganization sets the value of OtherRegisteredOrganization for the instance
func (instance *CIM_RegisteredSpecification) SetPropertyOtherRegisteredOrganization(value string) (err error) {
	return instance.SetProperty("OtherRegisteredOrganization", value)
}

// GetOtherRegisteredOrganization gets the value of OtherRegisteredOrganization for the instance
func (instance *CIM_RegisteredSpecification) GetPropertyOtherRegisteredOrganization() (value string, err error) {
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

// SetOtherSpecificationType sets the value of OtherSpecificationType for the instance
func (instance *CIM_RegisteredSpecification) SetPropertyOtherSpecificationType(value string) (err error) {
	return instance.SetProperty("OtherSpecificationType", value)
}

// GetOtherSpecificationType gets the value of OtherSpecificationType for the instance
func (instance *CIM_RegisteredSpecification) GetPropertyOtherSpecificationType() (value string, err error) {
	retValue, err := instance.GetProperty("OtherSpecificationType")
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
func (instance *CIM_RegisteredSpecification) SetPropertyRegisteredName(value string) (err error) {
	return instance.SetProperty("RegisteredName", value)
}

// GetRegisteredName gets the value of RegisteredName for the instance
func (instance *CIM_RegisteredSpecification) GetPropertyRegisteredName() (value string, err error) {
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
func (instance *CIM_RegisteredSpecification) SetPropertyRegisteredOrganization(value uint16) (err error) {
	return instance.SetProperty("RegisteredOrganization", value)
}

// GetRegisteredOrganization gets the value of RegisteredOrganization for the instance
func (instance *CIM_RegisteredSpecification) GetPropertyRegisteredOrganization() (value uint16, err error) {
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
func (instance *CIM_RegisteredSpecification) SetPropertyRegisteredVersion(value string) (err error) {
	return instance.SetProperty("RegisteredVersion", value)
}

// GetRegisteredVersion gets the value of RegisteredVersion for the instance
func (instance *CIM_RegisteredSpecification) GetPropertyRegisteredVersion() (value string, err error) {
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

// SetSpecificationType sets the value of SpecificationType for the instance
func (instance *CIM_RegisteredSpecification) SetPropertySpecificationType(value uint16) (err error) {
	return instance.SetProperty("SpecificationType", value)
}

// GetSpecificationType gets the value of SpecificationType for the instance
func (instance *CIM_RegisteredSpecification) GetPropertySpecificationType() (value uint16, err error) {
	retValue, err := instance.GetProperty("SpecificationType")
	if err != nil {
		return
	}
	value, ok := retValue.(uint16)
	if !ok {
		// TODO: Set an error
	}
	return
}
