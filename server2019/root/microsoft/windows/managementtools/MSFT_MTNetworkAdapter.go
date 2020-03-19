// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 3/19/2020 using wmigen
//      Source root.Microsoft.Windows.ManagementTools
//////////////////////////////////////////////
package managementtools

import (
	"github.com/microsoft/wmi/pkg/base/query"
	cim "github.com/microsoft/wmi/pkg/wmiinstance"
)

// MSFT_MTNetworkAdapter struct
type MSFT_MTNetworkAdapter struct {
	*CIM_ManagedElement

	//
	BytesReceived uint64

	//
	BytesReceivedPerInterval uint64

	//
	BytesReceivedThroughput float32

	//
	BytesSent uint64

	//
	BytesSentPerInterval uint64

	//
	BytesSentThroughput float32

	//
	BytesTotal uint64

	//
	BytesTotalPerInterval uint64

	//
	BytesTotalThroughput float32

	//
	CurrentIndex uint16

	//
	DNSName string

	//
	InterfaceDescription string

	//
	InterfaceGuid string

	//
	IntervalSeconds uint16

	//
	IPv4Address string

	//
	IPv6Address string

	//
	LinkSpeed uint64

	//
	MachineJoinedName string

	//
	MachineJoinedType uint16

	//
	MaxReceivedBitsPerSecond []float32

	//
	MaxSentBitsPerSecond []float32

	//
	Name string

	//
	NdisMedium uint32

	//
	NdisPhysicalMedium uint32

	//
	NonUniCastsReceived uint64

	//
	NonUniCastsReceivedPerInterval uint64

	//
	NonUniCastsSent uint64

	//
	NonUniCastsSentPerInterval uint64

	//
	NonUniCastsTotal uint64

	//
	NonUniCastsTotalPerInterval uint64

	//
	OperationStatus uint16

	//
	ReceivedBitsPerSecond []float32

	//
	ReceivedThroughput []float32

	//
	SentBitsPerSecond []float32

	//
	SentThroughput []float32

	//
	UniCastsReceived uint64

	//
	UniCastsReceivedPerInterval uint64

	//
	UniCastsSent uint64

	//
	UniCastsSentPerInterval uint64

	//
	UniCastsTotal uint64

	//
	UniCastsTotalPerInterval uint64

	//
	Utilization float32
}

func NewMSFT_MTNetworkAdapterEx1(instance *cim.WmiInstance) (newInstance *MSFT_MTNetworkAdapter, err error) {
	tmp, err := NewCIM_ManagedElementEx1(instance)

	if err != nil {
		return
	}
	newInstance = &MSFT_MTNetworkAdapter{
		CIM_ManagedElement: tmp,
	}
	return
}

func NewMSFT_MTNetworkAdapterEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *MSFT_MTNetworkAdapter, err error) {
	tmp, err := NewCIM_ManagedElementEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &MSFT_MTNetworkAdapter{
		CIM_ManagedElement: tmp,
	}
	return
}

// SetBytesReceived sets the value of BytesReceived for the instance
func (instance *MSFT_MTNetworkAdapter) SetPropertyBytesReceived(value uint64) (err error) {
	return instance.SetProperty("BytesReceived", value)
}

