// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 6/12/2024 using wmigen
//      Source root.WMI
//////////////////////////////////////////////
package wmi

import (
	"github.com/microsoft/wmi/pkg/base/query"
	"github.com/microsoft/wmi/pkg/errors"
	cim "github.com/microsoft/wmi/pkg/wmiinstance"
	"reflect"
)

// MSSerial_CommInfo struct
type MSSerial_CommInfo struct {
	*MSSerial

	//
	Active bool

	//
	BaudRate uint32

	//
	BitsPerByte uint32

	//
	InstanceName string

	//
	IsBusy bool

	//
	MaximumBaudRate uint32

	//
	MaximumInputBufferSize uint32

	//
	MaximumOutputBufferSize uint32

	//
	Parity uint32

	//
	ParityCheckEnable bool

	//
	SettableBaudRate bool

	//
	SettableDataBits bool

	//
	SettableFlowControl bool

	//
	SettableParity bool

	//
	SettableParityCheck bool

	//
	SettableStopBits bool

	//
	StopBits uint32

	//
	Support16BitMode bool

	//
	SupportDTRDSR bool

	//
	SupportIntervalTimeouts bool

	//
	SupportParityCheck bool

	//
	SupportRTSCTS bool

	//
	SupportXonXoff bool

	//
	XoffCharacter uint32

	//
	XoffXmitThreshold uint32

	//
	XonCharacter uint32

	//
	XonXmitThreshold uint32
}

func NewMSSerial_CommInfoEx1(instance *cim.WmiInstance) (newInstance *MSSerial_CommInfo, err error) {
	tmp, err := NewMSSerialEx1(instance)

	if err != nil {
		return
	}
	newInstance = &MSSerial_CommInfo{
		MSSerial: tmp,
	}
	return
}

func NewMSSerial_CommInfoEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *MSSerial_CommInfo, err error) {
	tmp, err := NewMSSerialEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &MSSerial_CommInfo{
		MSSerial: tmp,
	}
	return
}

// SetActive sets the value of Active for the instance
func (instance *MSSerial_CommInfo) SetPropertyActive(value bool) (err error) {
	return instance.SetProperty("Active", (value))
}

