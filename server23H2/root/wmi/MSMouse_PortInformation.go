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

// MSMouse_PortInformation struct
type MSMouse_PortInformation struct {
	*MSMouse

	//
	Active bool

	//
	Buttons uint32

	//
	ConnectorType uint32

	//
	DataQueueSize uint32

	//
	ErrorCount uint32

	//
	HardwareType uint32

	//
	InstanceName string
}

func NewMSMouse_PortInformationEx1(instance *cim.WmiInstance) (newInstance *MSMouse_PortInformation, err error) {
	tmp, err := NewMSMouseEx1(instance)

	if err != nil {
		return
	}
	newInstance = &MSMouse_PortInformation{
		MSMouse: tmp,
	}
	return
}

func NewMSMouse_PortInformationEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *MSMouse_PortInformation, err error) {
	tmp, err := NewMSMouseEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &MSMouse_PortInformation{
		MSMouse: tmp,
	}
	return
}

// SetActive sets the value of Active for the instance
func (instance *MSMouse_PortInformation) SetPropertyActive(value bool) (err error) {
	return instance.SetProperty("Active", (value))
}

// GetActive gets the value of Active for the instance
func (instance *MSMouse_PortInformation) GetPropertyActive() (value bool, err error) {
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

// SetButtons sets the value of Buttons for the instance
func (instance *MSMouse_PortInformation) SetPropertyButtons(value uint32) (err error) {
	return instance.SetProperty("Buttons", (value))
}

// GetButtons gets the value of Buttons for the instance
func (instance *MSMouse_PortInformation) GetPropertyButtons() (value uint32, err error) {
	retValue, err := instance.GetProperty("Buttons")
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

// SetConnectorType sets the value of ConnectorType for the instance
func (instance *MSMouse_PortInformation) SetPropertyConnectorType(value uint32) (err error) {
	return instance.SetProperty("ConnectorType", (value))
}

// GetConnectorType gets the value of ConnectorType for the instance
func (instance *MSMouse_PortInformation) GetPropertyConnectorType() (value uint32, err error) {
	retValue, err := instance.GetProperty("ConnectorType")
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

// SetDataQueueSize sets the value of DataQueueSize for the instance
func (instance *MSMouse_PortInformation) SetPropertyDataQueueSize(value uint32) (err error) {
	return instance.SetProperty("DataQueueSize", (value))
}

// GetDataQueueSize gets the value of DataQueueSize for the instance
func (instance *MSMouse_PortInformation) GetPropertyDataQueueSize() (value uint32, err error) {
	retValue, err := instance.GetProperty("DataQueueSize")
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

// SetErrorCount sets the value of ErrorCount for the instance
func (instance *MSMouse_PortInformation) SetPropertyErrorCount(value uint32) (err error) {
	return instance.SetProperty("ErrorCount", (value))
}

// GetErrorCount gets the value of ErrorCount for the instance
func (instance *MSMouse_PortInformation) GetPropertyErrorCount() (value uint32, err error) {
	retValue, err := instance.GetProperty("ErrorCount")
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

// SetHardwareType sets the value of HardwareType for the instance
func (instance *MSMouse_PortInformation) SetPropertyHardwareType(value uint32) (err error) {
	return instance.SetProperty("HardwareType", (value))
}

// GetHardwareType gets the value of HardwareType for the instance
func (instance *MSMouse_PortInformation) GetPropertyHardwareType() (value uint32, err error) {
	retValue, err := instance.GetProperty("HardwareType")
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
func (instance *MSMouse_PortInformation) SetPropertyInstanceName(value string) (err error) {
	return instance.SetProperty("InstanceName", (value))
}

// GetInstanceName gets the value of InstanceName for the instance
func (instance *MSMouse_PortInformation) GetPropertyInstanceName() (value string, err error) {
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
