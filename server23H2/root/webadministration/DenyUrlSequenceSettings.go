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

// DenyUrlSequenceSettings struct
type DenyUrlSequenceSettings struct {
	*EmbeddedObject

	//
	DenyUrlSequences []Sequence
}

func NewDenyUrlSequenceSettingsEx1(instance *cim.WmiInstance) (newInstance *DenyUrlSequenceSettings, err error) {
	tmp, err := NewEmbeddedObjectEx1(instance)

	if err != nil {
		return
	}
	newInstance = &DenyUrlSequenceSettings{
		EmbeddedObject: tmp,
	}
	return
}

func NewDenyUrlSequenceSettingsEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *DenyUrlSequenceSettings, err error) {
	tmp, err := NewEmbeddedObjectEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &DenyUrlSequenceSettings{
		EmbeddedObject: tmp,
	}
	return
}

// SetDenyUrlSequences sets the value of DenyUrlSequences for the instance
func (instance *DenyUrlSequenceSettings) SetPropertyDenyUrlSequences(value []Sequence) (err error) {
	return instance.SetProperty("DenyUrlSequences", (value))
}

// GetDenyUrlSequences gets the value of DenyUrlSequences for the instance
func (instance *DenyUrlSequenceSettings) GetPropertyDenyUrlSequences() (value []Sequence, err error) {
	retValue, err := instance.GetProperty("DenyUrlSequences")
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
