// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 6/13/2024 using wmigen
//      Source root.CIMV2
//////////////////////////////////////////////
package cimv2

import (
	"github.com/microsoft/wmi/pkg/base/query"
	"github.com/microsoft/wmi/pkg/errors"
	cim "github.com/microsoft/wmi/pkg/wmiinstance"
	"reflect"
)

// Win32_PerfRawData_MSSQLMICROSOFTNumberNumberWID_MSSQLMICROSOFTNumberNumberWIDDatabaseReplica struct
type Win32_PerfRawData_MSSQLMICROSOFTNumberNumberWID_MSSQLMICROSOFTNumberNumberWIDDatabaseReplica struct {
	*Win32_PerfRawData

	//
	DatabaseFlowControlDelay uint64

	//
	DatabaseFlowControlsPersec uint64

	//
	FileBytesReceivedPersec uint64

	//
	GroupCommitsPerSec uint64

	//
	GroupCommitTime uint64

	//
	LogBytesCompressedPersec uint64

	//
	LogBytesDecompressedPersec uint64

	//
	LogBytesReceivedPersec uint64

	//
	LogCompressionCachehitsPersec uint64

	//
	LogCompressionCachemissesPersec uint64

	//
	LogCompressionsPersec uint64

	//
	LogDecompressionsPersec uint64

	//
	Logremainingforundo uint64

	//
	LogSendQueue uint64

	//
	MirroredWriteTransactionsPersec uint64

	//
	RecoveryQueue uint64

	//
	RedoblockedPersec uint64

	//
	RedoBytesRemaining uint64

	//
	RedoneBytesPersec uint64

	//
	RedonesPersec uint64

	//
	TotalLogrequiringundo uint64

	//
	TransactionDelay uint64
}

func NewWin32_PerfRawData_MSSQLMICROSOFTNumberNumberWID_MSSQLMICROSOFTNumberNumberWIDDatabaseReplicaEx1(instance *cim.WmiInstance) (newInstance *Win32_PerfRawData_MSSQLMICROSOFTNumberNumberWID_MSSQLMICROSOFTNumberNumberWIDDatabaseReplica, err error) {
	tmp, err := NewWin32_PerfRawDataEx1(instance)

	if err != nil {
		return
	}
	newInstance = &Win32_PerfRawData_MSSQLMICROSOFTNumberNumberWID_MSSQLMICROSOFTNumberNumberWIDDatabaseReplica{
		Win32_PerfRawData: tmp,
	}
	return
}

func NewWin32_PerfRawData_MSSQLMICROSOFTNumberNumberWID_MSSQLMICROSOFTNumberNumberWIDDatabaseReplicaEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *Win32_PerfRawData_MSSQLMICROSOFTNumberNumberWID_MSSQLMICROSOFTNumberNumberWIDDatabaseReplica, err error) {
	tmp, err := NewWin32_PerfRawDataEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &Win32_PerfRawData_MSSQLMICROSOFTNumberNumberWID_MSSQLMICROSOFTNumberNumberWIDDatabaseReplica{
		Win32_PerfRawData: tmp,
	}
	return
}

// SetDatabaseFlowControlDelay sets the value of DatabaseFlowControlDelay for the instance
func (instance *Win32_PerfRawData_MSSQLMICROSOFTNumberNumberWID_MSSQLMICROSOFTNumberNumberWIDDatabaseReplica) SetPropertyDatabaseFlowControlDelay(value uint64) (err error) {
	return instance.SetProperty("DatabaseFlowControlDelay", (value))
}

