// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

// 
// Author:
//      Auto Generated on 6/6/2024 using wmigen
//      Source root.RSOP.NS9606B916_B14A_4EB6_84F0_165F56E1DD34.User
//////////////////////////////////////////////
package user
import (
 "github.com/microsoft/wmi/pkg/base/query"
cim "github.com/microsoft/wmi/pkg/wmiinstance"
 "github.com/microsoft/wmi/pkg/base/instance"
 "github.com/microsoft/wmi/pkg/errors"
 "reflect"
)

// RSOP_SOM struct
type RSOP_SOM struct { 
	*cim.WmiInstance

	// 
	blocked bool

	// 
	blocking bool

	// 
	id string

	// 
	reason uint32

	// 
	SOMOrder uint32

	// 
	_type uint32
}

	func NewRSOP_SOMEx1(instance *cim.WmiInstance) (newInstance *RSOP_SOM, err error) {tmp, err := instance, nil
		
	if err != nil { return }
	newInstance = &RSOP_SOM {
	WmiInstance: tmp,
	}
	return
	}
	

	func NewRSOP_SOMEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery ) (newInstance *RSOP_SOM, err error) {tmp, err := instance.GetWmiInstance(hostName, wmiNamespace, userName, password, domainName, query)
		
	if err != nil { return }
	newInstance = &RSOP_SOM {
	WmiInstance: tmp,
	}
	return
	}
	

// Setblocked sets the value of blocked for the instance
func (instance *RSOP_SOM) SetPropertyblocked(value bool) (err error) { 
    return instance.SetProperty("blocked", (value))
}

// Getblocked gets the value of blocked for the instance
func (instance *RSOP_SOM) GetPropertyblocked()(value bool, err error) { 
    retValue, err := instance.GetProperty("blocked")
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

// Setblocking sets the value of blocking for the instance
func (instance *RSOP_SOM) SetPropertyblocking(value bool) (err error) { 
    return instance.SetProperty("blocking", (value))
}

// Getblocking gets the value of blocking for the instance
func (instance *RSOP_SOM) GetPropertyblocking()(value bool, err error) { 
    retValue, err := instance.GetProperty("blocking")
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

// Setid sets the value of id for the instance
func (instance *RSOP_SOM) SetPropertyid(value string) (err error) { 
    return instance.SetProperty("id", (value))
}

// Getid gets the value of id for the instance
func (instance *RSOP_SOM) GetPropertyid()(value string, err error) { 
    retValue, err := instance.GetProperty("id")
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

// Setreason sets the value of reason for the instance
func (instance *RSOP_SOM) SetPropertyreason(value uint32) (err error) { 
    return instance.SetProperty("reason", (value))
}

// Getreason gets the value of reason for the instance
func (instance *RSOP_SOM) GetPropertyreason()(value uint32, err error) { 
    retValue, err := instance.GetProperty("reason")
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

// SetSOMOrder sets the value of SOMOrder for the instance
func (instance *RSOP_SOM) SetPropertySOMOrder(value uint32) (err error) { 
    return instance.SetProperty("SOMOrder", (value))
}

// GetSOMOrder gets the value of SOMOrder for the instance
func (instance *RSOP_SOM) GetPropertySOMOrder()(value uint32, err error) { 
    retValue, err := instance.GetProperty("SOMOrder")
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

// Settype sets the value of type for the instance
func (instance *RSOP_SOM) SetPropertytype(value uint32) (err error) { 
    return instance.SetProperty("type", (value))
}

// Gettype gets the value of type for the instance
func (instance *RSOP_SOM) GetPropertytype()(value uint32, err error) { 
    retValue, err := instance.GetProperty("type")
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

