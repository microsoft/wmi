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

// Win32_PerfRawData_TCPIPCounters_TCPIPExtendedPerformanceDiagnostics struct
type Win32_PerfRawData_TCPIPCounters_TCPIPExtendedPerformanceDiagnostics struct {
	*Win32_PerfRawData

	//
	InboundTCPconnectionattempts uint64

	//
	IPv4NBLsindicatedwithprevalidation uint64

	//
	IPv4NBLsPersecindicatedwithprevalidation uint64

	//
	IPv6NBLsindicatedwithprevalidation uint64

	//
	IPv6NBLsPersecindicatedwithprevalidation uint64

	//
	OutboundTCPconnectionattempts uint64

	//
	RSTsegmentssent uint64

	//
	SYNTCBscreated uint64

	//
	SYNTCBsdestroyed uint64

	//
	TCBscreated uint64

	//
	TCBsdestroyed uint64

	//
	TCPbadsegmentPersecreceived uint64

	//
	TCPbadsegmentsreceived uint64

	//
	TCPconnectionattemptsfailed uint64

	//
	TCPdeliverybufferingfailures uint64

	//
	TCPIPpendableRXNBLsaftercoalescing uint64

	//
	TCPIPpendableRXNBLsreceivedfromminiport uint64

	//
	TCPIPpendableRXNBLsreturnedtominiport uint64

	//
	TcpipRscDisabledMask uint32

	//
	TcpipUsoDisabledMask uint32

	//
	TCPLSObytePersecsent uint64

	//
	TCPLSOsegmentPersecsent uint64

	//
	TCPsegmentPersecreceived uint64

	//
	TCPsegmentPersecretransmitted uint64

	//
	TCPsegmentPersecsent uint64

	//
	TCPtuplesinserted uint64

	//
	TCPtuplesremoved uint64

	//
	TWTCBcollisions uint64

	//
	TWTCBscreated uint64

	//
	TWTCBsdestroyed uint64

	//
	TWTCBsstateshutdown uint64

	//
	UDPUSObytePersecsent uint64

	//
	UDPUSOsegmentPersecsent uint64

	//
	UroDisabledMask uint32
}

func NewWin32_PerfRawData_TCPIPCounters_TCPIPExtendedPerformanceDiagnosticsEx1(instance *cim.WmiInstance) (newInstance *Win32_PerfRawData_TCPIPCounters_TCPIPExtendedPerformanceDiagnostics, err error) {
	tmp, err := NewWin32_PerfRawDataEx1(instance)

	if err != nil {
		return
	}
	newInstance = &Win32_PerfRawData_TCPIPCounters_TCPIPExtendedPerformanceDiagnostics{
		Win32_PerfRawData: tmp,
	}
	return
}

func NewWin32_PerfRawData_TCPIPCounters_TCPIPExtendedPerformanceDiagnosticsEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *Win32_PerfRawData_TCPIPCounters_TCPIPExtendedPerformanceDiagnostics, err error) {
	tmp, err := NewWin32_PerfRawDataEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &Win32_PerfRawData_TCPIPCounters_TCPIPExtendedPerformanceDiagnostics{
		Win32_PerfRawData: tmp,
	}
	return
}

// SetInboundTCPconnectionattempts sets the value of InboundTCPconnectionattempts for the instance
func (instance *Win32_PerfRawData_TCPIPCounters_TCPIPExtendedPerformanceDiagnostics) SetPropertyInboundTCPconnectionattempts(value uint64) (err error) {
	return instance.SetProperty("InboundTCPconnectionattempts", (value))
}

// GetInboundTCPconnectionattempts gets the value of InboundTCPconnectionattempts for the instance
func (instance *Win32_PerfRawData_TCPIPCounters_TCPIPExtendedPerformanceDiagnostics) GetPropertyInboundTCPconnectionattempts() (value uint64, err error) {
	retValue, err := instance.GetProperty("InboundTCPconnectionattempts")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(uint64)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " uint64 is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = uint64(valuetmp)

	return
}

