// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 3/16/2020 using wmigen
//      Source root.Hardware
//////////////////////////////////////////////
package hardware

// CIM_Group struct
type CIM_Group struct {
	CIM_Collection

	//
	BusinessCategory string

	//
	CommonName string

	//
	CreationClassName string

	//
	Name string
}

// SetBusinessCategory sets the value of BusinessCategory for the instance
func (instance *CIM_Group) SetPropertyBusinessCategory(value string) (err error) {
	return instance.SetProperty("BusinessCategory", value)
}

// GetBusinessCategory gets the value of BusinessCategory for the instance
func (instance *CIM_Group) GetPropertyBusinessCategory() (value string, err error) {
	retValue, err := instance.GetProperty("BusinessCategory")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetCommonName sets the value of CommonName for the instance
func (instance *CIM_Group) SetPropertyCommonName(value string) (err error) {
	return instance.SetProperty("CommonName", value)
}

// GetCommonName gets the value of CommonName for the instance
func (instance *CIM_Group) GetPropertyCommonName() (value string, err error) {
	retValue, err := instance.GetProperty("CommonName")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetCreationClassName sets the value of CreationClassName for the instance
func (instance *CIM_Group) SetPropertyCreationClassName(value string) (err error) {
	return instance.SetProperty("CreationClassName", value)
}

// GetCreationClassName gets the value of CreationClassName for the instance
func (instance *CIM_Group) GetPropertyCreationClassName() (value string, err error) {
	retValue, err := instance.GetProperty("CreationClassName")
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
func (instance *CIM_Group) SetPropertyName(value string) (err error) {
	return instance.SetProperty("Name", value)
}

// GetName gets the value of Name for the instance
func (instance *CIM_Group) GetPropertyName() (value string, err error) {
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
