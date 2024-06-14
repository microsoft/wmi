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

// Win32_PerfFormattedData_MSSQLMICROSOFTNumberNumberWID_MSSQLMICROSOFTNumberNumberWIDDatabaseMirroring struct
type Win32_PerfFormattedData_MSSQLMICROSOFTNumberNumberWID_MSSQLMICROSOFTNumberNumberWIDDatabaseMirroring struct {
	*Win32_PerfFormattedData

	//
	BytesReceivedPersec uint64

	//
	BytesSentPersec uint64

	//
	LogBytesReceivedPersec uint64

	//
	LogBytesRedonefromCachePersec uint64

	//
	LogBytesSentfromCachePersec uint64

	//
	LogBytesSentPersec uint64

	//
	LogCompressedBytesRcvdPersec uint64

	//
	LogCompressedBytesSentPersec uint64

	//
	LogHardenTimems uint64

	//
	LogRemainingforUndoKB uint64

	//
	LogScannedforUndoKB uint64

	//
	LogSendFlowControlTimems uint64

	//
	LogSendQueueKB uint64

	//
	MirroredWriteTransactionsPersec uint64

	//
	PagesSentPersec uint64

	//
	ReceivesPersec uint64

	//
	RedoBytesPersec uint64

	//
	RedoQueueKB uint64

	//
	SendPerReceiveAckTime uint64

	//
	SendsPersec uint64

	//
	TransactionDelay uint64
}

func NewWin32_PerfFormattedData_MSSQLMICROSOFTNumberNumberWID_MSSQLMICROSOFTNumberNumberWIDDatabaseMirroringEx1(instance *cim.WmiInstance) (newInstance *Win32_PerfFormattedData_MSSQLMICROSOFTNumberNumberWID_MSSQLMICROSOFTNumberNumberWIDDatabaseMirroring, err error) {
	tmp, err := NewWin32_PerfFormattedDataEx1(instance)

	if err != nil {
		return
	}
	newInstance = &Win32_PerfFormattedData_MSSQLMICROSOFTNumberNumberWID_MSSQLMICROSOFTNumberNumberWIDDatabaseMirroring{
		Win32_PerfFormattedData: tmp,
	}
	return
}

func NewWin32_PerfFormattedData_MSSQLMICROSOFTNumberNumberWID_MSSQLMICROSOFTNumberNumberWIDDatabaseMirroringEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *Win32_PerfFormattedData_MSSQLMICROSOFTNumberNumberWID_MSSQLMICROSOFTNumberNumberWIDDatabaseMirroring, err error) {
	tmp, err := NewWin32_PerfFormattedDataEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &Win32_PerfFormattedData_MSSQLMICROSOFTNumberNumberWID_MSSQLMICROSOFTNumberNumberWIDDatabaseMirroring{
		Win32_PerfFormattedData: tmp,
	}
	return
}

// SetBytesReceivedPersec sets the value of BytesReceivedPersec for the instance
func (instance *Win32_PerfFormattedData_MSSQLMICROSOFTNumberNumberWID_MSSQLMICROSOFTNumberNumberWIDDatabaseMirroring) SetPropertyBytesReceivedPersec(value uint64) (err error) {
	return instance.SetProperty("BytesReceivedPersec", (value))
}

