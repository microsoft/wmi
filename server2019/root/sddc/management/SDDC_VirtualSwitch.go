// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 9/18/2020 using wmigen
//      Source root.SDDC.Management
//////////////////////////////////////////////
package management

import (
	"github.com/microsoft/wmi/pkg/base/instance"
	"github.com/microsoft/wmi/pkg/base/query"
	"github.com/microsoft/wmi/pkg/errors"
	cim "github.com/microsoft/wmi/pkg/wmiinstance"
	"reflect"
)

// SDDC_VirtualSwitch struct
type SDDC_VirtualSwitch struct {
	*cim.WmiInstance

	//
	AllowManagementOS bool

	//
	BandwidthReservationMode uint32

	//
	Extensions []string

	//
	HostName string

	//
	Id string

	//
	IsIovCapable bool

	//
	IsIovEnabled bool

	//
	IsTeamingEnabled bool

	//
	Name string

	//
	NetAdapterDescriptions []string

	//
	Notes string

	//
	SwitchType uint16
}

func NewSDDC_VirtualSwitchEx1(instance *cim.WmiInstance) (newInstance *SDDC_VirtualSwitch, err error) {
	tmp, err := instance, nil

	if err != nil {
		return
	}
	newInstance = &SDDC_VirtualSwitch{
		WmiInstance: tmp,
	}
	return
}

