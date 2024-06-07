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

// DurableOperationAttribute struct
type DurableOperationAttribute struct { 
	*Behavior

	// Specifies whether an Activation message can be processed by the Operation.
	CanCreateInstance bool

	// Specifies whether the runtime will complete the instance after the Operation.
	CompletesInstance bool
}

	func NewDurableOperationAttributeEx1(instance *cim.WmiInstance) (newInstance *DurableOperationAttribute, err error) {tmp, err := NewBehaviorEx1(instance)
		
	if err != nil { return }
	newInstance = &DurableOperationAttribute {
	Behavior: tmp,
	}
	return
	}
	

	func NewDurableOperationAttributeEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery ) (newInstance *DurableOperationAttribute, err error) {tmp, err := NewBehaviorEx6(hostName, wmiNamespace, userName, password, domainName, query)
		
	if err != nil { return }
	newInstance = &DurableOperationAttribute {
	Behavior: tmp,
	}
	return
	}
	

// SetCanCreateInstance sets the value of CanCreateInstance for the instance
func (instance *DurableOperationAttribute) SetPropertyCanCreateInstance(value bool) (err error) { 
    return instance.SetProperty("CanCreateInstance", (value))
}

// GetCanCreateInstance gets the value of CanCreateInstance for the instance
func (instance *DurableOperationAttribute) GetPropertyCanCreateInstance()(value bool, err error) { 
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

// SetCompletesInstance sets the value of CompletesInstance for the instance
func (instance *DurableOperationAttribute) SetPropertyCompletesInstance(value bool) (err error) { 
    return instance.SetProperty("CompletesInstance", (value))
}

// GetCompletesInstance gets the value of CompletesInstance for the instance
func (instance *DurableOperationAttribute) GetPropertyCompletesInstance()(value bool, err error) { 
    retValue, err := instance.GetProperty("CompletesInstance")
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

