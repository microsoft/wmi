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

// WildcardRedirectElement struct
type WildcardRedirectElement struct {
	*CollectionElement

	//
	Destination string

	//
	Wildcard string
}

func NewWildcardRedirectElementEx1(instance *cim.WmiInstance) (newInstance *WildcardRedirectElement, err error) {
	tmp, err := NewCollectionElementEx1(instance)

	if err != nil {
		return
	}
	newInstance = &WildcardRedirectElement{
		CollectionElement: tmp,
	}
	return
}

func NewWildcardRedirectElementEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *WildcardRedirectElement, err error) {
	tmp, err := NewCollectionElementEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &WildcardRedirectElement{
		CollectionElement: tmp,
	}
	return
}

// SetDestination sets the value of Destination for the instance
func (instance *WildcardRedirectElement) SetPropertyDestination(value string) (err error) {
	return instance.SetProperty("Destination", (value))
}

// GetDestination gets the value of Destination for the instance
func (instance *WildcardRedirectElement) GetPropertyDestination() (value string, err error) {
	retValue, err := instance.GetProperty("Destination")
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

// SetWildcard sets the value of Wildcard for the instance
func (instance *WildcardRedirectElement) SetPropertyWildcard(value string) (err error) {
	return instance.SetProperty("Wildcard", (value))
}

// GetWildcard gets the value of Wildcard for the instance
func (instance *WildcardRedirectElement) GetPropertyWildcard() (value string, err error) {
	retValue, err := instance.GetProperty("Wildcard")
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
