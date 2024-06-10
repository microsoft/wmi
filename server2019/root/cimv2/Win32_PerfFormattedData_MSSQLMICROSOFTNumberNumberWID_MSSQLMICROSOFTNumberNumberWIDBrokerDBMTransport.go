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

// Win32_PerfFormattedData_MSSQLMICROSOFTNumberNumberWID_MSSQLMICROSOFTNumberNumberWIDBrokerDBMTransport struct
type Win32_PerfFormattedData_MSSQLMICROSOFTNumberNumberWID_MSSQLMICROSOFTNumberNumberWIDBrokerDBMTransport struct {
	*Win32_PerfFormattedData

	//
	CurrentBytesforRecvIO uint64

	//
	CurrentBytesforSendIO uint64

	//
	CurrentMsgFragsforSendIO uint64

	//
	DecryptedIObytesPersec uint64

	//
	DecryptionIPerOsPersec uint64

	//
	EncryptedIObytesPersec uint64

	//
	EncryptionIPerOsPersec uint64

	//
	MessageFragmentP10SendsPersec uint64

	//
	MessageFragmentP1SendsPersec uint64

	//
	MessageFragmentP2SendsPersec uint64

	//
	MessageFragmentP3SendsPersec uint64

	//
	MessageFragmentP4SendsPersec uint64

	//
	MessageFragmentP5SendsPersec uint64

	//
	MessageFragmentP6SendsPersec uint64

	//
	MessageFragmentP7SendsPersec uint64

	//
	MessageFragmentP8SendsPersec uint64

	//
	MessageFragmentP9SendsPersec uint64

	//
	MessageFragmentReceivesPersec uint64

	//
	MessageFragmentSendsPersec uint64

	//
	MsgFragmentRecvSizeAvg uint64

	//
	MsgFragmentSendSizeAvg uint64

	//
	OpenConnectionCount uint64

	//
	PendingBytesforRecvIO uint64

	//
	PendingBytesforSendIO uint64

	//
	PendingMsgFragsforRecvIO uint64

	//
	PendingMsgFragsforSendIO uint64

	//
	ReceiveFlowControlEntersPersec uint64

	//
	ReceiveFlowControlExitsPersec uint64

	//
	ReceiveFlowControlGate uint64

	//
	ReceiveIObytesPersec uint64

	//
	ReceiveIOLenAvg uint64

	//
	ReceiveIPerOsPersec uint64

	//
	RecvIOBufferCopiesbytesPersec uint64

	//
	RecvIOBufferCopiesCount uint64

	//
	SendFlowControlEntersPersec uint64

	//
	SendFlowControlExitsPersec uint64

	//
	SendFlowControlGate uint64

	//
	SendIObytesPersec uint64

	//
	SendIOLenAvg uint64

	//
	SendIPerOsPersec uint64
}

func NewWin32_PerfFormattedData_MSSQLMICROSOFTNumberNumberWID_MSSQLMICROSOFTNumberNumberWIDBrokerDBMTransportEx1(instance *cim.WmiInstance) (newInstance *Win32_PerfFormattedData_MSSQLMICROSOFTNumberNumberWID_MSSQLMICROSOFTNumberNumberWIDBrokerDBMTransport, err error) {
	tmp, err := NewWin32_PerfFormattedDataEx1(instance)

	if err != nil {
		return
	}
	newInstance = &Win32_PerfFormattedData_MSSQLMICROSOFTNumberNumberWID_MSSQLMICROSOFTNumberNumberWIDBrokerDBMTransport{
		Win32_PerfFormattedData: tmp,
	}
	return
}

func NewWin32_PerfFormattedData_MSSQLMICROSOFTNumberNumberWID_MSSQLMICROSOFTNumberNumberWIDBrokerDBMTransportEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *Win32_PerfFormattedData_MSSQLMICROSOFTNumberNumberWID_MSSQLMICROSOFTNumberNumberWIDBrokerDBMTransport, err error) {
	tmp, err := NewWin32_PerfFormattedDataEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &Win32_PerfFormattedData_MSSQLMICROSOFTNumberNumberWID_MSSQLMICROSOFTNumberNumberWIDBrokerDBMTransport{
		Win32_PerfFormattedData: tmp,
	}
	return
}

// SetCurrentBytesforRecvIO sets the value of CurrentBytesforRecvIO for the instance
func (instance *Win32_PerfFormattedData_MSSQLMICROSOFTNumberNumberWID_MSSQLMICROSOFTNumberNumberWIDBrokerDBMTransport) SetPropertyCurrentBytesforRecvIO(value uint64) (err error) {
	return instance.SetProperty("CurrentBytesforRecvIO", (value))
}

