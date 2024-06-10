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

// Win32_PerfFormattedData_MSSQLMICROSOFTNumberNumberWID_MSSQLMICROSOFTNumberNumberWIDGeneralStatistics struct
type Win32_PerfFormattedData_MSSQLMICROSOFTNumberNumberWID_MSSQLMICROSOFTNumberNumberWIDGeneralStatistics struct {
	*Win32_PerfFormattedData

	//
	ActiveTempTables uint64

	//
	ConnectionResetPersec uint64

	//
	EventNotificationsDelayedDrop uint64

	//
	HTTPAuthenticatedRequests uint64

	//
	LogicalConnections uint64

	//
	LoginsPersec uint64

	//
	LogoutsPersec uint64

	//
	MarsDeadlocks uint64

	//
	Nonatomicyieldrate uint64

	//
	Processesblocked uint64

	//
	SOAPEmptyRequests uint64

	//
	SOAPMethodInvocations uint64

	//
	SOAPSessionInitiateRequests uint64

	//
	SOAPSessionTerminateRequests uint64

	//
	SOAPSQLRequests uint64

	//
	SOAPWSDLRequests uint64

	//
	SQLTraceIOProviderLockWaits uint64

	//
	Tempdbrecoveryunitid uint64

	//
	Tempdbrowsetid uint64

	//
	TempTablesCreationRate uint64

	//
	TempTablesForDestruction uint64

	//
	TraceEventNotificationQueue uint64

	//
	Transactions uint64

	//
	UserConnections uint64
}

func NewWin32_PerfFormattedData_MSSQLMICROSOFTNumberNumberWID_MSSQLMICROSOFTNumberNumberWIDGeneralStatisticsEx1(instance *cim.WmiInstance) (newInstance *Win32_PerfFormattedData_MSSQLMICROSOFTNumberNumberWID_MSSQLMICROSOFTNumberNumberWIDGeneralStatistics, err error) {
	tmp, err := NewWin32_PerfFormattedDataEx1(instance)

	if err != nil {
		return
	}
	newInstance = &Win32_PerfFormattedData_MSSQLMICROSOFTNumberNumberWID_MSSQLMICROSOFTNumberNumberWIDGeneralStatistics{
		Win32_PerfFormattedData: tmp,
	}
	return
}

func NewWin32_PerfFormattedData_MSSQLMICROSOFTNumberNumberWID_MSSQLMICROSOFTNumberNumberWIDGeneralStatisticsEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *Win32_PerfFormattedData_MSSQLMICROSOFTNumberNumberWID_MSSQLMICROSOFTNumberNumberWIDGeneralStatistics, err error) {
	tmp, err := NewWin32_PerfFormattedDataEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &Win32_PerfFormattedData_MSSQLMICROSOFTNumberNumberWID_MSSQLMICROSOFTNumberNumberWIDGeneralStatistics{
		Win32_PerfFormattedData: tmp,
	}
	return
}

// SetActiveTempTables sets the value of ActiveTempTables for the instance
func (instance *Win32_PerfFormattedData_MSSQLMICROSOFTNumberNumberWID_MSSQLMICROSOFTNumberNumberWIDGeneralStatistics) SetPropertyActiveTempTables(value uint64) (err error) {
	return instance.SetProperty("ActiveTempTables", (value))
}

