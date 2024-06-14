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

// ads_group struct
type ads_group struct {
	*ds_top

	//
	DS_accountNameHistory []string

	//
	DS_adminCount int32

	//
	DS_altSecurityIdentities []string

	//
	DS_controlAccessRights []Uint8Array

	//
	DS_desktopProfile string

	//
	DS_garbageCollPeriod int32

	//
	DS_gidNumber int32

	//
	DS_groupAttributes int32

	//
	DS_groupMembershipSAM Uint8Array

	//
	DS_groupType int32

	//
	DS_info string

	//
	DS_labeledURI []string

	//
	DS_legacyExchangeDN string

	//
	DS_mail string

	//
	DS_managedBy string

	//
	DS_member []string

	//
	DS_memberUid []string

	//
	DS_msDS_AzApplicationData string

	//
	DS_msDS_AzBizRule string

	//
	DS_msDS_AzBizRuleLanguage string

	//
	DS_msDS_AzGenericData string

	//
	DS_msDS_AzLastImportedBizRulePath string

	//
	DS_msDS_AzLDAPQuery string

	//
	DS_msDS_AzObjectGuid Uint8Array

	//
	DS_msDS_ExternalDirectoryObjectId string

	//
	DS_msDS_GeoCoordinatesAltitude int64

	//
	DS_msDS_GeoCoordinatesLatitude int64

	//
	DS_msDS_GeoCoordinatesLongitude int64

	//
	DS_msDS_KeyVersionNumber int32

	//
	DS_msDS_NonMembers []string

	//
	DS_msDS_PhoneticDisplayName string

	//
	DS_msDS_preferredDataLocation string

	//
	DS_msDS_PrimaryComputer []string

	//
	DS_msds_tokenGroupNames []string

	//
	DS_msds_tokenGroupNamesGlobalAndUniversal []string

	//
	DS_msds_tokenGroupNamesNoGCAcceptable []string

	//
	DS_msExchAssistantName string

	//
	DS_msExchLabeledURI []string

	//
	DS_msSFU30Name string

	//
	DS_msSFU30NisDomain string

	//
	DS_msSFU30PosixMember []string

	//
	DS_nonSecurityMember []string

	//
	DS_nTGroupMembers []Uint8Array

	//
	DS_objectSid Uint8Array

	//
	DS_operatorCount int32

	//
	DS_primaryGroupToken int32

	//
	DS_rid int32

	//
	DS_sAMAccountName string

	//
	DS_sAMAccountType int32

	//
	DS_secretary []string

	//
	DS_securityIdentifier Uint8Array

	//
	DS_showInAddressBook []string

	//
	DS_sIDHistory []Uint8Array

	//
	DS_supplementalCredentials []Uint8Array

	//
	DS_telephoneNumber string

	//
	DS_textEncodedORAddress string

	//
	DS_tokenGroups []Uint8Array

	//
	DS_tokenGroupsGlobalAndUniversal []Uint8Array

	//
	DS_tokenGroupsNoGCAcceptable []Uint8Array

	//
	DS_unixUserPassword []Uint8Array

	//
	DS_userCert Uint8Array

	//
	DS_userCertificate []Uint8Array

	//
	DS_userPassword []Uint8Array

	//
	DS_userSMIMECertificate []Uint8Array
}

func Newads_groupEx1(instance *cim.WmiInstance) (newInstance *ads_group, err error) {
	tmp, err := Newds_topEx1(instance)

	if err != nil {
		return
	}
	newInstance = &ads_group{
		ds_top: tmp,
	}
	return
}

func Newads_groupEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *ads_group, err error) {
	tmp, err := Newds_topEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &ads_group{
		ds_top: tmp,
	}
	return
}

// SetDS_accountNameHistory sets the value of DS_accountNameHistory for the instance
func (instance *ads_group) SetPropertyDS_accountNameHistory(value []string) (err error) {
	return instance.SetProperty("DS_accountNameHistory", (value))
}

