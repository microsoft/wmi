// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

// Author:
//
//	Auto Generated on 6/6/2024 using wmigen
//	Source root.HyperVCluster.v2
//
// ////////////////////////////////////////////
package v2

import (
	"github.com/microsoft/wmi/pkg/base/instance"
	"github.com/microsoft/wmi/pkg/base/query"
	"github.com/microsoft/wmi/pkg/errors"
	cim "github.com/microsoft/wmi/pkg/wmiinstance"
	"reflect"
)

// Msvm_SharedDiskSetReplicationRelationship struct
type Msvm_SharedDiskSetReplicationRelationship struct {
	*cim.WmiInstance

	//
	Caption string

	//
	CollectionID string

	//
	Description string

	//
	ElementName string

	//
	FailedOverReplicationType uint16

	//
	LastApplicationConsistentReplicationTime string

	//
	LastApplyTime string

	//
	LastReplicationTime string

	//
	LastReplicationType uint16

	//
	ReplicationHealth uint16

	//
	ReplicationState uint16
}

func NewMsvm_SharedDiskSetReplicationRelationshipEx1(instance *cim.WmiInstance) (newInstance *Msvm_SharedDiskSetReplicationRelationship, err error) {
	tmp, err := instance, nil

	if err != nil {
		return
	}
	newInstance = &Msvm_SharedDiskSetReplicationRelationship{
		WmiInstance: tmp,
	}
	return
}

func NewMsvm_SharedDiskSetReplicationRelationshipEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *Msvm_SharedDiskSetReplicationRelationship, err error) {
	tmp, err := instance.GetWmiInstance(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &Msvm_SharedDiskSetReplicationRelationship{
		WmiInstance: tmp,
	}
	return
}

// SetCaption sets the value of Caption for the instance
func (instance *Msvm_SharedDiskSetReplicationRelationship) SetPropertyCaption(value string) (err error) {
	return instance.SetProperty("Caption", (value))
}

// GetCaption gets the value of Caption for the instance
func (instance *Msvm_SharedDiskSetReplicationRelationship) GetPropertyCaption() (value string, err error) {
	retValue, err := instance.GetProperty("Caption")
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

// SetCollectionID sets the value of CollectionID for the instance
func (instance *Msvm_SharedDiskSetReplicationRelationship) SetPropertyCollectionID(value string) (err error) {
	return instance.SetProperty("CollectionID", (value))
}

// GetCollectionID gets the value of CollectionID for the instance
func (instance *Msvm_SharedDiskSetReplicationRelationship) GetPropertyCollectionID() (value string, err error) {
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

// SetDescription sets the value of Description for the instance
func (instance *Msvm_SharedDiskSetReplicationRelationship) SetPropertyDescription(value string) (err error) {
	return instance.SetProperty("Description", (value))
}

// GetDescription gets the value of Description for the instance
func (instance *Msvm_SharedDiskSetReplicationRelationship) GetPropertyDescription() (value string, err error) {
	retValue, err := instance.GetProperty("Description")
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

// SetElementName sets the value of ElementName for the instance
func (instance *Msvm_SharedDiskSetReplicationRelationship) SetPropertyElementName(value string) (err error) {
	return instance.SetProperty("ElementName", (value))
}

// GetElementName gets the value of ElementName for the instance
func (instance *Msvm_SharedDiskSetReplicationRelationship) GetPropertyElementName() (value string, err error) {
	retValue, err := instance.GetProperty("ElementName")
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

// SetFailedOverReplicationType sets the value of FailedOverReplicationType for the instance
func (instance *Msvm_SharedDiskSetReplicationRelationship) SetPropertyFailedOverReplicationType(value uint16) (err error) {
	return instance.SetProperty("FailedOverReplicationType", (value))
}

// GetFailedOverReplicationType gets the value of FailedOverReplicationType for the instance
func (instance *Msvm_SharedDiskSetReplicationRelationship) GetPropertyFailedOverReplicationType() (value uint16, err error) {
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

// SetLastApplicationConsistentReplicationTime sets the value of LastApplicationConsistentReplicationTime for the instance
func (instance *Msvm_SharedDiskSetReplicationRelationship) SetPropertyLastApplicationConsistentReplicationTime(value string) (err error) {
	return instance.SetProperty("LastApplicationConsistentReplicationTime", (value))
}

// GetLastApplicationConsistentReplicationTime gets the value of LastApplicationConsistentReplicationTime for the instance
func (instance *Msvm_SharedDiskSetReplicationRelationship) GetPropertyLastApplicationConsistentReplicationTime() (value string, err error) {
	retValue, err := instance.GetProperty("LastApplicationConsistentReplicationTime")
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

// SetLastApplyTime sets the value of LastApplyTime for the instance
func (instance *Msvm_SharedDiskSetReplicationRelationship) SetPropertyLastApplyTime(value string) (err error) {
	return instance.SetProperty("LastApplyTime", (value))
}

// GetLastApplyTime gets the value of LastApplyTime for the instance
func (instance *Msvm_SharedDiskSetReplicationRelationship) GetPropertyLastApplyTime() (value string, err error) {
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

// SetLastReplicationTime sets the value of LastReplicationTime for the instance
func (instance *Msvm_SharedDiskSetReplicationRelationship) SetPropertyLastReplicationTime(value string) (err error) {
	return instance.SetProperty("LastReplicationTime", (value))
}

// GetLastReplicationTime gets the value of LastReplicationTime for the instance
func (instance *Msvm_SharedDiskSetReplicationRelationship) GetPropertyLastReplicationTime() (value string, err error) {
	retValue, err := instance.GetProperty("LastReplicationTime")
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

// SetLastReplicationType sets the value of LastReplicationType for the instance
func (instance *Msvm_SharedDiskSetReplicationRelationship) SetPropertyLastReplicationType(value uint16) (err error) {
	return instance.SetProperty("LastReplicationType", (value))
}

// GetLastReplicationType gets the value of LastReplicationType for the instance
func (instance *Msvm_SharedDiskSetReplicationRelationship) GetPropertyLastReplicationType() (value uint16, err error) {
	retValue, err := instance.GetProperty("LastReplicationType")
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

// SetReplicationHealth sets the value of ReplicationHealth for the instance
func (instance *Msvm_SharedDiskSetReplicationRelationship) SetPropertyReplicationHealth(value uint16) (err error) {
	return instance.SetProperty("ReplicationHealth", (value))
}

// GetReplicationHealth gets the value of ReplicationHealth for the instance
func (instance *Msvm_SharedDiskSetReplicationRelationship) GetPropertyReplicationHealth() (value uint16, err error) {
	retValue, err := instance.GetProperty("ReplicationHealth")
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
func (instance *Msvm_SharedDiskSetReplicationRelationship) SetPropertyReplicationState(value uint16) (err error) {
	return instance.SetProperty("ReplicationState", (value))
}

// GetReplicationState gets the value of ReplicationState for the instance
func (instance *Msvm_SharedDiskSetReplicationRelationship) GetPropertyReplicationState() (value uint16, err error) {
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
