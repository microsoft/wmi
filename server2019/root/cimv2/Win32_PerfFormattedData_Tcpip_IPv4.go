// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 3/16/2020 using wmigen
//      Source root.CIMV2
//////////////////////////////////////////////
package cimv2

// Win32_PerfFormattedData_Tcpip_IPv4 struct
type Win32_PerfFormattedData_Tcpip_IPv4 struct {
	Win32_PerfFormattedData

	//
	DatagramsForwardedPersec uint32

	//
	DatagramsOutboundDiscarded uint32

	//
	DatagramsOutboundNoRoute uint32

	//
	DatagramsPersec uint32

	//
	DatagramsReceivedAddressErrors uint32

	//
	DatagramsReceivedDeliveredPersec uint32

	//
	DatagramsReceivedDiscarded uint32

	//
	DatagramsReceivedHeaderErrors uint32

	//
	DatagramsReceivedPersec uint32

	//
	DatagramsReceivedUnknownProtocol uint32

	//
	DatagramsSentPersec uint32

	//
	FragmentationFailures uint32

	//
	FragmentedDatagramsPersec uint32

	//
	FragmentReassemblyFailures uint32

	//
	FragmentsCreatedPersec uint32

	//
	FragmentsReassembledPersec uint32

	//
	FragmentsReceivedPersec uint32
}

// SetDatagramsForwardedPersec sets the value of DatagramsForwardedPersec for the instance
func (instance *Win32_PerfFormattedData_Tcpip_IPv4) SetPropertyDatagramsForwardedPersec(value uint32) (err error) {
	return instance.SetProperty("DatagramsForwardedPersec", value)
}

