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

// ads_user struct
type ads_user struct {
	*ads_organizationalperson

	//
	DS_accountExpires int64

	//
	DS_accountNameHistory []string

	//
	DS_aCSPolicyName string

	//
	DS_adminCount int32

	//
	DS_altSecurityIdentities []string

	//
	DS_audio []Uint8Array

	//
	DS_badPasswordTime int64

	//
	DS_badPwdCount int32

	//
	DS_businessCategory []string

	//
	DS_carLicense []string

	//
	DS_codePage int32

	//
	DS_controlAccessRights []Uint8Array

	//
	DS_dBCSPwd Uint8Array

	//
	DS_defaultClassStore []string

	//
	DS_departmentNumber []string

	//
	DS_desktopProfile string

	//
	DS_dynamicLDAPServer string

	//
	DS_employeeNumber string

	//
	DS_employeeType string

	//
	DS_garbageCollPeriod int32

	//
	DS_gecos string

	//
	DS_gidNumber int32

	//
	DS_groupMembershipSAM Uint8Array

	//
	DS_groupPriority []string

	//
	DS_groupsToIgnore []string

	//
	DS_homeDirectory string

	//
	DS_homeDrive string

	//
	DS_info string

	//
	DS_jpegPhoto []Uint8Array

	//
	DS_labeledURI []string

	//
	DS_lastLogoff int64

	//
	DS_lastLogon int64

	//
	DS_lastLogonTimestamp int64

	//
	DS_legacyExchangeDN string

	//
	DS_lmPwdHistory []Uint8Array

	//
	DS_localeID []int32

	//
	DS_lockoutTime int64

	//
	DS_loginShell string

	//
	DS_logonCount int32

	//
	DS_logonHours Uint8Array

	//
	DS_logonWorkstation Uint8Array

	//
	DS_maxStorage int64

	//
	DS_mS_DS_CreatorSID Uint8Array

	//
	DS_msCOM_UserPartitionSetLink string

	//
	DS_msDRM_IdentityCertificate []Uint8Array

	//
	DS_msDS_AssignedAuthNPolicy string

	//
	DS_msDS_AssignedAuthNPolicySilo string

	//
	DS_msDS_AuthenticatedAtDC []string

	//
	DS_msDS_AuthNPolicySiloMembersBL []string

	//
	DS_msDS_Cached_Membership Uint8Array

	//
	DS_msDS_Cached_Membership_Time_Stamp int64

	//
	DS_msDS_cloudExtensionAttribute1 string

	//
	DS_msDS_cloudExtensionAttribute10 string

	//
	DS_msDS_cloudExtensionAttribute11 string

	//
	DS_msDS_cloudExtensionAttribute12 string

	//
	DS_msDS_cloudExtensionAttribute13 string

	//
	DS_msDS_cloudExtensionAttribute14 string

	//
	DS_msDS_cloudExtensionAttribute15 string

	//
	DS_msDS_cloudExtensionAttribute16 string

	//
	DS_msDS_cloudExtensionAttribute17 string

	//
	DS_msDS_cloudExtensionAttribute18 string

	//
	DS_msDS_cloudExtensionAttribute19 string

	//
	DS_msDS_cloudExtensionAttribute2 string

	//
	DS_msDS_cloudExtensionAttribute20 string

	//
	DS_msDS_cloudExtensionAttribute3 string

	//
	DS_msDS_cloudExtensionAttribute4 string

	//
	DS_msDS_cloudExtensionAttribute5 string

	//
	DS_msDS_cloudExtensionAttribute6 string

	//
	DS_msDS_cloudExtensionAttribute7 string

	//
	DS_msDS_cloudExtensionAttribute8 string

	//
	DS_msDS_cloudExtensionAttribute9 string

	//
	DS_msDS_ExternalDirectoryObjectId string

	//
	DS_msDS_FailedInteractiveLogonCount int32

	//
	DS_msDS_FailedInteractiveLogonCountAtLastSuccessfulLogon int32

	//
	DS_msDS_GeoCoordinatesAltitude int64

	//
	DS_msDS_GeoCoordinatesLatitude int64

	//
	DS_msDS_GeoCoordinatesLongitude int64

	//
	DS_msDS_KeyCredentialLink []DN_With_Binary

	//
	DS_msDS_KeyPrincipalBL []string

	//
	DS_msDS_KeyVersionNumber int32

	//
	DS_msDS_LastFailedInteractiveLogonTime int64

	//
	DS_msDS_LastSuccessfulInteractiveLogonTime int64

	//
	DS_msDS_preferredDataLocation string

	//
	DS_msDS_PrimaryComputer []string

	//
	DS_msDS_ResultantPSO string

	//
	DS_msDS_SecondaryKrbTgtNumber int32

	//
	DS_msDS_Site_Affinity []Uint8Array

	//
	DS_msDS_SourceObjectDN string

	//
	DS_msDS_SupportedEncryptionTypes int32

	//
	DS_msDS_SyncServerUrl []string

	//
	DS_msds_tokenGroupNames []string

	//
	DS_msds_tokenGroupNamesGlobalAndUniversal []string

	//
	DS_msds_tokenGroupNamesNoGCAcceptable []string

	//
	DS_msDS_User_Account_Control_Computed int32

	//
	DS_msDS_UserPasswordExpiryTimeComputed int64

	//
	DS_msExchAssistantName string

	//
	DS_msExchLabeledURI []string

	//
	DS_msIIS_FTPDir string

	//
	DS_msIIS_FTPRoot string

	//
	DS_mSMQDigests []Uint8Array

	//
	DS_mSMQDigestsMig []Uint8Array

	//
	DS_mSMQSignCertificates Uint8Array

	//
	DS_mSMQSignCertificatesMig Uint8Array

	//
	DS_msNPAllowDialin bool

	//
	DS_msNPCallingStationID []string

	//
	DS_msNPSavedCallingStationID []string

	//
	DS_msPKI_CredentialRoamingTokens []DN_With_Binary

	//
	DS_msPKIAccountCredentials []DN_With_Binary

	//
	DS_msPKIDPAPIMasterKeys []DN_With_Binary

	//
	DS_msPKIRoamingTimeStamp Uint8Array

	//
	DS_msRADIUS_FramedInterfaceId string

	//
	DS_msRADIUS_FramedIpv6Prefix string

	//
	DS_msRADIUS_FramedIpv6Route []string

	//
	DS_msRADIUS_SavedFramedInterfaceId string

	//
	DS_msRADIUS_SavedFramedIpv6Prefix string

	//
	DS_msRADIUS_SavedFramedIpv6Route []string

	//
	DS_msRADIUSCallbackNumber string

	//
	DS_msRADIUSFramedIPAddress int32

	//
	DS_msRADIUSFramedRoute []string

	//
	DS_msRADIUSServiceType int32

	//
	DS_msRASSavedCallbackNumber string

	//
	DS_msRASSavedFramedIPAddress int32

	//
	DS_msRASSavedFramedRoute []string

	//
	DS_msSFU30Name string

	//
	DS_msSFU30NisDomain string

	//
	DS_msTSAllowLogon bool

	//
	DS_msTSBrokenConnectionAction bool

	//
	DS_msTSConnectClientDrives bool

	//
	DS_msTSConnectPrinterDrives bool

	//
	DS_msTSDefaultToMainPrinter bool

	//
	DS_msTSExpireDate string

	//
	DS_msTSExpireDate2 string

	//
	DS_msTSExpireDate3 string

	//
	DS_msTSExpireDate4 string

	//
	DS_msTSHomeDirectory string

	//
	DS_msTSHomeDrive string

	//
	DS_msTSInitialProgram string

	//
	DS_msTSLicenseVersion string

	//
	DS_msTSLicenseVersion2 string

	//
	DS_msTSLicenseVersion3 string

	//
	DS_msTSLicenseVersion4 string

	//
	DS_msTSLSProperty01 []string

	//
	DS_msTSLSProperty02 []string

	//
	DS_msTSManagingLS string

	//
	DS_msTSManagingLS2 string

	//
	DS_msTSManagingLS3 string

	//
	DS_msTSManagingLS4 string

	//
	DS_msTSMaxConnectionTime int32

	//
	DS_msTSMaxDisconnectionTime int32

	//
	DS_msTSMaxIdleTime int32

	//
	DS_msTSPrimaryDesktop string

	//
	DS_msTSProfilePath string

	//
	DS_msTSProperty01 []string

	//
	DS_msTSProperty02 []string

	//
	DS_msTSReconnectionAction bool

	//
	DS_msTSRemoteControl int32

	//
	DS_msTSSecondaryDesktops []string

	//
	DS_msTSWorkDirectory string

	//
	DS_networkAddress []string

	//
	DS_ntPwdHistory []Uint8Array

	//
	DS_objectSid Uint8Array

	//
	DS_operatorCount int32

	//
	DS_otherLoginWorkstations []string

	//
	DS_photo []Uint8Array

	//
	DS_preferredLanguage string

	//
	DS_preferredOU string

	//
	DS_primaryGroupID int32

	//
	DS_profilePath string

	//
	DS_pwdLastSet int64

	//
	DS_rid int32

	//
	DS_roomNumber []string

	//
	DS_sAMAccountName string

	//
	DS_sAMAccountType int32

	//
	DS_scriptPath string

	//
	DS_secretary []string

	//
	DS_securityIdentifier Uint8Array

	//
	DS_servicePrincipalName []string

	//
	DS_shadowExpire int32

	//
	DS_shadowFlag int32

	//
	DS_shadowInactive int32

	//
	DS_shadowLastChange int32

	//
	DS_shadowMax int32

	//
	DS_shadowMin int32

	//
	DS_shadowWarning int32

	//
	DS_showInAddressBook []string

	//
	DS_sIDHistory []Uint8Array

	//
	DS_supplementalCredentials []Uint8Array

	//
	DS_terminalServer Uint8Array

	//
	DS_textEncodedORAddress string

	//
	DS_tokenGroups []Uint8Array

	//
	DS_tokenGroupsGlobalAndUniversal []Uint8Array

	//
	DS_tokenGroupsNoGCAcceptable []Uint8Array

	//
	DS_uid []string

	//
	DS_uidNumber int32

	//
	DS_unicodePwd Uint8Array

	//
	DS_unixHomeDirectory string

	//
	DS_unixUserPassword []Uint8Array

	//
	DS_userAccountControl int32

	//
	DS_userCert Uint8Array

	//
	DS_userCertificate []Uint8Array

	//
	DS_userParameters string

	//
	DS_userPKCS12 []Uint8Array

	//
	DS_userPrincipalName string

	//
	DS_userSharedFolder string

	//
	DS_userSharedFolderOther []string

	//
	DS_userSMIMECertificate []Uint8Array

	//
	DS_userWorkstations string

	//
	DS_x500uniqueIdentifier []Uint8Array
}

func Newads_userEx1(instance *cim.WmiInstance) (newInstance *ads_user, err error) {
	tmp, err := Newads_organizationalpersonEx1(instance)

	if err != nil {
		return
	}
	newInstance = &ads_user{
		ads_organizationalperson: tmp,
	}
	return
}

func Newads_userEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *ads_user, err error) {
	tmp, err := Newads_organizationalpersonEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &ads_user{
		ads_organizationalperson: tmp,
	}
	return
}

// SetDS_accountExpires sets the value of DS_accountExpires for the instance
func (instance *ads_user) SetPropertyDS_accountExpires(value int64) (err error) {
	return instance.SetProperty("DS_accountExpires", (value))
}

