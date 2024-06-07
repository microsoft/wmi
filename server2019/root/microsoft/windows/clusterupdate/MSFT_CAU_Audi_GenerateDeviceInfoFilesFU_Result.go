// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

// 
// Author:
//      Auto Generated on 6/6/2024 using wmigen
//      Source root.Microsoft.Windows.ClusterUpdate
//////////////////////////////////////////////
package clusterupdate
import (
 "github.com/microsoft/wmi/pkg/base/query"
cim "github.com/microsoft/wmi/pkg/wmiinstance"
 "github.com/microsoft/wmi/pkg/base/instance"
 "github.com/microsoft/wmi/pkg/errors"
 "reflect"
)

// MSFT_CAU_Audi_GenerateDeviceInfoFilesFU_Result struct
type MSFT_CAU_Audi_GenerateDeviceInfoFilesFU_Result struct { 
	*cim.WmiInstance

	// 
	GenerateFilesHResult uint32
}

	func NewMSFT_CAU_Audi_GenerateDeviceInfoFilesFU_ResultEx1(instance *cim.WmiInstance) (newInstance *MSFT_CAU_Audi_GenerateDeviceInfoFilesFU_Result, err error) {tmp, err := instance, nil
		
	if err != nil { return }
	newInstance = &MSFT_CAU_Audi_GenerateDeviceInfoFilesFU_Result {
	WmiInstance: tmp,
	}
	return
	}
	

	func NewMSFT_CAU_Audi_GenerateDeviceInfoFilesFU_ResultEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery ) (newInstance *MSFT_CAU_Audi_GenerateDeviceInfoFilesFU_Result, err error) {tmp, err := instance.GetWmiInstance(hostName, wmiNamespace, userName, password, domainName, query)
		
	if err != nil { return }
	newInstance = &MSFT_CAU_Audi_GenerateDeviceInfoFilesFU_Result {
	WmiInstance: tmp,
	}
	return
	}
	

// SetGenerateFilesHResult sets the value of GenerateFilesHResult for the instance
func (instance *MSFT_CAU_Audi_GenerateDeviceInfoFilesFU_Result) SetPropertyGenerateFilesHResult(value uint32) (err error) { 
    return instance.SetProperty("GenerateFilesHResult", (value))
}

// GetGenerateFilesHResult gets the value of GenerateFilesHResult for the instance
func (instance *MSFT_CAU_Audi_GenerateDeviceInfoFilesFU_Result) GetPropertyGenerateFilesHResult()(value uint32, err error) { 
    retValue, err := instance.GetProperty("GenerateFilesHResult")
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

