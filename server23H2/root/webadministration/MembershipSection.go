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

// MembershipSection struct
type MembershipSection struct {
	*ConfigurationSectionWithCollection

	//
	DefaultProvider string

	//
	HashAlgorithmType string

	//
	Providers ProvidersSettings

	//
	UserIsOnlineTimeWindow string
}

func NewMembershipSectionEx1(instance *cim.WmiInstance) (newInstance *MembershipSection, err error) {
	tmp, err := NewConfigurationSectionWithCollectionEx1(instance)

	if err != nil {
		return
	}
	newInstance = &MembershipSection{
		ConfigurationSectionWithCollection: tmp,
	}
	return
}

func NewMembershipSectionEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *MembershipSection, err error) {
	tmp, err := NewConfigurationSectionWithCollectionEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &MembershipSection{
		ConfigurationSectionWithCollection: tmp,
	}
	return
}

// SetDefaultProvider sets the value of DefaultProvider for the instance
func (instance *MembershipSection) SetPropertyDefaultProvider(value string) (err error) {
	return instance.SetProperty("DefaultProvider", (value))
}

// GetDefaultProvider gets the value of DefaultProvider for the instance
func (instance *MembershipSection) GetPropertyDefaultProvider() (value string, err error) {
	retValue, err := instance.GetProperty("DefaultProvider")
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

// SetHashAlgorithmType sets the value of HashAlgorithmType for the instance
func (instance *MembershipSection) SetPropertyHashAlgorithmType(value string) (err error) {
	return instance.SetProperty("HashAlgorithmType", (value))
}

// GetHashAlgorithmType gets the value of HashAlgorithmType for the instance
func (instance *MembershipSection) GetPropertyHashAlgorithmType() (value string, err error) {
	retValue, err := instance.GetProperty("HashAlgorithmType")
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
func (instance *MembershipSection) SetPropertyProviders(value ProvidersSettings) (err error) {
	return instance.SetProperty("Providers", (value))
}

// GetProviders gets the value of Providers for the instance
func (instance *MembershipSection) GetPropertyProviders() (value ProvidersSettings, err error) {
	retValue, err := instance.GetProperty("Providers")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(ProvidersSettings)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " ProvidersSettings is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = ProvidersSettings(valuetmp)

	return
}

// SetUserIsOnlineTimeWindow sets the value of UserIsOnlineTimeWindow for the instance
func (instance *MembershipSection) SetPropertyUserIsOnlineTimeWindow(value string) (err error) {
	return instance.SetProperty("UserIsOnlineTimeWindow", (value))
}

// GetUserIsOnlineTimeWindow gets the value of UserIsOnlineTimeWindow for the instance
func (instance *MembershipSection) GetPropertyUserIsOnlineTimeWindow() (value string, err error) {
	retValue, err := instance.GetProperty("UserIsOnlineTimeWindow")
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
