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

// BgpRouteFlapDampeningConfig struct
type BgpRouteFlapDampeningConfig struct { 
	*cim.WmiInstance

	// 
	Ceiling uint32

	// 
	HalfLife uint32

	// 
	HalfLifeUnreachable uint32

	// 
	MaxSuppressTime uint32

	// 
	ReuseThreshold uint32

	// 
	RoutingDomain string

	// 
	SuppressThreshold uint32
}

	func NewBgpRouteFlapDampeningConfigEx1(instance *cim.WmiInstance) (newInstance *BgpRouteFlapDampeningConfig, err error) {tmp, err := instance, nil
		
	if err != nil { return }
	newInstance = &BgpRouteFlapDampeningConfig {
	WmiInstance: tmp,
	}
	return
	}
	

	func NewBgpRouteFlapDampeningConfigEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery ) (newInstance *BgpRouteFlapDampeningConfig, err error) {tmp, err := instance.GetWmiInstance(hostName, wmiNamespace, userName, password, domainName, query)
		
	if err != nil { return }
	newInstance = &BgpRouteFlapDampeningConfig {
	WmiInstance: tmp,
	}
	return
	}
	

// SetCeiling sets the value of Ceiling for the instance
func (instance *BgpRouteFlapDampeningConfig) SetPropertyCeiling(value uint32) (err error) { 
    return instance.SetProperty("Ceiling", (value))
}

// GetCeiling gets the value of Ceiling for the instance
func (instance *BgpRouteFlapDampeningConfig) GetPropertyCeiling()(value uint32, err error) { 
    retValue, err := instance.GetProperty("Ceiling")
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

// SetHalfLife sets the value of HalfLife for the instance
func (instance *BgpRouteFlapDampeningConfig) SetPropertyHalfLife(value uint32) (err error) { 
    return instance.SetProperty("HalfLife", (value))
}

// GetHalfLife gets the value of HalfLife for the instance
func (instance *BgpRouteFlapDampeningConfig) GetPropertyHalfLife()(value uint32, err error) { 
    retValue, err := instance.GetProperty("HalfLife")
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

// SetHalfLifeUnreachable sets the value of HalfLifeUnreachable for the instance
func (instance *BgpRouteFlapDampeningConfig) SetPropertyHalfLifeUnreachable(value uint32) (err error) { 
    return instance.SetProperty("HalfLifeUnreachable", (value))
}

// GetHalfLifeUnreachable gets the value of HalfLifeUnreachable for the instance
func (instance *BgpRouteFlapDampeningConfig) GetPropertyHalfLifeUnreachable()(value uint32, err error) { 
    retValue, err := instance.GetProperty("HalfLifeUnreachable")
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

// SetMaxSuppressTime sets the value of MaxSuppressTime for the instance
func (instance *BgpRouteFlapDampeningConfig) SetPropertyMaxSuppressTime(value uint32) (err error) { 
    return instance.SetProperty("MaxSuppressTime", (value))
}

// GetMaxSuppressTime gets the value of MaxSuppressTime for the instance
func (instance *BgpRouteFlapDampeningConfig) GetPropertyMaxSuppressTime()(value uint32, err error) { 
    retValue, err := instance.GetProperty("MaxSuppressTime")
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

// SetReuseThreshold sets the value of ReuseThreshold for the instance
func (instance *BgpRouteFlapDampeningConfig) SetPropertyReuseThreshold(value uint32) (err error) { 
    return instance.SetProperty("ReuseThreshold", (value))
}

// GetReuseThreshold gets the value of ReuseThreshold for the instance
func (instance *BgpRouteFlapDampeningConfig) GetPropertyReuseThreshold()(value uint32, err error) { 
    retValue, err := instance.GetProperty("ReuseThreshold")
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

// SetRoutingDomain sets the value of RoutingDomain for the instance
func (instance *BgpRouteFlapDampeningConfig) SetPropertyRoutingDomain(value string) (err error) { 
    return instance.SetProperty("RoutingDomain", (value))
}

// GetRoutingDomain gets the value of RoutingDomain for the instance
func (instance *BgpRouteFlapDampeningConfig) GetPropertyRoutingDomain()(value string, err error) { 
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

// SetSuppressThreshold sets the value of SuppressThreshold for the instance
func (instance *BgpRouteFlapDampeningConfig) SetPropertySuppressThreshold(value uint32) (err error) { 
    return instance.SetProperty("SuppressThreshold", (value))
}

// GetSuppressThreshold gets the value of SuppressThreshold for the instance
func (instance *BgpRouteFlapDampeningConfig) GetPropertySuppressThreshold()(value uint32, err error) { 
    retValue, err := instance.GetProperty("SuppressThreshold")
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

