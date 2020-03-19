// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 3/19/2020 using wmigen
//      Source root.RSOP.Computer
//////////////////////////////////////////////
package computer

import (
	"github.com/microsoft/wmi/pkg/base/instance"
	"github.com/microsoft/wmi/pkg/base/query"
	cim "github.com/microsoft/wmi/pkg/wmiinstance"
)

// RSOP_IEPrivacySettings struct
type RSOP_IEPrivacySettings struct {
	*cim.WmiInstance

	//
	firstPartyPrivacyType uint32

	//
	firstPartyPrivacyTypeText string

	//
	rsopID string

	//
	rsopPrecedence int32

	//
	thirdPartyPrivacyType uint32

	//
	thirdPartyPrivacyTypeText string

	//
	useAdvancedSettings bool
}

func NewRSOP_IEPrivacySettingsEx1(instance *cim.WmiInstance) (newInstance *RSOP_IEPrivacySettings, err error) {
	tmp, err := instance, nil

	if err != nil {
		return
	}
	newInstance = &RSOP_IEPrivacySettings{
		WmiInstance: tmp,
	}
	return
}

func NewRSOP_IEPrivacySettingsEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *RSOP_IEPrivacySettings, err error) {
	tmp, err := instance.GetWmiInstance(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &RSOP_IEPrivacySettings{
		WmiInstance: tmp,
	}
	return
}

// SetfirstPartyPrivacyType sets the value of firstPartyPrivacyType for the instance
func (instance *RSOP_IEPrivacySettings) SetPropertyfirstPartyPrivacyType(value uint32) (err error) {
	return instance.SetProperty("firstPartyPrivacyType", value)
}

// GetfirstPartyPrivacyType gets the value of firstPartyPrivacyType for the instance
func (instance *RSOP_IEPrivacySettings) GetPropertyfirstPartyPrivacyType() (value uint32, err error) {
	retValue, err := instance.GetProperty("firstPartyPrivacyType")
	if err != nil {
		return
	}
	value, ok := retValue.(uint32)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetfirstPartyPrivacyTypeText sets the value of firstPartyPrivacyTypeText for the instance
func (instance *RSOP_IEPrivacySettings) SetPropertyfirstPartyPrivacyTypeText(value string) (err error) {
	return instance.SetProperty("firstPartyPrivacyTypeText", value)
}

// GetfirstPartyPrivacyTypeText gets the value of firstPartyPrivacyTypeText for the instance
func (instance *RSOP_IEPrivacySettings) GetPropertyfirstPartyPrivacyTypeText() (value string, err error) {
	retValue, err := instance.GetProperty("firstPartyPrivacyTypeText")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetrsopID sets the value of rsopID for the instance
func (instance *RSOP_IEPrivacySettings) SetPropertyrsopID(value string) (err error) {
	return instance.SetProperty("rsopID", value)
}

// GetrsopID gets the value of rsopID for the instance
func (instance *RSOP_IEPrivacySettings) GetPropertyrsopID() (value string, err error) {
	retValue, err := instance.GetProperty("rsopID")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetrsopPrecedence sets the value of rsopPrecedence for the instance
func (instance *RSOP_IEPrivacySettings) SetPropertyrsopPrecedence(value int32) (err error) {
	return instance.SetProperty("rsopPrecedence", value)
}

// GetrsopPrecedence gets the value of rsopPrecedence for the instance
func (instance *RSOP_IEPrivacySettings) GetPropertyrsopPrecedence() (value int32, err error) {
	retValue, err := instance.GetProperty("rsopPrecedence")
	if err != nil {
		return
	}
	value, ok := retValue.(int32)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetthirdPartyPrivacyType sets the value of thirdPartyPrivacyType for the instance
func (instance *RSOP_IEPrivacySettings) SetPropertythirdPartyPrivacyType(value uint32) (err error) {
	return instance.SetProperty("thirdPartyPrivacyType", value)
}

// GetthirdPartyPrivacyType gets the value of thirdPartyPrivacyType for the instance
func (instance *RSOP_IEPrivacySettings) GetPropertythirdPartyPrivacyType() (value uint32, err error) {
	retValue, err := instance.GetProperty("thirdPartyPrivacyType")
	if err != nil {
		return
	}
	value, ok := retValue.(uint32)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetthirdPartyPrivacyTypeText sets the value of thirdPartyPrivacyTypeText for the instance
func (instance *RSOP_IEPrivacySettings) SetPropertythirdPartyPrivacyTypeText(value string) (err error) {
	return instance.SetProperty("thirdPartyPrivacyTypeText", value)
}

// GetthirdPartyPrivacyTypeText gets the value of thirdPartyPrivacyTypeText for the instance
func (instance *RSOP_IEPrivacySettings) GetPropertythirdPartyPrivacyTypeText() (value string, err error) {
	retValue, err := instance.GetProperty("thirdPartyPrivacyTypeText")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetuseAdvancedSettings sets the value of useAdvancedSettings for the instance
func (instance *RSOP_IEPrivacySettings) SetPropertyuseAdvancedSettings(value bool) (err error) {
	return instance.SetProperty("useAdvancedSettings", value)
}

// GetuseAdvancedSettings gets the value of useAdvancedSettings for the instance
func (instance *RSOP_IEPrivacySettings) GetPropertyuseAdvancedSettings() (value bool, err error) {
	retValue, err := instance.GetProperty("useAdvancedSettings")
	if err != nil {
		return
	}
	value, ok := retValue.(bool)
	if !ok {
		// TODO: Set an error
	}
	return
}
