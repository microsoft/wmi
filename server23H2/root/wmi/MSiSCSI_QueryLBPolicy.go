// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 6/12/2024 using wmigen
//      Source root.WMI
//////////////////////////////////////////////
package wmi

import (
	"github.com/microsoft/wmi/pkg/base/instance"
	"github.com/microsoft/wmi/pkg/base/query"
	"github.com/microsoft/wmi/pkg/errors"
	cim "github.com/microsoft/wmi/pkg/wmiinstance"
	"reflect"
)

// MSiSCSI_QueryLBPolicy struct
type MSiSCSI_QueryLBPolicy struct {
	*cim.WmiInstance

	//
	Active bool

	//
	InstanceName string

	// Load Balance Policy that is currently being used by iSCSI Initiator - one element for each session on the adapter
	LoadBalancePolicies []ISCSI_Supported_LB_Policies

	//
	Reserved uint32

	// Number of elements in LoadBalancePolicies array
	SessionCount uint32

	// Id that is globally unique to each instance of each adapter. Using the address of the Adapter Extension is a good idea.
	UniqueAdapterId uint64
}

func NewMSiSCSI_QueryLBPolicyEx1(instance *cim.WmiInstance) (newInstance *MSiSCSI_QueryLBPolicy, err error) {
	tmp, err := instance, nil

	if err != nil {
		return
	}
	newInstance = &MSiSCSI_QueryLBPolicy{
		WmiInstance: tmp,
	}
	return
}

func NewMSiSCSI_QueryLBPolicyEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *MSiSCSI_QueryLBPolicy, err error) {
	tmp, err := instance.GetWmiInstance(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &MSiSCSI_QueryLBPolicy{
		WmiInstance: tmp,
	}
	return
}

// SetActive sets the value of Active for the instance
func (instance *MSiSCSI_QueryLBPolicy) SetPropertyActive(value bool) (err error) {
	return instance.SetProperty("Active", (value))
}

// GetActive gets the value of Active for the instance
func (instance *MSiSCSI_QueryLBPolicy) GetPropertyActive() (value bool, err error) {
	retValue, err := instance.GetProperty("Active")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(bool)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " bool is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = bool(valuetmp)

	return
}

// SetInstanceName sets the value of InstanceName for the instance
func (instance *MSiSCSI_QueryLBPolicy) SetPropertyInstanceName(value string) (err error) {
	return instance.SetProperty("InstanceName", (value))
}

// GetInstanceName gets the value of InstanceName for the instance
func (instance *MSiSCSI_QueryLBPolicy) GetPropertyInstanceName() (value string, err error) {
	retValue, err := instance.GetProperty("InstanceName")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(string)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " string is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = string(valuetmp)

	return
}

// SetLoadBalancePolicies sets the value of LoadBalancePolicies for the instance
func (instance *MSiSCSI_QueryLBPolicy) SetPropertyLoadBalancePolicies(value []ISCSI_Supported_LB_Policies) (err error) {
	return instance.SetProperty("LoadBalancePolicies", (value))
}

// GetLoadBalancePolicies gets the value of LoadBalancePolicies for the instance
func (instance *MSiSCSI_QueryLBPolicy) GetPropertyLoadBalancePolicies() (value []ISCSI_Supported_LB_Policies, err error) {
	retValue, err := instance.GetProperty("LoadBalancePolicies")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	for _, interfaceValue := range retValue.([]interface{}) {
		valuetmp, ok := interfaceValue.(ISCSI_Supported_LB_Policies)
		if !ok {
			err = errors.Wrapf(errors.InvalidType, " ISCSI_Supported_LB_Policies is Invalid. Expected %s", reflect.TypeOf(interfaceValue))
			return
		}
		value = append(value, ISCSI_Supported_LB_Policies(valuetmp))
	}

	return
}

// SetReserved sets the value of Reserved for the instance
func (instance *MSiSCSI_QueryLBPolicy) SetPropertyReserved(value uint32) (err error) {
	return instance.SetProperty("Reserved", (value))
}

// GetReserved gets the value of Reserved for the instance
func (instance *MSiSCSI_QueryLBPolicy) GetPropertyReserved() (value uint32, err error) {
	retValue, err := instance.GetProperty("Reserved")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(uint32)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " uint32 is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = uint32(valuetmp)

	return
}

// SetSessionCount sets the value of SessionCount for the instance
func (instance *MSiSCSI_QueryLBPolicy) SetPropertySessionCount(value uint32) (err error) {
	return instance.SetProperty("SessionCount", (value))
}

// GetSessionCount gets the value of SessionCount for the instance
func (instance *MSiSCSI_QueryLBPolicy) GetPropertySessionCount() (value uint32, err error) {
	retValue, err := instance.GetProperty("SessionCount")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(uint32)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " uint32 is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = uint32(valuetmp)

	return
}

// SetUniqueAdapterId sets the value of UniqueAdapterId for the instance
func (instance *MSiSCSI_QueryLBPolicy) SetPropertyUniqueAdapterId(value uint64) (err error) {
	return instance.SetProperty("UniqueAdapterId", (value))
}

// GetUniqueAdapterId gets the value of UniqueAdapterId for the instance
func (instance *MSiSCSI_QueryLBPolicy) GetPropertyUniqueAdapterId() (value uint64, err error) {
	retValue, err := instance.GetProperty("UniqueAdapterId")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(uint64)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " uint64 is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = uint64(valuetmp)

	return
}
