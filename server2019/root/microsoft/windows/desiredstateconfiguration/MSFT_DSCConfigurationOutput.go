// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

// 
// Author:
//      Auto Generated on 6/6/2024 using wmigen
//      Source root.Microsoft.Windows.DesiredStateConfiguration
//////////////////////////////////////////////
package desiredstateconfiguration
import (
 "github.com/microsoft/wmi/pkg/base/query"
cim "github.com/microsoft/wmi/pkg/wmiinstance"
 "github.com/microsoft/wmi/pkg/base/instance"
 "github.com/microsoft/wmi/pkg/errors"
 "reflect"
)

// MSFT_DSCConfigurationOutput struct
type MSFT_DSCConfigurationOutput struct { 
	*cim.WmiInstance

	// 
	Bookmark []uint8

	// 
	JobId string
}

	func NewMSFT_DSCConfigurationOutputEx1(instance *cim.WmiInstance) (newInstance *MSFT_DSCConfigurationOutput, err error) {tmp, err := instance, nil
		
	if err != nil { return }
	newInstance = &MSFT_DSCConfigurationOutput {
	WmiInstance: tmp,
	}
	return
	}
	

	func NewMSFT_DSCConfigurationOutputEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery ) (newInstance *MSFT_DSCConfigurationOutput, err error) {tmp, err := instance.GetWmiInstance(hostName, wmiNamespace, userName, password, domainName, query)
		
	if err != nil { return }
	newInstance = &MSFT_DSCConfigurationOutput {
	WmiInstance: tmp,
	}
	return
	}
	

// SetBookmark sets the value of Bookmark for the instance
func (instance *MSFT_DSCConfigurationOutput) SetPropertyBookmark(value []uint8) (err error) { 
    return instance.SetProperty("Bookmark", (value))
}

// GetBookmark gets the value of Bookmark for the instance
func (instance *MSFT_DSCConfigurationOutput) GetPropertyBookmark()(value []uint8, err error) { 
    retValue, err := instance.GetProperty("Bookmark")
    if err != nil {
        return
    }
    if retValue == nil {
        // Doesn't have any value. Return empty
        return
    }
    
    for _, interfaceValue := range retValue.([]interface{}) {
        valuetmp, ok := interfaceValue.(uint8); 
        if !ok {
            err = errors.Wrapf(errors.InvalidType, " uint8 is Invalid. Expected %s", reflect.TypeOf(interfaceValue))
            return  
        }
        value = append(value, uint8(valuetmp))
    }

    return
}

// SetJobId sets the value of JobId for the instance
func (instance *MSFT_DSCConfigurationOutput) SetPropertyJobId(value string) (err error) { 
    return instance.SetProperty("JobId", (value))
}

// GetJobId gets the value of JobId for the instance
func (instance *MSFT_DSCConfigurationOutput) GetPropertyJobId()(value string, err error) { 
    retValue, err := instance.GetProperty("JobId")
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

