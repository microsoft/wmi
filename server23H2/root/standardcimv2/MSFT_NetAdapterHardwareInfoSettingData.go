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

// MSFT_NetAdapterHardwareInfoSettingData struct
type MSFT_NetAdapterHardwareInfoSettingData struct {
	*MSFT_NetAdapterSettingData

	//
	BusNumber uint32

	//
	DeviceNumber uint32

	//
	Dma64BitSupported bool

	//
	FunctionNumber uint32

	//
	LineBasedInterrupts bool

	//
	LineBasedInterruptSupported bool

	//
	LocationInformationString string

	//
	MaxInterruptMessages uint32

	//
	MsiEnabled bool

	//
	MsiInterruptSupported bool

	//
	MsiSupported bool

	//
	MsiXEnabled bool

	//
	MsiXInterruptSupported bool

	//
	MsixMessageAffinityArray []MSFT_NetAdapter_Group_Affinity

	//
	MsiXSupported bool

	//
	NoInterrupt bool

	//
	NumaNode uint16

	//
	NumMsiMessages uint32

	//
	NumMsixTableEntries uint32

	//
	PciCurrentSpeedAndMode uint32

	//
	PciDeviceLabelID uint32

	//
	PciDeviceLabelString string

	//
	PciDeviceType uint32

	//
	PciExpressCurrentLinkSpeedEncoded uint32

	//
	PciExpressCurrentLinkWidth uint32

	//
	PciExpressCurrentPayloadSize uint32

	//
	PciExpressMaxLinkSpeedEncoded uint32

	//
	PciExpressMaxLinkWidth uint32

	//
	PciExpressMaxPayloadSize uint32

	//
	PciExpressMaxReadRequestSize uint32

	//
	PciExpressVersion uint32

	//
	PciXCurrentSpeedAndMode uint32

	//
	S0WakeupSupported bool

	//
	SegmentNumber uint32

	//
	SlotNumber uint32

	//
	SriovSupport uint32
}

func NewMSFT_NetAdapterHardwareInfoSettingDataEx1(instance *cim.WmiInstance) (newInstance *MSFT_NetAdapterHardwareInfoSettingData, err error) {
	tmp, err := NewMSFT_NetAdapterSettingDataEx1(instance)

	if err != nil {
		return
	}
	newInstance = &MSFT_NetAdapterHardwareInfoSettingData{
		MSFT_NetAdapterSettingData: tmp,
	}
	return
}

func NewMSFT_NetAdapterHardwareInfoSettingDataEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *MSFT_NetAdapterHardwareInfoSettingData, err error) {
	tmp, err := NewMSFT_NetAdapterSettingDataEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &MSFT_NetAdapterHardwareInfoSettingData{
		MSFT_NetAdapterSettingData: tmp,
	}
	return
}

// SetBusNumber sets the value of BusNumber for the instance
func (instance *MSFT_NetAdapterHardwareInfoSettingData) SetPropertyBusNumber(value uint32) (err error) {
	return instance.SetProperty("BusNumber", (value))
}

