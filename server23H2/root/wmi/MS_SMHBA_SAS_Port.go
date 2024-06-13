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

// MS_SMHBA_SAS_Port struct
type MS_SMHBA_SAS_Port struct {
	*cim.WmiInstance

	//
	AttachedSASAddress []uint8

	//
	LocalSASAddress []uint8

	//
	NumberofDiscoveredPorts uint32

	//
	NumberofPhys uint32

	//
	PortProtocol uint32
}

func NewMS_SMHBA_SAS_PortEx1(instance *cim.WmiInstance) (newInstance *MS_SMHBA_SAS_Port, err error) {
	tmp, err := instance, nil

	if err != nil {
		return
	}
	newInstance = &MS_SMHBA_SAS_Port{
		WmiInstance: tmp,
	}
	return
}

func NewMS_SMHBA_SAS_PortEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *MS_SMHBA_SAS_Port, err error) {
	tmp, err := instance.GetWmiInstance(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &MS_SMHBA_SAS_Port{
		WmiInstance: tmp,
	}
	return
}

// SetAttachedSASAddress sets the value of AttachedSASAddress for the instance
func (instance *MS_SMHBA_SAS_Port) SetPropertyAttachedSASAddress(value []uint8) (err error) {
	return instance.SetProperty("AttachedSASAddress", (value))
}

// GetAttachedSASAddress gets the value of AttachedSASAddress for the instance
func (instance *MS_SMHBA_SAS_Port) GetPropertyAttachedSASAddress() (value []uint8, err error) {
	retValue, err := instance.GetProperty("AttachedSASAddress")
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

// SetLocalSASAddress sets the value of LocalSASAddress for the instance
func (instance *MS_SMHBA_SAS_Port) SetPropertyLocalSASAddress(value []uint8) (err error) {
	return instance.SetProperty("LocalSASAddress", (value))
}

// GetLocalSASAddress gets the value of LocalSASAddress for the instance
func (instance *MS_SMHBA_SAS_Port) GetPropertyLocalSASAddress() (value []uint8, err error) {
	retValue, err := instance.GetProperty("LocalSASAddress")
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

// SetNumberofDiscoveredPorts sets the value of NumberofDiscoveredPorts for the instance
func (instance *MS_SMHBA_SAS_Port) SetPropertyNumberofDiscoveredPorts(value uint32) (err error) {
	return instance.SetProperty("NumberofDiscoveredPorts", (value))
}

// GetNumberofDiscoveredPorts gets the value of NumberofDiscoveredPorts for the instance
func (instance *MS_SMHBA_SAS_Port) GetPropertyNumberofDiscoveredPorts() (value uint32, err error) {
	retValue, err := instance.GetProperty("NumberofDiscoveredPorts")
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

// SetNumberofPhys sets the value of NumberofPhys for the instance
func (instance *MS_SMHBA_SAS_Port) SetPropertyNumberofPhys(value uint32) (err error) {
	return instance.SetProperty("NumberofPhys", (value))
}

// GetNumberofPhys gets the value of NumberofPhys for the instance
func (instance *MS_SMHBA_SAS_Port) GetPropertyNumberofPhys() (value uint32, err error) {
	retValue, err := instance.GetProperty("NumberofPhys")
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

// SetPortProtocol sets the value of PortProtocol for the instance
func (instance *MS_SMHBA_SAS_Port) SetPropertyPortProtocol(value uint32) (err error) {
	return instance.SetProperty("PortProtocol", (value))
}

// GetPortProtocol gets the value of PortProtocol for the instance
func (instance *MS_SMHBA_SAS_Port) GetPropertyPortProtocol() (value uint32, err error) {
	retValue, err := instance.GetProperty("PortProtocol")
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
