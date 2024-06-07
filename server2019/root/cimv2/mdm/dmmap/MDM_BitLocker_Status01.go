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

// MDM_BitLocker_Status01 struct
type MDM_BitLocker_Status01 struct { 
	*cim.WmiInstance

	// 
	DeviceEncryptionStatus int32

	// 
	InstanceID string

	// 
	ParentID string

	// 
	RemovableDrivesEncryptionStatus int32

	// 
	RotateRecoveryPasswordsRequestID string

	// 
	RotateRecoveryPasswordsStatus int32
}

	func NewMDM_BitLocker_Status01Ex1(instance *cim.WmiInstance) (newInstance *MDM_BitLocker_Status01, err error) {tmp, err := instance, nil
		
	if err != nil { return }
	newInstance = &MDM_BitLocker_Status01 {
	WmiInstance: tmp,
	}
	return
	}
	

	func NewMDM_BitLocker_Status01Ex6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery ) (newInstance *MDM_BitLocker_Status01, err error) {tmp, err := instance.GetWmiInstance(hostName, wmiNamespace, userName, password, domainName, query)
		
	if err != nil { return }
	newInstance = &MDM_BitLocker_Status01 {
	WmiInstance: tmp,
	}
	return
	}
	

// SetDeviceEncryptionStatus sets the value of DeviceEncryptionStatus for the instance
func (instance *MDM_BitLocker_Status01) SetPropertyDeviceEncryptionStatus(value int32) (err error) { 
    return instance.SetProperty("DeviceEncryptionStatus", (value))
}

// GetDeviceEncryptionStatus gets the value of DeviceEncryptionStatus for the instance
func (instance *MDM_BitLocker_Status01) GetPropertyDeviceEncryptionStatus()(value int32, err error) { 
    retValue, err := instance.GetProperty("DeviceEncryptionStatus")
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
func (instance *MDM_BitLocker_Status01) SetPropertyInstanceID(value string) (err error) { 
    return instance.SetProperty("InstanceID", (value))
}

// GetInstanceID gets the value of InstanceID for the instance
func (instance *MDM_BitLocker_Status01) GetPropertyInstanceID()(value string, err error) { 
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

// SetParentID sets the value of ParentID for the instance
func (instance *MDM_BitLocker_Status01) SetPropertyParentID(value string) (err error) { 
    return instance.SetProperty("ParentID", (value))
}

// GetParentID gets the value of ParentID for the instance
func (instance *MDM_BitLocker_Status01) GetPropertyParentID()(value string, err error) { 
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

// SetRemovableDrivesEncryptionStatus sets the value of RemovableDrivesEncryptionStatus for the instance
func (instance *MDM_BitLocker_Status01) SetPropertyRemovableDrivesEncryptionStatus(value int32) (err error) { 
    return instance.SetProperty("RemovableDrivesEncryptionStatus", (value))
}

// GetRemovableDrivesEncryptionStatus gets the value of RemovableDrivesEncryptionStatus for the instance
func (instance *MDM_BitLocker_Status01) GetPropertyRemovableDrivesEncryptionStatus()(value int32, err error) { 
    retValue, err := instance.GetProperty("RemovableDrivesEncryptionStatus")
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

// SetRotateRecoveryPasswordsRequestID sets the value of RotateRecoveryPasswordsRequestID for the instance
func (instance *MDM_BitLocker_Status01) SetPropertyRotateRecoveryPasswordsRequestID(value string) (err error) { 
    return instance.SetProperty("RotateRecoveryPasswordsRequestID", (value))
}

// GetRotateRecoveryPasswordsRequestID gets the value of RotateRecoveryPasswordsRequestID for the instance
func (instance *MDM_BitLocker_Status01) GetPropertyRotateRecoveryPasswordsRequestID()(value string, err error) { 
    retValue, err := instance.GetProperty("RotateRecoveryPasswordsRequestID")
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

// SetRotateRecoveryPasswordsStatus sets the value of RotateRecoveryPasswordsStatus for the instance
func (instance *MDM_BitLocker_Status01) SetPropertyRotateRecoveryPasswordsStatus(value int32) (err error) { 
    return instance.SetProperty("RotateRecoveryPasswordsStatus", (value))
}

// GetRotateRecoveryPasswordsStatus gets the value of RotateRecoveryPasswordsStatus for the instance
func (instance *MDM_BitLocker_Status01) GetPropertyRotateRecoveryPasswordsStatus()(value int32, err error) { 
    retValue, err := instance.GetProperty("RotateRecoveryPasswordsStatus")
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

