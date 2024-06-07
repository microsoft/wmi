// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

// 
// Author:
//      Auto Generated on 6/6/2024 using wmigen
//      Source root.AccessLogging
//////////////////////////////////////////////
package accesslogging
import (
 "github.com/microsoft/wmi/pkg/base/query"
cim "github.com/microsoft/wmi/pkg/wmiinstance"
 "github.com/microsoft/wmi/pkg/base/instance"
 "github.com/microsoft/wmi/pkg/errors"
 "reflect"
)

// MsftUal_DailyDeviceAccess struct
type MsftUal_DailyDeviceAccess struct { 
	*cim.WmiInstance

	// The number of accesses of a role, or installed product, on the local server from a unique client device.
	AccessCount uint32

	// The date that a device accessed a role, or installed product, on the local server.
	AccessDate string

	// The IP address of a client device that accompanies the UAL payload from installed roles and products.
	IPAddress string

	// The name of the software parent product, or product line, that is providing User Access Logging data. This is also associated with a RoleName, and a RoleGuid.
	ProductName string

	// The UAL assigned, or registered, GUID representing the server role, or installed product. When a role or product reports its UAL data, this GUID accompanies the payload.
	RoleGuid string

	// The name of the role, component, or sub-product that is providing User Access Logging data. This is also associated with a ProductName, and a RoleGuid.
	RoleName string
}

	func NewMsftUal_DailyDeviceAccessEx1(instance *cim.WmiInstance) (newInstance *MsftUal_DailyDeviceAccess, err error) {tmp, err := instance, nil
		
	if err != nil { return }
	newInstance = &MsftUal_DailyDeviceAccess {
	WmiInstance: tmp,
	}
	return
	}
	

	func NewMsftUal_DailyDeviceAccessEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery ) (newInstance *MsftUal_DailyDeviceAccess, err error) {tmp, err := instance.GetWmiInstance(hostName, wmiNamespace, userName, password, domainName, query)
		
	if err != nil { return }
	newInstance = &MsftUal_DailyDeviceAccess {
	WmiInstance: tmp,
	}
	return
	}
	

// SetAccessCount sets the value of AccessCount for the instance
func (instance *MsftUal_DailyDeviceAccess) SetPropertyAccessCount(value uint32) (err error) { 
    return instance.SetProperty("AccessCount", (value))
}

// GetAccessCount gets the value of AccessCount for the instance
func (instance *MsftUal_DailyDeviceAccess) GetPropertyAccessCount()(value uint32, err error) { 
    retValue, err := instance.GetProperty("AccessCount")
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

// SetAccessDate sets the value of AccessDate for the instance
func (instance *MsftUal_DailyDeviceAccess) SetPropertyAccessDate(value string) (err error) { 
    return instance.SetProperty("AccessDate", (value))
}

// GetAccessDate gets the value of AccessDate for the instance
func (instance *MsftUal_DailyDeviceAccess) GetPropertyAccessDate()(value string, err error) { 
    retValue, err := instance.GetProperty("AccessDate")
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

// SetIPAddress sets the value of IPAddress for the instance
func (instance *MsftUal_DailyDeviceAccess) SetPropertyIPAddress(value string) (err error) { 
    return instance.SetProperty("IPAddress", (value))
}

// GetIPAddress gets the value of IPAddress for the instance
func (instance *MsftUal_DailyDeviceAccess) GetPropertyIPAddress()(value string, err error) { 
    retValue, err := instance.GetProperty("IPAddress")
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

// SetProductName sets the value of ProductName for the instance
func (instance *MsftUal_DailyDeviceAccess) SetPropertyProductName(value string) (err error) { 
    return instance.SetProperty("ProductName", (value))
}

// GetProductName gets the value of ProductName for the instance
func (instance *MsftUal_DailyDeviceAccess) GetPropertyProductName()(value string, err error) { 
    retValue, err := instance.GetProperty("ProductName")
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

// SetRoleGuid sets the value of RoleGuid for the instance
func (instance *MsftUal_DailyDeviceAccess) SetPropertyRoleGuid(value string) (err error) { 
    return instance.SetProperty("RoleGuid", (value))
}

// GetRoleGuid gets the value of RoleGuid for the instance
func (instance *MsftUal_DailyDeviceAccess) GetPropertyRoleGuid()(value string, err error) { 
    retValue, err := instance.GetProperty("RoleGuid")
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

// SetRoleName sets the value of RoleName for the instance
func (instance *MsftUal_DailyDeviceAccess) SetPropertyRoleName(value string) (err error) { 
    return instance.SetProperty("RoleName", (value))
}

// GetRoleName gets the value of RoleName for the instance
func (instance *MsftUal_DailyDeviceAccess) GetPropertyRoleName()(value string, err error) { 
    retValue, err := instance.GetProperty("RoleName")
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

