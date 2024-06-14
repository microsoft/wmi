// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 6/14/2024 using wmigen
//      Source root.WMI
//////////////////////////////////////////////
package wmi

import (
	"github.com/microsoft/wmi/pkg/base/query"
	"github.com/microsoft/wmi/pkg/errors"
	cim "github.com/microsoft/wmi/pkg/wmiinstance"
	"reflect"
)

// KerbSetPassword_End struct
type KerbSetPassword_End struct {
	*KerbSetPassword

	// Account Name
	AccountName string

	// Account Realm
	AccountRealm string

	// Client Name
	ClientName string

	// Client Realm
	ClientRealm string

	// KDC Address
	KdcAddress string

	// Status
	Status uint32
}

func NewKerbSetPassword_EndEx1(instance *cim.WmiInstance) (newInstance *KerbSetPassword_End, err error) {
	tmp, err := NewKerbSetPasswordEx1(instance)

	if err != nil {
		return
	}
	newInstance = &KerbSetPassword_End{
		KerbSetPassword: tmp,
	}
	return
}

func NewKerbSetPassword_EndEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *KerbSetPassword_End, err error) {
	tmp, err := NewKerbSetPasswordEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &KerbSetPassword_End{
		KerbSetPassword: tmp,
	}
	return
}

// SetAccountName sets the value of AccountName for the instance
func (instance *KerbSetPassword_End) SetPropertyAccountName(value string) (err error) {
	return instance.SetProperty("AccountName", (value))
}

// GetAccountName gets the value of AccountName for the instance
func (instance *KerbSetPassword_End) GetPropertyAccountName() (value string, err error) {
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

// SetAccountRealm sets the value of AccountRealm for the instance
func (instance *KerbSetPassword_End) SetPropertyAccountRealm(value string) (err error) {
	return instance.SetProperty("AccountRealm", (value))
}

// GetAccountRealm gets the value of AccountRealm for the instance
func (instance *KerbSetPassword_End) GetPropertyAccountRealm() (value string, err error) {
	retValue, err := instance.GetProperty("AccountRealm")
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

// SetClientName sets the value of ClientName for the instance
func (instance *KerbSetPassword_End) SetPropertyClientName(value string) (err error) {
	return instance.SetProperty("ClientName", (value))
}

// GetClientName gets the value of ClientName for the instance
func (instance *KerbSetPassword_End) GetPropertyClientName() (value string, err error) {
	retValue, err := instance.GetProperty("ClientName")
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

// SetClientRealm sets the value of ClientRealm for the instance
func (instance *KerbSetPassword_End) SetPropertyClientRealm(value string) (err error) {
	return instance.SetProperty("ClientRealm", (value))
}

// GetClientRealm gets the value of ClientRealm for the instance
func (instance *KerbSetPassword_End) GetPropertyClientRealm() (value string, err error) {
	retValue, err := instance.GetProperty("ClientRealm")
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

// SetKdcAddress sets the value of KdcAddress for the instance
func (instance *KerbSetPassword_End) SetPropertyKdcAddress(value string) (err error) {
	return instance.SetProperty("KdcAddress", (value))
}

// GetKdcAddress gets the value of KdcAddress for the instance
func (instance *KerbSetPassword_End) GetPropertyKdcAddress() (value string, err error) {
	retValue, err := instance.GetProperty("KdcAddress")
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
func (instance *KerbSetPassword_End) SetPropertyStatus(value uint32) (err error) {
	return instance.SetProperty("Status", (value))
}

// GetStatus gets the value of Status for the instance
func (instance *KerbSetPassword_End) GetPropertyStatus() (value uint32, err error) {
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
