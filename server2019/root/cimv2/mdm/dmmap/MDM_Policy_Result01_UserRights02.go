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

// MDM_Policy_Result01_UserRights02 struct
type MDM_Policy_Result01_UserRights02 struct {
	*cim.WmiInstance

	//
	AccessCredentialManagerAsTrustedCaller string

	//
	AccessFromNetwork string

	//
	ActAsPartOfTheOperatingSystem string

	//
	AllowLocalLogOn string

	//
	BackupFilesAndDirectories string

	//
	ChangeSystemTime string

	//
	CreateGlobalObjects string

	//
	CreatePageFile string

	//
	CreatePermanentSharedObjects string

	//
	CreateSymbolicLinks string

	//
	CreateToken string

	//
	DebugPrograms string

	//
	DenyAccessFromNetwork string

	//
	DenyLocalLogOn string

	//
	DenyRemoteDesktopServicesLogOn string

	//
	EnableDelegation string

	//
	GenerateSecurityAudits string

	//
	ImpersonateClient string

	//
	IncreaseSchedulingPriority string

	//
	InstanceID string

	//
	LoadUnloadDeviceDrivers string

	//
	LockMemory string

	//
	ManageAuditingAndSecurityLog string

	//
	ManageVolume string

	//
	ModifyFirmwareEnvironment string

	//
	ModifyObjectLabel string

	//
	ParentID string

	//
	ProfileSingleProcess string

	//
	RemoteShutdown string

	//
	RestoreFilesAndDirectories string

	//
	TakeOwnership string
}

func NewMDM_Policy_Result01_UserRights02Ex1(instance *cim.WmiInstance) (newInstance *MDM_Policy_Result01_UserRights02, err error) {
	tmp, err := instance, nil

	if err != nil {
		return
	}
	newInstance = &MDM_Policy_Result01_UserRights02{
		WmiInstance: tmp,
	}
	return
}

