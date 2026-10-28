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

// Win32_PerfRawData_TCPIPCounters_TCPIPTransportLayerPacketDropCounters struct
type Win32_PerfRawData_TCPIPCounters_TCPIPTransportLayerPacketDropCounters struct {
	*Win32_PerfRawData

	//
	InetDiscardAcceptInspection uint64

	//
	InetDiscardAcceptRedirection uint64

	//
	InetDiscardAckInvalid uint64

	//
	InetDiscardChecksumInvalid uint64

	//
	InetDiscardClosedWindow uint64

	//
	InetDiscardConnectedPath uint64

	//
	InetDiscardDestinationMulticast uint64

	//
	InetDiscardDuplicateSegment uint64

	//
	InetDiscardEndpointNotFound uint64

	//
	InetDiscardExpectedSyn uint64

	//
	InetDiscardFinReceived uint64

	//
	InetDiscardFinWait2 uint64

	//
	InetDiscardHeaderInvalid uint64

	//
	InetDiscardLandAttack uint64

	//
	InetDiscardListenerInvalidFlags uint64

	//
	InetDiscardMissedReset uint64

	//
	InetDiscardOutsideWindow uint64

	//
	InetDiscardPauseAccept uint64

	//
	InetDiscardPawsFailed uint64

	//
	InetDiscardReassemblyConflict uint64

	//
	InetDiscardReceiveInspection uint64

	//
	InetDiscardRst uint64

	//
	InetDiscardSessionState uint64

	//
	InetDiscardSimultaneousConnect uint64

	//
	InetDiscardSourceUnspecified uint64

	//
	InetDiscardSynAckWithFastopenCookieRequest uint64

	//
	InetDiscardSynAttack uint64

	//
	InetDiscardSynRcvdSyn uint64

	//
	InetDiscardTcbNotInTcbTable uint64

	//
	InetDiscardTcbRemoved uint64

	//
	InetDiscardTimeWaitTcb uint64

	//
	InetDiscardTimeWaitTcbReceivedRstOutsideWindow uint64

	//
	InetDiscardTimeWaitTcbSynAndOtherFlags uint64

	//
	InetDiscardUrgentDeliveryAllocationFailure uint64
}

func NewWin32_PerfRawData_TCPIPCounters_TCPIPTransportLayerPacketDropCountersEx1(instance *cim.WmiInstance) (newInstance *Win32_PerfRawData_TCPIPCounters_TCPIPTransportLayerPacketDropCounters, err error) {
	tmp, err := NewWin32_PerfRawDataEx1(instance)

	if err != nil {
		return
	}
	newInstance = &Win32_PerfRawData_TCPIPCounters_TCPIPTransportLayerPacketDropCounters{
		Win32_PerfRawData: tmp,
	}
	return
}

func NewWin32_PerfRawData_TCPIPCounters_TCPIPTransportLayerPacketDropCountersEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *Win32_PerfRawData_TCPIPCounters_TCPIPTransportLayerPacketDropCounters, err error) {
	tmp, err := NewWin32_PerfRawDataEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &Win32_PerfRawData_TCPIPCounters_TCPIPTransportLayerPacketDropCounters{
		Win32_PerfRawData: tmp,
	}
	return
}

// SetInetDiscardAcceptInspection sets the value of InetDiscardAcceptInspection for the instance
func (instance *Win32_PerfRawData_TCPIPCounters_TCPIPTransportLayerPacketDropCounters) SetPropertyInetDiscardAcceptInspection(value uint64) (err error) {
	return instance.SetProperty("InetDiscardAcceptInspection", (value))
}

// GetInetDiscardAcceptInspection gets the value of InetDiscardAcceptInspection for the instance
func (instance *Win32_PerfRawData_TCPIPCounters_TCPIPTransportLayerPacketDropCounters) GetPropertyInetDiscardAcceptInspection() (value uint64, err error) {
	retValue, err := instance.GetProperty("InetDiscardAcceptInspection")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(uint64)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " uint64 is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = uint64(valuetmp)

	return
}

