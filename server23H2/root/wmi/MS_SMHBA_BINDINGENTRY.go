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

// MS_SMHBA_BINDINGENTRY struct
type MS_SMHBA_BINDINGENTRY struct {
	*cim.WmiInstance

	//
	LUID []uint8

	//
	PortLun MS_SMHBA_PORTLUN

	//
	ScsiId HBAScsiID

	//
	Status uint32

	//
	_type uint32
}

func NewMS_SMHBA_BINDINGENTRYEx1(instance *cim.WmiInstance) (newInstance *MS_SMHBA_BINDINGENTRY, err error) {
	tmp, err := instance, nil

	if err != nil {
		return
	}
	newInstance = &MS_SMHBA_BINDINGENTRY{
		WmiInstance: tmp,
	}
	return
}

func NewMS_SMHBA_BINDINGENTRYEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *MS_SMHBA_BINDINGENTRY, err error) {
	tmp, err := instance.GetWmiInstance(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &MS_SMHBA_BINDINGENTRY{
		WmiInstance: tmp,
	}
	return
}

// SetLUID sets the value of LUID for the instance
func (instance *MS_SMHBA_BINDINGENTRY) SetPropertyLUID(value []uint8) (err error) {
	return instance.SetProperty("LUID", (value))
}

// GetLUID gets the value of LUID for the instance
func (instance *MS_SMHBA_BINDINGENTRY) GetPropertyLUID() (value []uint8, err error) {
	retValue, err := instance.GetProperty("LUID")
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

// SetPortLun sets the value of PortLun for the instance
func (instance *MS_SMHBA_BINDINGENTRY) SetPropertyPortLun(value MS_SMHBA_PORTLUN) (err error) {
	return instance.SetProperty("PortLun", (value))
}

// GetPortLun gets the value of PortLun for the instance
func (instance *MS_SMHBA_BINDINGENTRY) GetPropertyPortLun() (value MS_SMHBA_PORTLUN, err error) {
	retValue, err := instance.GetProperty("PortLun")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(MS_SMHBA_PORTLUN)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " MS_SMHBA_PORTLUN is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = MS_SMHBA_PORTLUN(valuetmp)

	return
}

// SetScsiId sets the value of ScsiId for the instance
func (instance *MS_SMHBA_BINDINGENTRY) SetPropertyScsiId(value HBAScsiID) (err error) {
	return instance.SetProperty("ScsiId", (value))
}

// GetScsiId gets the value of ScsiId for the instance
func (instance *MS_SMHBA_BINDINGENTRY) GetPropertyScsiId() (value HBAScsiID, err error) {
	retValue, err := instance.GetProperty("ScsiId")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(HBAScsiID)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " HBAScsiID is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = HBAScsiID(valuetmp)

	return
}

// SetStatus sets the value of Status for the instance
func (instance *MS_SMHBA_BINDINGENTRY) SetPropertyStatus(value uint32) (err error) {
	return instance.SetProperty("Status", (value))
}

// GetStatus gets the value of Status for the instance
func (instance *MS_SMHBA_BINDINGENTRY) GetPropertyStatus() (value uint32, err error) {
	retValue, err := instance.GetProperty("Status")
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

// Settype sets the value of type for the instance
func (instance *MS_SMHBA_BINDINGENTRY) SetPropertytype(value uint32) (err error) {
	return instance.SetProperty("type", (value))
}

// Gettype gets the value of type for the instance
func (instance *MS_SMHBA_BINDINGENTRY) GetPropertytype() (value uint32, err error) {
	retValue, err := instance.GetProperty("type")
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
