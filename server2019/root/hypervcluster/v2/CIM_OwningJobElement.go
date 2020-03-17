// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 3/16/2020 using wmigen
//      Source root.HyperVCluster.v2
//////////////////////////////////////////////
package v2

import (
	"github.com/microsoft/wmi/pkg/wmiinstance"
)

// CIM_OwningJobElement struct
type CIM_OwningJobElement struct {
	cim.WmiInstance

	// The Job created by the ManagedElement.
	OwnedElement CIM_Job

	// The ManagedElement responsible for the creation of the Job.
	OwningElement CIM_ManagedElement
}

// SetOwnedElement sets the value of OwnedElement for the instance
func (instance *CIM_OwningJobElement) SetPropertyOwnedElement(value CIM_Job) (err error) {
	return instance.SetProperty("OwnedElement", value)
}

// GetOwnedElement gets the value of OwnedElement for the instance
func (instance *CIM_OwningJobElement) GetPropertyOwnedElement() (value CIM_Job, err error) {
	retValue, err := instance.GetProperty("OwnedElement")
	if err != nil {
		return
	}
	value, ok := retValue.(CIM_Job)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetOwningElement sets the value of OwningElement for the instance
func (instance *CIM_OwningJobElement) SetPropertyOwningElement(value CIM_ManagedElement) (err error) {
	return instance.SetProperty("OwningElement", value)
}

// GetOwningElement gets the value of OwningElement for the instance
func (instance *CIM_OwningJobElement) GetPropertyOwningElement() (value CIM_ManagedElement, err error) {
	retValue, err := instance.GetProperty("OwningElement")
	if err != nil {
		return
	}
	value, ok := retValue.(CIM_ManagedElement)
	if !ok {
		// TODO: Set an error
	}
	return
}
