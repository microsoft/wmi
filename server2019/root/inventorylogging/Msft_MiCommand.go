// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 3/16/2020 using wmigen
//      Source root.InventoryLogging
//////////////////////////////////////////////
package inventorylogging

// Msft_MiCommand struct
type Msft_MiCommand struct {
	Msft_MiStream

	//
	ClassName string

	//
	Input interface{}

	//
	MethodName string

	//
	NamespaceName string

	//
	Parameters interface{}
}

// SetClassName sets the value of ClassName for the instance
func (instance *Msft_MiCommand) SetPropertyClassName(value string) (err error) {
	return instance.SetProperty("ClassName", value)
}

// GetClassName gets the value of ClassName for the instance
func (instance *Msft_MiCommand) GetPropertyClassName() (value string, err error) {
	retValue, err := instance.GetProperty("ClassName")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetInput sets the value of Input for the instance
func (instance *Msft_MiCommand) SetPropertyInput(value interface{}) (err error) {
	return instance.SetProperty("Input", value)
}

// GetInput gets the value of Input for the instance
func (instance *Msft_MiCommand) GetPropertyInput() (value interface{}, err error) {
	retValue, err := instance.GetProperty("Input")
	if err != nil {
		return
	}
	value, ok := retValue.(interface{})
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetMethodName sets the value of MethodName for the instance
func (instance *Msft_MiCommand) SetPropertyMethodName(value string) (err error) {
	return instance.SetProperty("MethodName", value)
}

// GetMethodName gets the value of MethodName for the instance
func (instance *Msft_MiCommand) GetPropertyMethodName() (value string, err error) {
	retValue, err := instance.GetProperty("MethodName")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetNamespaceName sets the value of NamespaceName for the instance
func (instance *Msft_MiCommand) SetPropertyNamespaceName(value string) (err error) {
	return instance.SetProperty("NamespaceName", value)
}

// GetNamespaceName gets the value of NamespaceName for the instance
func (instance *Msft_MiCommand) GetPropertyNamespaceName() (value string, err error) {
	retValue, err := instance.GetProperty("NamespaceName")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetParameters sets the value of Parameters for the instance
func (instance *Msft_MiCommand) SetPropertyParameters(value interface{}) (err error) {
	return instance.SetProperty("Parameters", value)
}

// GetParameters gets the value of Parameters for the instance
func (instance *Msft_MiCommand) GetPropertyParameters() (value interface{}, err error) {
	retValue, err := instance.GetProperty("Parameters")
	if err != nil {
		return
	}
	value, ok := retValue.(interface{})
	if !ok {
		// TODO: Set an error
	}
	return
}
