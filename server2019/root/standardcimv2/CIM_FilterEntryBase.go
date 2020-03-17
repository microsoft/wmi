// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 3/16/2020 using wmigen
//      Source root.StandardCimv2
//////////////////////////////////////////////
package standardcimv2

// CIM_FilterEntryBase struct
type CIM_FilterEntryBase struct {
	CIM_LogicalElement

	//
	CreationClassName string

	//
	IsNegated bool

	//
	SystemCreationClassName string

	//
	SystemName string
}

// SetCreationClassName sets the value of CreationClassName for the instance
func (instance *CIM_FilterEntryBase) SetPropertyCreationClassName(value string) (err error) {
	return instance.SetProperty("CreationClassName", value)
}

// GetCreationClassName gets the value of CreationClassName for the instance
func (instance *CIM_FilterEntryBase) GetPropertyCreationClassName() (value string, err error) {
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

// SetIsNegated sets the value of IsNegated for the instance
func (instance *CIM_FilterEntryBase) SetPropertyIsNegated(value bool) (err error) {
	return instance.SetProperty("IsNegated", value)
}

// GetIsNegated gets the value of IsNegated for the instance
func (instance *CIM_FilterEntryBase) GetPropertyIsNegated() (value bool, err error) {
	retValue, err := instance.GetProperty("IsNegated")
	if err != nil {
		return
	}
	value, ok := retValue.(bool)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetSystemCreationClassName sets the value of SystemCreationClassName for the instance
func (instance *CIM_FilterEntryBase) SetPropertySystemCreationClassName(value string) (err error) {
	return instance.SetProperty("SystemCreationClassName", value)
}

// GetSystemCreationClassName gets the value of SystemCreationClassName for the instance
func (instance *CIM_FilterEntryBase) GetPropertySystemCreationClassName() (value string, err error) {
	retValue, err := instance.GetProperty("SystemCreationClassName")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetSystemName sets the value of SystemName for the instance
func (instance *CIM_FilterEntryBase) SetPropertySystemName(value string) (err error) {
	return instance.SetProperty("SystemName", value)
}

// GetSystemName gets the value of SystemName for the instance
func (instance *CIM_FilterEntryBase) GetPropertySystemName() (value string, err error) {
	retValue, err := instance.GetProperty("SystemName")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}
