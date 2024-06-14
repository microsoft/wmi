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

// DenyQueryStringSequenceSettings struct
type DenyQueryStringSequenceSettings struct {
	*EmbeddedObject

	//
	DenyQueryStringSequences []Sequence
}

func NewDenyQueryStringSequenceSettingsEx1(instance *cim.WmiInstance) (newInstance *DenyQueryStringSequenceSettings, err error) {
	tmp, err := NewEmbeddedObjectEx1(instance)

	if err != nil {
		return
	}
	newInstance = &DenyQueryStringSequenceSettings{
		EmbeddedObject: tmp,
	}
	return
}

func NewDenyQueryStringSequenceSettingsEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *DenyQueryStringSequenceSettings, err error) {
	tmp, err := NewEmbeddedObjectEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &DenyQueryStringSequenceSettings{
		EmbeddedObject: tmp,
	}
	return
}

// SetDenyQueryStringSequences sets the value of DenyQueryStringSequences for the instance
func (instance *DenyQueryStringSequenceSettings) SetPropertyDenyQueryStringSequences(value []Sequence) (err error) {
	return instance.SetProperty("DenyQueryStringSequences", (value))
}

// GetDenyQueryStringSequences gets the value of DenyQueryStringSequences for the instance
func (instance *DenyQueryStringSequenceSettings) GetPropertyDenyQueryStringSequences() (value []Sequence, err error) {
	retValue, err := instance.GetProperty("DenyQueryStringSequences")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	for _, interfaceValue := range retValue.([]interface{}) {
		valuetmp, ok := interfaceValue.(Sequence)
		if !ok {
			err = errors.Wrapf(errors.InvalidType, " Sequence is Invalid. Expected %s", reflect.TypeOf(interfaceValue))
			return
		}
		value = append(value, Sequence(valuetmp))
	}

	return
}
