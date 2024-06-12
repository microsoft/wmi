// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 6/12/2024 using wmigen
//      Source root.SDDC.Management
//////////////////////////////////////////////
package management

import (
	"github.com/microsoft/wmi/pkg/base/instance"
	"github.com/microsoft/wmi/pkg/base/query"
	"github.com/microsoft/wmi/pkg/errors"
	cim "github.com/microsoft/wmi/pkg/wmiinstance"
	"reflect"
)

// SDDC_VmIntegrationService struct
type SDDC_VmIntegrationService struct {
	*cim.WmiInstance

	//
	IsEnabled bool

	//
	Name string
}

func NewSDDC_VmIntegrationServiceEx1(instance *cim.WmiInstance) (newInstance *SDDC_VmIntegrationService, err error) {
	tmp, err := instance, nil

	if err != nil {
		return
	}
	newInstance = &SDDC_VmIntegrationService{
		WmiInstance: tmp,
	}
	return
}

func NewSDDC_VmIntegrationServiceEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *SDDC_VmIntegrationService, err error) {
	tmp, err := instance.GetWmiInstance(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &SDDC_VmIntegrationService{
		WmiInstance: tmp,
	}
	return
}

// SetIsEnabled sets the value of IsEnabled for the instance
func (instance *SDDC_VmIntegrationService) SetPropertyIsEnabled(value bool) (err error) {
	return instance.SetProperty("IsEnabled", (value))
}

// GetIsEnabled gets the value of IsEnabled for the instance
func (instance *SDDC_VmIntegrationService) GetPropertyIsEnabled() (value bool, err error) {
	retValue, err := instance.GetProperty("IsEnabled")
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

// SetName sets the value of Name for the instance
func (instance *SDDC_VmIntegrationService) SetPropertyName(value string) (err error) {
	return instance.SetProperty("Name", (value))
}

// GetName gets the value of Name for the instance
func (instance *SDDC_VmIntegrationService) GetPropertyName() (value string, err error) {
	retValue, err := instance.GetProperty("Name")
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
