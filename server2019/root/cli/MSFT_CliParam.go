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

// MSFT_CliParam struct
type MSFT_CliParam struct { 
	*cim.WmiInstance

	// 
	Default string

	// 
	Description string

	// 
	Optional bool

	// 
	ParaId string

	// 
	Qualifiers []MSFT_CliQualifier

	// 
	Type string
}

	func NewMSFT_CliParamEx1(instance *cim.WmiInstance) (newInstance *MSFT_CliParam, err error) {tmp, err := instance, nil
		
	if err != nil { return }
	newInstance = &MSFT_CliParam {
	WmiInstance: tmp,
	}
	return
	}
	

	func NewMSFT_CliParamEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery ) (newInstance *MSFT_CliParam, err error) {tmp, err := instance.GetWmiInstance(hostName, wmiNamespace, userName, password, domainName, query)
		
	if err != nil { return }
	newInstance = &MSFT_CliParam {
	WmiInstance: tmp,
	}
	return
	}
	

// SetDefault sets the value of Default for the instance
func (instance *MSFT_CliParam) SetPropertyDefault(value string) (err error) { 
    return instance.SetProperty("Default", (value))
}

// GetDefault gets the value of Default for the instance
func (instance *MSFT_CliParam) GetPropertyDefault()(value string, err error) { 
    retValue, err := instance.GetProperty("Default")
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
func (instance *MSFT_CliParam) SetPropertyDescription(value string) (err error) { 
    return instance.SetProperty("Description", (value))
}

// GetDescription gets the value of Description for the instance
func (instance *MSFT_CliParam) GetPropertyDescription()(value string, err error) { 
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

// SetOptional sets the value of Optional for the instance
func (instance *MSFT_CliParam) SetPropertyOptional(value bool) (err error) { 
    return instance.SetProperty("Optional", (value))
}

// GetOptional gets the value of Optional for the instance
func (instance *MSFT_CliParam) GetPropertyOptional()(value bool, err error) { 
    retValue, err := instance.GetProperty("Optional")
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

// SetParaId sets the value of ParaId for the instance
func (instance *MSFT_CliParam) SetPropertyParaId(value string) (err error) { 
    return instance.SetProperty("ParaId", (value))
}

// GetParaId gets the value of ParaId for the instance
func (instance *MSFT_CliParam) GetPropertyParaId()(value string, err error) { 
    retValue, err := instance.GetProperty("ParaId")
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
func (instance *MSFT_CliParam) SetPropertyQualifiers(value []MSFT_CliQualifier) (err error) { 
    return instance.SetProperty("Qualifiers", (value))
}

// GetQualifiers gets the value of Qualifiers for the instance
func (instance *MSFT_CliParam) GetPropertyQualifiers()(value []MSFT_CliQualifier, err error) { 
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

// SetType sets the value of Type for the instance
func (instance *MSFT_CliParam) SetPropertyType(value string) (err error) { 
    return instance.SetProperty("Type", (value))
}

// GetType gets the value of Type for the instance
func (instance *MSFT_CliParam) GetPropertyType()(value string, err error) { 
    retValue, err := instance.GetProperty("Type")
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

