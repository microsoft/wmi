// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 6/12/2024 using wmigen
//      Source root.CIMV2
//////////////////////////////////////////////
package cimv2

import (
	"github.com/microsoft/wmi/pkg/base/query"
	"github.com/microsoft/wmi/pkg/errors"
	cim "github.com/microsoft/wmi/pkg/wmiinstance"
	"reflect"
)

// Win32_PerfRawData_TCPIPCounters_TCPIPPerformanceDiagnostics struct
type Win32_PerfRawData_TCPIPCounters_TCPIPPerformanceDiagnostics struct {
	*Win32_PerfRawData

	//
	Bytesoflostretransmitsretransmitted uint32

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
	NumberofSACKblocksdropped uint32

	//
	NumberofTCPRXfastpathbatchesinspected uint64

	//
	NumberofTCPRXfastpathbatchesnotinspected uint64

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
	TCPlossrecoveryepisodes uint32

	//
	TCPRSCbytesreceived uint32

	//
	TCPRSCevents uint32

	//
	TCPsuccessfullossrecoveryepisodes uint32

	//
	TCPtimeouts uint32

	//
	UDPdatagramscreatedviasoftwaresegmentation uint32

	//
	UDPURObytesreceived uint32

	//
	UDPUROevents uint32

	//
	UROsegmentationfailuresforrawsockets uint32

	//
	UROsegmentationsforrawsockets uint32

	//
	UROsegmentforwardingfailuresduringsoftwaresegmentation uint32

	//
	UROsegmentsforwardedviasoftwaresegmentation uint32

	//
	UROsegmentsforwardedviasoftwaresegmentationandchecksum uint32
}

func NewWin32_PerfRawData_TCPIPCounters_TCPIPPerformanceDiagnosticsEx1(instance *cim.WmiInstance) (newInstance *Win32_PerfRawData_TCPIPCounters_TCPIPPerformanceDiagnostics, err error) {
	tmp, err := NewWin32_PerfRawDataEx1(instance)

	if err != nil {
		return
	}
	newInstance = &Win32_PerfRawData_TCPIPCounters_TCPIPPerformanceDiagnostics{
		Win32_PerfRawData: tmp,
	}
	return
}

func NewWin32_PerfRawData_TCPIPCounters_TCPIPPerformanceDiagnosticsEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *Win32_PerfRawData_TCPIPCounters_TCPIPPerformanceDiagnostics, err error) {
	tmp, err := NewWin32_PerfRawDataEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &Win32_PerfRawData_TCPIPCounters_TCPIPPerformanceDiagnostics{
		Win32_PerfRawData: tmp,
	}
	return
}

// SetBytesoflostretransmitsretransmitted sets the value of Bytesoflostretransmitsretransmitted for the instance
func (instance *Win32_PerfRawData_TCPIPCounters_TCPIPPerformanceDiagnostics) SetPropertyBytesoflostretransmitsretransmitted(value uint32) (err error) {
	return instance.SetProperty("Bytesoflostretransmitsretransmitted", (value))
}

