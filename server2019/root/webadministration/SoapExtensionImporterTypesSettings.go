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

// SoapExtensionImporterTypesSettings struct
type SoapExtensionImporterTypesSettings struct { 
	*EmbeddedObject

	// 
	SoapExtensionImporterTypes []TypeElement
}

	func NewSoapExtensionImporterTypesSettingsEx1(instance *cim.WmiInstance) (newInstance *SoapExtensionImporterTypesSettings, err error) {tmp, err := NewEmbeddedObjectEx1(instance)
		
	if err != nil { return }
	newInstance = &SoapExtensionImporterTypesSettings {
	EmbeddedObject: tmp,
	}
	return
	}
	

	func NewSoapExtensionImporterTypesSettingsEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery ) (newInstance *SoapExtensionImporterTypesSettings, err error) {tmp, err := NewEmbeddedObjectEx6(hostName, wmiNamespace, userName, password, domainName, query)
		
	if err != nil { return }
	newInstance = &SoapExtensionImporterTypesSettings {
	EmbeddedObject: tmp,
	}
	return
	}
	

// SetSoapExtensionImporterTypes sets the value of SoapExtensionImporterTypes for the instance
func (instance *SoapExtensionImporterTypesSettings) SetPropertySoapExtensionImporterTypes(value []TypeElement) (err error) { 
    return instance.SetProperty("SoapExtensionImporterTypes", (value))
}

// GetSoapExtensionImporterTypes gets the value of SoapExtensionImporterTypes for the instance
func (instance *SoapExtensionImporterTypesSettings) GetPropertySoapExtensionImporterTypes()(value []TypeElement, err error) { 
    retValue, err := instance.GetProperty("SoapExtensionImporterTypes")
    if err != nil {
        return
    }
    if retValue == nil {
        // Doesn't have any value. Return empty
        return
    }
    
    for _, interfaceValue := range retValue.([]interface{}) {
        valuetmp, ok := interfaceValue.(TypeElement); 
        if !ok {
            err = errors.Wrapf(errors.InvalidType, " TypeElement is Invalid. Expected %s", reflect.TypeOf(interfaceValue))
            return  
        }
        value = append(value, TypeElement(valuetmp))
    }

    return
}

