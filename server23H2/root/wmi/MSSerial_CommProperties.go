// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 6/14/2024 using wmigen
//      Source root.WMI
//////////////////////////////////////////////
package wmi

import (
	"github.com/microsoft/wmi/pkg/base/query"
	"github.com/microsoft/wmi/pkg/errors"
	cim "github.com/microsoft/wmi/pkg/wmiinstance"
	"reflect"
)

// MSSerial_CommProperties struct
type MSSerial_CommProperties struct {
	*MSSerial

	//
	Active bool

	//
	dwCurrentRxQueue uint32

	//
	dwCurrentTxQueue uint32

	//
	dwMaxBaud uint32

	//
	dwMaxRxQueue uint32

	//
	dwMaxTxQueue uint32

	//
	dwProvCapabilities uint32

	//
	dwProvCharSize uint32

	//
	dwProvSpec1 uint32

	//
	dwProvSpec2 uint32

	//
	dwProvSubType uint32

	//
	dwReserved1 uint32

	//
	dwServiceMask uint32

	//
	dwSettableBaud uint32

	//
	dwSettableParams uint32

	//
	InstanceName string

	//
	wcProvChar []uint8

	//
	wPacketLength uint16

	//
	wPacketVersion uint16

	//
	wSettableData uint16

	//
	wSettableStopParity uint16
}

func NewMSSerial_CommPropertiesEx1(instance *cim.WmiInstance) (newInstance *MSSerial_CommProperties, err error) {
	tmp, err := NewMSSerialEx1(instance)

	if err != nil {
		return
	}
	newInstance = &MSSerial_CommProperties{
		MSSerial: tmp,
	}
	return
}

func NewMSSerial_CommPropertiesEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *MSSerial_CommProperties, err error) {
	tmp, err := NewMSSerialEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &MSSerial_CommProperties{
		MSSerial: tmp,
	}
	return
}

// SetActive sets the value of Active for the instance
func (instance *MSSerial_CommProperties) SetPropertyActive(value bool) (err error) {
	return instance.SetProperty("Active", (value))
}

