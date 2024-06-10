// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 9/18/2020 using wmigen
//      Source root.StandardCimv2
//////////////////////////////////////////////
package standardcimv2

import (
	"github.com/microsoft/wmi/pkg/base/query"
	"github.com/microsoft/wmi/pkg/errors"
	cim "github.com/microsoft/wmi/pkg/wmiinstance"
	"reflect"
)

// MSFT_NetAdapterRssSettingData struct
type MSFT_NetAdapterRssSettingData struct {
	*MSFT_NetAdapterSettingData

	//
	BaseProcessorGroup uint16

	//
	BaseProcessorNumber uint8

	//
	ClassificationAtDpcSupported bool

	//
	ClassificationAtIsrSupported bool

	//
	Enabled bool

	//
	HashKeySize uint16

	//
	IndirectionTable []MSFT_NetAdapter_ProcessorNumber

	//
	IndirectionTableEntryCount uint16

	//
	IPv4HashEnabled bool

	//
	IPv6ExtensionHashEnabled bool

	//
	IPv6HashEnabled bool

	//
	MaxProcessorGroup uint16

	//
	MaxProcessorNumber uint8

	//
	MaxProcessors uint32

	//
	MsiSupported bool

	//
	MsiXEnabled bool

	//
	MsiXSupported bool

	//
	NumaNode uint16

	//
	NumberOfInterruptMessages uint32

	//
	NumberOfReceiveQueues uint32

	//
	Profile uint32

	//
	RssOnPortsSupported bool

	//
	RssProcessorArray []MSFT_NetAdapter_RssProcessor

	//
	RssProcessorArraySize uint32

	//
	TcpIPv4HashEnabled bool

	//
	TcpIPv4HashSupported bool

	//
	TcpIPv6ExtensionHashEnabled bool

	//
	TcpIPv6ExtensionHashSupported bool

	//
	TcpIPv6HashEnabled bool

	//
	TcpIPv6HashSupported bool

	//
	ToeplitzHashFunctionEnabled bool

	//
	ToeplitzHashFunctionSupported bool

	//
	UdpIPv4HashEnabled bool

	//
	UdpIPv4HashSupported bool

	//
	UdpIPv6ExtensionHashEnabled bool

	//
	UdpIPv6ExtensionHashSupported bool

	//
	UdpIPv6HashEnabled bool

	//
	UdpIPv6HashSupported bool
}

func NewMSFT_NetAdapterRssSettingDataEx1(instance *cim.WmiInstance) (newInstance *MSFT_NetAdapterRssSettingData, err error) {
	tmp, err := NewMSFT_NetAdapterSettingDataEx1(instance)

	if err != nil {
		return
	}
	newInstance = &MSFT_NetAdapterRssSettingData{
		MSFT_NetAdapterSettingData: tmp,
	}
	return
}

func NewMSFT_NetAdapterRssSettingDataEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *MSFT_NetAdapterRssSettingData, err error) {
	tmp, err := NewMSFT_NetAdapterSettingDataEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &MSFT_NetAdapterRssSettingData{
		MSFT_NetAdapterSettingData: tmp,
	}
	return
}

// SetBaseProcessorGroup sets the value of BaseProcessorGroup for the instance
func (instance *MSFT_NetAdapterRssSettingData) SetPropertyBaseProcessorGroup(value uint16) (err error) {
	return instance.SetProperty("BaseProcessorGroup", (value))
}

// GetBaseProcessorGroup gets the value of BaseProcessorGroup for the instance
func (instance *MSFT_NetAdapterRssSettingData) GetPropertyBaseProcessorGroup() (value uint16, err error) {
	retValue, err := instance.GetProperty("BaseProcessorGroup")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(uint16)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " uint16 is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = uint16(valuetmp)

	return
}

// SetBaseProcessorNumber sets the value of BaseProcessorNumber for the instance
func (instance *MSFT_NetAdapterRssSettingData) SetPropertyBaseProcessorNumber(value uint8) (err error) {
	return instance.SetProperty("BaseProcessorNumber", (value))
}

