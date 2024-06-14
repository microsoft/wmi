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

// HBAFCPID struct
type HBAFCPID struct {
	*cim.WmiInstance

	//
	Fcid uint32

	//
	FcpLun uint64

	//
	NodeWWN []uint8

	//
	PortWWN []uint8
}

func NewHBAFCPIDEx1(instance *cim.WmiInstance) (newInstance *HBAFCPID, err error) {
	tmp, err := instance, nil

	if err != nil {
		return
	}
	newInstance = &HBAFCPID{
		WmiInstance: tmp,
	}
	return
}

func NewHBAFCPIDEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *HBAFCPID, err error) {
	tmp, err := instance.GetWmiInstance(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &HBAFCPID{
		WmiInstance: tmp,
	}
	return
}

// SetFcid sets the value of Fcid for the instance
func (instance *HBAFCPID) SetPropertyFcid(value uint32) (err error) {
	return instance.SetProperty("Fcid", (value))
}

// GetFcid gets the value of Fcid for the instance
func (instance *HBAFCPID) GetPropertyFcid() (value uint32, err error) {
	retValue, err := instance.GetProperty("Fcid")
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

// SetFcpLun sets the value of FcpLun for the instance
func (instance *HBAFCPID) SetPropertyFcpLun(value uint64) (err error) {
	return instance.SetProperty("FcpLun", (value))
}

// GetFcpLun gets the value of FcpLun for the instance
func (instance *HBAFCPID) GetPropertyFcpLun() (value uint64, err error) {
	retValue, err := instance.GetProperty("FcpLun")
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

// SetNodeWWN sets the value of NodeWWN for the instance
func (instance *HBAFCPID) SetPropertyNodeWWN(value []uint8) (err error) {
	return instance.SetProperty("NodeWWN", (value))
}

// GetNodeWWN gets the value of NodeWWN for the instance
func (instance *HBAFCPID) GetPropertyNodeWWN() (value []uint8, err error) {
	retValue, err := instance.GetProperty("NodeWWN")
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
func (instance *HBAFCPID) SetPropertyPortWWN(value []uint8) (err error) {
	return instance.SetProperty("PortWWN", (value))
}

// GetPortWWN gets the value of PortWWN for the instance
func (instance *HBAFCPID) GetPropertyPortWWN() (value []uint8, err error) {
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
