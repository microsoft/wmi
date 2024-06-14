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

// Win32_PerfRawData_MSSQLMICROSOFTNumberNumberWID_MSSQLMICROSOFTNumberNumberWIDBrokerStatistics struct
type Win32_PerfRawData_MSSQLMICROSOFTNumberNumberWID_MSSQLMICROSOFTNumberNumberWIDBrokerStatistics struct {
	*Win32_PerfRawData

	//
	ActivationErrorsTotal uint64

	//
	BrokerTransactionRollbacks uint64

	//
	CorruptedMessagesTotal uint64

	//
	DequeuedTransmissionQMsgsPersec uint64

	//
	DialogTimerEventCount uint64

	//
	DroppedMessagesTotal uint64

	//
	EnqueuedLocalMessagesPersec uint64

	//
	EnqueuedLocalMessagesTotal uint64

	//
	EnqueuedMessagesPersec uint64

	//
	EnqueuedMessagesTotal uint64

	//
	EnqueuedP10MessagesPersec uint64

	//
	EnqueuedP1MessagesPersec uint64

	//
	EnqueuedP2MessagesPersec uint64

	//
	EnqueuedP3MessagesPersec uint64

	//
	EnqueuedP4MessagesPersec uint64

	//
	EnqueuedP5MessagesPersec uint64

	//
	EnqueuedP6MessagesPersec uint64

	//
	EnqueuedP7MessagesPersec uint64

	//
	EnqueuedP8MessagesPersec uint64

	//
	EnqueuedP9MessagesPersec uint64

	//
	EnqueuedTransmissionQMsgsPersec uint64

	//
	EnqueuedTransportMsgFragsPersec uint64

	//
	EnqueuedTransportMsgFragTot uint64

	//
	EnqueuedTransportMsgsPersec uint64

	//
	EnqueuedTransportMsgsTotal uint64

	//
	ForwardedMessagesPersec uint64

	//
	ForwardedMessagesTotal uint64

	//
	ForwardedMsgBytesPersec uint64

	//
	ForwardedMsgByteTotal uint64

	//
	ForwardedMsgDiscardedTotal uint64

	//
	ForwardedMsgsDiscardedPersec uint64

	//
	ForwardedPendingMsgBytes uint64

	//
	ForwardedPendingMsgCount uint64

	//
	SQLRECEIVEsPersec uint64

	//
	SQLRECEIVETotal uint64

	//
	SQLSENDsPersec uint64

	//
	SQLSENDTotal uint64
}

func NewWin32_PerfRawData_MSSQLMICROSOFTNumberNumberWID_MSSQLMICROSOFTNumberNumberWIDBrokerStatisticsEx1(instance *cim.WmiInstance) (newInstance *Win32_PerfRawData_MSSQLMICROSOFTNumberNumberWID_MSSQLMICROSOFTNumberNumberWIDBrokerStatistics, err error) {
	tmp, err := NewWin32_PerfRawDataEx1(instance)

	if err != nil {
		return
	}
	newInstance = &Win32_PerfRawData_MSSQLMICROSOFTNumberNumberWID_MSSQLMICROSOFTNumberNumberWIDBrokerStatistics{
		Win32_PerfRawData: tmp,
	}
	return
}

func NewWin32_PerfRawData_MSSQLMICROSOFTNumberNumberWID_MSSQLMICROSOFTNumberNumberWIDBrokerStatisticsEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *Win32_PerfRawData_MSSQLMICROSOFTNumberNumberWID_MSSQLMICROSOFTNumberNumberWIDBrokerStatistics, err error) {
	tmp, err := NewWin32_PerfRawDataEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &Win32_PerfRawData_MSSQLMICROSOFTNumberNumberWID_MSSQLMICROSOFTNumberNumberWIDBrokerStatistics{
		Win32_PerfRawData: tmp,
	}
	return
}

// SetActivationErrorsTotal sets the value of ActivationErrorsTotal for the instance
func (instance *Win32_PerfRawData_MSSQLMICROSOFTNumberNumberWID_MSSQLMICROSOFTNumberNumberWIDBrokerStatistics) SetPropertyActivationErrorsTotal(value uint64) (err error) {
	return instance.SetProperty("ActivationErrorsTotal", (value))
}