// GetActive gets the value of Active for the instance
func (instance *MSSerial_CommInfo) GetPropertyActive() (value bool, err error) {
	retValue, err := instance.GetProperty("Active")
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

// SetBaudRate sets the value of BaudRate for the instance
func (instance *MSSerial_CommInfo) SetPropertyBaudRate(value uint32) (err error) {
	return instance.SetProperty("BaudRate", (value))
}

// GetBaudRate gets the value of BaudRate for the instance
func (instance *MSSerial_CommInfo) GetPropertyBaudRate() (value uint32, err error) {
	retValue, err := instance.GetProperty("BaudRate")
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

// SetBitsPerByte sets the value of BitsPerByte for the instance
func (instance *MSSerial_CommInfo) SetPropertyBitsPerByte(value uint32) (err error) {
	return instance.SetProperty("BitsPerByte", (value))
}

// GetBitsPerByte gets the value of BitsPerByte for the instance
func (instance *MSSerial_CommInfo) GetPropertyBitsPerByte() (value uint32, err error) {
	retValue, err := instance.GetProperty("BitsPerByte")
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

// SetInstanceName sets the value of InstanceName for the instance
func (instance *MSSerial_CommInfo) SetPropertyInstanceName(value string) (err error) {
	return instance.SetProperty("InstanceName", (value))
}

// GetInstanceName gets the value of InstanceName for the instance
func (instance *MSSerial_CommInfo) GetPropertyInstanceName() (value string, err error) {
	retValue, err := instance.GetProperty("InstanceName")
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

// SetIsBusy sets the value of IsBusy for the instance
func (instance *MSSerial_CommInfo) SetPropertyIsBusy(value bool) (err error) {
	return instance.SetProperty("IsBusy", (value))
}

// GetIsBusy gets the value of IsBusy for the instance
func (instance *MSSerial_CommInfo) GetPropertyIsBusy() (value bool, err error) {
	retValue, err := instance.GetProperty("IsBusy")
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

// SetMaximumBaudRate sets the value of MaximumBaudRate for the instance
func (instance *MSSerial_CommInfo) SetPropertyMaximumBaudRate(value uint32) (err error) {
	return instance.SetProperty("MaximumBaudRate", (value))
}

// GetMaximumBaudRate gets the value of MaximumBaudRate for the instance
func (instance *MSSerial_CommInfo) GetPropertyMaximumBaudRate() (value uint32, err error) {
	retValue, err := instance.GetProperty("MaximumBaudRate")
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

// SetMaximumInputBufferSize sets the value of MaximumInputBufferSize for the instance
func (instance *MSSerial_CommInfo) SetPropertyMaximumInputBufferSize(value uint32) (err error) {
	return instance.SetProperty("MaximumInputBufferSize", (value))
}

// GetMaximumInputBufferSize gets the value of MaximumInputBufferSize for the instance
func (instance *MSSerial_CommInfo) GetPropertyMaximumInputBufferSize() (value uint32, err error) {
	retValue, err := instance.GetProperty("MaximumInputBufferSize")
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

// SetMaximumOutputBufferSize sets the value of MaximumOutputBufferSize for the instance
func (instance *MSSerial_CommInfo) SetPropertyMaximumOutputBufferSize(value uint32) (err error) {
	return instance.SetProperty("MaximumOutputBufferSize", (value))
}

// GetMaximumOutputBufferSize gets the value of MaximumOutputBufferSize for the instance
func (instance *MSSerial_CommInfo) GetPropertyMaximumOutputBufferSize() (value uint32, err error) {
	retValue, err := instance.GetProperty("MaximumOutputBufferSize")
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

// SetParity sets the value of Parity for the instance
func (instance *MSSerial_CommInfo) SetPropertyParity(value uint32) (err error) {
	return instance.SetProperty("Parity", (value))
}

// GetParity gets the value of Parity for the instance
func (instance *MSSerial_CommInfo) GetPropertyParity() (value uint32, err error) {
	retValue, err := instance.GetProperty("Parity")
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

// SetParityCheckEnable sets the value of ParityCheckEnable for the instance
func (instance *MSSerial_CommInfo) SetPropertyParityCheckEnable(value bool) (err error) {
	return instance.SetProperty("ParityCheckEnable", (value))
}

// GetParityCheckEnable gets the value of ParityCheckEnable for the instance
func (instance *MSSerial_CommInfo) GetPropertyParityCheckEnable() (value bool, err error) {
	retValue, err := instance.GetProperty("ParityCheckEnable")
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

// SetSettableBaudRate sets the value of SettableBaudRate for the instance
func (instance *MSSerial_CommInfo) SetPropertySettableBaudRate(value bool) (err error) {
	return instance.SetProperty("SettableBaudRate", (value))
}

// GetSettableBaudRate gets the value of SettableBaudRate for the instance
func (instance *MSSerial_CommInfo) GetPropertySettableBaudRate() (value bool, err error) {
	retValue, err := instance.GetProperty("SettableBaudRate")
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

// SetSettableDataBits sets the value of SettableDataBits for the instance
func (instance *MSSerial_CommInfo) SetPropertySettableDataBits(value bool) (err error) {
	return instance.SetProperty("SettableDataBits", (value))
}

// GetSettableDataBits gets the value of SettableDataBits for the instance
func (instance *MSSerial_CommInfo) GetPropertySettableDataBits() (value bool, err error) {
	retValue, err := instance.GetProperty("SettableDataBits")
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

// SetSettableFlowControl sets the value of SettableFlowControl for the instance
func (instance *MSSerial_CommInfo) SetPropertySettableFlowControl(value bool) (err error) {
	return instance.SetProperty("SettableFlowControl", (value))
}

// GetSettableFlowControl gets the value of SettableFlowControl for the instance
func (instance *MSSerial_CommInfo) GetPropertySettableFlowControl() (value bool, err error) {
	retValue, err := instance.GetProperty("SettableFlowControl")
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

// SetSettableParity sets the value of SettableParity for the instance
func (instance *MSSerial_CommInfo) SetPropertySettableParity(value bool) (err error) {
	return instance.SetProperty("SettableParity", (value))
}

// GetSettableParity gets the value of SettableParity for the instance
func (instance *MSSerial_CommInfo) GetPropertySettableParity() (value bool, err error) {
	retValue, err := instance.GetProperty("SettableParity")
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

// SetSettableParityCheck sets the value of SettableParityCheck for the instance
func (instance *MSSerial_CommInfo) SetPropertySettableParityCheck(value bool) (err error) {
	return instance.SetProperty("SettableParityCheck", (value))
}

// GetSettableParityCheck gets the value of SettableParityCheck for the instance
func (instance *MSSerial_CommInfo) GetPropertySettableParityCheck() (value bool, err error) {
	retValue, err := instance.GetProperty("SettableParityCheck")
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

// SetSettableStopBits sets the value of SettableStopBits for the instance
func (instance *MSSerial_CommInfo) SetPropertySettableStopBits(value bool) (err error) {
	return instance.SetProperty("SettableStopBits", (value))
}

// GetSettableStopBits gets the value of SettableStopBits for the instance
func (instance *MSSerial_CommInfo) GetPropertySettableStopBits() (value bool, err error) {
	retValue, err := instance.GetProperty("SettableStopBits")
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

// SetStopBits sets the value of StopBits for the instance
func (instance *MSSerial_CommInfo) SetPropertyStopBits(value uint32) (err error) {
	return instance.SetProperty("StopBits", (value))
}

// GetStopBits gets the value of StopBits for the instance
func (instance *MSSerial_CommInfo) GetPropertyStopBits() (value uint32, err error) {
	retValue, err := instance.GetProperty("StopBits")
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

// SetSupport16BitMode sets the value of Support16BitMode for the instance
func (instance *MSSerial_CommInfo) SetPropertySupport16BitMode(value bool) (err error) {
	return instance.SetProperty("Support16BitMode", (value))
}

// GetSupport16BitMode gets the value of Support16BitMode for the instance
func (instance *MSSerial_CommInfo) GetPropertySupport16BitMode() (value bool, err error) {
	retValue, err := instance.GetProperty("Support16BitMode")
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

// SetSupportDTRDSR sets the value of SupportDTRDSR for the instance
func (instance *MSSerial_CommInfo) SetPropertySupportDTRDSR(value bool) (err error) {
	return instance.SetProperty("SupportDTRDSR", (value))
}

// GetSupportDTRDSR gets the value of SupportDTRDSR for the instance
func (instance *MSSerial_CommInfo) GetPropertySupportDTRDSR() (value bool, err error) {
	retValue, err := instance.GetProperty("SupportDTRDSR")
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

// SetSupportIntervalTimeouts sets the value of SupportIntervalTimeouts for the instance
func (instance *MSSerial_CommInfo) SetPropertySupportIntervalTimeouts(value bool) (err error) {
	return instance.SetProperty("SupportIntervalTimeouts", (value))
}

// GetSupportIntervalTimeouts gets the value of SupportIntervalTimeouts for the instance
func (instance *MSSerial_CommInfo) GetPropertySupportIntervalTimeouts() (value bool, err error) {
	retValue, err := instance.GetProperty("SupportIntervalTimeouts")
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

// SetSupportParityCheck sets the value of SupportParityCheck for the instance
func (instance *MSSerial_CommInfo) SetPropertySupportParityCheck(value bool) (err error) {
	return instance.SetProperty("SupportParityCheck", (value))
}

// GetSupportParityCheck gets the value of SupportParityCheck for the instance
func (instance *MSSerial_CommInfo) GetPropertySupportParityCheck() (value bool, err error) {
	retValue, err := instance.GetProperty("SupportParityCheck")
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

// SetSupportRTSCTS sets the value of SupportRTSCTS for the instance
func (instance *MSSerial_CommInfo) SetPropertySupportRTSCTS(value bool) (err error) {
	return instance.SetProperty("SupportRTSCTS", (value))
}

// GetSupportRTSCTS gets the value of SupportRTSCTS for the instance
func (instance *MSSerial_CommInfo) GetPropertySupportRTSCTS() (value bool, err error) {
	retValue, err := instance.GetProperty("SupportRTSCTS")
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

// SetSupportXonXoff sets the value of SupportXonXoff for the instance
func (instance *MSSerial_CommInfo) SetPropertySupportXonXoff(value bool) (err error) {
	return instance.SetProperty("SupportXonXoff", (value))
}

// GetSupportXonXoff gets the value of SupportXonXoff for the instance
func (instance *MSSerial_CommInfo) GetPropertySupportXonXoff() (value bool, err error) {
	retValue, err := instance.GetProperty("SupportXonXoff")
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

// SetXoffCharacter sets the value of XoffCharacter for the instance
func (instance *MSSerial_CommInfo) SetPropertyXoffCharacter(value uint32) (err error) {
	return instance.SetProperty("XoffCharacter", (value))
}

// GetXoffCharacter gets the value of XoffCharacter for the instance
func (instance *MSSerial_CommInfo) GetPropertyXoffCharacter() (value uint32, err error) {
	retValue, err := instance.GetProperty("XoffCharacter")
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

// SetXoffXmitThreshold sets the value of XoffXmitThreshold for the instance
func (instance *MSSerial_CommInfo) SetPropertyXoffXmitThreshold(value uint32) (err error) {
	return instance.SetProperty("XoffXmitThreshold", (value))
}

// GetXoffXmitThreshold gets the value of XoffXmitThreshold for the instance
func (instance *MSSerial_CommInfo) GetPropertyXoffXmitThreshold() (value uint32, err error) {
	retValue, err := instance.GetProperty("XoffXmitThreshold")
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

// SetXonCharacter sets the value of XonCharacter for the instance
func (instance *MSSerial_CommInfo) SetPropertyXonCharacter(value uint32) (err error) {
	return instance.SetProperty("XonCharacter", (value))
}

// GetXonCharacter gets the value of XonCharacter for the instance
func (instance *MSSerial_CommInfo) GetPropertyXonCharacter() (value uint32, err error) {
	retValue, err := instance.GetProperty("XonCharacter")
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

// SetXonXmitThreshold sets the value of XonXmitThreshold for the instance
func (instance *MSSerial_CommInfo) SetPropertyXonXmitThreshold(value uint32) (err error) {
	return instance.SetProperty("XonXmitThreshold", (value))
}

// GetXonXmitThreshold gets the value of XonXmitThreshold for the instance
func (instance *MSSerial_CommInfo) GetPropertyXonXmitThreshold() (value uint32, err error) {
	retValue, err := instance.GetProperty("XonXmitThreshold")
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
