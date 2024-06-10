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

// ads_computer struct
type ads_computer struct {
	*ads_user

	//
	DS_catalogs []string

	//
	DS_defaultLocalPolicyObject string

	//
	DS_dNSHostName string

	//
	DS_ipHostNumber []string

	//
	DS_localPolicyFlags int32

	//
	DS_location string

	//
	DS_machineRole int32

	//
	DS_managedBy string

	//
	DS_msDS_AdditionalDnsHostName []string

	//
	DS_msDS_AdditionalSamAccountName []string

	//
	DS_msDS_ExecuteScriptPassword Uint8Array

	//
	DS_msDS_GenerationId Uint8Array

	//
	DS_msDS_HostServiceAccount []string

	//
	DS_msDS_isGC bool

	//
	DS_msDS_isRODC bool

	//
	DS_msDS_IsUserCachableAtRodc int32

	//
	DS_msDS_KrbTgtLink string

	//
	DS_msDS_NeverRevealGroup []string

	//
	DS_msDS_PromotionSettings string

	//
	DS_msDS_RevealedList []DN_With_Binary

	//
	DS_msDS_RevealedUsers []DN_With_Binary

	//
	DS_msDS_RevealOnDemandGroup []string

	//
	DS_msDS_SiteName string

	//
	DS_msImaging_HashAlgorithm string

	//
	DS_msImaging_ThumbprintHash Uint8Array

	//
	DS_msSFU30Aliases []string

	//
	DS_msTPM_OwnerInformation string

	//
	DS_msTPM_TpmInformationForComputer string

	//
	DS_msTSEndpointData string

	//
	DS_msTSEndpointPlugin string

	//
	DS_msTSEndpointType int32

	//
	DS_msTSPrimaryDesktopBL []string

	//
	DS_msTSSecondaryDesktopBL []string

	//
	DS_netbootDUID Uint8Array

	//
	DS_netbootGUID Uint8Array

	//
	DS_netbootInitialization string

	//
	DS_netbootMachineFilePath string

	//
	DS_netbootMirrorDataFile []string

	//
	DS_netbootSIFFile []string

	//
	DS_nisMapName string

	//
	DS_operatingSystem string

	//
	DS_operatingSystemHotfix string

	//
	DS_operatingSystemServicePack string

	//
	DS_operatingSystemVersion string

	//
	DS_physicalLocationObject string

	//
	DS_policyReplicationFlags int32

	//
	DS_rIDSetReferences []string

	//
	DS_siteGUID Uint8Array

	//
	DS_volumeCount int32
}

func Newads_computerEx1(instance *cim.WmiInstance) (newInstance *ads_computer, err error) {
	tmp, err := Newads_userEx1(instance)

	if err != nil {
		return
	}
	newInstance = &ads_computer{
		ads_user: tmp,
	}
	return
}

func Newads_computerEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *ads_computer, err error) {
	tmp, err := Newads_userEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &ads_computer{
		ads_user: tmp,
	}
	return
}

// SetDS_catalogs sets the value of DS_catalogs for the instance
func (instance *ads_computer) SetPropertyDS_catalogs(value []string) (err error) {
	return instance.SetProperty("DS_catalogs", (value))
}

