// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

// 
// Author:
//      Auto Generated on 6/6/2024 using wmigen
//      Source root.directory.LDAP
//////////////////////////////////////////////
package ldap
import (
 "github.com/microsoft/wmi/pkg/base/query"
cim "github.com/microsoft/wmi/pkg/wmiinstance"
 "github.com/microsoft/wmi/pkg/errors"
 "reflect"
)

// ads_msds_device struct
type ads_msds_device struct { 
	*ds_top

	// 
	DS_altSecurityIdentities []string

	// 
	DS_msDS_ApproximateLastLogonTimeStamp int64

	// 
	DS_msDS_CloudIsManaged bool

	// 
	DS_msDS_ComputerSID Uint8Array

	// 
	DS_msDS_DeviceID Uint8Array

	// 
	DS_msDS_DeviceMDMStatus string

	// 
	DS_msDS_DeviceObjectVersion int32

	// 
	DS_msDS_DeviceOSType string

	// 
	DS_msDS_DeviceOSVersion string

	// 
	DS_msDS_DevicePhysicalIDs []string

	// 
	DS_msDS_DeviceTrustType int32

	// 
	DS_msDS_IsCompliant bool

	// 
	DS_msDS_IsEnabled bool

	// 
	DS_msDS_IsManaged bool

	// 
	DS_msDS_KeyCredentialLink []DN_With_Binary

	// 
	DS_msDS_RegisteredOwner Uint8Array

	// 
	DS_msDS_RegisteredUsers []Uint8Array
}

	func Newads_msds_deviceEx1(instance *cim.WmiInstance) (newInstance *ads_msds_device, err error) {tmp, err := Newds_topEx1(instance)
		
	if err != nil { return }
	newInstance = &ads_msds_device {
	ds_top: tmp,
	}
	return
	}
	

	func Newads_msds_deviceEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery ) (newInstance *ads_msds_device, err error) {tmp, err := Newds_topEx6(hostName, wmiNamespace, userName, password, domainName, query)
		
	if err != nil { return }
	newInstance = &ads_msds_device {
	ds_top: tmp,
	}
	return
	}
	

// SetDS_altSecurityIdentities sets the value of DS_altSecurityIdentities for the instance
func (instance *ads_msds_device) SetPropertyDS_altSecurityIdentities(value []string) (err error) { 
    return instance.SetProperty("DS_altSecurityIdentities", (value))
}

// GetDS_altSecurityIdentities gets the value of DS_altSecurityIdentities for the instance
func (instance *ads_msds_device) GetPropertyDS_altSecurityIdentities()(value []string, err error) { 
    retValue, err := instance.GetProperty("DS_altSecurityIdentities")
    if err != nil {
        return
    }
    if retValue == nil {
        // Doesn't have any value. Return empty
        return
    }
    
    for _, interfaceValue := range retValue.([]interface{}) {
        valuetmp, ok := interfaceValue.(string); 
        if !ok {
            err = errors.Wrapf(errors.InvalidType, " string is Invalid. Expected %s", reflect.TypeOf(interfaceValue))
            return  
        }
        value = append(value, string(valuetmp))
    }

    return
}

// SetDS_msDS_ApproximateLastLogonTimeStamp sets the value of DS_msDS_ApproximateLastLogonTimeStamp for the instance
func (instance *ads_msds_device) SetPropertyDS_msDS_ApproximateLastLogonTimeStamp(value int64) (err error) { 
    return instance.SetProperty("DS_msDS_ApproximateLastLogonTimeStamp", (value))
}

// GetDS_msDS_ApproximateLastLogonTimeStamp gets the value of DS_msDS_ApproximateLastLogonTimeStamp for the instance
func (instance *ads_msds_device) GetPropertyDS_msDS_ApproximateLastLogonTimeStamp()(value int64, err error) { 
    retValue, err := instance.GetProperty("DS_msDS_ApproximateLastLogonTimeStamp")
    if err != nil {
        return
    }
    if retValue == nil {
        // Doesn't have any value. Return empty
        return
    }
    
    valuetmp, ok := retValue.(int64); 
    if !ok {
        err = errors.Wrapf(errors.InvalidType, " int64 is Invalid. Expected %s", reflect.TypeOf(retValue))
        return  
    }

    value = int64(valuetmp)

    return
}

