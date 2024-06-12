// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 6/12/2024 using wmigen
//      Source root.HyperVCluster.v2
//////////////////////////////////////////////
package v2

import (
	"github.com/microsoft/wmi/pkg/base/query"
	"github.com/microsoft/wmi/pkg/errors"
	cim "github.com/microsoft/wmi/pkg/wmiinstance"
	"reflect"
)

// Msvm_CollectionRecoveryPoint struct
type Msvm_CollectionRecoveryPoint struct {
	*CIM_ManagedElement

	//
	CollectionID string

	//
	ConsistencyLevel uint16

	//
	CreationTime string

	//
	RecoveryPointIds []string

	//
	VirtualMachineIds []string
}

func NewMsvm_CollectionRecoveryPointEx1(instance *cim.WmiInstance) (newInstance *Msvm_CollectionRecoveryPoint, err error) {
	tmp, err := NewCIM_ManagedElementEx1(instance)

	if err != nil {
		return
	}
	newInstance = &Msvm_CollectionRecoveryPoint{
		CIM_ManagedElement: tmp,
	}
	return
}

func NewMsvm_CollectionRecoveryPointEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *Msvm_CollectionRecoveryPoint, err error) {
	tmp, err := NewCIM_ManagedElementEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &Msvm_CollectionRecoveryPoint{
		CIM_ManagedElement: tmp,
	}
	return
}

// SetCollectionID sets the value of CollectionID for the instance
func (instance *Msvm_CollectionRecoveryPoint) SetPropertyCollectionID(value string) (err error) {
	return instance.SetProperty("CollectionID", (value))
}

// GetCollectionID gets the value of CollectionID for the instance
func (instance *Msvm_CollectionRecoveryPoint) GetPropertyCollectionID() (value string, err error) {
	retValue, err := instance.GetProperty("CollectionID")
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

// SetConsistencyLevel sets the value of ConsistencyLevel for the instance
func (instance *Msvm_CollectionRecoveryPoint) SetPropertyConsistencyLevel(value uint16) (err error) {
	return instance.SetProperty("ConsistencyLevel", (value))
}

// GetConsistencyLevel gets the value of ConsistencyLevel for the instance
func (instance *Msvm_CollectionRecoveryPoint) GetPropertyConsistencyLevel() (value uint16, err error) {
	retValue, err := instance.GetProperty("ConsistencyLevel")
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

// SetCreationTime sets the value of CreationTime for the instance
func (instance *Msvm_CollectionRecoveryPoint) SetPropertyCreationTime(value string) (err error) {
	return instance.SetProperty("CreationTime", (value))
}

// GetCreationTime gets the value of CreationTime for the instance
func (instance *Msvm_CollectionRecoveryPoint) GetPropertyCreationTime() (value string, err error) {
	retValue, err := instance.GetProperty("CreationTime")
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

// SetRecoveryPointIds sets the value of RecoveryPointIds for the instance
func (instance *Msvm_CollectionRecoveryPoint) SetPropertyRecoveryPointIds(value []string) (err error) {
	return instance.SetProperty("RecoveryPointIds", (value))
}

// GetRecoveryPointIds gets the value of RecoveryPointIds for the instance
func (instance *Msvm_CollectionRecoveryPoint) GetPropertyRecoveryPointIds() (value []string, err error) {
	retValue, err := instance.GetProperty("RecoveryPointIds")
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

// SetVirtualMachineIds sets the value of VirtualMachineIds for the instance
func (instance *Msvm_CollectionRecoveryPoint) SetPropertyVirtualMachineIds(value []string) (err error) {
	return instance.SetProperty("VirtualMachineIds", (value))
}

// GetVirtualMachineIds gets the value of VirtualMachineIds for the instance
func (instance *Msvm_CollectionRecoveryPoint) GetPropertyVirtualMachineIds() (value []string, err error) {
	retValue, err := instance.GetProperty("VirtualMachineIds")
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
