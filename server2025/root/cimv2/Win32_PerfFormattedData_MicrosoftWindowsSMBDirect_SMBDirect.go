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

// Win32_PerfFormattedData_MicrosoftWindowsSMBDirect_SMBDirect struct
type Win32_PerfFormattedData_MicrosoftWindowsSMBDirect_SMBDirect struct {
	*Win32_PerfFormattedData

	//
	AllocationFailuresConnector uint32

	//
	AllocationFailuresCQ uint32

	//
	AllocationFailuresFRMR uint32

	//
	AllocationFailuresMR uint32

	//
	AllocationFailuresPD uint32

	//
	AllocationFailuresQP uint32

	//
	AllocationStallsConnector uint32

	//
	AllocationStallsCQ uint32

	//
	AllocationStallsFRMR uint32

	//
	AllocationStallsMR uint32

	//
	AllocationStallsPD uint32

	//
	AllocationStallsQP uint32

	//
	BytesRDMAReadPersec uint64

	//
	BytesRDMAWrittenPersec uint64

	//
	BytesReceivedPersec uint64

	//
	BytesRegistered uint32

	//
	BytesSentPersec uint64

	//
	FastregistrationsPersec uint64

	//
	FRMRPages uint32

	//
	FRMRs uint32

	//
	InboundConnectionFailures uint32

	//
	InboundConnectionRefusals uint32

	//
	InboundConnections uint32

	//
	InboundConnectionsMax uint32

	//
	InboundConnectionsPending uint32

	//
	InvalidationsPersec uint64

	//
	OutboundConnectionFailures uint32

	//
	OutboundConnections uint32

	//
	OutboundConnectionsMax uint32

	//
	ProtocolErrors uint32

	//
	RCQNotificationsPersec uint64

	//
	RCQNotificationsQueuedPersec uint64

	//
	RDMAReadsPersec uint64

	//
	RDMAWritesPersec uint64

	//
	ReceivesControlPersec uint64

	//
	ReceivesPersec uint64

	//
	RemoteInvalidationsPersec uint64

	//
	SCQNotificationsPersec uint64

	//
	SCQNotificationsQueuedPersec uint64

	//
	SendsControlPersec uint64

	//
	SendsPersec uint64

	//
	SQRequestsPersec uint64

	//
	SQRequestsPersecunsignaled uint64

	//
	StallsFRMRPersec uint64

	//
	StallsRDMAReadPersec uint64

	//
	StallsSendCreditPersec uint64

	//
	StallsSQPersec uint64

	//
	TotalConnections uint32

	//
	TotalConnectionsMax uint32
}

func NewWin32_PerfFormattedData_MicrosoftWindowsSMBDirect_SMBDirectEx1(instance *cim.WmiInstance) (newInstance *Win32_PerfFormattedData_MicrosoftWindowsSMBDirect_SMBDirect, err error) {
	tmp, err := NewWin32_PerfFormattedDataEx1(instance)

	if err != nil {
		return
	}
	newInstance = &Win32_PerfFormattedData_MicrosoftWindowsSMBDirect_SMBDirect{
		Win32_PerfFormattedData: tmp,
	}
	return
}

func NewWin32_PerfFormattedData_MicrosoftWindowsSMBDirect_SMBDirectEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *Win32_PerfFormattedData_MicrosoftWindowsSMBDirect_SMBDirect, err error) {
	tmp, err := NewWin32_PerfFormattedDataEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &Win32_PerfFormattedData_MicrosoftWindowsSMBDirect_SMBDirect{
		Win32_PerfFormattedData: tmp,
	}
	return
}

// SetAllocationFailuresConnector sets the value of AllocationFailuresConnector for the instance
func (instance *Win32_PerfFormattedData_MicrosoftWindowsSMBDirect_SMBDirect) SetPropertyAllocationFailuresConnector(value uint32) (err error) {
	return instance.SetProperty("AllocationFailuresConnector", (value))
}

