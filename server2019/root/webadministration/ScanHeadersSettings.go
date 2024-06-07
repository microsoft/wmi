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

// ScanHeadersSettings struct
type ScanHeadersSettings struct { 
	*EmbeddedObject

	// 
	ScanHeaders []ScanHeadersElement
}

	func NewScanHeadersSettingsEx1(instance *cim.WmiInstance) (newInstance *ScanHeadersSettings, err error) {tmp, err := NewEmbeddedObjectEx1(instance)
		
	if err != nil { return }
	newInstance = &ScanHeadersSettings {
	EmbeddedObject: tmp,
	}
	return
	}
	

	func NewScanHeadersSettingsEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery ) (newInstance *ScanHeadersSettings, err error) {tmp, err := NewEmbeddedObjectEx6(hostName, wmiNamespace, userName, password, domainName, query)
		
	if err != nil { return }
	newInstance = &ScanHeadersSettings {
	EmbeddedObject: tmp,
	}
	return
	}
	

// SetScanHeaders sets the value of ScanHeaders for the instance
func (instance *ScanHeadersSettings) SetPropertyScanHeaders(value []ScanHeadersElement) (err error) { 
    return instance.SetProperty("ScanHeaders", (value))
}

// GetScanHeaders gets the value of ScanHeaders for the instance
func (instance *ScanHeadersSettings) GetPropertyScanHeaders()(value []ScanHeadersElement, err error) { 
    retValue, err := instance.GetProperty("ScanHeaders")
    if err != nil {
        return
    }
    if retValue == nil {
        // Doesn't have any value. Return empty
        return
    }
    
    for _, interfaceValue := range retValue.([]interface{}) {
        valuetmp, ok := interfaceValue.(ScanHeadersElement); 
        if !ok {
            err = errors.Wrapf(errors.InvalidType, " ScanHeadersElement is Invalid. Expected %s", reflect.TypeOf(interfaceValue))
            return  
        }
        value = append(value, ScanHeadersElement(valuetmp))
    }

    return
}

