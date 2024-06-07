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

// MSNdis_80211_Configuration struct
type MSNdis_80211_Configuration struct { 
	*MSNdis

	// 
	Active bool

	// 
	InstanceName string

	// 
	Ndis80211Config MSNdis_80211_ConfigurationInfo
}

	func NewMSNdis_80211_ConfigurationEx1(instance *cim.WmiInstance) (newInstance *MSNdis_80211_Configuration, err error) {tmp, err := NewMSNdisEx1(instance)
		
	if err != nil { return }
	newInstance = &MSNdis_80211_Configuration {
	MSNdis: tmp,
	}
	return
	}
	

	func NewMSNdis_80211_ConfigurationEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery ) (newInstance *MSNdis_80211_Configuration, err error) {tmp, err := NewMSNdisEx6(hostName, wmiNamespace, userName, password, domainName, query)
		
	if err != nil { return }
	newInstance = &MSNdis_80211_Configuration {
	MSNdis: tmp,
	}
	return
	}
	

// SetActive sets the value of Active for the instance
func (instance *MSNdis_80211_Configuration) SetPropertyActive(value bool) (err error) { 
    return instance.SetProperty("Active", (value))
}

// GetActive gets the value of Active for the instance
func (instance *MSNdis_80211_Configuration) GetPropertyActive()(value bool, err error) { 
    retValue, err := instance.GetProperty("Active")
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

// SetInstanceName sets the value of InstanceName for the instance
func (instance *MSNdis_80211_Configuration) SetPropertyInstanceName(value string) (err error) { 
    return instance.SetProperty("InstanceName", (value))
}

// GetInstanceName gets the value of InstanceName for the instance
func (instance *MSNdis_80211_Configuration) GetPropertyInstanceName()(value string, err error) { 
    retValue, err := instance.GetProperty("InstanceName")
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

// SetNdis80211Config sets the value of Ndis80211Config for the instance
func (instance *MSNdis_80211_Configuration) SetPropertyNdis80211Config(value MSNdis_80211_ConfigurationInfo) (err error) { 
    return instance.SetProperty("Ndis80211Config", (value))
}

// GetNdis80211Config gets the value of Ndis80211Config for the instance
func (instance *MSNdis_80211_Configuration) GetPropertyNdis80211Config()(value MSNdis_80211_ConfigurationInfo, err error) { 
    retValue, err := instance.GetProperty("Ndis80211Config")
    if err != nil {
        return
    }
    if retValue == nil {
        // Doesn't have any value. Return empty
        return
    }
    
    valuetmp, ok := retValue.(MSNdis_80211_ConfigurationInfo); 
    if !ok {
        err = errors.Wrapf(errors.InvalidType, " MSNdis_80211_ConfigurationInfo is Invalid. Expected %s", reflect.TypeOf(retValue))
        return  
    }

    value = MSNdis_80211_ConfigurationInfo(valuetmp)

    return
}

