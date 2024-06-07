// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

// 
// Author:
//      Auto Generated on 6/6/2024 using wmigen
//      Source root.WMI
//////////////////////////////////////////////
package wmi
import (
 "github.com/microsoft/wmi/pkg/base/query"
cim "github.com/microsoft/wmi/pkg/wmiinstance"
 "github.com/microsoft/wmi/pkg/base/instance"
 "github.com/microsoft/wmi/pkg/errors"
 "reflect"
)

// MSStorageDriver_ClassErrorLog struct
type MSStorageDriver_ClassErrorLog struct { 
	*cim.WmiInstance

	// 
	Active bool

	// 
	InstanceName string

	// Error Log Array
	logEntries []MSStorageDriver_ClassErrorLogEntry

	// Number of Error Log Entries
	numEntries uint32
}

	func NewMSStorageDriver_ClassErrorLogEx1(instance *cim.WmiInstance) (newInstance *MSStorageDriver_ClassErrorLog, err error) {tmp, err := instance, nil
		
	if err != nil { return }
	newInstance = &MSStorageDriver_ClassErrorLog {
	WmiInstance: tmp,
	}
	return
	}
	

	func NewMSStorageDriver_ClassErrorLogEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery ) (newInstance *MSStorageDriver_ClassErrorLog, err error) {tmp, err := instance.GetWmiInstance(hostName, wmiNamespace, userName, password, domainName, query)
		
	if err != nil { return }
	newInstance = &MSStorageDriver_ClassErrorLog {
	WmiInstance: tmp,
	}
	return
	}
	

// SetActive sets the value of Active for the instance
func (instance *MSStorageDriver_ClassErrorLog) SetPropertyActive(value bool) (err error) { 
    return instance.SetProperty("Active", (value))
}

// GetActive gets the value of Active for the instance
func (instance *MSStorageDriver_ClassErrorLog) GetPropertyActive()(value bool, err error) { 
    retValue, err := instance.GetProperty("Active")
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

// SetInstanceName sets the value of InstanceName for the instance
func (instance *MSStorageDriver_ClassErrorLog) SetPropertyInstanceName(value string) (err error) { 
    return instance.SetProperty("InstanceName", (value))
}

// GetInstanceName gets the value of InstanceName for the instance
func (instance *MSStorageDriver_ClassErrorLog) GetPropertyInstanceName()(value string, err error) { 
    retValue, err := instance.GetProperty("InstanceName")
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

// SetlogEntries sets the value of logEntries for the instance
func (instance *MSStorageDriver_ClassErrorLog) SetPropertylogEntries(value []MSStorageDriver_ClassErrorLogEntry) (err error) { 
    return instance.SetProperty("logEntries", (value))
}

// GetlogEntries gets the value of logEntries for the instance
func (instance *MSStorageDriver_ClassErrorLog) GetPropertylogEntries()(value []MSStorageDriver_ClassErrorLogEntry, err error) { 
    retValue, err := instance.GetProperty("logEntries")
    if err != nil {
        return
    }
    if retValue == nil {
        // Doesn't have any value. Return empty
        return
    }
    
    for _, interfaceValue := range retValue.([]interface{}) {
        valuetmp, ok := interfaceValue.(MSStorageDriver_ClassErrorLogEntry); 
        if !ok {
            err = errors.Wrapf(errors.InvalidType, " MSStorageDriver_ClassErrorLogEntry is Invalid. Expected %s", reflect.TypeOf(interfaceValue))
            return  
        }
        value = append(value, MSStorageDriver_ClassErrorLogEntry(valuetmp))
    }

    return
}

// SetnumEntries sets the value of numEntries for the instance
func (instance *MSStorageDriver_ClassErrorLog) SetPropertynumEntries(value uint32) (err error) { 
    return instance.SetProperty("numEntries", (value))
}

// GetnumEntries gets the value of numEntries for the instance
func (instance *MSStorageDriver_ClassErrorLog) GetPropertynumEntries()(value uint32, err error) { 
    retValue, err := instance.GetProperty("numEntries")
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