// GetDatabaseFlowControlDelay gets the value of DatabaseFlowControlDelay for the instance
func (instance *Win32_PerfRawData_MSSQLMICROSOFTNumberNumberWID_MSSQLMICROSOFTNumberNumberWIDDatabaseReplica) GetPropertyDatabaseFlowControlDelay() (value uint64, err error) {
	retValue, err := instance.GetProperty("DatabaseFlowControlDelay")
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

// SetDatabaseFlowControlsPersec sets the value of DatabaseFlowControlsPersec for the instance
func (instance *Win32_PerfRawData_MSSQLMICROSOFTNumberNumberWID_MSSQLMICROSOFTNumberNumberWIDDatabaseReplica) SetPropertyDatabaseFlowControlsPersec(value uint64) (err error) {
	return instance.SetProperty("DatabaseFlowControlsPersec", (value))
}

// GetDatabaseFlowControlsPersec gets the value of DatabaseFlowControlsPersec for the instance
func (instance *Win32_PerfRawData_MSSQLMICROSOFTNumberNumberWID_MSSQLMICROSOFTNumberNumberWIDDatabaseReplica) GetPropertyDatabaseFlowControlsPersec() (value uint64, err error) {
	retValue, err := instance.GetProperty("DatabaseFlowControlsPersec")
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

// SetFileBytesReceivedPersec sets the value of FileBytesReceivedPersec for the instance
func (instance *Win32_PerfRawData_MSSQLMICROSOFTNumberNumberWID_MSSQLMICROSOFTNumberNumberWIDDatabaseReplica) SetPropertyFileBytesReceivedPersec(value uint64) (err error) {
	return instance.SetProperty("FileBytesReceivedPersec", (value))
}

// GetFileBytesReceivedPersec gets the value of FileBytesReceivedPersec for the instance
func (instance *Win32_PerfRawData_MSSQLMICROSOFTNumberNumberWID_MSSQLMICROSOFTNumberNumberWIDDatabaseReplica) GetPropertyFileBytesReceivedPersec() (value uint64, err error) {
	retValue, err := instance.GetProperty("FileBytesReceivedPersec")
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

// SetGroupCommitsPerSec sets the value of GroupCommitsPerSec for the instance
func (instance *Win32_PerfRawData_MSSQLMICROSOFTNumberNumberWID_MSSQLMICROSOFTNumberNumberWIDDatabaseReplica) SetPropertyGroupCommitsPerSec(value uint64) (err error) {
	return instance.SetProperty("GroupCommitsPerSec", (value))
}

// GetGroupCommitsPerSec gets the value of GroupCommitsPerSec for the instance
func (instance *Win32_PerfRawData_MSSQLMICROSOFTNumberNumberWID_MSSQLMICROSOFTNumberNumberWIDDatabaseReplica) GetPropertyGroupCommitsPerSec() (value uint64, err error) {
	retValue, err := instance.GetProperty("GroupCommitsPerSec")
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

// SetGroupCommitTime sets the value of GroupCommitTime for the instance
func (instance *Win32_PerfRawData_MSSQLMICROSOFTNumberNumberWID_MSSQLMICROSOFTNumberNumberWIDDatabaseReplica) SetPropertyGroupCommitTime(value uint64) (err error) {
	return instance.SetProperty("GroupCommitTime", (value))
}

// GetGroupCommitTime gets the value of GroupCommitTime for the instance
func (instance *Win32_PerfRawData_MSSQLMICROSOFTNumberNumberWID_MSSQLMICROSOFTNumberNumberWIDDatabaseReplica) GetPropertyGroupCommitTime() (value uint64, err error) {
	retValue, err := instance.GetProperty("GroupCommitTime")
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

// SetLogBytesCompressedPersec sets the value of LogBytesCompressedPersec for the instance
func (instance *Win32_PerfRawData_MSSQLMICROSOFTNumberNumberWID_MSSQLMICROSOFTNumberNumberWIDDatabaseReplica) SetPropertyLogBytesCompressedPersec(value uint64) (err error) {
	return instance.SetProperty("LogBytesCompressedPersec", (value))
}

// GetLogBytesCompressedPersec gets the value of LogBytesCompressedPersec for the instance
func (instance *Win32_PerfRawData_MSSQLMICROSOFTNumberNumberWID_MSSQLMICROSOFTNumberNumberWIDDatabaseReplica) GetPropertyLogBytesCompressedPersec() (value uint64, err error) {
	retValue, err := instance.GetProperty("LogBytesCompressedPersec")
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

// SetLogBytesDecompressedPersec sets the value of LogBytesDecompressedPersec for the instance
func (instance *Win32_PerfRawData_MSSQLMICROSOFTNumberNumberWID_MSSQLMICROSOFTNumberNumberWIDDatabaseReplica) SetPropertyLogBytesDecompressedPersec(value uint64) (err error) {
	return instance.SetProperty("LogBytesDecompressedPersec", (value))
}

// GetLogBytesDecompressedPersec gets the value of LogBytesDecompressedPersec for the instance
func (instance *Win32_PerfRawData_MSSQLMICROSOFTNumberNumberWID_MSSQLMICROSOFTNumberNumberWIDDatabaseReplica) GetPropertyLogBytesDecompressedPersec() (value uint64, err error) {
	retValue, err := instance.GetProperty("LogBytesDecompressedPersec")
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

// SetLogBytesReceivedPersec sets the value of LogBytesReceivedPersec for the instance
func (instance *Win32_PerfRawData_MSSQLMICROSOFTNumberNumberWID_MSSQLMICROSOFTNumberNumberWIDDatabaseReplica) SetPropertyLogBytesReceivedPersec(value uint64) (err error) {
	return instance.SetProperty("LogBytesReceivedPersec", (value))
}

// GetLogBytesReceivedPersec gets the value of LogBytesReceivedPersec for the instance
func (instance *Win32_PerfRawData_MSSQLMICROSOFTNumberNumberWID_MSSQLMICROSOFTNumberNumberWIDDatabaseReplica) GetPropertyLogBytesReceivedPersec() (value uint64, err error) {
	retValue, err := instance.GetProperty("LogBytesReceivedPersec")
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

// SetLogCompressionCachehitsPersec sets the value of LogCompressionCachehitsPersec for the instance
func (instance *Win32_PerfRawData_MSSQLMICROSOFTNumberNumberWID_MSSQLMICROSOFTNumberNumberWIDDatabaseReplica) SetPropertyLogCompressionCachehitsPersec(value uint64) (err error) {
	return instance.SetProperty("LogCompressionCachehitsPersec", (value))
}

// GetLogCompressionCachehitsPersec gets the value of LogCompressionCachehitsPersec for the instance
func (instance *Win32_PerfRawData_MSSQLMICROSOFTNumberNumberWID_MSSQLMICROSOFTNumberNumberWIDDatabaseReplica) GetPropertyLogCompressionCachehitsPersec() (value uint64, err error) {
	retValue, err := instance.GetProperty("LogCompressionCachehitsPersec")
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

// SetLogCompressionCachemissesPersec sets the value of LogCompressionCachemissesPersec for the instance
func (instance *Win32_PerfRawData_MSSQLMICROSOFTNumberNumberWID_MSSQLMICROSOFTNumberNumberWIDDatabaseReplica) SetPropertyLogCompressionCachemissesPersec(value uint64) (err error) {
	return instance.SetProperty("LogCompressionCachemissesPersec", (value))
}

// GetLogCompressionCachemissesPersec gets the value of LogCompressionCachemissesPersec for the instance
func (instance *Win32_PerfRawData_MSSQLMICROSOFTNumberNumberWID_MSSQLMICROSOFTNumberNumberWIDDatabaseReplica) GetPropertyLogCompressionCachemissesPersec() (value uint64, err error) {
	retValue, err := instance.GetProperty("LogCompressionCachemissesPersec")
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

// SetLogCompressionsPersec sets the value of LogCompressionsPersec for the instance
func (instance *Win32_PerfRawData_MSSQLMICROSOFTNumberNumberWID_MSSQLMICROSOFTNumberNumberWIDDatabaseReplica) SetPropertyLogCompressionsPersec(value uint64) (err error) {
	return instance.SetProperty("LogCompressionsPersec", (value))
}

// GetLogCompressionsPersec gets the value of LogCompressionsPersec for the instance
func (instance *Win32_PerfRawData_MSSQLMICROSOFTNumberNumberWID_MSSQLMICROSOFTNumberNumberWIDDatabaseReplica) GetPropertyLogCompressionsPersec() (value uint64, err error) {
	retValue, err := instance.GetProperty("LogCompressionsPersec")
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

// SetLogDecompressionsPersec sets the value of LogDecompressionsPersec for the instance
func (instance *Win32_PerfRawData_MSSQLMICROSOFTNumberNumberWID_MSSQLMICROSOFTNumberNumberWIDDatabaseReplica) SetPropertyLogDecompressionsPersec(value uint64) (err error) {
	return instance.SetProperty("LogDecompressionsPersec", (value))
}

// GetLogDecompressionsPersec gets the value of LogDecompressionsPersec for the instance
func (instance *Win32_PerfRawData_MSSQLMICROSOFTNumberNumberWID_MSSQLMICROSOFTNumberNumberWIDDatabaseReplica) GetPropertyLogDecompressionsPersec() (value uint64, err error) {
	retValue, err := instance.GetProperty("LogDecompressionsPersec")
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

// SetLogremainingforundo sets the value of Logremainingforundo for the instance
func (instance *Win32_PerfRawData_MSSQLMICROSOFTNumberNumberWID_MSSQLMICROSOFTNumberNumberWIDDatabaseReplica) SetPropertyLogremainingforundo(value uint64) (err error) {
	return instance.SetProperty("Logremainingforundo", (value))
}

// GetLogremainingforundo gets the value of Logremainingforundo for the instance
func (instance *Win32_PerfRawData_MSSQLMICROSOFTNumberNumberWID_MSSQLMICROSOFTNumberNumberWIDDatabaseReplica) GetPropertyLogremainingforundo() (value uint64, err error) {
	retValue, err := instance.GetProperty("Logremainingforundo")
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

// SetLogSendQueue sets the value of LogSendQueue for the instance
func (instance *Win32_PerfRawData_MSSQLMICROSOFTNumberNumberWID_MSSQLMICROSOFTNumberNumberWIDDatabaseReplica) SetPropertyLogSendQueue(value uint64) (err error) {
	return instance.SetProperty("LogSendQueue", (value))
}

// GetLogSendQueue gets the value of LogSendQueue for the instance
func (instance *Win32_PerfRawData_MSSQLMICROSOFTNumberNumberWID_MSSQLMICROSOFTNumberNumberWIDDatabaseReplica) GetPropertyLogSendQueue() (value uint64, err error) {
	retValue, err := instance.GetProperty("LogSendQueue")
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

// SetMirroredWriteTransactionsPersec sets the value of MirroredWriteTransactionsPersec for the instance
func (instance *Win32_PerfRawData_MSSQLMICROSOFTNumberNumberWID_MSSQLMICROSOFTNumberNumberWIDDatabaseReplica) SetPropertyMirroredWriteTransactionsPersec(value uint64) (err error) {
	return instance.SetProperty("MirroredWriteTransactionsPersec", (value))
}

// GetMirroredWriteTransactionsPersec gets the value of MirroredWriteTransactionsPersec for the instance
func (instance *Win32_PerfRawData_MSSQLMICROSOFTNumberNumberWID_MSSQLMICROSOFTNumberNumberWIDDatabaseReplica) GetPropertyMirroredWriteTransactionsPersec() (value uint64, err error) {
	retValue, err := instance.GetProperty("MirroredWriteTransactionsPersec")
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

// SetRecoveryQueue sets the value of RecoveryQueue for the instance
func (instance *Win32_PerfRawData_MSSQLMICROSOFTNumberNumberWID_MSSQLMICROSOFTNumberNumberWIDDatabaseReplica) SetPropertyRecoveryQueue(value uint64) (err error) {
	return instance.SetProperty("RecoveryQueue", (value))
}

// GetRecoveryQueue gets the value of RecoveryQueue for the instance
func (instance *Win32_PerfRawData_MSSQLMICROSOFTNumberNumberWID_MSSQLMICROSOFTNumberNumberWIDDatabaseReplica) GetPropertyRecoveryQueue() (value uint64, err error) {
	retValue, err := instance.GetProperty("RecoveryQueue")
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

// SetRedoblockedPersec sets the value of RedoblockedPersec for the instance
func (instance *Win32_PerfRawData_MSSQLMICROSOFTNumberNumberWID_MSSQLMICROSOFTNumberNumberWIDDatabaseReplica) SetPropertyRedoblockedPersec(value uint64) (err error) {
	return instance.SetProperty("RedoblockedPersec", (value))
}

// GetRedoblockedPersec gets the value of RedoblockedPersec for the instance
func (instance *Win32_PerfRawData_MSSQLMICROSOFTNumberNumberWID_MSSQLMICROSOFTNumberNumberWIDDatabaseReplica) GetPropertyRedoblockedPersec() (value uint64, err error) {
	retValue, err := instance.GetProperty("RedoblockedPersec")
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

// SetRedoBytesRemaining sets the value of RedoBytesRemaining for the instance
func (instance *Win32_PerfRawData_MSSQLMICROSOFTNumberNumberWID_MSSQLMICROSOFTNumberNumberWIDDatabaseReplica) SetPropertyRedoBytesRemaining(value uint64) (err error) {
	return instance.SetProperty("RedoBytesRemaining", (value))
}

// GetRedoBytesRemaining gets the value of RedoBytesRemaining for the instance
func (instance *Win32_PerfRawData_MSSQLMICROSOFTNumberNumberWID_MSSQLMICROSOFTNumberNumberWIDDatabaseReplica) GetPropertyRedoBytesRemaining() (value uint64, err error) {
	retValue, err := instance.GetProperty("RedoBytesRemaining")
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

// SetRedoneBytesPersec sets the value of RedoneBytesPersec for the instance
func (instance *Win32_PerfRawData_MSSQLMICROSOFTNumberNumberWID_MSSQLMICROSOFTNumberNumberWIDDatabaseReplica) SetPropertyRedoneBytesPersec(value uint64) (err error) {
	return instance.SetProperty("RedoneBytesPersec", (value))
}

// GetRedoneBytesPersec gets the value of RedoneBytesPersec for the instance
func (instance *Win32_PerfRawData_MSSQLMICROSOFTNumberNumberWID_MSSQLMICROSOFTNumberNumberWIDDatabaseReplica) GetPropertyRedoneBytesPersec() (value uint64, err error) {
	retValue, err := instance.GetProperty("RedoneBytesPersec")
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

// SetRedonesPersec sets the value of RedonesPersec for the instance
func (instance *Win32_PerfRawData_MSSQLMICROSOFTNumberNumberWID_MSSQLMICROSOFTNumberNumberWIDDatabaseReplica) SetPropertyRedonesPersec(value uint64) (err error) {
	return instance.SetProperty("RedonesPersec", (value))
}

// GetRedonesPersec gets the value of RedonesPersec for the instance
func (instance *Win32_PerfRawData_MSSQLMICROSOFTNumberNumberWID_MSSQLMICROSOFTNumberNumberWIDDatabaseReplica) GetPropertyRedonesPersec() (value uint64, err error) {
	retValue, err := instance.GetProperty("RedonesPersec")
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

// SetTotalLogrequiringundo sets the value of TotalLogrequiringundo for the instance
func (instance *Win32_PerfRawData_MSSQLMICROSOFTNumberNumberWID_MSSQLMICROSOFTNumberNumberWIDDatabaseReplica) SetPropertyTotalLogrequiringundo(value uint64) (err error) {
	return instance.SetProperty("TotalLogrequiringundo", (value))
}

// GetTotalLogrequiringundo gets the value of TotalLogrequiringundo for the instance
func (instance *Win32_PerfRawData_MSSQLMICROSOFTNumberNumberWID_MSSQLMICROSOFTNumberNumberWIDDatabaseReplica) GetPropertyTotalLogrequiringundo() (value uint64, err error) {
	retValue, err := instance.GetProperty("TotalLogrequiringundo")
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

// SetTransactionDelay sets the value of TransactionDelay for the instance
func (instance *Win32_PerfRawData_MSSQLMICROSOFTNumberNumberWID_MSSQLMICROSOFTNumberNumberWIDDatabaseReplica) SetPropertyTransactionDelay(value uint64) (err error) {
	return instance.SetProperty("TransactionDelay", (value))
}

// GetTransactionDelay gets the value of TransactionDelay for the instance
func (instance *Win32_PerfRawData_MSSQLMICROSOFTNumberNumberWID_MSSQLMICROSOFTNumberNumberWIDDatabaseReplica) GetPropertyTransactionDelay() (value uint64, err error) {
	retValue, err := instance.GetProperty("TransactionDelay")
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
