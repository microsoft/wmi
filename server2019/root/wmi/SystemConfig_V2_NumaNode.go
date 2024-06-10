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

// SystemConfig_V2_NumaNode struct
type SystemConfig_V2_NumaNode struct {
	*SystemConfig_V2

	//
	NodeCount uint32

	//
	NodeMap []uint64
}

func NewSystemConfig_V2_NumaNodeEx1(instance *cim.WmiInstance) (newInstance *SystemConfig_V2_NumaNode, err error) {
	tmp, err := NewSystemConfig_V2Ex1(instance)

	if err != nil {
		return
	}
	newInstance = &SystemConfig_V2_NumaNode{
		SystemConfig_V2: tmp,
	}
	return
}

func NewSystemConfig_V2_NumaNodeEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *SystemConfig_V2_NumaNode, err error) {
	tmp, err := NewSystemConfig_V2Ex6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &SystemConfig_V2_NumaNode{
		SystemConfig_V2: tmp,
	}
	return
}

// SetNodeCount sets the value of NodeCount for the instance
func (instance *SystemConfig_V2_NumaNode) SetPropertyNodeCount(value uint32) (err error) {
	return instance.SetProperty("NodeCount", (value))
}

// GetNodeCount gets the value of NodeCount for the instance
func (instance *SystemConfig_V2_NumaNode) GetPropertyNodeCount() (value uint32, err error) {
	retValue, err := instance.GetProperty("NodeCount")
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

// SetNodeMap sets the value of NodeMap for the instance
func (instance *SystemConfig_V2_NumaNode) SetPropertyNodeMap(value []uint64) (err error) {
	return instance.SetProperty("NodeMap", (value))
}

// GetNodeMap gets the value of NodeMap for the instance
func (instance *SystemConfig_V2_NumaNode) GetPropertyNodeMap() (value []uint64, err error) {
	retValue, err := instance.GetProperty("NodeMap")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	for _, interfaceValue := range retValue.([]interface{}) {
		valuetmp, ok := interfaceValue.(uint64)
		if !ok {
			err = errors.Wrapf(errors.InvalidType, " uint64 is Invalid. Expected %s", reflect.TypeOf(interfaceValue))
			return
		}
		value = append(value, uint64(valuetmp))
	}

	return
}
