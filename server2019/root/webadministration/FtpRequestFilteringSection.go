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

// FtpRequestFilteringSection struct
type FtpRequestFilteringSection struct {
	*ConfigurationSectionWithCollection

	//
	AllowHighBitCharacters bool

	//
	DenyUrlSequences DenyUrlSequenceSettings

	//
	FileExtensions FtpFileExtensionsSettings

	//
	HiddenSegments FtpHiddenSegmentSettings

	//
	RequestLimits FtpRequestLimitsSettings
}

func NewFtpRequestFilteringSectionEx1(instance *cim.WmiInstance) (newInstance *FtpRequestFilteringSection, err error) {
	tmp, err := NewConfigurationSectionWithCollectionEx1(instance)

	if err != nil {
		return
	}
	newInstance = &FtpRequestFilteringSection{
		ConfigurationSectionWithCollection: tmp,
	}
	return
}

func NewFtpRequestFilteringSectionEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *FtpRequestFilteringSection, err error) {
	tmp, err := NewConfigurationSectionWithCollectionEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &FtpRequestFilteringSection{
		ConfigurationSectionWithCollection: tmp,
	}
	return
}

// SetAllowHighBitCharacters sets the value of AllowHighBitCharacters for the instance
func (instance *FtpRequestFilteringSection) SetPropertyAllowHighBitCharacters(value bool) (err error) {
	return instance.SetProperty("AllowHighBitCharacters", (value))
}

// GetAllowHighBitCharacters gets the value of AllowHighBitCharacters for the instance
func (instance *FtpRequestFilteringSection) GetPropertyAllowHighBitCharacters() (value bool, err error) {
	retValue, err := instance.GetProperty("AllowHighBitCharacters")
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

// SetDenyUrlSequences sets the value of DenyUrlSequences for the instance
func (instance *FtpRequestFilteringSection) SetPropertyDenyUrlSequences(value DenyUrlSequenceSettings) (err error) {
	return instance.SetProperty("DenyUrlSequences", (value))
}

// GetDenyUrlSequences gets the value of DenyUrlSequences for the instance
func (instance *FtpRequestFilteringSection) GetPropertyDenyUrlSequences() (value DenyUrlSequenceSettings, err error) {
	retValue, err := instance.GetProperty("DenyUrlSequences")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(DenyUrlSequenceSettings)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " DenyUrlSequenceSettings is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = DenyUrlSequenceSettings(valuetmp)

	return
}

// SetFileExtensions sets the value of FileExtensions for the instance
func (instance *FtpRequestFilteringSection) SetPropertyFileExtensions(value FtpFileExtensionsSettings) (err error) {
	return instance.SetProperty("FileExtensions", (value))
}

// GetFileExtensions gets the value of FileExtensions for the instance
func (instance *FtpRequestFilteringSection) GetPropertyFileExtensions() (value FtpFileExtensionsSettings, err error) {
	retValue, err := instance.GetProperty("FileExtensions")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(FtpFileExtensionsSettings)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " FtpFileExtensionsSettings is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = FtpFileExtensionsSettings(valuetmp)

	return
}

// SetHiddenSegments sets the value of HiddenSegments for the instance
func (instance *FtpRequestFilteringSection) SetPropertyHiddenSegments(value FtpHiddenSegmentSettings) (err error) {
	return instance.SetProperty("HiddenSegments", (value))
}

// GetHiddenSegments gets the value of HiddenSegments for the instance
func (instance *FtpRequestFilteringSection) GetPropertyHiddenSegments() (value FtpHiddenSegmentSettings, err error) {
	retValue, err := instance.GetProperty("HiddenSegments")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(FtpHiddenSegmentSettings)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " FtpHiddenSegmentSettings is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = FtpHiddenSegmentSettings(valuetmp)

	return
}

// SetRequestLimits sets the value of RequestLimits for the instance
func (instance *FtpRequestFilteringSection) SetPropertyRequestLimits(value FtpRequestLimitsSettings) (err error) {
	return instance.SetProperty("RequestLimits", (value))
}

// GetRequestLimits gets the value of RequestLimits for the instance
func (instance *FtpRequestFilteringSection) GetPropertyRequestLimits() (value FtpRequestLimitsSettings, err error) {
	retValue, err := instance.GetProperty("RequestLimits")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(FtpRequestLimitsSettings)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " FtpRequestLimitsSettings is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = FtpRequestLimitsSettings(valuetmp)

	return
}
