// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 10/28/2026 using wmigen
//      Source root.CIMV2
//////////////////////////////////////////////
package cimv2

import (
	"github.com/microsoft/wmi/pkg/base/query"
	"github.com/microsoft/wmi/pkg/errors"
	cim "github.com/microsoft/wmi/pkg/wmiinstance"
	"reflect"
)

// Win32_PerfFormattedData_Counters_BluetoothRadio struct
type Win32_PerfFormattedData_Counters_BluetoothRadio struct {
	*Win32_PerfFormattedData

	//
	A2dpOffloadConnections uint32

	//
	ACLflusheventsPersec uint32

	//
	ActiveLEAdvertisementMonitorscount uint32

	//
	ClassicACLbytesreadPersec uint32

	//
	ClassicACLbyteswrittenPersec uint32

	//
	ClassicACLConnections uint32

	//
	ClassicACLwritecredits uint32

	//
	CreatedLEBIGCount uint32

	//
	EstablishedBIGSyncCount uint32

	//
	EstablishedLECISCount uint32

	//
	EstablishedLEPeriodicAdvertisingSyncCount uint32

	//
	InquiryScanDutyCyclePercent uint32

	//
	InquiryScanInterval uint32

	//
	InquiryScanWindow uint32

	//
	LEACLbytesreadPersec uint32

	//
	LEACLbyteswrittenPersec uint32

	//
	LEACLConnections uint32

	//
	LEACLwritecredits uint32

	//
	LEAdvertisingHandlesAllocated uint32

	//
	LEAdvertisingSetsEnabled uint32

	//
	LEBIGInfoReporteventsreceivedPersec uint32

	//
	LEPeriodicAdvertisingReportbytesreceivedPersec uint32

	//
	LEPeriodicAdvertisingReporteventsreceivedPersec uint32

	//
	LEScanDutyCyclePercentCodedPhy uint32

	//
	LEScanDutyCyclePercentUncoded1MPhy uint32

	//
	LEScanIntervalCodedPhy uint32

	//
	LEScanIntervalUncoded1MPhy uint32

	//
	LEScanWindowCodedPhy uint32

	//
	LEScanWindowUncoded1MPhy uint32

	//
	PageScanDutyCyclePercent uint32

	//
	PageScanInterval uint32

	//
	PageScanWindow uint32

	//
	ReportedLEMonitoredDevices uint32

	//
	SCObytesreadPersec uint32

	//
	SCObyteswrittenPersec uint32

	//
	SCOConnections uint32

	//
	SidebandSCOConnections uint32
}

func NewWin32_PerfFormattedData_Counters_BluetoothRadioEx1(instance *cim.WmiInstance) (newInstance *Win32_PerfFormattedData_Counters_BluetoothRadio, err error) {
	tmp, err := NewWin32_PerfFormattedDataEx1(instance)

	if err != nil {
		return
	}
	newInstance = &Win32_PerfFormattedData_Counters_BluetoothRadio{
		Win32_PerfFormattedData: tmp,
	}
	return
}

func NewWin32_PerfFormattedData_Counters_BluetoothRadioEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *Win32_PerfFormattedData_Counters_BluetoothRadio, err error) {
	tmp, err := NewWin32_PerfFormattedDataEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &Win32_PerfFormattedData_Counters_BluetoothRadio{
		Win32_PerfFormattedData: tmp,
	}
	return
}

// SetA2dpOffloadConnections sets the value of A2dpOffloadConnections for the instance
func (instance *Win32_PerfFormattedData_Counters_BluetoothRadio) SetPropertyA2dpOffloadConnections(value uint32) (err error) {
	return instance.SetProperty("A2dpOffloadConnections", (value))
}

