// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 6/14/2024 using wmigen
//      Source root.directory.LDAP
//////////////////////////////////////////////
package ldap

import (
	"github.com/microsoft/wmi/pkg/base/query"
	"github.com/microsoft/wmi/pkg/errors"
	cim "github.com/microsoft/wmi/pkg/wmiinstance"
	"reflect"
)

// ds_samdomain struct
type ds_samdomain struct {
	*ds_top

	//
	DS_auditingPolicy Uint8Array

	//
	DS_builtinCreationTime int64

	//
	DS_builtinModifiedCount int64

	//
	DS_cACertificate []Uint8Array

	//
	DS_controlAccessRights []Uint8Array

	//
	DS_creationTime int64

	//
	DS_defaultLocalPolicyObject string

	//
	DS_desktopProfile string

	//
	DS_domainPolicyObject string

	//
	DS_domainReplica string

	//
	DS_eFSPolicy []Uint8Array

	//
	DS_forceLogoff int64

	//
	DS_gPLink string

	//
	DS_gPOptions int32

	//
	DS_lockoutDuration int64

	//
	DS_lockOutObservationWindow int64

	//
	DS_lockoutThreshold int32

	//
	DS_lSACreationTime int64

	//
	DS_lSAModifiedCount int64

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
	DS_ms_DS_MachineAccountQuota int32

	//
	DS_msDS_AllUsersTrustQuota int32

	//
	DS_msDS_LogonTimeSyncInterval int32

	//
	DS_msDS_PerUserTrustQuota int32

	//
	DS_msDS_PerUserTrustTombstonesQuota int32

	//
	DS_nETBIOSName string

	//
	DS_nextRid int32

	//
	DS_nTMixedDomain int32

	//
	DS_objectSid Uint8Array

	//
	DS_oEMInformation string

	//
	DS_pekKeyChangeInterval int64

	//
	DS_pekList Uint8Array

	//
	DS_privateKey Uint8Array

	//
	DS_pwdHistoryLength int32

	//
	DS_pwdProperties int32

	//
	DS_replicaSource string

	//
	DS_rIDManagerReference string

	//
	DS_serverRole int32

	//
	DS_serverState int32

	//
	DS_treeName string

	//
	DS_uASCompat int32
}

func Newds_samdomainEx1(instance *cim.WmiInstance) (newInstance *ds_samdomain, err error) {
	tmp, err := Newds_topEx1(instance)

	if err != nil {
		return
	}
	newInstance = &ds_samdomain{
		ds_top: tmp,
	}
	return
}

func Newds_samdomainEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *ds_samdomain, err error) {
	tmp, err := Newds_topEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &ds_samdomain{
		ds_top: tmp,
	}
	return
}

// SetDS_auditingPolicy sets the value of DS_auditingPolicy for the instance
func (instance *ds_samdomain) SetPropertyDS_auditingPolicy(value Uint8Array) (err error) {
	return instance.SetProperty("DS_auditingPolicy", (value))
}