// SetInetDiscardAcceptRedirection sets the value of InetDiscardAcceptRedirection for the instance
func (instance *Win32_PerfRawData_TCPIPCounters_TCPIPTransportLayerPacketDropCounters) SetPropertyInetDiscardAcceptRedirection(value uint64) (err error) {
	return instance.SetProperty("InetDiscardAcceptRedirection", (value))
}

// GetInetDiscardAcceptRedirection gets the value of InetDiscardAcceptRedirection for the instance
func (instance *Win32_PerfRawData_TCPIPCounters_TCPIPTransportLayerPacketDropCounters) GetPropertyInetDiscardAcceptRedirection() (value uint64, err error) {
	retValue, err := instance.GetProperty("InetDiscardAcceptRedirection")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(uint64)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " uint64 is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = uint64(valuetmp)

	return
}

// SetInetDiscardAckInvalid sets the value of InetDiscardAckInvalid for the instance
func (instance *Win32_PerfRawData_TCPIPCounters_TCPIPTransportLayerPacketDropCounters) SetPropertyInetDiscardAckInvalid(value uint64) (err error) {
	return instance.SetProperty("InetDiscardAckInvalid", (value))
}

// GetInetDiscardAckInvalid gets the value of InetDiscardAckInvalid for the instance
func (instance *Win32_PerfRawData_TCPIPCounters_TCPIPTransportLayerPacketDropCounters) GetPropertyInetDiscardAckInvalid() (value uint64, err error) {
	retValue, err := instance.GetProperty("InetDiscardAckInvalid")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(uint64)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " uint64 is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = uint64(valuetmp)

	return
}

// SetInetDiscardChecksumInvalid sets the value of InetDiscardChecksumInvalid for the instance
func (instance *Win32_PerfRawData_TCPIPCounters_TCPIPTransportLayerPacketDropCounters) SetPropertyInetDiscardChecksumInvalid(value uint64) (err error) {
	return instance.SetProperty("InetDiscardChecksumInvalid", (value))
}

// GetInetDiscardChecksumInvalid gets the value of InetDiscardChecksumInvalid for the instance
func (instance *Win32_PerfRawData_TCPIPCounters_TCPIPTransportLayerPacketDropCounters) GetPropertyInetDiscardChecksumInvalid() (value uint64, err error) {
	retValue, err := instance.GetProperty("InetDiscardChecksumInvalid")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(uint64)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " uint64 is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = uint64(valuetmp)

	return
}

// SetInetDiscardClosedWindow sets the value of InetDiscardClosedWindow for the instance
func (instance *Win32_PerfRawData_TCPIPCounters_TCPIPTransportLayerPacketDropCounters) SetPropertyInetDiscardClosedWindow(value uint64) (err error) {
	return instance.SetProperty("InetDiscardClosedWindow", (value))
}

// GetInetDiscardClosedWindow gets the value of InetDiscardClosedWindow for the instance
func (instance *Win32_PerfRawData_TCPIPCounters_TCPIPTransportLayerPacketDropCounters) GetPropertyInetDiscardClosedWindow() (value uint64, err error) {
	retValue, err := instance.GetProperty("InetDiscardClosedWindow")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(uint64)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " uint64 is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = uint64(valuetmp)

	return
}

// SetInetDiscardConnectedPath sets the value of InetDiscardConnectedPath for the instance
func (instance *Win32_PerfRawData_TCPIPCounters_TCPIPTransportLayerPacketDropCounters) SetPropertyInetDiscardConnectedPath(value uint64) (err error) {
	return instance.SetProperty("InetDiscardConnectedPath", (value))
}

// GetInetDiscardConnectedPath gets the value of InetDiscardConnectedPath for the instance
func (instance *Win32_PerfRawData_TCPIPCounters_TCPIPTransportLayerPacketDropCounters) GetPropertyInetDiscardConnectedPath() (value uint64, err error) {
	retValue, err := instance.GetProperty("InetDiscardConnectedPath")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(uint64)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " uint64 is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = uint64(valuetmp)

	return
}

