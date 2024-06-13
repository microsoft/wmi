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

// ads_msds_authnpolicy struct
type ads_msds_authnpolicy struct {
	*ds_top

	//
	DS_msDS_AssignedAuthNPolicyBL []string

	//
	DS_msDS_AuthNPolicyEnforced bool

	//
	DS_msDS_ComputerAllowedToAuthenticateTo Uint8Array

	//
	DS_msDS_ComputerAuthNPolicyBL []string

	//
	DS_msDS_ComputerTGTLifetime int64

	//
	DS_msDS_ServiceAllowedNTLMNetworkAuthentication bool

	//
	DS_msDS_ServiceAllowedToAuthenticateFrom Uint8Array

	//
	DS_msDS_ServiceAllowedToAuthenticateTo Uint8Array

	//
	DS_msDS_ServiceAuthNPolicyBL []string

	//
	DS_msDS_ServiceTGTLifetime int64

	//
	DS_msDS_StrongNTLMPolicy int32

	//
	DS_msDS_UserAllowedNTLMNetworkAuthentication bool

	//
	DS_msDS_UserAllowedToAuthenticateFrom Uint8Array

	//
	DS_msDS_UserAllowedToAuthenticateTo Uint8Array

	//
	DS_msDS_UserAuthNPolicyBL []string

	//
	DS_msDS_UserTGTLifetime int64
}

func Newads_msds_authnpolicyEx1(instance *cim.WmiInstance) (newInstance *ads_msds_authnpolicy, err error) {
	tmp, err := Newds_topEx1(instance)

	if err != nil {
		return
	}
	newInstance = &ads_msds_authnpolicy{
		ds_top: tmp,
	}
	return
}

func Newads_msds_authnpolicyEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *ads_msds_authnpolicy, err error) {
	tmp, err := Newds_topEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &ads_msds_authnpolicy{
		ds_top: tmp,
	}
	return
}

// SetDS_msDS_AssignedAuthNPolicyBL sets the value of DS_msDS_AssignedAuthNPolicyBL for the instance
func (instance *ads_msds_authnpolicy) SetPropertyDS_msDS_AssignedAuthNPolicyBL(value []string) (err error) {
	return instance.SetProperty("DS_msDS_AssignedAuthNPolicyBL", (value))
}

