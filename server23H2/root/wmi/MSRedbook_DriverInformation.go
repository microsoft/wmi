// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 6/13/2024 using wmigen
//      Source root.WMI
//////////////////////////////////////////////
package wmi

import (
	"github.com/microsoft/wmi/pkg/base/instance"
	"github.com/microsoft/wmi/pkg/base/query"
	"github.com/microsoft/wmi/pkg/errors"
	cim "github.com/microsoft/wmi/pkg/wmiinstance"
	"reflect"
)

// MSRedbook_DriverInformation struct
type MSRedbook_DriverInformation struct {
	*cim.WmiInstance

	//
	Active bool

	//
	CDDAAccurate bool

	//
	CDDASupported bool

	//
	InstanceName string

	//
	MaximumSectorsPerRead uint32

	//
	NumberOfBuffers uint32

	//
	PlayEnabled bool

	//
	Reserved1 bool

	//
	SectorsPerRead uint32

	//
	SectorsPerReadMask uint32
}

func NewMSRedbook_DriverInformationEx1(instance *cim.WmiInstance) (newInstance *MSRedbook_DriverInformation, err error) {
	tmp, err := instance, nil

	if err != nil {
		return
	}
	newInstance = &MSRedbook_DriverInformation{
		WmiInstance: tmp,
	}
	return
}

func NewMSRedbook_DriverInformationEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *MSRedbook_DriverInformation, err error) {
	tmp, err := instance.GetWmiInstance(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &MSRedbook_DriverInformation{
		WmiInstance: tmp,
	}
	return
}

// SetActive sets the value of Active for the instance
func (instance *MSRedbook_DriverInformation) SetPropertyActive(value bool) (err error) {
	return instance.SetProperty("Active", (value))
}

// GetActive gets the value of Active for the instance
func (instance *MSRedbook_DriverInformation) GetPropertyActive() (value bool, err error) {
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

// SetCDDAAccurate sets the value of CDDAAccurate for the instance
func (instance *MSRedbook_DriverInformation) SetPropertyCDDAAccurate(value bool) (err error) {
	return instance.SetProperty("CDDAAccurate", (value))
}

// GetCDDAAccurate gets the value of CDDAAccurate for the instance
func (instance *MSRedbook_DriverInformation) GetPropertyCDDAAccurate() (value bool, err error) {
	retValue, err := instance.GetProperty("CDDAAccurate")
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

// SetCDDASupported sets the value of CDDASupported for the instance
func (instance *MSRedbook_DriverInformation) SetPropertyCDDASupported(value bool) (err error) {
	return instance.SetProperty("CDDASupported", (value))
}

// GetCDDASupported gets the value of CDDASupported for the instance
func (instance *MSRedbook_DriverInformation) GetPropertyCDDASupported() (value bool, err error) {
	retValue, err := instance.GetProperty("CDDASupported")
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
func (instance *MSRedbook_DriverInformation) SetPropertyInstanceName(value string) (err error) {
	return instance.SetProperty("InstanceName", (value))
}

// GetInstanceName gets the value of InstanceName for the instance
func (instance *MSRedbook_DriverInformation) GetPropertyInstanceName() (value string, err error) {
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

// SetMaximumSectorsPerRead sets the value of MaximumSectorsPerRead for the instance
func (instance *MSRedbook_DriverInformation) SetPropertyMaximumSectorsPerRead(value uint32) (err error) {
	return instance.SetProperty("MaximumSectorsPerRead", (value))
}

// GetMaximumSectorsPerRead gets the value of MaximumSectorsPerRead for the instance
func (instance *MSRedbook_DriverInformation) GetPropertyMaximumSectorsPerRead() (value uint32, err error) {
	retValue, err := instance.GetProperty("MaximumSectorsPerRead")
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

// SetNumberOfBuffers sets the value of NumberOfBuffers for the instance
func (instance *MSRedbook_DriverInformation) SetPropertyNumberOfBuffers(value uint32) (err error) {
	return instance.SetProperty("NumberOfBuffers", (value))
}

// GetNumberOfBuffers gets the value of NumberOfBuffers for the instance
func (instance *MSRedbook_DriverInformation) GetPropertyNumberOfBuffers() (value uint32, err error) {
	retValue, err := instance.GetProperty("NumberOfBuffers")
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

// SetPlayEnabled sets the value of PlayEnabled for the instance
func (instance *MSRedbook_DriverInformation) SetPropertyPlayEnabled(value bool) (err error) {
	return instance.SetProperty("PlayEnabled", (value))
}

// GetPlayEnabled gets the value of PlayEnabled for the instance
func (instance *MSRedbook_DriverInformation) GetPropertyPlayEnabled() (value bool, err error) {
	retValue, err := instance.GetProperty("PlayEnabled")
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

// SetReserved1 sets the value of Reserved1 for the instance
func (instance *MSRedbook_DriverInformation) SetPropertyReserved1(value bool) (err error) {
	return instance.SetProperty("Reserved1", (value))
}

// GetReserved1 gets the value of Reserved1 for the instance
func (instance *MSRedbook_DriverInformation) GetPropertyReserved1() (value bool, err error) {
	retValue, err := instance.GetProperty("Reserved1")
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

// SetSectorsPerRead sets the value of SectorsPerRead for the instance
func (instance *MSRedbook_DriverInformation) SetPropertySectorsPerRead(value uint32) (err error) {
	return instance.SetProperty("SectorsPerRead", (value))
}

// GetSectorsPerRead gets the value of SectorsPerRead for the instance
func (instance *MSRedbook_DriverInformation) GetPropertySectorsPerRead() (value uint32, err error) {
	retValue, err := instance.GetProperty("SectorsPerRead")
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

// SetSectorsPerReadMask sets the value of SectorsPerReadMask for the instance
func (instance *MSRedbook_DriverInformation) SetPropertySectorsPerReadMask(value uint32) (err error) {
	return instance.SetProperty("SectorsPerReadMask", (value))
}

// GetSectorsPerReadMask gets the value of SectorsPerReadMask for the instance
func (instance *MSRedbook_DriverInformation) GetPropertySectorsPerReadMask() (value uint32, err error) {
	retValue, err := instance.GetProperty("SectorsPerReadMask")
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