// SetIPv4NBLsindicatedwithprevalidation sets the value of IPv4NBLsindicatedwithprevalidation for the instance
func (instance *Win32_PerfRawData_TCPIPCounters_TCPIPExtendedPerformanceDiagnostics) SetPropertyIPv4NBLsindicatedwithprevalidation(value uint64) (err error) {
	return instance.SetProperty("IPv4NBLsindicatedwithprevalidation", (value))
}

// GetIPv4NBLsindicatedwithprevalidation gets the value of IPv4NBLsindicatedwithprevalidation for the instance
func (instance *Win32_PerfRawData_TCPIPCounters_TCPIPExtendedPerformanceDiagnostics) GetPropertyIPv4NBLsindicatedwithprevalidation() (value uint64, err error) {
	retValue, err := instance.GetProperty("IPv4NBLsindicatedwithprevalidation")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(uint64)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " uint64 is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = uint64(valuetmp)

	return
}

// SetIPv4NBLsPersecindicatedwithprevalidation sets the value of IPv4NBLsPersecindicatedwithprevalidation for the instance
func (instance *Win32_PerfRawData_TCPIPCounters_TCPIPExtendedPerformanceDiagnostics) SetPropertyIPv4NBLsPersecindicatedwithprevalidation(value uint64) (err error) {
	return instance.SetProperty("IPv4NBLsPersecindicatedwithprevalidation", (value))
}

// GetIPv4NBLsPersecindicatedwithprevalidation gets the value of IPv4NBLsPersecindicatedwithprevalidation for the instance
func (instance *Win32_PerfRawData_TCPIPCounters_TCPIPExtendedPerformanceDiagnostics) GetPropertyIPv4NBLsPersecindicatedwithprevalidation() (value uint64, err error) {
	retValue, err := instance.GetProperty("IPv4NBLsPersecindicatedwithprevalidation")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(uint64)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " uint64 is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = uint64(valuetmp)

	return
}

// SetIPv6NBLsindicatedwithprevalidation sets the value of IPv6NBLsindicatedwithprevalidation for the instance
func (instance *Win32_PerfRawData_TCPIPCounters_TCPIPExtendedPerformanceDiagnostics) SetPropertyIPv6NBLsindicatedwithprevalidation(value uint64) (err error) {
	return instance.SetProperty("IPv6NBLsindicatedwithprevalidation", (value))
}

// GetIPv6NBLsindicatedwithprevalidation gets the value of IPv6NBLsindicatedwithprevalidation for the instance
func (instance *Win32_PerfRawData_TCPIPCounters_TCPIPExtendedPerformanceDiagnostics) GetPropertyIPv6NBLsindicatedwithprevalidation() (value uint64, err error) {
	retValue, err := instance.GetProperty("IPv6NBLsindicatedwithprevalidation")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(uint64)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " uint64 is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = uint64(valuetmp)

	return
}

// SetIPv6NBLsPersecindicatedwithprevalidation sets the value of IPv6NBLsPersecindicatedwithprevalidation for the instance
func (instance *Win32_PerfRawData_TCPIPCounters_TCPIPExtendedPerformanceDiagnostics) SetPropertyIPv6NBLsPersecindicatedwithprevalidation(value uint64) (err error) {
	return instance.SetProperty("IPv6NBLsPersecindicatedwithprevalidation", (value))
}

// GetIPv6NBLsPersecindicatedwithprevalidation gets the value of IPv6NBLsPersecindicatedwithprevalidation for the instance
func (instance *Win32_PerfRawData_TCPIPCounters_TCPIPExtendedPerformanceDiagnostics) GetPropertyIPv6NBLsPersecindicatedwithprevalidation() (value uint64, err error) {
	retValue, err := instance.GetProperty("IPv6NBLsPersecindicatedwithprevalidation")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(uint64)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " uint64 is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = uint64(valuetmp)

	return
}

