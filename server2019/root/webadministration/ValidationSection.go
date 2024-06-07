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

// ValidationSection struct
type ValidationSection struct { 
	*ConfigurationSection

	// 
	ValidateIntegratedModeConfiguration bool
}

	func NewValidationSectionEx1(instance *cim.WmiInstance) (newInstance *ValidationSection, err error) {tmp, err := NewConfigurationSectionEx1(instance)
		
	if err != nil { return }
	newInstance = &ValidationSection {
	ConfigurationSection: tmp,
	}
	return
	}
	

	func NewValidationSectionEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery ) (newInstance *ValidationSection, err error) {tmp, err := NewConfigurationSectionEx6(hostName, wmiNamespace, userName, password, domainName, query)
		
	if err != nil { return }
	newInstance = &ValidationSection {
	ConfigurationSection: tmp,
	}
	return
	}
	

// SetValidateIntegratedModeConfiguration sets the value of ValidateIntegratedModeConfiguration for the instance
func (instance *ValidationSection) SetPropertyValidateIntegratedModeConfiguration(value bool) (err error) { 
    return instance.SetProperty("ValidateIntegratedModeConfiguration", (value))
}

// GetValidateIntegratedModeConfiguration gets the value of ValidateIntegratedModeConfiguration for the instance
func (instance *ValidationSection) GetPropertyValidateIntegratedModeConfiguration()(value bool, err error) { 
    retValue, err := instance.GetProperty("ValidateIntegratedModeConfiguration")
    if err != nil {
        return
    }
    if retValue == nil {
        // Doesn't have any value. Return empty
        return
    }
    
    valuetmp, ok := retValue.(bool); 
    if !ok {
        err = errors.Wrapf(errors.InvalidType, " bool is Invalid. Expected %s", reflect.TypeOf(retValue))
        return  
    }

    value = bool(valuetmp)

    return
}

