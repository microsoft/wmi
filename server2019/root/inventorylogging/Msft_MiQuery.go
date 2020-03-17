// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 3/16/2020 using wmigen
//      Source root.InventoryLogging
//////////////////////////////////////////////
package inventorylogging

// Msft_MiQuery struct
type Msft_MiQuery struct {
	Msft_MiStream

	//
	Dialect string

	//
	Expression string

	//
	NamespaceName string
}

// SetDialect sets the value of Dialect for the instance
func (instance *Msft_MiQuery) SetPropertyDialect(value string) (err error) {
	return instance.SetProperty("Dialect", value)
}

// GetDialect gets the value of Dialect for the instance
func (instance *Msft_MiQuery) GetPropertyDialect() (value string, err error) {
	retValue, err := instance.GetProperty("Dialect")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetExpression sets the value of Expression for the instance
func (instance *Msft_MiQuery) SetPropertyExpression(value string) (err error) {
	return instance.SetProperty("Expression", value)
}

// GetExpression gets the value of Expression for the instance
func (instance *Msft_MiQuery) GetPropertyExpression() (value string, err error) {
	retValue, err := instance.GetProperty("Expression")
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
func (instance *Msft_MiQuery) SetPropertyNamespaceName(value string) (err error) {
	return instance.SetProperty("NamespaceName", value)
}

// GetNamespaceName gets the value of NamespaceName for the instance
func (instance *Msft_MiQuery) GetPropertyNamespaceName() (value string, err error) {
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
