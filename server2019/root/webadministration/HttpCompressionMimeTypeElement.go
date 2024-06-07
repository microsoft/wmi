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

// HttpCompressionMimeTypeElement struct
type HttpCompressionMimeTypeElement struct { 
	*CollectionElement

	// 
	Enabled bool

	// 
	MimeType string
}

	func NewHttpCompressionMimeTypeElementEx1(instance *cim.WmiInstance) (newInstance *HttpCompressionMimeTypeElement, err error) {tmp, err := NewCollectionElementEx1(instance)
		
	if err != nil { return }
	newInstance = &HttpCompressionMimeTypeElement {
	CollectionElement: tmp,
	}
	return
	}
	

	func NewHttpCompressionMimeTypeElementEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery ) (newInstance *HttpCompressionMimeTypeElement, err error) {tmp, err := NewCollectionElementEx6(hostName, wmiNamespace, userName, password, domainName, query)
		
	if err != nil { return }
	newInstance = &HttpCompressionMimeTypeElement {
	CollectionElement: tmp,
	}
	return
	}
	

// SetEnabled sets the value of Enabled for the instance
func (instance *HttpCompressionMimeTypeElement) SetPropertyEnabled(value bool) (err error) { 
    return instance.SetProperty("Enabled", (value))
}

// GetEnabled gets the value of Enabled for the instance
func (instance *HttpCompressionMimeTypeElement) GetPropertyEnabled()(value bool, err error) { 
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

// SetMimeType sets the value of MimeType for the instance
func (instance *HttpCompressionMimeTypeElement) SetPropertyMimeType(value string) (err error) { 
    return instance.SetProperty("MimeType", (value))
}

// GetMimeType gets the value of MimeType for the instance
func (instance *HttpCompressionMimeTypeElement) GetPropertyMimeType()(value string, err error) { 
    retValue, err := instance.GetProperty("MimeType")
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