// SetInetDiscardDestinationMulticast sets the value of InetDiscardDestinationMulticast for the instance
func (instance *Win32_PerfRawData_TCPIPCounters_TCPIPTransportLayerPacketDropCounters) SetPropertyInetDiscardDestinationMulticast(value uint64) (err error) {
	return instance.SetProperty("InetDiscardDestinationMulticast", (value))
}

// GetInetDiscardDestinationMulticast gets the value of InetDiscardDestinationMulticast for the instance
func (instance *Win32_PerfRawData_TCPIPCounters_TCPIPTransportLayerPacketDropCounters) GetPropertyInetDiscardDestinationMulticast() (value uint64, err error) {
	retValue, err := instance.GetProperty("InetDiscardDestinationMulticast")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(uint64)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " uint64 is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = uint64(valuetmp)

	return
}

// SetInetDiscardDuplicateSegment sets the value of InetDiscardDuplicateSegment for the instance
func (instance *Win32_PerfRawData_TCPIPCounters_TCPIPTransportLayerPacketDropCounters) SetPropertyInetDiscardDuplicateSegment(value uint64) (err error) {
	return instance.SetProperty("InetDiscardDuplicateSegment", (value))
}

// GetInetDiscardDuplicateSegment gets the value of InetDiscardDuplicateSegment for the instance
func (instance *Win32_PerfRawData_TCPIPCounters_TCPIPTransportLayerPacketDropCounters) GetPropertyInetDiscardDuplicateSegment() (value uint64, err error) {
	retValue, err := instance.GetProperty("InetDiscardDuplicateSegment")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(uint64)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " uint64 is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = uint64(valuetmp)

	return
}

// SetInetDiscardEndpointNotFound sets the value of InetDiscardEndpointNotFound for the instance
func (instance *Win32_PerfRawData_TCPIPCounters_TCPIPTransportLayerPacketDropCounters) SetPropertyInetDiscardEndpointNotFound(value uint64) (err error) {
	return instance.SetProperty("InetDiscardEndpointNotFound", (value))
}

// GetInetDiscardEndpointNotFound gets the value of InetDiscardEndpointNotFound for the instance
func (instance *Win32_PerfRawData_TCPIPCounters_TCPIPTransportLayerPacketDropCounters) GetPropertyInetDiscardEndpointNotFound() (value uint64, err error) {
	retValue, err := instance.GetProperty("InetDiscardEndpointNotFound")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(uint64)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " uint64 is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = uint64(valuetmp)

	return
}

// SetInetDiscardExpectedSyn sets the value of InetDiscardExpectedSyn for the instance
func (instance *Win32_PerfRawData_TCPIPCounters_TCPIPTransportLayerPacketDropCounters) SetPropertyInetDiscardExpectedSyn(value uint64) (err error) {
	return instance.SetProperty("InetDiscardExpectedSyn", (value))
}

// GetInetDiscardExpectedSyn gets the value of InetDiscardExpectedSyn for the instance
func (instance *Win32_PerfRawData_TCPIPCounters_TCPIPTransportLayerPacketDropCounters) GetPropertyInetDiscardExpectedSyn() (value uint64, err error) {
	retValue, err := instance.GetProperty("InetDiscardExpectedSyn")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(uint64)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " uint64 is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = uint64(valuetmp)

	return
}

// SetInetDiscardFinReceived sets the value of InetDiscardFinReceived for the instance
func (instance *Win32_PerfRawData_TCPIPCounters_TCPIPTransportLayerPacketDropCounters) SetPropertyInetDiscardFinReceived(value uint64) (err error) {
	return instance.SetProperty("InetDiscardFinReceived", (value))
}

