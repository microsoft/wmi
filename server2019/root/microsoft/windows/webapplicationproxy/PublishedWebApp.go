// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

// 
// Author:
//      Auto Generated on 6/6/2024 using wmigen
//      Source root.Microsoft.Windows.WebApplicationProxy
//////////////////////////////////////////////
package webapplicationproxy
import (
 "github.com/microsoft/wmi/pkg/base/query"
cim "github.com/microsoft/wmi/pkg/wmiinstance"
 "github.com/microsoft/wmi/pkg/base/instance"
 "github.com/microsoft/wmi/pkg/errors"
 "reflect"
)

// PublishedWebApp struct
type PublishedWebApp struct { 
	*cim.WmiInstance

	// 
	ADFSRelyingPartyID string

	// 
	ADFSRelyingPartyName string

	// 
	ADFSUserCertificateStore string

	// 
	BackendServerAuthenticationMode string

	// 
	BackendServerAuthenticationSPN string

	// 
	BackendServerCertificateValidation string

	// 
	BackendServerUrl string

	// 
	ClientCertificateAuthenticationBindingMode string

	// 
	ClientCertificatePreauthenticationThumbprint string

	// 
	DisableHttpOnlyCookieProtection bool

	// 
	DisableTranslateUrlInRequestHeaders bool

	// 
	DisableTranslateUrlInResponseHeaders bool

	// 
	EnableHTTPRedirect bool

	// 
	EnableSignOut bool

	// 
	ExternalCertificateThumbprint string

	// 
	ExternalPreauthentication string

	// 
	ExternalUrl string

	// 
	ID string

	// 
	InactiveTransactionsTimeoutSec uint32

	// 
	Name string

	// 
	PersistentAccessCookieExpirationTimeSec uint32

	// 
	UseOAuthAuthentication bool
}

	func NewPublishedWebAppEx1(instance *cim.WmiInstance) (newInstance *PublishedWebApp, err error) {tmp, err := instance, nil
		
	if err != nil { return }
	newInstance = &PublishedWebApp {
	WmiInstance: tmp,
	}
	return
	}
	

	func NewPublishedWebAppEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery ) (newInstance *PublishedWebApp, err error) {tmp, err := instance.GetWmiInstance(hostName, wmiNamespace, userName, password, domainName, query)
		
	if err != nil { return }
	newInstance = &PublishedWebApp {
	WmiInstance: tmp,
	}
	return
	}
	

// SetADFSRelyingPartyID sets the value of ADFSRelyingPartyID for the instance
func (instance *PublishedWebApp) SetPropertyADFSRelyingPartyID(value string) (err error) { 
    return instance.SetProperty("ADFSRelyingPartyID", (value))
}

