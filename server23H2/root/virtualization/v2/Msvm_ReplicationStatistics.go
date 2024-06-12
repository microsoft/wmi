// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 6/12/2024 using wmigen
//      Source root.virtualization.v2
//////////////////////////////////////////////
package v2

import (
	"github.com/microsoft/wmi/pkg/base/query"
	"github.com/microsoft/wmi/pkg/errors"
	cim "github.com/microsoft/wmi/pkg/wmiinstance"
	"reflect"
)

// Msvm_ReplicationStatistics struct
type Msvm_ReplicationStatistics struct {
	*CIM_ManagedElement

	//
	ApplicationConsistentSnapshotFailureCount uint32

	//
	EndStatisticTime string

	//
	LastTestFailoverTime string

	//
	MaxReplicationLatency uint32

	//
	MaxReplicationSize uint64

	//
	NetworkFailureCount uint32

	//
	PendingReplicationSize uint64

	//
	ReplicationFailureCount uint32

	//
	ReplicationHealth uint32

	//
	ReplicationLatency uint32

	//
	ReplicationMissCount uint32

	//
	ReplicationSize uint64

	//
	ReplicationSuccessCount uint32

	//
	StartStatisticTime string
}

func NewMsvm_ReplicationStatisticsEx1(instance *cim.WmiInstance) (newInstance *Msvm_ReplicationStatistics, err error) {
	tmp, err := NewCIM_ManagedElementEx1(instance)

	if err != nil {
		return
	}
	newInstance = &Msvm_ReplicationStatistics{
		CIM_ManagedElement: tmp,
	}
	return
}

func NewMsvm_ReplicationStatisticsEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *Msvm_ReplicationStatistics, err error) {
	tmp, err := NewCIM_ManagedElementEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &Msvm_ReplicationStatistics{
		CIM_ManagedElement: tmp,
	}
	return
}

// SetApplicationConsistentSnapshotFailureCount sets the value of ApplicationConsistentSnapshotFailureCount for the instance
func (instance *Msvm_ReplicationStatistics) SetPropertyApplicationConsistentSnapshotFailureCount(value uint32) (err error) {
	return instance.SetProperty("ApplicationConsistentSnapshotFailureCount", (value))
}

