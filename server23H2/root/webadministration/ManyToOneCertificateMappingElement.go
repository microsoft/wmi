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

// ManyToOneCertificateMappingElement struct
type ManyToOneCertificateMappingElement struct {
	*CollectionElement

	//
	Description string

	//
	Enabled bool

	//
	Name string

	//
	Password string

	//
	PermissionMode int32

	//
	Rules ManyToOneCertificateRuleSettings

	//
	UserName string
}

func NewManyToOneCertificateMappingElementEx1(instance *cim.WmiInstance) (newInstance *ManyToOneCertificateMappingElement, err error) {
	tmp, err := NewCollectionElementEx1(instance)

	if err != nil {
		return
	}
	newInstance = &ManyToOneCertificateMappingElement{
		CollectionElement: tmp,
	}
	return
}

func NewManyToOneCertificateMappingElementEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *ManyToOneCertificateMappingElement, err error) {
	tmp, err := NewCollectionElementEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &ManyToOneCertificateMappingElement{
		CollectionElement: tmp,
	}
	return
}

// SetDescription sets the value of Description for the instance
func (instance *ManyToOneCertificateMappingElement) SetPropertyDescription(value string) (err error) {
	return instance.SetProperty("Description", (value))
}

// GetDescription gets the value of Description for the instance
func (instance *ManyToOneCertificateMappingElement) GetPropertyDescription() (value string, err error) {
	retValue, err := instance.GetProperty("Description")
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

// SetEnabled sets the value of Enabled for the instance
func (instance *ManyToOneCertificateMappingElement) SetPropertyEnabled(value bool) (err error) {
	return instance.SetProperty("Enabled", (value))
}

// GetEnabled gets the value of Enabled for the instance
func (instance *ManyToOneCertificateMappingElement) GetPropertyEnabled() (value bool, err error) {
	retValue, err := instance.GetProperty("Enabled")
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

// SetName sets the value of Name for the instance
func (instance *ManyToOneCertificateMappingElement) SetPropertyName(value string) (err error) {
	return instance.SetProperty("Name", (value))
}

// GetName gets the value of Name for the instance
func (instance *ManyToOneCertificateMappingElement) GetPropertyName() (value string, err error) {
	retValue, err := instance.GetProperty("Name")
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

// SetPassword sets the value of Password for the instance
func (instance *ManyToOneCertificateMappingElement) SetPropertyPassword(value string) (err error) {
	return instance.SetProperty("Password", (value))
}

// GetPassword gets the value of Password for the instance
func (instance *ManyToOneCertificateMappingElement) GetPropertyPassword() (value string, err error) {
	retValue, err := instance.GetProperty("Password")
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

// SetPermissionMode sets the value of PermissionMode for the instance
func (instance *ManyToOneCertificateMappingElement) SetPropertyPermissionMode(value int32) (err error) {
	return instance.SetProperty("PermissionMode", (value))
}

// GetPermissionMode gets the value of PermissionMode for the instance
func (instance *ManyToOneCertificateMappingElement) GetPropertyPermissionMode() (value int32, err error) {
	retValue, err := instance.GetProperty("PermissionMode")
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

// SetRules sets the value of Rules for the instance
func (instance *ManyToOneCertificateMappingElement) SetPropertyRules(value ManyToOneCertificateRuleSettings) (err error) {
	return instance.SetProperty("Rules", (value))
}

// GetRules gets the value of Rules for the instance
func (instance *ManyToOneCertificateMappingElement) GetPropertyRules() (value ManyToOneCertificateRuleSettings, err error) {
	retValue, err := instance.GetProperty("Rules")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(ManyToOneCertificateRuleSettings)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " ManyToOneCertificateRuleSettings is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = ManyToOneCertificateRuleSettings(valuetmp)

	return
}

// SetUserName sets the value of UserName for the instance
func (instance *ManyToOneCertificateMappingElement) SetPropertyUserName(value string) (err error) {
	return instance.SetProperty("UserName", (value))
}

// GetUserName gets the value of UserName for the instance
func (instance *ManyToOneCertificateMappingElement) GetPropertyUserName() (value string, err error) {
	retValue, err := instance.GetProperty("UserName")
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