// GetInetDiscardFinReceived gets the value of InetDiscardFinReceived for the instance
func (instance *Win32_PerfRawData_TCPIPCounters_TCPIPTransportLayerPacketDropCounters) GetPropertyInetDiscardFinReceived() (value uint64, err error) {
	retValue, err := instance.GetProperty("InetDiscardFinReceived")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(uint64)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " uint64 is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = uint64(valuetmp)

	return
}

// SetInetDiscardFinWait2 sets the value of InetDiscardFinWait2 for the instance
func (instance *Win32_PerfRawData_TCPIPCounters_TCPIPTransportLayerPacketDropCounters) SetPropertyInetDiscardFinWait2(value uint64) (err error) {
	return instance.SetProperty("InetDiscardFinWait2", (value))
}

// GetInetDiscardFinWait2 gets the value of InetDiscardFinWait2 for the instance
func (instance *Win32_PerfRawData_TCPIPCounters_TCPIPTransportLayerPacketDropCounters) GetPropertyInetDiscardFinWait2() (value uint64, err error) {
	retValue, err := instance.GetProperty("InetDiscardFinWait2")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(uint64)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " uint64 is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = uint64(valuetmp)

	return
}

// SetInetDiscardHeaderInvalid sets the value of InetDiscardHeaderInvalid for the instance
func (instance *Win32_PerfRawData_TCPIPCounters_TCPIPTransportLayerPacketDropCounters) SetPropertyInetDiscardHeaderInvalid(value uint64) (err error) {
	return instance.SetProperty("InetDiscardHeaderInvalid", (value))
}

// GetInetDiscardHeaderInvalid gets the value of InetDiscardHeaderInvalid for the instance
func (instance *Win32_PerfRawData_TCPIPCounters_TCPIPTransportLayerPacketDropCounters) GetPropertyInetDiscardHeaderInvalid() (value uint64, err error) {
	retValue, err := instance.GetProperty("InetDiscardHeaderInvalid")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(uint64)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " uint64 is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = uint64(valuetmp)

	return
}

// SetInetDiscardLandAttack sets the value of InetDiscardLandAttack for the instance
func (instance *Win32_PerfRawData_TCPIPCounters_TCPIPTransportLayerPacketDropCounters) SetPropertyInetDiscardLandAttack(value uint64) (err error) {
	return instance.SetProperty("InetDiscardLandAttack", (value))
}

// GetInetDiscardLandAttack gets the value of InetDiscardLandAttack for the instance
func (instance *Win32_PerfRawData_TCPIPCounters_TCPIPTransportLayerPacketDropCounters) GetPropertyInetDiscardLandAttack() (value uint64, err error) {
	retValue, err := instance.GetProperty("InetDiscardLandAttack")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(uint64)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " uint64 is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = uint64(valuetmp)

	return
}

// SetInetDiscardListenerInvalidFlags sets the value of InetDiscardListenerInvalidFlags for the instance
func (instance *Win32_PerfRawData_TCPIPCounters_TCPIPTransportLayerPacketDropCounters) SetPropertyInetDiscardListenerInvalidFlags(value uint64) (err error) {
	return instance.SetProperty("InetDiscardListenerInvalidFlags", (value))
}

// GetInetDiscardListenerInvalidFlags gets the value of InetDiscardListenerInvalidFlags for the instance
func (instance *Win32_PerfRawData_TCPIPCounters_TCPIPTransportLayerPacketDropCounters) GetPropertyInetDiscardListenerInvalidFlags() (value uint64, err error) {
	retValue, err := instance.GetProperty("InetDiscardListenerInvalidFlags")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(uint64)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " uint64 is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = uint64(valuetmp)

	return
}

// SetInetDiscardMissedReset sets the value of InetDiscardMissedReset for the instance
func (instance *Win32_PerfRawData_TCPIPCounters_TCPIPTransportLayerPacketDropCounters) SetPropertyInetDiscardMissedReset(value uint64) (err error) {
	return instance.SetProperty("InetDiscardMissedReset", (value))
}

