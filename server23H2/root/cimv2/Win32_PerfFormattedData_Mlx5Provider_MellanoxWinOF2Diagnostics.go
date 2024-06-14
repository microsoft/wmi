// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 6/14/2024 using wmigen
//      Source root.CIMV2
//////////////////////////////////////////////
package cimv2

import (
	"github.com/microsoft/wmi/pkg/base/query"
	"github.com/microsoft/wmi/pkg/errors"
	cim "github.com/microsoft/wmi/pkg/wmiinstance"
	"reflect"
)

// Win32_PerfFormattedData_Mlx5Provider_MellanoxWinOF2Diagnostics struct
type Win32_PerfFormattedData_Mlx5Provider_MellanoxWinOF2Diagnostics struct {
	*Win32_PerfFormattedData

	//
	AsyncEQOverrun uint32

	//
	BadChecksumPacketsInSlowPath uint64

	//
	CompletionEQOverrun uint64

	//
	CopiedSendPackets uint64

	//
	CorrectChecksumPacketsInSlowPath uint64

	//
	CPUMEMPages4KMappedByTPTForCQ uint64

	//
	CPUMEMPages4KMappedByTPTForEQ uint64

	//
	CPUMEMPages4KMappedByTPTForMR uint64

	//
	CPUMEMPages4KMappedByTPTForQP uint64

	//
	CQOverflow uint64

	//
	CriticalStallWatermarkReached uint64

	//
	Currentqueuesunderprocessorhandle uint64

	//
	DroplessModeEntries uint64

	//
	DroplessModeExits uint64

	//
	DumpMeNowCalls uint32

	//
	HeadofQueuetimeoutPacketdiscarded uint64

	//
	ImpliedNAKSequenceErrors uint64

	//
	LinkStateChangeDownEvents uint64

	//
	LinkStateChangeEvents uint64

	//
	LocalOperationErrors uint64

	//
	MinorStallWatermarkReached uint64

	//
	MTTEntriesUsedForCQ uint64

	//
	MTTEntriesUsedForEQ uint64

	//
	MTTEntriesUsedForMR uint64

	//
	MTTEntriesUsedForQP uint64

	//
	PacketsReceivedSteeringDropped uint64

	//
	PacketsReceivedWQEtoosmall uint64

	//
	QueuedSendPackets uint64

	//
	ReceiveCompletionsinPassivePerSec uint32

	//
	ReceivedRDMAReadrequests uint64

	//
	ReceivedRDMAWriterequests uint64

	//
	RequesterBadResponse uint64

	//
	RequesterCQEsflushedwitherror uint64

	//
	RequesterCQEswithError uint64

	//
	RequesterLocalLengthErrors uint64

	//
	RequesterLocalProtectionErrors uint64

	//
	RequesterMemoryWindowBindingErrors uint64

	//
	RequesterOutofOrderSequenceNAK uint64

	//
	RequesterQPTransportRetriesExceededErrors uint64

	//
	RequesterRemoteAccessErrors uint64

	//
	RequesterRemoteInvalidRequestErrors uint64

	//
	RequesterRemoteOperationErrors uint64

	//
	RequesterRetryExceededErrors uint64

	//
	RequesterRNRNAK uint64

	//
	RequesterRNRNAKRetriesExceededError uint64

	//
	RequesterTimeoutReceived uint64

	//
	ResetRequests uint64

	//
	ResponderCQEsflushedwitherror uint64

	//
	ResponderCQEswithError uint64

	//
	ResponderDuplicateRequestReceived uint64

	//
	ResponderLocalLengthErrors uint64

	//
	ResponderLocalProtectionErrors uint64

	//
	ResponderLocalQPOperationErrors uint64

	//
	ResponderOutofOrderSequenceReceived uint64

	//
	ResponderRemoteAccessErrors uint64

	//
	ResponderRemoteInvalidRequestErrors uint64

	//
	ResponderRNRNAK uint64

	//
	SendCompletionsinPassivePerSec uint32

	//
	Sendqueuespriority uint64

	//
	StalledStatePacketdiscarded uint64

	//
	Totalqueuesunderprocessorhandle uint64

	//
	TransmissionEngineHangEvents uint64

	//
	UndeterminedChecksumPacketsInSlowPath uint64

	//
	WatchDogExpiredPerSec uint32
}

func NewWin32_PerfFormattedData_Mlx5Provider_MellanoxWinOF2DiagnosticsEx1(instance *cim.WmiInstance) (newInstance *Win32_PerfFormattedData_Mlx5Provider_MellanoxWinOF2Diagnostics, err error) {
	tmp, err := NewWin32_PerfFormattedDataEx1(instance)

	if err != nil {
		return
	}
	newInstance = &Win32_PerfFormattedData_Mlx5Provider_MellanoxWinOF2Diagnostics{
		Win32_PerfFormattedData: tmp,
	}
	return
}

func NewWin32_PerfFormattedData_Mlx5Provider_MellanoxWinOF2DiagnosticsEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *Win32_PerfFormattedData_Mlx5Provider_MellanoxWinOF2Diagnostics, err error) {
	tmp, err := NewWin32_PerfFormattedDataEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &Win32_PerfFormattedData_Mlx5Provider_MellanoxWinOF2Diagnostics{
		Win32_PerfFormattedData: tmp,
	}
	return
}

// SetAsyncEQOverrun sets the value of AsyncEQOverrun for the instance
func (instance *Win32_PerfFormattedData_Mlx5Provider_MellanoxWinOF2Diagnostics) SetPropertyAsyncEQOverrun(value uint32) (err error) {
	return instance.SetProperty("AsyncEQOverrun", (value))
}

