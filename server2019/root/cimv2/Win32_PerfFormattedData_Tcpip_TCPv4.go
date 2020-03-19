// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 3/19/2020 using wmigen
//      Source root.CIMV2
//////////////////////////////////////////////
package cimv2

import (
	"github.com/microsoft/wmi/pkg/base/query"
	cim "github.com/microsoft/wmi/pkg/wmiinstance"
)

// Win32_PerfFormattedData_Tcpip_TCPv4 struct
type Win32_PerfFormattedData_Tcpip_TCPv4 struct {
	*Win32_PerfFormattedData

	//
	ConnectionFailures uint32

	//
	ConnectionsActive uint32

	//
	ConnectionsEstablished uint32

	//
	ConnectionsPassive uint32

	//
	ConnectionsReset uint32

	//
	SegmentsPersec uint32

	//
	SegmentsReceivedPersec uint32

	//
	SegmentsRetransmittedPersec uint32

	//
	SegmentsSentPersec uint32
}

func NewWin32_PerfFormattedData_Tcpip_TCPv4Ex1(instance *cim.WmiInstance) (newInstance *Win32_PerfFormattedData_Tcpip_TCPv4, err error) {
	tmp, err := NewWin32_PerfFormattedDataEx1(instance)

	if err != nil {
		return
	}
	newInstance = &Win32_PerfFormattedData_Tcpip_TCPv4{
		Win32_PerfFormattedData: tmp,
	}
	return
}

func NewWin32_PerfFormattedData_Tcpip_TCPv4Ex6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *Win32_PerfFormattedData_Tcpip_TCPv4, err error) {
	tmp, err := NewWin32_PerfFormattedDataEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &Win32_PerfFormattedData_Tcpip_TCPv4{
		Win32_PerfFormattedData: tmp,
	}
	return
}

// SetConnectionFailures sets the value of ConnectionFailures for the instance
func (instance *Win32_PerfFormattedData_Tcpip_TCPv4) SetPropertyConnectionFailures(value uint32) (err error) {
	return instance.SetProperty("ConnectionFailures", value)
}

