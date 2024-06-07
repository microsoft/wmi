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

// SystemConfig_V4_MobilePlatform struct
type SystemConfig_V4_MobilePlatform struct { 
	*SystemConfig_V4

	// 
	BspVersion string

	// 
	DeviceManufacturer string

	// 
	DeviceManufacturerDisplayName string

	// 
	DeviceModel string

	// 
	DeviceModelDisplayName string

	// 
	MobileOperator string

	// 
	SocVersion string
}

	func NewSystemConfig_V4_MobilePlatformEx1(instance *cim.WmiInstance) (newInstance *SystemConfig_V4_MobilePlatform, err error) {tmp, err := NewSystemConfig_V4Ex1(instance)
		
	if err != nil { return }
	newInstance = &SystemConfig_V4_MobilePlatform {
	SystemConfig_V4: tmp,
	}
	return
	}
	

	func NewSystemConfig_V4_MobilePlatformEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery ) (newInstance *SystemConfig_V4_MobilePlatform, err error) {tmp, err := NewSystemConfig_V4Ex6(hostName, wmiNamespace, userName, password, domainName, query)
		
	if err != nil { return }
	newInstance = &SystemConfig_V4_MobilePlatform {
	SystemConfig_V4: tmp,
	}
	return
	}
	

// SetBspVersion sets the value of BspVersion for the instance
func (instance *SystemConfig_V4_MobilePlatform) SetPropertyBspVersion(value string) (err error) { 
    return instance.SetProperty("BspVersion", (value))
}

// GetBspVersion gets the value of BspVersion for the instance
func (instance *SystemConfig_V4_MobilePlatform) GetPropertyBspVersion()(value string, err error) { 
    retValue, err := instance.GetProperty("BspVersion")
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

// SetDeviceManufacturer sets the value of DeviceManufacturer for the instance
func (instance *SystemConfig_V4_MobilePlatform) SetPropertyDeviceManufacturer(value string) (err error) { 
    return instance.SetProperty("DeviceManufacturer", (value))
}

// GetDeviceManufacturer gets the value of DeviceManufacturer for the instance
func (instance *SystemConfig_V4_MobilePlatform) GetPropertyDeviceManufacturer()(value string, err error) { 
    retValue, err := instance.GetProperty("DeviceManufacturer")
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

// SetDeviceManufacturerDisplayName sets the value of DeviceManufacturerDisplayName for the instance
func (instance *SystemConfig_V4_MobilePlatform) SetPropertyDeviceManufacturerDisplayName(value string) (err error) { 
    return instance.SetProperty("DeviceManufacturerDisplayName", (value))
}

// GetDeviceManufacturerDisplayName gets the value of DeviceManufacturerDisplayName for the instance
func (instance *SystemConfig_V4_MobilePlatform) GetPropertyDeviceManufacturerDisplayName()(value string, err error) { 
    retValue, err := instance.GetProperty("DeviceManufacturerDisplayName")
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

// SetDeviceModel sets the value of DeviceModel for the instance
func (instance *SystemConfig_V4_MobilePlatform) SetPropertyDeviceModel(value string) (err error) { 
    return instance.SetProperty("DeviceModel", (value))
}

// GetDeviceModel gets the value of DeviceModel for the instance
func (instance *SystemConfig_V4_MobilePlatform) GetPropertyDeviceModel()(value string, err error) { 
    retValue, err := instance.GetProperty("DeviceModel")
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

// SetDeviceModelDisplayName sets the value of DeviceModelDisplayName for the instance
func (instance *SystemConfig_V4_MobilePlatform) SetPropertyDeviceModelDisplayName(value string) (err error) { 
    return instance.SetProperty("DeviceModelDisplayName", (value))
}

// GetDeviceModelDisplayName gets the value of DeviceModelDisplayName for the instance
func (instance *SystemConfig_V4_MobilePlatform) GetPropertyDeviceModelDisplayName()(value string, err error) { 
    retValue, err := instance.GetProperty("DeviceModelDisplayName")
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

// SetMobileOperator sets the value of MobileOperator for the instance
func (instance *SystemConfig_V4_MobilePlatform) SetPropertyMobileOperator(value string) (err error) { 
    return instance.SetProperty("MobileOperator", (value))
}

// GetMobileOperator gets the value of MobileOperator for the instance
func (instance *SystemConfig_V4_MobilePlatform) GetPropertyMobileOperator()(value string, err error) { 
    retValue, err := instance.GetProperty("MobileOperator")
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

// SetSocVersion sets the value of SocVersion for the instance
func (instance *SystemConfig_V4_MobilePlatform) SetPropertySocVersion(value string) (err error) { 
    return instance.SetProperty("SocVersion", (value))
}

// GetSocVersion gets the value of SocVersion for the instance
func (instance *SystemConfig_V4_MobilePlatform) GetPropertySocVersion()(value string, err error) { 
    retValue, err := instance.GetProperty("SocVersion")
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