// GetBaseProcessorNumber gets the value of BaseProcessorNumber for the instance
func (instance *MSFT_NetAdapterRssSettingData) GetPropertyBaseProcessorNumber() (value uint8, err error) {
	retValue, err := instance.GetProperty("BaseProcessorNumber")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(uint8)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " uint8 is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = uint8(valuetmp)

	return
}

// SetClassificationAtDpcSupported sets the value of ClassificationAtDpcSupported for the instance
func (instance *MSFT_NetAdapterRssSettingData) SetPropertyClassificationAtDpcSupported(value bool) (err error) {
	return instance.SetProperty("ClassificationAtDpcSupported", (value))
}

// GetClassificationAtDpcSupported gets the value of ClassificationAtDpcSupported for the instance
func (instance *MSFT_NetAdapterRssSettingData) GetPropertyClassificationAtDpcSupported() (value bool, err error) {
	retValue, err := instance.GetProperty("ClassificationAtDpcSupported")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(bool)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " bool is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = bool(valuetmp)

	return
}

// SetClassificationAtIsrSupported sets the value of ClassificationAtIsrSupported for the instance
func (instance *MSFT_NetAdapterRssSettingData) SetPropertyClassificationAtIsrSupported(value bool) (err error) {
	return instance.SetProperty("ClassificationAtIsrSupported", (value))
}

// GetClassificationAtIsrSupported gets the value of ClassificationAtIsrSupported for the instance
func (instance *MSFT_NetAdapterRssSettingData) GetPropertyClassificationAtIsrSupported() (value bool, err error) {
	retValue, err := instance.GetProperty("ClassificationAtIsrSupported")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(bool)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " bool is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = bool(valuetmp)

	return
}

// SetEnabled sets the value of Enabled for the instance
func (instance *MSFT_NetAdapterRssSettingData) SetPropertyEnabled(value bool) (err error) {
	return instance.SetProperty("Enabled", (value))
}

// GetEnabled gets the value of Enabled for the instance
func (instance *MSFT_NetAdapterRssSettingData) GetPropertyEnabled() (value bool, err error) {
	retValue, err := instance.GetProperty("Enabled")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(bool)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " bool is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = bool(valuetmp)

	return
}

// SetHashKeySize sets the value of HashKeySize for the instance
func (instance *MSFT_NetAdapterRssSettingData) SetPropertyHashKeySize(value uint16) (err error) {
	return instance.SetProperty("HashKeySize", (value))
}

// GetHashKeySize gets the value of HashKeySize for the instance
func (instance *MSFT_NetAdapterRssSettingData) GetPropertyHashKeySize() (value uint16, err error) {
	retValue, err := instance.GetProperty("HashKeySize")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(uint16)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " uint16 is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = uint16(valuetmp)

	return
}

// SetIndirectionTable sets the value of IndirectionTable for the instance
func (instance *MSFT_NetAdapterRssSettingData) SetPropertyIndirectionTable(value []MSFT_NetAdapter_ProcessorNumber) (err error) {
	return instance.SetProperty("IndirectionTable", (value))
}

// GetIndirectionTable gets the value of IndirectionTable for the instance
func (instance *MSFT_NetAdapterRssSettingData) GetPropertyIndirectionTable() (value []MSFT_NetAdapter_ProcessorNumber, err error) {
	retValue, err := instance.GetProperty("IndirectionTable")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	for _, interfaceValue := range retValue.([]interface{}) {
		valuetmp, ok := interfaceValue.(MSFT_NetAdapter_ProcessorNumber)
		if !ok {
			err = errors.Wrapf(errors.InvalidType, " MSFT_NetAdapter_ProcessorNumber is Invalid. Expected %s", reflect.TypeOf(interfaceValue))
			return
		}
		value = append(value, MSFT_NetAdapter_ProcessorNumber(valuetmp))
	}

	return
}

// SetIndirectionTableEntryCount sets the value of IndirectionTableEntryCount for the instance
func (instance *MSFT_NetAdapterRssSettingData) SetPropertyIndirectionTableEntryCount(value uint16) (err error) {
	return instance.SetProperty("IndirectionTableEntryCount", (value))
}

// GetIndirectionTableEntryCount gets the value of IndirectionTableEntryCount for the instance
func (instance *MSFT_NetAdapterRssSettingData) GetPropertyIndirectionTableEntryCount() (value uint16, err error) {
	retValue, err := instance.GetProperty("IndirectionTableEntryCount")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(uint16)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " uint16 is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = uint16(valuetmp)

	return
}

