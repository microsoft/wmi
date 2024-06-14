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

// MustUnderstandBehavior struct
type MustUnderstandBehavior struct {
	*Behavior

	// When true, all SOAP header with the MustUnderstand attribute that are not handled will cause the behavior to throw.
	ValidateMustUnderstand bool
}

func NewMustUnderstandBehaviorEx1(instance *cim.WmiInstance) (newInstance *MustUnderstandBehavior, err error) {
	tmp, err := NewBehaviorEx1(instance)

	if err != nil {
		return
	}
	newInstance = &MustUnderstandBehavior{
		Behavior: tmp,
	}
	return
}

func NewMustUnderstandBehaviorEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *MustUnderstandBehavior, err error) {
	tmp, err := NewBehaviorEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &MustUnderstandBehavior{
		Behavior: tmp,
	}
	return
}

// SetValidateMustUnderstand sets the value of ValidateMustUnderstand for the instance
func (instance *MustUnderstandBehavior) SetPropertyValidateMustUnderstand(value bool) (err error) {
	return instance.SetProperty("ValidateMustUnderstand", (value))
}

// GetValidateMustUnderstand gets the value of ValidateMustUnderstand for the instance
func (instance *MustUnderstandBehavior) GetPropertyValidateMustUnderstand() (value bool, err error) {
	retValue, err := instance.GetProperty("ValidateMustUnderstand")
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
