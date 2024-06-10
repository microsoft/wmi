// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 9/18/2020 using wmigen
//      Source root.Microsoft.Windows.ManagementTools
//////////////////////////////////////////////
package managementtools

import (
	"github.com/microsoft/wmi/pkg/base/query"
	"github.com/microsoft/wmi/pkg/errors"
	cim "github.com/microsoft/wmi/pkg/wmiinstance"
	"reflect"
)

// MSFT_MTProcess struct
type MSFT_MTProcess struct {
	*CIM_ManagedElement

	//
	BasePriority uint32

	//
	CommandLine string

	//
	CommitCharge uint64

	//
	CpuPercent float32

	//
	CpuTime uint64

	//
	CreationDate string

	//
	CreationTime uint64

	//
	CyclePercent float32

	//
	CycleTime uint64

	//
	DataExecutionPrevention bool

	//
	DeltaPageFaults int32

	//
	DeltaWorkingSetSize int64

	//
	Elevated bool

	//
	ExecutablePath string

	//
	GdiObjects uint32

	//
	HandleCount uint32

	//
	IntervalSeconds uint16

	//
	IsImmersive bool

	//
	Name string

	//
	NonPagedPool uint64

	//
	OperatingSystemContext uint16

	//
	OtherOperationCount uint64

	//
	OtherTransferCount uint64

	//
	PagedPool uint64

	//
	PageFaults uint32

	//
	PeakWorkingSetSize uint64

	//
	Platform uint16

	//
	PrivateWorkingSetSize uint64

	//
	ProcessId uint32

	//
	ProcessStatus uint16

	//
	ReadOperationCount uint64

	//
	ReadTransferCount uint64

	//
	SessionId uint32

	//
	SharedWorkingSetSize uint64

	//
	ThreadCount uint32

	//
	UACVirtualization uint16

	//
	UserName string

	//
	UserObjects uint32

	//
	WorkingSetSize uint64

	//
	WriteOperationCount uint64

	//
	WriteTransferCount uint64
}

func NewMSFT_MTProcessEx1(instance *cim.WmiInstance) (newInstance *MSFT_MTProcess, err error) {
	tmp, err := NewCIM_ManagedElementEx1(instance)

	if err != nil {
		return
	}
	newInstance = &MSFT_MTProcess{
		CIM_ManagedElement: tmp,
	}
	return
}

func NewMSFT_MTProcessEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *MSFT_MTProcess, err error) {
	tmp, err := NewCIM_ManagedElementEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &MSFT_MTProcess{
		CIM_ManagedElement: tmp,
	}
	return
}

// SetBasePriority sets the value of BasePriority for the instance
func (instance *MSFT_MTProcess) SetPropertyBasePriority(value uint32) (err error) {
	return instance.SetProperty("BasePriority", (value))
}

