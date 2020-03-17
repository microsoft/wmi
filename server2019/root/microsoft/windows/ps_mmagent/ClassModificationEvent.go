// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 3/16/2020 using wmigen
//      Source root.Microsoft.Windows.PS_MMAgent
//////////////////////////////////////////////
package ps_mmagent

// __ClassModificationEvent struct
type __ClassModificationEvent struct {
	__ClassOperationEvent

	//
	PreviousClass interface{}
}

// SetPreviousClass sets the value of PreviousClass for the instance
func (instance *__ClassModificationEvent) SetPropertyPreviousClass(value interface{}) (err error) {
	return instance.SetProperty("PreviousClass", value)
}

// GetPreviousClass gets the value of PreviousClass for the instance
func (instance *__ClassModificationEvent) GetPropertyPreviousClass() (value interface{}, err error) {
	retValue, err := instance.GetProperty("PreviousClass")
	if err != nil {
		return
	}
	value, ok := retValue.(interface{})
	if !ok {
		// TODO: Set an error
	}
	return
}
