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

// Msvm_SummaryInformation struct
type Msvm_SummaryInformation struct {
	*Msvm_SummaryInformationBase

	//
	AllocatedGPU string

	//
	ApplicationHealth uint16

	//
	AsynchronousTasks []CIM_ConcreteJob

	//
	AvailableMemoryBuffer int32

	//
	GuestOperatingSystem string

	//
	Heartbeat uint16

	//
	HypervisorPartitionId uint64

	//
	IntegrationServicesVersionState uint16

	//
	MemoryAvailable int32

	//
	MemorySpansPhysicalNumaNodes bool

	//
	MemoryUsage uint64

	//
	ProcessorLoad uint16

	//
	ProcessorLoadHistory []uint16

	//
	ReplicationHealth uint16

	//
	ReplicationHealthEx []uint16

	//
	ReplicationMode uint16

	//
	ReplicationProviderId []string

	//
	ReplicationState uint16

	//
	ReplicationStateEx []uint16

	//
	Shielded bool

	//
	Snapshots []CIM_VirtualSystemSettingData

	//
	SwapFilesInUse bool

	//
	TestReplicaSystem CIM_ComputerSystem

	//
	ThumbnailImage []uint8

	//
	ThumbnailImageHeight uint16

	//
	ThumbnailImageWidth uint16
}

func NewMsvm_SummaryInformationEx1(instance *cim.WmiInstance) (newInstance *Msvm_SummaryInformation, err error) {
	tmp, err := NewMsvm_SummaryInformationBaseEx1(instance)

	if err != nil {
		return
	}
	newInstance = &Msvm_SummaryInformation{
		Msvm_SummaryInformationBase: tmp,
	}
	return
}

func NewMsvm_SummaryInformationEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *Msvm_SummaryInformation, err error) {
	tmp, err := NewMsvm_SummaryInformationBaseEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &Msvm_SummaryInformation{
		Msvm_SummaryInformationBase: tmp,
	}
	return
}

// SetAllocatedGPU sets the value of AllocatedGPU for the instance
func (instance *Msvm_SummaryInformation) SetPropertyAllocatedGPU(value string) (err error) {
	return instance.SetProperty("AllocatedGPU", (value))
}

