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

// FormsAuthenticationConfiguration struct
type FormsAuthenticationConfiguration struct { 
	*EmbeddedObject

	// 
	Cookieless int32

	// 
	Credentials FormsAuthenticationCredentials

	// 
	DefaultUrl string

	// 
	Domain string

	// 
	EnableCrossAppRedirects bool

	// 
	LoginUrl string

	// 
	Name string

	// 
	Path string

	// 
	Protection int32

	// 
	RequireSSL bool

	// 
	SlidingExpiration bool

	// 
	TicketCompatibilityMode string

	// 
	Timeout string
}

	func NewFormsAuthenticationConfigurationEx1(instance *cim.WmiInstance) (newInstance *FormsAuthenticationConfiguration, err error) {tmp, err := NewEmbeddedObjectEx1(instance)
		
	if err != nil { return }
	newInstance = &FormsAuthenticationConfiguration {
	EmbeddedObject: tmp,
	}
	return
	}
	

	func NewFormsAuthenticationConfigurationEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery ) (newInstance *FormsAuthenticationConfiguration, err error) {tmp, err := NewEmbeddedObjectEx6(hostName, wmiNamespace, userName, password, domainName, query)
		
	if err != nil { return }
	newInstance = &FormsAuthenticationConfiguration {
	EmbeddedObject: tmp,
	}
	return
	}
	

// SetCookieless sets the value of Cookieless for the instance
func (instance *FormsAuthenticationConfiguration) SetPropertyCookieless(value int32) (err error) { 
    return instance.SetProperty("Cookieless", (value))
}