// GetCurrentBytesforRecvIO gets the value of CurrentBytesforRecvIO for the instance
func (instance *Win32_PerfFormattedData_MSSQLMICROSOFTNumberNumberWID_MSSQLMICROSOFTNumberNumberWIDBrokerDBMTransport) GetPropertyCurrentBytesforRecvIO() (value uint64, err error) {
	retValue, err := instance.GetProperty("CurrentBytesforRecvIO")
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

// SetCurrentBytesforSendIO sets the value of CurrentBytesforSendIO for the instance
func (instance *Win32_PerfFormattedData_MSSQLMICROSOFTNumberNumberWID_MSSQLMICROSOFTNumberNumberWIDBrokerDBMTransport) SetPropertyCurrentBytesforSendIO(value uint64) (err error) {
	return instance.SetProperty("CurrentBytesforSendIO", (value))
}

// GetCurrentBytesforSendIO gets the value of CurrentBytesforSendIO for the instance
func (instance *Win32_PerfFormattedData_MSSQLMICROSOFTNumberNumberWID_MSSQLMICROSOFTNumberNumberWIDBrokerDBMTransport) GetPropertyCurrentBytesforSendIO() (value uint64, err error) {
	retValue, err := instance.GetProperty("CurrentBytesforSendIO")
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

// SetCurrentMsgFragsforSendIO sets the value of CurrentMsgFragsforSendIO for the instance
func (instance *Win32_PerfFormattedData_MSSQLMICROSOFTNumberNumberWID_MSSQLMICROSOFTNumberNumberWIDBrokerDBMTransport) SetPropertyCurrentMsgFragsforSendIO(value uint64) (err error) {
	return instance.SetProperty("CurrentMsgFragsforSendIO", (value))
}

// GetCurrentMsgFragsforSendIO gets the value of CurrentMsgFragsforSendIO for the instance
func (instance *Win32_PerfFormattedData_MSSQLMICROSOFTNumberNumberWID_MSSQLMICROSOFTNumberNumberWIDBrokerDBMTransport) GetPropertyCurrentMsgFragsforSendIO() (value uint64, err error) {
	retValue, err := instance.GetProperty("CurrentMsgFragsforSendIO")
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

// SetDecryptedIObytesPersec sets the value of DecryptedIObytesPersec for the instance
func (instance *Win32_PerfFormattedData_MSSQLMICROSOFTNumberNumberWID_MSSQLMICROSOFTNumberNumberWIDBrokerDBMTransport) SetPropertyDecryptedIObytesPersec(value uint64) (err error) {
	return instance.SetProperty("DecryptedIObytesPersec", (value))
}

// GetDecryptedIObytesPersec gets the value of DecryptedIObytesPersec for the instance
func (instance *Win32_PerfFormattedData_MSSQLMICROSOFTNumberNumberWID_MSSQLMICROSOFTNumberNumberWIDBrokerDBMTransport) GetPropertyDecryptedIObytesPersec() (value uint64, err error) {
	retValue, err := instance.GetProperty("DecryptedIObytesPersec")
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

// SetDecryptionIPerOsPersec sets the value of DecryptionIPerOsPersec for the instance
func (instance *Win32_PerfFormattedData_MSSQLMICROSOFTNumberNumberWID_MSSQLMICROSOFTNumberNumberWIDBrokerDBMTransport) SetPropertyDecryptionIPerOsPersec(value uint64) (err error) {
	return instance.SetProperty("DecryptionIPerOsPersec", (value))
}

// GetDecryptionIPerOsPersec gets the value of DecryptionIPerOsPersec for the instance
func (instance *Win32_PerfFormattedData_MSSQLMICROSOFTNumberNumberWID_MSSQLMICROSOFTNumberNumberWIDBrokerDBMTransport) GetPropertyDecryptionIPerOsPersec() (value uint64, err error) {
	retValue, err := instance.GetProperty("DecryptionIPerOsPersec")
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

// SetEncryptedIObytesPersec sets the value of EncryptedIObytesPersec for the instance
func (instance *Win32_PerfFormattedData_MSSQLMICROSOFTNumberNumberWID_MSSQLMICROSOFTNumberNumberWIDBrokerDBMTransport) SetPropertyEncryptedIObytesPersec(value uint64) (err error) {
	return instance.SetProperty("EncryptedIObytesPersec", (value))
}

// GetEncryptedIObytesPersec gets the value of EncryptedIObytesPersec for the instance
func (instance *Win32_PerfFormattedData_MSSQLMICROSOFTNumberNumberWID_MSSQLMICROSOFTNumberNumberWIDBrokerDBMTransport) GetPropertyEncryptedIObytesPersec() (value uint64, err error) {
	retValue, err := instance.GetProperty("EncryptedIObytesPersec")
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

// SetEncryptionIPerOsPersec sets the value of EncryptionIPerOsPersec for the instance
func (instance *Win32_PerfFormattedData_MSSQLMICROSOFTNumberNumberWID_MSSQLMICROSOFTNumberNumberWIDBrokerDBMTransport) SetPropertyEncryptionIPerOsPersec(value uint64) (err error) {
	return instance.SetProperty("EncryptionIPerOsPersec", (value))
}

// GetEncryptionIPerOsPersec gets the value of EncryptionIPerOsPersec for the instance
func (instance *Win32_PerfFormattedData_MSSQLMICROSOFTNumberNumberWID_MSSQLMICROSOFTNumberNumberWIDBrokerDBMTransport) GetPropertyEncryptionIPerOsPersec() (value uint64, err error) {
	retValue, err := instance.GetProperty("EncryptionIPerOsPersec")
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

// SetMessageFragmentP10SendsPersec sets the value of MessageFragmentP10SendsPersec for the instance
func (instance *Win32_PerfFormattedData_MSSQLMICROSOFTNumberNumberWID_MSSQLMICROSOFTNumberNumberWIDBrokerDBMTransport) SetPropertyMessageFragmentP10SendsPersec(value uint64) (err error) {
	return instance.SetProperty("MessageFragmentP10SendsPersec", (value))
}

// GetMessageFragmentP10SendsPersec gets the value of MessageFragmentP10SendsPersec for the instance
func (instance *Win32_PerfFormattedData_MSSQLMICROSOFTNumberNumberWID_MSSQLMICROSOFTNumberNumberWIDBrokerDBMTransport) GetPropertyMessageFragmentP10SendsPersec() (value uint64, err error) {
	retValue, err := instance.GetProperty("MessageFragmentP10SendsPersec")
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

// SetMessageFragmentP1SendsPersec sets the value of MessageFragmentP1SendsPersec for the instance
func (instance *Win32_PerfFormattedData_MSSQLMICROSOFTNumberNumberWID_MSSQLMICROSOFTNumberNumberWIDBrokerDBMTransport) SetPropertyMessageFragmentP1SendsPersec(value uint64) (err error) {
	return instance.SetProperty("MessageFragmentP1SendsPersec", (value))
}

// GetMessageFragmentP1SendsPersec gets the value of MessageFragmentP1SendsPersec for the instance
func (instance *Win32_PerfFormattedData_MSSQLMICROSOFTNumberNumberWID_MSSQLMICROSOFTNumberNumberWIDBrokerDBMTransport) GetPropertyMessageFragmentP1SendsPersec() (value uint64, err error) {
	retValue, err := instance.GetProperty("MessageFragmentP1SendsPersec")
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

// SetMessageFragmentP2SendsPersec sets the value of MessageFragmentP2SendsPersec for the instance
func (instance *Win32_PerfFormattedData_MSSQLMICROSOFTNumberNumberWID_MSSQLMICROSOFTNumberNumberWIDBrokerDBMTransport) SetPropertyMessageFragmentP2SendsPersec(value uint64) (err error) {
	return instance.SetProperty("MessageFragmentP2SendsPersec", (value))
}

// GetMessageFragmentP2SendsPersec gets the value of MessageFragmentP2SendsPersec for the instance
func (instance *Win32_PerfFormattedData_MSSQLMICROSOFTNumberNumberWID_MSSQLMICROSOFTNumberNumberWIDBrokerDBMTransport) GetPropertyMessageFragmentP2SendsPersec() (value uint64, err error) {
	retValue, err := instance.GetProperty("MessageFragmentP2SendsPersec")
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

// SetMessageFragmentP3SendsPersec sets the value of MessageFragmentP3SendsPersec for the instance
func (instance *Win32_PerfFormattedData_MSSQLMICROSOFTNumberNumberWID_MSSQLMICROSOFTNumberNumberWIDBrokerDBMTransport) SetPropertyMessageFragmentP3SendsPersec(value uint64) (err error) {
	return instance.SetProperty("MessageFragmentP3SendsPersec", (value))
}

// GetMessageFragmentP3SendsPersec gets the value of MessageFragmentP3SendsPersec for the instance
func (instance *Win32_PerfFormattedData_MSSQLMICROSOFTNumberNumberWID_MSSQLMICROSOFTNumberNumberWIDBrokerDBMTransport) GetPropertyMessageFragmentP3SendsPersec() (value uint64, err error) {
	retValue, err := instance.GetProperty("MessageFragmentP3SendsPersec")
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

// SetMessageFragmentP4SendsPersec sets the value of MessageFragmentP4SendsPersec for the instance
func (instance *Win32_PerfFormattedData_MSSQLMICROSOFTNumberNumberWID_MSSQLMICROSOFTNumberNumberWIDBrokerDBMTransport) SetPropertyMessageFragmentP4SendsPersec(value uint64) (err error) {
	return instance.SetProperty("MessageFragmentP4SendsPersec", (value))
}

// GetMessageFragmentP4SendsPersec gets the value of MessageFragmentP4SendsPersec for the instance
func (instance *Win32_PerfFormattedData_MSSQLMICROSOFTNumberNumberWID_MSSQLMICROSOFTNumberNumberWIDBrokerDBMTransport) GetPropertyMessageFragmentP4SendsPersec() (value uint64, err error) {
	retValue, err := instance.GetProperty("MessageFragmentP4SendsPersec")
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

// SetMessageFragmentP5SendsPersec sets the value of MessageFragmentP5SendsPersec for the instance
func (instance *Win32_PerfFormattedData_MSSQLMICROSOFTNumberNumberWID_MSSQLMICROSOFTNumberNumberWIDBrokerDBMTransport) SetPropertyMessageFragmentP5SendsPersec(value uint64) (err error) {
	return instance.SetProperty("MessageFragmentP5SendsPersec", (value))
}

// GetMessageFragmentP5SendsPersec gets the value of MessageFragmentP5SendsPersec for the instance
func (instance *Win32_PerfFormattedData_MSSQLMICROSOFTNumberNumberWID_MSSQLMICROSOFTNumberNumberWIDBrokerDBMTransport) GetPropertyMessageFragmentP5SendsPersec() (value uint64, err error) {
	retValue, err := instance.GetProperty("MessageFragmentP5SendsPersec")
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

// SetMessageFragmentP6SendsPersec sets the value of MessageFragmentP6SendsPersec for the instance
func (instance *Win32_PerfFormattedData_MSSQLMICROSOFTNumberNumberWID_MSSQLMICROSOFTNumberNumberWIDBrokerDBMTransport) SetPropertyMessageFragmentP6SendsPersec(value uint64) (err error) {
	return instance.SetProperty("MessageFragmentP6SendsPersec", (value))
}

// GetMessageFragmentP6SendsPersec gets the value of MessageFragmentP6SendsPersec for the instance
func (instance *Win32_PerfFormattedData_MSSQLMICROSOFTNumberNumberWID_MSSQLMICROSOFTNumberNumberWIDBrokerDBMTransport) GetPropertyMessageFragmentP6SendsPersec() (value uint64, err error) {
	retValue, err := instance.GetProperty("MessageFragmentP6SendsPersec")
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

// SetMessageFragmentP7SendsPersec sets the value of MessageFragmentP7SendsPersec for the instance
func (instance *Win32_PerfFormattedData_MSSQLMICROSOFTNumberNumberWID_MSSQLMICROSOFTNumberNumberWIDBrokerDBMTransport) SetPropertyMessageFragmentP7SendsPersec(value uint64) (err error) {
	return instance.SetProperty("MessageFragmentP7SendsPersec", (value))
}

// GetMessageFragmentP7SendsPersec gets the value of MessageFragmentP7SendsPersec for the instance
func (instance *Win32_PerfFormattedData_MSSQLMICROSOFTNumberNumberWID_MSSQLMICROSOFTNumberNumberWIDBrokerDBMTransport) GetPropertyMessageFragmentP7SendsPersec() (value uint64, err error) {
	retValue, err := instance.GetProperty("MessageFragmentP7SendsPersec")
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

// SetMessageFragmentP8SendsPersec sets the value of MessageFragmentP8SendsPersec for the instance
func (instance *Win32_PerfFormattedData_MSSQLMICROSOFTNumberNumberWID_MSSQLMICROSOFTNumberNumberWIDBrokerDBMTransport) SetPropertyMessageFragmentP8SendsPersec(value uint64) (err error) {
	return instance.SetProperty("MessageFragmentP8SendsPersec", (value))
}

// GetMessageFragmentP8SendsPersec gets the value of MessageFragmentP8SendsPersec for the instance
func (instance *Win32_PerfFormattedData_MSSQLMICROSOFTNumberNumberWID_MSSQLMICROSOFTNumberNumberWIDBrokerDBMTransport) GetPropertyMessageFragmentP8SendsPersec() (value uint64, err error) {
	retValue, err := instance.GetProperty("MessageFragmentP8SendsPersec")
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

// SetMessageFragmentP9SendsPersec sets the value of MessageFragmentP9SendsPersec for the instance
func (instance *Win32_PerfFormattedData_MSSQLMICROSOFTNumberNumberWID_MSSQLMICROSOFTNumberNumberWIDBrokerDBMTransport) SetPropertyMessageFragmentP9SendsPersec(value uint64) (err error) {
	return instance.SetProperty("MessageFragmentP9SendsPersec", (value))
}

// GetMessageFragmentP9SendsPersec gets the value of MessageFragmentP9SendsPersec for the instance
func (instance *Win32_PerfFormattedData_MSSQLMICROSOFTNumberNumberWID_MSSQLMICROSOFTNumberNumberWIDBrokerDBMTransport) GetPropertyMessageFragmentP9SendsPersec() (value uint64, err error) {
	retValue, err := instance.GetProperty("MessageFragmentP9SendsPersec")
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

// SetMessageFragmentReceivesPersec sets the value of MessageFragmentReceivesPersec for the instance
func (instance *Win32_PerfFormattedData_MSSQLMICROSOFTNumberNumberWID_MSSQLMICROSOFTNumberNumberWIDBrokerDBMTransport) SetPropertyMessageFragmentReceivesPersec(value uint64) (err error) {
	return instance.SetProperty("MessageFragmentReceivesPersec", (value))
}

// GetMessageFragmentReceivesPersec gets the value of MessageFragmentReceivesPersec for the instance
func (instance *Win32_PerfFormattedData_MSSQLMICROSOFTNumberNumberWID_MSSQLMICROSOFTNumberNumberWIDBrokerDBMTransport) GetPropertyMessageFragmentReceivesPersec() (value uint64, err error) {
	retValue, err := instance.GetProperty("MessageFragmentReceivesPersec")
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

// SetMessageFragmentSendsPersec sets the value of MessageFragmentSendsPersec for the instance
func (instance *Win32_PerfFormattedData_MSSQLMICROSOFTNumberNumberWID_MSSQLMICROSOFTNumberNumberWIDBrokerDBMTransport) SetPropertyMessageFragmentSendsPersec(value uint64) (err error) {
	return instance.SetProperty("MessageFragmentSendsPersec", (value))
}

// GetMessageFragmentSendsPersec gets the value of MessageFragmentSendsPersec for the instance
func (instance *Win32_PerfFormattedData_MSSQLMICROSOFTNumberNumberWID_MSSQLMICROSOFTNumberNumberWIDBrokerDBMTransport) GetPropertyMessageFragmentSendsPersec() (value uint64, err error) {
	retValue, err := instance.GetProperty("MessageFragmentSendsPersec")
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

// SetMsgFragmentRecvSizeAvg sets the value of MsgFragmentRecvSizeAvg for the instance
func (instance *Win32_PerfFormattedData_MSSQLMICROSOFTNumberNumberWID_MSSQLMICROSOFTNumberNumberWIDBrokerDBMTransport) SetPropertyMsgFragmentRecvSizeAvg(value uint64) (err error) {
	return instance.SetProperty("MsgFragmentRecvSizeAvg", (value))
}

// GetMsgFragmentRecvSizeAvg gets the value of MsgFragmentRecvSizeAvg for the instance
func (instance *Win32_PerfFormattedData_MSSQLMICROSOFTNumberNumberWID_MSSQLMICROSOFTNumberNumberWIDBrokerDBMTransport) GetPropertyMsgFragmentRecvSizeAvg() (value uint64, err error) {
	retValue, err := instance.GetProperty("MsgFragmentRecvSizeAvg")
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

// SetMsgFragmentSendSizeAvg sets the value of MsgFragmentSendSizeAvg for the instance
func (instance *Win32_PerfFormattedData_MSSQLMICROSOFTNumberNumberWID_MSSQLMICROSOFTNumberNumberWIDBrokerDBMTransport) SetPropertyMsgFragmentSendSizeAvg(value uint64) (err error) {
	return instance.SetProperty("MsgFragmentSendSizeAvg", (value))
}

// GetMsgFragmentSendSizeAvg gets the value of MsgFragmentSendSizeAvg for the instance
func (instance *Win32_PerfFormattedData_MSSQLMICROSOFTNumberNumberWID_MSSQLMICROSOFTNumberNumberWIDBrokerDBMTransport) GetPropertyMsgFragmentSendSizeAvg() (value uint64, err error) {
	retValue, err := instance.GetProperty("MsgFragmentSendSizeAvg")
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

// SetOpenConnectionCount sets the value of OpenConnectionCount for the instance
func (instance *Win32_PerfFormattedData_MSSQLMICROSOFTNumberNumberWID_MSSQLMICROSOFTNumberNumberWIDBrokerDBMTransport) SetPropertyOpenConnectionCount(value uint64) (err error) {
	return instance.SetProperty("OpenConnectionCount", (value))
}

// GetOpenConnectionCount gets the value of OpenConnectionCount for the instance
func (instance *Win32_PerfFormattedData_MSSQLMICROSOFTNumberNumberWID_MSSQLMICROSOFTNumberNumberWIDBrokerDBMTransport) GetPropertyOpenConnectionCount() (value uint64, err error) {
	retValue, err := instance.GetProperty("OpenConnectionCount")
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

// SetPendingBytesforRecvIO sets the value of PendingBytesforRecvIO for the instance
func (instance *Win32_PerfFormattedData_MSSQLMICROSOFTNumberNumberWID_MSSQLMICROSOFTNumberNumberWIDBrokerDBMTransport) SetPropertyPendingBytesforRecvIO(value uint64) (err error) {
	return instance.SetProperty("PendingBytesforRecvIO", (value))
}

// GetPendingBytesforRecvIO gets the value of PendingBytesforRecvIO for the instance
func (instance *Win32_PerfFormattedData_MSSQLMICROSOFTNumberNumberWID_MSSQLMICROSOFTNumberNumberWIDBrokerDBMTransport) GetPropertyPendingBytesforRecvIO() (value uint64, err error) {
	retValue, err := instance.GetProperty("PendingBytesforRecvIO")
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

// SetPendingBytesforSendIO sets the value of PendingBytesforSendIO for the instance
func (instance *Win32_PerfFormattedData_MSSQLMICROSOFTNumberNumberWID_MSSQLMICROSOFTNumberNumberWIDBrokerDBMTransport) SetPropertyPendingBytesforSendIO(value uint64) (err error) {
	return instance.SetProperty("PendingBytesforSendIO", (value))
}

// GetPendingBytesforSendIO gets the value of PendingBytesforSendIO for the instance
func (instance *Win32_PerfFormattedData_MSSQLMICROSOFTNumberNumberWID_MSSQLMICROSOFTNumberNumberWIDBrokerDBMTransport) GetPropertyPendingBytesforSendIO() (value uint64, err error) {
	retValue, err := instance.GetProperty("PendingBytesforSendIO")
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

// SetPendingMsgFragsforRecvIO sets the value of PendingMsgFragsforRecvIO for the instance
func (instance *Win32_PerfFormattedData_MSSQLMICROSOFTNumberNumberWID_MSSQLMICROSOFTNumberNumberWIDBrokerDBMTransport) SetPropertyPendingMsgFragsforRecvIO(value uint64) (err error) {
	return instance.SetProperty("PendingMsgFragsforRecvIO", (value))
}

// GetPendingMsgFragsforRecvIO gets the value of PendingMsgFragsforRecvIO for the instance
func (instance *Win32_PerfFormattedData_MSSQLMICROSOFTNumberNumberWID_MSSQLMICROSOFTNumberNumberWIDBrokerDBMTransport) GetPropertyPendingMsgFragsforRecvIO() (value uint64, err error) {
	retValue, err := instance.GetProperty("PendingMsgFragsforRecvIO")
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

// SetPendingMsgFragsforSendIO sets the value of PendingMsgFragsforSendIO for the instance
func (instance *Win32_PerfFormattedData_MSSQLMICROSOFTNumberNumberWID_MSSQLMICROSOFTNumberNumberWIDBrokerDBMTransport) SetPropertyPendingMsgFragsforSendIO(value uint64) (err error) {
	return instance.SetProperty("PendingMsgFragsforSendIO", (value))
}

// GetPendingMsgFragsforSendIO gets the value of PendingMsgFragsforSendIO for the instance
func (instance *Win32_PerfFormattedData_MSSQLMICROSOFTNumberNumberWID_MSSQLMICROSOFTNumberNumberWIDBrokerDBMTransport) GetPropertyPendingMsgFragsforSendIO() (value uint64, err error) {
	retValue, err := instance.GetProperty("PendingMsgFragsforSendIO")
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

// SetReceiveFlowControlEntersPersec sets the value of ReceiveFlowControlEntersPersec for the instance
func (instance *Win32_PerfFormattedData_MSSQLMICROSOFTNumberNumberWID_MSSQLMICROSOFTNumberNumberWIDBrokerDBMTransport) SetPropertyReceiveFlowControlEntersPersec(value uint64) (err error) {
	return instance.SetProperty("ReceiveFlowControlEntersPersec", (value))
}

// GetReceiveFlowControlEntersPersec gets the value of ReceiveFlowControlEntersPersec for the instance
func (instance *Win32_PerfFormattedData_MSSQLMICROSOFTNumberNumberWID_MSSQLMICROSOFTNumberNumberWIDBrokerDBMTransport) GetPropertyReceiveFlowControlEntersPersec() (value uint64, err error) {
	retValue, err := instance.GetProperty("ReceiveFlowControlEntersPersec")
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

// SetReceiveFlowControlExitsPersec sets the value of ReceiveFlowControlExitsPersec for the instance
func (instance *Win32_PerfFormattedData_MSSQLMICROSOFTNumberNumberWID_MSSQLMICROSOFTNumberNumberWIDBrokerDBMTransport) SetPropertyReceiveFlowControlExitsPersec(value uint64) (err error) {
	return instance.SetProperty("ReceiveFlowControlExitsPersec", (value))
}

// GetReceiveFlowControlExitsPersec gets the value of ReceiveFlowControlExitsPersec for the instance
func (instance *Win32_PerfFormattedData_MSSQLMICROSOFTNumberNumberWID_MSSQLMICROSOFTNumberNumberWIDBrokerDBMTransport) GetPropertyReceiveFlowControlExitsPersec() (value uint64, err error) {
	retValue, err := instance.GetProperty("ReceiveFlowControlExitsPersec")
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

// SetReceiveFlowControlGate sets the value of ReceiveFlowControlGate for the instance
func (instance *Win32_PerfFormattedData_MSSQLMICROSOFTNumberNumberWID_MSSQLMICROSOFTNumberNumberWIDBrokerDBMTransport) SetPropertyReceiveFlowControlGate(value uint64) (err error) {
	return instance.SetProperty("ReceiveFlowControlGate", (value))
}

// GetReceiveFlowControlGate gets the value of ReceiveFlowControlGate for the instance
func (instance *Win32_PerfFormattedData_MSSQLMICROSOFTNumberNumberWID_MSSQLMICROSOFTNumberNumberWIDBrokerDBMTransport) GetPropertyReceiveFlowControlGate() (value uint64, err error) {
	retValue, err := instance.GetProperty("ReceiveFlowControlGate")
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

// SetReceiveIObytesPersec sets the value of ReceiveIObytesPersec for the instance
func (instance *Win32_PerfFormattedData_MSSQLMICROSOFTNumberNumberWID_MSSQLMICROSOFTNumberNumberWIDBrokerDBMTransport) SetPropertyReceiveIObytesPersec(value uint64) (err error) {
	return instance.SetProperty("ReceiveIObytesPersec", (value))
}

// GetReceiveIObytesPersec gets the value of ReceiveIObytesPersec for the instance
func (instance *Win32_PerfFormattedData_MSSQLMICROSOFTNumberNumberWID_MSSQLMICROSOFTNumberNumberWIDBrokerDBMTransport) GetPropertyReceiveIObytesPersec() (value uint64, err error) {
	retValue, err := instance.GetProperty("ReceiveIObytesPersec")
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

// SetReceiveIOLenAvg sets the value of ReceiveIOLenAvg for the instance
func (instance *Win32_PerfFormattedData_MSSQLMICROSOFTNumberNumberWID_MSSQLMICROSOFTNumberNumberWIDBrokerDBMTransport) SetPropertyReceiveIOLenAvg(value uint64) (err error) {
	return instance.SetProperty("ReceiveIOLenAvg", (value))
}

// GetReceiveIOLenAvg gets the value of ReceiveIOLenAvg for the instance
func (instance *Win32_PerfFormattedData_MSSQLMICROSOFTNumberNumberWID_MSSQLMICROSOFTNumberNumberWIDBrokerDBMTransport) GetPropertyReceiveIOLenAvg() (value uint64, err error) {
	retValue, err := instance.GetProperty("ReceiveIOLenAvg")
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

// SetReceiveIPerOsPersec sets the value of ReceiveIPerOsPersec for the instance
func (instance *Win32_PerfFormattedData_MSSQLMICROSOFTNumberNumberWID_MSSQLMICROSOFTNumberNumberWIDBrokerDBMTransport) SetPropertyReceiveIPerOsPersec(value uint64) (err error) {
	return instance.SetProperty("ReceiveIPerOsPersec", (value))
}

// GetReceiveIPerOsPersec gets the value of ReceiveIPerOsPersec for the instance
func (instance *Win32_PerfFormattedData_MSSQLMICROSOFTNumberNumberWID_MSSQLMICROSOFTNumberNumberWIDBrokerDBMTransport) GetPropertyReceiveIPerOsPersec() (value uint64, err error) {
	retValue, err := instance.GetProperty("ReceiveIPerOsPersec")
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

// SetRecvIOBufferCopiesbytesPersec sets the value of RecvIOBufferCopiesbytesPersec for the instance
func (instance *Win32_PerfFormattedData_MSSQLMICROSOFTNumberNumberWID_MSSQLMICROSOFTNumberNumberWIDBrokerDBMTransport) SetPropertyRecvIOBufferCopiesbytesPersec(value uint64) (err error) {
	return instance.SetProperty("RecvIOBufferCopiesbytesPersec", (value))
}

// GetRecvIOBufferCopiesbytesPersec gets the value of RecvIOBufferCopiesbytesPersec for the instance
func (instance *Win32_PerfFormattedData_MSSQLMICROSOFTNumberNumberWID_MSSQLMICROSOFTNumberNumberWIDBrokerDBMTransport) GetPropertyRecvIOBufferCopiesbytesPersec() (value uint64, err error) {
	retValue, err := instance.GetProperty("RecvIOBufferCopiesbytesPersec")
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

// SetRecvIOBufferCopiesCount sets the value of RecvIOBufferCopiesCount for the instance
func (instance *Win32_PerfFormattedData_MSSQLMICROSOFTNumberNumberWID_MSSQLMICROSOFTNumberNumberWIDBrokerDBMTransport) SetPropertyRecvIOBufferCopiesCount(value uint64) (err error) {
	return instance.SetProperty("RecvIOBufferCopiesCount", (value))
}

// GetRecvIOBufferCopiesCount gets the value of RecvIOBufferCopiesCount for the instance
func (instance *Win32_PerfFormattedData_MSSQLMICROSOFTNumberNumberWID_MSSQLMICROSOFTNumberNumberWIDBrokerDBMTransport) GetPropertyRecvIOBufferCopiesCount() (value uint64, err error) {
	retValue, err := instance.GetProperty("RecvIOBufferCopiesCount")
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

// SetSendFlowControlEntersPersec sets the value of SendFlowControlEntersPersec for the instance
func (instance *Win32_PerfFormattedData_MSSQLMICROSOFTNumberNumberWID_MSSQLMICROSOFTNumberNumberWIDBrokerDBMTransport) SetPropertySendFlowControlEntersPersec(value uint64) (err error) {
	return instance.SetProperty("SendFlowControlEntersPersec", (value))
}

// GetSendFlowControlEntersPersec gets the value of SendFlowControlEntersPersec for the instance
func (instance *Win32_PerfFormattedData_MSSQLMICROSOFTNumberNumberWID_MSSQLMICROSOFTNumberNumberWIDBrokerDBMTransport) GetPropertySendFlowControlEntersPersec() (value uint64, err error) {
	retValue, err := instance.GetProperty("SendFlowControlEntersPersec")
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

// SetSendFlowControlExitsPersec sets the value of SendFlowControlExitsPersec for the instance
func (instance *Win32_PerfFormattedData_MSSQLMICROSOFTNumberNumberWID_MSSQLMICROSOFTNumberNumberWIDBrokerDBMTransport) SetPropertySendFlowControlExitsPersec(value uint64) (err error) {
	return instance.SetProperty("SendFlowControlExitsPersec", (value))
}

// GetSendFlowControlExitsPersec gets the value of SendFlowControlExitsPersec for the instance
func (instance *Win32_PerfFormattedData_MSSQLMICROSOFTNumberNumberWID_MSSQLMICROSOFTNumberNumberWIDBrokerDBMTransport) GetPropertySendFlowControlExitsPersec() (value uint64, err error) {
	retValue, err := instance.GetProperty("SendFlowControlExitsPersec")
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

// SetSendFlowControlGate sets the value of SendFlowControlGate for the instance
func (instance *Win32_PerfFormattedData_MSSQLMICROSOFTNumberNumberWID_MSSQLMICROSOFTNumberNumberWIDBrokerDBMTransport) SetPropertySendFlowControlGate(value uint64) (err error) {
	return instance.SetProperty("SendFlowControlGate", (value))
}

// GetSendFlowControlGate gets the value of SendFlowControlGate for the instance
func (instance *Win32_PerfFormattedData_MSSQLMICROSOFTNumberNumberWID_MSSQLMICROSOFTNumberNumberWIDBrokerDBMTransport) GetPropertySendFlowControlGate() (value uint64, err error) {
	retValue, err := instance.GetProperty("SendFlowControlGate")
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

// SetSendIObytesPersec sets the value of SendIObytesPersec for the instance
func (instance *Win32_PerfFormattedData_MSSQLMICROSOFTNumberNumberWID_MSSQLMICROSOFTNumberNumberWIDBrokerDBMTransport) SetPropertySendIObytesPersec(value uint64) (err error) {
	return instance.SetProperty("SendIObytesPersec", (value))
}

// GetSendIObytesPersec gets the value of SendIObytesPersec for the instance
func (instance *Win32_PerfFormattedData_MSSQLMICROSOFTNumberNumberWID_MSSQLMICROSOFTNumberNumberWIDBrokerDBMTransport) GetPropertySendIObytesPersec() (value uint64, err error) {
	retValue, err := instance.GetProperty("SendIObytesPersec")
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

// SetSendIOLenAvg sets the value of SendIOLenAvg for the instance
func (instance *Win32_PerfFormattedData_MSSQLMICROSOFTNumberNumberWID_MSSQLMICROSOFTNumberNumberWIDBrokerDBMTransport) SetPropertySendIOLenAvg(value uint64) (err error) {
	return instance.SetProperty("SendIOLenAvg", (value))
}

// GetSendIOLenAvg gets the value of SendIOLenAvg for the instance
func (instance *Win32_PerfFormattedData_MSSQLMICROSOFTNumberNumberWID_MSSQLMICROSOFTNumberNumberWIDBrokerDBMTransport) GetPropertySendIOLenAvg() (value uint64, err error) {
	retValue, err := instance.GetProperty("SendIOLenAvg")
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

// SetSendIPerOsPersec sets the value of SendIPerOsPersec for the instance
func (instance *Win32_PerfFormattedData_MSSQLMICROSOFTNumberNumberWID_MSSQLMICROSOFTNumberNumberWIDBrokerDBMTransport) SetPropertySendIPerOsPersec(value uint64) (err error) {
	return instance.SetProperty("SendIPerOsPersec", (value))
}

// GetSendIPerOsPersec gets the value of SendIPerOsPersec for the instance
func (instance *Win32_PerfFormattedData_MSSQLMICROSOFTNumberNumberWID_MSSQLMICROSOFTNumberNumberWIDBrokerDBMTransport) GetPropertySendIPerOsPersec() (value uint64, err error) {
	retValue, err := instance.GetProperty("SendIPerOsPersec")
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
