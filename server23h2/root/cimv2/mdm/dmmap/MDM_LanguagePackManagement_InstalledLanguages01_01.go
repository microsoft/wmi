// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 6/14/2024 using wmigen
//      Source root.CIMV2.mdm.dmmap
//////////////////////////////////////////////
package dmmap

import (
	"github.com/microsoft/wmi/pkg/base/instance"
	"github.com/microsoft/wmi/pkg/base/query"
	"github.com/microsoft/wmi/pkg/errors"
	cim "github.com/microsoft/wmi/pkg/wmiinstance"
	"reflect"
)

// MDM_LanguagePackManagement_InstalledLanguages01_01 struct
type MDM_LanguagePackManagement_InstalledLanguages01_01 struct {
	*cim.WmiInstance

	//
	InstanceID string

	//
	LanguageFeatures int32

	//
	ParentID string

	//
	Providers int32
}

func NewMDM_LanguagePackManagement_InstalledLanguages01_01Ex1(instance *cim.WmiInstance) (newInstance *MDM_LanguagePackManagement_InstalledLanguages01_01, err error) {
	tmp, err := instance, nil

	if err != nil {
		return
	}
	newInstance = &MDM_LanguagePackManagement_InstalledLanguages01_01{
		WmiInstance: tmp,
	}
	return
}

func NewMDM_LanguagePackManagement_InstalledLanguages01_01Ex6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *MDM_LanguagePackManagement_InstalledLanguages01_01, err error) {
	tmp, err := instance.GetWmiInstance(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &MDM_LanguagePackManagement_InstalledLanguages01_01{
		WmiInstance: tmp,
	}
	return
}

// SetInstanceID sets the value of InstanceID for the instance
func (instance *MDM_LanguagePackManagement_InstalledLanguages01_01) SetPropertyInstanceID(value string) (err error) {
	return instance.SetProperty("InstanceID", (value))
}

// GetInstanceID gets the value of InstanceID for the instance
func (instance *MDM_LanguagePackManagement_InstalledLanguages01_01) GetPropertyInstanceID() (value string, err error) {
	retValue, err := instance.GetProperty("InstanceID")
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

// SetLanguageFeatures sets the value of LanguageFeatures for the instance
func (instance *MDM_LanguagePackManagement_InstalledLanguages01_01) SetPropertyLanguageFeatures(value int32) (err error) {
	return instance.SetProperty("LanguageFeatures", (value))
}

// GetLanguageFeatures gets the value of LanguageFeatures for the instance
func (instance *MDM_LanguagePackManagement_InstalledLanguages01_01) GetPropertyLanguageFeatures() (value int32, err error) {
	retValue, err := instance.GetProperty("LanguageFeatures")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(int32)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " int32 is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = int32(valuetmp)

	return
}

// SetParentID sets the value of ParentID for the instance
func (instance *MDM_LanguagePackManagement_InstalledLanguages01_01) SetPropertyParentID(value string) (err error) {
	return instance.SetProperty("ParentID", (value))
}

// GetParentID gets the value of ParentID for the instance
func (instance *MDM_LanguagePackManagement_InstalledLanguages01_01) GetPropertyParentID() (value string, err error) {
	retValue, err := instance.GetProperty("ParentID")
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

// SetProviders sets the value of Providers for the instance
func (instance *MDM_LanguagePackManagement_InstalledLanguages01_01) SetPropertyProviders(value int32) (err error) {
	return instance.SetProperty("Providers", (value))
}

// GetProviders gets the value of Providers for the instance
func (instance *MDM_LanguagePackManagement_InstalledLanguages01_01) GetPropertyProviders() (value int32, err error) {
	retValue, err := instance.GetProperty("Providers")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(int32)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " int32 is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = int32(valuetmp)

	return
}
