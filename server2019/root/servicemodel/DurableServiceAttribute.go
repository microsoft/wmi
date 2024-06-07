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

// DurableServiceAttribute struct
type DurableServiceAttribute struct { 
	*Behavior

	// The SaveStateInOperationTransaction boolean in DurableServiceAttribute
	SaveStateInOperationTransaction bool
}

	func NewDurableServiceAttributeEx1(instance *cim.WmiInstance) (newInstance *DurableServiceAttribute, err error) {tmp, err := NewBehaviorEx1(instance)
		
	if err != nil { return }
	newInstance = &DurableServiceAttribute {
	Behavior: tmp,
	}
	return
	}
	

	func NewDurableServiceAttributeEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery ) (newInstance *DurableServiceAttribute, err error) {tmp, err := NewBehaviorEx6(hostName, wmiNamespace, userName, password, domainName, query)
		
	if err != nil { return }
	newInstance = &DurableServiceAttribute {
	Behavior: tmp,
	}
	return
	}
	

// SetSaveStateInOperationTransaction sets the value of SaveStateInOperationTransaction for the instance
func (instance *DurableServiceAttribute) SetPropertySaveStateInOperationTransaction(value bool) (err error) { 
    return instance.SetProperty("SaveStateInOperationTransaction", (value))
}

// GetSaveStateInOperationTransaction gets the value of SaveStateInOperationTransaction for the instance
func (instance *DurableServiceAttribute) GetPropertySaveStateInOperationTransaction()(value bool, err error) { 
    retValue, err := instance.GetProperty("SaveStateInOperationTransaction")
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

