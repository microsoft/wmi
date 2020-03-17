// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 3/16/2020 using wmigen
//      Source root.CIMV2
//////////////////////////////////////////////
package cimv2

// CIM_Service struct
type CIM_Service struct {
	CIM_LogicalElement

	//
	CreationClassName string

	//
	Started bool

	//
	StartMode string

	//
	SystemCreationClassName string

	//
	SystemName string
}

// SetCreationClassName sets the value of CreationClassName for the instance
func (instance *CIM_Service) SetPropertyCreationClassName(value string) (err error) {
	return instance.SetProperty("CreationClassName", value)
}

// GetCreationClassName gets the value of CreationClassName for the instance
func (instance *CIM_Service) GetPropertyCreationClassName() (value string, err error) {
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

// SetStarted sets the value of Started for the instance
func (instance *CIM_Service) SetPropertyStarted(value bool) (err error) {
	return instance.SetProperty("Started", value)
}

// GetStarted gets the value of Started for the instance
func (instance *CIM_Service) GetPropertyStarted() (value bool, err error) {
	retValue, err := instance.GetProperty("Started")
	if err != nil {
		return
	}
	value, ok := retValue.(bool)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetStartMode sets the value of StartMode for the instance
func (instance *CIM_Service) SetPropertyStartMode(value string) (err error) {
	return instance.SetProperty("StartMode", value)
}

// GetStartMode gets the value of StartMode for the instance
func (instance *CIM_Service) GetPropertyStartMode() (value string, err error) {
	retValue, err := instance.GetProperty("StartMode")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetSystemCreationClassName sets the value of SystemCreationClassName for the instance
func (instance *CIM_Service) SetPropertySystemCreationClassName(value string) (err error) {
	return instance.SetProperty("SystemCreationClassName", value)
}

// GetSystemCreationClassName gets the value of SystemCreationClassName for the instance
func (instance *CIM_Service) GetPropertySystemCreationClassName() (value string, err error) {
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
func (instance *CIM_Service) SetPropertySystemName(value string) (err error) {
	return instance.SetProperty("SystemName", value)
}

// GetSystemName gets the value of SystemName for the instance
func (instance *CIM_Service) GetPropertySystemName() (value string, err error) {
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

//

// <param name="ReturnValue" type="uint32 "></param>
func (instance *CIM_Service) StartService() (result uint32, err error) {
	retVal, err := instance.InvokeMethodWithReturn("StartService")
	if err != nil {
		return
	}
	result = uint32(retVal)
	return

}

//

// <param name="ReturnValue" type="uint32 "></param>
func (instance *CIM_Service) StopService() (result uint32, err error) {
	retVal, err := instance.InvokeMethodWithReturn("StopService")
	if err != nil {
		return
	}
	result = uint32(retVal)
	return

}
