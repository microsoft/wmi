// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 6/13/2024 using wmigen
//      Source root.directory.LDAP
//////////////////////////////////////////////
package ldap

import (
	"github.com/microsoft/wmi/pkg/base/query"
	"github.com/microsoft/wmi/pkg/errors"
	cim "github.com/microsoft/wmi/pkg/wmiinstance"
	"reflect"
)

// ads_msds_authnpolicysilo struct
type ads_msds_authnpolicysilo struct {
	*ds_top

	//
	DS_msDS_AssignedAuthNPolicySiloBL []string

	//
	DS_msDS_AuthNPolicySiloEnforced bool

	//
	DS_msDS_AuthNPolicySiloMembers []string

	//
	DS_msDS_ComputerAuthNPolicy string

	//
	DS_msDS_ServiceAuthNPolicy string

	//
	DS_msDS_UserAuthNPolicy string
}

func Newads_msds_authnpolicysiloEx1(instance *cim.WmiInstance) (newInstance *ads_msds_authnpolicysilo, err error) {
	tmp, err := Newds_topEx1(instance)

	if err != nil {
		return
	}
	newInstance = &ads_msds_authnpolicysilo{
		ds_top: tmp,
	}
	return
}

func Newads_msds_authnpolicysiloEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *ads_msds_authnpolicysilo, err error) {
	tmp, err := Newds_topEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &ads_msds_authnpolicysilo{
		ds_top: tmp,
	}
	return
}

// SetDS_msDS_AssignedAuthNPolicySiloBL sets the value of DS_msDS_AssignedAuthNPolicySiloBL for the instance
func (instance *ads_msds_authnpolicysilo) SetPropertyDS_msDS_AssignedAuthNPolicySiloBL(value []string) (err error) {
	return instance.SetProperty("DS_msDS_AssignedAuthNPolicySiloBL", (value))
}

// GetDS_msDS_AssignedAuthNPolicySiloBL gets the value of DS_msDS_AssignedAuthNPolicySiloBL for the instance
func (instance *ads_msds_authnpolicysilo) GetPropertyDS_msDS_AssignedAuthNPolicySiloBL() (value []string, err error) {
	retValue, err := instance.GetProperty("DS_msDS_AssignedAuthNPolicySiloBL")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	for _, interfaceValue := range retValue.([]interface{}) {
		valuetmp, ok := interfaceValue.(string)
		if !ok {
			err = errors.Wrapf(errors.InvalidType, " string is Invalid. Expected %s", reflect.TypeOf(interfaceValue))
			return
		}
		value = append(value, string(valuetmp))
	}

	return
}

// SetDS_msDS_AuthNPolicySiloEnforced sets the value of DS_msDS_AuthNPolicySiloEnforced for the instance
func (instance *ads_msds_authnpolicysilo) SetPropertyDS_msDS_AuthNPolicySiloEnforced(value bool) (err error) {
	return instance.SetProperty("DS_msDS_AuthNPolicySiloEnforced", (value))
}

// GetDS_msDS_AuthNPolicySiloEnforced gets the value of DS_msDS_AuthNPolicySiloEnforced for the instance
func (instance *ads_msds_authnpolicysilo) GetPropertyDS_msDS_AuthNPolicySiloEnforced() (value bool, err error) {
	retValue, err := instance.GetProperty("DS_msDS_AuthNPolicySiloEnforced")
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

// SetDS_msDS_AuthNPolicySiloMembers sets the value of DS_msDS_AuthNPolicySiloMembers for the instance
func (instance *ads_msds_authnpolicysilo) SetPropertyDS_msDS_AuthNPolicySiloMembers(value []string) (err error) {
	return instance.SetProperty("DS_msDS_AuthNPolicySiloMembers", (value))
}

// GetDS_msDS_AuthNPolicySiloMembers gets the value of DS_msDS_AuthNPolicySiloMembers for the instance
func (instance *ads_msds_authnpolicysilo) GetPropertyDS_msDS_AuthNPolicySiloMembers() (value []string, err error) {
	retValue, err := instance.GetProperty("DS_msDS_AuthNPolicySiloMembers")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	for _, interfaceValue := range retValue.([]interface{}) {
		valuetmp, ok := interfaceValue.(string)
		if !ok {
			err = errors.Wrapf(errors.InvalidType, " string is Invalid. Expected %s", reflect.TypeOf(interfaceValue))
			return
		}
		value = append(value, string(valuetmp))
	}

	return
}

// SetDS_msDS_ComputerAuthNPolicy sets the value of DS_msDS_ComputerAuthNPolicy for the instance
func (instance *ads_msds_authnpolicysilo) SetPropertyDS_msDS_ComputerAuthNPolicy(value string) (err error) {
	return instance.SetProperty("DS_msDS_ComputerAuthNPolicy", (value))
}

// GetDS_msDS_ComputerAuthNPolicy gets the value of DS_msDS_ComputerAuthNPolicy for the instance
func (instance *ads_msds_authnpolicysilo) GetPropertyDS_msDS_ComputerAuthNPolicy() (value string, err error) {
	retValue, err := instance.GetProperty("DS_msDS_ComputerAuthNPolicy")
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

// SetDS_msDS_ServiceAuthNPolicy sets the value of DS_msDS_ServiceAuthNPolicy for the instance
func (instance *ads_msds_authnpolicysilo) SetPropertyDS_msDS_ServiceAuthNPolicy(value string) (err error) {
	return instance.SetProperty("DS_msDS_ServiceAuthNPolicy", (value))
}

// GetDS_msDS_ServiceAuthNPolicy gets the value of DS_msDS_ServiceAuthNPolicy for the instance
func (instance *ads_msds_authnpolicysilo) GetPropertyDS_msDS_ServiceAuthNPolicy() (value string, err error) {
	retValue, err := instance.GetProperty("DS_msDS_ServiceAuthNPolicy")
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

// SetDS_msDS_UserAuthNPolicy sets the value of DS_msDS_UserAuthNPolicy for the instance
func (instance *ads_msds_authnpolicysilo) SetPropertyDS_msDS_UserAuthNPolicy(value string) (err error) {
	return instance.SetProperty("DS_msDS_UserAuthNPolicy", (value))
}

// GetDS_msDS_UserAuthNPolicy gets the value of DS_msDS_UserAuthNPolicy for the instance
func (instance *ads_msds_authnpolicysilo) GetPropertyDS_msDS_UserAuthNPolicy() (value string, err error) {
	retValue, err := instance.GetProperty("DS_msDS_UserAuthNPolicy")
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