// GetBasePriority gets the value of BasePriority for the instance
func (instance *MSFT_MTProcess) GetPropertyBasePriority() (value uint32, err error) {
	retValue, err := instance.GetProperty("BasePriority")
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

// SetCommandLine sets the value of CommandLine for the instance
func (instance *MSFT_MTProcess) SetPropertyCommandLine(value string) (err error) {
	return instance.SetProperty("CommandLine", (value))
}

// GetCommandLine gets the value of CommandLine for the instance
func (instance *MSFT_MTProcess) GetPropertyCommandLine() (value string, err error) {
	retValue, err := instance.GetProperty("CommandLine")
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

// SetCommitCharge sets the value of CommitCharge for the instance
func (instance *MSFT_MTProcess) SetPropertyCommitCharge(value uint64) (err error) {
	return instance.SetProperty("CommitCharge", (value))
}

// GetCommitCharge gets the value of CommitCharge for the instance
func (instance *MSFT_MTProcess) GetPropertyCommitCharge() (value uint64, err error) {
	retValue, err := instance.GetProperty("CommitCharge")
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

// SetCpuPercent sets the value of CpuPercent for the instance
func (instance *MSFT_MTProcess) SetPropertyCpuPercent(value float32) (err error) {
	return instance.SetProperty("CpuPercent", (value))
}

// GetCpuPercent gets the value of CpuPercent for the instance
func (instance *MSFT_MTProcess) GetPropertyCpuPercent() (value float32, err error) {
	retValue, err := instance.GetProperty("CpuPercent")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(float32)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " float32 is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = float32(valuetmp)

	return
}

// SetCpuTime sets the value of CpuTime for the instance
func (instance *MSFT_MTProcess) SetPropertyCpuTime(value uint64) (err error) {
	return instance.SetProperty("CpuTime", (value))
}

// GetCpuTime gets the value of CpuTime for the instance
func (instance *MSFT_MTProcess) GetPropertyCpuTime() (value uint64, err error) {
	retValue, err := instance.GetProperty("CpuTime")
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

// SetCreationDate sets the value of CreationDate for the instance
func (instance *MSFT_MTProcess) SetPropertyCreationDate(value string) (err error) {
	return instance.SetProperty("CreationDate", (value))
}

// GetCreationDate gets the value of CreationDate for the instance
func (instance *MSFT_MTProcess) GetPropertyCreationDate() (value string, err error) {
	retValue, err := instance.GetProperty("CreationDate")
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

// SetCreationTime sets the value of CreationTime for the instance
func (instance *MSFT_MTProcess) SetPropertyCreationTime(value uint64) (err error) {
	return instance.SetProperty("CreationTime", (value))
}

// GetCreationTime gets the value of CreationTime for the instance
func (instance *MSFT_MTProcess) GetPropertyCreationTime() (value uint64, err error) {
	retValue, err := instance.GetProperty("CreationTime")
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

// SetCyclePercent sets the value of CyclePercent for the instance
func (instance *MSFT_MTProcess) SetPropertyCyclePercent(value float32) (err error) {
	return instance.SetProperty("CyclePercent", (value))
}

// GetCyclePercent gets the value of CyclePercent for the instance
func (instance *MSFT_MTProcess) GetPropertyCyclePercent() (value float32, err error) {
	retValue, err := instance.GetProperty("CyclePercent")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(float32)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " float32 is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = float32(valuetmp)

	return
}

// SetCycleTime sets the value of CycleTime for the instance
func (instance *MSFT_MTProcess) SetPropertyCycleTime(value uint64) (err error) {
	return instance.SetProperty("CycleTime", (value))
}

// GetCycleTime gets the value of CycleTime for the instance
func (instance *MSFT_MTProcess) GetPropertyCycleTime() (value uint64, err error) {
	retValue, err := instance.GetProperty("CycleTime")
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

// SetDataExecutionPrevention sets the value of DataExecutionPrevention for the instance
func (instance *MSFT_MTProcess) SetPropertyDataExecutionPrevention(value bool) (err error) {
	return instance.SetProperty("DataExecutionPrevention", (value))
}

// GetDataExecutionPrevention gets the value of DataExecutionPrevention for the instance
func (instance *MSFT_MTProcess) GetPropertyDataExecutionPrevention() (value bool, err error) {
	retValue, err := instance.GetProperty("DataExecutionPrevention")
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

// SetDeltaPageFaults sets the value of DeltaPageFaults for the instance
func (instance *MSFT_MTProcess) SetPropertyDeltaPageFaults(value int32) (err error) {
	return instance.SetProperty("DeltaPageFaults", (value))
}

// GetDeltaPageFaults gets the value of DeltaPageFaults for the instance
func (instance *MSFT_MTProcess) GetPropertyDeltaPageFaults() (value int32, err error) {
	retValue, err := instance.GetProperty("DeltaPageFaults")
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

// SetDeltaWorkingSetSize sets the value of DeltaWorkingSetSize for the instance
func (instance *MSFT_MTProcess) SetPropertyDeltaWorkingSetSize(value int64) (err error) {
	return instance.SetProperty("DeltaWorkingSetSize", (value))
}

// GetDeltaWorkingSetSize gets the value of DeltaWorkingSetSize for the instance
func (instance *MSFT_MTProcess) GetPropertyDeltaWorkingSetSize() (value int64, err error) {
	retValue, err := instance.GetProperty("DeltaWorkingSetSize")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(int64)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " int64 is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = int64(valuetmp)

	return
}

// SetElevated sets the value of Elevated for the instance
func (instance *MSFT_MTProcess) SetPropertyElevated(value bool) (err error) {
	return instance.SetProperty("Elevated", (value))
}

// GetElevated gets the value of Elevated for the instance
func (instance *MSFT_MTProcess) GetPropertyElevated() (value bool, err error) {
	retValue, err := instance.GetProperty("Elevated")
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

// SetExecutablePath sets the value of ExecutablePath for the instance
func (instance *MSFT_MTProcess) SetPropertyExecutablePath(value string) (err error) {
	return instance.SetProperty("ExecutablePath", (value))
}

// GetExecutablePath gets the value of ExecutablePath for the instance
func (instance *MSFT_MTProcess) GetPropertyExecutablePath() (value string, err error) {
	retValue, err := instance.GetProperty("ExecutablePath")
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

// SetGdiObjects sets the value of GdiObjects for the instance
func (instance *MSFT_MTProcess) SetPropertyGdiObjects(value uint32) (err error) {
	return instance.SetProperty("GdiObjects", (value))
}

// GetGdiObjects gets the value of GdiObjects for the instance
func (instance *MSFT_MTProcess) GetPropertyGdiObjects() (value uint32, err error) {
	retValue, err := instance.GetProperty("GdiObjects")
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

// SetHandleCount sets the value of HandleCount for the instance
func (instance *MSFT_MTProcess) SetPropertyHandleCount(value uint32) (err error) {
	return instance.SetProperty("HandleCount", (value))
}

// GetHandleCount gets the value of HandleCount for the instance
func (instance *MSFT_MTProcess) GetPropertyHandleCount() (value uint32, err error) {
	retValue, err := instance.GetProperty("HandleCount")
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

// SetIntervalSeconds sets the value of IntervalSeconds for the instance
func (instance *MSFT_MTProcess) SetPropertyIntervalSeconds(value uint16) (err error) {
	return instance.SetProperty("IntervalSeconds", (value))
}

// GetIntervalSeconds gets the value of IntervalSeconds for the instance
func (instance *MSFT_MTProcess) GetPropertyIntervalSeconds() (value uint16, err error) {
	retValue, err := instance.GetProperty("IntervalSeconds")
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

// SetIsImmersive sets the value of IsImmersive for the instance
func (instance *MSFT_MTProcess) SetPropertyIsImmersive(value bool) (err error) {
	return instance.SetProperty("IsImmersive", (value))
}

// GetIsImmersive gets the value of IsImmersive for the instance
func (instance *MSFT_MTProcess) GetPropertyIsImmersive() (value bool, err error) {
	retValue, err := instance.GetProperty("IsImmersive")
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

// SetName sets the value of Name for the instance
func (instance *MSFT_MTProcess) SetPropertyName(value string) (err error) {
	return instance.SetProperty("Name", (value))
}

// GetName gets the value of Name for the instance
func (instance *MSFT_MTProcess) GetPropertyName() (value string, err error) {
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

// SetNonPagedPool sets the value of NonPagedPool for the instance
func (instance *MSFT_MTProcess) SetPropertyNonPagedPool(value uint64) (err error) {
	return instance.SetProperty("NonPagedPool", (value))
}

// GetNonPagedPool gets the value of NonPagedPool for the instance
func (instance *MSFT_MTProcess) GetPropertyNonPagedPool() (value uint64, err error) {
	retValue, err := instance.GetProperty("NonPagedPool")
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

// SetOperatingSystemContext sets the value of OperatingSystemContext for the instance
func (instance *MSFT_MTProcess) SetPropertyOperatingSystemContext(value uint16) (err error) {
	return instance.SetProperty("OperatingSystemContext", (value))
}

// GetOperatingSystemContext gets the value of OperatingSystemContext for the instance
func (instance *MSFT_MTProcess) GetPropertyOperatingSystemContext() (value uint16, err error) {
	retValue, err := instance.GetProperty("OperatingSystemContext")
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

// SetOtherOperationCount sets the value of OtherOperationCount for the instance
func (instance *MSFT_MTProcess) SetPropertyOtherOperationCount(value uint64) (err error) {
	return instance.SetProperty("OtherOperationCount", (value))
}

// GetOtherOperationCount gets the value of OtherOperationCount for the instance
func (instance *MSFT_MTProcess) GetPropertyOtherOperationCount() (value uint64, err error) {
	retValue, err := instance.GetProperty("OtherOperationCount")
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

// SetOtherTransferCount sets the value of OtherTransferCount for the instance
func (instance *MSFT_MTProcess) SetPropertyOtherTransferCount(value uint64) (err error) {
	return instance.SetProperty("OtherTransferCount", (value))
}

// GetOtherTransferCount gets the value of OtherTransferCount for the instance
func (instance *MSFT_MTProcess) GetPropertyOtherTransferCount() (value uint64, err error) {
	retValue, err := instance.GetProperty("OtherTransferCount")
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

// SetPagedPool sets the value of PagedPool for the instance
func (instance *MSFT_MTProcess) SetPropertyPagedPool(value uint64) (err error) {
	return instance.SetProperty("PagedPool", (value))
}

// GetPagedPool gets the value of PagedPool for the instance
func (instance *MSFT_MTProcess) GetPropertyPagedPool() (value uint64, err error) {
	retValue, err := instance.GetProperty("PagedPool")
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

// SetPageFaults sets the value of PageFaults for the instance
func (instance *MSFT_MTProcess) SetPropertyPageFaults(value uint32) (err error) {
	return instance.SetProperty("PageFaults", (value))
}

// GetPageFaults gets the value of PageFaults for the instance
func (instance *MSFT_MTProcess) GetPropertyPageFaults() (value uint32, err error) {
	retValue, err := instance.GetProperty("PageFaults")
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

// SetPeakWorkingSetSize sets the value of PeakWorkingSetSize for the instance
func (instance *MSFT_MTProcess) SetPropertyPeakWorkingSetSize(value uint64) (err error) {
	return instance.SetProperty("PeakWorkingSetSize", (value))
}

// GetPeakWorkingSetSize gets the value of PeakWorkingSetSize for the instance
func (instance *MSFT_MTProcess) GetPropertyPeakWorkingSetSize() (value uint64, err error) {
	retValue, err := instance.GetProperty("PeakWorkingSetSize")
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

// SetPlatform sets the value of Platform for the instance
func (instance *MSFT_MTProcess) SetPropertyPlatform(value uint16) (err error) {
	return instance.SetProperty("Platform", (value))
}

// GetPlatform gets the value of Platform for the instance
func (instance *MSFT_MTProcess) GetPropertyPlatform() (value uint16, err error) {
	retValue, err := instance.GetProperty("Platform")
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

// SetPrivateWorkingSetSize sets the value of PrivateWorkingSetSize for the instance
func (instance *MSFT_MTProcess) SetPropertyPrivateWorkingSetSize(value uint64) (err error) {
	return instance.SetProperty("PrivateWorkingSetSize", (value))
}

// GetPrivateWorkingSetSize gets the value of PrivateWorkingSetSize for the instance
func (instance *MSFT_MTProcess) GetPropertyPrivateWorkingSetSize() (value uint64, err error) {
	retValue, err := instance.GetProperty("PrivateWorkingSetSize")
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

// SetProcessId sets the value of ProcessId for the instance
func (instance *MSFT_MTProcess) SetPropertyProcessId(value uint32) (err error) {
	return instance.SetProperty("ProcessId", (value))
}

// GetProcessId gets the value of ProcessId for the instance
func (instance *MSFT_MTProcess) GetPropertyProcessId() (value uint32, err error) {
	retValue, err := instance.GetProperty("ProcessId")
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

// SetProcessStatus sets the value of ProcessStatus for the instance
func (instance *MSFT_MTProcess) SetPropertyProcessStatus(value uint16) (err error) {
	return instance.SetProperty("ProcessStatus", (value))
}

// GetProcessStatus gets the value of ProcessStatus for the instance
func (instance *MSFT_MTProcess) GetPropertyProcessStatus() (value uint16, err error) {
	retValue, err := instance.GetProperty("ProcessStatus")
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

// SetReadOperationCount sets the value of ReadOperationCount for the instance
func (instance *MSFT_MTProcess) SetPropertyReadOperationCount(value uint64) (err error) {
	return instance.SetProperty("ReadOperationCount", (value))
}

// GetReadOperationCount gets the value of ReadOperationCount for the instance
func (instance *MSFT_MTProcess) GetPropertyReadOperationCount() (value uint64, err error) {
	retValue, err := instance.GetProperty("ReadOperationCount")
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

// SetReadTransferCount sets the value of ReadTransferCount for the instance
func (instance *MSFT_MTProcess) SetPropertyReadTransferCount(value uint64) (err error) {
	return instance.SetProperty("ReadTransferCount", (value))
}

// GetReadTransferCount gets the value of ReadTransferCount for the instance
func (instance *MSFT_MTProcess) GetPropertyReadTransferCount() (value uint64, err error) {
	retValue, err := instance.GetProperty("ReadTransferCount")
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

// SetSessionId sets the value of SessionId for the instance
func (instance *MSFT_MTProcess) SetPropertySessionId(value uint32) (err error) {
	return instance.SetProperty("SessionId", (value))
}

// GetSessionId gets the value of SessionId for the instance
func (instance *MSFT_MTProcess) GetPropertySessionId() (value uint32, err error) {
	retValue, err := instance.GetProperty("SessionId")
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

// SetSharedWorkingSetSize sets the value of SharedWorkingSetSize for the instance
func (instance *MSFT_MTProcess) SetPropertySharedWorkingSetSize(value uint64) (err error) {
	return instance.SetProperty("SharedWorkingSetSize", (value))
}

// GetSharedWorkingSetSize gets the value of SharedWorkingSetSize for the instance
func (instance *MSFT_MTProcess) GetPropertySharedWorkingSetSize() (value uint64, err error) {
	retValue, err := instance.GetProperty("SharedWorkingSetSize")
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

// SetThreadCount sets the value of ThreadCount for the instance
func (instance *MSFT_MTProcess) SetPropertyThreadCount(value uint32) (err error) {
	return instance.SetProperty("ThreadCount", (value))
}

// GetThreadCount gets the value of ThreadCount for the instance
func (instance *MSFT_MTProcess) GetPropertyThreadCount() (value uint32, err error) {
	retValue, err := instance.GetProperty("ThreadCount")
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

// SetUACVirtualization sets the value of UACVirtualization for the instance
func (instance *MSFT_MTProcess) SetPropertyUACVirtualization(value uint16) (err error) {
	return instance.SetProperty("UACVirtualization", (value))
}

// GetUACVirtualization gets the value of UACVirtualization for the instance
func (instance *MSFT_MTProcess) GetPropertyUACVirtualization() (value uint16, err error) {
	retValue, err := instance.GetProperty("UACVirtualization")
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

// SetUserName sets the value of UserName for the instance
func (instance *MSFT_MTProcess) SetPropertyUserName(value string) (err error) {
	return instance.SetProperty("UserName", (value))
}

// GetUserName gets the value of UserName for the instance
func (instance *MSFT_MTProcess) GetPropertyUserName() (value string, err error) {
	retValue, err := instance.GetProperty("UserName")
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

// SetUserObjects sets the value of UserObjects for the instance
func (instance *MSFT_MTProcess) SetPropertyUserObjects(value uint32) (err error) {
	return instance.SetProperty("UserObjects", (value))
}

// GetUserObjects gets the value of UserObjects for the instance
func (instance *MSFT_MTProcess) GetPropertyUserObjects() (value uint32, err error) {
	retValue, err := instance.GetProperty("UserObjects")
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

// SetWorkingSetSize sets the value of WorkingSetSize for the instance
func (instance *MSFT_MTProcess) SetPropertyWorkingSetSize(value uint64) (err error) {
	return instance.SetProperty("WorkingSetSize", (value))
}

// GetWorkingSetSize gets the value of WorkingSetSize for the instance
func (instance *MSFT_MTProcess) GetPropertyWorkingSetSize() (value uint64, err error) {
	retValue, err := instance.GetProperty("WorkingSetSize")
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

// SetWriteOperationCount sets the value of WriteOperationCount for the instance
func (instance *MSFT_MTProcess) SetPropertyWriteOperationCount(value uint64) (err error) {
	return instance.SetProperty("WriteOperationCount", (value))
}

// GetWriteOperationCount gets the value of WriteOperationCount for the instance
func (instance *MSFT_MTProcess) GetPropertyWriteOperationCount() (value uint64, err error) {
	retValue, err := instance.GetProperty("WriteOperationCount")
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

// SetWriteTransferCount sets the value of WriteTransferCount for the instance
func (instance *MSFT_MTProcess) SetPropertyWriteTransferCount(value uint64) (err error) {
	return instance.SetProperty("WriteTransferCount", (value))
}

// GetWriteTransferCount gets the value of WriteTransferCount for the instance
func (instance *MSFT_MTProcess) GetPropertyWriteTransferCount() (value uint64, err error) {
	retValue, err := instance.GetProperty("WriteTransferCount")
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

//

// <param name="DumpFilePath" type="string "></param>
// <param name="ReturnValue" type="uint32 "></param>
func (instance *MSFT_MTProcess) CreateDump( /* OUT */ DumpFilePath string) (result uint32, err error) {
	retVal, err := instance.InvokeMethod("CreateDump")
	if err != nil {
		return
	}
	retValue := retVal[0].(int32)
	result = uint32(retValue)
	return

}

//

// <param name="CommandLine" type="string "></param>
// <param name="WaitMilliseconds" type="uint32 "></param>

// <param name="ActualCommandLine" type="string "></param>
// <param name="ProcessId" type="uint32 "></param>
// <param name="ReturnValue" type="uint32 "></param>
func (instance *MSFT_MTProcess) CreateProcess( /* IN */ CommandLine string,
	/* IN */ WaitMilliseconds uint32,
	/* OUT */ ProcessId uint32,
	/* OUT */ ActualCommandLine string) (result uint32, err error) {
	retVal, err := instance.InvokeMethod("CreateProcess", CommandLine, WaitMilliseconds)
	if err != nil {
		return
	}
	retValue := retVal[0].(int32)
	result = uint32(retValue)
	return

}
