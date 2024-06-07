// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

// 
// Author:
//      Auto Generated on 6/6/2024 using wmigen
//      Source root.WebAdministration
//////////////////////////////////////////////
package webadministration
import (
 "github.com/microsoft/wmi/pkg/base/query"
cim "github.com/microsoft/wmi/pkg/wmiinstance"
 "github.com/microsoft/wmi/pkg/base/instance"
 "github.com/microsoft/wmi/pkg/errors"
 "reflect"
)

// ConfigurationSection struct
type ConfigurationSection struct { 
	*cim.WmiInstance

	// 
	Location string

	// 
	Path string

	// 
	SectionInformation SectionInformation
}

	func NewConfigurationSectionEx1(instance *cim.WmiInstance) (newInstance *ConfigurationSection, err error) {tmp, err := instance, nil
		
	if err != nil { return }
	newInstance = &ConfigurationSection {
	WmiInstance: tmp,
	}
	return
	}
	

	func NewConfigurationSectionEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery ) (newInstance *ConfigurationSection, err error) {tmp, err := instance.GetWmiInstance(hostName, wmiNamespace, userName, password, domainName, query)
		
	if err != nil { return }
	newInstance = &ConfigurationSection {
	WmiInstance: tmp,
	}
	return
	}
	

// SetLocation sets the value of Location for the instance
func (instance *ConfigurationSection) SetPropertyLocation(value string) (err error) { 
    return instance.SetProperty("Location", (value))
}

// GetLocation gets the value of Location for the instance
func (instance *ConfigurationSection) GetPropertyLocation()(value string, err error) { 
    retValue, err := instance.GetProperty("Location")
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

// SetPath sets the value of Path for the instance
func (instance *ConfigurationSection) SetPropertyPath(value string) (err error) { 
    return instance.SetProperty("Path", (value))
}

// GetPath gets the value of Path for the instance
func (instance *ConfigurationSection) GetPropertyPath()(value string, err error) { 
    retValue, err := instance.GetProperty("Path")
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

// SetSectionInformation sets the value of SectionInformation for the instance
func (instance *ConfigurationSection) SetPropertySectionInformation(value SectionInformation) (err error) { 
    return instance.SetProperty("SectionInformation", (value))
}

// GetSectionInformation gets the value of SectionInformation for the instance
func (instance *ConfigurationSection) GetPropertySectionInformation()(value SectionInformation, err error) { 
    retValue, err := instance.GetProperty("SectionInformation")
    if err != nil {
        return
    }
    if retValue == nil {
        // Doesn't have any value. Return empty
        return
    }
    
    valuetmp, ok := retValue.(SectionInformation); 
    if !ok {
        err = errors.Wrapf(errors.InvalidType, " SectionInformation is Invalid. Expected %s", reflect.TypeOf(retValue))
        return  
    }

    value = SectionInformation(valuetmp)

    return
}

// 

// <param name="PropertyName" type="string "></param>
func (instance *ConfigurationSection) RevertToParent( /* OPTIONAL IN */ PropertyName string) (err error) {_, err = instance.InvokeMethodWithReturn("RevertToParent" , PropertyName);
	if err != nil { return }
	return
	
}


// 

// <param name="AllowDefinition" type="string "></param>
func (instance *ConfigurationSection) GetAllowDefinition( /* OUT */ AllowDefinition string) (err error) {_, err = instance.InvokeMethod("GetAllowDefinition" )
	if err != nil { return }
	return
	
}


// 

// <param name="AllowDefinition" type="string "></param>
func (instance *ConfigurationSection) SetAllowDefinition( /* IN */ AllowDefinition string) (err error) {_, err = instance.InvokeMethodWithReturn("SetAllowDefinition" , AllowDefinition);
	if err != nil { return }
	return
	
}


// 

// <param name="AllowLocation" type="string "></param>
func (instance *ConfigurationSection) GetAllowLocation( /* OUT */ AllowLocation string) (err error) {_, err = instance.InvokeMethod("GetAllowLocation" )
	if err != nil { return }
	return
	
}


// 

// <param name="AllowLocation" type="string "></param>
func (instance *ConfigurationSection) SetAllowLocation( /* IN */ AllowLocation string) (err error) {_, err = instance.InvokeMethodWithReturn("SetAllowLocation" , AllowLocation);
	if err != nil { return }
	return
	
}


