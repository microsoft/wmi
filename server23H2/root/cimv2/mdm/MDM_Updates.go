// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 6/12/2024 using wmigen
//      Source root.CIMV2.mdm
//////////////////////////////////////////////
package mdm

import (
	"github.com/microsoft/wmi/pkg/base/instance"
	"github.com/microsoft/wmi/pkg/base/query"
	"github.com/microsoft/wmi/pkg/errors"
	cim "github.com/microsoft/wmi/pkg/wmiinstance"
	"reflect"
)

// MDM_Updates struct
type MDM_Updates struct {
	*cim.WmiInstance

	//
	AutoUpdatePolicy uint32

	//
	key uint32
}

func NewMDM_UpdatesEx1(instance *cim.WmiInstance) (newInstance *MDM_Updates, err error) {
	tmp, err := instance, nil

	if err != nil {
		return
	}
	newInstance = &MDM_Updates{
		WmiInstance: tmp,
	}
	return
}

func NewMDM_UpdatesEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *MDM_Updates, err error) {
	tmp, err := instance.GetWmiInstance(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &MDM_Updates{
		WmiInstance: tmp,
	}
	return
}

// SetAutoUpdatePolicy sets the value of AutoUpdatePolicy for the instance
func (instance *MDM_Updates) SetPropertyAutoUpdatePolicy(value uint32) (err error) {
	return instance.SetProperty("AutoUpdatePolicy", (value))
}

// GetAutoUpdatePolicy gets the value of AutoUpdatePolicy for the instance
func (instance *MDM_Updates) GetPropertyAutoUpdatePolicy() (value uint32, err error) {
	retValue, err := instance.GetProperty("AutoUpdatePolicy")
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

// Setkey sets the value of key for the instance
func (instance *MDM_Updates) SetPropertykey(value uint32) (err error) {
	return instance.SetProperty("key", (value))
}

// Getkey gets the value of key for the instance
func (instance *MDM_Updates) GetPropertykey() (value uint32, err error) {
	retValue, err := instance.GetProperty("key")
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
