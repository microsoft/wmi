// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

// Author:
//
//	Auto Generated on 6/6/2024 using wmigen
//	Source root.CIMV2
//
// ////////////////////////////////////////////
package cimv2

import (
	"github.com/microsoft/wmi/pkg/base/query"
	"github.com/microsoft/wmi/pkg/errors"
	cim "github.com/microsoft/wmi/pkg/wmiinstance"
	"reflect"
)

// Win32_PerfRawData_MSSQLMICROSOFTNumberNumberWID_MSSQLMICROSOFTNumberNumberWIDDatabases struct
type Win32_PerfRawData_MSSQLMICROSOFTNumberNumberWID_MSSQLMICROSOFTNumberNumberWIDDatabases struct {
	*Win32_PerfRawData

	//
	ActiveTransactions uint64

	//
	BackupPerRestoreThroughputPersec uint64

	//
	BulkCopyRowsPersec uint64

	//
	BulkCopyThroughputPersec uint64

	//
	Committableentries uint64

	//
	DataFilesSizeKB uint64

	//
	DBCCLogicalScanBytesPersec uint64

	//
	GroupCommitTimePersec uint64

	//
	LogBytesFlushedPersec uint64

	//
	LogCacheHitRatio uint64

	//
	LogCacheHitRatio_Base uint64

	//
	LogCacheReadsPersec uint64

	//
	LogFilesSizeKB uint64

	//
	LogFilesUsedSizeKB uint64

	//
	LogFlushesPersec uint64

	//
	LogFlushWaitsPersec uint64

	//
	LogFlushWaitTime uint64

	//
	LogFlushWriteTimems uint64

	//
	LogGrowths uint64

	//
	LogPoolCacheMissesPersec uint64

	//
	LogPoolDiskReadsPersec uint64

	//
	LogPoolRequestsPersec uint64

	//
	LogShrinks uint64

	//
	LogTruncations uint64

	//
	PercentLogUsed uint64

	//
	ReplPendingXacts uint64

	//
	ReplTransRate uint64

	//
	ShrinkDataMovementBytesPersec uint64

	//
	TrackedtransactionsPersec uint64

	//
	TransactionsPersec uint64

	//
	WriteTransactionsPersec uint64

	//
	XTPMemoryUsedKB uint64
}

func NewWin32_PerfRawData_MSSQLMICROSOFTNumberNumberWID_MSSQLMICROSOFTNumberNumberWIDDatabasesEx1(instance *cim.WmiInstance) (newInstance *Win32_PerfRawData_MSSQLMICROSOFTNumberNumberWID_MSSQLMICROSOFTNumberNumberWIDDatabases, err error) {
	tmp, err := NewWin32_PerfRawDataEx1(instance)

	if err != nil {
		return
	}
	newInstance = &Win32_PerfRawData_MSSQLMICROSOFTNumberNumberWID_MSSQLMICROSOFTNumberNumberWIDDatabases{
		Win32_PerfRawData: tmp,
	}
	return
}

func NewWin32_PerfRawData_MSSQLMICROSOFTNumberNumberWID_MSSQLMICROSOFTNumberNumberWIDDatabasesEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *Win32_PerfRawData_MSSQLMICROSOFTNumberNumberWID_MSSQLMICROSOFTNumberNumberWIDDatabases, err error) {
	tmp, err := NewWin32_PerfRawDataEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &Win32_PerfRawData_MSSQLMICROSOFTNumberNumberWID_MSSQLMICROSOFTNumberNumberWIDDatabases{
		Win32_PerfRawData: tmp,
	}
	return
}

// SetActiveTransactions sets the value of ActiveTransactions for the instance
func (instance *Win32_PerfRawData_MSSQLMICROSOFTNumberNumberWID_MSSQLMICROSOFTNumberNumberWIDDatabases) SetPropertyActiveTransactions(value uint64) (err error) {
	return instance.SetProperty("ActiveTransactions", (value))
}

