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

// DispatcherSynchronizationBehavior struct
type DispatcherSynchronizationBehavior struct { 
	*Behavior

	// When enabled the reply on the channel will be send asynchronously.
	AsynchronousSendEnabled bool

	// Limits the maximum number of pending receives that may be queued on the channel.
	MaxPendingReceives int32
}

	func NewDispatcherSynchronizationBehaviorEx1(instance *cim.WmiInstance) (newInstance *DispatcherSynchronizationBehavior, err error) {tmp, err := NewBehaviorEx1(instance)
		
	if err != nil { return }
	newInstance = &DispatcherSynchronizationBehavior {
	Behavior: tmp,
	}
	return
	}
	

	func NewDispatcherSynchronizationBehaviorEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery ) (newInstance *DispatcherSynchronizationBehavior, err error) {tmp, err := NewBehaviorEx6(hostName, wmiNamespace, userName, password, domainName, query)
		
	if err != nil { return }
	newInstance = &DispatcherSynchronizationBehavior {
	Behavior: tmp,
	}
	return
	}
	

// SetAsynchronousSendEnabled sets the value of AsynchronousSendEnabled for the instance
func (instance *DispatcherSynchronizationBehavior) SetPropertyAsynchronousSendEnabled(value bool) (err error) { 
    return instance.SetProperty("AsynchronousSendEnabled", (value))
}

// GetAsynchronousSendEnabled gets the value of AsynchronousSendEnabled for the instance
func (instance *DispatcherSynchronizationBehavior) GetPropertyAsynchronousSendEnabled()(value bool, err error) { 
    retValue, err := instance.GetProperty("AsynchronousSendEnabled")
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

// SetMaxPendingReceives sets the value of MaxPendingReceives for the instance
func (instance *DispatcherSynchronizationBehavior) SetPropertyMaxPendingReceives(value int32) (err error) { 
    return instance.SetProperty("MaxPendingReceives", (value))
}

// GetMaxPendingReceives gets the value of MaxPendingReceives for the instance
func (instance *DispatcherSynchronizationBehavior) GetPropertyMaxPendingReceives()(value int32, err error) { 
    retValue, err := instance.GetProperty("MaxPendingReceives")
    if err != nil {
        return
    }
    if retValue == nil {
        // Doesn't have any value. Return empty
        return
    }
    
    valuetmp, ok := retValue.(int32); 
    if !ok {
        err = errors.Wrapf(errors.InvalidType, " int32 is Invalid. Expected %s", reflect.TypeOf(retValue))
        return  
    }

    value = int32(valuetmp)

    return
}