// GetBytesReceivedPersec gets the value of BytesReceivedPersec for the instance
func (instance *Win32_PerfFormattedData_MSSQLMICROSOFTNumberNumberWID_MSSQLMICROSOFTNumberNumberWIDDatabaseMirroring) GetPropertyBytesReceivedPersec() (value uint64, err error) {
	retValue, err := instance.GetProperty("BytesReceivedPersec")
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

// SetBytesSentPersec sets the value of BytesSentPersec for the instance
func (instance *Win32_PerfFormattedData_MSSQLMICROSOFTNumberNumberWID_MSSQLMICROSOFTNumberNumberWIDDatabaseMirroring) SetPropertyBytesSentPersec(value uint64) (err error) {
	return instance.SetProperty("BytesSentPersec", (value))
}

// GetBytesSentPersec gets the value of BytesSentPersec for the instance
func (instance *Win32_PerfFormattedData_MSSQLMICROSOFTNumberNumberWID_MSSQLMICROSOFTNumberNumberWIDDatabaseMirroring) GetPropertyBytesSentPersec() (value uint64, err error) {
	retValue, err := instance.GetProperty("BytesSentPersec")
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
func (instance *Win32_PerfFormattedData_MSSQLMICROSOFTNumberNumberWID_MSSQLMICROSOFTNumberNumberWIDDatabaseMirroring) SetPropertyLogBytesReceivedPersec(value uint64) (err error) {
	return instance.SetProperty("LogBytesReceivedPersec", (value))
}

// GetLogBytesReceivedPersec gets the value of LogBytesReceivedPersec for the instance
func (instance *Win32_PerfFormattedData_MSSQLMICROSOFTNumberNumberWID_MSSQLMICROSOFTNumberNumberWIDDatabaseMirroring) GetPropertyLogBytesReceivedPersec() (value uint64, err error) {
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

// SetLogBytesRedonefromCachePersec sets the value of LogBytesRedonefromCachePersec for the instance
func (instance *Win32_PerfFormattedData_MSSQLMICROSOFTNumberNumberWID_MSSQLMICROSOFTNumberNumberWIDDatabaseMirroring) SetPropertyLogBytesRedonefromCachePersec(value uint64) (err error) {
	return instance.SetProperty("LogBytesRedonefromCachePersec", (value))
}

// GetLogBytesRedonefromCachePersec gets the value of LogBytesRedonefromCachePersec for the instance
func (instance *Win32_PerfFormattedData_MSSQLMICROSOFTNumberNumberWID_MSSQLMICROSOFTNumberNumberWIDDatabaseMirroring) GetPropertyLogBytesRedonefromCachePersec() (value uint64, err error) {
	retValue, err := instance.GetProperty("LogBytesRedonefromCachePersec")
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

// SetLogBytesSentfromCachePersec sets the value of LogBytesSentfromCachePersec for the instance
func (instance *Win32_PerfFormattedData_MSSQLMICROSOFTNumberNumberWID_MSSQLMICROSOFTNumberNumberWIDDatabaseMirroring) SetPropertyLogBytesSentfromCachePersec(value uint64) (err error) {
	return instance.SetProperty("LogBytesSentfromCachePersec", (value))
}

// GetLogBytesSentfromCachePersec gets the value of LogBytesSentfromCachePersec for the instance
func (instance *Win32_PerfFormattedData_MSSQLMICROSOFTNumberNumberWID_MSSQLMICROSOFTNumberNumberWIDDatabaseMirroring) GetPropertyLogBytesSentfromCachePersec() (value uint64, err error) {
	retValue, err := instance.GetProperty("LogBytesSentfromCachePersec")
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

// SetLogBytesSentPersec sets the value of LogBytesSentPersec for the instance
func (instance *Win32_PerfFormattedData_MSSQLMICROSOFTNumberNumberWID_MSSQLMICROSOFTNumberNumberWIDDatabaseMirroring) SetPropertyLogBytesSentPersec(value uint64) (err error) {
	return instance.SetProperty("LogBytesSentPersec", (value))
}

// GetLogBytesSentPersec gets the value of LogBytesSentPersec for the instance
func (instance *Win32_PerfFormattedData_MSSQLMICROSOFTNumberNumberWID_MSSQLMICROSOFTNumberNumberWIDDatabaseMirroring) GetPropertyLogBytesSentPersec() (value uint64, err error) {
	retValue, err := instance.GetProperty("LogBytesSentPersec")
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

// SetLogCompressedBytesRcvdPersec sets the value of LogCompressedBytesRcvdPersec for the instance
func (instance *Win32_PerfFormattedData_MSSQLMICROSOFTNumberNumberWID_MSSQLMICROSOFTNumberNumberWIDDatabaseMirroring) SetPropertyLogCompressedBytesRcvdPersec(value uint64) (err error) {
	return instance.SetProperty("LogCompressedBytesRcvdPersec", (value))
}

// GetLogCompressedBytesRcvdPersec gets the value of LogCompressedBytesRcvdPersec for the instance
func (instance *Win32_PerfFormattedData_MSSQLMICROSOFTNumberNumberWID_MSSQLMICROSOFTNumberNumberWIDDatabaseMirroring) GetPropertyLogCompressedBytesRcvdPersec() (value uint64, err error) {
	retValue, err := instance.GetProperty("LogCompressedBytesRcvdPersec")
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

// SetLogCompressedBytesSentPersec sets the value of LogCompressedBytesSentPersec for the instance
func (instance *Win32_PerfFormattedData_MSSQLMICROSOFTNumberNumberWID_MSSQLMICROSOFTNumberNumberWIDDatabaseMirroring) SetPropertyLogCompressedBytesSentPersec(value uint64) (err error) {
	return instance.SetProperty("LogCompressedBytesSentPersec", (value))
}

// GetLogCompressedBytesSentPersec gets the value of LogCompressedBytesSentPersec for the instance
func (instance *Win32_PerfFormattedData_MSSQLMICROSOFTNumberNumberWID_MSSQLMICROSOFTNumberNumberWIDDatabaseMirroring) GetPropertyLogCompressedBytesSentPersec() (value uint64, err error) {
	retValue, err := instance.GetProperty("LogCompressedBytesSentPersec")
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

// SetLogHardenTimems sets the value of LogHardenTimems for the instance
func (instance *Win32_PerfFormattedData_MSSQLMICROSOFTNumberNumberWID_MSSQLMICROSOFTNumberNumberWIDDatabaseMirroring) SetPropertyLogHardenTimems(value uint64) (err error) {
	return instance.SetProperty("LogHardenTimems", (value))
}

// GetLogHardenTimems gets the value of LogHardenTimems for the instance
func (instance *Win32_PerfFormattedData_MSSQLMICROSOFTNumberNumberWID_MSSQLMICROSOFTNumberNumberWIDDatabaseMirroring) GetPropertyLogHardenTimems() (value uint64, err error) {
	retValue, err := instance.GetProperty("LogHardenTimems")
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

// SetLogRemainingforUndoKB sets the value of LogRemainingforUndoKB for the instance
func (instance *Win32_PerfFormattedData_MSSQLMICROSOFTNumberNumberWID_MSSQLMICROSOFTNumberNumberWIDDatabaseMirroring) SetPropertyLogRemainingforUndoKB(value uint64) (err error) {
	return instance.SetProperty("LogRemainingforUndoKB", (value))
}

// GetLogRemainingforUndoKB gets the value of LogRemainingforUndoKB for the instance
func (instance *Win32_PerfFormattedData_MSSQLMICROSOFTNumberNumberWID_MSSQLMICROSOFTNumberNumberWIDDatabaseMirroring) GetPropertyLogRemainingforUndoKB() (value uint64, err error) {
	retValue, err := instance.GetProperty("LogRemainingforUndoKB")
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

// SetLogScannedforUndoKB sets the value of LogScannedforUndoKB for the instance
func (instance *Win32_PerfFormattedData_MSSQLMICROSOFTNumberNumberWID_MSSQLMICROSOFTNumberNumberWIDDatabaseMirroring) SetPropertyLogScannedforUndoKB(value uint64) (err error) {
	return instance.SetProperty("LogScannedforUndoKB", (value))
}

// GetLogScannedforUndoKB gets the value of LogScannedforUndoKB for the instance
func (instance *Win32_PerfFormattedData_MSSQLMICROSOFTNumberNumberWID_MSSQLMICROSOFTNumberNumberWIDDatabaseMirroring) GetPropertyLogScannedforUndoKB() (value uint64, err error) {
	retValue, err := instance.GetProperty("LogScannedforUndoKB")
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

// SetLogSendFlowControlTimems sets the value of LogSendFlowControlTimems for the instance
func (instance *Win32_PerfFormattedData_MSSQLMICROSOFTNumberNumberWID_MSSQLMICROSOFTNumberNumberWIDDatabaseMirroring) SetPropertyLogSendFlowControlTimems(value uint64) (err error) {
	return instance.SetProperty("LogSendFlowControlTimems", (value))
}

// GetLogSendFlowControlTimems gets the value of LogSendFlowControlTimems for the instance
func (instance *Win32_PerfFormattedData_MSSQLMICROSOFTNumberNumberWID_MSSQLMICROSOFTNumberNumberWIDDatabaseMirroring) GetPropertyLogSendFlowControlTimems() (value uint64, err error) {
	retValue, err := instance.GetProperty("LogSendFlowControlTimems")
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

// SetLogSendQueueKB sets the value of LogSendQueueKB for the instance
func (instance *Win32_PerfFormattedData_MSSQLMICROSOFTNumberNumberWID_MSSQLMICROSOFTNumberNumberWIDDatabaseMirroring) SetPropertyLogSendQueueKB(value uint64) (err error) {
	return instance.SetProperty("LogSendQueueKB", (value))
}

// GetLogSendQueueKB gets the value of LogSendQueueKB for the instance
func (instance *Win32_PerfFormattedData_MSSQLMICROSOFTNumberNumberWID_MSSQLMICROSOFTNumberNumberWIDDatabaseMirroring) GetPropertyLogSendQueueKB() (value uint64, err error) {
	retValue, err := instance.GetProperty("LogSendQueueKB")
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
func (instance *Win32_PerfFormattedData_MSSQLMICROSOFTNumberNumberWID_MSSQLMICROSOFTNumberNumberWIDDatabaseMirroring) SetPropertyMirroredWriteTransactionsPersec(value uint64) (err error) {
	return instance.SetProperty("MirroredWriteTransactionsPersec", (value))
}

// GetMirroredWriteTransactionsPersec gets the value of MirroredWriteTransactionsPersec for the instance
func (instance *Win32_PerfFormattedData_MSSQLMICROSOFTNumberNumberWID_MSSQLMICROSOFTNumberNumberWIDDatabaseMirroring) GetPropertyMirroredWriteTransactionsPersec() (value uint64, err error) {
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

// SetPagesSentPersec sets the value of PagesSentPersec for the instance
func (instance *Win32_PerfFormattedData_MSSQLMICROSOFTNumberNumberWID_MSSQLMICROSOFTNumberNumberWIDDatabaseMirroring) SetPropertyPagesSentPersec(value uint64) (err error) {
	return instance.SetProperty("PagesSentPersec", (value))
}

// GetPagesSentPersec gets the value of PagesSentPersec for the instance
func (instance *Win32_PerfFormattedData_MSSQLMICROSOFTNumberNumberWID_MSSQLMICROSOFTNumberNumberWIDDatabaseMirroring) GetPropertyPagesSentPersec() (value uint64, err error) {
	retValue, err := instance.GetProperty("PagesSentPersec")
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

// SetReceivesPersec sets the value of ReceivesPersec for the instance
func (instance *Win32_PerfFormattedData_MSSQLMICROSOFTNumberNumberWID_MSSQLMICROSOFTNumberNumberWIDDatabaseMirroring) SetPropertyReceivesPersec(value uint64) (err error) {
	return instance.SetProperty("ReceivesPersec", (value))
}

// GetReceivesPersec gets the value of ReceivesPersec for the instance
func (instance *Win32_PerfFormattedData_MSSQLMICROSOFTNumberNumberWID_MSSQLMICROSOFTNumberNumberWIDDatabaseMirroring) GetPropertyReceivesPersec() (value uint64, err error) {
	retValue, err := instance.GetProperty("ReceivesPersec")
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

// SetRedoBytesPersec sets the value of RedoBytesPersec for the instance
func (instance *Win32_PerfFormattedData_MSSQLMICROSOFTNumberNumberWID_MSSQLMICROSOFTNumberNumberWIDDatabaseMirroring) SetPropertyRedoBytesPersec(value uint64) (err error) {
	return instance.SetProperty("RedoBytesPersec", (value))
}

// GetRedoBytesPersec gets the value of RedoBytesPersec for the instance
func (instance *Win32_PerfFormattedData_MSSQLMICROSOFTNumberNumberWID_MSSQLMICROSOFTNumberNumberWIDDatabaseMirroring) GetPropertyRedoBytesPersec() (value uint64, err error) {
	retValue, err := instance.GetProperty("RedoBytesPersec")
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

// SetRedoQueueKB sets the value of RedoQueueKB for the instance
func (instance *Win32_PerfFormattedData_MSSQLMICROSOFTNumberNumberWID_MSSQLMICROSOFTNumberNumberWIDDatabaseMirroring) SetPropertyRedoQueueKB(value uint64) (err error) {
	return instance.SetProperty("RedoQueueKB", (value))
}

// GetRedoQueueKB gets the value of RedoQueueKB for the instance
func (instance *Win32_PerfFormattedData_MSSQLMICROSOFTNumberNumberWID_MSSQLMICROSOFTNumberNumberWIDDatabaseMirroring) GetPropertyRedoQueueKB() (value uint64, err error) {
	retValue, err := instance.GetProperty("RedoQueueKB")
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

// SetSendPerReceiveAckTime sets the value of SendPerReceiveAckTime for the instance
func (instance *Win32_PerfFormattedData_MSSQLMICROSOFTNumberNumberWID_MSSQLMICROSOFTNumberNumberWIDDatabaseMirroring) SetPropertySendPerReceiveAckTime(value uint64) (err error) {
	return instance.SetProperty("SendPerReceiveAckTime", (value))
}

// GetSendPerReceiveAckTime gets the value of SendPerReceiveAckTime for the instance
func (instance *Win32_PerfFormattedData_MSSQLMICROSOFTNumberNumberWID_MSSQLMICROSOFTNumberNumberWIDDatabaseMirroring) GetPropertySendPerReceiveAckTime() (value uint64, err error) {
	retValue, err := instance.GetProperty("SendPerReceiveAckTime")
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

// SetSendsPersec sets the value of SendsPersec for the instance
func (instance *Win32_PerfFormattedData_MSSQLMICROSOFTNumberNumberWID_MSSQLMICROSOFTNumberNumberWIDDatabaseMirroring) SetPropertySendsPersec(value uint64) (err error) {
	return instance.SetProperty("SendsPersec", (value))
}

// GetSendsPersec gets the value of SendsPersec for the instance
func (instance *Win32_PerfFormattedData_MSSQLMICROSOFTNumberNumberWID_MSSQLMICROSOFTNumberNumberWIDDatabaseMirroring) GetPropertySendsPersec() (value uint64, err error) {
	retValue, err := instance.GetProperty("SendsPersec")
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
func (instance *Win32_PerfFormattedData_MSSQLMICROSOFTNumberNumberWID_MSSQLMICROSOFTNumberNumberWIDDatabaseMirroring) SetPropertyTransactionDelay(value uint64) (err error) {
	return instance.SetProperty("TransactionDelay", (value))
}

// GetTransactionDelay gets the value of TransactionDelay for the instance
func (instance *Win32_PerfFormattedData_MSSQLMICROSOFTNumberNumberWID_MSSQLMICROSOFTNumberNumberWIDDatabaseMirroring) GetPropertyTransactionDelay() (value uint64, err error) {
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
