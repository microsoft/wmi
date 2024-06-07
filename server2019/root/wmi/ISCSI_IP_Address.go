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

// ISCSI_IP_Address struct
type ISCSI_IP_Address struct { 
	*cim.WmiInstance

	// 
	IpV4Address uint32

	// 
	IpV6Address []uint8

	// 
	IpV6FlowInfo uint32

	// 
	IpV6ScopeId uint32

	// 
	TextAddress string

	// 
	Type Address_Type
}

	func NewISCSI_IP_AddressEx1(instance *cim.WmiInstance) (newInstance *ISCSI_IP_Address, err error) {tmp, err := instance, nil
		
	if err != nil { return }
	newInstance = &ISCSI_IP_Address {
	WmiInstance: tmp,
	}
	return
	}
	

	func NewISCSI_IP_AddressEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery ) (newInstance *ISCSI_IP_Address, err error) {tmp, err := instance.GetWmiInstance(hostName, wmiNamespace, userName, password, domainName, query)
		
	if err != nil { return }
	newInstance = &ISCSI_IP_Address {
	WmiInstance: tmp,
	}
	return
	}
	

// SetIpV4Address sets the value of IpV4Address for the instance
func (instance *ISCSI_IP_Address) SetPropertyIpV4Address(value uint32) (err error) { 
    return instance.SetProperty("IpV4Address", (value))
}

// GetIpV4Address gets the value of IpV4Address for the instance
func (instance *ISCSI_IP_Address) GetPropertyIpV4Address()(value uint32, err error) { 
    retValue, err := instance.GetProperty("IpV4Address")
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

// SetIpV6Address sets the value of IpV6Address for the instance
func (instance *ISCSI_IP_Address) SetPropertyIpV6Address(value []uint8) (err error) { 
    return instance.SetProperty("IpV6Address", (value))
}

// GetIpV6Address gets the value of IpV6Address for the instance
func (instance *ISCSI_IP_Address) GetPropertyIpV6Address()(value []uint8, err error) { 
    retValue, err := instance.GetProperty("IpV6Address")
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

// SetIpV6FlowInfo sets the value of IpV6FlowInfo for the instance
func (instance *ISCSI_IP_Address) SetPropertyIpV6FlowInfo(value uint32) (err error) { 
    return instance.SetProperty("IpV6FlowInfo", (value))
}

// GetIpV6FlowInfo gets the value of IpV6FlowInfo for the instance
func (instance *ISCSI_IP_Address) GetPropertyIpV6FlowInfo()(value uint32, err error) { 
    retValue, err := instance.GetProperty("IpV6FlowInfo")
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

// SetIpV6ScopeId sets the value of IpV6ScopeId for the instance
func (instance *ISCSI_IP_Address) SetPropertyIpV6ScopeId(value uint32) (err error) { 
    return instance.SetProperty("IpV6ScopeId", (value))
}

// GetIpV6ScopeId gets the value of IpV6ScopeId for the instance
func (instance *ISCSI_IP_Address) GetPropertyIpV6ScopeId()(value uint32, err error) { 
    retValue, err := instance.GetProperty("IpV6ScopeId")
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

// SetTextAddress sets the value of TextAddress for the instance
func (instance *ISCSI_IP_Address) SetPropertyTextAddress(value string) (err error) { 
    return instance.SetProperty("TextAddress", (value))
}

// GetTextAddress gets the value of TextAddress for the instance
func (instance *ISCSI_IP_Address) GetPropertyTextAddress()(value string, err error) { 
    retValue, err := instance.GetProperty("TextAddress")
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

// SetType sets the value of Type for the instance
func (instance *ISCSI_IP_Address) SetPropertyType(value Address_Type) (err error) { 
    return instance.SetProperty("Type", (value))
}

// GetType gets the value of Type for the instance
func (instance *ISCSI_IP_Address) GetPropertyType()(value Address_Type, err error) { 
    retValue, err := instance.GetProperty("Type")
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

    value = Address_Type(valuetmp)

    return
}

