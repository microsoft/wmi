// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

// 
// Author:
//      Auto Generated on 6/6/2024 using wmigen
//      Source root.StandardCimv2
//////////////////////////////////////////////
package standardcimv2
import (
 "github.com/microsoft/wmi/pkg/base/query"
cim "github.com/microsoft/wmi/pkg/wmiinstance"
 "github.com/microsoft/wmi/pkg/errors"
 "reflect"
)

// MSFT_DAConnectionStatus struct
type MSFT_DAConnectionStatus struct { 
	*MSFT_NetSettingData

	// 
	Status uint32

	// 
	Substatus uint32
}

	func NewMSFT_DAConnectionStatusEx1(instance *cim.WmiInstance) (newInstance *MSFT_DAConnectionStatus, err error) {tmp, err := NewMSFT_NetSettingDataEx1(instance)
		
	if err != nil { return }
	newInstance = &MSFT_DAConnectionStatus {
	MSFT_NetSettingData: tmp,
	}
	return
	}
	

	func NewMSFT_DAConnectionStatusEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery ) (newInstance *MSFT_DAConnectionStatus, err error) {tmp, err := NewMSFT_NetSettingDataEx6(hostName, wmiNamespace, userName, password, domainName, query)
		
	if err != nil { return }
	newInstance = &MSFT_DAConnectionStatus {
	MSFT_NetSettingData: tmp,
	}
	return
	}
	

// SetStatus sets the value of Status for the instance
func (instance *MSFT_DAConnectionStatus) SetPropertyStatus(value uint32) (err error) { 
    return instance.SetProperty("Status", (value))
}

// GetStatus gets the value of Status for the instance
func (instance *MSFT_DAConnectionStatus) GetPropertyStatus()(value uint32, err error) { 
    retValue, err := instance.GetProperty("Status")
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

// SetSubstatus sets the value of Substatus for the instance
func (instance *MSFT_DAConnectionStatus) SetPropertySubstatus(value uint32) (err error) { 
    return instance.SetProperty("Substatus", (value))
}

// GetSubstatus gets the value of Substatus for the instance
func (instance *MSFT_DAConnectionStatus) GetPropertySubstatus()(value uint32, err error) { 
    retValue, err := instance.GetProperty("Substatus")
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

