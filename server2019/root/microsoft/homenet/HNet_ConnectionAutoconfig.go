// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

// 
// Author:
//      Auto Generated on 6/6/2024 using wmigen
//      Source root.Microsoft.HomeNet
//////////////////////////////////////////////
package homenet
import (
 "github.com/microsoft/wmi/pkg/base/query"
cim "github.com/microsoft/wmi/pkg/wmiinstance"
 "github.com/microsoft/wmi/pkg/base/instance"
 "github.com/microsoft/wmi/pkg/errors"
 "reflect"
)

// HNet_ConnectionAutoconfig struct
type HNet_ConnectionAutoconfig struct { 
	*cim.WmiInstance

	// 
	Connection HNet_Connection
}

	func NewHNet_ConnectionAutoconfigEx1(instance *cim.WmiInstance) (newInstance *HNet_ConnectionAutoconfig, err error) {tmp, err := instance, nil
		
	if err != nil { return }
	newInstance = &HNet_ConnectionAutoconfig {
	WmiInstance: tmp,
	}
	return
	}
	

	func NewHNet_ConnectionAutoconfigEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery ) (newInstance *HNet_ConnectionAutoconfig, err error) {tmp, err := instance.GetWmiInstance(hostName, wmiNamespace, userName, password, domainName, query)
		
	if err != nil { return }
	newInstance = &HNet_ConnectionAutoconfig {
	WmiInstance: tmp,
	}
	return
	}
	

// SetConnection sets the value of Connection for the instance
func (instance *HNet_ConnectionAutoconfig) SetPropertyConnection(value HNet_Connection) (err error) { 
    return instance.SetProperty("Connection", (value))
}

// GetConnection gets the value of Connection for the instance
func (instance *HNet_ConnectionAutoconfig) GetPropertyConnection()(value HNet_Connection, err error) { 
    retValue, err := instance.GetProperty("Connection")
    if err != nil {
        return
    }
    if retValue == nil {
        // Doesn't have any value. Return empty
        return
    }
    
    valuetmp, ok := retValue.(HNet_Connection); 
    if !ok {
        err = errors.Wrapf(errors.InvalidType, " HNet_Connection is Invalid. Expected %s", reflect.TypeOf(retValue))
        return  
    }

    value = HNet_Connection(valuetmp)

    return
}