// GetConnectionFailures gets the value of ConnectionFailures for the instance
func (instance *Win32_PerfFormattedData_Tcpip_TCPv4) GetPropertyConnectionFailures() (value uint32, err error) {
	retValue, err := instance.GetProperty("ConnectionFailures")
	if err != nil {
		return
	}
	value, ok := retValue.(uint32)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetConnectionsActive sets the value of ConnectionsActive for the instance
func (instance *Win32_PerfFormattedData_Tcpip_TCPv4) SetPropertyConnectionsActive(value uint32) (err error) {
	return instance.SetProperty("ConnectionsActive", value)
}

// GetConnectionsActive gets the value of ConnectionsActive for the instance
func (instance *Win32_PerfFormattedData_Tcpip_TCPv4) GetPropertyConnectionsActive() (value uint32, err error) {
	retValue, err := instance.GetProperty("ConnectionsActive")
	if err != nil {
		return
	}
	value, ok := retValue.(uint32)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetConnectionsEstablished sets the value of ConnectionsEstablished for the instance
func (instance *Win32_PerfFormattedData_Tcpip_TCPv4) SetPropertyConnectionsEstablished(value uint32) (err error) {
	return instance.SetProperty("ConnectionsEstablished", value)
}

// GetConnectionsEstablished gets the value of ConnectionsEstablished for the instance
func (instance *Win32_PerfFormattedData_Tcpip_TCPv4) GetPropertyConnectionsEstablished() (value uint32, err error) {
	retValue, err := instance.GetProperty("ConnectionsEstablished")
	if err != nil {
		return
	}
	value, ok := retValue.(uint32)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetConnectionsPassive sets the value of ConnectionsPassive for the instance
func (instance *Win32_PerfFormattedData_Tcpip_TCPv4) SetPropertyConnectionsPassive(value uint32) (err error) {
	return instance.SetProperty("ConnectionsPassive", value)
}

// GetConnectionsPassive gets the value of ConnectionsPassive for the instance
func (instance *Win32_PerfFormattedData_Tcpip_TCPv4) GetPropertyConnectionsPassive() (value uint32, err error) {
	retValue, err := instance.GetProperty("ConnectionsPassive")
	if err != nil {
		return
	}
	value, ok := retValue.(uint32)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetConnectionsReset sets the value of ConnectionsReset for the instance
func (instance *Win32_PerfFormattedData_Tcpip_TCPv4) SetPropertyConnectionsReset(value uint32) (err error) {
	return instance.SetProperty("ConnectionsReset", value)
}

// GetConnectionsReset gets the value of ConnectionsReset for the instance
func (instance *Win32_PerfFormattedData_Tcpip_TCPv4) GetPropertyConnectionsReset() (value uint32, err error) {
	retValue, err := instance.GetProperty("ConnectionsReset")
	if err != nil {
		return
	}
	value, ok := retValue.(uint32)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetSegmentsPersec sets the value of SegmentsPersec for the instance
func (instance *Win32_PerfFormattedData_Tcpip_TCPv4) SetPropertySegmentsPersec(value uint32) (err error) {
	return instance.SetProperty("SegmentsPersec", value)
}

// GetSegmentsPersec gets the value of SegmentsPersec for the instance
func (instance *Win32_PerfFormattedData_Tcpip_TCPv4) GetPropertySegmentsPersec() (value uint32, err error) {
	retValue, err := instance.GetProperty("SegmentsPersec")
	if err != nil {
		return
	}
	value, ok := retValue.(uint32)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetSegmentsReceivedPersec sets the value of SegmentsReceivedPersec for the instance
func (instance *Win32_PerfFormattedData_Tcpip_TCPv4) SetPropertySegmentsReceivedPersec(value uint32) (err error) {
	return instance.SetProperty("SegmentsReceivedPersec", value)
}

// GetSegmentsReceivedPersec gets the value of SegmentsReceivedPersec for the instance
func (instance *Win32_PerfFormattedData_Tcpip_TCPv4) GetPropertySegmentsReceivedPersec() (value uint32, err error) {
	retValue, err := instance.GetProperty("SegmentsReceivedPersec")
	if err != nil {
		return
	}
	value, ok := retValue.(uint32)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetSegmentsRetransmittedPersec sets the value of SegmentsRetransmittedPersec for the instance
func (instance *Win32_PerfFormattedData_Tcpip_TCPv4) SetPropertySegmentsRetransmittedPersec(value uint32) (err error) {
	return instance.SetProperty("SegmentsRetransmittedPersec", value)
}

// GetSegmentsRetransmittedPersec gets the value of SegmentsRetransmittedPersec for the instance
func (instance *Win32_PerfFormattedData_Tcpip_TCPv4) GetPropertySegmentsRetransmittedPersec() (value uint32, err error) {
	retValue, err := instance.GetProperty("SegmentsRetransmittedPersec")
	if err != nil {
		return
	}
	value, ok := retValue.(uint32)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetSegmentsSentPersec sets the value of SegmentsSentPersec for the instance
func (instance *Win32_PerfFormattedData_Tcpip_TCPv4) SetPropertySegmentsSentPersec(value uint32) (err error) {
	return instance.SetProperty("SegmentsSentPersec", value)
}

// GetSegmentsSentPersec gets the value of SegmentsSentPersec for the instance
func (instance *Win32_PerfFormattedData_Tcpip_TCPv4) GetPropertySegmentsSentPersec() (value uint32, err error) {
	retValue, err := instance.GetProperty("SegmentsSentPersec")
	if err != nil {
		return
	}
	value, ok := retValue.(uint32)
	if !ok {
		// TODO: Set an error
	}
	return
}
