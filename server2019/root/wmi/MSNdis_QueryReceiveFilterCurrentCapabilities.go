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

// MSNdis_QueryReceiveFilterCurrentCapabilities struct
type MSNdis_QueryReceiveFilterCurrentCapabilities struct { 
	*MSNdis

	// 
	Active bool

	// 
	InstanceName string
}

	func NewMSNdis_QueryReceiveFilterCurrentCapabilitiesEx1(instance *cim.WmiInstance) (newInstance *MSNdis_QueryReceiveFilterCurrentCapabilities, err error) {tmp, err := NewMSNdisEx1(instance)
		
	if err != nil { return }
	newInstance = &MSNdis_QueryReceiveFilterCurrentCapabilities {
	MSNdis: tmp,
	}
	return
	}
	

	func NewMSNdis_QueryReceiveFilterCurrentCapabilitiesEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery ) (newInstance *MSNdis_QueryReceiveFilterCurrentCapabilities, err error) {tmp, err := NewMSNdisEx6(hostName, wmiNamespace, userName, password, domainName, query)
		
	if err != nil { return }
	newInstance = &MSNdis_QueryReceiveFilterCurrentCapabilities {
	MSNdis: tmp,
	}
	return
	}
	

// SetActive sets the value of Active for the instance
func (instance *MSNdis_QueryReceiveFilterCurrentCapabilities) SetPropertyActive(value bool) (err error) { 
    return instance.SetProperty("Active", (value))
}

// GetActive gets the value of Active for the instance
func (instance *MSNdis_QueryReceiveFilterCurrentCapabilities) GetPropertyActive()(value bool, err error) { 
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
func (instance *MSNdis_QueryReceiveFilterCurrentCapabilities) SetPropertyInstanceName(value string) (err error) { 
    return instance.SetProperty("InstanceName", (value))
}

// GetInstanceName gets the value of InstanceName for the instance
func (instance *MSNdis_QueryReceiveFilterCurrentCapabilities) GetPropertyInstanceName()(value string, err error) { 
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

// 

// <param name="Header" type="MSNdis_WmiMethodHeader "></param>

// <param name="ReceiveFilterCurrentCapabilities" type="MSNdis_ReceiveFilterCapabilities "></param>
func (instance *MSNdis_QueryReceiveFilterCurrentCapabilities) WmiQueryReceiveFilterCurrentCapabilities( /* IN */ Header MSNdis_WmiMethodHeader,
 /* OUT */ ReceiveFilterCurrentCapabilities MSNdis_ReceiveFilterCapabilities) (err error) {_, err = instance.InvokeMethod("WmiQueryReceiveFilterCurrentCapabilities" , Header)
	if err != nil { return }
	return
	
}


