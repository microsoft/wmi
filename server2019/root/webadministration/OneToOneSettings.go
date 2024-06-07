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

// OneToOneSettings struct
type OneToOneSettings struct { 
	*EmbeddedObject

	// 
	OneToOneMappings []OneToOneCertificateMappingElement
}

	func NewOneToOneSettingsEx1(instance *cim.WmiInstance) (newInstance *OneToOneSettings, err error) {tmp, err := NewEmbeddedObjectEx1(instance)
		
	if err != nil { return }
	newInstance = &OneToOneSettings {
	EmbeddedObject: tmp,
	}
	return
	}
	

	func NewOneToOneSettingsEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery ) (newInstance *OneToOneSettings, err error) {tmp, err := NewEmbeddedObjectEx6(hostName, wmiNamespace, userName, password, domainName, query)
		
	if err != nil { return }
	newInstance = &OneToOneSettings {
	EmbeddedObject: tmp,
	}
	return
	}
	

// SetOneToOneMappings sets the value of OneToOneMappings for the instance
func (instance *OneToOneSettings) SetPropertyOneToOneMappings(value []OneToOneCertificateMappingElement) (err error) { 
    return instance.SetProperty("OneToOneMappings", (value))
}

// GetOneToOneMappings gets the value of OneToOneMappings for the instance
func (instance *OneToOneSettings) GetPropertyOneToOneMappings()(value []OneToOneCertificateMappingElement, err error) { 
    retValue, err := instance.GetProperty("OneToOneMappings")
    if err != nil {
        return
    }
    if retValue == nil {
        // Doesn't have any value. Return empty
        return
    }
    
    for _, interfaceValue := range retValue.([]interface{}) {
        valuetmp, ok := interfaceValue.(OneToOneCertificateMappingElement); 
        if !ok {
            err = errors.Wrapf(errors.InvalidType, " OneToOneCertificateMappingElement is Invalid. Expected %s", reflect.TypeOf(interfaceValue))
            return  
        }
        value = append(value, OneToOneCertificateMappingElement(valuetmp))
    }

    return
}

