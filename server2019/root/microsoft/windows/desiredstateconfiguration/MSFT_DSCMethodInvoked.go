// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

// 
// Author:
//      Auto Generated on 6/6/2024 using wmigen
//      Source root.Microsoft.Windows.DesiredStateConfiguration
//////////////////////////////////////////////
package desiredstateconfiguration
import (
 "github.com/microsoft/wmi/pkg/base/query"
cim "github.com/microsoft/wmi/pkg/wmiinstance"
 "github.com/microsoft/wmi/pkg/base/instance"
 "github.com/microsoft/wmi/pkg/errors"
 "reflect"
)

// MSFT_DSCMethodInvoked struct
type MSFT_DSCMethodInvoked struct { 
	*cim.WmiInstance

	// 
	ConfigurationData []uint8

	// 
	ConfigurationNumber uint8

	// 
	Flags uint32

	// 
	Guid string

	// 
	MethodNumber uint8

	// 
	ModuleName string

	// 
	ResourceName string

	// 
	UserSid string
}

	func NewMSFT_DSCMethodInvokedEx1(instance *cim.WmiInstance) (newInstance *MSFT_DSCMethodInvoked, err error) {tmp, err := instance, nil
		
	if err != nil { return }
	newInstance = &MSFT_DSCMethodInvoked {
	WmiInstance: tmp,
	}
	return
	}
	

	func NewMSFT_DSCMethodInvokedEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery ) (newInstance *MSFT_DSCMethodInvoked, err error) {tmp, err := instance.GetWmiInstance(hostName, wmiNamespace, userName, password, domainName, query)
		
	if err != nil { return }
	newInstance = &MSFT_DSCMethodInvoked {
	WmiInstance: tmp,
	}
	return
	}
	

// SetConfigurationData sets the value of ConfigurationData for the instance
func (instance *MSFT_DSCMethodInvoked) SetPropertyConfigurationData(value []uint8) (err error) { 
    return instance.SetProperty("ConfigurationData", (value))
}

// GetConfigurationData gets the value of ConfigurationData for the instance
func (instance *MSFT_DSCMethodInvoked) GetPropertyConfigurationData()(value []uint8, err error) { 
    retValue, err := instance.GetProperty("ConfigurationData")
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

// SetConfigurationNumber sets the value of ConfigurationNumber for the instance
func (instance *MSFT_DSCMethodInvoked) SetPropertyConfigurationNumber(value uint8) (err error) { 
    return instance.SetProperty("ConfigurationNumber", (value))
}

// GetConfigurationNumber gets the value of ConfigurationNumber for the instance
func (instance *MSFT_DSCMethodInvoked) GetPropertyConfigurationNumber()(value uint8, err error) { 
    retValue, err := instance.GetProperty("ConfigurationNumber")
    if err != nil {
        return
    }
    if retValue == nil {
        // Doesn't have any value. Return empty
        return
    }
    
    valuetmp, ok := retValue.(uint8); 
    if !ok {
        err = errors.Wrapf(errors.InvalidType, " uint8 is Invalid. Expected %s", reflect.TypeOf(retValue))
        return  
    }

    value = uint8(valuetmp)

    return
}

// SetFlags sets the value of Flags for the instance
func (instance *MSFT_DSCMethodInvoked) SetPropertyFlags(value uint32) (err error) { 
    return instance.SetProperty("Flags", (value))
}

// GetFlags gets the value of Flags for the instance
func (instance *MSFT_DSCMethodInvoked) GetPropertyFlags()(value uint32, err error) { 
    retValue, err := instance.GetProperty("Flags")
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

// SetGuid sets the value of Guid for the instance
func (instance *MSFT_DSCMethodInvoked) SetPropertyGuid(value string) (err error) { 
    return instance.SetProperty("Guid", (value))
}

// GetGuid gets the value of Guid for the instance
func (instance *MSFT_DSCMethodInvoked) GetPropertyGuid()(value string, err error) { 
    retValue, err := instance.GetProperty("Guid")
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

// SetMethodNumber sets the value of MethodNumber for the instance
func (instance *MSFT_DSCMethodInvoked) SetPropertyMethodNumber(value uint8) (err error) { 
    return instance.SetProperty("MethodNumber", (value))
}

// GetMethodNumber gets the value of MethodNumber for the instance
func (instance *MSFT_DSCMethodInvoked) GetPropertyMethodNumber()(value uint8, err error) { 
    retValue, err := instance.GetProperty("MethodNumber")
    if err != nil {
        return
    }
    if retValue == nil {
        // Doesn't have any value. Return empty
        return
    }
    
    valuetmp, ok := retValue.(uint8); 
    if !ok {
        err = errors.Wrapf(errors.InvalidType, " uint8 is Invalid. Expected %s", reflect.TypeOf(retValue))
        return  
    }

    value = uint8(valuetmp)

    return
}

// SetModuleName sets the value of ModuleName for the instance
func (instance *MSFT_DSCMethodInvoked) SetPropertyModuleName(value string) (err error) { 
    return instance.SetProperty("ModuleName", (value))
}

// GetModuleName gets the value of ModuleName for the instance
func (instance *MSFT_DSCMethodInvoked) GetPropertyModuleName()(value string, err error) { 
    retValue, err := instance.GetProperty("ModuleName")
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

// SetResourceName sets the value of ResourceName for the instance
func (instance *MSFT_DSCMethodInvoked) SetPropertyResourceName(value string) (err error) { 
    return instance.SetProperty("ResourceName", (value))
}

// GetResourceName gets the value of ResourceName for the instance
func (instance *MSFT_DSCMethodInvoked) GetPropertyResourceName()(value string, err error) { 
    retValue, err := instance.GetProperty("ResourceName")
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

// SetUserSid sets the value of UserSid for the instance
func (instance *MSFT_DSCMethodInvoked) SetPropertyUserSid(value string) (err error) { 
    return instance.SetProperty("UserSid", (value))
}

// GetUserSid gets the value of UserSid for the instance
func (instance *MSFT_DSCMethodInvoked) GetPropertyUserSid()(value string, err error) { 
    retValue, err := instance.GetProperty("UserSid")
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