// GetInetDiscardMissedReset gets the value of InetDiscardMissedReset for the instance
func (instance *Win32_PerfRawData_TCPIPCounters_TCPIPTransportLayerPacketDropCounters) GetPropertyInetDiscardMissedReset() (value uint64, err error) {
	retValue, err := instance.GetProperty("InetDiscardMissedReset")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(uint64)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " uint64 is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = uint64(valuetmp)

	return
}

// SetInetDiscardOutsideWindow sets the value of InetDiscardOutsideWindow for the instance
func (instance *Win32_PerfRawData_TCPIPCounters_TCPIPTransportLayerPacketDropCounters) SetPropertyInetDiscardOutsideWindow(value uint64) (err error) {
	return instance.SetProperty("InetDiscardOutsideWindow", (value))
}

// GetInetDiscardOutsideWindow gets the value of InetDiscardOutsideWindow for the instance
func (instance *Win32_PerfRawData_TCPIPCounters_TCPIPTransportLayerPacketDropCounters) GetPropertyInetDiscardOutsideWindow() (value uint64, err error) {
	retValue, err := instance.GetProperty("InetDiscardOutsideWindow")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(uint64)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " uint64 is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = uint64(valuetmp)

	return
}

// SetInetDiscardPauseAccept sets the value of InetDiscardPauseAccept for the instance
func (instance *Win32_PerfRawData_TCPIPCounters_TCPIPTransportLayerPacketDropCounters) SetPropertyInetDiscardPauseAccept(value uint64) (err error) {
	return instance.SetProperty("InetDiscardPauseAccept", (value))
}

// GetInetDiscardPauseAccept gets the value of InetDiscardPauseAccept for the instance
func (instance *Win32_PerfRawData_TCPIPCounters_TCPIPTransportLayerPacketDropCounters) GetPropertyInetDiscardPauseAccept() (value uint64, err error) {
	retValue, err := instance.GetProperty("InetDiscardPauseAccept")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(uint64)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " uint64 is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = uint64(valuetmp)

	return
}

// SetInetDiscardPawsFailed sets the value of InetDiscardPawsFailed for the instance
func (instance *Win32_PerfRawData_TCPIPCounters_TCPIPTransportLayerPacketDropCounters) SetPropertyInetDiscardPawsFailed(value uint64) (err error) {
	return instance.SetProperty("InetDiscardPawsFailed", (value))
}

// GetInetDiscardPawsFailed gets the value of InetDiscardPawsFailed for the instance
func (instance *Win32_PerfRawData_TCPIPCounters_TCPIPTransportLayerPacketDropCounters) GetPropertyInetDiscardPawsFailed() (value uint64, err error) {
	retValue, err := instance.GetProperty("InetDiscardPawsFailed")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(uint64)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " uint64 is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = uint64(valuetmp)

	return
}

// SetInetDiscardReassemblyConflict sets the value of InetDiscardReassemblyConflict for the instance
func (instance *Win32_PerfRawData_TCPIPCounters_TCPIPTransportLayerPacketDropCounters) SetPropertyInetDiscardReassemblyConflict(value uint64) (err error) {
	return instance.SetProperty("InetDiscardReassemblyConflict", (value))
}

// GetInetDiscardReassemblyConflict gets the value of InetDiscardReassemblyConflict for the instance
func (instance *Win32_PerfRawData_TCPIPCounters_TCPIPTransportLayerPacketDropCounters) GetPropertyInetDiscardReassemblyConflict() (value uint64, err error) {
	retValue, err := instance.GetProperty("InetDiscardReassemblyConflict")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(uint64)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " uint64 is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = uint64(valuetmp)

	return
}

// SetInetDiscardReceiveInspection sets the value of InetDiscardReceiveInspection for the instance
func (instance *Win32_PerfRawData_TCPIPCounters_TCPIPTransportLayerPacketDropCounters) SetPropertyInetDiscardReceiveInspection(value uint64) (err error) {
	return instance.SetProperty("InetDiscardReceiveInspection", (value))
}

