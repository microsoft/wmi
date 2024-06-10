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

// ads_ntdsdsa struct
type ads_ntdsdsa struct {
	*ds_applicationsettings

	//
	DS_dMDLocation string

	//
	DS_fRSRootPath string

	//
	DS_hasMasterNCs []string

	//
	DS_hasPartialReplicaNCs []string

	//
	DS_invocationId Uint8Array

	//
	DS_lastBackupRestorationTime int64

	//
	DS_managedBy string

	//
	DS_msDS_Behavior_Version int32

	//
	DS_msDS_EnabledFeature []string

	//
	DS_msDS_HasDomainNCs []string

	//
	DS_msDS_hasFullReplicaNCs []string

	//
	DS_msDS_HasInstantiatedNCs []DN_With_Binary

	//
	DS_msDS_hasMasterNCs []string

	//
	DS_msDS_isGC bool

	//
	DS_msDS_isRODC bool

	//
	DS_msDS_IsUserCachableAtRodc int32

	//
	DS_msDS_NeverRevealGroup []string

	//
	DS_msDS_ReplicationEpoch int32

	//
	DS_msDS_RetiredReplNCSignatures Uint8Array

	//
	DS_msDS_RevealedUsers []DN_With_Binary

	//
	DS_msDS_RevealOnDemandGroup []string

	//
	DS_msDS_SiteName string

	//
	DS_networkAddress []string

	//
	DS_options int32

	//
	DS_queryPolicyObject string

	//
	DS_retiredReplDSASignatures Uint8Array

	//
	DS_serverReference string
}

func Newads_ntdsdsaEx1(instance *cim.WmiInstance) (newInstance *ads_ntdsdsa, err error) {
	tmp, err := Newds_applicationsettingsEx1(instance)

	if err != nil {
		return
	}
	newInstance = &ads_ntdsdsa{
		ds_applicationsettings: tmp,
	}
	return
}

func Newads_ntdsdsaEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *ads_ntdsdsa, err error) {
	tmp, err := Newds_applicationsettingsEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &ads_ntdsdsa{
		ds_applicationsettings: tmp,
	}
	return
}

// SetDS_dMDLocation sets the value of DS_dMDLocation for the instance
func (instance *ads_ntdsdsa) SetPropertyDS_dMDLocation(value string) (err error) {
	return instance.SetProperty("DS_dMDLocation", (value))
}

