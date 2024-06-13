// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 6/13/2024 using wmigen
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

// MDM_Policy_Result01_LocalPoliciesSecurityOptions02 struct
type MDM_Policy_Result01_LocalPoliciesSecurityOptions02 struct {
	*cim.WmiInstance

	//
	Accounts_BlockMicrosoftAccounts int32

	//
	Accounts_EnableAdministratorAccountStatus int32

	//
	Accounts_EnableGuestAccountStatus int32

	//
	Accounts_LimitLocalAccountUseOfBlankPasswordsToConsoleLogonOnly int32

	//
	Accounts_RenameAdministratorAccount string

	//
	Accounts_RenameGuestAccount string

	//
	Devices_AllowedToFormatAndEjectRemovableMedia string

	//
	Devices_AllowUndockWithoutHavingToLogon int32

	//
	Devices_PreventUsersFromInstallingPrinterDriversWhenConnectingToSharedPrinters int32

	//
	Devices_RestrictCDROMAccessToLocallyLoggedOnUserOnly string

	//
	InstanceID string

	//
	InteractiveLogon_DisplayUserInformationWhenTheSessionIsLocked int32

	//
	InteractiveLogon_DoNotDisplayLastSignedIn int32

	//
	InteractiveLogon_DoNotDisplayUsernameAtSignIn int32

	//
	InteractiveLogon_DoNotRequireCTRLALTDEL int32

	//
	InteractiveLogon_MachineInactivityLimit int32

	//
	InteractiveLogon_MessageTextForUsersAttemptingToLogOn string

	//
	InteractiveLogon_MessageTitleForUsersAttemptingToLogOn string

	//
	InteractiveLogon_SmartCardRemovalBehavior string

	//
	MicrosoftNetworkClient_DigitallySignCommunicationsAlways int32

	//
	MicrosoftNetworkClient_DigitallySignCommunicationsIfServerAgrees int32

	//
	MicrosoftNetworkClient_SendUnencryptedPasswordToThirdPartySMBServers int32

	//
	MicrosoftNetworkServer_DigitallySignCommunicationsAlways int32

	//
	MicrosoftNetworkServer_DigitallySignCommunicationsIfClientAgrees int32

	//
	NetworkAccess_DoNotAllowAnonymousEnumerationOfSAMAccounts int32

	//
	NetworkAccess_DoNotAllowAnonymousEnumerationOfSamAccountsAndShares int32

	//
	NetworkAccess_RestrictAnonymousAccessToNamedPipesAndShares int32

	//
	NetworkAccess_RestrictClientsAllowedToMakeRemoteCallsToSAM string

	//
	NetworkSecurity_AllowLocalSystemToUseComputerIdentityForNTLM int32

	//
	NetworkSecurity_AllowPKU2UAuthenticationRequests int32

	//
	NetworkSecurity_DoNotStoreLANManagerHashValueOnNextPasswordChange int32

	//
	NetworkSecurity_LANManagerAuthenticationLevel int32

	//
	NetworkSecurity_MinimumSessionSecurityForNTLMSSPBasedClients int32

	//
	NetworkSecurity_MinimumSessionSecurityForNTLMSSPBasedServers int32

	//
	NetworkSecurity_RestrictNTLM_AddRemoteServerExceptionsForNTLMAuthentication string

	//
	NetworkSecurity_RestrictNTLM_AuditIncomingNTLMTraffic int32

	//
	NetworkSecurity_RestrictNTLM_IncomingNTLMTraffic int32

	//
	NetworkSecurity_RestrictNTLM_OutgoingNTLMTrafficToRemoteServers int32

	//
	ParentID string

	//
	Shutdown_AllowSystemToBeShutDownWithoutHavingToLogOn int32

	//
	Shutdown_ClearVirtualMemoryPageFile int32

	//
	UserAccountControl_AllowUIAccessApplicationsToPromptForElevation int32

	//
	UserAccountControl_BehaviorOfTheElevationPromptForAdministrators int32

	//
	UserAccountControl_BehaviorOfTheElevationPromptForStandardUsers int32

	//
	UserAccountControl_DetectApplicationInstallationsAndPromptForElevation int32

	//
	UserAccountControl_OnlyElevateExecutableFilesThatAreSignedAndValidated int32

	//
	UserAccountControl_OnlyElevateUIAccessApplicationsThatAreInstalledInSecureLocations int32

	//
	UserAccountControl_RunAllAdministratorsInAdminApprovalMode int32

	//
	UserAccountControl_SwitchToTheSecureDesktopWhenPromptingForElevation int32

	//
	UserAccountControl_UseAdminApprovalMode int32

	//
	UserAccountControl_VirtualizeFileAndRegistryWriteFailuresToPerUserLocations int32
}

func NewMDM_Policy_Result01_LocalPoliciesSecurityOptions02Ex1(instance *cim.WmiInstance) (newInstance *MDM_Policy_Result01_LocalPoliciesSecurityOptions02, err error) {
	tmp, err := instance, nil

	if err != nil {
		return
	}
	newInstance = &MDM_Policy_Result01_LocalPoliciesSecurityOptions02{
		WmiInstance: tmp,
	}
	return
}