func NewMDM_Policy_Result01_UserRights02Ex6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *MDM_Policy_Result01_UserRights02, err error) {
	tmp, err := instance.GetWmiInstance(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &MDM_Policy_Result01_UserRights02{
		WmiInstance: tmp,
	}
	return
}

// SetAccessCredentialManagerAsTrustedCaller sets the value of AccessCredentialManagerAsTrustedCaller for the instance
func (instance *MDM_Policy_Result01_UserRights02) SetPropertyAccessCredentialManagerAsTrustedCaller(value string) (err error) {
	return instance.SetProperty("AccessCredentialManagerAsTrustedCaller", (value))
}

// GetAccessCredentialManagerAsTrustedCaller gets the value of AccessCredentialManagerAsTrustedCaller for the instance
func (instance *MDM_Policy_Result01_UserRights02) GetPropertyAccessCredentialManagerAsTrustedCaller() (value string, err error) {
	retValue, err := instance.GetProperty("AccessCredentialManagerAsTrustedCaller")
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

// SetAccessFromNetwork sets the value of AccessFromNetwork for the instance
func (instance *MDM_Policy_Result01_UserRights02) SetPropertyAccessFromNetwork(value string) (err error) {
	return instance.SetProperty("AccessFromNetwork", (value))
}

// GetAccessFromNetwork gets the value of AccessFromNetwork for the instance
func (instance *MDM_Policy_Result01_UserRights02) GetPropertyAccessFromNetwork() (value string, err error) {
	retValue, err := instance.GetProperty("AccessFromNetwork")
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

// SetActAsPartOfTheOperatingSystem sets the value of ActAsPartOfTheOperatingSystem for the instance
func (instance *MDM_Policy_Result01_UserRights02) SetPropertyActAsPartOfTheOperatingSystem(value string) (err error) {
	return instance.SetProperty("ActAsPartOfTheOperatingSystem", (value))
}

// GetActAsPartOfTheOperatingSystem gets the value of ActAsPartOfTheOperatingSystem for the instance
func (instance *MDM_Policy_Result01_UserRights02) GetPropertyActAsPartOfTheOperatingSystem() (value string, err error) {
	retValue, err := instance.GetProperty("ActAsPartOfTheOperatingSystem")
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

// SetAllowLocalLogOn sets the value of AllowLocalLogOn for the instance
func (instance *MDM_Policy_Result01_UserRights02) SetPropertyAllowLocalLogOn(value string) (err error) {
	return instance.SetProperty("AllowLocalLogOn", (value))
}

// GetAllowLocalLogOn gets the value of AllowLocalLogOn for the instance
func (instance *MDM_Policy_Result01_UserRights02) GetPropertyAllowLocalLogOn() (value string, err error) {
	retValue, err := instance.GetProperty("AllowLocalLogOn")
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

// SetBackupFilesAndDirectories sets the value of BackupFilesAndDirectories for the instance
func (instance *MDM_Policy_Result01_UserRights02) SetPropertyBackupFilesAndDirectories(value string) (err error) {
	return instance.SetProperty("BackupFilesAndDirectories", (value))
}

// GetBackupFilesAndDirectories gets the value of BackupFilesAndDirectories for the instance
func (instance *MDM_Policy_Result01_UserRights02) GetPropertyBackupFilesAndDirectories() (value string, err error) {
	retValue, err := instance.GetProperty("BackupFilesAndDirectories")
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

// SetChangeSystemTime sets the value of ChangeSystemTime for the instance
func (instance *MDM_Policy_Result01_UserRights02) SetPropertyChangeSystemTime(value string) (err error) {
	return instance.SetProperty("ChangeSystemTime", (value))
}

// GetChangeSystemTime gets the value of ChangeSystemTime for the instance
func (instance *MDM_Policy_Result01_UserRights02) GetPropertyChangeSystemTime() (value string, err error) {
	retValue, err := instance.GetProperty("ChangeSystemTime")
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

// SetCreateGlobalObjects sets the value of CreateGlobalObjects for the instance
func (instance *MDM_Policy_Result01_UserRights02) SetPropertyCreateGlobalObjects(value string) (err error) {
	return instance.SetProperty("CreateGlobalObjects", (value))
}

// GetCreateGlobalObjects gets the value of CreateGlobalObjects for the instance
func (instance *MDM_Policy_Result01_UserRights02) GetPropertyCreateGlobalObjects() (value string, err error) {
	retValue, err := instance.GetProperty("CreateGlobalObjects")
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

// SetCreatePageFile sets the value of CreatePageFile for the instance
func (instance *MDM_Policy_Result01_UserRights02) SetPropertyCreatePageFile(value string) (err error) {
	return instance.SetProperty("CreatePageFile", (value))
}

// GetCreatePageFile gets the value of CreatePageFile for the instance
func (instance *MDM_Policy_Result01_UserRights02) GetPropertyCreatePageFile() (value string, err error) {
	retValue, err := instance.GetProperty("CreatePageFile")
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

// SetCreatePermanentSharedObjects sets the value of CreatePermanentSharedObjects for the instance
func (instance *MDM_Policy_Result01_UserRights02) SetPropertyCreatePermanentSharedObjects(value string) (err error) {
	return instance.SetProperty("CreatePermanentSharedObjects", (value))
}

// GetCreatePermanentSharedObjects gets the value of CreatePermanentSharedObjects for the instance
func (instance *MDM_Policy_Result01_UserRights02) GetPropertyCreatePermanentSharedObjects() (value string, err error) {
	retValue, err := instance.GetProperty("CreatePermanentSharedObjects")
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

// SetCreateSymbolicLinks sets the value of CreateSymbolicLinks for the instance
func (instance *MDM_Policy_Result01_UserRights02) SetPropertyCreateSymbolicLinks(value string) (err error) {
	return instance.SetProperty("CreateSymbolicLinks", (value))
}

// GetCreateSymbolicLinks gets the value of CreateSymbolicLinks for the instance
func (instance *MDM_Policy_Result01_UserRights02) GetPropertyCreateSymbolicLinks() (value string, err error) {
	retValue, err := instance.GetProperty("CreateSymbolicLinks")
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

// SetCreateToken sets the value of CreateToken for the instance
func (instance *MDM_Policy_Result01_UserRights02) SetPropertyCreateToken(value string) (err error) {
	return instance.SetProperty("CreateToken", (value))
}

// GetCreateToken gets the value of CreateToken for the instance
func (instance *MDM_Policy_Result01_UserRights02) GetPropertyCreateToken() (value string, err error) {
	retValue, err := instance.GetProperty("CreateToken")
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

// SetDebugPrograms sets the value of DebugPrograms for the instance
func (instance *MDM_Policy_Result01_UserRights02) SetPropertyDebugPrograms(value string) (err error) {
	return instance.SetProperty("DebugPrograms", (value))
}

// GetDebugPrograms gets the value of DebugPrograms for the instance
func (instance *MDM_Policy_Result01_UserRights02) GetPropertyDebugPrograms() (value string, err error) {
	retValue, err := instance.GetProperty("DebugPrograms")
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

// SetDenyAccessFromNetwork sets the value of DenyAccessFromNetwork for the instance
func (instance *MDM_Policy_Result01_UserRights02) SetPropertyDenyAccessFromNetwork(value string) (err error) {
	return instance.SetProperty("DenyAccessFromNetwork", (value))
}

// GetDenyAccessFromNetwork gets the value of DenyAccessFromNetwork for the instance
func (instance *MDM_Policy_Result01_UserRights02) GetPropertyDenyAccessFromNetwork() (value string, err error) {
	retValue, err := instance.GetProperty("DenyAccessFromNetwork")
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

// SetDenyLocalLogOn sets the value of DenyLocalLogOn for the instance
func (instance *MDM_Policy_Result01_UserRights02) SetPropertyDenyLocalLogOn(value string) (err error) {
	return instance.SetProperty("DenyLocalLogOn", (value))
}

// GetDenyLocalLogOn gets the value of DenyLocalLogOn for the instance
func (instance *MDM_Policy_Result01_UserRights02) GetPropertyDenyLocalLogOn() (value string, err error) {
	retValue, err := instance.GetProperty("DenyLocalLogOn")
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

// SetDenyRemoteDesktopServicesLogOn sets the value of DenyRemoteDesktopServicesLogOn for the instance
func (instance *MDM_Policy_Result01_UserRights02) SetPropertyDenyRemoteDesktopServicesLogOn(value string) (err error) {
	return instance.SetProperty("DenyRemoteDesktopServicesLogOn", (value))
}

// GetDenyRemoteDesktopServicesLogOn gets the value of DenyRemoteDesktopServicesLogOn for the instance
func (instance *MDM_Policy_Result01_UserRights02) GetPropertyDenyRemoteDesktopServicesLogOn() (value string, err error) {
	retValue, err := instance.GetProperty("DenyRemoteDesktopServicesLogOn")
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

// SetEnableDelegation sets the value of EnableDelegation for the instance
func (instance *MDM_Policy_Result01_UserRights02) SetPropertyEnableDelegation(value string) (err error) {
	return instance.SetProperty("EnableDelegation", (value))
}

// GetEnableDelegation gets the value of EnableDelegation for the instance
func (instance *MDM_Policy_Result01_UserRights02) GetPropertyEnableDelegation() (value string, err error) {
	retValue, err := instance.GetProperty("EnableDelegation")
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

// SetGenerateSecurityAudits sets the value of GenerateSecurityAudits for the instance
func (instance *MDM_Policy_Result01_UserRights02) SetPropertyGenerateSecurityAudits(value string) (err error) {
	return instance.SetProperty("GenerateSecurityAudits", (value))
}

// GetGenerateSecurityAudits gets the value of GenerateSecurityAudits for the instance
func (instance *MDM_Policy_Result01_UserRights02) GetPropertyGenerateSecurityAudits() (value string, err error) {
	retValue, err := instance.GetProperty("GenerateSecurityAudits")
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

// SetImpersonateClient sets the value of ImpersonateClient for the instance
func (instance *MDM_Policy_Result01_UserRights02) SetPropertyImpersonateClient(value string) (err error) {
	return instance.SetProperty("ImpersonateClient", (value))
}

// GetImpersonateClient gets the value of ImpersonateClient for the instance
func (instance *MDM_Policy_Result01_UserRights02) GetPropertyImpersonateClient() (value string, err error) {
	retValue, err := instance.GetProperty("ImpersonateClient")
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

// SetIncreaseSchedulingPriority sets the value of IncreaseSchedulingPriority for the instance
func (instance *MDM_Policy_Result01_UserRights02) SetPropertyIncreaseSchedulingPriority(value string) (err error) {
	return instance.SetProperty("IncreaseSchedulingPriority", (value))
}

// GetIncreaseSchedulingPriority gets the value of IncreaseSchedulingPriority for the instance
func (instance *MDM_Policy_Result01_UserRights02) GetPropertyIncreaseSchedulingPriority() (value string, err error) {
	retValue, err := instance.GetProperty("IncreaseSchedulingPriority")
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
func (instance *MDM_Policy_Result01_UserRights02) SetPropertyInstanceID(value string) (err error) {
	return instance.SetProperty("InstanceID", (value))
}

// GetInstanceID gets the value of InstanceID for the instance
func (instance *MDM_Policy_Result01_UserRights02) GetPropertyInstanceID() (value string, err error) {
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

// SetLoadUnloadDeviceDrivers sets the value of LoadUnloadDeviceDrivers for the instance
func (instance *MDM_Policy_Result01_UserRights02) SetPropertyLoadUnloadDeviceDrivers(value string) (err error) {
	return instance.SetProperty("LoadUnloadDeviceDrivers", (value))
}

// GetLoadUnloadDeviceDrivers gets the value of LoadUnloadDeviceDrivers for the instance
func (instance *MDM_Policy_Result01_UserRights02) GetPropertyLoadUnloadDeviceDrivers() (value string, err error) {
	retValue, err := instance.GetProperty("LoadUnloadDeviceDrivers")
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

// SetLockMemory sets the value of LockMemory for the instance
func (instance *MDM_Policy_Result01_UserRights02) SetPropertyLockMemory(value string) (err error) {
	return instance.SetProperty("LockMemory", (value))
}

// GetLockMemory gets the value of LockMemory for the instance
func (instance *MDM_Policy_Result01_UserRights02) GetPropertyLockMemory() (value string, err error) {
	retValue, err := instance.GetProperty("LockMemory")
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

// SetManageAuditingAndSecurityLog sets the value of ManageAuditingAndSecurityLog for the instance
func (instance *MDM_Policy_Result01_UserRights02) SetPropertyManageAuditingAndSecurityLog(value string) (err error) {
	return instance.SetProperty("ManageAuditingAndSecurityLog", (value))
}

// GetManageAuditingAndSecurityLog gets the value of ManageAuditingAndSecurityLog for the instance
func (instance *MDM_Policy_Result01_UserRights02) GetPropertyManageAuditingAndSecurityLog() (value string, err error) {
	retValue, err := instance.GetProperty("ManageAuditingAndSecurityLog")
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

// SetManageVolume sets the value of ManageVolume for the instance
func (instance *MDM_Policy_Result01_UserRights02) SetPropertyManageVolume(value string) (err error) {
	return instance.SetProperty("ManageVolume", (value))
}

// GetManageVolume gets the value of ManageVolume for the instance
func (instance *MDM_Policy_Result01_UserRights02) GetPropertyManageVolume() (value string, err error) {
	retValue, err := instance.GetProperty("ManageVolume")
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

// SetModifyFirmwareEnvironment sets the value of ModifyFirmwareEnvironment for the instance
func (instance *MDM_Policy_Result01_UserRights02) SetPropertyModifyFirmwareEnvironment(value string) (err error) {
	return instance.SetProperty("ModifyFirmwareEnvironment", (value))
}

// GetModifyFirmwareEnvironment gets the value of ModifyFirmwareEnvironment for the instance
func (instance *MDM_Policy_Result01_UserRights02) GetPropertyModifyFirmwareEnvironment() (value string, err error) {
	retValue, err := instance.GetProperty("ModifyFirmwareEnvironment")
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

// SetModifyObjectLabel sets the value of ModifyObjectLabel for the instance
func (instance *MDM_Policy_Result01_UserRights02) SetPropertyModifyObjectLabel(value string) (err error) {
	return instance.SetProperty("ModifyObjectLabel", (value))
}

// GetModifyObjectLabel gets the value of ModifyObjectLabel for the instance
func (instance *MDM_Policy_Result01_UserRights02) GetPropertyModifyObjectLabel() (value string, err error) {
	retValue, err := instance.GetProperty("ModifyObjectLabel")
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

// SetParentID sets the value of ParentID for the instance
func (instance *MDM_Policy_Result01_UserRights02) SetPropertyParentID(value string) (err error) {
	return instance.SetProperty("ParentID", (value))
}

// GetParentID gets the value of ParentID for the instance
func (instance *MDM_Policy_Result01_UserRights02) GetPropertyParentID() (value string, err error) {
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

// SetProfileSingleProcess sets the value of ProfileSingleProcess for the instance
func (instance *MDM_Policy_Result01_UserRights02) SetPropertyProfileSingleProcess(value string) (err error) {
	return instance.SetProperty("ProfileSingleProcess", (value))
}

// GetProfileSingleProcess gets the value of ProfileSingleProcess for the instance
func (instance *MDM_Policy_Result01_UserRights02) GetPropertyProfileSingleProcess() (value string, err error) {
	retValue, err := instance.GetProperty("ProfileSingleProcess")
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

// SetRemoteShutdown sets the value of RemoteShutdown for the instance
func (instance *MDM_Policy_Result01_UserRights02) SetPropertyRemoteShutdown(value string) (err error) {
	return instance.SetProperty("RemoteShutdown", (value))
}

// GetRemoteShutdown gets the value of RemoteShutdown for the instance
func (instance *MDM_Policy_Result01_UserRights02) GetPropertyRemoteShutdown() (value string, err error) {
	retValue, err := instance.GetProperty("RemoteShutdown")
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

// SetRestoreFilesAndDirectories sets the value of RestoreFilesAndDirectories for the instance
func (instance *MDM_Policy_Result01_UserRights02) SetPropertyRestoreFilesAndDirectories(value string) (err error) {
	return instance.SetProperty("RestoreFilesAndDirectories", (value))
}

// GetRestoreFilesAndDirectories gets the value of RestoreFilesAndDirectories for the instance
func (instance *MDM_Policy_Result01_UserRights02) GetPropertyRestoreFilesAndDirectories() (value string, err error) {
	retValue, err := instance.GetProperty("RestoreFilesAndDirectories")
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

// SetTakeOwnership sets the value of TakeOwnership for the instance
func (instance *MDM_Policy_Result01_UserRights02) SetPropertyTakeOwnership(value string) (err error) {
	return instance.SetProperty("TakeOwnership", (value))
}

// GetTakeOwnership gets the value of TakeOwnership for the instance
func (instance *MDM_Policy_Result01_UserRights02) GetPropertyTakeOwnership() (value string, err error) {
	retValue, err := instance.GetProperty("TakeOwnership")
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