// GetDS_catalogs gets the value of DS_catalogs for the instance
func (instance *ads_computer) GetPropertyDS_catalogs() (value []string, err error) {
	retValue, err := instance.GetProperty("DS_catalogs")
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

// SetDS_defaultLocalPolicyObject sets the value of DS_defaultLocalPolicyObject for the instance
func (instance *ads_computer) SetPropertyDS_defaultLocalPolicyObject(value string) (err error) {
	return instance.SetProperty("DS_defaultLocalPolicyObject", (value))
}

// GetDS_defaultLocalPolicyObject gets the value of DS_defaultLocalPolicyObject for the instance
func (instance *ads_computer) GetPropertyDS_defaultLocalPolicyObject() (value string, err error) {
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

// SetDS_dNSHostName sets the value of DS_dNSHostName for the instance
func (instance *ads_computer) SetPropertyDS_dNSHostName(value string) (err error) {
	return instance.SetProperty("DS_dNSHostName", (value))
}

// GetDS_dNSHostName gets the value of DS_dNSHostName for the instance
func (instance *ads_computer) GetPropertyDS_dNSHostName() (value string, err error) {
	retValue, err := instance.GetProperty("DS_dNSHostName")
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

// SetDS_ipHostNumber sets the value of DS_ipHostNumber for the instance
func (instance *ads_computer) SetPropertyDS_ipHostNumber(value []string) (err error) {
	return instance.SetProperty("DS_ipHostNumber", (value))
}

// GetDS_ipHostNumber gets the value of DS_ipHostNumber for the instance
func (instance *ads_computer) GetPropertyDS_ipHostNumber() (value []string, err error) {
	retValue, err := instance.GetProperty("DS_ipHostNumber")
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

// SetDS_localPolicyFlags sets the value of DS_localPolicyFlags for the instance
func (instance *ads_computer) SetPropertyDS_localPolicyFlags(value int32) (err error) {
	return instance.SetProperty("DS_localPolicyFlags", (value))
}

// GetDS_localPolicyFlags gets the value of DS_localPolicyFlags for the instance
func (instance *ads_computer) GetPropertyDS_localPolicyFlags() (value int32, err error) {
	retValue, err := instance.GetProperty("DS_localPolicyFlags")
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

// SetDS_location sets the value of DS_location for the instance
func (instance *ads_computer) SetPropertyDS_location(value string) (err error) {
	return instance.SetProperty("DS_location", (value))
}

// GetDS_location gets the value of DS_location for the instance
func (instance *ads_computer) GetPropertyDS_location() (value string, err error) {
	retValue, err := instance.GetProperty("DS_location")
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

// SetDS_machineRole sets the value of DS_machineRole for the instance
func (instance *ads_computer) SetPropertyDS_machineRole(value int32) (err error) {
	return instance.SetProperty("DS_machineRole", (value))
}

// GetDS_machineRole gets the value of DS_machineRole for the instance
func (instance *ads_computer) GetPropertyDS_machineRole() (value int32, err error) {
	retValue, err := instance.GetProperty("DS_machineRole")
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
func (instance *ads_computer) SetPropertyDS_managedBy(value string) (err error) {
	return instance.SetProperty("DS_managedBy", (value))
}

// GetDS_managedBy gets the value of DS_managedBy for the instance
func (instance *ads_computer) GetPropertyDS_managedBy() (value string, err error) {
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

// SetDS_msDS_AdditionalDnsHostName sets the value of DS_msDS_AdditionalDnsHostName for the instance
func (instance *ads_computer) SetPropertyDS_msDS_AdditionalDnsHostName(value []string) (err error) {
	return instance.SetProperty("DS_msDS_AdditionalDnsHostName", (value))
}

// GetDS_msDS_AdditionalDnsHostName gets the value of DS_msDS_AdditionalDnsHostName for the instance
func (instance *ads_computer) GetPropertyDS_msDS_AdditionalDnsHostName() (value []string, err error) {
	retValue, err := instance.GetProperty("DS_msDS_AdditionalDnsHostName")
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

// SetDS_msDS_AdditionalSamAccountName sets the value of DS_msDS_AdditionalSamAccountName for the instance
func (instance *ads_computer) SetPropertyDS_msDS_AdditionalSamAccountName(value []string) (err error) {
	return instance.SetProperty("DS_msDS_AdditionalSamAccountName", (value))
}

// GetDS_msDS_AdditionalSamAccountName gets the value of DS_msDS_AdditionalSamAccountName for the instance
func (instance *ads_computer) GetPropertyDS_msDS_AdditionalSamAccountName() (value []string, err error) {
	retValue, err := instance.GetProperty("DS_msDS_AdditionalSamAccountName")
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

// SetDS_msDS_ExecuteScriptPassword sets the value of DS_msDS_ExecuteScriptPassword for the instance
func (instance *ads_computer) SetPropertyDS_msDS_ExecuteScriptPassword(value Uint8Array) (err error) {
	return instance.SetProperty("DS_msDS_ExecuteScriptPassword", (value))
}

// GetDS_msDS_ExecuteScriptPassword gets the value of DS_msDS_ExecuteScriptPassword for the instance
func (instance *ads_computer) GetPropertyDS_msDS_ExecuteScriptPassword() (value Uint8Array, err error) {
	retValue, err := instance.GetProperty("DS_msDS_ExecuteScriptPassword")
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

// SetDS_msDS_GenerationId sets the value of DS_msDS_GenerationId for the instance
func (instance *ads_computer) SetPropertyDS_msDS_GenerationId(value Uint8Array) (err error) {
	return instance.SetProperty("DS_msDS_GenerationId", (value))
}

// GetDS_msDS_GenerationId gets the value of DS_msDS_GenerationId for the instance
func (instance *ads_computer) GetPropertyDS_msDS_GenerationId() (value Uint8Array, err error) {
	retValue, err := instance.GetProperty("DS_msDS_GenerationId")
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

// SetDS_msDS_HostServiceAccount sets the value of DS_msDS_HostServiceAccount for the instance
func (instance *ads_computer) SetPropertyDS_msDS_HostServiceAccount(value []string) (err error) {
	return instance.SetProperty("DS_msDS_HostServiceAccount", (value))
}

// GetDS_msDS_HostServiceAccount gets the value of DS_msDS_HostServiceAccount for the instance
func (instance *ads_computer) GetPropertyDS_msDS_HostServiceAccount() (value []string, err error) {
	retValue, err := instance.GetProperty("DS_msDS_HostServiceAccount")
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
func (instance *ads_computer) SetPropertyDS_msDS_isGC(value bool) (err error) {
	return instance.SetProperty("DS_msDS_isGC", (value))
}

// GetDS_msDS_isGC gets the value of DS_msDS_isGC for the instance
func (instance *ads_computer) GetPropertyDS_msDS_isGC() (value bool, err error) {
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
func (instance *ads_computer) SetPropertyDS_msDS_isRODC(value bool) (err error) {
	return instance.SetProperty("DS_msDS_isRODC", (value))
}

// GetDS_msDS_isRODC gets the value of DS_msDS_isRODC for the instance
func (instance *ads_computer) GetPropertyDS_msDS_isRODC() (value bool, err error) {
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
func (instance *ads_computer) SetPropertyDS_msDS_IsUserCachableAtRodc(value int32) (err error) {
	return instance.SetProperty("DS_msDS_IsUserCachableAtRodc", (value))
}

// GetDS_msDS_IsUserCachableAtRodc gets the value of DS_msDS_IsUserCachableAtRodc for the instance
func (instance *ads_computer) GetPropertyDS_msDS_IsUserCachableAtRodc() (value int32, err error) {
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

// SetDS_msDS_KrbTgtLink sets the value of DS_msDS_KrbTgtLink for the instance
func (instance *ads_computer) SetPropertyDS_msDS_KrbTgtLink(value string) (err error) {
	return instance.SetProperty("DS_msDS_KrbTgtLink", (value))
}

// GetDS_msDS_KrbTgtLink gets the value of DS_msDS_KrbTgtLink for the instance
func (instance *ads_computer) GetPropertyDS_msDS_KrbTgtLink() (value string, err error) {
	retValue, err := instance.GetProperty("DS_msDS_KrbTgtLink")
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

// SetDS_msDS_NeverRevealGroup sets the value of DS_msDS_NeverRevealGroup for the instance
func (instance *ads_computer) SetPropertyDS_msDS_NeverRevealGroup(value []string) (err error) {
	return instance.SetProperty("DS_msDS_NeverRevealGroup", (value))
}

// GetDS_msDS_NeverRevealGroup gets the value of DS_msDS_NeverRevealGroup for the instance
func (instance *ads_computer) GetPropertyDS_msDS_NeverRevealGroup() (value []string, err error) {
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

// SetDS_msDS_PromotionSettings sets the value of DS_msDS_PromotionSettings for the instance
func (instance *ads_computer) SetPropertyDS_msDS_PromotionSettings(value string) (err error) {
	return instance.SetProperty("DS_msDS_PromotionSettings", (value))
}

// GetDS_msDS_PromotionSettings gets the value of DS_msDS_PromotionSettings for the instance
func (instance *ads_computer) GetPropertyDS_msDS_PromotionSettings() (value string, err error) {
	retValue, err := instance.GetProperty("DS_msDS_PromotionSettings")
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

// SetDS_msDS_RevealedList sets the value of DS_msDS_RevealedList for the instance
func (instance *ads_computer) SetPropertyDS_msDS_RevealedList(value []DN_With_Binary) (err error) {
	return instance.SetProperty("DS_msDS_RevealedList", (value))
}

// GetDS_msDS_RevealedList gets the value of DS_msDS_RevealedList for the instance
func (instance *ads_computer) GetPropertyDS_msDS_RevealedList() (value []DN_With_Binary, err error) {
	retValue, err := instance.GetProperty("DS_msDS_RevealedList")
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

// SetDS_msDS_RevealedUsers sets the value of DS_msDS_RevealedUsers for the instance
func (instance *ads_computer) SetPropertyDS_msDS_RevealedUsers(value []DN_With_Binary) (err error) {
	return instance.SetProperty("DS_msDS_RevealedUsers", (value))
}

// GetDS_msDS_RevealedUsers gets the value of DS_msDS_RevealedUsers for the instance
func (instance *ads_computer) GetPropertyDS_msDS_RevealedUsers() (value []DN_With_Binary, err error) {
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
func (instance *ads_computer) SetPropertyDS_msDS_RevealOnDemandGroup(value []string) (err error) {
	return instance.SetProperty("DS_msDS_RevealOnDemandGroup", (value))
}

// GetDS_msDS_RevealOnDemandGroup gets the value of DS_msDS_RevealOnDemandGroup for the instance
func (instance *ads_computer) GetPropertyDS_msDS_RevealOnDemandGroup() (value []string, err error) {
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
func (instance *ads_computer) SetPropertyDS_msDS_SiteName(value string) (err error) {
	return instance.SetProperty("DS_msDS_SiteName", (value))
}

// GetDS_msDS_SiteName gets the value of DS_msDS_SiteName for the instance
func (instance *ads_computer) GetPropertyDS_msDS_SiteName() (value string, err error) {
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

// SetDS_msImaging_HashAlgorithm sets the value of DS_msImaging_HashAlgorithm for the instance
func (instance *ads_computer) SetPropertyDS_msImaging_HashAlgorithm(value string) (err error) {
	return instance.SetProperty("DS_msImaging_HashAlgorithm", (value))
}

// GetDS_msImaging_HashAlgorithm gets the value of DS_msImaging_HashAlgorithm for the instance
func (instance *ads_computer) GetPropertyDS_msImaging_HashAlgorithm() (value string, err error) {
	retValue, err := instance.GetProperty("DS_msImaging_HashAlgorithm")
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

// SetDS_msImaging_ThumbprintHash sets the value of DS_msImaging_ThumbprintHash for the instance
func (instance *ads_computer) SetPropertyDS_msImaging_ThumbprintHash(value Uint8Array) (err error) {
	return instance.SetProperty("DS_msImaging_ThumbprintHash", (value))
}

// GetDS_msImaging_ThumbprintHash gets the value of DS_msImaging_ThumbprintHash for the instance
func (instance *ads_computer) GetPropertyDS_msImaging_ThumbprintHash() (value Uint8Array, err error) {
	retValue, err := instance.GetProperty("DS_msImaging_ThumbprintHash")
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

// SetDS_msSFU30Aliases sets the value of DS_msSFU30Aliases for the instance
func (instance *ads_computer) SetPropertyDS_msSFU30Aliases(value []string) (err error) {
	return instance.SetProperty("DS_msSFU30Aliases", (value))
}

// GetDS_msSFU30Aliases gets the value of DS_msSFU30Aliases for the instance
func (instance *ads_computer) GetPropertyDS_msSFU30Aliases() (value []string, err error) {
	retValue, err := instance.GetProperty("DS_msSFU30Aliases")
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

// SetDS_msTPM_OwnerInformation sets the value of DS_msTPM_OwnerInformation for the instance
func (instance *ads_computer) SetPropertyDS_msTPM_OwnerInformation(value string) (err error) {
	return instance.SetProperty("DS_msTPM_OwnerInformation", (value))
}

// GetDS_msTPM_OwnerInformation gets the value of DS_msTPM_OwnerInformation for the instance
func (instance *ads_computer) GetPropertyDS_msTPM_OwnerInformation() (value string, err error) {
	retValue, err := instance.GetProperty("DS_msTPM_OwnerInformation")
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

// SetDS_msTPM_TpmInformationForComputer sets the value of DS_msTPM_TpmInformationForComputer for the instance
func (instance *ads_computer) SetPropertyDS_msTPM_TpmInformationForComputer(value string) (err error) {
	return instance.SetProperty("DS_msTPM_TpmInformationForComputer", (value))
}

// GetDS_msTPM_TpmInformationForComputer gets the value of DS_msTPM_TpmInformationForComputer for the instance
func (instance *ads_computer) GetPropertyDS_msTPM_TpmInformationForComputer() (value string, err error) {
	retValue, err := instance.GetProperty("DS_msTPM_TpmInformationForComputer")
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

// SetDS_msTSEndpointData sets the value of DS_msTSEndpointData for the instance
func (instance *ads_computer) SetPropertyDS_msTSEndpointData(value string) (err error) {
	return instance.SetProperty("DS_msTSEndpointData", (value))
}

// GetDS_msTSEndpointData gets the value of DS_msTSEndpointData for the instance
func (instance *ads_computer) GetPropertyDS_msTSEndpointData() (value string, err error) {
	retValue, err := instance.GetProperty("DS_msTSEndpointData")
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

// SetDS_msTSEndpointPlugin sets the value of DS_msTSEndpointPlugin for the instance
func (instance *ads_computer) SetPropertyDS_msTSEndpointPlugin(value string) (err error) {
	return instance.SetProperty("DS_msTSEndpointPlugin", (value))
}

// GetDS_msTSEndpointPlugin gets the value of DS_msTSEndpointPlugin for the instance
func (instance *ads_computer) GetPropertyDS_msTSEndpointPlugin() (value string, err error) {
	retValue, err := instance.GetProperty("DS_msTSEndpointPlugin")
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

// SetDS_msTSEndpointType sets the value of DS_msTSEndpointType for the instance
func (instance *ads_computer) SetPropertyDS_msTSEndpointType(value int32) (err error) {
	return instance.SetProperty("DS_msTSEndpointType", (value))
}

// GetDS_msTSEndpointType gets the value of DS_msTSEndpointType for the instance
func (instance *ads_computer) GetPropertyDS_msTSEndpointType() (value int32, err error) {
	retValue, err := instance.GetProperty("DS_msTSEndpointType")
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

// SetDS_msTSPrimaryDesktopBL sets the value of DS_msTSPrimaryDesktopBL for the instance
func (instance *ads_computer) SetPropertyDS_msTSPrimaryDesktopBL(value []string) (err error) {
	return instance.SetProperty("DS_msTSPrimaryDesktopBL", (value))
}

// GetDS_msTSPrimaryDesktopBL gets the value of DS_msTSPrimaryDesktopBL for the instance
func (instance *ads_computer) GetPropertyDS_msTSPrimaryDesktopBL() (value []string, err error) {
	retValue, err := instance.GetProperty("DS_msTSPrimaryDesktopBL")
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

// SetDS_msTSSecondaryDesktopBL sets the value of DS_msTSSecondaryDesktopBL for the instance
func (instance *ads_computer) SetPropertyDS_msTSSecondaryDesktopBL(value []string) (err error) {
	return instance.SetProperty("DS_msTSSecondaryDesktopBL", (value))
}

// GetDS_msTSSecondaryDesktopBL gets the value of DS_msTSSecondaryDesktopBL for the instance
func (instance *ads_computer) GetPropertyDS_msTSSecondaryDesktopBL() (value []string, err error) {
	retValue, err := instance.GetProperty("DS_msTSSecondaryDesktopBL")
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

// SetDS_netbootDUID sets the value of DS_netbootDUID for the instance
func (instance *ads_computer) SetPropertyDS_netbootDUID(value Uint8Array) (err error) {
	return instance.SetProperty("DS_netbootDUID", (value))
}

// GetDS_netbootDUID gets the value of DS_netbootDUID for the instance
func (instance *ads_computer) GetPropertyDS_netbootDUID() (value Uint8Array, err error) {
	retValue, err := instance.GetProperty("DS_netbootDUID")
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

// SetDS_netbootGUID sets the value of DS_netbootGUID for the instance
func (instance *ads_computer) SetPropertyDS_netbootGUID(value Uint8Array) (err error) {
	return instance.SetProperty("DS_netbootGUID", (value))
}

// GetDS_netbootGUID gets the value of DS_netbootGUID for the instance
func (instance *ads_computer) GetPropertyDS_netbootGUID() (value Uint8Array, err error) {
	retValue, err := instance.GetProperty("DS_netbootGUID")
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

// SetDS_netbootInitialization sets the value of DS_netbootInitialization for the instance
func (instance *ads_computer) SetPropertyDS_netbootInitialization(value string) (err error) {
	return instance.SetProperty("DS_netbootInitialization", (value))
}

// GetDS_netbootInitialization gets the value of DS_netbootInitialization for the instance
func (instance *ads_computer) GetPropertyDS_netbootInitialization() (value string, err error) {
	retValue, err := instance.GetProperty("DS_netbootInitialization")
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

// SetDS_netbootMachineFilePath sets the value of DS_netbootMachineFilePath for the instance
func (instance *ads_computer) SetPropertyDS_netbootMachineFilePath(value string) (err error) {
	return instance.SetProperty("DS_netbootMachineFilePath", (value))
}

// GetDS_netbootMachineFilePath gets the value of DS_netbootMachineFilePath for the instance
func (instance *ads_computer) GetPropertyDS_netbootMachineFilePath() (value string, err error) {
	retValue, err := instance.GetProperty("DS_netbootMachineFilePath")
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

// SetDS_netbootMirrorDataFile sets the value of DS_netbootMirrorDataFile for the instance
func (instance *ads_computer) SetPropertyDS_netbootMirrorDataFile(value []string) (err error) {
	return instance.SetProperty("DS_netbootMirrorDataFile", (value))
}

// GetDS_netbootMirrorDataFile gets the value of DS_netbootMirrorDataFile for the instance
func (instance *ads_computer) GetPropertyDS_netbootMirrorDataFile() (value []string, err error) {
	retValue, err := instance.GetProperty("DS_netbootMirrorDataFile")
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

// SetDS_netbootSIFFile sets the value of DS_netbootSIFFile for the instance
func (instance *ads_computer) SetPropertyDS_netbootSIFFile(value []string) (err error) {
	return instance.SetProperty("DS_netbootSIFFile", (value))
}

// GetDS_netbootSIFFile gets the value of DS_netbootSIFFile for the instance
func (instance *ads_computer) GetPropertyDS_netbootSIFFile() (value []string, err error) {
	retValue, err := instance.GetProperty("DS_netbootSIFFile")
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

// SetDS_nisMapName sets the value of DS_nisMapName for the instance
func (instance *ads_computer) SetPropertyDS_nisMapName(value string) (err error) {
	return instance.SetProperty("DS_nisMapName", (value))
}

// GetDS_nisMapName gets the value of DS_nisMapName for the instance
func (instance *ads_computer) GetPropertyDS_nisMapName() (value string, err error) {
	retValue, err := instance.GetProperty("DS_nisMapName")
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

// SetDS_operatingSystem sets the value of DS_operatingSystem for the instance
func (instance *ads_computer) SetPropertyDS_operatingSystem(value string) (err error) {
	return instance.SetProperty("DS_operatingSystem", (value))
}

// GetDS_operatingSystem gets the value of DS_operatingSystem for the instance
func (instance *ads_computer) GetPropertyDS_operatingSystem() (value string, err error) {
	retValue, err := instance.GetProperty("DS_operatingSystem")
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

// SetDS_operatingSystemHotfix sets the value of DS_operatingSystemHotfix for the instance
func (instance *ads_computer) SetPropertyDS_operatingSystemHotfix(value string) (err error) {
	return instance.SetProperty("DS_operatingSystemHotfix", (value))
}

// GetDS_operatingSystemHotfix gets the value of DS_operatingSystemHotfix for the instance
func (instance *ads_computer) GetPropertyDS_operatingSystemHotfix() (value string, err error) {
	retValue, err := instance.GetProperty("DS_operatingSystemHotfix")
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

// SetDS_operatingSystemServicePack sets the value of DS_operatingSystemServicePack for the instance
func (instance *ads_computer) SetPropertyDS_operatingSystemServicePack(value string) (err error) {
	return instance.SetProperty("DS_operatingSystemServicePack", (value))
}

// GetDS_operatingSystemServicePack gets the value of DS_operatingSystemServicePack for the instance
func (instance *ads_computer) GetPropertyDS_operatingSystemServicePack() (value string, err error) {
	retValue, err := instance.GetProperty("DS_operatingSystemServicePack")
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

// SetDS_operatingSystemVersion sets the value of DS_operatingSystemVersion for the instance
func (instance *ads_computer) SetPropertyDS_operatingSystemVersion(value string) (err error) {
	return instance.SetProperty("DS_operatingSystemVersion", (value))
}

// GetDS_operatingSystemVersion gets the value of DS_operatingSystemVersion for the instance
func (instance *ads_computer) GetPropertyDS_operatingSystemVersion() (value string, err error) {
	retValue, err := instance.GetProperty("DS_operatingSystemVersion")
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

// SetDS_physicalLocationObject sets the value of DS_physicalLocationObject for the instance
func (instance *ads_computer) SetPropertyDS_physicalLocationObject(value string) (err error) {
	return instance.SetProperty("DS_physicalLocationObject", (value))
}

// GetDS_physicalLocationObject gets the value of DS_physicalLocationObject for the instance
func (instance *ads_computer) GetPropertyDS_physicalLocationObject() (value string, err error) {
	retValue, err := instance.GetProperty("DS_physicalLocationObject")
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

// SetDS_policyReplicationFlags sets the value of DS_policyReplicationFlags for the instance
func (instance *ads_computer) SetPropertyDS_policyReplicationFlags(value int32) (err error) {
	return instance.SetProperty("DS_policyReplicationFlags", (value))
}

// GetDS_policyReplicationFlags gets the value of DS_policyReplicationFlags for the instance
func (instance *ads_computer) GetPropertyDS_policyReplicationFlags() (value int32, err error) {
	retValue, err := instance.GetProperty("DS_policyReplicationFlags")
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

// SetDS_rIDSetReferences sets the value of DS_rIDSetReferences for the instance
func (instance *ads_computer) SetPropertyDS_rIDSetReferences(value []string) (err error) {
	return instance.SetProperty("DS_rIDSetReferences", (value))
}

// GetDS_rIDSetReferences gets the value of DS_rIDSetReferences for the instance
func (instance *ads_computer) GetPropertyDS_rIDSetReferences() (value []string, err error) {
	retValue, err := instance.GetProperty("DS_rIDSetReferences")
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

// SetDS_siteGUID sets the value of DS_siteGUID for the instance
func (instance *ads_computer) SetPropertyDS_siteGUID(value Uint8Array) (err error) {
	return instance.SetProperty("DS_siteGUID", (value))
}

// GetDS_siteGUID gets the value of DS_siteGUID for the instance
func (instance *ads_computer) GetPropertyDS_siteGUID() (value Uint8Array, err error) {
	retValue, err := instance.GetProperty("DS_siteGUID")
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

// SetDS_volumeCount sets the value of DS_volumeCount for the instance
func (instance *ads_computer) SetPropertyDS_volumeCount(value int32) (err error) {
	return instance.SetProperty("DS_volumeCount", (value))
}

// GetDS_volumeCount gets the value of DS_volumeCount for the instance
func (instance *ads_computer) GetPropertyDS_volumeCount() (value int32, err error) {
	retValue, err := instance.GetProperty("DS_volumeCount")
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
