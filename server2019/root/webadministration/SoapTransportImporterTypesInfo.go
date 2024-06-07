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

// SoapTransportImporterTypesInfo struct
type SoapTransportImporterTypesInfo struct { 
	*EmbeddedObject

	// 
	SoapTransportImporterTypes []SoapTransportImporterTypeElement
}

	func NewSoapTransportImporterTypesInfoEx1(instance *cim.WmiInstance) (newInstance *SoapTransportImporterTypesInfo, err error) {tmp, err := NewEmbeddedObjectEx1(instance)
		
	if err != nil { return }
	newInstance = &SoapTransportImporterTypesInfo {
	EmbeddedObject: tmp,
	}
	return
	}
	

	func NewSoapTransportImporterTypesInfoEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery ) (newInstance *SoapTransportImporterTypesInfo, err error) {tmp, err := NewEmbeddedObjectEx6(hostName, wmiNamespace, userName, password, domainName, query)
		
	if err != nil { return }
	newInstance = &SoapTransportImporterTypesInfo {
	EmbeddedObject: tmp,
	}
	return
	}
	

// SetSoapTransportImporterTypes sets the value of SoapTransportImporterTypes for the instance
func (instance *SoapTransportImporterTypesInfo) SetPropertySoapTransportImporterTypes(value []SoapTransportImporterTypeElement) (err error) { 
    return instance.SetProperty("SoapTransportImporterTypes", (value))
}

// GetSoapTransportImporterTypes gets the value of SoapTransportImporterTypes for the instance
func (instance *SoapTransportImporterTypesInfo) GetPropertySoapTransportImporterTypes()(value []SoapTransportImporterTypeElement, err error) { 
    retValue, err := instance.GetProperty("SoapTransportImporterTypes")
    if err != nil {
        return
    }
    if retValue == nil {
        // Doesn't have any value. Return empty
        return
    }
    
    for _, interfaceValue := range retValue.([]interface{}) {
        valuetmp, ok := interfaceValue.(SoapTransportImporterTypeElement); 
        if !ok {
            err = errors.Wrapf(errors.InvalidType, " SoapTransportImporterTypeElement is Invalid. Expected %s", reflect.TypeOf(interfaceValue))
            return  
        }
        value = append(value, SoapTransportImporterTypeElement(valuetmp))
    }

    return
}

