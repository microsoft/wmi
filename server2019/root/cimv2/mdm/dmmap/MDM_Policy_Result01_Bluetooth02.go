// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

// 
// Author:
//      Auto Generated on 6/6/2024 using wmigen
//      Source root.CIMV2.mdm.dmmap
//////////////////////////////////////////////
package dmmap
import (
 "github.com/microsoft/wmi/pkg/base/query"
cim "github.com/microsoft/wmi/pkg/wmiinstance"
 "github.com/microsoft/wmi/pkg/base/instance"
 "github.com/microsoft/wmi/pkg/errors"
 "reflect"
)

// MDM_Policy_Result01_Bluetooth02 struct
type MDM_Policy_Result01_Bluetooth02 struct { 
	*cim.WmiInstance

	// 
	AllowAdvertising int32

	// 
	AllowDiscoverableMode int32

	// 
	AllowPrepairing int32

	// 
	AllowPromptedProximalConnections int32

	// 
	InstanceID string

	// 
	LocalDeviceName string

	// 
	ParentID string

	// 
	ServicesAllowedList string

	// 
	SetMinimumEncryptionKeySize int32
}

	func NewMDM_Policy_Result01_Bluetooth02Ex1(instance *cim.WmiInstance) (newInstance *MDM_Policy_Result01_Bluetooth02, err error) {tmp, err := instance, nil
		
	if err != nil { return }
	newInstance = &MDM_Policy_Result01_Bluetooth02 {
	WmiInstance: tmp,
	}
	return
	}
	

	func NewMDM_Policy_Result01_Bluetooth02Ex6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery ) (newInstance *MDM_Policy_Result01_Bluetooth02, err error) {tmp, err := instance.GetWmiInstance(hostName, wmiNamespace, userName, password, domainName, query)
		
	if err != nil { return }
	newInstance = &MDM_Policy_Result01_Bluetooth02 {
	WmiInstance: tmp,
	}
	return
	}
	

// SetAllowAdvertising sets the value of AllowAdvertising for the instance
func (instance *MDM_Policy_Result01_Bluetooth02) SetPropertyAllowAdvertising(value int32) (err error) { 
    return instance.SetProperty("AllowAdvertising", (value))
}

// GetAllowAdvertising gets the value of AllowAdvertising for the instance
func (instance *MDM_Policy_Result01_Bluetooth02) GetPropertyAllowAdvertising()(value int32, err error) { 
    retValue, err := instance.GetProperty("AllowAdvertising")
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

// SetAllowDiscoverableMode sets the value of AllowDiscoverableMode for the instance
func (instance *MDM_Policy_Result01_Bluetooth02) SetPropertyAllowDiscoverableMode(value int32) (err error) { 
    return instance.SetProperty("AllowDiscoverableMode", (value))
}

// GetAllowDiscoverableMode gets the value of AllowDiscoverableMode for the instance
func (instance *MDM_Policy_Result01_Bluetooth02) GetPropertyAllowDiscoverableMode()(value int32, err error) { 
    retValue, err := instance.GetProperty("AllowDiscoverableMode")
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

// SetAllowPrepairing sets the value of AllowPrepairing for the instance
func (instance *MDM_Policy_Result01_Bluetooth02) SetPropertyAllowPrepairing(value int32) (err error) { 
    return instance.SetProperty("AllowPrepairing", (value))
}

// GetAllowPrepairing gets the value of AllowPrepairing for the instance
func (instance *MDM_Policy_Result01_Bluetooth02) GetPropertyAllowPrepairing()(value int32, err error) { 
    retValue, err := instance.GetProperty("AllowPrepairing")
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

// SetAllowPromptedProximalConnections sets the value of AllowPromptedProximalConnections for the instance
func (instance *MDM_Policy_Result01_Bluetooth02) SetPropertyAllowPromptedProximalConnections(value int32) (err error) { 
    return instance.SetProperty("AllowPromptedProximalConnections", (value))
}

// GetAllowPromptedProximalConnections gets the value of AllowPromptedProximalConnections for the instance
func (instance *MDM_Policy_Result01_Bluetooth02) GetPropertyAllowPromptedProximalConnections()(value int32, err error) { 
    retValue, err := instance.GetProperty("AllowPromptedProximalConnections")
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

// SetInstanceID sets the value of InstanceID for the instance
func (instance *MDM_Policy_Result01_Bluetooth02) SetPropertyInstanceID(value string) (err error) { 
    return instance.SetProperty("InstanceID", (value))
}

// GetInstanceID gets the value of InstanceID for the instance
func (instance *MDM_Policy_Result01_Bluetooth02) GetPropertyInstanceID()(value string, err error) { 
    retValue, err := instance.GetProperty("InstanceID")
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

// SetLocalDeviceName sets the value of LocalDeviceName for the instance
func (instance *MDM_Policy_Result01_Bluetooth02) SetPropertyLocalDeviceName(value string) (err error) { 
    return instance.SetProperty("LocalDeviceName", (value))
}

// GetLocalDeviceName gets the value of LocalDeviceName for the instance
func (instance *MDM_Policy_Result01_Bluetooth02) GetPropertyLocalDeviceName()(value string, err error) { 
    retValue, err := instance.GetProperty("LocalDeviceName")
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

// SetParentID sets the value of ParentID for the instance
func (instance *MDM_Policy_Result01_Bluetooth02) SetPropertyParentID(value string) (err error) { 
    return instance.SetProperty("ParentID", (value))
}

// GetParentID gets the value of ParentID for the instance
func (instance *MDM_Policy_Result01_Bluetooth02) GetPropertyParentID()(value string, err error) { 
    retValue, err := instance.GetProperty("ParentID")
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

// SetServicesAllowedList sets the value of ServicesAllowedList for the instance
func (instance *MDM_Policy_Result01_Bluetooth02) SetPropertyServicesAllowedList(value string) (err error) { 
    return instance.SetProperty("ServicesAllowedList", (value))
}

// GetServicesAllowedList gets the value of ServicesAllowedList for the instance
func (instance *MDM_Policy_Result01_Bluetooth02) GetPropertyServicesAllowedList()(value string, err error) { 
    retValue, err := instance.GetProperty("ServicesAllowedList")
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

// SetSetMinimumEncryptionKeySize sets the value of SetMinimumEncryptionKeySize for the instance
func (instance *MDM_Policy_Result01_Bluetooth02) SetPropertySetMinimumEncryptionKeySize(value int32) (err error) { 
    return instance.SetProperty("SetMinimumEncryptionKeySize", (value))
}

// GetSetMinimumEncryptionKeySize gets the value of SetMinimumEncryptionKeySize for the instance
func (instance *MDM_Policy_Result01_Bluetooth02) GetPropertySetMinimumEncryptionKeySize()(value int32, err error) { 
    retValue, err := instance.GetProperty("SetMinimumEncryptionKeySize")
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