// GetDS_auditingPolicy gets the value of DS_auditingPolicy for the instance
func (instance *ds_samdomain) GetPropertyDS_auditingPolicy() (value Uint8Array, err error) {
	retValue, err := instance.GetProperty("DS_auditingPolicy")
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

// SetDS_builtinCreationTime sets the value of DS_builtinCreationTime for the instance
func (instance *ds_samdomain) SetPropertyDS_builtinCreationTime(value int64) (err error) {
	return instance.SetProperty("DS_builtinCreationTime", (value))
}

// GetDS_builtinCreationTime gets the value of DS_builtinCreationTime for the instance
func (instance *ds_samdomain) GetPropertyDS_builtinCreationTime() (value int64, err error) {
	retValue, err := instance.GetProperty("DS_builtinCreationTime")
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

// SetDS_builtinModifiedCount sets the value of DS_builtinModifiedCount for the instance
func (instance *ds_samdomain) SetPropertyDS_builtinModifiedCount(value int64) (err error) {
	return instance.SetProperty("DS_builtinModifiedCount", (value))
}

// GetDS_builtinModifiedCount gets the value of DS_builtinModifiedCount for the instance
func (instance *ds_samdomain) GetPropertyDS_builtinModifiedCount() (value int64, err error) {
	retValue, err := instance.GetProperty("DS_builtinModifiedCount")
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

// SetDS_cACertificate sets the value of DS_cACertificate for the instance
func (instance *ds_samdomain) SetPropertyDS_cACertificate(value []Uint8Array) (err error) {
	return instance.SetProperty("DS_cACertificate", (value))
}

// GetDS_cACertificate gets the value of DS_cACertificate for the instance
func (instance *ds_samdomain) GetPropertyDS_cACertificate() (value []Uint8Array, err error) {
	retValue, err := instance.GetProperty("DS_cACertificate")
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

// SetDS_controlAccessRights sets the value of DS_controlAccessRights for the instance
func (instance *ds_samdomain) SetPropertyDS_controlAccessRights(value []Uint8Array) (err error) {
	return instance.SetProperty("DS_controlAccessRights", (value))
}

// GetDS_controlAccessRights gets the value of DS_controlAccessRights for the instance
func (instance *ds_samdomain) GetPropertyDS_controlAccessRights() (value []Uint8Array, err error) {
	retValue, err := instance.GetProperty("DS_controlAccessRights")
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

// SetDS_creationTime sets the value of DS_creationTime for the instance
func (instance *ds_samdomain) SetPropertyDS_creationTime(value int64) (err error) {
	return instance.SetProperty("DS_creationTime", (value))
}

// GetDS_creationTime gets the value of DS_creationTime for the instance
func (instance *ds_samdomain) GetPropertyDS_creationTime() (value int64, err error) {
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

// SetDS_defaultLocalPolicyObject sets the value of DS_defaultLocalPolicyObject for the instance
func (instance *ds_samdomain) SetPropertyDS_defaultLocalPolicyObject(value string) (err error) {
	return instance.SetProperty("DS_defaultLocalPolicyObject", (value))
}

// GetDS_defaultLocalPolicyObject gets the value of DS_defaultLocalPolicyObject for the instance
func (instance *ds_samdomain) GetPropertyDS_defaultLocalPolicyObject() (value string, err error) {
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

// SetDS_desktopProfile sets the value of DS_desktopProfile for the instance
func (instance *ds_samdomain) SetPropertyDS_desktopProfile(value string) (err error) {
	return instance.SetProperty("DS_desktopProfile", (value))
}

// GetDS_desktopProfile gets the value of DS_desktopProfile for the instance
func (instance *ds_samdomain) GetPropertyDS_desktopProfile() (value string, err error) {
	retValue, err := instance.GetProperty("DS_desktopProfile")
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

// SetDS_domainPolicyObject sets the value of DS_domainPolicyObject for the instance
func (instance *ds_samdomain) SetPropertyDS_domainPolicyObject(value string) (err error) {
	return instance.SetProperty("DS_domainPolicyObject", (value))
}

// GetDS_domainPolicyObject gets the value of DS_domainPolicyObject for the instance
func (instance *ds_samdomain) GetPropertyDS_domainPolicyObject() (value string, err error) {
	retValue, err := instance.GetProperty("DS_domainPolicyObject")
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

// SetDS_domainReplica sets the value of DS_domainReplica for the instance
func (instance *ds_samdomain) SetPropertyDS_domainReplica(value string) (err error) {
	return instance.SetProperty("DS_domainReplica", (value))
}

// GetDS_domainReplica gets the value of DS_domainReplica for the instance
func (instance *ds_samdomain) GetPropertyDS_domainReplica() (value string, err error) {
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

// SetDS_eFSPolicy sets the value of DS_eFSPolicy for the instance
func (instance *ds_samdomain) SetPropertyDS_eFSPolicy(value []Uint8Array) (err error) {
	return instance.SetProperty("DS_eFSPolicy", (value))
}

// GetDS_eFSPolicy gets the value of DS_eFSPolicy for the instance
func (instance *ds_samdomain) GetPropertyDS_eFSPolicy() (value []Uint8Array, err error) {
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
func (instance *ds_samdomain) SetPropertyDS_forceLogoff(value int64) (err error) {
	return instance.SetProperty("DS_forceLogoff", (value))
}

// GetDS_forceLogoff gets the value of DS_forceLogoff for the instance
func (instance *ds_samdomain) GetPropertyDS_forceLogoff() (value int64, err error) {
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

// SetDS_gPLink sets the value of DS_gPLink for the instance
func (instance *ds_samdomain) SetPropertyDS_gPLink(value string) (err error) {
	return instance.SetProperty("DS_gPLink", (value))
}

// GetDS_gPLink gets the value of DS_gPLink for the instance
func (instance *ds_samdomain) GetPropertyDS_gPLink() (value string, err error) {
	retValue, err := instance.GetProperty("DS_gPLink")
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

// SetDS_gPOptions sets the value of DS_gPOptions for the instance
func (instance *ds_samdomain) SetPropertyDS_gPOptions(value int32) (err error) {
	return instance.SetProperty("DS_gPOptions", (value))
}

// GetDS_gPOptions gets the value of DS_gPOptions for the instance
func (instance *ds_samdomain) GetPropertyDS_gPOptions() (value int32, err error) {
	retValue, err := instance.GetProperty("DS_gPOptions")
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

// SetDS_lockoutDuration sets the value of DS_lockoutDuration for the instance
func (instance *ds_samdomain) SetPropertyDS_lockoutDuration(value int64) (err error) {
	return instance.SetProperty("DS_lockoutDuration", (value))
}

// GetDS_lockoutDuration gets the value of DS_lockoutDuration for the instance
func (instance *ds_samdomain) GetPropertyDS_lockoutDuration() (value int64, err error) {
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
func (instance *ds_samdomain) SetPropertyDS_lockOutObservationWindow(value int64) (err error) {
	return instance.SetProperty("DS_lockOutObservationWindow", (value))
}

// GetDS_lockOutObservationWindow gets the value of DS_lockOutObservationWindow for the instance
func (instance *ds_samdomain) GetPropertyDS_lockOutObservationWindow() (value int64, err error) {
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
func (instance *ds_samdomain) SetPropertyDS_lockoutThreshold(value int32) (err error) {
	return instance.SetProperty("DS_lockoutThreshold", (value))
}

// GetDS_lockoutThreshold gets the value of DS_lockoutThreshold for the instance
func (instance *ds_samdomain) GetPropertyDS_lockoutThreshold() (value int32, err error) {
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

// SetDS_lSACreationTime sets the value of DS_lSACreationTime for the instance
func (instance *ds_samdomain) SetPropertyDS_lSACreationTime(value int64) (err error) {
	return instance.SetProperty("DS_lSACreationTime", (value))
}

// GetDS_lSACreationTime gets the value of DS_lSACreationTime for the instance
func (instance *ds_samdomain) GetPropertyDS_lSACreationTime() (value int64, err error) {
	retValue, err := instance.GetProperty("DS_lSACreationTime")
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

// SetDS_lSAModifiedCount sets the value of DS_lSAModifiedCount for the instance
func (instance *ds_samdomain) SetPropertyDS_lSAModifiedCount(value int64) (err error) {
	return instance.SetProperty("DS_lSAModifiedCount", (value))
}

// GetDS_lSAModifiedCount gets the value of DS_lSAModifiedCount for the instance
func (instance *ds_samdomain) GetPropertyDS_lSAModifiedCount() (value int64, err error) {
	retValue, err := instance.GetProperty("DS_lSAModifiedCount")
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

// SetDS_maxPwdAge sets the value of DS_maxPwdAge for the instance
func (instance *ds_samdomain) SetPropertyDS_maxPwdAge(value int64) (err error) {
	return instance.SetProperty("DS_maxPwdAge", (value))
}

// GetDS_maxPwdAge gets the value of DS_maxPwdAge for the instance
func (instance *ds_samdomain) GetPropertyDS_maxPwdAge() (value int64, err error) {
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
func (instance *ds_samdomain) SetPropertyDS_minPwdAge(value int64) (err error) {
	return instance.SetProperty("DS_minPwdAge", (value))
}

// GetDS_minPwdAge gets the value of DS_minPwdAge for the instance
func (instance *ds_samdomain) GetPropertyDS_minPwdAge() (value int64, err error) {
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
func (instance *ds_samdomain) SetPropertyDS_minPwdLength(value int32) (err error) {
	return instance.SetProperty("DS_minPwdLength", (value))
}

// GetDS_minPwdLength gets the value of DS_minPwdLength for the instance
func (instance *ds_samdomain) GetPropertyDS_minPwdLength() (value int32, err error) {
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
func (instance *ds_samdomain) SetPropertyDS_modifiedCount(value int64) (err error) {
	return instance.SetProperty("DS_modifiedCount", (value))
}

// GetDS_modifiedCount gets the value of DS_modifiedCount for the instance
func (instance *ds_samdomain) GetPropertyDS_modifiedCount() (value int64, err error) {
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
func (instance *ds_samdomain) SetPropertyDS_modifiedCountAtLastProm(value int64) (err error) {
	return instance.SetProperty("DS_modifiedCountAtLastProm", (value))
}

// GetDS_modifiedCountAtLastProm gets the value of DS_modifiedCountAtLastProm for the instance
func (instance *ds_samdomain) GetPropertyDS_modifiedCountAtLastProm() (value int64, err error) {
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

// SetDS_ms_DS_MachineAccountQuota sets the value of DS_ms_DS_MachineAccountQuota for the instance
func (instance *ds_samdomain) SetPropertyDS_ms_DS_MachineAccountQuota(value int32) (err error) {
	return instance.SetProperty("DS_ms_DS_MachineAccountQuota", (value))
}

// GetDS_ms_DS_MachineAccountQuota gets the value of DS_ms_DS_MachineAccountQuota for the instance
func (instance *ds_samdomain) GetPropertyDS_ms_DS_MachineAccountQuota() (value int32, err error) {
	retValue, err := instance.GetProperty("DS_ms_DS_MachineAccountQuota")
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

// SetDS_msDS_AllUsersTrustQuota sets the value of DS_msDS_AllUsersTrustQuota for the instance
func (instance *ds_samdomain) SetPropertyDS_msDS_AllUsersTrustQuota(value int32) (err error) {
	return instance.SetProperty("DS_msDS_AllUsersTrustQuota", (value))
}

// GetDS_msDS_AllUsersTrustQuota gets the value of DS_msDS_AllUsersTrustQuota for the instance
func (instance *ds_samdomain) GetPropertyDS_msDS_AllUsersTrustQuota() (value int32, err error) {
	retValue, err := instance.GetProperty("DS_msDS_AllUsersTrustQuota")
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

// SetDS_msDS_LogonTimeSyncInterval sets the value of DS_msDS_LogonTimeSyncInterval for the instance
func (instance *ds_samdomain) SetPropertyDS_msDS_LogonTimeSyncInterval(value int32) (err error) {
	return instance.SetProperty("DS_msDS_LogonTimeSyncInterval", (value))
}

// GetDS_msDS_LogonTimeSyncInterval gets the value of DS_msDS_LogonTimeSyncInterval for the instance
func (instance *ds_samdomain) GetPropertyDS_msDS_LogonTimeSyncInterval() (value int32, err error) {
	retValue, err := instance.GetProperty("DS_msDS_LogonTimeSyncInterval")
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

// SetDS_msDS_PerUserTrustQuota sets the value of DS_msDS_PerUserTrustQuota for the instance
func (instance *ds_samdomain) SetPropertyDS_msDS_PerUserTrustQuota(value int32) (err error) {
	return instance.SetProperty("DS_msDS_PerUserTrustQuota", (value))
}

// GetDS_msDS_PerUserTrustQuota gets the value of DS_msDS_PerUserTrustQuota for the instance
func (instance *ds_samdomain) GetPropertyDS_msDS_PerUserTrustQuota() (value int32, err error) {
	retValue, err := instance.GetProperty("DS_msDS_PerUserTrustQuota")
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

// SetDS_msDS_PerUserTrustTombstonesQuota sets the value of DS_msDS_PerUserTrustTombstonesQuota for the instance
func (instance *ds_samdomain) SetPropertyDS_msDS_PerUserTrustTombstonesQuota(value int32) (err error) {
	return instance.SetProperty("DS_msDS_PerUserTrustTombstonesQuota", (value))
}

// GetDS_msDS_PerUserTrustTombstonesQuota gets the value of DS_msDS_PerUserTrustTombstonesQuota for the instance
func (instance *ds_samdomain) GetPropertyDS_msDS_PerUserTrustTombstonesQuota() (value int32, err error) {
	retValue, err := instance.GetProperty("DS_msDS_PerUserTrustTombstonesQuota")
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

// SetDS_nETBIOSName sets the value of DS_nETBIOSName for the instance
func (instance *ds_samdomain) SetPropertyDS_nETBIOSName(value string) (err error) {
	return instance.SetProperty("DS_nETBIOSName", (value))
}

// GetDS_nETBIOSName gets the value of DS_nETBIOSName for the instance
func (instance *ds_samdomain) GetPropertyDS_nETBIOSName() (value string, err error) {
	retValue, err := instance.GetProperty("DS_nETBIOSName")
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

// SetDS_nextRid sets the value of DS_nextRid for the instance
func (instance *ds_samdomain) SetPropertyDS_nextRid(value int32) (err error) {
	return instance.SetProperty("DS_nextRid", (value))
}

// GetDS_nextRid gets the value of DS_nextRid for the instance
func (instance *ds_samdomain) GetPropertyDS_nextRid() (value int32, err error) {
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

// SetDS_nTMixedDomain sets the value of DS_nTMixedDomain for the instance
func (instance *ds_samdomain) SetPropertyDS_nTMixedDomain(value int32) (err error) {
	return instance.SetProperty("DS_nTMixedDomain", (value))
}

// GetDS_nTMixedDomain gets the value of DS_nTMixedDomain for the instance
func (instance *ds_samdomain) GetPropertyDS_nTMixedDomain() (value int32, err error) {
	retValue, err := instance.GetProperty("DS_nTMixedDomain")
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
func (instance *ds_samdomain) SetPropertyDS_objectSid(value Uint8Array) (err error) {
	return instance.SetProperty("DS_objectSid", (value))
}

// GetDS_objectSid gets the value of DS_objectSid for the instance
func (instance *ds_samdomain) GetPropertyDS_objectSid() (value Uint8Array, err error) {
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
func (instance *ds_samdomain) SetPropertyDS_oEMInformation(value string) (err error) {
	return instance.SetProperty("DS_oEMInformation", (value))
}

// GetDS_oEMInformation gets the value of DS_oEMInformation for the instance
func (instance *ds_samdomain) GetPropertyDS_oEMInformation() (value string, err error) {
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

// SetDS_pekKeyChangeInterval sets the value of DS_pekKeyChangeInterval for the instance
func (instance *ds_samdomain) SetPropertyDS_pekKeyChangeInterval(value int64) (err error) {
	return instance.SetProperty("DS_pekKeyChangeInterval", (value))
}

// GetDS_pekKeyChangeInterval gets the value of DS_pekKeyChangeInterval for the instance
func (instance *ds_samdomain) GetPropertyDS_pekKeyChangeInterval() (value int64, err error) {
	retValue, err := instance.GetProperty("DS_pekKeyChangeInterval")
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

// SetDS_pekList sets the value of DS_pekList for the instance
func (instance *ds_samdomain) SetPropertyDS_pekList(value Uint8Array) (err error) {
	return instance.SetProperty("DS_pekList", (value))
}

// GetDS_pekList gets the value of DS_pekList for the instance
func (instance *ds_samdomain) GetPropertyDS_pekList() (value Uint8Array, err error) {
	retValue, err := instance.GetProperty("DS_pekList")
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

// SetDS_privateKey sets the value of DS_privateKey for the instance
func (instance *ds_samdomain) SetPropertyDS_privateKey(value Uint8Array) (err error) {
	return instance.SetProperty("DS_privateKey", (value))
}

// GetDS_privateKey gets the value of DS_privateKey for the instance
func (instance *ds_samdomain) GetPropertyDS_privateKey() (value Uint8Array, err error) {
	retValue, err := instance.GetProperty("DS_privateKey")
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
func (instance *ds_samdomain) SetPropertyDS_pwdHistoryLength(value int32) (err error) {
	return instance.SetProperty("DS_pwdHistoryLength", (value))
}

// GetDS_pwdHistoryLength gets the value of DS_pwdHistoryLength for the instance
func (instance *ds_samdomain) GetPropertyDS_pwdHistoryLength() (value int32, err error) {
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
func (instance *ds_samdomain) SetPropertyDS_pwdProperties(value int32) (err error) {
	return instance.SetProperty("DS_pwdProperties", (value))
}

// GetDS_pwdProperties gets the value of DS_pwdProperties for the instance
func (instance *ds_samdomain) GetPropertyDS_pwdProperties() (value int32, err error) {
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

// SetDS_replicaSource sets the value of DS_replicaSource for the instance
func (instance *ds_samdomain) SetPropertyDS_replicaSource(value string) (err error) {
	return instance.SetProperty("DS_replicaSource", (value))
}

// GetDS_replicaSource gets the value of DS_replicaSource for the instance
func (instance *ds_samdomain) GetPropertyDS_replicaSource() (value string, err error) {
	retValue, err := instance.GetProperty("DS_replicaSource")
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

// SetDS_rIDManagerReference sets the value of DS_rIDManagerReference for the instance
func (instance *ds_samdomain) SetPropertyDS_rIDManagerReference(value string) (err error) {
	return instance.SetProperty("DS_rIDManagerReference", (value))
}

// GetDS_rIDManagerReference gets the value of DS_rIDManagerReference for the instance
func (instance *ds_samdomain) GetPropertyDS_rIDManagerReference() (value string, err error) {
	retValue, err := instance.GetProperty("DS_rIDManagerReference")
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

// SetDS_serverRole sets the value of DS_serverRole for the instance
func (instance *ds_samdomain) SetPropertyDS_serverRole(value int32) (err error) {
	return instance.SetProperty("DS_serverRole", (value))
}

// GetDS_serverRole gets the value of DS_serverRole for the instance
func (instance *ds_samdomain) GetPropertyDS_serverRole() (value int32, err error) {
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
func (instance *ds_samdomain) SetPropertyDS_serverState(value int32) (err error) {
	return instance.SetProperty("DS_serverState", (value))
}

// GetDS_serverState gets the value of DS_serverState for the instance
func (instance *ds_samdomain) GetPropertyDS_serverState() (value int32, err error) {
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

// SetDS_treeName sets the value of DS_treeName for the instance
func (instance *ds_samdomain) SetPropertyDS_treeName(value string) (err error) {
	return instance.SetProperty("DS_treeName", (value))
}

// GetDS_treeName gets the value of DS_treeName for the instance
func (instance *ds_samdomain) GetPropertyDS_treeName() (value string, err error) {
	retValue, err := instance.GetProperty("DS_treeName")
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

// SetDS_uASCompat sets the value of DS_uASCompat for the instance
func (instance *ds_samdomain) SetPropertyDS_uASCompat(value int32) (err error) {
	return instance.SetProperty("DS_uASCompat", (value))
}

// GetDS_uASCompat gets the value of DS_uASCompat for the instance
func (instance *ds_samdomain) GetPropertyDS_uASCompat() (value int32, err error) {
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
