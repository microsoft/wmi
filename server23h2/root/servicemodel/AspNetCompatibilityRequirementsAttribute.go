// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 6/14/2024 using wmigen
//      Source root.ServiceModel
//////////////////////////////////////////////
package servicemodel

import (
	"github.com/microsoft/wmi/pkg/base/query"
	"github.com/microsoft/wmi/pkg/errors"
	cim "github.com/microsoft/wmi/pkg/wmiinstance"
	"reflect"
)

// AspNetCompatibilityRequirementsAttribute struct
type AspNetCompatibilityRequirementsAttribute struct {
	*Behavior

	// Indicates if Asp.Net compatibility mode is active.
	RequirementsMode string
}

func NewAspNetCompatibilityRequirementsAttributeEx1(instance *cim.WmiInstance) (newInstance *AspNetCompatibilityRequirementsAttribute, err error) {
	tmp, err := NewBehaviorEx1(instance)

	if err != nil {
		return
	}
	newInstance = &AspNetCompatibilityRequirementsAttribute{
		Behavior: tmp,
	}
	return
}

func NewAspNetCompatibilityRequirementsAttributeEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *AspNetCompatibilityRequirementsAttribute, err error) {
	tmp, err := NewBehaviorEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &AspNetCompatibilityRequirementsAttribute{
		Behavior: tmp,
	}
	return
}

// SetRequirementsMode sets the value of RequirementsMode for the instance
func (instance *AspNetCompatibilityRequirementsAttribute) SetPropertyRequirementsMode(value string) (err error) {
	return instance.SetProperty("RequirementsMode", (value))
}

// GetRequirementsMode gets the value of RequirementsMode for the instance
func (instance *AspNetCompatibilityRequirementsAttribute) GetPropertyRequirementsMode() (value string, err error) {
	retValue, err := instance.GetProperty("RequirementsMode")
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
