// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 6/14/2024 using wmigen
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

// MDM_Policy_Result01_MSSLegacy02 struct
type MDM_Policy_Result01_MSSLegacy02 struct {
	*cim.WmiInstance

	//
	AllowICMPRedirectsToOverrideOSPFGeneratedRoutes string

	//
	AllowTheComputerToIgnoreNetBIOSNameReleaseRequestsExceptFromWINSServers string

	//
	InstanceID string

	//
	IPSourceRoutingProtectionLevel string

	//
	IPv6SourceRoutingProtectionLevel string

	//
	ParentID string
}

func NewMDM_Policy_Result01_MSSLegacy02Ex1(instance *cim.WmiInstance) (newInstance *MDM_Policy_Result01_MSSLegacy02, err error) {
	tmp, err := instance, nil

	if err != nil {
		return
	}
	newInstance = &MDM_Policy_Result01_MSSLegacy02{
		WmiInstance: tmp,
	}
	return
}

func NewMDM_Policy_Result01_MSSLegacy02Ex6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *MDM_Policy_Result01_MSSLegacy02, err error) {
	tmp, err := instance.GetWmiInstance(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &MDM_Policy_Result01_MSSLegacy02{
		WmiInstance: tmp,
	}
	return
}

// SetAllowICMPRedirectsToOverrideOSPFGeneratedRoutes sets the value of AllowICMPRedirectsToOverrideOSPFGeneratedRoutes for the instance
func (instance *MDM_Policy_Result01_MSSLegacy02) SetPropertyAllowICMPRedirectsToOverrideOSPFGeneratedRoutes(value string) (err error) {
	return instance.SetProperty("AllowICMPRedirectsToOverrideOSPFGeneratedRoutes", (value))
}

// GetAllowICMPRedirectsToOverrideOSPFGeneratedRoutes gets the value of AllowICMPRedirectsToOverrideOSPFGeneratedRoutes for the instance
func (instance *MDM_Policy_Result01_MSSLegacy02) GetPropertyAllowICMPRedirectsToOverrideOSPFGeneratedRoutes() (value string, err error) {
	retValue, err := instance.GetProperty("AllowICMPRedirectsToOverrideOSPFGeneratedRoutes")
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

// SetAllowTheComputerToIgnoreNetBIOSNameReleaseRequestsExceptFromWINSServers sets the value of AllowTheComputerToIgnoreNetBIOSNameReleaseRequestsExceptFromWINSServers for the instance
func (instance *MDM_Policy_Result01_MSSLegacy02) SetPropertyAllowTheComputerToIgnoreNetBIOSNameReleaseRequestsExceptFromWINSServers(value string) (err error) {
	return instance.SetProperty("AllowTheComputerToIgnoreNetBIOSNameReleaseRequestsExceptFromWINSServers", (value))
}

// GetAllowTheComputerToIgnoreNetBIOSNameReleaseRequestsExceptFromWINSServers gets the value of AllowTheComputerToIgnoreNetBIOSNameReleaseRequestsExceptFromWINSServers for the instance
func (instance *MDM_Policy_Result01_MSSLegacy02) GetPropertyAllowTheComputerToIgnoreNetBIOSNameReleaseRequestsExceptFromWINSServers() (value string, err error) {
	retValue, err := instance.GetProperty("AllowTheComputerToIgnoreNetBIOSNameReleaseRequestsExceptFromWINSServers")
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
func (instance *MDM_Policy_Result01_MSSLegacy02) SetPropertyInstanceID(value string) (err error) {
	return instance.SetProperty("InstanceID", (value))
}

// GetInstanceID gets the value of InstanceID for the instance
func (instance *MDM_Policy_Result01_MSSLegacy02) GetPropertyInstanceID() (value string, err error) {
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

// SetIPSourceRoutingProtectionLevel sets the value of IPSourceRoutingProtectionLevel for the instance
func (instance *MDM_Policy_Result01_MSSLegacy02) SetPropertyIPSourceRoutingProtectionLevel(value string) (err error) {
	return instance.SetProperty("IPSourceRoutingProtectionLevel", (value))
}

// GetIPSourceRoutingProtectionLevel gets the value of IPSourceRoutingProtectionLevel for the instance
func (instance *MDM_Policy_Result01_MSSLegacy02) GetPropertyIPSourceRoutingProtectionLevel() (value string, err error) {
	retValue, err := instance.GetProperty("IPSourceRoutingProtectionLevel")
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

// SetIPv6SourceRoutingProtectionLevel sets the value of IPv6SourceRoutingProtectionLevel for the instance
func (instance *MDM_Policy_Result01_MSSLegacy02) SetPropertyIPv6SourceRoutingProtectionLevel(value string) (err error) {
	return instance.SetProperty("IPv6SourceRoutingProtectionLevel", (value))
}

// GetIPv6SourceRoutingProtectionLevel gets the value of IPv6SourceRoutingProtectionLevel for the instance
func (instance *MDM_Policy_Result01_MSSLegacy02) GetPropertyIPv6SourceRoutingProtectionLevel() (value string, err error) {
	retValue, err := instance.GetProperty("IPv6SourceRoutingProtectionLevel")
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
func (instance *MDM_Policy_Result01_MSSLegacy02) SetPropertyParentID(value string) (err error) {
	return instance.SetProperty("ParentID", (value))
}

// GetParentID gets the value of ParentID for the instance
func (instance *MDM_Policy_Result01_MSSLegacy02) GetPropertyParentID() (value string, err error) {
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