// GetActive gets the value of Active for the instance
func (instance *MSSerial_CommProperties) GetPropertyActive() (value bool, err error) {
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

// SetdwCurrentRxQueue sets the value of dwCurrentRxQueue for the instance
func (instance *MSSerial_CommProperties) SetPropertydwCurrentRxQueue(value uint32) (err error) {
	return instance.SetProperty("dwCurrentRxQueue", (value))
}

// GetdwCurrentRxQueue gets the value of dwCurrentRxQueue for the instance
func (instance *MSSerial_CommProperties) GetPropertydwCurrentRxQueue() (value uint32, err error) {
	retValue, err := instance.GetProperty("dwCurrentRxQueue")
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

// SetdwCurrentTxQueue sets the value of dwCurrentTxQueue for the instance
func (instance *MSSerial_CommProperties) SetPropertydwCurrentTxQueue(value uint32) (err error) {
	return instance.SetProperty("dwCurrentTxQueue", (value))
}

// GetdwCurrentTxQueue gets the value of dwCurrentTxQueue for the instance
func (instance *MSSerial_CommProperties) GetPropertydwCurrentTxQueue() (value uint32, err error) {
	retValue, err := instance.GetProperty("dwCurrentTxQueue")
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

// SetdwMaxBaud sets the value of dwMaxBaud for the instance
func (instance *MSSerial_CommProperties) SetPropertydwMaxBaud(value uint32) (err error) {
	return instance.SetProperty("dwMaxBaud", (value))
}

// GetdwMaxBaud gets the value of dwMaxBaud for the instance
func (instance *MSSerial_CommProperties) GetPropertydwMaxBaud() (value uint32, err error) {
	retValue, err := instance.GetProperty("dwMaxBaud")
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

// SetdwMaxRxQueue sets the value of dwMaxRxQueue for the instance
func (instance *MSSerial_CommProperties) SetPropertydwMaxRxQueue(value uint32) (err error) {
	return instance.SetProperty("dwMaxRxQueue", (value))
}

// GetdwMaxRxQueue gets the value of dwMaxRxQueue for the instance
func (instance *MSSerial_CommProperties) GetPropertydwMaxRxQueue() (value uint32, err error) {
	retValue, err := instance.GetProperty("dwMaxRxQueue")
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

// SetdwMaxTxQueue sets the value of dwMaxTxQueue for the instance
func (instance *MSSerial_CommProperties) SetPropertydwMaxTxQueue(value uint32) (err error) {
	return instance.SetProperty("dwMaxTxQueue", (value))
}

// GetdwMaxTxQueue gets the value of dwMaxTxQueue for the instance
func (instance *MSSerial_CommProperties) GetPropertydwMaxTxQueue() (value uint32, err error) {
	retValue, err := instance.GetProperty("dwMaxTxQueue")
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

// SetdwProvCapabilities sets the value of dwProvCapabilities for the instance
func (instance *MSSerial_CommProperties) SetPropertydwProvCapabilities(value uint32) (err error) {
	return instance.SetProperty("dwProvCapabilities", (value))
}

// GetdwProvCapabilities gets the value of dwProvCapabilities for the instance
func (instance *MSSerial_CommProperties) GetPropertydwProvCapabilities() (value uint32, err error) {
	retValue, err := instance.GetProperty("dwProvCapabilities")
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

// SetdwProvCharSize sets the value of dwProvCharSize for the instance
func (instance *MSSerial_CommProperties) SetPropertydwProvCharSize(value uint32) (err error) {
	return instance.SetProperty("dwProvCharSize", (value))
}

// GetdwProvCharSize gets the value of dwProvCharSize for the instance
func (instance *MSSerial_CommProperties) GetPropertydwProvCharSize() (value uint32, err error) {
	retValue, err := instance.GetProperty("dwProvCharSize")
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

// SetdwProvSpec1 sets the value of dwProvSpec1 for the instance
func (instance *MSSerial_CommProperties) SetPropertydwProvSpec1(value uint32) (err error) {
	return instance.SetProperty("dwProvSpec1", (value))
}

// GetdwProvSpec1 gets the value of dwProvSpec1 for the instance
func (instance *MSSerial_CommProperties) GetPropertydwProvSpec1() (value uint32, err error) {
	retValue, err := instance.GetProperty("dwProvSpec1")
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

// SetdwProvSpec2 sets the value of dwProvSpec2 for the instance
func (instance *MSSerial_CommProperties) SetPropertydwProvSpec2(value uint32) (err error) {
	return instance.SetProperty("dwProvSpec2", (value))
}

// GetdwProvSpec2 gets the value of dwProvSpec2 for the instance
func (instance *MSSerial_CommProperties) GetPropertydwProvSpec2() (value uint32, err error) {
	retValue, err := instance.GetProperty("dwProvSpec2")
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

// SetdwProvSubType sets the value of dwProvSubType for the instance
func (instance *MSSerial_CommProperties) SetPropertydwProvSubType(value uint32) (err error) {
	return instance.SetProperty("dwProvSubType", (value))
}

// GetdwProvSubType gets the value of dwProvSubType for the instance
func (instance *MSSerial_CommProperties) GetPropertydwProvSubType() (value uint32, err error) {
	retValue, err := instance.GetProperty("dwProvSubType")
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

// SetdwReserved1 sets the value of dwReserved1 for the instance
func (instance *MSSerial_CommProperties) SetPropertydwReserved1(value uint32) (err error) {
	return instance.SetProperty("dwReserved1", (value))
}

// GetdwReserved1 gets the value of dwReserved1 for the instance
func (instance *MSSerial_CommProperties) GetPropertydwReserved1() (value uint32, err error) {
	retValue, err := instance.GetProperty("dwReserved1")
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

// SetdwServiceMask sets the value of dwServiceMask for the instance
func (instance *MSSerial_CommProperties) SetPropertydwServiceMask(value uint32) (err error) {
	return instance.SetProperty("dwServiceMask", (value))
}

// GetdwServiceMask gets the value of dwServiceMask for the instance
func (instance *MSSerial_CommProperties) GetPropertydwServiceMask() (value uint32, err error) {
	retValue, err := instance.GetProperty("dwServiceMask")
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

// SetdwSettableBaud sets the value of dwSettableBaud for the instance
func (instance *MSSerial_CommProperties) SetPropertydwSettableBaud(value uint32) (err error) {
	return instance.SetProperty("dwSettableBaud", (value))
}

// GetdwSettableBaud gets the value of dwSettableBaud for the instance
func (instance *MSSerial_CommProperties) GetPropertydwSettableBaud() (value uint32, err error) {
	retValue, err := instance.GetProperty("dwSettableBaud")
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

// SetdwSettableParams sets the value of dwSettableParams for the instance
func (instance *MSSerial_CommProperties) SetPropertydwSettableParams(value uint32) (err error) {
	return instance.SetProperty("dwSettableParams", (value))
}

// GetdwSettableParams gets the value of dwSettableParams for the instance
func (instance *MSSerial_CommProperties) GetPropertydwSettableParams() (value uint32, err error) {
	retValue, err := instance.GetProperty("dwSettableParams")
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
func (instance *MSSerial_CommProperties) SetPropertyInstanceName(value string) (err error) {
	return instance.SetProperty("InstanceName", (value))
}

// GetInstanceName gets the value of InstanceName for the instance
func (instance *MSSerial_CommProperties) GetPropertyInstanceName() (value string, err error) {
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

// SetwcProvChar sets the value of wcProvChar for the instance
func (instance *MSSerial_CommProperties) SetPropertywcProvChar(value []uint8) (err error) {
	return instance.SetProperty("wcProvChar", (value))
}

// GetwcProvChar gets the value of wcProvChar for the instance
func (instance *MSSerial_CommProperties) GetPropertywcProvChar() (value []uint8, err error) {
	retValue, err := instance.GetProperty("wcProvChar")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	for _, interfaceValue := range retValue.([]interface{}) {
		valuetmp, ok := interfaceValue.(uint8)
		if !ok {
			err = errors.Wrapf(errors.InvalidType, " uint8 is Invalid. Expected %s", reflect.TypeOf(interfaceValue))
			return
		}
		value = append(value, uint8(valuetmp))
	}

	return
}

// SetwPacketLength sets the value of wPacketLength for the instance
func (instance *MSSerial_CommProperties) SetPropertywPacketLength(value uint16) (err error) {
	return instance.SetProperty("wPacketLength", (value))
}

// GetwPacketLength gets the value of wPacketLength for the instance
func (instance *MSSerial_CommProperties) GetPropertywPacketLength() (value uint16, err error) {
	retValue, err := instance.GetProperty("wPacketLength")
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

// SetwPacketVersion sets the value of wPacketVersion for the instance
func (instance *MSSerial_CommProperties) SetPropertywPacketVersion(value uint16) (err error) {
	return instance.SetProperty("wPacketVersion", (value))
}

// GetwPacketVersion gets the value of wPacketVersion for the instance
func (instance *MSSerial_CommProperties) GetPropertywPacketVersion() (value uint16, err error) {
	retValue, err := instance.GetProperty("wPacketVersion")
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

// SetwSettableData sets the value of wSettableData for the instance
func (instance *MSSerial_CommProperties) SetPropertywSettableData(value uint16) (err error) {
	return instance.SetProperty("wSettableData", (value))
}

// GetwSettableData gets the value of wSettableData for the instance
func (instance *MSSerial_CommProperties) GetPropertywSettableData() (value uint16, err error) {
	retValue, err := instance.GetProperty("wSettableData")
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

// SetwSettableStopParity sets the value of wSettableStopParity for the instance
func (instance *MSSerial_CommProperties) SetPropertywSettableStopParity(value uint16) (err error) {
	return instance.SetProperty("wSettableStopParity", (value))
}

// GetwSettableStopParity gets the value of wSettableStopParity for the instance
func (instance *MSSerial_CommProperties) GetPropertywSettableStopParity() (value uint16, err error) {
	retValue, err := instance.GetProperty("wSettableStopParity")
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
