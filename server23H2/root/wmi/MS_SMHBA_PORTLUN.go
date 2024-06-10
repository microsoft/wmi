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
	"github.com/microsoft/wmi/pkg/base/instance"
	"github.com/microsoft/wmi/pkg/base/query"
	"github.com/microsoft/wmi/pkg/errors"
	cim "github.com/microsoft/wmi/pkg/wmiinstance"
	"reflect"
)

// MS_SMHBA_PORTLUN struct
type MS_SMHBA_PORTLUN struct {
	*cim.WmiInstance

	//
	domainPortWWN []uint8

	//
	PortWWN []uint8

	//
	TargetLun uint64
}

func NewMS_SMHBA_PORTLUNEx1(instance *cim.WmiInstance) (newInstance *MS_SMHBA_PORTLUN, err error) {
	tmp, err := instance, nil

	if err != nil {
		return
	}
	newInstance = &MS_SMHBA_PORTLUN{
		WmiInstance: tmp,
	}
	return
}

func NewMS_SMHBA_PORTLUNEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *MS_SMHBA_PORTLUN, err error) {
	tmp, err := instance.GetWmiInstance(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &MS_SMHBA_PORTLUN{
		WmiInstance: tmp,
	}
	return
}

// SetdomainPortWWN sets the value of domainPortWWN for the instance
func (instance *MS_SMHBA_PORTLUN) SetPropertydomainPortWWN(value []uint8) (err error) {
	return instance.SetProperty("domainPortWWN", (value))
}

// GetdomainPortWWN gets the value of domainPortWWN for the instance
func (instance *MS_SMHBA_PORTLUN) GetPropertydomainPortWWN() (value []uint8, err error) {
	retValue, err := instance.GetProperty("domainPortWWN")
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

// SetPortWWN sets the value of PortWWN for the instance
func (instance *MS_SMHBA_PORTLUN) SetPropertyPortWWN(value []uint8) (err error) {
	return instance.SetProperty("PortWWN", (value))
}

// GetPortWWN gets the value of PortWWN for the instance
func (instance *MS_SMHBA_PORTLUN) GetPropertyPortWWN() (value []uint8, err error) {
	retValue, err := instance.GetProperty("PortWWN")
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

// SetTargetLun sets the value of TargetLun for the instance
func (instance *MS_SMHBA_PORTLUN) SetPropertyTargetLun(value uint64) (err error) {
	return instance.SetProperty("TargetLun", (value))
}

// GetTargetLun gets the value of TargetLun for the instance
func (instance *MS_SMHBA_PORTLUN) GetPropertyTargetLun() (value uint64, err error) {
	retValue, err := instance.GetProperty("TargetLun")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(uint64)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " uint64 is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = uint64(valuetmp)

	return
}
