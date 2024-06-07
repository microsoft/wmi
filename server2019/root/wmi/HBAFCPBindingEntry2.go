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

// HBAFCPBindingEntry2 struct
type HBAFCPBindingEntry2 struct { 
	*cim.WmiInstance

	// 
	FCPId HBAFCPID

	// 
	Luid []uint8

	// 
	ScsiId HBAScsiID

	// 
	Type uint32
}

	func NewHBAFCPBindingEntry2Ex1(instance *cim.WmiInstance) (newInstance *HBAFCPBindingEntry2, err error) {tmp, err := instance, nil
		
	if err != nil { return }
	newInstance = &HBAFCPBindingEntry2 {
	WmiInstance: tmp,
	}
	return
	}
	

	func NewHBAFCPBindingEntry2Ex6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery ) (newInstance *HBAFCPBindingEntry2, err error) {tmp, err := instance.GetWmiInstance(hostName, wmiNamespace, userName, password, domainName, query)
		
	if err != nil { return }
	newInstance = &HBAFCPBindingEntry2 {
	WmiInstance: tmp,
	}
	return
	}
	

// SetFCPId sets the value of FCPId for the instance
func (instance *HBAFCPBindingEntry2) SetPropertyFCPId(value HBAFCPID) (err error) { 
    return instance.SetProperty("FCPId", (value))
}

// GetFCPId gets the value of FCPId for the instance
func (instance *HBAFCPBindingEntry2) GetPropertyFCPId()(value HBAFCPID, err error) { 
    retValue, err := instance.GetProperty("FCPId")
    if err != nil {
        return
    }
    if retValue == nil {
        // Doesn't have any value. Return empty
        return
    }
    
    valuetmp, ok := retValue.(HBAFCPID); 
    if !ok {
        err = errors.Wrapf(errors.InvalidType, " HBAFCPID is Invalid. Expected %s", reflect.TypeOf(retValue))
        return  
    }

    value = HBAFCPID(valuetmp)

    return
}

// SetLuid sets the value of Luid for the instance
func (instance *HBAFCPBindingEntry2) SetPropertyLuid(value []uint8) (err error) { 
    return instance.SetProperty("Luid", (value))
}

// GetLuid gets the value of Luid for the instance
func (instance *HBAFCPBindingEntry2) GetPropertyLuid()(value []uint8, err error) { 
    retValue, err := instance.GetProperty("Luid")
    if err != nil {
        return
    }
    if retValue == nil {
        // Doesn't have any value. Return empty
        return
    }
    
    for _, interfaceValue := range retValue.([]interface{}) {
        valuetmp, ok := interfaceValue.(uint8); 
        if !ok {
            err = errors.Wrapf(errors.InvalidType, " uint8 is Invalid. Expected %s", reflect.TypeOf(interfaceValue))
            return  
        }
        value = append(value, uint8(valuetmp))
    }

    return
}

// SetScsiId sets the value of ScsiId for the instance
func (instance *HBAFCPBindingEntry2) SetPropertyScsiId(value HBAScsiID) (err error) { 
    return instance.SetProperty("ScsiId", (value))
}

// GetScsiId gets the value of ScsiId for the instance
func (instance *HBAFCPBindingEntry2) GetPropertyScsiId()(value HBAScsiID, err error) { 
    retValue, err := instance.GetProperty("ScsiId")
    if err != nil {
        return
    }
    if retValue == nil {
        // Doesn't have any value. Return empty
        return
    }
    
    valuetmp, ok := retValue.(HBAScsiID); 
    if !ok {
        err = errors.Wrapf(errors.InvalidType, " HBAScsiID is Invalid. Expected %s", reflect.TypeOf(retValue))
        return  
    }

    value = HBAScsiID(valuetmp)

    return
}

// SetType sets the value of Type for the instance
func (instance *HBAFCPBindingEntry2) SetPropertyType(value uint32) (err error) { 
    return instance.SetProperty("Type", (value))
}

// GetType gets the value of Type for the instance
func (instance *HBAFCPBindingEntry2) GetPropertyType()(value uint32, err error) { 
    retValue, err := instance.GetProperty("Type")
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

