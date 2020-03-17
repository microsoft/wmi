// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 3/16/2020 using wmigen
//      Source root.InventoryLogging
//////////////////////////////////////////////
package inventorylogging

// __NamespaceOperationEvent struct
type __NamespaceOperationEvent struct {
	__Event

	//
	TargetNamespace __Namespace
}

// SetTargetNamespace sets the value of TargetNamespace for the instance
func (instance *__NamespaceOperationEvent) SetPropertyTargetNamespace(value __Namespace) (err error) {
	return instance.SetProperty("TargetNamespace", value)
}

// GetTargetNamespace gets the value of TargetNamespace for the instance
func (instance *__NamespaceOperationEvent) GetPropertyTargetNamespace() (value __Namespace, err error) {
	retValue, err := instance.GetProperty("TargetNamespace")
	if err != nil {
		return
	}
	value, ok := retValue.(__Namespace)
	if !ok {
		// TODO: Set an error
	}
	return
}
