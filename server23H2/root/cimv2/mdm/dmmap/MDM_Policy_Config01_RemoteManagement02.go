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

// MDM_Policy_Config01_RemoteManagement02 struct
type MDM_Policy_Config01_RemoteManagement02 struct {
	*cim.WmiInstance

	//
	AllowBasicAuthentication_Client string

	//
	AllowBasicAuthentication_Service string

	//
	AllowCredSSPAuthenticationClient string

	//
	AllowCredSSPAuthenticationService string

	//
	AllowRemoteServerManagement string

	//
	AllowUnencryptedTraffic_Client string

	//
	AllowUnencryptedTraffic_Service string

	//
	DisallowDigestAuthentication string

	//
	DisallowNegotiateAuthenticationClient string

	//
	DisallowNegotiateAuthenticationService string

	//
	DisallowStoringOfRunAsCredentials string

	//
	InstanceID string

	//
	ParentID string

	//
	SpecifyChannelBindingTokenHardeningLevel string

	//
	TrustedHosts string

	//
	TurnOnCompatibilityHTTPListener string

	//
	TurnOnCompatibilityHTTPSListener string
}

func NewMDM_Policy_Config01_RemoteManagement02Ex1(instance *cim.WmiInstance) (newInstance *MDM_Policy_Config01_RemoteManagement02, err error) {
	tmp, err := instance, nil

	if err != nil {
		return
	}
	newInstance = &MDM_Policy_Config01_RemoteManagement02{
		WmiInstance: tmp,
	}
	return
}

