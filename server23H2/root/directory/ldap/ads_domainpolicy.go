// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

// Author:
//
//	Auto Generated on 6/6/2024 using wmigen
//	Source root.directory.LDAP
//
// ////////////////////////////////////////////
package ldap

import (
	"github.com/microsoft/wmi/pkg/base/query"
	"github.com/microsoft/wmi/pkg/errors"
	cim "github.com/microsoft/wmi/pkg/wmiinstance"
	"reflect"
)

// ads_domainpolicy struct
type ads_domainpolicy struct {
	*ds_leaf

	//
	DS_authenticationOptions int32

	//
	DS_defaultLocalPolicyObject string

	//
	DS_domainCAs []string

	//
	DS_domainPolicyReference string

	//
	DS_domainWidePolicy []Uint8Array

	//
	DS_eFSPolicy []Uint8Array

	//
	DS_forceLogoff int64

	//
	DS_ipsecPolicyReference string

	//
	DS_lockoutDuration int64

	//
	DS_lockOutObservationWindow int64

	//
	DS_lockoutThreshold int32

	//
	DS_managedBy string

	//
	DS_maxPwdAge int64

	//
	DS_maxRenewAge int64

	//
	DS_maxTicketAge int64

	//
	DS_minPwdAge int64

	//
	DS_minPwdLength int32

	//
	DS_minTicketAge int64

	//
	DS_proxyLifetime int64

	//
	DS_publicKeyPolicy Uint8Array

	//
	DS_pwdHistoryLength int32

	//
	DS_pwdProperties int32

	//
	DS_qualityOfService int32
}

func Newads_domainpolicyEx1(instance *cim.WmiInstance) (newInstance *ads_domainpolicy, err error) {
	tmp, err := Newds_leafEx1(instance)

	if err != nil {
		return
	}
	newInstance = &ads_domainpolicy{
		ds_leaf: tmp,
	}
	return
}

func Newads_domainpolicyEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *ads_domainpolicy, err error) {
	tmp, err := Newds_leafEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &ads_domainpolicy{
		ds_leaf: tmp,
	}
	return
}

// SetDS_authenticationOptions sets the value of DS_authenticationOptions for the instance
func (instance *ads_domainpolicy) SetPropertyDS_authenticationOptions(value int32) (err error) {
	return instance.SetProperty("DS_authenticationOptions", (value))
}

