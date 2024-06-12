// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 6/12/2024 using wmigen
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

// MSiSCSIInitiator_InitiatorClass struct
type MSiSCSIInitiator_InitiatorClass struct {
	*cim.WmiInstance

	//
	InitiatorName string
}

func NewMSiSCSIInitiator_InitiatorClassEx1(instance *cim.WmiInstance) (newInstance *MSiSCSIInitiator_InitiatorClass, err error) {
	tmp, err := instance, nil

	if err != nil {
		return
	}
	newInstance = &MSiSCSIInitiator_InitiatorClass{
		WmiInstance: tmp,
	}
	return
}

func NewMSiSCSIInitiator_InitiatorClassEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *MSiSCSIInitiator_InitiatorClass, err error) {
	tmp, err := instance.GetWmiInstance(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &MSiSCSIInitiator_InitiatorClass{
		WmiInstance: tmp,
	}
	return
}

// SetInitiatorName sets the value of InitiatorName for the instance
func (instance *MSiSCSIInitiator_InitiatorClass) SetPropertyInitiatorName(value string) (err error) {
	return instance.SetProperty("InitiatorName", (value))
}

// GetInitiatorName gets the value of InitiatorName for the instance
func (instance *MSiSCSIInitiator_InitiatorClass) GetPropertyInitiatorName() (value string, err error) {
	retValue, err := instance.GetProperty("InitiatorName")
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