// GetAllocatedGPU gets the value of AllocatedGPU for the instance
func (instance *Msvm_SummaryInformation) GetPropertyAllocatedGPU() (value string, err error) {
	retValue, err := instance.GetProperty("AllocatedGPU")
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

// SetApplicationHealth sets the value of ApplicationHealth for the instance
func (instance *Msvm_SummaryInformation) SetPropertyApplicationHealth(value uint16) (err error) {
	return instance.SetProperty("ApplicationHealth", (value))
}

// GetApplicationHealth gets the value of ApplicationHealth for the instance
func (instance *Msvm_SummaryInformation) GetPropertyApplicationHealth() (value uint16, err error) {
	retValue, err := instance.GetProperty("ApplicationHealth")
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

// SetAsynchronousTasks sets the value of AsynchronousTasks for the instance
func (instance *Msvm_SummaryInformation) SetPropertyAsynchronousTasks(value []CIM_ConcreteJob) (err error) {
	return instance.SetProperty("AsynchronousTasks", (value))
}

// GetAsynchronousTasks gets the value of AsynchronousTasks for the instance
func (instance *Msvm_SummaryInformation) GetPropertyAsynchronousTasks() (value []CIM_ConcreteJob, err error) {
	retValue, err := instance.GetProperty("AsynchronousTasks")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	for _, interfaceValue := range retValue.([]interface{}) {
		valuetmp, ok := interfaceValue.(CIM_ConcreteJob)
		if !ok {
			err = errors.Wrapf(errors.InvalidType, " CIM_ConcreteJob is Invalid. Expected %s", reflect.TypeOf(interfaceValue))
			return
		}
		value = append(value, CIM_ConcreteJob(valuetmp))
	}

	return
}

// SetAvailableMemoryBuffer sets the value of AvailableMemoryBuffer for the instance
func (instance *Msvm_SummaryInformation) SetPropertyAvailableMemoryBuffer(value int32) (err error) {
	return instance.SetProperty("AvailableMemoryBuffer", (value))
}

// GetAvailableMemoryBuffer gets the value of AvailableMemoryBuffer for the instance
func (instance *Msvm_SummaryInformation) GetPropertyAvailableMemoryBuffer() (value int32, err error) {
	retValue, err := instance.GetProperty("AvailableMemoryBuffer")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(int32)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " int32 is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = int32(valuetmp)

	return
}

// SetGuestOperatingSystem sets the value of GuestOperatingSystem for the instance
func (instance *Msvm_SummaryInformation) SetPropertyGuestOperatingSystem(value string) (err error) {
	return instance.SetProperty("GuestOperatingSystem", (value))
}

// GetGuestOperatingSystem gets the value of GuestOperatingSystem for the instance
func (instance *Msvm_SummaryInformation) GetPropertyGuestOperatingSystem() (value string, err error) {
	retValue, err := instance.GetProperty("GuestOperatingSystem")
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

// SetHeartbeat sets the value of Heartbeat for the instance
func (instance *Msvm_SummaryInformation) SetPropertyHeartbeat(value uint16) (err error) {
	return instance.SetProperty("Heartbeat", (value))
}

// GetHeartbeat gets the value of Heartbeat for the instance
func (instance *Msvm_SummaryInformation) GetPropertyHeartbeat() (value uint16, err error) {
	retValue, err := instance.GetProperty("Heartbeat")
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

// SetHypervisorPartitionId sets the value of HypervisorPartitionId for the instance
func (instance *Msvm_SummaryInformation) SetPropertyHypervisorPartitionId(value uint64) (err error) {
	return instance.SetProperty("HypervisorPartitionId", (value))
}

// GetHypervisorPartitionId gets the value of HypervisorPartitionId for the instance
func (instance *Msvm_SummaryInformation) GetPropertyHypervisorPartitionId() (value uint64, err error) {
	retValue, err := instance.GetProperty("HypervisorPartitionId")
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

// SetIntegrationServicesVersionState sets the value of IntegrationServicesVersionState for the instance
func (instance *Msvm_SummaryInformation) SetPropertyIntegrationServicesVersionState(value uint16) (err error) {
	return instance.SetProperty("IntegrationServicesVersionState", (value))
}

// GetIntegrationServicesVersionState gets the value of IntegrationServicesVersionState for the instance
func (instance *Msvm_SummaryInformation) GetPropertyIntegrationServicesVersionState() (value uint16, err error) {
	retValue, err := instance.GetProperty("IntegrationServicesVersionState")
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

// SetMemoryAvailable sets the value of MemoryAvailable for the instance
func (instance *Msvm_SummaryInformation) SetPropertyMemoryAvailable(value int32) (err error) {
	return instance.SetProperty("MemoryAvailable", (value))
}

// GetMemoryAvailable gets the value of MemoryAvailable for the instance
func (instance *Msvm_SummaryInformation) GetPropertyMemoryAvailable() (value int32, err error) {
	retValue, err := instance.GetProperty("MemoryAvailable")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(int32)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " int32 is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = int32(valuetmp)

	return
}

// SetMemorySpansPhysicalNumaNodes sets the value of MemorySpansPhysicalNumaNodes for the instance
func (instance *Msvm_SummaryInformation) SetPropertyMemorySpansPhysicalNumaNodes(value bool) (err error) {
	return instance.SetProperty("MemorySpansPhysicalNumaNodes", (value))
}

// GetMemorySpansPhysicalNumaNodes gets the value of MemorySpansPhysicalNumaNodes for the instance
func (instance *Msvm_SummaryInformation) GetPropertyMemorySpansPhysicalNumaNodes() (value bool, err error) {
	retValue, err := instance.GetProperty("MemorySpansPhysicalNumaNodes")
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

// SetMemoryUsage sets the value of MemoryUsage for the instance
func (instance *Msvm_SummaryInformation) SetPropertyMemoryUsage(value uint64) (err error) {
	return instance.SetProperty("MemoryUsage", (value))
}

// GetMemoryUsage gets the value of MemoryUsage for the instance
func (instance *Msvm_SummaryInformation) GetPropertyMemoryUsage() (value uint64, err error) {
	retValue, err := instance.GetProperty("MemoryUsage")
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

// SetProcessorLoad sets the value of ProcessorLoad for the instance
func (instance *Msvm_SummaryInformation) SetPropertyProcessorLoad(value uint16) (err error) {
	return instance.SetProperty("ProcessorLoad", (value))
}

// GetProcessorLoad gets the value of ProcessorLoad for the instance
func (instance *Msvm_SummaryInformation) GetPropertyProcessorLoad() (value uint16, err error) {
	retValue, err := instance.GetProperty("ProcessorLoad")
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

// SetProcessorLoadHistory sets the value of ProcessorLoadHistory for the instance
func (instance *Msvm_SummaryInformation) SetPropertyProcessorLoadHistory(value []uint16) (err error) {
	return instance.SetProperty("ProcessorLoadHistory", (value))
}

// GetProcessorLoadHistory gets the value of ProcessorLoadHistory for the instance
func (instance *Msvm_SummaryInformation) GetPropertyProcessorLoadHistory() (value []uint16, err error) {
	retValue, err := instance.GetProperty("ProcessorLoadHistory")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	for _, interfaceValue := range retValue.([]interface{}) {
		valuetmp, ok := interfaceValue.(uint16)
		if !ok {
			err = errors.Wrapf(errors.InvalidType, " uint16 is Invalid. Expected %s", reflect.TypeOf(interfaceValue))
			return
		}
		value = append(value, uint16(valuetmp))
	}

	return
}

// SetReplicationHealth sets the value of ReplicationHealth for the instance
func (instance *Msvm_SummaryInformation) SetPropertyReplicationHealth(value uint16) (err error) {
	return instance.SetProperty("ReplicationHealth", (value))
}

// GetReplicationHealth gets the value of ReplicationHealth for the instance
func (instance *Msvm_SummaryInformation) GetPropertyReplicationHealth() (value uint16, err error) {
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

// SetReplicationHealthEx sets the value of ReplicationHealthEx for the instance
func (instance *Msvm_SummaryInformation) SetPropertyReplicationHealthEx(value []uint16) (err error) {
	return instance.SetProperty("ReplicationHealthEx", (value))
}

// GetReplicationHealthEx gets the value of ReplicationHealthEx for the instance
func (instance *Msvm_SummaryInformation) GetPropertyReplicationHealthEx() (value []uint16, err error) {
	retValue, err := instance.GetProperty("ReplicationHealthEx")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	for _, interfaceValue := range retValue.([]interface{}) {
		valuetmp, ok := interfaceValue.(uint16)
		if !ok {
			err = errors.Wrapf(errors.InvalidType, " uint16 is Invalid. Expected %s", reflect.TypeOf(interfaceValue))
			return
		}
		value = append(value, uint16(valuetmp))
	}

	return
}

// SetReplicationMode sets the value of ReplicationMode for the instance
func (instance *Msvm_SummaryInformation) SetPropertyReplicationMode(value uint16) (err error) {
	return instance.SetProperty("ReplicationMode", (value))
}

// GetReplicationMode gets the value of ReplicationMode for the instance
func (instance *Msvm_SummaryInformation) GetPropertyReplicationMode() (value uint16, err error) {
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

// SetReplicationProviderId sets the value of ReplicationProviderId for the instance
func (instance *Msvm_SummaryInformation) SetPropertyReplicationProviderId(value []string) (err error) {
	return instance.SetProperty("ReplicationProviderId", (value))
}

// GetReplicationProviderId gets the value of ReplicationProviderId for the instance
func (instance *Msvm_SummaryInformation) GetPropertyReplicationProviderId() (value []string, err error) {
	retValue, err := instance.GetProperty("ReplicationProviderId")
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

// SetReplicationState sets the value of ReplicationState for the instance
func (instance *Msvm_SummaryInformation) SetPropertyReplicationState(value uint16) (err error) {
	return instance.SetProperty("ReplicationState", (value))
}

// GetReplicationState gets the value of ReplicationState for the instance
func (instance *Msvm_SummaryInformation) GetPropertyReplicationState() (value uint16, err error) {
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

// SetReplicationStateEx sets the value of ReplicationStateEx for the instance
func (instance *Msvm_SummaryInformation) SetPropertyReplicationStateEx(value []uint16) (err error) {
	return instance.SetProperty("ReplicationStateEx", (value))
}

// GetReplicationStateEx gets the value of ReplicationStateEx for the instance
func (instance *Msvm_SummaryInformation) GetPropertyReplicationStateEx() (value []uint16, err error) {
	retValue, err := instance.GetProperty("ReplicationStateEx")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	for _, interfaceValue := range retValue.([]interface{}) {
		valuetmp, ok := interfaceValue.(uint16)
		if !ok {
			err = errors.Wrapf(errors.InvalidType, " uint16 is Invalid. Expected %s", reflect.TypeOf(interfaceValue))
			return
		}
		value = append(value, uint16(valuetmp))
	}

	return
}

// SetShielded sets the value of Shielded for the instance
func (instance *Msvm_SummaryInformation) SetPropertyShielded(value bool) (err error) {
	return instance.SetProperty("Shielded", (value))
}

// GetShielded gets the value of Shielded for the instance
func (instance *Msvm_SummaryInformation) GetPropertyShielded() (value bool, err error) {
	retValue, err := instance.GetProperty("Shielded")
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

// SetSnapshots sets the value of Snapshots for the instance
func (instance *Msvm_SummaryInformation) SetPropertySnapshots(value []CIM_VirtualSystemSettingData) (err error) {
	return instance.SetProperty("Snapshots", (value))
}

// GetSnapshots gets the value of Snapshots for the instance
func (instance *Msvm_SummaryInformation) GetPropertySnapshots() (value []CIM_VirtualSystemSettingData, err error) {
	retValue, err := instance.GetProperty("Snapshots")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	for _, interfaceValue := range retValue.([]interface{}) {
		valuetmp, ok := interfaceValue.(CIM_VirtualSystemSettingData)
		if !ok {
			err = errors.Wrapf(errors.InvalidType, " CIM_VirtualSystemSettingData is Invalid. Expected %s", reflect.TypeOf(interfaceValue))
			return
		}
		value = append(value, CIM_VirtualSystemSettingData(valuetmp))
	}

	return
}

// SetSwapFilesInUse sets the value of SwapFilesInUse for the instance
func (instance *Msvm_SummaryInformation) SetPropertySwapFilesInUse(value bool) (err error) {
	return instance.SetProperty("SwapFilesInUse", (value))
}

// GetSwapFilesInUse gets the value of SwapFilesInUse for the instance
func (instance *Msvm_SummaryInformation) GetPropertySwapFilesInUse() (value bool, err error) {
	retValue, err := instance.GetProperty("SwapFilesInUse")
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

// SetTestReplicaSystem sets the value of TestReplicaSystem for the instance
func (instance *Msvm_SummaryInformation) SetPropertyTestReplicaSystem(value CIM_ComputerSystem) (err error) {
	return instance.SetProperty("TestReplicaSystem", (value))
}

// GetTestReplicaSystem gets the value of TestReplicaSystem for the instance
func (instance *Msvm_SummaryInformation) GetPropertyTestReplicaSystem() (value CIM_ComputerSystem, err error) {
	retValue, err := instance.GetProperty("TestReplicaSystem")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(CIM_ComputerSystem)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " CIM_ComputerSystem is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = CIM_ComputerSystem(valuetmp)

	return
}

// SetThumbnailImage sets the value of ThumbnailImage for the instance
func (instance *Msvm_SummaryInformation) SetPropertyThumbnailImage(value []uint8) (err error) {
	return instance.SetProperty("ThumbnailImage", (value))
}

// GetThumbnailImage gets the value of ThumbnailImage for the instance
func (instance *Msvm_SummaryInformation) GetPropertyThumbnailImage() (value []uint8, err error) {
	retValue, err := instance.GetProperty("ThumbnailImage")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	for _, interfaceValue := range retValue.([]interface{}) {
		valuetmp, ok := interfaceValue.(uint8)
		if !ok {
			err = errors.Wrapf(errors.InvalidType, " uint8 is Invalid. Expected %s", reflect.TypeOf(interfaceValue))
			return
		}
		value = append(value, uint8(valuetmp))
	}

	return
}

// SetThumbnailImageHeight sets the value of ThumbnailImageHeight for the instance
func (instance *Msvm_SummaryInformation) SetPropertyThumbnailImageHeight(value uint16) (err error) {
	return instance.SetProperty("ThumbnailImageHeight", (value))
}

// GetThumbnailImageHeight gets the value of ThumbnailImageHeight for the instance
func (instance *Msvm_SummaryInformation) GetPropertyThumbnailImageHeight() (value uint16, err error) {
	retValue, err := instance.GetProperty("ThumbnailImageHeight")
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

// SetThumbnailImageWidth sets the value of ThumbnailImageWidth for the instance
func (instance *Msvm_SummaryInformation) SetPropertyThumbnailImageWidth(value uint16) (err error) {
	return instance.SetProperty("ThumbnailImageWidth", (value))
}

// GetThumbnailImageWidth gets the value of ThumbnailImageWidth for the instance
func (instance *Msvm_SummaryInformation) GetPropertyThumbnailImageWidth() (value uint16, err error) {
	retValue, err := instance.GetProperty("ThumbnailImageWidth")
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
