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

// MSFT_NetQosDcbxSettingData struct
type MSFT_NetQosDcbxSettingData struct { 
	*MSFT_NetSettingData

	// 
	InterfaceAlias string

	// 
	InterfaceIndex uint32

	// 
	PolicySet uint8

	// 
	Willing bool
}

	func NewMSFT_NetQosDcbxSettingDataEx1(instance *cim.WmiInstance) (newInstance *MSFT_NetQosDcbxSettingData, err error) {tmp, err := NewMSFT_NetSettingDataEx1(instance)
		
	if err != nil { return }
	newInstance = &MSFT_NetQosDcbxSettingData {
	MSFT_NetSettingData: tmp,
	}
	return
	}
	

	func NewMSFT_NetQosDcbxSettingDataEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery ) (newInstance *MSFT_NetQosDcbxSettingData, err error) {tmp, err := NewMSFT_NetSettingDataEx6(hostName, wmiNamespace, userName, password, domainName, query)
		
	if err != nil { return }
	newInstance = &MSFT_NetQosDcbxSettingData {
	MSFT_NetSettingData: tmp,
	}
	return
	}
	

// SetInterfaceAlias sets the value of InterfaceAlias for the instance
func (instance *MSFT_NetQosDcbxSettingData) SetPropertyInterfaceAlias(value string) (err error) { 
    return instance.SetProperty("InterfaceAlias", (value))
}

// GetInterfaceAlias gets the value of InterfaceAlias for the instance
func (instance *MSFT_NetQosDcbxSettingData) GetPropertyInterfaceAlias()(value string, err error) { 
    retValue, err := instance.GetProperty("InterfaceAlias")
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

// SetInterfaceIndex sets the value of InterfaceIndex for the instance
func (instance *MSFT_NetQosDcbxSettingData) SetPropertyInterfaceIndex(value uint32) (err error) { 
    return instance.SetProperty("InterfaceIndex", (value))
}

// GetInterfaceIndex gets the value of InterfaceIndex for the instance
func (instance *MSFT_NetQosDcbxSettingData) GetPropertyInterfaceIndex()(value uint32, err error) { 
    retValue, err := instance.GetProperty("InterfaceIndex")
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

// SetPolicySet sets the value of PolicySet for the instance
func (instance *MSFT_NetQosDcbxSettingData) SetPropertyPolicySet(value uint8) (err error) { 
    return instance.SetProperty("PolicySet", (value))
}

// GetPolicySet gets the value of PolicySet for the instance
func (instance *MSFT_NetQosDcbxSettingData) GetPropertyPolicySet()(value uint8, err error) { 
    retValue, err := instance.GetProperty("PolicySet")
    if err != nil {
        return
    }
    if retValue == nil {
        // Doesn't have any value. Return empty
        return
    }
    
    valuetmp, ok := retValue.(uint8); 
    if !ok {
        err = errors.Wrapf(errors.InvalidType, " uint8 is Invalid. Expected %s", reflect.TypeOf(retValue))
        return  
    }

    value = uint8(valuetmp)

    return
}

// SetWilling sets the value of Willing for the instance
func (instance *MSFT_NetQosDcbxSettingData) SetPropertyWilling(value bool) (err error) { 
    return instance.SetProperty("Willing", (value))
}

// GetWilling gets the value of Willing for the instance
func (instance *MSFT_NetQosDcbxSettingData) GetPropertyWilling()(value bool, err error) { 
    retValue, err := instance.GetProperty("Willing")
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

// 

// <param name="InterfaceAlias" type="string "></param>
// <param name="InterfaceIndex" type="uint32 "></param>
// <param name="PolicySet" type="uint8 "></param>

// <param name="ReturnValue" type="uint32 "></param>
func (instance *MSFT_NetQosDcbxSettingData) SwitchPolicySet( /* IN */ PolicySet uint8,
 /* IN */ InterfaceAlias string,
 /* IN */ InterfaceIndex uint32) (result uint32, err error) {retVal, err := instance.InvokeMethodWithReturn("SwitchPolicySet" , PolicySet, InterfaceAlias, InterfaceIndex);
	if err != nil { return }
	result = uint32(retVal)
	return
	
}


