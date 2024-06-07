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

// MSNdis_80211_ReceiveAntennaSelected struct
type MSNdis_80211_ReceiveAntennaSelected struct { 
	*MSNdis

	// 
	Active bool

	// 
	InstanceName string

	// 
	Ndis80211ReceiveAntennaSelected uint32
}

	func NewMSNdis_80211_ReceiveAntennaSelectedEx1(instance *cim.WmiInstance) (newInstance *MSNdis_80211_ReceiveAntennaSelected, err error) {tmp, err := NewMSNdisEx1(instance)
		
	if err != nil { return }
	newInstance = &MSNdis_80211_ReceiveAntennaSelected {
	MSNdis: tmp,
	}
	return
	}
	

	func NewMSNdis_80211_ReceiveAntennaSelectedEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery ) (newInstance *MSNdis_80211_ReceiveAntennaSelected, err error) {tmp, err := NewMSNdisEx6(hostName, wmiNamespace, userName, password, domainName, query)
		
	if err != nil { return }
	newInstance = &MSNdis_80211_ReceiveAntennaSelected {
	MSNdis: tmp,
	}
	return
	}
	

// SetActive sets the value of Active for the instance
func (instance *MSNdis_80211_ReceiveAntennaSelected) SetPropertyActive(value bool) (err error) { 
    return instance.SetProperty("Active", (value))
}

// GetActive gets the value of Active for the instance
func (instance *MSNdis_80211_ReceiveAntennaSelected) GetPropertyActive()(value bool, err error) { 
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
func (instance *MSNdis_80211_ReceiveAntennaSelected) SetPropertyInstanceName(value string) (err error) { 
    return instance.SetProperty("InstanceName", (value))
}

// GetInstanceName gets the value of InstanceName for the instance
func (instance *MSNdis_80211_ReceiveAntennaSelected) GetPropertyInstanceName()(value string, err error) { 
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

// SetNdis80211ReceiveAntennaSelected sets the value of Ndis80211ReceiveAntennaSelected for the instance
func (instance *MSNdis_80211_ReceiveAntennaSelected) SetPropertyNdis80211ReceiveAntennaSelected(value uint32) (err error) { 
    return instance.SetProperty("Ndis80211ReceiveAntennaSelected", (value))
}

// GetNdis80211ReceiveAntennaSelected gets the value of Ndis80211ReceiveAntennaSelected for the instance
func (instance *MSNdis_80211_ReceiveAntennaSelected) GetPropertyNdis80211ReceiveAntennaSelected()(value uint32, err error) { 
    retValue, err := instance.GetProperty("Ndis80211ReceiveAntennaSelected")
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

