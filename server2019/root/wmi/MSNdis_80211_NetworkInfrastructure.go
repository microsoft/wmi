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

// MSNdis_80211_NetworkInfrastructure struct
type MSNdis_80211_NetworkInfrastructure struct { 
	*MSNdis

	// 
	Ndis80211NetworkInfrastructure NetworkInfrastructure_Ndis80211NetworkInfrastructure
}

	func NewMSNdis_80211_NetworkInfrastructureEx1(instance *cim.WmiInstance) (newInstance *MSNdis_80211_NetworkInfrastructure, err error) {tmp, err := NewMSNdisEx1(instance)
		
	if err != nil { return }
	newInstance = &MSNdis_80211_NetworkInfrastructure {
	MSNdis: tmp,
	}
	return
	}
	

	func NewMSNdis_80211_NetworkInfrastructureEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery ) (newInstance *MSNdis_80211_NetworkInfrastructure, err error) {tmp, err := NewMSNdisEx6(hostName, wmiNamespace, userName, password, domainName, query)
		
	if err != nil { return }
	newInstance = &MSNdis_80211_NetworkInfrastructure {
	MSNdis: tmp,
	}
	return
	}
	

// SetNdis80211NetworkInfrastructure sets the value of Ndis80211NetworkInfrastructure for the instance
func (instance *MSNdis_80211_NetworkInfrastructure) SetPropertyNdis80211NetworkInfrastructure(value NetworkInfrastructure_Ndis80211NetworkInfrastructure) (err error) { 
    return instance.SetProperty("Ndis80211NetworkInfrastructure", (value))
}

// GetNdis80211NetworkInfrastructure gets the value of Ndis80211NetworkInfrastructure for the instance
func (instance *MSNdis_80211_NetworkInfrastructure) GetPropertyNdis80211NetworkInfrastructure()(value NetworkInfrastructure_Ndis80211NetworkInfrastructure, err error) { 
    retValue, err := instance.GetProperty("Ndis80211NetworkInfrastructure")
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

    value = NetworkInfrastructure_Ndis80211NetworkInfrastructure(valuetmp)

    return
}

