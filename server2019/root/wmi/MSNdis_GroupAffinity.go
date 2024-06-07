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
 "github.com/microsoft/wmi/pkg/errors"
 "reflect"
)

// MSNdis_GroupAffinity struct
type MSNdis_GroupAffinity struct { 
	*MSNdis

	// 
	Group uint16

	// 
	Mask uint64

	// 
	Reserved []uint16
}

	func NewMSNdis_GroupAffinityEx1(instance *cim.WmiInstance) (newInstance *MSNdis_GroupAffinity, err error) {tmp, err := NewMSNdisEx1(instance)
		
	if err != nil { return }
	newInstance = &MSNdis_GroupAffinity {
	MSNdis: tmp,
	}
	return
	}
	

	func NewMSNdis_GroupAffinityEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery ) (newInstance *MSNdis_GroupAffinity, err error) {tmp, err := NewMSNdisEx6(hostName, wmiNamespace, userName, password, domainName, query)
		
	if err != nil { return }
	newInstance = &MSNdis_GroupAffinity {
	MSNdis: tmp,
	}
	return
	}
	

// SetGroup sets the value of Group for the instance
func (instance *MSNdis_GroupAffinity) SetPropertyGroup(value uint16) (err error) { 
    return instance.SetProperty("Group", (value))
}

// GetGroup gets the value of Group for the instance
func (instance *MSNdis_GroupAffinity) GetPropertyGroup()(value uint16, err error) { 
    retValue, err := instance.GetProperty("Group")
    if err != nil {
        return
    }
    if retValue == nil {
        // Doesn't have any value. Return empty
        return
    }
    
    valuetmp, ok := retValue.(uint16); 
    if !ok {
        err = errors.Wrapf(errors.InvalidType, " uint16 is Invalid. Expected %s", reflect.TypeOf(retValue))
        return  
    }

    value = uint16(valuetmp)

    return
}

// SetMask sets the value of Mask for the instance
func (instance *MSNdis_GroupAffinity) SetPropertyMask(value uint64) (err error) { 
    return instance.SetProperty("Mask", (value))
}

// GetMask gets the value of Mask for the instance
func (instance *MSNdis_GroupAffinity) GetPropertyMask()(value uint64, err error) { 
    retValue, err := instance.GetProperty("Mask")
    if err != nil {
        return
    }
    if retValue == nil {
        // Doesn't have any value. Return empty
        return
    }
    
    valuetmp, ok := retValue.(uint64); 
    if !ok {
        err = errors.Wrapf(errors.InvalidType, " uint64 is Invalid. Expected %s", reflect.TypeOf(retValue))
        return  
    }

    value = uint64(valuetmp)

    return
}

// SetReserved sets the value of Reserved for the instance
func (instance *MSNdis_GroupAffinity) SetPropertyReserved(value []uint16) (err error) { 
    return instance.SetProperty("Reserved", (value))
}

// GetReserved gets the value of Reserved for the instance
func (instance *MSNdis_GroupAffinity) GetPropertyReserved()(value []uint16, err error) { 
    retValue, err := instance.GetProperty("Reserved")
    if err != nil {
        return
    }
    if retValue == nil {
        // Doesn't have any value. Return empty
        return
    }
    
    for _, interfaceValue := range retValue.([]interface{}) {
        valuetmp, ok := interfaceValue.(uint16); 
        if !ok {
            err = errors.Wrapf(errors.InvalidType, " uint16 is Invalid. Expected %s", reflect.TypeOf(interfaceValue))
            return  
        }
        value = append(value, uint16(valuetmp))
    }

    return
}

