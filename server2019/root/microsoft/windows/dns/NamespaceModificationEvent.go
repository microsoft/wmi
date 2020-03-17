// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 3/16/2020 using wmigen
//      Source root.Microsoft.Windows.Dns
//////////////////////////////////////////////
package dns

// __NamespaceModificationEvent struct
type __NamespaceModificationEvent struct {
	__NamespaceOperationEvent

	//
	PreviousNamespace __Namespace
}

// SetPreviousNamespace sets the value of PreviousNamespace for the instance
func (instance *__NamespaceModificationEvent) SetPropertyPreviousNamespace(value __Namespace) (err error) {
	return instance.SetProperty("PreviousNamespace", value)
}

// GetPreviousNamespace gets the value of PreviousNamespace for the instance
func (instance *__NamespaceModificationEvent) GetPropertyPreviousNamespace() (value __Namespace, err error) {
	retValue, err := instance.GetProperty("PreviousNamespace")
	if err != nil {
		return
	}
	value, ok := retValue.(__Namespace)
	if !ok {
		// TODO: Set an error
	}
	return
}
