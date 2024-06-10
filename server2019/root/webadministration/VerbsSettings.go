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

// VerbsSettings struct
type VerbsSettings struct {
	*EmbeddedObject

	//
	AllowUnlisted bool

	//
	ApplyToWebDAV bool

	//
	Verbs []VerbElement
}

func NewVerbsSettingsEx1(instance *cim.WmiInstance) (newInstance *VerbsSettings, err error) {
	tmp, err := NewEmbeddedObjectEx1(instance)

	if err != nil {
		return
	}
	newInstance = &VerbsSettings{
		EmbeddedObject: tmp,
	}
	return
}

func NewVerbsSettingsEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *VerbsSettings, err error) {
	tmp, err := NewEmbeddedObjectEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &VerbsSettings{
		EmbeddedObject: tmp,
	}
	return
}

// SetAllowUnlisted sets the value of AllowUnlisted for the instance
func (instance *VerbsSettings) SetPropertyAllowUnlisted(value bool) (err error) {
	return instance.SetProperty("AllowUnlisted", (value))
}

// GetAllowUnlisted gets the value of AllowUnlisted for the instance
func (instance *VerbsSettings) GetPropertyAllowUnlisted() (value bool, err error) {
	retValue, err := instance.GetProperty("AllowUnlisted")
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

// SetApplyToWebDAV sets the value of ApplyToWebDAV for the instance
func (instance *VerbsSettings) SetPropertyApplyToWebDAV(value bool) (err error) {
	return instance.SetProperty("ApplyToWebDAV", (value))
}

// GetApplyToWebDAV gets the value of ApplyToWebDAV for the instance
func (instance *VerbsSettings) GetPropertyApplyToWebDAV() (value bool, err error) {
	retValue, err := instance.GetProperty("ApplyToWebDAV")
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

// SetVerbs sets the value of Verbs for the instance
func (instance *VerbsSettings) SetPropertyVerbs(value []VerbElement) (err error) {
	return instance.SetProperty("Verbs", (value))
}

// GetVerbs gets the value of Verbs for the instance
func (instance *VerbsSettings) GetPropertyVerbs() (value []VerbElement, err error) {
	retValue, err := instance.GetProperty("Verbs")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	for _, interfaceValue := range retValue.([]interface{}) {
		valuetmp, ok := interfaceValue.(VerbElement)
		if !ok {
			err = errors.Wrapf(errors.InvalidType, " VerbElement is Invalid. Expected %s", reflect.TypeOf(interfaceValue))
			return
		}
		value = append(value, VerbElement(valuetmp))
	}

	return
}
