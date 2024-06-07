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

// MSNdis_FddiDownstreamNodeLong struct
type MSNdis_FddiDownstreamNodeLong struct { 
	*MSNdis

	// 
	Active bool

	// 
	InstanceName string

	// 
	NdisFddiDownstreamNodeLong MSNdis_NetworkAddress
}

	func NewMSNdis_FddiDownstreamNodeLongEx1(instance *cim.WmiInstance) (newInstance *MSNdis_FddiDownstreamNodeLong, err error) {tmp, err := NewMSNdisEx1(instance)
		
	if err != nil { return }
	newInstance = &MSNdis_FddiDownstreamNodeLong {
	MSNdis: tmp,
	}
	return
	}
	

	func NewMSNdis_FddiDownstreamNodeLongEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery ) (newInstance *MSNdis_FddiDownstreamNodeLong, err error) {tmp, err := NewMSNdisEx6(hostName, wmiNamespace, userName, password, domainName, query)
		
	if err != nil { return }
	newInstance = &MSNdis_FddiDownstreamNodeLong {
	MSNdis: tmp,
	}
	return
	}
	

// SetActive sets the value of Active for the instance
func (instance *MSNdis_FddiDownstreamNodeLong) SetPropertyActive(value bool) (err error) { 
    return instance.SetProperty("Active", (value))
}

// GetActive gets the value of Active for the instance
func (instance *MSNdis_FddiDownstreamNodeLong) GetPropertyActive()(value bool, err error) { 
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
func (instance *MSNdis_FddiDownstreamNodeLong) SetPropertyInstanceName(value string) (err error) { 
    return instance.SetProperty("InstanceName", (value))
}

// GetInstanceName gets the value of InstanceName for the instance
func (instance *MSNdis_FddiDownstreamNodeLong) GetPropertyInstanceName()(value string, err error) { 
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

// SetNdisFddiDownstreamNodeLong sets the value of NdisFddiDownstreamNodeLong for the instance
func (instance *MSNdis_FddiDownstreamNodeLong) SetPropertyNdisFddiDownstreamNodeLong(value MSNdis_NetworkAddress) (err error) { 
    return instance.SetProperty("NdisFddiDownstreamNodeLong", (value))
}

// GetNdisFddiDownstreamNodeLong gets the value of NdisFddiDownstreamNodeLong for the instance
func (instance *MSNdis_FddiDownstreamNodeLong) GetPropertyNdisFddiDownstreamNodeLong()(value MSNdis_NetworkAddress, err error) { 
    retValue, err := instance.GetProperty("NdisFddiDownstreamNodeLong")
    if err != nil {
        return
    }
    if retValue == nil {
        // Doesn't have any value. Return empty
        return
    }
    
    valuetmp, ok := retValue.(MSNdis_NetworkAddress); 
    if !ok {
        err = errors.Wrapf(errors.InvalidType, " MSNdis_NetworkAddress is Invalid. Expected %s", reflect.TypeOf(retValue))
        return  
    }

    value = MSNdis_NetworkAddress(valuetmp)

    return
}