// SetIPv4HashEnabled sets the value of IPv4HashEnabled for the instance
func (instance *MSFT_NetAdapterRssSettingData) SetPropertyIPv4HashEnabled(value bool) (err error) {
	return instance.SetProperty("IPv4HashEnabled", (value))
}

// GetIPv4HashEnabled gets the value of IPv4HashEnabled for the instance
func (instance *MSFT_NetAdapterRssSettingData) GetPropertyIPv4HashEnabled() (value bool, err error) {
	retValue, err := instance.GetProperty("IPv4HashEnabled")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(bool)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " bool is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = bool(valuetmp)

	return
}

// SetIPv6ExtensionHashEnabled sets the value of IPv6ExtensionHashEnabled for the instance
func (instance *MSFT_NetAdapterRssSettingData) SetPropertyIPv6ExtensionHashEnabled(value bool) (err error) {
	return instance.SetProperty("IPv6ExtensionHashEnabled", (value))
}

// GetIPv6ExtensionHashEnabled gets the value of IPv6ExtensionHashEnabled for the instance
func (instance *MSFT_NetAdapterRssSettingData) GetPropertyIPv6ExtensionHashEnabled() (value bool, err error) {
	retValue, err := instance.GetProperty("IPv6ExtensionHashEnabled")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(bool)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " bool is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = bool(valuetmp)

	return
}

// SetIPv6HashEnabled sets the value of IPv6HashEnabled for the instance
func (instance *MSFT_NetAdapterRssSettingData) SetPropertyIPv6HashEnabled(value bool) (err error) {
	return instance.SetProperty("IPv6HashEnabled", (value))
}

// GetIPv6HashEnabled gets the value of IPv6HashEnabled for the instance
func (instance *MSFT_NetAdapterRssSettingData) GetPropertyIPv6HashEnabled() (value bool, err error) {
	retValue, err := instance.GetProperty("IPv6HashEnabled")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(bool)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " bool is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = bool(valuetmp)

	return
}

// SetMaxProcessorGroup sets the value of MaxProcessorGroup for the instance
func (instance *MSFT_NetAdapterRssSettingData) SetPropertyMaxProcessorGroup(value uint16) (err error) {
	return instance.SetProperty("MaxProcessorGroup", (value))
}

// GetMaxProcessorGroup gets the value of MaxProcessorGroup for the instance
func (instance *MSFT_NetAdapterRssSettingData) GetPropertyMaxProcessorGroup() (value uint16, err error) {
	retValue, err := instance.GetProperty("MaxProcessorGroup")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(uint16)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " uint16 is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = uint16(valuetmp)

	return
}

// SetMaxProcessorNumber sets the value of MaxProcessorNumber for the instance
func (instance *MSFT_NetAdapterRssSettingData) SetPropertyMaxProcessorNumber(value uint8) (err error) {
	return instance.SetProperty("MaxProcessorNumber", (value))
}

// GetMaxProcessorNumber gets the value of MaxProcessorNumber for the instance
func (instance *MSFT_NetAdapterRssSettingData) GetPropertyMaxProcessorNumber() (value uint8, err error) {
	retValue, err := instance.GetProperty("MaxProcessorNumber")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(uint8)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " uint8 is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = uint8(valuetmp)

	return
}

// SetMaxProcessors sets the value of MaxProcessors for the instance
func (instance *MSFT_NetAdapterRssSettingData) SetPropertyMaxProcessors(value uint32) (err error) {
	return instance.SetProperty("MaxProcessors", (value))
}