// GetDatagramsForwardedPersec gets the value of DatagramsForwardedPersec for the instance
func (instance *Win32_PerfFormattedData_Tcpip_IPv4) GetPropertyDatagramsForwardedPersec() (value uint32, err error) {
	retValue, err := instance.GetProperty("DatagramsForwardedPersec")
	if err != nil {
		return
	}
	value, ok := retValue.(uint32)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetDatagramsOutboundDiscarded sets the value of DatagramsOutboundDiscarded for the instance
func (instance *Win32_PerfFormattedData_Tcpip_IPv4) SetPropertyDatagramsOutboundDiscarded(value uint32) (err error) {
	return instance.SetProperty("DatagramsOutboundDiscarded", value)
}

// GetDatagramsOutboundDiscarded gets the value of DatagramsOutboundDiscarded for the instance
func (instance *Win32_PerfFormattedData_Tcpip_IPv4) GetPropertyDatagramsOutboundDiscarded() (value uint32, err error) {
	retValue, err := instance.GetProperty("DatagramsOutboundDiscarded")
	if err != nil {
		return
	}
	value, ok := retValue.(uint32)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetDatagramsOutboundNoRoute sets the value of DatagramsOutboundNoRoute for the instance
func (instance *Win32_PerfFormattedData_Tcpip_IPv4) SetPropertyDatagramsOutboundNoRoute(value uint32) (err error) {
	return instance.SetProperty("DatagramsOutboundNoRoute", value)
}

// GetDatagramsOutboundNoRoute gets the value of DatagramsOutboundNoRoute for the instance
func (instance *Win32_PerfFormattedData_Tcpip_IPv4) GetPropertyDatagramsOutboundNoRoute() (value uint32, err error) {
	retValue, err := instance.GetProperty("DatagramsOutboundNoRoute")
	if err != nil {
		return
	}
	value, ok := retValue.(uint32)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetDatagramsPersec sets the value of DatagramsPersec for the instance
func (instance *Win32_PerfFormattedData_Tcpip_IPv4) SetPropertyDatagramsPersec(value uint32) (err error) {
	return instance.SetProperty("DatagramsPersec", value)
}

// GetDatagramsPersec gets the value of DatagramsPersec for the instance
func (instance *Win32_PerfFormattedData_Tcpip_IPv4) GetPropertyDatagramsPersec() (value uint32, err error) {
	retValue, err := instance.GetProperty("DatagramsPersec")
	if err != nil {
		return
	}
	value, ok := retValue.(uint32)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetDatagramsReceivedAddressErrors sets the value of DatagramsReceivedAddressErrors for the instance
func (instance *Win32_PerfFormattedData_Tcpip_IPv4) SetPropertyDatagramsReceivedAddressErrors(value uint32) (err error) {
	return instance.SetProperty("DatagramsReceivedAddressErrors", value)
}

// GetDatagramsReceivedAddressErrors gets the value of DatagramsReceivedAddressErrors for the instance
func (instance *Win32_PerfFormattedData_Tcpip_IPv4) GetPropertyDatagramsReceivedAddressErrors() (value uint32, err error) {
	retValue, err := instance.GetProperty("DatagramsReceivedAddressErrors")
	if err != nil {
		return
	}
	value, ok := retValue.(uint32)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetDatagramsReceivedDeliveredPersec sets the value of DatagramsReceivedDeliveredPersec for the instance
func (instance *Win32_PerfFormattedData_Tcpip_IPv4) SetPropertyDatagramsReceivedDeliveredPersec(value uint32) (err error) {
	return instance.SetProperty("DatagramsReceivedDeliveredPersec", value)
}

// GetDatagramsReceivedDeliveredPersec gets the value of DatagramsReceivedDeliveredPersec for the instance
func (instance *Win32_PerfFormattedData_Tcpip_IPv4) GetPropertyDatagramsReceivedDeliveredPersec() (value uint32, err error) {
	retValue, err := instance.GetProperty("DatagramsReceivedDeliveredPersec")
	if err != nil {
		return
	}
	value, ok := retValue.(uint32)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetDatagramsReceivedDiscarded sets the value of DatagramsReceivedDiscarded for the instance
func (instance *Win32_PerfFormattedData_Tcpip_IPv4) SetPropertyDatagramsReceivedDiscarded(value uint32) (err error) {
	return instance.SetProperty("DatagramsReceivedDiscarded", value)
}

// GetDatagramsReceivedDiscarded gets the value of DatagramsReceivedDiscarded for the instance
func (instance *Win32_PerfFormattedData_Tcpip_IPv4) GetPropertyDatagramsReceivedDiscarded() (value uint32, err error) {
	retValue, err := instance.GetProperty("DatagramsReceivedDiscarded")
	if err != nil {
		return
	}
	value, ok := retValue.(uint32)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetDatagramsReceivedHeaderErrors sets the value of DatagramsReceivedHeaderErrors for the instance
func (instance *Win32_PerfFormattedData_Tcpip_IPv4) SetPropertyDatagramsReceivedHeaderErrors(value uint32) (err error) {
	return instance.SetProperty("DatagramsReceivedHeaderErrors", value)
}

// GetDatagramsReceivedHeaderErrors gets the value of DatagramsReceivedHeaderErrors for the instance
func (instance *Win32_PerfFormattedData_Tcpip_IPv4) GetPropertyDatagramsReceivedHeaderErrors() (value uint32, err error) {
	retValue, err := instance.GetProperty("DatagramsReceivedHeaderErrors")
	if err != nil {
		return
	}
	value, ok := retValue.(uint32)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetDatagramsReceivedPersec sets the value of DatagramsReceivedPersec for the instance
func (instance *Win32_PerfFormattedData_Tcpip_IPv4) SetPropertyDatagramsReceivedPersec(value uint32) (err error) {
	return instance.SetProperty("DatagramsReceivedPersec", value)
}

// GetDatagramsReceivedPersec gets the value of DatagramsReceivedPersec for the instance
func (instance *Win32_PerfFormattedData_Tcpip_IPv4) GetPropertyDatagramsReceivedPersec() (value uint32, err error) {
	retValue, err := instance.GetProperty("DatagramsReceivedPersec")
	if err != nil {
		return
	}
	value, ok := retValue.(uint32)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetDatagramsReceivedUnknownProtocol sets the value of DatagramsReceivedUnknownProtocol for the instance
func (instance *Win32_PerfFormattedData_Tcpip_IPv4) SetPropertyDatagramsReceivedUnknownProtocol(value uint32) (err error) {
	return instance.SetProperty("DatagramsReceivedUnknownProtocol", value)
}

// GetDatagramsReceivedUnknownProtocol gets the value of DatagramsReceivedUnknownProtocol for the instance
func (instance *Win32_PerfFormattedData_Tcpip_IPv4) GetPropertyDatagramsReceivedUnknownProtocol() (value uint32, err error) {
	retValue, err := instance.GetProperty("DatagramsReceivedUnknownProtocol")
	if err != nil {
		return
	}
	value, ok := retValue.(uint32)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetDatagramsSentPersec sets the value of DatagramsSentPersec for the instance
func (instance *Win32_PerfFormattedData_Tcpip_IPv4) SetPropertyDatagramsSentPersec(value uint32) (err error) {
	return instance.SetProperty("DatagramsSentPersec", value)
}

// GetDatagramsSentPersec gets the value of DatagramsSentPersec for the instance
func (instance *Win32_PerfFormattedData_Tcpip_IPv4) GetPropertyDatagramsSentPersec() (value uint32, err error) {
	retValue, err := instance.GetProperty("DatagramsSentPersec")
	if err != nil {
		return
	}
	value, ok := retValue.(uint32)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetFragmentationFailures sets the value of FragmentationFailures for the instance
func (instance *Win32_PerfFormattedData_Tcpip_IPv4) SetPropertyFragmentationFailures(value uint32) (err error) {
	return instance.SetProperty("FragmentationFailures", value)
}

// GetFragmentationFailures gets the value of FragmentationFailures for the instance
func (instance *Win32_PerfFormattedData_Tcpip_IPv4) GetPropertyFragmentationFailures() (value uint32, err error) {
	retValue, err := instance.GetProperty("FragmentationFailures")
	if err != nil {
		return
	}
	value, ok := retValue.(uint32)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetFragmentedDatagramsPersec sets the value of FragmentedDatagramsPersec for the instance
func (instance *Win32_PerfFormattedData_Tcpip_IPv4) SetPropertyFragmentedDatagramsPersec(value uint32) (err error) {
	return instance.SetProperty("FragmentedDatagramsPersec", value)
}

// GetFragmentedDatagramsPersec gets the value of FragmentedDatagramsPersec for the instance
func (instance *Win32_PerfFormattedData_Tcpip_IPv4) GetPropertyFragmentedDatagramsPersec() (value uint32, err error) {
	retValue, err := instance.GetProperty("FragmentedDatagramsPersec")
	if err != nil {
		return
	}
	value, ok := retValue.(uint32)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetFragmentReassemblyFailures sets the value of FragmentReassemblyFailures for the instance
func (instance *Win32_PerfFormattedData_Tcpip_IPv4) SetPropertyFragmentReassemblyFailures(value uint32) (err error) {
	return instance.SetProperty("FragmentReassemblyFailures", value)
}

// GetFragmentReassemblyFailures gets the value of FragmentReassemblyFailures for the instance
func (instance *Win32_PerfFormattedData_Tcpip_IPv4) GetPropertyFragmentReassemblyFailures() (value uint32, err error) {
	retValue, err := instance.GetProperty("FragmentReassemblyFailures")
	if err != nil {
		return
	}
	value, ok := retValue.(uint32)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetFragmentsCreatedPersec sets the value of FragmentsCreatedPersec for the instance
func (instance *Win32_PerfFormattedData_Tcpip_IPv4) SetPropertyFragmentsCreatedPersec(value uint32) (err error) {
	return instance.SetProperty("FragmentsCreatedPersec", value)
}

// GetFragmentsCreatedPersec gets the value of FragmentsCreatedPersec for the instance
func (instance *Win32_PerfFormattedData_Tcpip_IPv4) GetPropertyFragmentsCreatedPersec() (value uint32, err error) {
	retValue, err := instance.GetProperty("FragmentsCreatedPersec")
	if err != nil {
		return
	}
	value, ok := retValue.(uint32)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetFragmentsReassembledPersec sets the value of FragmentsReassembledPersec for the instance
func (instance *Win32_PerfFormattedData_Tcpip_IPv4) SetPropertyFragmentsReassembledPersec(value uint32) (err error) {
	return instance.SetProperty("FragmentsReassembledPersec", value)
}

// GetFragmentsReassembledPersec gets the value of FragmentsReassembledPersec for the instance
func (instance *Win32_PerfFormattedData_Tcpip_IPv4) GetPropertyFragmentsReassembledPersec() (value uint32, err error) {
	retValue, err := instance.GetProperty("FragmentsReassembledPersec")
	if err != nil {
		return
	}
	value, ok := retValue.(uint32)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetFragmentsReceivedPersec sets the value of FragmentsReceivedPersec for the instance
func (instance *Win32_PerfFormattedData_Tcpip_IPv4) SetPropertyFragmentsReceivedPersec(value uint32) (err error) {
	return instance.SetProperty("FragmentsReceivedPersec", value)
}

// GetFragmentsReceivedPersec gets the value of FragmentsReceivedPersec for the instance
func (instance *Win32_PerfFormattedData_Tcpip_IPv4) GetPropertyFragmentsReceivedPersec() (value uint32, err error) {
	retValue, err := instance.GetProperty("FragmentsReceivedPersec")
	if err != nil {
		return
	}
	value, ok := retValue.(uint32)
	if !ok {
		// TODO: Set an error
	}
	return
}