// GetBytesoflostretransmitsretransmitted gets the value of Bytesoflostretransmitsretransmitted for the instance
func (instance *Win32_PerfRawData_TCPIPCounters_TCPIPPerformanceDiagnostics) GetPropertyBytesoflostretransmitsretransmitted() (value uint32, err error) {
	retValue, err := instance.GetProperty("Bytesoflostretransmitsretransmitted")
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

// SetDeniedconnectorsendrequestsinlowpowermode sets the value of Deniedconnectorsendrequestsinlowpowermode for the instance
func (instance *Win32_PerfRawData_TCPIPCounters_TCPIPPerformanceDiagnostics) SetPropertyDeniedconnectorsendrequestsinlowpowermode(value uint32) (err error) {
	return instance.SetProperty("Deniedconnectorsendrequestsinlowpowermode", (value))
}

// GetDeniedconnectorsendrequestsinlowpowermode gets the value of Deniedconnectorsendrequestsinlowpowermode for the instance
func (instance *Win32_PerfRawData_TCPIPCounters_TCPIPPerformanceDiagnostics) GetPropertyDeniedconnectorsendrequestsinlowpowermode() (value uint32, err error) {
	retValue, err := instance.GetProperty("Deniedconnectorsendrequestsinlowpowermode")
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

// SetIPv4NBLsindicatedwithlowresourceflag sets the value of IPv4NBLsindicatedwithlowresourceflag for the instance
func (instance *Win32_PerfRawData_TCPIPCounters_TCPIPPerformanceDiagnostics) SetPropertyIPv4NBLsindicatedwithlowresourceflag(value uint32) (err error) {
	return instance.SetProperty("IPv4NBLsindicatedwithlowresourceflag", (value))
}

// GetIPv4NBLsindicatedwithlowresourceflag gets the value of IPv4NBLsindicatedwithlowresourceflag for the instance
func (instance *Win32_PerfRawData_TCPIPCounters_TCPIPPerformanceDiagnostics) GetPropertyIPv4NBLsindicatedwithlowresourceflag() (value uint32, err error) {
	retValue, err := instance.GetProperty("IPv4NBLsindicatedwithlowresourceflag")
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

// SetIPv4NBLsindicatedwithoutprevalidation sets the value of IPv4NBLsindicatedwithoutprevalidation for the instance
func (instance *Win32_PerfRawData_TCPIPCounters_TCPIPPerformanceDiagnostics) SetPropertyIPv4NBLsindicatedwithoutprevalidation(value uint32) (err error) {
	return instance.SetProperty("IPv4NBLsindicatedwithoutprevalidation", (value))
}

// GetIPv4NBLsindicatedwithoutprevalidation gets the value of IPv4NBLsindicatedwithoutprevalidation for the instance
func (instance *Win32_PerfRawData_TCPIPCounters_TCPIPPerformanceDiagnostics) GetPropertyIPv4NBLsindicatedwithoutprevalidation() (value uint32, err error) {
	retValue, err := instance.GetProperty("IPv4NBLsindicatedwithoutprevalidation")
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

// SetIPv4NBLsPersecindicatedwithlowresourceflag sets the value of IPv4NBLsPersecindicatedwithlowresourceflag for the instance
func (instance *Win32_PerfRawData_TCPIPCounters_TCPIPPerformanceDiagnostics) SetPropertyIPv4NBLsPersecindicatedwithlowresourceflag(value uint32) (err error) {
	return instance.SetProperty("IPv4NBLsPersecindicatedwithlowresourceflag", (value))
}

// GetIPv4NBLsPersecindicatedwithlowresourceflag gets the value of IPv4NBLsPersecindicatedwithlowresourceflag for the instance
func (instance *Win32_PerfRawData_TCPIPCounters_TCPIPPerformanceDiagnostics) GetPropertyIPv4NBLsPersecindicatedwithlowresourceflag() (value uint32, err error) {
	retValue, err := instance.GetProperty("IPv4NBLsPersecindicatedwithlowresourceflag")
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

// SetIPv4NBLsPersecindicatedwithoutprevalidation sets the value of IPv4NBLsPersecindicatedwithoutprevalidation for the instance
func (instance *Win32_PerfRawData_TCPIPCounters_TCPIPPerformanceDiagnostics) SetPropertyIPv4NBLsPersecindicatedwithoutprevalidation(value uint32) (err error) {
	return instance.SetProperty("IPv4NBLsPersecindicatedwithoutprevalidation", (value))
}

// GetIPv4NBLsPersecindicatedwithoutprevalidation gets the value of IPv4NBLsPersecindicatedwithoutprevalidation for the instance
func (instance *Win32_PerfRawData_TCPIPCounters_TCPIPPerformanceDiagnostics) GetPropertyIPv4NBLsPersecindicatedwithoutprevalidation() (value uint32, err error) {
	retValue, err := instance.GetProperty("IPv4NBLsPersecindicatedwithoutprevalidation")
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

// SetIPv4NBLsPersectreatedasnonprevalidated sets the value of IPv4NBLsPersectreatedasnonprevalidated for the instance
func (instance *Win32_PerfRawData_TCPIPCounters_TCPIPPerformanceDiagnostics) SetPropertyIPv4NBLsPersectreatedasnonprevalidated(value uint32) (err error) {
	return instance.SetProperty("IPv4NBLsPersectreatedasnonprevalidated", (value))
}

// GetIPv4NBLsPersectreatedasnonprevalidated gets the value of IPv4NBLsPersectreatedasnonprevalidated for the instance
func (instance *Win32_PerfRawData_TCPIPCounters_TCPIPPerformanceDiagnostics) GetPropertyIPv4NBLsPersectreatedasnonprevalidated() (value uint32, err error) {
	retValue, err := instance.GetProperty("IPv4NBLsPersectreatedasnonprevalidated")
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

// SetIPv4NBLstreatedasnonprevalidated sets the value of IPv4NBLstreatedasnonprevalidated for the instance
func (instance *Win32_PerfRawData_TCPIPCounters_TCPIPPerformanceDiagnostics) SetPropertyIPv4NBLstreatedasnonprevalidated(value uint32) (err error) {
	return instance.SetProperty("IPv4NBLstreatedasnonprevalidated", (value))
}

// GetIPv4NBLstreatedasnonprevalidated gets the value of IPv4NBLstreatedasnonprevalidated for the instance
func (instance *Win32_PerfRawData_TCPIPCounters_TCPIPPerformanceDiagnostics) GetPropertyIPv4NBLstreatedasnonprevalidated() (value uint32, err error) {
	retValue, err := instance.GetProperty("IPv4NBLstreatedasnonprevalidated")
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

// SetIPv4outboundNBLsnotprocessedviafastpath sets the value of IPv4outboundNBLsnotprocessedviafastpath for the instance
func (instance *Win32_PerfRawData_TCPIPCounters_TCPIPPerformanceDiagnostics) SetPropertyIPv4outboundNBLsnotprocessedviafastpath(value uint32) (err error) {
	return instance.SetProperty("IPv4outboundNBLsnotprocessedviafastpath", (value))
}

// GetIPv4outboundNBLsnotprocessedviafastpath gets the value of IPv4outboundNBLsnotprocessedviafastpath for the instance
func (instance *Win32_PerfRawData_TCPIPCounters_TCPIPPerformanceDiagnostics) GetPropertyIPv4outboundNBLsnotprocessedviafastpath() (value uint32, err error) {
	retValue, err := instance.GetProperty("IPv4outboundNBLsnotprocessedviafastpath")
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

// SetIPv4outboundNBLsPersecnotprocessedviafastpath sets the value of IPv4outboundNBLsPersecnotprocessedviafastpath for the instance
func (instance *Win32_PerfRawData_TCPIPCounters_TCPIPPerformanceDiagnostics) SetPropertyIPv4outboundNBLsPersecnotprocessedviafastpath(value uint32) (err error) {
	return instance.SetProperty("IPv4outboundNBLsPersecnotprocessedviafastpath", (value))
}

// GetIPv4outboundNBLsPersecnotprocessedviafastpath gets the value of IPv4outboundNBLsPersecnotprocessedviafastpath for the instance
func (instance *Win32_PerfRawData_TCPIPCounters_TCPIPPerformanceDiagnostics) GetPropertyIPv4outboundNBLsPersecnotprocessedviafastpath() (value uint32, err error) {
	retValue, err := instance.GetProperty("IPv4outboundNBLsPersecnotprocessedviafastpath")
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

// SetIPv6NBLsindicatedwithlowresourceflag sets the value of IPv6NBLsindicatedwithlowresourceflag for the instance
func (instance *Win32_PerfRawData_TCPIPCounters_TCPIPPerformanceDiagnostics) SetPropertyIPv6NBLsindicatedwithlowresourceflag(value uint32) (err error) {
	return instance.SetProperty("IPv6NBLsindicatedwithlowresourceflag", (value))
}

// GetIPv6NBLsindicatedwithlowresourceflag gets the value of IPv6NBLsindicatedwithlowresourceflag for the instance
func (instance *Win32_PerfRawData_TCPIPCounters_TCPIPPerformanceDiagnostics) GetPropertyIPv6NBLsindicatedwithlowresourceflag() (value uint32, err error) {
	retValue, err := instance.GetProperty("IPv6NBLsindicatedwithlowresourceflag")
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

// SetIPv6NBLsindicatedwithoutprevalidation sets the value of IPv6NBLsindicatedwithoutprevalidation for the instance
func (instance *Win32_PerfRawData_TCPIPCounters_TCPIPPerformanceDiagnostics) SetPropertyIPv6NBLsindicatedwithoutprevalidation(value uint32) (err error) {
	return instance.SetProperty("IPv6NBLsindicatedwithoutprevalidation", (value))
}

// GetIPv6NBLsindicatedwithoutprevalidation gets the value of IPv6NBLsindicatedwithoutprevalidation for the instance
func (instance *Win32_PerfRawData_TCPIPCounters_TCPIPPerformanceDiagnostics) GetPropertyIPv6NBLsindicatedwithoutprevalidation() (value uint32, err error) {
	retValue, err := instance.GetProperty("IPv6NBLsindicatedwithoutprevalidation")
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

// SetIPv6NBLsPersecindicatedwithlowresourceflag sets the value of IPv6NBLsPersecindicatedwithlowresourceflag for the instance
func (instance *Win32_PerfRawData_TCPIPCounters_TCPIPPerformanceDiagnostics) SetPropertyIPv6NBLsPersecindicatedwithlowresourceflag(value uint32) (err error) {
	return instance.SetProperty("IPv6NBLsPersecindicatedwithlowresourceflag", (value))
}

// GetIPv6NBLsPersecindicatedwithlowresourceflag gets the value of IPv6NBLsPersecindicatedwithlowresourceflag for the instance
func (instance *Win32_PerfRawData_TCPIPCounters_TCPIPPerformanceDiagnostics) GetPropertyIPv6NBLsPersecindicatedwithlowresourceflag() (value uint32, err error) {
	retValue, err := instance.GetProperty("IPv6NBLsPersecindicatedwithlowresourceflag")
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

// SetIPv6NBLsPersecindicatedwithoutprevalidation sets the value of IPv6NBLsPersecindicatedwithoutprevalidation for the instance
func (instance *Win32_PerfRawData_TCPIPCounters_TCPIPPerformanceDiagnostics) SetPropertyIPv6NBLsPersecindicatedwithoutprevalidation(value uint32) (err error) {
	return instance.SetProperty("IPv6NBLsPersecindicatedwithoutprevalidation", (value))
}

// GetIPv6NBLsPersecindicatedwithoutprevalidation gets the value of IPv6NBLsPersecindicatedwithoutprevalidation for the instance
func (instance *Win32_PerfRawData_TCPIPCounters_TCPIPPerformanceDiagnostics) GetPropertyIPv6NBLsPersecindicatedwithoutprevalidation() (value uint32, err error) {
	retValue, err := instance.GetProperty("IPv6NBLsPersecindicatedwithoutprevalidation")
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

// SetIPv6NBLsPersectreatedasnonprevalidated sets the value of IPv6NBLsPersectreatedasnonprevalidated for the instance
func (instance *Win32_PerfRawData_TCPIPCounters_TCPIPPerformanceDiagnostics) SetPropertyIPv6NBLsPersectreatedasnonprevalidated(value uint32) (err error) {
	return instance.SetProperty("IPv6NBLsPersectreatedasnonprevalidated", (value))
}

// GetIPv6NBLsPersectreatedasnonprevalidated gets the value of IPv6NBLsPersectreatedasnonprevalidated for the instance
func (instance *Win32_PerfRawData_TCPIPCounters_TCPIPPerformanceDiagnostics) GetPropertyIPv6NBLsPersectreatedasnonprevalidated() (value uint32, err error) {
	retValue, err := instance.GetProperty("IPv6NBLsPersectreatedasnonprevalidated")
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

// SetIPv6NBLstreatedasnonprevalidated sets the value of IPv6NBLstreatedasnonprevalidated for the instance
func (instance *Win32_PerfRawData_TCPIPCounters_TCPIPPerformanceDiagnostics) SetPropertyIPv6NBLstreatedasnonprevalidated(value uint32) (err error) {
	return instance.SetProperty("IPv6NBLstreatedasnonprevalidated", (value))
}

// GetIPv6NBLstreatedasnonprevalidated gets the value of IPv6NBLstreatedasnonprevalidated for the instance
func (instance *Win32_PerfRawData_TCPIPCounters_TCPIPPerformanceDiagnostics) GetPropertyIPv6NBLstreatedasnonprevalidated() (value uint32, err error) {
	retValue, err := instance.GetProperty("IPv6NBLstreatedasnonprevalidated")
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

// SetIPv6outboundNBLsnotprocessedviafastpath sets the value of IPv6outboundNBLsnotprocessedviafastpath for the instance
func (instance *Win32_PerfRawData_TCPIPCounters_TCPIPPerformanceDiagnostics) SetPropertyIPv6outboundNBLsnotprocessedviafastpath(value uint32) (err error) {
	return instance.SetProperty("IPv6outboundNBLsnotprocessedviafastpath", (value))
}

// GetIPv6outboundNBLsnotprocessedviafastpath gets the value of IPv6outboundNBLsnotprocessedviafastpath for the instance
func (instance *Win32_PerfRawData_TCPIPCounters_TCPIPPerformanceDiagnostics) GetPropertyIPv6outboundNBLsnotprocessedviafastpath() (value uint32, err error) {
	retValue, err := instance.GetProperty("IPv6outboundNBLsnotprocessedviafastpath")
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

// SetIPv6outboundNBLsPersecnotprocessedviafastpath sets the value of IPv6outboundNBLsPersecnotprocessedviafastpath for the instance
func (instance *Win32_PerfRawData_TCPIPCounters_TCPIPPerformanceDiagnostics) SetPropertyIPv6outboundNBLsPersecnotprocessedviafastpath(value uint32) (err error) {
	return instance.SetProperty("IPv6outboundNBLsPersecnotprocessedviafastpath", (value))
}

// GetIPv6outboundNBLsPersecnotprocessedviafastpath gets the value of IPv6outboundNBLsPersecnotprocessedviafastpath for the instance
func (instance *Win32_PerfRawData_TCPIPCounters_TCPIPPerformanceDiagnostics) GetPropertyIPv6outboundNBLsPersecnotprocessedviafastpath() (value uint32, err error) {
	retValue, err := instance.GetProperty("IPv6outboundNBLsPersecnotprocessedviafastpath")
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

// SetNumberofSACKblocksdropped sets the value of NumberofSACKblocksdropped for the instance
func (instance *Win32_PerfRawData_TCPIPCounters_TCPIPPerformanceDiagnostics) SetPropertyNumberofSACKblocksdropped(value uint32) (err error) {
	return instance.SetProperty("NumberofSACKblocksdropped", (value))
}

// GetNumberofSACKblocksdropped gets the value of NumberofSACKblocksdropped for the instance
func (instance *Win32_PerfRawData_TCPIPCounters_TCPIPPerformanceDiagnostics) GetPropertyNumberofSACKblocksdropped() (value uint32, err error) {
	retValue, err := instance.GetProperty("NumberofSACKblocksdropped")
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

// SetNumberofTCPRXfastpathbatchesinspected sets the value of NumberofTCPRXfastpathbatchesinspected for the instance
func (instance *Win32_PerfRawData_TCPIPCounters_TCPIPPerformanceDiagnostics) SetPropertyNumberofTCPRXfastpathbatchesinspected(value uint64) (err error) {
	return instance.SetProperty("NumberofTCPRXfastpathbatchesinspected", (value))
}

// GetNumberofTCPRXfastpathbatchesinspected gets the value of NumberofTCPRXfastpathbatchesinspected for the instance
func (instance *Win32_PerfRawData_TCPIPCounters_TCPIPPerformanceDiagnostics) GetPropertyNumberofTCPRXfastpathbatchesinspected() (value uint64, err error) {
	retValue, err := instance.GetProperty("NumberofTCPRXfastpathbatchesinspected")
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

// SetNumberofTCPRXfastpathbatchesnotinspected sets the value of NumberofTCPRXfastpathbatchesnotinspected for the instance
func (instance *Win32_PerfRawData_TCPIPCounters_TCPIPPerformanceDiagnostics) SetPropertyNumberofTCPRXfastpathbatchesnotinspected(value uint64) (err error) {
	return instance.SetProperty("NumberofTCPRXfastpathbatchesnotinspected", (value))
}

// GetNumberofTCPRXfastpathbatchesnotinspected gets the value of NumberofTCPRXfastpathbatchesnotinspected for the instance
func (instance *Win32_PerfRawData_TCPIPCounters_TCPIPPerformanceDiagnostics) GetPropertyNumberofTCPRXfastpathbatchesnotinspected() (value uint64, err error) {
	retValue, err := instance.GetProperty("NumberofTCPRXfastpathbatchesnotinspected")
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

// SetRSCsegmentforwardingfailuresduringsoftwaresegmentation sets the value of RSCsegmentforwardingfailuresduringsoftwaresegmentation for the instance
func (instance *Win32_PerfRawData_TCPIPCounters_TCPIPPerformanceDiagnostics) SetPropertyRSCsegmentforwardingfailuresduringsoftwaresegmentation(value uint32) (err error) {
	return instance.SetProperty("RSCsegmentforwardingfailuresduringsoftwaresegmentation", (value))
}

// GetRSCsegmentforwardingfailuresduringsoftwaresegmentation gets the value of RSCsegmentforwardingfailuresduringsoftwaresegmentation for the instance
func (instance *Win32_PerfRawData_TCPIPCounters_TCPIPPerformanceDiagnostics) GetPropertyRSCsegmentforwardingfailuresduringsoftwaresegmentation() (value uint32, err error) {
	retValue, err := instance.GetProperty("RSCsegmentforwardingfailuresduringsoftwaresegmentation")
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

// SetRSCsegmentsforwardedviaLSO sets the value of RSCsegmentsforwardedviaLSO for the instance
func (instance *Win32_PerfRawData_TCPIPCounters_TCPIPPerformanceDiagnostics) SetPropertyRSCsegmentsforwardedviaLSO(value uint32) (err error) {
	return instance.SetProperty("RSCsegmentsforwardedviaLSO", (value))
}

// GetRSCsegmentsforwardedviaLSO gets the value of RSCsegmentsforwardedviaLSO for the instance
func (instance *Win32_PerfRawData_TCPIPCounters_TCPIPPerformanceDiagnostics) GetPropertyRSCsegmentsforwardedviaLSO() (value uint32, err error) {
	retValue, err := instance.GetProperty("RSCsegmentsforwardedviaLSO")
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

// SetRSCsegmentsforwardedviasoftwaresegmentation sets the value of RSCsegmentsforwardedviasoftwaresegmentation for the instance
func (instance *Win32_PerfRawData_TCPIPCounters_TCPIPPerformanceDiagnostics) SetPropertyRSCsegmentsforwardedviasoftwaresegmentation(value uint32) (err error) {
	return instance.SetProperty("RSCsegmentsforwardedviasoftwaresegmentation", (value))
}

// GetRSCsegmentsforwardedviasoftwaresegmentation gets the value of RSCsegmentsforwardedviasoftwaresegmentation for the instance
func (instance *Win32_PerfRawData_TCPIPCounters_TCPIPPerformanceDiagnostics) GetPropertyRSCsegmentsforwardedviasoftwaresegmentation() (value uint32, err error) {
	retValue, err := instance.GetProperty("RSCsegmentsforwardedviasoftwaresegmentation")
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

// SetRSCsegmentsforwardedviasoftwaresegmentationandchecksum sets the value of RSCsegmentsforwardedviasoftwaresegmentationandchecksum for the instance
func (instance *Win32_PerfRawData_TCPIPCounters_TCPIPPerformanceDiagnostics) SetPropertyRSCsegmentsforwardedviasoftwaresegmentationandchecksum(value uint32) (err error) {
	return instance.SetProperty("RSCsegmentsforwardedviasoftwaresegmentationandchecksum", (value))
}

// GetRSCsegmentsforwardedviasoftwaresegmentationandchecksum gets the value of RSCsegmentsforwardedviasoftwaresegmentationandchecksum for the instance
func (instance *Win32_PerfRawData_TCPIPCounters_TCPIPPerformanceDiagnostics) GetPropertyRSCsegmentsforwardedviasoftwaresegmentationandchecksum() (value uint32, err error) {
	retValue, err := instance.GetProperty("RSCsegmentsforwardedviasoftwaresegmentationandchecksum")
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

// SetTCPchecksumerrors sets the value of TCPchecksumerrors for the instance
func (instance *Win32_PerfRawData_TCPIPCounters_TCPIPPerformanceDiagnostics) SetPropertyTCPchecksumerrors(value uint32) (err error) {
	return instance.SetProperty("TCPchecksumerrors", (value))
}

// GetTCPchecksumerrors gets the value of TCPchecksumerrors for the instance
func (instance *Win32_PerfRawData_TCPIPCounters_TCPIPPerformanceDiagnostics) GetPropertyTCPchecksumerrors() (value uint32, err error) {
	retValue, err := instance.GetProperty("TCPchecksumerrors")
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

// SetTCPconnectrequestsfallenoffloopbackfastpath sets the value of TCPconnectrequestsfallenoffloopbackfastpath for the instance
func (instance *Win32_PerfRawData_TCPIPCounters_TCPIPPerformanceDiagnostics) SetPropertyTCPconnectrequestsfallenoffloopbackfastpath(value uint32) (err error) {
	return instance.SetProperty("TCPconnectrequestsfallenoffloopbackfastpath", (value))
}

// GetTCPconnectrequestsfallenoffloopbackfastpath gets the value of TCPconnectrequestsfallenoffloopbackfastpath for the instance
func (instance *Win32_PerfRawData_TCPIPCounters_TCPIPPerformanceDiagnostics) GetPropertyTCPconnectrequestsfallenoffloopbackfastpath() (value uint32, err error) {
	retValue, err := instance.GetProperty("TCPconnectrequestsfallenoffloopbackfastpath")
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

// SetTCPconnectrequestsPersecfallenoffloopbackfastpath sets the value of TCPconnectrequestsPersecfallenoffloopbackfastpath for the instance
func (instance *Win32_PerfRawData_TCPIPCounters_TCPIPPerformanceDiagnostics) SetPropertyTCPconnectrequestsPersecfallenoffloopbackfastpath(value uint32) (err error) {
	return instance.SetProperty("TCPconnectrequestsPersecfallenoffloopbackfastpath", (value))
}

// GetTCPconnectrequestsPersecfallenoffloopbackfastpath gets the value of TCPconnectrequestsPersecfallenoffloopbackfastpath for the instance
func (instance *Win32_PerfRawData_TCPIPCounters_TCPIPPerformanceDiagnostics) GetPropertyTCPconnectrequestsPersecfallenoffloopbackfastpath() (value uint32, err error) {
	retValue, err := instance.GetProperty("TCPconnectrequestsPersecfallenoffloopbackfastpath")
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

// SetTCPinboundsegmentsnotprocessedviafastpath sets the value of TCPinboundsegmentsnotprocessedviafastpath for the instance
func (instance *Win32_PerfRawData_TCPIPCounters_TCPIPPerformanceDiagnostics) SetPropertyTCPinboundsegmentsnotprocessedviafastpath(value uint32) (err error) {
	return instance.SetProperty("TCPinboundsegmentsnotprocessedviafastpath", (value))
}

// GetTCPinboundsegmentsnotprocessedviafastpath gets the value of TCPinboundsegmentsnotprocessedviafastpath for the instance
func (instance *Win32_PerfRawData_TCPIPCounters_TCPIPPerformanceDiagnostics) GetPropertyTCPinboundsegmentsnotprocessedviafastpath() (value uint32, err error) {
	retValue, err := instance.GetProperty("TCPinboundsegmentsnotprocessedviafastpath")
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

// SetTCPinboundsegmentsPersecnotprocessedviafastpath sets the value of TCPinboundsegmentsPersecnotprocessedviafastpath for the instance
func (instance *Win32_PerfRawData_TCPIPCounters_TCPIPPerformanceDiagnostics) SetPropertyTCPinboundsegmentsPersecnotprocessedviafastpath(value uint32) (err error) {
	return instance.SetProperty("TCPinboundsegmentsPersecnotprocessedviafastpath", (value))
}

// GetTCPinboundsegmentsPersecnotprocessedviafastpath gets the value of TCPinboundsegmentsPersecnotprocessedviafastpath for the instance
func (instance *Win32_PerfRawData_TCPIPCounters_TCPIPPerformanceDiagnostics) GetPropertyTCPinboundsegmentsPersecnotprocessedviafastpath() (value uint32, err error) {
	retValue, err := instance.GetProperty("TCPinboundsegmentsPersecnotprocessedviafastpath")
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

// SetTCPlossrecoveryepisodes sets the value of TCPlossrecoveryepisodes for the instance
func (instance *Win32_PerfRawData_TCPIPCounters_TCPIPPerformanceDiagnostics) SetPropertyTCPlossrecoveryepisodes(value uint32) (err error) {
	return instance.SetProperty("TCPlossrecoveryepisodes", (value))
}

// GetTCPlossrecoveryepisodes gets the value of TCPlossrecoveryepisodes for the instance
func (instance *Win32_PerfRawData_TCPIPCounters_TCPIPPerformanceDiagnostics) GetPropertyTCPlossrecoveryepisodes() (value uint32, err error) {
	retValue, err := instance.GetProperty("TCPlossrecoveryepisodes")
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

// SetTCPRSCbytesreceived sets the value of TCPRSCbytesreceived for the instance
func (instance *Win32_PerfRawData_TCPIPCounters_TCPIPPerformanceDiagnostics) SetPropertyTCPRSCbytesreceived(value uint32) (err error) {
	return instance.SetProperty("TCPRSCbytesreceived", (value))
}

// GetTCPRSCbytesreceived gets the value of TCPRSCbytesreceived for the instance
func (instance *Win32_PerfRawData_TCPIPCounters_TCPIPPerformanceDiagnostics) GetPropertyTCPRSCbytesreceived() (value uint32, err error) {
	retValue, err := instance.GetProperty("TCPRSCbytesreceived")
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

// SetTCPRSCevents sets the value of TCPRSCevents for the instance
func (instance *Win32_PerfRawData_TCPIPCounters_TCPIPPerformanceDiagnostics) SetPropertyTCPRSCevents(value uint32) (err error) {
	return instance.SetProperty("TCPRSCevents", (value))
}

// GetTCPRSCevents gets the value of TCPRSCevents for the instance
func (instance *Win32_PerfRawData_TCPIPCounters_TCPIPPerformanceDiagnostics) GetPropertyTCPRSCevents() (value uint32, err error) {
	retValue, err := instance.GetProperty("TCPRSCevents")
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

// SetTCPsuccessfullossrecoveryepisodes sets the value of TCPsuccessfullossrecoveryepisodes for the instance
func (instance *Win32_PerfRawData_TCPIPCounters_TCPIPPerformanceDiagnostics) SetPropertyTCPsuccessfullossrecoveryepisodes(value uint32) (err error) {
	return instance.SetProperty("TCPsuccessfullossrecoveryepisodes", (value))
}

// GetTCPsuccessfullossrecoveryepisodes gets the value of TCPsuccessfullossrecoveryepisodes for the instance
func (instance *Win32_PerfRawData_TCPIPCounters_TCPIPPerformanceDiagnostics) GetPropertyTCPsuccessfullossrecoveryepisodes() (value uint32, err error) {
	retValue, err := instance.GetProperty("TCPsuccessfullossrecoveryepisodes")
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

// SetTCPtimeouts sets the value of TCPtimeouts for the instance
func (instance *Win32_PerfRawData_TCPIPCounters_TCPIPPerformanceDiagnostics) SetPropertyTCPtimeouts(value uint32) (err error) {
	return instance.SetProperty("TCPtimeouts", (value))
}

// GetTCPtimeouts gets the value of TCPtimeouts for the instance
func (instance *Win32_PerfRawData_TCPIPCounters_TCPIPPerformanceDiagnostics) GetPropertyTCPtimeouts() (value uint32, err error) {
	retValue, err := instance.GetProperty("TCPtimeouts")
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

// SetUDPdatagramscreatedviasoftwaresegmentation sets the value of UDPdatagramscreatedviasoftwaresegmentation for the instance
func (instance *Win32_PerfRawData_TCPIPCounters_TCPIPPerformanceDiagnostics) SetPropertyUDPdatagramscreatedviasoftwaresegmentation(value uint32) (err error) {
	return instance.SetProperty("UDPdatagramscreatedviasoftwaresegmentation", (value))
}

// GetUDPdatagramscreatedviasoftwaresegmentation gets the value of UDPdatagramscreatedviasoftwaresegmentation for the instance
func (instance *Win32_PerfRawData_TCPIPCounters_TCPIPPerformanceDiagnostics) GetPropertyUDPdatagramscreatedviasoftwaresegmentation() (value uint32, err error) {
	retValue, err := instance.GetProperty("UDPdatagramscreatedviasoftwaresegmentation")
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

// SetUDPURObytesreceived sets the value of UDPURObytesreceived for the instance
func (instance *Win32_PerfRawData_TCPIPCounters_TCPIPPerformanceDiagnostics) SetPropertyUDPURObytesreceived(value uint32) (err error) {
	return instance.SetProperty("UDPURObytesreceived", (value))
}

// GetUDPURObytesreceived gets the value of UDPURObytesreceived for the instance
func (instance *Win32_PerfRawData_TCPIPCounters_TCPIPPerformanceDiagnostics) GetPropertyUDPURObytesreceived() (value uint32, err error) {
	retValue, err := instance.GetProperty("UDPURObytesreceived")
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

// SetUDPUROevents sets the value of UDPUROevents for the instance
func (instance *Win32_PerfRawData_TCPIPCounters_TCPIPPerformanceDiagnostics) SetPropertyUDPUROevents(value uint32) (err error) {
	return instance.SetProperty("UDPUROevents", (value))
}

// GetUDPUROevents gets the value of UDPUROevents for the instance
func (instance *Win32_PerfRawData_TCPIPCounters_TCPIPPerformanceDiagnostics) GetPropertyUDPUROevents() (value uint32, err error) {
	retValue, err := instance.GetProperty("UDPUROevents")
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

// SetUROsegmentationfailuresforrawsockets sets the value of UROsegmentationfailuresforrawsockets for the instance
func (instance *Win32_PerfRawData_TCPIPCounters_TCPIPPerformanceDiagnostics) SetPropertyUROsegmentationfailuresforrawsockets(value uint32) (err error) {
	return instance.SetProperty("UROsegmentationfailuresforrawsockets", (value))
}

// GetUROsegmentationfailuresforrawsockets gets the value of UROsegmentationfailuresforrawsockets for the instance
func (instance *Win32_PerfRawData_TCPIPCounters_TCPIPPerformanceDiagnostics) GetPropertyUROsegmentationfailuresforrawsockets() (value uint32, err error) {
	retValue, err := instance.GetProperty("UROsegmentationfailuresforrawsockets")
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

// SetUROsegmentationsforrawsockets sets the value of UROsegmentationsforrawsockets for the instance
func (instance *Win32_PerfRawData_TCPIPCounters_TCPIPPerformanceDiagnostics) SetPropertyUROsegmentationsforrawsockets(value uint32) (err error) {
	return instance.SetProperty("UROsegmentationsforrawsockets", (value))
}

// GetUROsegmentationsforrawsockets gets the value of UROsegmentationsforrawsockets for the instance
func (instance *Win32_PerfRawData_TCPIPCounters_TCPIPPerformanceDiagnostics) GetPropertyUROsegmentationsforrawsockets() (value uint32, err error) {
	retValue, err := instance.GetProperty("UROsegmentationsforrawsockets")
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

// SetUROsegmentforwardingfailuresduringsoftwaresegmentation sets the value of UROsegmentforwardingfailuresduringsoftwaresegmentation for the instance
func (instance *Win32_PerfRawData_TCPIPCounters_TCPIPPerformanceDiagnostics) SetPropertyUROsegmentforwardingfailuresduringsoftwaresegmentation(value uint32) (err error) {
	return instance.SetProperty("UROsegmentforwardingfailuresduringsoftwaresegmentation", (value))
}

// GetUROsegmentforwardingfailuresduringsoftwaresegmentation gets the value of UROsegmentforwardingfailuresduringsoftwaresegmentation for the instance
func (instance *Win32_PerfRawData_TCPIPCounters_TCPIPPerformanceDiagnostics) GetPropertyUROsegmentforwardingfailuresduringsoftwaresegmentation() (value uint32, err error) {
	retValue, err := instance.GetProperty("UROsegmentforwardingfailuresduringsoftwaresegmentation")
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

// SetUROsegmentsforwardedviasoftwaresegmentation sets the value of UROsegmentsforwardedviasoftwaresegmentation for the instance
func (instance *Win32_PerfRawData_TCPIPCounters_TCPIPPerformanceDiagnostics) SetPropertyUROsegmentsforwardedviasoftwaresegmentation(value uint32) (err error) {
	return instance.SetProperty("UROsegmentsforwardedviasoftwaresegmentation", (value))
}

// GetUROsegmentsforwardedviasoftwaresegmentation gets the value of UROsegmentsforwardedviasoftwaresegmentation for the instance
func (instance *Win32_PerfRawData_TCPIPCounters_TCPIPPerformanceDiagnostics) GetPropertyUROsegmentsforwardedviasoftwaresegmentation() (value uint32, err error) {
	retValue, err := instance.GetProperty("UROsegmentsforwardedviasoftwaresegmentation")
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

// SetUROsegmentsforwardedviasoftwaresegmentationandchecksum sets the value of UROsegmentsforwardedviasoftwaresegmentationandchecksum for the instance
func (instance *Win32_PerfRawData_TCPIPCounters_TCPIPPerformanceDiagnostics) SetPropertyUROsegmentsforwardedviasoftwaresegmentationandchecksum(value uint32) (err error) {
	return instance.SetProperty("UROsegmentsforwardedviasoftwaresegmentationandchecksum", (value))
}

// GetUROsegmentsforwardedviasoftwaresegmentationandchecksum gets the value of UROsegmentsforwardedviasoftwaresegmentationandchecksum for the instance
func (instance *Win32_PerfRawData_TCPIPCounters_TCPIPPerformanceDiagnostics) GetPropertyUROsegmentsforwardedviasoftwaresegmentationandchecksum() (value uint32, err error) {
	retValue, err := instance.GetProperty("UROsegmentsforwardedviasoftwaresegmentationandchecksum")
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
