// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

// 
// Author:
//      Auto Generated on 6/6/2024 using wmigen
//      Source root.DEFAULT
//////////////////////////////////////////////
package default
import (
 "github.com/microsoft/wmi/pkg/base/query"
cim "github.com/microsoft/wmi/pkg/wmiinstance"
 "github.com/microsoft/wmi/pkg/errors"
 "reflect"
)

// __AdapStatus struct
type __AdapStatus struct { 
	*__SystemClass

	// 
	LastStartTime string

	// 
	LastStopTime string

	// 
	Status uint32
}

	func New__AdapStatusEx1(instance *cim.WmiInstance) (newInstance *__AdapStatus, err error) {tmp, err := New__SystemClassEx1(instance)
		
	if err != nil { return }
	newInstance = &__AdapStatus {
	__SystemClass: tmp,
	}
	return
	}
	

	func New__AdapStatusEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery ) (newInstance *__AdapStatus, err error) {tmp, err := New__SystemClassEx6(hostName, wmiNamespace, userName, password, domainName, query)
		
	if err != nil { return }
	newInstance = &__AdapStatus {
	__SystemClass: tmp,
	}
	return
	}
	

// SetLastStartTime sets the value of LastStartTime for the instance
func (instance *__AdapStatus) SetPropertyLastStartTime(value string) (err error) { 
    return instance.SetProperty("LastStartTime", (value))
}

// GetLastStartTime gets the value of LastStartTime for the instance
func (instance *__AdapStatus) GetPropertyLastStartTime()(value string, err error) { 
    retValue, err := instance.GetProperty("LastStartTime")
    if err != nil {
        return
    }
    if retValue == nil {
        // Doesn't have any value. Return empty
        return
    }
    
    valuetmp, ok := retValue.(string); 
    if !ok {
        err = errors.Wrapf(errors.InvalidType, " string is Invalid. Expected %s", reflect.TypeOf(retValue))
        return  
    }

    value = string(valuetmp)

    return
}

// SetLastStopTime sets the value of LastStopTime for the instance
func (instance *__AdapStatus) SetPropertyLastStopTime(value string) (err error) { 
    return instance.SetProperty("LastStopTime", (value))
}

// GetLastStopTime gets the value of LastStopTime for the instance
func (instance *__AdapStatus) GetPropertyLastStopTime()(value string, err error) { 
    retValue, err := instance.GetProperty("LastStopTime")
    if err != nil {
        return
    }
    if retValue == nil {
        // Doesn't have any value. Return empty
        return
    }
    
    valuetmp, ok := retValue.(string); 
    if !ok {
        err = errors.Wrapf(errors.InvalidType, " string is Invalid. Expected %s", reflect.TypeOf(retValue))
        return  
    }

    value = string(valuetmp)

    return
}

// SetStatus sets the value of Status for the instance
func (instance *__AdapStatus) SetPropertyStatus(value uint32) (err error) { 
    return instance.SetProperty("Status", (value))
}

// GetStatus gets the value of Status for the instance
func (instance *__AdapStatus) GetPropertyStatus()(value uint32, err error) { 
    retValue, err := instance.GetProperty("Status")
    if err != nil {
        return
    }
    if retValue == nil {
        // Doesn't have any value. Return empty
        return
    }
    
    valuetmp, ok := retValue.(uint32); 
    if !ok {
        err = errors.Wrapf(errors.InvalidType, " uint32 is Invalid. Expected %s", reflect.TypeOf(retValue))
        return  
    }

    value = uint32(valuetmp)

    return
}

