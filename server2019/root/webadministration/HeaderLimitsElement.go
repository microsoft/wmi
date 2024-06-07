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

// HeaderLimitsElement struct
type HeaderLimitsElement struct { 
	*CollectionElement

	// 
	Header string

	// 
	SizeLimit uint32
}

	func NewHeaderLimitsElementEx1(instance *cim.WmiInstance) (newInstance *HeaderLimitsElement, err error) {tmp, err := NewCollectionElementEx1(instance)
		
	if err != nil { return }
	newInstance = &HeaderLimitsElement {
	CollectionElement: tmp,
	}
	return
	}
	

	func NewHeaderLimitsElementEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery ) (newInstance *HeaderLimitsElement, err error) {tmp, err := NewCollectionElementEx6(hostName, wmiNamespace, userName, password, domainName, query)
		
	if err != nil { return }
	newInstance = &HeaderLimitsElement {
	CollectionElement: tmp,
	}
	return
	}
	

// SetHeader sets the value of Header for the instance
func (instance *HeaderLimitsElement) SetPropertyHeader(value string) (err error) { 
    return instance.SetProperty("Header", (value))
}

// GetHeader gets the value of Header for the instance
func (instance *HeaderLimitsElement) GetPropertyHeader()(value string, err error) { 
    retValue, err := instance.GetProperty("Header")
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

// SetSizeLimit sets the value of SizeLimit for the instance
func (instance *HeaderLimitsElement) SetPropertySizeLimit(value uint32) (err error) { 
    return instance.SetProperty("SizeLimit", (value))
}

// GetSizeLimit gets the value of SizeLimit for the instance
func (instance *HeaderLimitsElement) GetPropertySizeLimit()(value uint32, err error) { 
    retValue, err := instance.GetProperty("SizeLimit")
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

