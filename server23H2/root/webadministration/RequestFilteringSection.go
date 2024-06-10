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

// RequestFilteringSection struct
type RequestFilteringSection struct {
	*ConfigurationSectionWithCollection

	//
	AllowDoubleEscaping bool

	//
	AllowHighBitCharacters bool

	//
	AlwaysAllowedQueryStrings AlwaysAllowedQueryStringSettings

	//
	AlwaysAllowedUrls AwaysAllowedUrlSettings

	//
	DenyQueryStringSequences DenyQueryStringSequenceSettings

	//
	DenyUrlSequences DenyUrlSequenceSettings

	//
	FileExtensions FileExtensionsSettings

	//
	FilteringRules FilteringRulesSettings

	//
	HiddenSegments HiddenSegmentSettings

	//
	RequestLimits RequestLimitsElement

	//
	UnescapeQueryString bool

	//
	Verbs VerbsSettings
}

func NewRequestFilteringSectionEx1(instance *cim.WmiInstance) (newInstance *RequestFilteringSection, err error) {
	tmp, err := NewConfigurationSectionWithCollectionEx1(instance)

	if err != nil {
		return
	}
	newInstance = &RequestFilteringSection{
		ConfigurationSectionWithCollection: tmp,
	}
	return
}

func NewRequestFilteringSectionEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *RequestFilteringSection, err error) {
	tmp, err := NewConfigurationSectionWithCollectionEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &RequestFilteringSection{
		ConfigurationSectionWithCollection: tmp,
	}
	return
}

// SetAllowDoubleEscaping sets the value of AllowDoubleEscaping for the instance
func (instance *RequestFilteringSection) SetPropertyAllowDoubleEscaping(value bool) (err error) {
	return instance.SetProperty("AllowDoubleEscaping", (value))
}