// SetOutboundTCPconnectionattempts sets the value of OutboundTCPconnectionattempts for the instance
func (instance *Win32_PerfRawData_TCPIPCounters_TCPIPExtendedPerformanceDiagnostics) SetPropertyOutboundTCPconnectionattempts(value uint64) (err error) {
	return instance.SetProperty("OutboundTCPconnectionattempts", (value))
}

// GetOutboundTCPconnectionattempts gets the value of OutboundTCPconnectionattempts for the instance
func (instance *Win32_PerfRawData_TCPIPCounters_TCPIPExtendedPerformanceDiagnostics) GetPropertyOutboundTCPconnectionattempts() (value uint64, err error) {
	retValue, err := instance.GetProperty("OutboundTCPconnectionattempts")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(uint64)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " uint64 is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = uint64(valuetmp)

	return
}

// SetRSTsegmentssent sets the value of RSTsegmentssent for the instance
func (instance *Win32_PerfRawData_TCPIPCounters_TCPIPExtendedPerformanceDiagnostics) SetPropertyRSTsegmentssent(value uint64) (err error) {
	return instance.SetProperty("RSTsegmentssent", (value))
}

// GetRSTsegmentssent gets the value of RSTsegmentssent for the instance
func (instance *Win32_PerfRawData_TCPIPCounters_TCPIPExtendedPerformanceDiagnostics) GetPropertyRSTsegmentssent() (value uint64, err error) {
	retValue, err := instance.GetProperty("RSTsegmentssent")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(uint64)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " uint64 is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = uint64(valuetmp)

	return
}

// SetSYNTCBscreated sets the value of SYNTCBscreated for the instance
func (instance *Win32_PerfRawData_TCPIPCounters_TCPIPExtendedPerformanceDiagnostics) SetPropertySYNTCBscreated(value uint64) (err error) {
	return instance.SetProperty("SYNTCBscreated", (value))
}

// GetSYNTCBscreated gets the value of SYNTCBscreated for the instance
func (instance *Win32_PerfRawData_TCPIPCounters_TCPIPExtendedPerformanceDiagnostics) GetPropertySYNTCBscreated() (value uint64, err error) {
	retValue, err := instance.GetProperty("SYNTCBscreated")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(uint64)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " uint64 is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = uint64(valuetmp)

	return
}

// SetSYNTCBsdestroyed sets the value of SYNTCBsdestroyed for the instance
func (instance *Win32_PerfRawData_TCPIPCounters_TCPIPExtendedPerformanceDiagnostics) SetPropertySYNTCBsdestroyed(value uint64) (err error) {
	return instance.SetProperty("SYNTCBsdestroyed", (value))
}

// GetSYNTCBsdestroyed gets the value of SYNTCBsdestroyed for the instance
func (instance *Win32_PerfRawData_TCPIPCounters_TCPIPExtendedPerformanceDiagnostics) GetPropertySYNTCBsdestroyed() (value uint64, err error) {
	retValue, err := instance.GetProperty("SYNTCBsdestroyed")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(uint64)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " uint64 is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = uint64(valuetmp)

	return
}

// SetTCBscreated sets the value of TCBscreated for the instance
func (instance *Win32_PerfRawData_TCPIPCounters_TCPIPExtendedPerformanceDiagnostics) SetPropertyTCBscreated(value uint64) (err error) {
	return instance.SetProperty("TCBscreated", (value))
}

// GetTCBscreated gets the value of TCBscreated for the instance
func (instance *Win32_PerfRawData_TCPIPCounters_TCPIPExtendedPerformanceDiagnostics) GetPropertyTCBscreated() (value uint64, err error) {
	retValue, err := instance.GetProperty("TCBscreated")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(uint64)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " uint64 is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = uint64(valuetmp)

	return
}