// GetDS_accountNameHistory gets the value of DS_accountNameHistory for the instance
func (instance *ads_group) GetPropertyDS_accountNameHistory() (value []string, err error) {
	retValue, err := instance.GetProperty("DS_accountNameHistory")
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

// SetDS_adminCount sets the value of DS_adminCount for the instance
func (instance *ads_group) SetPropertyDS_adminCount(value int32) (err error) {
	return instance.SetProperty("DS_adminCount", (value))
}

// GetDS_adminCount gets the value of DS_adminCount for the instance
func (instance *ads_group) GetPropertyDS_adminCount() (value int32, err error) {
	retValue, err := instance.GetProperty("DS_adminCount")
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

// SetDS_altSecurityIdentities sets the value of DS_altSecurityIdentities for the instance
func (instance *ads_group) SetPropertyDS_altSecurityIdentities(value []string) (err error) {
	return instance.SetProperty("DS_altSecurityIdentities", (value))
}

// GetDS_altSecurityIdentities gets the value of DS_altSecurityIdentities for the instance
func (instance *ads_group) GetPropertyDS_altSecurityIdentities() (value []string, err error) {
	retValue, err := instance.GetProperty("DS_altSecurityIdentities")
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

// SetDS_controlAccessRights sets the value of DS_controlAccessRights for the instance
func (instance *ads_group) SetPropertyDS_controlAccessRights(value []Uint8Array) (err error) {
	return instance.SetProperty("DS_controlAccessRights", (value))
}

// GetDS_controlAccessRights gets the value of DS_controlAccessRights for the instance
func (instance *ads_group) GetPropertyDS_controlAccessRights() (value []Uint8Array, err error) {
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

// SetDS_desktopProfile sets the value of DS_desktopProfile for the instance
func (instance *ads_group) SetPropertyDS_desktopProfile(value string) (err error) {
	return instance.SetProperty("DS_desktopProfile", (value))
}

// GetDS_desktopProfile gets the value of DS_desktopProfile for the instance
func (instance *ads_group) GetPropertyDS_desktopProfile() (value string, err error) {
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

// SetDS_garbageCollPeriod sets the value of DS_garbageCollPeriod for the instance
func (instance *ads_group) SetPropertyDS_garbageCollPeriod(value int32) (err error) {
	return instance.SetProperty("DS_garbageCollPeriod", (value))
}

// GetDS_garbageCollPeriod gets the value of DS_garbageCollPeriod for the instance
func (instance *ads_group) GetPropertyDS_garbageCollPeriod() (value int32, err error) {
	retValue, err := instance.GetProperty("DS_garbageCollPeriod")
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

// SetDS_gidNumber sets the value of DS_gidNumber for the instance
func (instance *ads_group) SetPropertyDS_gidNumber(value int32) (err error) {
	return instance.SetProperty("DS_gidNumber", (value))
}

// GetDS_gidNumber gets the value of DS_gidNumber for the instance
func (instance *ads_group) GetPropertyDS_gidNumber() (value int32, err error) {
	retValue, err := instance.GetProperty("DS_gidNumber")
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

// SetDS_groupAttributes sets the value of DS_groupAttributes for the instance
func (instance *ads_group) SetPropertyDS_groupAttributes(value int32) (err error) {
	return instance.SetProperty("DS_groupAttributes", (value))
}

// GetDS_groupAttributes gets the value of DS_groupAttributes for the instance
func (instance *ads_group) GetPropertyDS_groupAttributes() (value int32, err error) {
	retValue, err := instance.GetProperty("DS_groupAttributes")
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

// SetDS_groupMembershipSAM sets the value of DS_groupMembershipSAM for the instance
func (instance *ads_group) SetPropertyDS_groupMembershipSAM(value Uint8Array) (err error) {
	return instance.SetProperty("DS_groupMembershipSAM", (value))
}

// GetDS_groupMembershipSAM gets the value of DS_groupMembershipSAM for the instance
func (instance *ads_group) GetPropertyDS_groupMembershipSAM() (value Uint8Array, err error) {
	retValue, err := instance.GetProperty("DS_groupMembershipSAM")
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

// SetDS_groupType sets the value of DS_groupType for the instance
func (instance *ads_group) SetPropertyDS_groupType(value int32) (err error) {
	return instance.SetProperty("DS_groupType", (value))
}

// GetDS_groupType gets the value of DS_groupType for the instance
func (instance *ads_group) GetPropertyDS_groupType() (value int32, err error) {
	retValue, err := instance.GetProperty("DS_groupType")
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

// SetDS_info sets the value of DS_info for the instance
func (instance *ads_group) SetPropertyDS_info(value string) (err error) {
	return instance.SetProperty("DS_info", (value))
}

// GetDS_info gets the value of DS_info for the instance
func (instance *ads_group) GetPropertyDS_info() (value string, err error) {
	retValue, err := instance.GetProperty("DS_info")
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

// SetDS_labeledURI sets the value of DS_labeledURI for the instance
func (instance *ads_group) SetPropertyDS_labeledURI(value []string) (err error) {
	return instance.SetProperty("DS_labeledURI", (value))
}

// GetDS_labeledURI gets the value of DS_labeledURI for the instance
func (instance *ads_group) GetPropertyDS_labeledURI() (value []string, err error) {
	retValue, err := instance.GetProperty("DS_labeledURI")
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

// SetDS_legacyExchangeDN sets the value of DS_legacyExchangeDN for the instance
func (instance *ads_group) SetPropertyDS_legacyExchangeDN(value string) (err error) {
	return instance.SetProperty("DS_legacyExchangeDN", (value))
}

// GetDS_legacyExchangeDN gets the value of DS_legacyExchangeDN for the instance
func (instance *ads_group) GetPropertyDS_legacyExchangeDN() (value string, err error) {
	retValue, err := instance.GetProperty("DS_legacyExchangeDN")
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

// SetDS_mail sets the value of DS_mail for the instance
func (instance *ads_group) SetPropertyDS_mail(value string) (err error) {
	return instance.SetProperty("DS_mail", (value))
}

// GetDS_mail gets the value of DS_mail for the instance
func (instance *ads_group) GetPropertyDS_mail() (value string, err error) {
	retValue, err := instance.GetProperty("DS_mail")
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

// SetDS_managedBy sets the value of DS_managedBy for the instance
func (instance *ads_group) SetPropertyDS_managedBy(value string) (err error) {
	return instance.SetProperty("DS_managedBy", (value))
}

// GetDS_managedBy gets the value of DS_managedBy for the instance
func (instance *ads_group) GetPropertyDS_managedBy() (value string, err error) {
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

// SetDS_member sets the value of DS_member for the instance
func (instance *ads_group) SetPropertyDS_member(value []string) (err error) {
	return instance.SetProperty("DS_member", (value))
}

// GetDS_member gets the value of DS_member for the instance
func (instance *ads_group) GetPropertyDS_member() (value []string, err error) {
	retValue, err := instance.GetProperty("DS_member")
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

// SetDS_memberUid sets the value of DS_memberUid for the instance
func (instance *ads_group) SetPropertyDS_memberUid(value []string) (err error) {
	return instance.SetProperty("DS_memberUid", (value))
}

// GetDS_memberUid gets the value of DS_memberUid for the instance
func (instance *ads_group) GetPropertyDS_memberUid() (value []string, err error) {
	retValue, err := instance.GetProperty("DS_memberUid")
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

// SetDS_msDS_AzApplicationData sets the value of DS_msDS_AzApplicationData for the instance
func (instance *ads_group) SetPropertyDS_msDS_AzApplicationData(value string) (err error) {
	return instance.SetProperty("DS_msDS_AzApplicationData", (value))
}

// GetDS_msDS_AzApplicationData gets the value of DS_msDS_AzApplicationData for the instance
func (instance *ads_group) GetPropertyDS_msDS_AzApplicationData() (value string, err error) {
	retValue, err := instance.GetProperty("DS_msDS_AzApplicationData")
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

// SetDS_msDS_AzBizRule sets the value of DS_msDS_AzBizRule for the instance
func (instance *ads_group) SetPropertyDS_msDS_AzBizRule(value string) (err error) {
	return instance.SetProperty("DS_msDS_AzBizRule", (value))
}

// GetDS_msDS_AzBizRule gets the value of DS_msDS_AzBizRule for the instance
func (instance *ads_group) GetPropertyDS_msDS_AzBizRule() (value string, err error) {
	retValue, err := instance.GetProperty("DS_msDS_AzBizRule")
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

// SetDS_msDS_AzBizRuleLanguage sets the value of DS_msDS_AzBizRuleLanguage for the instance
func (instance *ads_group) SetPropertyDS_msDS_AzBizRuleLanguage(value string) (err error) {
	return instance.SetProperty("DS_msDS_AzBizRuleLanguage", (value))
}

// GetDS_msDS_AzBizRuleLanguage gets the value of DS_msDS_AzBizRuleLanguage for the instance
func (instance *ads_group) GetPropertyDS_msDS_AzBizRuleLanguage() (value string, err error) {
	retValue, err := instance.GetProperty("DS_msDS_AzBizRuleLanguage")
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

// SetDS_msDS_AzGenericData sets the value of DS_msDS_AzGenericData for the instance
func (instance *ads_group) SetPropertyDS_msDS_AzGenericData(value string) (err error) {
	return instance.SetProperty("DS_msDS_AzGenericData", (value))
}

// GetDS_msDS_AzGenericData gets the value of DS_msDS_AzGenericData for the instance
func (instance *ads_group) GetPropertyDS_msDS_AzGenericData() (value string, err error) {
	retValue, err := instance.GetProperty("DS_msDS_AzGenericData")
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

// SetDS_msDS_AzLastImportedBizRulePath sets the value of DS_msDS_AzLastImportedBizRulePath for the instance
func (instance *ads_group) SetPropertyDS_msDS_AzLastImportedBizRulePath(value string) (err error) {
	return instance.SetProperty("DS_msDS_AzLastImportedBizRulePath", (value))
}

// GetDS_msDS_AzLastImportedBizRulePath gets the value of DS_msDS_AzLastImportedBizRulePath for the instance
func (instance *ads_group) GetPropertyDS_msDS_AzLastImportedBizRulePath() (value string, err error) {
	retValue, err := instance.GetProperty("DS_msDS_AzLastImportedBizRulePath")
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

// SetDS_msDS_AzLDAPQuery sets the value of DS_msDS_AzLDAPQuery for the instance
func (instance *ads_group) SetPropertyDS_msDS_AzLDAPQuery(value string) (err error) {
	return instance.SetProperty("DS_msDS_AzLDAPQuery", (value))
}

// GetDS_msDS_AzLDAPQuery gets the value of DS_msDS_AzLDAPQuery for the instance
func (instance *ads_group) GetPropertyDS_msDS_AzLDAPQuery() (value string, err error) {
	retValue, err := instance.GetProperty("DS_msDS_AzLDAPQuery")
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

// SetDS_msDS_AzObjectGuid sets the value of DS_msDS_AzObjectGuid for the instance
func (instance *ads_group) SetPropertyDS_msDS_AzObjectGuid(value Uint8Array) (err error) {
	return instance.SetProperty("DS_msDS_AzObjectGuid", (value))
}

// GetDS_msDS_AzObjectGuid gets the value of DS_msDS_AzObjectGuid for the instance
func (instance *ads_group) GetPropertyDS_msDS_AzObjectGuid() (value Uint8Array, err error) {
	retValue, err := instance.GetProperty("DS_msDS_AzObjectGuid")
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

// SetDS_msDS_ExternalDirectoryObjectId sets the value of DS_msDS_ExternalDirectoryObjectId for the instance
func (instance *ads_group) SetPropertyDS_msDS_ExternalDirectoryObjectId(value string) (err error) {
	return instance.SetProperty("DS_msDS_ExternalDirectoryObjectId", (value))
}

// GetDS_msDS_ExternalDirectoryObjectId gets the value of DS_msDS_ExternalDirectoryObjectId for the instance
func (instance *ads_group) GetPropertyDS_msDS_ExternalDirectoryObjectId() (value string, err error) {
	retValue, err := instance.GetProperty("DS_msDS_ExternalDirectoryObjectId")
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

// SetDS_msDS_GeoCoordinatesAltitude sets the value of DS_msDS_GeoCoordinatesAltitude for the instance
func (instance *ads_group) SetPropertyDS_msDS_GeoCoordinatesAltitude(value int64) (err error) {
	return instance.SetProperty("DS_msDS_GeoCoordinatesAltitude", (value))
}

// GetDS_msDS_GeoCoordinatesAltitude gets the value of DS_msDS_GeoCoordinatesAltitude for the instance
func (instance *ads_group) GetPropertyDS_msDS_GeoCoordinatesAltitude() (value int64, err error) {
	retValue, err := instance.GetProperty("DS_msDS_GeoCoordinatesAltitude")
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

// SetDS_msDS_GeoCoordinatesLatitude sets the value of DS_msDS_GeoCoordinatesLatitude for the instance
func (instance *ads_group) SetPropertyDS_msDS_GeoCoordinatesLatitude(value int64) (err error) {
	return instance.SetProperty("DS_msDS_GeoCoordinatesLatitude", (value))
}

// GetDS_msDS_GeoCoordinatesLatitude gets the value of DS_msDS_GeoCoordinatesLatitude for the instance
func (instance *ads_group) GetPropertyDS_msDS_GeoCoordinatesLatitude() (value int64, err error) {
	retValue, err := instance.GetProperty("DS_msDS_GeoCoordinatesLatitude")
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

// SetDS_msDS_GeoCoordinatesLongitude sets the value of DS_msDS_GeoCoordinatesLongitude for the instance
func (instance *ads_group) SetPropertyDS_msDS_GeoCoordinatesLongitude(value int64) (err error) {
	return instance.SetProperty("DS_msDS_GeoCoordinatesLongitude", (value))
}

// GetDS_msDS_GeoCoordinatesLongitude gets the value of DS_msDS_GeoCoordinatesLongitude for the instance
func (instance *ads_group) GetPropertyDS_msDS_GeoCoordinatesLongitude() (value int64, err error) {
	retValue, err := instance.GetProperty("DS_msDS_GeoCoordinatesLongitude")
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

// SetDS_msDS_KeyVersionNumber sets the value of DS_msDS_KeyVersionNumber for the instance
func (instance *ads_group) SetPropertyDS_msDS_KeyVersionNumber(value int32) (err error) {
	return instance.SetProperty("DS_msDS_KeyVersionNumber", (value))
}

// GetDS_msDS_KeyVersionNumber gets the value of DS_msDS_KeyVersionNumber for the instance
func (instance *ads_group) GetPropertyDS_msDS_KeyVersionNumber() (value int32, err error) {
	retValue, err := instance.GetProperty("DS_msDS_KeyVersionNumber")
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

// SetDS_msDS_NonMembers sets the value of DS_msDS_NonMembers for the instance
func (instance *ads_group) SetPropertyDS_msDS_NonMembers(value []string) (err error) {
	return instance.SetProperty("DS_msDS_NonMembers", (value))
}

// GetDS_msDS_NonMembers gets the value of DS_msDS_NonMembers for the instance
func (instance *ads_group) GetPropertyDS_msDS_NonMembers() (value []string, err error) {
	retValue, err := instance.GetProperty("DS_msDS_NonMembers")
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

// SetDS_msDS_PhoneticDisplayName sets the value of DS_msDS_PhoneticDisplayName for the instance
func (instance *ads_group) SetPropertyDS_msDS_PhoneticDisplayName(value string) (err error) {
	return instance.SetProperty("DS_msDS_PhoneticDisplayName", (value))
}

// GetDS_msDS_PhoneticDisplayName gets the value of DS_msDS_PhoneticDisplayName for the instance
func (instance *ads_group) GetPropertyDS_msDS_PhoneticDisplayName() (value string, err error) {
	retValue, err := instance.GetProperty("DS_msDS_PhoneticDisplayName")
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

// SetDS_msDS_preferredDataLocation sets the value of DS_msDS_preferredDataLocation for the instance
func (instance *ads_group) SetPropertyDS_msDS_preferredDataLocation(value string) (err error) {
	return instance.SetProperty("DS_msDS_preferredDataLocation", (value))
}

// GetDS_msDS_preferredDataLocation gets the value of DS_msDS_preferredDataLocation for the instance
func (instance *ads_group) GetPropertyDS_msDS_preferredDataLocation() (value string, err error) {
	retValue, err := instance.GetProperty("DS_msDS_preferredDataLocation")
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

// SetDS_msDS_PrimaryComputer sets the value of DS_msDS_PrimaryComputer for the instance
func (instance *ads_group) SetPropertyDS_msDS_PrimaryComputer(value []string) (err error) {
	return instance.SetProperty("DS_msDS_PrimaryComputer", (value))
}

// GetDS_msDS_PrimaryComputer gets the value of DS_msDS_PrimaryComputer for the instance
func (instance *ads_group) GetPropertyDS_msDS_PrimaryComputer() (value []string, err error) {
	retValue, err := instance.GetProperty("DS_msDS_PrimaryComputer")
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

// SetDS_msds_tokenGroupNames sets the value of DS_msds_tokenGroupNames for the instance
func (instance *ads_group) SetPropertyDS_msds_tokenGroupNames(value []string) (err error) {
	return instance.SetProperty("DS_msds_tokenGroupNames", (value))
}

// GetDS_msds_tokenGroupNames gets the value of DS_msds_tokenGroupNames for the instance
func (instance *ads_group) GetPropertyDS_msds_tokenGroupNames() (value []string, err error) {
	retValue, err := instance.GetProperty("DS_msds_tokenGroupNames")
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

// SetDS_msds_tokenGroupNamesGlobalAndUniversal sets the value of DS_msds_tokenGroupNamesGlobalAndUniversal for the instance
func (instance *ads_group) SetPropertyDS_msds_tokenGroupNamesGlobalAndUniversal(value []string) (err error) {
	return instance.SetProperty("DS_msds_tokenGroupNamesGlobalAndUniversal", (value))
}

// GetDS_msds_tokenGroupNamesGlobalAndUniversal gets the value of DS_msds_tokenGroupNamesGlobalAndUniversal for the instance
func (instance *ads_group) GetPropertyDS_msds_tokenGroupNamesGlobalAndUniversal() (value []string, err error) {
	retValue, err := instance.GetProperty("DS_msds_tokenGroupNamesGlobalAndUniversal")
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

// SetDS_msds_tokenGroupNamesNoGCAcceptable sets the value of DS_msds_tokenGroupNamesNoGCAcceptable for the instance
func (instance *ads_group) SetPropertyDS_msds_tokenGroupNamesNoGCAcceptable(value []string) (err error) {
	return instance.SetProperty("DS_msds_tokenGroupNamesNoGCAcceptable", (value))
}

// GetDS_msds_tokenGroupNamesNoGCAcceptable gets the value of DS_msds_tokenGroupNamesNoGCAcceptable for the instance
func (instance *ads_group) GetPropertyDS_msds_tokenGroupNamesNoGCAcceptable() (value []string, err error) {
	retValue, err := instance.GetProperty("DS_msds_tokenGroupNamesNoGCAcceptable")
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

// SetDS_msExchAssistantName sets the value of DS_msExchAssistantName for the instance
func (instance *ads_group) SetPropertyDS_msExchAssistantName(value string) (err error) {
	return instance.SetProperty("DS_msExchAssistantName", (value))
}

// GetDS_msExchAssistantName gets the value of DS_msExchAssistantName for the instance
func (instance *ads_group) GetPropertyDS_msExchAssistantName() (value string, err error) {
	retValue, err := instance.GetProperty("DS_msExchAssistantName")
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

// SetDS_msExchLabeledURI sets the value of DS_msExchLabeledURI for the instance
func (instance *ads_group) SetPropertyDS_msExchLabeledURI(value []string) (err error) {
	return instance.SetProperty("DS_msExchLabeledURI", (value))
}

// GetDS_msExchLabeledURI gets the value of DS_msExchLabeledURI for the instance
func (instance *ads_group) GetPropertyDS_msExchLabeledURI() (value []string, err error) {
	retValue, err := instance.GetProperty("DS_msExchLabeledURI")
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

// SetDS_msSFU30Name sets the value of DS_msSFU30Name for the instance
func (instance *ads_group) SetPropertyDS_msSFU30Name(value string) (err error) {
	return instance.SetProperty("DS_msSFU30Name", (value))
}

// GetDS_msSFU30Name gets the value of DS_msSFU30Name for the instance
func (instance *ads_group) GetPropertyDS_msSFU30Name() (value string, err error) {
	retValue, err := instance.GetProperty("DS_msSFU30Name")
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

// SetDS_msSFU30NisDomain sets the value of DS_msSFU30NisDomain for the instance
func (instance *ads_group) SetPropertyDS_msSFU30NisDomain(value string) (err error) {
	return instance.SetProperty("DS_msSFU30NisDomain", (value))
}

// GetDS_msSFU30NisDomain gets the value of DS_msSFU30NisDomain for the instance
func (instance *ads_group) GetPropertyDS_msSFU30NisDomain() (value string, err error) {
	retValue, err := instance.GetProperty("DS_msSFU30NisDomain")
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

// SetDS_msSFU30PosixMember sets the value of DS_msSFU30PosixMember for the instance
func (instance *ads_group) SetPropertyDS_msSFU30PosixMember(value []string) (err error) {
	return instance.SetProperty("DS_msSFU30PosixMember", (value))
}

// GetDS_msSFU30PosixMember gets the value of DS_msSFU30PosixMember for the instance
func (instance *ads_group) GetPropertyDS_msSFU30PosixMember() (value []string, err error) {
	retValue, err := instance.GetProperty("DS_msSFU30PosixMember")
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

// SetDS_nonSecurityMember sets the value of DS_nonSecurityMember for the instance
func (instance *ads_group) SetPropertyDS_nonSecurityMember(value []string) (err error) {
	return instance.SetProperty("DS_nonSecurityMember", (value))
}

// GetDS_nonSecurityMember gets the value of DS_nonSecurityMember for the instance
func (instance *ads_group) GetPropertyDS_nonSecurityMember() (value []string, err error) {
	retValue, err := instance.GetProperty("DS_nonSecurityMember")
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

// SetDS_nTGroupMembers sets the value of DS_nTGroupMembers for the instance
func (instance *ads_group) SetPropertyDS_nTGroupMembers(value []Uint8Array) (err error) {
	return instance.SetProperty("DS_nTGroupMembers", (value))
}

// GetDS_nTGroupMembers gets the value of DS_nTGroupMembers for the instance
func (instance *ads_group) GetPropertyDS_nTGroupMembers() (value []Uint8Array, err error) {
	retValue, err := instance.GetProperty("DS_nTGroupMembers")
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

// SetDS_objectSid sets the value of DS_objectSid for the instance
func (instance *ads_group) SetPropertyDS_objectSid(value Uint8Array) (err error) {
	return instance.SetProperty("DS_objectSid", (value))
}

// GetDS_objectSid gets the value of DS_objectSid for the instance
func (instance *ads_group) GetPropertyDS_objectSid() (value Uint8Array, err error) {
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

// SetDS_operatorCount sets the value of DS_operatorCount for the instance
func (instance *ads_group) SetPropertyDS_operatorCount(value int32) (err error) {
	return instance.SetProperty("DS_operatorCount", (value))
}

// GetDS_operatorCount gets the value of DS_operatorCount for the instance
func (instance *ads_group) GetPropertyDS_operatorCount() (value int32, err error) {
	retValue, err := instance.GetProperty("DS_operatorCount")
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

// SetDS_primaryGroupToken sets the value of DS_primaryGroupToken for the instance
func (instance *ads_group) SetPropertyDS_primaryGroupToken(value int32) (err error) {
	return instance.SetProperty("DS_primaryGroupToken", (value))
}

// GetDS_primaryGroupToken gets the value of DS_primaryGroupToken for the instance
func (instance *ads_group) GetPropertyDS_primaryGroupToken() (value int32, err error) {
	retValue, err := instance.GetProperty("DS_primaryGroupToken")
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

// SetDS_rid sets the value of DS_rid for the instance
func (instance *ads_group) SetPropertyDS_rid(value int32) (err error) {
	return instance.SetProperty("DS_rid", (value))
}

// GetDS_rid gets the value of DS_rid for the instance
func (instance *ads_group) GetPropertyDS_rid() (value int32, err error) {
	retValue, err := instance.GetProperty("DS_rid")
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

// SetDS_sAMAccountName sets the value of DS_sAMAccountName for the instance
func (instance *ads_group) SetPropertyDS_sAMAccountName(value string) (err error) {
	return instance.SetProperty("DS_sAMAccountName", (value))
}

// GetDS_sAMAccountName gets the value of DS_sAMAccountName for the instance
func (instance *ads_group) GetPropertyDS_sAMAccountName() (value string, err error) {
	retValue, err := instance.GetProperty("DS_sAMAccountName")
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

// SetDS_sAMAccountType sets the value of DS_sAMAccountType for the instance
func (instance *ads_group) SetPropertyDS_sAMAccountType(value int32) (err error) {
	return instance.SetProperty("DS_sAMAccountType", (value))
}

// GetDS_sAMAccountType gets the value of DS_sAMAccountType for the instance
func (instance *ads_group) GetPropertyDS_sAMAccountType() (value int32, err error) {
	retValue, err := instance.GetProperty("DS_sAMAccountType")
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

// SetDS_secretary sets the value of DS_secretary for the instance
func (instance *ads_group) SetPropertyDS_secretary(value []string) (err error) {
	return instance.SetProperty("DS_secretary", (value))
}

// GetDS_secretary gets the value of DS_secretary for the instance
func (instance *ads_group) GetPropertyDS_secretary() (value []string, err error) {
	retValue, err := instance.GetProperty("DS_secretary")
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

// SetDS_securityIdentifier sets the value of DS_securityIdentifier for the instance
func (instance *ads_group) SetPropertyDS_securityIdentifier(value Uint8Array) (err error) {
	return instance.SetProperty("DS_securityIdentifier", (value))
}

// GetDS_securityIdentifier gets the value of DS_securityIdentifier for the instance
func (instance *ads_group) GetPropertyDS_securityIdentifier() (value Uint8Array, err error) {
	retValue, err := instance.GetProperty("DS_securityIdentifier")
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

// SetDS_showInAddressBook sets the value of DS_showInAddressBook for the instance
func (instance *ads_group) SetPropertyDS_showInAddressBook(value []string) (err error) {
	return instance.SetProperty("DS_showInAddressBook", (value))
}

// GetDS_showInAddressBook gets the value of DS_showInAddressBook for the instance
func (instance *ads_group) GetPropertyDS_showInAddressBook() (value []string, err error) {
	retValue, err := instance.GetProperty("DS_showInAddressBook")
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

// SetDS_sIDHistory sets the value of DS_sIDHistory for the instance
func (instance *ads_group) SetPropertyDS_sIDHistory(value []Uint8Array) (err error) {
	return instance.SetProperty("DS_sIDHistory", (value))
}

// GetDS_sIDHistory gets the value of DS_sIDHistory for the instance
func (instance *ads_group) GetPropertyDS_sIDHistory() (value []Uint8Array, err error) {
	retValue, err := instance.GetProperty("DS_sIDHistory")
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

// SetDS_supplementalCredentials sets the value of DS_supplementalCredentials for the instance
func (instance *ads_group) SetPropertyDS_supplementalCredentials(value []Uint8Array) (err error) {
	return instance.SetProperty("DS_supplementalCredentials", (value))
}

// GetDS_supplementalCredentials gets the value of DS_supplementalCredentials for the instance
func (instance *ads_group) GetPropertyDS_supplementalCredentials() (value []Uint8Array, err error) {
	retValue, err := instance.GetProperty("DS_supplementalCredentials")
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

// SetDS_telephoneNumber sets the value of DS_telephoneNumber for the instance
func (instance *ads_group) SetPropertyDS_telephoneNumber(value string) (err error) {
	return instance.SetProperty("DS_telephoneNumber", (value))
}

// GetDS_telephoneNumber gets the value of DS_telephoneNumber for the instance
func (instance *ads_group) GetPropertyDS_telephoneNumber() (value string, err error) {
	retValue, err := instance.GetProperty("DS_telephoneNumber")
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

// SetDS_textEncodedORAddress sets the value of DS_textEncodedORAddress for the instance
func (instance *ads_group) SetPropertyDS_textEncodedORAddress(value string) (err error) {
	return instance.SetProperty("DS_textEncodedORAddress", (value))
}

// GetDS_textEncodedORAddress gets the value of DS_textEncodedORAddress for the instance
func (instance *ads_group) GetPropertyDS_textEncodedORAddress() (value string, err error) {
	retValue, err := instance.GetProperty("DS_textEncodedORAddress")
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

// SetDS_tokenGroups sets the value of DS_tokenGroups for the instance
func (instance *ads_group) SetPropertyDS_tokenGroups(value []Uint8Array) (err error) {
	return instance.SetProperty("DS_tokenGroups", (value))
}

// GetDS_tokenGroups gets the value of DS_tokenGroups for the instance
func (instance *ads_group) GetPropertyDS_tokenGroups() (value []Uint8Array, err error) {
	retValue, err := instance.GetProperty("DS_tokenGroups")
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

// SetDS_tokenGroupsGlobalAndUniversal sets the value of DS_tokenGroupsGlobalAndUniversal for the instance
func (instance *ads_group) SetPropertyDS_tokenGroupsGlobalAndUniversal(value []Uint8Array) (err error) {
	return instance.SetProperty("DS_tokenGroupsGlobalAndUniversal", (value))
}

// GetDS_tokenGroupsGlobalAndUniversal gets the value of DS_tokenGroupsGlobalAndUniversal for the instance
func (instance *ads_group) GetPropertyDS_tokenGroupsGlobalAndUniversal() (value []Uint8Array, err error) {
	retValue, err := instance.GetProperty("DS_tokenGroupsGlobalAndUniversal")
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

// SetDS_tokenGroupsNoGCAcceptable sets the value of DS_tokenGroupsNoGCAcceptable for the instance
func (instance *ads_group) SetPropertyDS_tokenGroupsNoGCAcceptable(value []Uint8Array) (err error) {
	return instance.SetProperty("DS_tokenGroupsNoGCAcceptable", (value))
}

// GetDS_tokenGroupsNoGCAcceptable gets the value of DS_tokenGroupsNoGCAcceptable for the instance
func (instance *ads_group) GetPropertyDS_tokenGroupsNoGCAcceptable() (value []Uint8Array, err error) {
	retValue, err := instance.GetProperty("DS_tokenGroupsNoGCAcceptable")
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

// SetDS_unixUserPassword sets the value of DS_unixUserPassword for the instance
func (instance *ads_group) SetPropertyDS_unixUserPassword(value []Uint8Array) (err error) {
	return instance.SetProperty("DS_unixUserPassword", (value))
}

// GetDS_unixUserPassword gets the value of DS_unixUserPassword for the instance
func (instance *ads_group) GetPropertyDS_unixUserPassword() (value []Uint8Array, err error) {
	retValue, err := instance.GetProperty("DS_unixUserPassword")
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

// SetDS_userCert sets the value of DS_userCert for the instance
func (instance *ads_group) SetPropertyDS_userCert(value Uint8Array) (err error) {
	return instance.SetProperty("DS_userCert", (value))
}

// GetDS_userCert gets the value of DS_userCert for the instance
func (instance *ads_group) GetPropertyDS_userCert() (value Uint8Array, err error) {
	retValue, err := instance.GetProperty("DS_userCert")
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

// SetDS_userCertificate sets the value of DS_userCertificate for the instance
func (instance *ads_group) SetPropertyDS_userCertificate(value []Uint8Array) (err error) {
	return instance.SetProperty("DS_userCertificate", (value))
}

// GetDS_userCertificate gets the value of DS_userCertificate for the instance
func (instance *ads_group) GetPropertyDS_userCertificate() (value []Uint8Array, err error) {
	retValue, err := instance.GetProperty("DS_userCertificate")
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

// SetDS_userPassword sets the value of DS_userPassword for the instance
func (instance *ads_group) SetPropertyDS_userPassword(value []Uint8Array) (err error) {
	return instance.SetProperty("DS_userPassword", (value))
}

// GetDS_userPassword gets the value of DS_userPassword for the instance
func (instance *ads_group) GetPropertyDS_userPassword() (value []Uint8Array, err error) {
	retValue, err := instance.GetProperty("DS_userPassword")
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

// SetDS_userSMIMECertificate sets the value of DS_userSMIMECertificate for the instance
func (instance *ads_group) SetPropertyDS_userSMIMECertificate(value []Uint8Array) (err error) {
	return instance.SetProperty("DS_userSMIMECertificate", (value))
}

// GetDS_userSMIMECertificate gets the value of DS_userSMIMECertificate for the instance
func (instance *ads_group) GetPropertyDS_userSMIMECertificate() (value []Uint8Array, err error) {
	retValue, err := instance.GetProperty("DS_userSMIMECertificate")
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
