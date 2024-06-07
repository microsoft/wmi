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

// BcdDeviceLocateStringData struct
type BcdDeviceLocateStringData struct { 
	*BcdDeviceLocateData

	// This provides the locate device element.
	Path string
}

	func NewBcdDeviceLocateStringDataEx1(instance *cim.WmiInstance) (newInstance *BcdDeviceLocateStringData, err error) {tmp, err := NewBcdDeviceLocateDataEx1(instance)
		
	if err != nil { return }
	newInstance = &BcdDeviceLocateStringData {
	BcdDeviceLocateData: tmp,
	}
	return
	}
	

	func NewBcdDeviceLocateStringDataEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery ) (newInstance *BcdDeviceLocateStringData, err error) {tmp, err := NewBcdDeviceLocateDataEx6(hostName, wmiNamespace, userName, password, domainName, query)
		
	if err != nil { return }
	newInstance = &BcdDeviceLocateStringData {
	BcdDeviceLocateData: tmp,
	}
	return
	}
	

// SetPath sets the value of Path for the instance
func (instance *BcdDeviceLocateStringData) SetPropertyPath(value string) (err error) { 
    return instance.SetProperty("Path", (value))
}

// GetPath gets the value of Path for the instance
func (instance *BcdDeviceLocateStringData) GetPropertyPath()(value string, err error) { 
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