// SetTCBsdestroyed sets the value of TCBsdestroyed for the instance
func (instance *Win32_PerfRawData_TCPIPCounters_TCPIPExtendedPerformanceDiagnostics) SetPropertyTCBsdestroyed(value uint64) (err error) {
	return instance.SetProperty("TCBsdestroyed", (value))
}

// GetTCBsdestroyed gets the value of TCBsdestroyed for the instance
func (instance *Win32_PerfRawData_TCPIPCounters_TCPIPExtendedPerformanceDiagnostics) GetPropertyTCBsdestroyed() (value uint64, err error) {
	retValue, err := instance.GetProperty("TCBsdestroyed")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(uint64)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " uint64 is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = uint64(valuetmp)

	return
}

// SetTCPbadsegmentPersecreceived sets the value of TCPbadsegmentPersecreceived for the instance
func (instance *Win32_PerfRawData_TCPIPCounters_TCPIPExtendedPerformanceDiagnostics) SetPropertyTCPbadsegmentPersecreceived(value uint64) (err error) {
	return instance.SetProperty("TCPbadsegmentPersecreceived", (value))
}

// GetTCPbadsegmentPersecreceived gets the value of TCPbadsegmentPersecreceived for the instance
func (instance *Win32_PerfRawData_TCPIPCounters_TCPIPExtendedPerformanceDiagnostics) GetPropertyTCPbadsegmentPersecreceived() (value uint64, err error) {
	retValue, err := instance.GetProperty("TCPbadsegmentPersecreceived")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(uint64)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " uint64 is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = uint64(valuetmp)

	return
}

// SetTCPbadsegmentsreceived sets the value of TCPbadsegmentsreceived for the instance
func (instance *Win32_PerfRawData_TCPIPCounters_TCPIPExtendedPerformanceDiagnostics) SetPropertyTCPbadsegmentsreceived(value uint64) (err error) {
	return instance.SetProperty("TCPbadsegmentsreceived", (value))
}

// GetTCPbadsegmentsreceived gets the value of TCPbadsegmentsreceived for the instance
func (instance *Win32_PerfRawData_TCPIPCounters_TCPIPExtendedPerformanceDiagnostics) GetPropertyTCPbadsegmentsreceived() (value uint64, err error) {
	retValue, err := instance.GetProperty("TCPbadsegmentsreceived")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(uint64)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " uint64 is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = uint64(valuetmp)

	return
}

// SetTCPconnectionattemptsfailed sets the value of TCPconnectionattemptsfailed for the instance
func (instance *Win32_PerfRawData_TCPIPCounters_TCPIPExtendedPerformanceDiagnostics) SetPropertyTCPconnectionattemptsfailed(value uint64) (err error) {
	return instance.SetProperty("TCPconnectionattemptsfailed", (value))
}

// GetTCPconnectionattemptsfailed gets the value of TCPconnectionattemptsfailed for the instance
func (instance *Win32_PerfRawData_TCPIPCounters_TCPIPExtendedPerformanceDiagnostics) GetPropertyTCPconnectionattemptsfailed() (value uint64, err error) {
	retValue, err := instance.GetProperty("TCPconnectionattemptsfailed")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(uint64)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " uint64 is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = uint64(valuetmp)

	return
}

// SetTCPdeliverybufferingfailures sets the value of TCPdeliverybufferingfailures for the instance
func (instance *Win32_PerfRawData_TCPIPCounters_TCPIPExtendedPerformanceDiagnostics) SetPropertyTCPdeliverybufferingfailures(value uint64) (err error) {
	return instance.SetProperty("TCPdeliverybufferingfailures", (value))
}

// GetTCPdeliverybufferingfailures gets the value of TCPdeliverybufferingfailures for the instance
func (instance *Win32_PerfRawData_TCPIPCounters_TCPIPExtendedPerformanceDiagnostics) GetPropertyTCPdeliverybufferingfailures() (value uint64, err error) {
	retValue, err := instance.GetProperty("TCPdeliverybufferingfailures")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(uint64)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " uint64 is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = uint64(valuetmp)

	return
}

