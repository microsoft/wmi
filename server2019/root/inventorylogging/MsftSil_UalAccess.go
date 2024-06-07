// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

// 
// Author:
//      Auto Generated on 6/6/2024 using wmigen
//      Source root.InventoryLogging
//////////////////////////////////////////////
package inventorylogging
import (
 "github.com/microsoft/wmi/pkg/base/query"
cim "github.com/microsoft/wmi/pkg/wmiinstance"
 "github.com/microsoft/wmi/pkg/errors"
 "reflect"
)

// MsftSil_UalAccess struct
type MsftSil_UalAccess struct { 
	*MsftSil_Data

	// 
	ProductName string

	// 
	RoleGuid string

	// 
	RoleName string

	// 
	SampleDate string

	// 
	UniqueDeviceAccessCount uint32

	// 
	UniqueUserAccessCount uint32
}

	func NewMsftSil_UalAccessEx1(instance *cim.WmiInstance) (newInstance *MsftSil_UalAccess, err error) {tmp, err := NewMsftSil_DataEx1(instance)
		
	if err != nil { return }
	newInstance = &MsftSil_UalAccess {
	MsftSil_Data: tmp,
	}
	return
	}
	

	func NewMsftSil_UalAccessEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery ) (newInstance *MsftSil_UalAccess, err error) {tmp, err := NewMsftSil_DataEx6(hostName, wmiNamespace, userName, password, domainName, query)
		
	if err != nil { return }
	newInstance = &MsftSil_UalAccess {
	MsftSil_Data: tmp,
	}
	return
	}
	

// SetProductName sets the value of ProductName for the instance
func (instance *MsftSil_UalAccess) SetPropertyProductName(value string) (err error) { 
    return instance.SetProperty("ProductName", (value))
}

// GetProductName gets the value of ProductName for the instance
func (instance *MsftSil_UalAccess) GetPropertyProductName()(value string, err error) { 
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
func (instance *MsftSil_UalAccess) SetPropertyRoleGuid(value string) (err error) { 
    return instance.SetProperty("RoleGuid", (value))
}

// GetRoleGuid gets the value of RoleGuid for the instance
func (instance *MsftSil_UalAccess) GetPropertyRoleGuid()(value string, err error) { 
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
func (instance *MsftSil_UalAccess) SetPropertyRoleName(value string) (err error) { 
    return instance.SetProperty("RoleName", (value))
}

// GetRoleName gets the value of RoleName for the instance
func (instance *MsftSil_UalAccess) GetPropertyRoleName()(value string, err error) { 
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

// SetSampleDate sets the value of SampleDate for the instance
func (instance *MsftSil_UalAccess) SetPropertySampleDate(value string) (err error) { 
    return instance.SetProperty("SampleDate", (value))
}

// GetSampleDate gets the value of SampleDate for the instance
func (instance *MsftSil_UalAccess) GetPropertySampleDate()(value string, err error) { 
    retValue, err := instance.GetProperty("SampleDate")
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

// SetUniqueDeviceAccessCount sets the value of UniqueDeviceAccessCount for the instance
func (instance *MsftSil_UalAccess) SetPropertyUniqueDeviceAccessCount(value uint32) (err error) { 
    return instance.SetProperty("UniqueDeviceAccessCount", (value))
}

// GetUniqueDeviceAccessCount gets the value of UniqueDeviceAccessCount for the instance
func (instance *MsftSil_UalAccess) GetPropertyUniqueDeviceAccessCount()(value uint32, err error) { 
    retValue, err := instance.GetProperty("UniqueDeviceAccessCount")
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

// SetUniqueUserAccessCount sets the value of UniqueUserAccessCount for the instance
func (instance *MsftSil_UalAccess) SetPropertyUniqueUserAccessCount(value uint32) (err error) { 
    return instance.SetProperty("UniqueUserAccessCount", (value))
}

// GetUniqueUserAccessCount gets the value of UniqueUserAccessCount for the instance
func (instance *MsftSil_UalAccess) GetPropertyUniqueUserAccessCount()(value uint32, err error) { 
    retValue, err := instance.GetProperty("UniqueUserAccessCount")
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

