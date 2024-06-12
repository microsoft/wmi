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

// MSSmBios_RawSMBiosTables struct
type MSSmBios_RawSMBiosTables struct {
	*MS_SmBios

	//
	Active bool

	//
	DmiRevision uint8

	//
	InstanceName string

	//
	Size uint32

	//
	SMBiosData []uint8

	//
	SmbiosMajorVersion uint8

	//
	SmbiosMinorVersion uint8

	//
	Used20CallingMethod bool
}

func NewMSSmBios_RawSMBiosTablesEx1(instance *cim.WmiInstance) (newInstance *MSSmBios_RawSMBiosTables, err error) {
	tmp, err := NewMS_SmBiosEx1(instance)

	if err != nil {
		return
	}
	newInstance = &MSSmBios_RawSMBiosTables{
		MS_SmBios: tmp,
	}
	return
}

func NewMSSmBios_RawSMBiosTablesEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *MSSmBios_RawSMBiosTables, err error) {
	tmp, err := NewMS_SmBiosEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &MSSmBios_RawSMBiosTables{
		MS_SmBios: tmp,
	}
	return
}

// SetActive sets the value of Active for the instance
func (instance *MSSmBios_RawSMBiosTables) SetPropertyActive(value bool) (err error) {
	return instance.SetProperty("Active", (value))
}

// GetActive gets the value of Active for the instance
func (instance *MSSmBios_RawSMBiosTables) GetPropertyActive() (value bool, err error) {
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

// SetDmiRevision sets the value of DmiRevision for the instance
func (instance *MSSmBios_RawSMBiosTables) SetPropertyDmiRevision(value uint8) (err error) {
	return instance.SetProperty("DmiRevision", (value))
}

// GetDmiRevision gets the value of DmiRevision for the instance
func (instance *MSSmBios_RawSMBiosTables) GetPropertyDmiRevision() (value uint8, err error) {
	retValue, err := instance.GetProperty("DmiRevision")
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

// SetInstanceName sets the value of InstanceName for the instance
func (instance *MSSmBios_RawSMBiosTables) SetPropertyInstanceName(value string) (err error) {
	return instance.SetProperty("InstanceName", (value))
}

// GetInstanceName gets the value of InstanceName for the instance
func (instance *MSSmBios_RawSMBiosTables) GetPropertyInstanceName() (value string, err error) {
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

// SetSize sets the value of Size for the instance
func (instance *MSSmBios_RawSMBiosTables) SetPropertySize(value uint32) (err error) {
	return instance.SetProperty("Size", (value))
}

// GetSize gets the value of Size for the instance
func (instance *MSSmBios_RawSMBiosTables) GetPropertySize() (value uint32, err error) {
	retValue, err := instance.GetProperty("Size")
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

// SetSMBiosData sets the value of SMBiosData for the instance
func (instance *MSSmBios_RawSMBiosTables) SetPropertySMBiosData(value []uint8) (err error) {
	return instance.SetProperty("SMBiosData", (value))
}

// GetSMBiosData gets the value of SMBiosData for the instance
func (instance *MSSmBios_RawSMBiosTables) GetPropertySMBiosData() (value []uint8, err error) {
	retValue, err := instance.GetProperty("SMBiosData")
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

// SetSmbiosMajorVersion sets the value of SmbiosMajorVersion for the instance
func (instance *MSSmBios_RawSMBiosTables) SetPropertySmbiosMajorVersion(value uint8) (err error) {
	return instance.SetProperty("SmbiosMajorVersion", (value))
}

// GetSmbiosMajorVersion gets the value of SmbiosMajorVersion for the instance
func (instance *MSSmBios_RawSMBiosTables) GetPropertySmbiosMajorVersion() (value uint8, err error) {
	retValue, err := instance.GetProperty("SmbiosMajorVersion")
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

// SetSmbiosMinorVersion sets the value of SmbiosMinorVersion for the instance
func (instance *MSSmBios_RawSMBiosTables) SetPropertySmbiosMinorVersion(value uint8) (err error) {
	return instance.SetProperty("SmbiosMinorVersion", (value))
}

// GetSmbiosMinorVersion gets the value of SmbiosMinorVersion for the instance
func (instance *MSSmBios_RawSMBiosTables) GetPropertySmbiosMinorVersion() (value uint8, err error) {
	retValue, err := instance.GetProperty("SmbiosMinorVersion")
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

// SetUsed20CallingMethod sets the value of Used20CallingMethod for the instance
func (instance *MSSmBios_RawSMBiosTables) SetPropertyUsed20CallingMethod(value bool) (err error) {
	return instance.SetProperty("Used20CallingMethod", (value))
}

// GetUsed20CallingMethod gets the value of Used20CallingMethod for the instance
func (instance *MSSmBios_RawSMBiosTables) GetPropertyUsed20CallingMethod() (value bool, err error) {
	retValue, err := instance.GetProperty("Used20CallingMethod")
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