// GetApplicationConsistentSnapshotFailureCount gets the value of ApplicationConsistentSnapshotFailureCount for the instance
func (instance *Msvm_ReplicationStatistics) GetPropertyApplicationConsistentSnapshotFailureCount() (value uint32, err error) {
	retValue, err := instance.GetProperty("ApplicationConsistentSnapshotFailureCount")
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

// SetEndStatisticTime sets the value of EndStatisticTime for the instance
func (instance *Msvm_ReplicationStatistics) SetPropertyEndStatisticTime(value string) (err error) {
	return instance.SetProperty("EndStatisticTime", (value))
}

// GetEndStatisticTime gets the value of EndStatisticTime for the instance
func (instance *Msvm_ReplicationStatistics) GetPropertyEndStatisticTime() (value string, err error) {
	retValue, err := instance.GetProperty("EndStatisticTime")
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

// SetLastTestFailoverTime sets the value of LastTestFailoverTime for the instance
func (instance *Msvm_ReplicationStatistics) SetPropertyLastTestFailoverTime(value string) (err error) {
	return instance.SetProperty("LastTestFailoverTime", (value))
}

// GetLastTestFailoverTime gets the value of LastTestFailoverTime for the instance
func (instance *Msvm_ReplicationStatistics) GetPropertyLastTestFailoverTime() (value string, err error) {
	retValue, err := instance.GetProperty("LastTestFailoverTime")
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

// SetMaxReplicationLatency sets the value of MaxReplicationLatency for the instance
func (instance *Msvm_ReplicationStatistics) SetPropertyMaxReplicationLatency(value uint32) (err error) {
	return instance.SetProperty("MaxReplicationLatency", (value))
}

// GetMaxReplicationLatency gets the value of MaxReplicationLatency for the instance
func (instance *Msvm_ReplicationStatistics) GetPropertyMaxReplicationLatency() (value uint32, err error) {
	retValue, err := instance.GetProperty("MaxReplicationLatency")
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

// SetMaxReplicationSize sets the value of MaxReplicationSize for the instance
func (instance *Msvm_ReplicationStatistics) SetPropertyMaxReplicationSize(value uint64) (err error) {
	return instance.SetProperty("MaxReplicationSize", (value))
}

// GetMaxReplicationSize gets the value of MaxReplicationSize for the instance
func (instance *Msvm_ReplicationStatistics) GetPropertyMaxReplicationSize() (value uint64, err error) {
	retValue, err := instance.GetProperty("MaxReplicationSize")
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

// SetNetworkFailureCount sets the value of NetworkFailureCount for the instance
func (instance *Msvm_ReplicationStatistics) SetPropertyNetworkFailureCount(value uint32) (err error) {
	return instance.SetProperty("NetworkFailureCount", (value))
}

// GetNetworkFailureCount gets the value of NetworkFailureCount for the instance
func (instance *Msvm_ReplicationStatistics) GetPropertyNetworkFailureCount() (value uint32, err error) {
	retValue, err := instance.GetProperty("NetworkFailureCount")
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

// SetPendingReplicationSize sets the value of PendingReplicationSize for the instance
func (instance *Msvm_ReplicationStatistics) SetPropertyPendingReplicationSize(value uint64) (err error) {
	return instance.SetProperty("PendingReplicationSize", (value))
}

// GetPendingReplicationSize gets the value of PendingReplicationSize for the instance
func (instance *Msvm_ReplicationStatistics) GetPropertyPendingReplicationSize() (value uint64, err error) {
	retValue, err := instance.GetProperty("PendingReplicationSize")
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

// SetReplicationFailureCount sets the value of ReplicationFailureCount for the instance
func (instance *Msvm_ReplicationStatistics) SetPropertyReplicationFailureCount(value uint32) (err error) {
	return instance.SetProperty("ReplicationFailureCount", (value))
}

// GetReplicationFailureCount gets the value of ReplicationFailureCount for the instance
func (instance *Msvm_ReplicationStatistics) GetPropertyReplicationFailureCount() (value uint32, err error) {
	retValue, err := instance.GetProperty("ReplicationFailureCount")
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

// SetReplicationHealth sets the value of ReplicationHealth for the instance
func (instance *Msvm_ReplicationStatistics) SetPropertyReplicationHealth(value uint32) (err error) {
	return instance.SetProperty("ReplicationHealth", (value))
}

// GetReplicationHealth gets the value of ReplicationHealth for the instance
func (instance *Msvm_ReplicationStatistics) GetPropertyReplicationHealth() (value uint32, err error) {
	retValue, err := instance.GetProperty("ReplicationHealth")
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

// SetReplicationLatency sets the value of ReplicationLatency for the instance
func (instance *Msvm_ReplicationStatistics) SetPropertyReplicationLatency(value uint32) (err error) {
	return instance.SetProperty("ReplicationLatency", (value))
}

// GetReplicationLatency gets the value of ReplicationLatency for the instance
func (instance *Msvm_ReplicationStatistics) GetPropertyReplicationLatency() (value uint32, err error) {
	retValue, err := instance.GetProperty("ReplicationLatency")
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

// SetReplicationMissCount sets the value of ReplicationMissCount for the instance
func (instance *Msvm_ReplicationStatistics) SetPropertyReplicationMissCount(value uint32) (err error) {
	return instance.SetProperty("ReplicationMissCount", (value))
}

// GetReplicationMissCount gets the value of ReplicationMissCount for the instance
func (instance *Msvm_ReplicationStatistics) GetPropertyReplicationMissCount() (value uint32, err error) {
	retValue, err := instance.GetProperty("ReplicationMissCount")
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

// SetReplicationSize sets the value of ReplicationSize for the instance
func (instance *Msvm_ReplicationStatistics) SetPropertyReplicationSize(value uint64) (err error) {
	return instance.SetProperty("ReplicationSize", (value))
}

// GetReplicationSize gets the value of ReplicationSize for the instance
func (instance *Msvm_ReplicationStatistics) GetPropertyReplicationSize() (value uint64, err error) {
	retValue, err := instance.GetProperty("ReplicationSize")
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

// SetReplicationSuccessCount sets the value of ReplicationSuccessCount for the instance
func (instance *Msvm_ReplicationStatistics) SetPropertyReplicationSuccessCount(value uint32) (err error) {
	return instance.SetProperty("ReplicationSuccessCount", (value))
}

// GetReplicationSuccessCount gets the value of ReplicationSuccessCount for the instance
func (instance *Msvm_ReplicationStatistics) GetPropertyReplicationSuccessCount() (value uint32, err error) {
	retValue, err := instance.GetProperty("ReplicationSuccessCount")
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

// SetStartStatisticTime sets the value of StartStatisticTime for the instance
func (instance *Msvm_ReplicationStatistics) SetPropertyStartStatisticTime(value string) (err error) {
	return instance.SetProperty("StartStatisticTime", (value))
}

// GetStartStatisticTime gets the value of StartStatisticTime for the instance
func (instance *Msvm_ReplicationStatistics) GetPropertyStartStatisticTime() (value string, err error) {
	retValue, err := instance.GetProperty("StartStatisticTime")
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