func NewMDM_Policy_Config01_RemoteManagement02Ex6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *MDM_Policy_Config01_RemoteManagement02, err error) {
	tmp, err := instance.GetWmiInstance(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &MDM_Policy_Config01_RemoteManagement02{
		WmiInstance: tmp,
	}
	return
}

// SetAllowBasicAuthentication_Client sets the value of AllowBasicAuthentication_Client for the instance
func (instance *MDM_Policy_Config01_RemoteManagement02) SetPropertyAllowBasicAuthentication_Client(value string) (err error) {
	return instance.SetProperty("AllowBasicAuthentication_Client", (value))
}

// GetAllowBasicAuthentication_Client gets the value of AllowBasicAuthentication_Client for the instance
func (instance *MDM_Policy_Config01_RemoteManagement02) GetPropertyAllowBasicAuthentication_Client() (value string, err error) {
	retValue, err := instance.GetProperty("AllowBasicAuthentication_Client")
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

// SetAllowBasicAuthentication_Service sets the value of AllowBasicAuthentication_Service for the instance
func (instance *MDM_Policy_Config01_RemoteManagement02) SetPropertyAllowBasicAuthentication_Service(value string) (err error) {
	return instance.SetProperty("AllowBasicAuthentication_Service", (value))
}

// GetAllowBasicAuthentication_Service gets the value of AllowBasicAuthentication_Service for the instance
func (instance *MDM_Policy_Config01_RemoteManagement02) GetPropertyAllowBasicAuthentication_Service() (value string, err error) {
	retValue, err := instance.GetProperty("AllowBasicAuthentication_Service")
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

// SetAllowCredSSPAuthenticationClient sets the value of AllowCredSSPAuthenticationClient for the instance
func (instance *MDM_Policy_Config01_RemoteManagement02) SetPropertyAllowCredSSPAuthenticationClient(value string) (err error) {
	return instance.SetProperty("AllowCredSSPAuthenticationClient", (value))
}

// GetAllowCredSSPAuthenticationClient gets the value of AllowCredSSPAuthenticationClient for the instance
func (instance *MDM_Policy_Config01_RemoteManagement02) GetPropertyAllowCredSSPAuthenticationClient() (value string, err error) {
	retValue, err := instance.GetProperty("AllowCredSSPAuthenticationClient")
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

// SetAllowCredSSPAuthenticationService sets the value of AllowCredSSPAuthenticationService for the instance
func (instance *MDM_Policy_Config01_RemoteManagement02) SetPropertyAllowCredSSPAuthenticationService(value string) (err error) {
	return instance.SetProperty("AllowCredSSPAuthenticationService", (value))
}

// GetAllowCredSSPAuthenticationService gets the value of AllowCredSSPAuthenticationService for the instance
func (instance *MDM_Policy_Config01_RemoteManagement02) GetPropertyAllowCredSSPAuthenticationService() (value string, err error) {
	retValue, err := instance.GetProperty("AllowCredSSPAuthenticationService")
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

// SetAllowRemoteServerManagement sets the value of AllowRemoteServerManagement for the instance
func (instance *MDM_Policy_Config01_RemoteManagement02) SetPropertyAllowRemoteServerManagement(value string) (err error) {
	return instance.SetProperty("AllowRemoteServerManagement", (value))
}

// GetAllowRemoteServerManagement gets the value of AllowRemoteServerManagement for the instance
func (instance *MDM_Policy_Config01_RemoteManagement02) GetPropertyAllowRemoteServerManagement() (value string, err error) {
	retValue, err := instance.GetProperty("AllowRemoteServerManagement")
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

// SetAllowUnencryptedTraffic_Client sets the value of AllowUnencryptedTraffic_Client for the instance
func (instance *MDM_Policy_Config01_RemoteManagement02) SetPropertyAllowUnencryptedTraffic_Client(value string) (err error) {
	return instance.SetProperty("AllowUnencryptedTraffic_Client", (value))
}

// GetAllowUnencryptedTraffic_Client gets the value of AllowUnencryptedTraffic_Client for the instance
func (instance *MDM_Policy_Config01_RemoteManagement02) GetPropertyAllowUnencryptedTraffic_Client() (value string, err error) {
	retValue, err := instance.GetProperty("AllowUnencryptedTraffic_Client")
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

// SetAllowUnencryptedTraffic_Service sets the value of AllowUnencryptedTraffic_Service for the instance
func (instance *MDM_Policy_Config01_RemoteManagement02) SetPropertyAllowUnencryptedTraffic_Service(value string) (err error) {
	return instance.SetProperty("AllowUnencryptedTraffic_Service", (value))
}

// GetAllowUnencryptedTraffic_Service gets the value of AllowUnencryptedTraffic_Service for the instance
func (instance *MDM_Policy_Config01_RemoteManagement02) GetPropertyAllowUnencryptedTraffic_Service() (value string, err error) {
	retValue, err := instance.GetProperty("AllowUnencryptedTraffic_Service")
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

// SetDisallowDigestAuthentication sets the value of DisallowDigestAuthentication for the instance
func (instance *MDM_Policy_Config01_RemoteManagement02) SetPropertyDisallowDigestAuthentication(value string) (err error) {
	return instance.SetProperty("DisallowDigestAuthentication", (value))
}

// GetDisallowDigestAuthentication gets the value of DisallowDigestAuthentication for the instance
func (instance *MDM_Policy_Config01_RemoteManagement02) GetPropertyDisallowDigestAuthentication() (value string, err error) {
	retValue, err := instance.GetProperty("DisallowDigestAuthentication")
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

// SetDisallowNegotiateAuthenticationClient sets the value of DisallowNegotiateAuthenticationClient for the instance
func (instance *MDM_Policy_Config01_RemoteManagement02) SetPropertyDisallowNegotiateAuthenticationClient(value string) (err error) {
	return instance.SetProperty("DisallowNegotiateAuthenticationClient", (value))
}

// GetDisallowNegotiateAuthenticationClient gets the value of DisallowNegotiateAuthenticationClient for the instance
func (instance *MDM_Policy_Config01_RemoteManagement02) GetPropertyDisallowNegotiateAuthenticationClient() (value string, err error) {
	retValue, err := instance.GetProperty("DisallowNegotiateAuthenticationClient")
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

// SetDisallowNegotiateAuthenticationService sets the value of DisallowNegotiateAuthenticationService for the instance
func (instance *MDM_Policy_Config01_RemoteManagement02) SetPropertyDisallowNegotiateAuthenticationService(value string) (err error) {
	return instance.SetProperty("DisallowNegotiateAuthenticationService", (value))
}

// GetDisallowNegotiateAuthenticationService gets the value of DisallowNegotiateAuthenticationService for the instance
func (instance *MDM_Policy_Config01_RemoteManagement02) GetPropertyDisallowNegotiateAuthenticationService() (value string, err error) {
	retValue, err := instance.GetProperty("DisallowNegotiateAuthenticationService")
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

// SetDisallowStoringOfRunAsCredentials sets the value of DisallowStoringOfRunAsCredentials for the instance
func (instance *MDM_Policy_Config01_RemoteManagement02) SetPropertyDisallowStoringOfRunAsCredentials(value string) (err error) {
	return instance.SetProperty("DisallowStoringOfRunAsCredentials", (value))
}

// GetDisallowStoringOfRunAsCredentials gets the value of DisallowStoringOfRunAsCredentials for the instance
func (instance *MDM_Policy_Config01_RemoteManagement02) GetPropertyDisallowStoringOfRunAsCredentials() (value string, err error) {
	retValue, err := instance.GetProperty("DisallowStoringOfRunAsCredentials")
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
func (instance *MDM_Policy_Config01_RemoteManagement02) SetPropertyInstanceID(value string) (err error) {
	return instance.SetProperty("InstanceID", (value))
}

// GetInstanceID gets the value of InstanceID for the instance
func (instance *MDM_Policy_Config01_RemoteManagement02) GetPropertyInstanceID() (value string, err error) {
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
func (instance *MDM_Policy_Config01_RemoteManagement02) SetPropertyParentID(value string) (err error) {
	return instance.SetProperty("ParentID", (value))
}

// GetParentID gets the value of ParentID for the instance
func (instance *MDM_Policy_Config01_RemoteManagement02) GetPropertyParentID() (value string, err error) {
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

// SetSpecifyChannelBindingTokenHardeningLevel sets the value of SpecifyChannelBindingTokenHardeningLevel for the instance
func (instance *MDM_Policy_Config01_RemoteManagement02) SetPropertySpecifyChannelBindingTokenHardeningLevel(value string) (err error) {
	return instance.SetProperty("SpecifyChannelBindingTokenHardeningLevel", (value))
}

// GetSpecifyChannelBindingTokenHardeningLevel gets the value of SpecifyChannelBindingTokenHardeningLevel for the instance
func (instance *MDM_Policy_Config01_RemoteManagement02) GetPropertySpecifyChannelBindingTokenHardeningLevel() (value string, err error) {
	retValue, err := instance.GetProperty("SpecifyChannelBindingTokenHardeningLevel")
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

// SetTrustedHosts sets the value of TrustedHosts for the instance
func (instance *MDM_Policy_Config01_RemoteManagement02) SetPropertyTrustedHosts(value string) (err error) {
	return instance.SetProperty("TrustedHosts", (value))
}

// GetTrustedHosts gets the value of TrustedHosts for the instance
func (instance *MDM_Policy_Config01_RemoteManagement02) GetPropertyTrustedHosts() (value string, err error) {
	retValue, err := instance.GetProperty("TrustedHosts")
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

// SetTurnOnCompatibilityHTTPListener sets the value of TurnOnCompatibilityHTTPListener for the instance
func (instance *MDM_Policy_Config01_RemoteManagement02) SetPropertyTurnOnCompatibilityHTTPListener(value string) (err error) {
	return instance.SetProperty("TurnOnCompatibilityHTTPListener", (value))
}

// GetTurnOnCompatibilityHTTPListener gets the value of TurnOnCompatibilityHTTPListener for the instance
func (instance *MDM_Policy_Config01_RemoteManagement02) GetPropertyTurnOnCompatibilityHTTPListener() (value string, err error) {
	retValue, err := instance.GetProperty("TurnOnCompatibilityHTTPListener")
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

// SetTurnOnCompatibilityHTTPSListener sets the value of TurnOnCompatibilityHTTPSListener for the instance
func (instance *MDM_Policy_Config01_RemoteManagement02) SetPropertyTurnOnCompatibilityHTTPSListener(value string) (err error) {
	return instance.SetProperty("TurnOnCompatibilityHTTPSListener", (value))
}

// GetTurnOnCompatibilityHTTPSListener gets the value of TurnOnCompatibilityHTTPSListener for the instance
func (instance *MDM_Policy_Config01_RemoteManagement02) GetPropertyTurnOnCompatibilityHTTPSListener() (value string, err error) {
	retValue, err := instance.GetProperty("TurnOnCompatibilityHTTPSListener")
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
