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

// MSNdis_NetworkLinkSpeed struct
type MSNdis_NetworkLinkSpeed struct { 
	*MSNdis

	// 
	Inbound uint32

	// 
	Outbound uint32
}

	func NewMSNdis_NetworkLinkSpeedEx1(instance *cim.WmiInstance) (newInstance *MSNdis_NetworkLinkSpeed, err error) {tmp, err := NewMSNdisEx1(instance)
		
	if err != nil { return }
	newInstance = &MSNdis_NetworkLinkSpeed {
	MSNdis: tmp,
	}
	return
	}
	

	func NewMSNdis_NetworkLinkSpeedEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery ) (newInstance *MSNdis_NetworkLinkSpeed, err error) {tmp, err := NewMSNdisEx6(hostName, wmiNamespace, userName, password, domainName, query)
		
	if err != nil { return }
	newInstance = &MSNdis_NetworkLinkSpeed {
	MSNdis: tmp,
	}
	return
	}
	

// SetInbound sets the value of Inbound for the instance
func (instance *MSNdis_NetworkLinkSpeed) SetPropertyInbound(value uint32) (err error) { 
    return instance.SetProperty("Inbound", (value))
}

// GetInbound gets the value of Inbound for the instance
func (instance *MSNdis_NetworkLinkSpeed) GetPropertyInbound()(value uint32, err error) { 
    retValue, err := instance.GetProperty("Inbound")
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

// SetOutbound sets the value of Outbound for the instance
func (instance *MSNdis_NetworkLinkSpeed) SetPropertyOutbound(value uint32) (err error) { 
    return instance.SetProperty("Outbound", (value))
}

// GetOutbound gets the value of Outbound for the instance
func (instance *MSNdis_NetworkLinkSpeed) GetPropertyOutbound()(value uint32, err error) { 
    retValue, err := instance.GetProperty("Outbound")
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

