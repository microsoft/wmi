// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 3/16/2020 using wmigen
//      Source root.Interop
//////////////////////////////////////////////
package interop

// CIM_ClassModification struct
type CIM_ClassModification struct {
	CIM_ClassIndication

	// A copy of the 'previous' class definition whose change generated the Indication. PreviousClassDefinition contains an 'older' copy of the class' information, as compared to what is found in the ClassDefinition property (inherited from ClassIndication).
	PreviousClassDefinition interface{}
}

// SetPreviousClassDefinition sets the value of PreviousClassDefinition for the instance
func (instance *CIM_ClassModification) SetPropertyPreviousClassDefinition(value interface{}) (err error) {
	return instance.SetProperty("PreviousClassDefinition", value)
}

// GetPreviousClassDefinition gets the value of PreviousClassDefinition for the instance
func (instance *CIM_ClassModification) GetPropertyPreviousClassDefinition() (value interface{}, err error) {
	retValue, err := instance.GetProperty("PreviousClassDefinition")
	if err != nil {
		return
	}
	value, ok := retValue.(interface{})
	if !ok {
		// TODO: Set an error
	}
	return
}
