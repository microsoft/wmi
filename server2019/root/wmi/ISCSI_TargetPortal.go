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

// ISCSI_TargetPortal struct
type ISCSI_TargetPortal struct { 
	*cim.WmiInstance

	// 
	Address ISCSI_IP_Address

	// 
	Reserved uint32

	// 
	Socket uint16
}

	func NewISCSI_TargetPortalEx1(instance *cim.WmiInstance) (newInstance *ISCSI_TargetPortal, err error) {tmp, err := instance, nil
		
	if err != nil { return }
	newInstance = &ISCSI_TargetPortal {
	WmiInstance: tmp,
	}
	return
	}
	

	func NewISCSI_TargetPortalEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery ) (newInstance *ISCSI_TargetPortal, err error) {tmp, err := instance.GetWmiInstance(hostName, wmiNamespace, userName, password, domainName, query)
		
	if err != nil { return }
	newInstance = &ISCSI_TargetPortal {
	WmiInstance: tmp,
	}
	return
	}
	

// SetAddress sets the value of Address for the instance
func (instance *ISCSI_TargetPortal) SetPropertyAddress(value ISCSI_IP_Address) (err error) { 
    return instance.SetProperty("Address", (value))
}

// GetAddress gets the value of Address for the instance
func (instance *ISCSI_TargetPortal) GetPropertyAddress()(value ISCSI_IP_Address, err error) { 
    retValue, err := instance.GetProperty("Address")
    if err != nil {
        return
    }
    if retValue == nil {
        // Doesn't have any value. Return empty
        return
    }
    
    valuetmp, ok := retValue.(ISCSI_IP_Address); 
    if !ok {
        err = errors.Wrapf(errors.InvalidType, " ISCSI_IP_Address is Invalid. Expected %s", reflect.TypeOf(retValue))
        return  
    }

    value = ISCSI_IP_Address(valuetmp)

    return
}

// SetReserved sets the value of Reserved for the instance
func (instance *ISCSI_TargetPortal) SetPropertyReserved(value uint32) (err error) { 
    return instance.SetProperty("Reserved", (value))
}

// GetReserved gets the value of Reserved for the instance
func (instance *ISCSI_TargetPortal) GetPropertyReserved()(value uint32, err error) { 
    retValue, err := instance.GetProperty("Reserved")
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

// SetSocket sets the value of Socket for the instance
func (instance *ISCSI_TargetPortal) SetPropertySocket(value uint16) (err error) { 
    return instance.SetProperty("Socket", (value))
}

// GetSocket gets the value of Socket for the instance
func (instance *ISCSI_TargetPortal) GetPropertySocket()(value uint16, err error) { 
    retValue, err := instance.GetProperty("Socket")
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

