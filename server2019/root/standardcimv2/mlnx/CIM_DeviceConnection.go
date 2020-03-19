// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 3/19/2020 using wmigen
//      Source root.StandardCimv2.mlnx
//////////////////////////////////////////////
package mlnx

import (
	"github.com/microsoft/wmi/pkg/base/query"
	cim "github.com/microsoft/wmi/pkg/wmiinstance"
)

// CIM_DeviceConnection struct
type CIM_DeviceConnection struct {
	*CIM_Dependency

	//
	NegotiatedDataWidth uint32

	//
	NegotiatedSpeed uint64
}

func NewCIM_DeviceConnectionEx1(instance *cim.WmiInstance) (newInstance *CIM_DeviceConnection, err error) {
	tmp, err := NewCIM_DependencyEx1(instance)

	if err != nil {
		return
	}
	newInstance = &CIM_DeviceConnection{
		CIM_Dependency: tmp,
	}
	return
}

func NewCIM_DeviceConnectionEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *CIM_DeviceConnection, err error) {
	tmp, err := NewCIM_DependencyEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &CIM_DeviceConnection{
		CIM_Dependency: tmp,
	}
	return
}

// SetNegotiatedDataWidth sets the value of NegotiatedDataWidth for the instance
func (instance *CIM_DeviceConnection) SetPropertyNegotiatedDataWidth(value uint32) (err error) {
	return instance.SetProperty("NegotiatedDataWidth", value)
}

// GetNegotiatedDataWidth gets the value of NegotiatedDataWidth for the instance
func (instance *CIM_DeviceConnection) GetPropertyNegotiatedDataWidth() (value uint32, err error) {
	retValue, err := instance.GetProperty("NegotiatedDataWidth")
	if err != nil {
		return
	}
	value, ok := retValue.(uint32)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetNegotiatedSpeed sets the value of NegotiatedSpeed for the instance
func (instance *CIM_DeviceConnection) SetPropertyNegotiatedSpeed(value uint64) (err error) {
	return instance.SetProperty("NegotiatedSpeed", value)
}

// GetNegotiatedSpeed gets the value of NegotiatedSpeed for the instance
func (instance *CIM_DeviceConnection) GetPropertyNegotiatedSpeed() (value uint64, err error) {
	retValue, err := instance.GetProperty("NegotiatedSpeed")
	if err != nil {
		return
	}
	value, ok := retValue.(uint64)
	if !ok {
		// TODO: Set an error
	}
	return
}