// GetInetDiscardReceiveInspection gets the value of InetDiscardReceiveInspection for the instance
func (instance *Win32_PerfRawData_TCPIPCounters_TCPIPTransportLayerPacketDropCounters) GetPropertyInetDiscardReceiveInspection() (value uint64, err error) {
	retValue, err := instance.GetProperty("InetDiscardReceiveInspection")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(uint64)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " uint64 is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = uint64(valuetmp)

	return
}

// SetInetDiscardRst sets the value of InetDiscardRst for the instance
func (instance *Win32_PerfRawData_TCPIPCounters_TCPIPTransportLayerPacketDropCounters) SetPropertyInetDiscardRst(value uint64) (err error) {
	return instance.SetProperty("InetDiscardRst", (value))
}

// GetInetDiscardRst gets the value of InetDiscardRst for the instance
func (instance *Win32_PerfRawData_TCPIPCounters_TCPIPTransportLayerPacketDropCounters) GetPropertyInetDiscardRst() (value uint64, err error) {
	retValue, err := instance.GetProperty("InetDiscardRst")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(uint64)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " uint64 is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = uint64(valuetmp)

	return
}

// SetInetDiscardSessionState sets the value of InetDiscardSessionState for the instance
func (instance *Win32_PerfRawData_TCPIPCounters_TCPIPTransportLayerPacketDropCounters) SetPropertyInetDiscardSessionState(value uint64) (err error) {
	return instance.SetProperty("InetDiscardSessionState", (value))
}

// GetInetDiscardSessionState gets the value of InetDiscardSessionState for the instance
func (instance *Win32_PerfRawData_TCPIPCounters_TCPIPTransportLayerPacketDropCounters) GetPropertyInetDiscardSessionState() (value uint64, err error) {
	retValue, err := instance.GetProperty("InetDiscardSessionState")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(uint64)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " uint64 is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = uint64(valuetmp)

	return
}

// SetInetDiscardSimultaneousConnect sets the value of InetDiscardSimultaneousConnect for the instance
func (instance *Win32_PerfRawData_TCPIPCounters_TCPIPTransportLayerPacketDropCounters) SetPropertyInetDiscardSimultaneousConnect(value uint64) (err error) {
	return instance.SetProperty("InetDiscardSimultaneousConnect", (value))
}

// GetInetDiscardSimultaneousConnect gets the value of InetDiscardSimultaneousConnect for the instance
func (instance *Win32_PerfRawData_TCPIPCounters_TCPIPTransportLayerPacketDropCounters) GetPropertyInetDiscardSimultaneousConnect() (value uint64, err error) {
	retValue, err := instance.GetProperty("InetDiscardSimultaneousConnect")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(uint64)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " uint64 is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = uint64(valuetmp)

	return
}

// SetInetDiscardSourceUnspecified sets the value of InetDiscardSourceUnspecified for the instance
func (instance *Win32_PerfRawData_TCPIPCounters_TCPIPTransportLayerPacketDropCounters) SetPropertyInetDiscardSourceUnspecified(value uint64) (err error) {
	return instance.SetProperty("InetDiscardSourceUnspecified", (value))
}

// GetInetDiscardSourceUnspecified gets the value of InetDiscardSourceUnspecified for the instance
func (instance *Win32_PerfRawData_TCPIPCounters_TCPIPTransportLayerPacketDropCounters) GetPropertyInetDiscardSourceUnspecified() (value uint64, err error) {
	retValue, err := instance.GetProperty("InetDiscardSourceUnspecified")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(uint64)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " uint64 is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = uint64(valuetmp)

	return
}

// SetInetDiscardSynAckWithFastopenCookieRequest sets the value of InetDiscardSynAckWithFastopenCookieRequest for the instance
func (instance *Win32_PerfRawData_TCPIPCounters_TCPIPTransportLayerPacketDropCounters) SetPropertyInetDiscardSynAckWithFastopenCookieRequest(value uint64) (err error) {
	return instance.SetProperty("InetDiscardSynAckWithFastopenCookieRequest", (value))
}

