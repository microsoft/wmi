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

// MimeMapElement struct
type MimeMapElement struct { 
	*CollectionElement

	// 
	FileExtension string

	// 
	MimeType string
}

	func NewMimeMapElementEx1(instance *cim.WmiInstance) (newInstance *MimeMapElement, err error) {tmp, err := NewCollectionElementEx1(instance)
		
	if err != nil { return }
	newInstance = &MimeMapElement {
	CollectionElement: tmp,
	}
	return
	}
	

	func NewMimeMapElementEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery ) (newInstance *MimeMapElement, err error) {tmp, err := NewCollectionElementEx6(hostName, wmiNamespace, userName, password, domainName, query)
		
	if err != nil { return }
	newInstance = &MimeMapElement {
	CollectionElement: tmp,
	}
	return
	}
	

// SetFileExtension sets the value of FileExtension for the instance
func (instance *MimeMapElement) SetPropertyFileExtension(value string) (err error) { 
    return instance.SetProperty("FileExtension", (value))
}

// GetFileExtension gets the value of FileExtension for the instance
func (instance *MimeMapElement) GetPropertyFileExtension()(value string, err error) { 
    retValue, err := instance.GetProperty("FileExtension")
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

// SetMimeType sets the value of MimeType for the instance
func (instance *MimeMapElement) SetPropertyMimeType(value string) (err error) { 
    return instance.SetProperty("MimeType", (value))
}

// GetMimeType gets the value of MimeType for the instance
func (instance *MimeMapElement) GetPropertyMimeType()(value string, err error) { 
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

