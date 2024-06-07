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

// MSiSCSI_LB_Operations struct
type MSiSCSI_LB_Operations struct { 
	*cim.WmiInstance

	// 
	Active bool

	// 
	InstanceName string
}

	func NewMSiSCSI_LB_OperationsEx1(instance *cim.WmiInstance) (newInstance *MSiSCSI_LB_Operations, err error) {tmp, err := instance, nil
		
	if err != nil { return }
	newInstance = &MSiSCSI_LB_Operations {
	WmiInstance: tmp,
	}
	return
	}
	

	func NewMSiSCSI_LB_OperationsEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery ) (newInstance *MSiSCSI_LB_Operations, err error) {tmp, err := instance.GetWmiInstance(hostName, wmiNamespace, userName, password, domainName, query)
		
	if err != nil { return }
	newInstance = &MSiSCSI_LB_Operations {
	WmiInstance: tmp,
	}
	return
	}
	

// SetActive sets the value of Active for the instance
func (instance *MSiSCSI_LB_Operations) SetPropertyActive(value bool) (err error) { 
    return instance.SetProperty("Active", (value))
}

// GetActive gets the value of Active for the instance
func (instance *MSiSCSI_LB_Operations) GetPropertyActive()(value bool, err error) { 
    retValue, err := instance.GetProperty("Active")
    if err != nil {
        return
    }
    if retValue == nil {
        // Doesn't have any value. Return empty
        return
    }
    
    valuetmp, ok := retValue.(bool); 
    if !ok {
        err = errors.Wrapf(errors.InvalidType, " bool is Invalid. Expected %s", reflect.TypeOf(retValue))
        return  
    }

    value = bool(valuetmp)

    return
}

// SetInstanceName sets the value of InstanceName for the instance
func (instance *MSiSCSI_LB_Operations) SetPropertyInstanceName(value string) (err error) { 
    return instance.SetProperty("InstanceName", (value))
}

// GetInstanceName gets the value of InstanceName for the instance
func (instance *MSiSCSI_LB_Operations) GetPropertyInstanceName()(value string, err error) { 
    retValue, err := instance.GetProperty("InstanceName")
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

// Sets Load Balance Policy for the iSCSI Initiator

// <param name="LoadBalancePolicies" type="ISCSI_Supported_LB_Policies ">New Load Balance policy to be set</param>

// <param name="Status" type="uint32 ">Status of the operation</param>
func (instance *MSiSCSI_LB_Operations) SetLoadBalancePolicy( /* IN */ LoadBalancePolicies ISCSI_Supported_LB_Policies,
 /* OUT */ Status uint32) (err error) {_, err = instance.InvokeMethod("SetLoadBalancePolicy" , LoadBalancePolicies)
	if err != nil { return }
	return
	
}