// GetAsyncEQOverrun gets the value of AsyncEQOverrun for the instance
func (instance *Win32_PerfFormattedData_Mlx5Provider_MellanoxWinOF2Diagnostics) GetPropertyAsyncEQOverrun() (value uint32, err error) {
	retValue, err := instance.GetProperty("AsyncEQOverrun")
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

// SetBadChecksumPacketsInSlowPath sets the value of BadChecksumPacketsInSlowPath for the instance
func (instance *Win32_PerfFormattedData_Mlx5Provider_MellanoxWinOF2Diagnostics) SetPropertyBadChecksumPacketsInSlowPath(value uint64) (err error) {
	return instance.SetProperty("BadChecksumPacketsInSlowPath", (value))
}

// GetBadChecksumPacketsInSlowPath gets the value of BadChecksumPacketsInSlowPath for the instance
func (instance *Win32_PerfFormattedData_Mlx5Provider_MellanoxWinOF2Diagnostics) GetPropertyBadChecksumPacketsInSlowPath() (value uint64, err error) {
	retValue, err := instance.GetProperty("BadChecksumPacketsInSlowPath")
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

// SetCompletionEQOverrun sets the value of CompletionEQOverrun for the instance
func (instance *Win32_PerfFormattedData_Mlx5Provider_MellanoxWinOF2Diagnostics) SetPropertyCompletionEQOverrun(value uint64) (err error) {
	return instance.SetProperty("CompletionEQOverrun", (value))
}

// GetCompletionEQOverrun gets the value of CompletionEQOverrun for the instance
func (instance *Win32_PerfFormattedData_Mlx5Provider_MellanoxWinOF2Diagnostics) GetPropertyCompletionEQOverrun() (value uint64, err error) {
	retValue, err := instance.GetProperty("CompletionEQOverrun")
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

// SetCopiedSendPackets sets the value of CopiedSendPackets for the instance
func (instance *Win32_PerfFormattedData_Mlx5Provider_MellanoxWinOF2Diagnostics) SetPropertyCopiedSendPackets(value uint64) (err error) {
	return instance.SetProperty("CopiedSendPackets", (value))
}

// GetCopiedSendPackets gets the value of CopiedSendPackets for the instance
func (instance *Win32_PerfFormattedData_Mlx5Provider_MellanoxWinOF2Diagnostics) GetPropertyCopiedSendPackets() (value uint64, err error) {
	retValue, err := instance.GetProperty("CopiedSendPackets")
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

// SetCorrectChecksumPacketsInSlowPath sets the value of CorrectChecksumPacketsInSlowPath for the instance
func (instance *Win32_PerfFormattedData_Mlx5Provider_MellanoxWinOF2Diagnostics) SetPropertyCorrectChecksumPacketsInSlowPath(value uint64) (err error) {
	return instance.SetProperty("CorrectChecksumPacketsInSlowPath", (value))
}

// GetCorrectChecksumPacketsInSlowPath gets the value of CorrectChecksumPacketsInSlowPath for the instance
func (instance *Win32_PerfFormattedData_Mlx5Provider_MellanoxWinOF2Diagnostics) GetPropertyCorrectChecksumPacketsInSlowPath() (value uint64, err error) {
	retValue, err := instance.GetProperty("CorrectChecksumPacketsInSlowPath")
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

// SetCPUMEMPages4KMappedByTPTForCQ sets the value of CPUMEMPages4KMappedByTPTForCQ for the instance
func (instance *Win32_PerfFormattedData_Mlx5Provider_MellanoxWinOF2Diagnostics) SetPropertyCPUMEMPages4KMappedByTPTForCQ(value uint64) (err error) {
	return instance.SetProperty("CPUMEMPages4KMappedByTPTForCQ", (value))
}

// GetCPUMEMPages4KMappedByTPTForCQ gets the value of CPUMEMPages4KMappedByTPTForCQ for the instance
func (instance *Win32_PerfFormattedData_Mlx5Provider_MellanoxWinOF2Diagnostics) GetPropertyCPUMEMPages4KMappedByTPTForCQ() (value uint64, err error) {
	retValue, err := instance.GetProperty("CPUMEMPages4KMappedByTPTForCQ")
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

// SetCPUMEMPages4KMappedByTPTForEQ sets the value of CPUMEMPages4KMappedByTPTForEQ for the instance
func (instance *Win32_PerfFormattedData_Mlx5Provider_MellanoxWinOF2Diagnostics) SetPropertyCPUMEMPages4KMappedByTPTForEQ(value uint64) (err error) {
	return instance.SetProperty("CPUMEMPages4KMappedByTPTForEQ", (value))
}

// GetCPUMEMPages4KMappedByTPTForEQ gets the value of CPUMEMPages4KMappedByTPTForEQ for the instance
func (instance *Win32_PerfFormattedData_Mlx5Provider_MellanoxWinOF2Diagnostics) GetPropertyCPUMEMPages4KMappedByTPTForEQ() (value uint64, err error) {
	retValue, err := instance.GetProperty("CPUMEMPages4KMappedByTPTForEQ")
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

// SetCPUMEMPages4KMappedByTPTForMR sets the value of CPUMEMPages4KMappedByTPTForMR for the instance
func (instance *Win32_PerfFormattedData_Mlx5Provider_MellanoxWinOF2Diagnostics) SetPropertyCPUMEMPages4KMappedByTPTForMR(value uint64) (err error) {
	return instance.SetProperty("CPUMEMPages4KMappedByTPTForMR", (value))
}

// GetCPUMEMPages4KMappedByTPTForMR gets the value of CPUMEMPages4KMappedByTPTForMR for the instance
func (instance *Win32_PerfFormattedData_Mlx5Provider_MellanoxWinOF2Diagnostics) GetPropertyCPUMEMPages4KMappedByTPTForMR() (value uint64, err error) {
	retValue, err := instance.GetProperty("CPUMEMPages4KMappedByTPTForMR")
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

// SetCPUMEMPages4KMappedByTPTForQP sets the value of CPUMEMPages4KMappedByTPTForQP for the instance
func (instance *Win32_PerfFormattedData_Mlx5Provider_MellanoxWinOF2Diagnostics) SetPropertyCPUMEMPages4KMappedByTPTForQP(value uint64) (err error) {
	return instance.SetProperty("CPUMEMPages4KMappedByTPTForQP", (value))
}

// GetCPUMEMPages4KMappedByTPTForQP gets the value of CPUMEMPages4KMappedByTPTForQP for the instance
func (instance *Win32_PerfFormattedData_Mlx5Provider_MellanoxWinOF2Diagnostics) GetPropertyCPUMEMPages4KMappedByTPTForQP() (value uint64, err error) {
	retValue, err := instance.GetProperty("CPUMEMPages4KMappedByTPTForQP")
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

// SetCQOverflow sets the value of CQOverflow for the instance
func (instance *Win32_PerfFormattedData_Mlx5Provider_MellanoxWinOF2Diagnostics) SetPropertyCQOverflow(value uint64) (err error) {
	return instance.SetProperty("CQOverflow", (value))
}

// GetCQOverflow gets the value of CQOverflow for the instance
func (instance *Win32_PerfFormattedData_Mlx5Provider_MellanoxWinOF2Diagnostics) GetPropertyCQOverflow() (value uint64, err error) {
	retValue, err := instance.GetProperty("CQOverflow")
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

// SetCriticalStallWatermarkReached sets the value of CriticalStallWatermarkReached for the instance
func (instance *Win32_PerfFormattedData_Mlx5Provider_MellanoxWinOF2Diagnostics) SetPropertyCriticalStallWatermarkReached(value uint64) (err error) {
	return instance.SetProperty("CriticalStallWatermarkReached", (value))
}

// GetCriticalStallWatermarkReached gets the value of CriticalStallWatermarkReached for the instance
func (instance *Win32_PerfFormattedData_Mlx5Provider_MellanoxWinOF2Diagnostics) GetPropertyCriticalStallWatermarkReached() (value uint64, err error) {
	retValue, err := instance.GetProperty("CriticalStallWatermarkReached")
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

// SetCurrentqueuesunderprocessorhandle sets the value of Currentqueuesunderprocessorhandle for the instance
func (instance *Win32_PerfFormattedData_Mlx5Provider_MellanoxWinOF2Diagnostics) SetPropertyCurrentqueuesunderprocessorhandle(value uint64) (err error) {
	return instance.SetProperty("Currentqueuesunderprocessorhandle", (value))
}

// GetCurrentqueuesunderprocessorhandle gets the value of Currentqueuesunderprocessorhandle for the instance
func (instance *Win32_PerfFormattedData_Mlx5Provider_MellanoxWinOF2Diagnostics) GetPropertyCurrentqueuesunderprocessorhandle() (value uint64, err error) {
	retValue, err := instance.GetProperty("Currentqueuesunderprocessorhandle")
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

// SetDroplessModeEntries sets the value of DroplessModeEntries for the instance
func (instance *Win32_PerfFormattedData_Mlx5Provider_MellanoxWinOF2Diagnostics) SetPropertyDroplessModeEntries(value uint64) (err error) {
	return instance.SetProperty("DroplessModeEntries", (value))
}

// GetDroplessModeEntries gets the value of DroplessModeEntries for the instance
func (instance *Win32_PerfFormattedData_Mlx5Provider_MellanoxWinOF2Diagnostics) GetPropertyDroplessModeEntries() (value uint64, err error) {
	retValue, err := instance.GetProperty("DroplessModeEntries")
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

// SetDroplessModeExits sets the value of DroplessModeExits for the instance
func (instance *Win32_PerfFormattedData_Mlx5Provider_MellanoxWinOF2Diagnostics) SetPropertyDroplessModeExits(value uint64) (err error) {
	return instance.SetProperty("DroplessModeExits", (value))
}

// GetDroplessModeExits gets the value of DroplessModeExits for the instance
func (instance *Win32_PerfFormattedData_Mlx5Provider_MellanoxWinOF2Diagnostics) GetPropertyDroplessModeExits() (value uint64, err error) {
	retValue, err := instance.GetProperty("DroplessModeExits")
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

// SetDumpMeNowCalls sets the value of DumpMeNowCalls for the instance
func (instance *Win32_PerfFormattedData_Mlx5Provider_MellanoxWinOF2Diagnostics) SetPropertyDumpMeNowCalls(value uint32) (err error) {
	return instance.SetProperty("DumpMeNowCalls", (value))
}

// GetDumpMeNowCalls gets the value of DumpMeNowCalls for the instance
func (instance *Win32_PerfFormattedData_Mlx5Provider_MellanoxWinOF2Diagnostics) GetPropertyDumpMeNowCalls() (value uint32, err error) {
	retValue, err := instance.GetProperty("DumpMeNowCalls")
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

// SetHeadofQueuetimeoutPacketdiscarded sets the value of HeadofQueuetimeoutPacketdiscarded for the instance
func (instance *Win32_PerfFormattedData_Mlx5Provider_MellanoxWinOF2Diagnostics) SetPropertyHeadofQueuetimeoutPacketdiscarded(value uint64) (err error) {
	return instance.SetProperty("HeadofQueuetimeoutPacketdiscarded", (value))
}

// GetHeadofQueuetimeoutPacketdiscarded gets the value of HeadofQueuetimeoutPacketdiscarded for the instance
func (instance *Win32_PerfFormattedData_Mlx5Provider_MellanoxWinOF2Diagnostics) GetPropertyHeadofQueuetimeoutPacketdiscarded() (value uint64, err error) {
	retValue, err := instance.GetProperty("HeadofQueuetimeoutPacketdiscarded")
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

// SetImpliedNAKSequenceErrors sets the value of ImpliedNAKSequenceErrors for the instance
func (instance *Win32_PerfFormattedData_Mlx5Provider_MellanoxWinOF2Diagnostics) SetPropertyImpliedNAKSequenceErrors(value uint64) (err error) {
	return instance.SetProperty("ImpliedNAKSequenceErrors", (value))
}

// GetImpliedNAKSequenceErrors gets the value of ImpliedNAKSequenceErrors for the instance
func (instance *Win32_PerfFormattedData_Mlx5Provider_MellanoxWinOF2Diagnostics) GetPropertyImpliedNAKSequenceErrors() (value uint64, err error) {
	retValue, err := instance.GetProperty("ImpliedNAKSequenceErrors")
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

// SetLinkStateChangeDownEvents sets the value of LinkStateChangeDownEvents for the instance
func (instance *Win32_PerfFormattedData_Mlx5Provider_MellanoxWinOF2Diagnostics) SetPropertyLinkStateChangeDownEvents(value uint64) (err error) {
	return instance.SetProperty("LinkStateChangeDownEvents", (value))
}

// GetLinkStateChangeDownEvents gets the value of LinkStateChangeDownEvents for the instance
func (instance *Win32_PerfFormattedData_Mlx5Provider_MellanoxWinOF2Diagnostics) GetPropertyLinkStateChangeDownEvents() (value uint64, err error) {
	retValue, err := instance.GetProperty("LinkStateChangeDownEvents")
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

// SetLinkStateChangeEvents sets the value of LinkStateChangeEvents for the instance
func (instance *Win32_PerfFormattedData_Mlx5Provider_MellanoxWinOF2Diagnostics) SetPropertyLinkStateChangeEvents(value uint64) (err error) {
	return instance.SetProperty("LinkStateChangeEvents", (value))
}

// GetLinkStateChangeEvents gets the value of LinkStateChangeEvents for the instance
func (instance *Win32_PerfFormattedData_Mlx5Provider_MellanoxWinOF2Diagnostics) GetPropertyLinkStateChangeEvents() (value uint64, err error) {
	retValue, err := instance.GetProperty("LinkStateChangeEvents")
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

// SetLocalOperationErrors sets the value of LocalOperationErrors for the instance
func (instance *Win32_PerfFormattedData_Mlx5Provider_MellanoxWinOF2Diagnostics) SetPropertyLocalOperationErrors(value uint64) (err error) {
	return instance.SetProperty("LocalOperationErrors", (value))
}

// GetLocalOperationErrors gets the value of LocalOperationErrors for the instance
func (instance *Win32_PerfFormattedData_Mlx5Provider_MellanoxWinOF2Diagnostics) GetPropertyLocalOperationErrors() (value uint64, err error) {
	retValue, err := instance.GetProperty("LocalOperationErrors")
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

// SetMinorStallWatermarkReached sets the value of MinorStallWatermarkReached for the instance
func (instance *Win32_PerfFormattedData_Mlx5Provider_MellanoxWinOF2Diagnostics) SetPropertyMinorStallWatermarkReached(value uint64) (err error) {
	return instance.SetProperty("MinorStallWatermarkReached", (value))
}

// GetMinorStallWatermarkReached gets the value of MinorStallWatermarkReached for the instance
func (instance *Win32_PerfFormattedData_Mlx5Provider_MellanoxWinOF2Diagnostics) GetPropertyMinorStallWatermarkReached() (value uint64, err error) {
	retValue, err := instance.GetProperty("MinorStallWatermarkReached")
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

// SetMTTEntriesUsedForCQ sets the value of MTTEntriesUsedForCQ for the instance
func (instance *Win32_PerfFormattedData_Mlx5Provider_MellanoxWinOF2Diagnostics) SetPropertyMTTEntriesUsedForCQ(value uint64) (err error) {
	return instance.SetProperty("MTTEntriesUsedForCQ", (value))
}

// GetMTTEntriesUsedForCQ gets the value of MTTEntriesUsedForCQ for the instance
func (instance *Win32_PerfFormattedData_Mlx5Provider_MellanoxWinOF2Diagnostics) GetPropertyMTTEntriesUsedForCQ() (value uint64, err error) {
	retValue, err := instance.GetProperty("MTTEntriesUsedForCQ")
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

// SetMTTEntriesUsedForEQ sets the value of MTTEntriesUsedForEQ for the instance
func (instance *Win32_PerfFormattedData_Mlx5Provider_MellanoxWinOF2Diagnostics) SetPropertyMTTEntriesUsedForEQ(value uint64) (err error) {
	return instance.SetProperty("MTTEntriesUsedForEQ", (value))
}

// GetMTTEntriesUsedForEQ gets the value of MTTEntriesUsedForEQ for the instance
func (instance *Win32_PerfFormattedData_Mlx5Provider_MellanoxWinOF2Diagnostics) GetPropertyMTTEntriesUsedForEQ() (value uint64, err error) {
	retValue, err := instance.GetProperty("MTTEntriesUsedForEQ")
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

// SetMTTEntriesUsedForMR sets the value of MTTEntriesUsedForMR for the instance
func (instance *Win32_PerfFormattedData_Mlx5Provider_MellanoxWinOF2Diagnostics) SetPropertyMTTEntriesUsedForMR(value uint64) (err error) {
	return instance.SetProperty("MTTEntriesUsedForMR", (value))
}

// GetMTTEntriesUsedForMR gets the value of MTTEntriesUsedForMR for the instance
func (instance *Win32_PerfFormattedData_Mlx5Provider_MellanoxWinOF2Diagnostics) GetPropertyMTTEntriesUsedForMR() (value uint64, err error) {
	retValue, err := instance.GetProperty("MTTEntriesUsedForMR")
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

// SetMTTEntriesUsedForQP sets the value of MTTEntriesUsedForQP for the instance
func (instance *Win32_PerfFormattedData_Mlx5Provider_MellanoxWinOF2Diagnostics) SetPropertyMTTEntriesUsedForQP(value uint64) (err error) {
	return instance.SetProperty("MTTEntriesUsedForQP", (value))
}

// GetMTTEntriesUsedForQP gets the value of MTTEntriesUsedForQP for the instance
func (instance *Win32_PerfFormattedData_Mlx5Provider_MellanoxWinOF2Diagnostics) GetPropertyMTTEntriesUsedForQP() (value uint64, err error) {
	retValue, err := instance.GetProperty("MTTEntriesUsedForQP")
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

// SetPacketsReceivedSteeringDropped sets the value of PacketsReceivedSteeringDropped for the instance
func (instance *Win32_PerfFormattedData_Mlx5Provider_MellanoxWinOF2Diagnostics) SetPropertyPacketsReceivedSteeringDropped(value uint64) (err error) {
	return instance.SetProperty("PacketsReceivedSteeringDropped", (value))
}

// GetPacketsReceivedSteeringDropped gets the value of PacketsReceivedSteeringDropped for the instance
func (instance *Win32_PerfFormattedData_Mlx5Provider_MellanoxWinOF2Diagnostics) GetPropertyPacketsReceivedSteeringDropped() (value uint64, err error) {
	retValue, err := instance.GetProperty("PacketsReceivedSteeringDropped")
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

// SetPacketsReceivedWQEtoosmall sets the value of PacketsReceivedWQEtoosmall for the instance
func (instance *Win32_PerfFormattedData_Mlx5Provider_MellanoxWinOF2Diagnostics) SetPropertyPacketsReceivedWQEtoosmall(value uint64) (err error) {
	return instance.SetProperty("PacketsReceivedWQEtoosmall", (value))
}

// GetPacketsReceivedWQEtoosmall gets the value of PacketsReceivedWQEtoosmall for the instance
func (instance *Win32_PerfFormattedData_Mlx5Provider_MellanoxWinOF2Diagnostics) GetPropertyPacketsReceivedWQEtoosmall() (value uint64, err error) {
	retValue, err := instance.GetProperty("PacketsReceivedWQEtoosmall")
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

// SetQueuedSendPackets sets the value of QueuedSendPackets for the instance
func (instance *Win32_PerfFormattedData_Mlx5Provider_MellanoxWinOF2Diagnostics) SetPropertyQueuedSendPackets(value uint64) (err error) {
	return instance.SetProperty("QueuedSendPackets", (value))
}

// GetQueuedSendPackets gets the value of QueuedSendPackets for the instance
func (instance *Win32_PerfFormattedData_Mlx5Provider_MellanoxWinOF2Diagnostics) GetPropertyQueuedSendPackets() (value uint64, err error) {
	retValue, err := instance.GetProperty("QueuedSendPackets")
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

// SetReceiveCompletionsinPassivePerSec sets the value of ReceiveCompletionsinPassivePerSec for the instance
func (instance *Win32_PerfFormattedData_Mlx5Provider_MellanoxWinOF2Diagnostics) SetPropertyReceiveCompletionsinPassivePerSec(value uint32) (err error) {
	return instance.SetProperty("ReceiveCompletionsinPassivePerSec", (value))
}

// GetReceiveCompletionsinPassivePerSec gets the value of ReceiveCompletionsinPassivePerSec for the instance
func (instance *Win32_PerfFormattedData_Mlx5Provider_MellanoxWinOF2Diagnostics) GetPropertyReceiveCompletionsinPassivePerSec() (value uint32, err error) {
	retValue, err := instance.GetProperty("ReceiveCompletionsinPassivePerSec")
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

// SetReceivedRDMAReadrequests sets the value of ReceivedRDMAReadrequests for the instance
func (instance *Win32_PerfFormattedData_Mlx5Provider_MellanoxWinOF2Diagnostics) SetPropertyReceivedRDMAReadrequests(value uint64) (err error) {
	return instance.SetProperty("ReceivedRDMAReadrequests", (value))
}

// GetReceivedRDMAReadrequests gets the value of ReceivedRDMAReadrequests for the instance
func (instance *Win32_PerfFormattedData_Mlx5Provider_MellanoxWinOF2Diagnostics) GetPropertyReceivedRDMAReadrequests() (value uint64, err error) {
	retValue, err := instance.GetProperty("ReceivedRDMAReadrequests")
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

// SetReceivedRDMAWriterequests sets the value of ReceivedRDMAWriterequests for the instance
func (instance *Win32_PerfFormattedData_Mlx5Provider_MellanoxWinOF2Diagnostics) SetPropertyReceivedRDMAWriterequests(value uint64) (err error) {
	return instance.SetProperty("ReceivedRDMAWriterequests", (value))
}

// GetReceivedRDMAWriterequests gets the value of ReceivedRDMAWriterequests for the instance
func (instance *Win32_PerfFormattedData_Mlx5Provider_MellanoxWinOF2Diagnostics) GetPropertyReceivedRDMAWriterequests() (value uint64, err error) {
	retValue, err := instance.GetProperty("ReceivedRDMAWriterequests")
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

// SetRequesterBadResponse sets the value of RequesterBadResponse for the instance
func (instance *Win32_PerfFormattedData_Mlx5Provider_MellanoxWinOF2Diagnostics) SetPropertyRequesterBadResponse(value uint64) (err error) {
	return instance.SetProperty("RequesterBadResponse", (value))
}

// GetRequesterBadResponse gets the value of RequesterBadResponse for the instance
func (instance *Win32_PerfFormattedData_Mlx5Provider_MellanoxWinOF2Diagnostics) GetPropertyRequesterBadResponse() (value uint64, err error) {
	retValue, err := instance.GetProperty("RequesterBadResponse")
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

// SetRequesterCQEsflushedwitherror sets the value of RequesterCQEsflushedwitherror for the instance
func (instance *Win32_PerfFormattedData_Mlx5Provider_MellanoxWinOF2Diagnostics) SetPropertyRequesterCQEsflushedwitherror(value uint64) (err error) {
	return instance.SetProperty("RequesterCQEsflushedwitherror", (value))
}

// GetRequesterCQEsflushedwitherror gets the value of RequesterCQEsflushedwitherror for the instance
func (instance *Win32_PerfFormattedData_Mlx5Provider_MellanoxWinOF2Diagnostics) GetPropertyRequesterCQEsflushedwitherror() (value uint64, err error) {
	retValue, err := instance.GetProperty("RequesterCQEsflushedwitherror")
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

// SetRequesterCQEswithError sets the value of RequesterCQEswithError for the instance
func (instance *Win32_PerfFormattedData_Mlx5Provider_MellanoxWinOF2Diagnostics) SetPropertyRequesterCQEswithError(value uint64) (err error) {
	return instance.SetProperty("RequesterCQEswithError", (value))
}

// GetRequesterCQEswithError gets the value of RequesterCQEswithError for the instance
func (instance *Win32_PerfFormattedData_Mlx5Provider_MellanoxWinOF2Diagnostics) GetPropertyRequesterCQEswithError() (value uint64, err error) {
	retValue, err := instance.GetProperty("RequesterCQEswithError")
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

// SetRequesterLocalLengthErrors sets the value of RequesterLocalLengthErrors for the instance
func (instance *Win32_PerfFormattedData_Mlx5Provider_MellanoxWinOF2Diagnostics) SetPropertyRequesterLocalLengthErrors(value uint64) (err error) {
	return instance.SetProperty("RequesterLocalLengthErrors", (value))
}

// GetRequesterLocalLengthErrors gets the value of RequesterLocalLengthErrors for the instance
func (instance *Win32_PerfFormattedData_Mlx5Provider_MellanoxWinOF2Diagnostics) GetPropertyRequesterLocalLengthErrors() (value uint64, err error) {
	retValue, err := instance.GetProperty("RequesterLocalLengthErrors")
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

// SetRequesterLocalProtectionErrors sets the value of RequesterLocalProtectionErrors for the instance
func (instance *Win32_PerfFormattedData_Mlx5Provider_MellanoxWinOF2Diagnostics) SetPropertyRequesterLocalProtectionErrors(value uint64) (err error) {
	return instance.SetProperty("RequesterLocalProtectionErrors", (value))
}

// GetRequesterLocalProtectionErrors gets the value of RequesterLocalProtectionErrors for the instance
func (instance *Win32_PerfFormattedData_Mlx5Provider_MellanoxWinOF2Diagnostics) GetPropertyRequesterLocalProtectionErrors() (value uint64, err error) {
	retValue, err := instance.GetProperty("RequesterLocalProtectionErrors")
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

// SetRequesterMemoryWindowBindingErrors sets the value of RequesterMemoryWindowBindingErrors for the instance
func (instance *Win32_PerfFormattedData_Mlx5Provider_MellanoxWinOF2Diagnostics) SetPropertyRequesterMemoryWindowBindingErrors(value uint64) (err error) {
	return instance.SetProperty("RequesterMemoryWindowBindingErrors", (value))
}

// GetRequesterMemoryWindowBindingErrors gets the value of RequesterMemoryWindowBindingErrors for the instance
func (instance *Win32_PerfFormattedData_Mlx5Provider_MellanoxWinOF2Diagnostics) GetPropertyRequesterMemoryWindowBindingErrors() (value uint64, err error) {
	retValue, err := instance.GetProperty("RequesterMemoryWindowBindingErrors")
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

// SetRequesterOutofOrderSequenceNAK sets the value of RequesterOutofOrderSequenceNAK for the instance
func (instance *Win32_PerfFormattedData_Mlx5Provider_MellanoxWinOF2Diagnostics) SetPropertyRequesterOutofOrderSequenceNAK(value uint64) (err error) {
	return instance.SetProperty("RequesterOutofOrderSequenceNAK", (value))
}

// GetRequesterOutofOrderSequenceNAK gets the value of RequesterOutofOrderSequenceNAK for the instance
func (instance *Win32_PerfFormattedData_Mlx5Provider_MellanoxWinOF2Diagnostics) GetPropertyRequesterOutofOrderSequenceNAK() (value uint64, err error) {
	retValue, err := instance.GetProperty("RequesterOutofOrderSequenceNAK")
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

// SetRequesterQPTransportRetriesExceededErrors sets the value of RequesterQPTransportRetriesExceededErrors for the instance
func (instance *Win32_PerfFormattedData_Mlx5Provider_MellanoxWinOF2Diagnostics) SetPropertyRequesterQPTransportRetriesExceededErrors(value uint64) (err error) {
	return instance.SetProperty("RequesterQPTransportRetriesExceededErrors", (value))
}

// GetRequesterQPTransportRetriesExceededErrors gets the value of RequesterQPTransportRetriesExceededErrors for the instance
func (instance *Win32_PerfFormattedData_Mlx5Provider_MellanoxWinOF2Diagnostics) GetPropertyRequesterQPTransportRetriesExceededErrors() (value uint64, err error) {
	retValue, err := instance.GetProperty("RequesterQPTransportRetriesExceededErrors")
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

// SetRequesterRemoteAccessErrors sets the value of RequesterRemoteAccessErrors for the instance
func (instance *Win32_PerfFormattedData_Mlx5Provider_MellanoxWinOF2Diagnostics) SetPropertyRequesterRemoteAccessErrors(value uint64) (err error) {
	return instance.SetProperty("RequesterRemoteAccessErrors", (value))
}

// GetRequesterRemoteAccessErrors gets the value of RequesterRemoteAccessErrors for the instance
func (instance *Win32_PerfFormattedData_Mlx5Provider_MellanoxWinOF2Diagnostics) GetPropertyRequesterRemoteAccessErrors() (value uint64, err error) {
	retValue, err := instance.GetProperty("RequesterRemoteAccessErrors")
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

// SetRequesterRemoteInvalidRequestErrors sets the value of RequesterRemoteInvalidRequestErrors for the instance
func (instance *Win32_PerfFormattedData_Mlx5Provider_MellanoxWinOF2Diagnostics) SetPropertyRequesterRemoteInvalidRequestErrors(value uint64) (err error) {
	return instance.SetProperty("RequesterRemoteInvalidRequestErrors", (value))
}

// GetRequesterRemoteInvalidRequestErrors gets the value of RequesterRemoteInvalidRequestErrors for the instance
func (instance *Win32_PerfFormattedData_Mlx5Provider_MellanoxWinOF2Diagnostics) GetPropertyRequesterRemoteInvalidRequestErrors() (value uint64, err error) {
	retValue, err := instance.GetProperty("RequesterRemoteInvalidRequestErrors")
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

// SetRequesterRemoteOperationErrors sets the value of RequesterRemoteOperationErrors for the instance
func (instance *Win32_PerfFormattedData_Mlx5Provider_MellanoxWinOF2Diagnostics) SetPropertyRequesterRemoteOperationErrors(value uint64) (err error) {
	return instance.SetProperty("RequesterRemoteOperationErrors", (value))
}

// GetRequesterRemoteOperationErrors gets the value of RequesterRemoteOperationErrors for the instance
func (instance *Win32_PerfFormattedData_Mlx5Provider_MellanoxWinOF2Diagnostics) GetPropertyRequesterRemoteOperationErrors() (value uint64, err error) {
	retValue, err := instance.GetProperty("RequesterRemoteOperationErrors")
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

// SetRequesterRetryExceededErrors sets the value of RequesterRetryExceededErrors for the instance
func (instance *Win32_PerfFormattedData_Mlx5Provider_MellanoxWinOF2Diagnostics) SetPropertyRequesterRetryExceededErrors(value uint64) (err error) {
	return instance.SetProperty("RequesterRetryExceededErrors", (value))
}

// GetRequesterRetryExceededErrors gets the value of RequesterRetryExceededErrors for the instance
func (instance *Win32_PerfFormattedData_Mlx5Provider_MellanoxWinOF2Diagnostics) GetPropertyRequesterRetryExceededErrors() (value uint64, err error) {
	retValue, err := instance.GetProperty("RequesterRetryExceededErrors")
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

// SetRequesterRNRNAK sets the value of RequesterRNRNAK for the instance
func (instance *Win32_PerfFormattedData_Mlx5Provider_MellanoxWinOF2Diagnostics) SetPropertyRequesterRNRNAK(value uint64) (err error) {
	return instance.SetProperty("RequesterRNRNAK", (value))
}

// GetRequesterRNRNAK gets the value of RequesterRNRNAK for the instance
func (instance *Win32_PerfFormattedData_Mlx5Provider_MellanoxWinOF2Diagnostics) GetPropertyRequesterRNRNAK() (value uint64, err error) {
	retValue, err := instance.GetProperty("RequesterRNRNAK")
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

// SetRequesterRNRNAKRetriesExceededError sets the value of RequesterRNRNAKRetriesExceededError for the instance
func (instance *Win32_PerfFormattedData_Mlx5Provider_MellanoxWinOF2Diagnostics) SetPropertyRequesterRNRNAKRetriesExceededError(value uint64) (err error) {
	return instance.SetProperty("RequesterRNRNAKRetriesExceededError", (value))
}

// GetRequesterRNRNAKRetriesExceededError gets the value of RequesterRNRNAKRetriesExceededError for the instance
func (instance *Win32_PerfFormattedData_Mlx5Provider_MellanoxWinOF2Diagnostics) GetPropertyRequesterRNRNAKRetriesExceededError() (value uint64, err error) {
	retValue, err := instance.GetProperty("RequesterRNRNAKRetriesExceededError")
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

// SetRequesterTimeoutReceived sets the value of RequesterTimeoutReceived for the instance
func (instance *Win32_PerfFormattedData_Mlx5Provider_MellanoxWinOF2Diagnostics) SetPropertyRequesterTimeoutReceived(value uint64) (err error) {
	return instance.SetProperty("RequesterTimeoutReceived", (value))
}

// GetRequesterTimeoutReceived gets the value of RequesterTimeoutReceived for the instance
func (instance *Win32_PerfFormattedData_Mlx5Provider_MellanoxWinOF2Diagnostics) GetPropertyRequesterTimeoutReceived() (value uint64, err error) {
	retValue, err := instance.GetProperty("RequesterTimeoutReceived")
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

// SetResetRequests sets the value of ResetRequests for the instance
func (instance *Win32_PerfFormattedData_Mlx5Provider_MellanoxWinOF2Diagnostics) SetPropertyResetRequests(value uint64) (err error) {
	return instance.SetProperty("ResetRequests", (value))
}

// GetResetRequests gets the value of ResetRequests for the instance
func (instance *Win32_PerfFormattedData_Mlx5Provider_MellanoxWinOF2Diagnostics) GetPropertyResetRequests() (value uint64, err error) {
	retValue, err := instance.GetProperty("ResetRequests")
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

// SetResponderCQEsflushedwitherror sets the value of ResponderCQEsflushedwitherror for the instance
func (instance *Win32_PerfFormattedData_Mlx5Provider_MellanoxWinOF2Diagnostics) SetPropertyResponderCQEsflushedwitherror(value uint64) (err error) {
	return instance.SetProperty("ResponderCQEsflushedwitherror", (value))
}

// GetResponderCQEsflushedwitherror gets the value of ResponderCQEsflushedwitherror for the instance
func (instance *Win32_PerfFormattedData_Mlx5Provider_MellanoxWinOF2Diagnostics) GetPropertyResponderCQEsflushedwitherror() (value uint64, err error) {
	retValue, err := instance.GetProperty("ResponderCQEsflushedwitherror")
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

// SetResponderCQEswithError sets the value of ResponderCQEswithError for the instance
func (instance *Win32_PerfFormattedData_Mlx5Provider_MellanoxWinOF2Diagnostics) SetPropertyResponderCQEswithError(value uint64) (err error) {
	return instance.SetProperty("ResponderCQEswithError", (value))
}

// GetResponderCQEswithError gets the value of ResponderCQEswithError for the instance
func (instance *Win32_PerfFormattedData_Mlx5Provider_MellanoxWinOF2Diagnostics) GetPropertyResponderCQEswithError() (value uint64, err error) {
	retValue, err := instance.GetProperty("ResponderCQEswithError")
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

// SetResponderDuplicateRequestReceived sets the value of ResponderDuplicateRequestReceived for the instance
func (instance *Win32_PerfFormattedData_Mlx5Provider_MellanoxWinOF2Diagnostics) SetPropertyResponderDuplicateRequestReceived(value uint64) (err error) {
	return instance.SetProperty("ResponderDuplicateRequestReceived", (value))
}

// GetResponderDuplicateRequestReceived gets the value of ResponderDuplicateRequestReceived for the instance
func (instance *Win32_PerfFormattedData_Mlx5Provider_MellanoxWinOF2Diagnostics) GetPropertyResponderDuplicateRequestReceived() (value uint64, err error) {
	retValue, err := instance.GetProperty("ResponderDuplicateRequestReceived")
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

// SetResponderLocalLengthErrors sets the value of ResponderLocalLengthErrors for the instance
func (instance *Win32_PerfFormattedData_Mlx5Provider_MellanoxWinOF2Diagnostics) SetPropertyResponderLocalLengthErrors(value uint64) (err error) {
	return instance.SetProperty("ResponderLocalLengthErrors", (value))
}

// GetResponderLocalLengthErrors gets the value of ResponderLocalLengthErrors for the instance
func (instance *Win32_PerfFormattedData_Mlx5Provider_MellanoxWinOF2Diagnostics) GetPropertyResponderLocalLengthErrors() (value uint64, err error) {
	retValue, err := instance.GetProperty("ResponderLocalLengthErrors")
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

// SetResponderLocalProtectionErrors sets the value of ResponderLocalProtectionErrors for the instance
func (instance *Win32_PerfFormattedData_Mlx5Provider_MellanoxWinOF2Diagnostics) SetPropertyResponderLocalProtectionErrors(value uint64) (err error) {
	return instance.SetProperty("ResponderLocalProtectionErrors", (value))
}

// GetResponderLocalProtectionErrors gets the value of ResponderLocalProtectionErrors for the instance
func (instance *Win32_PerfFormattedData_Mlx5Provider_MellanoxWinOF2Diagnostics) GetPropertyResponderLocalProtectionErrors() (value uint64, err error) {
	retValue, err := instance.GetProperty("ResponderLocalProtectionErrors")
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

// SetResponderLocalQPOperationErrors sets the value of ResponderLocalQPOperationErrors for the instance
func (instance *Win32_PerfFormattedData_Mlx5Provider_MellanoxWinOF2Diagnostics) SetPropertyResponderLocalQPOperationErrors(value uint64) (err error) {
	return instance.SetProperty("ResponderLocalQPOperationErrors", (value))
}

// GetResponderLocalQPOperationErrors gets the value of ResponderLocalQPOperationErrors for the instance
func (instance *Win32_PerfFormattedData_Mlx5Provider_MellanoxWinOF2Diagnostics) GetPropertyResponderLocalQPOperationErrors() (value uint64, err error) {
	retValue, err := instance.GetProperty("ResponderLocalQPOperationErrors")
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

// SetResponderOutofOrderSequenceReceived sets the value of ResponderOutofOrderSequenceReceived for the instance
func (instance *Win32_PerfFormattedData_Mlx5Provider_MellanoxWinOF2Diagnostics) SetPropertyResponderOutofOrderSequenceReceived(value uint64) (err error) {
	return instance.SetProperty("ResponderOutofOrderSequenceReceived", (value))
}

// GetResponderOutofOrderSequenceReceived gets the value of ResponderOutofOrderSequenceReceived for the instance
func (instance *Win32_PerfFormattedData_Mlx5Provider_MellanoxWinOF2Diagnostics) GetPropertyResponderOutofOrderSequenceReceived() (value uint64, err error) {
	retValue, err := instance.GetProperty("ResponderOutofOrderSequenceReceived")
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

// SetResponderRemoteAccessErrors sets the value of ResponderRemoteAccessErrors for the instance
func (instance *Win32_PerfFormattedData_Mlx5Provider_MellanoxWinOF2Diagnostics) SetPropertyResponderRemoteAccessErrors(value uint64) (err error) {
	return instance.SetProperty("ResponderRemoteAccessErrors", (value))
}

// GetResponderRemoteAccessErrors gets the value of ResponderRemoteAccessErrors for the instance
func (instance *Win32_PerfFormattedData_Mlx5Provider_MellanoxWinOF2Diagnostics) GetPropertyResponderRemoteAccessErrors() (value uint64, err error) {
	retValue, err := instance.GetProperty("ResponderRemoteAccessErrors")
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

// SetResponderRemoteInvalidRequestErrors sets the value of ResponderRemoteInvalidRequestErrors for the instance
func (instance *Win32_PerfFormattedData_Mlx5Provider_MellanoxWinOF2Diagnostics) SetPropertyResponderRemoteInvalidRequestErrors(value uint64) (err error) {
	return instance.SetProperty("ResponderRemoteInvalidRequestErrors", (value))
}

// GetResponderRemoteInvalidRequestErrors gets the value of ResponderRemoteInvalidRequestErrors for the instance
func (instance *Win32_PerfFormattedData_Mlx5Provider_MellanoxWinOF2Diagnostics) GetPropertyResponderRemoteInvalidRequestErrors() (value uint64, err error) {
	retValue, err := instance.GetProperty("ResponderRemoteInvalidRequestErrors")
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

// SetResponderRNRNAK sets the value of ResponderRNRNAK for the instance
func (instance *Win32_PerfFormattedData_Mlx5Provider_MellanoxWinOF2Diagnostics) SetPropertyResponderRNRNAK(value uint64) (err error) {
	return instance.SetProperty("ResponderRNRNAK", (value))
}

// GetResponderRNRNAK gets the value of ResponderRNRNAK for the instance
func (instance *Win32_PerfFormattedData_Mlx5Provider_MellanoxWinOF2Diagnostics) GetPropertyResponderRNRNAK() (value uint64, err error) {
	retValue, err := instance.GetProperty("ResponderRNRNAK")
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

// SetSendCompletionsinPassivePerSec sets the value of SendCompletionsinPassivePerSec for the instance
func (instance *Win32_PerfFormattedData_Mlx5Provider_MellanoxWinOF2Diagnostics) SetPropertySendCompletionsinPassivePerSec(value uint32) (err error) {
	return instance.SetProperty("SendCompletionsinPassivePerSec", (value))
}

// GetSendCompletionsinPassivePerSec gets the value of SendCompletionsinPassivePerSec for the instance
func (instance *Win32_PerfFormattedData_Mlx5Provider_MellanoxWinOF2Diagnostics) GetPropertySendCompletionsinPassivePerSec() (value uint32, err error) {
	retValue, err := instance.GetProperty("SendCompletionsinPassivePerSec")
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

// SetSendqueuespriority sets the value of Sendqueuespriority for the instance
func (instance *Win32_PerfFormattedData_Mlx5Provider_MellanoxWinOF2Diagnostics) SetPropertySendqueuespriority(value uint64) (err error) {
	return instance.SetProperty("Sendqueuespriority", (value))
}

// GetSendqueuespriority gets the value of Sendqueuespriority for the instance
func (instance *Win32_PerfFormattedData_Mlx5Provider_MellanoxWinOF2Diagnostics) GetPropertySendqueuespriority() (value uint64, err error) {
	retValue, err := instance.GetProperty("Sendqueuespriority")
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

// SetStalledStatePacketdiscarded sets the value of StalledStatePacketdiscarded for the instance
func (instance *Win32_PerfFormattedData_Mlx5Provider_MellanoxWinOF2Diagnostics) SetPropertyStalledStatePacketdiscarded(value uint64) (err error) {
	return instance.SetProperty("StalledStatePacketdiscarded", (value))
}

// GetStalledStatePacketdiscarded gets the value of StalledStatePacketdiscarded for the instance
func (instance *Win32_PerfFormattedData_Mlx5Provider_MellanoxWinOF2Diagnostics) GetPropertyStalledStatePacketdiscarded() (value uint64, err error) {
	retValue, err := instance.GetProperty("StalledStatePacketdiscarded")
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

// SetTotalqueuesunderprocessorhandle sets the value of Totalqueuesunderprocessorhandle for the instance
func (instance *Win32_PerfFormattedData_Mlx5Provider_MellanoxWinOF2Diagnostics) SetPropertyTotalqueuesunderprocessorhandle(value uint64) (err error) {
	return instance.SetProperty("Totalqueuesunderprocessorhandle", (value))
}

// GetTotalqueuesunderprocessorhandle gets the value of Totalqueuesunderprocessorhandle for the instance
func (instance *Win32_PerfFormattedData_Mlx5Provider_MellanoxWinOF2Diagnostics) GetPropertyTotalqueuesunderprocessorhandle() (value uint64, err error) {
	retValue, err := instance.GetProperty("Totalqueuesunderprocessorhandle")
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

// SetTransmissionEngineHangEvents sets the value of TransmissionEngineHangEvents for the instance
func (instance *Win32_PerfFormattedData_Mlx5Provider_MellanoxWinOF2Diagnostics) SetPropertyTransmissionEngineHangEvents(value uint64) (err error) {
	return instance.SetProperty("TransmissionEngineHangEvents", (value))
}

// GetTransmissionEngineHangEvents gets the value of TransmissionEngineHangEvents for the instance
func (instance *Win32_PerfFormattedData_Mlx5Provider_MellanoxWinOF2Diagnostics) GetPropertyTransmissionEngineHangEvents() (value uint64, err error) {
	retValue, err := instance.GetProperty("TransmissionEngineHangEvents")
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

// SetUndeterminedChecksumPacketsInSlowPath sets the value of UndeterminedChecksumPacketsInSlowPath for the instance
func (instance *Win32_PerfFormattedData_Mlx5Provider_MellanoxWinOF2Diagnostics) SetPropertyUndeterminedChecksumPacketsInSlowPath(value uint64) (err error) {
	return instance.SetProperty("UndeterminedChecksumPacketsInSlowPath", (value))
}

// GetUndeterminedChecksumPacketsInSlowPath gets the value of UndeterminedChecksumPacketsInSlowPath for the instance
func (instance *Win32_PerfFormattedData_Mlx5Provider_MellanoxWinOF2Diagnostics) GetPropertyUndeterminedChecksumPacketsInSlowPath() (value uint64, err error) {
	retValue, err := instance.GetProperty("UndeterminedChecksumPacketsInSlowPath")
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

// SetWatchDogExpiredPerSec sets the value of WatchDogExpiredPerSec for the instance
func (instance *Win32_PerfFormattedData_Mlx5Provider_MellanoxWinOF2Diagnostics) SetPropertyWatchDogExpiredPerSec(value uint32) (err error) {
	return instance.SetProperty("WatchDogExpiredPerSec", (value))
}

// GetWatchDogExpiredPerSec gets the value of WatchDogExpiredPerSec for the instance
func (instance *Win32_PerfFormattedData_Mlx5Provider_MellanoxWinOF2Diagnostics) GetPropertyWatchDogExpiredPerSec() (value uint32, err error) {
	retValue, err := instance.GetProperty("WatchDogExpiredPerSec")
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