// GetAllowDoubleEscaping gets the value of AllowDoubleEscaping for the instance
func (instance *RequestFilteringSection) GetPropertyAllowDoubleEscaping() (value bool, err error) {
	retValue, err := instance.GetProperty("AllowDoubleEscaping")
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

// SetAllowHighBitCharacters sets the value of AllowHighBitCharacters for the instance
func (instance *RequestFilteringSection) SetPropertyAllowHighBitCharacters(value bool) (err error) {
	return instance.SetProperty("AllowHighBitCharacters", (value))
}

// GetAllowHighBitCharacters gets the value of AllowHighBitCharacters for the instance
func (instance *RequestFilteringSection) GetPropertyAllowHighBitCharacters() (value bool, err error) {
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

// SetAlwaysAllowedQueryStrings sets the value of AlwaysAllowedQueryStrings for the instance
func (instance *RequestFilteringSection) SetPropertyAlwaysAllowedQueryStrings(value AlwaysAllowedQueryStringSettings) (err error) {
	return instance.SetProperty("AlwaysAllowedQueryStrings", (value))
}

// GetAlwaysAllowedQueryStrings gets the value of AlwaysAllowedQueryStrings for the instance
func (instance *RequestFilteringSection) GetPropertyAlwaysAllowedQueryStrings() (value AlwaysAllowedQueryStringSettings, err error) {
	retValue, err := instance.GetProperty("AlwaysAllowedQueryStrings")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(AlwaysAllowedQueryStringSettings)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " AlwaysAllowedQueryStringSettings is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = AlwaysAllowedQueryStringSettings(valuetmp)

	return
}

// SetAlwaysAllowedUrls sets the value of AlwaysAllowedUrls for the instance
func (instance *RequestFilteringSection) SetPropertyAlwaysAllowedUrls(value AwaysAllowedUrlSettings) (err error) {
	return instance.SetProperty("AlwaysAllowedUrls", (value))
}

// GetAlwaysAllowedUrls gets the value of AlwaysAllowedUrls for the instance
func (instance *RequestFilteringSection) GetPropertyAlwaysAllowedUrls() (value AwaysAllowedUrlSettings, err error) {
	retValue, err := instance.GetProperty("AlwaysAllowedUrls")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(AwaysAllowedUrlSettings)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " AwaysAllowedUrlSettings is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = AwaysAllowedUrlSettings(valuetmp)

	return
}

// SetDenyQueryStringSequences sets the value of DenyQueryStringSequences for the instance
func (instance *RequestFilteringSection) SetPropertyDenyQueryStringSequences(value DenyQueryStringSequenceSettings) (err error) {
	return instance.SetProperty("DenyQueryStringSequences", (value))
}

// GetDenyQueryStringSequences gets the value of DenyQueryStringSequences for the instance
func (instance *RequestFilteringSection) GetPropertyDenyQueryStringSequences() (value DenyQueryStringSequenceSettings, err error) {
	retValue, err := instance.GetProperty("DenyQueryStringSequences")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(DenyQueryStringSequenceSettings)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " DenyQueryStringSequenceSettings is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = DenyQueryStringSequenceSettings(valuetmp)

	return
}

// SetDenyUrlSequences sets the value of DenyUrlSequences for the instance
func (instance *RequestFilteringSection) SetPropertyDenyUrlSequences(value DenyUrlSequenceSettings) (err error) {
	return instance.SetProperty("DenyUrlSequences", (value))
}

// GetDenyUrlSequences gets the value of DenyUrlSequences for the instance
func (instance *RequestFilteringSection) GetPropertyDenyUrlSequences() (value DenyUrlSequenceSettings, err error) {
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
func (instance *RequestFilteringSection) SetPropertyFileExtensions(value FileExtensionsSettings) (err error) {
	return instance.SetProperty("FileExtensions", (value))
}

// GetFileExtensions gets the value of FileExtensions for the instance
func (instance *RequestFilteringSection) GetPropertyFileExtensions() (value FileExtensionsSettings, err error) {
	retValue, err := instance.GetProperty("FileExtensions")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(FileExtensionsSettings)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " FileExtensionsSettings is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = FileExtensionsSettings(valuetmp)

	return
}

// SetFilteringRules sets the value of FilteringRules for the instance
func (instance *RequestFilteringSection) SetPropertyFilteringRules(value FilteringRulesSettings) (err error) {
	return instance.SetProperty("FilteringRules", (value))
}

// GetFilteringRules gets the value of FilteringRules for the instance
func (instance *RequestFilteringSection) GetPropertyFilteringRules() (value FilteringRulesSettings, err error) {
	retValue, err := instance.GetProperty("FilteringRules")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(FilteringRulesSettings)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " FilteringRulesSettings is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = FilteringRulesSettings(valuetmp)

	return
}

// SetHiddenSegments sets the value of HiddenSegments for the instance
func (instance *RequestFilteringSection) SetPropertyHiddenSegments(value HiddenSegmentSettings) (err error) {
	return instance.SetProperty("HiddenSegments", (value))
}

// GetHiddenSegments gets the value of HiddenSegments for the instance
func (instance *RequestFilteringSection) GetPropertyHiddenSegments() (value HiddenSegmentSettings, err error) {
	retValue, err := instance.GetProperty("HiddenSegments")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(HiddenSegmentSettings)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " HiddenSegmentSettings is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = HiddenSegmentSettings(valuetmp)

	return
}

// SetRequestLimits sets the value of RequestLimits for the instance
func (instance *RequestFilteringSection) SetPropertyRequestLimits(value RequestLimitsElement) (err error) {
	return instance.SetProperty("RequestLimits", (value))
}

// GetRequestLimits gets the value of RequestLimits for the instance
func (instance *RequestFilteringSection) GetPropertyRequestLimits() (value RequestLimitsElement, err error) {
	retValue, err := instance.GetProperty("RequestLimits")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(RequestLimitsElement)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " RequestLimitsElement is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = RequestLimitsElement(valuetmp)

	return
}

// SetUnescapeQueryString sets the value of UnescapeQueryString for the instance
func (instance *RequestFilteringSection) SetPropertyUnescapeQueryString(value bool) (err error) {
	return instance.SetProperty("UnescapeQueryString", (value))
}

// GetUnescapeQueryString gets the value of UnescapeQueryString for the instance
func (instance *RequestFilteringSection) GetPropertyUnescapeQueryString() (value bool, err error) {
	retValue, err := instance.GetProperty("UnescapeQueryString")
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
func (instance *RequestFilteringSection) SetPropertyVerbs(value VerbsSettings) (err error) {
	return instance.SetProperty("Verbs", (value))
}

// GetVerbs gets the value of Verbs for the instance
func (instance *RequestFilteringSection) GetPropertyVerbs() (value VerbsSettings, err error) {
	retValue, err := instance.GetProperty("Verbs")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(VerbsSettings)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " VerbsSettings is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = VerbsSettings(valuetmp)

	return
}
