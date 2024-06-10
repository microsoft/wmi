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

// RequestCachingSection struct
type RequestCachingSection struct {
	*ConfigurationSection

	//
	DefaultFtpCachePolicy FtpCachePolicySettings

	//
	DefaultHttpCachePolicy HttpCachePolicySettings

	//
	DefaultPolicyLevel int32

	//
	DisableAllCaching bool

	//
	IsPrivateCache bool

	//
	UnspecifiedMaximumAge string
}

func NewRequestCachingSectionEx1(instance *cim.WmiInstance) (newInstance *RequestCachingSection, err error) {
	tmp, err := NewConfigurationSectionEx1(instance)

	if err != nil {
		return
	}
	newInstance = &RequestCachingSection{
		ConfigurationSection: tmp,
	}
	return
}

func NewRequestCachingSectionEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *RequestCachingSection, err error) {
	tmp, err := NewConfigurationSectionEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &RequestCachingSection{
		ConfigurationSection: tmp,
	}
	return
}

// SetDefaultFtpCachePolicy sets the value of DefaultFtpCachePolicy for the instance
func (instance *RequestCachingSection) SetPropertyDefaultFtpCachePolicy(value FtpCachePolicySettings) (err error) {
	return instance.SetProperty("DefaultFtpCachePolicy", (value))
}

// GetDefaultFtpCachePolicy gets the value of DefaultFtpCachePolicy for the instance
func (instance *RequestCachingSection) GetPropertyDefaultFtpCachePolicy() (value FtpCachePolicySettings, err error) {
	retValue, err := instance.GetProperty("DefaultFtpCachePolicy")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(FtpCachePolicySettings)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " FtpCachePolicySettings is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = FtpCachePolicySettings(valuetmp)

	return
}

// SetDefaultHttpCachePolicy sets the value of DefaultHttpCachePolicy for the instance
func (instance *RequestCachingSection) SetPropertyDefaultHttpCachePolicy(value HttpCachePolicySettings) (err error) {
	return instance.SetProperty("DefaultHttpCachePolicy", (value))
}

// GetDefaultHttpCachePolicy gets the value of DefaultHttpCachePolicy for the instance
func (instance *RequestCachingSection) GetPropertyDefaultHttpCachePolicy() (value HttpCachePolicySettings, err error) {
	retValue, err := instance.GetProperty("DefaultHttpCachePolicy")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(HttpCachePolicySettings)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " HttpCachePolicySettings is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = HttpCachePolicySettings(valuetmp)

	return
}

// SetDefaultPolicyLevel sets the value of DefaultPolicyLevel for the instance
func (instance *RequestCachingSection) SetPropertyDefaultPolicyLevel(value int32) (err error) {
	return instance.SetProperty("DefaultPolicyLevel", (value))
}

// GetDefaultPolicyLevel gets the value of DefaultPolicyLevel for the instance
func (instance *RequestCachingSection) GetPropertyDefaultPolicyLevel() (value int32, err error) {
	retValue, err := instance.GetProperty("DefaultPolicyLevel")
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

// SetDisableAllCaching sets the value of DisableAllCaching for the instance
func (instance *RequestCachingSection) SetPropertyDisableAllCaching(value bool) (err error) {
	return instance.SetProperty("DisableAllCaching", (value))
}

// GetDisableAllCaching gets the value of DisableAllCaching for the instance
func (instance *RequestCachingSection) GetPropertyDisableAllCaching() (value bool, err error) {
	retValue, err := instance.GetProperty("DisableAllCaching")
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

// SetIsPrivateCache sets the value of IsPrivateCache for the instance
func (instance *RequestCachingSection) SetPropertyIsPrivateCache(value bool) (err error) {
	return instance.SetProperty("IsPrivateCache", (value))
}

// GetIsPrivateCache gets the value of IsPrivateCache for the instance
func (instance *RequestCachingSection) GetPropertyIsPrivateCache() (value bool, err error) {
	retValue, err := instance.GetProperty("IsPrivateCache")
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

// SetUnspecifiedMaximumAge sets the value of UnspecifiedMaximumAge for the instance
func (instance *RequestCachingSection) SetPropertyUnspecifiedMaximumAge(value string) (err error) {
	return instance.SetProperty("UnspecifiedMaximumAge", (value))
}

// GetUnspecifiedMaximumAge gets the value of UnspecifiedMaximumAge for the instance
func (instance *RequestCachingSection) GetPropertyUnspecifiedMaximumAge() (value string, err error) {
	retValue, err := instance.GetProperty("UnspecifiedMaximumAge")
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