// GetInetDiscardSynAckWithFastopenCookieRequest gets the value of InetDiscardSynAckWithFastopenCookieRequest for the instance
func (instance *Win32_PerfRawData_TCPIPCounters_TCPIPTransportLayerPacketDropCounters) GetPropertyInetDiscardSynAckWithFastopenCookieRequest() (value uint64, err error) {
	retValue, err := instance.GetProperty("InetDiscardSynAckWithFastopenCookieRequest")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(uint64)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " uint64 is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = uint64(valuetmp)

	return
}

// SetInetDiscardSynAttack sets the value of InetDiscardSynAttack for the instance
func (instance *Win32_PerfRawData_TCPIPCounters_TCPIPTransportLayerPacketDropCounters) SetPropertyInetDiscardSynAttack(value uint64) (err error) {
	return instance.SetProperty("InetDiscardSynAttack", (value))
}

// GetInetDiscardSynAttack gets the value of InetDiscardSynAttack for the instance
func (instance *Win32_PerfRawData_TCPIPCounters_TCPIPTransportLayerPacketDropCounters) GetPropertyInetDiscardSynAttack() (value uint64, err error) {
	retValue, err := instance.GetProperty("InetDiscardSynAttack")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(uint64)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " uint64 is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = uint64(valuetmp)

	return
}

// SetInetDiscardSynRcvdSyn sets the value of InetDiscardSynRcvdSyn for the instance
func (instance *Win32_PerfRawData_TCPIPCounters_TCPIPTransportLayerPacketDropCounters) SetPropertyInetDiscardSynRcvdSyn(value uint64) (err error) {
	return instance.SetProperty("InetDiscardSynRcvdSyn", (value))
}

// GetInetDiscardSynRcvdSyn gets the value of InetDiscardSynRcvdSyn for the instance
func (instance *Win32_PerfRawData_TCPIPCounters_TCPIPTransportLayerPacketDropCounters) GetPropertyInetDiscardSynRcvdSyn() (value uint64, err error) {
	retValue, err := instance.GetProperty("InetDiscardSynRcvdSyn")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(uint64)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " uint64 is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = uint64(valuetmp)

	return
}

// SetInetDiscardTcbNotInTcbTable sets the value of InetDiscardTcbNotInTcbTable for the instance
func (instance *Win32_PerfRawData_TCPIPCounters_TCPIPTransportLayerPacketDropCounters) SetPropertyInetDiscardTcbNotInTcbTable(value uint64) (err error) {
	return instance.SetProperty("InetDiscardTcbNotInTcbTable", (value))
}

// GetInetDiscardTcbNotInTcbTable gets the value of InetDiscardTcbNotInTcbTable for the instance
func (instance *Win32_PerfRawData_TCPIPCounters_TCPIPTransportLayerPacketDropCounters) GetPropertyInetDiscardTcbNotInTcbTable() (value uint64, err error) {
	retValue, err := instance.GetProperty("InetDiscardTcbNotInTcbTable")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(uint64)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " uint64 is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = uint64(valuetmp)

	return
}

// SetInetDiscardTcbRemoved sets the value of InetDiscardTcbRemoved for the instance
func (instance *Win32_PerfRawData_TCPIPCounters_TCPIPTransportLayerPacketDropCounters) SetPropertyInetDiscardTcbRemoved(value uint64) (err error) {
	return instance.SetProperty("InetDiscardTcbRemoved", (value))
}

// GetInetDiscardTcbRemoved gets the value of InetDiscardTcbRemoved for the instance
func (instance *Win32_PerfRawData_TCPIPCounters_TCPIPTransportLayerPacketDropCounters) GetPropertyInetDiscardTcbRemoved() (value uint64, err error) {
	retValue, err := instance.GetProperty("InetDiscardTcbRemoved")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(uint64)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " uint64 is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = uint64(valuetmp)

	return
}

