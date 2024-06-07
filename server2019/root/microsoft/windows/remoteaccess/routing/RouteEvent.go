// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

// 
// Author:
//      Auto Generated on 6/6/2024 using wmigen
//      Source root.Microsoft.Windows.RemoteAccess.Routing
//////////////////////////////////////////////
package routing
import (
 "github.com/microsoft/wmi/pkg/base/query"
cim "github.com/microsoft/wmi/pkg/wmiinstance"
 "github.com/microsoft/wmi/pkg/base/instance"
 "github.com/microsoft/wmi/pkg/errors"
 "reflect"
)

// RouteEvent struct
type RouteEvent struct { 
	*cim.WmiInstance

	// 
	InitialUpdate bool

	// 
	RouteUpdates []RouteUpdate

	// 
	RoutingDomainID string
}

	func NewRouteEventEx1(instance *cim.WmiInstance) (newInstance *RouteEvent, err error) {tmp, err := instance, nil
		
	if err != nil { return }
	newInstance = &RouteEvent {
	WmiInstance: tmp,
	}
	return
	}
	

	func NewRouteEventEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery ) (newInstance *RouteEvent, err error) {tmp, err := instance.GetWmiInstance(hostName, wmiNamespace, userName, password, domainName, query)
		
	if err != nil { return }
	newInstance = &RouteEvent {
	WmiInstance: tmp,
	}
	return
	}
	

// SetInitialUpdate sets the value of InitialUpdate for the instance
func (instance *RouteEvent) SetPropertyInitialUpdate(value bool) (err error) { 
    return instance.SetProperty("InitialUpdate", (value))
}

// GetInitialUpdate gets the value of InitialUpdate for the instance
func (instance *RouteEvent) GetPropertyInitialUpdate()(value bool, err error) { 
    retValue, err := instance.GetProperty("InitialUpdate")
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

// SetRouteUpdates sets the value of RouteUpdates for the instance
func (instance *RouteEvent) SetPropertyRouteUpdates(value []RouteUpdate) (err error) { 
    return instance.SetProperty("RouteUpdates", (value))
}

// GetRouteUpdates gets the value of RouteUpdates for the instance
func (instance *RouteEvent) GetPropertyRouteUpdates()(value []RouteUpdate, err error) { 
    retValue, err := instance.GetProperty("RouteUpdates")
    if err != nil {
        return
    }
    if retValue == nil {
        // Doesn't have any value. Return empty
        return
    }
    
    for _, interfaceValue := range retValue.([]interface{}) {
        valuetmp, ok := interfaceValue.(RouteUpdate); 
        if !ok {
            err = errors.Wrapf(errors.InvalidType, " RouteUpdate is Invalid. Expected %s", reflect.TypeOf(interfaceValue))
            return  
        }
        value = append(value, RouteUpdate(valuetmp))
    }

    return
}

// SetRoutingDomainID sets the value of RoutingDomainID for the instance
func (instance *RouteEvent) SetPropertyRoutingDomainID(value string) (err error) { 
    return instance.SetProperty("RoutingDomainID", (value))
}

// GetRoutingDomainID gets the value of RoutingDomainID for the instance
func (instance *RouteEvent) GetPropertyRoutingDomainID()(value string, err error) { 
    retValue, err := instance.GetProperty("RoutingDomainID")
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

