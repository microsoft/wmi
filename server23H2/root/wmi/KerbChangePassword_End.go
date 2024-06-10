// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

// Author:
//
//	Auto Generated on 6/6/2024 using wmigen
//	Source root.WMI
//
// ////////////////////////////////////////////
package wmi

import (
	"github.com/microsoft/wmi/pkg/base/query"
	"github.com/microsoft/wmi/pkg/errors"
	cim "github.com/microsoft/wmi/pkg/wmiinstance"
	"reflect"
)

// KerbChangePassword_End struct
type KerbChangePassword_End struct {
	*KerbChangePassword

	// Account Name
	AccountName string

	// Account Realm
	DomainName string

	// Status
	Status uint32
}

func NewKerbChangePassword_EndEx1(instance *cim.WmiInstance) (newInstance *KerbChangePassword_End, err error) {
	tmp, err := NewKerbChangePasswordEx1(instance)

	if err != nil {
		return
	}
	newInstance = &KerbChangePassword_End{
		KerbChangePassword: tmp,
	}
	return
}

func NewKerbChangePassword_EndEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *KerbChangePassword_End, err error) {
	tmp, err := NewKerbChangePasswordEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &KerbChangePassword_End{
		KerbChangePassword: tmp,
	}
	return
}

// SetAccountName sets the value of AccountName for the instance
func (instance *KerbChangePassword_End) SetPropertyAccountName(value string) (err error) {
	return instance.SetProperty("AccountName", (value))
}

// GetAccountName gets the value of AccountName for the instance
func (instance *KerbChangePassword_End) GetPropertyAccountName() (value string, err error) {
	retValue, err := instance.GetProperty("AccountName")
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

// SetDomainName sets the value of DomainName for the instance
func (instance *KerbChangePassword_End) SetPropertyDomainName(value string) (err error) {
	return instance.SetProperty("DomainName", (value))
}

// GetDomainName gets the value of DomainName for the instance
func (instance *KerbChangePassword_End) GetPropertyDomainName() (value string, err error) {
	retValue, err := instance.GetProperty("DomainName")
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

// SetStatus sets the value of Status for the instance
func (instance *KerbChangePassword_End) SetPropertyStatus(value uint32) (err error) {
	return instance.SetProperty("Status", (value))
}

// GetStatus gets the value of Status for the instance
func (instance *KerbChangePassword_End) GetPropertyStatus() (value uint32, err error) {
	retValue, err := instance.GetProperty("Status")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(uint32)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " uint32 is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = uint32(valuetmp)

	return
}
