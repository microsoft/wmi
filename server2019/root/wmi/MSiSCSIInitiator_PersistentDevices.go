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
 "github.com/microsoft/wmi/pkg/base/instance"
 "github.com/microsoft/wmi/pkg/errors"
 "reflect"
)

// MSiSCSIInitiator_PersistentDevices struct
type MSiSCSIInitiator_PersistentDevices struct { 
	*cim.WmiInstance

	// 
	DevicePath string
}

	func NewMSiSCSIInitiator_PersistentDevicesEx1(instance *cim.WmiInstance) (newInstance *MSiSCSIInitiator_PersistentDevices, err error) {tmp, err := instance, nil
		
	if err != nil { return }
	newInstance = &MSiSCSIInitiator_PersistentDevices {
	WmiInstance: tmp,
	}
	return
	}
	

	func NewMSiSCSIInitiator_PersistentDevicesEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery ) (newInstance *MSiSCSIInitiator_PersistentDevices, err error) {tmp, err := instance.GetWmiInstance(hostName, wmiNamespace, userName, password, domainName, query)
		
	if err != nil { return }
	newInstance = &MSiSCSIInitiator_PersistentDevices {
	WmiInstance: tmp,
	}
	return
	}
	

// SetDevicePath sets the value of DevicePath for the instance
func (instance *MSiSCSIInitiator_PersistentDevices) SetPropertyDevicePath(value string) (err error) { 
    return instance.SetProperty("DevicePath", (value))
}

// GetDevicePath gets the value of DevicePath for the instance
func (instance *MSiSCSIInitiator_PersistentDevices) GetPropertyDevicePath()(value string, err error) { 
    retValue, err := instance.GetProperty("DevicePath")
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

