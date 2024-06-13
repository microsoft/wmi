// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 6/13/2024 using wmigen
//      Source root.HyperVCluster.v2
//////////////////////////////////////////////
package v2

import (
	"github.com/microsoft/wmi/pkg/base/query"
	"github.com/microsoft/wmi/pkg/errors"
	cim "github.com/microsoft/wmi/pkg/wmiinstance"
	"reflect"
)

// CIM_DeviceConnection struct
type CIM_DeviceConnection struct {
	*CIM_Dependency

	// When several bus and connection data widths are possible, the NegotiatedDataWidth property defines the one that is in use between the Devices. Data width is specified in bits. If data width is not negotiated, or if this information is not available or not important to Device management, the property should be set to 0.
	NegotiatedDataWidth uint32

	// When several bus and connection speeds are possible, the NegotiatedSpeed property defines the one that is in use between the Devices. Speed is specified in bits per second. If connection or bus speeds are not negotiated, or if this information is not available or not important to Device management, the property should be set to 0.
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
	return instance.SetProperty("NegotiatedDataWidth", (value))
}

// GetNegotiatedDataWidth gets the value of NegotiatedDataWidth for the instance
func (instance *CIM_DeviceConnection) GetPropertyNegotiatedDataWidth() (value uint32, err error) {
	retValue, err := instance.GetProperty("NegotiatedDataWidth")
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

// SetNegotiatedSpeed sets the value of NegotiatedSpeed for the instance
func (instance *CIM_DeviceConnection) SetPropertyNegotiatedSpeed(value uint64) (err error) {
	return instance.SetProperty("NegotiatedSpeed", (value))
}

// GetNegotiatedSpeed gets the value of NegotiatedSpeed for the instance
func (instance *CIM_DeviceConnection) GetPropertyNegotiatedSpeed() (value uint64, err error) {
	retValue, err := instance.GetProperty("NegotiatedSpeed")
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
