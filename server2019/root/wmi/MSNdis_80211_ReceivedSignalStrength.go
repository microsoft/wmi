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

// MSNdis_80211_ReceivedSignalStrength struct
type MSNdis_80211_ReceivedSignalStrength struct { 
	*MSNdis

	// 
	Active bool

	// 
	InstanceName string

	// 
	Ndis80211ReceivedSignalStrength int32
}

	func NewMSNdis_80211_ReceivedSignalStrengthEx1(instance *cim.WmiInstance) (newInstance *MSNdis_80211_ReceivedSignalStrength, err error) {tmp, err := NewMSNdisEx1(instance)
		
	if err != nil { return }
	newInstance = &MSNdis_80211_ReceivedSignalStrength {
	MSNdis: tmp,
	}
	return
	}
	

	func NewMSNdis_80211_ReceivedSignalStrengthEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery ) (newInstance *MSNdis_80211_ReceivedSignalStrength, err error) {tmp, err := NewMSNdisEx6(hostName, wmiNamespace, userName, password, domainName, query)
		
	if err != nil { return }
	newInstance = &MSNdis_80211_ReceivedSignalStrength {
	MSNdis: tmp,
	}
	return
	}
	

// SetActive sets the value of Active for the instance
func (instance *MSNdis_80211_ReceivedSignalStrength) SetPropertyActive(value bool) (err error) { 
    return instance.SetProperty("Active", (value))
}

// GetActive gets the value of Active for the instance
func (instance *MSNdis_80211_ReceivedSignalStrength) GetPropertyActive()(value bool, err error) { 
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
func (instance *MSNdis_80211_ReceivedSignalStrength) SetPropertyInstanceName(value string) (err error) { 
    return instance.SetProperty("InstanceName", (value))
}

// GetInstanceName gets the value of InstanceName for the instance
func (instance *MSNdis_80211_ReceivedSignalStrength) GetPropertyInstanceName()(value string, err error) { 
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

// SetNdis80211ReceivedSignalStrength sets the value of Ndis80211ReceivedSignalStrength for the instance
func (instance *MSNdis_80211_ReceivedSignalStrength) SetPropertyNdis80211ReceivedSignalStrength(value int32) (err error) { 
    return instance.SetProperty("Ndis80211ReceivedSignalStrength", (value))
}

// GetNdis80211ReceivedSignalStrength gets the value of Ndis80211ReceivedSignalStrength for the instance
func (instance *MSNdis_80211_ReceivedSignalStrength) GetPropertyNdis80211ReceivedSignalStrength()(value int32, err error) { 
    retValue, err := instance.GetProperty("Ndis80211ReceivedSignalStrength")
    if err != nil {
        return
    }
    if retValue == nil {
        // Doesn't have any value. Return empty
        return
    }
    
    valuetmp, ok := retValue.(int32); 
    if !ok {
        err = errors.Wrapf(errors.InvalidType, " int32 is Invalid. Expected %s", reflect.TypeOf(retValue))
        return  
    }

    value = int32(valuetmp)

    return
}

