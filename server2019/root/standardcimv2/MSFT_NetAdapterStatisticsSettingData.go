// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 3/19/2020 using wmigen
//      Source root.StandardCimv2
//////////////////////////////////////////////
package standardcimv2

import (
	"github.com/microsoft/wmi/pkg/base/query"
	cim "github.com/microsoft/wmi/pkg/wmiinstance"
)

// MSFT_NetAdapterStatisticsSettingData struct
type MSFT_NetAdapterStatisticsSettingData struct {
	*MSFT_NetAdapterSettingData

	//
	OutboundDiscardedPackets uint64

	//
	OutboundPacketErrors uint64

	//
	RdmaStatistics MSFT_NetAdapter_RdmaStatistics

	//
	ReceivedBroadcastBytes uint64

	//
	ReceivedBroadcastPackets uint64

	//
	ReceivedBytes uint64

	//
	ReceivedDiscardedPackets uint64

	//
	ReceivedMulticastBytes uint64

	//
	ReceivedMulticastPackets uint64

	//
	ReceivedPacketErrors uint64

	//
	ReceivedUnicastBytes uint64

	//
	ReceivedUnicastPackets uint64

	//
	RscStatistics MSFT_NetAdapter_RscStatistics

	//
	SentBroadcastBytes uint64

	//
	SentBroadcastPackets uint64

	//
	SentBytes uint64

	//
	SentMulticastBytes uint64

	//
	SentMulticastPackets uint64

	//
	SentUnicastBytes uint64

	//
	SentUnicastPackets uint64

	//
	SupportedStatistics uint32
}

func NewMSFT_NetAdapterStatisticsSettingDataEx1(instance *cim.WmiInstance) (newInstance *MSFT_NetAdapterStatisticsSettingData, err error) {
	tmp, err := NewMSFT_NetAdapterSettingDataEx1(instance)

	if err != nil {
		return
	}
	newInstance = &MSFT_NetAdapterStatisticsSettingData{
		MSFT_NetAdapterSettingData: tmp,
	}
	return
}

func NewMSFT_NetAdapterStatisticsSettingDataEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *MSFT_NetAdapterStatisticsSettingData, err error) {
	tmp, err := NewMSFT_NetAdapterSettingDataEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &MSFT_NetAdapterStatisticsSettingData{
		MSFT_NetAdapterSettingData: tmp,
	}
	return
}

// SetOutboundDiscardedPackets sets the value of OutboundDiscardedPackets for the instance
func (instance *MSFT_NetAdapterStatisticsSettingData) SetPropertyOutboundDiscardedPackets(value uint64) (err error) {
	return instance.SetProperty("OutboundDiscardedPackets", value)
}

