// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 6/13/2024 using wmigen
//      Source root.WMI
//////////////////////////////////////////////
package wmi

import (
	"github.com/microsoft/wmi/pkg/base/query"
	"github.com/microsoft/wmi/pkg/errors"
	cim "github.com/microsoft/wmi/pkg/wmiinstance"
	"reflect"
)

// MSSmBios_SMBiosEventlog struct
type MSSmBios_SMBiosEventlog struct {
	*MS_SmBios

	//
	AccessMethod uint8

	//
	AccessMethodAddress uint32

	//
	Active bool

	//
	InstanceName string

	//
	LengthEachLogTypeDesc uint8

	//
	ListLogTypeDesc []uint8

	//
	LogArea []uint8

	//
	LogAreaLength uint16

	//
	LogChangeToken uint32

	//
	LogDataStart uint16

	//
	LogHeaderDescExists bool

	//
	LogHeaderFormat uint8

	//
	LogHeaderStart uint16

	//
	LogStatus uint8

	//
	LogTypeDescLength uint16

	//
	NumberLogTypeDesc uint8

	//
	Reserved uint8
}

func NewMSSmBios_SMBiosEventlogEx1(instance *cim.WmiInstance) (newInstance *MSSmBios_SMBiosEventlog, err error) {
	tmp, err := NewMS_SmBiosEx1(instance)

	if err != nil {
		return
	}
	newInstance = &MSSmBios_SMBiosEventlog{
		MS_SmBios: tmp,
	}
	return
}

func NewMSSmBios_SMBiosEventlogEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *MSSmBios_SMBiosEventlog, err error) {
	tmp, err := NewMS_SmBiosEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &MSSmBios_SMBiosEventlog{
		MS_SmBios: tmp,
	}
	return
}

// SetAccessMethod sets the value of AccessMethod for the instance
func (instance *MSSmBios_SMBiosEventlog) SetPropertyAccessMethod(value uint8) (err error) {
	return instance.SetProperty("AccessMethod", (value))
}

