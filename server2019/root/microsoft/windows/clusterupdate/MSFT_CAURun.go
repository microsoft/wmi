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

// MSFT_CAURun struct
type MSFT_CAURun struct { 
	*cim.WmiInstance

	// 
	OrchestratorGuid string
}

	func NewMSFT_CAURunEx1(instance *cim.WmiInstance) (newInstance *MSFT_CAURun, err error) {tmp, err := instance, nil
		
	if err != nil { return }
	newInstance = &MSFT_CAURun {
	WmiInstance: tmp,
	}
	return
	}
	

	func NewMSFT_CAURunEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery ) (newInstance *MSFT_CAURun, err error) {tmp, err := instance.GetWmiInstance(hostName, wmiNamespace, userName, password, domainName, query)
		
	if err != nil { return }
	newInstance = &MSFT_CAURun {
	WmiInstance: tmp,
	}
	return
	}
	

// SetOrchestratorGuid sets the value of OrchestratorGuid for the instance
func (instance *MSFT_CAURun) SetPropertyOrchestratorGuid(value string) (err error) { 
    return instance.SetProperty("OrchestratorGuid", (value))
}

// GetOrchestratorGuid gets the value of OrchestratorGuid for the instance
func (instance *MSFT_CAURun) GetPropertyOrchestratorGuid()(value string, err error) { 
    retValue, err := instance.GetProperty("OrchestratorGuid")
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

