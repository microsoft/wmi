// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 6/13/2024 using wmigen
//      Source root.WMI
//////////////////////////////////////////////
package wmi

import (
	"github.com/microsoft/wmi/pkg/base/query"
	"github.com/microsoft/wmi/pkg/errors"
	cim "github.com/microsoft/wmi/pkg/wmiinstance"
	"reflect"
)

// MSNdis_AtmMaxAal1PacketSize struct
type MSNdis_AtmMaxAal1PacketSize struct {
	*MSNdis

	//
	Active bool

	//
	InstanceName string

	//
	NdisAtmMaxAal1PacketSize uint32
}

func NewMSNdis_AtmMaxAal1PacketSizeEx1(instance *cim.WmiInstance) (newInstance *MSNdis_AtmMaxAal1PacketSize, err error) {
	tmp, err := NewMSNdisEx1(instance)

	if err != nil {
		return
	}
	newInstance = &MSNdis_AtmMaxAal1PacketSize{
		MSNdis: tmp,
	}
	return
}

func NewMSNdis_AtmMaxAal1PacketSizeEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *MSNdis_AtmMaxAal1PacketSize, err error) {
	tmp, err := NewMSNdisEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &MSNdis_AtmMaxAal1PacketSize{
		MSNdis: tmp,
	}
	return
}

// SetActive sets the value of Active for the instance
func (instance *MSNdis_AtmMaxAal1PacketSize) SetPropertyActive(value bool) (err error) {
	return instance.SetProperty("Active", (value))
}

// GetActive gets the value of Active for the instance
func (instance *MSNdis_AtmMaxAal1PacketSize) GetPropertyActive() (value bool, err error) {
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
func (instance *MSNdis_AtmMaxAal1PacketSize) SetPropertyInstanceName(value string) (err error) {
	return instance.SetProperty("InstanceName", (value))
}

// GetInstanceName gets the value of InstanceName for the instance
func (instance *MSNdis_AtmMaxAal1PacketSize) GetPropertyInstanceName() (value string, err error) {
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

// SetNdisAtmMaxAal1PacketSize sets the value of NdisAtmMaxAal1PacketSize for the instance
func (instance *MSNdis_AtmMaxAal1PacketSize) SetPropertyNdisAtmMaxAal1PacketSize(value uint32) (err error) {
	return instance.SetProperty("NdisAtmMaxAal1PacketSize", (value))
}

// GetNdisAtmMaxAal1PacketSize gets the value of NdisAtmMaxAal1PacketSize for the instance
func (instance *MSNdis_AtmMaxAal1PacketSize) GetPropertyNdisAtmMaxAal1PacketSize() (value uint32, err error) {
	retValue, err := instance.GetProperty("NdisAtmMaxAal1PacketSize")
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
