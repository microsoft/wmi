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

// Win32_PerfFormattedData_TCPIPCounters_TCPIPPerformanceDiagnostics struct
type Win32_PerfFormattedData_TCPIPCounters_TCPIPPerformanceDiagnostics struct {
	*Win32_PerfFormattedData

	//
	Deniedconnectorsendrequestsinlowpowermode uint32

	//
	IPv4NBLsindicatedwithlowresourceflag uint32

	//
	IPv4NBLsindicatedwithoutprevalidation uint32

	//
	IPv4NBLsPersecindicatedwithlowresourceflag uint32

	//
	IPv4NBLsPersecindicatedwithoutprevalidation uint32

	//
	IPv4NBLsPersectreatedasnonprevalidated uint32

	//
	IPv4NBLstreatedasnonprevalidated uint32

	//
	IPv4outboundNBLsnotprocessedviafastpath uint32

	//
	IPv4outboundNBLsPersecnotprocessedviafastpath uint32

	//
	IPv6NBLsindicatedwithlowresourceflag uint32

	//
	IPv6NBLsindicatedwithoutprevalidation uint32

	//
	IPv6NBLsPersecindicatedwithlowresourceflag uint32

	//
	IPv6NBLsPersecindicatedwithoutprevalidation uint32

	//
	IPv6NBLsPersectreatedasnonprevalidated uint32

	//
	IPv6NBLstreatedasnonprevalidated uint32

	//
	IPv6outboundNBLsnotprocessedviafastpath uint32

	//
	IPv6outboundNBLsPersecnotprocessedviafastpath uint32

	//
	RSCsegmentforwardingfailuresduringsoftwaresegmentation uint32

	//
	RSCsegmentsforwardedviaLSO uint32

	//
	RSCsegmentsforwardedviasoftwaresegmentation uint32

	//
	RSCsegmentsforwardedviasoftwaresegmentationandchecksum uint32

	//
	TCPchecksumerrors uint32

	//
	TCPconnectrequestsfallenoffloopbackfastpath uint32

	//
	TCPconnectrequestsPersecfallenoffloopbackfastpath uint32

	//
	TCPinboundsegmentsnotprocessedviafastpath uint32

	//
	TCPinboundsegmentsPersecnotprocessedviafastpath uint32

	//
	TCPRSCbytesreceived uint32

	//
	TCPRSCevents uint32

	//
	TCPtimeouts uint32

	//
	UDPdatagramscreatedviasoftwaresegmentation uint32

	//
	UDPURObytesreceived uint32

	//
	UDPUROevents uint32

	//
	UROsegmentforwardingfailuresduringsoftwaresegmentation uint32

	//
	UROsegmentsforwardedviasoftwaresegmentation uint32

	//
	UROsegmentsforwardedviasoftwaresegmentationandchecksum uint32
}

func NewWin32_PerfFormattedData_TCPIPCounters_TCPIPPerformanceDiagnosticsEx1(instance *cim.WmiInstance) (newInstance *Win32_PerfFormattedData_TCPIPCounters_TCPIPPerformanceDiagnostics, err error) {
	tmp, err := NewWin32_PerfFormattedDataEx1(instance)

	if err != nil {
		return
	}
	newInstance = &Win32_PerfFormattedData_TCPIPCounters_TCPIPPerformanceDiagnostics{
		Win32_PerfFormattedData: tmp,
	}
	return
}

func NewWin32_PerfFormattedData_TCPIPCounters_TCPIPPerformanceDiagnosticsEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *Win32_PerfFormattedData_TCPIPCounters_TCPIPPerformanceDiagnostics, err error) {
	tmp, err := NewWin32_PerfFormattedDataEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &Win32_PerfFormattedData_TCPIPCounters_TCPIPPerformanceDiagnostics{
		Win32_PerfFormattedData: tmp,
	}
	return
}

// SetDeniedconnectorsendrequestsinlowpowermode sets the value of Deniedconnectorsendrequestsinlowpowermode for the instance
func (instance *Win32_PerfFormattedData_TCPIPCounters_TCPIPPerformanceDiagnostics) SetPropertyDeniedconnectorsendrequestsinlowpowermode(value uint32) (err error) {
	return instance.SetProperty("Deniedconnectorsendrequestsinlowpowermode", value)
}