// SetInetDiscardTimeWaitTcb sets the value of InetDiscardTimeWaitTcb for the instance
func (instance *Win32_PerfRawData_TCPIPCounters_TCPIPTransportLayerPacketDropCounters) SetPropertyInetDiscardTimeWaitTcb(value uint64) (err error) {
	return instance.SetProperty("InetDiscardTimeWaitTcb", (value))
}

// GetInetDiscardTimeWaitTcb gets the value of InetDiscardTimeWaitTcb for the instance
func (instance *Win32_PerfRawData_TCPIPCounters_TCPIPTransportLayerPacketDropCounters) GetPropertyInetDiscardTimeWaitTcb() (value uint64, err error) {
	retValue, err := instance.GetProperty("InetDiscardTimeWaitTcb")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(uint64)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " uint64 is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = uint64(valuetmp)

	return
}

// SetInetDiscardTimeWaitTcbReceivedRstOutsideWindow sets the value of InetDiscardTimeWaitTcbReceivedRstOutsideWindow for the instance
func (instance *Win32_PerfRawData_TCPIPCounters_TCPIPTransportLayerPacketDropCounters) SetPropertyInetDiscardTimeWaitTcbReceivedRstOutsideWindow(value uint64) (err error) {
	return instance.SetProperty("InetDiscardTimeWaitTcbReceivedRstOutsideWindow", (value))
}

// GetInetDiscardTimeWaitTcbReceivedRstOutsideWindow gets the value of InetDiscardTimeWaitTcbReceivedRstOutsideWindow for the instance
func (instance *Win32_PerfRawData_TCPIPCounters_TCPIPTransportLayerPacketDropCounters) GetPropertyInetDiscardTimeWaitTcbReceivedRstOutsideWindow() (value uint64, err error) {
	retValue, err := instance.GetProperty("InetDiscardTimeWaitTcbReceivedRstOutsideWindow")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(uint64)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " uint64 is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = uint64(valuetmp)

	return
}

// SetInetDiscardTimeWaitTcbSynAndOtherFlags sets the value of InetDiscardTimeWaitTcbSynAndOtherFlags for the instance
func (instance *Win32_PerfRawData_TCPIPCounters_TCPIPTransportLayerPacketDropCounters) SetPropertyInetDiscardTimeWaitTcbSynAndOtherFlags(value uint64) (err error) {
	return instance.SetProperty("InetDiscardTimeWaitTcbSynAndOtherFlags", (value))
}

// GetInetDiscardTimeWaitTcbSynAndOtherFlags gets the value of InetDiscardTimeWaitTcbSynAndOtherFlags for the instance
func (instance *Win32_PerfRawData_TCPIPCounters_TCPIPTransportLayerPacketDropCounters) GetPropertyInetDiscardTimeWaitTcbSynAndOtherFlags() (value uint64, err error) {
	retValue, err := instance.GetProperty("InetDiscardTimeWaitTcbSynAndOtherFlags")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(uint64)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " uint64 is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = uint64(valuetmp)

	return
}

// SetInetDiscardUrgentDeliveryAllocationFailure sets the value of InetDiscardUrgentDeliveryAllocationFailure for the instance
func (instance *Win32_PerfRawData_TCPIPCounters_TCPIPTransportLayerPacketDropCounters) SetPropertyInetDiscardUrgentDeliveryAllocationFailure(value uint64) (err error) {
	return instance.SetProperty("InetDiscardUrgentDeliveryAllocationFailure", (value))
}

// GetInetDiscardUrgentDeliveryAllocationFailure gets the value of InetDiscardUrgentDeliveryAllocationFailure for the instance
func (instance *Win32_PerfRawData_TCPIPCounters_TCPIPTransportLayerPacketDropCounters) GetPropertyInetDiscardUrgentDeliveryAllocationFailure() (value uint64, err error) {
	retValue, err := instance.GetProperty("InetDiscardUrgentDeliveryAllocationFailure")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(uint64)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " uint64 is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = uint64(valuetmp)

	return
}