// SetDS_msDS_CloudIsManaged sets the value of DS_msDS_CloudIsManaged for the instance
func (instance *ads_msds_device) SetPropertyDS_msDS_CloudIsManaged(value bool) (err error) { 
    return instance.SetProperty("DS_msDS_CloudIsManaged", (value))
}

// GetDS_msDS_CloudIsManaged gets the value of DS_msDS_CloudIsManaged for the instance
func (instance *ads_msds_device) GetPropertyDS_msDS_CloudIsManaged()(value bool, err error) { 
    retValue, err := instance.GetProperty("DS_msDS_CloudIsManaged")
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

// SetDS_msDS_ComputerSID sets the value of DS_msDS_ComputerSID for the instance
func (instance *ads_msds_device) SetPropertyDS_msDS_ComputerSID(value Uint8Array) (err error) { 
    return instance.SetProperty("DS_msDS_ComputerSID", (value))
}

// GetDS_msDS_ComputerSID gets the value of DS_msDS_ComputerSID for the instance
func (instance *ads_msds_device) GetPropertyDS_msDS_ComputerSID()(value Uint8Array, err error) { 
    retValue, err := instance.GetProperty("DS_msDS_ComputerSID")
    if err != nil {
        return
    }
    if retValue == nil {
        // Doesn't have any value. Return empty
        return
    }
    
    valuetmp, ok := retValue.(Uint8Array); 
    if !ok {
        err = errors.Wrapf(errors.InvalidType, " Uint8Array is Invalid. Expected %s", reflect.TypeOf(retValue))
        return  
    }

    value = Uint8Array(valuetmp)

    return
}

// SetDS_msDS_DeviceID sets the value of DS_msDS_DeviceID for the instance
func (instance *ads_msds_device) SetPropertyDS_msDS_DeviceID(value Uint8Array) (err error) { 
    return instance.SetProperty("DS_msDS_DeviceID", (value))
}

// GetDS_msDS_DeviceID gets the value of DS_msDS_DeviceID for the instance
func (instance *ads_msds_device) GetPropertyDS_msDS_DeviceID()(value Uint8Array, err error) { 
    retValue, err := instance.GetProperty("DS_msDS_DeviceID")
    if err != nil {
        return
    }
    if retValue == nil {
        // Doesn't have any value. Return empty
        return
    }
    
    valuetmp, ok := retValue.(Uint8Array); 
    if !ok {
        err = errors.Wrapf(errors.InvalidType, " Uint8Array is Invalid. Expected %s", reflect.TypeOf(retValue))
        return  
    }

    value = Uint8Array(valuetmp)

    return
}

// SetDS_msDS_DeviceMDMStatus sets the value of DS_msDS_DeviceMDMStatus for the instance
func (instance *ads_msds_device) SetPropertyDS_msDS_DeviceMDMStatus(value string) (err error) { 
    return instance.SetProperty("DS_msDS_DeviceMDMStatus", (value))
}

// GetDS_msDS_DeviceMDMStatus gets the value of DS_msDS_DeviceMDMStatus for the instance
func (instance *ads_msds_device) GetPropertyDS_msDS_DeviceMDMStatus()(value string, err error) { 
    retValue, err := instance.GetProperty("DS_msDS_DeviceMDMStatus")
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

// SetDS_msDS_DeviceObjectVersion sets the value of DS_msDS_DeviceObjectVersion for the instance
func (instance *ads_msds_device) SetPropertyDS_msDS_DeviceObjectVersion(value int32) (err error) { 
    return instance.SetProperty("DS_msDS_DeviceObjectVersion", (value))
}

// GetDS_msDS_DeviceObjectVersion gets the value of DS_msDS_DeviceObjectVersion for the instance
func (instance *ads_msds_device) GetPropertyDS_msDS_DeviceObjectVersion()(value int32, err error) { 
    retValue, err := instance.GetProperty("DS_msDS_DeviceObjectVersion")
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

    value = int32(valuetmp)

    return
}

// SetDS_msDS_DeviceOSType sets the value of DS_msDS_DeviceOSType for the instance
func (instance *ads_msds_device) SetPropertyDS_msDS_DeviceOSType(value string) (err error) { 
    return instance.SetProperty("DS_msDS_DeviceOSType", (value))
}

// GetDS_msDS_DeviceOSType gets the value of DS_msDS_DeviceOSType for the instance
func (instance *ads_msds_device) GetPropertyDS_msDS_DeviceOSType()(value string, err error) { 
    retValue, err := instance.GetProperty("DS_msDS_DeviceOSType")
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

// SetDS_msDS_DeviceOSVersion sets the value of DS_msDS_DeviceOSVersion for the instance
func (instance *ads_msds_device) SetPropertyDS_msDS_DeviceOSVersion(value string) (err error) { 
    return instance.SetProperty("DS_msDS_DeviceOSVersion", (value))
}

// GetDS_msDS_DeviceOSVersion gets the value of DS_msDS_DeviceOSVersion for the instance
func (instance *ads_msds_device) GetPropertyDS_msDS_DeviceOSVersion()(value string, err error) { 
    retValue, err := instance.GetProperty("DS_msDS_DeviceOSVersion")
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

// SetDS_msDS_DevicePhysicalIDs sets the value of DS_msDS_DevicePhysicalIDs for the instance
func (instance *ads_msds_device) SetPropertyDS_msDS_DevicePhysicalIDs(value []string) (err error) { 
    return instance.SetProperty("DS_msDS_DevicePhysicalIDs", (value))
}

// GetDS_msDS_DevicePhysicalIDs gets the value of DS_msDS_DevicePhysicalIDs for the instance
func (instance *ads_msds_device) GetPropertyDS_msDS_DevicePhysicalIDs()(value []string, err error) { 
    retValue, err := instance.GetProperty("DS_msDS_DevicePhysicalIDs")
    if err != nil {
        return
    }
    if retValue == nil {
        // Doesn't have any value. Return empty
        return
    }
    
    for _, interfaceValue := range retValue.([]interface{}) {
        valuetmp, ok := interfaceValue.(string); 
        if !ok {
            err = errors.Wrapf(errors.InvalidType, " string is Invalid. Expected %s", reflect.TypeOf(interfaceValue))
            return  
        }
        value = append(value, string(valuetmp))
    }

    return
}

// SetDS_msDS_DeviceTrustType sets the value of DS_msDS_DeviceTrustType for the instance
func (instance *ads_msds_device) SetPropertyDS_msDS_DeviceTrustType(value int32) (err error) { 
    return instance.SetProperty("DS_msDS_DeviceTrustType", (value))
}

// GetDS_msDS_DeviceTrustType gets the value of DS_msDS_DeviceTrustType for the instance
func (instance *ads_msds_device) GetPropertyDS_msDS_DeviceTrustType()(value int32, err error) { 
    retValue, err := instance.GetProperty("DS_msDS_DeviceTrustType")
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

    value = int32(valuetmp)

    return
}

// SetDS_msDS_IsCompliant sets the value of DS_msDS_IsCompliant for the instance
func (instance *ads_msds_device) SetPropertyDS_msDS_IsCompliant(value bool) (err error) { 
    return instance.SetProperty("DS_msDS_IsCompliant", (value))
}

// GetDS_msDS_IsCompliant gets the value of DS_msDS_IsCompliant for the instance
func (instance *ads_msds_device) GetPropertyDS_msDS_IsCompliant()(value bool, err error) { 
    retValue, err := instance.GetProperty("DS_msDS_IsCompliant")
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

// SetDS_msDS_IsEnabled sets the value of DS_msDS_IsEnabled for the instance
func (instance *ads_msds_device) SetPropertyDS_msDS_IsEnabled(value bool) (err error) { 
    return instance.SetProperty("DS_msDS_IsEnabled", (value))
}

// GetDS_msDS_IsEnabled gets the value of DS_msDS_IsEnabled for the instance
func (instance *ads_msds_device) GetPropertyDS_msDS_IsEnabled()(value bool, err error) { 
    retValue, err := instance.GetProperty("DS_msDS_IsEnabled")
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

// SetDS_msDS_IsManaged sets the value of DS_msDS_IsManaged for the instance
func (instance *ads_msds_device) SetPropertyDS_msDS_IsManaged(value bool) (err error) { 
    return instance.SetProperty("DS_msDS_IsManaged", (value))
}

// GetDS_msDS_IsManaged gets the value of DS_msDS_IsManaged for the instance
func (instance *ads_msds_device) GetPropertyDS_msDS_IsManaged()(value bool, err error) { 
    retValue, err := instance.GetProperty("DS_msDS_IsManaged")
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

// SetDS_msDS_KeyCredentialLink sets the value of DS_msDS_KeyCredentialLink for the instance
func (instance *ads_msds_device) SetPropertyDS_msDS_KeyCredentialLink(value []DN_With_Binary) (err error) { 
    return instance.SetProperty("DS_msDS_KeyCredentialLink", (value))
}

// GetDS_msDS_KeyCredentialLink gets the value of DS_msDS_KeyCredentialLink for the instance
func (instance *ads_msds_device) GetPropertyDS_msDS_KeyCredentialLink()(value []DN_With_Binary, err error) { 
    retValue, err := instance.GetProperty("DS_msDS_KeyCredentialLink")
    if err != nil {
        return
    }
    if retValue == nil {
        // Doesn't have any value. Return empty
        return
    }
    
    for _, interfaceValue := range retValue.([]interface{}) {
        valuetmp, ok := interfaceValue.(DN_With_Binary); 
        if !ok {
            err = errors.Wrapf(errors.InvalidType, " DN_With_Binary is Invalid. Expected %s", reflect.TypeOf(interfaceValue))
            return  
        }
        value = append(value, DN_With_Binary(valuetmp))
    }

    return
}

// SetDS_msDS_RegisteredOwner sets the value of DS_msDS_RegisteredOwner for the instance
func (instance *ads_msds_device) SetPropertyDS_msDS_RegisteredOwner(value Uint8Array) (err error) { 
    return instance.SetProperty("DS_msDS_RegisteredOwner", (value))
}

// GetDS_msDS_RegisteredOwner gets the value of DS_msDS_RegisteredOwner for the instance
func (instance *ads_msds_device) GetPropertyDS_msDS_RegisteredOwner()(value Uint8Array, err error) { 
    retValue, err := instance.GetProperty("DS_msDS_RegisteredOwner")
    if err != nil {
        return
    }
    if retValue == nil {
        // Doesn't have any value. Return empty
        return
    }
    
    valuetmp, ok := retValue.(Uint8Array); 
    if !ok {
        err = errors.Wrapf(errors.InvalidType, " Uint8Array is Invalid. Expected %s", reflect.TypeOf(retValue))
        return  
    }

    value = Uint8Array(valuetmp)

    return
}

// SetDS_msDS_RegisteredUsers sets the value of DS_msDS_RegisteredUsers for the instance
func (instance *ads_msds_device) SetPropertyDS_msDS_RegisteredUsers(value []Uint8Array) (err error) { 
    return instance.SetProperty("DS_msDS_RegisteredUsers", (value))
}

// GetDS_msDS_RegisteredUsers gets the value of DS_msDS_RegisteredUsers for the instance
func (instance *ads_msds_device) GetPropertyDS_msDS_RegisteredUsers()(value []Uint8Array, err error) { 
    retValue, err := instance.GetProperty("DS_msDS_RegisteredUsers")
    if err != nil {
        return
    }
    if retValue == nil {
        // Doesn't have any value. Return empty
        return
    }
    
    for _, interfaceValue := range retValue.([]interface{}) {
        valuetmp, ok := interfaceValue.(Uint8Array); 
        if !ok {
            err = errors.Wrapf(errors.InvalidType, " Uint8Array is Invalid. Expected %s", reflect.TypeOf(interfaceValue))
            return  
        }
        value = append(value, Uint8Array(valuetmp))
    }

    return
}

