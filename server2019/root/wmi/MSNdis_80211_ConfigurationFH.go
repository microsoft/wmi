// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

// 
// Author:
//      Auto Generated on 6/6/2024 using wmigen
//      Source root.WMI
//////////////////////////////////////////////
package wmi
import (
 "github.com/microsoft/wmi/pkg/base/query"
cim "github.com/microsoft/wmi/pkg/wmiinstance"
 "github.com/microsoft/wmi/pkg/errors"
 "reflect"
)

// MSNdis_80211_ConfigurationFH struct
type MSNdis_80211_ConfigurationFH struct { 
	*MSNdis

	// 
	DwellTime uint32

	// 
	FHLength uint32

	// 
	HopPattern uint32

	// 
	HopSet uint32
}

	func NewMSNdis_80211_ConfigurationFHEx1(instance *cim.WmiInstance) (newInstance *MSNdis_80211_ConfigurationFH, err error) {tmp, err := NewMSNdisEx1(instance)
		
	if err != nil { return }
	newInstance = &MSNdis_80211_ConfigurationFH {
	MSNdis: tmp,
	}
	return
	}
	

	func NewMSNdis_80211_ConfigurationFHEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery ) (newInstance *MSNdis_80211_ConfigurationFH, err error) {tmp, err := NewMSNdisEx6(hostName, wmiNamespace, userName, password, domainName, query)
		
	if err != nil { return }
	newInstance = &MSNdis_80211_ConfigurationFH {
	MSNdis: tmp,
	}
	return
	}
	

// SetDwellTime sets the value of DwellTime for the instance
func (instance *MSNdis_80211_ConfigurationFH) SetPropertyDwellTime(value uint32) (err error) { 
    return instance.SetProperty("DwellTime", (value))
}

// GetDwellTime gets the value of DwellTime for the instance
func (instance *MSNdis_80211_ConfigurationFH) GetPropertyDwellTime()(value uint32, err error) { 
    retValue, err := instance.GetProperty("DwellTime")
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

// SetFHLength sets the value of FHLength for the instance
func (instance *MSNdis_80211_ConfigurationFH) SetPropertyFHLength(value uint32) (err error) { 
    return instance.SetProperty("FHLength", (value))
}

// GetFHLength gets the value of FHLength for the instance
func (instance *MSNdis_80211_ConfigurationFH) GetPropertyFHLength()(value uint32, err error) { 
    retValue, err := instance.GetProperty("FHLength")
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

// SetHopPattern sets the value of HopPattern for the instance
func (instance *MSNdis_80211_ConfigurationFH) SetPropertyHopPattern(value uint32) (err error) { 
    return instance.SetProperty("HopPattern", (value))
}

// GetHopPattern gets the value of HopPattern for the instance
func (instance *MSNdis_80211_ConfigurationFH) GetPropertyHopPattern()(value uint32, err error) { 
    retValue, err := instance.GetProperty("HopPattern")
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

// SetHopSet sets the value of HopSet for the instance
func (instance *MSNdis_80211_ConfigurationFH) SetPropertyHopSet(value uint32) (err error) { 
    return instance.SetProperty("HopSet", (value))
}

// GetHopSet gets the value of HopSet for the instance
func (instance *MSNdis_80211_ConfigurationFH) GetPropertyHopSet()(value uint32, err error) { 
    retValue, err := instance.GetProperty("HopSet")
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

