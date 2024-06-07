// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

// 
// Author:
//      Auto Generated on 6/6/2024 using wmigen
//      Source root.Microsoft.Windows.RemoteAccess.Server
//////////////////////////////////////////////
package server
import (
 "github.com/microsoft/wmi/pkg/base/query"
cim "github.com/microsoft/wmi/pkg/wmiinstance"
 "github.com/microsoft/wmi/pkg/base/instance"
 "github.com/microsoft/wmi/pkg/errors"
 "reflect"
)

// DASiteAddresses struct
type DASiteAddresses struct { 
	*cim.WmiInstance

	// 
	Addresses []string
}

	func NewDASiteAddressesEx1(instance *cim.WmiInstance) (newInstance *DASiteAddresses, err error) {tmp, err := instance, nil
		
	if err != nil { return }
	newInstance = &DASiteAddresses {
	WmiInstance: tmp,
	}
	return
	}
	

	func NewDASiteAddressesEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery ) (newInstance *DASiteAddresses, err error) {tmp, err := instance.GetWmiInstance(hostName, wmiNamespace, userName, password, domainName, query)
		
	if err != nil { return }
	newInstance = &DASiteAddresses {
	WmiInstance: tmp,
	}
	return
	}
	

// SetAddresses sets the value of Addresses for the instance
func (instance *DASiteAddresses) SetPropertyAddresses(value []string) (err error) { 
    return instance.SetProperty("Addresses", (value))
}

// GetAddresses gets the value of Addresses for the instance
func (instance *DASiteAddresses) GetPropertyAddresses()(value []string, err error) { 
    retValue, err := instance.GetProperty("Addresses")
    if err != nil {
        return
    }
    if retValue == nil {
        // Doesn't have any value. Return empty
        return
    }
    
    for _, interfaceValue := range retValue.([]interface{}) {
        valuetmp, ok := interfaceValue.(string); 
        if !ok {
            err = errors.Wrapf(errors.InvalidType, " string is Invalid. Expected %s", reflect.TypeOf(interfaceValue))
            return  
        }
        value = append(value, string(valuetmp))
    }

    return
}

