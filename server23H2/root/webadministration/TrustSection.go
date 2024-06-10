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

// TrustSection struct
type TrustSection struct {
	*ConfigurationSection

	//
	LegacyCasModel bool

	//
	Level string

	//
	OriginUrl string

	//
	ProcessRequestInApplicationTrust bool
}

func NewTrustSectionEx1(instance *cim.WmiInstance) (newInstance *TrustSection, err error) {
	tmp, err := NewConfigurationSectionEx1(instance)

	if err != nil {
		return
	}
	newInstance = &TrustSection{
		ConfigurationSection: tmp,
	}
	return
}

func NewTrustSectionEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *TrustSection, err error) {
	tmp, err := NewConfigurationSectionEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &TrustSection{
		ConfigurationSection: tmp,
	}
	return
}

// SetLegacyCasModel sets the value of LegacyCasModel for the instance
func (instance *TrustSection) SetPropertyLegacyCasModel(value bool) (err error) {
	return instance.SetProperty("LegacyCasModel", (value))
}

// GetLegacyCasModel gets the value of LegacyCasModel for the instance
func (instance *TrustSection) GetPropertyLegacyCasModel() (value bool, err error) {
	retValue, err := instance.GetProperty("LegacyCasModel")
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

// SetLevel sets the value of Level for the instance
func (instance *TrustSection) SetPropertyLevel(value string) (err error) {
	return instance.SetProperty("Level", (value))
}

// GetLevel gets the value of Level for the instance
func (instance *TrustSection) GetPropertyLevel() (value string, err error) {
	retValue, err := instance.GetProperty("Level")
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

// SetOriginUrl sets the value of OriginUrl for the instance
func (instance *TrustSection) SetPropertyOriginUrl(value string) (err error) {
	return instance.SetProperty("OriginUrl", (value))
}

// GetOriginUrl gets the value of OriginUrl for the instance
func (instance *TrustSection) GetPropertyOriginUrl() (value string, err error) {
	retValue, err := instance.GetProperty("OriginUrl")
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

// SetProcessRequestInApplicationTrust sets the value of ProcessRequestInApplicationTrust for the instance
func (instance *TrustSection) SetPropertyProcessRequestInApplicationTrust(value bool) (err error) {
	return instance.SetProperty("ProcessRequestInApplicationTrust", (value))
}

// GetProcessRequestInApplicationTrust gets the value of ProcessRequestInApplicationTrust for the instance
func (instance *TrustSection) GetPropertyProcessRequestInApplicationTrust() (value bool, err error) {
	retValue, err := instance.GetProperty("ProcessRequestInApplicationTrust")
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
