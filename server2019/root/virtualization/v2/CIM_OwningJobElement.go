// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

// 
// Author:
//      Auto Generated on 6/6/2024 using wmigen
//      Source root.virtualization.v2
//////////////////////////////////////////////
package v2
import (
 "github.com/microsoft/wmi/pkg/base/query"
cim "github.com/microsoft/wmi/pkg/wmiinstance"
 "github.com/microsoft/wmi/pkg/base/instance"
 "github.com/microsoft/wmi/pkg/errors"
 "reflect"
)

// CIM_OwningJobElement struct
type CIM_OwningJobElement struct { 
	*cim.WmiInstance

	// The Job created by the ManagedElement.
	OwnedElement CIM_Job

	// The ManagedElement responsible for the creation of the Job.
	OwningElement CIM_ManagedElement
}

	func NewCIM_OwningJobElementEx1(instance *cim.WmiInstance) (newInstance *CIM_OwningJobElement, err error) {tmp, err := instance, nil
		
	if err != nil { return }
	newInstance = &CIM_OwningJobElement {
	WmiInstance: tmp,
	}
	return
	}
	

	func NewCIM_OwningJobElementEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery ) (newInstance *CIM_OwningJobElement, err error) {tmp, err := instance.GetWmiInstance(hostName, wmiNamespace, userName, password, domainName, query)
		
	if err != nil { return }
	newInstance = &CIM_OwningJobElement {
	WmiInstance: tmp,
	}
	return
	}
	

// SetOwnedElement sets the value of OwnedElement for the instance
func (instance *CIM_OwningJobElement) SetPropertyOwnedElement(value CIM_Job) (err error) { 
    return instance.SetProperty("OwnedElement", (value))
}

// GetOwnedElement gets the value of OwnedElement for the instance
func (instance *CIM_OwningJobElement) GetPropertyOwnedElement()(value CIM_Job, err error) { 
    retValue, err := instance.GetProperty("OwnedElement")
    if err != nil {
        return
    }
    if retValue == nil {
        // Doesn't have any value. Return empty
        return
    }
    
    valuetmp, ok := retValue.(CIM_Job); 
    if !ok {
        err = errors.Wrapf(errors.InvalidType, " CIM_Job is Invalid. Expected %s", reflect.TypeOf(retValue))
        return  
    }

    value = CIM_Job(valuetmp)

    return
}

// SetOwningElement sets the value of OwningElement for the instance
func (instance *CIM_OwningJobElement) SetPropertyOwningElement(value CIM_ManagedElement) (err error) { 
    return instance.SetProperty("OwningElement", (value))
}

// GetOwningElement gets the value of OwningElement for the instance
func (instance *CIM_OwningJobElement) GetPropertyOwningElement()(value CIM_ManagedElement, err error) { 
    retValue, err := instance.GetProperty("OwningElement")
    if err != nil {
        return
    }
    if retValue == nil {
        // Doesn't have any value. Return empty
        return
    }
    
    valuetmp, ok := retValue.(CIM_ManagedElement); 
    if !ok {
        err = errors.Wrapf(errors.InvalidType, " CIM_ManagedElement is Invalid. Expected %s", reflect.TypeOf(retValue))
        return  
    }

    value = CIM_ManagedElement(valuetmp)

    return
}