// GetBusNumber gets the value of BusNumber for the instance
func (instance *MSFT_NetAdapterHardwareInfoSettingData) GetPropertyBusNumber() (value uint32, err error) {
	retValue, err := instance.GetProperty("BusNumber")
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

// SetDeviceNumber sets the value of DeviceNumber for the instance
func (instance *MSFT_NetAdapterHardwareInfoSettingData) SetPropertyDeviceNumber(value uint32) (err error) {
	return instance.SetProperty("DeviceNumber", (value))
}

// GetDeviceNumber gets the value of DeviceNumber for the instance
func (instance *MSFT_NetAdapterHardwareInfoSettingData) GetPropertyDeviceNumber() (value uint32, err error) {
	retValue, err := instance.GetProperty("DeviceNumber")
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

// SetDma64BitSupported sets the value of Dma64BitSupported for the instance
func (instance *MSFT_NetAdapterHardwareInfoSettingData) SetPropertyDma64BitSupported(value bool) (err error) {
	return instance.SetProperty("Dma64BitSupported", (value))
}

// GetDma64BitSupported gets the value of Dma64BitSupported for the instance
func (instance *MSFT_NetAdapterHardwareInfoSettingData) GetPropertyDma64BitSupported() (value bool, err error) {
	retValue, err := instance.GetProperty("Dma64BitSupported")
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

// SetFunctionNumber sets the value of FunctionNumber for the instance
func (instance *MSFT_NetAdapterHardwareInfoSettingData) SetPropertyFunctionNumber(value uint32) (err error) {
	return instance.SetProperty("FunctionNumber", (value))
}

// GetFunctionNumber gets the value of FunctionNumber for the instance
func (instance *MSFT_NetAdapterHardwareInfoSettingData) GetPropertyFunctionNumber() (value uint32, err error) {
	retValue, err := instance.GetProperty("FunctionNumber")
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

// SetLineBasedInterrupts sets the value of LineBasedInterrupts for the instance
func (instance *MSFT_NetAdapterHardwareInfoSettingData) SetPropertyLineBasedInterrupts(value bool) (err error) {
	return instance.SetProperty("LineBasedInterrupts", (value))
}

// GetLineBasedInterrupts gets the value of LineBasedInterrupts for the instance
func (instance *MSFT_NetAdapterHardwareInfoSettingData) GetPropertyLineBasedInterrupts() (value bool, err error) {
	retValue, err := instance.GetProperty("LineBasedInterrupts")
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

// SetLineBasedInterruptSupported sets the value of LineBasedInterruptSupported for the instance
func (instance *MSFT_NetAdapterHardwareInfoSettingData) SetPropertyLineBasedInterruptSupported(value bool) (err error) {
	return instance.SetProperty("LineBasedInterruptSupported", (value))
}

// GetLineBasedInterruptSupported gets the value of LineBasedInterruptSupported for the instance
func (instance *MSFT_NetAdapterHardwareInfoSettingData) GetPropertyLineBasedInterruptSupported() (value bool, err error) {
	retValue, err := instance.GetProperty("LineBasedInterruptSupported")
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

// SetLocationInformationString sets the value of LocationInformationString for the instance
func (instance *MSFT_NetAdapterHardwareInfoSettingData) SetPropertyLocationInformationString(value string) (err error) {
	return instance.SetProperty("LocationInformationString", (value))
}

// GetLocationInformationString gets the value of LocationInformationString for the instance
func (instance *MSFT_NetAdapterHardwareInfoSettingData) GetPropertyLocationInformationString() (value string, err error) {
	retValue, err := instance.GetProperty("LocationInformationString")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(string)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " string is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = string(valuetmp)

	return
}

// SetMaxInterruptMessages sets the value of MaxInterruptMessages for the instance
func (instance *MSFT_NetAdapterHardwareInfoSettingData) SetPropertyMaxInterruptMessages(value uint32) (err error) {
	return instance.SetProperty("MaxInterruptMessages", (value))
}

// GetMaxInterruptMessages gets the value of MaxInterruptMessages for the instance
func (instance *MSFT_NetAdapterHardwareInfoSettingData) GetPropertyMaxInterruptMessages() (value uint32, err error) {
	retValue, err := instance.GetProperty("MaxInterruptMessages")
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

// SetMsiEnabled sets the value of MsiEnabled for the instance
func (instance *MSFT_NetAdapterHardwareInfoSettingData) SetPropertyMsiEnabled(value bool) (err error) {
	return instance.SetProperty("MsiEnabled", (value))
}

// GetMsiEnabled gets the value of MsiEnabled for the instance
func (instance *MSFT_NetAdapterHardwareInfoSettingData) GetPropertyMsiEnabled() (value bool, err error) {
	retValue, err := instance.GetProperty("MsiEnabled")
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

// SetMsiInterruptSupported sets the value of MsiInterruptSupported for the instance
func (instance *MSFT_NetAdapterHardwareInfoSettingData) SetPropertyMsiInterruptSupported(value bool) (err error) {
	return instance.SetProperty("MsiInterruptSupported", (value))
}

// GetMsiInterruptSupported gets the value of MsiInterruptSupported for the instance
func (instance *MSFT_NetAdapterHardwareInfoSettingData) GetPropertyMsiInterruptSupported() (value bool, err error) {
	retValue, err := instance.GetProperty("MsiInterruptSupported")
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

// SetMsiSupported sets the value of MsiSupported for the instance
func (instance *MSFT_NetAdapterHardwareInfoSettingData) SetPropertyMsiSupported(value bool) (err error) {
	return instance.SetProperty("MsiSupported", (value))
}

// GetMsiSupported gets the value of MsiSupported for the instance
func (instance *MSFT_NetAdapterHardwareInfoSettingData) GetPropertyMsiSupported() (value bool, err error) {
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
func (instance *MSFT_NetAdapterHardwareInfoSettingData) SetPropertyMsiXEnabled(value bool) (err error) {
	return instance.SetProperty("MsiXEnabled", (value))
}

// GetMsiXEnabled gets the value of MsiXEnabled for the instance
func (instance *MSFT_NetAdapterHardwareInfoSettingData) GetPropertyMsiXEnabled() (value bool, err error) {
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

// SetMsiXInterruptSupported sets the value of MsiXInterruptSupported for the instance
func (instance *MSFT_NetAdapterHardwareInfoSettingData) SetPropertyMsiXInterruptSupported(value bool) (err error) {
	return instance.SetProperty("MsiXInterruptSupported", (value))
}

// GetMsiXInterruptSupported gets the value of MsiXInterruptSupported for the instance
func (instance *MSFT_NetAdapterHardwareInfoSettingData) GetPropertyMsiXInterruptSupported() (value bool, err error) {
	retValue, err := instance.GetProperty("MsiXInterruptSupported")
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

// SetMsixMessageAffinityArray sets the value of MsixMessageAffinityArray for the instance
func (instance *MSFT_NetAdapterHardwareInfoSettingData) SetPropertyMsixMessageAffinityArray(value []MSFT_NetAdapter_Group_Affinity) (err error) {
	return instance.SetProperty("MsixMessageAffinityArray", (value))
}

// GetMsixMessageAffinityArray gets the value of MsixMessageAffinityArray for the instance
func (instance *MSFT_NetAdapterHardwareInfoSettingData) GetPropertyMsixMessageAffinityArray() (value []MSFT_NetAdapter_Group_Affinity, err error) {
	retValue, err := instance.GetProperty("MsixMessageAffinityArray")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	for _, interfaceValue := range retValue.([]interface{}) {
		valuetmp, ok := interfaceValue.(MSFT_NetAdapter_Group_Affinity)
		if !ok {
			err = errors.Wrapf(errors.InvalidType, " MSFT_NetAdapter_Group_Affinity is Invalid. Expected %s", reflect.TypeOf(interfaceValue))
			return
		}
		value = append(value, MSFT_NetAdapter_Group_Affinity(valuetmp))
	}

	return
}

// SetMsiXSupported sets the value of MsiXSupported for the instance
func (instance *MSFT_NetAdapterHardwareInfoSettingData) SetPropertyMsiXSupported(value bool) (err error) {
	return instance.SetProperty("MsiXSupported", (value))
}

// GetMsiXSupported gets the value of MsiXSupported for the instance
func (instance *MSFT_NetAdapterHardwareInfoSettingData) GetPropertyMsiXSupported() (value bool, err error) {
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

// SetNoInterrupt sets the value of NoInterrupt for the instance
func (instance *MSFT_NetAdapterHardwareInfoSettingData) SetPropertyNoInterrupt(value bool) (err error) {
	return instance.SetProperty("NoInterrupt", (value))
}

// GetNoInterrupt gets the value of NoInterrupt for the instance
func (instance *MSFT_NetAdapterHardwareInfoSettingData) GetPropertyNoInterrupt() (value bool, err error) {
	retValue, err := instance.GetProperty("NoInterrupt")
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
func (instance *MSFT_NetAdapterHardwareInfoSettingData) SetPropertyNumaNode(value uint16) (err error) {
	return instance.SetProperty("NumaNode", (value))
}

// GetNumaNode gets the value of NumaNode for the instance
func (instance *MSFT_NetAdapterHardwareInfoSettingData) GetPropertyNumaNode() (value uint16, err error) {
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

// SetNumMsiMessages sets the value of NumMsiMessages for the instance
func (instance *MSFT_NetAdapterHardwareInfoSettingData) SetPropertyNumMsiMessages(value uint32) (err error) {
	return instance.SetProperty("NumMsiMessages", (value))
}

// GetNumMsiMessages gets the value of NumMsiMessages for the instance
func (instance *MSFT_NetAdapterHardwareInfoSettingData) GetPropertyNumMsiMessages() (value uint32, err error) {
	retValue, err := instance.GetProperty("NumMsiMessages")
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

// SetNumMsixTableEntries sets the value of NumMsixTableEntries for the instance
func (instance *MSFT_NetAdapterHardwareInfoSettingData) SetPropertyNumMsixTableEntries(value uint32) (err error) {
	return instance.SetProperty("NumMsixTableEntries", (value))
}

// GetNumMsixTableEntries gets the value of NumMsixTableEntries for the instance
func (instance *MSFT_NetAdapterHardwareInfoSettingData) GetPropertyNumMsixTableEntries() (value uint32, err error) {
	retValue, err := instance.GetProperty("NumMsixTableEntries")
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

// SetPciCurrentSpeedAndMode sets the value of PciCurrentSpeedAndMode for the instance
func (instance *MSFT_NetAdapterHardwareInfoSettingData) SetPropertyPciCurrentSpeedAndMode(value uint32) (err error) {
	return instance.SetProperty("PciCurrentSpeedAndMode", (value))
}

// GetPciCurrentSpeedAndMode gets the value of PciCurrentSpeedAndMode for the instance
func (instance *MSFT_NetAdapterHardwareInfoSettingData) GetPropertyPciCurrentSpeedAndMode() (value uint32, err error) {
	retValue, err := instance.GetProperty("PciCurrentSpeedAndMode")
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

// SetPciDeviceLabelID sets the value of PciDeviceLabelID for the instance
func (instance *MSFT_NetAdapterHardwareInfoSettingData) SetPropertyPciDeviceLabelID(value uint32) (err error) {
	return instance.SetProperty("PciDeviceLabelID", (value))
}

// GetPciDeviceLabelID gets the value of PciDeviceLabelID for the instance
func (instance *MSFT_NetAdapterHardwareInfoSettingData) GetPropertyPciDeviceLabelID() (value uint32, err error) {
	retValue, err := instance.GetProperty("PciDeviceLabelID")
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

// SetPciDeviceLabelString sets the value of PciDeviceLabelString for the instance
func (instance *MSFT_NetAdapterHardwareInfoSettingData) SetPropertyPciDeviceLabelString(value string) (err error) {
	return instance.SetProperty("PciDeviceLabelString", (value))
}

// GetPciDeviceLabelString gets the value of PciDeviceLabelString for the instance
func (instance *MSFT_NetAdapterHardwareInfoSettingData) GetPropertyPciDeviceLabelString() (value string, err error) {
	retValue, err := instance.GetProperty("PciDeviceLabelString")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(string)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " string is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = string(valuetmp)

	return
}

// SetPciDeviceType sets the value of PciDeviceType for the instance
func (instance *MSFT_NetAdapterHardwareInfoSettingData) SetPropertyPciDeviceType(value uint32) (err error) {
	return instance.SetProperty("PciDeviceType", (value))
}

// GetPciDeviceType gets the value of PciDeviceType for the instance
func (instance *MSFT_NetAdapterHardwareInfoSettingData) GetPropertyPciDeviceType() (value uint32, err error) {
	retValue, err := instance.GetProperty("PciDeviceType")
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

// SetPciExpressCurrentLinkSpeedEncoded sets the value of PciExpressCurrentLinkSpeedEncoded for the instance
func (instance *MSFT_NetAdapterHardwareInfoSettingData) SetPropertyPciExpressCurrentLinkSpeedEncoded(value uint32) (err error) {
	return instance.SetProperty("PciExpressCurrentLinkSpeedEncoded", (value))
}

// GetPciExpressCurrentLinkSpeedEncoded gets the value of PciExpressCurrentLinkSpeedEncoded for the instance
func (instance *MSFT_NetAdapterHardwareInfoSettingData) GetPropertyPciExpressCurrentLinkSpeedEncoded() (value uint32, err error) {
	retValue, err := instance.GetProperty("PciExpressCurrentLinkSpeedEncoded")
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

// SetPciExpressCurrentLinkWidth sets the value of PciExpressCurrentLinkWidth for the instance
func (instance *MSFT_NetAdapterHardwareInfoSettingData) SetPropertyPciExpressCurrentLinkWidth(value uint32) (err error) {
	return instance.SetProperty("PciExpressCurrentLinkWidth", (value))
}

// GetPciExpressCurrentLinkWidth gets the value of PciExpressCurrentLinkWidth for the instance
func (instance *MSFT_NetAdapterHardwareInfoSettingData) GetPropertyPciExpressCurrentLinkWidth() (value uint32, err error) {
	retValue, err := instance.GetProperty("PciExpressCurrentLinkWidth")
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

// SetPciExpressCurrentPayloadSize sets the value of PciExpressCurrentPayloadSize for the instance
func (instance *MSFT_NetAdapterHardwareInfoSettingData) SetPropertyPciExpressCurrentPayloadSize(value uint32) (err error) {
	return instance.SetProperty("PciExpressCurrentPayloadSize", (value))
}

// GetPciExpressCurrentPayloadSize gets the value of PciExpressCurrentPayloadSize for the instance
func (instance *MSFT_NetAdapterHardwareInfoSettingData) GetPropertyPciExpressCurrentPayloadSize() (value uint32, err error) {
	retValue, err := instance.GetProperty("PciExpressCurrentPayloadSize")
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

// SetPciExpressMaxLinkSpeedEncoded sets the value of PciExpressMaxLinkSpeedEncoded for the instance
func (instance *MSFT_NetAdapterHardwareInfoSettingData) SetPropertyPciExpressMaxLinkSpeedEncoded(value uint32) (err error) {
	return instance.SetProperty("PciExpressMaxLinkSpeedEncoded", (value))
}

// GetPciExpressMaxLinkSpeedEncoded gets the value of PciExpressMaxLinkSpeedEncoded for the instance
func (instance *MSFT_NetAdapterHardwareInfoSettingData) GetPropertyPciExpressMaxLinkSpeedEncoded() (value uint32, err error) {
	retValue, err := instance.GetProperty("PciExpressMaxLinkSpeedEncoded")
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

// SetPciExpressMaxLinkWidth sets the value of PciExpressMaxLinkWidth for the instance
func (instance *MSFT_NetAdapterHardwareInfoSettingData) SetPropertyPciExpressMaxLinkWidth(value uint32) (err error) {
	return instance.SetProperty("PciExpressMaxLinkWidth", (value))
}

// GetPciExpressMaxLinkWidth gets the value of PciExpressMaxLinkWidth for the instance
func (instance *MSFT_NetAdapterHardwareInfoSettingData) GetPropertyPciExpressMaxLinkWidth() (value uint32, err error) {
	retValue, err := instance.GetProperty("PciExpressMaxLinkWidth")
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

// SetPciExpressMaxPayloadSize sets the value of PciExpressMaxPayloadSize for the instance
func (instance *MSFT_NetAdapterHardwareInfoSettingData) SetPropertyPciExpressMaxPayloadSize(value uint32) (err error) {
	return instance.SetProperty("PciExpressMaxPayloadSize", (value))
}

// GetPciExpressMaxPayloadSize gets the value of PciExpressMaxPayloadSize for the instance
func (instance *MSFT_NetAdapterHardwareInfoSettingData) GetPropertyPciExpressMaxPayloadSize() (value uint32, err error) {
	retValue, err := instance.GetProperty("PciExpressMaxPayloadSize")
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

// SetPciExpressMaxReadRequestSize sets the value of PciExpressMaxReadRequestSize for the instance
func (instance *MSFT_NetAdapterHardwareInfoSettingData) SetPropertyPciExpressMaxReadRequestSize(value uint32) (err error) {
	return instance.SetProperty("PciExpressMaxReadRequestSize", (value))
}

// GetPciExpressMaxReadRequestSize gets the value of PciExpressMaxReadRequestSize for the instance
func (instance *MSFT_NetAdapterHardwareInfoSettingData) GetPropertyPciExpressMaxReadRequestSize() (value uint32, err error) {
	retValue, err := instance.GetProperty("PciExpressMaxReadRequestSize")
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

// SetPciExpressVersion sets the value of PciExpressVersion for the instance
func (instance *MSFT_NetAdapterHardwareInfoSettingData) SetPropertyPciExpressVersion(value uint32) (err error) {
	return instance.SetProperty("PciExpressVersion", (value))
}

// GetPciExpressVersion gets the value of PciExpressVersion for the instance
func (instance *MSFT_NetAdapterHardwareInfoSettingData) GetPropertyPciExpressVersion() (value uint32, err error) {
	retValue, err := instance.GetProperty("PciExpressVersion")
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

// SetPciXCurrentSpeedAndMode sets the value of PciXCurrentSpeedAndMode for the instance
func (instance *MSFT_NetAdapterHardwareInfoSettingData) SetPropertyPciXCurrentSpeedAndMode(value uint32) (err error) {
	return instance.SetProperty("PciXCurrentSpeedAndMode", (value))
}

// GetPciXCurrentSpeedAndMode gets the value of PciXCurrentSpeedAndMode for the instance
func (instance *MSFT_NetAdapterHardwareInfoSettingData) GetPropertyPciXCurrentSpeedAndMode() (value uint32, err error) {
	retValue, err := instance.GetProperty("PciXCurrentSpeedAndMode")
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

// SetS0WakeupSupported sets the value of S0WakeupSupported for the instance
func (instance *MSFT_NetAdapterHardwareInfoSettingData) SetPropertyS0WakeupSupported(value bool) (err error) {
	return instance.SetProperty("S0WakeupSupported", (value))
}

// GetS0WakeupSupported gets the value of S0WakeupSupported for the instance
func (instance *MSFT_NetAdapterHardwareInfoSettingData) GetPropertyS0WakeupSupported() (value bool, err error) {
	retValue, err := instance.GetProperty("S0WakeupSupported")
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

// SetSegmentNumber sets the value of SegmentNumber for the instance
func (instance *MSFT_NetAdapterHardwareInfoSettingData) SetPropertySegmentNumber(value uint32) (err error) {
	return instance.SetProperty("SegmentNumber", (value))
}

// GetSegmentNumber gets the value of SegmentNumber for the instance
func (instance *MSFT_NetAdapterHardwareInfoSettingData) GetPropertySegmentNumber() (value uint32, err error) {
	retValue, err := instance.GetProperty("SegmentNumber")
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

// SetSlotNumber sets the value of SlotNumber for the instance
func (instance *MSFT_NetAdapterHardwareInfoSettingData) SetPropertySlotNumber(value uint32) (err error) {
	return instance.SetProperty("SlotNumber", (value))
}

// GetSlotNumber gets the value of SlotNumber for the instance
func (instance *MSFT_NetAdapterHardwareInfoSettingData) GetPropertySlotNumber() (value uint32, err error) {
	retValue, err := instance.GetProperty("SlotNumber")
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

// SetSriovSupport sets the value of SriovSupport for the instance
func (instance *MSFT_NetAdapterHardwareInfoSettingData) SetPropertySriovSupport(value uint32) (err error) {
	return instance.SetProperty("SriovSupport", (value))
}

// GetSriovSupport gets the value of SriovSupport for the instance
func (instance *MSFT_NetAdapterHardwareInfoSettingData) GetPropertySriovSupport() (value uint32, err error) {
	retValue, err := instance.GetProperty("SriovSupport")
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
