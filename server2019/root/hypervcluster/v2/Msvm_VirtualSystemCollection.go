// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 9/18/2020 using wmigen
//      Source root.HyperVCluster.v2
//////////////////////////////////////////////
package v2

import (
	"github.com/microsoft/wmi/pkg/base/query"
	"github.com/microsoft/wmi/pkg/errors"
	cim "github.com/microsoft/wmi/pkg/wmiinstance"
	"reflect"
)

// Msvm_VirtualSystemCollection struct
type Msvm_VirtualSystemCollection struct {
	*CIM_CollectionOfMSEs

	//
	FailedOverReplicationType uint16

	//
	LastApplyConsistencyLevel uint16

	//
	LastApplyTime string

	//
	LastApplyVirtualMachineIds []string

	//
	ReplicationMode uint16

	//
	ReplicationState uint16
}

func NewMsvm_VirtualSystemCollectionEx1(instance *cim.WmiInstance) (newInstance *Msvm_VirtualSystemCollection, err error) {
	tmp, err := NewCIM_CollectionOfMSEsEx1(instance)

	if err != nil {
		return
	}
	newInstance = &Msvm_VirtualSystemCollection{
		CIM_CollectionOfMSEs: tmp,
	}
	return
}

func NewMsvm_VirtualSystemCollectionEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *Msvm_VirtualSystemCollection, err error) {
	tmp, err := NewCIM_CollectionOfMSEsEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &Msvm_VirtualSystemCollection{
		CIM_CollectionOfMSEs: tmp,
	}
	return
}

// SetFailedOverReplicationType sets the value of FailedOverReplicationType for the instance
func (instance *Msvm_VirtualSystemCollection) SetPropertyFailedOverReplicationType(value uint16) (err error) {
	return instance.SetProperty("FailedOverReplicationType", (value))
}

// GetFailedOverReplicationType gets the value of FailedOverReplicationType for the instance
func (instance *Msvm_VirtualSystemCollection) GetPropertyFailedOverReplicationType() (value uint16, err error) {
	retValue, err := instance.GetProperty("FailedOverReplicationType")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(uint16)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " uint16 is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = uint16(valuetmp)

	return
}

// SetLastApplyConsistencyLevel sets the value of LastApplyConsistencyLevel for the instance
func (instance *Msvm_VirtualSystemCollection) SetPropertyLastApplyConsistencyLevel(value uint16) (err error) {
	return instance.SetProperty("LastApplyConsistencyLevel", (value))
}

// GetLastApplyConsistencyLevel gets the value of LastApplyConsistencyLevel for the instance
func (instance *Msvm_VirtualSystemCollection) GetPropertyLastApplyConsistencyLevel() (value uint16, err error) {
	retValue, err := instance.GetProperty("LastApplyConsistencyLevel")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(uint16)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " uint16 is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = uint16(valuetmp)

	return
}

// SetLastApplyTime sets the value of LastApplyTime for the instance
func (instance *Msvm_VirtualSystemCollection) SetPropertyLastApplyTime(value string) (err error) {
	return instance.SetProperty("LastApplyTime", (value))
}

// GetLastApplyTime gets the value of LastApplyTime for the instance
func (instance *Msvm_VirtualSystemCollection) GetPropertyLastApplyTime() (value string, err error) {
	retValue, err := instance.GetProperty("LastApplyTime")
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

// SetLastApplyVirtualMachineIds sets the value of LastApplyVirtualMachineIds for the instance
func (instance *Msvm_VirtualSystemCollection) SetPropertyLastApplyVirtualMachineIds(value []string) (err error) {
	return instance.SetProperty("LastApplyVirtualMachineIds", (value))
}

// GetLastApplyVirtualMachineIds gets the value of LastApplyVirtualMachineIds for the instance
func (instance *Msvm_VirtualSystemCollection) GetPropertyLastApplyVirtualMachineIds() (value []string, err error) {
	retValue, err := instance.GetProperty("LastApplyVirtualMachineIds")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	for _, interfaceValue := range retValue.([]interface{}) {
		valuetmp, ok := interfaceValue.(string)
		if !ok {
			err = errors.Wrapf(errors.InvalidType, " string is Invalid. Expected %s", reflect.TypeOf(interfaceValue))
			return
		}
		value = append(value, string(valuetmp))
	}

	return
}

// SetReplicationMode sets the value of ReplicationMode for the instance
func (instance *Msvm_VirtualSystemCollection) SetPropertyReplicationMode(value uint16) (err error) {
	return instance.SetProperty("ReplicationMode", (value))
}

// GetReplicationMode gets the value of ReplicationMode for the instance
func (instance *Msvm_VirtualSystemCollection) GetPropertyReplicationMode() (value uint16, err error) {
	retValue, err := instance.GetProperty("ReplicationMode")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(uint16)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " uint16 is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = uint16(valuetmp)

	return
}

// SetReplicationState sets the value of ReplicationState for the instance
func (instance *Msvm_VirtualSystemCollection) SetPropertyReplicationState(value uint16) (err error) {
	return instance.SetProperty("ReplicationState", (value))
}

// GetReplicationState gets the value of ReplicationState for the instance
func (instance *Msvm_VirtualSystemCollection) GetPropertyReplicationState() (value uint16, err error) {
	retValue, err := instance.GetProperty("ReplicationState")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(uint16)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " uint16 is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = uint16(valuetmp)

	return
}
