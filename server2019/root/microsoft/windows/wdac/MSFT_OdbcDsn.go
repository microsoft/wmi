// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

// 
// Author:
//      Auto Generated on 6/6/2024 using wmigen
//      Source root.Microsoft.Windows.Wdac
//////////////////////////////////////////////
package wdac
import (
 "github.com/microsoft/wmi/pkg/base/query"
cim "github.com/microsoft/wmi/pkg/wmiinstance"
 "github.com/microsoft/wmi/pkg/base/instance"
 "github.com/microsoft/wmi/pkg/errors"
 "reflect"
)

// MSFT_OdbcDsn struct
type MSFT_OdbcDsn struct { 
	*cim.WmiInstance

	// 
	DriverName string

	// 
	DsnType string

	// 
	KeyValuePair []MSFT_OdbcKeyValuePair

	// 
	Name string

	// 
	Platform string
}

	func NewMSFT_OdbcDsnEx1(instance *cim.WmiInstance) (newInstance *MSFT_OdbcDsn, err error) {tmp, err := instance, nil
		
	if err != nil { return }
	newInstance = &MSFT_OdbcDsn {
	WmiInstance: tmp,
	}
	return
	}
	

	func NewMSFT_OdbcDsnEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery ) (newInstance *MSFT_OdbcDsn, err error) {tmp, err := instance.GetWmiInstance(hostName, wmiNamespace, userName, password, domainName, query)
		
	if err != nil { return }
	newInstance = &MSFT_OdbcDsn {
	WmiInstance: tmp,
	}
	return
	}
	

// SetDriverName sets the value of DriverName for the instance
func (instance *MSFT_OdbcDsn) SetPropertyDriverName(value string) (err error) { 
    return instance.SetProperty("DriverName", (value))
}

// GetDriverName gets the value of DriverName for the instance
func (instance *MSFT_OdbcDsn) GetPropertyDriverName()(value string, err error) { 
    retValue, err := instance.GetProperty("DriverName")
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

// SetDsnType sets the value of DsnType for the instance
func (instance *MSFT_OdbcDsn) SetPropertyDsnType(value string) (err error) { 
    return instance.SetProperty("DsnType", (value))
}

// GetDsnType gets the value of DsnType for the instance
func (instance *MSFT_OdbcDsn) GetPropertyDsnType()(value string, err error) { 
    retValue, err := instance.GetProperty("DsnType")
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

// SetKeyValuePair sets the value of KeyValuePair for the instance
func (instance *MSFT_OdbcDsn) SetPropertyKeyValuePair(value []MSFT_OdbcKeyValuePair) (err error) { 
    return instance.SetProperty("KeyValuePair", (value))
}

// GetKeyValuePair gets the value of KeyValuePair for the instance
func (instance *MSFT_OdbcDsn) GetPropertyKeyValuePair()(value []MSFT_OdbcKeyValuePair, err error) { 
    retValue, err := instance.GetProperty("KeyValuePair")
    if err != nil {
        return
    }
    if retValue == nil {
        // Doesn't have any value. Return empty
        return
    }
    
    for _, interfaceValue := range retValue.([]interface{}) {
        valuetmp, ok := interfaceValue.(MSFT_OdbcKeyValuePair); 
        if !ok {
            err = errors.Wrapf(errors.InvalidType, " MSFT_OdbcKeyValuePair is Invalid. Expected %s", reflect.TypeOf(interfaceValue))
            return  
        }
        value = append(value, MSFT_OdbcKeyValuePair(valuetmp))
    }

    return
}

// SetName sets the value of Name for the instance
func (instance *MSFT_OdbcDsn) SetPropertyName(value string) (err error) { 
    return instance.SetProperty("Name", (value))
}

// GetName gets the value of Name for the instance
func (instance *MSFT_OdbcDsn) GetPropertyName()(value string, err error) { 
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

// SetPlatform sets the value of Platform for the instance
func (instance *MSFT_OdbcDsn) SetPropertyPlatform(value string) (err error) { 
    return instance.SetProperty("Platform", (value))
}

// GetPlatform gets the value of Platform for the instance
func (instance *MSFT_OdbcDsn) GetPropertyPlatform()(value string, err error) { 
    retValue, err := instance.GetProperty("Platform")
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

