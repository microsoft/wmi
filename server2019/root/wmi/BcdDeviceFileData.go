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

// BcdDeviceFileData struct
type BcdDeviceFileData struct { 
	*BcdDeviceData

	// This is the parent device of this file device.
	Parent BcdDeviceData

	// This is the device path.
	Path string
}

	func NewBcdDeviceFileDataEx1(instance *cim.WmiInstance) (newInstance *BcdDeviceFileData, err error) {tmp, err := NewBcdDeviceDataEx1(instance)
		
	if err != nil { return }
	newInstance = &BcdDeviceFileData {
	BcdDeviceData: tmp,
	}
	return
	}
	

	func NewBcdDeviceFileDataEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery ) (newInstance *BcdDeviceFileData, err error) {tmp, err := NewBcdDeviceDataEx6(hostName, wmiNamespace, userName, password, domainName, query)
		
	if err != nil { return }
	newInstance = &BcdDeviceFileData {
	BcdDeviceData: tmp,
	}
	return
	}
	

// SetParent sets the value of Parent for the instance
func (instance *BcdDeviceFileData) SetPropertyParent(value BcdDeviceData) (err error) { 
    return instance.SetProperty("Parent", (value))
}

// GetParent gets the value of Parent for the instance
func (instance *BcdDeviceFileData) GetPropertyParent()(value BcdDeviceData, err error) { 
    retValue, err := instance.GetProperty("Parent")
    if err != nil {
        return
    }
    if retValue == nil {
        // Doesn't have any value. Return empty
        return
    }
    
    valuetmp, ok := retValue.(BcdDeviceData); 
    if !ok {
        err = errors.Wrapf(errors.InvalidType, " BcdDeviceData is Invalid. Expected %s", reflect.TypeOf(retValue))
        return  
    }

    value = BcdDeviceData(valuetmp)

    return
}

// SetPath sets the value of Path for the instance
func (instance *BcdDeviceFileData) SetPropertyPath(value string) (err error) { 
    return instance.SetProperty("Path", (value))
}

// GetPath gets the value of Path for the instance
func (instance *BcdDeviceFileData) GetPropertyPath()(value string, err error) { 
    retValue, err := instance.GetProperty("Path")
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

