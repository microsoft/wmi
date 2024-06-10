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

// ads_builtindomain struct
type ads_builtindomain struct {
	*ds_top

	//
	DS_creationTime int64

	//
	DS_domainReplica string

	//
	DS_forceLogoff int64

	//
	DS_lockoutDuration int64

	//
	DS_lockOutObservationWindow int64

	//
	DS_lockoutThreshold int32

	//
	DS_maxPwdAge int64

	//
	DS_minPwdAge int64

	//
	DS_minPwdLength int32

	//
	DS_modifiedCount int64

	//
	DS_modifiedCountAtLastProm int64

	//
	DS_nextRid int32

	//
	DS_objectSid Uint8Array

	//
	DS_oEMInformation string

	//
	DS_pwdHistoryLength int32

	//
	DS_pwdProperties int32

	//
	DS_serverRole int32

	//
	DS_serverState int32

	//
	DS_uASCompat int32
}

func Newads_builtindomainEx1(instance *cim.WmiInstance) (newInstance *ads_builtindomain, err error) {
	tmp, err := Newds_topEx1(instance)

	if err != nil {
		return
	}
	newInstance = &ads_builtindomain{
		ds_top: tmp,
	}
	return
}

func Newads_builtindomainEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *ads_builtindomain, err error) {
	tmp, err := Newds_topEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &ads_builtindomain{
		ds_top: tmp,
	}
	return
}

// SetDS_creationTime sets the value of DS_creationTime for the instance
func (instance *ads_builtindomain) SetPropertyDS_creationTime(value int64) (err error) {
	return instance.SetProperty("DS_creationTime", (value))
}

