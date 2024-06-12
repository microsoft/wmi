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

// MSiSCSIInitiator_PortalGroup struct
type MSiSCSIInitiator_PortalGroup struct {
	*cim.WmiInstance

	//
	Index uint32

	//
	Portals []MSiSCSIInitiator_Portal
}

func NewMSiSCSIInitiator_PortalGroupEx1(instance *cim.WmiInstance) (newInstance *MSiSCSIInitiator_PortalGroup, err error) {
	tmp, err := instance, nil

	if err != nil {
		return
	}
	newInstance = &MSiSCSIInitiator_PortalGroup{
		WmiInstance: tmp,
	}
	return
}

func NewMSiSCSIInitiator_PortalGroupEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *MSiSCSIInitiator_PortalGroup, err error) {
	tmp, err := instance.GetWmiInstance(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &MSiSCSIInitiator_PortalGroup{
		WmiInstance: tmp,
	}
	return
}

// SetIndex sets the value of Index for the instance
func (instance *MSiSCSIInitiator_PortalGroup) SetPropertyIndex(value uint32) (err error) {
	return instance.SetProperty("Index", (value))
}

// GetIndex gets the value of Index for the instance
func (instance *MSiSCSIInitiator_PortalGroup) GetPropertyIndex() (value uint32, err error) {
	retValue, err := instance.GetProperty("Index")
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

// SetPortals sets the value of Portals for the instance
func (instance *MSiSCSIInitiator_PortalGroup) SetPropertyPortals(value []MSiSCSIInitiator_Portal) (err error) {
	return instance.SetProperty("Portals", (value))
}

// GetPortals gets the value of Portals for the instance
func (instance *MSiSCSIInitiator_PortalGroup) GetPropertyPortals() (value []MSiSCSIInitiator_Portal, err error) {
	retValue, err := instance.GetProperty("Portals")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	for _, interfaceValue := range retValue.([]interface{}) {
		valuetmp, ok := interfaceValue.(MSiSCSIInitiator_Portal)
		if !ok {
			err = errors.Wrapf(errors.InvalidType, " MSiSCSIInitiator_Portal is Invalid. Expected %s", reflect.TypeOf(interfaceValue))
			return
		}
		value = append(value, MSiSCSIInitiator_Portal(valuetmp))
	}

	return
}