// GetAccessMethod gets the value of AccessMethod for the instance
func (instance *MSSmBios_SMBiosEventlog) GetPropertyAccessMethod() (value uint8, err error) {
	retValue, err := instance.GetProperty("AccessMethod")
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

// SetAccessMethodAddress sets the value of AccessMethodAddress for the instance
func (instance *MSSmBios_SMBiosEventlog) SetPropertyAccessMethodAddress(value uint32) (err error) {
	return instance.SetProperty("AccessMethodAddress", (value))
}

// GetAccessMethodAddress gets the value of AccessMethodAddress for the instance
func (instance *MSSmBios_SMBiosEventlog) GetPropertyAccessMethodAddress() (value uint32, err error) {
	retValue, err := instance.GetProperty("AccessMethodAddress")
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

// SetActive sets the value of Active for the instance
func (instance *MSSmBios_SMBiosEventlog) SetPropertyActive(value bool) (err error) {
	return instance.SetProperty("Active", (value))
}

// GetActive gets the value of Active for the instance
func (instance *MSSmBios_SMBiosEventlog) GetPropertyActive() (value bool, err error) {
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

// SetInstanceName sets the value of InstanceName for the instance
func (instance *MSSmBios_SMBiosEventlog) SetPropertyInstanceName(value string) (err error) {
	return instance.SetProperty("InstanceName", (value))
}

// GetInstanceName gets the value of InstanceName for the instance
func (instance *MSSmBios_SMBiosEventlog) GetPropertyInstanceName() (value string, err error) {
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

// SetLengthEachLogTypeDesc sets the value of LengthEachLogTypeDesc for the instance
func (instance *MSSmBios_SMBiosEventlog) SetPropertyLengthEachLogTypeDesc(value uint8) (err error) {
	return instance.SetProperty("LengthEachLogTypeDesc", (value))
}

// GetLengthEachLogTypeDesc gets the value of LengthEachLogTypeDesc for the instance
func (instance *MSSmBios_SMBiosEventlog) GetPropertyLengthEachLogTypeDesc() (value uint8, err error) {
	retValue, err := instance.GetProperty("LengthEachLogTypeDesc")
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

// SetListLogTypeDesc sets the value of ListLogTypeDesc for the instance
func (instance *MSSmBios_SMBiosEventlog) SetPropertyListLogTypeDesc(value []uint8) (err error) {
	return instance.SetProperty("ListLogTypeDesc", (value))
}

// GetListLogTypeDesc gets the value of ListLogTypeDesc for the instance
func (instance *MSSmBios_SMBiosEventlog) GetPropertyListLogTypeDesc() (value []uint8, err error) {
	retValue, err := instance.GetProperty("ListLogTypeDesc")
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

// SetLogArea sets the value of LogArea for the instance
func (instance *MSSmBios_SMBiosEventlog) SetPropertyLogArea(value []uint8) (err error) {
	return instance.SetProperty("LogArea", (value))
}

// GetLogArea gets the value of LogArea for the instance
func (instance *MSSmBios_SMBiosEventlog) GetPropertyLogArea() (value []uint8, err error) {
	retValue, err := instance.GetProperty("LogArea")
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

// SetLogAreaLength sets the value of LogAreaLength for the instance
func (instance *MSSmBios_SMBiosEventlog) SetPropertyLogAreaLength(value uint16) (err error) {
	return instance.SetProperty("LogAreaLength", (value))
}

// GetLogAreaLength gets the value of LogAreaLength for the instance
func (instance *MSSmBios_SMBiosEventlog) GetPropertyLogAreaLength() (value uint16, err error) {
	retValue, err := instance.GetProperty("LogAreaLength")
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

// SetLogChangeToken sets the value of LogChangeToken for the instance
func (instance *MSSmBios_SMBiosEventlog) SetPropertyLogChangeToken(value uint32) (err error) {
	return instance.SetProperty("LogChangeToken", (value))
}

// GetLogChangeToken gets the value of LogChangeToken for the instance
func (instance *MSSmBios_SMBiosEventlog) GetPropertyLogChangeToken() (value uint32, err error) {
	retValue, err := instance.GetProperty("LogChangeToken")
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

// SetLogDataStart sets the value of LogDataStart for the instance
func (instance *MSSmBios_SMBiosEventlog) SetPropertyLogDataStart(value uint16) (err error) {
	return instance.SetProperty("LogDataStart", (value))
}

// GetLogDataStart gets the value of LogDataStart for the instance
func (instance *MSSmBios_SMBiosEventlog) GetPropertyLogDataStart() (value uint16, err error) {
	retValue, err := instance.GetProperty("LogDataStart")
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

// SetLogHeaderDescExists sets the value of LogHeaderDescExists for the instance
func (instance *MSSmBios_SMBiosEventlog) SetPropertyLogHeaderDescExists(value bool) (err error) {
	return instance.SetProperty("LogHeaderDescExists", (value))
}

// GetLogHeaderDescExists gets the value of LogHeaderDescExists for the instance
func (instance *MSSmBios_SMBiosEventlog) GetPropertyLogHeaderDescExists() (value bool, err error) {
	retValue, err := instance.GetProperty("LogHeaderDescExists")
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

// SetLogHeaderFormat sets the value of LogHeaderFormat for the instance
func (instance *MSSmBios_SMBiosEventlog) SetPropertyLogHeaderFormat(value uint8) (err error) {
	return instance.SetProperty("LogHeaderFormat", (value))
}

// GetLogHeaderFormat gets the value of LogHeaderFormat for the instance
func (instance *MSSmBios_SMBiosEventlog) GetPropertyLogHeaderFormat() (value uint8, err error) {
	retValue, err := instance.GetProperty("LogHeaderFormat")
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

// SetLogHeaderStart sets the value of LogHeaderStart for the instance
func (instance *MSSmBios_SMBiosEventlog) SetPropertyLogHeaderStart(value uint16) (err error) {
	return instance.SetProperty("LogHeaderStart", (value))
}

// GetLogHeaderStart gets the value of LogHeaderStart for the instance
func (instance *MSSmBios_SMBiosEventlog) GetPropertyLogHeaderStart() (value uint16, err error) {
	retValue, err := instance.GetProperty("LogHeaderStart")
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

// SetLogStatus sets the value of LogStatus for the instance
func (instance *MSSmBios_SMBiosEventlog) SetPropertyLogStatus(value uint8) (err error) {
	return instance.SetProperty("LogStatus", (value))
}

// GetLogStatus gets the value of LogStatus for the instance
func (instance *MSSmBios_SMBiosEventlog) GetPropertyLogStatus() (value uint8, err error) {
	retValue, err := instance.GetProperty("LogStatus")
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

// SetLogTypeDescLength sets the value of LogTypeDescLength for the instance
func (instance *MSSmBios_SMBiosEventlog) SetPropertyLogTypeDescLength(value uint16) (err error) {
	return instance.SetProperty("LogTypeDescLength", (value))
}

// GetLogTypeDescLength gets the value of LogTypeDescLength for the instance
func (instance *MSSmBios_SMBiosEventlog) GetPropertyLogTypeDescLength() (value uint16, err error) {
	retValue, err := instance.GetProperty("LogTypeDescLength")
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

// SetNumberLogTypeDesc sets the value of NumberLogTypeDesc for the instance
func (instance *MSSmBios_SMBiosEventlog) SetPropertyNumberLogTypeDesc(value uint8) (err error) {
	return instance.SetProperty("NumberLogTypeDesc", (value))
}

// GetNumberLogTypeDesc gets the value of NumberLogTypeDesc for the instance
func (instance *MSSmBios_SMBiosEventlog) GetPropertyNumberLogTypeDesc() (value uint8, err error) {
	retValue, err := instance.GetProperty("NumberLogTypeDesc")
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

// SetReserved sets the value of Reserved for the instance
func (instance *MSSmBios_SMBiosEventlog) SetPropertyReserved(value uint8) (err error) {
	return instance.SetProperty("Reserved", (value))
}

// GetReserved gets the value of Reserved for the instance
func (instance *MSSmBios_SMBiosEventlog) GetPropertyReserved() (value uint8, err error) {
	retValue, err := instance.GetProperty("Reserved")
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