// GetCookieless gets the value of Cookieless for the instance
func (instance *FormsAuthenticationConfiguration) GetPropertyCookieless()(value int32, err error) { 
    retValue, err := instance.GetProperty("Cookieless")
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

// SetCredentials sets the value of Credentials for the instance
func (instance *FormsAuthenticationConfiguration) SetPropertyCredentials(value FormsAuthenticationCredentials) (err error) { 
    return instance.SetProperty("Credentials", (value))
}

// GetCredentials gets the value of Credentials for the instance
func (instance *FormsAuthenticationConfiguration) GetPropertyCredentials()(value FormsAuthenticationCredentials, err error) { 
    retValue, err := instance.GetProperty("Credentials")
    if err != nil {
        return
    }
    if retValue == nil {
        // Doesn't have any value. Return empty
        return
    }
    
    valuetmp, ok := retValue.(FormsAuthenticationCredentials); 
    if !ok {
        err = errors.Wrapf(errors.InvalidType, " FormsAuthenticationCredentials is Invalid. Expected %s", reflect.TypeOf(retValue))
        return  
    }

    value = FormsAuthenticationCredentials(valuetmp)

    return
}

// SetDefaultUrl sets the value of DefaultUrl for the instance
func (instance *FormsAuthenticationConfiguration) SetPropertyDefaultUrl(value string) (err error) { 
    return instance.SetProperty("DefaultUrl", (value))
}

// GetDefaultUrl gets the value of DefaultUrl for the instance
func (instance *FormsAuthenticationConfiguration) GetPropertyDefaultUrl()(value string, err error) { 
    retValue, err := instance.GetProperty("DefaultUrl")
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

// SetDomain sets the value of Domain for the instance
func (instance *FormsAuthenticationConfiguration) SetPropertyDomain(value string) (err error) { 
    return instance.SetProperty("Domain", (value))
}

// GetDomain gets the value of Domain for the instance
func (instance *FormsAuthenticationConfiguration) GetPropertyDomain()(value string, err error) { 
    retValue, err := instance.GetProperty("Domain")
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

// SetEnableCrossAppRedirects sets the value of EnableCrossAppRedirects for the instance
func (instance *FormsAuthenticationConfiguration) SetPropertyEnableCrossAppRedirects(value bool) (err error) { 
    return instance.SetProperty("EnableCrossAppRedirects", (value))
}

// GetEnableCrossAppRedirects gets the value of EnableCrossAppRedirects for the instance
func (instance *FormsAuthenticationConfiguration) GetPropertyEnableCrossAppRedirects()(value bool, err error) { 
    retValue, err := instance.GetProperty("EnableCrossAppRedirects")
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

// SetLoginUrl sets the value of LoginUrl for the instance
func (instance *FormsAuthenticationConfiguration) SetPropertyLoginUrl(value string) (err error) { 
    return instance.SetProperty("LoginUrl", (value))
}

// GetLoginUrl gets the value of LoginUrl for the instance
func (instance *FormsAuthenticationConfiguration) GetPropertyLoginUrl()(value string, err error) { 
    retValue, err := instance.GetProperty("LoginUrl")
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

// SetName sets the value of Name for the instance
func (instance *FormsAuthenticationConfiguration) SetPropertyName(value string) (err error) { 
    return instance.SetProperty("Name", (value))
}

// GetName gets the value of Name for the instance
func (instance *FormsAuthenticationConfiguration) GetPropertyName()(value string, err error) { 
    retValue, err := instance.GetProperty("Name")
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
func (instance *FormsAuthenticationConfiguration) SetPropertyPath(value string) (err error) { 
    return instance.SetProperty("Path", (value))
}

// GetPath gets the value of Path for the instance
func (instance *FormsAuthenticationConfiguration) GetPropertyPath()(value string, err error) { 
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

// SetProtection sets the value of Protection for the instance
func (instance *FormsAuthenticationConfiguration) SetPropertyProtection(value int32) (err error) { 
    return instance.SetProperty("Protection", (value))
}

// GetProtection gets the value of Protection for the instance
func (instance *FormsAuthenticationConfiguration) GetPropertyProtection()(value int32, err error) { 
    retValue, err := instance.GetProperty("Protection")
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

// SetRequireSSL sets the value of RequireSSL for the instance
func (instance *FormsAuthenticationConfiguration) SetPropertyRequireSSL(value bool) (err error) { 
    return instance.SetProperty("RequireSSL", (value))
}

// GetRequireSSL gets the value of RequireSSL for the instance
func (instance *FormsAuthenticationConfiguration) GetPropertyRequireSSL()(value bool, err error) { 
    retValue, err := instance.GetProperty("RequireSSL")
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

// SetSlidingExpiration sets the value of SlidingExpiration for the instance
func (instance *FormsAuthenticationConfiguration) SetPropertySlidingExpiration(value bool) (err error) { 
    return instance.SetProperty("SlidingExpiration", (value))
}

// GetSlidingExpiration gets the value of SlidingExpiration for the instance
func (instance *FormsAuthenticationConfiguration) GetPropertySlidingExpiration()(value bool, err error) { 
    retValue, err := instance.GetProperty("SlidingExpiration")
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

// SetTicketCompatibilityMode sets the value of TicketCompatibilityMode for the instance
func (instance *FormsAuthenticationConfiguration) SetPropertyTicketCompatibilityMode(value string) (err error) { 
    return instance.SetProperty("TicketCompatibilityMode", (value))
}

// GetTicketCompatibilityMode gets the value of TicketCompatibilityMode for the instance
func (instance *FormsAuthenticationConfiguration) GetPropertyTicketCompatibilityMode()(value string, err error) { 
    retValue, err := instance.GetProperty("TicketCompatibilityMode")
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

// SetTimeout sets the value of Timeout for the instance
func (instance *FormsAuthenticationConfiguration) SetPropertyTimeout(value string) (err error) { 
    return instance.SetProperty("Timeout", (value))
}

// GetTimeout gets the value of Timeout for the instance
func (instance *FormsAuthenticationConfiguration) GetPropertyTimeout()(value string, err error) { 
    retValue, err := instance.GetProperty("Timeout")
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

