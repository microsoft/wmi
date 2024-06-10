// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

// Author:
//
//	Auto Generated on 6/6/2024 using wmigen
//	Source root.Microsoft.Windows.Cluster.Scaleout
//
// ////////////////////////////////////////////
package scaleout

import (
	"github.com/microsoft/wmi/pkg/base/instance"
	"github.com/microsoft/wmi/pkg/base/query"
	"github.com/microsoft/wmi/pkg/errors"
	cim "github.com/microsoft/wmi/pkg/wmiinstance"
	"reflect"
)

// MSFT_ClusterSetNode struct
type MSFT_ClusterSetNode struct {
	*cim.WmiInstance

	//
	AvailableMemory uint64

	//
	AvailableMemoryAfterReclaimation uint64

	//
	AverageCpuUsage uint32

	//
	FreeCpuReserve uint64

	//
	Id uint64

	//
	LocalDiskFreeSpaceInMB uint32

	//
	LocalDiskTotalSpaceInMB uint32

	//
	MaxCpuReserve uint64

	//
	Member string

	//
	MemberId uint64

	//
	Name string

	//
	NodeId uint32

	//
	NodeLPCount uint32

	//
	ReserveCpu uint64

	//
	ReservedLocalDiskUsage uint32

	//
	ReservedMemory uint64

	//
	State uint32

	//
	TotalMemory uint64
}

func NewMSFT_ClusterSetNodeEx1(instance *cim.WmiInstance) (newInstance *MSFT_ClusterSetNode, err error) {
	tmp, err := instance, nil

	if err != nil {
		return
	}
	newInstance = &MSFT_ClusterSetNode{
		WmiInstance: tmp,
	}
	return
}

