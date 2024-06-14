// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 6/14/2024 using wmigen
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

// MSiSCSI_NICConfig struct
type MSiSCSI_NICConfig struct {
	*cim.WmiInstance

	//
	Active bool

	//
	InstanceName string

	// Speed of network link in megabits per second.
	LinkSpeed uint32

	// Link State **typedef**
	LinkState NICConfig_LinkState

	// Ethernet MAC Address
	MacAddress []uint8

	// Maximum frame size
	MaxFrameSize uint32

	// Maximum Speed of network link in megabits per second.
	MaxLinkSpeed uint32
}

func NewMSiSCSI_NICConfigEx1(instance *cim.WmiInstance) (newInstance *MSiSCSI_NICConfig, err error) {
	tmp, err := instance, nil

	if err != nil {
		return
	}
	newInstance = &MSiSCSI_NICConfig{
		WmiInstance: tmp,
	}
	return
}

func NewMSiSCSI_NICConfigEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *MSiSCSI_NICConfig, err error) {
	tmp, err := instance.GetWmiInstance(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &MSiSCSI_NICConfig{
		WmiInstance: tmp,
	}
	return
}

// SetActive sets the value of Active for the instance
func (instance *MSiSCSI_NICConfig) SetPropertyActive(value bool) (err error) {
	return instance.SetProperty("Active", (value))
}

// GetActive gets the value of Active for the instance
func (instance *MSiSCSI_NICConfig) GetPropertyActive() (value bool, err error) {
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
func (instance *MSiSCSI_NICConfig) SetPropertyInstanceName(value string) (err error) {
	return instance.SetProperty("InstanceName", (value))
}

// GetInstanceName gets the value of InstanceName for the instance
func (instance *MSiSCSI_NICConfig) GetPropertyInstanceName() (value string, err error) {
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

// SetLinkSpeed sets the value of LinkSpeed for the instance
func (instance *MSiSCSI_NICConfig) SetPropertyLinkSpeed(value uint32) (err error) {
	return instance.SetProperty("LinkSpeed", (value))
}

// GetLinkSpeed gets the value of LinkSpeed for the instance
func (instance *MSiSCSI_NICConfig) GetPropertyLinkSpeed() (value uint32, err error) {
	retValue, err := instance.GetProperty("LinkSpeed")
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

// SetLinkState sets the value of LinkState for the instance
func (instance *MSiSCSI_NICConfig) SetPropertyLinkState(value NICConfig_LinkState) (err error) {
	return instance.SetProperty("LinkState", (value))
}

// GetLinkState gets the value of LinkState for the instance
func (instance *MSiSCSI_NICConfig) GetPropertyLinkState() (value NICConfig_LinkState, err error) {
	retValue, err := instance.GetProperty("LinkState")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(int32)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " int32 is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = NICConfig_LinkState(valuetmp)

	return
}

// SetMacAddress sets the value of MacAddress for the instance
func (instance *MSiSCSI_NICConfig) SetPropertyMacAddress(value []uint8) (err error) {
	return instance.SetProperty("MacAddress", (value))
}

// GetMacAddress gets the value of MacAddress for the instance
func (instance *MSiSCSI_NICConfig) GetPropertyMacAddress() (value []uint8, err error) {
	retValue, err := instance.GetProperty("MacAddress")
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

// SetMaxFrameSize sets the value of MaxFrameSize for the instance
func (instance *MSiSCSI_NICConfig) SetPropertyMaxFrameSize(value uint32) (err error) {
	return instance.SetProperty("MaxFrameSize", (value))
}

// GetMaxFrameSize gets the value of MaxFrameSize for the instance
func (instance *MSiSCSI_NICConfig) GetPropertyMaxFrameSize() (value uint32, err error) {
	retValue, err := instance.GetProperty("MaxFrameSize")
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

// SetMaxLinkSpeed sets the value of MaxLinkSpeed for the instance
func (instance *MSiSCSI_NICConfig) SetPropertyMaxLinkSpeed(value uint32) (err error) {
	return instance.SetProperty("MaxLinkSpeed", (value))
}

// GetMaxLinkSpeed gets the value of MaxLinkSpeed for the instance
func (instance *MSiSCSI_NICConfig) GetPropertyMaxLinkSpeed() (value uint32, err error) {
	retValue, err := instance.GetProperty("MaxLinkSpeed")
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
