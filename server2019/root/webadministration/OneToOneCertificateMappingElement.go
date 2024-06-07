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

// OneToOneCertificateMappingElement struct
type OneToOneCertificateMappingElement struct { 
	*CollectionElement

	// 
	Certificate string

	// 
	Enabled bool

	// 
	Password string

	// 
	UserName string
}

	func NewOneToOneCertificateMappingElementEx1(instance *cim.WmiInstance) (newInstance *OneToOneCertificateMappingElement, err error) {tmp, err := NewCollectionElementEx1(instance)
		
	if err != nil { return }
	newInstance = &OneToOneCertificateMappingElement {
	CollectionElement: tmp,
	}
	return
	}
	

	func NewOneToOneCertificateMappingElementEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery ) (newInstance *OneToOneCertificateMappingElement, err error) {tmp, err := NewCollectionElementEx6(hostName, wmiNamespace, userName, password, domainName, query)
		
	if err != nil { return }
	newInstance = &OneToOneCertificateMappingElement {
	CollectionElement: tmp,
	}
	return
	}
	

// SetCertificate sets the value of Certificate for the instance
func (instance *OneToOneCertificateMappingElement) SetPropertyCertificate(value string) (err error) { 
    return instance.SetProperty("Certificate", (value))
}

// GetCertificate gets the value of Certificate for the instance
func (instance *OneToOneCertificateMappingElement) GetPropertyCertificate()(value string, err error) { 
    retValue, err := instance.GetProperty("Certificate")
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

// SetEnabled sets the value of Enabled for the instance
func (instance *OneToOneCertificateMappingElement) SetPropertyEnabled(value bool) (err error) { 
    return instance.SetProperty("Enabled", (value))
}

// GetEnabled gets the value of Enabled for the instance
func (instance *OneToOneCertificateMappingElement) GetPropertyEnabled()(value bool, err error) { 
    retValue, err := instance.GetProperty("Enabled")
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

// SetPassword sets the value of Password for the instance
func (instance *OneToOneCertificateMappingElement) SetPropertyPassword(value string) (err error) { 
    return instance.SetProperty("Password", (value))
}

// GetPassword gets the value of Password for the instance
func (instance *OneToOneCertificateMappingElement) GetPropertyPassword()(value string, err error) { 
    retValue, err := instance.GetProperty("Password")
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

// SetUserName sets the value of UserName for the instance
func (instance *OneToOneCertificateMappingElement) SetPropertyUserName(value string) (err error) { 
    return instance.SetProperty("UserName", (value))
}

// GetUserName gets the value of UserName for the instance
func (instance *OneToOneCertificateMappingElement) GetPropertyUserName()(value string, err error) { 
    retValue, err := instance.GetProperty("UserName")
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