// GetMaxProcessors gets the value of MaxProcessors for the instance
func (instance *MSFT_NetAdapterRssSettingData) GetPropertyMaxProcessors() (value uint32, err error) {
	retValue, err := instance.GetProperty("MaxProcessors")
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

// SetMsiSupported sets the value of MsiSupported for the instance
func (instance *MSFT_NetAdapterRssSettingData) SetPropertyMsiSupported(value bool) (err error) {
	return instance.SetProperty("MsiSupported", (value))
}

// GetMsiSupported gets the value of MsiSupported for the instance
func (instance *MSFT_NetAdapterRssSettingData) GetPropertyMsiSupported() (value bool, err error) {
	retValue, err := instance.GetProperty("MsiSupported")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(bool)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " bool is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = bool(valuetmp)

	return
}

// SetMsiXEnabled sets the value of MsiXEnabled for the instance
func (instance *MSFT_NetAdapterRssSettingData) SetPropertyMsiXEnabled(value bool) (err error) {
	return instance.SetProperty("MsiXEnabled", (value))
}

// GetMsiXEnabled gets the value of MsiXEnabled for the instance
func (instance *MSFT_NetAdapterRssSettingData) GetPropertyMsiXEnabled() (value bool, err error) {
	retValue, err := instance.GetProperty("MsiXEnabled")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(bool)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " bool is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = bool(valuetmp)

	return
}

// SetMsiXSupported sets the value of MsiXSupported for the instance
func (instance *MSFT_NetAdapterRssSettingData) SetPropertyMsiXSupported(value bool) (err error) {
	return instance.SetProperty("MsiXSupported", (value))
}

// GetMsiXSupported gets the value of MsiXSupported for the instance
func (instance *MSFT_NetAdapterRssSettingData) GetPropertyMsiXSupported() (value bool, err error) {
	retValue, err := instance.GetProperty("MsiXSupported")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(bool)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " bool is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = bool(valuetmp)

	return
}

// SetNumaNode sets the value of NumaNode for the instance
func (instance *MSFT_NetAdapterRssSettingData) SetPropertyNumaNode(value uint16) (err error) {
	return instance.SetProperty("NumaNode", (value))
}

// GetNumaNode gets the value of NumaNode for the instance
func (instance *MSFT_NetAdapterRssSettingData) GetPropertyNumaNode() (value uint16, err error) {
	retValue, err := instance.GetProperty("NumaNode")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(uint16)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " uint16 is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = uint16(valuetmp)

	return
}

// SetNumberOfInterruptMessages sets the value of NumberOfInterruptMessages for the instance
func (instance *MSFT_NetAdapterRssSettingData) SetPropertyNumberOfInterruptMessages(value uint32) (err error) {
	return instance.SetProperty("NumberOfInterruptMessages", (value))
}

// GetNumberOfInterruptMessages gets the value of NumberOfInterruptMessages for the instance
func (instance *MSFT_NetAdapterRssSettingData) GetPropertyNumberOfInterruptMessages() (value uint32, err error) {
	retValue, err := instance.GetProperty("NumberOfInterruptMessages")
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

// SetNumberOfReceiveQueues sets the value of NumberOfReceiveQueues for the instance
func (instance *MSFT_NetAdapterRssSettingData) SetPropertyNumberOfReceiveQueues(value uint32) (err error) {
	return instance.SetProperty("NumberOfReceiveQueues", (value))
}

// GetNumberOfReceiveQueues gets the value of NumberOfReceiveQueues for the instance
func (instance *MSFT_NetAdapterRssSettingData) GetPropertyNumberOfReceiveQueues() (value uint32, err error) {
	retValue, err := instance.GetProperty("NumberOfReceiveQueues")
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

// SetProfile sets the value of Profile for the instance
func (instance *MSFT_NetAdapterRssSettingData) SetPropertyProfile(value uint32) (err error) {
	return instance.SetProperty("Profile", (value))
}

// GetProfile gets the value of Profile for the instance
func (instance *MSFT_NetAdapterRssSettingData) GetPropertyProfile() (value uint32, err error) {
	retValue, err := instance.GetProperty("Profile")
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

// SetRssOnPortsSupported sets the value of RssOnPortsSupported for the instance
func (instance *MSFT_NetAdapterRssSettingData) SetPropertyRssOnPortsSupported(value bool) (err error) {
	return instance.SetProperty("RssOnPortsSupported", (value))
}

// GetRssOnPortsSupported gets the value of RssOnPortsSupported for the instance
func (instance *MSFT_NetAdapterRssSettingData) GetPropertyRssOnPortsSupported() (value bool, err error) {
	retValue, err := instance.GetProperty("RssOnPortsSupported")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(bool)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " bool is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = bool(valuetmp)

	return
}

// SetRssProcessorArray sets the value of RssProcessorArray for the instance
func (instance *MSFT_NetAdapterRssSettingData) SetPropertyRssProcessorArray(value []MSFT_NetAdapter_RssProcessor) (err error) {
	return instance.SetProperty("RssProcessorArray", (value))
}

// GetRssProcessorArray gets the value of RssProcessorArray for the instance
func (instance *MSFT_NetAdapterRssSettingData) GetPropertyRssProcessorArray() (value []MSFT_NetAdapter_RssProcessor, err error) {
	retValue, err := instance.GetProperty("RssProcessorArray")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	for _, interfaceValue := range retValue.([]interface{}) {
		valuetmp, ok := interfaceValue.(MSFT_NetAdapter_RssProcessor)
		if !ok {
			err = errors.Wrapf(errors.InvalidType, " MSFT_NetAdapter_RssProcessor is Invalid. Expected %s", reflect.TypeOf(interfaceValue))
			return
		}
		value = append(value, MSFT_NetAdapter_RssProcessor(valuetmp))
	}

	return
}

// SetRssProcessorArraySize sets the value of RssProcessorArraySize for the instance
func (instance *MSFT_NetAdapterRssSettingData) SetPropertyRssProcessorArraySize(value uint32) (err error) {
	return instance.SetProperty("RssProcessorArraySize", (value))
}

// GetRssProcessorArraySize gets the value of RssProcessorArraySize for the instance
func (instance *MSFT_NetAdapterRssSettingData) GetPropertyRssProcessorArraySize() (value uint32, err error) {
	retValue, err := instance.GetProperty("RssProcessorArraySize")
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

// SetTcpIPv4HashEnabled sets the value of TcpIPv4HashEnabled for the instance
func (instance *MSFT_NetAdapterRssSettingData) SetPropertyTcpIPv4HashEnabled(value bool) (err error) {
	return instance.SetProperty("TcpIPv4HashEnabled", (value))
}

// GetTcpIPv4HashEnabled gets the value of TcpIPv4HashEnabled for the instance
func (instance *MSFT_NetAdapterRssSettingData) GetPropertyTcpIPv4HashEnabled() (value bool, err error) {
	retValue, err := instance.GetProperty("TcpIPv4HashEnabled")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(bool)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " bool is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = bool(valuetmp)

	return
}

// SetTcpIPv4HashSupported sets the value of TcpIPv4HashSupported for the instance
func (instance *MSFT_NetAdapterRssSettingData) SetPropertyTcpIPv4HashSupported(value bool) (err error) {
	return instance.SetProperty("TcpIPv4HashSupported", (value))
}

// GetTcpIPv4HashSupported gets the value of TcpIPv4HashSupported for the instance
func (instance *MSFT_NetAdapterRssSettingData) GetPropertyTcpIPv4HashSupported() (value bool, err error) {
	retValue, err := instance.GetProperty("TcpIPv4HashSupported")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(bool)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " bool is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = bool(valuetmp)

	return
}

// SetTcpIPv6ExtensionHashEnabled sets the value of TcpIPv6ExtensionHashEnabled for the instance
func (instance *MSFT_NetAdapterRssSettingData) SetPropertyTcpIPv6ExtensionHashEnabled(value bool) (err error) {
	return instance.SetProperty("TcpIPv6ExtensionHashEnabled", (value))
}

// GetTcpIPv6ExtensionHashEnabled gets the value of TcpIPv6ExtensionHashEnabled for the instance
func (instance *MSFT_NetAdapterRssSettingData) GetPropertyTcpIPv6ExtensionHashEnabled() (value bool, err error) {
	retValue, err := instance.GetProperty("TcpIPv6ExtensionHashEnabled")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(bool)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " bool is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = bool(valuetmp)

	return
}

// SetTcpIPv6ExtensionHashSupported sets the value of TcpIPv6ExtensionHashSupported for the instance
func (instance *MSFT_NetAdapterRssSettingData) SetPropertyTcpIPv6ExtensionHashSupported(value bool) (err error) {
	return instance.SetProperty("TcpIPv6ExtensionHashSupported", (value))
}

// GetTcpIPv6ExtensionHashSupported gets the value of TcpIPv6ExtensionHashSupported for the instance
func (instance *MSFT_NetAdapterRssSettingData) GetPropertyTcpIPv6ExtensionHashSupported() (value bool, err error) {
	retValue, err := instance.GetProperty("TcpIPv6ExtensionHashSupported")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(bool)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " bool is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = bool(valuetmp)

	return
}

// SetTcpIPv6HashEnabled sets the value of TcpIPv6HashEnabled for the instance
func (instance *MSFT_NetAdapterRssSettingData) SetPropertyTcpIPv6HashEnabled(value bool) (err error) {
	return instance.SetProperty("TcpIPv6HashEnabled", (value))
}

// GetTcpIPv6HashEnabled gets the value of TcpIPv6HashEnabled for the instance
func (instance *MSFT_NetAdapterRssSettingData) GetPropertyTcpIPv6HashEnabled() (value bool, err error) {
	retValue, err := instance.GetProperty("TcpIPv6HashEnabled")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(bool)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " bool is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = bool(valuetmp)

	return
}

// SetTcpIPv6HashSupported sets the value of TcpIPv6HashSupported for the instance
func (instance *MSFT_NetAdapterRssSettingData) SetPropertyTcpIPv6HashSupported(value bool) (err error) {
	return instance.SetProperty("TcpIPv6HashSupported", (value))
}

// GetTcpIPv6HashSupported gets the value of TcpIPv6HashSupported for the instance
func (instance *MSFT_NetAdapterRssSettingData) GetPropertyTcpIPv6HashSupported() (value bool, err error) {
	retValue, err := instance.GetProperty("TcpIPv6HashSupported")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(bool)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " bool is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = bool(valuetmp)

	return
}

// SetToeplitzHashFunctionEnabled sets the value of ToeplitzHashFunctionEnabled for the instance
func (instance *MSFT_NetAdapterRssSettingData) SetPropertyToeplitzHashFunctionEnabled(value bool) (err error) {
	return instance.SetProperty("ToeplitzHashFunctionEnabled", (value))
}

// GetToeplitzHashFunctionEnabled gets the value of ToeplitzHashFunctionEnabled for the instance
func (instance *MSFT_NetAdapterRssSettingData) GetPropertyToeplitzHashFunctionEnabled() (value bool, err error) {
	retValue, err := instance.GetProperty("ToeplitzHashFunctionEnabled")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(bool)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " bool is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = bool(valuetmp)

	return
}

// SetToeplitzHashFunctionSupported sets the value of ToeplitzHashFunctionSupported for the instance
func (instance *MSFT_NetAdapterRssSettingData) SetPropertyToeplitzHashFunctionSupported(value bool) (err error) {
	return instance.SetProperty("ToeplitzHashFunctionSupported", (value))
}

// GetToeplitzHashFunctionSupported gets the value of ToeplitzHashFunctionSupported for the instance
func (instance *MSFT_NetAdapterRssSettingData) GetPropertyToeplitzHashFunctionSupported() (value bool, err error) {
	retValue, err := instance.GetProperty("ToeplitzHashFunctionSupported")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(bool)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " bool is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = bool(valuetmp)

	return
}

// SetUdpIPv4HashEnabled sets the value of UdpIPv4HashEnabled for the instance
func (instance *MSFT_NetAdapterRssSettingData) SetPropertyUdpIPv4HashEnabled(value bool) (err error) {
	return instance.SetProperty("UdpIPv4HashEnabled", (value))
}

// GetUdpIPv4HashEnabled gets the value of UdpIPv4HashEnabled for the instance
func (instance *MSFT_NetAdapterRssSettingData) GetPropertyUdpIPv4HashEnabled() (value bool, err error) {
	retValue, err := instance.GetProperty("UdpIPv4HashEnabled")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(bool)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " bool is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = bool(valuetmp)

	return
}

// SetUdpIPv4HashSupported sets the value of UdpIPv4HashSupported for the instance
func (instance *MSFT_NetAdapterRssSettingData) SetPropertyUdpIPv4HashSupported(value bool) (err error) {
	return instance.SetProperty("UdpIPv4HashSupported", (value))
}

// GetUdpIPv4HashSupported gets the value of UdpIPv4HashSupported for the instance
func (instance *MSFT_NetAdapterRssSettingData) GetPropertyUdpIPv4HashSupported() (value bool, err error) {
	retValue, err := instance.GetProperty("UdpIPv4HashSupported")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(bool)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " bool is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = bool(valuetmp)

	return
}

// SetUdpIPv6ExtensionHashEnabled sets the value of UdpIPv6ExtensionHashEnabled for the instance
func (instance *MSFT_NetAdapterRssSettingData) SetPropertyUdpIPv6ExtensionHashEnabled(value bool) (err error) {
	return instance.SetProperty("UdpIPv6ExtensionHashEnabled", (value))
}

// GetUdpIPv6ExtensionHashEnabled gets the value of UdpIPv6ExtensionHashEnabled for the instance
func (instance *MSFT_NetAdapterRssSettingData) GetPropertyUdpIPv6ExtensionHashEnabled() (value bool, err error) {
	retValue, err := instance.GetProperty("UdpIPv6ExtensionHashEnabled")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(bool)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " bool is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = bool(valuetmp)

	return
}

// SetUdpIPv6ExtensionHashSupported sets the value of UdpIPv6ExtensionHashSupported for the instance
func (instance *MSFT_NetAdapterRssSettingData) SetPropertyUdpIPv6ExtensionHashSupported(value bool) (err error) {
	return instance.SetProperty("UdpIPv6ExtensionHashSupported", (value))
}

// GetUdpIPv6ExtensionHashSupported gets the value of UdpIPv6ExtensionHashSupported for the instance
func (instance *MSFT_NetAdapterRssSettingData) GetPropertyUdpIPv6ExtensionHashSupported() (value bool, err error) {
	retValue, err := instance.GetProperty("UdpIPv6ExtensionHashSupported")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(bool)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " bool is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = bool(valuetmp)

	return
}

// SetUdpIPv6HashEnabled sets the value of UdpIPv6HashEnabled for the instance
func (instance *MSFT_NetAdapterRssSettingData) SetPropertyUdpIPv6HashEnabled(value bool) (err error) {
	return instance.SetProperty("UdpIPv6HashEnabled", (value))
}

// GetUdpIPv6HashEnabled gets the value of UdpIPv6HashEnabled for the instance
func (instance *MSFT_NetAdapterRssSettingData) GetPropertyUdpIPv6HashEnabled() (value bool, err error) {
	retValue, err := instance.GetProperty("UdpIPv6HashEnabled")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(bool)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " bool is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = bool(valuetmp)

	return
}

// SetUdpIPv6HashSupported sets the value of UdpIPv6HashSupported for the instance
func (instance *MSFT_NetAdapterRssSettingData) SetPropertyUdpIPv6HashSupported(value bool) (err error) {
	return instance.SetProperty("UdpIPv6HashSupported", (value))
}

// GetUdpIPv6HashSupported gets the value of UdpIPv6HashSupported for the instance
func (instance *MSFT_NetAdapterRssSettingData) GetPropertyUdpIPv6HashSupported() (value bool, err error) {
	retValue, err := instance.GetProperty("UdpIPv6HashSupported")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(bool)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " bool is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = bool(valuetmp)

	return
}

//

// <param name="cmdletOutput" type="MSFT_NetAdapterRssSettingData "></param>
// <param name="ReturnValue" type="uint32 "></param>
func (instance *MSFT_NetAdapterRssSettingData) Enable( /* OUT */ cmdletOutput MSFT_NetAdapterRssSettingData) (result uint32, err error) {
	retVal, err := instance.InvokeMethod("Enable")
	if err != nil {
		return
	}
	retValue := retVal[0].(int32)
	result = uint32(retValue)
	return

}

//

// <param name="cmdletOutput" type="MSFT_NetAdapterRssSettingData "></param>
// <param name="ReturnValue" type="uint32 "></param>
func (instance *MSFT_NetAdapterRssSettingData) Disable( /* OUT */ cmdletOutput MSFT_NetAdapterRssSettingData) (result uint32, err error) {
	retVal, err := instance.InvokeMethod("Disable")
	if err != nil {
		return
	}
	retValue := retVal[0].(int32)
	result = uint32(retValue)
	return

}