// GetA2dpOffloadConnections gets the value of A2dpOffloadConnections for the instance
func (instance *Win32_PerfFormattedData_Counters_BluetoothRadio) GetPropertyA2dpOffloadConnections() (value uint32, err error) {
	retValue, err := instance.GetProperty("A2dpOffloadConnections")
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

// SetACLflusheventsPersec sets the value of ACLflusheventsPersec for the instance
func (instance *Win32_PerfFormattedData_Counters_BluetoothRadio) SetPropertyACLflusheventsPersec(value uint32) (err error) {
	return instance.SetProperty("ACLflusheventsPersec", (value))
}

// GetACLflusheventsPersec gets the value of ACLflusheventsPersec for the instance
func (instance *Win32_PerfFormattedData_Counters_BluetoothRadio) GetPropertyACLflusheventsPersec() (value uint32, err error) {
	retValue, err := instance.GetProperty("ACLflusheventsPersec")
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

// SetActiveLEAdvertisementMonitorscount sets the value of ActiveLEAdvertisementMonitorscount for the instance
func (instance *Win32_PerfFormattedData_Counters_BluetoothRadio) SetPropertyActiveLEAdvertisementMonitorscount(value uint32) (err error) {
	return instance.SetProperty("ActiveLEAdvertisementMonitorscount", (value))
}

// GetActiveLEAdvertisementMonitorscount gets the value of ActiveLEAdvertisementMonitorscount for the instance
func (instance *Win32_PerfFormattedData_Counters_BluetoothRadio) GetPropertyActiveLEAdvertisementMonitorscount() (value uint32, err error) {
	retValue, err := instance.GetProperty("ActiveLEAdvertisementMonitorscount")
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

// SetClassicACLbytesreadPersec sets the value of ClassicACLbytesreadPersec for the instance
func (instance *Win32_PerfFormattedData_Counters_BluetoothRadio) SetPropertyClassicACLbytesreadPersec(value uint32) (err error) {
	return instance.SetProperty("ClassicACLbytesreadPersec", (value))
}

// GetClassicACLbytesreadPersec gets the value of ClassicACLbytesreadPersec for the instance
func (instance *Win32_PerfFormattedData_Counters_BluetoothRadio) GetPropertyClassicACLbytesreadPersec() (value uint32, err error) {
	retValue, err := instance.GetProperty("ClassicACLbytesreadPersec")
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

// SetClassicACLbyteswrittenPersec sets the value of ClassicACLbyteswrittenPersec for the instance
func (instance *Win32_PerfFormattedData_Counters_BluetoothRadio) SetPropertyClassicACLbyteswrittenPersec(value uint32) (err error) {
	return instance.SetProperty("ClassicACLbyteswrittenPersec", (value))
}

// GetClassicACLbyteswrittenPersec gets the value of ClassicACLbyteswrittenPersec for the instance
func (instance *Win32_PerfFormattedData_Counters_BluetoothRadio) GetPropertyClassicACLbyteswrittenPersec() (value uint32, err error) {
	retValue, err := instance.GetProperty("ClassicACLbyteswrittenPersec")
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

// SetClassicACLConnections sets the value of ClassicACLConnections for the instance
func (instance *Win32_PerfFormattedData_Counters_BluetoothRadio) SetPropertyClassicACLConnections(value uint32) (err error) {
	return instance.SetProperty("ClassicACLConnections", (value))
}

// GetClassicACLConnections gets the value of ClassicACLConnections for the instance
func (instance *Win32_PerfFormattedData_Counters_BluetoothRadio) GetPropertyClassicACLConnections() (value uint32, err error) {
	retValue, err := instance.GetProperty("ClassicACLConnections")
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

// SetClassicACLwritecredits sets the value of ClassicACLwritecredits for the instance
func (instance *Win32_PerfFormattedData_Counters_BluetoothRadio) SetPropertyClassicACLwritecredits(value uint32) (err error) {
	return instance.SetProperty("ClassicACLwritecredits", (value))
}

// GetClassicACLwritecredits gets the value of ClassicACLwritecredits for the instance
func (instance *Win32_PerfFormattedData_Counters_BluetoothRadio) GetPropertyClassicACLwritecredits() (value uint32, err error) {
	retValue, err := instance.GetProperty("ClassicACLwritecredits")
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

// SetCreatedLEBIGCount sets the value of CreatedLEBIGCount for the instance
func (instance *Win32_PerfFormattedData_Counters_BluetoothRadio) SetPropertyCreatedLEBIGCount(value uint32) (err error) {
	return instance.SetProperty("CreatedLEBIGCount", (value))
}

// GetCreatedLEBIGCount gets the value of CreatedLEBIGCount for the instance
func (instance *Win32_PerfFormattedData_Counters_BluetoothRadio) GetPropertyCreatedLEBIGCount() (value uint32, err error) {
	retValue, err := instance.GetProperty("CreatedLEBIGCount")
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

// SetEstablishedBIGSyncCount sets the value of EstablishedBIGSyncCount for the instance
func (instance *Win32_PerfFormattedData_Counters_BluetoothRadio) SetPropertyEstablishedBIGSyncCount(value uint32) (err error) {
	return instance.SetProperty("EstablishedBIGSyncCount", (value))
}

// GetEstablishedBIGSyncCount gets the value of EstablishedBIGSyncCount for the instance
func (instance *Win32_PerfFormattedData_Counters_BluetoothRadio) GetPropertyEstablishedBIGSyncCount() (value uint32, err error) {
	retValue, err := instance.GetProperty("EstablishedBIGSyncCount")
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

// SetEstablishedLECISCount sets the value of EstablishedLECISCount for the instance
func (instance *Win32_PerfFormattedData_Counters_BluetoothRadio) SetPropertyEstablishedLECISCount(value uint32) (err error) {
	return instance.SetProperty("EstablishedLECISCount", (value))
}

// GetEstablishedLECISCount gets the value of EstablishedLECISCount for the instance
func (instance *Win32_PerfFormattedData_Counters_BluetoothRadio) GetPropertyEstablishedLECISCount() (value uint32, err error) {
	retValue, err := instance.GetProperty("EstablishedLECISCount")
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

// SetEstablishedLEPeriodicAdvertisingSyncCount sets the value of EstablishedLEPeriodicAdvertisingSyncCount for the instance
func (instance *Win32_PerfFormattedData_Counters_BluetoothRadio) SetPropertyEstablishedLEPeriodicAdvertisingSyncCount(value uint32) (err error) {
	return instance.SetProperty("EstablishedLEPeriodicAdvertisingSyncCount", (value))
}

// GetEstablishedLEPeriodicAdvertisingSyncCount gets the value of EstablishedLEPeriodicAdvertisingSyncCount for the instance
func (instance *Win32_PerfFormattedData_Counters_BluetoothRadio) GetPropertyEstablishedLEPeriodicAdvertisingSyncCount() (value uint32, err error) {
	retValue, err := instance.GetProperty("EstablishedLEPeriodicAdvertisingSyncCount")
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

// SetInquiryScanDutyCyclePercent sets the value of InquiryScanDutyCyclePercent for the instance
func (instance *Win32_PerfFormattedData_Counters_BluetoothRadio) SetPropertyInquiryScanDutyCyclePercent(value uint32) (err error) {
	return instance.SetProperty("InquiryScanDutyCyclePercent", (value))
}

// GetInquiryScanDutyCyclePercent gets the value of InquiryScanDutyCyclePercent for the instance
func (instance *Win32_PerfFormattedData_Counters_BluetoothRadio) GetPropertyInquiryScanDutyCyclePercent() (value uint32, err error) {
	retValue, err := instance.GetProperty("InquiryScanDutyCyclePercent")
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

// SetInquiryScanInterval sets the value of InquiryScanInterval for the instance
func (instance *Win32_PerfFormattedData_Counters_BluetoothRadio) SetPropertyInquiryScanInterval(value uint32) (err error) {
	return instance.SetProperty("InquiryScanInterval", (value))
}

// GetInquiryScanInterval gets the value of InquiryScanInterval for the instance
func (instance *Win32_PerfFormattedData_Counters_BluetoothRadio) GetPropertyInquiryScanInterval() (value uint32, err error) {
	retValue, err := instance.GetProperty("InquiryScanInterval")
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

// SetInquiryScanWindow sets the value of InquiryScanWindow for the instance
func (instance *Win32_PerfFormattedData_Counters_BluetoothRadio) SetPropertyInquiryScanWindow(value uint32) (err error) {
	return instance.SetProperty("InquiryScanWindow", (value))
}

// GetInquiryScanWindow gets the value of InquiryScanWindow for the instance
func (instance *Win32_PerfFormattedData_Counters_BluetoothRadio) GetPropertyInquiryScanWindow() (value uint32, err error) {
	retValue, err := instance.GetProperty("InquiryScanWindow")
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

// SetLEACLbytesreadPersec sets the value of LEACLbytesreadPersec for the instance
func (instance *Win32_PerfFormattedData_Counters_BluetoothRadio) SetPropertyLEACLbytesreadPersec(value uint32) (err error) {
	return instance.SetProperty("LEACLbytesreadPersec", (value))
}

// GetLEACLbytesreadPersec gets the value of LEACLbytesreadPersec for the instance
func (instance *Win32_PerfFormattedData_Counters_BluetoothRadio) GetPropertyLEACLbytesreadPersec() (value uint32, err error) {
	retValue, err := instance.GetProperty("LEACLbytesreadPersec")
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

// SetLEACLbyteswrittenPersec sets the value of LEACLbyteswrittenPersec for the instance
func (instance *Win32_PerfFormattedData_Counters_BluetoothRadio) SetPropertyLEACLbyteswrittenPersec(value uint32) (err error) {
	return instance.SetProperty("LEACLbyteswrittenPersec", (value))
}

// GetLEACLbyteswrittenPersec gets the value of LEACLbyteswrittenPersec for the instance
func (instance *Win32_PerfFormattedData_Counters_BluetoothRadio) GetPropertyLEACLbyteswrittenPersec() (value uint32, err error) {
	retValue, err := instance.GetProperty("LEACLbyteswrittenPersec")
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

// SetLEACLConnections sets the value of LEACLConnections for the instance
func (instance *Win32_PerfFormattedData_Counters_BluetoothRadio) SetPropertyLEACLConnections(value uint32) (err error) {
	return instance.SetProperty("LEACLConnections", (value))
}

// GetLEACLConnections gets the value of LEACLConnections for the instance
func (instance *Win32_PerfFormattedData_Counters_BluetoothRadio) GetPropertyLEACLConnections() (value uint32, err error) {
	retValue, err := instance.GetProperty("LEACLConnections")
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

// SetLEACLwritecredits sets the value of LEACLwritecredits for the instance
func (instance *Win32_PerfFormattedData_Counters_BluetoothRadio) SetPropertyLEACLwritecredits(value uint32) (err error) {
	return instance.SetProperty("LEACLwritecredits", (value))
}

// GetLEACLwritecredits gets the value of LEACLwritecredits for the instance
func (instance *Win32_PerfFormattedData_Counters_BluetoothRadio) GetPropertyLEACLwritecredits() (value uint32, err error) {
	retValue, err := instance.GetProperty("LEACLwritecredits")
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

// SetLEAdvertisingHandlesAllocated sets the value of LEAdvertisingHandlesAllocated for the instance
func (instance *Win32_PerfFormattedData_Counters_BluetoothRadio) SetPropertyLEAdvertisingHandlesAllocated(value uint32) (err error) {
	return instance.SetProperty("LEAdvertisingHandlesAllocated", (value))
}

// GetLEAdvertisingHandlesAllocated gets the value of LEAdvertisingHandlesAllocated for the instance
func (instance *Win32_PerfFormattedData_Counters_BluetoothRadio) GetPropertyLEAdvertisingHandlesAllocated() (value uint32, err error) {
	retValue, err := instance.GetProperty("LEAdvertisingHandlesAllocated")
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

// SetLEAdvertisingSetsEnabled sets the value of LEAdvertisingSetsEnabled for the instance
func (instance *Win32_PerfFormattedData_Counters_BluetoothRadio) SetPropertyLEAdvertisingSetsEnabled(value uint32) (err error) {
	return instance.SetProperty("LEAdvertisingSetsEnabled", (value))
}

// GetLEAdvertisingSetsEnabled gets the value of LEAdvertisingSetsEnabled for the instance
func (instance *Win32_PerfFormattedData_Counters_BluetoothRadio) GetPropertyLEAdvertisingSetsEnabled() (value uint32, err error) {
	retValue, err := instance.GetProperty("LEAdvertisingSetsEnabled")
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

// SetLEBIGInfoReporteventsreceivedPersec sets the value of LEBIGInfoReporteventsreceivedPersec for the instance
func (instance *Win32_PerfFormattedData_Counters_BluetoothRadio) SetPropertyLEBIGInfoReporteventsreceivedPersec(value uint32) (err error) {
	return instance.SetProperty("LEBIGInfoReporteventsreceivedPersec", (value))
}

// GetLEBIGInfoReporteventsreceivedPersec gets the value of LEBIGInfoReporteventsreceivedPersec for the instance
func (instance *Win32_PerfFormattedData_Counters_BluetoothRadio) GetPropertyLEBIGInfoReporteventsreceivedPersec() (value uint32, err error) {
	retValue, err := instance.GetProperty("LEBIGInfoReporteventsreceivedPersec")
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

// SetLEPeriodicAdvertisingReportbytesreceivedPersec sets the value of LEPeriodicAdvertisingReportbytesreceivedPersec for the instance
func (instance *Win32_PerfFormattedData_Counters_BluetoothRadio) SetPropertyLEPeriodicAdvertisingReportbytesreceivedPersec(value uint32) (err error) {
	return instance.SetProperty("LEPeriodicAdvertisingReportbytesreceivedPersec", (value))
}

// GetLEPeriodicAdvertisingReportbytesreceivedPersec gets the value of LEPeriodicAdvertisingReportbytesreceivedPersec for the instance
func (instance *Win32_PerfFormattedData_Counters_BluetoothRadio) GetPropertyLEPeriodicAdvertisingReportbytesreceivedPersec() (value uint32, err error) {
	retValue, err := instance.GetProperty("LEPeriodicAdvertisingReportbytesreceivedPersec")
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

// SetLEPeriodicAdvertisingReporteventsreceivedPersec sets the value of LEPeriodicAdvertisingReporteventsreceivedPersec for the instance
func (instance *Win32_PerfFormattedData_Counters_BluetoothRadio) SetPropertyLEPeriodicAdvertisingReporteventsreceivedPersec(value uint32) (err error) {
	return instance.SetProperty("LEPeriodicAdvertisingReporteventsreceivedPersec", (value))
}

// GetLEPeriodicAdvertisingReporteventsreceivedPersec gets the value of LEPeriodicAdvertisingReporteventsreceivedPersec for the instance
func (instance *Win32_PerfFormattedData_Counters_BluetoothRadio) GetPropertyLEPeriodicAdvertisingReporteventsreceivedPersec() (value uint32, err error) {
	retValue, err := instance.GetProperty("LEPeriodicAdvertisingReporteventsreceivedPersec")
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

// SetLEScanDutyCyclePercentCodedPhy sets the value of LEScanDutyCyclePercentCodedPhy for the instance
func (instance *Win32_PerfFormattedData_Counters_BluetoothRadio) SetPropertyLEScanDutyCyclePercentCodedPhy(value uint32) (err error) {
	return instance.SetProperty("LEScanDutyCyclePercentCodedPhy", (value))
}

// GetLEScanDutyCyclePercentCodedPhy gets the value of LEScanDutyCyclePercentCodedPhy for the instance
func (instance *Win32_PerfFormattedData_Counters_BluetoothRadio) GetPropertyLEScanDutyCyclePercentCodedPhy() (value uint32, err error) {
	retValue, err := instance.GetProperty("LEScanDutyCyclePercentCodedPhy")
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

// SetLEScanDutyCyclePercentUncoded1MPhy sets the value of LEScanDutyCyclePercentUncoded1MPhy for the instance
func (instance *Win32_PerfFormattedData_Counters_BluetoothRadio) SetPropertyLEScanDutyCyclePercentUncoded1MPhy(value uint32) (err error) {
	return instance.SetProperty("LEScanDutyCyclePercentUncoded1MPhy", (value))
}

// GetLEScanDutyCyclePercentUncoded1MPhy gets the value of LEScanDutyCyclePercentUncoded1MPhy for the instance
func (instance *Win32_PerfFormattedData_Counters_BluetoothRadio) GetPropertyLEScanDutyCyclePercentUncoded1MPhy() (value uint32, err error) {
	retValue, err := instance.GetProperty("LEScanDutyCyclePercentUncoded1MPhy")
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

// SetLEScanIntervalCodedPhy sets the value of LEScanIntervalCodedPhy for the instance
func (instance *Win32_PerfFormattedData_Counters_BluetoothRadio) SetPropertyLEScanIntervalCodedPhy(value uint32) (err error) {
	return instance.SetProperty("LEScanIntervalCodedPhy", (value))
}

// GetLEScanIntervalCodedPhy gets the value of LEScanIntervalCodedPhy for the instance
func (instance *Win32_PerfFormattedData_Counters_BluetoothRadio) GetPropertyLEScanIntervalCodedPhy() (value uint32, err error) {
	retValue, err := instance.GetProperty("LEScanIntervalCodedPhy")
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

// SetLEScanIntervalUncoded1MPhy sets the value of LEScanIntervalUncoded1MPhy for the instance
func (instance *Win32_PerfFormattedData_Counters_BluetoothRadio) SetPropertyLEScanIntervalUncoded1MPhy(value uint32) (err error) {
	return instance.SetProperty("LEScanIntervalUncoded1MPhy", (value))
}

// GetLEScanIntervalUncoded1MPhy gets the value of LEScanIntervalUncoded1MPhy for the instance
func (instance *Win32_PerfFormattedData_Counters_BluetoothRadio) GetPropertyLEScanIntervalUncoded1MPhy() (value uint32, err error) {
	retValue, err := instance.GetProperty("LEScanIntervalUncoded1MPhy")
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

// SetLEScanWindowCodedPhy sets the value of LEScanWindowCodedPhy for the instance
func (instance *Win32_PerfFormattedData_Counters_BluetoothRadio) SetPropertyLEScanWindowCodedPhy(value uint32) (err error) {
	return instance.SetProperty("LEScanWindowCodedPhy", (value))
}

// GetLEScanWindowCodedPhy gets the value of LEScanWindowCodedPhy for the instance
func (instance *Win32_PerfFormattedData_Counters_BluetoothRadio) GetPropertyLEScanWindowCodedPhy() (value uint32, err error) {
	retValue, err := instance.GetProperty("LEScanWindowCodedPhy")
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

// SetLEScanWindowUncoded1MPhy sets the value of LEScanWindowUncoded1MPhy for the instance
func (instance *Win32_PerfFormattedData_Counters_BluetoothRadio) SetPropertyLEScanWindowUncoded1MPhy(value uint32) (err error) {
	return instance.SetProperty("LEScanWindowUncoded1MPhy", (value))
}

// GetLEScanWindowUncoded1MPhy gets the value of LEScanWindowUncoded1MPhy for the instance
func (instance *Win32_PerfFormattedData_Counters_BluetoothRadio) GetPropertyLEScanWindowUncoded1MPhy() (value uint32, err error) {
	retValue, err := instance.GetProperty("LEScanWindowUncoded1MPhy")
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

// SetPageScanDutyCyclePercent sets the value of PageScanDutyCyclePercent for the instance
func (instance *Win32_PerfFormattedData_Counters_BluetoothRadio) SetPropertyPageScanDutyCyclePercent(value uint32) (err error) {
	return instance.SetProperty("PageScanDutyCyclePercent", (value))
}

// GetPageScanDutyCyclePercent gets the value of PageScanDutyCyclePercent for the instance
func (instance *Win32_PerfFormattedData_Counters_BluetoothRadio) GetPropertyPageScanDutyCyclePercent() (value uint32, err error) {
	retValue, err := instance.GetProperty("PageScanDutyCyclePercent")
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

// SetPageScanInterval sets the value of PageScanInterval for the instance
func (instance *Win32_PerfFormattedData_Counters_BluetoothRadio) SetPropertyPageScanInterval(value uint32) (err error) {
	return instance.SetProperty("PageScanInterval", (value))
}

// GetPageScanInterval gets the value of PageScanInterval for the instance
func (instance *Win32_PerfFormattedData_Counters_BluetoothRadio) GetPropertyPageScanInterval() (value uint32, err error) {
	retValue, err := instance.GetProperty("PageScanInterval")
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

// SetPageScanWindow sets the value of PageScanWindow for the instance
func (instance *Win32_PerfFormattedData_Counters_BluetoothRadio) SetPropertyPageScanWindow(value uint32) (err error) {
	return instance.SetProperty("PageScanWindow", (value))
}

// GetPageScanWindow gets the value of PageScanWindow for the instance
func (instance *Win32_PerfFormattedData_Counters_BluetoothRadio) GetPropertyPageScanWindow() (value uint32, err error) {
	retValue, err := instance.GetProperty("PageScanWindow")
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

// SetReportedLEMonitoredDevices sets the value of ReportedLEMonitoredDevices for the instance
func (instance *Win32_PerfFormattedData_Counters_BluetoothRadio) SetPropertyReportedLEMonitoredDevices(value uint32) (err error) {
	return instance.SetProperty("ReportedLEMonitoredDevices", (value))
}

// GetReportedLEMonitoredDevices gets the value of ReportedLEMonitoredDevices for the instance
func (instance *Win32_PerfFormattedData_Counters_BluetoothRadio) GetPropertyReportedLEMonitoredDevices() (value uint32, err error) {
	retValue, err := instance.GetProperty("ReportedLEMonitoredDevices")
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

// SetSCObytesreadPersec sets the value of SCObytesreadPersec for the instance
func (instance *Win32_PerfFormattedData_Counters_BluetoothRadio) SetPropertySCObytesreadPersec(value uint32) (err error) {
	return instance.SetProperty("SCObytesreadPersec", (value))
}

// GetSCObytesreadPersec gets the value of SCObytesreadPersec for the instance
func (instance *Win32_PerfFormattedData_Counters_BluetoothRadio) GetPropertySCObytesreadPersec() (value uint32, err error) {
	retValue, err := instance.GetProperty("SCObytesreadPersec")
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

// SetSCObyteswrittenPersec sets the value of SCObyteswrittenPersec for the instance
func (instance *Win32_PerfFormattedData_Counters_BluetoothRadio) SetPropertySCObyteswrittenPersec(value uint32) (err error) {
	return instance.SetProperty("SCObyteswrittenPersec", (value))
}

// GetSCObyteswrittenPersec gets the value of SCObyteswrittenPersec for the instance
func (instance *Win32_PerfFormattedData_Counters_BluetoothRadio) GetPropertySCObyteswrittenPersec() (value uint32, err error) {
	retValue, err := instance.GetProperty("SCObyteswrittenPersec")
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

// SetSCOConnections sets the value of SCOConnections for the instance
func (instance *Win32_PerfFormattedData_Counters_BluetoothRadio) SetPropertySCOConnections(value uint32) (err error) {
	return instance.SetProperty("SCOConnections", (value))
}

// GetSCOConnections gets the value of SCOConnections for the instance
func (instance *Win32_PerfFormattedData_Counters_BluetoothRadio) GetPropertySCOConnections() (value uint32, err error) {
	retValue, err := instance.GetProperty("SCOConnections")
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

// SetSidebandSCOConnections sets the value of SidebandSCOConnections for the instance
func (instance *Win32_PerfFormattedData_Counters_BluetoothRadio) SetPropertySidebandSCOConnections(value uint32) (err error) {
	return instance.SetProperty("SidebandSCOConnections", (value))
}

// GetSidebandSCOConnections gets the value of SidebandSCOConnections for the instance
func (instance *Win32_PerfFormattedData_Counters_BluetoothRadio) GetPropertySidebandSCOConnections() (value uint32, err error) {
	retValue, err := instance.GetProperty("SidebandSCOConnections")
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