func NewMDM_Policy_Result01_LocalPoliciesSecurityOptions02Ex6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *MDM_Policy_Result01_LocalPoliciesSecurityOptions02, err error) {
	tmp, err := instance.GetWmiInstance(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &MDM_Policy_Result01_LocalPoliciesSecurityOptions02{
		WmiInstance: tmp,
	}
	return
}

// SetAccounts_BlockMicrosoftAccounts sets the value of Accounts_BlockMicrosoftAccounts for the instance
func (instance *MDM_Policy_Result01_LocalPoliciesSecurityOptions02) SetPropertyAccounts_BlockMicrosoftAccounts(value int32) (err error) {
	return instance.SetProperty("Accounts_BlockMicrosoftAccounts", (value))
}

// GetAccounts_BlockMicrosoftAccounts gets the value of Accounts_BlockMicrosoftAccounts for the instance
func (instance *MDM_Policy_Result01_LocalPoliciesSecurityOptions02) GetPropertyAccounts_BlockMicrosoftAccounts() (value int32, err error) {
	retValue, err := instance.GetProperty("Accounts_BlockMicrosoftAccounts")
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

// SetAccounts_EnableAdministratorAccountStatus sets the value of Accounts_EnableAdministratorAccountStatus for the instance
func (instance *MDM_Policy_Result01_LocalPoliciesSecurityOptions02) SetPropertyAccounts_EnableAdministratorAccountStatus(value int32) (err error) {
	return instance.SetProperty("Accounts_EnableAdministratorAccountStatus", (value))
}

// GetAccounts_EnableAdministratorAccountStatus gets the value of Accounts_EnableAdministratorAccountStatus for the instance
func (instance *MDM_Policy_Result01_LocalPoliciesSecurityOptions02) GetPropertyAccounts_EnableAdministratorAccountStatus() (value int32, err error) {
	retValue, err := instance.GetProperty("Accounts_EnableAdministratorAccountStatus")
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

// SetAccounts_EnableGuestAccountStatus sets the value of Accounts_EnableGuestAccountStatus for the instance
func (instance *MDM_Policy_Result01_LocalPoliciesSecurityOptions02) SetPropertyAccounts_EnableGuestAccountStatus(value int32) (err error) {
	return instance.SetProperty("Accounts_EnableGuestAccountStatus", (value))
}

// GetAccounts_EnableGuestAccountStatus gets the value of Accounts_EnableGuestAccountStatus for the instance
func (instance *MDM_Policy_Result01_LocalPoliciesSecurityOptions02) GetPropertyAccounts_EnableGuestAccountStatus() (value int32, err error) {
	retValue, err := instance.GetProperty("Accounts_EnableGuestAccountStatus")
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

// SetAccounts_LimitLocalAccountUseOfBlankPasswordsToConsoleLogonOnly sets the value of Accounts_LimitLocalAccountUseOfBlankPasswordsToConsoleLogonOnly for the instance
func (instance *MDM_Policy_Result01_LocalPoliciesSecurityOptions02) SetPropertyAccounts_LimitLocalAccountUseOfBlankPasswordsToConsoleLogonOnly(value int32) (err error) {
	return instance.SetProperty("Accounts_LimitLocalAccountUseOfBlankPasswordsToConsoleLogonOnly", (value))
}

// GetAccounts_LimitLocalAccountUseOfBlankPasswordsToConsoleLogonOnly gets the value of Accounts_LimitLocalAccountUseOfBlankPasswordsToConsoleLogonOnly for the instance
func (instance *MDM_Policy_Result01_LocalPoliciesSecurityOptions02) GetPropertyAccounts_LimitLocalAccountUseOfBlankPasswordsToConsoleLogonOnly() (value int32, err error) {
	retValue, err := instance.GetProperty("Accounts_LimitLocalAccountUseOfBlankPasswordsToConsoleLogonOnly")
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

// SetAccounts_RenameAdministratorAccount sets the value of Accounts_RenameAdministratorAccount for the instance
func (instance *MDM_Policy_Result01_LocalPoliciesSecurityOptions02) SetPropertyAccounts_RenameAdministratorAccount(value string) (err error) {
	return instance.SetProperty("Accounts_RenameAdministratorAccount", (value))
}

// GetAccounts_RenameAdministratorAccount gets the value of Accounts_RenameAdministratorAccount for the instance
func (instance *MDM_Policy_Result01_LocalPoliciesSecurityOptions02) GetPropertyAccounts_RenameAdministratorAccount() (value string, err error) {
	retValue, err := instance.GetProperty("Accounts_RenameAdministratorAccount")
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

// SetAccounts_RenameGuestAccount sets the value of Accounts_RenameGuestAccount for the instance
func (instance *MDM_Policy_Result01_LocalPoliciesSecurityOptions02) SetPropertyAccounts_RenameGuestAccount(value string) (err error) {
	return instance.SetProperty("Accounts_RenameGuestAccount", (value))
}

// GetAccounts_RenameGuestAccount gets the value of Accounts_RenameGuestAccount for the instance
func (instance *MDM_Policy_Result01_LocalPoliciesSecurityOptions02) GetPropertyAccounts_RenameGuestAccount() (value string, err error) {
	retValue, err := instance.GetProperty("Accounts_RenameGuestAccount")
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

// SetDevices_AllowedToFormatAndEjectRemovableMedia sets the value of Devices_AllowedToFormatAndEjectRemovableMedia for the instance
func (instance *MDM_Policy_Result01_LocalPoliciesSecurityOptions02) SetPropertyDevices_AllowedToFormatAndEjectRemovableMedia(value string) (err error) {
	return instance.SetProperty("Devices_AllowedToFormatAndEjectRemovableMedia", (value))
}

// GetDevices_AllowedToFormatAndEjectRemovableMedia gets the value of Devices_AllowedToFormatAndEjectRemovableMedia for the instance
func (instance *MDM_Policy_Result01_LocalPoliciesSecurityOptions02) GetPropertyDevices_AllowedToFormatAndEjectRemovableMedia() (value string, err error) {
	retValue, err := instance.GetProperty("Devices_AllowedToFormatAndEjectRemovableMedia")
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

// SetDevices_AllowUndockWithoutHavingToLogon sets the value of Devices_AllowUndockWithoutHavingToLogon for the instance
func (instance *MDM_Policy_Result01_LocalPoliciesSecurityOptions02) SetPropertyDevices_AllowUndockWithoutHavingToLogon(value int32) (err error) {
	return instance.SetProperty("Devices_AllowUndockWithoutHavingToLogon", (value))
}

// GetDevices_AllowUndockWithoutHavingToLogon gets the value of Devices_AllowUndockWithoutHavingToLogon for the instance
func (instance *MDM_Policy_Result01_LocalPoliciesSecurityOptions02) GetPropertyDevices_AllowUndockWithoutHavingToLogon() (value int32, err error) {
	retValue, err := instance.GetProperty("Devices_AllowUndockWithoutHavingToLogon")
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

// SetDevices_PreventUsersFromInstallingPrinterDriversWhenConnectingToSharedPrinters sets the value of Devices_PreventUsersFromInstallingPrinterDriversWhenConnectingToSharedPrinters for the instance
func (instance *MDM_Policy_Result01_LocalPoliciesSecurityOptions02) SetPropertyDevices_PreventUsersFromInstallingPrinterDriversWhenConnectingToSharedPrinters(value int32) (err error) {
	return instance.SetProperty("Devices_PreventUsersFromInstallingPrinterDriversWhenConnectingToSharedPrinters", (value))
}

// GetDevices_PreventUsersFromInstallingPrinterDriversWhenConnectingToSharedPrinters gets the value of Devices_PreventUsersFromInstallingPrinterDriversWhenConnectingToSharedPrinters for the instance
func (instance *MDM_Policy_Result01_LocalPoliciesSecurityOptions02) GetPropertyDevices_PreventUsersFromInstallingPrinterDriversWhenConnectingToSharedPrinters() (value int32, err error) {
	retValue, err := instance.GetProperty("Devices_PreventUsersFromInstallingPrinterDriversWhenConnectingToSharedPrinters")
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

// SetDevices_RestrictCDROMAccessToLocallyLoggedOnUserOnly sets the value of Devices_RestrictCDROMAccessToLocallyLoggedOnUserOnly for the instance
func (instance *MDM_Policy_Result01_LocalPoliciesSecurityOptions02) SetPropertyDevices_RestrictCDROMAccessToLocallyLoggedOnUserOnly(value string) (err error) {
	return instance.SetProperty("Devices_RestrictCDROMAccessToLocallyLoggedOnUserOnly", (value))
}

// GetDevices_RestrictCDROMAccessToLocallyLoggedOnUserOnly gets the value of Devices_RestrictCDROMAccessToLocallyLoggedOnUserOnly for the instance
func (instance *MDM_Policy_Result01_LocalPoliciesSecurityOptions02) GetPropertyDevices_RestrictCDROMAccessToLocallyLoggedOnUserOnly() (value string, err error) {
	retValue, err := instance.GetProperty("Devices_RestrictCDROMAccessToLocallyLoggedOnUserOnly")
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

// SetInstanceID sets the value of InstanceID for the instance
func (instance *MDM_Policy_Result01_LocalPoliciesSecurityOptions02) SetPropertyInstanceID(value string) (err error) {
	return instance.SetProperty("InstanceID", (value))
}

// GetInstanceID gets the value of InstanceID for the instance
func (instance *MDM_Policy_Result01_LocalPoliciesSecurityOptions02) GetPropertyInstanceID() (value string, err error) {
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

// SetInteractiveLogon_DisplayUserInformationWhenTheSessionIsLocked sets the value of InteractiveLogon_DisplayUserInformationWhenTheSessionIsLocked for the instance
func (instance *MDM_Policy_Result01_LocalPoliciesSecurityOptions02) SetPropertyInteractiveLogon_DisplayUserInformationWhenTheSessionIsLocked(value int32) (err error) {
	return instance.SetProperty("InteractiveLogon_DisplayUserInformationWhenTheSessionIsLocked", (value))
}

// GetInteractiveLogon_DisplayUserInformationWhenTheSessionIsLocked gets the value of InteractiveLogon_DisplayUserInformationWhenTheSessionIsLocked for the instance
func (instance *MDM_Policy_Result01_LocalPoliciesSecurityOptions02) GetPropertyInteractiveLogon_DisplayUserInformationWhenTheSessionIsLocked() (value int32, err error) {
	retValue, err := instance.GetProperty("InteractiveLogon_DisplayUserInformationWhenTheSessionIsLocked")
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

// SetInteractiveLogon_DoNotDisplayLastSignedIn sets the value of InteractiveLogon_DoNotDisplayLastSignedIn for the instance
func (instance *MDM_Policy_Result01_LocalPoliciesSecurityOptions02) SetPropertyInteractiveLogon_DoNotDisplayLastSignedIn(value int32) (err error) {
	return instance.SetProperty("InteractiveLogon_DoNotDisplayLastSignedIn", (value))
}

// GetInteractiveLogon_DoNotDisplayLastSignedIn gets the value of InteractiveLogon_DoNotDisplayLastSignedIn for the instance
func (instance *MDM_Policy_Result01_LocalPoliciesSecurityOptions02) GetPropertyInteractiveLogon_DoNotDisplayLastSignedIn() (value int32, err error) {
	retValue, err := instance.GetProperty("InteractiveLogon_DoNotDisplayLastSignedIn")
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

// SetInteractiveLogon_DoNotDisplayUsernameAtSignIn sets the value of InteractiveLogon_DoNotDisplayUsernameAtSignIn for the instance
func (instance *MDM_Policy_Result01_LocalPoliciesSecurityOptions02) SetPropertyInteractiveLogon_DoNotDisplayUsernameAtSignIn(value int32) (err error) {
	return instance.SetProperty("InteractiveLogon_DoNotDisplayUsernameAtSignIn", (value))
}

// GetInteractiveLogon_DoNotDisplayUsernameAtSignIn gets the value of InteractiveLogon_DoNotDisplayUsernameAtSignIn for the instance
func (instance *MDM_Policy_Result01_LocalPoliciesSecurityOptions02) GetPropertyInteractiveLogon_DoNotDisplayUsernameAtSignIn() (value int32, err error) {
	retValue, err := instance.GetProperty("InteractiveLogon_DoNotDisplayUsernameAtSignIn")
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

// SetInteractiveLogon_DoNotRequireCTRLALTDEL sets the value of InteractiveLogon_DoNotRequireCTRLALTDEL for the instance
func (instance *MDM_Policy_Result01_LocalPoliciesSecurityOptions02) SetPropertyInteractiveLogon_DoNotRequireCTRLALTDEL(value int32) (err error) {
	return instance.SetProperty("InteractiveLogon_DoNotRequireCTRLALTDEL", (value))
}

// GetInteractiveLogon_DoNotRequireCTRLALTDEL gets the value of InteractiveLogon_DoNotRequireCTRLALTDEL for the instance
func (instance *MDM_Policy_Result01_LocalPoliciesSecurityOptions02) GetPropertyInteractiveLogon_DoNotRequireCTRLALTDEL() (value int32, err error) {
	retValue, err := instance.GetProperty("InteractiveLogon_DoNotRequireCTRLALTDEL")
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

// SetInteractiveLogon_MachineInactivityLimit sets the value of InteractiveLogon_MachineInactivityLimit for the instance
func (instance *MDM_Policy_Result01_LocalPoliciesSecurityOptions02) SetPropertyInteractiveLogon_MachineInactivityLimit(value int32) (err error) {
	return instance.SetProperty("InteractiveLogon_MachineInactivityLimit", (value))
}

// GetInteractiveLogon_MachineInactivityLimit gets the value of InteractiveLogon_MachineInactivityLimit for the instance
func (instance *MDM_Policy_Result01_LocalPoliciesSecurityOptions02) GetPropertyInteractiveLogon_MachineInactivityLimit() (value int32, err error) {
	retValue, err := instance.GetProperty("InteractiveLogon_MachineInactivityLimit")
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

// SetInteractiveLogon_MessageTextForUsersAttemptingToLogOn sets the value of InteractiveLogon_MessageTextForUsersAttemptingToLogOn for the instance
func (instance *MDM_Policy_Result01_LocalPoliciesSecurityOptions02) SetPropertyInteractiveLogon_MessageTextForUsersAttemptingToLogOn(value string) (err error) {
	return instance.SetProperty("InteractiveLogon_MessageTextForUsersAttemptingToLogOn", (value))
}

// GetInteractiveLogon_MessageTextForUsersAttemptingToLogOn gets the value of InteractiveLogon_MessageTextForUsersAttemptingToLogOn for the instance
func (instance *MDM_Policy_Result01_LocalPoliciesSecurityOptions02) GetPropertyInteractiveLogon_MessageTextForUsersAttemptingToLogOn() (value string, err error) {
	retValue, err := instance.GetProperty("InteractiveLogon_MessageTextForUsersAttemptingToLogOn")
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

// SetInteractiveLogon_MessageTitleForUsersAttemptingToLogOn sets the value of InteractiveLogon_MessageTitleForUsersAttemptingToLogOn for the instance
func (instance *MDM_Policy_Result01_LocalPoliciesSecurityOptions02) SetPropertyInteractiveLogon_MessageTitleForUsersAttemptingToLogOn(value string) (err error) {
	return instance.SetProperty("InteractiveLogon_MessageTitleForUsersAttemptingToLogOn", (value))
}

// GetInteractiveLogon_MessageTitleForUsersAttemptingToLogOn gets the value of InteractiveLogon_MessageTitleForUsersAttemptingToLogOn for the instance
func (instance *MDM_Policy_Result01_LocalPoliciesSecurityOptions02) GetPropertyInteractiveLogon_MessageTitleForUsersAttemptingToLogOn() (value string, err error) {
	retValue, err := instance.GetProperty("InteractiveLogon_MessageTitleForUsersAttemptingToLogOn")
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

// SetInteractiveLogon_SmartCardRemovalBehavior sets the value of InteractiveLogon_SmartCardRemovalBehavior for the instance
func (instance *MDM_Policy_Result01_LocalPoliciesSecurityOptions02) SetPropertyInteractiveLogon_SmartCardRemovalBehavior(value string) (err error) {
	return instance.SetProperty("InteractiveLogon_SmartCardRemovalBehavior", (value))
}

// GetInteractiveLogon_SmartCardRemovalBehavior gets the value of InteractiveLogon_SmartCardRemovalBehavior for the instance
func (instance *MDM_Policy_Result01_LocalPoliciesSecurityOptions02) GetPropertyInteractiveLogon_SmartCardRemovalBehavior() (value string, err error) {
	retValue, err := instance.GetProperty("InteractiveLogon_SmartCardRemovalBehavior")
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

// SetMicrosoftNetworkClient_DigitallySignCommunicationsAlways sets the value of MicrosoftNetworkClient_DigitallySignCommunicationsAlways for the instance
func (instance *MDM_Policy_Result01_LocalPoliciesSecurityOptions02) SetPropertyMicrosoftNetworkClient_DigitallySignCommunicationsAlways(value int32) (err error) {
	return instance.SetProperty("MicrosoftNetworkClient_DigitallySignCommunicationsAlways", (value))
}

// GetMicrosoftNetworkClient_DigitallySignCommunicationsAlways gets the value of MicrosoftNetworkClient_DigitallySignCommunicationsAlways for the instance
func (instance *MDM_Policy_Result01_LocalPoliciesSecurityOptions02) GetPropertyMicrosoftNetworkClient_DigitallySignCommunicationsAlways() (value int32, err error) {
	retValue, err := instance.GetProperty("MicrosoftNetworkClient_DigitallySignCommunicationsAlways")
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

// SetMicrosoftNetworkClient_DigitallySignCommunicationsIfServerAgrees sets the value of MicrosoftNetworkClient_DigitallySignCommunicationsIfServerAgrees for the instance
func (instance *MDM_Policy_Result01_LocalPoliciesSecurityOptions02) SetPropertyMicrosoftNetworkClient_DigitallySignCommunicationsIfServerAgrees(value int32) (err error) {
	return instance.SetProperty("MicrosoftNetworkClient_DigitallySignCommunicationsIfServerAgrees", (value))
}

// GetMicrosoftNetworkClient_DigitallySignCommunicationsIfServerAgrees gets the value of MicrosoftNetworkClient_DigitallySignCommunicationsIfServerAgrees for the instance
func (instance *MDM_Policy_Result01_LocalPoliciesSecurityOptions02) GetPropertyMicrosoftNetworkClient_DigitallySignCommunicationsIfServerAgrees() (value int32, err error) {
	retValue, err := instance.GetProperty("MicrosoftNetworkClient_DigitallySignCommunicationsIfServerAgrees")
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

// SetMicrosoftNetworkClient_SendUnencryptedPasswordToThirdPartySMBServers sets the value of MicrosoftNetworkClient_SendUnencryptedPasswordToThirdPartySMBServers for the instance
func (instance *MDM_Policy_Result01_LocalPoliciesSecurityOptions02) SetPropertyMicrosoftNetworkClient_SendUnencryptedPasswordToThirdPartySMBServers(value int32) (err error) {
	return instance.SetProperty("MicrosoftNetworkClient_SendUnencryptedPasswordToThirdPartySMBServers", (value))
}

// GetMicrosoftNetworkClient_SendUnencryptedPasswordToThirdPartySMBServers gets the value of MicrosoftNetworkClient_SendUnencryptedPasswordToThirdPartySMBServers for the instance
func (instance *MDM_Policy_Result01_LocalPoliciesSecurityOptions02) GetPropertyMicrosoftNetworkClient_SendUnencryptedPasswordToThirdPartySMBServers() (value int32, err error) {
	retValue, err := instance.GetProperty("MicrosoftNetworkClient_SendUnencryptedPasswordToThirdPartySMBServers")
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

// SetMicrosoftNetworkServer_DigitallySignCommunicationsAlways sets the value of MicrosoftNetworkServer_DigitallySignCommunicationsAlways for the instance
func (instance *MDM_Policy_Result01_LocalPoliciesSecurityOptions02) SetPropertyMicrosoftNetworkServer_DigitallySignCommunicationsAlways(value int32) (err error) {
	return instance.SetProperty("MicrosoftNetworkServer_DigitallySignCommunicationsAlways", (value))
}

// GetMicrosoftNetworkServer_DigitallySignCommunicationsAlways gets the value of MicrosoftNetworkServer_DigitallySignCommunicationsAlways for the instance
func (instance *MDM_Policy_Result01_LocalPoliciesSecurityOptions02) GetPropertyMicrosoftNetworkServer_DigitallySignCommunicationsAlways() (value int32, err error) {
	retValue, err := instance.GetProperty("MicrosoftNetworkServer_DigitallySignCommunicationsAlways")
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

// SetMicrosoftNetworkServer_DigitallySignCommunicationsIfClientAgrees sets the value of MicrosoftNetworkServer_DigitallySignCommunicationsIfClientAgrees for the instance
func (instance *MDM_Policy_Result01_LocalPoliciesSecurityOptions02) SetPropertyMicrosoftNetworkServer_DigitallySignCommunicationsIfClientAgrees(value int32) (err error) {
	return instance.SetProperty("MicrosoftNetworkServer_DigitallySignCommunicationsIfClientAgrees", (value))
}

// GetMicrosoftNetworkServer_DigitallySignCommunicationsIfClientAgrees gets the value of MicrosoftNetworkServer_DigitallySignCommunicationsIfClientAgrees for the instance
func (instance *MDM_Policy_Result01_LocalPoliciesSecurityOptions02) GetPropertyMicrosoftNetworkServer_DigitallySignCommunicationsIfClientAgrees() (value int32, err error) {
	retValue, err := instance.GetProperty("MicrosoftNetworkServer_DigitallySignCommunicationsIfClientAgrees")
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

// SetNetworkAccess_DoNotAllowAnonymousEnumerationOfSAMAccounts sets the value of NetworkAccess_DoNotAllowAnonymousEnumerationOfSAMAccounts for the instance
func (instance *MDM_Policy_Result01_LocalPoliciesSecurityOptions02) SetPropertyNetworkAccess_DoNotAllowAnonymousEnumerationOfSAMAccounts(value int32) (err error) {
	return instance.SetProperty("NetworkAccess_DoNotAllowAnonymousEnumerationOfSAMAccounts", (value))
}

// GetNetworkAccess_DoNotAllowAnonymousEnumerationOfSAMAccounts gets the value of NetworkAccess_DoNotAllowAnonymousEnumerationOfSAMAccounts for the instance
func (instance *MDM_Policy_Result01_LocalPoliciesSecurityOptions02) GetPropertyNetworkAccess_DoNotAllowAnonymousEnumerationOfSAMAccounts() (value int32, err error) {
	retValue, err := instance.GetProperty("NetworkAccess_DoNotAllowAnonymousEnumerationOfSAMAccounts")
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

// SetNetworkAccess_DoNotAllowAnonymousEnumerationOfSamAccountsAndShares sets the value of NetworkAccess_DoNotAllowAnonymousEnumerationOfSamAccountsAndShares for the instance
func (instance *MDM_Policy_Result01_LocalPoliciesSecurityOptions02) SetPropertyNetworkAccess_DoNotAllowAnonymousEnumerationOfSamAccountsAndShares(value int32) (err error) {
	return instance.SetProperty("NetworkAccess_DoNotAllowAnonymousEnumerationOfSamAccountsAndShares", (value))
}

// GetNetworkAccess_DoNotAllowAnonymousEnumerationOfSamAccountsAndShares gets the value of NetworkAccess_DoNotAllowAnonymousEnumerationOfSamAccountsAndShares for the instance
func (instance *MDM_Policy_Result01_LocalPoliciesSecurityOptions02) GetPropertyNetworkAccess_DoNotAllowAnonymousEnumerationOfSamAccountsAndShares() (value int32, err error) {
	retValue, err := instance.GetProperty("NetworkAccess_DoNotAllowAnonymousEnumerationOfSamAccountsAndShares")
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

// SetNetworkAccess_RestrictAnonymousAccessToNamedPipesAndShares sets the value of NetworkAccess_RestrictAnonymousAccessToNamedPipesAndShares for the instance
func (instance *MDM_Policy_Result01_LocalPoliciesSecurityOptions02) SetPropertyNetworkAccess_RestrictAnonymousAccessToNamedPipesAndShares(value int32) (err error) {
	return instance.SetProperty("NetworkAccess_RestrictAnonymousAccessToNamedPipesAndShares", (value))
}

// GetNetworkAccess_RestrictAnonymousAccessToNamedPipesAndShares gets the value of NetworkAccess_RestrictAnonymousAccessToNamedPipesAndShares for the instance
func (instance *MDM_Policy_Result01_LocalPoliciesSecurityOptions02) GetPropertyNetworkAccess_RestrictAnonymousAccessToNamedPipesAndShares() (value int32, err error) {
	retValue, err := instance.GetProperty("NetworkAccess_RestrictAnonymousAccessToNamedPipesAndShares")
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

// SetNetworkAccess_RestrictClientsAllowedToMakeRemoteCallsToSAM sets the value of NetworkAccess_RestrictClientsAllowedToMakeRemoteCallsToSAM for the instance
func (instance *MDM_Policy_Result01_LocalPoliciesSecurityOptions02) SetPropertyNetworkAccess_RestrictClientsAllowedToMakeRemoteCallsToSAM(value string) (err error) {
	return instance.SetProperty("NetworkAccess_RestrictClientsAllowedToMakeRemoteCallsToSAM", (value))
}

// GetNetworkAccess_RestrictClientsAllowedToMakeRemoteCallsToSAM gets the value of NetworkAccess_RestrictClientsAllowedToMakeRemoteCallsToSAM for the instance
func (instance *MDM_Policy_Result01_LocalPoliciesSecurityOptions02) GetPropertyNetworkAccess_RestrictClientsAllowedToMakeRemoteCallsToSAM() (value string, err error) {
	retValue, err := instance.GetProperty("NetworkAccess_RestrictClientsAllowedToMakeRemoteCallsToSAM")
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

// SetNetworkSecurity_AllowLocalSystemToUseComputerIdentityForNTLM sets the value of NetworkSecurity_AllowLocalSystemToUseComputerIdentityForNTLM for the instance
func (instance *MDM_Policy_Result01_LocalPoliciesSecurityOptions02) SetPropertyNetworkSecurity_AllowLocalSystemToUseComputerIdentityForNTLM(value int32) (err error) {
	return instance.SetProperty("NetworkSecurity_AllowLocalSystemToUseComputerIdentityForNTLM", (value))
}

// GetNetworkSecurity_AllowLocalSystemToUseComputerIdentityForNTLM gets the value of NetworkSecurity_AllowLocalSystemToUseComputerIdentityForNTLM for the instance
func (instance *MDM_Policy_Result01_LocalPoliciesSecurityOptions02) GetPropertyNetworkSecurity_AllowLocalSystemToUseComputerIdentityForNTLM() (value int32, err error) {
	retValue, err := instance.GetProperty("NetworkSecurity_AllowLocalSystemToUseComputerIdentityForNTLM")
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

// SetNetworkSecurity_AllowPKU2UAuthenticationRequests sets the value of NetworkSecurity_AllowPKU2UAuthenticationRequests for the instance
func (instance *MDM_Policy_Result01_LocalPoliciesSecurityOptions02) SetPropertyNetworkSecurity_AllowPKU2UAuthenticationRequests(value int32) (err error) {
	return instance.SetProperty("NetworkSecurity_AllowPKU2UAuthenticationRequests", (value))
}

// GetNetworkSecurity_AllowPKU2UAuthenticationRequests gets the value of NetworkSecurity_AllowPKU2UAuthenticationRequests for the instance
func (instance *MDM_Policy_Result01_LocalPoliciesSecurityOptions02) GetPropertyNetworkSecurity_AllowPKU2UAuthenticationRequests() (value int32, err error) {
	retValue, err := instance.GetProperty("NetworkSecurity_AllowPKU2UAuthenticationRequests")
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

// SetNetworkSecurity_DoNotStoreLANManagerHashValueOnNextPasswordChange sets the value of NetworkSecurity_DoNotStoreLANManagerHashValueOnNextPasswordChange for the instance
func (instance *MDM_Policy_Result01_LocalPoliciesSecurityOptions02) SetPropertyNetworkSecurity_DoNotStoreLANManagerHashValueOnNextPasswordChange(value int32) (err error) {
	return instance.SetProperty("NetworkSecurity_DoNotStoreLANManagerHashValueOnNextPasswordChange", (value))
}

// GetNetworkSecurity_DoNotStoreLANManagerHashValueOnNextPasswordChange gets the value of NetworkSecurity_DoNotStoreLANManagerHashValueOnNextPasswordChange for the instance
func (instance *MDM_Policy_Result01_LocalPoliciesSecurityOptions02) GetPropertyNetworkSecurity_DoNotStoreLANManagerHashValueOnNextPasswordChange() (value int32, err error) {
	retValue, err := instance.GetProperty("NetworkSecurity_DoNotStoreLANManagerHashValueOnNextPasswordChange")
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

// SetNetworkSecurity_LANManagerAuthenticationLevel sets the value of NetworkSecurity_LANManagerAuthenticationLevel for the instance
func (instance *MDM_Policy_Result01_LocalPoliciesSecurityOptions02) SetPropertyNetworkSecurity_LANManagerAuthenticationLevel(value int32) (err error) {
	return instance.SetProperty("NetworkSecurity_LANManagerAuthenticationLevel", (value))
}

// GetNetworkSecurity_LANManagerAuthenticationLevel gets the value of NetworkSecurity_LANManagerAuthenticationLevel for the instance
func (instance *MDM_Policy_Result01_LocalPoliciesSecurityOptions02) GetPropertyNetworkSecurity_LANManagerAuthenticationLevel() (value int32, err error) {
	retValue, err := instance.GetProperty("NetworkSecurity_LANManagerAuthenticationLevel")
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

// SetNetworkSecurity_MinimumSessionSecurityForNTLMSSPBasedClients sets the value of NetworkSecurity_MinimumSessionSecurityForNTLMSSPBasedClients for the instance
func (instance *MDM_Policy_Result01_LocalPoliciesSecurityOptions02) SetPropertyNetworkSecurity_MinimumSessionSecurityForNTLMSSPBasedClients(value int32) (err error) {
	return instance.SetProperty("NetworkSecurity_MinimumSessionSecurityForNTLMSSPBasedClients", (value))
}

// GetNetworkSecurity_MinimumSessionSecurityForNTLMSSPBasedClients gets the value of NetworkSecurity_MinimumSessionSecurityForNTLMSSPBasedClients for the instance
func (instance *MDM_Policy_Result01_LocalPoliciesSecurityOptions02) GetPropertyNetworkSecurity_MinimumSessionSecurityForNTLMSSPBasedClients() (value int32, err error) {
	retValue, err := instance.GetProperty("NetworkSecurity_MinimumSessionSecurityForNTLMSSPBasedClients")
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

// SetNetworkSecurity_MinimumSessionSecurityForNTLMSSPBasedServers sets the value of NetworkSecurity_MinimumSessionSecurityForNTLMSSPBasedServers for the instance
func (instance *MDM_Policy_Result01_LocalPoliciesSecurityOptions02) SetPropertyNetworkSecurity_MinimumSessionSecurityForNTLMSSPBasedServers(value int32) (err error) {
	return instance.SetProperty("NetworkSecurity_MinimumSessionSecurityForNTLMSSPBasedServers", (value))
}

// GetNetworkSecurity_MinimumSessionSecurityForNTLMSSPBasedServers gets the value of NetworkSecurity_MinimumSessionSecurityForNTLMSSPBasedServers for the instance
func (instance *MDM_Policy_Result01_LocalPoliciesSecurityOptions02) GetPropertyNetworkSecurity_MinimumSessionSecurityForNTLMSSPBasedServers() (value int32, err error) {
	retValue, err := instance.GetProperty("NetworkSecurity_MinimumSessionSecurityForNTLMSSPBasedServers")
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

// SetNetworkSecurity_RestrictNTLM_AddRemoteServerExceptionsForNTLMAuthentication sets the value of NetworkSecurity_RestrictNTLM_AddRemoteServerExceptionsForNTLMAuthentication for the instance
func (instance *MDM_Policy_Result01_LocalPoliciesSecurityOptions02) SetPropertyNetworkSecurity_RestrictNTLM_AddRemoteServerExceptionsForNTLMAuthentication(value string) (err error) {
	return instance.SetProperty("NetworkSecurity_RestrictNTLM_AddRemoteServerExceptionsForNTLMAuthentication", (value))
}

// GetNetworkSecurity_RestrictNTLM_AddRemoteServerExceptionsForNTLMAuthentication gets the value of NetworkSecurity_RestrictNTLM_AddRemoteServerExceptionsForNTLMAuthentication for the instance
func (instance *MDM_Policy_Result01_LocalPoliciesSecurityOptions02) GetPropertyNetworkSecurity_RestrictNTLM_AddRemoteServerExceptionsForNTLMAuthentication() (value string, err error) {
	retValue, err := instance.GetProperty("NetworkSecurity_RestrictNTLM_AddRemoteServerExceptionsForNTLMAuthentication")
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

// SetNetworkSecurity_RestrictNTLM_AuditIncomingNTLMTraffic sets the value of NetworkSecurity_RestrictNTLM_AuditIncomingNTLMTraffic for the instance
func (instance *MDM_Policy_Result01_LocalPoliciesSecurityOptions02) SetPropertyNetworkSecurity_RestrictNTLM_AuditIncomingNTLMTraffic(value int32) (err error) {
	return instance.SetProperty("NetworkSecurity_RestrictNTLM_AuditIncomingNTLMTraffic", (value))
}

// GetNetworkSecurity_RestrictNTLM_AuditIncomingNTLMTraffic gets the value of NetworkSecurity_RestrictNTLM_AuditIncomingNTLMTraffic for the instance
func (instance *MDM_Policy_Result01_LocalPoliciesSecurityOptions02) GetPropertyNetworkSecurity_RestrictNTLM_AuditIncomingNTLMTraffic() (value int32, err error) {
	retValue, err := instance.GetProperty("NetworkSecurity_RestrictNTLM_AuditIncomingNTLMTraffic")
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

// SetNetworkSecurity_RestrictNTLM_IncomingNTLMTraffic sets the value of NetworkSecurity_RestrictNTLM_IncomingNTLMTraffic for the instance
func (instance *MDM_Policy_Result01_LocalPoliciesSecurityOptions02) SetPropertyNetworkSecurity_RestrictNTLM_IncomingNTLMTraffic(value int32) (err error) {
	return instance.SetProperty("NetworkSecurity_RestrictNTLM_IncomingNTLMTraffic", (value))
}

// GetNetworkSecurity_RestrictNTLM_IncomingNTLMTraffic gets the value of NetworkSecurity_RestrictNTLM_IncomingNTLMTraffic for the instance
func (instance *MDM_Policy_Result01_LocalPoliciesSecurityOptions02) GetPropertyNetworkSecurity_RestrictNTLM_IncomingNTLMTraffic() (value int32, err error) {
	retValue, err := instance.GetProperty("NetworkSecurity_RestrictNTLM_IncomingNTLMTraffic")
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

// SetNetworkSecurity_RestrictNTLM_OutgoingNTLMTrafficToRemoteServers sets the value of NetworkSecurity_RestrictNTLM_OutgoingNTLMTrafficToRemoteServers for the instance
func (instance *MDM_Policy_Result01_LocalPoliciesSecurityOptions02) SetPropertyNetworkSecurity_RestrictNTLM_OutgoingNTLMTrafficToRemoteServers(value int32) (err error) {
	return instance.SetProperty("NetworkSecurity_RestrictNTLM_OutgoingNTLMTrafficToRemoteServers", (value))
}

// GetNetworkSecurity_RestrictNTLM_OutgoingNTLMTrafficToRemoteServers gets the value of NetworkSecurity_RestrictNTLM_OutgoingNTLMTrafficToRemoteServers for the instance
func (instance *MDM_Policy_Result01_LocalPoliciesSecurityOptions02) GetPropertyNetworkSecurity_RestrictNTLM_OutgoingNTLMTrafficToRemoteServers() (value int32, err error) {
	retValue, err := instance.GetProperty("NetworkSecurity_RestrictNTLM_OutgoingNTLMTrafficToRemoteServers")
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
func (instance *MDM_Policy_Result01_LocalPoliciesSecurityOptions02) SetPropertyParentID(value string) (err error) {
	return instance.SetProperty("ParentID", (value))
}

// GetParentID gets the value of ParentID for the instance
func (instance *MDM_Policy_Result01_LocalPoliciesSecurityOptions02) GetPropertyParentID() (value string, err error) {
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

// SetShutdown_AllowSystemToBeShutDownWithoutHavingToLogOn sets the value of Shutdown_AllowSystemToBeShutDownWithoutHavingToLogOn for the instance
func (instance *MDM_Policy_Result01_LocalPoliciesSecurityOptions02) SetPropertyShutdown_AllowSystemToBeShutDownWithoutHavingToLogOn(value int32) (err error) {
	return instance.SetProperty("Shutdown_AllowSystemToBeShutDownWithoutHavingToLogOn", (value))
}

// GetShutdown_AllowSystemToBeShutDownWithoutHavingToLogOn gets the value of Shutdown_AllowSystemToBeShutDownWithoutHavingToLogOn for the instance
func (instance *MDM_Policy_Result01_LocalPoliciesSecurityOptions02) GetPropertyShutdown_AllowSystemToBeShutDownWithoutHavingToLogOn() (value int32, err error) {
	retValue, err := instance.GetProperty("Shutdown_AllowSystemToBeShutDownWithoutHavingToLogOn")
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

// SetShutdown_ClearVirtualMemoryPageFile sets the value of Shutdown_ClearVirtualMemoryPageFile for the instance
func (instance *MDM_Policy_Result01_LocalPoliciesSecurityOptions02) SetPropertyShutdown_ClearVirtualMemoryPageFile(value int32) (err error) {
	return instance.SetProperty("Shutdown_ClearVirtualMemoryPageFile", (value))
}

// GetShutdown_ClearVirtualMemoryPageFile gets the value of Shutdown_ClearVirtualMemoryPageFile for the instance
func (instance *MDM_Policy_Result01_LocalPoliciesSecurityOptions02) GetPropertyShutdown_ClearVirtualMemoryPageFile() (value int32, err error) {
	retValue, err := instance.GetProperty("Shutdown_ClearVirtualMemoryPageFile")
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

// SetUserAccountControl_AllowUIAccessApplicationsToPromptForElevation sets the value of UserAccountControl_AllowUIAccessApplicationsToPromptForElevation for the instance
func (instance *MDM_Policy_Result01_LocalPoliciesSecurityOptions02) SetPropertyUserAccountControl_AllowUIAccessApplicationsToPromptForElevation(value int32) (err error) {
	return instance.SetProperty("UserAccountControl_AllowUIAccessApplicationsToPromptForElevation", (value))
}

// GetUserAccountControl_AllowUIAccessApplicationsToPromptForElevation gets the value of UserAccountControl_AllowUIAccessApplicationsToPromptForElevation for the instance
func (instance *MDM_Policy_Result01_LocalPoliciesSecurityOptions02) GetPropertyUserAccountControl_AllowUIAccessApplicationsToPromptForElevation() (value int32, err error) {
	retValue, err := instance.GetProperty("UserAccountControl_AllowUIAccessApplicationsToPromptForElevation")
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

// SetUserAccountControl_BehaviorOfTheElevationPromptForAdministrators sets the value of UserAccountControl_BehaviorOfTheElevationPromptForAdministrators for the instance
func (instance *MDM_Policy_Result01_LocalPoliciesSecurityOptions02) SetPropertyUserAccountControl_BehaviorOfTheElevationPromptForAdministrators(value int32) (err error) {
	return instance.SetProperty("UserAccountControl_BehaviorOfTheElevationPromptForAdministrators", (value))
}

// GetUserAccountControl_BehaviorOfTheElevationPromptForAdministrators gets the value of UserAccountControl_BehaviorOfTheElevationPromptForAdministrators for the instance
func (instance *MDM_Policy_Result01_LocalPoliciesSecurityOptions02) GetPropertyUserAccountControl_BehaviorOfTheElevationPromptForAdministrators() (value int32, err error) {
	retValue, err := instance.GetProperty("UserAccountControl_BehaviorOfTheElevationPromptForAdministrators")
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

// SetUserAccountControl_BehaviorOfTheElevationPromptForStandardUsers sets the value of UserAccountControl_BehaviorOfTheElevationPromptForStandardUsers for the instance
func (instance *MDM_Policy_Result01_LocalPoliciesSecurityOptions02) SetPropertyUserAccountControl_BehaviorOfTheElevationPromptForStandardUsers(value int32) (err error) {
	return instance.SetProperty("UserAccountControl_BehaviorOfTheElevationPromptForStandardUsers", (value))
}

// GetUserAccountControl_BehaviorOfTheElevationPromptForStandardUsers gets the value of UserAccountControl_BehaviorOfTheElevationPromptForStandardUsers for the instance
func (instance *MDM_Policy_Result01_LocalPoliciesSecurityOptions02) GetPropertyUserAccountControl_BehaviorOfTheElevationPromptForStandardUsers() (value int32, err error) {
	retValue, err := instance.GetProperty("UserAccountControl_BehaviorOfTheElevationPromptForStandardUsers")
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

// SetUserAccountControl_DetectApplicationInstallationsAndPromptForElevation sets the value of UserAccountControl_DetectApplicationInstallationsAndPromptForElevation for the instance
func (instance *MDM_Policy_Result01_LocalPoliciesSecurityOptions02) SetPropertyUserAccountControl_DetectApplicationInstallationsAndPromptForElevation(value int32) (err error) {
	return instance.SetProperty("UserAccountControl_DetectApplicationInstallationsAndPromptForElevation", (value))
}

// GetUserAccountControl_DetectApplicationInstallationsAndPromptForElevation gets the value of UserAccountControl_DetectApplicationInstallationsAndPromptForElevation for the instance
func (instance *MDM_Policy_Result01_LocalPoliciesSecurityOptions02) GetPropertyUserAccountControl_DetectApplicationInstallationsAndPromptForElevation() (value int32, err error) {
	retValue, err := instance.GetProperty("UserAccountControl_DetectApplicationInstallationsAndPromptForElevation")
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

// SetUserAccountControl_OnlyElevateExecutableFilesThatAreSignedAndValidated sets the value of UserAccountControl_OnlyElevateExecutableFilesThatAreSignedAndValidated for the instance
func (instance *MDM_Policy_Result01_LocalPoliciesSecurityOptions02) SetPropertyUserAccountControl_OnlyElevateExecutableFilesThatAreSignedAndValidated(value int32) (err error) {
	return instance.SetProperty("UserAccountControl_OnlyElevateExecutableFilesThatAreSignedAndValidated", (value))
}

// GetUserAccountControl_OnlyElevateExecutableFilesThatAreSignedAndValidated gets the value of UserAccountControl_OnlyElevateExecutableFilesThatAreSignedAndValidated for the instance
func (instance *MDM_Policy_Result01_LocalPoliciesSecurityOptions02) GetPropertyUserAccountControl_OnlyElevateExecutableFilesThatAreSignedAndValidated() (value int32, err error) {
	retValue, err := instance.GetProperty("UserAccountControl_OnlyElevateExecutableFilesThatAreSignedAndValidated")
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

// SetUserAccountControl_OnlyElevateUIAccessApplicationsThatAreInstalledInSecureLocations sets the value of UserAccountControl_OnlyElevateUIAccessApplicationsThatAreInstalledInSecureLocations for the instance
func (instance *MDM_Policy_Result01_LocalPoliciesSecurityOptions02) SetPropertyUserAccountControl_OnlyElevateUIAccessApplicationsThatAreInstalledInSecureLocations(value int32) (err error) {
	return instance.SetProperty("UserAccountControl_OnlyElevateUIAccessApplicationsThatAreInstalledInSecureLocations", (value))
}

// GetUserAccountControl_OnlyElevateUIAccessApplicationsThatAreInstalledInSecureLocations gets the value of UserAccountControl_OnlyElevateUIAccessApplicationsThatAreInstalledInSecureLocations for the instance
func (instance *MDM_Policy_Result01_LocalPoliciesSecurityOptions02) GetPropertyUserAccountControl_OnlyElevateUIAccessApplicationsThatAreInstalledInSecureLocations() (value int32, err error) {
	retValue, err := instance.GetProperty("UserAccountControl_OnlyElevateUIAccessApplicationsThatAreInstalledInSecureLocations")
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

// SetUserAccountControl_RunAllAdministratorsInAdminApprovalMode sets the value of UserAccountControl_RunAllAdministratorsInAdminApprovalMode for the instance
func (instance *MDM_Policy_Result01_LocalPoliciesSecurityOptions02) SetPropertyUserAccountControl_RunAllAdministratorsInAdminApprovalMode(value int32) (err error) {
	return instance.SetProperty("UserAccountControl_RunAllAdministratorsInAdminApprovalMode", (value))
}

// GetUserAccountControl_RunAllAdministratorsInAdminApprovalMode gets the value of UserAccountControl_RunAllAdministratorsInAdminApprovalMode for the instance
func (instance *MDM_Policy_Result01_LocalPoliciesSecurityOptions02) GetPropertyUserAccountControl_RunAllAdministratorsInAdminApprovalMode() (value int32, err error) {
	retValue, err := instance.GetProperty("UserAccountControl_RunAllAdministratorsInAdminApprovalMode")
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

// SetUserAccountControl_SwitchToTheSecureDesktopWhenPromptingForElevation sets the value of UserAccountControl_SwitchToTheSecureDesktopWhenPromptingForElevation for the instance
func (instance *MDM_Policy_Result01_LocalPoliciesSecurityOptions02) SetPropertyUserAccountControl_SwitchToTheSecureDesktopWhenPromptingForElevation(value int32) (err error) {
	return instance.SetProperty("UserAccountControl_SwitchToTheSecureDesktopWhenPromptingForElevation", (value))
}

// GetUserAccountControl_SwitchToTheSecureDesktopWhenPromptingForElevation gets the value of UserAccountControl_SwitchToTheSecureDesktopWhenPromptingForElevation for the instance
func (instance *MDM_Policy_Result01_LocalPoliciesSecurityOptions02) GetPropertyUserAccountControl_SwitchToTheSecureDesktopWhenPromptingForElevation() (value int32, err error) {
	retValue, err := instance.GetProperty("UserAccountControl_SwitchToTheSecureDesktopWhenPromptingForElevation")
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

// SetUserAccountControl_UseAdminApprovalMode sets the value of UserAccountControl_UseAdminApprovalMode for the instance
func (instance *MDM_Policy_Result01_LocalPoliciesSecurityOptions02) SetPropertyUserAccountControl_UseAdminApprovalMode(value int32) (err error) {
	return instance.SetProperty("UserAccountControl_UseAdminApprovalMode", (value))
}

// GetUserAccountControl_UseAdminApprovalMode gets the value of UserAccountControl_UseAdminApprovalMode for the instance
func (instance *MDM_Policy_Result01_LocalPoliciesSecurityOptions02) GetPropertyUserAccountControl_UseAdminApprovalMode() (value int32, err error) {
	retValue, err := instance.GetProperty("UserAccountControl_UseAdminApprovalMode")
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

// SetUserAccountControl_VirtualizeFileAndRegistryWriteFailuresToPerUserLocations sets the value of UserAccountControl_VirtualizeFileAndRegistryWriteFailuresToPerUserLocations for the instance
func (instance *MDM_Policy_Result01_LocalPoliciesSecurityOptions02) SetPropertyUserAccountControl_VirtualizeFileAndRegistryWriteFailuresToPerUserLocations(value int32) (err error) {
	return instance.SetProperty("UserAccountControl_VirtualizeFileAndRegistryWriteFailuresToPerUserLocations", (value))
}

// GetUserAccountControl_VirtualizeFileAndRegistryWriteFailuresToPerUserLocations gets the value of UserAccountControl_VirtualizeFileAndRegistryWriteFailuresToPerUserLocations for the instance
func (instance *MDM_Policy_Result01_LocalPoliciesSecurityOptions02) GetPropertyUserAccountControl_VirtualizeFileAndRegistryWriteFailuresToPerUserLocations() (value int32, err error) {
	retValue, err := instance.GetProperty("UserAccountControl_VirtualizeFileAndRegistryWriteFailuresToPerUserLocations")
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
