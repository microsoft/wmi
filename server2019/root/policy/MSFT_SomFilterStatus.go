// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

// Author:
//
//	Auto Generated on 6/6/2024 using wmigen
//	Source root.Policy
//
// ////////////////////////////////////////////
package policy

import (
	"github.com/microsoft/wmi/pkg/base/instance"
	"github.com/microsoft/wmi/pkg/base/query"
	"github.com/microsoft/wmi/pkg/errors"
	cim "github.com/microsoft/wmi/pkg/wmiinstance"
	"reflect"
)

// MSFT_SomFilterStatus struct
type MSFT_SomFilterStatus struct {
	*cim.WmiInstance

	//
	ContainerAvailable bool

	//
	Domain string

	//
	SchemaAvailable bool
}

func NewMSFT_SomFilterStatusEx1(instance *cim.WmiInstance) (newInstance *MSFT_SomFilterStatus, err error) {
	tmp, err := instance, nil

	if err != nil {
		return
	}
	newInstance = &MSFT_SomFilterStatus{
		WmiInstance: tmp,
	}
	return
}

func NewMSFT_SomFilterStatusEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *MSFT_SomFilterStatus, err error) {
	tmp, err := instance.GetWmiInstance(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &MSFT_SomFilterStatus{
		WmiInstance: tmp,
	}
	return
}

// SetContainerAvailable sets the value of ContainerAvailable for the instance
func (instance *MSFT_SomFilterStatus) SetPropertyContainerAvailable(value bool) (err error) {
	return instance.SetProperty("ContainerAvailable", (value))
}

// GetContainerAvailable gets the value of ContainerAvailable for the instance
func (instance *MSFT_SomFilterStatus) GetPropertyContainerAvailable() (value bool, err error) {
	retValue, err := instance.GetProperty("ContainerAvailable")
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

// SetDomain sets the value of Domain for the instance
func (instance *MSFT_SomFilterStatus) SetPropertyDomain(value string) (err error) {
	return instance.SetProperty("Domain", (value))
}

// GetDomain gets the value of Domain for the instance
func (instance *MSFT_SomFilterStatus) GetPropertyDomain() (value string, err error) {
	retValue, err := instance.GetProperty("Domain")
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

// SetSchemaAvailable sets the value of SchemaAvailable for the instance
func (instance *MSFT_SomFilterStatus) SetPropertySchemaAvailable(value bool) (err error) {
	return instance.SetProperty("SchemaAvailable", (value))
}

// GetSchemaAvailable gets the value of SchemaAvailable for the instance
func (instance *MSFT_SomFilterStatus) GetPropertySchemaAvailable() (value bool, err error) {
	retValue, err := instance.GetProperty("SchemaAvailable")
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
