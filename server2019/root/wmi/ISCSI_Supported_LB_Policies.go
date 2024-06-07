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

// ISCSI_Supported_LB_Policies struct
type ISCSI_Supported_LB_Policies struct { 
	*cim.WmiInstance

	// Number of entries in MSiSCSI_Paths array
	iSCSI_PathCount uint32

	// Describes iSCSI Initiator Paths
	iSCSI_Paths []ISCSI_Path

	// Load Balance policy supported by the iSCSI Initiator
	LoadBalancePolicy Policies_LoadBalancePolicy

	// Id that is unique to this session within this adapter. 
	UniqueSessionId uint64
}

	func NewISCSI_Supported_LB_PoliciesEx1(instance *cim.WmiInstance) (newInstance *ISCSI_Supported_LB_Policies, err error) {tmp, err := instance, nil
		
	if err != nil { return }
	newInstance = &ISCSI_Supported_LB_Policies {
	WmiInstance: tmp,
	}
	return
	}
	

	func NewISCSI_Supported_LB_PoliciesEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery ) (newInstance *ISCSI_Supported_LB_Policies, err error) {tmp, err := instance.GetWmiInstance(hostName, wmiNamespace, userName, password, domainName, query)
		
	if err != nil { return }
	newInstance = &ISCSI_Supported_LB_Policies {
	WmiInstance: tmp,
	}
	return
	}
	

// SetiSCSI_PathCount sets the value of iSCSI_PathCount for the instance
func (instance *ISCSI_Supported_LB_Policies) SetPropertyiSCSI_PathCount(value uint32) (err error) { 
    return instance.SetProperty("iSCSI_PathCount", (value))
}

// GetiSCSI_PathCount gets the value of iSCSI_PathCount for the instance
func (instance *ISCSI_Supported_LB_Policies) GetPropertyiSCSI_PathCount()(value uint32, err error) { 
    retValue, err := instance.GetProperty("iSCSI_PathCount")
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

// SetiSCSI_Paths sets the value of iSCSI_Paths for the instance
func (instance *ISCSI_Supported_LB_Policies) SetPropertyiSCSI_Paths(value []ISCSI_Path) (err error) { 
    return instance.SetProperty("iSCSI_Paths", (value))
}

// GetiSCSI_Paths gets the value of iSCSI_Paths for the instance
func (instance *ISCSI_Supported_LB_Policies) GetPropertyiSCSI_Paths()(value []ISCSI_Path, err error) { 
    retValue, err := instance.GetProperty("iSCSI_Paths")
    if err != nil {
        return
    }
    if retValue == nil {
        // Doesn't have any value. Return empty
        return
    }
    
    for _, interfaceValue := range retValue.([]interface{}) {
        valuetmp, ok := interfaceValue.(ISCSI_Path); 
        if !ok {
            err = errors.Wrapf(errors.InvalidType, " ISCSI_Path is Invalid. Expected %s", reflect.TypeOf(interfaceValue))
            return  
        }
        value = append(value, ISCSI_Path(valuetmp))
    }

    return
}

// SetLoadBalancePolicy sets the value of LoadBalancePolicy for the instance
func (instance *ISCSI_Supported_LB_Policies) SetPropertyLoadBalancePolicy(value Policies_LoadBalancePolicy) (err error) { 
    return instance.SetProperty("LoadBalancePolicy", (value))
}

// GetLoadBalancePolicy gets the value of LoadBalancePolicy for the instance
func (instance *ISCSI_Supported_LB_Policies) GetPropertyLoadBalancePolicy()(value Policies_LoadBalancePolicy, err error) { 
    retValue, err := instance.GetProperty("LoadBalancePolicy")
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

    value = Policies_LoadBalancePolicy(valuetmp)

    return
}

// SetUniqueSessionId sets the value of UniqueSessionId for the instance
func (instance *ISCSI_Supported_LB_Policies) SetPropertyUniqueSessionId(value uint64) (err error) { 
    return instance.SetProperty("UniqueSessionId", (value))
}

// GetUniqueSessionId gets the value of UniqueSessionId for the instance
func (instance *ISCSI_Supported_LB_Policies) GetPropertyUniqueSessionId()(value uint64, err error) { 
    retValue, err := instance.GetProperty("UniqueSessionId")
    if err != nil {
        return
    }
    if retValue == nil {
        // Doesn't have any value. Return empty
        return
    }
    
    valuetmp, ok := retValue.(uint64); 
    if !ok {
        err = errors.Wrapf(errors.InvalidType, " uint64 is Invalid. Expected %s", reflect.TypeOf(retValue))
        return  
    }

    value = uint64(valuetmp)

    return
}