// GetActiveTempTables gets the value of ActiveTempTables for the instance
func (instance *Win32_PerfFormattedData_MSSQLMICROSOFTNumberNumberWID_MSSQLMICROSOFTNumberNumberWIDGeneralStatistics) GetPropertyActiveTempTables() (value uint64, err error) {
	retValue, err := instance.GetProperty("ActiveTempTables")
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

// SetConnectionResetPersec sets the value of ConnectionResetPersec for the instance
func (instance *Win32_PerfFormattedData_MSSQLMICROSOFTNumberNumberWID_MSSQLMICROSOFTNumberNumberWIDGeneralStatistics) SetPropertyConnectionResetPersec(value uint64) (err error) {
	return instance.SetProperty("ConnectionResetPersec", (value))
}

// GetConnectionResetPersec gets the value of ConnectionResetPersec for the instance
func (instance *Win32_PerfFormattedData_MSSQLMICROSOFTNumberNumberWID_MSSQLMICROSOFTNumberNumberWIDGeneralStatistics) GetPropertyConnectionResetPersec() (value uint64, err error) {
	retValue, err := instance.GetProperty("ConnectionResetPersec")
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

// SetEventNotificationsDelayedDrop sets the value of EventNotificationsDelayedDrop for the instance
func (instance *Win32_PerfFormattedData_MSSQLMICROSOFTNumberNumberWID_MSSQLMICROSOFTNumberNumberWIDGeneralStatistics) SetPropertyEventNotificationsDelayedDrop(value uint64) (err error) {
	return instance.SetProperty("EventNotificationsDelayedDrop", (value))
}

// GetEventNotificationsDelayedDrop gets the value of EventNotificationsDelayedDrop for the instance
func (instance *Win32_PerfFormattedData_MSSQLMICROSOFTNumberNumberWID_MSSQLMICROSOFTNumberNumberWIDGeneralStatistics) GetPropertyEventNotificationsDelayedDrop() (value uint64, err error) {
	retValue, err := instance.GetProperty("EventNotificationsDelayedDrop")
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

// SetHTTPAuthenticatedRequests sets the value of HTTPAuthenticatedRequests for the instance
func (instance *Win32_PerfFormattedData_MSSQLMICROSOFTNumberNumberWID_MSSQLMICROSOFTNumberNumberWIDGeneralStatistics) SetPropertyHTTPAuthenticatedRequests(value uint64) (err error) {
	return instance.SetProperty("HTTPAuthenticatedRequests", (value))
}

// GetHTTPAuthenticatedRequests gets the value of HTTPAuthenticatedRequests for the instance
func (instance *Win32_PerfFormattedData_MSSQLMICROSOFTNumberNumberWID_MSSQLMICROSOFTNumberNumberWIDGeneralStatistics) GetPropertyHTTPAuthenticatedRequests() (value uint64, err error) {
	retValue, err := instance.GetProperty("HTTPAuthenticatedRequests")
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

// SetLogicalConnections sets the value of LogicalConnections for the instance
func (instance *Win32_PerfFormattedData_MSSQLMICROSOFTNumberNumberWID_MSSQLMICROSOFTNumberNumberWIDGeneralStatistics) SetPropertyLogicalConnections(value uint64) (err error) {
	return instance.SetProperty("LogicalConnections", (value))
}

// GetLogicalConnections gets the value of LogicalConnections for the instance
func (instance *Win32_PerfFormattedData_MSSQLMICROSOFTNumberNumberWID_MSSQLMICROSOFTNumberNumberWIDGeneralStatistics) GetPropertyLogicalConnections() (value uint64, err error) {
	retValue, err := instance.GetProperty("LogicalConnections")
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

// SetLoginsPersec sets the value of LoginsPersec for the instance
func (instance *Win32_PerfFormattedData_MSSQLMICROSOFTNumberNumberWID_MSSQLMICROSOFTNumberNumberWIDGeneralStatistics) SetPropertyLoginsPersec(value uint64) (err error) {
	return instance.SetProperty("LoginsPersec", (value))
}

// GetLoginsPersec gets the value of LoginsPersec for the instance
func (instance *Win32_PerfFormattedData_MSSQLMICROSOFTNumberNumberWID_MSSQLMICROSOFTNumberNumberWIDGeneralStatistics) GetPropertyLoginsPersec() (value uint64, err error) {
	retValue, err := instance.GetProperty("LoginsPersec")
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

// SetLogoutsPersec sets the value of LogoutsPersec for the instance
func (instance *Win32_PerfFormattedData_MSSQLMICROSOFTNumberNumberWID_MSSQLMICROSOFTNumberNumberWIDGeneralStatistics) SetPropertyLogoutsPersec(value uint64) (err error) {
	return instance.SetProperty("LogoutsPersec", (value))
}

// GetLogoutsPersec gets the value of LogoutsPersec for the instance
func (instance *Win32_PerfFormattedData_MSSQLMICROSOFTNumberNumberWID_MSSQLMICROSOFTNumberNumberWIDGeneralStatistics) GetPropertyLogoutsPersec() (value uint64, err error) {
	retValue, err := instance.GetProperty("LogoutsPersec")
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

// SetMarsDeadlocks sets the value of MarsDeadlocks for the instance
func (instance *Win32_PerfFormattedData_MSSQLMICROSOFTNumberNumberWID_MSSQLMICROSOFTNumberNumberWIDGeneralStatistics) SetPropertyMarsDeadlocks(value uint64) (err error) {
	return instance.SetProperty("MarsDeadlocks", (value))
}

// GetMarsDeadlocks gets the value of MarsDeadlocks for the instance
func (instance *Win32_PerfFormattedData_MSSQLMICROSOFTNumberNumberWID_MSSQLMICROSOFTNumberNumberWIDGeneralStatistics) GetPropertyMarsDeadlocks() (value uint64, err error) {
	retValue, err := instance.GetProperty("MarsDeadlocks")
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

// SetNonatomicyieldrate sets the value of Nonatomicyieldrate for the instance
func (instance *Win32_PerfFormattedData_MSSQLMICROSOFTNumberNumberWID_MSSQLMICROSOFTNumberNumberWIDGeneralStatistics) SetPropertyNonatomicyieldrate(value uint64) (err error) {
	return instance.SetProperty("Nonatomicyieldrate", (value))
}

// GetNonatomicyieldrate gets the value of Nonatomicyieldrate for the instance
func (instance *Win32_PerfFormattedData_MSSQLMICROSOFTNumberNumberWID_MSSQLMICROSOFTNumberNumberWIDGeneralStatistics) GetPropertyNonatomicyieldrate() (value uint64, err error) {
	retValue, err := instance.GetProperty("Nonatomicyieldrate")
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

// SetProcessesblocked sets the value of Processesblocked for the instance
func (instance *Win32_PerfFormattedData_MSSQLMICROSOFTNumberNumberWID_MSSQLMICROSOFTNumberNumberWIDGeneralStatistics) SetPropertyProcessesblocked(value uint64) (err error) {
	return instance.SetProperty("Processesblocked", (value))
}

// GetProcessesblocked gets the value of Processesblocked for the instance
func (instance *Win32_PerfFormattedData_MSSQLMICROSOFTNumberNumberWID_MSSQLMICROSOFTNumberNumberWIDGeneralStatistics) GetPropertyProcessesblocked() (value uint64, err error) {
	retValue, err := instance.GetProperty("Processesblocked")
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

// SetSOAPEmptyRequests sets the value of SOAPEmptyRequests for the instance
func (instance *Win32_PerfFormattedData_MSSQLMICROSOFTNumberNumberWID_MSSQLMICROSOFTNumberNumberWIDGeneralStatistics) SetPropertySOAPEmptyRequests(value uint64) (err error) {
	return instance.SetProperty("SOAPEmptyRequests", (value))
}

// GetSOAPEmptyRequests gets the value of SOAPEmptyRequests for the instance
func (instance *Win32_PerfFormattedData_MSSQLMICROSOFTNumberNumberWID_MSSQLMICROSOFTNumberNumberWIDGeneralStatistics) GetPropertySOAPEmptyRequests() (value uint64, err error) {
	retValue, err := instance.GetProperty("SOAPEmptyRequests")
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

// SetSOAPMethodInvocations sets the value of SOAPMethodInvocations for the instance
func (instance *Win32_PerfFormattedData_MSSQLMICROSOFTNumberNumberWID_MSSQLMICROSOFTNumberNumberWIDGeneralStatistics) SetPropertySOAPMethodInvocations(value uint64) (err error) {
	return instance.SetProperty("SOAPMethodInvocations", (value))
}

// GetSOAPMethodInvocations gets the value of SOAPMethodInvocations for the instance
func (instance *Win32_PerfFormattedData_MSSQLMICROSOFTNumberNumberWID_MSSQLMICROSOFTNumberNumberWIDGeneralStatistics) GetPropertySOAPMethodInvocations() (value uint64, err error) {
	retValue, err := instance.GetProperty("SOAPMethodInvocations")
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

// SetSOAPSessionInitiateRequests sets the value of SOAPSessionInitiateRequests for the instance
func (instance *Win32_PerfFormattedData_MSSQLMICROSOFTNumberNumberWID_MSSQLMICROSOFTNumberNumberWIDGeneralStatistics) SetPropertySOAPSessionInitiateRequests(value uint64) (err error) {
	return instance.SetProperty("SOAPSessionInitiateRequests", (value))
}

// GetSOAPSessionInitiateRequests gets the value of SOAPSessionInitiateRequests for the instance
func (instance *Win32_PerfFormattedData_MSSQLMICROSOFTNumberNumberWID_MSSQLMICROSOFTNumberNumberWIDGeneralStatistics) GetPropertySOAPSessionInitiateRequests() (value uint64, err error) {
	retValue, err := instance.GetProperty("SOAPSessionInitiateRequests")
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

// SetSOAPSessionTerminateRequests sets the value of SOAPSessionTerminateRequests for the instance
func (instance *Win32_PerfFormattedData_MSSQLMICROSOFTNumberNumberWID_MSSQLMICROSOFTNumberNumberWIDGeneralStatistics) SetPropertySOAPSessionTerminateRequests(value uint64) (err error) {
	return instance.SetProperty("SOAPSessionTerminateRequests", (value))
}

// GetSOAPSessionTerminateRequests gets the value of SOAPSessionTerminateRequests for the instance
func (instance *Win32_PerfFormattedData_MSSQLMICROSOFTNumberNumberWID_MSSQLMICROSOFTNumberNumberWIDGeneralStatistics) GetPropertySOAPSessionTerminateRequests() (value uint64, err error) {
	retValue, err := instance.GetProperty("SOAPSessionTerminateRequests")
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

// SetSOAPSQLRequests sets the value of SOAPSQLRequests for the instance
func (instance *Win32_PerfFormattedData_MSSQLMICROSOFTNumberNumberWID_MSSQLMICROSOFTNumberNumberWIDGeneralStatistics) SetPropertySOAPSQLRequests(value uint64) (err error) {
	return instance.SetProperty("SOAPSQLRequests", (value))
}

// GetSOAPSQLRequests gets the value of SOAPSQLRequests for the instance
func (instance *Win32_PerfFormattedData_MSSQLMICROSOFTNumberNumberWID_MSSQLMICROSOFTNumberNumberWIDGeneralStatistics) GetPropertySOAPSQLRequests() (value uint64, err error) {
	retValue, err := instance.GetProperty("SOAPSQLRequests")
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

// SetSOAPWSDLRequests sets the value of SOAPWSDLRequests for the instance
func (instance *Win32_PerfFormattedData_MSSQLMICROSOFTNumberNumberWID_MSSQLMICROSOFTNumberNumberWIDGeneralStatistics) SetPropertySOAPWSDLRequests(value uint64) (err error) {
	return instance.SetProperty("SOAPWSDLRequests", (value))
}

// GetSOAPWSDLRequests gets the value of SOAPWSDLRequests for the instance
func (instance *Win32_PerfFormattedData_MSSQLMICROSOFTNumberNumberWID_MSSQLMICROSOFTNumberNumberWIDGeneralStatistics) GetPropertySOAPWSDLRequests() (value uint64, err error) {
	retValue, err := instance.GetProperty("SOAPWSDLRequests")
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

// SetSQLTraceIOProviderLockWaits sets the value of SQLTraceIOProviderLockWaits for the instance
func (instance *Win32_PerfFormattedData_MSSQLMICROSOFTNumberNumberWID_MSSQLMICROSOFTNumberNumberWIDGeneralStatistics) SetPropertySQLTraceIOProviderLockWaits(value uint64) (err error) {
	return instance.SetProperty("SQLTraceIOProviderLockWaits", (value))
}

// GetSQLTraceIOProviderLockWaits gets the value of SQLTraceIOProviderLockWaits for the instance
func (instance *Win32_PerfFormattedData_MSSQLMICROSOFTNumberNumberWID_MSSQLMICROSOFTNumberNumberWIDGeneralStatistics) GetPropertySQLTraceIOProviderLockWaits() (value uint64, err error) {
	retValue, err := instance.GetProperty("SQLTraceIOProviderLockWaits")
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

// SetTempdbrecoveryunitid sets the value of Tempdbrecoveryunitid for the instance
func (instance *Win32_PerfFormattedData_MSSQLMICROSOFTNumberNumberWID_MSSQLMICROSOFTNumberNumberWIDGeneralStatistics) SetPropertyTempdbrecoveryunitid(value uint64) (err error) {
	return instance.SetProperty("Tempdbrecoveryunitid", (value))
}

// GetTempdbrecoveryunitid gets the value of Tempdbrecoveryunitid for the instance
func (instance *Win32_PerfFormattedData_MSSQLMICROSOFTNumberNumberWID_MSSQLMICROSOFTNumberNumberWIDGeneralStatistics) GetPropertyTempdbrecoveryunitid() (value uint64, err error) {
	retValue, err := instance.GetProperty("Tempdbrecoveryunitid")
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

// SetTempdbrowsetid sets the value of Tempdbrowsetid for the instance
func (instance *Win32_PerfFormattedData_MSSQLMICROSOFTNumberNumberWID_MSSQLMICROSOFTNumberNumberWIDGeneralStatistics) SetPropertyTempdbrowsetid(value uint64) (err error) {
	return instance.SetProperty("Tempdbrowsetid", (value))
}

// GetTempdbrowsetid gets the value of Tempdbrowsetid for the instance
func (instance *Win32_PerfFormattedData_MSSQLMICROSOFTNumberNumberWID_MSSQLMICROSOFTNumberNumberWIDGeneralStatistics) GetPropertyTempdbrowsetid() (value uint64, err error) {
	retValue, err := instance.GetProperty("Tempdbrowsetid")
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

// SetTempTablesCreationRate sets the value of TempTablesCreationRate for the instance
func (instance *Win32_PerfFormattedData_MSSQLMICROSOFTNumberNumberWID_MSSQLMICROSOFTNumberNumberWIDGeneralStatistics) SetPropertyTempTablesCreationRate(value uint64) (err error) {
	return instance.SetProperty("TempTablesCreationRate", (value))
}

// GetTempTablesCreationRate gets the value of TempTablesCreationRate for the instance
func (instance *Win32_PerfFormattedData_MSSQLMICROSOFTNumberNumberWID_MSSQLMICROSOFTNumberNumberWIDGeneralStatistics) GetPropertyTempTablesCreationRate() (value uint64, err error) {
	retValue, err := instance.GetProperty("TempTablesCreationRate")
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

// SetTempTablesForDestruction sets the value of TempTablesForDestruction for the instance
func (instance *Win32_PerfFormattedData_MSSQLMICROSOFTNumberNumberWID_MSSQLMICROSOFTNumberNumberWIDGeneralStatistics) SetPropertyTempTablesForDestruction(value uint64) (err error) {
	return instance.SetProperty("TempTablesForDestruction", (value))
}

// GetTempTablesForDestruction gets the value of TempTablesForDestruction for the instance
func (instance *Win32_PerfFormattedData_MSSQLMICROSOFTNumberNumberWID_MSSQLMICROSOFTNumberNumberWIDGeneralStatistics) GetPropertyTempTablesForDestruction() (value uint64, err error) {
	retValue, err := instance.GetProperty("TempTablesForDestruction")
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

// SetTraceEventNotificationQueue sets the value of TraceEventNotificationQueue for the instance
func (instance *Win32_PerfFormattedData_MSSQLMICROSOFTNumberNumberWID_MSSQLMICROSOFTNumberNumberWIDGeneralStatistics) SetPropertyTraceEventNotificationQueue(value uint64) (err error) {
	return instance.SetProperty("TraceEventNotificationQueue", (value))
}

// GetTraceEventNotificationQueue gets the value of TraceEventNotificationQueue for the instance
func (instance *Win32_PerfFormattedData_MSSQLMICROSOFTNumberNumberWID_MSSQLMICROSOFTNumberNumberWIDGeneralStatistics) GetPropertyTraceEventNotificationQueue() (value uint64, err error) {
	retValue, err := instance.GetProperty("TraceEventNotificationQueue")
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

// SetTransactions sets the value of Transactions for the instance
func (instance *Win32_PerfFormattedData_MSSQLMICROSOFTNumberNumberWID_MSSQLMICROSOFTNumberNumberWIDGeneralStatistics) SetPropertyTransactions(value uint64) (err error) {
	return instance.SetProperty("Transactions", (value))
}

// GetTransactions gets the value of Transactions for the instance
func (instance *Win32_PerfFormattedData_MSSQLMICROSOFTNumberNumberWID_MSSQLMICROSOFTNumberNumberWIDGeneralStatistics) GetPropertyTransactions() (value uint64, err error) {
	retValue, err := instance.GetProperty("Transactions")
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

// SetUserConnections sets the value of UserConnections for the instance
func (instance *Win32_PerfFormattedData_MSSQLMICROSOFTNumberNumberWID_MSSQLMICROSOFTNumberNumberWIDGeneralStatistics) SetPropertyUserConnections(value uint64) (err error) {
	return instance.SetProperty("UserConnections", (value))
}

// GetUserConnections gets the value of UserConnections for the instance
func (instance *Win32_PerfFormattedData_MSSQLMICROSOFTNumberNumberWID_MSSQLMICROSOFTNumberNumberWIDGeneralStatistics) GetPropertyUserConnections() (value uint64, err error) {
	retValue, err := instance.GetProperty("UserConnections")
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
