// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 6/12/2024 using wmigen
//      Source root.CIMV2.mdm.dmmap
//////////////////////////////////////////////
package dmmap

import (
	"github.com/microsoft/wmi/pkg/base/instance"
	"github.com/microsoft/wmi/pkg/base/query"
	"github.com/microsoft/wmi/pkg/errors"
	cim "github.com/microsoft/wmi/pkg/wmiinstance"
	"reflect"
)

// MDM_Policy_Result01_MSSecurityGuide02 struct
type MDM_Policy_Result01_MSSecurityGuide02 struct {
	*cim.WmiInstance

	//
	ApplyUACRestrictionsToLocalAccountsOnNetworkLogon string

	//
	ConfigureSMBV1ClientDriver string

	//
	ConfigureSMBV1Server string

	//
	EnableStructuredExceptionHandlingOverwriteProtection string

	//
	InstanceID string

	//
	ParentID string

	//
	TurnOnWindowsDefenderProtectionAgainstPotentiallyUnwantedApplications string

	//
	WDigestAuthentication string
}

func NewMDM_Policy_Result01_MSSecurityGuide02Ex1(instance *cim.WmiInstance) (newInstance *MDM_Policy_Result01_MSSecurityGuide02, err error) {
	tmp, err := instance, nil

	if err != nil {
		return
	}
	newInstance = &MDM_Policy_Result01_MSSecurityGuide02{
		WmiInstance: tmp,
	}
	return
}

func NewMDM_Policy_Result01_MSSecurityGuide02Ex6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *MDM_Policy_Result01_MSSecurityGuide02, err error) {
	tmp, err := instance.GetWmiInstance(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &MDM_Policy_Result01_MSSecurityGuide02{
		WmiInstance: tmp,
	}
	return
}

// SetApplyUACRestrictionsToLocalAccountsOnNetworkLogon sets the value of ApplyUACRestrictionsToLocalAccountsOnNetworkLogon for the instance
func (instance *MDM_Policy_Result01_MSSecurityGuide02) SetPropertyApplyUACRestrictionsToLocalAccountsOnNetworkLogon(value string) (err error) {
	return instance.SetProperty("ApplyUACRestrictionsToLocalAccountsOnNetworkLogon", (value))
}

// GetApplyUACRestrictionsToLocalAccountsOnNetworkLogon gets the value of ApplyUACRestrictionsToLocalAccountsOnNetworkLogon for the instance
func (instance *MDM_Policy_Result01_MSSecurityGuide02) GetPropertyApplyUACRestrictionsToLocalAccountsOnNetworkLogon() (value string, err error) {
	retValue, err := instance.GetProperty("ApplyUACRestrictionsToLocalAccountsOnNetworkLogon")
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

// SetConfigureSMBV1ClientDriver sets the value of ConfigureSMBV1ClientDriver for the instance
func (instance *MDM_Policy_Result01_MSSecurityGuide02) SetPropertyConfigureSMBV1ClientDriver(value string) (err error) {
	return instance.SetProperty("ConfigureSMBV1ClientDriver", (value))
}

// GetConfigureSMBV1ClientDriver gets the value of ConfigureSMBV1ClientDriver for the instance
func (instance *MDM_Policy_Result01_MSSecurityGuide02) GetPropertyConfigureSMBV1ClientDriver() (value string, err error) {
	retValue, err := instance.GetProperty("ConfigureSMBV1ClientDriver")
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

// SetConfigureSMBV1Server sets the value of ConfigureSMBV1Server for the instance
func (instance *MDM_Policy_Result01_MSSecurityGuide02) SetPropertyConfigureSMBV1Server(value string) (err error) {
	return instance.SetProperty("ConfigureSMBV1Server", (value))
}

// GetConfigureSMBV1Server gets the value of ConfigureSMBV1Server for the instance
func (instance *MDM_Policy_Result01_MSSecurityGuide02) GetPropertyConfigureSMBV1Server() (value string, err error) {
	retValue, err := instance.GetProperty("ConfigureSMBV1Server")
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

// SetEnableStructuredExceptionHandlingOverwriteProtection sets the value of EnableStructuredExceptionHandlingOverwriteProtection for the instance
func (instance *MDM_Policy_Result01_MSSecurityGuide02) SetPropertyEnableStructuredExceptionHandlingOverwriteProtection(value string) (err error) {
	return instance.SetProperty("EnableStructuredExceptionHandlingOverwriteProtection", (value))
}

// GetEnableStructuredExceptionHandlingOverwriteProtection gets the value of EnableStructuredExceptionHandlingOverwriteProtection for the instance
func (instance *MDM_Policy_Result01_MSSecurityGuide02) GetPropertyEnableStructuredExceptionHandlingOverwriteProtection() (value string, err error) {
	retValue, err := instance.GetProperty("EnableStructuredExceptionHandlingOverwriteProtection")
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

// SetInstanceID sets the value of InstanceID for the instance
func (instance *MDM_Policy_Result01_MSSecurityGuide02) SetPropertyInstanceID(value string) (err error) {
	return instance.SetProperty("InstanceID", (value))
}

// GetInstanceID gets the value of InstanceID for the instance
func (instance *MDM_Policy_Result01_MSSecurityGuide02) GetPropertyInstanceID() (value string, err error) {
	retValue, err := instance.GetProperty("InstanceID")
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

// SetParentID sets the value of ParentID for the instance
func (instance *MDM_Policy_Result01_MSSecurityGuide02) SetPropertyParentID(value string) (err error) {
	return instance.SetProperty("ParentID", (value))
}

// GetParentID gets the value of ParentID for the instance
func (instance *MDM_Policy_Result01_MSSecurityGuide02) GetPropertyParentID() (value string, err error) {
	retValue, err := instance.GetProperty("ParentID")
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

// SetTurnOnWindowsDefenderProtectionAgainstPotentiallyUnwantedApplications sets the value of TurnOnWindowsDefenderProtectionAgainstPotentiallyUnwantedApplications for the instance
func (instance *MDM_Policy_Result01_MSSecurityGuide02) SetPropertyTurnOnWindowsDefenderProtectionAgainstPotentiallyUnwantedApplications(value string) (err error) {
	return instance.SetProperty("TurnOnWindowsDefenderProtectionAgainstPotentiallyUnwantedApplications", (value))
}

// GetTurnOnWindowsDefenderProtectionAgainstPotentiallyUnwantedApplications gets the value of TurnOnWindowsDefenderProtectionAgainstPotentiallyUnwantedApplications for the instance
func (instance *MDM_Policy_Result01_MSSecurityGuide02) GetPropertyTurnOnWindowsDefenderProtectionAgainstPotentiallyUnwantedApplications() (value string, err error) {
	retValue, err := instance.GetProperty("TurnOnWindowsDefenderProtectionAgainstPotentiallyUnwantedApplications")
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

// SetWDigestAuthentication sets the value of WDigestAuthentication for the instance
func (instance *MDM_Policy_Result01_MSSecurityGuide02) SetPropertyWDigestAuthentication(value string) (err error) {
	return instance.SetProperty("WDigestAuthentication", (value))
}

// GetWDigestAuthentication gets the value of WDigestAuthentication for the instance
func (instance *MDM_Policy_Result01_MSSecurityGuide02) GetPropertyWDigestAuthentication() (value string, err error) {
	retValue, err := instance.GetProperty("WDigestAuthentication")
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
