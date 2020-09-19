// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 9/18/2020 using wmigen
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

// MDM_Policy_Config01_Audit02 struct
type MDM_Policy_Config01_Audit02 struct {
	*cim.WmiInstance

	//
	AccountLogon_AuditCredentialValidation int32

	//
	AccountLogon_AuditKerberosAuthenticationService int32

	//
	AccountLogon_AuditKerberosServiceTicketOperations int32

	//
	AccountLogon_AuditOtherAccountLogonEvents int32

	//
	AccountLogonLogoff_AuditAccountLockout int32

	//
	AccountLogonLogoff_AuditGroupMembership int32

	//
	AccountLogonLogoff_AuditIPsecExtendedMode int32

	//
	AccountLogonLogoff_AuditIPsecMainMode int32

	//
	AccountLogonLogoff_AuditIPsecQuickMode int32

	//
	AccountLogonLogoff_AuditLogoff int32

	//
	AccountLogonLogoff_AuditLogon int32

	//
	AccountLogonLogoff_AuditNetworkPolicyServer int32

	//
	AccountLogonLogoff_AuditOtherLogonLogoffEvents int32

	//
	AccountLogonLogoff_AuditSpecialLogon int32

	//
	AccountLogonLogoff_AuditUserDeviceClaims int32

	//
	AccountManagement_AuditApplicationGroupManagement int32

	//
	AccountManagement_AuditComputerAccountManagement int32

	//
	AccountManagement_AuditDistributionGroupManagement int32

	//
	AccountManagement_AuditOtherAccountManagementEvents int32

	//
	AccountManagement_AuditSecurityGroupManagement int32

	//
	AccountManagement_AuditUserAccountManagement int32

	//
	DetailedTracking_AuditDPAPIActivity int32

	//
	DetailedTracking_AuditPNPActivity int32

	//
	DetailedTracking_AuditProcessCreation int32

	//
	DetailedTracking_AuditProcessTermination int32

	//
	DetailedTracking_AuditRPCEvents int32

	//
	DetailedTracking_AuditTokenRightAdjusted int32

	//
	DSAccess_AuditDetailedDirectoryServiceReplication int32

	//
	DSAccess_AuditDirectoryServiceAccess int32

	//
	DSAccess_AuditDirectoryServiceChanges int32

	//
	DSAccess_AuditDirectoryServiceReplication int32

	//
	InstanceID string

	//
	ObjectAccess_AuditApplicationGenerated int32

	//
	ObjectAccess_AuditCentralAccessPolicyStaging int32

	//
	ObjectAccess_AuditCertificationServices int32

	//
	ObjectAccess_AuditDetailedFileShare int32

	//
	ObjectAccess_AuditFileShare int32

	//
	ObjectAccess_AuditFileSystem int32

	//
	ObjectAccess_AuditFilteringPlatformConnection int32

	//
	ObjectAccess_AuditFilteringPlatformPacketDrop int32

	//
	ObjectAccess_AuditHandleManipulation int32

	//
	ObjectAccess_AuditKernelObject int32

	//
	ObjectAccess_AuditOtherObjectAccessEvents int32

	//
	ObjectAccess_AuditRegistry int32

	//
	ObjectAccess_AuditRemovableStorage int32

	//
	ObjectAccess_AuditSAM int32

	//
	ParentID string

	//
	PolicyChange_AuditAuthenticationPolicyChange int32

	//
	PolicyChange_AuditAuthorizationPolicyChange int32

	//
	PolicyChange_AuditFilteringPlatformPolicyChange int32

	//
	PolicyChange_AuditMPSSVCRuleLevelPolicyChange int32

	//
	PolicyChange_AuditOtherPolicyChangeEvents int32

	//
	PolicyChange_AuditPolicyChange int32

	//
	PrivilegeUse_AuditNonSensitivePrivilegeUse int32

	//
	PrivilegeUse_AuditOtherPrivilegeUseEvents int32

	//
	PrivilegeUse_AuditSensitivePrivilegeUse int32

	//
	System_AuditIPsecDriver int32

	//
	System_AuditOtherSystemEvents int32

	//
	System_AuditSecurityStateChange int32

	//
	System_AuditSecuritySystemExtension int32

	//
	System_AuditSystemIntegrity int32
}

func NewMDM_Policy_Config01_Audit02Ex1(instance *cim.WmiInstance) (newInstance *MDM_Policy_Config01_Audit02, err error) {
	tmp, err := instance, nil

	if err != nil {
		return
	}
	newInstance = &MDM_Policy_Config01_Audit02{
		WmiInstance: tmp,
	}
	return
}

