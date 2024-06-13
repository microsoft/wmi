// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 6/13/2024 using wmigen
//      Source root.CIMV2.mdm
//////////////////////////////////////////////
package mdm

import (
	"github.com/microsoft/wmi/pkg/base/instance"
	"github.com/microsoft/wmi/pkg/base/query"
	"github.com/microsoft/wmi/pkg/errors"
	cim "github.com/microsoft/wmi/pkg/wmiinstance"
	"reflect"
)

// MDM_BrowserSettings struct
type MDM_BrowserSettings struct {
	*cim.WmiInstance

	//
	AlwaysSendDoNotTrackHeader bool

	//
	AutofillEnabled bool

	//
	ForceFraudWarning bool

	//
	GoToIntranetForSingleWord bool

	//
	InternetBlockPopups bool

	//
	InternetPluginsEnabled bool

	//
	InternetProtectedModeEnabled bool

	//
	InternetScriptingEnabled bool

	//
	InternetZoneSecurityLevel uint32

	//
	IntranetSecurityZoneEnabled bool

	//
	IntranetZoneSecurityLevel uint32

	//
	key uint32

	//
	RestrictedSitesZoneSecurityLevel uint32

	//
	TrustedSitesZoneSecurityLevel uint32
}

func NewMDM_BrowserSettingsEx1(instance *cim.WmiInstance) (newInstance *MDM_BrowserSettings, err error) {
	tmp, err := instance, nil

	if err != nil {
		return
	}
	newInstance = &MDM_BrowserSettings{
		WmiInstance: tmp,
	}
	return
}