// GetDS_authenticationOptions gets the value of DS_authenticationOptions for the instance
func (instance *ads_domainpolicy) GetPropertyDS_authenticationOptions() (value int32, err error) {
	retValue, err := instance.GetProperty("DS_authenticationOptions")
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

// SetDS_defaultLocalPolicyObject sets the value of DS_defaultLocalPolicyObject for the instance
func (instance *ads_domainpolicy) SetPropertyDS_defaultLocalPolicyObject(value string) (err error) {
	return instance.SetProperty("DS_defaultLocalPolicyObject", (value))
}

// GetDS_defaultLocalPolicyObject gets the value of DS_defaultLocalPolicyObject for the instance
func (instance *ads_domainpolicy) GetPropertyDS_defaultLocalPolicyObject() (value string, err error) {
	retValue, err := instance.GetProperty("DS_defaultLocalPolicyObject")
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

// SetDS_domainCAs sets the value of DS_domainCAs for the instance
func (instance *ads_domainpolicy) SetPropertyDS_domainCAs(value []string) (err error) {
	return instance.SetProperty("DS_domainCAs", (value))
}

// GetDS_domainCAs gets the value of DS_domainCAs for the instance
func (instance *ads_domainpolicy) GetPropertyDS_domainCAs() (value []string, err error) {
	retValue, err := instance.GetProperty("DS_domainCAs")
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

// SetDS_domainPolicyReference sets the value of DS_domainPolicyReference for the instance
func (instance *ads_domainpolicy) SetPropertyDS_domainPolicyReference(value string) (err error) {
	return instance.SetProperty("DS_domainPolicyReference", (value))
}

// GetDS_domainPolicyReference gets the value of DS_domainPolicyReference for the instance
func (instance *ads_domainpolicy) GetPropertyDS_domainPolicyReference() (value string, err error) {
	retValue, err := instance.GetProperty("DS_domainPolicyReference")
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

// SetDS_domainWidePolicy sets the value of DS_domainWidePolicy for the instance
func (instance *ads_domainpolicy) SetPropertyDS_domainWidePolicy(value []Uint8Array) (err error) {
	return instance.SetProperty("DS_domainWidePolicy", (value))
}

// GetDS_domainWidePolicy gets the value of DS_domainWidePolicy for the instance
func (instance *ads_domainpolicy) GetPropertyDS_domainWidePolicy() (value []Uint8Array, err error) {
	retValue, err := instance.GetProperty("DS_domainWidePolicy")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	for _, interfaceValue := range retValue.([]interface{}) {
		valuetmp, ok := interfaceValue.(Uint8Array)
		if !ok {
			err = errors.Wrapf(errors.InvalidType, " Uint8Array is Invalid. Expected %s", reflect.TypeOf(interfaceValue))
			return
		}
		value = append(value, Uint8Array(valuetmp))
	}

	return
}

// SetDS_eFSPolicy sets the value of DS_eFSPolicy for the instance
func (instance *ads_domainpolicy) SetPropertyDS_eFSPolicy(value []Uint8Array) (err error) {
	return instance.SetProperty("DS_eFSPolicy", (value))
}

// GetDS_eFSPolicy gets the value of DS_eFSPolicy for the instance
func (instance *ads_domainpolicy) GetPropertyDS_eFSPolicy() (value []Uint8Array, err error) {
	retValue, err := instance.GetProperty("DS_eFSPolicy")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	for _, interfaceValue := range retValue.([]interface{}) {
		valuetmp, ok := interfaceValue.(Uint8Array)
		if !ok {
			err = errors.Wrapf(errors.InvalidType, " Uint8Array is Invalid. Expected %s", reflect.TypeOf(interfaceValue))
			return
		}
		value = append(value, Uint8Array(valuetmp))
	}

	return
}

// SetDS_forceLogoff sets the value of DS_forceLogoff for the instance
func (instance *ads_domainpolicy) SetPropertyDS_forceLogoff(value int64) (err error) {
	return instance.SetProperty("DS_forceLogoff", (value))
}

// GetDS_forceLogoff gets the value of DS_forceLogoff for the instance
func (instance *ads_domainpolicy) GetPropertyDS_forceLogoff() (value int64, err error) {
	retValue, err := instance.GetProperty("DS_forceLogoff")
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

// SetDS_ipsecPolicyReference sets the value of DS_ipsecPolicyReference for the instance
func (instance *ads_domainpolicy) SetPropertyDS_ipsecPolicyReference(value string) (err error) {
	return instance.SetProperty("DS_ipsecPolicyReference", (value))
}

// GetDS_ipsecPolicyReference gets the value of DS_ipsecPolicyReference for the instance
func (instance *ads_domainpolicy) GetPropertyDS_ipsecPolicyReference() (value string, err error) {
	retValue, err := instance.GetProperty("DS_ipsecPolicyReference")
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

// SetDS_lockoutDuration sets the value of DS_lockoutDuration for the instance
func (instance *ads_domainpolicy) SetPropertyDS_lockoutDuration(value int64) (err error) {
	return instance.SetProperty("DS_lockoutDuration", (value))
}

// GetDS_lockoutDuration gets the value of DS_lockoutDuration for the instance
func (instance *ads_domainpolicy) GetPropertyDS_lockoutDuration() (value int64, err error) {
	retValue, err := instance.GetProperty("DS_lockoutDuration")
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

// SetDS_lockOutObservationWindow sets the value of DS_lockOutObservationWindow for the instance
func (instance *ads_domainpolicy) SetPropertyDS_lockOutObservationWindow(value int64) (err error) {
	return instance.SetProperty("DS_lockOutObservationWindow", (value))
}

// GetDS_lockOutObservationWindow gets the value of DS_lockOutObservationWindow for the instance
func (instance *ads_domainpolicy) GetPropertyDS_lockOutObservationWindow() (value int64, err error) {
	retValue, err := instance.GetProperty("DS_lockOutObservationWindow")
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

// SetDS_lockoutThreshold sets the value of DS_lockoutThreshold for the instance
func (instance *ads_domainpolicy) SetPropertyDS_lockoutThreshold(value int32) (err error) {
	return instance.SetProperty("DS_lockoutThreshold", (value))
}

// GetDS_lockoutThreshold gets the value of DS_lockoutThreshold for the instance
func (instance *ads_domainpolicy) GetPropertyDS_lockoutThreshold() (value int32, err error) {
	retValue, err := instance.GetProperty("DS_lockoutThreshold")
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

// SetDS_managedBy sets the value of DS_managedBy for the instance
func (instance *ads_domainpolicy) SetPropertyDS_managedBy(value string) (err error) {
	return instance.SetProperty("DS_managedBy", (value))
}

// GetDS_managedBy gets the value of DS_managedBy for the instance
func (instance *ads_domainpolicy) GetPropertyDS_managedBy() (value string, err error) {
	retValue, err := instance.GetProperty("DS_managedBy")
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

// SetDS_maxPwdAge sets the value of DS_maxPwdAge for the instance
func (instance *ads_domainpolicy) SetPropertyDS_maxPwdAge(value int64) (err error) {
	return instance.SetProperty("DS_maxPwdAge", (value))
}

// GetDS_maxPwdAge gets the value of DS_maxPwdAge for the instance
func (instance *ads_domainpolicy) GetPropertyDS_maxPwdAge() (value int64, err error) {
	retValue, err := instance.GetProperty("DS_maxPwdAge")
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

// SetDS_maxRenewAge sets the value of DS_maxRenewAge for the instance
func (instance *ads_domainpolicy) SetPropertyDS_maxRenewAge(value int64) (err error) {
	return instance.SetProperty("DS_maxRenewAge", (value))
}

// GetDS_maxRenewAge gets the value of DS_maxRenewAge for the instance
func (instance *ads_domainpolicy) GetPropertyDS_maxRenewAge() (value int64, err error) {
	retValue, err := instance.GetProperty("DS_maxRenewAge")
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

// SetDS_maxTicketAge sets the value of DS_maxTicketAge for the instance
func (instance *ads_domainpolicy) SetPropertyDS_maxTicketAge(value int64) (err error) {
	return instance.SetProperty("DS_maxTicketAge", (value))
}

// GetDS_maxTicketAge gets the value of DS_maxTicketAge for the instance
func (instance *ads_domainpolicy) GetPropertyDS_maxTicketAge() (value int64, err error) {
	retValue, err := instance.GetProperty("DS_maxTicketAge")
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

// SetDS_minPwdAge sets the value of DS_minPwdAge for the instance
func (instance *ads_domainpolicy) SetPropertyDS_minPwdAge(value int64) (err error) {
	return instance.SetProperty("DS_minPwdAge", (value))
}

// GetDS_minPwdAge gets the value of DS_minPwdAge for the instance
func (instance *ads_domainpolicy) GetPropertyDS_minPwdAge() (value int64, err error) {
	retValue, err := instance.GetProperty("DS_minPwdAge")
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

// SetDS_minPwdLength sets the value of DS_minPwdLength for the instance
func (instance *ads_domainpolicy) SetPropertyDS_minPwdLength(value int32) (err error) {
	return instance.SetProperty("DS_minPwdLength", (value))
}

// GetDS_minPwdLength gets the value of DS_minPwdLength for the instance
func (instance *ads_domainpolicy) GetPropertyDS_minPwdLength() (value int32, err error) {
	retValue, err := instance.GetProperty("DS_minPwdLength")
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

// SetDS_minTicketAge sets the value of DS_minTicketAge for the instance
func (instance *ads_domainpolicy) SetPropertyDS_minTicketAge(value int64) (err error) {
	return instance.SetProperty("DS_minTicketAge", (value))
}

// GetDS_minTicketAge gets the value of DS_minTicketAge for the instance
func (instance *ads_domainpolicy) GetPropertyDS_minTicketAge() (value int64, err error) {
	retValue, err := instance.GetProperty("DS_minTicketAge")
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

// SetDS_proxyLifetime sets the value of DS_proxyLifetime for the instance
func (instance *ads_domainpolicy) SetPropertyDS_proxyLifetime(value int64) (err error) {
	return instance.SetProperty("DS_proxyLifetime", (value))
}

// GetDS_proxyLifetime gets the value of DS_proxyLifetime for the instance
func (instance *ads_domainpolicy) GetPropertyDS_proxyLifetime() (value int64, err error) {
	retValue, err := instance.GetProperty("DS_proxyLifetime")
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

// SetDS_publicKeyPolicy sets the value of DS_publicKeyPolicy for the instance
func (instance *ads_domainpolicy) SetPropertyDS_publicKeyPolicy(value Uint8Array) (err error) {
	return instance.SetProperty("DS_publicKeyPolicy", (value))
}

// GetDS_publicKeyPolicy gets the value of DS_publicKeyPolicy for the instance
func (instance *ads_domainpolicy) GetPropertyDS_publicKeyPolicy() (value Uint8Array, err error) {
	retValue, err := instance.GetProperty("DS_publicKeyPolicy")
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

// SetDS_pwdHistoryLength sets the value of DS_pwdHistoryLength for the instance
func (instance *ads_domainpolicy) SetPropertyDS_pwdHistoryLength(value int32) (err error) {
	return instance.SetProperty("DS_pwdHistoryLength", (value))
}

// GetDS_pwdHistoryLength gets the value of DS_pwdHistoryLength for the instance
func (instance *ads_domainpolicy) GetPropertyDS_pwdHistoryLength() (value int32, err error) {
	retValue, err := instance.GetProperty("DS_pwdHistoryLength")
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

// SetDS_pwdProperties sets the value of DS_pwdProperties for the instance
func (instance *ads_domainpolicy) SetPropertyDS_pwdProperties(value int32) (err error) {
	return instance.SetProperty("DS_pwdProperties", (value))
}

// GetDS_pwdProperties gets the value of DS_pwdProperties for the instance
func (instance *ads_domainpolicy) GetPropertyDS_pwdProperties() (value int32, err error) {
	retValue, err := instance.GetProperty("DS_pwdProperties")
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

// SetDS_qualityOfService sets the value of DS_qualityOfService for the instance
func (instance *ads_domainpolicy) SetPropertyDS_qualityOfService(value int32) (err error) {
	return instance.SetProperty("DS_qualityOfService", (value))
}

// GetDS_qualityOfService gets the value of DS_qualityOfService for the instance
func (instance *ads_domainpolicy) GetPropertyDS_qualityOfService() (value int32, err error) {
	retValue, err := instance.GetProperty("DS_qualityOfService")
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