func NewMDM_Policy_Config01_Audit02Ex6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *MDM_Policy_Config01_Audit02, err error) {
	tmp, err := instance.GetWmiInstance(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &MDM_Policy_Config01_Audit02{
		WmiInstance: tmp,
	}
	return
}

// SetAccountLogon_AuditCredentialValidation sets the value of AccountLogon_AuditCredentialValidation for the instance
func (instance *MDM_Policy_Config01_Audit02) SetPropertyAccountLogon_AuditCredentialValidation(value int32) (err error) {
	return instance.SetProperty("AccountLogon_AuditCredentialValidation", (value))
}

// GetAccountLogon_AuditCredentialValidation gets the value of AccountLogon_AuditCredentialValidation for the instance
func (instance *MDM_Policy_Config01_Audit02) GetPropertyAccountLogon_AuditCredentialValidation() (value int32, err error) {
	retValue, err := instance.GetProperty("AccountLogon_AuditCredentialValidation")
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

// SetAccountLogon_AuditKerberosAuthenticationService sets the value of AccountLogon_AuditKerberosAuthenticationService for the instance
func (instance *MDM_Policy_Config01_Audit02) SetPropertyAccountLogon_AuditKerberosAuthenticationService(value int32) (err error) {
	return instance.SetProperty("AccountLogon_AuditKerberosAuthenticationService", (value))
}

// GetAccountLogon_AuditKerberosAuthenticationService gets the value of AccountLogon_AuditKerberosAuthenticationService for the instance
func (instance *MDM_Policy_Config01_Audit02) GetPropertyAccountLogon_AuditKerberosAuthenticationService() (value int32, err error) {
	retValue, err := instance.GetProperty("AccountLogon_AuditKerberosAuthenticationService")
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

// SetAccountLogon_AuditKerberosServiceTicketOperations sets the value of AccountLogon_AuditKerberosServiceTicketOperations for the instance
func (instance *MDM_Policy_Config01_Audit02) SetPropertyAccountLogon_AuditKerberosServiceTicketOperations(value int32) (err error) {
	return instance.SetProperty("AccountLogon_AuditKerberosServiceTicketOperations", (value))
}

// GetAccountLogon_AuditKerberosServiceTicketOperations gets the value of AccountLogon_AuditKerberosServiceTicketOperations for the instance
func (instance *MDM_Policy_Config01_Audit02) GetPropertyAccountLogon_AuditKerberosServiceTicketOperations() (value int32, err error) {
	retValue, err := instance.GetProperty("AccountLogon_AuditKerberosServiceTicketOperations")
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

// SetAccountLogon_AuditOtherAccountLogonEvents sets the value of AccountLogon_AuditOtherAccountLogonEvents for the instance
func (instance *MDM_Policy_Config01_Audit02) SetPropertyAccountLogon_AuditOtherAccountLogonEvents(value int32) (err error) {
	return instance.SetProperty("AccountLogon_AuditOtherAccountLogonEvents", (value))
}

// GetAccountLogon_AuditOtherAccountLogonEvents gets the value of AccountLogon_AuditOtherAccountLogonEvents for the instance
func (instance *MDM_Policy_Config01_Audit02) GetPropertyAccountLogon_AuditOtherAccountLogonEvents() (value int32, err error) {
	retValue, err := instance.GetProperty("AccountLogon_AuditOtherAccountLogonEvents")
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

// SetAccountLogonLogoff_AuditAccountLockout sets the value of AccountLogonLogoff_AuditAccountLockout for the instance
func (instance *MDM_Policy_Config01_Audit02) SetPropertyAccountLogonLogoff_AuditAccountLockout(value int32) (err error) {
	return instance.SetProperty("AccountLogonLogoff_AuditAccountLockout", (value))
}

// GetAccountLogonLogoff_AuditAccountLockout gets the value of AccountLogonLogoff_AuditAccountLockout for the instance
func (instance *MDM_Policy_Config01_Audit02) GetPropertyAccountLogonLogoff_AuditAccountLockout() (value int32, err error) {
	retValue, err := instance.GetProperty("AccountLogonLogoff_AuditAccountLockout")
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

// SetAccountLogonLogoff_AuditGroupMembership sets the value of AccountLogonLogoff_AuditGroupMembership for the instance
func (instance *MDM_Policy_Config01_Audit02) SetPropertyAccountLogonLogoff_AuditGroupMembership(value int32) (err error) {
	return instance.SetProperty("AccountLogonLogoff_AuditGroupMembership", (value))
}

// GetAccountLogonLogoff_AuditGroupMembership gets the value of AccountLogonLogoff_AuditGroupMembership for the instance
func (instance *MDM_Policy_Config01_Audit02) GetPropertyAccountLogonLogoff_AuditGroupMembership() (value int32, err error) {
	retValue, err := instance.GetProperty("AccountLogonLogoff_AuditGroupMembership")
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

// SetAccountLogonLogoff_AuditIPsecExtendedMode sets the value of AccountLogonLogoff_AuditIPsecExtendedMode for the instance
func (instance *MDM_Policy_Config01_Audit02) SetPropertyAccountLogonLogoff_AuditIPsecExtendedMode(value int32) (err error) {
	return instance.SetProperty("AccountLogonLogoff_AuditIPsecExtendedMode", (value))
}

// GetAccountLogonLogoff_AuditIPsecExtendedMode gets the value of AccountLogonLogoff_AuditIPsecExtendedMode for the instance
func (instance *MDM_Policy_Config01_Audit02) GetPropertyAccountLogonLogoff_AuditIPsecExtendedMode() (value int32, err error) {
	retValue, err := instance.GetProperty("AccountLogonLogoff_AuditIPsecExtendedMode")
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

// SetAccountLogonLogoff_AuditIPsecMainMode sets the value of AccountLogonLogoff_AuditIPsecMainMode for the instance
func (instance *MDM_Policy_Config01_Audit02) SetPropertyAccountLogonLogoff_AuditIPsecMainMode(value int32) (err error) {
	return instance.SetProperty("AccountLogonLogoff_AuditIPsecMainMode", (value))
}

// GetAccountLogonLogoff_AuditIPsecMainMode gets the value of AccountLogonLogoff_AuditIPsecMainMode for the instance
func (instance *MDM_Policy_Config01_Audit02) GetPropertyAccountLogonLogoff_AuditIPsecMainMode() (value int32, err error) {
	retValue, err := instance.GetProperty("AccountLogonLogoff_AuditIPsecMainMode")
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

// SetAccountLogonLogoff_AuditIPsecQuickMode sets the value of AccountLogonLogoff_AuditIPsecQuickMode for the instance
func (instance *MDM_Policy_Config01_Audit02) SetPropertyAccountLogonLogoff_AuditIPsecQuickMode(value int32) (err error) {
	return instance.SetProperty("AccountLogonLogoff_AuditIPsecQuickMode", (value))
}

// GetAccountLogonLogoff_AuditIPsecQuickMode gets the value of AccountLogonLogoff_AuditIPsecQuickMode for the instance
func (instance *MDM_Policy_Config01_Audit02) GetPropertyAccountLogonLogoff_AuditIPsecQuickMode() (value int32, err error) {
	retValue, err := instance.GetProperty("AccountLogonLogoff_AuditIPsecQuickMode")
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

// SetAccountLogonLogoff_AuditLogoff sets the value of AccountLogonLogoff_AuditLogoff for the instance
func (instance *MDM_Policy_Config01_Audit02) SetPropertyAccountLogonLogoff_AuditLogoff(value int32) (err error) {
	return instance.SetProperty("AccountLogonLogoff_AuditLogoff", (value))
}

// GetAccountLogonLogoff_AuditLogoff gets the value of AccountLogonLogoff_AuditLogoff for the instance
func (instance *MDM_Policy_Config01_Audit02) GetPropertyAccountLogonLogoff_AuditLogoff() (value int32, err error) {
	retValue, err := instance.GetProperty("AccountLogonLogoff_AuditLogoff")
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

// SetAccountLogonLogoff_AuditLogon sets the value of AccountLogonLogoff_AuditLogon for the instance
func (instance *MDM_Policy_Config01_Audit02) SetPropertyAccountLogonLogoff_AuditLogon(value int32) (err error) {
	return instance.SetProperty("AccountLogonLogoff_AuditLogon", (value))
}

// GetAccountLogonLogoff_AuditLogon gets the value of AccountLogonLogoff_AuditLogon for the instance
func (instance *MDM_Policy_Config01_Audit02) GetPropertyAccountLogonLogoff_AuditLogon() (value int32, err error) {
	retValue, err := instance.GetProperty("AccountLogonLogoff_AuditLogon")
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

// SetAccountLogonLogoff_AuditNetworkPolicyServer sets the value of AccountLogonLogoff_AuditNetworkPolicyServer for the instance
func (instance *MDM_Policy_Config01_Audit02) SetPropertyAccountLogonLogoff_AuditNetworkPolicyServer(value int32) (err error) {
	return instance.SetProperty("AccountLogonLogoff_AuditNetworkPolicyServer", (value))
}

// GetAccountLogonLogoff_AuditNetworkPolicyServer gets the value of AccountLogonLogoff_AuditNetworkPolicyServer for the instance
func (instance *MDM_Policy_Config01_Audit02) GetPropertyAccountLogonLogoff_AuditNetworkPolicyServer() (value int32, err error) {
	retValue, err := instance.GetProperty("AccountLogonLogoff_AuditNetworkPolicyServer")
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

// SetAccountLogonLogoff_AuditOtherLogonLogoffEvents sets the value of AccountLogonLogoff_AuditOtherLogonLogoffEvents for the instance
func (instance *MDM_Policy_Config01_Audit02) SetPropertyAccountLogonLogoff_AuditOtherLogonLogoffEvents(value int32) (err error) {
	return instance.SetProperty("AccountLogonLogoff_AuditOtherLogonLogoffEvents", (value))
}

// GetAccountLogonLogoff_AuditOtherLogonLogoffEvents gets the value of AccountLogonLogoff_AuditOtherLogonLogoffEvents for the instance
func (instance *MDM_Policy_Config01_Audit02) GetPropertyAccountLogonLogoff_AuditOtherLogonLogoffEvents() (value int32, err error) {
	retValue, err := instance.GetProperty("AccountLogonLogoff_AuditOtherLogonLogoffEvents")
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

// SetAccountLogonLogoff_AuditSpecialLogon sets the value of AccountLogonLogoff_AuditSpecialLogon for the instance
func (instance *MDM_Policy_Config01_Audit02) SetPropertyAccountLogonLogoff_AuditSpecialLogon(value int32) (err error) {
	return instance.SetProperty("AccountLogonLogoff_AuditSpecialLogon", (value))
}

// GetAccountLogonLogoff_AuditSpecialLogon gets the value of AccountLogonLogoff_AuditSpecialLogon for the instance
func (instance *MDM_Policy_Config01_Audit02) GetPropertyAccountLogonLogoff_AuditSpecialLogon() (value int32, err error) {
	retValue, err := instance.GetProperty("AccountLogonLogoff_AuditSpecialLogon")
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

// SetAccountLogonLogoff_AuditUserDeviceClaims sets the value of AccountLogonLogoff_AuditUserDeviceClaims for the instance
func (instance *MDM_Policy_Config01_Audit02) SetPropertyAccountLogonLogoff_AuditUserDeviceClaims(value int32) (err error) {
	return instance.SetProperty("AccountLogonLogoff_AuditUserDeviceClaims", (value))
}

// GetAccountLogonLogoff_AuditUserDeviceClaims gets the value of AccountLogonLogoff_AuditUserDeviceClaims for the instance
func (instance *MDM_Policy_Config01_Audit02) GetPropertyAccountLogonLogoff_AuditUserDeviceClaims() (value int32, err error) {
	retValue, err := instance.GetProperty("AccountLogonLogoff_AuditUserDeviceClaims")
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

// SetAccountManagement_AuditApplicationGroupManagement sets the value of AccountManagement_AuditApplicationGroupManagement for the instance
func (instance *MDM_Policy_Config01_Audit02) SetPropertyAccountManagement_AuditApplicationGroupManagement(value int32) (err error) {
	return instance.SetProperty("AccountManagement_AuditApplicationGroupManagement", (value))
}

// GetAccountManagement_AuditApplicationGroupManagement gets the value of AccountManagement_AuditApplicationGroupManagement for the instance
func (instance *MDM_Policy_Config01_Audit02) GetPropertyAccountManagement_AuditApplicationGroupManagement() (value int32, err error) {
	retValue, err := instance.GetProperty("AccountManagement_AuditApplicationGroupManagement")
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

// SetAccountManagement_AuditComputerAccountManagement sets the value of AccountManagement_AuditComputerAccountManagement for the instance
func (instance *MDM_Policy_Config01_Audit02) SetPropertyAccountManagement_AuditComputerAccountManagement(value int32) (err error) {
	return instance.SetProperty("AccountManagement_AuditComputerAccountManagement", (value))
}

// GetAccountManagement_AuditComputerAccountManagement gets the value of AccountManagement_AuditComputerAccountManagement for the instance
func (instance *MDM_Policy_Config01_Audit02) GetPropertyAccountManagement_AuditComputerAccountManagement() (value int32, err error) {
	retValue, err := instance.GetProperty("AccountManagement_AuditComputerAccountManagement")
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

// SetAccountManagement_AuditDistributionGroupManagement sets the value of AccountManagement_AuditDistributionGroupManagement for the instance
func (instance *MDM_Policy_Config01_Audit02) SetPropertyAccountManagement_AuditDistributionGroupManagement(value int32) (err error) {
	return instance.SetProperty("AccountManagement_AuditDistributionGroupManagement", (value))
}

// GetAccountManagement_AuditDistributionGroupManagement gets the value of AccountManagement_AuditDistributionGroupManagement for the instance
func (instance *MDM_Policy_Config01_Audit02) GetPropertyAccountManagement_AuditDistributionGroupManagement() (value int32, err error) {
	retValue, err := instance.GetProperty("AccountManagement_AuditDistributionGroupManagement")
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

// SetAccountManagement_AuditOtherAccountManagementEvents sets the value of AccountManagement_AuditOtherAccountManagementEvents for the instance
func (instance *MDM_Policy_Config01_Audit02) SetPropertyAccountManagement_AuditOtherAccountManagementEvents(value int32) (err error) {
	return instance.SetProperty("AccountManagement_AuditOtherAccountManagementEvents", (value))
}

// GetAccountManagement_AuditOtherAccountManagementEvents gets the value of AccountManagement_AuditOtherAccountManagementEvents for the instance
func (instance *MDM_Policy_Config01_Audit02) GetPropertyAccountManagement_AuditOtherAccountManagementEvents() (value int32, err error) {
	retValue, err := instance.GetProperty("AccountManagement_AuditOtherAccountManagementEvents")
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

// SetAccountManagement_AuditSecurityGroupManagement sets the value of AccountManagement_AuditSecurityGroupManagement for the instance
func (instance *MDM_Policy_Config01_Audit02) SetPropertyAccountManagement_AuditSecurityGroupManagement(value int32) (err error) {
	return instance.SetProperty("AccountManagement_AuditSecurityGroupManagement", (value))
}

// GetAccountManagement_AuditSecurityGroupManagement gets the value of AccountManagement_AuditSecurityGroupManagement for the instance
func (instance *MDM_Policy_Config01_Audit02) GetPropertyAccountManagement_AuditSecurityGroupManagement() (value int32, err error) {
	retValue, err := instance.GetProperty("AccountManagement_AuditSecurityGroupManagement")
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

// SetAccountManagement_AuditUserAccountManagement sets the value of AccountManagement_AuditUserAccountManagement for the instance
func (instance *MDM_Policy_Config01_Audit02) SetPropertyAccountManagement_AuditUserAccountManagement(value int32) (err error) {
	return instance.SetProperty("AccountManagement_AuditUserAccountManagement", (value))
}

// GetAccountManagement_AuditUserAccountManagement gets the value of AccountManagement_AuditUserAccountManagement for the instance
func (instance *MDM_Policy_Config01_Audit02) GetPropertyAccountManagement_AuditUserAccountManagement() (value int32, err error) {
	retValue, err := instance.GetProperty("AccountManagement_AuditUserAccountManagement")
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

// SetDetailedTracking_AuditDPAPIActivity sets the value of DetailedTracking_AuditDPAPIActivity for the instance
func (instance *MDM_Policy_Config01_Audit02) SetPropertyDetailedTracking_AuditDPAPIActivity(value int32) (err error) {
	return instance.SetProperty("DetailedTracking_AuditDPAPIActivity", (value))
}

// GetDetailedTracking_AuditDPAPIActivity gets the value of DetailedTracking_AuditDPAPIActivity for the instance
func (instance *MDM_Policy_Config01_Audit02) GetPropertyDetailedTracking_AuditDPAPIActivity() (value int32, err error) {
	retValue, err := instance.GetProperty("DetailedTracking_AuditDPAPIActivity")
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

// SetDetailedTracking_AuditPNPActivity sets the value of DetailedTracking_AuditPNPActivity for the instance
func (instance *MDM_Policy_Config01_Audit02) SetPropertyDetailedTracking_AuditPNPActivity(value int32) (err error) {
	return instance.SetProperty("DetailedTracking_AuditPNPActivity", (value))
}

// GetDetailedTracking_AuditPNPActivity gets the value of DetailedTracking_AuditPNPActivity for the instance
func (instance *MDM_Policy_Config01_Audit02) GetPropertyDetailedTracking_AuditPNPActivity() (value int32, err error) {
	retValue, err := instance.GetProperty("DetailedTracking_AuditPNPActivity")
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

// SetDetailedTracking_AuditProcessCreation sets the value of DetailedTracking_AuditProcessCreation for the instance
func (instance *MDM_Policy_Config01_Audit02) SetPropertyDetailedTracking_AuditProcessCreation(value int32) (err error) {
	return instance.SetProperty("DetailedTracking_AuditProcessCreation", (value))
}

// GetDetailedTracking_AuditProcessCreation gets the value of DetailedTracking_AuditProcessCreation for the instance
func (instance *MDM_Policy_Config01_Audit02) GetPropertyDetailedTracking_AuditProcessCreation() (value int32, err error) {
	retValue, err := instance.GetProperty("DetailedTracking_AuditProcessCreation")
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

// SetDetailedTracking_AuditProcessTermination sets the value of DetailedTracking_AuditProcessTermination for the instance
func (instance *MDM_Policy_Config01_Audit02) SetPropertyDetailedTracking_AuditProcessTermination(value int32) (err error) {
	return instance.SetProperty("DetailedTracking_AuditProcessTermination", (value))
}

// GetDetailedTracking_AuditProcessTermination gets the value of DetailedTracking_AuditProcessTermination for the instance
func (instance *MDM_Policy_Config01_Audit02) GetPropertyDetailedTracking_AuditProcessTermination() (value int32, err error) {
	retValue, err := instance.GetProperty("DetailedTracking_AuditProcessTermination")
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

// SetDetailedTracking_AuditRPCEvents sets the value of DetailedTracking_AuditRPCEvents for the instance
func (instance *MDM_Policy_Config01_Audit02) SetPropertyDetailedTracking_AuditRPCEvents(value int32) (err error) {
	return instance.SetProperty("DetailedTracking_AuditRPCEvents", (value))
}

// GetDetailedTracking_AuditRPCEvents gets the value of DetailedTracking_AuditRPCEvents for the instance
func (instance *MDM_Policy_Config01_Audit02) GetPropertyDetailedTracking_AuditRPCEvents() (value int32, err error) {
	retValue, err := instance.GetProperty("DetailedTracking_AuditRPCEvents")
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

// SetDetailedTracking_AuditTokenRightAdjusted sets the value of DetailedTracking_AuditTokenRightAdjusted for the instance
func (instance *MDM_Policy_Config01_Audit02) SetPropertyDetailedTracking_AuditTokenRightAdjusted(value int32) (err error) {
	return instance.SetProperty("DetailedTracking_AuditTokenRightAdjusted", (value))
}

// GetDetailedTracking_AuditTokenRightAdjusted gets the value of DetailedTracking_AuditTokenRightAdjusted for the instance
func (instance *MDM_Policy_Config01_Audit02) GetPropertyDetailedTracking_AuditTokenRightAdjusted() (value int32, err error) {
	retValue, err := instance.GetProperty("DetailedTracking_AuditTokenRightAdjusted")
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

// SetDSAccess_AuditDetailedDirectoryServiceReplication sets the value of DSAccess_AuditDetailedDirectoryServiceReplication for the instance
func (instance *MDM_Policy_Config01_Audit02) SetPropertyDSAccess_AuditDetailedDirectoryServiceReplication(value int32) (err error) {
	return instance.SetProperty("DSAccess_AuditDetailedDirectoryServiceReplication", (value))
}

// GetDSAccess_AuditDetailedDirectoryServiceReplication gets the value of DSAccess_AuditDetailedDirectoryServiceReplication for the instance
func (instance *MDM_Policy_Config01_Audit02) GetPropertyDSAccess_AuditDetailedDirectoryServiceReplication() (value int32, err error) {
	retValue, err := instance.GetProperty("DSAccess_AuditDetailedDirectoryServiceReplication")
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

// SetDSAccess_AuditDirectoryServiceAccess sets the value of DSAccess_AuditDirectoryServiceAccess for the instance
func (instance *MDM_Policy_Config01_Audit02) SetPropertyDSAccess_AuditDirectoryServiceAccess(value int32) (err error) {
	return instance.SetProperty("DSAccess_AuditDirectoryServiceAccess", (value))
}

// GetDSAccess_AuditDirectoryServiceAccess gets the value of DSAccess_AuditDirectoryServiceAccess for the instance
func (instance *MDM_Policy_Config01_Audit02) GetPropertyDSAccess_AuditDirectoryServiceAccess() (value int32, err error) {
	retValue, err := instance.GetProperty("DSAccess_AuditDirectoryServiceAccess")
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

// SetDSAccess_AuditDirectoryServiceChanges sets the value of DSAccess_AuditDirectoryServiceChanges for the instance
func (instance *MDM_Policy_Config01_Audit02) SetPropertyDSAccess_AuditDirectoryServiceChanges(value int32) (err error) {
	return instance.SetProperty("DSAccess_AuditDirectoryServiceChanges", (value))
}

// GetDSAccess_AuditDirectoryServiceChanges gets the value of DSAccess_AuditDirectoryServiceChanges for the instance
func (instance *MDM_Policy_Config01_Audit02) GetPropertyDSAccess_AuditDirectoryServiceChanges() (value int32, err error) {
	retValue, err := instance.GetProperty("DSAccess_AuditDirectoryServiceChanges")
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

// SetDSAccess_AuditDirectoryServiceReplication sets the value of DSAccess_AuditDirectoryServiceReplication for the instance
func (instance *MDM_Policy_Config01_Audit02) SetPropertyDSAccess_AuditDirectoryServiceReplication(value int32) (err error) {
	return instance.SetProperty("DSAccess_AuditDirectoryServiceReplication", (value))
}

// GetDSAccess_AuditDirectoryServiceReplication gets the value of DSAccess_AuditDirectoryServiceReplication for the instance
func (instance *MDM_Policy_Config01_Audit02) GetPropertyDSAccess_AuditDirectoryServiceReplication() (value int32, err error) {
	retValue, err := instance.GetProperty("DSAccess_AuditDirectoryServiceReplication")
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

// SetInstanceID sets the value of InstanceID for the instance
func (instance *MDM_Policy_Config01_Audit02) SetPropertyInstanceID(value string) (err error) {
	return instance.SetProperty("InstanceID", (value))
}

// GetInstanceID gets the value of InstanceID for the instance
func (instance *MDM_Policy_Config01_Audit02) GetPropertyInstanceID() (value string, err error) {
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

// SetObjectAccess_AuditApplicationGenerated sets the value of ObjectAccess_AuditApplicationGenerated for the instance
func (instance *MDM_Policy_Config01_Audit02) SetPropertyObjectAccess_AuditApplicationGenerated(value int32) (err error) {
	return instance.SetProperty("ObjectAccess_AuditApplicationGenerated", (value))
}

// GetObjectAccess_AuditApplicationGenerated gets the value of ObjectAccess_AuditApplicationGenerated for the instance
func (instance *MDM_Policy_Config01_Audit02) GetPropertyObjectAccess_AuditApplicationGenerated() (value int32, err error) {
	retValue, err := instance.GetProperty("ObjectAccess_AuditApplicationGenerated")
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

// SetObjectAccess_AuditCentralAccessPolicyStaging sets the value of ObjectAccess_AuditCentralAccessPolicyStaging for the instance
func (instance *MDM_Policy_Config01_Audit02) SetPropertyObjectAccess_AuditCentralAccessPolicyStaging(value int32) (err error) {
	return instance.SetProperty("ObjectAccess_AuditCentralAccessPolicyStaging", (value))
}

// GetObjectAccess_AuditCentralAccessPolicyStaging gets the value of ObjectAccess_AuditCentralAccessPolicyStaging for the instance
func (instance *MDM_Policy_Config01_Audit02) GetPropertyObjectAccess_AuditCentralAccessPolicyStaging() (value int32, err error) {
	retValue, err := instance.GetProperty("ObjectAccess_AuditCentralAccessPolicyStaging")
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

// SetObjectAccess_AuditCertificationServices sets the value of ObjectAccess_AuditCertificationServices for the instance
func (instance *MDM_Policy_Config01_Audit02) SetPropertyObjectAccess_AuditCertificationServices(value int32) (err error) {
	return instance.SetProperty("ObjectAccess_AuditCertificationServices", (value))
}

// GetObjectAccess_AuditCertificationServices gets the value of ObjectAccess_AuditCertificationServices for the instance
func (instance *MDM_Policy_Config01_Audit02) GetPropertyObjectAccess_AuditCertificationServices() (value int32, err error) {
	retValue, err := instance.GetProperty("ObjectAccess_AuditCertificationServices")
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

// SetObjectAccess_AuditDetailedFileShare sets the value of ObjectAccess_AuditDetailedFileShare for the instance
func (instance *MDM_Policy_Config01_Audit02) SetPropertyObjectAccess_AuditDetailedFileShare(value int32) (err error) {
	return instance.SetProperty("ObjectAccess_AuditDetailedFileShare", (value))
}

// GetObjectAccess_AuditDetailedFileShare gets the value of ObjectAccess_AuditDetailedFileShare for the instance
func (instance *MDM_Policy_Config01_Audit02) GetPropertyObjectAccess_AuditDetailedFileShare() (value int32, err error) {
	retValue, err := instance.GetProperty("ObjectAccess_AuditDetailedFileShare")
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

// SetObjectAccess_AuditFileShare sets the value of ObjectAccess_AuditFileShare for the instance
func (instance *MDM_Policy_Config01_Audit02) SetPropertyObjectAccess_AuditFileShare(value int32) (err error) {
	return instance.SetProperty("ObjectAccess_AuditFileShare", (value))
}

// GetObjectAccess_AuditFileShare gets the value of ObjectAccess_AuditFileShare for the instance
func (instance *MDM_Policy_Config01_Audit02) GetPropertyObjectAccess_AuditFileShare() (value int32, err error) {
	retValue, err := instance.GetProperty("ObjectAccess_AuditFileShare")
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

// SetObjectAccess_AuditFileSystem sets the value of ObjectAccess_AuditFileSystem for the instance
func (instance *MDM_Policy_Config01_Audit02) SetPropertyObjectAccess_AuditFileSystem(value int32) (err error) {
	return instance.SetProperty("ObjectAccess_AuditFileSystem", (value))
}

// GetObjectAccess_AuditFileSystem gets the value of ObjectAccess_AuditFileSystem for the instance
func (instance *MDM_Policy_Config01_Audit02) GetPropertyObjectAccess_AuditFileSystem() (value int32, err error) {
	retValue, err := instance.GetProperty("ObjectAccess_AuditFileSystem")
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

// SetObjectAccess_AuditFilteringPlatformConnection sets the value of ObjectAccess_AuditFilteringPlatformConnection for the instance
func (instance *MDM_Policy_Config01_Audit02) SetPropertyObjectAccess_AuditFilteringPlatformConnection(value int32) (err error) {
	return instance.SetProperty("ObjectAccess_AuditFilteringPlatformConnection", (value))
}

// GetObjectAccess_AuditFilteringPlatformConnection gets the value of ObjectAccess_AuditFilteringPlatformConnection for the instance
func (instance *MDM_Policy_Config01_Audit02) GetPropertyObjectAccess_AuditFilteringPlatformConnection() (value int32, err error) {
	retValue, err := instance.GetProperty("ObjectAccess_AuditFilteringPlatformConnection")
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

// SetObjectAccess_AuditFilteringPlatformPacketDrop sets the value of ObjectAccess_AuditFilteringPlatformPacketDrop for the instance
func (instance *MDM_Policy_Config01_Audit02) SetPropertyObjectAccess_AuditFilteringPlatformPacketDrop(value int32) (err error) {
	return instance.SetProperty("ObjectAccess_AuditFilteringPlatformPacketDrop", (value))
}

// GetObjectAccess_AuditFilteringPlatformPacketDrop gets the value of ObjectAccess_AuditFilteringPlatformPacketDrop for the instance
func (instance *MDM_Policy_Config01_Audit02) GetPropertyObjectAccess_AuditFilteringPlatformPacketDrop() (value int32, err error) {
	retValue, err := instance.GetProperty("ObjectAccess_AuditFilteringPlatformPacketDrop")
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

// SetObjectAccess_AuditHandleManipulation sets the value of ObjectAccess_AuditHandleManipulation for the instance
func (instance *MDM_Policy_Config01_Audit02) SetPropertyObjectAccess_AuditHandleManipulation(value int32) (err error) {
	return instance.SetProperty("ObjectAccess_AuditHandleManipulation", (value))
}

// GetObjectAccess_AuditHandleManipulation gets the value of ObjectAccess_AuditHandleManipulation for the instance
func (instance *MDM_Policy_Config01_Audit02) GetPropertyObjectAccess_AuditHandleManipulation() (value int32, err error) {
	retValue, err := instance.GetProperty("ObjectAccess_AuditHandleManipulation")
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

// SetObjectAccess_AuditKernelObject sets the value of ObjectAccess_AuditKernelObject for the instance
func (instance *MDM_Policy_Config01_Audit02) SetPropertyObjectAccess_AuditKernelObject(value int32) (err error) {
	return instance.SetProperty("ObjectAccess_AuditKernelObject", (value))
}

// GetObjectAccess_AuditKernelObject gets the value of ObjectAccess_AuditKernelObject for the instance
func (instance *MDM_Policy_Config01_Audit02) GetPropertyObjectAccess_AuditKernelObject() (value int32, err error) {
	retValue, err := instance.GetProperty("ObjectAccess_AuditKernelObject")
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

// SetObjectAccess_AuditOtherObjectAccessEvents sets the value of ObjectAccess_AuditOtherObjectAccessEvents for the instance
func (instance *MDM_Policy_Config01_Audit02) SetPropertyObjectAccess_AuditOtherObjectAccessEvents(value int32) (err error) {
	return instance.SetProperty("ObjectAccess_AuditOtherObjectAccessEvents", (value))
}

// GetObjectAccess_AuditOtherObjectAccessEvents gets the value of ObjectAccess_AuditOtherObjectAccessEvents for the instance
func (instance *MDM_Policy_Config01_Audit02) GetPropertyObjectAccess_AuditOtherObjectAccessEvents() (value int32, err error) {
	retValue, err := instance.GetProperty("ObjectAccess_AuditOtherObjectAccessEvents")
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

// SetObjectAccess_AuditRegistry sets the value of ObjectAccess_AuditRegistry for the instance
func (instance *MDM_Policy_Config01_Audit02) SetPropertyObjectAccess_AuditRegistry(value int32) (err error) {
	return instance.SetProperty("ObjectAccess_AuditRegistry", (value))
}

// GetObjectAccess_AuditRegistry gets the value of ObjectAccess_AuditRegistry for the instance
func (instance *MDM_Policy_Config01_Audit02) GetPropertyObjectAccess_AuditRegistry() (value int32, err error) {
	retValue, err := instance.GetProperty("ObjectAccess_AuditRegistry")
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

// SetObjectAccess_AuditRemovableStorage sets the value of ObjectAccess_AuditRemovableStorage for the instance
func (instance *MDM_Policy_Config01_Audit02) SetPropertyObjectAccess_AuditRemovableStorage(value int32) (err error) {
	return instance.SetProperty("ObjectAccess_AuditRemovableStorage", (value))
}

// GetObjectAccess_AuditRemovableStorage gets the value of ObjectAccess_AuditRemovableStorage for the instance
func (instance *MDM_Policy_Config01_Audit02) GetPropertyObjectAccess_AuditRemovableStorage() (value int32, err error) {
	retValue, err := instance.GetProperty("ObjectAccess_AuditRemovableStorage")
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

// SetObjectAccess_AuditSAM sets the value of ObjectAccess_AuditSAM for the instance
func (instance *MDM_Policy_Config01_Audit02) SetPropertyObjectAccess_AuditSAM(value int32) (err error) {
	return instance.SetProperty("ObjectAccess_AuditSAM", (value))
}

// GetObjectAccess_AuditSAM gets the value of ObjectAccess_AuditSAM for the instance
func (instance *MDM_Policy_Config01_Audit02) GetPropertyObjectAccess_AuditSAM() (value int32, err error) {
	retValue, err := instance.GetProperty("ObjectAccess_AuditSAM")
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

// SetParentID sets the value of ParentID for the instance
func (instance *MDM_Policy_Config01_Audit02) SetPropertyParentID(value string) (err error) {
	return instance.SetProperty("ParentID", (value))
}

// GetParentID gets the value of ParentID for the instance
func (instance *MDM_Policy_Config01_Audit02) GetPropertyParentID() (value string, err error) {
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

// SetPolicyChange_AuditAuthenticationPolicyChange sets the value of PolicyChange_AuditAuthenticationPolicyChange for the instance
func (instance *MDM_Policy_Config01_Audit02) SetPropertyPolicyChange_AuditAuthenticationPolicyChange(value int32) (err error) {
	return instance.SetProperty("PolicyChange_AuditAuthenticationPolicyChange", (value))
}

// GetPolicyChange_AuditAuthenticationPolicyChange gets the value of PolicyChange_AuditAuthenticationPolicyChange for the instance
func (instance *MDM_Policy_Config01_Audit02) GetPropertyPolicyChange_AuditAuthenticationPolicyChange() (value int32, err error) {
	retValue, err := instance.GetProperty("PolicyChange_AuditAuthenticationPolicyChange")
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

// SetPolicyChange_AuditAuthorizationPolicyChange sets the value of PolicyChange_AuditAuthorizationPolicyChange for the instance
func (instance *MDM_Policy_Config01_Audit02) SetPropertyPolicyChange_AuditAuthorizationPolicyChange(value int32) (err error) {
	return instance.SetProperty("PolicyChange_AuditAuthorizationPolicyChange", (value))
}

// GetPolicyChange_AuditAuthorizationPolicyChange gets the value of PolicyChange_AuditAuthorizationPolicyChange for the instance
func (instance *MDM_Policy_Config01_Audit02) GetPropertyPolicyChange_AuditAuthorizationPolicyChange() (value int32, err error) {
	retValue, err := instance.GetProperty("PolicyChange_AuditAuthorizationPolicyChange")
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

// SetPolicyChange_AuditFilteringPlatformPolicyChange sets the value of PolicyChange_AuditFilteringPlatformPolicyChange for the instance
func (instance *MDM_Policy_Config01_Audit02) SetPropertyPolicyChange_AuditFilteringPlatformPolicyChange(value int32) (err error) {
	return instance.SetProperty("PolicyChange_AuditFilteringPlatformPolicyChange", (value))
}

// GetPolicyChange_AuditFilteringPlatformPolicyChange gets the value of PolicyChange_AuditFilteringPlatformPolicyChange for the instance
func (instance *MDM_Policy_Config01_Audit02) GetPropertyPolicyChange_AuditFilteringPlatformPolicyChange() (value int32, err error) {
	retValue, err := instance.GetProperty("PolicyChange_AuditFilteringPlatformPolicyChange")
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

// SetPolicyChange_AuditMPSSVCRuleLevelPolicyChange sets the value of PolicyChange_AuditMPSSVCRuleLevelPolicyChange for the instance
func (instance *MDM_Policy_Config01_Audit02) SetPropertyPolicyChange_AuditMPSSVCRuleLevelPolicyChange(value int32) (err error) {
	return instance.SetProperty("PolicyChange_AuditMPSSVCRuleLevelPolicyChange", (value))
}

// GetPolicyChange_AuditMPSSVCRuleLevelPolicyChange gets the value of PolicyChange_AuditMPSSVCRuleLevelPolicyChange for the instance
func (instance *MDM_Policy_Config01_Audit02) GetPropertyPolicyChange_AuditMPSSVCRuleLevelPolicyChange() (value int32, err error) {
	retValue, err := instance.GetProperty("PolicyChange_AuditMPSSVCRuleLevelPolicyChange")
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

// SetPolicyChange_AuditOtherPolicyChangeEvents sets the value of PolicyChange_AuditOtherPolicyChangeEvents for the instance
func (instance *MDM_Policy_Config01_Audit02) SetPropertyPolicyChange_AuditOtherPolicyChangeEvents(value int32) (err error) {
	return instance.SetProperty("PolicyChange_AuditOtherPolicyChangeEvents", (value))
}

// GetPolicyChange_AuditOtherPolicyChangeEvents gets the value of PolicyChange_AuditOtherPolicyChangeEvents for the instance
func (instance *MDM_Policy_Config01_Audit02) GetPropertyPolicyChange_AuditOtherPolicyChangeEvents() (value int32, err error) {
	retValue, err := instance.GetProperty("PolicyChange_AuditOtherPolicyChangeEvents")
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

// SetPolicyChange_AuditPolicyChange sets the value of PolicyChange_AuditPolicyChange for the instance
func (instance *MDM_Policy_Config01_Audit02) SetPropertyPolicyChange_AuditPolicyChange(value int32) (err error) {
	return instance.SetProperty("PolicyChange_AuditPolicyChange", (value))
}

// GetPolicyChange_AuditPolicyChange gets the value of PolicyChange_AuditPolicyChange for the instance
func (instance *MDM_Policy_Config01_Audit02) GetPropertyPolicyChange_AuditPolicyChange() (value int32, err error) {
	retValue, err := instance.GetProperty("PolicyChange_AuditPolicyChange")
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

// SetPrivilegeUse_AuditNonSensitivePrivilegeUse sets the value of PrivilegeUse_AuditNonSensitivePrivilegeUse for the instance
func (instance *MDM_Policy_Config01_Audit02) SetPropertyPrivilegeUse_AuditNonSensitivePrivilegeUse(value int32) (err error) {
	return instance.SetProperty("PrivilegeUse_AuditNonSensitivePrivilegeUse", (value))
}

// GetPrivilegeUse_AuditNonSensitivePrivilegeUse gets the value of PrivilegeUse_AuditNonSensitivePrivilegeUse for the instance
func (instance *MDM_Policy_Config01_Audit02) GetPropertyPrivilegeUse_AuditNonSensitivePrivilegeUse() (value int32, err error) {
	retValue, err := instance.GetProperty("PrivilegeUse_AuditNonSensitivePrivilegeUse")
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

// SetPrivilegeUse_AuditOtherPrivilegeUseEvents sets the value of PrivilegeUse_AuditOtherPrivilegeUseEvents for the instance
func (instance *MDM_Policy_Config01_Audit02) SetPropertyPrivilegeUse_AuditOtherPrivilegeUseEvents(value int32) (err error) {
	return instance.SetProperty("PrivilegeUse_AuditOtherPrivilegeUseEvents", (value))
}

// GetPrivilegeUse_AuditOtherPrivilegeUseEvents gets the value of PrivilegeUse_AuditOtherPrivilegeUseEvents for the instance
func (instance *MDM_Policy_Config01_Audit02) GetPropertyPrivilegeUse_AuditOtherPrivilegeUseEvents() (value int32, err error) {
	retValue, err := instance.GetProperty("PrivilegeUse_AuditOtherPrivilegeUseEvents")
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

// SetPrivilegeUse_AuditSensitivePrivilegeUse sets the value of PrivilegeUse_AuditSensitivePrivilegeUse for the instance
func (instance *MDM_Policy_Config01_Audit02) SetPropertyPrivilegeUse_AuditSensitivePrivilegeUse(value int32) (err error) {
	return instance.SetProperty("PrivilegeUse_AuditSensitivePrivilegeUse", (value))
}

// GetPrivilegeUse_AuditSensitivePrivilegeUse gets the value of PrivilegeUse_AuditSensitivePrivilegeUse for the instance
func (instance *MDM_Policy_Config01_Audit02) GetPropertyPrivilegeUse_AuditSensitivePrivilegeUse() (value int32, err error) {
	retValue, err := instance.GetProperty("PrivilegeUse_AuditSensitivePrivilegeUse")
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

// SetSystem_AuditIPsecDriver sets the value of System_AuditIPsecDriver for the instance
func (instance *MDM_Policy_Config01_Audit02) SetPropertySystem_AuditIPsecDriver(value int32) (err error) {
	return instance.SetProperty("System_AuditIPsecDriver", (value))
}

// GetSystem_AuditIPsecDriver gets the value of System_AuditIPsecDriver for the instance
func (instance *MDM_Policy_Config01_Audit02) GetPropertySystem_AuditIPsecDriver() (value int32, err error) {
	retValue, err := instance.GetProperty("System_AuditIPsecDriver")
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

// SetSystem_AuditOtherSystemEvents sets the value of System_AuditOtherSystemEvents for the instance
func (instance *MDM_Policy_Config01_Audit02) SetPropertySystem_AuditOtherSystemEvents(value int32) (err error) {
	return instance.SetProperty("System_AuditOtherSystemEvents", (value))
}

// GetSystem_AuditOtherSystemEvents gets the value of System_AuditOtherSystemEvents for the instance
func (instance *MDM_Policy_Config01_Audit02) GetPropertySystem_AuditOtherSystemEvents() (value int32, err error) {
	retValue, err := instance.GetProperty("System_AuditOtherSystemEvents")
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

// SetSystem_AuditSecurityStateChange sets the value of System_AuditSecurityStateChange for the instance
func (instance *MDM_Policy_Config01_Audit02) SetPropertySystem_AuditSecurityStateChange(value int32) (err error) {
	return instance.SetProperty("System_AuditSecurityStateChange", (value))
}

// GetSystem_AuditSecurityStateChange gets the value of System_AuditSecurityStateChange for the instance
func (instance *MDM_Policy_Config01_Audit02) GetPropertySystem_AuditSecurityStateChange() (value int32, err error) {
	retValue, err := instance.GetProperty("System_AuditSecurityStateChange")
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

// SetSystem_AuditSecuritySystemExtension sets the value of System_AuditSecuritySystemExtension for the instance
func (instance *MDM_Policy_Config01_Audit02) SetPropertySystem_AuditSecuritySystemExtension(value int32) (err error) {
	return instance.SetProperty("System_AuditSecuritySystemExtension", (value))
}

// GetSystem_AuditSecuritySystemExtension gets the value of System_AuditSecuritySystemExtension for the instance
func (instance *MDM_Policy_Config01_Audit02) GetPropertySystem_AuditSecuritySystemExtension() (value int32, err error) {
	retValue, err := instance.GetProperty("System_AuditSecuritySystemExtension")
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

// SetSystem_AuditSystemIntegrity sets the value of System_AuditSystemIntegrity for the instance
func (instance *MDM_Policy_Config01_Audit02) SetPropertySystem_AuditSystemIntegrity(value int32) (err error) {
	return instance.SetProperty("System_AuditSystemIntegrity", (value))
}

// GetSystem_AuditSystemIntegrity gets the value of System_AuditSystemIntegrity for the instance
func (instance *MDM_Policy_Config01_Audit02) GetPropertySystem_AuditSystemIntegrity() (value int32, err error) {
	retValue, err := instance.GetProperty("System_AuditSystemIntegrity")
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
