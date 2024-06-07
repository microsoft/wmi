// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

// 
// Author:
//      Auto Generated on 6/6/2024 using wmigen
//      Source root.CIMV2.mdm
//////////////////////////////////////////////
package mdm
import (
 "github.com/microsoft/wmi/pkg/base/query"
cim "github.com/microsoft/wmi/pkg/wmiinstance"
 "github.com/microsoft/wmi/pkg/base/instance"
 "github.com/microsoft/wmi/pkg/errors"
 "reflect"
)

// MDM_BrowserSecurityZones struct
type MDM_BrowserSecurityZones struct { 
	*cim.WmiInstance

	// 
	Exists bool

	// 
	Namespace string

	// 
	Zone uint32
}

	func NewMDM_BrowserSecurityZonesEx1(instance *cim.WmiInstance) (newInstance *MDM_BrowserSecurityZones, err error) {tmp, err := instance, nil
		
	if err != nil { return }
	newInstance = &MDM_BrowserSecurityZones {
	WmiInstance: tmp,
	}
	return
	}
	

	func NewMDM_BrowserSecurityZonesEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery ) (newInstance *MDM_BrowserSecurityZones, err error) {tmp, err := instance.GetWmiInstance(hostName, wmiNamespace, userName, password, domainName, query)
		
	if err != nil { return }
	newInstance = &MDM_BrowserSecurityZones {
	WmiInstance: tmp,
	}
	return
	}
	

// SetExists sets the value of Exists for the instance
func (instance *MDM_BrowserSecurityZones) SetPropertyExists(value bool) (err error) { 
    return instance.SetProperty("Exists", (value))
}

// GetExists gets the value of Exists for the instance
func (instance *MDM_BrowserSecurityZones) GetPropertyExists()(value bool, err error) { 
    retValue, err := instance.GetProperty("Exists")
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

// SetNamespace sets the value of Namespace for the instance
func (instance *MDM_BrowserSecurityZones) SetPropertyNamespace(value string) (err error) { 
    return instance.SetProperty("Namespace", (value))
}

// GetNamespace gets the value of Namespace for the instance
func (instance *MDM_BrowserSecurityZones) GetPropertyNamespace()(value string, err error) { 
    retValue, err := instance.GetProperty("Namespace")
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

// SetZone sets the value of Zone for the instance
func (instance *MDM_BrowserSecurityZones) SetPropertyZone(value uint32) (err error) { 
    return instance.SetProperty("Zone", (value))
}

// GetZone gets the value of Zone for the instance
func (instance *MDM_BrowserSecurityZones) GetPropertyZone()(value uint32, err error) { 
    retValue, err := instance.GetProperty("Zone")
    if err != nil {
        return
    }
    if retValue == nil {
        // Doesn't have any value. Return empty
        return
    }
    
    valuetmp, ok := retValue.(uint32); 
    if !ok {
        err = errors.Wrapf(errors.InvalidType, " uint32 is Invalid. Expected %s", reflect.TypeOf(retValue))
        return  
    }

    value = uint32(valuetmp)

    return
}

