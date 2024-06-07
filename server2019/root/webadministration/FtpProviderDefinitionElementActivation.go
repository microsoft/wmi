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

// FtpProviderDefinitionElementActivation struct
type FtpProviderDefinitionElementActivation struct { 
	*EmbeddedObject

	// 
	Activation []FtpProviderDataSettings
}

	func NewFtpProviderDefinitionElementActivationEx1(instance *cim.WmiInstance) (newInstance *FtpProviderDefinitionElementActivation, err error) {tmp, err := NewEmbeddedObjectEx1(instance)
		
	if err != nil { return }
	newInstance = &FtpProviderDefinitionElementActivation {
	EmbeddedObject: tmp,
	}
	return
	}
	

	func NewFtpProviderDefinitionElementActivationEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery ) (newInstance *FtpProviderDefinitionElementActivation, err error) {tmp, err := NewEmbeddedObjectEx6(hostName, wmiNamespace, userName, password, domainName, query)
		
	if err != nil { return }
	newInstance = &FtpProviderDefinitionElementActivation {
	EmbeddedObject: tmp,
	}
	return
	}
	

// SetActivation sets the value of Activation for the instance
func (instance *FtpProviderDefinitionElementActivation) SetPropertyActivation(value []FtpProviderDataSettings) (err error) { 
    return instance.SetProperty("Activation", (value))
}

// GetActivation gets the value of Activation for the instance
func (instance *FtpProviderDefinitionElementActivation) GetPropertyActivation()(value []FtpProviderDataSettings, err error) { 
    retValue, err := instance.GetProperty("Activation")
    if err != nil {
        return
    }
    if retValue == nil {
        // Doesn't have any value. Return empty
        return
    }
    
    for _, interfaceValue := range retValue.([]interface{}) {
        valuetmp, ok := interfaceValue.(FtpProviderDataSettings); 
        if !ok {
            err = errors.Wrapf(errors.InvalidType, " FtpProviderDataSettings is Invalid. Expected %s", reflect.TypeOf(interfaceValue))
            return  
        }
        value = append(value, FtpProviderDataSettings(valuetmp))
    }

    return
}