// GetAllocationFailuresConnector gets the value of AllocationFailuresConnector for the instance
func (instance *Win32_PerfFormattedData_MicrosoftWindowsSMBDirect_SMBDirect) GetPropertyAllocationFailuresConnector() (value uint32, err error) {
	retValue, err := instance.GetProperty("AllocationFailuresConnector")
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

// SetAllocationFailuresCQ sets the value of AllocationFailuresCQ for the instance
func (instance *Win32_PerfFormattedData_MicrosoftWindowsSMBDirect_SMBDirect) SetPropertyAllocationFailuresCQ(value uint32) (err error) {
	return instance.SetProperty("AllocationFailuresCQ", (value))
}

// GetAllocationFailuresCQ gets the value of AllocationFailuresCQ for the instance
func (instance *Win32_PerfFormattedData_MicrosoftWindowsSMBDirect_SMBDirect) GetPropertyAllocationFailuresCQ() (value uint32, err error) {
	retValue, err := instance.GetProperty("AllocationFailuresCQ")
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

// SetAllocationFailuresFRMR sets the value of AllocationFailuresFRMR for the instance
func (instance *Win32_PerfFormattedData_MicrosoftWindowsSMBDirect_SMBDirect) SetPropertyAllocationFailuresFRMR(value uint32) (err error) {
	return instance.SetProperty("AllocationFailuresFRMR", (value))
}

// GetAllocationFailuresFRMR gets the value of AllocationFailuresFRMR for the instance
func (instance *Win32_PerfFormattedData_MicrosoftWindowsSMBDirect_SMBDirect) GetPropertyAllocationFailuresFRMR() (value uint32, err error) {
	retValue, err := instance.GetProperty("AllocationFailuresFRMR")
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

// SetAllocationFailuresMR sets the value of AllocationFailuresMR for the instance
func (instance *Win32_PerfFormattedData_MicrosoftWindowsSMBDirect_SMBDirect) SetPropertyAllocationFailuresMR(value uint32) (err error) {
	return instance.SetProperty("AllocationFailuresMR", (value))
}

// GetAllocationFailuresMR gets the value of AllocationFailuresMR for the instance
func (instance *Win32_PerfFormattedData_MicrosoftWindowsSMBDirect_SMBDirect) GetPropertyAllocationFailuresMR() (value uint32, err error) {
	retValue, err := instance.GetProperty("AllocationFailuresMR")
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

// SetAllocationFailuresPD sets the value of AllocationFailuresPD for the instance
func (instance *Win32_PerfFormattedData_MicrosoftWindowsSMBDirect_SMBDirect) SetPropertyAllocationFailuresPD(value uint32) (err error) {
	return instance.SetProperty("AllocationFailuresPD", (value))
}

// GetAllocationFailuresPD gets the value of AllocationFailuresPD for the instance
func (instance *Win32_PerfFormattedData_MicrosoftWindowsSMBDirect_SMBDirect) GetPropertyAllocationFailuresPD() (value uint32, err error) {
	retValue, err := instance.GetProperty("AllocationFailuresPD")
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

// SetAllocationFailuresQP sets the value of AllocationFailuresQP for the instance
func (instance *Win32_PerfFormattedData_MicrosoftWindowsSMBDirect_SMBDirect) SetPropertyAllocationFailuresQP(value uint32) (err error) {
	return instance.SetProperty("AllocationFailuresQP", (value))
}

// GetAllocationFailuresQP gets the value of AllocationFailuresQP for the instance
func (instance *Win32_PerfFormattedData_MicrosoftWindowsSMBDirect_SMBDirect) GetPropertyAllocationFailuresQP() (value uint32, err error) {
	retValue, err := instance.GetProperty("AllocationFailuresQP")
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

// SetAllocationStallsConnector sets the value of AllocationStallsConnector for the instance
func (instance *Win32_PerfFormattedData_MicrosoftWindowsSMBDirect_SMBDirect) SetPropertyAllocationStallsConnector(value uint32) (err error) {
	return instance.SetProperty("AllocationStallsConnector", (value))
}

// GetAllocationStallsConnector gets the value of AllocationStallsConnector for the instance
func (instance *Win32_PerfFormattedData_MicrosoftWindowsSMBDirect_SMBDirect) GetPropertyAllocationStallsConnector() (value uint32, err error) {
	retValue, err := instance.GetProperty("AllocationStallsConnector")
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

// SetAllocationStallsCQ sets the value of AllocationStallsCQ for the instance
func (instance *Win32_PerfFormattedData_MicrosoftWindowsSMBDirect_SMBDirect) SetPropertyAllocationStallsCQ(value uint32) (err error) {
	return instance.SetProperty("AllocationStallsCQ", (value))
}

// GetAllocationStallsCQ gets the value of AllocationStallsCQ for the instance
func (instance *Win32_PerfFormattedData_MicrosoftWindowsSMBDirect_SMBDirect) GetPropertyAllocationStallsCQ() (value uint32, err error) {
	retValue, err := instance.GetProperty("AllocationStallsCQ")
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

// SetAllocationStallsFRMR sets the value of AllocationStallsFRMR for the instance
func (instance *Win32_PerfFormattedData_MicrosoftWindowsSMBDirect_SMBDirect) SetPropertyAllocationStallsFRMR(value uint32) (err error) {
	return instance.SetProperty("AllocationStallsFRMR", (value))
}

// GetAllocationStallsFRMR gets the value of AllocationStallsFRMR for the instance
func (instance *Win32_PerfFormattedData_MicrosoftWindowsSMBDirect_SMBDirect) GetPropertyAllocationStallsFRMR() (value uint32, err error) {
	retValue, err := instance.GetProperty("AllocationStallsFRMR")
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

// SetAllocationStallsMR sets the value of AllocationStallsMR for the instance
func (instance *Win32_PerfFormattedData_MicrosoftWindowsSMBDirect_SMBDirect) SetPropertyAllocationStallsMR(value uint32) (err error) {
	return instance.SetProperty("AllocationStallsMR", (value))
}

// GetAllocationStallsMR gets the value of AllocationStallsMR for the instance
func (instance *Win32_PerfFormattedData_MicrosoftWindowsSMBDirect_SMBDirect) GetPropertyAllocationStallsMR() (value uint32, err error) {
	retValue, err := instance.GetProperty("AllocationStallsMR")
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

// SetAllocationStallsPD sets the value of AllocationStallsPD for the instance
func (instance *Win32_PerfFormattedData_MicrosoftWindowsSMBDirect_SMBDirect) SetPropertyAllocationStallsPD(value uint32) (err error) {
	return instance.SetProperty("AllocationStallsPD", (value))
}

// GetAllocationStallsPD gets the value of AllocationStallsPD for the instance
func (instance *Win32_PerfFormattedData_MicrosoftWindowsSMBDirect_SMBDirect) GetPropertyAllocationStallsPD() (value uint32, err error) {
	retValue, err := instance.GetProperty("AllocationStallsPD")
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

// SetAllocationStallsQP sets the value of AllocationStallsQP for the instance
func (instance *Win32_PerfFormattedData_MicrosoftWindowsSMBDirect_SMBDirect) SetPropertyAllocationStallsQP(value uint32) (err error) {
	return instance.SetProperty("AllocationStallsQP", (value))
}

// GetAllocationStallsQP gets the value of AllocationStallsQP for the instance
func (instance *Win32_PerfFormattedData_MicrosoftWindowsSMBDirect_SMBDirect) GetPropertyAllocationStallsQP() (value uint32, err error) {
	retValue, err := instance.GetProperty("AllocationStallsQP")
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

// SetBytesRDMAReadPersec sets the value of BytesRDMAReadPersec for the instance
func (instance *Win32_PerfFormattedData_MicrosoftWindowsSMBDirect_SMBDirect) SetPropertyBytesRDMAReadPersec(value uint64) (err error) {
	return instance.SetProperty("BytesRDMAReadPersec", (value))
}

// GetBytesRDMAReadPersec gets the value of BytesRDMAReadPersec for the instance
func (instance *Win32_PerfFormattedData_MicrosoftWindowsSMBDirect_SMBDirect) GetPropertyBytesRDMAReadPersec() (value uint64, err error) {
	retValue, err := instance.GetProperty("BytesRDMAReadPersec")
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

// SetBytesRDMAWrittenPersec sets the value of BytesRDMAWrittenPersec for the instance
func (instance *Win32_PerfFormattedData_MicrosoftWindowsSMBDirect_SMBDirect) SetPropertyBytesRDMAWrittenPersec(value uint64) (err error) {
	return instance.SetProperty("BytesRDMAWrittenPersec", (value))
}

// GetBytesRDMAWrittenPersec gets the value of BytesRDMAWrittenPersec for the instance
func (instance *Win32_PerfFormattedData_MicrosoftWindowsSMBDirect_SMBDirect) GetPropertyBytesRDMAWrittenPersec() (value uint64, err error) {
	retValue, err := instance.GetProperty("BytesRDMAWrittenPersec")
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

// SetBytesReceivedPersec sets the value of BytesReceivedPersec for the instance
func (instance *Win32_PerfFormattedData_MicrosoftWindowsSMBDirect_SMBDirect) SetPropertyBytesReceivedPersec(value uint64) (err error) {
	return instance.SetProperty("BytesReceivedPersec", (value))
}

// GetBytesReceivedPersec gets the value of BytesReceivedPersec for the instance
func (instance *Win32_PerfFormattedData_MicrosoftWindowsSMBDirect_SMBDirect) GetPropertyBytesReceivedPersec() (value uint64, err error) {
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

// SetBytesRegistered sets the value of BytesRegistered for the instance
func (instance *Win32_PerfFormattedData_MicrosoftWindowsSMBDirect_SMBDirect) SetPropertyBytesRegistered(value uint32) (err error) {
	return instance.SetProperty("BytesRegistered", (value))
}

// GetBytesRegistered gets the value of BytesRegistered for the instance
func (instance *Win32_PerfFormattedData_MicrosoftWindowsSMBDirect_SMBDirect) GetPropertyBytesRegistered() (value uint32, err error) {
	retValue, err := instance.GetProperty("BytesRegistered")
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

// SetBytesSentPersec sets the value of BytesSentPersec for the instance
func (instance *Win32_PerfFormattedData_MicrosoftWindowsSMBDirect_SMBDirect) SetPropertyBytesSentPersec(value uint64) (err error) {
	return instance.SetProperty("BytesSentPersec", (value))
}

// GetBytesSentPersec gets the value of BytesSentPersec for the instance
func (instance *Win32_PerfFormattedData_MicrosoftWindowsSMBDirect_SMBDirect) GetPropertyBytesSentPersec() (value uint64, err error) {
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

// SetFastregistrationsPersec sets the value of FastregistrationsPersec for the instance
func (instance *Win32_PerfFormattedData_MicrosoftWindowsSMBDirect_SMBDirect) SetPropertyFastregistrationsPersec(value uint64) (err error) {
	return instance.SetProperty("FastregistrationsPersec", (value))
}

// GetFastregistrationsPersec gets the value of FastregistrationsPersec for the instance
func (instance *Win32_PerfFormattedData_MicrosoftWindowsSMBDirect_SMBDirect) GetPropertyFastregistrationsPersec() (value uint64, err error) {
	retValue, err := instance.GetProperty("FastregistrationsPersec")
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

// SetFRMRPages sets the value of FRMRPages for the instance
func (instance *Win32_PerfFormattedData_MicrosoftWindowsSMBDirect_SMBDirect) SetPropertyFRMRPages(value uint32) (err error) {
	return instance.SetProperty("FRMRPages", (value))
}

// GetFRMRPages gets the value of FRMRPages for the instance
func (instance *Win32_PerfFormattedData_MicrosoftWindowsSMBDirect_SMBDirect) GetPropertyFRMRPages() (value uint32, err error) {
	retValue, err := instance.GetProperty("FRMRPages")
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

// SetFRMRs sets the value of FRMRs for the instance
func (instance *Win32_PerfFormattedData_MicrosoftWindowsSMBDirect_SMBDirect) SetPropertyFRMRs(value uint32) (err error) {
	return instance.SetProperty("FRMRs", (value))
}

// GetFRMRs gets the value of FRMRs for the instance
func (instance *Win32_PerfFormattedData_MicrosoftWindowsSMBDirect_SMBDirect) GetPropertyFRMRs() (value uint32, err error) {
	retValue, err := instance.GetProperty("FRMRs")
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

// SetInboundConnectionFailures sets the value of InboundConnectionFailures for the instance
func (instance *Win32_PerfFormattedData_MicrosoftWindowsSMBDirect_SMBDirect) SetPropertyInboundConnectionFailures(value uint32) (err error) {
	return instance.SetProperty("InboundConnectionFailures", (value))
}

// GetInboundConnectionFailures gets the value of InboundConnectionFailures for the instance
func (instance *Win32_PerfFormattedData_MicrosoftWindowsSMBDirect_SMBDirect) GetPropertyInboundConnectionFailures() (value uint32, err error) {
	retValue, err := instance.GetProperty("InboundConnectionFailures")
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

// SetInboundConnectionRefusals sets the value of InboundConnectionRefusals for the instance
func (instance *Win32_PerfFormattedData_MicrosoftWindowsSMBDirect_SMBDirect) SetPropertyInboundConnectionRefusals(value uint32) (err error) {
	return instance.SetProperty("InboundConnectionRefusals", (value))
}

// GetInboundConnectionRefusals gets the value of InboundConnectionRefusals for the instance
func (instance *Win32_PerfFormattedData_MicrosoftWindowsSMBDirect_SMBDirect) GetPropertyInboundConnectionRefusals() (value uint32, err error) {
	retValue, err := instance.GetProperty("InboundConnectionRefusals")
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

// SetInboundConnections sets the value of InboundConnections for the instance
func (instance *Win32_PerfFormattedData_MicrosoftWindowsSMBDirect_SMBDirect) SetPropertyInboundConnections(value uint32) (err error) {
	return instance.SetProperty("InboundConnections", (value))
}

// GetInboundConnections gets the value of InboundConnections for the instance
func (instance *Win32_PerfFormattedData_MicrosoftWindowsSMBDirect_SMBDirect) GetPropertyInboundConnections() (value uint32, err error) {
	retValue, err := instance.GetProperty("InboundConnections")
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

// SetInboundConnectionsMax sets the value of InboundConnectionsMax for the instance
func (instance *Win32_PerfFormattedData_MicrosoftWindowsSMBDirect_SMBDirect) SetPropertyInboundConnectionsMax(value uint32) (err error) {
	return instance.SetProperty("InboundConnectionsMax", (value))
}

// GetInboundConnectionsMax gets the value of InboundConnectionsMax for the instance
func (instance *Win32_PerfFormattedData_MicrosoftWindowsSMBDirect_SMBDirect) GetPropertyInboundConnectionsMax() (value uint32, err error) {
	retValue, err := instance.GetProperty("InboundConnectionsMax")
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

// SetInboundConnectionsPending sets the value of InboundConnectionsPending for the instance
func (instance *Win32_PerfFormattedData_MicrosoftWindowsSMBDirect_SMBDirect) SetPropertyInboundConnectionsPending(value uint32) (err error) {
	return instance.SetProperty("InboundConnectionsPending", (value))
}

// GetInboundConnectionsPending gets the value of InboundConnectionsPending for the instance
func (instance *Win32_PerfFormattedData_MicrosoftWindowsSMBDirect_SMBDirect) GetPropertyInboundConnectionsPending() (value uint32, err error) {
	retValue, err := instance.GetProperty("InboundConnectionsPending")
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

// SetInvalidationsPersec sets the value of InvalidationsPersec for the instance
func (instance *Win32_PerfFormattedData_MicrosoftWindowsSMBDirect_SMBDirect) SetPropertyInvalidationsPersec(value uint64) (err error) {
	return instance.SetProperty("InvalidationsPersec", (value))
}

// GetInvalidationsPersec gets the value of InvalidationsPersec for the instance
func (instance *Win32_PerfFormattedData_MicrosoftWindowsSMBDirect_SMBDirect) GetPropertyInvalidationsPersec() (value uint64, err error) {
	retValue, err := instance.GetProperty("InvalidationsPersec")
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

// SetOutboundConnectionFailures sets the value of OutboundConnectionFailures for the instance
func (instance *Win32_PerfFormattedData_MicrosoftWindowsSMBDirect_SMBDirect) SetPropertyOutboundConnectionFailures(value uint32) (err error) {
	return instance.SetProperty("OutboundConnectionFailures", (value))
}

// GetOutboundConnectionFailures gets the value of OutboundConnectionFailures for the instance
func (instance *Win32_PerfFormattedData_MicrosoftWindowsSMBDirect_SMBDirect) GetPropertyOutboundConnectionFailures() (value uint32, err error) {
	retValue, err := instance.GetProperty("OutboundConnectionFailures")
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

// SetOutboundConnections sets the value of OutboundConnections for the instance
func (instance *Win32_PerfFormattedData_MicrosoftWindowsSMBDirect_SMBDirect) SetPropertyOutboundConnections(value uint32) (err error) {
	return instance.SetProperty("OutboundConnections", (value))
}

// GetOutboundConnections gets the value of OutboundConnections for the instance
func (instance *Win32_PerfFormattedData_MicrosoftWindowsSMBDirect_SMBDirect) GetPropertyOutboundConnections() (value uint32, err error) {
	retValue, err := instance.GetProperty("OutboundConnections")
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

// SetOutboundConnectionsMax sets the value of OutboundConnectionsMax for the instance
func (instance *Win32_PerfFormattedData_MicrosoftWindowsSMBDirect_SMBDirect) SetPropertyOutboundConnectionsMax(value uint32) (err error) {
	return instance.SetProperty("OutboundConnectionsMax", (value))
}

// GetOutboundConnectionsMax gets the value of OutboundConnectionsMax for the instance
func (instance *Win32_PerfFormattedData_MicrosoftWindowsSMBDirect_SMBDirect) GetPropertyOutboundConnectionsMax() (value uint32, err error) {
	retValue, err := instance.GetProperty("OutboundConnectionsMax")
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

// SetProtocolErrors sets the value of ProtocolErrors for the instance
func (instance *Win32_PerfFormattedData_MicrosoftWindowsSMBDirect_SMBDirect) SetPropertyProtocolErrors(value uint32) (err error) {
	return instance.SetProperty("ProtocolErrors", (value))
}

// GetProtocolErrors gets the value of ProtocolErrors for the instance
func (instance *Win32_PerfFormattedData_MicrosoftWindowsSMBDirect_SMBDirect) GetPropertyProtocolErrors() (value uint32, err error) {
	retValue, err := instance.GetProperty("ProtocolErrors")
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

// SetRCQNotificationsPersec sets the value of RCQNotificationsPersec for the instance
func (instance *Win32_PerfFormattedData_MicrosoftWindowsSMBDirect_SMBDirect) SetPropertyRCQNotificationsPersec(value uint64) (err error) {
	return instance.SetProperty("RCQNotificationsPersec", (value))
}

// GetRCQNotificationsPersec gets the value of RCQNotificationsPersec for the instance
func (instance *Win32_PerfFormattedData_MicrosoftWindowsSMBDirect_SMBDirect) GetPropertyRCQNotificationsPersec() (value uint64, err error) {
	retValue, err := instance.GetProperty("RCQNotificationsPersec")
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

// SetRCQNotificationsQueuedPersec sets the value of RCQNotificationsQueuedPersec for the instance
func (instance *Win32_PerfFormattedData_MicrosoftWindowsSMBDirect_SMBDirect) SetPropertyRCQNotificationsQueuedPersec(value uint64) (err error) {
	return instance.SetProperty("RCQNotificationsQueuedPersec", (value))
}

// GetRCQNotificationsQueuedPersec gets the value of RCQNotificationsQueuedPersec for the instance
func (instance *Win32_PerfFormattedData_MicrosoftWindowsSMBDirect_SMBDirect) GetPropertyRCQNotificationsQueuedPersec() (value uint64, err error) {
	retValue, err := instance.GetProperty("RCQNotificationsQueuedPersec")
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

// SetRDMAReadsPersec sets the value of RDMAReadsPersec for the instance
func (instance *Win32_PerfFormattedData_MicrosoftWindowsSMBDirect_SMBDirect) SetPropertyRDMAReadsPersec(value uint64) (err error) {
	return instance.SetProperty("RDMAReadsPersec", (value))
}

// GetRDMAReadsPersec gets the value of RDMAReadsPersec for the instance
func (instance *Win32_PerfFormattedData_MicrosoftWindowsSMBDirect_SMBDirect) GetPropertyRDMAReadsPersec() (value uint64, err error) {
	retValue, err := instance.GetProperty("RDMAReadsPersec")
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

// SetRDMAWritesPersec sets the value of RDMAWritesPersec for the instance
func (instance *Win32_PerfFormattedData_MicrosoftWindowsSMBDirect_SMBDirect) SetPropertyRDMAWritesPersec(value uint64) (err error) {
	return instance.SetProperty("RDMAWritesPersec", (value))
}

// GetRDMAWritesPersec gets the value of RDMAWritesPersec for the instance
func (instance *Win32_PerfFormattedData_MicrosoftWindowsSMBDirect_SMBDirect) GetPropertyRDMAWritesPersec() (value uint64, err error) {
	retValue, err := instance.GetProperty("RDMAWritesPersec")
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

// SetReceivesControlPersec sets the value of ReceivesControlPersec for the instance
func (instance *Win32_PerfFormattedData_MicrosoftWindowsSMBDirect_SMBDirect) SetPropertyReceivesControlPersec(value uint64) (err error) {
	return instance.SetProperty("ReceivesControlPersec", (value))
}

// GetReceivesControlPersec gets the value of ReceivesControlPersec for the instance
func (instance *Win32_PerfFormattedData_MicrosoftWindowsSMBDirect_SMBDirect) GetPropertyReceivesControlPersec() (value uint64, err error) {
	retValue, err := instance.GetProperty("ReceivesControlPersec")
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
func (instance *Win32_PerfFormattedData_MicrosoftWindowsSMBDirect_SMBDirect) SetPropertyReceivesPersec(value uint64) (err error) {
	return instance.SetProperty("ReceivesPersec", (value))
}

// GetReceivesPersec gets the value of ReceivesPersec for the instance
func (instance *Win32_PerfFormattedData_MicrosoftWindowsSMBDirect_SMBDirect) GetPropertyReceivesPersec() (value uint64, err error) {
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

// SetRemoteInvalidationsPersec sets the value of RemoteInvalidationsPersec for the instance
func (instance *Win32_PerfFormattedData_MicrosoftWindowsSMBDirect_SMBDirect) SetPropertyRemoteInvalidationsPersec(value uint64) (err error) {
	return instance.SetProperty("RemoteInvalidationsPersec", (value))
}

// GetRemoteInvalidationsPersec gets the value of RemoteInvalidationsPersec for the instance
func (instance *Win32_PerfFormattedData_MicrosoftWindowsSMBDirect_SMBDirect) GetPropertyRemoteInvalidationsPersec() (value uint64, err error) {
	retValue, err := instance.GetProperty("RemoteInvalidationsPersec")
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

// SetSCQNotificationsPersec sets the value of SCQNotificationsPersec for the instance
func (instance *Win32_PerfFormattedData_MicrosoftWindowsSMBDirect_SMBDirect) SetPropertySCQNotificationsPersec(value uint64) (err error) {
	return instance.SetProperty("SCQNotificationsPersec", (value))
}

// GetSCQNotificationsPersec gets the value of SCQNotificationsPersec for the instance
func (instance *Win32_PerfFormattedData_MicrosoftWindowsSMBDirect_SMBDirect) GetPropertySCQNotificationsPersec() (value uint64, err error) {
	retValue, err := instance.GetProperty("SCQNotificationsPersec")
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

// SetSCQNotificationsQueuedPersec sets the value of SCQNotificationsQueuedPersec for the instance
func (instance *Win32_PerfFormattedData_MicrosoftWindowsSMBDirect_SMBDirect) SetPropertySCQNotificationsQueuedPersec(value uint64) (err error) {
	return instance.SetProperty("SCQNotificationsQueuedPersec", (value))
}

// GetSCQNotificationsQueuedPersec gets the value of SCQNotificationsQueuedPersec for the instance
func (instance *Win32_PerfFormattedData_MicrosoftWindowsSMBDirect_SMBDirect) GetPropertySCQNotificationsQueuedPersec() (value uint64, err error) {
	retValue, err := instance.GetProperty("SCQNotificationsQueuedPersec")
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

// SetSendsControlPersec sets the value of SendsControlPersec for the instance
func (instance *Win32_PerfFormattedData_MicrosoftWindowsSMBDirect_SMBDirect) SetPropertySendsControlPersec(value uint64) (err error) {
	return instance.SetProperty("SendsControlPersec", (value))
}

// GetSendsControlPersec gets the value of SendsControlPersec for the instance
func (instance *Win32_PerfFormattedData_MicrosoftWindowsSMBDirect_SMBDirect) GetPropertySendsControlPersec() (value uint64, err error) {
	retValue, err := instance.GetProperty("SendsControlPersec")
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
func (instance *Win32_PerfFormattedData_MicrosoftWindowsSMBDirect_SMBDirect) SetPropertySendsPersec(value uint64) (err error) {
	return instance.SetProperty("SendsPersec", (value))
}

// GetSendsPersec gets the value of SendsPersec for the instance
func (instance *Win32_PerfFormattedData_MicrosoftWindowsSMBDirect_SMBDirect) GetPropertySendsPersec() (value uint64, err error) {
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

// SetSQRequestsPersec sets the value of SQRequestsPersec for the instance
func (instance *Win32_PerfFormattedData_MicrosoftWindowsSMBDirect_SMBDirect) SetPropertySQRequestsPersec(value uint64) (err error) {
	return instance.SetProperty("SQRequestsPersec", (value))
}

// GetSQRequestsPersec gets the value of SQRequestsPersec for the instance
func (instance *Win32_PerfFormattedData_MicrosoftWindowsSMBDirect_SMBDirect) GetPropertySQRequestsPersec() (value uint64, err error) {
	retValue, err := instance.GetProperty("SQRequestsPersec")
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

// SetSQRequestsPersecunsignaled sets the value of SQRequestsPersecunsignaled for the instance
func (instance *Win32_PerfFormattedData_MicrosoftWindowsSMBDirect_SMBDirect) SetPropertySQRequestsPersecunsignaled(value uint64) (err error) {
	return instance.SetProperty("SQRequestsPersecunsignaled", (value))
}

// GetSQRequestsPersecunsignaled gets the value of SQRequestsPersecunsignaled for the instance
func (instance *Win32_PerfFormattedData_MicrosoftWindowsSMBDirect_SMBDirect) GetPropertySQRequestsPersecunsignaled() (value uint64, err error) {
	retValue, err := instance.GetProperty("SQRequestsPersecunsignaled")
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

// SetStallsFRMRPersec sets the value of StallsFRMRPersec for the instance
func (instance *Win32_PerfFormattedData_MicrosoftWindowsSMBDirect_SMBDirect) SetPropertyStallsFRMRPersec(value uint64) (err error) {
	return instance.SetProperty("StallsFRMRPersec", (value))
}

// GetStallsFRMRPersec gets the value of StallsFRMRPersec for the instance
func (instance *Win32_PerfFormattedData_MicrosoftWindowsSMBDirect_SMBDirect) GetPropertyStallsFRMRPersec() (value uint64, err error) {
	retValue, err := instance.GetProperty("StallsFRMRPersec")
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

// SetStallsRDMAReadPersec sets the value of StallsRDMAReadPersec for the instance
func (instance *Win32_PerfFormattedData_MicrosoftWindowsSMBDirect_SMBDirect) SetPropertyStallsRDMAReadPersec(value uint64) (err error) {
	return instance.SetProperty("StallsRDMAReadPersec", (value))
}

// GetStallsRDMAReadPersec gets the value of StallsRDMAReadPersec for the instance
func (instance *Win32_PerfFormattedData_MicrosoftWindowsSMBDirect_SMBDirect) GetPropertyStallsRDMAReadPersec() (value uint64, err error) {
	retValue, err := instance.GetProperty("StallsRDMAReadPersec")
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

// SetStallsSendCreditPersec sets the value of StallsSendCreditPersec for the instance
func (instance *Win32_PerfFormattedData_MicrosoftWindowsSMBDirect_SMBDirect) SetPropertyStallsSendCreditPersec(value uint64) (err error) {
	return instance.SetProperty("StallsSendCreditPersec", (value))
}

// GetStallsSendCreditPersec gets the value of StallsSendCreditPersec for the instance
func (instance *Win32_PerfFormattedData_MicrosoftWindowsSMBDirect_SMBDirect) GetPropertyStallsSendCreditPersec() (value uint64, err error) {
	retValue, err := instance.GetProperty("StallsSendCreditPersec")
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

// SetStallsSQPersec sets the value of StallsSQPersec for the instance
func (instance *Win32_PerfFormattedData_MicrosoftWindowsSMBDirect_SMBDirect) SetPropertyStallsSQPersec(value uint64) (err error) {
	return instance.SetProperty("StallsSQPersec", (value))
}

// GetStallsSQPersec gets the value of StallsSQPersec for the instance
func (instance *Win32_PerfFormattedData_MicrosoftWindowsSMBDirect_SMBDirect) GetPropertyStallsSQPersec() (value uint64, err error) {
	retValue, err := instance.GetProperty("StallsSQPersec")
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

// SetTotalConnections sets the value of TotalConnections for the instance
func (instance *Win32_PerfFormattedData_MicrosoftWindowsSMBDirect_SMBDirect) SetPropertyTotalConnections(value uint32) (err error) {
	return instance.SetProperty("TotalConnections", (value))
}

// GetTotalConnections gets the value of TotalConnections for the instance
func (instance *Win32_PerfFormattedData_MicrosoftWindowsSMBDirect_SMBDirect) GetPropertyTotalConnections() (value uint32, err error) {
	retValue, err := instance.GetProperty("TotalConnections")
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

// SetTotalConnectionsMax sets the value of TotalConnectionsMax for the instance
func (instance *Win32_PerfFormattedData_MicrosoftWindowsSMBDirect_SMBDirect) SetPropertyTotalConnectionsMax(value uint32) (err error) {
	return instance.SetProperty("TotalConnectionsMax", (value))
}

// GetTotalConnectionsMax gets the value of TotalConnectionsMax for the instance
func (instance *Win32_PerfFormattedData_MicrosoftWindowsSMBDirect_SMBDirect) GetPropertyTotalConnectionsMax() (value uint32, err error) {
	retValue, err := instance.GetProperty("TotalConnectionsMax")
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
