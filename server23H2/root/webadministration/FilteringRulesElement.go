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

// FilteringRulesElement struct
type FilteringRulesElement struct {
	*CollectionElement

	//
	AppliesTo AppliesToSettings

	//
	DenyStrings DenyStringsSettings

	//
	DenyUnescapedPercent bool

	//
	Name string

	//
	ScanAllRaw bool

	//
	ScanHeaders ScanHeadersSettings

	//
	ScanQueryString bool

	//
	ScanUrl bool
}

func NewFilteringRulesElementEx1(instance *cim.WmiInstance) (newInstance *FilteringRulesElement, err error) {
	tmp, err := NewCollectionElementEx1(instance)

	if err != nil {
		return
	}
	newInstance = &FilteringRulesElement{
		CollectionElement: tmp,
	}
	return
}

func NewFilteringRulesElementEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *FilteringRulesElement, err error) {
	tmp, err := NewCollectionElementEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &FilteringRulesElement{
		CollectionElement: tmp,
	}
	return
}

// SetAppliesTo sets the value of AppliesTo for the instance
func (instance *FilteringRulesElement) SetPropertyAppliesTo(value AppliesToSettings) (err error) {
	return instance.SetProperty("AppliesTo", (value))
}

// GetAppliesTo gets the value of AppliesTo for the instance
func (instance *FilteringRulesElement) GetPropertyAppliesTo() (value AppliesToSettings, err error) {
	retValue, err := instance.GetProperty("AppliesTo")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(AppliesToSettings)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " AppliesToSettings is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = AppliesToSettings(valuetmp)

	return
}

// SetDenyStrings sets the value of DenyStrings for the instance
func (instance *FilteringRulesElement) SetPropertyDenyStrings(value DenyStringsSettings) (err error) {
	return instance.SetProperty("DenyStrings", (value))
}

// GetDenyStrings gets the value of DenyStrings for the instance
func (instance *FilteringRulesElement) GetPropertyDenyStrings() (value DenyStringsSettings, err error) {
	retValue, err := instance.GetProperty("DenyStrings")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(DenyStringsSettings)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " DenyStringsSettings is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = DenyStringsSettings(valuetmp)

	return
}

// SetDenyUnescapedPercent sets the value of DenyUnescapedPercent for the instance
func (instance *FilteringRulesElement) SetPropertyDenyUnescapedPercent(value bool) (err error) {
	return instance.SetProperty("DenyUnescapedPercent", (value))
}

// GetDenyUnescapedPercent gets the value of DenyUnescapedPercent for the instance
func (instance *FilteringRulesElement) GetPropertyDenyUnescapedPercent() (value bool, err error) {
	retValue, err := instance.GetProperty("DenyUnescapedPercent")
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
func (instance *FilteringRulesElement) SetPropertyName(value string) (err error) {
	return instance.SetProperty("Name", (value))
}

// GetName gets the value of Name for the instance
func (instance *FilteringRulesElement) GetPropertyName() (value string, err error) {
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

// SetScanAllRaw sets the value of ScanAllRaw for the instance
func (instance *FilteringRulesElement) SetPropertyScanAllRaw(value bool) (err error) {
	return instance.SetProperty("ScanAllRaw", (value))
}

// GetScanAllRaw gets the value of ScanAllRaw for the instance
func (instance *FilteringRulesElement) GetPropertyScanAllRaw() (value bool, err error) {
	retValue, err := instance.GetProperty("ScanAllRaw")
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

// SetScanHeaders sets the value of ScanHeaders for the instance
func (instance *FilteringRulesElement) SetPropertyScanHeaders(value ScanHeadersSettings) (err error) {
	return instance.SetProperty("ScanHeaders", (value))
}

// GetScanHeaders gets the value of ScanHeaders for the instance
func (instance *FilteringRulesElement) GetPropertyScanHeaders() (value ScanHeadersSettings, err error) {
	retValue, err := instance.GetProperty("ScanHeaders")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(ScanHeadersSettings)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " ScanHeadersSettings is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = ScanHeadersSettings(valuetmp)

	return
}

// SetScanQueryString sets the value of ScanQueryString for the instance
func (instance *FilteringRulesElement) SetPropertyScanQueryString(value bool) (err error) {
	return instance.SetProperty("ScanQueryString", (value))
}

// GetScanQueryString gets the value of ScanQueryString for the instance
func (instance *FilteringRulesElement) GetPropertyScanQueryString() (value bool, err error) {
	retValue, err := instance.GetProperty("ScanQueryString")
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

// SetScanUrl sets the value of ScanUrl for the instance
func (instance *FilteringRulesElement) SetPropertyScanUrl(value bool) (err error) {
	return instance.SetProperty("ScanUrl", (value))
}

// GetScanUrl gets the value of ScanUrl for the instance
func (instance *FilteringRulesElement) GetPropertyScanUrl() (value bool, err error) {
	retValue, err := instance.GetProperty("ScanUrl")
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
