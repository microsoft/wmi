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

// MSNdis_80211_ConfigurationInfo struct
type MSNdis_80211_ConfigurationInfo struct { 
	*MSNdis

	// 
	ATIMWindow uint32

	// 
	BeaconPeriod uint32

	// 
	ConfigLength uint32

	// 
	DSConfig uint32

	// 
	FHConfig MSNdis_80211_ConfigurationFH
}

	func NewMSNdis_80211_ConfigurationInfoEx1(instance *cim.WmiInstance) (newInstance *MSNdis_80211_ConfigurationInfo, err error) {tmp, err := NewMSNdisEx1(instance)
		
	if err != nil { return }
	newInstance = &MSNdis_80211_ConfigurationInfo {
	MSNdis: tmp,
	}
	return
	}
	

	func NewMSNdis_80211_ConfigurationInfoEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery ) (newInstance *MSNdis_80211_ConfigurationInfo, err error) {tmp, err := NewMSNdisEx6(hostName, wmiNamespace, userName, password, domainName, query)
		
	if err != nil { return }
	newInstance = &MSNdis_80211_ConfigurationInfo {
	MSNdis: tmp,
	}
	return
	}
	

// SetATIMWindow sets the value of ATIMWindow for the instance
func (instance *MSNdis_80211_ConfigurationInfo) SetPropertyATIMWindow(value uint32) (err error) { 
    return instance.SetProperty("ATIMWindow", (value))
}

// GetATIMWindow gets the value of ATIMWindow for the instance
func (instance *MSNdis_80211_ConfigurationInfo) GetPropertyATIMWindow()(value uint32, err error) { 
    retValue, err := instance.GetProperty("ATIMWindow")
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

// SetBeaconPeriod sets the value of BeaconPeriod for the instance
func (instance *MSNdis_80211_ConfigurationInfo) SetPropertyBeaconPeriod(value uint32) (err error) { 
    return instance.SetProperty("BeaconPeriod", (value))
}

// GetBeaconPeriod gets the value of BeaconPeriod for the instance
func (instance *MSNdis_80211_ConfigurationInfo) GetPropertyBeaconPeriod()(value uint32, err error) { 
    retValue, err := instance.GetProperty("BeaconPeriod")
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

// SetConfigLength sets the value of ConfigLength for the instance
func (instance *MSNdis_80211_ConfigurationInfo) SetPropertyConfigLength(value uint32) (err error) { 
    return instance.SetProperty("ConfigLength", (value))
}

// GetConfigLength gets the value of ConfigLength for the instance
func (instance *MSNdis_80211_ConfigurationInfo) GetPropertyConfigLength()(value uint32, err error) { 
    retValue, err := instance.GetProperty("ConfigLength")
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

// SetDSConfig sets the value of DSConfig for the instance
func (instance *MSNdis_80211_ConfigurationInfo) SetPropertyDSConfig(value uint32) (err error) { 
    return instance.SetProperty("DSConfig", (value))
}

// GetDSConfig gets the value of DSConfig for the instance
func (instance *MSNdis_80211_ConfigurationInfo) GetPropertyDSConfig()(value uint32, err error) { 
    retValue, err := instance.GetProperty("DSConfig")
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

// SetFHConfig sets the value of FHConfig for the instance
func (instance *MSNdis_80211_ConfigurationInfo) SetPropertyFHConfig(value MSNdis_80211_ConfigurationFH) (err error) { 
    return instance.SetProperty("FHConfig", (value))
}

// GetFHConfig gets the value of FHConfig for the instance
func (instance *MSNdis_80211_ConfigurationInfo) GetPropertyFHConfig()(value MSNdis_80211_ConfigurationFH, err error) { 
    retValue, err := instance.GetProperty("FHConfig")
    if err != nil {
        return
    }
    if retValue == nil {
        // Doesn't have any value. Return empty
        return
    }
    
    valuetmp, ok := retValue.(MSNdis_80211_ConfigurationFH); 
    if !ok {
        err = errors.Wrapf(errors.InvalidType, " MSNdis_80211_ConfigurationFH is Invalid. Expected %s", reflect.TypeOf(retValue))
        return  
    }

    value = MSNdis_80211_ConfigurationFH(valuetmp)

    return
}