// GetADFSRelyingPartyID gets the value of ADFSRelyingPartyID for the instance
func (instance *PublishedWebApp) GetPropertyADFSRelyingPartyID()(value string, err error) { 
    retValue, err := instance.GetProperty("ADFSRelyingPartyID")
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

// SetADFSRelyingPartyName sets the value of ADFSRelyingPartyName for the instance
func (instance *PublishedWebApp) SetPropertyADFSRelyingPartyName(value string) (err error) { 
    return instance.SetProperty("ADFSRelyingPartyName", (value))
}

// GetADFSRelyingPartyName gets the value of ADFSRelyingPartyName for the instance
func (instance *PublishedWebApp) GetPropertyADFSRelyingPartyName()(value string, err error) { 
    retValue, err := instance.GetProperty("ADFSRelyingPartyName")
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

// SetADFSUserCertificateStore sets the value of ADFSUserCertificateStore for the instance
func (instance *PublishedWebApp) SetPropertyADFSUserCertificateStore(value string) (err error) { 
    return instance.SetProperty("ADFSUserCertificateStore", (value))
}

// GetADFSUserCertificateStore gets the value of ADFSUserCertificateStore for the instance
func (instance *PublishedWebApp) GetPropertyADFSUserCertificateStore()(value string, err error) { 
    retValue, err := instance.GetProperty("ADFSUserCertificateStore")
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

// SetBackendServerAuthenticationMode sets the value of BackendServerAuthenticationMode for the instance
func (instance *PublishedWebApp) SetPropertyBackendServerAuthenticationMode(value string) (err error) { 
    return instance.SetProperty("BackendServerAuthenticationMode", (value))
}

// GetBackendServerAuthenticationMode gets the value of BackendServerAuthenticationMode for the instance
func (instance *PublishedWebApp) GetPropertyBackendServerAuthenticationMode()(value string, err error) { 
    retValue, err := instance.GetProperty("BackendServerAuthenticationMode")
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

// SetBackendServerAuthenticationSPN sets the value of BackendServerAuthenticationSPN for the instance
func (instance *PublishedWebApp) SetPropertyBackendServerAuthenticationSPN(value string) (err error) { 
    return instance.SetProperty("BackendServerAuthenticationSPN", (value))
}

// GetBackendServerAuthenticationSPN gets the value of BackendServerAuthenticationSPN for the instance
func (instance *PublishedWebApp) GetPropertyBackendServerAuthenticationSPN()(value string, err error) { 
    retValue, err := instance.GetProperty("BackendServerAuthenticationSPN")
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

// SetBackendServerCertificateValidation sets the value of BackendServerCertificateValidation for the instance
func (instance *PublishedWebApp) SetPropertyBackendServerCertificateValidation(value string) (err error) { 
    return instance.SetProperty("BackendServerCertificateValidation", (value))
}

// GetBackendServerCertificateValidation gets the value of BackendServerCertificateValidation for the instance
func (instance *PublishedWebApp) GetPropertyBackendServerCertificateValidation()(value string, err error) { 
    retValue, err := instance.GetProperty("BackendServerCertificateValidation")
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

// SetBackendServerUrl sets the value of BackendServerUrl for the instance
func (instance *PublishedWebApp) SetPropertyBackendServerUrl(value string) (err error) { 
    return instance.SetProperty("BackendServerUrl", (value))
}

// GetBackendServerUrl gets the value of BackendServerUrl for the instance
func (instance *PublishedWebApp) GetPropertyBackendServerUrl()(value string, err error) { 
    retValue, err := instance.GetProperty("BackendServerUrl")
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

// SetClientCertificateAuthenticationBindingMode sets the value of ClientCertificateAuthenticationBindingMode for the instance
func (instance *PublishedWebApp) SetPropertyClientCertificateAuthenticationBindingMode(value string) (err error) { 
    return instance.SetProperty("ClientCertificateAuthenticationBindingMode", (value))
}

// GetClientCertificateAuthenticationBindingMode gets the value of ClientCertificateAuthenticationBindingMode for the instance
func (instance *PublishedWebApp) GetPropertyClientCertificateAuthenticationBindingMode()(value string, err error) { 
    retValue, err := instance.GetProperty("ClientCertificateAuthenticationBindingMode")
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

// SetClientCertificatePreauthenticationThumbprint sets the value of ClientCertificatePreauthenticationThumbprint for the instance
func (instance *PublishedWebApp) SetPropertyClientCertificatePreauthenticationThumbprint(value string) (err error) { 
    return instance.SetProperty("ClientCertificatePreauthenticationThumbprint", (value))
}

// GetClientCertificatePreauthenticationThumbprint gets the value of ClientCertificatePreauthenticationThumbprint for the instance
func (instance *PublishedWebApp) GetPropertyClientCertificatePreauthenticationThumbprint()(value string, err error) { 
    retValue, err := instance.GetProperty("ClientCertificatePreauthenticationThumbprint")
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

// SetDisableHttpOnlyCookieProtection sets the value of DisableHttpOnlyCookieProtection for the instance
func (instance *PublishedWebApp) SetPropertyDisableHttpOnlyCookieProtection(value bool) (err error) { 
    return instance.SetProperty("DisableHttpOnlyCookieProtection", (value))
}

// GetDisableHttpOnlyCookieProtection gets the value of DisableHttpOnlyCookieProtection for the instance
func (instance *PublishedWebApp) GetPropertyDisableHttpOnlyCookieProtection()(value bool, err error) { 
    retValue, err := instance.GetProperty("DisableHttpOnlyCookieProtection")
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

// SetDisableTranslateUrlInRequestHeaders sets the value of DisableTranslateUrlInRequestHeaders for the instance
func (instance *PublishedWebApp) SetPropertyDisableTranslateUrlInRequestHeaders(value bool) (err error) { 
    return instance.SetProperty("DisableTranslateUrlInRequestHeaders", (value))
}

// GetDisableTranslateUrlInRequestHeaders gets the value of DisableTranslateUrlInRequestHeaders for the instance
func (instance *PublishedWebApp) GetPropertyDisableTranslateUrlInRequestHeaders()(value bool, err error) { 
    retValue, err := instance.GetProperty("DisableTranslateUrlInRequestHeaders")
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

// SetDisableTranslateUrlInResponseHeaders sets the value of DisableTranslateUrlInResponseHeaders for the instance
func (instance *PublishedWebApp) SetPropertyDisableTranslateUrlInResponseHeaders(value bool) (err error) { 
    return instance.SetProperty("DisableTranslateUrlInResponseHeaders", (value))
}

// GetDisableTranslateUrlInResponseHeaders gets the value of DisableTranslateUrlInResponseHeaders for the instance
func (instance *PublishedWebApp) GetPropertyDisableTranslateUrlInResponseHeaders()(value bool, err error) { 
    retValue, err := instance.GetProperty("DisableTranslateUrlInResponseHeaders")
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

// SetEnableHTTPRedirect sets the value of EnableHTTPRedirect for the instance
func (instance *PublishedWebApp) SetPropertyEnableHTTPRedirect(value bool) (err error) { 
    return instance.SetProperty("EnableHTTPRedirect", (value))
}

// GetEnableHTTPRedirect gets the value of EnableHTTPRedirect for the instance
func (instance *PublishedWebApp) GetPropertyEnableHTTPRedirect()(value bool, err error) { 
    retValue, err := instance.GetProperty("EnableHTTPRedirect")
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

// SetEnableSignOut sets the value of EnableSignOut for the instance
func (instance *PublishedWebApp) SetPropertyEnableSignOut(value bool) (err error) { 
    return instance.SetProperty("EnableSignOut", (value))
}

// GetEnableSignOut gets the value of EnableSignOut for the instance
func (instance *PublishedWebApp) GetPropertyEnableSignOut()(value bool, err error) { 
    retValue, err := instance.GetProperty("EnableSignOut")
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

// SetExternalCertificateThumbprint sets the value of ExternalCertificateThumbprint for the instance
func (instance *PublishedWebApp) SetPropertyExternalCertificateThumbprint(value string) (err error) { 
    return instance.SetProperty("ExternalCertificateThumbprint", (value))
}

// GetExternalCertificateThumbprint gets the value of ExternalCertificateThumbprint for the instance
func (instance *PublishedWebApp) GetPropertyExternalCertificateThumbprint()(value string, err error) { 
    retValue, err := instance.GetProperty("ExternalCertificateThumbprint")
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

// SetExternalPreauthentication sets the value of ExternalPreauthentication for the instance
func (instance *PublishedWebApp) SetPropertyExternalPreauthentication(value string) (err error) { 
    return instance.SetProperty("ExternalPreauthentication", (value))
}

// GetExternalPreauthentication gets the value of ExternalPreauthentication for the instance
func (instance *PublishedWebApp) GetPropertyExternalPreauthentication()(value string, err error) { 
    retValue, err := instance.GetProperty("ExternalPreauthentication")
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

// SetExternalUrl sets the value of ExternalUrl for the instance
func (instance *PublishedWebApp) SetPropertyExternalUrl(value string) (err error) { 
    return instance.SetProperty("ExternalUrl", (value))
}

// GetExternalUrl gets the value of ExternalUrl for the instance
func (instance *PublishedWebApp) GetPropertyExternalUrl()(value string, err error) { 
    retValue, err := instance.GetProperty("ExternalUrl")
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

// SetID sets the value of ID for the instance
func (instance *PublishedWebApp) SetPropertyID(value string) (err error) { 
    return instance.SetProperty("ID", (value))
}

// GetID gets the value of ID for the instance
func (instance *PublishedWebApp) GetPropertyID()(value string, err error) { 
    retValue, err := instance.GetProperty("ID")
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

// SetInactiveTransactionsTimeoutSec sets the value of InactiveTransactionsTimeoutSec for the instance
func (instance *PublishedWebApp) SetPropertyInactiveTransactionsTimeoutSec(value uint32) (err error) { 
    return instance.SetProperty("InactiveTransactionsTimeoutSec", (value))
}

// GetInactiveTransactionsTimeoutSec gets the value of InactiveTransactionsTimeoutSec for the instance
func (instance *PublishedWebApp) GetPropertyInactiveTransactionsTimeoutSec()(value uint32, err error) { 
    retValue, err := instance.GetProperty("InactiveTransactionsTimeoutSec")
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

// SetName sets the value of Name for the instance
func (instance *PublishedWebApp) SetPropertyName(value string) (err error) { 
    return instance.SetProperty("Name", (value))
}

// GetName gets the value of Name for the instance
func (instance *PublishedWebApp) GetPropertyName()(value string, err error) { 
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

// SetPersistentAccessCookieExpirationTimeSec sets the value of PersistentAccessCookieExpirationTimeSec for the instance
func (instance *PublishedWebApp) SetPropertyPersistentAccessCookieExpirationTimeSec(value uint32) (err error) { 
    return instance.SetProperty("PersistentAccessCookieExpirationTimeSec", (value))
}

// GetPersistentAccessCookieExpirationTimeSec gets the value of PersistentAccessCookieExpirationTimeSec for the instance
func (instance *PublishedWebApp) GetPropertyPersistentAccessCookieExpirationTimeSec()(value uint32, err error) { 
    retValue, err := instance.GetProperty("PersistentAccessCookieExpirationTimeSec")
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

// SetUseOAuthAuthentication sets the value of UseOAuthAuthentication for the instance
func (instance *PublishedWebApp) SetPropertyUseOAuthAuthentication(value bool) (err error) { 
    return instance.SetProperty("UseOAuthAuthentication", (value))
}

// GetUseOAuthAuthentication gets the value of UseOAuthAuthentication for the instance
func (instance *PublishedWebApp) GetPropertyUseOAuthAuthentication()(value bool, err error) { 
    retValue, err := instance.GetProperty("UseOAuthAuthentication")
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

