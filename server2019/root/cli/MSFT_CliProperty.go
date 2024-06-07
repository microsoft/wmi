// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

// 
// Author:
//      Auto Generated on 6/6/2024 using wmigen
//      Source root.Cli
//////////////////////////////////////////////
package cli
import (
 "github.com/microsoft/wmi/pkg/base/query"
cim "github.com/microsoft/wmi/pkg/wmiinstance"
 "github.com/microsoft/wmi/pkg/base/instance"
 "github.com/microsoft/wmi/pkg/errors"
 "reflect"
)

// MSFT_CliProperty struct
type MSFT_CliProperty struct { 
	*cim.WmiInstance

	// 
	Derivation string

	// 
	Description string

	// 
	Name string

	// 
	Qualifiers []MSFT_CliQualifier
}

	func NewMSFT_CliPropertyEx1(instance *cim.WmiInstance) (newInstance *MSFT_CliProperty, err error) {tmp, err := instance, nil
		
	if err != nil { return }
	newInstance = &MSFT_CliProperty {
	WmiInstance: tmp,
	}
	return
	}
	

	func NewMSFT_CliPropertyEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery ) (newInstance *MSFT_CliProperty, err error) {tmp, err := instance.GetWmiInstance(hostName, wmiNamespace, userName, password, domainName, query)
		
	if err != nil { return }
	newInstance = &MSFT_CliProperty {
	WmiInstance: tmp,
	}
	return
	}
	

// SetDerivation sets the value of Derivation for the instance
func (instance *MSFT_CliProperty) SetPropertyDerivation(value string) (err error) { 
    return instance.SetProperty("Derivation", (value))
}

// GetDerivation gets the value of Derivation for the instance
func (instance *MSFT_CliProperty) GetPropertyDerivation()(value string, err error) { 
    retValue, err := instance.GetProperty("Derivation")
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

// SetDescription sets the value of Description for the instance
func (instance *MSFT_CliProperty) SetPropertyDescription(value string) (err error) { 
    return instance.SetProperty("Description", (value))
}

// GetDescription gets the value of Description for the instance
func (instance *MSFT_CliProperty) GetPropertyDescription()(value string, err error) { 
    retValue, err := instance.GetProperty("Description")
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

// SetName sets the value of Name for the instance
func (instance *MSFT_CliProperty) SetPropertyName(value string) (err error) { 
    return instance.SetProperty("Name", (value))
}

// GetName gets the value of Name for the instance
func (instance *MSFT_CliProperty) GetPropertyName()(value string, err error) { 
    retValue, err := instance.GetProperty("Name")
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

// SetQualifiers sets the value of Qualifiers for the instance
func (instance *MSFT_CliProperty) SetPropertyQualifiers(value []MSFT_CliQualifier) (err error) { 
    return instance.SetProperty("Qualifiers", (value))
}

// GetQualifiers gets the value of Qualifiers for the instance
func (instance *MSFT_CliProperty) GetPropertyQualifiers()(value []MSFT_CliQualifier, err error) { 
    retValue, err := instance.GetProperty("Qualifiers")
    if err != nil {
        return
    }
    if retValue == nil {
        // Doesn't have any value. Return empty
        return
    }
    
    for _, interfaceValue := range retValue.([]interface{}) {
        valuetmp, ok := interfaceValue.(MSFT_CliQualifier); 
        if !ok {
            err = errors.Wrapf(errors.InvalidType, " MSFT_CliQualifier is Invalid. Expected %s", reflect.TypeOf(interfaceValue))
            return  
        }
        value = append(value, MSFT_CliQualifier(valuetmp))
    }

    return
}