// GetOutboundDiscardedPackets gets the value of OutboundDiscardedPackets for the instance
func (instance *MSFT_NetAdapterStatisticsSettingData) GetPropertyOutboundDiscardedPackets() (value uint64, err error) {
	retValue, err := instance.GetProperty("OutboundDiscardedPackets")
	if err != nil {
		return
	}
	value, ok := retValue.(uint64)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetOutboundPacketErrors sets the value of OutboundPacketErrors for the instance
func (instance *MSFT_NetAdapterStatisticsSettingData) SetPropertyOutboundPacketErrors(value uint64) (err error) {
	return instance.SetProperty("OutboundPacketErrors", value)
}

// GetOutboundPacketErrors gets the value of OutboundPacketErrors for the instance
func (instance *MSFT_NetAdapterStatisticsSettingData) GetPropertyOutboundPacketErrors() (value uint64, err error) {
	retValue, err := instance.GetProperty("OutboundPacketErrors")
	if err != nil {
		return
	}
	value, ok := retValue.(uint64)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetRdmaStatistics sets the value of RdmaStatistics for the instance
func (instance *MSFT_NetAdapterStatisticsSettingData) SetPropertyRdmaStatistics(value MSFT_NetAdapter_RdmaStatistics) (err error) {
	return instance.SetProperty("RdmaStatistics", value)
}

// GetRdmaStatistics gets the value of RdmaStatistics for the instance
func (instance *MSFT_NetAdapterStatisticsSettingData) GetPropertyRdmaStatistics() (value MSFT_NetAdapter_RdmaStatistics, err error) {
	retValue, err := instance.GetProperty("RdmaStatistics")
	if err != nil {
		return
	}
	value, ok := retValue.(MSFT_NetAdapter_RdmaStatistics)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetReceivedBroadcastBytes sets the value of ReceivedBroadcastBytes for the instance
func (instance *MSFT_NetAdapterStatisticsSettingData) SetPropertyReceivedBroadcastBytes(value uint64) (err error) {
	return instance.SetProperty("ReceivedBroadcastBytes", value)
}

// GetReceivedBroadcastBytes gets the value of ReceivedBroadcastBytes for the instance
func (instance *MSFT_NetAdapterStatisticsSettingData) GetPropertyReceivedBroadcastBytes() (value uint64, err error) {
	retValue, err := instance.GetProperty("ReceivedBroadcastBytes")
	if err != nil {
		return
	}
	value, ok := retValue.(uint64)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetReceivedBroadcastPackets sets the value of ReceivedBroadcastPackets for the instance
func (instance *MSFT_NetAdapterStatisticsSettingData) SetPropertyReceivedBroadcastPackets(value uint64) (err error) {
	return instance.SetProperty("ReceivedBroadcastPackets", value)
}

// GetReceivedBroadcastPackets gets the value of ReceivedBroadcastPackets for the instance
func (instance *MSFT_NetAdapterStatisticsSettingData) GetPropertyReceivedBroadcastPackets() (value uint64, err error) {
	retValue, err := instance.GetProperty("ReceivedBroadcastPackets")
	if err != nil {
		return
	}
	value, ok := retValue.(uint64)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetReceivedBytes sets the value of ReceivedBytes for the instance
func (instance *MSFT_NetAdapterStatisticsSettingData) SetPropertyReceivedBytes(value uint64) (err error) {
	return instance.SetProperty("ReceivedBytes", value)
}

// GetReceivedBytes gets the value of ReceivedBytes for the instance
func (instance *MSFT_NetAdapterStatisticsSettingData) GetPropertyReceivedBytes() (value uint64, err error) {
	retValue, err := instance.GetProperty("ReceivedBytes")
	if err != nil {
		return
	}
	value, ok := retValue.(uint64)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetReceivedDiscardedPackets sets the value of ReceivedDiscardedPackets for the instance
func (instance *MSFT_NetAdapterStatisticsSettingData) SetPropertyReceivedDiscardedPackets(value uint64) (err error) {
	return instance.SetProperty("ReceivedDiscardedPackets", value)
}

// GetReceivedDiscardedPackets gets the value of ReceivedDiscardedPackets for the instance
func (instance *MSFT_NetAdapterStatisticsSettingData) GetPropertyReceivedDiscardedPackets() (value uint64, err error) {
	retValue, err := instance.GetProperty("ReceivedDiscardedPackets")
	if err != nil {
		return
	}
	value, ok := retValue.(uint64)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetReceivedMulticastBytes sets the value of ReceivedMulticastBytes for the instance
func (instance *MSFT_NetAdapterStatisticsSettingData) SetPropertyReceivedMulticastBytes(value uint64) (err error) {
	return instance.SetProperty("ReceivedMulticastBytes", value)
}

// GetReceivedMulticastBytes gets the value of ReceivedMulticastBytes for the instance
func (instance *MSFT_NetAdapterStatisticsSettingData) GetPropertyReceivedMulticastBytes() (value uint64, err error) {
	retValue, err := instance.GetProperty("ReceivedMulticastBytes")
	if err != nil {
		return
	}
	value, ok := retValue.(uint64)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetReceivedMulticastPackets sets the value of ReceivedMulticastPackets for the instance
func (instance *MSFT_NetAdapterStatisticsSettingData) SetPropertyReceivedMulticastPackets(value uint64) (err error) {
	return instance.SetProperty("ReceivedMulticastPackets", value)
}

// GetReceivedMulticastPackets gets the value of ReceivedMulticastPackets for the instance
func (instance *MSFT_NetAdapterStatisticsSettingData) GetPropertyReceivedMulticastPackets() (value uint64, err error) {
	retValue, err := instance.GetProperty("ReceivedMulticastPackets")
	if err != nil {
		return
	}
	value, ok := retValue.(uint64)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetReceivedPacketErrors sets the value of ReceivedPacketErrors for the instance
func (instance *MSFT_NetAdapterStatisticsSettingData) SetPropertyReceivedPacketErrors(value uint64) (err error) {
	return instance.SetProperty("ReceivedPacketErrors", value)
}

// GetReceivedPacketErrors gets the value of ReceivedPacketErrors for the instance
func (instance *MSFT_NetAdapterStatisticsSettingData) GetPropertyReceivedPacketErrors() (value uint64, err error) {
	retValue, err := instance.GetProperty("ReceivedPacketErrors")
	if err != nil {
		return
	}
	value, ok := retValue.(uint64)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetReceivedUnicastBytes sets the value of ReceivedUnicastBytes for the instance
func (instance *MSFT_NetAdapterStatisticsSettingData) SetPropertyReceivedUnicastBytes(value uint64) (err error) {
	return instance.SetProperty("ReceivedUnicastBytes", value)
}

// GetReceivedUnicastBytes gets the value of ReceivedUnicastBytes for the instance
func (instance *MSFT_NetAdapterStatisticsSettingData) GetPropertyReceivedUnicastBytes() (value uint64, err error) {
	retValue, err := instance.GetProperty("ReceivedUnicastBytes")
	if err != nil {
		return
	}
	value, ok := retValue.(uint64)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetReceivedUnicastPackets sets the value of ReceivedUnicastPackets for the instance
func (instance *MSFT_NetAdapterStatisticsSettingData) SetPropertyReceivedUnicastPackets(value uint64) (err error) {
	return instance.SetProperty("ReceivedUnicastPackets", value)
}

// GetReceivedUnicastPackets gets the value of ReceivedUnicastPackets for the instance
func (instance *MSFT_NetAdapterStatisticsSettingData) GetPropertyReceivedUnicastPackets() (value uint64, err error) {
	retValue, err := instance.GetProperty("ReceivedUnicastPackets")
	if err != nil {
		return
	}
	value, ok := retValue.(uint64)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetRscStatistics sets the value of RscStatistics for the instance
func (instance *MSFT_NetAdapterStatisticsSettingData) SetPropertyRscStatistics(value MSFT_NetAdapter_RscStatistics) (err error) {
	return instance.SetProperty("RscStatistics", value)
}

// GetRscStatistics gets the value of RscStatistics for the instance
func (instance *MSFT_NetAdapterStatisticsSettingData) GetPropertyRscStatistics() (value MSFT_NetAdapter_RscStatistics, err error) {
	retValue, err := instance.GetProperty("RscStatistics")
	if err != nil {
		return
	}
	value, ok := retValue.(MSFT_NetAdapter_RscStatistics)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetSentBroadcastBytes sets the value of SentBroadcastBytes for the instance
func (instance *MSFT_NetAdapterStatisticsSettingData) SetPropertySentBroadcastBytes(value uint64) (err error) {
	return instance.SetProperty("SentBroadcastBytes", value)
}

// GetSentBroadcastBytes gets the value of SentBroadcastBytes for the instance
func (instance *MSFT_NetAdapterStatisticsSettingData) GetPropertySentBroadcastBytes() (value uint64, err error) {
	retValue, err := instance.GetProperty("SentBroadcastBytes")
	if err != nil {
		return
	}
	value, ok := retValue.(uint64)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetSentBroadcastPackets sets the value of SentBroadcastPackets for the instance
func (instance *MSFT_NetAdapterStatisticsSettingData) SetPropertySentBroadcastPackets(value uint64) (err error) {
	return instance.SetProperty("SentBroadcastPackets", value)
}

// GetSentBroadcastPackets gets the value of SentBroadcastPackets for the instance
func (instance *MSFT_NetAdapterStatisticsSettingData) GetPropertySentBroadcastPackets() (value uint64, err error) {
	retValue, err := instance.GetProperty("SentBroadcastPackets")
	if err != nil {
		return
	}
	value, ok := retValue.(uint64)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetSentBytes sets the value of SentBytes for the instance
func (instance *MSFT_NetAdapterStatisticsSettingData) SetPropertySentBytes(value uint64) (err error) {
	return instance.SetProperty("SentBytes", value)
}

// GetSentBytes gets the value of SentBytes for the instance
func (instance *MSFT_NetAdapterStatisticsSettingData) GetPropertySentBytes() (value uint64, err error) {
	retValue, err := instance.GetProperty("SentBytes")
	if err != nil {
		return
	}
	value, ok := retValue.(uint64)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetSentMulticastBytes sets the value of SentMulticastBytes for the instance
func (instance *MSFT_NetAdapterStatisticsSettingData) SetPropertySentMulticastBytes(value uint64) (err error) {
	return instance.SetProperty("SentMulticastBytes", value)
}

// GetSentMulticastBytes gets the value of SentMulticastBytes for the instance
func (instance *MSFT_NetAdapterStatisticsSettingData) GetPropertySentMulticastBytes() (value uint64, err error) {
	retValue, err := instance.GetProperty("SentMulticastBytes")
	if err != nil {
		return
	}
	value, ok := retValue.(uint64)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetSentMulticastPackets sets the value of SentMulticastPackets for the instance
func (instance *MSFT_NetAdapterStatisticsSettingData) SetPropertySentMulticastPackets(value uint64) (err error) {
	return instance.SetProperty("SentMulticastPackets", value)
}

// GetSentMulticastPackets gets the value of SentMulticastPackets for the instance
func (instance *MSFT_NetAdapterStatisticsSettingData) GetPropertySentMulticastPackets() (value uint64, err error) {
	retValue, err := instance.GetProperty("SentMulticastPackets")
	if err != nil {
		return
	}
	value, ok := retValue.(uint64)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetSentUnicastBytes sets the value of SentUnicastBytes for the instance
func (instance *MSFT_NetAdapterStatisticsSettingData) SetPropertySentUnicastBytes(value uint64) (err error) {
	return instance.SetProperty("SentUnicastBytes", value)
}

// GetSentUnicastBytes gets the value of SentUnicastBytes for the instance
func (instance *MSFT_NetAdapterStatisticsSettingData) GetPropertySentUnicastBytes() (value uint64, err error) {
	retValue, err := instance.GetProperty("SentUnicastBytes")
	if err != nil {
		return
	}
	value, ok := retValue.(uint64)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetSentUnicastPackets sets the value of SentUnicastPackets for the instance
func (instance *MSFT_NetAdapterStatisticsSettingData) SetPropertySentUnicastPackets(value uint64) (err error) {
	return instance.SetProperty("SentUnicastPackets", value)
}

// GetSentUnicastPackets gets the value of SentUnicastPackets for the instance
func (instance *MSFT_NetAdapterStatisticsSettingData) GetPropertySentUnicastPackets() (value uint64, err error) {
	retValue, err := instance.GetProperty("SentUnicastPackets")
	if err != nil {
		return
	}
	value, ok := retValue.(uint64)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetSupportedStatistics sets the value of SupportedStatistics for the instance
func (instance *MSFT_NetAdapterStatisticsSettingData) SetPropertySupportedStatistics(value uint32) (err error) {
	return instance.SetProperty("SupportedStatistics", value)
}

// GetSupportedStatistics gets the value of SupportedStatistics for the instance
func (instance *MSFT_NetAdapterStatisticsSettingData) GetPropertySupportedStatistics() (value uint32, err error) {
	retValue, err := instance.GetProperty("SupportedStatistics")
	if err != nil {
		return
	}
	value, ok := retValue.(uint32)
	if !ok {
		// TODO: Set an error
	}
	return
}
