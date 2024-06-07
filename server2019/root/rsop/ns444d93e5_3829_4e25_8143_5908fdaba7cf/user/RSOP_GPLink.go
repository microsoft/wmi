// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

// 
// Author:
//      Auto Generated on 6/6/2024 using wmigen
//      Source root.RSOP.NS444D93E5_3829_4E25_8143_5908FDABA7CF.User
//////////////////////////////////////////////
package user
import (
 "github.com/microsoft/wmi/pkg/base/query"
cim "github.com/microsoft/wmi/pkg/wmiinstance"
 "github.com/microsoft/wmi/pkg/base/instance"
 "github.com/microsoft/wmi/pkg/errors"
 "reflect"
)

// RSOP_GPLink struct
type RSOP_GPLink struct { 
	*cim.WmiInstance

	// 
	appliedOrder uint32

	// 
	enabled bool

	// 
	GPO RSOP_GPO

	// 
	linkOrder uint32

	// 
	noOverride bool

	// 
	SOM RSOP_SOM

	// 
	somOrder uint32
}

	func NewRSOP_GPLinkEx1(instance *cim.WmiInstance) (newInstance *RSOP_GPLink, err error) {tmp, err := instance, nil
		
	if err != nil { return }
	newInstance = &RSOP_GPLink {
	WmiInstance: tmp,
	}
	return
	}
	

	func NewRSOP_GPLinkEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery ) (newInstance *RSOP_GPLink, err error) {tmp, err := instance.GetWmiInstance(hostName, wmiNamespace, userName, password, domainName, query)
		
	if err != nil { return }
	newInstance = &RSOP_GPLink {
	WmiInstance: tmp,
	}
	return
	}
	

// SetappliedOrder sets the value of appliedOrder for the instance
func (instance *RSOP_GPLink) SetPropertyappliedOrder(value uint32) (err error) { 
    return instance.SetProperty("appliedOrder", (value))
}

// GetappliedOrder gets the value of appliedOrder for the instance
func (instance *RSOP_GPLink) GetPropertyappliedOrder()(value uint32, err error) { 
    retValue, err := instance.GetProperty("appliedOrder")
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

// Setenabled sets the value of enabled for the instance
func (instance *RSOP_GPLink) SetPropertyenabled(value bool) (err error) { 
    return instance.SetProperty("enabled", (value))
}

// Getenabled gets the value of enabled for the instance
func (instance *RSOP_GPLink) GetPropertyenabled()(value bool, err error) { 
    retValue, err := instance.GetProperty("enabled")
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

// SetGPO sets the value of GPO for the instance
func (instance *RSOP_GPLink) SetPropertyGPO(value RSOP_GPO) (err error) { 
    return instance.SetProperty("GPO", (value))
}

// GetGPO gets the value of GPO for the instance
func (instance *RSOP_GPLink) GetPropertyGPO()(value RSOP_GPO, err error) { 
    retValue, err := instance.GetProperty("GPO")
    if err != nil {
        return
    }
    if retValue == nil {
        // Doesn't have any value. Return empty
        return
    }
    
    valuetmp, ok := retValue.(RSOP_GPO); 
    if !ok {
        err = errors.Wrapf(errors.InvalidType, " RSOP_GPO is Invalid. Expected %s", reflect.TypeOf(retValue))
        return  
    }

    value = RSOP_GPO(valuetmp)

    return
}

// SetlinkOrder sets the value of linkOrder for the instance
func (instance *RSOP_GPLink) SetPropertylinkOrder(value uint32) (err error) { 
    return instance.SetProperty("linkOrder", (value))
}

// GetlinkOrder gets the value of linkOrder for the instance
func (instance *RSOP_GPLink) GetPropertylinkOrder()(value uint32, err error) { 
    retValue, err := instance.GetProperty("linkOrder")
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

// SetnoOverride sets the value of noOverride for the instance
func (instance *RSOP_GPLink) SetPropertynoOverride(value bool) (err error) { 
    return instance.SetProperty("noOverride", (value))
}

// GetnoOverride gets the value of noOverride for the instance
func (instance *RSOP_GPLink) GetPropertynoOverride()(value bool, err error) { 
    retValue, err := instance.GetProperty("noOverride")
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

// SetSOM sets the value of SOM for the instance
func (instance *RSOP_GPLink) SetPropertySOM(value RSOP_SOM) (err error) { 
    return instance.SetProperty("SOM", (value))
}

// GetSOM gets the value of SOM for the instance
func (instance *RSOP_GPLink) GetPropertySOM()(value RSOP_SOM, err error) { 
    retValue, err := instance.GetProperty("SOM")
    if err != nil {
        return
    }
    if retValue == nil {
        // Doesn't have any value. Return empty
        return
    }
    
    valuetmp, ok := retValue.(RSOP_SOM); 
    if !ok {
        err = errors.Wrapf(errors.InvalidType, " RSOP_SOM is Invalid. Expected %s", reflect.TypeOf(retValue))
        return  
    }

    value = RSOP_SOM(valuetmp)

    return
}

// SetsomOrder sets the value of somOrder for the instance
func (instance *RSOP_GPLink) SetPropertysomOrder(value uint32) (err error) { 
    return instance.SetProperty("somOrder", (value))
}

// GetsomOrder gets the value of somOrder for the instance
func (instance *RSOP_GPLink) GetPropertysomOrder()(value uint32, err error) { 
    retValue, err := instance.GetProperty("somOrder")
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