func NewMDM_BrowserSettingsEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *MDM_BrowserSettings, err error) {
	tmp, err := instance.GetWmiInstance(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &MDM_BrowserSettings{
		WmiInstance: tmp,
	}
	return
}

// SetAlwaysSendDoNotTrackHeader sets the value of AlwaysSendDoNotTrackHeader for the instance
func (instance *MDM_BrowserSettings) SetPropertyAlwaysSendDoNotTrackHeader(value bool) (err error) {
	return instance.SetProperty("AlwaysSendDoNotTrackHeader", (value))
}

// GetAlwaysSendDoNotTrackHeader gets the value of AlwaysSendDoNotTrackHeader for the instance
func (instance *MDM_BrowserSettings) GetPropertyAlwaysSendDoNotTrackHeader() (value bool, err error) {
	retValue, err := instance.GetProperty("AlwaysSendDoNotTrackHeader")
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

// SetAutofillEnabled sets the value of AutofillEnabled for the instance
func (instance *MDM_BrowserSettings) SetPropertyAutofillEnabled(value bool) (err error) {
	return instance.SetProperty("AutofillEnabled", (value))
}

// GetAutofillEnabled gets the value of AutofillEnabled for the instance
func (instance *MDM_BrowserSettings) GetPropertyAutofillEnabled() (value bool, err error) {
	retValue, err := instance.GetProperty("AutofillEnabled")
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

// SetForceFraudWarning sets the value of ForceFraudWarning for the instance
func (instance *MDM_BrowserSettings) SetPropertyForceFraudWarning(value bool) (err error) {
	return instance.SetProperty("ForceFraudWarning", (value))
}

// GetForceFraudWarning gets the value of ForceFraudWarning for the instance
func (instance *MDM_BrowserSettings) GetPropertyForceFraudWarning() (value bool, err error) {
	retValue, err := instance.GetProperty("ForceFraudWarning")
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

// SetGoToIntranetForSingleWord sets the value of GoToIntranetForSingleWord for the instance
func (instance *MDM_BrowserSettings) SetPropertyGoToIntranetForSingleWord(value bool) (err error) {
	return instance.SetProperty("GoToIntranetForSingleWord", (value))
}

// GetGoToIntranetForSingleWord gets the value of GoToIntranetForSingleWord for the instance
func (instance *MDM_BrowserSettings) GetPropertyGoToIntranetForSingleWord() (value bool, err error) {
	retValue, err := instance.GetProperty("GoToIntranetForSingleWord")
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

// SetInternetBlockPopups sets the value of InternetBlockPopups for the instance
func (instance *MDM_BrowserSettings) SetPropertyInternetBlockPopups(value bool) (err error) {
	return instance.SetProperty("InternetBlockPopups", (value))
}

// GetInternetBlockPopups gets the value of InternetBlockPopups for the instance
func (instance *MDM_BrowserSettings) GetPropertyInternetBlockPopups() (value bool, err error) {
	retValue, err := instance.GetProperty("InternetBlockPopups")
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

// SetInternetPluginsEnabled sets the value of InternetPluginsEnabled for the instance
func (instance *MDM_BrowserSettings) SetPropertyInternetPluginsEnabled(value bool) (err error) {
	return instance.SetProperty("InternetPluginsEnabled", (value))
}

// GetInternetPluginsEnabled gets the value of InternetPluginsEnabled for the instance
func (instance *MDM_BrowserSettings) GetPropertyInternetPluginsEnabled() (value bool, err error) {
	retValue, err := instance.GetProperty("InternetPluginsEnabled")
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

// SetInternetProtectedModeEnabled sets the value of InternetProtectedModeEnabled for the instance
func (instance *MDM_BrowserSettings) SetPropertyInternetProtectedModeEnabled(value bool) (err error) {
	return instance.SetProperty("InternetProtectedModeEnabled", (value))
}

// GetInternetProtectedModeEnabled gets the value of InternetProtectedModeEnabled for the instance
func (instance *MDM_BrowserSettings) GetPropertyInternetProtectedModeEnabled() (value bool, err error) {
	retValue, err := instance.GetProperty("InternetProtectedModeEnabled")
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

// SetInternetScriptingEnabled sets the value of InternetScriptingEnabled for the instance
func (instance *MDM_BrowserSettings) SetPropertyInternetScriptingEnabled(value bool) (err error) {
	return instance.SetProperty("InternetScriptingEnabled", (value))
}

// GetInternetScriptingEnabled gets the value of InternetScriptingEnabled for the instance
func (instance *MDM_BrowserSettings) GetPropertyInternetScriptingEnabled() (value bool, err error) {
	retValue, err := instance.GetProperty("InternetScriptingEnabled")
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

// SetInternetZoneSecurityLevel sets the value of InternetZoneSecurityLevel for the instance
func (instance *MDM_BrowserSettings) SetPropertyInternetZoneSecurityLevel(value uint32) (err error) {
	return instance.SetProperty("InternetZoneSecurityLevel", (value))
}

// GetInternetZoneSecurityLevel gets the value of InternetZoneSecurityLevel for the instance
func (instance *MDM_BrowserSettings) GetPropertyInternetZoneSecurityLevel() (value uint32, err error) {
	retValue, err := instance.GetProperty("InternetZoneSecurityLevel")
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

// SetIntranetSecurityZoneEnabled sets the value of IntranetSecurityZoneEnabled for the instance
func (instance *MDM_BrowserSettings) SetPropertyIntranetSecurityZoneEnabled(value bool) (err error) {
	return instance.SetProperty("IntranetSecurityZoneEnabled", (value))
}

// GetIntranetSecurityZoneEnabled gets the value of IntranetSecurityZoneEnabled for the instance
func (instance *MDM_BrowserSettings) GetPropertyIntranetSecurityZoneEnabled() (value bool, err error) {
	retValue, err := instance.GetProperty("IntranetSecurityZoneEnabled")
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

// SetIntranetZoneSecurityLevel sets the value of IntranetZoneSecurityLevel for the instance
func (instance *MDM_BrowserSettings) SetPropertyIntranetZoneSecurityLevel(value uint32) (err error) {
	return instance.SetProperty("IntranetZoneSecurityLevel", (value))
}

// GetIntranetZoneSecurityLevel gets the value of IntranetZoneSecurityLevel for the instance
func (instance *MDM_BrowserSettings) GetPropertyIntranetZoneSecurityLevel() (value uint32, err error) {
	retValue, err := instance.GetProperty("IntranetZoneSecurityLevel")
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

// Setkey sets the value of key for the instance
func (instance *MDM_BrowserSettings) SetPropertykey(value uint32) (err error) {
	return instance.SetProperty("key", (value))
}

// Getkey gets the value of key for the instance
func (instance *MDM_BrowserSettings) GetPropertykey() (value uint32, err error) {
	retValue, err := instance.GetProperty("key")
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

// SetRestrictedSitesZoneSecurityLevel sets the value of RestrictedSitesZoneSecurityLevel for the instance
func (instance *MDM_BrowserSettings) SetPropertyRestrictedSitesZoneSecurityLevel(value uint32) (err error) {
	return instance.SetProperty("RestrictedSitesZoneSecurityLevel", (value))
}

// GetRestrictedSitesZoneSecurityLevel gets the value of RestrictedSitesZoneSecurityLevel for the instance
func (instance *MDM_BrowserSettings) GetPropertyRestrictedSitesZoneSecurityLevel() (value uint32, err error) {
	retValue, err := instance.GetProperty("RestrictedSitesZoneSecurityLevel")
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

// SetTrustedSitesZoneSecurityLevel sets the value of TrustedSitesZoneSecurityLevel for the instance
func (instance *MDM_BrowserSettings) SetPropertyTrustedSitesZoneSecurityLevel(value uint32) (err error) {
	return instance.SetProperty("TrustedSitesZoneSecurityLevel", (value))
}

// GetTrustedSitesZoneSecurityLevel gets the value of TrustedSitesZoneSecurityLevel for the instance
func (instance *MDM_BrowserSettings) GetPropertyTrustedSitesZoneSecurityLevel() (value uint32, err error) {
	retValue, err := instance.GetProperty("TrustedSitesZoneSecurityLevel")
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