// SetTCPIPpendableRXNBLsaftercoalescing sets the value of TCPIPpendableRXNBLsaftercoalescing for the instance
func (instance *Win32_PerfRawData_TCPIPCounters_TCPIPExtendedPerformanceDiagnostics) SetPropertyTCPIPpendableRXNBLsaftercoalescing(value uint64) (err error) {
	return instance.SetProperty("TCPIPpendableRXNBLsaftercoalescing", (value))
}

// GetTCPIPpendableRXNBLsaftercoalescing gets the value of TCPIPpendableRXNBLsaftercoalescing for the instance
func (instance *Win32_PerfRawData_TCPIPCounters_TCPIPExtendedPerformanceDiagnostics) GetPropertyTCPIPpendableRXNBLsaftercoalescing() (value uint64, err error) {
	retValue, err := instance.GetProperty("TCPIPpendableRXNBLsaftercoalescing")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(uint64)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " uint64 is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = uint64(valuetmp)

	return
}

// SetTCPIPpendableRXNBLsreceivedfromminiport sets the value of TCPIPpendableRXNBLsreceivedfromminiport for the instance
func (instance *Win32_PerfRawData_TCPIPCounters_TCPIPExtendedPerformanceDiagnostics) SetPropertyTCPIPpendableRXNBLsreceivedfromminiport(value uint64) (err error) {
	return instance.SetProperty("TCPIPpendableRXNBLsreceivedfromminiport", (value))
}

// GetTCPIPpendableRXNBLsreceivedfromminiport gets the value of TCPIPpendableRXNBLsreceivedfromminiport for the instance
func (instance *Win32_PerfRawData_TCPIPCounters_TCPIPExtendedPerformanceDiagnostics) GetPropertyTCPIPpendableRXNBLsreceivedfromminiport() (value uint64, err error) {
	retValue, err := instance.GetProperty("TCPIPpendableRXNBLsreceivedfromminiport")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(uint64)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " uint64 is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = uint64(valuetmp)

	return
}

// SetTCPIPpendableRXNBLsreturnedtominiport sets the value of TCPIPpendableRXNBLsreturnedtominiport for the instance
func (instance *Win32_PerfRawData_TCPIPCounters_TCPIPExtendedPerformanceDiagnostics) SetPropertyTCPIPpendableRXNBLsreturnedtominiport(value uint64) (err error) {
	return instance.SetProperty("TCPIPpendableRXNBLsreturnedtominiport", (value))
}

// GetTCPIPpendableRXNBLsreturnedtominiport gets the value of TCPIPpendableRXNBLsreturnedtominiport for the instance
func (instance *Win32_PerfRawData_TCPIPCounters_TCPIPExtendedPerformanceDiagnostics) GetPropertyTCPIPpendableRXNBLsreturnedtominiport() (value uint64, err error) {
	retValue, err := instance.GetProperty("TCPIPpendableRXNBLsreturnedtominiport")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(uint64)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " uint64 is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = uint64(valuetmp)

	return
}

// SetTcpipRscDisabledMask sets the value of TcpipRscDisabledMask for the instance
func (instance *Win32_PerfRawData_TCPIPCounters_TCPIPExtendedPerformanceDiagnostics) SetPropertyTcpipRscDisabledMask(value uint32) (err error) {
	return instance.SetProperty("TcpipRscDisabledMask", (value))
}