// GetDeniedconnectorsendrequestsinlowpowermode gets the value of Deniedconnectorsendrequestsinlowpowermode for the instance
func (instance *Win32_PerfFormattedData_TCPIPCounters_TCPIPPerformanceDiagnostics) GetPropertyDeniedconnectorsendrequestsinlowpowermode() (value uint32, err error) {
	retValue, err := instance.GetProperty("Deniedconnectorsendrequestsinlowpowermode")
	if err != nil {
		return
	}
	value, ok := retValue.(uint32)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetIPv4NBLsindicatedwithlowresourceflag sets the value of IPv4NBLsindicatedwithlowresourceflag for the instance
func (instance *Win32_PerfFormattedData_TCPIPCounters_TCPIPPerformanceDiagnostics) SetPropertyIPv4NBLsindicatedwithlowresourceflag(value uint32) (err error) {
	return instance.SetProperty("IPv4NBLsindicatedwithlowresourceflag", value)
}

// GetIPv4NBLsindicatedwithlowresourceflag gets the value of IPv4NBLsindicatedwithlowresourceflag for the instance
func (instance *Win32_PerfFormattedData_TCPIPCounters_TCPIPPerformanceDiagnostics) GetPropertyIPv4NBLsindicatedwithlowresourceflag() (value uint32, err error) {
	retValue, err := instance.GetProperty("IPv4NBLsindicatedwithlowresourceflag")
	if err != nil {
		return
	}
	value, ok := retValue.(uint32)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetIPv4NBLsindicatedwithoutprevalidation sets the value of IPv4NBLsindicatedwithoutprevalidation for the instance
func (instance *Win32_PerfFormattedData_TCPIPCounters_TCPIPPerformanceDiagnostics) SetPropertyIPv4NBLsindicatedwithoutprevalidation(value uint32) (err error) {
	return instance.SetProperty("IPv4NBLsindicatedwithoutprevalidation", value)
}

// GetIPv4NBLsindicatedwithoutprevalidation gets the value of IPv4NBLsindicatedwithoutprevalidation for the instance
func (instance *Win32_PerfFormattedData_TCPIPCounters_TCPIPPerformanceDiagnostics) GetPropertyIPv4NBLsindicatedwithoutprevalidation() (value uint32, err error) {
	retValue, err := instance.GetProperty("IPv4NBLsindicatedwithoutprevalidation")
	if err != nil {
		return
	}
	value, ok := retValue.(uint32)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetIPv4NBLsPersecindicatedwithlowresourceflag sets the value of IPv4NBLsPersecindicatedwithlowresourceflag for the instance
func (instance *Win32_PerfFormattedData_TCPIPCounters_TCPIPPerformanceDiagnostics) SetPropertyIPv4NBLsPersecindicatedwithlowresourceflag(value uint32) (err error) {
	return instance.SetProperty("IPv4NBLsPersecindicatedwithlowresourceflag", value)
}

// GetIPv4NBLsPersecindicatedwithlowresourceflag gets the value of IPv4NBLsPersecindicatedwithlowresourceflag for the instance
func (instance *Win32_PerfFormattedData_TCPIPCounters_TCPIPPerformanceDiagnostics) GetPropertyIPv4NBLsPersecindicatedwithlowresourceflag() (value uint32, err error) {
	retValue, err := instance.GetProperty("IPv4NBLsPersecindicatedwithlowresourceflag")
	if err != nil {
		return
	}
	value, ok := retValue.(uint32)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetIPv4NBLsPersecindicatedwithoutprevalidation sets the value of IPv4NBLsPersecindicatedwithoutprevalidation for the instance
func (instance *Win32_PerfFormattedData_TCPIPCounters_TCPIPPerformanceDiagnostics) SetPropertyIPv4NBLsPersecindicatedwithoutprevalidation(value uint32) (err error) {
	return instance.SetProperty("IPv4NBLsPersecindicatedwithoutprevalidation", value)
}

// GetIPv4NBLsPersecindicatedwithoutprevalidation gets the value of IPv4NBLsPersecindicatedwithoutprevalidation for the instance
func (instance *Win32_PerfFormattedData_TCPIPCounters_TCPIPPerformanceDiagnostics) GetPropertyIPv4NBLsPersecindicatedwithoutprevalidation() (value uint32, err error) {
	retValue, err := instance.GetProperty("IPv4NBLsPersecindicatedwithoutprevalidation")
	if err != nil {
		return
	}
	value, ok := retValue.(uint32)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetIPv4NBLsPersectreatedasnonprevalidated sets the value of IPv4NBLsPersectreatedasnonprevalidated for the instance
func (instance *Win32_PerfFormattedData_TCPIPCounters_TCPIPPerformanceDiagnostics) SetPropertyIPv4NBLsPersectreatedasnonprevalidated(value uint32) (err error) {
	return instance.SetProperty("IPv4NBLsPersectreatedasnonprevalidated", value)
}

// GetIPv4NBLsPersectreatedasnonprevalidated gets the value of IPv4NBLsPersectreatedasnonprevalidated for the instance
func (instance *Win32_PerfFormattedData_TCPIPCounters_TCPIPPerformanceDiagnostics) GetPropertyIPv4NBLsPersectreatedasnonprevalidated() (value uint32, err error) {
	retValue, err := instance.GetProperty("IPv4NBLsPersectreatedasnonprevalidated")
	if err != nil {
		return
	}
	value, ok := retValue.(uint32)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetIPv4NBLstreatedasnonprevalidated sets the value of IPv4NBLstreatedasnonprevalidated for the instance
func (instance *Win32_PerfFormattedData_TCPIPCounters_TCPIPPerformanceDiagnostics) SetPropertyIPv4NBLstreatedasnonprevalidated(value uint32) (err error) {
	return instance.SetProperty("IPv4NBLstreatedasnonprevalidated", value)
}

// GetIPv4NBLstreatedasnonprevalidated gets the value of IPv4NBLstreatedasnonprevalidated for the instance
func (instance *Win32_PerfFormattedData_TCPIPCounters_TCPIPPerformanceDiagnostics) GetPropertyIPv4NBLstreatedasnonprevalidated() (value uint32, err error) {
	retValue, err := instance.GetProperty("IPv4NBLstreatedasnonprevalidated")
	if err != nil {
		return
	}
	value, ok := retValue.(uint32)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetIPv4outboundNBLsnotprocessedviafastpath sets the value of IPv4outboundNBLsnotprocessedviafastpath for the instance
func (instance *Win32_PerfFormattedData_TCPIPCounters_TCPIPPerformanceDiagnostics) SetPropertyIPv4outboundNBLsnotprocessedviafastpath(value uint32) (err error) {
	return instance.SetProperty("IPv4outboundNBLsnotprocessedviafastpath", value)
}

// GetIPv4outboundNBLsnotprocessedviafastpath gets the value of IPv4outboundNBLsnotprocessedviafastpath for the instance
func (instance *Win32_PerfFormattedData_TCPIPCounters_TCPIPPerformanceDiagnostics) GetPropertyIPv4outboundNBLsnotprocessedviafastpath() (value uint32, err error) {
	retValue, err := instance.GetProperty("IPv4outboundNBLsnotprocessedviafastpath")
	if err != nil {
		return
	}
	value, ok := retValue.(uint32)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetIPv4outboundNBLsPersecnotprocessedviafastpath sets the value of IPv4outboundNBLsPersecnotprocessedviafastpath for the instance
func (instance *Win32_PerfFormattedData_TCPIPCounters_TCPIPPerformanceDiagnostics) SetPropertyIPv4outboundNBLsPersecnotprocessedviafastpath(value uint32) (err error) {
	return instance.SetProperty("IPv4outboundNBLsPersecnotprocessedviafastpath", value)
}

// GetIPv4outboundNBLsPersecnotprocessedviafastpath gets the value of IPv4outboundNBLsPersecnotprocessedviafastpath for the instance
func (instance *Win32_PerfFormattedData_TCPIPCounters_TCPIPPerformanceDiagnostics) GetPropertyIPv4outboundNBLsPersecnotprocessedviafastpath() (value uint32, err error) {
	retValue, err := instance.GetProperty("IPv4outboundNBLsPersecnotprocessedviafastpath")
	if err != nil {
		return
	}
	value, ok := retValue.(uint32)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetIPv6NBLsindicatedwithlowresourceflag sets the value of IPv6NBLsindicatedwithlowresourceflag for the instance
func (instance *Win32_PerfFormattedData_TCPIPCounters_TCPIPPerformanceDiagnostics) SetPropertyIPv6NBLsindicatedwithlowresourceflag(value uint32) (err error) {
	return instance.SetProperty("IPv6NBLsindicatedwithlowresourceflag", value)
}

// GetIPv6NBLsindicatedwithlowresourceflag gets the value of IPv6NBLsindicatedwithlowresourceflag for the instance
func (instance *Win32_PerfFormattedData_TCPIPCounters_TCPIPPerformanceDiagnostics) GetPropertyIPv6NBLsindicatedwithlowresourceflag() (value uint32, err error) {
	retValue, err := instance.GetProperty("IPv6NBLsindicatedwithlowresourceflag")
	if err != nil {
		return
	}
	value, ok := retValue.(uint32)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetIPv6NBLsindicatedwithoutprevalidation sets the value of IPv6NBLsindicatedwithoutprevalidation for the instance
func (instance *Win32_PerfFormattedData_TCPIPCounters_TCPIPPerformanceDiagnostics) SetPropertyIPv6NBLsindicatedwithoutprevalidation(value uint32) (err error) {
	return instance.SetProperty("IPv6NBLsindicatedwithoutprevalidation", value)
}

// GetIPv6NBLsindicatedwithoutprevalidation gets the value of IPv6NBLsindicatedwithoutprevalidation for the instance
func (instance *Win32_PerfFormattedData_TCPIPCounters_TCPIPPerformanceDiagnostics) GetPropertyIPv6NBLsindicatedwithoutprevalidation() (value uint32, err error) {
	retValue, err := instance.GetProperty("IPv6NBLsindicatedwithoutprevalidation")
	if err != nil {
		return
	}
	value, ok := retValue.(uint32)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetIPv6NBLsPersecindicatedwithlowresourceflag sets the value of IPv6NBLsPersecindicatedwithlowresourceflag for the instance
func (instance *Win32_PerfFormattedData_TCPIPCounters_TCPIPPerformanceDiagnostics) SetPropertyIPv6NBLsPersecindicatedwithlowresourceflag(value uint32) (err error) {
	return instance.SetProperty("IPv6NBLsPersecindicatedwithlowresourceflag", value)
}

// GetIPv6NBLsPersecindicatedwithlowresourceflag gets the value of IPv6NBLsPersecindicatedwithlowresourceflag for the instance
func (instance *Win32_PerfFormattedData_TCPIPCounters_TCPIPPerformanceDiagnostics) GetPropertyIPv6NBLsPersecindicatedwithlowresourceflag() (value uint32, err error) {
	retValue, err := instance.GetProperty("IPv6NBLsPersecindicatedwithlowresourceflag")
	if err != nil {
		return
	}
	value, ok := retValue.(uint32)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetIPv6NBLsPersecindicatedwithoutprevalidation sets the value of IPv6NBLsPersecindicatedwithoutprevalidation for the instance
func (instance *Win32_PerfFormattedData_TCPIPCounters_TCPIPPerformanceDiagnostics) SetPropertyIPv6NBLsPersecindicatedwithoutprevalidation(value uint32) (err error) {
	return instance.SetProperty("IPv6NBLsPersecindicatedwithoutprevalidation", value)
}

// GetIPv6NBLsPersecindicatedwithoutprevalidation gets the value of IPv6NBLsPersecindicatedwithoutprevalidation for the instance
func (instance *Win32_PerfFormattedData_TCPIPCounters_TCPIPPerformanceDiagnostics) GetPropertyIPv6NBLsPersecindicatedwithoutprevalidation() (value uint32, err error) {
	retValue, err := instance.GetProperty("IPv6NBLsPersecindicatedwithoutprevalidation")
	if err != nil {
		return
	}
	value, ok := retValue.(uint32)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetIPv6NBLsPersectreatedasnonprevalidated sets the value of IPv6NBLsPersectreatedasnonprevalidated for the instance
func (instance *Win32_PerfFormattedData_TCPIPCounters_TCPIPPerformanceDiagnostics) SetPropertyIPv6NBLsPersectreatedasnonprevalidated(value uint32) (err error) {
	return instance.SetProperty("IPv6NBLsPersectreatedasnonprevalidated", value)
}

// GetIPv6NBLsPersectreatedasnonprevalidated gets the value of IPv6NBLsPersectreatedasnonprevalidated for the instance
func (instance *Win32_PerfFormattedData_TCPIPCounters_TCPIPPerformanceDiagnostics) GetPropertyIPv6NBLsPersectreatedasnonprevalidated() (value uint32, err error) {
	retValue, err := instance.GetProperty("IPv6NBLsPersectreatedasnonprevalidated")
	if err != nil {
		return
	}
	value, ok := retValue.(uint32)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetIPv6NBLstreatedasnonprevalidated sets the value of IPv6NBLstreatedasnonprevalidated for the instance
func (instance *Win32_PerfFormattedData_TCPIPCounters_TCPIPPerformanceDiagnostics) SetPropertyIPv6NBLstreatedasnonprevalidated(value uint32) (err error) {
	return instance.SetProperty("IPv6NBLstreatedasnonprevalidated", value)
}

// GetIPv6NBLstreatedasnonprevalidated gets the value of IPv6NBLstreatedasnonprevalidated for the instance
func (instance *Win32_PerfFormattedData_TCPIPCounters_TCPIPPerformanceDiagnostics) GetPropertyIPv6NBLstreatedasnonprevalidated() (value uint32, err error) {
	retValue, err := instance.GetProperty("IPv6NBLstreatedasnonprevalidated")
	if err != nil {
		return
	}
	value, ok := retValue.(uint32)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetIPv6outboundNBLsnotprocessedviafastpath sets the value of IPv6outboundNBLsnotprocessedviafastpath for the instance
func (instance *Win32_PerfFormattedData_TCPIPCounters_TCPIPPerformanceDiagnostics) SetPropertyIPv6outboundNBLsnotprocessedviafastpath(value uint32) (err error) {
	return instance.SetProperty("IPv6outboundNBLsnotprocessedviafastpath", value)
}

// GetIPv6outboundNBLsnotprocessedviafastpath gets the value of IPv6outboundNBLsnotprocessedviafastpath for the instance
func (instance *Win32_PerfFormattedData_TCPIPCounters_TCPIPPerformanceDiagnostics) GetPropertyIPv6outboundNBLsnotprocessedviafastpath() (value uint32, err error) {
	retValue, err := instance.GetProperty("IPv6outboundNBLsnotprocessedviafastpath")
	if err != nil {
		return
	}
	value, ok := retValue.(uint32)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetIPv6outboundNBLsPersecnotprocessedviafastpath sets the value of IPv6outboundNBLsPersecnotprocessedviafastpath for the instance
func (instance *Win32_PerfFormattedData_TCPIPCounters_TCPIPPerformanceDiagnostics) SetPropertyIPv6outboundNBLsPersecnotprocessedviafastpath(value uint32) (err error) {
	return instance.SetProperty("IPv6outboundNBLsPersecnotprocessedviafastpath", value)
}

// GetIPv6outboundNBLsPersecnotprocessedviafastpath gets the value of IPv6outboundNBLsPersecnotprocessedviafastpath for the instance
func (instance *Win32_PerfFormattedData_TCPIPCounters_TCPIPPerformanceDiagnostics) GetPropertyIPv6outboundNBLsPersecnotprocessedviafastpath() (value uint32, err error) {
	retValue, err := instance.GetProperty("IPv6outboundNBLsPersecnotprocessedviafastpath")
	if err != nil {
		return
	}
	value, ok := retValue.(uint32)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetRSCsegmentforwardingfailuresduringsoftwaresegmentation sets the value of RSCsegmentforwardingfailuresduringsoftwaresegmentation for the instance
func (instance *Win32_PerfFormattedData_TCPIPCounters_TCPIPPerformanceDiagnostics) SetPropertyRSCsegmentforwardingfailuresduringsoftwaresegmentation(value uint32) (err error) {
	return instance.SetProperty("RSCsegmentforwardingfailuresduringsoftwaresegmentation", value)
}

// GetRSCsegmentforwardingfailuresduringsoftwaresegmentation gets the value of RSCsegmentforwardingfailuresduringsoftwaresegmentation for the instance
func (instance *Win32_PerfFormattedData_TCPIPCounters_TCPIPPerformanceDiagnostics) GetPropertyRSCsegmentforwardingfailuresduringsoftwaresegmentation() (value uint32, err error) {
	retValue, err := instance.GetProperty("RSCsegmentforwardingfailuresduringsoftwaresegmentation")
	if err != nil {
		return
	}
	value, ok := retValue.(uint32)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetRSCsegmentsforwardedviaLSO sets the value of RSCsegmentsforwardedviaLSO for the instance
func (instance *Win32_PerfFormattedData_TCPIPCounters_TCPIPPerformanceDiagnostics) SetPropertyRSCsegmentsforwardedviaLSO(value uint32) (err error) {
	return instance.SetProperty("RSCsegmentsforwardedviaLSO", value)
}

// GetRSCsegmentsforwardedviaLSO gets the value of RSCsegmentsforwardedviaLSO for the instance
func (instance *Win32_PerfFormattedData_TCPIPCounters_TCPIPPerformanceDiagnostics) GetPropertyRSCsegmentsforwardedviaLSO() (value uint32, err error) {
	retValue, err := instance.GetProperty("RSCsegmentsforwardedviaLSO")
	if err != nil {
		return
	}
	value, ok := retValue.(uint32)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetRSCsegmentsforwardedviasoftwaresegmentation sets the value of RSCsegmentsforwardedviasoftwaresegmentation for the instance
func (instance *Win32_PerfFormattedData_TCPIPCounters_TCPIPPerformanceDiagnostics) SetPropertyRSCsegmentsforwardedviasoftwaresegmentation(value uint32) (err error) {
	return instance.SetProperty("RSCsegmentsforwardedviasoftwaresegmentation", value)
}

// GetRSCsegmentsforwardedviasoftwaresegmentation gets the value of RSCsegmentsforwardedviasoftwaresegmentation for the instance
func (instance *Win32_PerfFormattedData_TCPIPCounters_TCPIPPerformanceDiagnostics) GetPropertyRSCsegmentsforwardedviasoftwaresegmentation() (value uint32, err error) {
	retValue, err := instance.GetProperty("RSCsegmentsforwardedviasoftwaresegmentation")
	if err != nil {
		return
	}
	value, ok := retValue.(uint32)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetRSCsegmentsforwardedviasoftwaresegmentationandchecksum sets the value of RSCsegmentsforwardedviasoftwaresegmentationandchecksum for the instance
func (instance *Win32_PerfFormattedData_TCPIPCounters_TCPIPPerformanceDiagnostics) SetPropertyRSCsegmentsforwardedviasoftwaresegmentationandchecksum(value uint32) (err error) {
	return instance.SetProperty("RSCsegmentsforwardedviasoftwaresegmentationandchecksum", value)
}

// GetRSCsegmentsforwardedviasoftwaresegmentationandchecksum gets the value of RSCsegmentsforwardedviasoftwaresegmentationandchecksum for the instance
func (instance *Win32_PerfFormattedData_TCPIPCounters_TCPIPPerformanceDiagnostics) GetPropertyRSCsegmentsforwardedviasoftwaresegmentationandchecksum() (value uint32, err error) {
	retValue, err := instance.GetProperty("RSCsegmentsforwardedviasoftwaresegmentationandchecksum")
	if err != nil {
		return
	}
	value, ok := retValue.(uint32)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetTCPchecksumerrors sets the value of TCPchecksumerrors for the instance
func (instance *Win32_PerfFormattedData_TCPIPCounters_TCPIPPerformanceDiagnostics) SetPropertyTCPchecksumerrors(value uint32) (err error) {
	return instance.SetProperty("TCPchecksumerrors", value)
}

// GetTCPchecksumerrors gets the value of TCPchecksumerrors for the instance
func (instance *Win32_PerfFormattedData_TCPIPCounters_TCPIPPerformanceDiagnostics) GetPropertyTCPchecksumerrors() (value uint32, err error) {
	retValue, err := instance.GetProperty("TCPchecksumerrors")
	if err != nil {
		return
	}
	value, ok := retValue.(uint32)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetTCPconnectrequestsfallenoffloopbackfastpath sets the value of TCPconnectrequestsfallenoffloopbackfastpath for the instance
func (instance *Win32_PerfFormattedData_TCPIPCounters_TCPIPPerformanceDiagnostics) SetPropertyTCPconnectrequestsfallenoffloopbackfastpath(value uint32) (err error) {
	return instance.SetProperty("TCPconnectrequestsfallenoffloopbackfastpath", value)
}

// GetTCPconnectrequestsfallenoffloopbackfastpath gets the value of TCPconnectrequestsfallenoffloopbackfastpath for the instance
func (instance *Win32_PerfFormattedData_TCPIPCounters_TCPIPPerformanceDiagnostics) GetPropertyTCPconnectrequestsfallenoffloopbackfastpath() (value uint32, err error) {
	retValue, err := instance.GetProperty("TCPconnectrequestsfallenoffloopbackfastpath")
	if err != nil {
		return
	}
	value, ok := retValue.(uint32)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetTCPconnectrequestsPersecfallenoffloopbackfastpath sets the value of TCPconnectrequestsPersecfallenoffloopbackfastpath for the instance
func (instance *Win32_PerfFormattedData_TCPIPCounters_TCPIPPerformanceDiagnostics) SetPropertyTCPconnectrequestsPersecfallenoffloopbackfastpath(value uint32) (err error) {
	return instance.SetProperty("TCPconnectrequestsPersecfallenoffloopbackfastpath", value)
}

// GetTCPconnectrequestsPersecfallenoffloopbackfastpath gets the value of TCPconnectrequestsPersecfallenoffloopbackfastpath for the instance
func (instance *Win32_PerfFormattedData_TCPIPCounters_TCPIPPerformanceDiagnostics) GetPropertyTCPconnectrequestsPersecfallenoffloopbackfastpath() (value uint32, err error) {
	retValue, err := instance.GetProperty("TCPconnectrequestsPersecfallenoffloopbackfastpath")
	if err != nil {
		return
	}
	value, ok := retValue.(uint32)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetTCPinboundsegmentsnotprocessedviafastpath sets the value of TCPinboundsegmentsnotprocessedviafastpath for the instance
func (instance *Win32_PerfFormattedData_TCPIPCounters_TCPIPPerformanceDiagnostics) SetPropertyTCPinboundsegmentsnotprocessedviafastpath(value uint32) (err error) {
	return instance.SetProperty("TCPinboundsegmentsnotprocessedviafastpath", value)
}

// GetTCPinboundsegmentsnotprocessedviafastpath gets the value of TCPinboundsegmentsnotprocessedviafastpath for the instance
func (instance *Win32_PerfFormattedData_TCPIPCounters_TCPIPPerformanceDiagnostics) GetPropertyTCPinboundsegmentsnotprocessedviafastpath() (value uint32, err error) {
	retValue, err := instance.GetProperty("TCPinboundsegmentsnotprocessedviafastpath")
	if err != nil {
		return
	}
	value, ok := retValue.(uint32)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetTCPinboundsegmentsPersecnotprocessedviafastpath sets the value of TCPinboundsegmentsPersecnotprocessedviafastpath for the instance
func (instance *Win32_PerfFormattedData_TCPIPCounters_TCPIPPerformanceDiagnostics) SetPropertyTCPinboundsegmentsPersecnotprocessedviafastpath(value uint32) (err error) {
	return instance.SetProperty("TCPinboundsegmentsPersecnotprocessedviafastpath", value)
}

// GetTCPinboundsegmentsPersecnotprocessedviafastpath gets the value of TCPinboundsegmentsPersecnotprocessedviafastpath for the instance
func (instance *Win32_PerfFormattedData_TCPIPCounters_TCPIPPerformanceDiagnostics) GetPropertyTCPinboundsegmentsPersecnotprocessedviafastpath() (value uint32, err error) {
	retValue, err := instance.GetProperty("TCPinboundsegmentsPersecnotprocessedviafastpath")
	if err != nil {
		return
	}
	value, ok := retValue.(uint32)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetTCPRSCbytesreceived sets the value of TCPRSCbytesreceived for the instance
func (instance *Win32_PerfFormattedData_TCPIPCounters_TCPIPPerformanceDiagnostics) SetPropertyTCPRSCbytesreceived(value uint32) (err error) {
	return instance.SetProperty("TCPRSCbytesreceived", value)
}

// GetTCPRSCbytesreceived gets the value of TCPRSCbytesreceived for the instance
func (instance *Win32_PerfFormattedData_TCPIPCounters_TCPIPPerformanceDiagnostics) GetPropertyTCPRSCbytesreceived() (value uint32, err error) {
	retValue, err := instance.GetProperty("TCPRSCbytesreceived")
	if err != nil {
		return
	}
	value, ok := retValue.(uint32)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetTCPRSCevents sets the value of TCPRSCevents for the instance
func (instance *Win32_PerfFormattedData_TCPIPCounters_TCPIPPerformanceDiagnostics) SetPropertyTCPRSCevents(value uint32) (err error) {
	return instance.SetProperty("TCPRSCevents", value)
}

// GetTCPRSCevents gets the value of TCPRSCevents for the instance
func (instance *Win32_PerfFormattedData_TCPIPCounters_TCPIPPerformanceDiagnostics) GetPropertyTCPRSCevents() (value uint32, err error) {
	retValue, err := instance.GetProperty("TCPRSCevents")
	if err != nil {
		return
	}
	value, ok := retValue.(uint32)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetTCPtimeouts sets the value of TCPtimeouts for the instance
func (instance *Win32_PerfFormattedData_TCPIPCounters_TCPIPPerformanceDiagnostics) SetPropertyTCPtimeouts(value uint32) (err error) {
	return instance.SetProperty("TCPtimeouts", value)
}

// GetTCPtimeouts gets the value of TCPtimeouts for the instance
func (instance *Win32_PerfFormattedData_TCPIPCounters_TCPIPPerformanceDiagnostics) GetPropertyTCPtimeouts() (value uint32, err error) {
	retValue, err := instance.GetProperty("TCPtimeouts")
	if err != nil {
		return
	}
	value, ok := retValue.(uint32)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetUDPdatagramscreatedviasoftwaresegmentation sets the value of UDPdatagramscreatedviasoftwaresegmentation for the instance
func (instance *Win32_PerfFormattedData_TCPIPCounters_TCPIPPerformanceDiagnostics) SetPropertyUDPdatagramscreatedviasoftwaresegmentation(value uint32) (err error) {
	return instance.SetProperty("UDPdatagramscreatedviasoftwaresegmentation", value)
}

// GetUDPdatagramscreatedviasoftwaresegmentation gets the value of UDPdatagramscreatedviasoftwaresegmentation for the instance
func (instance *Win32_PerfFormattedData_TCPIPCounters_TCPIPPerformanceDiagnostics) GetPropertyUDPdatagramscreatedviasoftwaresegmentation() (value uint32, err error) {
	retValue, err := instance.GetProperty("UDPdatagramscreatedviasoftwaresegmentation")
	if err != nil {
		return
	}
	value, ok := retValue.(uint32)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetUDPURObytesreceived sets the value of UDPURObytesreceived for the instance
func (instance *Win32_PerfFormattedData_TCPIPCounters_TCPIPPerformanceDiagnostics) SetPropertyUDPURObytesreceived(value uint32) (err error) {
	return instance.SetProperty("UDPURObytesreceived", value)
}

// GetUDPURObytesreceived gets the value of UDPURObytesreceived for the instance
func (instance *Win32_PerfFormattedData_TCPIPCounters_TCPIPPerformanceDiagnostics) GetPropertyUDPURObytesreceived() (value uint32, err error) {
	retValue, err := instance.GetProperty("UDPURObytesreceived")
	if err != nil {
		return
	}
	value, ok := retValue.(uint32)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetUDPUROevents sets the value of UDPUROevents for the instance
func (instance *Win32_PerfFormattedData_TCPIPCounters_TCPIPPerformanceDiagnostics) SetPropertyUDPUROevents(value uint32) (err error) {
	return instance.SetProperty("UDPUROevents", value)
}

// GetUDPUROevents gets the value of UDPUROevents for the instance
func (instance *Win32_PerfFormattedData_TCPIPCounters_TCPIPPerformanceDiagnostics) GetPropertyUDPUROevents() (value uint32, err error) {
	retValue, err := instance.GetProperty("UDPUROevents")
	if err != nil {
		return
	}
	value, ok := retValue.(uint32)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetUROsegmentforwardingfailuresduringsoftwaresegmentation sets the value of UROsegmentforwardingfailuresduringsoftwaresegmentation for the instance
func (instance *Win32_PerfFormattedData_TCPIPCounters_TCPIPPerformanceDiagnostics) SetPropertyUROsegmentforwardingfailuresduringsoftwaresegmentation(value uint32) (err error) {
	return instance.SetProperty("UROsegmentforwardingfailuresduringsoftwaresegmentation", value)
}

// GetUROsegmentforwardingfailuresduringsoftwaresegmentation gets the value of UROsegmentforwardingfailuresduringsoftwaresegmentation for the instance
func (instance *Win32_PerfFormattedData_TCPIPCounters_TCPIPPerformanceDiagnostics) GetPropertyUROsegmentforwardingfailuresduringsoftwaresegmentation() (value uint32, err error) {
	retValue, err := instance.GetProperty("UROsegmentforwardingfailuresduringsoftwaresegmentation")
	if err != nil {
		return
	}
	value, ok := retValue.(uint32)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetUROsegmentsforwardedviasoftwaresegmentation sets the value of UROsegmentsforwardedviasoftwaresegmentation for the instance
func (instance *Win32_PerfFormattedData_TCPIPCounters_TCPIPPerformanceDiagnostics) SetPropertyUROsegmentsforwardedviasoftwaresegmentation(value uint32) (err error) {
	return instance.SetProperty("UROsegmentsforwardedviasoftwaresegmentation", value)
}

// GetUROsegmentsforwardedviasoftwaresegmentation gets the value of UROsegmentsforwardedviasoftwaresegmentation for the instance
func (instance *Win32_PerfFormattedData_TCPIPCounters_TCPIPPerformanceDiagnostics) GetPropertyUROsegmentsforwardedviasoftwaresegmentation() (value uint32, err error) {
	retValue, err := instance.GetProperty("UROsegmentsforwardedviasoftwaresegmentation")
	if err != nil {
		return
	}
	value, ok := retValue.(uint32)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetUROsegmentsforwardedviasoftwaresegmentationandchecksum sets the value of UROsegmentsforwardedviasoftwaresegmentationandchecksum for the instance
func (instance *Win32_PerfFormattedData_TCPIPCounters_TCPIPPerformanceDiagnostics) SetPropertyUROsegmentsforwardedviasoftwaresegmentationandchecksum(value uint32) (err error) {
	return instance.SetProperty("UROsegmentsforwardedviasoftwaresegmentationandchecksum", value)
}

// GetUROsegmentsforwardedviasoftwaresegmentationandchecksum gets the value of UROsegmentsforwardedviasoftwaresegmentationandchecksum for the instance
func (instance *Win32_PerfFormattedData_TCPIPCounters_TCPIPPerformanceDiagnostics) GetPropertyUROsegmentsforwardedviasoftwaresegmentationandchecksum() (value uint32, err error) {
	retValue, err := instance.GetProperty("UROsegmentsforwardedviasoftwaresegmentationandchecksum")
	if err != nil {
		return
	}
	value, ok := retValue.(uint32)
	if !ok {
		// TODO: Set an error
	}
	return
}