// GetDS_creationTime gets the value of DS_creationTime for the instance
func (instance *ads_builtindomain) GetPropertyDS_creationTime() (value int64, err error) {
	retValue, err := instance.GetProperty("DS_creationTime")
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

// SetDS_domainReplica sets the value of DS_domainReplica for the instance
func (instance *ads_builtindomain) SetPropertyDS_domainReplica(value string) (err error) {
	return instance.SetProperty("DS_domainReplica", (value))
}

// GetDS_domainReplica gets the value of DS_domainReplica for the instance
func (instance *ads_builtindomain) GetPropertyDS_domainReplica() (value string, err error) {
	retValue, err := instance.GetProperty("DS_domainReplica")
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

// SetDS_forceLogoff sets the value of DS_forceLogoff for the instance
func (instance *ads_builtindomain) SetPropertyDS_forceLogoff(value int64) (err error) {
	return instance.SetProperty("DS_forceLogoff", (value))
}

// GetDS_forceLogoff gets the value of DS_forceLogoff for the instance
func (instance *ads_builtindomain) GetPropertyDS_forceLogoff() (value int64, err error) {
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

// SetDS_lockoutDuration sets the value of DS_lockoutDuration for the instance
func (instance *ads_builtindomain) SetPropertyDS_lockoutDuration(value int64) (err error) {
	return instance.SetProperty("DS_lockoutDuration", (value))
}

// GetDS_lockoutDuration gets the value of DS_lockoutDuration for the instance
func (instance *ads_builtindomain) GetPropertyDS_lockoutDuration() (value int64, err error) {
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
func (instance *ads_builtindomain) SetPropertyDS_lockOutObservationWindow(value int64) (err error) {
	return instance.SetProperty("DS_lockOutObservationWindow", (value))
}

// GetDS_lockOutObservationWindow gets the value of DS_lockOutObservationWindow for the instance
func (instance *ads_builtindomain) GetPropertyDS_lockOutObservationWindow() (value int64, err error) {
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
func (instance *ads_builtindomain) SetPropertyDS_lockoutThreshold(value int32) (err error) {
	return instance.SetProperty("DS_lockoutThreshold", (value))
}

// GetDS_lockoutThreshold gets the value of DS_lockoutThreshold for the instance
func (instance *ads_builtindomain) GetPropertyDS_lockoutThreshold() (value int32, err error) {
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

// SetDS_maxPwdAge sets the value of DS_maxPwdAge for the instance
func (instance *ads_builtindomain) SetPropertyDS_maxPwdAge(value int64) (err error) {
	return instance.SetProperty("DS_maxPwdAge", (value))
}

// GetDS_maxPwdAge gets the value of DS_maxPwdAge for the instance
func (instance *ads_builtindomain) GetPropertyDS_maxPwdAge() (value int64, err error) {
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

// SetDS_minPwdAge sets the value of DS_minPwdAge for the instance
func (instance *ads_builtindomain) SetPropertyDS_minPwdAge(value int64) (err error) {
	return instance.SetProperty("DS_minPwdAge", (value))
}

// GetDS_minPwdAge gets the value of DS_minPwdAge for the instance
func (instance *ads_builtindomain) GetPropertyDS_minPwdAge() (value int64, err error) {
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
func (instance *ads_builtindomain) SetPropertyDS_minPwdLength(value int32) (err error) {
	return instance.SetProperty("DS_minPwdLength", (value))
}

// GetDS_minPwdLength gets the value of DS_minPwdLength for the instance
func (instance *ads_builtindomain) GetPropertyDS_minPwdLength() (value int32, err error) {
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

// SetDS_modifiedCount sets the value of DS_modifiedCount for the instance
func (instance *ads_builtindomain) SetPropertyDS_modifiedCount(value int64) (err error) {
	return instance.SetProperty("DS_modifiedCount", (value))
}

// GetDS_modifiedCount gets the value of DS_modifiedCount for the instance
func (instance *ads_builtindomain) GetPropertyDS_modifiedCount() (value int64, err error) {
	retValue, err := instance.GetProperty("DS_modifiedCount")
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

// SetDS_modifiedCountAtLastProm sets the value of DS_modifiedCountAtLastProm for the instance
func (instance *ads_builtindomain) SetPropertyDS_modifiedCountAtLastProm(value int64) (err error) {
	return instance.SetProperty("DS_modifiedCountAtLastProm", (value))
}

// GetDS_modifiedCountAtLastProm gets the value of DS_modifiedCountAtLastProm for the instance
func (instance *ads_builtindomain) GetPropertyDS_modifiedCountAtLastProm() (value int64, err error) {
	retValue, err := instance.GetProperty("DS_modifiedCountAtLastProm")
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

// SetDS_nextRid sets the value of DS_nextRid for the instance
func (instance *ads_builtindomain) SetPropertyDS_nextRid(value int32) (err error) {
	return instance.SetProperty("DS_nextRid", (value))
}

// GetDS_nextRid gets the value of DS_nextRid for the instance
func (instance *ads_builtindomain) GetPropertyDS_nextRid() (value int32, err error) {
	retValue, err := instance.GetProperty("DS_nextRid")
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

// SetDS_objectSid sets the value of DS_objectSid for the instance
func (instance *ads_builtindomain) SetPropertyDS_objectSid(value Uint8Array) (err error) {
	return instance.SetProperty("DS_objectSid", (value))
}

// GetDS_objectSid gets the value of DS_objectSid for the instance
func (instance *ads_builtindomain) GetPropertyDS_objectSid() (value Uint8Array, err error) {
	retValue, err := instance.GetProperty("DS_objectSid")
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

// SetDS_oEMInformation sets the value of DS_oEMInformation for the instance
func (instance *ads_builtindomain) SetPropertyDS_oEMInformation(value string) (err error) {
	return instance.SetProperty("DS_oEMInformation", (value))
}

// GetDS_oEMInformation gets the value of DS_oEMInformation for the instance
func (instance *ads_builtindomain) GetPropertyDS_oEMInformation() (value string, err error) {
	retValue, err := instance.GetProperty("DS_oEMInformation")
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

// SetDS_pwdHistoryLength sets the value of DS_pwdHistoryLength for the instance
func (instance *ads_builtindomain) SetPropertyDS_pwdHistoryLength(value int32) (err error) {
	return instance.SetProperty("DS_pwdHistoryLength", (value))
}

// GetDS_pwdHistoryLength gets the value of DS_pwdHistoryLength for the instance
func (instance *ads_builtindomain) GetPropertyDS_pwdHistoryLength() (value int32, err error) {
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
func (instance *ads_builtindomain) SetPropertyDS_pwdProperties(value int32) (err error) {
	return instance.SetProperty("DS_pwdProperties", (value))
}

// GetDS_pwdProperties gets the value of DS_pwdProperties for the instance
func (instance *ads_builtindomain) GetPropertyDS_pwdProperties() (value int32, err error) {
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

// SetDS_serverRole sets the value of DS_serverRole for the instance
func (instance *ads_builtindomain) SetPropertyDS_serverRole(value int32) (err error) {
	return instance.SetProperty("DS_serverRole", (value))
}

// GetDS_serverRole gets the value of DS_serverRole for the instance
func (instance *ads_builtindomain) GetPropertyDS_serverRole() (value int32, err error) {
	retValue, err := instance.GetProperty("DS_serverRole")
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

// SetDS_serverState sets the value of DS_serverState for the instance
func (instance *ads_builtindomain) SetPropertyDS_serverState(value int32) (err error) {
	return instance.SetProperty("DS_serverState", (value))
}

// GetDS_serverState gets the value of DS_serverState for the instance
func (instance *ads_builtindomain) GetPropertyDS_serverState() (value int32, err error) {
	retValue, err := instance.GetProperty("DS_serverState")
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

// SetDS_uASCompat sets the value of DS_uASCompat for the instance
func (instance *ads_builtindomain) SetPropertyDS_uASCompat(value int32) (err error) {
	return instance.SetProperty("DS_uASCompat", (value))
}

// GetDS_uASCompat gets the value of DS_uASCompat for the instance
func (instance *ads_builtindomain) GetPropertyDS_uASCompat() (value int32, err error) {
	retValue, err := instance.GetProperty("DS_uASCompat")
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