func NewMSFT_ClusterSetNodeEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *MSFT_ClusterSetNode, err error) {
	tmp, err := instance.GetWmiInstance(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &MSFT_ClusterSetNode{
		WmiInstance: tmp,
	}
	return
}

// SetAvailableMemory sets the value of AvailableMemory for the instance
func (instance *MSFT_ClusterSetNode) SetPropertyAvailableMemory(value uint64) (err error) {
	return instance.SetProperty("AvailableMemory", (value))
}

// GetAvailableMemory gets the value of AvailableMemory for the instance
func (instance *MSFT_ClusterSetNode) GetPropertyAvailableMemory() (value uint64, err error) {
	retValue, err := instance.GetProperty("AvailableMemory")
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

// SetAvailableMemoryAfterReclaimation sets the value of AvailableMemoryAfterReclaimation for the instance
func (instance *MSFT_ClusterSetNode) SetPropertyAvailableMemoryAfterReclaimation(value uint64) (err error) {
	return instance.SetProperty("AvailableMemoryAfterReclaimation", (value))
}

// GetAvailableMemoryAfterReclaimation gets the value of AvailableMemoryAfterReclaimation for the instance
func (instance *MSFT_ClusterSetNode) GetPropertyAvailableMemoryAfterReclaimation() (value uint64, err error) {
	retValue, err := instance.GetProperty("AvailableMemoryAfterReclaimation")
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

// SetAverageCpuUsage sets the value of AverageCpuUsage for the instance
func (instance *MSFT_ClusterSetNode) SetPropertyAverageCpuUsage(value uint32) (err error) {
	return instance.SetProperty("AverageCpuUsage", (value))
}

// GetAverageCpuUsage gets the value of AverageCpuUsage for the instance
func (instance *MSFT_ClusterSetNode) GetPropertyAverageCpuUsage() (value uint32, err error) {
	retValue, err := instance.GetProperty("AverageCpuUsage")
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

// SetFreeCpuReserve sets the value of FreeCpuReserve for the instance
func (instance *MSFT_ClusterSetNode) SetPropertyFreeCpuReserve(value uint64) (err error) {
	return instance.SetProperty("FreeCpuReserve", (value))
}

// GetFreeCpuReserve gets the value of FreeCpuReserve for the instance
func (instance *MSFT_ClusterSetNode) GetPropertyFreeCpuReserve() (value uint64, err error) {
	retValue, err := instance.GetProperty("FreeCpuReserve")
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

// SetId sets the value of Id for the instance
func (instance *MSFT_ClusterSetNode) SetPropertyId(value uint64) (err error) {
	return instance.SetProperty("Id", (value))
}

// GetId gets the value of Id for the instance
func (instance *MSFT_ClusterSetNode) GetPropertyId() (value uint64, err error) {
	retValue, err := instance.GetProperty("Id")
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

// SetLocalDiskFreeSpaceInMB sets the value of LocalDiskFreeSpaceInMB for the instance
func (instance *MSFT_ClusterSetNode) SetPropertyLocalDiskFreeSpaceInMB(value uint32) (err error) {
	return instance.SetProperty("LocalDiskFreeSpaceInMB", (value))
}

// GetLocalDiskFreeSpaceInMB gets the value of LocalDiskFreeSpaceInMB for the instance
func (instance *MSFT_ClusterSetNode) GetPropertyLocalDiskFreeSpaceInMB() (value uint32, err error) {
	retValue, err := instance.GetProperty("LocalDiskFreeSpaceInMB")
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

// SetLocalDiskTotalSpaceInMB sets the value of LocalDiskTotalSpaceInMB for the instance
func (instance *MSFT_ClusterSetNode) SetPropertyLocalDiskTotalSpaceInMB(value uint32) (err error) {
	return instance.SetProperty("LocalDiskTotalSpaceInMB", (value))
}

// GetLocalDiskTotalSpaceInMB gets the value of LocalDiskTotalSpaceInMB for the instance
func (instance *MSFT_ClusterSetNode) GetPropertyLocalDiskTotalSpaceInMB() (value uint32, err error) {
	retValue, err := instance.GetProperty("LocalDiskTotalSpaceInMB")
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

// SetMaxCpuReserve sets the value of MaxCpuReserve for the instance
func (instance *MSFT_ClusterSetNode) SetPropertyMaxCpuReserve(value uint64) (err error) {
	return instance.SetProperty("MaxCpuReserve", (value))
}

// GetMaxCpuReserve gets the value of MaxCpuReserve for the instance
func (instance *MSFT_ClusterSetNode) GetPropertyMaxCpuReserve() (value uint64, err error) {
	retValue, err := instance.GetProperty("MaxCpuReserve")
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

// SetMember sets the value of Member for the instance
func (instance *MSFT_ClusterSetNode) SetPropertyMember(value string) (err error) {
	return instance.SetProperty("Member", (value))
}

// GetMember gets the value of Member for the instance
func (instance *MSFT_ClusterSetNode) GetPropertyMember() (value string, err error) {
	retValue, err := instance.GetProperty("Member")
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

// SetMemberId sets the value of MemberId for the instance
func (instance *MSFT_ClusterSetNode) SetPropertyMemberId(value uint64) (err error) {
	return instance.SetProperty("MemberId", (value))
}

// GetMemberId gets the value of MemberId for the instance
func (instance *MSFT_ClusterSetNode) GetPropertyMemberId() (value uint64, err error) {
	retValue, err := instance.GetProperty("MemberId")
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

// SetName sets the value of Name for the instance
func (instance *MSFT_ClusterSetNode) SetPropertyName(value string) (err error) {
	return instance.SetProperty("Name", (value))
}

// GetName gets the value of Name for the instance
func (instance *MSFT_ClusterSetNode) GetPropertyName() (value string, err error) {
	retValue, err := instance.GetProperty("Name")
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

// SetNodeId sets the value of NodeId for the instance
func (instance *MSFT_ClusterSetNode) SetPropertyNodeId(value uint32) (err error) {
	return instance.SetProperty("NodeId", (value))
}

// GetNodeId gets the value of NodeId for the instance
func (instance *MSFT_ClusterSetNode) GetPropertyNodeId() (value uint32, err error) {
	retValue, err := instance.GetProperty("NodeId")
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

// SetNodeLPCount sets the value of NodeLPCount for the instance
func (instance *MSFT_ClusterSetNode) SetPropertyNodeLPCount(value uint32) (err error) {
	return instance.SetProperty("NodeLPCount", (value))
}

// GetNodeLPCount gets the value of NodeLPCount for the instance
func (instance *MSFT_ClusterSetNode) GetPropertyNodeLPCount() (value uint32, err error) {
	retValue, err := instance.GetProperty("NodeLPCount")
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

// SetReserveCpu sets the value of ReserveCpu for the instance
func (instance *MSFT_ClusterSetNode) SetPropertyReserveCpu(value uint64) (err error) {
	return instance.SetProperty("ReserveCpu", (value))
}

// GetReserveCpu gets the value of ReserveCpu for the instance
func (instance *MSFT_ClusterSetNode) GetPropertyReserveCpu() (value uint64, err error) {
	retValue, err := instance.GetProperty("ReserveCpu")
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

// SetReservedLocalDiskUsage sets the value of ReservedLocalDiskUsage for the instance
func (instance *MSFT_ClusterSetNode) SetPropertyReservedLocalDiskUsage(value uint32) (err error) {
	return instance.SetProperty("ReservedLocalDiskUsage", (value))
}

// GetReservedLocalDiskUsage gets the value of ReservedLocalDiskUsage for the instance
func (instance *MSFT_ClusterSetNode) GetPropertyReservedLocalDiskUsage() (value uint32, err error) {
	retValue, err := instance.GetProperty("ReservedLocalDiskUsage")
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

// SetReservedMemory sets the value of ReservedMemory for the instance
func (instance *MSFT_ClusterSetNode) SetPropertyReservedMemory(value uint64) (err error) {
	return instance.SetProperty("ReservedMemory", (value))
}

// GetReservedMemory gets the value of ReservedMemory for the instance
func (instance *MSFT_ClusterSetNode) GetPropertyReservedMemory() (value uint64, err error) {
	retValue, err := instance.GetProperty("ReservedMemory")
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

// SetState sets the value of State for the instance
func (instance *MSFT_ClusterSetNode) SetPropertyState(value uint32) (err error) {
	return instance.SetProperty("State", (value))
}

// GetState gets the value of State for the instance
func (instance *MSFT_ClusterSetNode) GetPropertyState() (value uint32, err error) {
	retValue, err := instance.GetProperty("State")
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

// SetTotalMemory sets the value of TotalMemory for the instance
func (instance *MSFT_ClusterSetNode) SetPropertyTotalMemory(value uint64) (err error) {
	return instance.SetProperty("TotalMemory", (value))
}

// GetTotalMemory gets the value of TotalMemory for the instance
func (instance *MSFT_ClusterSetNode) GetPropertyTotalMemory() (value uint64, err error) {
	retValue, err := instance.GetProperty("TotalMemory")
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
