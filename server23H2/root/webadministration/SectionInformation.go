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

// SectionInformation struct
type SectionInformation struct {
	*EmbeddedObject

	//
	EffectiveOverrideMode string

	//
	IsLocked bool

	//
	LockItem bool

	//
	OverrideMode string
}

func NewSectionInformationEx1(instance *cim.WmiInstance) (newInstance *SectionInformation, err error) {
	tmp, err := NewEmbeddedObjectEx1(instance)

	if err != nil {
		return
	}
	newInstance = &SectionInformation{
		EmbeddedObject: tmp,
	}
	return
}

func NewSectionInformationEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *SectionInformation, err error) {
	tmp, err := NewEmbeddedObjectEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &SectionInformation{
		EmbeddedObject: tmp,
	}
	return
}

// SetEffectiveOverrideMode sets the value of EffectiveOverrideMode for the instance
func (instance *SectionInformation) SetPropertyEffectiveOverrideMode(value string) (err error) {
	return instance.SetProperty("EffectiveOverrideMode", (value))
}

// GetEffectiveOverrideMode gets the value of EffectiveOverrideMode for the instance
func (instance *SectionInformation) GetPropertyEffectiveOverrideMode() (value string, err error) {
	retValue, err := instance.GetProperty("EffectiveOverrideMode")
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

// SetIsLocked sets the value of IsLocked for the instance
func (instance *SectionInformation) SetPropertyIsLocked(value bool) (err error) {
	return instance.SetProperty("IsLocked", (value))
}

// GetIsLocked gets the value of IsLocked for the instance
func (instance *SectionInformation) GetPropertyIsLocked() (value bool, err error) {
	retValue, err := instance.GetProperty("IsLocked")
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

// SetLockItem sets the value of LockItem for the instance
func (instance *SectionInformation) SetPropertyLockItem(value bool) (err error) {
	return instance.SetProperty("LockItem", (value))
}

// GetLockItem gets the value of LockItem for the instance
func (instance *SectionInformation) GetPropertyLockItem() (value bool, err error) {
	retValue, err := instance.GetProperty("LockItem")
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

// SetOverrideMode sets the value of OverrideMode for the instance
func (instance *SectionInformation) SetPropertyOverrideMode(value string) (err error) {
	return instance.SetProperty("OverrideMode", (value))
}

// GetOverrideMode gets the value of OverrideMode for the instance
func (instance *SectionInformation) GetPropertyOverrideMode() (value string, err error) {
	retValue, err := instance.GetProperty("OverrideMode")
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
