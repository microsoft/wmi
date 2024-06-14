// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 6/13/2024 using wmigen
//      Source root.WebAdministration
//////////////////////////////////////////////
package webadministration

import (
	"github.com/microsoft/wmi/pkg/base/query"
	"github.com/microsoft/wmi/pkg/errors"
	cim "github.com/microsoft/wmi/pkg/wmiinstance"
	"reflect"
)

// SystemWebDeploymentSection struct
type SystemWebDeploymentSection struct {
	*ConfigurationSection

	//
	Retail bool
}

func NewSystemWebDeploymentSectionEx1(instance *cim.WmiInstance) (newInstance *SystemWebDeploymentSection, err error) {
	tmp, err := NewConfigurationSectionEx1(instance)

	if err != nil {
		return
	}
	newInstance = &SystemWebDeploymentSection{
		ConfigurationSection: tmp,
	}
	return
}

func NewSystemWebDeploymentSectionEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *SystemWebDeploymentSection, err error) {
	tmp, err := NewConfigurationSectionEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &SystemWebDeploymentSection{
		ConfigurationSection: tmp,
	}
	return
}

// SetRetail sets the value of Retail for the instance
func (instance *SystemWebDeploymentSection) SetPropertyRetail(value bool) (err error) {
	return instance.SetProperty("Retail", (value))
}

// GetRetail gets the value of Retail for the instance
func (instance *SystemWebDeploymentSection) GetPropertyRetail() (value bool, err error) {
	retValue, err := instance.GetProperty("Retail")
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
