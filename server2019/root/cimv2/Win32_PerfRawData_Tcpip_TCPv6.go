// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 3/16/2020 using wmigen
//      Source root.CIMV2
//////////////////////////////////////////////
package cimv2

// Win32_PerfRawData_Tcpip_TCPv6 struct
type Win32_PerfRawData_Tcpip_TCPv6 struct {
	Win32_PerfRawData

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

// SetConnectionFailures sets the value of ConnectionFailures for the instance
func (instance *Win32_PerfRawData_Tcpip_TCPv6) SetPropertyConnectionFailures(value uint32) (err error) {
	return instance.SetProperty("ConnectionFailures", value)
}

// GetConnectionFailures gets the value of ConnectionFailures for the instance
func (instance *Win32_PerfRawData_Tcpip_TCPv6) GetPropertyConnectionFailures() (value uint32, err error) {
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
func (instance *Win32_PerfRawData_Tcpip_TCPv6) SetPropertyConnectionsActive(value uint32) (err error) {
	return instance.SetProperty("ConnectionsActive", value)
}

// GetConnectionsActive gets the value of ConnectionsActive for the instance
func (instance *Win32_PerfRawData_Tcpip_TCPv6) GetPropertyConnectionsActive() (value uint32, err error) {
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
func (instance *Win32_PerfRawData_Tcpip_TCPv6) SetPropertyConnectionsEstablished(value uint32) (err error) {
	return instance.SetProperty("ConnectionsEstablished", value)
}

// GetConnectionsEstablished gets the value of ConnectionsEstablished for the instance
func (instance *Win32_PerfRawData_Tcpip_TCPv6) GetPropertyConnectionsEstablished() (value uint32, err error) {
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
func (instance *Win32_PerfRawData_Tcpip_TCPv6) SetPropertyConnectionsPassive(value uint32) (err error) {
	return instance.SetProperty("ConnectionsPassive", value)
}

// GetConnectionsPassive gets the value of ConnectionsPassive for the instance
func (instance *Win32_PerfRawData_Tcpip_TCPv6) GetPropertyConnectionsPassive() (value uint32, err error) {
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
func (instance *Win32_PerfRawData_Tcpip_TCPv6) SetPropertyConnectionsReset(value uint32) (err error) {
	return instance.SetProperty("ConnectionsReset", value)
}

// GetConnectionsReset gets the value of ConnectionsReset for the instance
func (instance *Win32_PerfRawData_Tcpip_TCPv6) GetPropertyConnectionsReset() (value uint32, err error) {
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
func (instance *Win32_PerfRawData_Tcpip_TCPv6) SetPropertySegmentsPersec(value uint32) (err error) {
	return instance.SetProperty("SegmentsPersec", value)
}

// GetSegmentsPersec gets the value of SegmentsPersec for the instance
func (instance *Win32_PerfRawData_Tcpip_TCPv6) GetPropertySegmentsPersec() (value uint32, err error) {
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
func (instance *Win32_PerfRawData_Tcpip_TCPv6) SetPropertySegmentsReceivedPersec(value uint32) (err error) {
	return instance.SetProperty("SegmentsReceivedPersec", value)
}

// GetSegmentsReceivedPersec gets the value of SegmentsReceivedPersec for the instance
func (instance *Win32_PerfRawData_Tcpip_TCPv6) GetPropertySegmentsReceivedPersec() (value uint32, err error) {
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
func (instance *Win32_PerfRawData_Tcpip_TCPv6) SetPropertySegmentsRetransmittedPersec(value uint32) (err error) {
	return instance.SetProperty("SegmentsRetransmittedPersec", value)
}

// GetSegmentsRetransmittedPersec gets the value of SegmentsRetransmittedPersec for the instance
func (instance *Win32_PerfRawData_Tcpip_TCPv6) GetPropertySegmentsRetransmittedPersec() (value uint32, err error) {
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
func (instance *Win32_PerfRawData_Tcpip_TCPv6) SetPropertySegmentsSentPersec(value uint32) (err error) {
	return instance.SetProperty("SegmentsSentPersec", value)
}

// GetSegmentsSentPersec gets the value of SegmentsSentPersec for the instance
func (instance *Win32_PerfRawData_Tcpip_TCPv6) GetPropertySegmentsSentPersec() (value uint32, err error) {
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
