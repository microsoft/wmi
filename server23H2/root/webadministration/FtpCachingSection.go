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

// FtpCachingSection struct
type FtpCachingSection struct {
	*ConfigurationSectionWithCollection

	//
	CredentialsCache FtpCredentialsCacheSettings
}

func NewFtpCachingSectionEx1(instance *cim.WmiInstance) (newInstance *FtpCachingSection, err error) {
	tmp, err := NewConfigurationSectionWithCollectionEx1(instance)

	if err != nil {
		return
	}
	newInstance = &FtpCachingSection{
		ConfigurationSectionWithCollection: tmp,
	}
	return
}

func NewFtpCachingSectionEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *FtpCachingSection, err error) {
	tmp, err := NewConfigurationSectionWithCollectionEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &FtpCachingSection{
		ConfigurationSectionWithCollection: tmp,
	}
	return
}

// SetCredentialsCache sets the value of CredentialsCache for the instance
func (instance *FtpCachingSection) SetPropertyCredentialsCache(value FtpCredentialsCacheSettings) (err error) {
	return instance.SetProperty("CredentialsCache", (value))
}

// GetCredentialsCache gets the value of CredentialsCache for the instance
func (instance *FtpCachingSection) GetPropertyCredentialsCache() (value FtpCredentialsCacheSettings, err error) {
	retValue, err := instance.GetProperty("CredentialsCache")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(FtpCredentialsCacheSettings)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " FtpCredentialsCacheSettings is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = FtpCredentialsCacheSettings(valuetmp)

	return
}
