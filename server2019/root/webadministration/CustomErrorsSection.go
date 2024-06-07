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
 "github.com/microsoft/wmi/pkg/errors"
 "reflect"
)

// CustomErrorsSection struct
type CustomErrorsSection struct { 
	*ConfigurationSectionWithCollection

	// 
	AllowNestedErrors bool

	// 
	CustomErrors []CustomError

	// 
	DefaultRedirect string

	// 
	Mode int32

	// 
	RedirectMode string
}

	func NewCustomErrorsSectionEx1(instance *cim.WmiInstance) (newInstance *CustomErrorsSection, err error) {tmp, err := NewConfigurationSectionWithCollectionEx1(instance)
		
	if err != nil { return }
	newInstance = &CustomErrorsSection {
	ConfigurationSectionWithCollection: tmp,
	}
	return
	}
	

	func NewCustomErrorsSectionEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery ) (newInstance *CustomErrorsSection, err error) {tmp, err := NewConfigurationSectionWithCollectionEx6(hostName, wmiNamespace, userName, password, domainName, query)
		
	if err != nil { return }
	newInstance = &CustomErrorsSection {
	ConfigurationSectionWithCollection: tmp,
	}
	return
	}
	

// SetAllowNestedErrors sets the value of AllowNestedErrors for the instance
func (instance *CustomErrorsSection) SetPropertyAllowNestedErrors(value bool) (err error) { 
    return instance.SetProperty("AllowNestedErrors", (value))
}

// GetAllowNestedErrors gets the value of AllowNestedErrors for the instance
func (instance *CustomErrorsSection) GetPropertyAllowNestedErrors()(value bool, err error) { 
    retValue, err := instance.GetProperty("AllowNestedErrors")
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

// SetCustomErrors sets the value of CustomErrors for the instance
func (instance *CustomErrorsSection) SetPropertyCustomErrors(value []CustomError) (err error) { 
    return instance.SetProperty("CustomErrors", (value))
}

// GetCustomErrors gets the value of CustomErrors for the instance
func (instance *CustomErrorsSection) GetPropertyCustomErrors()(value []CustomError, err error) { 
    retValue, err := instance.GetProperty("CustomErrors")
    if err != nil {
        return
    }
    if retValue == nil {
        // Doesn't have any value. Return empty
        return
    }
    
    for _, interfaceValue := range retValue.([]interface{}) {
        valuetmp, ok := interfaceValue.(CustomError); 
        if !ok {
            err = errors.Wrapf(errors.InvalidType, " CustomError is Invalid. Expected %s", reflect.TypeOf(interfaceValue))
            return  
        }
        value = append(value, CustomError(valuetmp))
    }

    return
}

// SetDefaultRedirect sets the value of DefaultRedirect for the instance
func (instance *CustomErrorsSection) SetPropertyDefaultRedirect(value string) (err error) { 
    return instance.SetProperty("DefaultRedirect", (value))
}

// GetDefaultRedirect gets the value of DefaultRedirect for the instance
func (instance *CustomErrorsSection) GetPropertyDefaultRedirect()(value string, err error) { 
    retValue, err := instance.GetProperty("DefaultRedirect")
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

// SetMode sets the value of Mode for the instance
func (instance *CustomErrorsSection) SetPropertyMode(value int32) (err error) { 
    return instance.SetProperty("Mode", (value))
}

// GetMode gets the value of Mode for the instance
func (instance *CustomErrorsSection) GetPropertyMode()(value int32, err error) { 
    retValue, err := instance.GetProperty("Mode")
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

// SetRedirectMode sets the value of RedirectMode for the instance
func (instance *CustomErrorsSection) SetPropertyRedirectMode(value string) (err error) { 
    return instance.SetProperty("RedirectMode", (value))
}

// GetRedirectMode gets the value of RedirectMode for the instance
func (instance *CustomErrorsSection) GetPropertyRedirectMode()(value string, err error) { 
    retValue, err := instance.GetProperty("RedirectMode")
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

