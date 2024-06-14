// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 6/14/2024 using wmigen
//      Source root.StandardCimv2
//////////////////////////////////////////////
package standardcimv2

import (
	"github.com/microsoft/wmi/pkg/base/query"
	"github.com/microsoft/wmi/pkg/errors"
	cim "github.com/microsoft/wmi/pkg/wmiinstance"
	"reflect"
)

// MSFT_NetAdapterPowerManagementSettingData struct
type MSFT_NetAdapterPowerManagementSettingData struct {
	*MSFT_NetAdapterSettingData

	//
	AllowComputerToTurnOffDevice uint32

	//
	ArpOffload uint32

	//
	D0PacketCoalescing uint32

	//
	DeviceSleepOnDisconnect uint32

	//
	NSOffload uint32

	//
	OffloadParameters []MSFT_NetAdapterPowerManagement_Offload

	//
	RsnRekeyOffload uint32

	//
	SelectiveSuspend uint32

	//
	WakeOnMagicPacket uint32

	//
	WakeOnPattern uint32

	//
	WakePatterns []MSFT_NetAdapterPowerManagement_WakePattern
}

func NewMSFT_NetAdapterPowerManagementSettingDataEx1(instance *cim.WmiInstance) (newInstance *MSFT_NetAdapterPowerManagementSettingData, err error) {
	tmp, err := NewMSFT_NetAdapterSettingDataEx1(instance)

	if err != nil {
		return
	}
	newInstance = &MSFT_NetAdapterPowerManagementSettingData{
		MSFT_NetAdapterSettingData: tmp,
	}
	return
}

func NewMSFT_NetAdapterPowerManagementSettingDataEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *MSFT_NetAdapterPowerManagementSettingData, err error) {
	tmp, err := NewMSFT_NetAdapterSettingDataEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &MSFT_NetAdapterPowerManagementSettingData{
		MSFT_NetAdapterSettingData: tmp,
	}
	return
}

// SetAllowComputerToTurnOffDevice sets the value of AllowComputerToTurnOffDevice for the instance
func (instance *MSFT_NetAdapterPowerManagementSettingData) SetPropertyAllowComputerToTurnOffDevice(value uint32) (err error) {
	return instance.SetProperty("AllowComputerToTurnOffDevice", (value))
}