// GetTcpipRscDisabledMask gets the value of TcpipRscDisabledMask for the instance
func (instance *Win32_PerfRawData_TCPIPCounters_TCPIPExtendedPerformanceDiagnostics) GetPropertyTcpipRscDisabledMask() (value uint32, err error) {
	retValue, err := instance.GetProperty("TcpipRscDisabledMask")
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

// SetTcpipUsoDisabledMask sets the value of TcpipUsoDisabledMask for the instance
func (instance *Win32_PerfRawData_TCPIPCounters_TCPIPExtendedPerformanceDiagnostics) SetPropertyTcpipUsoDisabledMask(value uint32) (err error) {
	return instance.SetProperty("TcpipUsoDisabledMask", (value))
}

// GetTcpipUsoDisabledMask gets the value of TcpipUsoDisabledMask for the instance
func (instance *Win32_PerfRawData_TCPIPCounters_TCPIPExtendedPerformanceDiagnostics) GetPropertyTcpipUsoDisabledMask() (value uint32, err error) {
	retValue, err := instance.GetProperty("TcpipUsoDisabledMask")
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

// SetTCPLSObytePersecsent sets the value of TCPLSObytePersecsent for the instance
func (instance *Win32_PerfRawData_TCPIPCounters_TCPIPExtendedPerformanceDiagnostics) SetPropertyTCPLSObytePersecsent(value uint64) (err error) {
	return instance.SetProperty("TCPLSObytePersecsent", (value))
}

// GetTCPLSObytePersecsent gets the value of TCPLSObytePersecsent for the instance
func (instance *Win32_PerfRawData_TCPIPCounters_TCPIPExtendedPerformanceDiagnostics) GetPropertyTCPLSObytePersecsent() (value uint64, err error) {
	retValue, err := instance.GetProperty("TCPLSObytePersecsent")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(uint64)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " uint64 is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = uint64(valuetmp)

	return
}

// SetTCPLSOsegmentPersecsent sets the value of TCPLSOsegmentPersecsent for the instance
func (instance *Win32_PerfRawData_TCPIPCounters_TCPIPExtendedPerformanceDiagnostics) SetPropertyTCPLSOsegmentPersecsent(value uint64) (err error) {
	return instance.SetProperty("TCPLSOsegmentPersecsent", (value))
}

// GetTCPLSOsegmentPersecsent gets the value of TCPLSOsegmentPersecsent for the instance
func (instance *Win32_PerfRawData_TCPIPCounters_TCPIPExtendedPerformanceDiagnostics) GetPropertyTCPLSOsegmentPersecsent() (value uint64, err error) {
	retValue, err := instance.GetProperty("TCPLSOsegmentPersecsent")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(uint64)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " uint64 is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = uint64(valuetmp)

	return
}

// SetTCPsegmentPersecreceived sets the value of TCPsegmentPersecreceived for the instance
func (instance *Win32_PerfRawData_TCPIPCounters_TCPIPExtendedPerformanceDiagnostics) SetPropertyTCPsegmentPersecreceived(value uint64) (err error) {
	return instance.SetProperty("TCPsegmentPersecreceived", (value))
}

// GetTCPsegmentPersecreceived gets the value of TCPsegmentPersecreceived for the instance
func (instance *Win32_PerfRawData_TCPIPCounters_TCPIPExtendedPerformanceDiagnostics) GetPropertyTCPsegmentPersecreceived() (value uint64, err error) {
	retValue, err := instance.GetProperty("TCPsegmentPersecreceived")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(uint64)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " uint64 is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = uint64(valuetmp)

	return
}

// SetTCPsegmentPersecretransmitted sets the value of TCPsegmentPersecretransmitted for the instance
func (instance *Win32_PerfRawData_TCPIPCounters_TCPIPExtendedPerformanceDiagnostics) SetPropertyTCPsegmentPersecretransmitted(value uint64) (err error) {
	return instance.SetProperty("TCPsegmentPersecretransmitted", (value))
}

// GetTCPsegmentPersecretransmitted gets the value of TCPsegmentPersecretransmitted for the instance
func (instance *Win32_PerfRawData_TCPIPCounters_TCPIPExtendedPerformanceDiagnostics) GetPropertyTCPsegmentPersecretransmitted() (value uint64, err error) {
	retValue, err := instance.GetProperty("TCPsegmentPersecretransmitted")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(uint64)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " uint64 is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = uint64(valuetmp)

	return
}

// SetTCPsegmentPersecsent sets the value of TCPsegmentPersecsent for the instance
func (instance *Win32_PerfRawData_TCPIPCounters_TCPIPExtendedPerformanceDiagnostics) SetPropertyTCPsegmentPersecsent(value uint64) (err error) {
	return instance.SetProperty("TCPsegmentPersecsent", (value))
}

// GetTCPsegmentPersecsent gets the value of TCPsegmentPersecsent for the instance
func (instance *Win32_PerfRawData_TCPIPCounters_TCPIPExtendedPerformanceDiagnostics) GetPropertyTCPsegmentPersecsent() (value uint64, err error) {
	retValue, err := instance.GetProperty("TCPsegmentPersecsent")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(uint64)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " uint64 is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = uint64(valuetmp)

	return
}

// SetTCPtuplesinserted sets the value of TCPtuplesinserted for the instance
func (instance *Win32_PerfRawData_TCPIPCounters_TCPIPExtendedPerformanceDiagnostics) SetPropertyTCPtuplesinserted(value uint64) (err error) {
	return instance.SetProperty("TCPtuplesinserted", (value))
}

// GetTCPtuplesinserted gets the value of TCPtuplesinserted for the instance
func (instance *Win32_PerfRawData_TCPIPCounters_TCPIPExtendedPerformanceDiagnostics) GetPropertyTCPtuplesinserted() (value uint64, err error) {
	retValue, err := instance.GetProperty("TCPtuplesinserted")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(uint64)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " uint64 is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = uint64(valuetmp)

	return
}

// SetTCPtuplesremoved sets the value of TCPtuplesremoved for the instance
func (instance *Win32_PerfRawData_TCPIPCounters_TCPIPExtendedPerformanceDiagnostics) SetPropertyTCPtuplesremoved(value uint64) (err error) {
	return instance.SetProperty("TCPtuplesremoved", (value))
}

// GetTCPtuplesremoved gets the value of TCPtuplesremoved for the instance
func (instance *Win32_PerfRawData_TCPIPCounters_TCPIPExtendedPerformanceDiagnostics) GetPropertyTCPtuplesremoved() (value uint64, err error) {
	retValue, err := instance.GetProperty("TCPtuplesremoved")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(uint64)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " uint64 is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = uint64(valuetmp)

	return
}

// SetTWTCBcollisions sets the value of TWTCBcollisions for the instance
func (instance *Win32_PerfRawData_TCPIPCounters_TCPIPExtendedPerformanceDiagnostics) SetPropertyTWTCBcollisions(value uint64) (err error) {
	return instance.SetProperty("TWTCBcollisions", (value))
}

// GetTWTCBcollisions gets the value of TWTCBcollisions for the instance
func (instance *Win32_PerfRawData_TCPIPCounters_TCPIPExtendedPerformanceDiagnostics) GetPropertyTWTCBcollisions() (value uint64, err error) {
	retValue, err := instance.GetProperty("TWTCBcollisions")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(uint64)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " uint64 is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = uint64(valuetmp)

	return
}

// SetTWTCBscreated sets the value of TWTCBscreated for the instance
func (instance *Win32_PerfRawData_TCPIPCounters_TCPIPExtendedPerformanceDiagnostics) SetPropertyTWTCBscreated(value uint64) (err error) {
	return instance.SetProperty("TWTCBscreated", (value))
}

// GetTWTCBscreated gets the value of TWTCBscreated for the instance
func (instance *Win32_PerfRawData_TCPIPCounters_TCPIPExtendedPerformanceDiagnostics) GetPropertyTWTCBscreated() (value uint64, err error) {
	retValue, err := instance.GetProperty("TWTCBscreated")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(uint64)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " uint64 is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = uint64(valuetmp)

	return
}

// SetTWTCBsdestroyed sets the value of TWTCBsdestroyed for the instance
func (instance *Win32_PerfRawData_TCPIPCounters_TCPIPExtendedPerformanceDiagnostics) SetPropertyTWTCBsdestroyed(value uint64) (err error) {
	return instance.SetProperty("TWTCBsdestroyed", (value))
}

// GetTWTCBsdestroyed gets the value of TWTCBsdestroyed for the instance
func (instance *Win32_PerfRawData_TCPIPCounters_TCPIPExtendedPerformanceDiagnostics) GetPropertyTWTCBsdestroyed() (value uint64, err error) {
	retValue, err := instance.GetProperty("TWTCBsdestroyed")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(uint64)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " uint64 is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = uint64(valuetmp)

	return
}

// SetTWTCBsstateshutdown sets the value of TWTCBsstateshutdown for the instance
func (instance *Win32_PerfRawData_TCPIPCounters_TCPIPExtendedPerformanceDiagnostics) SetPropertyTWTCBsstateshutdown(value uint64) (err error) {
	return instance.SetProperty("TWTCBsstateshutdown", (value))
}

// GetTWTCBsstateshutdown gets the value of TWTCBsstateshutdown for the instance
func (instance *Win32_PerfRawData_TCPIPCounters_TCPIPExtendedPerformanceDiagnostics) GetPropertyTWTCBsstateshutdown() (value uint64, err error) {
	retValue, err := instance.GetProperty("TWTCBsstateshutdown")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(uint64)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " uint64 is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = uint64(valuetmp)

	return
}

// SetUDPUSObytePersecsent sets the value of UDPUSObytePersecsent for the instance
func (instance *Win32_PerfRawData_TCPIPCounters_TCPIPExtendedPerformanceDiagnostics) SetPropertyUDPUSObytePersecsent(value uint64) (err error) {
	return instance.SetProperty("UDPUSObytePersecsent", (value))
}

// GetUDPUSObytePersecsent gets the value of UDPUSObytePersecsent for the instance
func (instance *Win32_PerfRawData_TCPIPCounters_TCPIPExtendedPerformanceDiagnostics) GetPropertyUDPUSObytePersecsent() (value uint64, err error) {
	retValue, err := instance.GetProperty("UDPUSObytePersecsent")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(uint64)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " uint64 is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = uint64(valuetmp)

	return
}

// SetUDPUSOsegmentPersecsent sets the value of UDPUSOsegmentPersecsent for the instance
func (instance *Win32_PerfRawData_TCPIPCounters_TCPIPExtendedPerformanceDiagnostics) SetPropertyUDPUSOsegmentPersecsent(value uint64) (err error) {
	return instance.SetProperty("UDPUSOsegmentPersecsent", (value))
}

// GetUDPUSOsegmentPersecsent gets the value of UDPUSOsegmentPersecsent for the instance
func (instance *Win32_PerfRawData_TCPIPCounters_TCPIPExtendedPerformanceDiagnostics) GetPropertyUDPUSOsegmentPersecsent() (value uint64, err error) {
	retValue, err := instance.GetProperty("UDPUSOsegmentPersecsent")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(uint64)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " uint64 is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = uint64(valuetmp)

	return
}

// SetUroDisabledMask sets the value of UroDisabledMask for the instance
func (instance *Win32_PerfRawData_TCPIPCounters_TCPIPExtendedPerformanceDiagnostics) SetPropertyUroDisabledMask(value uint32) (err error) {
	return instance.SetProperty("UroDisabledMask", (value))
}

// GetUroDisabledMask gets the value of UroDisabledMask for the instance
func (instance *Win32_PerfRawData_TCPIPCounters_TCPIPExtendedPerformanceDiagnostics) GetPropertyUroDisabledMask() (value uint32, err error) {
	retValue, err := instance.GetProperty("UroDisabledMask")
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