// GetDS_dMDLocation gets the value of DS_dMDLocation for the instance
func (instance *ads_ntdsdsa) GetPropertyDS_dMDLocation() (value string, err error) {
	retValue, err := instance.GetProperty("DS_dMDLocation")
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

// SetDS_fRSRootPath sets the value of DS_fRSRootPath for the instance
func (instance *ads_ntdsdsa) SetPropertyDS_fRSRootPath(value string) (err error) {
	return instance.SetProperty("DS_fRSRootPath", (value))
}

// GetDS_fRSRootPath gets the value of DS_fRSRootPath for the instance
func (instance *ads_ntdsdsa) GetPropertyDS_fRSRootPath() (value string, err error) {
	retValue, err := instance.GetProperty("DS_fRSRootPath")
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

// SetDS_hasMasterNCs sets the value of DS_hasMasterNCs for the instance
func (instance *ads_ntdsdsa) SetPropertyDS_hasMasterNCs(value []string) (err error) {
	return instance.SetProperty("DS_hasMasterNCs", (value))
}

// GetDS_hasMasterNCs gets the value of DS_hasMasterNCs for the instance
func (instance *ads_ntdsdsa) GetPropertyDS_hasMasterNCs() (value []string, err error) {
	retValue, err := instance.GetProperty("DS_hasMasterNCs")
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

// SetDS_hasPartialReplicaNCs sets the value of DS_hasPartialReplicaNCs for the instance
func (instance *ads_ntdsdsa) SetPropertyDS_hasPartialReplicaNCs(value []string) (err error) {
	return instance.SetProperty("DS_hasPartialReplicaNCs", (value))
}

// GetDS_hasPartialReplicaNCs gets the value of DS_hasPartialReplicaNCs for the instance
func (instance *ads_ntdsdsa) GetPropertyDS_hasPartialReplicaNCs() (value []string, err error) {
	retValue, err := instance.GetProperty("DS_hasPartialReplicaNCs")
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

// SetDS_invocationId sets the value of DS_invocationId for the instance
func (instance *ads_ntdsdsa) SetPropertyDS_invocationId(value Uint8Array) (err error) {
	return instance.SetProperty("DS_invocationId", (value))
}

// GetDS_invocationId gets the value of DS_invocationId for the instance
func (instance *ads_ntdsdsa) GetPropertyDS_invocationId() (value Uint8Array, err error) {
	retValue, err := instance.GetProperty("DS_invocationId")
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

// SetDS_lastBackupRestorationTime sets the value of DS_lastBackupRestorationTime for the instance
func (instance *ads_ntdsdsa) SetPropertyDS_lastBackupRestorationTime(value int64) (err error) {
	return instance.SetProperty("DS_lastBackupRestorationTime", (value))
}

// GetDS_lastBackupRestorationTime gets the value of DS_lastBackupRestorationTime for the instance
func (instance *ads_ntdsdsa) GetPropertyDS_lastBackupRestorationTime() (value int64, err error) {
	retValue, err := instance.GetProperty("DS_lastBackupRestorationTime")
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

// SetDS_managedBy sets the value of DS_managedBy for the instance
func (instance *ads_ntdsdsa) SetPropertyDS_managedBy(value string) (err error) {
	return instance.SetProperty("DS_managedBy", (value))
}

// GetDS_managedBy gets the value of DS_managedBy for the instance
func (instance *ads_ntdsdsa) GetPropertyDS_managedBy() (value string, err error) {
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

// SetDS_msDS_Behavior_Version sets the value of DS_msDS_Behavior_Version for the instance
func (instance *ads_ntdsdsa) SetPropertyDS_msDS_Behavior_Version(value int32) (err error) {
	return instance.SetProperty("DS_msDS_Behavior_Version", (value))
}

// GetDS_msDS_Behavior_Version gets the value of DS_msDS_Behavior_Version for the instance
func (instance *ads_ntdsdsa) GetPropertyDS_msDS_Behavior_Version() (value int32, err error) {
	retValue, err := instance.GetProperty("DS_msDS_Behavior_Version")
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

// SetDS_msDS_EnabledFeature sets the value of DS_msDS_EnabledFeature for the instance
func (instance *ads_ntdsdsa) SetPropertyDS_msDS_EnabledFeature(value []string) (err error) {
	return instance.SetProperty("DS_msDS_EnabledFeature", (value))
}

// GetDS_msDS_EnabledFeature gets the value of DS_msDS_EnabledFeature for the instance
func (instance *ads_ntdsdsa) GetPropertyDS_msDS_EnabledFeature() (value []string, err error) {
	retValue, err := instance.GetProperty("DS_msDS_EnabledFeature")
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

// SetDS_msDS_HasDomainNCs sets the value of DS_msDS_HasDomainNCs for the instance
func (instance *ads_ntdsdsa) SetPropertyDS_msDS_HasDomainNCs(value []string) (err error) {
	return instance.SetProperty("DS_msDS_HasDomainNCs", (value))
}

// GetDS_msDS_HasDomainNCs gets the value of DS_msDS_HasDomainNCs for the instance
func (instance *ads_ntdsdsa) GetPropertyDS_msDS_HasDomainNCs() (value []string, err error) {
	retValue, err := instance.GetProperty("DS_msDS_HasDomainNCs")
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

// SetDS_msDS_hasFullReplicaNCs sets the value of DS_msDS_hasFullReplicaNCs for the instance
func (instance *ads_ntdsdsa) SetPropertyDS_msDS_hasFullReplicaNCs(value []string) (err error) {
	return instance.SetProperty("DS_msDS_hasFullReplicaNCs", (value))
}

// GetDS_msDS_hasFullReplicaNCs gets the value of DS_msDS_hasFullReplicaNCs for the instance
func (instance *ads_ntdsdsa) GetPropertyDS_msDS_hasFullReplicaNCs() (value []string, err error) {
	retValue, err := instance.GetProperty("DS_msDS_hasFullReplicaNCs")
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

// SetDS_msDS_HasInstantiatedNCs sets the value of DS_msDS_HasInstantiatedNCs for the instance
func (instance *ads_ntdsdsa) SetPropertyDS_msDS_HasInstantiatedNCs(value []DN_With_Binary) (err error) {
	return instance.SetProperty("DS_msDS_HasInstantiatedNCs", (value))
}

// GetDS_msDS_HasInstantiatedNCs gets the value of DS_msDS_HasInstantiatedNCs for the instance
func (instance *ads_ntdsdsa) GetPropertyDS_msDS_HasInstantiatedNCs() (value []DN_With_Binary, err error) {
	retValue, err := instance.GetProperty("DS_msDS_HasInstantiatedNCs")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	for _, interfaceValue := range retValue.([]interface{}) {
		valuetmp, ok := interfaceValue.(DN_With_Binary)
		if !ok {
			err = errors.Wrapf(errors.InvalidType, " DN_With_Binary is Invalid. Expected %s", reflect.TypeOf(interfaceValue))
			return
		}
		value = append(value, DN_With_Binary(valuetmp))
	}

	return
}

// SetDS_msDS_hasMasterNCs sets the value of DS_msDS_hasMasterNCs for the instance
func (instance *ads_ntdsdsa) SetPropertyDS_msDS_hasMasterNCs(value []string) (err error) {
	return instance.SetProperty("DS_msDS_hasMasterNCs", (value))
}

// GetDS_msDS_hasMasterNCs gets the value of DS_msDS_hasMasterNCs for the instance
func (instance *ads_ntdsdsa) GetPropertyDS_msDS_hasMasterNCs() (value []string, err error) {
	retValue, err := instance.GetProperty("DS_msDS_hasMasterNCs")
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

// SetDS_msDS_isGC sets the value of DS_msDS_isGC for the instance
func (instance *ads_ntdsdsa) SetPropertyDS_msDS_isGC(value bool) (err error) {
	return instance.SetProperty("DS_msDS_isGC", (value))
}

// GetDS_msDS_isGC gets the value of DS_msDS_isGC for the instance
func (instance *ads_ntdsdsa) GetPropertyDS_msDS_isGC() (value bool, err error) {
	retValue, err := instance.GetProperty("DS_msDS_isGC")
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

// SetDS_msDS_isRODC sets the value of DS_msDS_isRODC for the instance
func (instance *ads_ntdsdsa) SetPropertyDS_msDS_isRODC(value bool) (err error) {
	return instance.SetProperty("DS_msDS_isRODC", (value))
}

// GetDS_msDS_isRODC gets the value of DS_msDS_isRODC for the instance
func (instance *ads_ntdsdsa) GetPropertyDS_msDS_isRODC() (value bool, err error) {
	retValue, err := instance.GetProperty("DS_msDS_isRODC")
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

// SetDS_msDS_IsUserCachableAtRodc sets the value of DS_msDS_IsUserCachableAtRodc for the instance
func (instance *ads_ntdsdsa) SetPropertyDS_msDS_IsUserCachableAtRodc(value int32) (err error) {
	return instance.SetProperty("DS_msDS_IsUserCachableAtRodc", (value))
}

// GetDS_msDS_IsUserCachableAtRodc gets the value of DS_msDS_IsUserCachableAtRodc for the instance
func (instance *ads_ntdsdsa) GetPropertyDS_msDS_IsUserCachableAtRodc() (value int32, err error) {
	retValue, err := instance.GetProperty("DS_msDS_IsUserCachableAtRodc")
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

// SetDS_msDS_NeverRevealGroup sets the value of DS_msDS_NeverRevealGroup for the instance
func (instance *ads_ntdsdsa) SetPropertyDS_msDS_NeverRevealGroup(value []string) (err error) {
	return instance.SetProperty("DS_msDS_NeverRevealGroup", (value))
}

// GetDS_msDS_NeverRevealGroup gets the value of DS_msDS_NeverRevealGroup for the instance
func (instance *ads_ntdsdsa) GetPropertyDS_msDS_NeverRevealGroup() (value []string, err error) {
	retValue, err := instance.GetProperty("DS_msDS_NeverRevealGroup")
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

// SetDS_msDS_ReplicationEpoch sets the value of DS_msDS_ReplicationEpoch for the instance
func (instance *ads_ntdsdsa) SetPropertyDS_msDS_ReplicationEpoch(value int32) (err error) {
	return instance.SetProperty("DS_msDS_ReplicationEpoch", (value))
}

// GetDS_msDS_ReplicationEpoch gets the value of DS_msDS_ReplicationEpoch for the instance
func (instance *ads_ntdsdsa) GetPropertyDS_msDS_ReplicationEpoch() (value int32, err error) {
	retValue, err := instance.GetProperty("DS_msDS_ReplicationEpoch")
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

// SetDS_msDS_RetiredReplNCSignatures sets the value of DS_msDS_RetiredReplNCSignatures for the instance
func (instance *ads_ntdsdsa) SetPropertyDS_msDS_RetiredReplNCSignatures(value Uint8Array) (err error) {
	return instance.SetProperty("DS_msDS_RetiredReplNCSignatures", (value))
}

// GetDS_msDS_RetiredReplNCSignatures gets the value of DS_msDS_RetiredReplNCSignatures for the instance
func (instance *ads_ntdsdsa) GetPropertyDS_msDS_RetiredReplNCSignatures() (value Uint8Array, err error) {
	retValue, err := instance.GetProperty("DS_msDS_RetiredReplNCSignatures")
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

// SetDS_msDS_RevealedUsers sets the value of DS_msDS_RevealedUsers for the instance
func (instance *ads_ntdsdsa) SetPropertyDS_msDS_RevealedUsers(value []DN_With_Binary) (err error) {
	return instance.SetProperty("DS_msDS_RevealedUsers", (value))
}

// GetDS_msDS_RevealedUsers gets the value of DS_msDS_RevealedUsers for the instance
func (instance *ads_ntdsdsa) GetPropertyDS_msDS_RevealedUsers() (value []DN_With_Binary, err error) {
	retValue, err := instance.GetProperty("DS_msDS_RevealedUsers")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	for _, interfaceValue := range retValue.([]interface{}) {
		valuetmp, ok := interfaceValue.(DN_With_Binary)
		if !ok {
			err = errors.Wrapf(errors.InvalidType, " DN_With_Binary is Invalid. Expected %s", reflect.TypeOf(interfaceValue))
			return
		}
		value = append(value, DN_With_Binary(valuetmp))
	}

	return
}

// SetDS_msDS_RevealOnDemandGroup sets the value of DS_msDS_RevealOnDemandGroup for the instance
func (instance *ads_ntdsdsa) SetPropertyDS_msDS_RevealOnDemandGroup(value []string) (err error) {
	return instance.SetProperty("DS_msDS_RevealOnDemandGroup", (value))
}

// GetDS_msDS_RevealOnDemandGroup gets the value of DS_msDS_RevealOnDemandGroup for the instance
func (instance *ads_ntdsdsa) GetPropertyDS_msDS_RevealOnDemandGroup() (value []string, err error) {
	retValue, err := instance.GetProperty("DS_msDS_RevealOnDemandGroup")
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

// SetDS_msDS_SiteName sets the value of DS_msDS_SiteName for the instance
func (instance *ads_ntdsdsa) SetPropertyDS_msDS_SiteName(value string) (err error) {
	return instance.SetProperty("DS_msDS_SiteName", (value))
}

// GetDS_msDS_SiteName gets the value of DS_msDS_SiteName for the instance
func (instance *ads_ntdsdsa) GetPropertyDS_msDS_SiteName() (value string, err error) {
	retValue, err := instance.GetProperty("DS_msDS_SiteName")
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

// SetDS_networkAddress sets the value of DS_networkAddress for the instance
func (instance *ads_ntdsdsa) SetPropertyDS_networkAddress(value []string) (err error) {
	return instance.SetProperty("DS_networkAddress", (value))
}

// GetDS_networkAddress gets the value of DS_networkAddress for the instance
func (instance *ads_ntdsdsa) GetPropertyDS_networkAddress() (value []string, err error) {
	retValue, err := instance.GetProperty("DS_networkAddress")
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

// SetDS_options sets the value of DS_options for the instance
func (instance *ads_ntdsdsa) SetPropertyDS_options(value int32) (err error) {
	return instance.SetProperty("DS_options", (value))
}

// GetDS_options gets the value of DS_options for the instance
func (instance *ads_ntdsdsa) GetPropertyDS_options() (value int32, err error) {
	retValue, err := instance.GetProperty("DS_options")
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

// SetDS_queryPolicyObject sets the value of DS_queryPolicyObject for the instance
func (instance *ads_ntdsdsa) SetPropertyDS_queryPolicyObject(value string) (err error) {
	return instance.SetProperty("DS_queryPolicyObject", (value))
}

// GetDS_queryPolicyObject gets the value of DS_queryPolicyObject for the instance
func (instance *ads_ntdsdsa) GetPropertyDS_queryPolicyObject() (value string, err error) {
	retValue, err := instance.GetProperty("DS_queryPolicyObject")
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

// SetDS_retiredReplDSASignatures sets the value of DS_retiredReplDSASignatures for the instance
func (instance *ads_ntdsdsa) SetPropertyDS_retiredReplDSASignatures(value Uint8Array) (err error) {
	return instance.SetProperty("DS_retiredReplDSASignatures", (value))
}

// GetDS_retiredReplDSASignatures gets the value of DS_retiredReplDSASignatures for the instance
func (instance *ads_ntdsdsa) GetPropertyDS_retiredReplDSASignatures() (value Uint8Array, err error) {
	retValue, err := instance.GetProperty("DS_retiredReplDSASignatures")
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

// SetDS_serverReference sets the value of DS_serverReference for the instance
func (instance *ads_ntdsdsa) SetPropertyDS_serverReference(value string) (err error) {
	return instance.SetProperty("DS_serverReference", (value))
}

// GetDS_serverReference gets the value of DS_serverReference for the instance
func (instance *ads_ntdsdsa) GetPropertyDS_serverReference() (value string, err error) {
	retValue, err := instance.GetProperty("DS_serverReference")
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