// GetAllowComputerToTurnOffDevice gets the value of AllowComputerToTurnOffDevice for the instance
func (instance *MSFT_NetAdapterPowerManagementSettingData) GetPropertyAllowComputerToTurnOffDevice() (value uint32, err error) {
	retValue, err := instance.GetProperty("AllowComputerToTurnOffDevice")
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

// SetArpOffload sets the value of ArpOffload for the instance
func (instance *MSFT_NetAdapterPowerManagementSettingData) SetPropertyArpOffload(value uint32) (err error) {
	return instance.SetProperty("ArpOffload", (value))
}

// GetArpOffload gets the value of ArpOffload for the instance
func (instance *MSFT_NetAdapterPowerManagementSettingData) GetPropertyArpOffload() (value uint32, err error) {
	retValue, err := instance.GetProperty("ArpOffload")
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

// SetD0PacketCoalescing sets the value of D0PacketCoalescing for the instance
func (instance *MSFT_NetAdapterPowerManagementSettingData) SetPropertyD0PacketCoalescing(value uint32) (err error) {
	return instance.SetProperty("D0PacketCoalescing", (value))
}

// GetD0PacketCoalescing gets the value of D0PacketCoalescing for the instance
func (instance *MSFT_NetAdapterPowerManagementSettingData) GetPropertyD0PacketCoalescing() (value uint32, err error) {
	retValue, err := instance.GetProperty("D0PacketCoalescing")
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

// SetDeviceSleepOnDisconnect sets the value of DeviceSleepOnDisconnect for the instance
func (instance *MSFT_NetAdapterPowerManagementSettingData) SetPropertyDeviceSleepOnDisconnect(value uint32) (err error) {
	return instance.SetProperty("DeviceSleepOnDisconnect", (value))
}

// GetDeviceSleepOnDisconnect gets the value of DeviceSleepOnDisconnect for the instance
func (instance *MSFT_NetAdapterPowerManagementSettingData) GetPropertyDeviceSleepOnDisconnect() (value uint32, err error) {
	retValue, err := instance.GetProperty("DeviceSleepOnDisconnect")
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

// SetNSOffload sets the value of NSOffload for the instance
func (instance *MSFT_NetAdapterPowerManagementSettingData) SetPropertyNSOffload(value uint32) (err error) {
	return instance.SetProperty("NSOffload", (value))
}

// GetNSOffload gets the value of NSOffload for the instance
func (instance *MSFT_NetAdapterPowerManagementSettingData) GetPropertyNSOffload() (value uint32, err error) {
	retValue, err := instance.GetProperty("NSOffload")
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

// SetOffloadParameters sets the value of OffloadParameters for the instance
func (instance *MSFT_NetAdapterPowerManagementSettingData) SetPropertyOffloadParameters(value []MSFT_NetAdapterPowerManagement_Offload) (err error) {
	return instance.SetProperty("OffloadParameters", (value))
}

// GetOffloadParameters gets the value of OffloadParameters for the instance
func (instance *MSFT_NetAdapterPowerManagementSettingData) GetPropertyOffloadParameters() (value []MSFT_NetAdapterPowerManagement_Offload, err error) {
	retValue, err := instance.GetProperty("OffloadParameters")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	for _, interfaceValue := range retValue.([]interface{}) {
		valuetmp, ok := interfaceValue.(MSFT_NetAdapterPowerManagement_Offload)
		if !ok {
			err = errors.Wrapf(errors.InvalidType, " MSFT_NetAdapterPowerManagement_Offload is Invalid. Expected %s", reflect.TypeOf(interfaceValue))
			return
		}
		value = append(value, MSFT_NetAdapterPowerManagement_Offload(valuetmp))
	}

	return
}

// SetRsnRekeyOffload sets the value of RsnRekeyOffload for the instance
func (instance *MSFT_NetAdapterPowerManagementSettingData) SetPropertyRsnRekeyOffload(value uint32) (err error) {
	return instance.SetProperty("RsnRekeyOffload", (value))
}

// GetRsnRekeyOffload gets the value of RsnRekeyOffload for the instance
func (instance *MSFT_NetAdapterPowerManagementSettingData) GetPropertyRsnRekeyOffload() (value uint32, err error) {
	retValue, err := instance.GetProperty("RsnRekeyOffload")
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

// SetSelectiveSuspend sets the value of SelectiveSuspend for the instance
func (instance *MSFT_NetAdapterPowerManagementSettingData) SetPropertySelectiveSuspend(value uint32) (err error) {
	return instance.SetProperty("SelectiveSuspend", (value))
}

// GetSelectiveSuspend gets the value of SelectiveSuspend for the instance
func (instance *MSFT_NetAdapterPowerManagementSettingData) GetPropertySelectiveSuspend() (value uint32, err error) {
	retValue, err := instance.GetProperty("SelectiveSuspend")
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

// SetWakeOnMagicPacket sets the value of WakeOnMagicPacket for the instance
func (instance *MSFT_NetAdapterPowerManagementSettingData) SetPropertyWakeOnMagicPacket(value uint32) (err error) {
	return instance.SetProperty("WakeOnMagicPacket", (value))
}

// GetWakeOnMagicPacket gets the value of WakeOnMagicPacket for the instance
func (instance *MSFT_NetAdapterPowerManagementSettingData) GetPropertyWakeOnMagicPacket() (value uint32, err error) {
	retValue, err := instance.GetProperty("WakeOnMagicPacket")
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

// SetWakeOnPattern sets the value of WakeOnPattern for the instance
func (instance *MSFT_NetAdapterPowerManagementSettingData) SetPropertyWakeOnPattern(value uint32) (err error) {
	return instance.SetProperty("WakeOnPattern", (value))
}

// GetWakeOnPattern gets the value of WakeOnPattern for the instance
func (instance *MSFT_NetAdapterPowerManagementSettingData) GetPropertyWakeOnPattern() (value uint32, err error) {
	retValue, err := instance.GetProperty("WakeOnPattern")
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

// SetWakePatterns sets the value of WakePatterns for the instance
func (instance *MSFT_NetAdapterPowerManagementSettingData) SetPropertyWakePatterns(value []MSFT_NetAdapterPowerManagement_WakePattern) (err error) {
	return instance.SetProperty("WakePatterns", (value))
}

// GetWakePatterns gets the value of WakePatterns for the instance
func (instance *MSFT_NetAdapterPowerManagementSettingData) GetPropertyWakePatterns() (value []MSFT_NetAdapterPowerManagement_WakePattern, err error) {
	retValue, err := instance.GetProperty("WakePatterns")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	for _, interfaceValue := range retValue.([]interface{}) {
		valuetmp, ok := interfaceValue.(MSFT_NetAdapterPowerManagement_WakePattern)
		if !ok {
			err = errors.Wrapf(errors.InvalidType, " MSFT_NetAdapterPowerManagement_WakePattern is Invalid. Expected %s", reflect.TypeOf(interfaceValue))
			return
		}
		value = append(value, MSFT_NetAdapterPowerManagement_WakePattern(valuetmp))
	}

	return
}

//

// <param name="ArpOffload" type="bool "></param>
// <param name="D0PacketCoalescing" type="bool "></param>
// <param name="DeviceSleepOnDisconnect" type="bool "></param>
// <param name="NSOffload" type="bool "></param>
// <param name="RsnRekeyOffload" type="bool "></param>
// <param name="SelectiveSuspend" type="bool "></param>
// <param name="WakeOnMagicPacket" type="bool "></param>
// <param name="WakeOnPattern" type="bool "></param>

// <param name="cmdletOutput" type="MSFT_NetAdapterPowerManagementSettingData "></param>
// <param name="ReturnValue" type="uint32 "></param>
func (instance *MSFT_NetAdapterPowerManagementSettingData) Enable( /* OUT */ cmdletOutput MSFT_NetAdapterPowerManagementSettingData,
	/* OPTIONAL IN */ ArpOffload bool,
	/* OPTIONAL IN */ D0PacketCoalescing bool,
	/* OPTIONAL IN */ DeviceSleepOnDisconnect bool,
	/* OPTIONAL IN */ NSOffload bool,
	/* OPTIONAL IN */ RsnRekeyOffload bool,
	/* OPTIONAL IN */ SelectiveSuspend bool,
	/* OPTIONAL IN */ WakeOnMagicPacket bool,
	/* OPTIONAL IN */ WakeOnPattern bool) (result uint32, err error) {
	retVal, err := instance.InvokeMethod("Enable", ArpOffload, D0PacketCoalescing, DeviceSleepOnDisconnect, NSOffload, RsnRekeyOffload, SelectiveSuspend, WakeOnMagicPacket, WakeOnPattern)
	if err != nil {
		return
	}
	retValue := retVal[0].(int32)
	result = uint32(retValue)
	return

}

//

// <param name="ArpOffload" type="bool "></param>
// <param name="D0PacketCoalescing" type="bool "></param>
// <param name="DeviceSleepOnDisconnect" type="bool "></param>
// <param name="NSOffload" type="bool "></param>
// <param name="RsnRekeyOffload" type="bool "></param>
// <param name="SelectiveSuspend" type="bool "></param>
// <param name="WakeOnMagicPacket" type="bool "></param>
// <param name="WakeOnPattern" type="bool "></param>

// <param name="cmdletOutput" type="MSFT_NetAdapterPowerManagementSettingData "></param>
// <param name="ReturnValue" type="uint32 "></param>
func (instance *MSFT_NetAdapterPowerManagementSettingData) Disable( /* OUT */ cmdletOutput MSFT_NetAdapterPowerManagementSettingData,
	/* OPTIONAL IN */ ArpOffload bool,
	/* OPTIONAL IN */ D0PacketCoalescing bool,
	/* OPTIONAL IN */ DeviceSleepOnDisconnect bool,
	/* OPTIONAL IN */ NSOffload bool,
	/* OPTIONAL IN */ RsnRekeyOffload bool,
	/* OPTIONAL IN */ SelectiveSuspend bool,
	/* OPTIONAL IN */ WakeOnMagicPacket bool,
	/* OPTIONAL IN */ WakeOnPattern bool) (result uint32, err error) {
	retVal, err := instance.InvokeMethod("Disable", ArpOffload, D0PacketCoalescing, DeviceSleepOnDisconnect, NSOffload, RsnRekeyOffload, SelectiveSuspend, WakeOnMagicPacket, WakeOnPattern)
	if err != nil {
		return
	}
	retValue := retVal[0].(int32)
	result = uint32(retValue)
	return

}
