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

// ScanHeadersElement struct
type ScanHeadersElement struct { 
	*CollectionElement

	// 
	RequestHeader string
}

	func NewScanHeadersElementEx1(instance *cim.WmiInstance) (newInstance *ScanHeadersElement, err error) {tmp, err := NewCollectionElementEx1(instance)
		
	if err != nil { return }
	newInstance = &ScanHeadersElement {
	CollectionElement: tmp,
	}
	return
	}
	

	func NewScanHeadersElementEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery ) (newInstance *ScanHeadersElement, err error) {tmp, err := NewCollectionElementEx6(hostName, wmiNamespace, userName, password, domainName, query)
		
	if err != nil { return }
	newInstance = &ScanHeadersElement {
	CollectionElement: tmp,
	}
	return
	}
	

// SetRequestHeader sets the value of RequestHeader for the instance
func (instance *ScanHeadersElement) SetPropertyRequestHeader(value string) (err error) { 
    return instance.SetProperty("RequestHeader", (value))
}

// GetRequestHeader gets the value of RequestHeader for the instance
func (instance *ScanHeadersElement) GetPropertyRequestHeader()(value string, err error) { 
    retValue, err := instance.GetProperty("RequestHeader")
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

