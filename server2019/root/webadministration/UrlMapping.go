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

// UrlMapping struct
type UrlMapping struct { 
	*CollectionElement

	// 
	MappedUrl string

	// 
	Url string
}

	func NewUrlMappingEx1(instance *cim.WmiInstance) (newInstance *UrlMapping, err error) {tmp, err := NewCollectionElementEx1(instance)
		
	if err != nil { return }
	newInstance = &UrlMapping {
	CollectionElement: tmp,
	}
	return
	}
	

	func NewUrlMappingEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery ) (newInstance *UrlMapping, err error) {tmp, err := NewCollectionElementEx6(hostName, wmiNamespace, userName, password, domainName, query)
		
	if err != nil { return }
	newInstance = &UrlMapping {
	CollectionElement: tmp,
	}
	return
	}
	

// SetMappedUrl sets the value of MappedUrl for the instance
func (instance *UrlMapping) SetPropertyMappedUrl(value string) (err error) { 
    return instance.SetProperty("MappedUrl", (value))
}

// GetMappedUrl gets the value of MappedUrl for the instance
func (instance *UrlMapping) GetPropertyMappedUrl()(value string, err error) { 
    retValue, err := instance.GetProperty("MappedUrl")
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

// SetUrl sets the value of Url for the instance
func (instance *UrlMapping) SetPropertyUrl(value string) (err error) { 
    return instance.SetProperty("Url", (value))
}

// GetUrl gets the value of Url for the instance
func (instance *UrlMapping) GetPropertyUrl()(value string, err error) { 
    retValue, err := instance.GetProperty("Url")
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

