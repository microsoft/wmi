// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

// 
// Author:
//      Auto Generated on 6/6/2024 using wmigen
//      Source root.Microsoft.Windows.RemoteAccess
//////////////////////////////////////////////
package remoteaccess
import (
 "github.com/microsoft/wmi/pkg/base/query"
cim "github.com/microsoft/wmi/pkg/wmiinstance"
 "github.com/microsoft/wmi/pkg/base/instance"
 "github.com/microsoft/wmi/pkg/errors"
 "reflect"
)

// BgpCustomNetworkInfo struct
type BgpCustomNetworkInfo struct { 
	*cim.WmiInstance

	// 
	Interface []string

	// 
	Network []string

	// 
	RoutingDomain string
}

	func NewBgpCustomNetworkInfoEx1(instance *cim.WmiInstance) (newInstance *BgpCustomNetworkInfo, err error) {tmp, err := instance, nil
		
	if err != nil { return }
	newInstance = &BgpCustomNetworkInfo {
	WmiInstance: tmp,
	}
	return
	}
	

	func NewBgpCustomNetworkInfoEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery ) (newInstance *BgpCustomNetworkInfo, err error) {tmp, err := instance.GetWmiInstance(hostName, wmiNamespace, userName, password, domainName, query)
		
	if err != nil { return }
	newInstance = &BgpCustomNetworkInfo {
	WmiInstance: tmp,
	}
	return
	}
	

// SetInterface sets the value of Interface for the instance
func (instance *BgpCustomNetworkInfo) SetPropertyInterface(value []string) (err error) { 
    return instance.SetProperty("Interface", (value))
}

// GetInterface gets the value of Interface for the instance
func (instance *BgpCustomNetworkInfo) GetPropertyInterface()(value []string, err error) { 
    retValue, err := instance.GetProperty("Interface")
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

// SetNetwork sets the value of Network for the instance
func (instance *BgpCustomNetworkInfo) SetPropertyNetwork(value []string) (err error) { 
    return instance.SetProperty("Network", (value))
}

// GetNetwork gets the value of Network for the instance
func (instance *BgpCustomNetworkInfo) GetPropertyNetwork()(value []string, err error) { 
    retValue, err := instance.GetProperty("Network")
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

// SetRoutingDomain sets the value of RoutingDomain for the instance
func (instance *BgpCustomNetworkInfo) SetPropertyRoutingDomain(value string) (err error) { 
    return instance.SetProperty("RoutingDomain", (value))
}

// GetRoutingDomain gets the value of RoutingDomain for the instance
func (instance *BgpCustomNetworkInfo) GetPropertyRoutingDomain()(value string, err error) { 
    retValue, err := instance.GetProperty("RoutingDomain")
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