func NewSDDC_VirtualSwitchEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *SDDC_VirtualSwitch, err error) {
	tmp, err := instance.GetWmiInstance(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &SDDC_VirtualSwitch{
		WmiInstance: tmp,
	}
	return
}

// SetAllowManagementOS sets the value of AllowManagementOS for the instance
func (instance *SDDC_VirtualSwitch) SetPropertyAllowManagementOS(value bool) (err error) {
	return instance.SetProperty("AllowManagementOS", (value))
}

// GetAllowManagementOS gets the value of AllowManagementOS for the instance
func (instance *SDDC_VirtualSwitch) GetPropertyAllowManagementOS() (value bool, err error) {
	retValue, err := instance.GetProperty("AllowManagementOS")
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

// SetBandwidthReservationMode sets the value of BandwidthReservationMode for the instance
func (instance *SDDC_VirtualSwitch) SetPropertyBandwidthReservationMode(value uint32) (err error) {
	return instance.SetProperty("BandwidthReservationMode", (value))
}

// GetBandwidthReservationMode gets the value of BandwidthReservationMode for the instance
func (instance *SDDC_VirtualSwitch) GetPropertyBandwidthReservationMode() (value uint32, err error) {
	retValue, err := instance.GetProperty("BandwidthReservationMode")
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

// SetExtensions sets the value of Extensions for the instance
func (instance *SDDC_VirtualSwitch) SetPropertyExtensions(value []string) (err error) {
	return instance.SetProperty("Extensions", (value))
}

// GetExtensions gets the value of Extensions for the instance
func (instance *SDDC_VirtualSwitch) GetPropertyExtensions() (value []string, err error) {
	retValue, err := instance.GetProperty("Extensions")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	for _, interfaceValue := range retValue.([]interface{}) {
		valuetmp, ok := interfaceValue.(string)
		if !ok {
			err = errors.Wrapf(errors.InvalidType, " string is Invalid. Expected %s", reflect.TypeOf(interfaceValue))
			return
		}
		value = append(value, string(valuetmp))
	}

	return
}

// SetHostName sets the value of HostName for the instance
func (instance *SDDC_VirtualSwitch) SetPropertyHostName(value string) (err error) {
	return instance.SetProperty("HostName", (value))
}

// GetHostName gets the value of HostName for the instance
func (instance *SDDC_VirtualSwitch) GetPropertyHostName() (value string, err error) {
	retValue, err := instance.GetProperty("HostName")
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

// SetId sets the value of Id for the instance
func (instance *SDDC_VirtualSwitch) SetPropertyId(value string) (err error) {
	return instance.SetProperty("Id", (value))
}

// GetId gets the value of Id for the instance
func (instance *SDDC_VirtualSwitch) GetPropertyId() (value string, err error) {
	retValue, err := instance.GetProperty("Id")
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

// SetIsIovCapable sets the value of IsIovCapable for the instance
func (instance *SDDC_VirtualSwitch) SetPropertyIsIovCapable(value bool) (err error) {
	return instance.SetProperty("IsIovCapable", (value))
}

// GetIsIovCapable gets the value of IsIovCapable for the instance
func (instance *SDDC_VirtualSwitch) GetPropertyIsIovCapable() (value bool, err error) {
	retValue, err := instance.GetProperty("IsIovCapable")
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

// SetIsIovEnabled sets the value of IsIovEnabled for the instance
func (instance *SDDC_VirtualSwitch) SetPropertyIsIovEnabled(value bool) (err error) {
	return instance.SetProperty("IsIovEnabled", (value))
}

// GetIsIovEnabled gets the value of IsIovEnabled for the instance
func (instance *SDDC_VirtualSwitch) GetPropertyIsIovEnabled() (value bool, err error) {
	retValue, err := instance.GetProperty("IsIovEnabled")
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

// SetIsTeamingEnabled sets the value of IsTeamingEnabled for the instance
func (instance *SDDC_VirtualSwitch) SetPropertyIsTeamingEnabled(value bool) (err error) {
	return instance.SetProperty("IsTeamingEnabled", (value))
}

// GetIsTeamingEnabled gets the value of IsTeamingEnabled for the instance
func (instance *SDDC_VirtualSwitch) GetPropertyIsTeamingEnabled() (value bool, err error) {
	retValue, err := instance.GetProperty("IsTeamingEnabled")
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

// SetName sets the value of Name for the instance
func (instance *SDDC_VirtualSwitch) SetPropertyName(value string) (err error) {
	return instance.SetProperty("Name", (value))
}

// GetName gets the value of Name for the instance
func (instance *SDDC_VirtualSwitch) GetPropertyName() (value string, err error) {
	retValue, err := instance.GetProperty("Name")
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

// SetNetAdapterDescriptions sets the value of NetAdapterDescriptions for the instance
func (instance *SDDC_VirtualSwitch) SetPropertyNetAdapterDescriptions(value []string) (err error) {
	return instance.SetProperty("NetAdapterDescriptions", (value))
}

// GetNetAdapterDescriptions gets the value of NetAdapterDescriptions for the instance
func (instance *SDDC_VirtualSwitch) GetPropertyNetAdapterDescriptions() (value []string, err error) {
	retValue, err := instance.GetProperty("NetAdapterDescriptions")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	for _, interfaceValue := range retValue.([]interface{}) {
		valuetmp, ok := interfaceValue.(string)
		if !ok {
			err = errors.Wrapf(errors.InvalidType, " string is Invalid. Expected %s", reflect.TypeOf(interfaceValue))
			return
		}
		value = append(value, string(valuetmp))
	}

	return
}

// SetNotes sets the value of Notes for the instance
func (instance *SDDC_VirtualSwitch) SetPropertyNotes(value string) (err error) {
	return instance.SetProperty("Notes", (value))
}

// GetNotes gets the value of Notes for the instance
func (instance *SDDC_VirtualSwitch) GetPropertyNotes() (value string, err error) {
	retValue, err := instance.GetProperty("Notes")
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

// SetSwitchType sets the value of SwitchType for the instance
func (instance *SDDC_VirtualSwitch) SetPropertySwitchType(value uint16) (err error) {
	return instance.SetProperty("SwitchType", (value))
}

// GetSwitchType gets the value of SwitchType for the instance
func (instance *SDDC_VirtualSwitch) GetPropertySwitchType() (value uint16, err error) {
	retValue, err := instance.GetProperty("SwitchType")
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

//

// <param name="ReturnValue" type="uint32 "></param>
func (instance *SDDC_VirtualSwitch) Refresh() (result uint32, err error) {
	retVal, err := instance.InvokeMethodWithReturn("Refresh")
	if err != nil {
		return
	}
	result = uint32(retVal)
	return

}
