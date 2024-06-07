// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

// 
// Author:
//      Auto Generated on 6/6/2024 using wmigen
//      Source root.Microsoft.Windows.DesiredStateConfiguration
//////////////////////////////////////////////
package desiredstateconfiguration
import (
 "github.com/microsoft/wmi/pkg/base/query"
cim "github.com/microsoft/wmi/pkg/wmiinstance"
 "github.com/microsoft/wmi/pkg/base/instance"
 "github.com/microsoft/wmi/pkg/errors"
 "reflect"
)

// MSFT_KeyValuePair struct
type MSFT_KeyValuePair struct { 
	*cim.WmiInstance

	// 
	key string

	// 
	Value string
}

	func NewMSFT_KeyValuePairEx1(instance *cim.WmiInstance) (newInstance *MSFT_KeyValuePair, err error) {tmp, err := instance, nil
		
	if err != nil { return }
	newInstance = &MSFT_KeyValuePair {
	WmiInstance: tmp,
	}
	return
	}
	

	func NewMSFT_KeyValuePairEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery ) (newInstance *MSFT_KeyValuePair, err error) {tmp, err := instance.GetWmiInstance(hostName, wmiNamespace, userName, password, domainName, query)
		
	if err != nil { return }
	newInstance = &MSFT_KeyValuePair {
	WmiInstance: tmp,
	}
	return
	}
	

// Setkey sets the value of key for the instance
func (instance *MSFT_KeyValuePair) SetPropertykey(value string) (err error) { 
    return instance.SetProperty("key", (value))
}

// Getkey gets the value of key for the instance
func (instance *MSFT_KeyValuePair) GetPropertykey()(value string, err error) { 
    retValue, err := instance.GetProperty("key")
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

// SetValue sets the value of Value for the instance
func (instance *MSFT_KeyValuePair) SetPropertyValue(value string) (err error) { 
    return instance.SetProperty("Value", (value))
}

// GetValue gets the value of Value for the instance
func (instance *MSFT_KeyValuePair) GetPropertyValue()(value string, err error) { 
    retValue, err := instance.GetProperty("Value")
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

