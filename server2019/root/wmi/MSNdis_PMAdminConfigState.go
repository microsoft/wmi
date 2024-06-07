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

// MSNdis_PMAdminConfigState struct
type MSNdis_PMAdminConfigState struct { 
	*MSNdis

	// 
	NdisPMAdminConfigState PMAdminConfigState_NdisPMAdminConfigState
}

	func NewMSNdis_PMAdminConfigStateEx1(instance *cim.WmiInstance) (newInstance *MSNdis_PMAdminConfigState, err error) {tmp, err := NewMSNdisEx1(instance)
		
	if err != nil { return }
	newInstance = &MSNdis_PMAdminConfigState {
	MSNdis: tmp,
	}
	return
	}
	

	func NewMSNdis_PMAdminConfigStateEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery ) (newInstance *MSNdis_PMAdminConfigState, err error) {tmp, err := NewMSNdisEx6(hostName, wmiNamespace, userName, password, domainName, query)
		
	if err != nil { return }
	newInstance = &MSNdis_PMAdminConfigState {
	MSNdis: tmp,
	}
	return
	}
	

// SetNdisPMAdminConfigState sets the value of NdisPMAdminConfigState for the instance
func (instance *MSNdis_PMAdminConfigState) SetPropertyNdisPMAdminConfigState(value PMAdminConfigState_NdisPMAdminConfigState) (err error) { 
    return instance.SetProperty("NdisPMAdminConfigState", (value))
}

// GetNdisPMAdminConfigState gets the value of NdisPMAdminConfigState for the instance
func (instance *MSNdis_PMAdminConfigState) GetPropertyNdisPMAdminConfigState()(value PMAdminConfigState_NdisPMAdminConfigState, err error) { 
    retValue, err := instance.GetProperty("NdisPMAdminConfigState")
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

    value = PMAdminConfigState_NdisPMAdminConfigState(valuetmp)

    return
}

