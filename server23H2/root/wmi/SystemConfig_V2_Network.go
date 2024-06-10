// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

// Author:
//
//	Auto Generated on 6/6/2024 using wmigen
//	Source root.WMI
//
// ////////////////////////////////////////////
package wmi

import (
	"github.com/microsoft/wmi/pkg/base/query"
	"github.com/microsoft/wmi/pkg/errors"
	cim "github.com/microsoft/wmi/pkg/wmiinstance"
	"reflect"
)

// SystemConfig_V2_Network struct
type SystemConfig_V2_Network struct {
	*SystemConfig_V2

	//
	MaxHashTableSize uint32

	//
	MaxUserPort uint32

	//
	TcbTablePartitions uint32

	//
	TcpTimedWaitDelay uint32
}

func NewSystemConfig_V2_NetworkEx1(instance *cim.WmiInstance) (newInstance *SystemConfig_V2_Network, err error) {
	tmp, err := NewSystemConfig_V2Ex1(instance)

	if err != nil {
		return
	}
	newInstance = &SystemConfig_V2_Network{
		SystemConfig_V2: tmp,
	}
	return
}

func NewSystemConfig_V2_NetworkEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *SystemConfig_V2_Network, err error) {
	tmp, err := NewSystemConfig_V2Ex6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &SystemConfig_V2_Network{
		SystemConfig_V2: tmp,
	}
	return
}

// SetMaxHashTableSize sets the value of MaxHashTableSize for the instance
func (instance *SystemConfig_V2_Network) SetPropertyMaxHashTableSize(value uint32) (err error) {
	return instance.SetProperty("MaxHashTableSize", (value))
}

// GetMaxHashTableSize gets the value of MaxHashTableSize for the instance
func (instance *SystemConfig_V2_Network) GetPropertyMaxHashTableSize() (value uint32, err error) {
	retValue, err := instance.GetProperty("MaxHashTableSize")
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

// SetMaxUserPort sets the value of MaxUserPort for the instance
func (instance *SystemConfig_V2_Network) SetPropertyMaxUserPort(value uint32) (err error) {
	return instance.SetProperty("MaxUserPort", (value))
}

// GetMaxUserPort gets the value of MaxUserPort for the instance
func (instance *SystemConfig_V2_Network) GetPropertyMaxUserPort() (value uint32, err error) {
	retValue, err := instance.GetProperty("MaxUserPort")
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

// SetTcbTablePartitions sets the value of TcbTablePartitions for the instance
func (instance *SystemConfig_V2_Network) SetPropertyTcbTablePartitions(value uint32) (err error) {
	return instance.SetProperty("TcbTablePartitions", (value))
}

// GetTcbTablePartitions gets the value of TcbTablePartitions for the instance
func (instance *SystemConfig_V2_Network) GetPropertyTcbTablePartitions() (value uint32, err error) {
	retValue, err := instance.GetProperty("TcbTablePartitions")
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

// SetTcpTimedWaitDelay sets the value of TcpTimedWaitDelay for the instance
func (instance *SystemConfig_V2_Network) SetPropertyTcpTimedWaitDelay(value uint32) (err error) {
	return instance.SetProperty("TcpTimedWaitDelay", (value))
}

// GetTcpTimedWaitDelay gets the value of TcpTimedWaitDelay for the instance
func (instance *SystemConfig_V2_Network) GetPropertyTcpTimedWaitDelay() (value uint32, err error) {
	retValue, err := instance.GetProperty("TcpTimedWaitDelay")
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