// GetDS_msDS_AssignedAuthNPolicyBL gets the value of DS_msDS_AssignedAuthNPolicyBL for the instance
func (instance *ads_msds_authnpolicy) GetPropertyDS_msDS_AssignedAuthNPolicyBL() (value []string, err error) {
	retValue, err := instance.GetProperty("DS_msDS_AssignedAuthNPolicyBL")
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

// SetDS_msDS_AuthNPolicyEnforced sets the value of DS_msDS_AuthNPolicyEnforced for the instance
func (instance *ads_msds_authnpolicy) SetPropertyDS_msDS_AuthNPolicyEnforced(value bool) (err error) {
	return instance.SetProperty("DS_msDS_AuthNPolicyEnforced", (value))
}

// GetDS_msDS_AuthNPolicyEnforced gets the value of DS_msDS_AuthNPolicyEnforced for the instance
func (instance *ads_msds_authnpolicy) GetPropertyDS_msDS_AuthNPolicyEnforced() (value bool, err error) {
	retValue, err := instance.GetProperty("DS_msDS_AuthNPolicyEnforced")
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

// SetDS_msDS_ComputerAllowedToAuthenticateTo sets the value of DS_msDS_ComputerAllowedToAuthenticateTo for the instance
func (instance *ads_msds_authnpolicy) SetPropertyDS_msDS_ComputerAllowedToAuthenticateTo(value Uint8Array) (err error) {
	return instance.SetProperty("DS_msDS_ComputerAllowedToAuthenticateTo", (value))
}

// GetDS_msDS_ComputerAllowedToAuthenticateTo gets the value of DS_msDS_ComputerAllowedToAuthenticateTo for the instance
func (instance *ads_msds_authnpolicy) GetPropertyDS_msDS_ComputerAllowedToAuthenticateTo() (value Uint8Array, err error) {
	retValue, err := instance.GetProperty("DS_msDS_ComputerAllowedToAuthenticateTo")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(Uint8Array)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " Uint8Array is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = Uint8Array(valuetmp)

	return
}

// SetDS_msDS_ComputerAuthNPolicyBL sets the value of DS_msDS_ComputerAuthNPolicyBL for the instance
func (instance *ads_msds_authnpolicy) SetPropertyDS_msDS_ComputerAuthNPolicyBL(value []string) (err error) {
	return instance.SetProperty("DS_msDS_ComputerAuthNPolicyBL", (value))
}

// GetDS_msDS_ComputerAuthNPolicyBL gets the value of DS_msDS_ComputerAuthNPolicyBL for the instance
func (instance *ads_msds_authnpolicy) GetPropertyDS_msDS_ComputerAuthNPolicyBL() (value []string, err error) {
	retValue, err := instance.GetProperty("DS_msDS_ComputerAuthNPolicyBL")
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

// SetDS_msDS_ComputerTGTLifetime sets the value of DS_msDS_ComputerTGTLifetime for the instance
func (instance *ads_msds_authnpolicy) SetPropertyDS_msDS_ComputerTGTLifetime(value int64) (err error) {
	return instance.SetProperty("DS_msDS_ComputerTGTLifetime", (value))
}

// GetDS_msDS_ComputerTGTLifetime gets the value of DS_msDS_ComputerTGTLifetime for the instance
func (instance *ads_msds_authnpolicy) GetPropertyDS_msDS_ComputerTGTLifetime() (value int64, err error) {
	retValue, err := instance.GetProperty("DS_msDS_ComputerTGTLifetime")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(int64)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " int64 is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = int64(valuetmp)

	return
}

// SetDS_msDS_ServiceAllowedNTLMNetworkAuthentication sets the value of DS_msDS_ServiceAllowedNTLMNetworkAuthentication for the instance
func (instance *ads_msds_authnpolicy) SetPropertyDS_msDS_ServiceAllowedNTLMNetworkAuthentication(value bool) (err error) {
	return instance.SetProperty("DS_msDS_ServiceAllowedNTLMNetworkAuthentication", (value))
}

// GetDS_msDS_ServiceAllowedNTLMNetworkAuthentication gets the value of DS_msDS_ServiceAllowedNTLMNetworkAuthentication for the instance
func (instance *ads_msds_authnpolicy) GetPropertyDS_msDS_ServiceAllowedNTLMNetworkAuthentication() (value bool, err error) {
	retValue, err := instance.GetProperty("DS_msDS_ServiceAllowedNTLMNetworkAuthentication")
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

// SetDS_msDS_ServiceAllowedToAuthenticateFrom sets the value of DS_msDS_ServiceAllowedToAuthenticateFrom for the instance
func (instance *ads_msds_authnpolicy) SetPropertyDS_msDS_ServiceAllowedToAuthenticateFrom(value Uint8Array) (err error) {
	return instance.SetProperty("DS_msDS_ServiceAllowedToAuthenticateFrom", (value))
}

// GetDS_msDS_ServiceAllowedToAuthenticateFrom gets the value of DS_msDS_ServiceAllowedToAuthenticateFrom for the instance
func (instance *ads_msds_authnpolicy) GetPropertyDS_msDS_ServiceAllowedToAuthenticateFrom() (value Uint8Array, err error) {
	retValue, err := instance.GetProperty("DS_msDS_ServiceAllowedToAuthenticateFrom")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(Uint8Array)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " Uint8Array is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = Uint8Array(valuetmp)

	return
}

// SetDS_msDS_ServiceAllowedToAuthenticateTo sets the value of DS_msDS_ServiceAllowedToAuthenticateTo for the instance
func (instance *ads_msds_authnpolicy) SetPropertyDS_msDS_ServiceAllowedToAuthenticateTo(value Uint8Array) (err error) {
	return instance.SetProperty("DS_msDS_ServiceAllowedToAuthenticateTo", (value))
}

// GetDS_msDS_ServiceAllowedToAuthenticateTo gets the value of DS_msDS_ServiceAllowedToAuthenticateTo for the instance
func (instance *ads_msds_authnpolicy) GetPropertyDS_msDS_ServiceAllowedToAuthenticateTo() (value Uint8Array, err error) {
	retValue, err := instance.GetProperty("DS_msDS_ServiceAllowedToAuthenticateTo")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(Uint8Array)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " Uint8Array is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = Uint8Array(valuetmp)

	return
}

// SetDS_msDS_ServiceAuthNPolicyBL sets the value of DS_msDS_ServiceAuthNPolicyBL for the instance
func (instance *ads_msds_authnpolicy) SetPropertyDS_msDS_ServiceAuthNPolicyBL(value []string) (err error) {
	return instance.SetProperty("DS_msDS_ServiceAuthNPolicyBL", (value))
}

// GetDS_msDS_ServiceAuthNPolicyBL gets the value of DS_msDS_ServiceAuthNPolicyBL for the instance
func (instance *ads_msds_authnpolicy) GetPropertyDS_msDS_ServiceAuthNPolicyBL() (value []string, err error) {
	retValue, err := instance.GetProperty("DS_msDS_ServiceAuthNPolicyBL")
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

// SetDS_msDS_ServiceTGTLifetime sets the value of DS_msDS_ServiceTGTLifetime for the instance
func (instance *ads_msds_authnpolicy) SetPropertyDS_msDS_ServiceTGTLifetime(value int64) (err error) {
	return instance.SetProperty("DS_msDS_ServiceTGTLifetime", (value))
}

// GetDS_msDS_ServiceTGTLifetime gets the value of DS_msDS_ServiceTGTLifetime for the instance
func (instance *ads_msds_authnpolicy) GetPropertyDS_msDS_ServiceTGTLifetime() (value int64, err error) {
	retValue, err := instance.GetProperty("DS_msDS_ServiceTGTLifetime")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(int64)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " int64 is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = int64(valuetmp)

	return
}

// SetDS_msDS_StrongNTLMPolicy sets the value of DS_msDS_StrongNTLMPolicy for the instance
func (instance *ads_msds_authnpolicy) SetPropertyDS_msDS_StrongNTLMPolicy(value int32) (err error) {
	return instance.SetProperty("DS_msDS_StrongNTLMPolicy", (value))
}

// GetDS_msDS_StrongNTLMPolicy gets the value of DS_msDS_StrongNTLMPolicy for the instance
func (instance *ads_msds_authnpolicy) GetPropertyDS_msDS_StrongNTLMPolicy() (value int32, err error) {
	retValue, err := instance.GetProperty("DS_msDS_StrongNTLMPolicy")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(int32)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " int32 is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = int32(valuetmp)

	return
}

// SetDS_msDS_UserAllowedNTLMNetworkAuthentication sets the value of DS_msDS_UserAllowedNTLMNetworkAuthentication for the instance
func (instance *ads_msds_authnpolicy) SetPropertyDS_msDS_UserAllowedNTLMNetworkAuthentication(value bool) (err error) {
	return instance.SetProperty("DS_msDS_UserAllowedNTLMNetworkAuthentication", (value))
}

// GetDS_msDS_UserAllowedNTLMNetworkAuthentication gets the value of DS_msDS_UserAllowedNTLMNetworkAuthentication for the instance
func (instance *ads_msds_authnpolicy) GetPropertyDS_msDS_UserAllowedNTLMNetworkAuthentication() (value bool, err error) {
	retValue, err := instance.GetProperty("DS_msDS_UserAllowedNTLMNetworkAuthentication")
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

// SetDS_msDS_UserAllowedToAuthenticateFrom sets the value of DS_msDS_UserAllowedToAuthenticateFrom for the instance
func (instance *ads_msds_authnpolicy) SetPropertyDS_msDS_UserAllowedToAuthenticateFrom(value Uint8Array) (err error) {
	return instance.SetProperty("DS_msDS_UserAllowedToAuthenticateFrom", (value))
}

// GetDS_msDS_UserAllowedToAuthenticateFrom gets the value of DS_msDS_UserAllowedToAuthenticateFrom for the instance
func (instance *ads_msds_authnpolicy) GetPropertyDS_msDS_UserAllowedToAuthenticateFrom() (value Uint8Array, err error) {
	retValue, err := instance.GetProperty("DS_msDS_UserAllowedToAuthenticateFrom")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(Uint8Array)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " Uint8Array is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = Uint8Array(valuetmp)

	return
}

// SetDS_msDS_UserAllowedToAuthenticateTo sets the value of DS_msDS_UserAllowedToAuthenticateTo for the instance
func (instance *ads_msds_authnpolicy) SetPropertyDS_msDS_UserAllowedToAuthenticateTo(value Uint8Array) (err error) {
	return instance.SetProperty("DS_msDS_UserAllowedToAuthenticateTo", (value))
}

// GetDS_msDS_UserAllowedToAuthenticateTo gets the value of DS_msDS_UserAllowedToAuthenticateTo for the instance
func (instance *ads_msds_authnpolicy) GetPropertyDS_msDS_UserAllowedToAuthenticateTo() (value Uint8Array, err error) {
	retValue, err := instance.GetProperty("DS_msDS_UserAllowedToAuthenticateTo")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(Uint8Array)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " Uint8Array is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = Uint8Array(valuetmp)

	return
}

// SetDS_msDS_UserAuthNPolicyBL sets the value of DS_msDS_UserAuthNPolicyBL for the instance
func (instance *ads_msds_authnpolicy) SetPropertyDS_msDS_UserAuthNPolicyBL(value []string) (err error) {
	return instance.SetProperty("DS_msDS_UserAuthNPolicyBL", (value))
}

// GetDS_msDS_UserAuthNPolicyBL gets the value of DS_msDS_UserAuthNPolicyBL for the instance
func (instance *ads_msds_authnpolicy) GetPropertyDS_msDS_UserAuthNPolicyBL() (value []string, err error) {
	retValue, err := instance.GetProperty("DS_msDS_UserAuthNPolicyBL")
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

// SetDS_msDS_UserTGTLifetime sets the value of DS_msDS_UserTGTLifetime for the instance
func (instance *ads_msds_authnpolicy) SetPropertyDS_msDS_UserTGTLifetime(value int64) (err error) {
	return instance.SetProperty("DS_msDS_UserTGTLifetime", (value))
}

// GetDS_msDS_UserTGTLifetime gets the value of DS_msDS_UserTGTLifetime for the instance
func (instance *ads_msds_authnpolicy) GetPropertyDS_msDS_UserTGTLifetime() (value int64, err error) {
	retValue, err := instance.GetProperty("DS_msDS_UserTGTLifetime")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(int64)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " int64 is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = int64(valuetmp)

	return
}
