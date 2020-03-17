// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 3/16/2020 using wmigen
//      Source root.StandardCimv2
//////////////////////////////////////////////
package standardcimv2

// CIM_PolicyActionStructure struct
type CIM_PolicyActionStructure struct {
	CIM_PolicyComponent

	//
	ActionOrder uint16
}

// SetActionOrder sets the value of ActionOrder for the instance
func (instance *CIM_PolicyActionStructure) SetPropertyActionOrder(value uint16) (err error) {
	return instance.SetProperty("ActionOrder", value)
}

// GetActionOrder gets the value of ActionOrder for the instance
func (instance *CIM_PolicyActionStructure) GetPropertyActionOrder() (value uint16, err error) {
	retValue, err := instance.GetProperty("ActionOrder")
	if err != nil {
		return
	}
	value, ok := retValue.(uint16)
	if !ok {
		// TODO: Set an error
	}
	return
}
