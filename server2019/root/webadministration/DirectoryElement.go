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

// DirectoryElement struct
type DirectoryElement struct { 
	*CollectionElement

	// 
	DirectoryName string
}

	func NewDirectoryElementEx1(instance *cim.WmiInstance) (newInstance *DirectoryElement, err error) {tmp, err := NewCollectionElementEx1(instance)
		
	if err != nil { return }
	newInstance = &DirectoryElement {
	CollectionElement: tmp,
	}
	return
	}
	

	func NewDirectoryElementEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery ) (newInstance *DirectoryElement, err error) {tmp, err := NewCollectionElementEx6(hostName, wmiNamespace, userName, password, domainName, query)
		
	if err != nil { return }
	newInstance = &DirectoryElement {
	CollectionElement: tmp,
	}
	return
	}
	

// SetDirectoryName sets the value of DirectoryName for the instance
func (instance *DirectoryElement) SetPropertyDirectoryName(value string) (err error) { 
    return instance.SetProperty("DirectoryName", (value))
}

// GetDirectoryName gets the value of DirectoryName for the instance
func (instance *DirectoryElement) GetPropertyDirectoryName()(value string, err error) { 
    retValue, err := instance.GetProperty("DirectoryName")
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

