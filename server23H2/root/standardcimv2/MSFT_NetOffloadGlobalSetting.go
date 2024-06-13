// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 6/13/2024 using wmigen
//      Source root.StandardCimv2
//////////////////////////////////////////////
package standardcimv2

import (
	"github.com/microsoft/wmi/pkg/base/query"
	"github.com/microsoft/wmi/pkg/errors"
	cim "github.com/microsoft/wmi/pkg/wmiinstance"
	"reflect"
)

// MSFT_NetOffloadGlobalSetting struct
type MSFT_NetOffloadGlobalSetting struct {
	*MSFT_NetSettingData

	//
	Chimney uint8

	//
	NetworkDirect uint8

	//
	NetworkDirectAcrossIPSubnets uint8

	//
	PacketCoalescingFilter uint8

	//
	ReceiveSegmentCoalescing uint8

	//
	ReceiveSideScaling uint8

	//
	TaskOffload uint8
}

func NewMSFT_NetOffloadGlobalSettingEx1(instance *cim.WmiInstance) (newInstance *MSFT_NetOffloadGlobalSetting, err error) {
	tmp, err := NewMSFT_NetSettingDataEx1(instance)

	if err != nil {
		return
	}
	newInstance = &MSFT_NetOffloadGlobalSetting{
		MSFT_NetSettingData: tmp,
	}
	return
}

func NewMSFT_NetOffloadGlobalSettingEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *MSFT_NetOffloadGlobalSetting, err error) {
	tmp, err := NewMSFT_NetSettingDataEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &MSFT_NetOffloadGlobalSetting{
		MSFT_NetSettingData: tmp,
	}
	return
}

// SetChimney sets the value of Chimney for the instance
func (instance *MSFT_NetOffloadGlobalSetting) SetPropertyChimney(value uint8) (err error) {
	return instance.SetProperty("Chimney", (value))
}

// GetChimney gets the value of Chimney for the instance
func (instance *MSFT_NetOffloadGlobalSetting) GetPropertyChimney() (value uint8, err error) {
	retValue, err := instance.GetProperty("Chimney")
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

// SetNetworkDirect sets the value of NetworkDirect for the instance
func (instance *MSFT_NetOffloadGlobalSetting) SetPropertyNetworkDirect(value uint8) (err error) {
	return instance.SetProperty("NetworkDirect", (value))
}

// GetNetworkDirect gets the value of NetworkDirect for the instance
func (instance *MSFT_NetOffloadGlobalSetting) GetPropertyNetworkDirect() (value uint8, err error) {
	retValue, err := instance.GetProperty("NetworkDirect")
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

// SetNetworkDirectAcrossIPSubnets sets the value of NetworkDirectAcrossIPSubnets for the instance
func (instance *MSFT_NetOffloadGlobalSetting) SetPropertyNetworkDirectAcrossIPSubnets(value uint8) (err error) {
	return instance.SetProperty("NetworkDirectAcrossIPSubnets", (value))
}

// GetNetworkDirectAcrossIPSubnets gets the value of NetworkDirectAcrossIPSubnets for the instance
func (instance *MSFT_NetOffloadGlobalSetting) GetPropertyNetworkDirectAcrossIPSubnets() (value uint8, err error) {
	retValue, err := instance.GetProperty("NetworkDirectAcrossIPSubnets")
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

// SetPacketCoalescingFilter sets the value of PacketCoalescingFilter for the instance
func (instance *MSFT_NetOffloadGlobalSetting) SetPropertyPacketCoalescingFilter(value uint8) (err error) {
	return instance.SetProperty("PacketCoalescingFilter", (value))
}

// GetPacketCoalescingFilter gets the value of PacketCoalescingFilter for the instance
func (instance *MSFT_NetOffloadGlobalSetting) GetPropertyPacketCoalescingFilter() (value uint8, err error) {
	retValue, err := instance.GetProperty("PacketCoalescingFilter")
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

// SetReceiveSegmentCoalescing sets the value of ReceiveSegmentCoalescing for the instance
func (instance *MSFT_NetOffloadGlobalSetting) SetPropertyReceiveSegmentCoalescing(value uint8) (err error) {
	return instance.SetProperty("ReceiveSegmentCoalescing", (value))
}

// GetReceiveSegmentCoalescing gets the value of ReceiveSegmentCoalescing for the instance
func (instance *MSFT_NetOffloadGlobalSetting) GetPropertyReceiveSegmentCoalescing() (value uint8, err error) {
	retValue, err := instance.GetProperty("ReceiveSegmentCoalescing")
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

// SetReceiveSideScaling sets the value of ReceiveSideScaling for the instance
func (instance *MSFT_NetOffloadGlobalSetting) SetPropertyReceiveSideScaling(value uint8) (err error) {
	return instance.SetProperty("ReceiveSideScaling", (value))
}

// GetReceiveSideScaling gets the value of ReceiveSideScaling for the instance
func (instance *MSFT_NetOffloadGlobalSetting) GetPropertyReceiveSideScaling() (value uint8, err error) {
	retValue, err := instance.GetProperty("ReceiveSideScaling")
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

// SetTaskOffload sets the value of TaskOffload for the instance
func (instance *MSFT_NetOffloadGlobalSetting) SetPropertyTaskOffload(value uint8) (err error) {
	return instance.SetProperty("TaskOffload", (value))
}

// GetTaskOffload gets the value of TaskOffload for the instance
func (instance *MSFT_NetOffloadGlobalSetting) GetPropertyTaskOffload() (value uint8, err error) {
	retValue, err := instance.GetProperty("TaskOffload")
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
