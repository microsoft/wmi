// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 3/16/2020 using wmigen
//      Source root.CIMV2
//////////////////////////////////////////////
package cimv2

// Win32_PublishComponentAction struct
type Win32_PublishComponentAction struct {
	CIM_Action

	//
	AppData string

	//
	ComponentID string

	//
	Qual string
}

// SetAppData sets the value of AppData for the instance
func (instance *Win32_PublishComponentAction) SetPropertyAppData(value string) (err error) {
	return instance.SetProperty("AppData", value)
}

// GetAppData gets the value of AppData for the instance
func (instance *Win32_PublishComponentAction) GetPropertyAppData() (value string, err error) {
	retValue, err := instance.GetProperty("AppData")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetComponentID sets the value of ComponentID for the instance
func (instance *Win32_PublishComponentAction) SetPropertyComponentID(value string) (err error) {
	return instance.SetProperty("ComponentID", value)
}

// GetComponentID gets the value of ComponentID for the instance
func (instance *Win32_PublishComponentAction) GetPropertyComponentID() (value string, err error) {
	retValue, err := instance.GetProperty("ComponentID")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetQual sets the value of Qual for the instance
func (instance *Win32_PublishComponentAction) SetPropertyQual(value string) (err error) {
	return instance.SetProperty("Qual", value)
}

// GetQual gets the value of Qual for the instance
func (instance *Win32_PublishComponentAction) GetPropertyQual() (value string, err error) {
	retValue, err := instance.GetProperty("Qual")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}