// GetDS_accountExpires gets the value of DS_accountExpires for the instance
func (instance *ads_user) GetPropertyDS_accountExpires() (value int64, err error) {
	retValue, err := instance.GetProperty("DS_accountExpires")
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

// SetDS_accountNameHistory sets the value of DS_accountNameHistory for the instance
func (instance *ads_user) SetPropertyDS_accountNameHistory(value []string) (err error) {
	return instance.SetProperty("DS_accountNameHistory", (value))
}

// GetDS_accountNameHistory gets the value of DS_accountNameHistory for the instance
func (instance *ads_user) GetPropertyDS_accountNameHistory() (value []string, err error) {
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

// SetDS_aCSPolicyName sets the value of DS_aCSPolicyName for the instance
func (instance *ads_user) SetPropertyDS_aCSPolicyName(value string) (err error) {
	return instance.SetProperty("DS_aCSPolicyName", (value))
}

// GetDS_aCSPolicyName gets the value of DS_aCSPolicyName for the instance
func (instance *ads_user) GetPropertyDS_aCSPolicyName() (value string, err error) {
	retValue, err := instance.GetProperty("DS_aCSPolicyName")
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

// SetDS_adminCount sets the value of DS_adminCount for the instance
func (instance *ads_user) SetPropertyDS_adminCount(value int32) (err error) {
	return instance.SetProperty("DS_adminCount", (value))
}

// GetDS_adminCount gets the value of DS_adminCount for the instance
func (instance *ads_user) GetPropertyDS_adminCount() (value int32, err error) {
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
func (instance *ads_user) SetPropertyDS_altSecurityIdentities(value []string) (err error) {
	return instance.SetProperty("DS_altSecurityIdentities", (value))
}

// GetDS_altSecurityIdentities gets the value of DS_altSecurityIdentities for the instance
func (instance *ads_user) GetPropertyDS_altSecurityIdentities() (value []string, err error) {
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

// SetDS_audio sets the value of DS_audio for the instance
func (instance *ads_user) SetPropertyDS_audio(value []Uint8Array) (err error) {
	return instance.SetProperty("DS_audio", (value))
}

// GetDS_audio gets the value of DS_audio for the instance
func (instance *ads_user) GetPropertyDS_audio() (value []Uint8Array, err error) {
	retValue, err := instance.GetProperty("DS_audio")
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

// SetDS_badPasswordTime sets the value of DS_badPasswordTime for the instance
func (instance *ads_user) SetPropertyDS_badPasswordTime(value int64) (err error) {
	return instance.SetProperty("DS_badPasswordTime", (value))
}

// GetDS_badPasswordTime gets the value of DS_badPasswordTime for the instance
func (instance *ads_user) GetPropertyDS_badPasswordTime() (value int64, err error) {
	retValue, err := instance.GetProperty("DS_badPasswordTime")
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

// SetDS_badPwdCount sets the value of DS_badPwdCount for the instance
func (instance *ads_user) SetPropertyDS_badPwdCount(value int32) (err error) {
	return instance.SetProperty("DS_badPwdCount", (value))
}

// GetDS_badPwdCount gets the value of DS_badPwdCount for the instance
func (instance *ads_user) GetPropertyDS_badPwdCount() (value int32, err error) {
	retValue, err := instance.GetProperty("DS_badPwdCount")
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

// SetDS_businessCategory sets the value of DS_businessCategory for the instance
func (instance *ads_user) SetPropertyDS_businessCategory(value []string) (err error) {
	return instance.SetProperty("DS_businessCategory", (value))
}

// GetDS_businessCategory gets the value of DS_businessCategory for the instance
func (instance *ads_user) GetPropertyDS_businessCategory() (value []string, err error) {
	retValue, err := instance.GetProperty("DS_businessCategory")
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

// SetDS_carLicense sets the value of DS_carLicense for the instance
func (instance *ads_user) SetPropertyDS_carLicense(value []string) (err error) {
	return instance.SetProperty("DS_carLicense", (value))
}

// GetDS_carLicense gets the value of DS_carLicense for the instance
func (instance *ads_user) GetPropertyDS_carLicense() (value []string, err error) {
	retValue, err := instance.GetProperty("DS_carLicense")
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

// SetDS_codePage sets the value of DS_codePage for the instance
func (instance *ads_user) SetPropertyDS_codePage(value int32) (err error) {
	return instance.SetProperty("DS_codePage", (value))
}

// GetDS_codePage gets the value of DS_codePage for the instance
func (instance *ads_user) GetPropertyDS_codePage() (value int32, err error) {
	retValue, err := instance.GetProperty("DS_codePage")
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

// SetDS_controlAccessRights sets the value of DS_controlAccessRights for the instance
func (instance *ads_user) SetPropertyDS_controlAccessRights(value []Uint8Array) (err error) {
	return instance.SetProperty("DS_controlAccessRights", (value))
}

// GetDS_controlAccessRights gets the value of DS_controlAccessRights for the instance
func (instance *ads_user) GetPropertyDS_controlAccessRights() (value []Uint8Array, err error) {
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

// SetDS_dBCSPwd sets the value of DS_dBCSPwd for the instance
func (instance *ads_user) SetPropertyDS_dBCSPwd(value Uint8Array) (err error) {
	return instance.SetProperty("DS_dBCSPwd", (value))
}

// GetDS_dBCSPwd gets the value of DS_dBCSPwd for the instance
func (instance *ads_user) GetPropertyDS_dBCSPwd() (value Uint8Array, err error) {
	retValue, err := instance.GetProperty("DS_dBCSPwd")
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

// SetDS_defaultClassStore sets the value of DS_defaultClassStore for the instance
func (instance *ads_user) SetPropertyDS_defaultClassStore(value []string) (err error) {
	return instance.SetProperty("DS_defaultClassStore", (value))
}

// GetDS_defaultClassStore gets the value of DS_defaultClassStore for the instance
func (instance *ads_user) GetPropertyDS_defaultClassStore() (value []string, err error) {
	retValue, err := instance.GetProperty("DS_defaultClassStore")
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

// SetDS_departmentNumber sets the value of DS_departmentNumber for the instance
func (instance *ads_user) SetPropertyDS_departmentNumber(value []string) (err error) {
	return instance.SetProperty("DS_departmentNumber", (value))
}

// GetDS_departmentNumber gets the value of DS_departmentNumber for the instance
func (instance *ads_user) GetPropertyDS_departmentNumber() (value []string, err error) {
	retValue, err := instance.GetProperty("DS_departmentNumber")
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

// SetDS_desktopProfile sets the value of DS_desktopProfile for the instance
func (instance *ads_user) SetPropertyDS_desktopProfile(value string) (err error) {
	return instance.SetProperty("DS_desktopProfile", (value))
}

// GetDS_desktopProfile gets the value of DS_desktopProfile for the instance
func (instance *ads_user) GetPropertyDS_desktopProfile() (value string, err error) {
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

// SetDS_dynamicLDAPServer sets the value of DS_dynamicLDAPServer for the instance
func (instance *ads_user) SetPropertyDS_dynamicLDAPServer(value string) (err error) {
	return instance.SetProperty("DS_dynamicLDAPServer", (value))
}

// GetDS_dynamicLDAPServer gets the value of DS_dynamicLDAPServer for the instance
func (instance *ads_user) GetPropertyDS_dynamicLDAPServer() (value string, err error) {
	retValue, err := instance.GetProperty("DS_dynamicLDAPServer")
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

// SetDS_employeeNumber sets the value of DS_employeeNumber for the instance
func (instance *ads_user) SetPropertyDS_employeeNumber(value string) (err error) {
	return instance.SetProperty("DS_employeeNumber", (value))
}

// GetDS_employeeNumber gets the value of DS_employeeNumber for the instance
func (instance *ads_user) GetPropertyDS_employeeNumber() (value string, err error) {
	retValue, err := instance.GetProperty("DS_employeeNumber")
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

// SetDS_employeeType sets the value of DS_employeeType for the instance
func (instance *ads_user) SetPropertyDS_employeeType(value string) (err error) {
	return instance.SetProperty("DS_employeeType", (value))
}

// GetDS_employeeType gets the value of DS_employeeType for the instance
func (instance *ads_user) GetPropertyDS_employeeType() (value string, err error) {
	retValue, err := instance.GetProperty("DS_employeeType")
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
func (instance *ads_user) SetPropertyDS_garbageCollPeriod(value int32) (err error) {
	return instance.SetProperty("DS_garbageCollPeriod", (value))
}

// GetDS_garbageCollPeriod gets the value of DS_garbageCollPeriod for the instance
func (instance *ads_user) GetPropertyDS_garbageCollPeriod() (value int32, err error) {
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

// SetDS_gecos sets the value of DS_gecos for the instance
func (instance *ads_user) SetPropertyDS_gecos(value string) (err error) {
	return instance.SetProperty("DS_gecos", (value))
}

// GetDS_gecos gets the value of DS_gecos for the instance
func (instance *ads_user) GetPropertyDS_gecos() (value string, err error) {
	retValue, err := instance.GetProperty("DS_gecos")
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

// SetDS_gidNumber sets the value of DS_gidNumber for the instance
func (instance *ads_user) SetPropertyDS_gidNumber(value int32) (err error) {
	return instance.SetProperty("DS_gidNumber", (value))
}

// GetDS_gidNumber gets the value of DS_gidNumber for the instance
func (instance *ads_user) GetPropertyDS_gidNumber() (value int32, err error) {
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

// SetDS_groupMembershipSAM sets the value of DS_groupMembershipSAM for the instance
func (instance *ads_user) SetPropertyDS_groupMembershipSAM(value Uint8Array) (err error) {
	return instance.SetProperty("DS_groupMembershipSAM", (value))
}

// GetDS_groupMembershipSAM gets the value of DS_groupMembershipSAM for the instance
func (instance *ads_user) GetPropertyDS_groupMembershipSAM() (value Uint8Array, err error) {
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

// SetDS_groupPriority sets the value of DS_groupPriority for the instance
func (instance *ads_user) SetPropertyDS_groupPriority(value []string) (err error) {
	return instance.SetProperty("DS_groupPriority", (value))
}

// GetDS_groupPriority gets the value of DS_groupPriority for the instance
func (instance *ads_user) GetPropertyDS_groupPriority() (value []string, err error) {
	retValue, err := instance.GetProperty("DS_groupPriority")
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

// SetDS_groupsToIgnore sets the value of DS_groupsToIgnore for the instance
func (instance *ads_user) SetPropertyDS_groupsToIgnore(value []string) (err error) {
	return instance.SetProperty("DS_groupsToIgnore", (value))
}

// GetDS_groupsToIgnore gets the value of DS_groupsToIgnore for the instance
func (instance *ads_user) GetPropertyDS_groupsToIgnore() (value []string, err error) {
	retValue, err := instance.GetProperty("DS_groupsToIgnore")
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

// SetDS_homeDirectory sets the value of DS_homeDirectory for the instance
func (instance *ads_user) SetPropertyDS_homeDirectory(value string) (err error) {
	return instance.SetProperty("DS_homeDirectory", (value))
}

// GetDS_homeDirectory gets the value of DS_homeDirectory for the instance
func (instance *ads_user) GetPropertyDS_homeDirectory() (value string, err error) {
	retValue, err := instance.GetProperty("DS_homeDirectory")
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

// SetDS_homeDrive sets the value of DS_homeDrive for the instance
func (instance *ads_user) SetPropertyDS_homeDrive(value string) (err error) {
	return instance.SetProperty("DS_homeDrive", (value))
}

// GetDS_homeDrive gets the value of DS_homeDrive for the instance
func (instance *ads_user) GetPropertyDS_homeDrive() (value string, err error) {
	retValue, err := instance.GetProperty("DS_homeDrive")
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

// SetDS_info sets the value of DS_info for the instance
func (instance *ads_user) SetPropertyDS_info(value string) (err error) {
	return instance.SetProperty("DS_info", (value))
}

// GetDS_info gets the value of DS_info for the instance
func (instance *ads_user) GetPropertyDS_info() (value string, err error) {
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

// SetDS_jpegPhoto sets the value of DS_jpegPhoto for the instance
func (instance *ads_user) SetPropertyDS_jpegPhoto(value []Uint8Array) (err error) {
	return instance.SetProperty("DS_jpegPhoto", (value))
}

// GetDS_jpegPhoto gets the value of DS_jpegPhoto for the instance
func (instance *ads_user) GetPropertyDS_jpegPhoto() (value []Uint8Array, err error) {
	retValue, err := instance.GetProperty("DS_jpegPhoto")
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

// SetDS_labeledURI sets the value of DS_labeledURI for the instance
func (instance *ads_user) SetPropertyDS_labeledURI(value []string) (err error) {
	return instance.SetProperty("DS_labeledURI", (value))
}

// GetDS_labeledURI gets the value of DS_labeledURI for the instance
func (instance *ads_user) GetPropertyDS_labeledURI() (value []string, err error) {
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

// SetDS_lastLogoff sets the value of DS_lastLogoff for the instance
func (instance *ads_user) SetPropertyDS_lastLogoff(value int64) (err error) {
	return instance.SetProperty("DS_lastLogoff", (value))
}

// GetDS_lastLogoff gets the value of DS_lastLogoff for the instance
func (instance *ads_user) GetPropertyDS_lastLogoff() (value int64, err error) {
	retValue, err := instance.GetProperty("DS_lastLogoff")
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

// SetDS_lastLogon sets the value of DS_lastLogon for the instance
func (instance *ads_user) SetPropertyDS_lastLogon(value int64) (err error) {
	return instance.SetProperty("DS_lastLogon", (value))
}

// GetDS_lastLogon gets the value of DS_lastLogon for the instance
func (instance *ads_user) GetPropertyDS_lastLogon() (value int64, err error) {
	retValue, err := instance.GetProperty("DS_lastLogon")
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

// SetDS_lastLogonTimestamp sets the value of DS_lastLogonTimestamp for the instance
func (instance *ads_user) SetPropertyDS_lastLogonTimestamp(value int64) (err error) {
	return instance.SetProperty("DS_lastLogonTimestamp", (value))
}

// GetDS_lastLogonTimestamp gets the value of DS_lastLogonTimestamp for the instance
func (instance *ads_user) GetPropertyDS_lastLogonTimestamp() (value int64, err error) {
	retValue, err := instance.GetProperty("DS_lastLogonTimestamp")
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

// SetDS_legacyExchangeDN sets the value of DS_legacyExchangeDN for the instance
func (instance *ads_user) SetPropertyDS_legacyExchangeDN(value string) (err error) {
	return instance.SetProperty("DS_legacyExchangeDN", (value))
}

// GetDS_legacyExchangeDN gets the value of DS_legacyExchangeDN for the instance
func (instance *ads_user) GetPropertyDS_legacyExchangeDN() (value string, err error) {
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

// SetDS_lmPwdHistory sets the value of DS_lmPwdHistory for the instance
func (instance *ads_user) SetPropertyDS_lmPwdHistory(value []Uint8Array) (err error) {
	return instance.SetProperty("DS_lmPwdHistory", (value))
}

// GetDS_lmPwdHistory gets the value of DS_lmPwdHistory for the instance
func (instance *ads_user) GetPropertyDS_lmPwdHistory() (value []Uint8Array, err error) {
	retValue, err := instance.GetProperty("DS_lmPwdHistory")
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

// SetDS_localeID sets the value of DS_localeID for the instance
func (instance *ads_user) SetPropertyDS_localeID(value []int32) (err error) {
	return instance.SetProperty("DS_localeID", (value))
}

// GetDS_localeID gets the value of DS_localeID for the instance
func (instance *ads_user) GetPropertyDS_localeID() (value []int32, err error) {
	retValue, err := instance.GetProperty("DS_localeID")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	for _, interfaceValue := range retValue.([]interface{}) {
		valuetmp, ok := interfaceValue.(int32)
		if !ok {
			err = errors.Wrapf(errors.InvalidType, " int32 is Invalid. Expected %s", reflect.TypeOf(interfaceValue))
			return
		}
		value = append(value, int32(valuetmp))
	}

	return
}

// SetDS_lockoutTime sets the value of DS_lockoutTime for the instance
func (instance *ads_user) SetPropertyDS_lockoutTime(value int64) (err error) {
	return instance.SetProperty("DS_lockoutTime", (value))
}

// GetDS_lockoutTime gets the value of DS_lockoutTime for the instance
func (instance *ads_user) GetPropertyDS_lockoutTime() (value int64, err error) {
	retValue, err := instance.GetProperty("DS_lockoutTime")
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

// SetDS_loginShell sets the value of DS_loginShell for the instance
func (instance *ads_user) SetPropertyDS_loginShell(value string) (err error) {
	return instance.SetProperty("DS_loginShell", (value))
}

// GetDS_loginShell gets the value of DS_loginShell for the instance
func (instance *ads_user) GetPropertyDS_loginShell() (value string, err error) {
	retValue, err := instance.GetProperty("DS_loginShell")
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

// SetDS_logonCount sets the value of DS_logonCount for the instance
func (instance *ads_user) SetPropertyDS_logonCount(value int32) (err error) {
	return instance.SetProperty("DS_logonCount", (value))
}

// GetDS_logonCount gets the value of DS_logonCount for the instance
func (instance *ads_user) GetPropertyDS_logonCount() (value int32, err error) {
	retValue, err := instance.GetProperty("DS_logonCount")
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

// SetDS_logonHours sets the value of DS_logonHours for the instance
func (instance *ads_user) SetPropertyDS_logonHours(value Uint8Array) (err error) {
	return instance.SetProperty("DS_logonHours", (value))
}

// GetDS_logonHours gets the value of DS_logonHours for the instance
func (instance *ads_user) GetPropertyDS_logonHours() (value Uint8Array, err error) {
	retValue, err := instance.GetProperty("DS_logonHours")
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

// SetDS_logonWorkstation sets the value of DS_logonWorkstation for the instance
func (instance *ads_user) SetPropertyDS_logonWorkstation(value Uint8Array) (err error) {
	return instance.SetProperty("DS_logonWorkstation", (value))
}

// GetDS_logonWorkstation gets the value of DS_logonWorkstation for the instance
func (instance *ads_user) GetPropertyDS_logonWorkstation() (value Uint8Array, err error) {
	retValue, err := instance.GetProperty("DS_logonWorkstation")
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

// SetDS_maxStorage sets the value of DS_maxStorage for the instance
func (instance *ads_user) SetPropertyDS_maxStorage(value int64) (err error) {
	return instance.SetProperty("DS_maxStorage", (value))
}

// GetDS_maxStorage gets the value of DS_maxStorage for the instance
func (instance *ads_user) GetPropertyDS_maxStorage() (value int64, err error) {
	retValue, err := instance.GetProperty("DS_maxStorage")
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

// SetDS_mS_DS_CreatorSID sets the value of DS_mS_DS_CreatorSID for the instance
func (instance *ads_user) SetPropertyDS_mS_DS_CreatorSID(value Uint8Array) (err error) {
	return instance.SetProperty("DS_mS_DS_CreatorSID", (value))
}

// GetDS_mS_DS_CreatorSID gets the value of DS_mS_DS_CreatorSID for the instance
func (instance *ads_user) GetPropertyDS_mS_DS_CreatorSID() (value Uint8Array, err error) {
	retValue, err := instance.GetProperty("DS_mS_DS_CreatorSID")
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

// SetDS_msCOM_UserPartitionSetLink sets the value of DS_msCOM_UserPartitionSetLink for the instance
func (instance *ads_user) SetPropertyDS_msCOM_UserPartitionSetLink(value string) (err error) {
	return instance.SetProperty("DS_msCOM_UserPartitionSetLink", (value))
}

// GetDS_msCOM_UserPartitionSetLink gets the value of DS_msCOM_UserPartitionSetLink for the instance
func (instance *ads_user) GetPropertyDS_msCOM_UserPartitionSetLink() (value string, err error) {
	retValue, err := instance.GetProperty("DS_msCOM_UserPartitionSetLink")
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

// SetDS_msDRM_IdentityCertificate sets the value of DS_msDRM_IdentityCertificate for the instance
func (instance *ads_user) SetPropertyDS_msDRM_IdentityCertificate(value []Uint8Array) (err error) {
	return instance.SetProperty("DS_msDRM_IdentityCertificate", (value))
}

// GetDS_msDRM_IdentityCertificate gets the value of DS_msDRM_IdentityCertificate for the instance
func (instance *ads_user) GetPropertyDS_msDRM_IdentityCertificate() (value []Uint8Array, err error) {
	retValue, err := instance.GetProperty("DS_msDRM_IdentityCertificate")
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

// SetDS_msDS_AssignedAuthNPolicy sets the value of DS_msDS_AssignedAuthNPolicy for the instance
func (instance *ads_user) SetPropertyDS_msDS_AssignedAuthNPolicy(value string) (err error) {
	return instance.SetProperty("DS_msDS_AssignedAuthNPolicy", (value))
}

// GetDS_msDS_AssignedAuthNPolicy gets the value of DS_msDS_AssignedAuthNPolicy for the instance
func (instance *ads_user) GetPropertyDS_msDS_AssignedAuthNPolicy() (value string, err error) {
	retValue, err := instance.GetProperty("DS_msDS_AssignedAuthNPolicy")
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

// SetDS_msDS_AssignedAuthNPolicySilo sets the value of DS_msDS_AssignedAuthNPolicySilo for the instance
func (instance *ads_user) SetPropertyDS_msDS_AssignedAuthNPolicySilo(value string) (err error) {
	return instance.SetProperty("DS_msDS_AssignedAuthNPolicySilo", (value))
}

// GetDS_msDS_AssignedAuthNPolicySilo gets the value of DS_msDS_AssignedAuthNPolicySilo for the instance
func (instance *ads_user) GetPropertyDS_msDS_AssignedAuthNPolicySilo() (value string, err error) {
	retValue, err := instance.GetProperty("DS_msDS_AssignedAuthNPolicySilo")
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

// SetDS_msDS_AuthenticatedAtDC sets the value of DS_msDS_AuthenticatedAtDC for the instance
func (instance *ads_user) SetPropertyDS_msDS_AuthenticatedAtDC(value []string) (err error) {
	return instance.SetProperty("DS_msDS_AuthenticatedAtDC", (value))
}

// GetDS_msDS_AuthenticatedAtDC gets the value of DS_msDS_AuthenticatedAtDC for the instance
func (instance *ads_user) GetPropertyDS_msDS_AuthenticatedAtDC() (value []string, err error) {
	retValue, err := instance.GetProperty("DS_msDS_AuthenticatedAtDC")
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

// SetDS_msDS_AuthNPolicySiloMembersBL sets the value of DS_msDS_AuthNPolicySiloMembersBL for the instance
func (instance *ads_user) SetPropertyDS_msDS_AuthNPolicySiloMembersBL(value []string) (err error) {
	return instance.SetProperty("DS_msDS_AuthNPolicySiloMembersBL", (value))
}

// GetDS_msDS_AuthNPolicySiloMembersBL gets the value of DS_msDS_AuthNPolicySiloMembersBL for the instance
func (instance *ads_user) GetPropertyDS_msDS_AuthNPolicySiloMembersBL() (value []string, err error) {
	retValue, err := instance.GetProperty("DS_msDS_AuthNPolicySiloMembersBL")
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

// SetDS_msDS_Cached_Membership sets the value of DS_msDS_Cached_Membership for the instance
func (instance *ads_user) SetPropertyDS_msDS_Cached_Membership(value Uint8Array) (err error) {
	return instance.SetProperty("DS_msDS_Cached_Membership", (value))
}

// GetDS_msDS_Cached_Membership gets the value of DS_msDS_Cached_Membership for the instance
func (instance *ads_user) GetPropertyDS_msDS_Cached_Membership() (value Uint8Array, err error) {
	retValue, err := instance.GetProperty("DS_msDS_Cached_Membership")
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

// SetDS_msDS_Cached_Membership_Time_Stamp sets the value of DS_msDS_Cached_Membership_Time_Stamp for the instance
func (instance *ads_user) SetPropertyDS_msDS_Cached_Membership_Time_Stamp(value int64) (err error) {
	return instance.SetProperty("DS_msDS_Cached_Membership_Time_Stamp", (value))
}

// GetDS_msDS_Cached_Membership_Time_Stamp gets the value of DS_msDS_Cached_Membership_Time_Stamp for the instance
func (instance *ads_user) GetPropertyDS_msDS_Cached_Membership_Time_Stamp() (value int64, err error) {
	retValue, err := instance.GetProperty("DS_msDS_Cached_Membership_Time_Stamp")
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

// SetDS_msDS_cloudExtensionAttribute1 sets the value of DS_msDS_cloudExtensionAttribute1 for the instance
func (instance *ads_user) SetPropertyDS_msDS_cloudExtensionAttribute1(value string) (err error) {
	return instance.SetProperty("DS_msDS_cloudExtensionAttribute1", (value))
}

// GetDS_msDS_cloudExtensionAttribute1 gets the value of DS_msDS_cloudExtensionAttribute1 for the instance
func (instance *ads_user) GetPropertyDS_msDS_cloudExtensionAttribute1() (value string, err error) {
	retValue, err := instance.GetProperty("DS_msDS_cloudExtensionAttribute1")
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

// SetDS_msDS_cloudExtensionAttribute10 sets the value of DS_msDS_cloudExtensionAttribute10 for the instance
func (instance *ads_user) SetPropertyDS_msDS_cloudExtensionAttribute10(value string) (err error) {
	return instance.SetProperty("DS_msDS_cloudExtensionAttribute10", (value))
}

// GetDS_msDS_cloudExtensionAttribute10 gets the value of DS_msDS_cloudExtensionAttribute10 for the instance
func (instance *ads_user) GetPropertyDS_msDS_cloudExtensionAttribute10() (value string, err error) {
	retValue, err := instance.GetProperty("DS_msDS_cloudExtensionAttribute10")
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

// SetDS_msDS_cloudExtensionAttribute11 sets the value of DS_msDS_cloudExtensionAttribute11 for the instance
func (instance *ads_user) SetPropertyDS_msDS_cloudExtensionAttribute11(value string) (err error) {
	return instance.SetProperty("DS_msDS_cloudExtensionAttribute11", (value))
}

// GetDS_msDS_cloudExtensionAttribute11 gets the value of DS_msDS_cloudExtensionAttribute11 for the instance
func (instance *ads_user) GetPropertyDS_msDS_cloudExtensionAttribute11() (value string, err error) {
	retValue, err := instance.GetProperty("DS_msDS_cloudExtensionAttribute11")
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

// SetDS_msDS_cloudExtensionAttribute12 sets the value of DS_msDS_cloudExtensionAttribute12 for the instance
func (instance *ads_user) SetPropertyDS_msDS_cloudExtensionAttribute12(value string) (err error) {
	return instance.SetProperty("DS_msDS_cloudExtensionAttribute12", (value))
}

// GetDS_msDS_cloudExtensionAttribute12 gets the value of DS_msDS_cloudExtensionAttribute12 for the instance
func (instance *ads_user) GetPropertyDS_msDS_cloudExtensionAttribute12() (value string, err error) {
	retValue, err := instance.GetProperty("DS_msDS_cloudExtensionAttribute12")
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

// SetDS_msDS_cloudExtensionAttribute13 sets the value of DS_msDS_cloudExtensionAttribute13 for the instance
func (instance *ads_user) SetPropertyDS_msDS_cloudExtensionAttribute13(value string) (err error) {
	return instance.SetProperty("DS_msDS_cloudExtensionAttribute13", (value))
}

// GetDS_msDS_cloudExtensionAttribute13 gets the value of DS_msDS_cloudExtensionAttribute13 for the instance
func (instance *ads_user) GetPropertyDS_msDS_cloudExtensionAttribute13() (value string, err error) {
	retValue, err := instance.GetProperty("DS_msDS_cloudExtensionAttribute13")
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

// SetDS_msDS_cloudExtensionAttribute14 sets the value of DS_msDS_cloudExtensionAttribute14 for the instance
func (instance *ads_user) SetPropertyDS_msDS_cloudExtensionAttribute14(value string) (err error) {
	return instance.SetProperty("DS_msDS_cloudExtensionAttribute14", (value))
}

// GetDS_msDS_cloudExtensionAttribute14 gets the value of DS_msDS_cloudExtensionAttribute14 for the instance
func (instance *ads_user) GetPropertyDS_msDS_cloudExtensionAttribute14() (value string, err error) {
	retValue, err := instance.GetProperty("DS_msDS_cloudExtensionAttribute14")
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

// SetDS_msDS_cloudExtensionAttribute15 sets the value of DS_msDS_cloudExtensionAttribute15 for the instance
func (instance *ads_user) SetPropertyDS_msDS_cloudExtensionAttribute15(value string) (err error) {
	return instance.SetProperty("DS_msDS_cloudExtensionAttribute15", (value))
}

// GetDS_msDS_cloudExtensionAttribute15 gets the value of DS_msDS_cloudExtensionAttribute15 for the instance
func (instance *ads_user) GetPropertyDS_msDS_cloudExtensionAttribute15() (value string, err error) {
	retValue, err := instance.GetProperty("DS_msDS_cloudExtensionAttribute15")
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

// SetDS_msDS_cloudExtensionAttribute16 sets the value of DS_msDS_cloudExtensionAttribute16 for the instance
func (instance *ads_user) SetPropertyDS_msDS_cloudExtensionAttribute16(value string) (err error) {
	return instance.SetProperty("DS_msDS_cloudExtensionAttribute16", (value))
}

// GetDS_msDS_cloudExtensionAttribute16 gets the value of DS_msDS_cloudExtensionAttribute16 for the instance
func (instance *ads_user) GetPropertyDS_msDS_cloudExtensionAttribute16() (value string, err error) {
	retValue, err := instance.GetProperty("DS_msDS_cloudExtensionAttribute16")
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

// SetDS_msDS_cloudExtensionAttribute17 sets the value of DS_msDS_cloudExtensionAttribute17 for the instance
func (instance *ads_user) SetPropertyDS_msDS_cloudExtensionAttribute17(value string) (err error) {
	return instance.SetProperty("DS_msDS_cloudExtensionAttribute17", (value))
}

// GetDS_msDS_cloudExtensionAttribute17 gets the value of DS_msDS_cloudExtensionAttribute17 for the instance
func (instance *ads_user) GetPropertyDS_msDS_cloudExtensionAttribute17() (value string, err error) {
	retValue, err := instance.GetProperty("DS_msDS_cloudExtensionAttribute17")
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

// SetDS_msDS_cloudExtensionAttribute18 sets the value of DS_msDS_cloudExtensionAttribute18 for the instance
func (instance *ads_user) SetPropertyDS_msDS_cloudExtensionAttribute18(value string) (err error) {
	return instance.SetProperty("DS_msDS_cloudExtensionAttribute18", (value))
}

// GetDS_msDS_cloudExtensionAttribute18 gets the value of DS_msDS_cloudExtensionAttribute18 for the instance
func (instance *ads_user) GetPropertyDS_msDS_cloudExtensionAttribute18() (value string, err error) {
	retValue, err := instance.GetProperty("DS_msDS_cloudExtensionAttribute18")
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

// SetDS_msDS_cloudExtensionAttribute19 sets the value of DS_msDS_cloudExtensionAttribute19 for the instance
func (instance *ads_user) SetPropertyDS_msDS_cloudExtensionAttribute19(value string) (err error) {
	return instance.SetProperty("DS_msDS_cloudExtensionAttribute19", (value))
}

// GetDS_msDS_cloudExtensionAttribute19 gets the value of DS_msDS_cloudExtensionAttribute19 for the instance
func (instance *ads_user) GetPropertyDS_msDS_cloudExtensionAttribute19() (value string, err error) {
	retValue, err := instance.GetProperty("DS_msDS_cloudExtensionAttribute19")
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

// SetDS_msDS_cloudExtensionAttribute2 sets the value of DS_msDS_cloudExtensionAttribute2 for the instance
func (instance *ads_user) SetPropertyDS_msDS_cloudExtensionAttribute2(value string) (err error) {
	return instance.SetProperty("DS_msDS_cloudExtensionAttribute2", (value))
}

// GetDS_msDS_cloudExtensionAttribute2 gets the value of DS_msDS_cloudExtensionAttribute2 for the instance
func (instance *ads_user) GetPropertyDS_msDS_cloudExtensionAttribute2() (value string, err error) {
	retValue, err := instance.GetProperty("DS_msDS_cloudExtensionAttribute2")
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

// SetDS_msDS_cloudExtensionAttribute20 sets the value of DS_msDS_cloudExtensionAttribute20 for the instance
func (instance *ads_user) SetPropertyDS_msDS_cloudExtensionAttribute20(value string) (err error) {
	return instance.SetProperty("DS_msDS_cloudExtensionAttribute20", (value))
}

// GetDS_msDS_cloudExtensionAttribute20 gets the value of DS_msDS_cloudExtensionAttribute20 for the instance
func (instance *ads_user) GetPropertyDS_msDS_cloudExtensionAttribute20() (value string, err error) {
	retValue, err := instance.GetProperty("DS_msDS_cloudExtensionAttribute20")
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

// SetDS_msDS_cloudExtensionAttribute3 sets the value of DS_msDS_cloudExtensionAttribute3 for the instance
func (instance *ads_user) SetPropertyDS_msDS_cloudExtensionAttribute3(value string) (err error) {
	return instance.SetProperty("DS_msDS_cloudExtensionAttribute3", (value))
}

// GetDS_msDS_cloudExtensionAttribute3 gets the value of DS_msDS_cloudExtensionAttribute3 for the instance
func (instance *ads_user) GetPropertyDS_msDS_cloudExtensionAttribute3() (value string, err error) {
	retValue, err := instance.GetProperty("DS_msDS_cloudExtensionAttribute3")
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

// SetDS_msDS_cloudExtensionAttribute4 sets the value of DS_msDS_cloudExtensionAttribute4 for the instance
func (instance *ads_user) SetPropertyDS_msDS_cloudExtensionAttribute4(value string) (err error) {
	return instance.SetProperty("DS_msDS_cloudExtensionAttribute4", (value))
}

// GetDS_msDS_cloudExtensionAttribute4 gets the value of DS_msDS_cloudExtensionAttribute4 for the instance
func (instance *ads_user) GetPropertyDS_msDS_cloudExtensionAttribute4() (value string, err error) {
	retValue, err := instance.GetProperty("DS_msDS_cloudExtensionAttribute4")
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

// SetDS_msDS_cloudExtensionAttribute5 sets the value of DS_msDS_cloudExtensionAttribute5 for the instance
func (instance *ads_user) SetPropertyDS_msDS_cloudExtensionAttribute5(value string) (err error) {
	return instance.SetProperty("DS_msDS_cloudExtensionAttribute5", (value))
}

// GetDS_msDS_cloudExtensionAttribute5 gets the value of DS_msDS_cloudExtensionAttribute5 for the instance
func (instance *ads_user) GetPropertyDS_msDS_cloudExtensionAttribute5() (value string, err error) {
	retValue, err := instance.GetProperty("DS_msDS_cloudExtensionAttribute5")
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

// SetDS_msDS_cloudExtensionAttribute6 sets the value of DS_msDS_cloudExtensionAttribute6 for the instance
func (instance *ads_user) SetPropertyDS_msDS_cloudExtensionAttribute6(value string) (err error) {
	return instance.SetProperty("DS_msDS_cloudExtensionAttribute6", (value))
}

// GetDS_msDS_cloudExtensionAttribute6 gets the value of DS_msDS_cloudExtensionAttribute6 for the instance
func (instance *ads_user) GetPropertyDS_msDS_cloudExtensionAttribute6() (value string, err error) {
	retValue, err := instance.GetProperty("DS_msDS_cloudExtensionAttribute6")
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

// SetDS_msDS_cloudExtensionAttribute7 sets the value of DS_msDS_cloudExtensionAttribute7 for the instance
func (instance *ads_user) SetPropertyDS_msDS_cloudExtensionAttribute7(value string) (err error) {
	return instance.SetProperty("DS_msDS_cloudExtensionAttribute7", (value))
}

// GetDS_msDS_cloudExtensionAttribute7 gets the value of DS_msDS_cloudExtensionAttribute7 for the instance
func (instance *ads_user) GetPropertyDS_msDS_cloudExtensionAttribute7() (value string, err error) {
	retValue, err := instance.GetProperty("DS_msDS_cloudExtensionAttribute7")
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

// SetDS_msDS_cloudExtensionAttribute8 sets the value of DS_msDS_cloudExtensionAttribute8 for the instance
func (instance *ads_user) SetPropertyDS_msDS_cloudExtensionAttribute8(value string) (err error) {
	return instance.SetProperty("DS_msDS_cloudExtensionAttribute8", (value))
}

// GetDS_msDS_cloudExtensionAttribute8 gets the value of DS_msDS_cloudExtensionAttribute8 for the instance
func (instance *ads_user) GetPropertyDS_msDS_cloudExtensionAttribute8() (value string, err error) {
	retValue, err := instance.GetProperty("DS_msDS_cloudExtensionAttribute8")
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

// SetDS_msDS_cloudExtensionAttribute9 sets the value of DS_msDS_cloudExtensionAttribute9 for the instance
func (instance *ads_user) SetPropertyDS_msDS_cloudExtensionAttribute9(value string) (err error) {
	return instance.SetProperty("DS_msDS_cloudExtensionAttribute9", (value))
}

// GetDS_msDS_cloudExtensionAttribute9 gets the value of DS_msDS_cloudExtensionAttribute9 for the instance
func (instance *ads_user) GetPropertyDS_msDS_cloudExtensionAttribute9() (value string, err error) {
	retValue, err := instance.GetProperty("DS_msDS_cloudExtensionAttribute9")
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

// SetDS_msDS_ExternalDirectoryObjectId sets the value of DS_msDS_ExternalDirectoryObjectId for the instance
func (instance *ads_user) SetPropertyDS_msDS_ExternalDirectoryObjectId(value string) (err error) {
	return instance.SetProperty("DS_msDS_ExternalDirectoryObjectId", (value))
}

// GetDS_msDS_ExternalDirectoryObjectId gets the value of DS_msDS_ExternalDirectoryObjectId for the instance
func (instance *ads_user) GetPropertyDS_msDS_ExternalDirectoryObjectId() (value string, err error) {
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

// SetDS_msDS_FailedInteractiveLogonCount sets the value of DS_msDS_FailedInteractiveLogonCount for the instance
func (instance *ads_user) SetPropertyDS_msDS_FailedInteractiveLogonCount(value int32) (err error) {
	return instance.SetProperty("DS_msDS_FailedInteractiveLogonCount", (value))
}

// GetDS_msDS_FailedInteractiveLogonCount gets the value of DS_msDS_FailedInteractiveLogonCount for the instance
func (instance *ads_user) GetPropertyDS_msDS_FailedInteractiveLogonCount() (value int32, err error) {
	retValue, err := instance.GetProperty("DS_msDS_FailedInteractiveLogonCount")
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

// SetDS_msDS_FailedInteractiveLogonCountAtLastSuccessfulLogon sets the value of DS_msDS_FailedInteractiveLogonCountAtLastSuccessfulLogon for the instance
func (instance *ads_user) SetPropertyDS_msDS_FailedInteractiveLogonCountAtLastSuccessfulLogon(value int32) (err error) {
	return instance.SetProperty("DS_msDS_FailedInteractiveLogonCountAtLastSuccessfulLogon", (value))
}

// GetDS_msDS_FailedInteractiveLogonCountAtLastSuccessfulLogon gets the value of DS_msDS_FailedInteractiveLogonCountAtLastSuccessfulLogon for the instance
func (instance *ads_user) GetPropertyDS_msDS_FailedInteractiveLogonCountAtLastSuccessfulLogon() (value int32, err error) {
	retValue, err := instance.GetProperty("DS_msDS_FailedInteractiveLogonCountAtLastSuccessfulLogon")
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

// SetDS_msDS_GeoCoordinatesAltitude sets the value of DS_msDS_GeoCoordinatesAltitude for the instance
func (instance *ads_user) SetPropertyDS_msDS_GeoCoordinatesAltitude(value int64) (err error) {
	return instance.SetProperty("DS_msDS_GeoCoordinatesAltitude", (value))
}

// GetDS_msDS_GeoCoordinatesAltitude gets the value of DS_msDS_GeoCoordinatesAltitude for the instance
func (instance *ads_user) GetPropertyDS_msDS_GeoCoordinatesAltitude() (value int64, err error) {
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
func (instance *ads_user) SetPropertyDS_msDS_GeoCoordinatesLatitude(value int64) (err error) {
	return instance.SetProperty("DS_msDS_GeoCoordinatesLatitude", (value))
}

// GetDS_msDS_GeoCoordinatesLatitude gets the value of DS_msDS_GeoCoordinatesLatitude for the instance
func (instance *ads_user) GetPropertyDS_msDS_GeoCoordinatesLatitude() (value int64, err error) {
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
func (instance *ads_user) SetPropertyDS_msDS_GeoCoordinatesLongitude(value int64) (err error) {
	return instance.SetProperty("DS_msDS_GeoCoordinatesLongitude", (value))
}

// GetDS_msDS_GeoCoordinatesLongitude gets the value of DS_msDS_GeoCoordinatesLongitude for the instance
func (instance *ads_user) GetPropertyDS_msDS_GeoCoordinatesLongitude() (value int64, err error) {
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

// SetDS_msDS_KeyCredentialLink sets the value of DS_msDS_KeyCredentialLink for the instance
func (instance *ads_user) SetPropertyDS_msDS_KeyCredentialLink(value []DN_With_Binary) (err error) {
	return instance.SetProperty("DS_msDS_KeyCredentialLink", (value))
}

// GetDS_msDS_KeyCredentialLink gets the value of DS_msDS_KeyCredentialLink for the instance
func (instance *ads_user) GetPropertyDS_msDS_KeyCredentialLink() (value []DN_With_Binary, err error) {
	retValue, err := instance.GetProperty("DS_msDS_KeyCredentialLink")
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

// SetDS_msDS_KeyPrincipalBL sets the value of DS_msDS_KeyPrincipalBL for the instance
func (instance *ads_user) SetPropertyDS_msDS_KeyPrincipalBL(value []string) (err error) {
	return instance.SetProperty("DS_msDS_KeyPrincipalBL", (value))
}

// GetDS_msDS_KeyPrincipalBL gets the value of DS_msDS_KeyPrincipalBL for the instance
func (instance *ads_user) GetPropertyDS_msDS_KeyPrincipalBL() (value []string, err error) {
	retValue, err := instance.GetProperty("DS_msDS_KeyPrincipalBL")
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

// SetDS_msDS_KeyVersionNumber sets the value of DS_msDS_KeyVersionNumber for the instance
func (instance *ads_user) SetPropertyDS_msDS_KeyVersionNumber(value int32) (err error) {
	return instance.SetProperty("DS_msDS_KeyVersionNumber", (value))
}

// GetDS_msDS_KeyVersionNumber gets the value of DS_msDS_KeyVersionNumber for the instance
func (instance *ads_user) GetPropertyDS_msDS_KeyVersionNumber() (value int32, err error) {
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

// SetDS_msDS_LastFailedInteractiveLogonTime sets the value of DS_msDS_LastFailedInteractiveLogonTime for the instance
func (instance *ads_user) SetPropertyDS_msDS_LastFailedInteractiveLogonTime(value int64) (err error) {
	return instance.SetProperty("DS_msDS_LastFailedInteractiveLogonTime", (value))
}

// GetDS_msDS_LastFailedInteractiveLogonTime gets the value of DS_msDS_LastFailedInteractiveLogonTime for the instance
func (instance *ads_user) GetPropertyDS_msDS_LastFailedInteractiveLogonTime() (value int64, err error) {
	retValue, err := instance.GetProperty("DS_msDS_LastFailedInteractiveLogonTime")
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

// SetDS_msDS_LastSuccessfulInteractiveLogonTime sets the value of DS_msDS_LastSuccessfulInteractiveLogonTime for the instance
func (instance *ads_user) SetPropertyDS_msDS_LastSuccessfulInteractiveLogonTime(value int64) (err error) {
	return instance.SetProperty("DS_msDS_LastSuccessfulInteractiveLogonTime", (value))
}

// GetDS_msDS_LastSuccessfulInteractiveLogonTime gets the value of DS_msDS_LastSuccessfulInteractiveLogonTime for the instance
func (instance *ads_user) GetPropertyDS_msDS_LastSuccessfulInteractiveLogonTime() (value int64, err error) {
	retValue, err := instance.GetProperty("DS_msDS_LastSuccessfulInteractiveLogonTime")
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

// SetDS_msDS_preferredDataLocation sets the value of DS_msDS_preferredDataLocation for the instance
func (instance *ads_user) SetPropertyDS_msDS_preferredDataLocation(value string) (err error) {
	return instance.SetProperty("DS_msDS_preferredDataLocation", (value))
}

// GetDS_msDS_preferredDataLocation gets the value of DS_msDS_preferredDataLocation for the instance
func (instance *ads_user) GetPropertyDS_msDS_preferredDataLocation() (value string, err error) {
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
func (instance *ads_user) SetPropertyDS_msDS_PrimaryComputer(value []string) (err error) {
	return instance.SetProperty("DS_msDS_PrimaryComputer", (value))
}

// GetDS_msDS_PrimaryComputer gets the value of DS_msDS_PrimaryComputer for the instance
func (instance *ads_user) GetPropertyDS_msDS_PrimaryComputer() (value []string, err error) {
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

// SetDS_msDS_ResultantPSO sets the value of DS_msDS_ResultantPSO for the instance
func (instance *ads_user) SetPropertyDS_msDS_ResultantPSO(value string) (err error) {
	return instance.SetProperty("DS_msDS_ResultantPSO", (value))
}

// GetDS_msDS_ResultantPSO gets the value of DS_msDS_ResultantPSO for the instance
func (instance *ads_user) GetPropertyDS_msDS_ResultantPSO() (value string, err error) {
	retValue, err := instance.GetProperty("DS_msDS_ResultantPSO")
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

// SetDS_msDS_SecondaryKrbTgtNumber sets the value of DS_msDS_SecondaryKrbTgtNumber for the instance
func (instance *ads_user) SetPropertyDS_msDS_SecondaryKrbTgtNumber(value int32) (err error) {
	return instance.SetProperty("DS_msDS_SecondaryKrbTgtNumber", (value))
}

// GetDS_msDS_SecondaryKrbTgtNumber gets the value of DS_msDS_SecondaryKrbTgtNumber for the instance
func (instance *ads_user) GetPropertyDS_msDS_SecondaryKrbTgtNumber() (value int32, err error) {
	retValue, err := instance.GetProperty("DS_msDS_SecondaryKrbTgtNumber")
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

// SetDS_msDS_Site_Affinity sets the value of DS_msDS_Site_Affinity for the instance
func (instance *ads_user) SetPropertyDS_msDS_Site_Affinity(value []Uint8Array) (err error) {
	return instance.SetProperty("DS_msDS_Site_Affinity", (value))
}

// GetDS_msDS_Site_Affinity gets the value of DS_msDS_Site_Affinity for the instance
func (instance *ads_user) GetPropertyDS_msDS_Site_Affinity() (value []Uint8Array, err error) {
	retValue, err := instance.GetProperty("DS_msDS_Site_Affinity")
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

// SetDS_msDS_SourceObjectDN sets the value of DS_msDS_SourceObjectDN for the instance
func (instance *ads_user) SetPropertyDS_msDS_SourceObjectDN(value string) (err error) {
	return instance.SetProperty("DS_msDS_SourceObjectDN", (value))
}

// GetDS_msDS_SourceObjectDN gets the value of DS_msDS_SourceObjectDN for the instance
func (instance *ads_user) GetPropertyDS_msDS_SourceObjectDN() (value string, err error) {
	retValue, err := instance.GetProperty("DS_msDS_SourceObjectDN")
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

// SetDS_msDS_SupportedEncryptionTypes sets the value of DS_msDS_SupportedEncryptionTypes for the instance
func (instance *ads_user) SetPropertyDS_msDS_SupportedEncryptionTypes(value int32) (err error) {
	return instance.SetProperty("DS_msDS_SupportedEncryptionTypes", (value))
}

// GetDS_msDS_SupportedEncryptionTypes gets the value of DS_msDS_SupportedEncryptionTypes for the instance
func (instance *ads_user) GetPropertyDS_msDS_SupportedEncryptionTypes() (value int32, err error) {
	retValue, err := instance.GetProperty("DS_msDS_SupportedEncryptionTypes")
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

// SetDS_msDS_SyncServerUrl sets the value of DS_msDS_SyncServerUrl for the instance
func (instance *ads_user) SetPropertyDS_msDS_SyncServerUrl(value []string) (err error) {
	return instance.SetProperty("DS_msDS_SyncServerUrl", (value))
}

// GetDS_msDS_SyncServerUrl gets the value of DS_msDS_SyncServerUrl for the instance
func (instance *ads_user) GetPropertyDS_msDS_SyncServerUrl() (value []string, err error) {
	retValue, err := instance.GetProperty("DS_msDS_SyncServerUrl")
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
func (instance *ads_user) SetPropertyDS_msds_tokenGroupNames(value []string) (err error) {
	return instance.SetProperty("DS_msds_tokenGroupNames", (value))
}

// GetDS_msds_tokenGroupNames gets the value of DS_msds_tokenGroupNames for the instance
func (instance *ads_user) GetPropertyDS_msds_tokenGroupNames() (value []string, err error) {
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
func (instance *ads_user) SetPropertyDS_msds_tokenGroupNamesGlobalAndUniversal(value []string) (err error) {
	return instance.SetProperty("DS_msds_tokenGroupNamesGlobalAndUniversal", (value))
}

// GetDS_msds_tokenGroupNamesGlobalAndUniversal gets the value of DS_msds_tokenGroupNamesGlobalAndUniversal for the instance
func (instance *ads_user) GetPropertyDS_msds_tokenGroupNamesGlobalAndUniversal() (value []string, err error) {
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
func (instance *ads_user) SetPropertyDS_msds_tokenGroupNamesNoGCAcceptable(value []string) (err error) {
	return instance.SetProperty("DS_msds_tokenGroupNamesNoGCAcceptable", (value))
}

// GetDS_msds_tokenGroupNamesNoGCAcceptable gets the value of DS_msds_tokenGroupNamesNoGCAcceptable for the instance
func (instance *ads_user) GetPropertyDS_msds_tokenGroupNamesNoGCAcceptable() (value []string, err error) {
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

// SetDS_msDS_User_Account_Control_Computed sets the value of DS_msDS_User_Account_Control_Computed for the instance
func (instance *ads_user) SetPropertyDS_msDS_User_Account_Control_Computed(value int32) (err error) {
	return instance.SetProperty("DS_msDS_User_Account_Control_Computed", (value))
}

// GetDS_msDS_User_Account_Control_Computed gets the value of DS_msDS_User_Account_Control_Computed for the instance
func (instance *ads_user) GetPropertyDS_msDS_User_Account_Control_Computed() (value int32, err error) {
	retValue, err := instance.GetProperty("DS_msDS_User_Account_Control_Computed")
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

// SetDS_msDS_UserPasswordExpiryTimeComputed sets the value of DS_msDS_UserPasswordExpiryTimeComputed for the instance
func (instance *ads_user) SetPropertyDS_msDS_UserPasswordExpiryTimeComputed(value int64) (err error) {
	return instance.SetProperty("DS_msDS_UserPasswordExpiryTimeComputed", (value))
}

// GetDS_msDS_UserPasswordExpiryTimeComputed gets the value of DS_msDS_UserPasswordExpiryTimeComputed for the instance
func (instance *ads_user) GetPropertyDS_msDS_UserPasswordExpiryTimeComputed() (value int64, err error) {
	retValue, err := instance.GetProperty("DS_msDS_UserPasswordExpiryTimeComputed")
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

// SetDS_msExchAssistantName sets the value of DS_msExchAssistantName for the instance
func (instance *ads_user) SetPropertyDS_msExchAssistantName(value string) (err error) {
	return instance.SetProperty("DS_msExchAssistantName", (value))
}

// GetDS_msExchAssistantName gets the value of DS_msExchAssistantName for the instance
func (instance *ads_user) GetPropertyDS_msExchAssistantName() (value string, err error) {
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
func (instance *ads_user) SetPropertyDS_msExchLabeledURI(value []string) (err error) {
	return instance.SetProperty("DS_msExchLabeledURI", (value))
}

// GetDS_msExchLabeledURI gets the value of DS_msExchLabeledURI for the instance
func (instance *ads_user) GetPropertyDS_msExchLabeledURI() (value []string, err error) {
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

// SetDS_msIIS_FTPDir sets the value of DS_msIIS_FTPDir for the instance
func (instance *ads_user) SetPropertyDS_msIIS_FTPDir(value string) (err error) {
	return instance.SetProperty("DS_msIIS_FTPDir", (value))
}

// GetDS_msIIS_FTPDir gets the value of DS_msIIS_FTPDir for the instance
func (instance *ads_user) GetPropertyDS_msIIS_FTPDir() (value string, err error) {
	retValue, err := instance.GetProperty("DS_msIIS_FTPDir")
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

// SetDS_msIIS_FTPRoot sets the value of DS_msIIS_FTPRoot for the instance
func (instance *ads_user) SetPropertyDS_msIIS_FTPRoot(value string) (err error) {
	return instance.SetProperty("DS_msIIS_FTPRoot", (value))
}

// GetDS_msIIS_FTPRoot gets the value of DS_msIIS_FTPRoot for the instance
func (instance *ads_user) GetPropertyDS_msIIS_FTPRoot() (value string, err error) {
	retValue, err := instance.GetProperty("DS_msIIS_FTPRoot")
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

// SetDS_mSMQDigests sets the value of DS_mSMQDigests for the instance
func (instance *ads_user) SetPropertyDS_mSMQDigests(value []Uint8Array) (err error) {
	return instance.SetProperty("DS_mSMQDigests", (value))
}

// GetDS_mSMQDigests gets the value of DS_mSMQDigests for the instance
func (instance *ads_user) GetPropertyDS_mSMQDigests() (value []Uint8Array, err error) {
	retValue, err := instance.GetProperty("DS_mSMQDigests")
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

// SetDS_mSMQDigestsMig sets the value of DS_mSMQDigestsMig for the instance
func (instance *ads_user) SetPropertyDS_mSMQDigestsMig(value []Uint8Array) (err error) {
	return instance.SetProperty("DS_mSMQDigestsMig", (value))
}

// GetDS_mSMQDigestsMig gets the value of DS_mSMQDigestsMig for the instance
func (instance *ads_user) GetPropertyDS_mSMQDigestsMig() (value []Uint8Array, err error) {
	retValue, err := instance.GetProperty("DS_mSMQDigestsMig")
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

// SetDS_mSMQSignCertificates sets the value of DS_mSMQSignCertificates for the instance
func (instance *ads_user) SetPropertyDS_mSMQSignCertificates(value Uint8Array) (err error) {
	return instance.SetProperty("DS_mSMQSignCertificates", (value))
}

// GetDS_mSMQSignCertificates gets the value of DS_mSMQSignCertificates for the instance
func (instance *ads_user) GetPropertyDS_mSMQSignCertificates() (value Uint8Array, err error) {
	retValue, err := instance.GetProperty("DS_mSMQSignCertificates")
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

// SetDS_mSMQSignCertificatesMig sets the value of DS_mSMQSignCertificatesMig for the instance
func (instance *ads_user) SetPropertyDS_mSMQSignCertificatesMig(value Uint8Array) (err error) {
	return instance.SetProperty("DS_mSMQSignCertificatesMig", (value))
}

// GetDS_mSMQSignCertificatesMig gets the value of DS_mSMQSignCertificatesMig for the instance
func (instance *ads_user) GetPropertyDS_mSMQSignCertificatesMig() (value Uint8Array, err error) {
	retValue, err := instance.GetProperty("DS_mSMQSignCertificatesMig")
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

// SetDS_msNPAllowDialin sets the value of DS_msNPAllowDialin for the instance
func (instance *ads_user) SetPropertyDS_msNPAllowDialin(value bool) (err error) {
	return instance.SetProperty("DS_msNPAllowDialin", (value))
}

// GetDS_msNPAllowDialin gets the value of DS_msNPAllowDialin for the instance
func (instance *ads_user) GetPropertyDS_msNPAllowDialin() (value bool, err error) {
	retValue, err := instance.GetProperty("DS_msNPAllowDialin")
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

// SetDS_msNPCallingStationID sets the value of DS_msNPCallingStationID for the instance
func (instance *ads_user) SetPropertyDS_msNPCallingStationID(value []string) (err error) {
	return instance.SetProperty("DS_msNPCallingStationID", (value))
}

// GetDS_msNPCallingStationID gets the value of DS_msNPCallingStationID for the instance
func (instance *ads_user) GetPropertyDS_msNPCallingStationID() (value []string, err error) {
	retValue, err := instance.GetProperty("DS_msNPCallingStationID")
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

// SetDS_msNPSavedCallingStationID sets the value of DS_msNPSavedCallingStationID for the instance
func (instance *ads_user) SetPropertyDS_msNPSavedCallingStationID(value []string) (err error) {
	return instance.SetProperty("DS_msNPSavedCallingStationID", (value))
}

// GetDS_msNPSavedCallingStationID gets the value of DS_msNPSavedCallingStationID for the instance
func (instance *ads_user) GetPropertyDS_msNPSavedCallingStationID() (value []string, err error) {
	retValue, err := instance.GetProperty("DS_msNPSavedCallingStationID")
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

// SetDS_msPKI_CredentialRoamingTokens sets the value of DS_msPKI_CredentialRoamingTokens for the instance
func (instance *ads_user) SetPropertyDS_msPKI_CredentialRoamingTokens(value []DN_With_Binary) (err error) {
	return instance.SetProperty("DS_msPKI_CredentialRoamingTokens", (value))
}

// GetDS_msPKI_CredentialRoamingTokens gets the value of DS_msPKI_CredentialRoamingTokens for the instance
func (instance *ads_user) GetPropertyDS_msPKI_CredentialRoamingTokens() (value []DN_With_Binary, err error) {
	retValue, err := instance.GetProperty("DS_msPKI_CredentialRoamingTokens")
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

// SetDS_msPKIAccountCredentials sets the value of DS_msPKIAccountCredentials for the instance
func (instance *ads_user) SetPropertyDS_msPKIAccountCredentials(value []DN_With_Binary) (err error) {
	return instance.SetProperty("DS_msPKIAccountCredentials", (value))
}

// GetDS_msPKIAccountCredentials gets the value of DS_msPKIAccountCredentials for the instance
func (instance *ads_user) GetPropertyDS_msPKIAccountCredentials() (value []DN_With_Binary, err error) {
	retValue, err := instance.GetProperty("DS_msPKIAccountCredentials")
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

// SetDS_msPKIDPAPIMasterKeys sets the value of DS_msPKIDPAPIMasterKeys for the instance
func (instance *ads_user) SetPropertyDS_msPKIDPAPIMasterKeys(value []DN_With_Binary) (err error) {
	return instance.SetProperty("DS_msPKIDPAPIMasterKeys", (value))
}

// GetDS_msPKIDPAPIMasterKeys gets the value of DS_msPKIDPAPIMasterKeys for the instance
func (instance *ads_user) GetPropertyDS_msPKIDPAPIMasterKeys() (value []DN_With_Binary, err error) {
	retValue, err := instance.GetProperty("DS_msPKIDPAPIMasterKeys")
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

// SetDS_msPKIRoamingTimeStamp sets the value of DS_msPKIRoamingTimeStamp for the instance
func (instance *ads_user) SetPropertyDS_msPKIRoamingTimeStamp(value Uint8Array) (err error) {
	return instance.SetProperty("DS_msPKIRoamingTimeStamp", (value))
}

// GetDS_msPKIRoamingTimeStamp gets the value of DS_msPKIRoamingTimeStamp for the instance
func (instance *ads_user) GetPropertyDS_msPKIRoamingTimeStamp() (value Uint8Array, err error) {
	retValue, err := instance.GetProperty("DS_msPKIRoamingTimeStamp")
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

// SetDS_msRADIUS_FramedInterfaceId sets the value of DS_msRADIUS_FramedInterfaceId for the instance
func (instance *ads_user) SetPropertyDS_msRADIUS_FramedInterfaceId(value string) (err error) {
	return instance.SetProperty("DS_msRADIUS_FramedInterfaceId", (value))
}

// GetDS_msRADIUS_FramedInterfaceId gets the value of DS_msRADIUS_FramedInterfaceId for the instance
func (instance *ads_user) GetPropertyDS_msRADIUS_FramedInterfaceId() (value string, err error) {
	retValue, err := instance.GetProperty("DS_msRADIUS_FramedInterfaceId")
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

// SetDS_msRADIUS_FramedIpv6Prefix sets the value of DS_msRADIUS_FramedIpv6Prefix for the instance
func (instance *ads_user) SetPropertyDS_msRADIUS_FramedIpv6Prefix(value string) (err error) {
	return instance.SetProperty("DS_msRADIUS_FramedIpv6Prefix", (value))
}

// GetDS_msRADIUS_FramedIpv6Prefix gets the value of DS_msRADIUS_FramedIpv6Prefix for the instance
func (instance *ads_user) GetPropertyDS_msRADIUS_FramedIpv6Prefix() (value string, err error) {
	retValue, err := instance.GetProperty("DS_msRADIUS_FramedIpv6Prefix")
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

// SetDS_msRADIUS_FramedIpv6Route sets the value of DS_msRADIUS_FramedIpv6Route for the instance
func (instance *ads_user) SetPropertyDS_msRADIUS_FramedIpv6Route(value []string) (err error) {
	return instance.SetProperty("DS_msRADIUS_FramedIpv6Route", (value))
}

// GetDS_msRADIUS_FramedIpv6Route gets the value of DS_msRADIUS_FramedIpv6Route for the instance
func (instance *ads_user) GetPropertyDS_msRADIUS_FramedIpv6Route() (value []string, err error) {
	retValue, err := instance.GetProperty("DS_msRADIUS_FramedIpv6Route")
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

// SetDS_msRADIUS_SavedFramedInterfaceId sets the value of DS_msRADIUS_SavedFramedInterfaceId for the instance
func (instance *ads_user) SetPropertyDS_msRADIUS_SavedFramedInterfaceId(value string) (err error) {
	return instance.SetProperty("DS_msRADIUS_SavedFramedInterfaceId", (value))
}

// GetDS_msRADIUS_SavedFramedInterfaceId gets the value of DS_msRADIUS_SavedFramedInterfaceId for the instance
func (instance *ads_user) GetPropertyDS_msRADIUS_SavedFramedInterfaceId() (value string, err error) {
	retValue, err := instance.GetProperty("DS_msRADIUS_SavedFramedInterfaceId")
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

// SetDS_msRADIUS_SavedFramedIpv6Prefix sets the value of DS_msRADIUS_SavedFramedIpv6Prefix for the instance
func (instance *ads_user) SetPropertyDS_msRADIUS_SavedFramedIpv6Prefix(value string) (err error) {
	return instance.SetProperty("DS_msRADIUS_SavedFramedIpv6Prefix", (value))
}

// GetDS_msRADIUS_SavedFramedIpv6Prefix gets the value of DS_msRADIUS_SavedFramedIpv6Prefix for the instance
func (instance *ads_user) GetPropertyDS_msRADIUS_SavedFramedIpv6Prefix() (value string, err error) {
	retValue, err := instance.GetProperty("DS_msRADIUS_SavedFramedIpv6Prefix")
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

// SetDS_msRADIUS_SavedFramedIpv6Route sets the value of DS_msRADIUS_SavedFramedIpv6Route for the instance
func (instance *ads_user) SetPropertyDS_msRADIUS_SavedFramedIpv6Route(value []string) (err error) {
	return instance.SetProperty("DS_msRADIUS_SavedFramedIpv6Route", (value))
}

// GetDS_msRADIUS_SavedFramedIpv6Route gets the value of DS_msRADIUS_SavedFramedIpv6Route for the instance
func (instance *ads_user) GetPropertyDS_msRADIUS_SavedFramedIpv6Route() (value []string, err error) {
	retValue, err := instance.GetProperty("DS_msRADIUS_SavedFramedIpv6Route")
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

// SetDS_msRADIUSCallbackNumber sets the value of DS_msRADIUSCallbackNumber for the instance
func (instance *ads_user) SetPropertyDS_msRADIUSCallbackNumber(value string) (err error) {
	return instance.SetProperty("DS_msRADIUSCallbackNumber", (value))
}

// GetDS_msRADIUSCallbackNumber gets the value of DS_msRADIUSCallbackNumber for the instance
func (instance *ads_user) GetPropertyDS_msRADIUSCallbackNumber() (value string, err error) {
	retValue, err := instance.GetProperty("DS_msRADIUSCallbackNumber")
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

// SetDS_msRADIUSFramedIPAddress sets the value of DS_msRADIUSFramedIPAddress for the instance
func (instance *ads_user) SetPropertyDS_msRADIUSFramedIPAddress(value int32) (err error) {
	return instance.SetProperty("DS_msRADIUSFramedIPAddress", (value))
}

// GetDS_msRADIUSFramedIPAddress gets the value of DS_msRADIUSFramedIPAddress for the instance
func (instance *ads_user) GetPropertyDS_msRADIUSFramedIPAddress() (value int32, err error) {
	retValue, err := instance.GetProperty("DS_msRADIUSFramedIPAddress")
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

// SetDS_msRADIUSFramedRoute sets the value of DS_msRADIUSFramedRoute for the instance
func (instance *ads_user) SetPropertyDS_msRADIUSFramedRoute(value []string) (err error) {
	return instance.SetProperty("DS_msRADIUSFramedRoute", (value))
}

// GetDS_msRADIUSFramedRoute gets the value of DS_msRADIUSFramedRoute for the instance
func (instance *ads_user) GetPropertyDS_msRADIUSFramedRoute() (value []string, err error) {
	retValue, err := instance.GetProperty("DS_msRADIUSFramedRoute")
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

// SetDS_msRADIUSServiceType sets the value of DS_msRADIUSServiceType for the instance
func (instance *ads_user) SetPropertyDS_msRADIUSServiceType(value int32) (err error) {
	return instance.SetProperty("DS_msRADIUSServiceType", (value))
}

// GetDS_msRADIUSServiceType gets the value of DS_msRADIUSServiceType for the instance
func (instance *ads_user) GetPropertyDS_msRADIUSServiceType() (value int32, err error) {
	retValue, err := instance.GetProperty("DS_msRADIUSServiceType")
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

// SetDS_msRASSavedCallbackNumber sets the value of DS_msRASSavedCallbackNumber for the instance
func (instance *ads_user) SetPropertyDS_msRASSavedCallbackNumber(value string) (err error) {
	return instance.SetProperty("DS_msRASSavedCallbackNumber", (value))
}

// GetDS_msRASSavedCallbackNumber gets the value of DS_msRASSavedCallbackNumber for the instance
func (instance *ads_user) GetPropertyDS_msRASSavedCallbackNumber() (value string, err error) {
	retValue, err := instance.GetProperty("DS_msRASSavedCallbackNumber")
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

// SetDS_msRASSavedFramedIPAddress sets the value of DS_msRASSavedFramedIPAddress for the instance
func (instance *ads_user) SetPropertyDS_msRASSavedFramedIPAddress(value int32) (err error) {
	return instance.SetProperty("DS_msRASSavedFramedIPAddress", (value))
}

// GetDS_msRASSavedFramedIPAddress gets the value of DS_msRASSavedFramedIPAddress for the instance
func (instance *ads_user) GetPropertyDS_msRASSavedFramedIPAddress() (value int32, err error) {
	retValue, err := instance.GetProperty("DS_msRASSavedFramedIPAddress")
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

// SetDS_msRASSavedFramedRoute sets the value of DS_msRASSavedFramedRoute for the instance
func (instance *ads_user) SetPropertyDS_msRASSavedFramedRoute(value []string) (err error) {
	return instance.SetProperty("DS_msRASSavedFramedRoute", (value))
}

// GetDS_msRASSavedFramedRoute gets the value of DS_msRASSavedFramedRoute for the instance
func (instance *ads_user) GetPropertyDS_msRASSavedFramedRoute() (value []string, err error) {
	retValue, err := instance.GetProperty("DS_msRASSavedFramedRoute")
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
func (instance *ads_user) SetPropertyDS_msSFU30Name(value string) (err error) {
	return instance.SetProperty("DS_msSFU30Name", (value))
}

// GetDS_msSFU30Name gets the value of DS_msSFU30Name for the instance
func (instance *ads_user) GetPropertyDS_msSFU30Name() (value string, err error) {
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
func (instance *ads_user) SetPropertyDS_msSFU30NisDomain(value string) (err error) {
	return instance.SetProperty("DS_msSFU30NisDomain", (value))
}

// GetDS_msSFU30NisDomain gets the value of DS_msSFU30NisDomain for the instance
func (instance *ads_user) GetPropertyDS_msSFU30NisDomain() (value string, err error) {
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

// SetDS_msTSAllowLogon sets the value of DS_msTSAllowLogon for the instance
func (instance *ads_user) SetPropertyDS_msTSAllowLogon(value bool) (err error) {
	return instance.SetProperty("DS_msTSAllowLogon", (value))
}

// GetDS_msTSAllowLogon gets the value of DS_msTSAllowLogon for the instance
func (instance *ads_user) GetPropertyDS_msTSAllowLogon() (value bool, err error) {
	retValue, err := instance.GetProperty("DS_msTSAllowLogon")
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

// SetDS_msTSBrokenConnectionAction sets the value of DS_msTSBrokenConnectionAction for the instance
func (instance *ads_user) SetPropertyDS_msTSBrokenConnectionAction(value bool) (err error) {
	return instance.SetProperty("DS_msTSBrokenConnectionAction", (value))
}

// GetDS_msTSBrokenConnectionAction gets the value of DS_msTSBrokenConnectionAction for the instance
func (instance *ads_user) GetPropertyDS_msTSBrokenConnectionAction() (value bool, err error) {
	retValue, err := instance.GetProperty("DS_msTSBrokenConnectionAction")
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

// SetDS_msTSConnectClientDrives sets the value of DS_msTSConnectClientDrives for the instance
func (instance *ads_user) SetPropertyDS_msTSConnectClientDrives(value bool) (err error) {
	return instance.SetProperty("DS_msTSConnectClientDrives", (value))
}

// GetDS_msTSConnectClientDrives gets the value of DS_msTSConnectClientDrives for the instance
func (instance *ads_user) GetPropertyDS_msTSConnectClientDrives() (value bool, err error) {
	retValue, err := instance.GetProperty("DS_msTSConnectClientDrives")
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

// SetDS_msTSConnectPrinterDrives sets the value of DS_msTSConnectPrinterDrives for the instance
func (instance *ads_user) SetPropertyDS_msTSConnectPrinterDrives(value bool) (err error) {
	return instance.SetProperty("DS_msTSConnectPrinterDrives", (value))
}

// GetDS_msTSConnectPrinterDrives gets the value of DS_msTSConnectPrinterDrives for the instance
func (instance *ads_user) GetPropertyDS_msTSConnectPrinterDrives() (value bool, err error) {
	retValue, err := instance.GetProperty("DS_msTSConnectPrinterDrives")
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

// SetDS_msTSDefaultToMainPrinter sets the value of DS_msTSDefaultToMainPrinter for the instance
func (instance *ads_user) SetPropertyDS_msTSDefaultToMainPrinter(value bool) (err error) {
	return instance.SetProperty("DS_msTSDefaultToMainPrinter", (value))
}

// GetDS_msTSDefaultToMainPrinter gets the value of DS_msTSDefaultToMainPrinter for the instance
func (instance *ads_user) GetPropertyDS_msTSDefaultToMainPrinter() (value bool, err error) {
	retValue, err := instance.GetProperty("DS_msTSDefaultToMainPrinter")
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

// SetDS_msTSExpireDate sets the value of DS_msTSExpireDate for the instance
func (instance *ads_user) SetPropertyDS_msTSExpireDate(value string) (err error) {
	return instance.SetProperty("DS_msTSExpireDate", (value))
}

// GetDS_msTSExpireDate gets the value of DS_msTSExpireDate for the instance
func (instance *ads_user) GetPropertyDS_msTSExpireDate() (value string, err error) {
	retValue, err := instance.GetProperty("DS_msTSExpireDate")
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

// SetDS_msTSExpireDate2 sets the value of DS_msTSExpireDate2 for the instance
func (instance *ads_user) SetPropertyDS_msTSExpireDate2(value string) (err error) {
	return instance.SetProperty("DS_msTSExpireDate2", (value))
}

// GetDS_msTSExpireDate2 gets the value of DS_msTSExpireDate2 for the instance
func (instance *ads_user) GetPropertyDS_msTSExpireDate2() (value string, err error) {
	retValue, err := instance.GetProperty("DS_msTSExpireDate2")
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

// SetDS_msTSExpireDate3 sets the value of DS_msTSExpireDate3 for the instance
func (instance *ads_user) SetPropertyDS_msTSExpireDate3(value string) (err error) {
	return instance.SetProperty("DS_msTSExpireDate3", (value))
}

// GetDS_msTSExpireDate3 gets the value of DS_msTSExpireDate3 for the instance
func (instance *ads_user) GetPropertyDS_msTSExpireDate3() (value string, err error) {
	retValue, err := instance.GetProperty("DS_msTSExpireDate3")
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

// SetDS_msTSExpireDate4 sets the value of DS_msTSExpireDate4 for the instance
func (instance *ads_user) SetPropertyDS_msTSExpireDate4(value string) (err error) {
	return instance.SetProperty("DS_msTSExpireDate4", (value))
}

// GetDS_msTSExpireDate4 gets the value of DS_msTSExpireDate4 for the instance
func (instance *ads_user) GetPropertyDS_msTSExpireDate4() (value string, err error) {
	retValue, err := instance.GetProperty("DS_msTSExpireDate4")
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

// SetDS_msTSHomeDirectory sets the value of DS_msTSHomeDirectory for the instance
func (instance *ads_user) SetPropertyDS_msTSHomeDirectory(value string) (err error) {
	return instance.SetProperty("DS_msTSHomeDirectory", (value))
}

// GetDS_msTSHomeDirectory gets the value of DS_msTSHomeDirectory for the instance
func (instance *ads_user) GetPropertyDS_msTSHomeDirectory() (value string, err error) {
	retValue, err := instance.GetProperty("DS_msTSHomeDirectory")
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

// SetDS_msTSHomeDrive sets the value of DS_msTSHomeDrive for the instance
func (instance *ads_user) SetPropertyDS_msTSHomeDrive(value string) (err error) {
	return instance.SetProperty("DS_msTSHomeDrive", (value))
}

// GetDS_msTSHomeDrive gets the value of DS_msTSHomeDrive for the instance
func (instance *ads_user) GetPropertyDS_msTSHomeDrive() (value string, err error) {
	retValue, err := instance.GetProperty("DS_msTSHomeDrive")
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

// SetDS_msTSInitialProgram sets the value of DS_msTSInitialProgram for the instance
func (instance *ads_user) SetPropertyDS_msTSInitialProgram(value string) (err error) {
	return instance.SetProperty("DS_msTSInitialProgram", (value))
}

// GetDS_msTSInitialProgram gets the value of DS_msTSInitialProgram for the instance
func (instance *ads_user) GetPropertyDS_msTSInitialProgram() (value string, err error) {
	retValue, err := instance.GetProperty("DS_msTSInitialProgram")
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

// SetDS_msTSLicenseVersion sets the value of DS_msTSLicenseVersion for the instance
func (instance *ads_user) SetPropertyDS_msTSLicenseVersion(value string) (err error) {
	return instance.SetProperty("DS_msTSLicenseVersion", (value))
}

// GetDS_msTSLicenseVersion gets the value of DS_msTSLicenseVersion for the instance
func (instance *ads_user) GetPropertyDS_msTSLicenseVersion() (value string, err error) {
	retValue, err := instance.GetProperty("DS_msTSLicenseVersion")
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

// SetDS_msTSLicenseVersion2 sets the value of DS_msTSLicenseVersion2 for the instance
func (instance *ads_user) SetPropertyDS_msTSLicenseVersion2(value string) (err error) {
	return instance.SetProperty("DS_msTSLicenseVersion2", (value))
}

// GetDS_msTSLicenseVersion2 gets the value of DS_msTSLicenseVersion2 for the instance
func (instance *ads_user) GetPropertyDS_msTSLicenseVersion2() (value string, err error) {
	retValue, err := instance.GetProperty("DS_msTSLicenseVersion2")
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

// SetDS_msTSLicenseVersion3 sets the value of DS_msTSLicenseVersion3 for the instance
func (instance *ads_user) SetPropertyDS_msTSLicenseVersion3(value string) (err error) {
	return instance.SetProperty("DS_msTSLicenseVersion3", (value))
}

// GetDS_msTSLicenseVersion3 gets the value of DS_msTSLicenseVersion3 for the instance
func (instance *ads_user) GetPropertyDS_msTSLicenseVersion3() (value string, err error) {
	retValue, err := instance.GetProperty("DS_msTSLicenseVersion3")
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

// SetDS_msTSLicenseVersion4 sets the value of DS_msTSLicenseVersion4 for the instance
func (instance *ads_user) SetPropertyDS_msTSLicenseVersion4(value string) (err error) {
	return instance.SetProperty("DS_msTSLicenseVersion4", (value))
}

// GetDS_msTSLicenseVersion4 gets the value of DS_msTSLicenseVersion4 for the instance
func (instance *ads_user) GetPropertyDS_msTSLicenseVersion4() (value string, err error) {
	retValue, err := instance.GetProperty("DS_msTSLicenseVersion4")
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

// SetDS_msTSLSProperty01 sets the value of DS_msTSLSProperty01 for the instance
func (instance *ads_user) SetPropertyDS_msTSLSProperty01(value []string) (err error) {
	return instance.SetProperty("DS_msTSLSProperty01", (value))
}

// GetDS_msTSLSProperty01 gets the value of DS_msTSLSProperty01 for the instance
func (instance *ads_user) GetPropertyDS_msTSLSProperty01() (value []string, err error) {
	retValue, err := instance.GetProperty("DS_msTSLSProperty01")
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

// SetDS_msTSLSProperty02 sets the value of DS_msTSLSProperty02 for the instance
func (instance *ads_user) SetPropertyDS_msTSLSProperty02(value []string) (err error) {
	return instance.SetProperty("DS_msTSLSProperty02", (value))
}

// GetDS_msTSLSProperty02 gets the value of DS_msTSLSProperty02 for the instance
func (instance *ads_user) GetPropertyDS_msTSLSProperty02() (value []string, err error) {
	retValue, err := instance.GetProperty("DS_msTSLSProperty02")
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

// SetDS_msTSManagingLS sets the value of DS_msTSManagingLS for the instance
func (instance *ads_user) SetPropertyDS_msTSManagingLS(value string) (err error) {
	return instance.SetProperty("DS_msTSManagingLS", (value))
}

// GetDS_msTSManagingLS gets the value of DS_msTSManagingLS for the instance
func (instance *ads_user) GetPropertyDS_msTSManagingLS() (value string, err error) {
	retValue, err := instance.GetProperty("DS_msTSManagingLS")
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

// SetDS_msTSManagingLS2 sets the value of DS_msTSManagingLS2 for the instance
func (instance *ads_user) SetPropertyDS_msTSManagingLS2(value string) (err error) {
	return instance.SetProperty("DS_msTSManagingLS2", (value))
}

// GetDS_msTSManagingLS2 gets the value of DS_msTSManagingLS2 for the instance
func (instance *ads_user) GetPropertyDS_msTSManagingLS2() (value string, err error) {
	retValue, err := instance.GetProperty("DS_msTSManagingLS2")
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

// SetDS_msTSManagingLS3 sets the value of DS_msTSManagingLS3 for the instance
func (instance *ads_user) SetPropertyDS_msTSManagingLS3(value string) (err error) {
	return instance.SetProperty("DS_msTSManagingLS3", (value))
}

// GetDS_msTSManagingLS3 gets the value of DS_msTSManagingLS3 for the instance
func (instance *ads_user) GetPropertyDS_msTSManagingLS3() (value string, err error) {
	retValue, err := instance.GetProperty("DS_msTSManagingLS3")
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

// SetDS_msTSManagingLS4 sets the value of DS_msTSManagingLS4 for the instance
func (instance *ads_user) SetPropertyDS_msTSManagingLS4(value string) (err error) {
	return instance.SetProperty("DS_msTSManagingLS4", (value))
}

// GetDS_msTSManagingLS4 gets the value of DS_msTSManagingLS4 for the instance
func (instance *ads_user) GetPropertyDS_msTSManagingLS4() (value string, err error) {
	retValue, err := instance.GetProperty("DS_msTSManagingLS4")
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

// SetDS_msTSMaxConnectionTime sets the value of DS_msTSMaxConnectionTime for the instance
func (instance *ads_user) SetPropertyDS_msTSMaxConnectionTime(value int32) (err error) {
	return instance.SetProperty("DS_msTSMaxConnectionTime", (value))
}

// GetDS_msTSMaxConnectionTime gets the value of DS_msTSMaxConnectionTime for the instance
func (instance *ads_user) GetPropertyDS_msTSMaxConnectionTime() (value int32, err error) {
	retValue, err := instance.GetProperty("DS_msTSMaxConnectionTime")
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

// SetDS_msTSMaxDisconnectionTime sets the value of DS_msTSMaxDisconnectionTime for the instance
func (instance *ads_user) SetPropertyDS_msTSMaxDisconnectionTime(value int32) (err error) {
	return instance.SetProperty("DS_msTSMaxDisconnectionTime", (value))
}

// GetDS_msTSMaxDisconnectionTime gets the value of DS_msTSMaxDisconnectionTime for the instance
func (instance *ads_user) GetPropertyDS_msTSMaxDisconnectionTime() (value int32, err error) {
	retValue, err := instance.GetProperty("DS_msTSMaxDisconnectionTime")
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

// SetDS_msTSMaxIdleTime sets the value of DS_msTSMaxIdleTime for the instance
func (instance *ads_user) SetPropertyDS_msTSMaxIdleTime(value int32) (err error) {
	return instance.SetProperty("DS_msTSMaxIdleTime", (value))
}

// GetDS_msTSMaxIdleTime gets the value of DS_msTSMaxIdleTime for the instance
func (instance *ads_user) GetPropertyDS_msTSMaxIdleTime() (value int32, err error) {
	retValue, err := instance.GetProperty("DS_msTSMaxIdleTime")
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

// SetDS_msTSPrimaryDesktop sets the value of DS_msTSPrimaryDesktop for the instance
func (instance *ads_user) SetPropertyDS_msTSPrimaryDesktop(value string) (err error) {
	return instance.SetProperty("DS_msTSPrimaryDesktop", (value))
}

// GetDS_msTSPrimaryDesktop gets the value of DS_msTSPrimaryDesktop for the instance
func (instance *ads_user) GetPropertyDS_msTSPrimaryDesktop() (value string, err error) {
	retValue, err := instance.GetProperty("DS_msTSPrimaryDesktop")
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

// SetDS_msTSProfilePath sets the value of DS_msTSProfilePath for the instance
func (instance *ads_user) SetPropertyDS_msTSProfilePath(value string) (err error) {
	return instance.SetProperty("DS_msTSProfilePath", (value))
}

// GetDS_msTSProfilePath gets the value of DS_msTSProfilePath for the instance
func (instance *ads_user) GetPropertyDS_msTSProfilePath() (value string, err error) {
	retValue, err := instance.GetProperty("DS_msTSProfilePath")
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

// SetDS_msTSProperty01 sets the value of DS_msTSProperty01 for the instance
func (instance *ads_user) SetPropertyDS_msTSProperty01(value []string) (err error) {
	return instance.SetProperty("DS_msTSProperty01", (value))
}

// GetDS_msTSProperty01 gets the value of DS_msTSProperty01 for the instance
func (instance *ads_user) GetPropertyDS_msTSProperty01() (value []string, err error) {
	retValue, err := instance.GetProperty("DS_msTSProperty01")
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

// SetDS_msTSProperty02 sets the value of DS_msTSProperty02 for the instance
func (instance *ads_user) SetPropertyDS_msTSProperty02(value []string) (err error) {
	return instance.SetProperty("DS_msTSProperty02", (value))
}

// GetDS_msTSProperty02 gets the value of DS_msTSProperty02 for the instance
func (instance *ads_user) GetPropertyDS_msTSProperty02() (value []string, err error) {
	retValue, err := instance.GetProperty("DS_msTSProperty02")
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

// SetDS_msTSReconnectionAction sets the value of DS_msTSReconnectionAction for the instance
func (instance *ads_user) SetPropertyDS_msTSReconnectionAction(value bool) (err error) {
	return instance.SetProperty("DS_msTSReconnectionAction", (value))
}

// GetDS_msTSReconnectionAction gets the value of DS_msTSReconnectionAction for the instance
func (instance *ads_user) GetPropertyDS_msTSReconnectionAction() (value bool, err error) {
	retValue, err := instance.GetProperty("DS_msTSReconnectionAction")
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

// SetDS_msTSRemoteControl sets the value of DS_msTSRemoteControl for the instance
func (instance *ads_user) SetPropertyDS_msTSRemoteControl(value int32) (err error) {
	return instance.SetProperty("DS_msTSRemoteControl", (value))
}

// GetDS_msTSRemoteControl gets the value of DS_msTSRemoteControl for the instance
func (instance *ads_user) GetPropertyDS_msTSRemoteControl() (value int32, err error) {
	retValue, err := instance.GetProperty("DS_msTSRemoteControl")
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

// SetDS_msTSSecondaryDesktops sets the value of DS_msTSSecondaryDesktops for the instance
func (instance *ads_user) SetPropertyDS_msTSSecondaryDesktops(value []string) (err error) {
	return instance.SetProperty("DS_msTSSecondaryDesktops", (value))
}

// GetDS_msTSSecondaryDesktops gets the value of DS_msTSSecondaryDesktops for the instance
func (instance *ads_user) GetPropertyDS_msTSSecondaryDesktops() (value []string, err error) {
	retValue, err := instance.GetProperty("DS_msTSSecondaryDesktops")
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

// SetDS_msTSWorkDirectory sets the value of DS_msTSWorkDirectory for the instance
func (instance *ads_user) SetPropertyDS_msTSWorkDirectory(value string) (err error) {
	return instance.SetProperty("DS_msTSWorkDirectory", (value))
}

// GetDS_msTSWorkDirectory gets the value of DS_msTSWorkDirectory for the instance
func (instance *ads_user) GetPropertyDS_msTSWorkDirectory() (value string, err error) {
	retValue, err := instance.GetProperty("DS_msTSWorkDirectory")
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
func (instance *ads_user) SetPropertyDS_networkAddress(value []string) (err error) {
	return instance.SetProperty("DS_networkAddress", (value))
}

// GetDS_networkAddress gets the value of DS_networkAddress for the instance
func (instance *ads_user) GetPropertyDS_networkAddress() (value []string, err error) {
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

// SetDS_ntPwdHistory sets the value of DS_ntPwdHistory for the instance
func (instance *ads_user) SetPropertyDS_ntPwdHistory(value []Uint8Array) (err error) {
	return instance.SetProperty("DS_ntPwdHistory", (value))
}

// GetDS_ntPwdHistory gets the value of DS_ntPwdHistory for the instance
func (instance *ads_user) GetPropertyDS_ntPwdHistory() (value []Uint8Array, err error) {
	retValue, err := instance.GetProperty("DS_ntPwdHistory")
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
func (instance *ads_user) SetPropertyDS_objectSid(value Uint8Array) (err error) {
	return instance.SetProperty("DS_objectSid", (value))
}

// GetDS_objectSid gets the value of DS_objectSid for the instance
func (instance *ads_user) GetPropertyDS_objectSid() (value Uint8Array, err error) {
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
func (instance *ads_user) SetPropertyDS_operatorCount(value int32) (err error) {
	return instance.SetProperty("DS_operatorCount", (value))
}

// GetDS_operatorCount gets the value of DS_operatorCount for the instance
func (instance *ads_user) GetPropertyDS_operatorCount() (value int32, err error) {
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

// SetDS_otherLoginWorkstations sets the value of DS_otherLoginWorkstations for the instance
func (instance *ads_user) SetPropertyDS_otherLoginWorkstations(value []string) (err error) {
	return instance.SetProperty("DS_otherLoginWorkstations", (value))
}

// GetDS_otherLoginWorkstations gets the value of DS_otherLoginWorkstations for the instance
func (instance *ads_user) GetPropertyDS_otherLoginWorkstations() (value []string, err error) {
	retValue, err := instance.GetProperty("DS_otherLoginWorkstations")
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

// SetDS_photo sets the value of DS_photo for the instance
func (instance *ads_user) SetPropertyDS_photo(value []Uint8Array) (err error) {
	return instance.SetProperty("DS_photo", (value))
}

// GetDS_photo gets the value of DS_photo for the instance
func (instance *ads_user) GetPropertyDS_photo() (value []Uint8Array, err error) {
	retValue, err := instance.GetProperty("DS_photo")
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

// SetDS_preferredLanguage sets the value of DS_preferredLanguage for the instance
func (instance *ads_user) SetPropertyDS_preferredLanguage(value string) (err error) {
	return instance.SetProperty("DS_preferredLanguage", (value))
}

// GetDS_preferredLanguage gets the value of DS_preferredLanguage for the instance
func (instance *ads_user) GetPropertyDS_preferredLanguage() (value string, err error) {
	retValue, err := instance.GetProperty("DS_preferredLanguage")
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

// SetDS_preferredOU sets the value of DS_preferredOU for the instance
func (instance *ads_user) SetPropertyDS_preferredOU(value string) (err error) {
	return instance.SetProperty("DS_preferredOU", (value))
}

// GetDS_preferredOU gets the value of DS_preferredOU for the instance
func (instance *ads_user) GetPropertyDS_preferredOU() (value string, err error) {
	retValue, err := instance.GetProperty("DS_preferredOU")
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

// SetDS_primaryGroupID sets the value of DS_primaryGroupID for the instance
func (instance *ads_user) SetPropertyDS_primaryGroupID(value int32) (err error) {
	return instance.SetProperty("DS_primaryGroupID", (value))
}

// GetDS_primaryGroupID gets the value of DS_primaryGroupID for the instance
func (instance *ads_user) GetPropertyDS_primaryGroupID() (value int32, err error) {
	retValue, err := instance.GetProperty("DS_primaryGroupID")
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

// SetDS_profilePath sets the value of DS_profilePath for the instance
func (instance *ads_user) SetPropertyDS_profilePath(value string) (err error) {
	return instance.SetProperty("DS_profilePath", (value))
}

// GetDS_profilePath gets the value of DS_profilePath for the instance
func (instance *ads_user) GetPropertyDS_profilePath() (value string, err error) {
	retValue, err := instance.GetProperty("DS_profilePath")
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

// SetDS_pwdLastSet sets the value of DS_pwdLastSet for the instance
func (instance *ads_user) SetPropertyDS_pwdLastSet(value int64) (err error) {
	return instance.SetProperty("DS_pwdLastSet", (value))
}

// GetDS_pwdLastSet gets the value of DS_pwdLastSet for the instance
func (instance *ads_user) GetPropertyDS_pwdLastSet() (value int64, err error) {
	retValue, err := instance.GetProperty("DS_pwdLastSet")
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

// SetDS_rid sets the value of DS_rid for the instance
func (instance *ads_user) SetPropertyDS_rid(value int32) (err error) {
	return instance.SetProperty("DS_rid", (value))
}

// GetDS_rid gets the value of DS_rid for the instance
func (instance *ads_user) GetPropertyDS_rid() (value int32, err error) {
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

// SetDS_roomNumber sets the value of DS_roomNumber for the instance
func (instance *ads_user) SetPropertyDS_roomNumber(value []string) (err error) {
	return instance.SetProperty("DS_roomNumber", (value))
}

// GetDS_roomNumber gets the value of DS_roomNumber for the instance
func (instance *ads_user) GetPropertyDS_roomNumber() (value []string, err error) {
	retValue, err := instance.GetProperty("DS_roomNumber")
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

// SetDS_sAMAccountName sets the value of DS_sAMAccountName for the instance
func (instance *ads_user) SetPropertyDS_sAMAccountName(value string) (err error) {
	return instance.SetProperty("DS_sAMAccountName", (value))
}

// GetDS_sAMAccountName gets the value of DS_sAMAccountName for the instance
func (instance *ads_user) GetPropertyDS_sAMAccountName() (value string, err error) {
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
func (instance *ads_user) SetPropertyDS_sAMAccountType(value int32) (err error) {
	return instance.SetProperty("DS_sAMAccountType", (value))
}

// GetDS_sAMAccountType gets the value of DS_sAMAccountType for the instance
func (instance *ads_user) GetPropertyDS_sAMAccountType() (value int32, err error) {
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

// SetDS_scriptPath sets the value of DS_scriptPath for the instance
func (instance *ads_user) SetPropertyDS_scriptPath(value string) (err error) {
	return instance.SetProperty("DS_scriptPath", (value))
}

// GetDS_scriptPath gets the value of DS_scriptPath for the instance
func (instance *ads_user) GetPropertyDS_scriptPath() (value string, err error) {
	retValue, err := instance.GetProperty("DS_scriptPath")
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

// SetDS_secretary sets the value of DS_secretary for the instance
func (instance *ads_user) SetPropertyDS_secretary(value []string) (err error) {
	return instance.SetProperty("DS_secretary", (value))
}

// GetDS_secretary gets the value of DS_secretary for the instance
func (instance *ads_user) GetPropertyDS_secretary() (value []string, err error) {
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
func (instance *ads_user) SetPropertyDS_securityIdentifier(value Uint8Array) (err error) {
	return instance.SetProperty("DS_securityIdentifier", (value))
}

// GetDS_securityIdentifier gets the value of DS_securityIdentifier for the instance
func (instance *ads_user) GetPropertyDS_securityIdentifier() (value Uint8Array, err error) {
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

// SetDS_servicePrincipalName sets the value of DS_servicePrincipalName for the instance
func (instance *ads_user) SetPropertyDS_servicePrincipalName(value []string) (err error) {
	return instance.SetProperty("DS_servicePrincipalName", (value))
}

// GetDS_servicePrincipalName gets the value of DS_servicePrincipalName for the instance
func (instance *ads_user) GetPropertyDS_servicePrincipalName() (value []string, err error) {
	retValue, err := instance.GetProperty("DS_servicePrincipalName")
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

// SetDS_shadowExpire sets the value of DS_shadowExpire for the instance
func (instance *ads_user) SetPropertyDS_shadowExpire(value int32) (err error) {
	return instance.SetProperty("DS_shadowExpire", (value))
}

// GetDS_shadowExpire gets the value of DS_shadowExpire for the instance
func (instance *ads_user) GetPropertyDS_shadowExpire() (value int32, err error) {
	retValue, err := instance.GetProperty("DS_shadowExpire")
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

// SetDS_shadowFlag sets the value of DS_shadowFlag for the instance
func (instance *ads_user) SetPropertyDS_shadowFlag(value int32) (err error) {
	return instance.SetProperty("DS_shadowFlag", (value))
}

// GetDS_shadowFlag gets the value of DS_shadowFlag for the instance
func (instance *ads_user) GetPropertyDS_shadowFlag() (value int32, err error) {
	retValue, err := instance.GetProperty("DS_shadowFlag")
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

// SetDS_shadowInactive sets the value of DS_shadowInactive for the instance
func (instance *ads_user) SetPropertyDS_shadowInactive(value int32) (err error) {
	return instance.SetProperty("DS_shadowInactive", (value))
}

// GetDS_shadowInactive gets the value of DS_shadowInactive for the instance
func (instance *ads_user) GetPropertyDS_shadowInactive() (value int32, err error) {
	retValue, err := instance.GetProperty("DS_shadowInactive")
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

// SetDS_shadowLastChange sets the value of DS_shadowLastChange for the instance
func (instance *ads_user) SetPropertyDS_shadowLastChange(value int32) (err error) {
	return instance.SetProperty("DS_shadowLastChange", (value))
}

// GetDS_shadowLastChange gets the value of DS_shadowLastChange for the instance
func (instance *ads_user) GetPropertyDS_shadowLastChange() (value int32, err error) {
	retValue, err := instance.GetProperty("DS_shadowLastChange")
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

// SetDS_shadowMax sets the value of DS_shadowMax for the instance
func (instance *ads_user) SetPropertyDS_shadowMax(value int32) (err error) {
	return instance.SetProperty("DS_shadowMax", (value))
}

// GetDS_shadowMax gets the value of DS_shadowMax for the instance
func (instance *ads_user) GetPropertyDS_shadowMax() (value int32, err error) {
	retValue, err := instance.GetProperty("DS_shadowMax")
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

// SetDS_shadowMin sets the value of DS_shadowMin for the instance
func (instance *ads_user) SetPropertyDS_shadowMin(value int32) (err error) {
	return instance.SetProperty("DS_shadowMin", (value))
}

// GetDS_shadowMin gets the value of DS_shadowMin for the instance
func (instance *ads_user) GetPropertyDS_shadowMin() (value int32, err error) {
	retValue, err := instance.GetProperty("DS_shadowMin")
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

// SetDS_shadowWarning sets the value of DS_shadowWarning for the instance
func (instance *ads_user) SetPropertyDS_shadowWarning(value int32) (err error) {
	return instance.SetProperty("DS_shadowWarning", (value))
}

// GetDS_shadowWarning gets the value of DS_shadowWarning for the instance
func (instance *ads_user) GetPropertyDS_shadowWarning() (value int32, err error) {
	retValue, err := instance.GetProperty("DS_shadowWarning")
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

// SetDS_showInAddressBook sets the value of DS_showInAddressBook for the instance
func (instance *ads_user) SetPropertyDS_showInAddressBook(value []string) (err error) {
	return instance.SetProperty("DS_showInAddressBook", (value))
}

// GetDS_showInAddressBook gets the value of DS_showInAddressBook for the instance
func (instance *ads_user) GetPropertyDS_showInAddressBook() (value []string, err error) {
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
func (instance *ads_user) SetPropertyDS_sIDHistory(value []Uint8Array) (err error) {
	return instance.SetProperty("DS_sIDHistory", (value))
}

// GetDS_sIDHistory gets the value of DS_sIDHistory for the instance
func (instance *ads_user) GetPropertyDS_sIDHistory() (value []Uint8Array, err error) {
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
func (instance *ads_user) SetPropertyDS_supplementalCredentials(value []Uint8Array) (err error) {
	return instance.SetProperty("DS_supplementalCredentials", (value))
}

// GetDS_supplementalCredentials gets the value of DS_supplementalCredentials for the instance
func (instance *ads_user) GetPropertyDS_supplementalCredentials() (value []Uint8Array, err error) {
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

// SetDS_terminalServer sets the value of DS_terminalServer for the instance
func (instance *ads_user) SetPropertyDS_terminalServer(value Uint8Array) (err error) {
	return instance.SetProperty("DS_terminalServer", (value))
}

// GetDS_terminalServer gets the value of DS_terminalServer for the instance
func (instance *ads_user) GetPropertyDS_terminalServer() (value Uint8Array, err error) {
	retValue, err := instance.GetProperty("DS_terminalServer")
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

// SetDS_textEncodedORAddress sets the value of DS_textEncodedORAddress for the instance
func (instance *ads_user) SetPropertyDS_textEncodedORAddress(value string) (err error) {
	return instance.SetProperty("DS_textEncodedORAddress", (value))
}

// GetDS_textEncodedORAddress gets the value of DS_textEncodedORAddress for the instance
func (instance *ads_user) GetPropertyDS_textEncodedORAddress() (value string, err error) {
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
func (instance *ads_user) SetPropertyDS_tokenGroups(value []Uint8Array) (err error) {
	return instance.SetProperty("DS_tokenGroups", (value))
}

// GetDS_tokenGroups gets the value of DS_tokenGroups for the instance
func (instance *ads_user) GetPropertyDS_tokenGroups() (value []Uint8Array, err error) {
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
func (instance *ads_user) SetPropertyDS_tokenGroupsGlobalAndUniversal(value []Uint8Array) (err error) {
	return instance.SetProperty("DS_tokenGroupsGlobalAndUniversal", (value))
}

// GetDS_tokenGroupsGlobalAndUniversal gets the value of DS_tokenGroupsGlobalAndUniversal for the instance
func (instance *ads_user) GetPropertyDS_tokenGroupsGlobalAndUniversal() (value []Uint8Array, err error) {
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
func (instance *ads_user) SetPropertyDS_tokenGroupsNoGCAcceptable(value []Uint8Array) (err error) {
	return instance.SetProperty("DS_tokenGroupsNoGCAcceptable", (value))
}

// GetDS_tokenGroupsNoGCAcceptable gets the value of DS_tokenGroupsNoGCAcceptable for the instance
func (instance *ads_user) GetPropertyDS_tokenGroupsNoGCAcceptable() (value []Uint8Array, err error) {
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

// SetDS_uid sets the value of DS_uid for the instance
func (instance *ads_user) SetPropertyDS_uid(value []string) (err error) {
	return instance.SetProperty("DS_uid", (value))
}

// GetDS_uid gets the value of DS_uid for the instance
func (instance *ads_user) GetPropertyDS_uid() (value []string, err error) {
	retValue, err := instance.GetProperty("DS_uid")
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

// SetDS_uidNumber sets the value of DS_uidNumber for the instance
func (instance *ads_user) SetPropertyDS_uidNumber(value int32) (err error) {
	return instance.SetProperty("DS_uidNumber", (value))
}

// GetDS_uidNumber gets the value of DS_uidNumber for the instance
func (instance *ads_user) GetPropertyDS_uidNumber() (value int32, err error) {
	retValue, err := instance.GetProperty("DS_uidNumber")
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

// SetDS_unicodePwd sets the value of DS_unicodePwd for the instance
func (instance *ads_user) SetPropertyDS_unicodePwd(value Uint8Array) (err error) {
	return instance.SetProperty("DS_unicodePwd", (value))
}

// GetDS_unicodePwd gets the value of DS_unicodePwd for the instance
func (instance *ads_user) GetPropertyDS_unicodePwd() (value Uint8Array, err error) {
	retValue, err := instance.GetProperty("DS_unicodePwd")
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

// SetDS_unixHomeDirectory sets the value of DS_unixHomeDirectory for the instance
func (instance *ads_user) SetPropertyDS_unixHomeDirectory(value string) (err error) {
	return instance.SetProperty("DS_unixHomeDirectory", (value))
}

// GetDS_unixHomeDirectory gets the value of DS_unixHomeDirectory for the instance
func (instance *ads_user) GetPropertyDS_unixHomeDirectory() (value string, err error) {
	retValue, err := instance.GetProperty("DS_unixHomeDirectory")
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

// SetDS_unixUserPassword sets the value of DS_unixUserPassword for the instance
func (instance *ads_user) SetPropertyDS_unixUserPassword(value []Uint8Array) (err error) {
	return instance.SetProperty("DS_unixUserPassword", (value))
}

// GetDS_unixUserPassword gets the value of DS_unixUserPassword for the instance
func (instance *ads_user) GetPropertyDS_unixUserPassword() (value []Uint8Array, err error) {
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

// SetDS_userAccountControl sets the value of DS_userAccountControl for the instance
func (instance *ads_user) SetPropertyDS_userAccountControl(value int32) (err error) {
	return instance.SetProperty("DS_userAccountControl", (value))
}

// GetDS_userAccountControl gets the value of DS_userAccountControl for the instance
func (instance *ads_user) GetPropertyDS_userAccountControl() (value int32, err error) {
	retValue, err := instance.GetProperty("DS_userAccountControl")
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

// SetDS_userCert sets the value of DS_userCert for the instance
func (instance *ads_user) SetPropertyDS_userCert(value Uint8Array) (err error) {
	return instance.SetProperty("DS_userCert", (value))
}

// GetDS_userCert gets the value of DS_userCert for the instance
func (instance *ads_user) GetPropertyDS_userCert() (value Uint8Array, err error) {
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
func (instance *ads_user) SetPropertyDS_userCertificate(value []Uint8Array) (err error) {
	return instance.SetProperty("DS_userCertificate", (value))
}

// GetDS_userCertificate gets the value of DS_userCertificate for the instance
func (instance *ads_user) GetPropertyDS_userCertificate() (value []Uint8Array, err error) {
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

// SetDS_userParameters sets the value of DS_userParameters for the instance
func (instance *ads_user) SetPropertyDS_userParameters(value string) (err error) {
	return instance.SetProperty("DS_userParameters", (value))
}

// GetDS_userParameters gets the value of DS_userParameters for the instance
func (instance *ads_user) GetPropertyDS_userParameters() (value string, err error) {
	retValue, err := instance.GetProperty("DS_userParameters")
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

// SetDS_userPKCS12 sets the value of DS_userPKCS12 for the instance
func (instance *ads_user) SetPropertyDS_userPKCS12(value []Uint8Array) (err error) {
	return instance.SetProperty("DS_userPKCS12", (value))
}

// GetDS_userPKCS12 gets the value of DS_userPKCS12 for the instance
func (instance *ads_user) GetPropertyDS_userPKCS12() (value []Uint8Array, err error) {
	retValue, err := instance.GetProperty("DS_userPKCS12")
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

// SetDS_userPrincipalName sets the value of DS_userPrincipalName for the instance
func (instance *ads_user) SetPropertyDS_userPrincipalName(value string) (err error) {
	return instance.SetProperty("DS_userPrincipalName", (value))
}

// GetDS_userPrincipalName gets the value of DS_userPrincipalName for the instance
func (instance *ads_user) GetPropertyDS_userPrincipalName() (value string, err error) {
	retValue, err := instance.GetProperty("DS_userPrincipalName")
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

// SetDS_userSharedFolder sets the value of DS_userSharedFolder for the instance
func (instance *ads_user) SetPropertyDS_userSharedFolder(value string) (err error) {
	return instance.SetProperty("DS_userSharedFolder", (value))
}

// GetDS_userSharedFolder gets the value of DS_userSharedFolder for the instance
func (instance *ads_user) GetPropertyDS_userSharedFolder() (value string, err error) {
	retValue, err := instance.GetProperty("DS_userSharedFolder")
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

// SetDS_userSharedFolderOther sets the value of DS_userSharedFolderOther for the instance
func (instance *ads_user) SetPropertyDS_userSharedFolderOther(value []string) (err error) {
	return instance.SetProperty("DS_userSharedFolderOther", (value))
}

// GetDS_userSharedFolderOther gets the value of DS_userSharedFolderOther for the instance
func (instance *ads_user) GetPropertyDS_userSharedFolderOther() (value []string, err error) {
	retValue, err := instance.GetProperty("DS_userSharedFolderOther")
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

// SetDS_userSMIMECertificate sets the value of DS_userSMIMECertificate for the instance
func (instance *ads_user) SetPropertyDS_userSMIMECertificate(value []Uint8Array) (err error) {
	return instance.SetProperty("DS_userSMIMECertificate", (value))
}

// GetDS_userSMIMECertificate gets the value of DS_userSMIMECertificate for the instance
func (instance *ads_user) GetPropertyDS_userSMIMECertificate() (value []Uint8Array, err error) {
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

// SetDS_userWorkstations sets the value of DS_userWorkstations for the instance
func (instance *ads_user) SetPropertyDS_userWorkstations(value string) (err error) {
	return instance.SetProperty("DS_userWorkstations", (value))
}

// GetDS_userWorkstations gets the value of DS_userWorkstations for the instance
func (instance *ads_user) GetPropertyDS_userWorkstations() (value string, err error) {
	retValue, err := instance.GetProperty("DS_userWorkstations")
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

// SetDS_x500uniqueIdentifier sets the value of DS_x500uniqueIdentifier for the instance
func (instance *ads_user) SetPropertyDS_x500uniqueIdentifier(value []Uint8Array) (err error) {
	return instance.SetProperty("DS_x500uniqueIdentifier", (value))
}

// GetDS_x500uniqueIdentifier gets the value of DS_x500uniqueIdentifier for the instance
func (instance *ads_user) GetPropertyDS_x500uniqueIdentifier() (value []Uint8Array, err error) {
	retValue, err := instance.GetProperty("DS_x500uniqueIdentifier")
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
