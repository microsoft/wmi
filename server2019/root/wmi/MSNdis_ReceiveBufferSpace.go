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

// MSNdis_ReceiveBufferSpace struct
type MSNdis_ReceiveBufferSpace struct { 
	*MSNdis

	// 
	Active bool

	// 
	InstanceName string

	// 
	NdisReceiveBufferSpace uint32
}

	func NewMSNdis_ReceiveBufferSpaceEx1(instance *cim.WmiInstance) (newInstance *MSNdis_ReceiveBufferSpace, err error) {tmp, err := NewMSNdisEx1(instance)
		
	if err != nil { return }
	newInstance = &MSNdis_ReceiveBufferSpace {
	MSNdis: tmp,
	}
	return
	}
	

	func NewMSNdis_ReceiveBufferSpaceEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery ) (newInstance *MSNdis_ReceiveBufferSpace, err error) {tmp, err := NewMSNdisEx6(hostName, wmiNamespace, userName, password, domainName, query)
		
	if err != nil { return }
	newInstance = &MSNdis_ReceiveBufferSpace {
	MSNdis: tmp,
	}
	return
	}
	

// SetActive sets the value of Active for the instance
func (instance *MSNdis_ReceiveBufferSpace) SetPropertyActive(value bool) (err error) { 
    return instance.SetProperty("Active", (value))
}

// GetActive gets the value of Active for the instance
func (instance *MSNdis_ReceiveBufferSpace) GetPropertyActive()(value bool, err error) { 
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
func (instance *MSNdis_ReceiveBufferSpace) SetPropertyInstanceName(value string) (err error) { 
    return instance.SetProperty("InstanceName", (value))
}

// GetInstanceName gets the value of InstanceName for the instance
func (instance *MSNdis_ReceiveBufferSpace) GetPropertyInstanceName()(value string, err error) { 
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

// SetNdisReceiveBufferSpace sets the value of NdisReceiveBufferSpace for the instance
func (instance *MSNdis_ReceiveBufferSpace) SetPropertyNdisReceiveBufferSpace(value uint32) (err error) { 
    return instance.SetProperty("NdisReceiveBufferSpace", (value))
}

// GetNdisReceiveBufferSpace gets the value of NdisReceiveBufferSpace for the instance
func (instance *MSNdis_ReceiveBufferSpace) GetPropertyNdisReceiveBufferSpace()(value uint32, err error) { 
    retValue, err := instance.GetProperty("NdisReceiveBufferSpace")
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