// GetBytesReceived gets the value of BytesReceived for the instance
func (instance *MSFT_MTNetworkAdapter) GetPropertyBytesReceived() (value uint64, err error) {
	retValue, err := instance.GetProperty("BytesReceived")
	if err != nil {
		return
	}
	value, ok := retValue.(uint64)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetBytesReceivedPerInterval sets the value of BytesReceivedPerInterval for the instance
func (instance *MSFT_MTNetworkAdapter) SetPropertyBytesReceivedPerInterval(value uint64) (err error) {
	return instance.SetProperty("BytesReceivedPerInterval", value)
}

// GetBytesReceivedPerInterval gets the value of BytesReceivedPerInterval for the instance
func (instance *MSFT_MTNetworkAdapter) GetPropertyBytesReceivedPerInterval() (value uint64, err error) {
	retValue, err := instance.GetProperty("BytesReceivedPerInterval")
	if err != nil {
		return
	}
	value, ok := retValue.(uint64)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetBytesReceivedThroughput sets the value of BytesReceivedThroughput for the instance
func (instance *MSFT_MTNetworkAdapter) SetPropertyBytesReceivedThroughput(value float32) (err error) {
	return instance.SetProperty("BytesReceivedThroughput", value)
}

// GetBytesReceivedThroughput gets the value of BytesReceivedThroughput for the instance
func (instance *MSFT_MTNetworkAdapter) GetPropertyBytesReceivedThroughput() (value float32, err error) {
	retValue, err := instance.GetProperty("BytesReceivedThroughput")
	if err != nil {
		return
	}
	value, ok := retValue.(float32)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetBytesSent sets the value of BytesSent for the instance
func (instance *MSFT_MTNetworkAdapter) SetPropertyBytesSent(value uint64) (err error) {
	return instance.SetProperty("BytesSent", value)
}

// GetBytesSent gets the value of BytesSent for the instance
func (instance *MSFT_MTNetworkAdapter) GetPropertyBytesSent() (value uint64, err error) {
	retValue, err := instance.GetProperty("BytesSent")
	if err != nil {
		return
	}
	value, ok := retValue.(uint64)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetBytesSentPerInterval sets the value of BytesSentPerInterval for the instance
func (instance *MSFT_MTNetworkAdapter) SetPropertyBytesSentPerInterval(value uint64) (err error) {
	return instance.SetProperty("BytesSentPerInterval", value)
}

// GetBytesSentPerInterval gets the value of BytesSentPerInterval for the instance
func (instance *MSFT_MTNetworkAdapter) GetPropertyBytesSentPerInterval() (value uint64, err error) {
	retValue, err := instance.GetProperty("BytesSentPerInterval")
	if err != nil {
		return
	}
	value, ok := retValue.(uint64)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetBytesSentThroughput sets the value of BytesSentThroughput for the instance
func (instance *MSFT_MTNetworkAdapter) SetPropertyBytesSentThroughput(value float32) (err error) {
	return instance.SetProperty("BytesSentThroughput", value)
}

// GetBytesSentThroughput gets the value of BytesSentThroughput for the instance
func (instance *MSFT_MTNetworkAdapter) GetPropertyBytesSentThroughput() (value float32, err error) {
	retValue, err := instance.GetProperty("BytesSentThroughput")
	if err != nil {
		return
	}
	value, ok := retValue.(float32)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetBytesTotal sets the value of BytesTotal for the instance
func (instance *MSFT_MTNetworkAdapter) SetPropertyBytesTotal(value uint64) (err error) {
	return instance.SetProperty("BytesTotal", value)
}

// GetBytesTotal gets the value of BytesTotal for the instance
func (instance *MSFT_MTNetworkAdapter) GetPropertyBytesTotal() (value uint64, err error) {
	retValue, err := instance.GetProperty("BytesTotal")
	if err != nil {
		return
	}
	value, ok := retValue.(uint64)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetBytesTotalPerInterval sets the value of BytesTotalPerInterval for the instance
func (instance *MSFT_MTNetworkAdapter) SetPropertyBytesTotalPerInterval(value uint64) (err error) {
	return instance.SetProperty("BytesTotalPerInterval", value)
}

// GetBytesTotalPerInterval gets the value of BytesTotalPerInterval for the instance
func (instance *MSFT_MTNetworkAdapter) GetPropertyBytesTotalPerInterval() (value uint64, err error) {
	retValue, err := instance.GetProperty("BytesTotalPerInterval")
	if err != nil {
		return
	}
	value, ok := retValue.(uint64)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetBytesTotalThroughput sets the value of BytesTotalThroughput for the instance
func (instance *MSFT_MTNetworkAdapter) SetPropertyBytesTotalThroughput(value float32) (err error) {
	return instance.SetProperty("BytesTotalThroughput", value)
}

// GetBytesTotalThroughput gets the value of BytesTotalThroughput for the instance
func (instance *MSFT_MTNetworkAdapter) GetPropertyBytesTotalThroughput() (value float32, err error) {
	retValue, err := instance.GetProperty("BytesTotalThroughput")
	if err != nil {
		return
	}
	value, ok := retValue.(float32)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetCurrentIndex sets the value of CurrentIndex for the instance
func (instance *MSFT_MTNetworkAdapter) SetPropertyCurrentIndex(value uint16) (err error) {
	return instance.SetProperty("CurrentIndex", value)
}

// GetCurrentIndex gets the value of CurrentIndex for the instance
func (instance *MSFT_MTNetworkAdapter) GetPropertyCurrentIndex() (value uint16, err error) {
	retValue, err := instance.GetProperty("CurrentIndex")
	if err != nil {
		return
	}
	value, ok := retValue.(uint16)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetDNSName sets the value of DNSName for the instance
func (instance *MSFT_MTNetworkAdapter) SetPropertyDNSName(value string) (err error) {
	return instance.SetProperty("DNSName", value)
}

// GetDNSName gets the value of DNSName for the instance
func (instance *MSFT_MTNetworkAdapter) GetPropertyDNSName() (value string, err error) {
	retValue, err := instance.GetProperty("DNSName")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetInterfaceDescription sets the value of InterfaceDescription for the instance
func (instance *MSFT_MTNetworkAdapter) SetPropertyInterfaceDescription(value string) (err error) {
	return instance.SetProperty("InterfaceDescription", value)
}

// GetInterfaceDescription gets the value of InterfaceDescription for the instance
func (instance *MSFT_MTNetworkAdapter) GetPropertyInterfaceDescription() (value string, err error) {
	retValue, err := instance.GetProperty("InterfaceDescription")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetInterfaceGuid sets the value of InterfaceGuid for the instance
func (instance *MSFT_MTNetworkAdapter) SetPropertyInterfaceGuid(value string) (err error) {
	return instance.SetProperty("InterfaceGuid", value)
}

// GetInterfaceGuid gets the value of InterfaceGuid for the instance
func (instance *MSFT_MTNetworkAdapter) GetPropertyInterfaceGuid() (value string, err error) {
	retValue, err := instance.GetProperty("InterfaceGuid")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetIntervalSeconds sets the value of IntervalSeconds for the instance
func (instance *MSFT_MTNetworkAdapter) SetPropertyIntervalSeconds(value uint16) (err error) {
	return instance.SetProperty("IntervalSeconds", value)
}

// GetIntervalSeconds gets the value of IntervalSeconds for the instance
func (instance *MSFT_MTNetworkAdapter) GetPropertyIntervalSeconds() (value uint16, err error) {
	retValue, err := instance.GetProperty("IntervalSeconds")
	if err != nil {
		return
	}
	value, ok := retValue.(uint16)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetIPv4Address sets the value of IPv4Address for the instance
func (instance *MSFT_MTNetworkAdapter) SetPropertyIPv4Address(value string) (err error) {
	return instance.SetProperty("IPv4Address", value)
}

// GetIPv4Address gets the value of IPv4Address for the instance
func (instance *MSFT_MTNetworkAdapter) GetPropertyIPv4Address() (value string, err error) {
	retValue, err := instance.GetProperty("IPv4Address")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetIPv6Address sets the value of IPv6Address for the instance
func (instance *MSFT_MTNetworkAdapter) SetPropertyIPv6Address(value string) (err error) {
	return instance.SetProperty("IPv6Address", value)
}

// GetIPv6Address gets the value of IPv6Address for the instance
func (instance *MSFT_MTNetworkAdapter) GetPropertyIPv6Address() (value string, err error) {
	retValue, err := instance.GetProperty("IPv6Address")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetLinkSpeed sets the value of LinkSpeed for the instance
func (instance *MSFT_MTNetworkAdapter) SetPropertyLinkSpeed(value uint64) (err error) {
	return instance.SetProperty("LinkSpeed", value)
}

// GetLinkSpeed gets the value of LinkSpeed for the instance
func (instance *MSFT_MTNetworkAdapter) GetPropertyLinkSpeed() (value uint64, err error) {
	retValue, err := instance.GetProperty("LinkSpeed")
	if err != nil {
		return
	}
	value, ok := retValue.(uint64)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetMachineJoinedName sets the value of MachineJoinedName for the instance
func (instance *MSFT_MTNetworkAdapter) SetPropertyMachineJoinedName(value string) (err error) {
	return instance.SetProperty("MachineJoinedName", value)
}

// GetMachineJoinedName gets the value of MachineJoinedName for the instance
func (instance *MSFT_MTNetworkAdapter) GetPropertyMachineJoinedName() (value string, err error) {
	retValue, err := instance.GetProperty("MachineJoinedName")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetMachineJoinedType sets the value of MachineJoinedType for the instance
func (instance *MSFT_MTNetworkAdapter) SetPropertyMachineJoinedType(value uint16) (err error) {
	return instance.SetProperty("MachineJoinedType", value)
}

// GetMachineJoinedType gets the value of MachineJoinedType for the instance
func (instance *MSFT_MTNetworkAdapter) GetPropertyMachineJoinedType() (value uint16, err error) {
	retValue, err := instance.GetProperty("MachineJoinedType")
	if err != nil {
		return
	}
	value, ok := retValue.(uint16)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetMaxReceivedBitsPerSecond sets the value of MaxReceivedBitsPerSecond for the instance
func (instance *MSFT_MTNetworkAdapter) SetPropertyMaxReceivedBitsPerSecond(value []float32) (err error) {
	return instance.SetProperty("MaxReceivedBitsPerSecond", value)
}

// GetMaxReceivedBitsPerSecond gets the value of MaxReceivedBitsPerSecond for the instance
func (instance *MSFT_MTNetworkAdapter) GetPropertyMaxReceivedBitsPerSecond() (value []float32, err error) {
	retValue, err := instance.GetProperty("MaxReceivedBitsPerSecond")
	if err != nil {
		return
	}
	value, ok := retValue.([]float32)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetMaxSentBitsPerSecond sets the value of MaxSentBitsPerSecond for the instance
func (instance *MSFT_MTNetworkAdapter) SetPropertyMaxSentBitsPerSecond(value []float32) (err error) {
	return instance.SetProperty("MaxSentBitsPerSecond", value)
}

// GetMaxSentBitsPerSecond gets the value of MaxSentBitsPerSecond for the instance
func (instance *MSFT_MTNetworkAdapter) GetPropertyMaxSentBitsPerSecond() (value []float32, err error) {
	retValue, err := instance.GetProperty("MaxSentBitsPerSecond")
	if err != nil {
		return
	}
	value, ok := retValue.([]float32)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetName sets the value of Name for the instance
func (instance *MSFT_MTNetworkAdapter) SetPropertyName(value string) (err error) {
	return instance.SetProperty("Name", value)
}

// GetName gets the value of Name for the instance
func (instance *MSFT_MTNetworkAdapter) GetPropertyName() (value string, err error) {
	retValue, err := instance.GetProperty("Name")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetNdisMedium sets the value of NdisMedium for the instance
func (instance *MSFT_MTNetworkAdapter) SetPropertyNdisMedium(value uint32) (err error) {
	return instance.SetProperty("NdisMedium", value)
}

// GetNdisMedium gets the value of NdisMedium for the instance
func (instance *MSFT_MTNetworkAdapter) GetPropertyNdisMedium() (value uint32, err error) {
	retValue, err := instance.GetProperty("NdisMedium")
	if err != nil {
		return
	}
	value, ok := retValue.(uint32)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetNdisPhysicalMedium sets the value of NdisPhysicalMedium for the instance
func (instance *MSFT_MTNetworkAdapter) SetPropertyNdisPhysicalMedium(value uint32) (err error) {
	return instance.SetProperty("NdisPhysicalMedium", value)
}

// GetNdisPhysicalMedium gets the value of NdisPhysicalMedium for the instance
func (instance *MSFT_MTNetworkAdapter) GetPropertyNdisPhysicalMedium() (value uint32, err error) {
	retValue, err := instance.GetProperty("NdisPhysicalMedium")
	if err != nil {
		return
	}
	value, ok := retValue.(uint32)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetNonUniCastsReceived sets the value of NonUniCastsReceived for the instance
func (instance *MSFT_MTNetworkAdapter) SetPropertyNonUniCastsReceived(value uint64) (err error) {
	return instance.SetProperty("NonUniCastsReceived", value)
}

// GetNonUniCastsReceived gets the value of NonUniCastsReceived for the instance
func (instance *MSFT_MTNetworkAdapter) GetPropertyNonUniCastsReceived() (value uint64, err error) {
	retValue, err := instance.GetProperty("NonUniCastsReceived")
	if err != nil {
		return
	}
	value, ok := retValue.(uint64)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetNonUniCastsReceivedPerInterval sets the value of NonUniCastsReceivedPerInterval for the instance
func (instance *MSFT_MTNetworkAdapter) SetPropertyNonUniCastsReceivedPerInterval(value uint64) (err error) {
	return instance.SetProperty("NonUniCastsReceivedPerInterval", value)
}

// GetNonUniCastsReceivedPerInterval gets the value of NonUniCastsReceivedPerInterval for the instance
func (instance *MSFT_MTNetworkAdapter) GetPropertyNonUniCastsReceivedPerInterval() (value uint64, err error) {
	retValue, err := instance.GetProperty("NonUniCastsReceivedPerInterval")
	if err != nil {
		return
	}
	value, ok := retValue.(uint64)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetNonUniCastsSent sets the value of NonUniCastsSent for the instance
func (instance *MSFT_MTNetworkAdapter) SetPropertyNonUniCastsSent(value uint64) (err error) {
	return instance.SetProperty("NonUniCastsSent", value)
}

// GetNonUniCastsSent gets the value of NonUniCastsSent for the instance
func (instance *MSFT_MTNetworkAdapter) GetPropertyNonUniCastsSent() (value uint64, err error) {
	retValue, err := instance.GetProperty("NonUniCastsSent")
	if err != nil {
		return
	}
	value, ok := retValue.(uint64)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetNonUniCastsSentPerInterval sets the value of NonUniCastsSentPerInterval for the instance
func (instance *MSFT_MTNetworkAdapter) SetPropertyNonUniCastsSentPerInterval(value uint64) (err error) {
	return instance.SetProperty("NonUniCastsSentPerInterval", value)
}

// GetNonUniCastsSentPerInterval gets the value of NonUniCastsSentPerInterval for the instance
func (instance *MSFT_MTNetworkAdapter) GetPropertyNonUniCastsSentPerInterval() (value uint64, err error) {
	retValue, err := instance.GetProperty("NonUniCastsSentPerInterval")
	if err != nil {
		return
	}
	value, ok := retValue.(uint64)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetNonUniCastsTotal sets the value of NonUniCastsTotal for the instance
func (instance *MSFT_MTNetworkAdapter) SetPropertyNonUniCastsTotal(value uint64) (err error) {
	return instance.SetProperty("NonUniCastsTotal", value)
}

// GetNonUniCastsTotal gets the value of NonUniCastsTotal for the instance
func (instance *MSFT_MTNetworkAdapter) GetPropertyNonUniCastsTotal() (value uint64, err error) {
	retValue, err := instance.GetProperty("NonUniCastsTotal")
	if err != nil {
		return
	}
	value, ok := retValue.(uint64)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetNonUniCastsTotalPerInterval sets the value of NonUniCastsTotalPerInterval for the instance
func (instance *MSFT_MTNetworkAdapter) SetPropertyNonUniCastsTotalPerInterval(value uint64) (err error) {
	return instance.SetProperty("NonUniCastsTotalPerInterval", value)
}

// GetNonUniCastsTotalPerInterval gets the value of NonUniCastsTotalPerInterval for the instance
func (instance *MSFT_MTNetworkAdapter) GetPropertyNonUniCastsTotalPerInterval() (value uint64, err error) {
	retValue, err := instance.GetProperty("NonUniCastsTotalPerInterval")
	if err != nil {
		return
	}
	value, ok := retValue.(uint64)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetOperationStatus sets the value of OperationStatus for the instance
func (instance *MSFT_MTNetworkAdapter) SetPropertyOperationStatus(value uint16) (err error) {
	return instance.SetProperty("OperationStatus", value)
}

// GetOperationStatus gets the value of OperationStatus for the instance
func (instance *MSFT_MTNetworkAdapter) GetPropertyOperationStatus() (value uint16, err error) {
	retValue, err := instance.GetProperty("OperationStatus")
	if err != nil {
		return
	}
	value, ok := retValue.(uint16)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetReceivedBitsPerSecond sets the value of ReceivedBitsPerSecond for the instance
func (instance *MSFT_MTNetworkAdapter) SetPropertyReceivedBitsPerSecond(value []float32) (err error) {
	return instance.SetProperty("ReceivedBitsPerSecond", value)
}

// GetReceivedBitsPerSecond gets the value of ReceivedBitsPerSecond for the instance
func (instance *MSFT_MTNetworkAdapter) GetPropertyReceivedBitsPerSecond() (value []float32, err error) {
	retValue, err := instance.GetProperty("ReceivedBitsPerSecond")
	if err != nil {
		return
	}
	value, ok := retValue.([]float32)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetReceivedThroughput sets the value of ReceivedThroughput for the instance
func (instance *MSFT_MTNetworkAdapter) SetPropertyReceivedThroughput(value []float32) (err error) {
	return instance.SetProperty("ReceivedThroughput", value)
}

// GetReceivedThroughput gets the value of ReceivedThroughput for the instance
func (instance *MSFT_MTNetworkAdapter) GetPropertyReceivedThroughput() (value []float32, err error) {
	retValue, err := instance.GetProperty("ReceivedThroughput")
	if err != nil {
		return
	}
	value, ok := retValue.([]float32)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetSentBitsPerSecond sets the value of SentBitsPerSecond for the instance
func (instance *MSFT_MTNetworkAdapter) SetPropertySentBitsPerSecond(value []float32) (err error) {
	return instance.SetProperty("SentBitsPerSecond", value)
}

// GetSentBitsPerSecond gets the value of SentBitsPerSecond for the instance
func (instance *MSFT_MTNetworkAdapter) GetPropertySentBitsPerSecond() (value []float32, err error) {
	retValue, err := instance.GetProperty("SentBitsPerSecond")
	if err != nil {
		return
	}
	value, ok := retValue.([]float32)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetSentThroughput sets the value of SentThroughput for the instance
func (instance *MSFT_MTNetworkAdapter) SetPropertySentThroughput(value []float32) (err error) {
	return instance.SetProperty("SentThroughput", value)
}

// GetSentThroughput gets the value of SentThroughput for the instance
func (instance *MSFT_MTNetworkAdapter) GetPropertySentThroughput() (value []float32, err error) {
	retValue, err := instance.GetProperty("SentThroughput")
	if err != nil {
		return
	}
	value, ok := retValue.([]float32)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetUniCastsReceived sets the value of UniCastsReceived for the instance
func (instance *MSFT_MTNetworkAdapter) SetPropertyUniCastsReceived(value uint64) (err error) {
	return instance.SetProperty("UniCastsReceived", value)
}

// GetUniCastsReceived gets the value of UniCastsReceived for the instance
func (instance *MSFT_MTNetworkAdapter) GetPropertyUniCastsReceived() (value uint64, err error) {
	retValue, err := instance.GetProperty("UniCastsReceived")
	if err != nil {
		return
	}
	value, ok := retValue.(uint64)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetUniCastsReceivedPerInterval sets the value of UniCastsReceivedPerInterval for the instance
func (instance *MSFT_MTNetworkAdapter) SetPropertyUniCastsReceivedPerInterval(value uint64) (err error) {
	return instance.SetProperty("UniCastsReceivedPerInterval", value)
}

// GetUniCastsReceivedPerInterval gets the value of UniCastsReceivedPerInterval for the instance
func (instance *MSFT_MTNetworkAdapter) GetPropertyUniCastsReceivedPerInterval() (value uint64, err error) {
	retValue, err := instance.GetProperty("UniCastsReceivedPerInterval")
	if err != nil {
		return
	}
	value, ok := retValue.(uint64)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetUniCastsSent sets the value of UniCastsSent for the instance
func (instance *MSFT_MTNetworkAdapter) SetPropertyUniCastsSent(value uint64) (err error) {
	return instance.SetProperty("UniCastsSent", value)
}

// GetUniCastsSent gets the value of UniCastsSent for the instance
func (instance *MSFT_MTNetworkAdapter) GetPropertyUniCastsSent() (value uint64, err error) {
	retValue, err := instance.GetProperty("UniCastsSent")
	if err != nil {
		return
	}
	value, ok := retValue.(uint64)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetUniCastsSentPerInterval sets the value of UniCastsSentPerInterval for the instance
func (instance *MSFT_MTNetworkAdapter) SetPropertyUniCastsSentPerInterval(value uint64) (err error) {
	return instance.SetProperty("UniCastsSentPerInterval", value)
}

// GetUniCastsSentPerInterval gets the value of UniCastsSentPerInterval for the instance
func (instance *MSFT_MTNetworkAdapter) GetPropertyUniCastsSentPerInterval() (value uint64, err error) {
	retValue, err := instance.GetProperty("UniCastsSentPerInterval")
	if err != nil {
		return
	}
	value, ok := retValue.(uint64)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetUniCastsTotal sets the value of UniCastsTotal for the instance
func (instance *MSFT_MTNetworkAdapter) SetPropertyUniCastsTotal(value uint64) (err error) {
	return instance.SetProperty("UniCastsTotal", value)
}

// GetUniCastsTotal gets the value of UniCastsTotal for the instance
func (instance *MSFT_MTNetworkAdapter) GetPropertyUniCastsTotal() (value uint64, err error) {
	retValue, err := instance.GetProperty("UniCastsTotal")
	if err != nil {
		return
	}
	value, ok := retValue.(uint64)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetUniCastsTotalPerInterval sets the value of UniCastsTotalPerInterval for the instance
func (instance *MSFT_MTNetworkAdapter) SetPropertyUniCastsTotalPerInterval(value uint64) (err error) {
	return instance.SetProperty("UniCastsTotalPerInterval", value)
}

// GetUniCastsTotalPerInterval gets the value of UniCastsTotalPerInterval for the instance
func (instance *MSFT_MTNetworkAdapter) GetPropertyUniCastsTotalPerInterval() (value uint64, err error) {
	retValue, err := instance.GetProperty("UniCastsTotalPerInterval")
	if err != nil {
		return
	}
	value, ok := retValue.(uint64)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetUtilization sets the value of Utilization for the instance
func (instance *MSFT_MTNetworkAdapter) SetPropertyUtilization(value float32) (err error) {
	return instance.SetProperty("Utilization", value)
}

// GetUtilization gets the value of Utilization for the instance
func (instance *MSFT_MTNetworkAdapter) GetPropertyUtilization() (value float32, err error) {
	retValue, err := instance.GetProperty("Utilization")
	if err != nil {
		return
	}
	value, ok := retValue.(float32)
	if !ok {
		// TODO: Set an error
	}
	return
}
