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

// VerbElement struct
type VerbElement struct {
	*CollectionElement

	//
	Allowed bool

	//
	Verb string
}

func NewVerbElementEx1(instance *cim.WmiInstance) (newInstance *VerbElement, err error) {
	tmp, err := NewCollectionElementEx1(instance)

	if err != nil {
		return
	}
	newInstance = &VerbElement{
		CollectionElement: tmp,
	}
	return
}

func NewVerbElementEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *VerbElement, err error) {
	tmp, err := NewCollectionElementEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &VerbElement{
		CollectionElement: tmp,
	}
	return
}

// SetAllowed sets the value of Allowed for the instance
func (instance *VerbElement) SetPropertyAllowed(value bool) (err error) {
	return instance.SetProperty("Allowed", (value))
}

// GetAllowed gets the value of Allowed for the instance
func (instance *VerbElement) GetPropertyAllowed() (value bool, err error) {
	retValue, err := instance.GetProperty("Allowed")
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

// SetVerb sets the value of Verb for the instance
func (instance *VerbElement) SetPropertyVerb(value string) (err error) {
	return instance.SetProperty("Verb", (value))
}

// GetVerb gets the value of Verb for the instance
func (instance *VerbElement) GetPropertyVerb() (value string, err error) {
	retValue, err := instance.GetProperty("Verb")
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