// GetActiveTransactions gets the value of ActiveTransactions for the instance
func (instance *Win32_PerfRawData_MSSQLMICROSOFTNumberNumberWID_MSSQLMICROSOFTNumberNumberWIDDatabases) GetPropertyActiveTransactions() (value uint64, err error) {
	retValue, err := instance.GetProperty("ActiveTransactions")
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

// SetBackupPerRestoreThroughputPersec sets the value of BackupPerRestoreThroughputPersec for the instance
func (instance *Win32_PerfRawData_MSSQLMICROSOFTNumberNumberWID_MSSQLMICROSOFTNumberNumberWIDDatabases) SetPropertyBackupPerRestoreThroughputPersec(value uint64) (err error) {
	return instance.SetProperty("BackupPerRestoreThroughputPersec", (value))
}

// GetBackupPerRestoreThroughputPersec gets the value of BackupPerRestoreThroughputPersec for the instance
func (instance *Win32_PerfRawData_MSSQLMICROSOFTNumberNumberWID_MSSQLMICROSOFTNumberNumberWIDDatabases) GetPropertyBackupPerRestoreThroughputPersec() (value uint64, err error) {
	retValue, err := instance.GetProperty("BackupPerRestoreThroughputPersec")
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

// SetBulkCopyRowsPersec sets the value of BulkCopyRowsPersec for the instance
func (instance *Win32_PerfRawData_MSSQLMICROSOFTNumberNumberWID_MSSQLMICROSOFTNumberNumberWIDDatabases) SetPropertyBulkCopyRowsPersec(value uint64) (err error) {
	return instance.SetProperty("BulkCopyRowsPersec", (value))
}

// GetBulkCopyRowsPersec gets the value of BulkCopyRowsPersec for the instance
func (instance *Win32_PerfRawData_MSSQLMICROSOFTNumberNumberWID_MSSQLMICROSOFTNumberNumberWIDDatabases) GetPropertyBulkCopyRowsPersec() (value uint64, err error) {
	retValue, err := instance.GetProperty("BulkCopyRowsPersec")
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

// SetBulkCopyThroughputPersec sets the value of BulkCopyThroughputPersec for the instance
func (instance *Win32_PerfRawData_MSSQLMICROSOFTNumberNumberWID_MSSQLMICROSOFTNumberNumberWIDDatabases) SetPropertyBulkCopyThroughputPersec(value uint64) (err error) {
	return instance.SetProperty("BulkCopyThroughputPersec", (value))
}

// GetBulkCopyThroughputPersec gets the value of BulkCopyThroughputPersec for the instance
func (instance *Win32_PerfRawData_MSSQLMICROSOFTNumberNumberWID_MSSQLMICROSOFTNumberNumberWIDDatabases) GetPropertyBulkCopyThroughputPersec() (value uint64, err error) {
	retValue, err := instance.GetProperty("BulkCopyThroughputPersec")
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

// SetCommittableentries sets the value of Committableentries for the instance
func (instance *Win32_PerfRawData_MSSQLMICROSOFTNumberNumberWID_MSSQLMICROSOFTNumberNumberWIDDatabases) SetPropertyCommittableentries(value uint64) (err error) {
	return instance.SetProperty("Committableentries", (value))
}

// GetCommittableentries gets the value of Committableentries for the instance
func (instance *Win32_PerfRawData_MSSQLMICROSOFTNumberNumberWID_MSSQLMICROSOFTNumberNumberWIDDatabases) GetPropertyCommittableentries() (value uint64, err error) {
	retValue, err := instance.GetProperty("Committableentries")
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

// SetDataFilesSizeKB sets the value of DataFilesSizeKB for the instance
func (instance *Win32_PerfRawData_MSSQLMICROSOFTNumberNumberWID_MSSQLMICROSOFTNumberNumberWIDDatabases) SetPropertyDataFilesSizeKB(value uint64) (err error) {
	return instance.SetProperty("DataFilesSizeKB", (value))
}

// GetDataFilesSizeKB gets the value of DataFilesSizeKB for the instance
func (instance *Win32_PerfRawData_MSSQLMICROSOFTNumberNumberWID_MSSQLMICROSOFTNumberNumberWIDDatabases) GetPropertyDataFilesSizeKB() (value uint64, err error) {
	retValue, err := instance.GetProperty("DataFilesSizeKB")
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

// SetDBCCLogicalScanBytesPersec sets the value of DBCCLogicalScanBytesPersec for the instance
func (instance *Win32_PerfRawData_MSSQLMICROSOFTNumberNumberWID_MSSQLMICROSOFTNumberNumberWIDDatabases) SetPropertyDBCCLogicalScanBytesPersec(value uint64) (err error) {
	return instance.SetProperty("DBCCLogicalScanBytesPersec", (value))
}

// GetDBCCLogicalScanBytesPersec gets the value of DBCCLogicalScanBytesPersec for the instance
func (instance *Win32_PerfRawData_MSSQLMICROSOFTNumberNumberWID_MSSQLMICROSOFTNumberNumberWIDDatabases) GetPropertyDBCCLogicalScanBytesPersec() (value uint64, err error) {
	retValue, err := instance.GetProperty("DBCCLogicalScanBytesPersec")
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

// SetGroupCommitTimePersec sets the value of GroupCommitTimePersec for the instance
func (instance *Win32_PerfRawData_MSSQLMICROSOFTNumberNumberWID_MSSQLMICROSOFTNumberNumberWIDDatabases) SetPropertyGroupCommitTimePersec(value uint64) (err error) {
	return instance.SetProperty("GroupCommitTimePersec", (value))
}

// GetGroupCommitTimePersec gets the value of GroupCommitTimePersec for the instance
func (instance *Win32_PerfRawData_MSSQLMICROSOFTNumberNumberWID_MSSQLMICROSOFTNumberNumberWIDDatabases) GetPropertyGroupCommitTimePersec() (value uint64, err error) {
	retValue, err := instance.GetProperty("GroupCommitTimePersec")
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

// SetLogBytesFlushedPersec sets the value of LogBytesFlushedPersec for the instance
func (instance *Win32_PerfRawData_MSSQLMICROSOFTNumberNumberWID_MSSQLMICROSOFTNumberNumberWIDDatabases) SetPropertyLogBytesFlushedPersec(value uint64) (err error) {
	return instance.SetProperty("LogBytesFlushedPersec", (value))
}

// GetLogBytesFlushedPersec gets the value of LogBytesFlushedPersec for the instance
func (instance *Win32_PerfRawData_MSSQLMICROSOFTNumberNumberWID_MSSQLMICROSOFTNumberNumberWIDDatabases) GetPropertyLogBytesFlushedPersec() (value uint64, err error) {
	retValue, err := instance.GetProperty("LogBytesFlushedPersec")
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

// SetLogCacheHitRatio sets the value of LogCacheHitRatio for the instance
func (instance *Win32_PerfRawData_MSSQLMICROSOFTNumberNumberWID_MSSQLMICROSOFTNumberNumberWIDDatabases) SetPropertyLogCacheHitRatio(value uint64) (err error) {
	return instance.SetProperty("LogCacheHitRatio", (value))
}

// GetLogCacheHitRatio gets the value of LogCacheHitRatio for the instance
func (instance *Win32_PerfRawData_MSSQLMICROSOFTNumberNumberWID_MSSQLMICROSOFTNumberNumberWIDDatabases) GetPropertyLogCacheHitRatio() (value uint64, err error) {
	retValue, err := instance.GetProperty("LogCacheHitRatio")
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

// SetLogCacheHitRatio_Base sets the value of LogCacheHitRatio_Base for the instance
func (instance *Win32_PerfRawData_MSSQLMICROSOFTNumberNumberWID_MSSQLMICROSOFTNumberNumberWIDDatabases) SetPropertyLogCacheHitRatio_Base(value uint64) (err error) {
	return instance.SetProperty("LogCacheHitRatio_Base", (value))
}

// GetLogCacheHitRatio_Base gets the value of LogCacheHitRatio_Base for the instance
func (instance *Win32_PerfRawData_MSSQLMICROSOFTNumberNumberWID_MSSQLMICROSOFTNumberNumberWIDDatabases) GetPropertyLogCacheHitRatio_Base() (value uint64, err error) {
	retValue, err := instance.GetProperty("LogCacheHitRatio_Base")
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

// SetLogCacheReadsPersec sets the value of LogCacheReadsPersec for the instance
func (instance *Win32_PerfRawData_MSSQLMICROSOFTNumberNumberWID_MSSQLMICROSOFTNumberNumberWIDDatabases) SetPropertyLogCacheReadsPersec(value uint64) (err error) {
	return instance.SetProperty("LogCacheReadsPersec", (value))
}

// GetLogCacheReadsPersec gets the value of LogCacheReadsPersec for the instance
func (instance *Win32_PerfRawData_MSSQLMICROSOFTNumberNumberWID_MSSQLMICROSOFTNumberNumberWIDDatabases) GetPropertyLogCacheReadsPersec() (value uint64, err error) {
	retValue, err := instance.GetProperty("LogCacheReadsPersec")
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

// SetLogFilesSizeKB sets the value of LogFilesSizeKB for the instance
func (instance *Win32_PerfRawData_MSSQLMICROSOFTNumberNumberWID_MSSQLMICROSOFTNumberNumberWIDDatabases) SetPropertyLogFilesSizeKB(value uint64) (err error) {
	return instance.SetProperty("LogFilesSizeKB", (value))
}

// GetLogFilesSizeKB gets the value of LogFilesSizeKB for the instance
func (instance *Win32_PerfRawData_MSSQLMICROSOFTNumberNumberWID_MSSQLMICROSOFTNumberNumberWIDDatabases) GetPropertyLogFilesSizeKB() (value uint64, err error) {
	retValue, err := instance.GetProperty("LogFilesSizeKB")
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

// SetLogFilesUsedSizeKB sets the value of LogFilesUsedSizeKB for the instance
func (instance *Win32_PerfRawData_MSSQLMICROSOFTNumberNumberWID_MSSQLMICROSOFTNumberNumberWIDDatabases) SetPropertyLogFilesUsedSizeKB(value uint64) (err error) {
	return instance.SetProperty("LogFilesUsedSizeKB", (value))
}

// GetLogFilesUsedSizeKB gets the value of LogFilesUsedSizeKB for the instance
func (instance *Win32_PerfRawData_MSSQLMICROSOFTNumberNumberWID_MSSQLMICROSOFTNumberNumberWIDDatabases) GetPropertyLogFilesUsedSizeKB() (value uint64, err error) {
	retValue, err := instance.GetProperty("LogFilesUsedSizeKB")
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

// SetLogFlushesPersec sets the value of LogFlushesPersec for the instance
func (instance *Win32_PerfRawData_MSSQLMICROSOFTNumberNumberWID_MSSQLMICROSOFTNumberNumberWIDDatabases) SetPropertyLogFlushesPersec(value uint64) (err error) {
	return instance.SetProperty("LogFlushesPersec", (value))
}

// GetLogFlushesPersec gets the value of LogFlushesPersec for the instance
func (instance *Win32_PerfRawData_MSSQLMICROSOFTNumberNumberWID_MSSQLMICROSOFTNumberNumberWIDDatabases) GetPropertyLogFlushesPersec() (value uint64, err error) {
	retValue, err := instance.GetProperty("LogFlushesPersec")
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

// SetLogFlushWaitsPersec sets the value of LogFlushWaitsPersec for the instance
func (instance *Win32_PerfRawData_MSSQLMICROSOFTNumberNumberWID_MSSQLMICROSOFTNumberNumberWIDDatabases) SetPropertyLogFlushWaitsPersec(value uint64) (err error) {
	return instance.SetProperty("LogFlushWaitsPersec", (value))
}

// GetLogFlushWaitsPersec gets the value of LogFlushWaitsPersec for the instance
func (instance *Win32_PerfRawData_MSSQLMICROSOFTNumberNumberWID_MSSQLMICROSOFTNumberNumberWIDDatabases) GetPropertyLogFlushWaitsPersec() (value uint64, err error) {
	retValue, err := instance.GetProperty("LogFlushWaitsPersec")
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

// SetLogFlushWaitTime sets the value of LogFlushWaitTime for the instance
func (instance *Win32_PerfRawData_MSSQLMICROSOFTNumberNumberWID_MSSQLMICROSOFTNumberNumberWIDDatabases) SetPropertyLogFlushWaitTime(value uint64) (err error) {
	return instance.SetProperty("LogFlushWaitTime", (value))
}

// GetLogFlushWaitTime gets the value of LogFlushWaitTime for the instance
func (instance *Win32_PerfRawData_MSSQLMICROSOFTNumberNumberWID_MSSQLMICROSOFTNumberNumberWIDDatabases) GetPropertyLogFlushWaitTime() (value uint64, err error) {
	retValue, err := instance.GetProperty("LogFlushWaitTime")
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

// SetLogFlushWriteTimems sets the value of LogFlushWriteTimems for the instance
func (instance *Win32_PerfRawData_MSSQLMICROSOFTNumberNumberWID_MSSQLMICROSOFTNumberNumberWIDDatabases) SetPropertyLogFlushWriteTimems(value uint64) (err error) {
	return instance.SetProperty("LogFlushWriteTimems", (value))
}

// GetLogFlushWriteTimems gets the value of LogFlushWriteTimems for the instance
func (instance *Win32_PerfRawData_MSSQLMICROSOFTNumberNumberWID_MSSQLMICROSOFTNumberNumberWIDDatabases) GetPropertyLogFlushWriteTimems() (value uint64, err error) {
	retValue, err := instance.GetProperty("LogFlushWriteTimems")
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

// SetLogGrowths sets the value of LogGrowths for the instance
func (instance *Win32_PerfRawData_MSSQLMICROSOFTNumberNumberWID_MSSQLMICROSOFTNumberNumberWIDDatabases) SetPropertyLogGrowths(value uint64) (err error) {
	return instance.SetProperty("LogGrowths", (value))
}

// GetLogGrowths gets the value of LogGrowths for the instance
func (instance *Win32_PerfRawData_MSSQLMICROSOFTNumberNumberWID_MSSQLMICROSOFTNumberNumberWIDDatabases) GetPropertyLogGrowths() (value uint64, err error) {
	retValue, err := instance.GetProperty("LogGrowths")
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

// SetLogPoolCacheMissesPersec sets the value of LogPoolCacheMissesPersec for the instance
func (instance *Win32_PerfRawData_MSSQLMICROSOFTNumberNumberWID_MSSQLMICROSOFTNumberNumberWIDDatabases) SetPropertyLogPoolCacheMissesPersec(value uint64) (err error) {
	return instance.SetProperty("LogPoolCacheMissesPersec", (value))
}

// GetLogPoolCacheMissesPersec gets the value of LogPoolCacheMissesPersec for the instance
func (instance *Win32_PerfRawData_MSSQLMICROSOFTNumberNumberWID_MSSQLMICROSOFTNumberNumberWIDDatabases) GetPropertyLogPoolCacheMissesPersec() (value uint64, err error) {
	retValue, err := instance.GetProperty("LogPoolCacheMissesPersec")
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

// SetLogPoolDiskReadsPersec sets the value of LogPoolDiskReadsPersec for the instance
func (instance *Win32_PerfRawData_MSSQLMICROSOFTNumberNumberWID_MSSQLMICROSOFTNumberNumberWIDDatabases) SetPropertyLogPoolDiskReadsPersec(value uint64) (err error) {
	return instance.SetProperty("LogPoolDiskReadsPersec", (value))
}

// GetLogPoolDiskReadsPersec gets the value of LogPoolDiskReadsPersec for the instance
func (instance *Win32_PerfRawData_MSSQLMICROSOFTNumberNumberWID_MSSQLMICROSOFTNumberNumberWIDDatabases) GetPropertyLogPoolDiskReadsPersec() (value uint64, err error) {
	retValue, err := instance.GetProperty("LogPoolDiskReadsPersec")
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

// SetLogPoolRequestsPersec sets the value of LogPoolRequestsPersec for the instance
func (instance *Win32_PerfRawData_MSSQLMICROSOFTNumberNumberWID_MSSQLMICROSOFTNumberNumberWIDDatabases) SetPropertyLogPoolRequestsPersec(value uint64) (err error) {
	return instance.SetProperty("LogPoolRequestsPersec", (value))
}

// GetLogPoolRequestsPersec gets the value of LogPoolRequestsPersec for the instance
func (instance *Win32_PerfRawData_MSSQLMICROSOFTNumberNumberWID_MSSQLMICROSOFTNumberNumberWIDDatabases) GetPropertyLogPoolRequestsPersec() (value uint64, err error) {
	retValue, err := instance.GetProperty("LogPoolRequestsPersec")
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

// SetLogShrinks sets the value of LogShrinks for the instance
func (instance *Win32_PerfRawData_MSSQLMICROSOFTNumberNumberWID_MSSQLMICROSOFTNumberNumberWIDDatabases) SetPropertyLogShrinks(value uint64) (err error) {
	return instance.SetProperty("LogShrinks", (value))
}

// GetLogShrinks gets the value of LogShrinks for the instance
func (instance *Win32_PerfRawData_MSSQLMICROSOFTNumberNumberWID_MSSQLMICROSOFTNumberNumberWIDDatabases) GetPropertyLogShrinks() (value uint64, err error) {
	retValue, err := instance.GetProperty("LogShrinks")
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

// SetLogTruncations sets the value of LogTruncations for the instance
func (instance *Win32_PerfRawData_MSSQLMICROSOFTNumberNumberWID_MSSQLMICROSOFTNumberNumberWIDDatabases) SetPropertyLogTruncations(value uint64) (err error) {
	return instance.SetProperty("LogTruncations", (value))
}

// GetLogTruncations gets the value of LogTruncations for the instance
func (instance *Win32_PerfRawData_MSSQLMICROSOFTNumberNumberWID_MSSQLMICROSOFTNumberNumberWIDDatabases) GetPropertyLogTruncations() (value uint64, err error) {
	retValue, err := instance.GetProperty("LogTruncations")
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

// SetPercentLogUsed sets the value of PercentLogUsed for the instance
func (instance *Win32_PerfRawData_MSSQLMICROSOFTNumberNumberWID_MSSQLMICROSOFTNumberNumberWIDDatabases) SetPropertyPercentLogUsed(value uint64) (err error) {
	return instance.SetProperty("PercentLogUsed", (value))
}

// GetPercentLogUsed gets the value of PercentLogUsed for the instance
func (instance *Win32_PerfRawData_MSSQLMICROSOFTNumberNumberWID_MSSQLMICROSOFTNumberNumberWIDDatabases) GetPropertyPercentLogUsed() (value uint64, err error) {
	retValue, err := instance.GetProperty("PercentLogUsed")
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

// SetReplPendingXacts sets the value of ReplPendingXacts for the instance
func (instance *Win32_PerfRawData_MSSQLMICROSOFTNumberNumberWID_MSSQLMICROSOFTNumberNumberWIDDatabases) SetPropertyReplPendingXacts(value uint64) (err error) {
	return instance.SetProperty("ReplPendingXacts", (value))
}

// GetReplPendingXacts gets the value of ReplPendingXacts for the instance
func (instance *Win32_PerfRawData_MSSQLMICROSOFTNumberNumberWID_MSSQLMICROSOFTNumberNumberWIDDatabases) GetPropertyReplPendingXacts() (value uint64, err error) {
	retValue, err := instance.GetProperty("ReplPendingXacts")
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

// SetReplTransRate sets the value of ReplTransRate for the instance
func (instance *Win32_PerfRawData_MSSQLMICROSOFTNumberNumberWID_MSSQLMICROSOFTNumberNumberWIDDatabases) SetPropertyReplTransRate(value uint64) (err error) {
	return instance.SetProperty("ReplTransRate", (value))
}

// GetReplTransRate gets the value of ReplTransRate for the instance
func (instance *Win32_PerfRawData_MSSQLMICROSOFTNumberNumberWID_MSSQLMICROSOFTNumberNumberWIDDatabases) GetPropertyReplTransRate() (value uint64, err error) {
	retValue, err := instance.GetProperty("ReplTransRate")
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

// SetShrinkDataMovementBytesPersec sets the value of ShrinkDataMovementBytesPersec for the instance
func (instance *Win32_PerfRawData_MSSQLMICROSOFTNumberNumberWID_MSSQLMICROSOFTNumberNumberWIDDatabases) SetPropertyShrinkDataMovementBytesPersec(value uint64) (err error) {
	return instance.SetProperty("ShrinkDataMovementBytesPersec", (value))
}

// GetShrinkDataMovementBytesPersec gets the value of ShrinkDataMovementBytesPersec for the instance
func (instance *Win32_PerfRawData_MSSQLMICROSOFTNumberNumberWID_MSSQLMICROSOFTNumberNumberWIDDatabases) GetPropertyShrinkDataMovementBytesPersec() (value uint64, err error) {
	retValue, err := instance.GetProperty("ShrinkDataMovementBytesPersec")
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

// SetTrackedtransactionsPersec sets the value of TrackedtransactionsPersec for the instance
func (instance *Win32_PerfRawData_MSSQLMICROSOFTNumberNumberWID_MSSQLMICROSOFTNumberNumberWIDDatabases) SetPropertyTrackedtransactionsPersec(value uint64) (err error) {
	return instance.SetProperty("TrackedtransactionsPersec", (value))
}

// GetTrackedtransactionsPersec gets the value of TrackedtransactionsPersec for the instance
func (instance *Win32_PerfRawData_MSSQLMICROSOFTNumberNumberWID_MSSQLMICROSOFTNumberNumberWIDDatabases) GetPropertyTrackedtransactionsPersec() (value uint64, err error) {
	retValue, err := instance.GetProperty("TrackedtransactionsPersec")
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

// SetTransactionsPersec sets the value of TransactionsPersec for the instance
func (instance *Win32_PerfRawData_MSSQLMICROSOFTNumberNumberWID_MSSQLMICROSOFTNumberNumberWIDDatabases) SetPropertyTransactionsPersec(value uint64) (err error) {
	return instance.SetProperty("TransactionsPersec", (value))
}

// GetTransactionsPersec gets the value of TransactionsPersec for the instance
func (instance *Win32_PerfRawData_MSSQLMICROSOFTNumberNumberWID_MSSQLMICROSOFTNumberNumberWIDDatabases) GetPropertyTransactionsPersec() (value uint64, err error) {
	retValue, err := instance.GetProperty("TransactionsPersec")
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

// SetWriteTransactionsPersec sets the value of WriteTransactionsPersec for the instance
func (instance *Win32_PerfRawData_MSSQLMICROSOFTNumberNumberWID_MSSQLMICROSOFTNumberNumberWIDDatabases) SetPropertyWriteTransactionsPersec(value uint64) (err error) {
	return instance.SetProperty("WriteTransactionsPersec", (value))
}

// GetWriteTransactionsPersec gets the value of WriteTransactionsPersec for the instance
func (instance *Win32_PerfRawData_MSSQLMICROSOFTNumberNumberWID_MSSQLMICROSOFTNumberNumberWIDDatabases) GetPropertyWriteTransactionsPersec() (value uint64, err error) {
	retValue, err := instance.GetProperty("WriteTransactionsPersec")
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

// SetXTPMemoryUsedKB sets the value of XTPMemoryUsedKB for the instance
func (instance *Win32_PerfRawData_MSSQLMICROSOFTNumberNumberWID_MSSQLMICROSOFTNumberNumberWIDDatabases) SetPropertyXTPMemoryUsedKB(value uint64) (err error) {
	return instance.SetProperty("XTPMemoryUsedKB", (value))
}

// GetXTPMemoryUsedKB gets the value of XTPMemoryUsedKB for the instance
func (instance *Win32_PerfRawData_MSSQLMICROSOFTNumberNumberWID_MSSQLMICROSOFTNumberNumberWIDDatabases) GetPropertyXTPMemoryUsedKB() (value uint64, err error) {
	retValue, err := instance.GetProperty("XTPMemoryUsedKB")
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
