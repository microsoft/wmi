// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

// 
// Author:
//      Auto Generated on 6/6/2024 using wmigen
//      Source root.ServiceModel
//////////////////////////////////////////////
package servicemodel
import (
 "github.com/microsoft/wmi/pkg/base/query"
cim "github.com/microsoft/wmi/pkg/wmiinstance"
 "github.com/microsoft/wmi/pkg/errors"
 "reflect"
)

// WorkflowOperationBehavior struct
type WorkflowOperationBehavior struct { 
	*Behavior

	// Specifies whether an Activation message can be processed by the Operation.
	CanCreateInstance bool

	// Specifies whether the Operation is invoked Synchronously or not.
	SynchronousDispatch bool
}

	func NewWorkflowOperationBehaviorEx1(instance *cim.WmiInstance) (newInstance *WorkflowOperationBehavior, err error) {tmp, err := NewBehaviorEx1(instance)
		
	if err != nil { return }
	newInstance = &WorkflowOperationBehavior {
	Behavior: tmp,
	}
	return
	}
	

	func NewWorkflowOperationBehaviorEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery ) (newInstance *WorkflowOperationBehavior, err error) {tmp, err := NewBehaviorEx6(hostName, wmiNamespace, userName, password, domainName, query)
		
	if err != nil { return }
	newInstance = &WorkflowOperationBehavior {
	Behavior: tmp,
	}
	return
	}
	

// SetCanCreateInstance sets the value of CanCreateInstance for the instance
func (instance *WorkflowOperationBehavior) SetPropertyCanCreateInstance(value bool) (err error) { 
    return instance.SetProperty("CanCreateInstance", (value))
}

// GetCanCreateInstance gets the value of CanCreateInstance for the instance
func (instance *WorkflowOperationBehavior) GetPropertyCanCreateInstance()(value bool, err error) { 
    retValue, err := instance.GetProperty("CanCreateInstance")
    if err != nil {
        return
    }
    if retValue == nil {
        // Doesn't have any value. Return empty
        return
    }
    
    valuetmp, ok := retValue.(bool); 
    if !ok {
        err = errors.Wrapf(errors.InvalidType, " bool is Invalid. Expected %s", reflect.TypeOf(retValue))
        return  
    }

    value = bool(valuetmp)

    return
}

// SetSynchronousDispatch sets the value of SynchronousDispatch for the instance
func (instance *WorkflowOperationBehavior) SetPropertySynchronousDispatch(value bool) (err error) { 
    return instance.SetProperty("SynchronousDispatch", (value))
}

// GetSynchronousDispatch gets the value of SynchronousDispatch for the instance
func (instance *WorkflowOperationBehavior) GetPropertySynchronousDispatch()(value bool, err error) { 
    retValue, err := instance.GetProperty("SynchronousDispatch")
    if err != nil {
        return
    }
    if retValue == nil {
        // Doesn't have any value. Return empty
        return
    }
    
    valuetmp, ok := retValue.(bool); 
    if !ok {
        err = errors.Wrapf(errors.InvalidType, " bool is Invalid. Expected %s", reflect.TypeOf(retValue))
        return  
    }

    value = bool(valuetmp)

    return
}

