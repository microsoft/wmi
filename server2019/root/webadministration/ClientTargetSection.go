// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

// Author:
//
//	Auto Generated on 6/6/2024 using wmigen
//	Source root.WebAdministration
//
// ////////////////////////////////////////////
package webadministration

import (
	"github.com/microsoft/wmi/pkg/base/query"
	"github.com/microsoft/wmi/pkg/errors"
	cim "github.com/microsoft/wmi/pkg/wmiinstance"
	"reflect"
)

// ClientTargetSection struct
type ClientTargetSection struct {
	*ConfigurationSectionWithCollection

	//
	ClientTarget []ClientTarget
}

func NewClientTargetSectionEx1(instance *cim.WmiInstance) (newInstance *ClientTargetSection, err error) {
	tmp, err := NewConfigurationSectionWithCollectionEx1(instance)

	if err != nil {
		return
	}
	newInstance = &ClientTargetSection{
		ConfigurationSectionWithCollection: tmp,
	}
	return
}

func NewClientTargetSectionEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *ClientTargetSection, err error) {
	tmp, err := NewConfigurationSectionWithCollectionEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &ClientTargetSection{
		ConfigurationSectionWithCollection: tmp,
	}
	return
}

// SetClientTarget sets the value of ClientTarget for the instance
func (instance *ClientTargetSection) SetPropertyClientTarget(value []ClientTarget) (err error) {
	return instance.SetProperty("ClientTarget", (value))
}

// GetClientTarget gets the value of ClientTarget for the instance
func (instance *ClientTargetSection) GetPropertyClientTarget() (value []ClientTarget, err error) {
	retValue, err := instance.GetProperty("ClientTarget")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	for _, interfaceValue := range retValue.([]interface{}) {
		valuetmp, ok := interfaceValue.(ClientTarget)
		if !ok {
			err = errors.Wrapf(errors.InvalidType, " ClientTarget is Invalid. Expected %s", reflect.TypeOf(interfaceValue))
			return
		}
		value = append(value, ClientTarget(valuetmp))
	}

	return
}
