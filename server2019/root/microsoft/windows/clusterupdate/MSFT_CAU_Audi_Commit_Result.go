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

// MSFT_CAU_Audi_Commit_Result struct
type MSFT_CAU_Audi_Commit_Result struct { 
	*cim.WmiInstance

	// 
	AudiCommitHResult uint32
}

	func NewMSFT_CAU_Audi_Commit_ResultEx1(instance *cim.WmiInstance) (newInstance *MSFT_CAU_Audi_Commit_Result, err error) {tmp, err := instance, nil
		
	if err != nil { return }
	newInstance = &MSFT_CAU_Audi_Commit_Result {
	WmiInstance: tmp,
	}
	return
	}
	

	func NewMSFT_CAU_Audi_Commit_ResultEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery ) (newInstance *MSFT_CAU_Audi_Commit_Result, err error) {tmp, err := instance.GetWmiInstance(hostName, wmiNamespace, userName, password, domainName, query)
		
	if err != nil { return }
	newInstance = &MSFT_CAU_Audi_Commit_Result {
	WmiInstance: tmp,
	}
	return
	}
	

// SetAudiCommitHResult sets the value of AudiCommitHResult for the instance
func (instance *MSFT_CAU_Audi_Commit_Result) SetPropertyAudiCommitHResult(value uint32) (err error) { 
    return instance.SetProperty("AudiCommitHResult", (value))
}

// GetAudiCommitHResult gets the value of AudiCommitHResult for the instance
func (instance *MSFT_CAU_Audi_Commit_Result) GetPropertyAudiCommitHResult()(value uint32, err error) { 
    retValue, err := instance.GetProperty("AudiCommitHResult")
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