// GetActivationErrorsTotal gets the value of ActivationErrorsTotal for the instance
func (instance *Win32_PerfRawData_MSSQLMICROSOFTNumberNumberWID_MSSQLMICROSOFTNumberNumberWIDBrokerStatistics) GetPropertyActivationErrorsTotal() (value uint64, err error) {
	retValue, err := instance.GetProperty("ActivationErrorsTotal")
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

// SetBrokerTransactionRollbacks sets the value of BrokerTransactionRollbacks for the instance
func (instance *Win32_PerfRawData_MSSQLMICROSOFTNumberNumberWID_MSSQLMICROSOFTNumberNumberWIDBrokerStatistics) SetPropertyBrokerTransactionRollbacks(value uint64) (err error) {
	return instance.SetProperty("BrokerTransactionRollbacks", (value))
}

// GetBrokerTransactionRollbacks gets the value of BrokerTransactionRollbacks for the instance
func (instance *Win32_PerfRawData_MSSQLMICROSOFTNumberNumberWID_MSSQLMICROSOFTNumberNumberWIDBrokerStatistics) GetPropertyBrokerTransactionRollbacks() (value uint64, err error) {
	retValue, err := instance.GetProperty("BrokerTransactionRollbacks")
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

// SetCorruptedMessagesTotal sets the value of CorruptedMessagesTotal for the instance
func (instance *Win32_PerfRawData_MSSQLMICROSOFTNumberNumberWID_MSSQLMICROSOFTNumberNumberWIDBrokerStatistics) SetPropertyCorruptedMessagesTotal(value uint64) (err error) {
	return instance.SetProperty("CorruptedMessagesTotal", (value))
}

// GetCorruptedMessagesTotal gets the value of CorruptedMessagesTotal for the instance
func (instance *Win32_PerfRawData_MSSQLMICROSOFTNumberNumberWID_MSSQLMICROSOFTNumberNumberWIDBrokerStatistics) GetPropertyCorruptedMessagesTotal() (value uint64, err error) {
	retValue, err := instance.GetProperty("CorruptedMessagesTotal")
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

// SetDequeuedTransmissionQMsgsPersec sets the value of DequeuedTransmissionQMsgsPersec for the instance
func (instance *Win32_PerfRawData_MSSQLMICROSOFTNumberNumberWID_MSSQLMICROSOFTNumberNumberWIDBrokerStatistics) SetPropertyDequeuedTransmissionQMsgsPersec(value uint64) (err error) {
	return instance.SetProperty("DequeuedTransmissionQMsgsPersec", (value))
}

// GetDequeuedTransmissionQMsgsPersec gets the value of DequeuedTransmissionQMsgsPersec for the instance
func (instance *Win32_PerfRawData_MSSQLMICROSOFTNumberNumberWID_MSSQLMICROSOFTNumberNumberWIDBrokerStatistics) GetPropertyDequeuedTransmissionQMsgsPersec() (value uint64, err error) {
	retValue, err := instance.GetProperty("DequeuedTransmissionQMsgsPersec")
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

// SetDialogTimerEventCount sets the value of DialogTimerEventCount for the instance
func (instance *Win32_PerfRawData_MSSQLMICROSOFTNumberNumberWID_MSSQLMICROSOFTNumberNumberWIDBrokerStatistics) SetPropertyDialogTimerEventCount(value uint64) (err error) {
	return instance.SetProperty("DialogTimerEventCount", (value))
}

// GetDialogTimerEventCount gets the value of DialogTimerEventCount for the instance
func (instance *Win32_PerfRawData_MSSQLMICROSOFTNumberNumberWID_MSSQLMICROSOFTNumberNumberWIDBrokerStatistics) GetPropertyDialogTimerEventCount() (value uint64, err error) {
	retValue, err := instance.GetProperty("DialogTimerEventCount")
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

// SetDroppedMessagesTotal sets the value of DroppedMessagesTotal for the instance
func (instance *Win32_PerfRawData_MSSQLMICROSOFTNumberNumberWID_MSSQLMICROSOFTNumberNumberWIDBrokerStatistics) SetPropertyDroppedMessagesTotal(value uint64) (err error) {
	return instance.SetProperty("DroppedMessagesTotal", (value))
}

// GetDroppedMessagesTotal gets the value of DroppedMessagesTotal for the instance
func (instance *Win32_PerfRawData_MSSQLMICROSOFTNumberNumberWID_MSSQLMICROSOFTNumberNumberWIDBrokerStatistics) GetPropertyDroppedMessagesTotal() (value uint64, err error) {
	retValue, err := instance.GetProperty("DroppedMessagesTotal")
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

// SetEnqueuedLocalMessagesPersec sets the value of EnqueuedLocalMessagesPersec for the instance
func (instance *Win32_PerfRawData_MSSQLMICROSOFTNumberNumberWID_MSSQLMICROSOFTNumberNumberWIDBrokerStatistics) SetPropertyEnqueuedLocalMessagesPersec(value uint64) (err error) {
	return instance.SetProperty("EnqueuedLocalMessagesPersec", (value))
}

// GetEnqueuedLocalMessagesPersec gets the value of EnqueuedLocalMessagesPersec for the instance
func (instance *Win32_PerfRawData_MSSQLMICROSOFTNumberNumberWID_MSSQLMICROSOFTNumberNumberWIDBrokerStatistics) GetPropertyEnqueuedLocalMessagesPersec() (value uint64, err error) {
	retValue, err := instance.GetProperty("EnqueuedLocalMessagesPersec")
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

// SetEnqueuedLocalMessagesTotal sets the value of EnqueuedLocalMessagesTotal for the instance
func (instance *Win32_PerfRawData_MSSQLMICROSOFTNumberNumberWID_MSSQLMICROSOFTNumberNumberWIDBrokerStatistics) SetPropertyEnqueuedLocalMessagesTotal(value uint64) (err error) {
	return instance.SetProperty("EnqueuedLocalMessagesTotal", (value))
}

// GetEnqueuedLocalMessagesTotal gets the value of EnqueuedLocalMessagesTotal for the instance
func (instance *Win32_PerfRawData_MSSQLMICROSOFTNumberNumberWID_MSSQLMICROSOFTNumberNumberWIDBrokerStatistics) GetPropertyEnqueuedLocalMessagesTotal() (value uint64, err error) {
	retValue, err := instance.GetProperty("EnqueuedLocalMessagesTotal")
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

// SetEnqueuedMessagesPersec sets the value of EnqueuedMessagesPersec for the instance
func (instance *Win32_PerfRawData_MSSQLMICROSOFTNumberNumberWID_MSSQLMICROSOFTNumberNumberWIDBrokerStatistics) SetPropertyEnqueuedMessagesPersec(value uint64) (err error) {
	return instance.SetProperty("EnqueuedMessagesPersec", (value))
}

// GetEnqueuedMessagesPersec gets the value of EnqueuedMessagesPersec for the instance
func (instance *Win32_PerfRawData_MSSQLMICROSOFTNumberNumberWID_MSSQLMICROSOFTNumberNumberWIDBrokerStatistics) GetPropertyEnqueuedMessagesPersec() (value uint64, err error) {
	retValue, err := instance.GetProperty("EnqueuedMessagesPersec")
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

// SetEnqueuedMessagesTotal sets the value of EnqueuedMessagesTotal for the instance
func (instance *Win32_PerfRawData_MSSQLMICROSOFTNumberNumberWID_MSSQLMICROSOFTNumberNumberWIDBrokerStatistics) SetPropertyEnqueuedMessagesTotal(value uint64) (err error) {
	return instance.SetProperty("EnqueuedMessagesTotal", (value))
}

// GetEnqueuedMessagesTotal gets the value of EnqueuedMessagesTotal for the instance
func (instance *Win32_PerfRawData_MSSQLMICROSOFTNumberNumberWID_MSSQLMICROSOFTNumberNumberWIDBrokerStatistics) GetPropertyEnqueuedMessagesTotal() (value uint64, err error) {
	retValue, err := instance.GetProperty("EnqueuedMessagesTotal")
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

// SetEnqueuedP10MessagesPersec sets the value of EnqueuedP10MessagesPersec for the instance
func (instance *Win32_PerfRawData_MSSQLMICROSOFTNumberNumberWID_MSSQLMICROSOFTNumberNumberWIDBrokerStatistics) SetPropertyEnqueuedP10MessagesPersec(value uint64) (err error) {
	return instance.SetProperty("EnqueuedP10MessagesPersec", (value))
}

// GetEnqueuedP10MessagesPersec gets the value of EnqueuedP10MessagesPersec for the instance
func (instance *Win32_PerfRawData_MSSQLMICROSOFTNumberNumberWID_MSSQLMICROSOFTNumberNumberWIDBrokerStatistics) GetPropertyEnqueuedP10MessagesPersec() (value uint64, err error) {
	retValue, err := instance.GetProperty("EnqueuedP10MessagesPersec")
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

// SetEnqueuedP1MessagesPersec sets the value of EnqueuedP1MessagesPersec for the instance
func (instance *Win32_PerfRawData_MSSQLMICROSOFTNumberNumberWID_MSSQLMICROSOFTNumberNumberWIDBrokerStatistics) SetPropertyEnqueuedP1MessagesPersec(value uint64) (err error) {
	return instance.SetProperty("EnqueuedP1MessagesPersec", (value))
}

// GetEnqueuedP1MessagesPersec gets the value of EnqueuedP1MessagesPersec for the instance
func (instance *Win32_PerfRawData_MSSQLMICROSOFTNumberNumberWID_MSSQLMICROSOFTNumberNumberWIDBrokerStatistics) GetPropertyEnqueuedP1MessagesPersec() (value uint64, err error) {
	retValue, err := instance.GetProperty("EnqueuedP1MessagesPersec")
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

// SetEnqueuedP2MessagesPersec sets the value of EnqueuedP2MessagesPersec for the instance
func (instance *Win32_PerfRawData_MSSQLMICROSOFTNumberNumberWID_MSSQLMICROSOFTNumberNumberWIDBrokerStatistics) SetPropertyEnqueuedP2MessagesPersec(value uint64) (err error) {
	return instance.SetProperty("EnqueuedP2MessagesPersec", (value))
}

// GetEnqueuedP2MessagesPersec gets the value of EnqueuedP2MessagesPersec for the instance
func (instance *Win32_PerfRawData_MSSQLMICROSOFTNumberNumberWID_MSSQLMICROSOFTNumberNumberWIDBrokerStatistics) GetPropertyEnqueuedP2MessagesPersec() (value uint64, err error) {
	retValue, err := instance.GetProperty("EnqueuedP2MessagesPersec")
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

// SetEnqueuedP3MessagesPersec sets the value of EnqueuedP3MessagesPersec for the instance
func (instance *Win32_PerfRawData_MSSQLMICROSOFTNumberNumberWID_MSSQLMICROSOFTNumberNumberWIDBrokerStatistics) SetPropertyEnqueuedP3MessagesPersec(value uint64) (err error) {
	return instance.SetProperty("EnqueuedP3MessagesPersec", (value))
}

// GetEnqueuedP3MessagesPersec gets the value of EnqueuedP3MessagesPersec for the instance
func (instance *Win32_PerfRawData_MSSQLMICROSOFTNumberNumberWID_MSSQLMICROSOFTNumberNumberWIDBrokerStatistics) GetPropertyEnqueuedP3MessagesPersec() (value uint64, err error) {
	retValue, err := instance.GetProperty("EnqueuedP3MessagesPersec")
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

// SetEnqueuedP4MessagesPersec sets the value of EnqueuedP4MessagesPersec for the instance
func (instance *Win32_PerfRawData_MSSQLMICROSOFTNumberNumberWID_MSSQLMICROSOFTNumberNumberWIDBrokerStatistics) SetPropertyEnqueuedP4MessagesPersec(value uint64) (err error) {
	return instance.SetProperty("EnqueuedP4MessagesPersec", (value))
}

// GetEnqueuedP4MessagesPersec gets the value of EnqueuedP4MessagesPersec for the instance
func (instance *Win32_PerfRawData_MSSQLMICROSOFTNumberNumberWID_MSSQLMICROSOFTNumberNumberWIDBrokerStatistics) GetPropertyEnqueuedP4MessagesPersec() (value uint64, err error) {
	retValue, err := instance.GetProperty("EnqueuedP4MessagesPersec")
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

// SetEnqueuedP5MessagesPersec sets the value of EnqueuedP5MessagesPersec for the instance
func (instance *Win32_PerfRawData_MSSQLMICROSOFTNumberNumberWID_MSSQLMICROSOFTNumberNumberWIDBrokerStatistics) SetPropertyEnqueuedP5MessagesPersec(value uint64) (err error) {
	return instance.SetProperty("EnqueuedP5MessagesPersec", (value))
}

// GetEnqueuedP5MessagesPersec gets the value of EnqueuedP5MessagesPersec for the instance
func (instance *Win32_PerfRawData_MSSQLMICROSOFTNumberNumberWID_MSSQLMICROSOFTNumberNumberWIDBrokerStatistics) GetPropertyEnqueuedP5MessagesPersec() (value uint64, err error) {
	retValue, err := instance.GetProperty("EnqueuedP5MessagesPersec")
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

// SetEnqueuedP6MessagesPersec sets the value of EnqueuedP6MessagesPersec for the instance
func (instance *Win32_PerfRawData_MSSQLMICROSOFTNumberNumberWID_MSSQLMICROSOFTNumberNumberWIDBrokerStatistics) SetPropertyEnqueuedP6MessagesPersec(value uint64) (err error) {
	return instance.SetProperty("EnqueuedP6MessagesPersec", (value))
}

// GetEnqueuedP6MessagesPersec gets the value of EnqueuedP6MessagesPersec for the instance
func (instance *Win32_PerfRawData_MSSQLMICROSOFTNumberNumberWID_MSSQLMICROSOFTNumberNumberWIDBrokerStatistics) GetPropertyEnqueuedP6MessagesPersec() (value uint64, err error) {
	retValue, err := instance.GetProperty("EnqueuedP6MessagesPersec")
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

// SetEnqueuedP7MessagesPersec sets the value of EnqueuedP7MessagesPersec for the instance
func (instance *Win32_PerfRawData_MSSQLMICROSOFTNumberNumberWID_MSSQLMICROSOFTNumberNumberWIDBrokerStatistics) SetPropertyEnqueuedP7MessagesPersec(value uint64) (err error) {
	return instance.SetProperty("EnqueuedP7MessagesPersec", (value))
}

// GetEnqueuedP7MessagesPersec gets the value of EnqueuedP7MessagesPersec for the instance
func (instance *Win32_PerfRawData_MSSQLMICROSOFTNumberNumberWID_MSSQLMICROSOFTNumberNumberWIDBrokerStatistics) GetPropertyEnqueuedP7MessagesPersec() (value uint64, err error) {
	retValue, err := instance.GetProperty("EnqueuedP7MessagesPersec")
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

// SetEnqueuedP8MessagesPersec sets the value of EnqueuedP8MessagesPersec for the instance
func (instance *Win32_PerfRawData_MSSQLMICROSOFTNumberNumberWID_MSSQLMICROSOFTNumberNumberWIDBrokerStatistics) SetPropertyEnqueuedP8MessagesPersec(value uint64) (err error) {
	return instance.SetProperty("EnqueuedP8MessagesPersec", (value))
}

// GetEnqueuedP8MessagesPersec gets the value of EnqueuedP8MessagesPersec for the instance
func (instance *Win32_PerfRawData_MSSQLMICROSOFTNumberNumberWID_MSSQLMICROSOFTNumberNumberWIDBrokerStatistics) GetPropertyEnqueuedP8MessagesPersec() (value uint64, err error) {
	retValue, err := instance.GetProperty("EnqueuedP8MessagesPersec")
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

// SetEnqueuedP9MessagesPersec sets the value of EnqueuedP9MessagesPersec for the instance
func (instance *Win32_PerfRawData_MSSQLMICROSOFTNumberNumberWID_MSSQLMICROSOFTNumberNumberWIDBrokerStatistics) SetPropertyEnqueuedP9MessagesPersec(value uint64) (err error) {
	return instance.SetProperty("EnqueuedP9MessagesPersec", (value))
}

// GetEnqueuedP9MessagesPersec gets the value of EnqueuedP9MessagesPersec for the instance
func (instance *Win32_PerfRawData_MSSQLMICROSOFTNumberNumberWID_MSSQLMICROSOFTNumberNumberWIDBrokerStatistics) GetPropertyEnqueuedP9MessagesPersec() (value uint64, err error) {
	retValue, err := instance.GetProperty("EnqueuedP9MessagesPersec")
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

// SetEnqueuedTransmissionQMsgsPersec sets the value of EnqueuedTransmissionQMsgsPersec for the instance
func (instance *Win32_PerfRawData_MSSQLMICROSOFTNumberNumberWID_MSSQLMICROSOFTNumberNumberWIDBrokerStatistics) SetPropertyEnqueuedTransmissionQMsgsPersec(value uint64) (err error) {
	return instance.SetProperty("EnqueuedTransmissionQMsgsPersec", (value))
}

// GetEnqueuedTransmissionQMsgsPersec gets the value of EnqueuedTransmissionQMsgsPersec for the instance
func (instance *Win32_PerfRawData_MSSQLMICROSOFTNumberNumberWID_MSSQLMICROSOFTNumberNumberWIDBrokerStatistics) GetPropertyEnqueuedTransmissionQMsgsPersec() (value uint64, err error) {
	retValue, err := instance.GetProperty("EnqueuedTransmissionQMsgsPersec")
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

// SetEnqueuedTransportMsgFragsPersec sets the value of EnqueuedTransportMsgFragsPersec for the instance
func (instance *Win32_PerfRawData_MSSQLMICROSOFTNumberNumberWID_MSSQLMICROSOFTNumberNumberWIDBrokerStatistics) SetPropertyEnqueuedTransportMsgFragsPersec(value uint64) (err error) {
	return instance.SetProperty("EnqueuedTransportMsgFragsPersec", (value))
}

// GetEnqueuedTransportMsgFragsPersec gets the value of EnqueuedTransportMsgFragsPersec for the instance
func (instance *Win32_PerfRawData_MSSQLMICROSOFTNumberNumberWID_MSSQLMICROSOFTNumberNumberWIDBrokerStatistics) GetPropertyEnqueuedTransportMsgFragsPersec() (value uint64, err error) {
	retValue, err := instance.GetProperty("EnqueuedTransportMsgFragsPersec")
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

// SetEnqueuedTransportMsgFragTot sets the value of EnqueuedTransportMsgFragTot for the instance
func (instance *Win32_PerfRawData_MSSQLMICROSOFTNumberNumberWID_MSSQLMICROSOFTNumberNumberWIDBrokerStatistics) SetPropertyEnqueuedTransportMsgFragTot(value uint64) (err error) {
	return instance.SetProperty("EnqueuedTransportMsgFragTot", (value))
}

// GetEnqueuedTransportMsgFragTot gets the value of EnqueuedTransportMsgFragTot for the instance
func (instance *Win32_PerfRawData_MSSQLMICROSOFTNumberNumberWID_MSSQLMICROSOFTNumberNumberWIDBrokerStatistics) GetPropertyEnqueuedTransportMsgFragTot() (value uint64, err error) {
	retValue, err := instance.GetProperty("EnqueuedTransportMsgFragTot")
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

// SetEnqueuedTransportMsgsPersec sets the value of EnqueuedTransportMsgsPersec for the instance
func (instance *Win32_PerfRawData_MSSQLMICROSOFTNumberNumberWID_MSSQLMICROSOFTNumberNumberWIDBrokerStatistics) SetPropertyEnqueuedTransportMsgsPersec(value uint64) (err error) {
	return instance.SetProperty("EnqueuedTransportMsgsPersec", (value))
}

// GetEnqueuedTransportMsgsPersec gets the value of EnqueuedTransportMsgsPersec for the instance
func (instance *Win32_PerfRawData_MSSQLMICROSOFTNumberNumberWID_MSSQLMICROSOFTNumberNumberWIDBrokerStatistics) GetPropertyEnqueuedTransportMsgsPersec() (value uint64, err error) {
	retValue, err := instance.GetProperty("EnqueuedTransportMsgsPersec")
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

// SetEnqueuedTransportMsgsTotal sets the value of EnqueuedTransportMsgsTotal for the instance
func (instance *Win32_PerfRawData_MSSQLMICROSOFTNumberNumberWID_MSSQLMICROSOFTNumberNumberWIDBrokerStatistics) SetPropertyEnqueuedTransportMsgsTotal(value uint64) (err error) {
	return instance.SetProperty("EnqueuedTransportMsgsTotal", (value))
}

// GetEnqueuedTransportMsgsTotal gets the value of EnqueuedTransportMsgsTotal for the instance
func (instance *Win32_PerfRawData_MSSQLMICROSOFTNumberNumberWID_MSSQLMICROSOFTNumberNumberWIDBrokerStatistics) GetPropertyEnqueuedTransportMsgsTotal() (value uint64, err error) {
	retValue, err := instance.GetProperty("EnqueuedTransportMsgsTotal")
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

// SetForwardedMessagesPersec sets the value of ForwardedMessagesPersec for the instance
func (instance *Win32_PerfRawData_MSSQLMICROSOFTNumberNumberWID_MSSQLMICROSOFTNumberNumberWIDBrokerStatistics) SetPropertyForwardedMessagesPersec(value uint64) (err error) {
	return instance.SetProperty("ForwardedMessagesPersec", (value))
}

// GetForwardedMessagesPersec gets the value of ForwardedMessagesPersec for the instance
func (instance *Win32_PerfRawData_MSSQLMICROSOFTNumberNumberWID_MSSQLMICROSOFTNumberNumberWIDBrokerStatistics) GetPropertyForwardedMessagesPersec() (value uint64, err error) {
	retValue, err := instance.GetProperty("ForwardedMessagesPersec")
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

// SetForwardedMessagesTotal sets the value of ForwardedMessagesTotal for the instance
func (instance *Win32_PerfRawData_MSSQLMICROSOFTNumberNumberWID_MSSQLMICROSOFTNumberNumberWIDBrokerStatistics) SetPropertyForwardedMessagesTotal(value uint64) (err error) {
	return instance.SetProperty("ForwardedMessagesTotal", (value))
}

// GetForwardedMessagesTotal gets the value of ForwardedMessagesTotal for the instance
func (instance *Win32_PerfRawData_MSSQLMICROSOFTNumberNumberWID_MSSQLMICROSOFTNumberNumberWIDBrokerStatistics) GetPropertyForwardedMessagesTotal() (value uint64, err error) {
	retValue, err := instance.GetProperty("ForwardedMessagesTotal")
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

// SetForwardedMsgBytesPersec sets the value of ForwardedMsgBytesPersec for the instance
func (instance *Win32_PerfRawData_MSSQLMICROSOFTNumberNumberWID_MSSQLMICROSOFTNumberNumberWIDBrokerStatistics) SetPropertyForwardedMsgBytesPersec(value uint64) (err error) {
	return instance.SetProperty("ForwardedMsgBytesPersec", (value))
}

// GetForwardedMsgBytesPersec gets the value of ForwardedMsgBytesPersec for the instance
func (instance *Win32_PerfRawData_MSSQLMICROSOFTNumberNumberWID_MSSQLMICROSOFTNumberNumberWIDBrokerStatistics) GetPropertyForwardedMsgBytesPersec() (value uint64, err error) {
	retValue, err := instance.GetProperty("ForwardedMsgBytesPersec")
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

// SetForwardedMsgByteTotal sets the value of ForwardedMsgByteTotal for the instance
func (instance *Win32_PerfRawData_MSSQLMICROSOFTNumberNumberWID_MSSQLMICROSOFTNumberNumberWIDBrokerStatistics) SetPropertyForwardedMsgByteTotal(value uint64) (err error) {
	return instance.SetProperty("ForwardedMsgByteTotal", (value))
}

// GetForwardedMsgByteTotal gets the value of ForwardedMsgByteTotal for the instance
func (instance *Win32_PerfRawData_MSSQLMICROSOFTNumberNumberWID_MSSQLMICROSOFTNumberNumberWIDBrokerStatistics) GetPropertyForwardedMsgByteTotal() (value uint64, err error) {
	retValue, err := instance.GetProperty("ForwardedMsgByteTotal")
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

// SetForwardedMsgDiscardedTotal sets the value of ForwardedMsgDiscardedTotal for the instance
func (instance *Win32_PerfRawData_MSSQLMICROSOFTNumberNumberWID_MSSQLMICROSOFTNumberNumberWIDBrokerStatistics) SetPropertyForwardedMsgDiscardedTotal(value uint64) (err error) {
	return instance.SetProperty("ForwardedMsgDiscardedTotal", (value))
}

// GetForwardedMsgDiscardedTotal gets the value of ForwardedMsgDiscardedTotal for the instance
func (instance *Win32_PerfRawData_MSSQLMICROSOFTNumberNumberWID_MSSQLMICROSOFTNumberNumberWIDBrokerStatistics) GetPropertyForwardedMsgDiscardedTotal() (value uint64, err error) {
	retValue, err := instance.GetProperty("ForwardedMsgDiscardedTotal")
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

// SetForwardedMsgsDiscardedPersec sets the value of ForwardedMsgsDiscardedPersec for the instance
func (instance *Win32_PerfRawData_MSSQLMICROSOFTNumberNumberWID_MSSQLMICROSOFTNumberNumberWIDBrokerStatistics) SetPropertyForwardedMsgsDiscardedPersec(value uint64) (err error) {
	return instance.SetProperty("ForwardedMsgsDiscardedPersec", (value))
}

// GetForwardedMsgsDiscardedPersec gets the value of ForwardedMsgsDiscardedPersec for the instance
func (instance *Win32_PerfRawData_MSSQLMICROSOFTNumberNumberWID_MSSQLMICROSOFTNumberNumberWIDBrokerStatistics) GetPropertyForwardedMsgsDiscardedPersec() (value uint64, err error) {
	retValue, err := instance.GetProperty("ForwardedMsgsDiscardedPersec")
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

// SetForwardedPendingMsgBytes sets the value of ForwardedPendingMsgBytes for the instance
func (instance *Win32_PerfRawData_MSSQLMICROSOFTNumberNumberWID_MSSQLMICROSOFTNumberNumberWIDBrokerStatistics) SetPropertyForwardedPendingMsgBytes(value uint64) (err error) {
	return instance.SetProperty("ForwardedPendingMsgBytes", (value))
}

// GetForwardedPendingMsgBytes gets the value of ForwardedPendingMsgBytes for the instance
func (instance *Win32_PerfRawData_MSSQLMICROSOFTNumberNumberWID_MSSQLMICROSOFTNumberNumberWIDBrokerStatistics) GetPropertyForwardedPendingMsgBytes() (value uint64, err error) {
	retValue, err := instance.GetProperty("ForwardedPendingMsgBytes")
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

// SetForwardedPendingMsgCount sets the value of ForwardedPendingMsgCount for the instance
func (instance *Win32_PerfRawData_MSSQLMICROSOFTNumberNumberWID_MSSQLMICROSOFTNumberNumberWIDBrokerStatistics) SetPropertyForwardedPendingMsgCount(value uint64) (err error) {
	return instance.SetProperty("ForwardedPendingMsgCount", (value))
}

// GetForwardedPendingMsgCount gets the value of ForwardedPendingMsgCount for the instance
func (instance *Win32_PerfRawData_MSSQLMICROSOFTNumberNumberWID_MSSQLMICROSOFTNumberNumberWIDBrokerStatistics) GetPropertyForwardedPendingMsgCount() (value uint64, err error) {
	retValue, err := instance.GetProperty("ForwardedPendingMsgCount")
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

// SetSQLRECEIVEsPersec sets the value of SQLRECEIVEsPersec for the instance
func (instance *Win32_PerfRawData_MSSQLMICROSOFTNumberNumberWID_MSSQLMICROSOFTNumberNumberWIDBrokerStatistics) SetPropertySQLRECEIVEsPersec(value uint64) (err error) {
	return instance.SetProperty("SQLRECEIVEsPersec", (value))
}

// GetSQLRECEIVEsPersec gets the value of SQLRECEIVEsPersec for the instance
func (instance *Win32_PerfRawData_MSSQLMICROSOFTNumberNumberWID_MSSQLMICROSOFTNumberNumberWIDBrokerStatistics) GetPropertySQLRECEIVEsPersec() (value uint64, err error) {
	retValue, err := instance.GetProperty("SQLRECEIVEsPersec")
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

// SetSQLRECEIVETotal sets the value of SQLRECEIVETotal for the instance
func (instance *Win32_PerfRawData_MSSQLMICROSOFTNumberNumberWID_MSSQLMICROSOFTNumberNumberWIDBrokerStatistics) SetPropertySQLRECEIVETotal(value uint64) (err error) {
	return instance.SetProperty("SQLRECEIVETotal", (value))
}

// GetSQLRECEIVETotal gets the value of SQLRECEIVETotal for the instance
func (instance *Win32_PerfRawData_MSSQLMICROSOFTNumberNumberWID_MSSQLMICROSOFTNumberNumberWIDBrokerStatistics) GetPropertySQLRECEIVETotal() (value uint64, err error) {
	retValue, err := instance.GetProperty("SQLRECEIVETotal")
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

// SetSQLSENDsPersec sets the value of SQLSENDsPersec for the instance
func (instance *Win32_PerfRawData_MSSQLMICROSOFTNumberNumberWID_MSSQLMICROSOFTNumberNumberWIDBrokerStatistics) SetPropertySQLSENDsPersec(value uint64) (err error) {
	return instance.SetProperty("SQLSENDsPersec", (value))
}

// GetSQLSENDsPersec gets the value of SQLSENDsPersec for the instance
func (instance *Win32_PerfRawData_MSSQLMICROSOFTNumberNumberWID_MSSQLMICROSOFTNumberNumberWIDBrokerStatistics) GetPropertySQLSENDsPersec() (value uint64, err error) {
	retValue, err := instance.GetProperty("SQLSENDsPersec")
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

// SetSQLSENDTotal sets the value of SQLSENDTotal for the instance
func (instance *Win32_PerfRawData_MSSQLMICROSOFTNumberNumberWID_MSSQLMICROSOFTNumberNumberWIDBrokerStatistics) SetPropertySQLSENDTotal(value uint64) (err error) {
	return instance.SetProperty("SQLSENDTotal", (value))
}

// GetSQLSENDTotal gets the value of SQLSENDTotal for the instance
func (instance *Win32_PerfRawData_MSSQLMICROSOFTNumberNumberWID_MSSQLMICROSOFTNumberNumberWIDBrokerStatistics) GetPropertySQLSENDTotal() (value uint64, err error) {
	retValue, err := instance.GetProperty("SQLSENDTotal")
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
