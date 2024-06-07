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

// RouteUpdate struct
type RouteUpdate struct { 
	*cim.WmiInstance

	// 
	RouteInfo RouteInfo

	// 
	UpdateType string
}

	func NewRouteUpdateEx1(instance *cim.WmiInstance) (newInstance *RouteUpdate, err error) {tmp, err := instance, nil
		
	if err != nil { return }
	newInstance = &RouteUpdate {
	WmiInstance: tmp,
	}
	return
	}
	

	func NewRouteUpdateEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery ) (newInstance *RouteUpdate, err error) {tmp, err := instance.GetWmiInstance(hostName, wmiNamespace, userName, password, domainName, query)
		
	if err != nil { return }
	newInstance = &RouteUpdate {
	WmiInstance: tmp,
	}
	return
	}
	

// SetRouteInfo sets the value of RouteInfo for the instance
func (instance *RouteUpdate) SetPropertyRouteInfo(value RouteInfo) (err error) { 
    return instance.SetProperty("RouteInfo", (value))
}

// GetRouteInfo gets the value of RouteInfo for the instance
func (instance *RouteUpdate) GetPropertyRouteInfo()(value RouteInfo, err error) { 
    retValue, err := instance.GetProperty("RouteInfo")
    if err != nil {
        return
    }
    if retValue == nil {
        // Doesn't have any value. Return empty
        return
    }
    
    valuetmp, ok := retValue.(RouteInfo); 
    if !ok {
        err = errors.Wrapf(errors.InvalidType, " RouteInfo is Invalid. Expected %s", reflect.TypeOf(retValue))
        return  
    }

    value = RouteInfo(valuetmp)

    return
}

// SetUpdateType sets the value of UpdateType for the instance
func (instance *RouteUpdate) SetPropertyUpdateType(value string) (err error) { 
    return instance.SetProperty("UpdateType", (value))
}

// GetUpdateType gets the value of UpdateType for the instance
func (instance *RouteUpdate) GetPropertyUpdateType()(value string, err error) { 
    retValue, err := instance.GetProperty("UpdateType")
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

