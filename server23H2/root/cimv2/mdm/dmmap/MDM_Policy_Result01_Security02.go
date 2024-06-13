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

// MDM_Policy_Result01_Security02 struct
type MDM_Policy_Result01_Security02 struct {
	*cim.WmiInstance

	//
	AllowAddProvisioningPackage int32

	//
	AllowRemoveProvisioningPackage int32

	//
	ClearTPMIfNotReady int32

	//
	ConfigureWindowsPasswords int32

	//
	InstanceID string

	//
	ParentID string

	//
	PreventAutomaticDeviceEncryptionForAzureADJoinedDevices int32

	//
	RecoveryEnvironmentAuthentication int32

	//
	RequireDeviceEncryption int32

	//
	RequireProvisioningPackageSignature int32

	//
	RequireRetrieveHealthCertificateOnBoot int32
}

func NewMDM_Policy_Result01_Security02Ex1(instance *cim.WmiInstance) (newInstance *MDM_Policy_Result01_Security02, err error) {
	tmp, err := instance, nil

	if err != nil {
		return
	}
	newInstance = &MDM_Policy_Result01_Security02{
		WmiInstance: tmp,
	}
	return
}

func NewMDM_Policy_Result01_Security02Ex6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *MDM_Policy_Result01_Security02, err error) {
	tmp, err := instance.GetWmiInstance(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &MDM_Policy_Result01_Security02{
		WmiInstance: tmp,
	}
	return
}

// SetAllowAddProvisioningPackage sets the value of AllowAddProvisioningPackage for the instance
func (instance *MDM_Policy_Result01_Security02) SetPropertyAllowAddProvisioningPackage(value int32) (err error) {
	return instance.SetProperty("AllowAddProvisioningPackage", (value))
}

// GetAllowAddProvisioningPackage gets the value of AllowAddProvisioningPackage for the instance
func (instance *MDM_Policy_Result01_Security02) GetPropertyAllowAddProvisioningPackage() (value int32, err error) {
	retValue, err := instance.GetProperty("AllowAddProvisioningPackage")
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

// SetAllowRemoveProvisioningPackage sets the value of AllowRemoveProvisioningPackage for the instance
func (instance *MDM_Policy_Result01_Security02) SetPropertyAllowRemoveProvisioningPackage(value int32) (err error) {
	return instance.SetProperty("AllowRemoveProvisioningPackage", (value))
}

// GetAllowRemoveProvisioningPackage gets the value of AllowRemoveProvisioningPackage for the instance
func (instance *MDM_Policy_Result01_Security02) GetPropertyAllowRemoveProvisioningPackage() (value int32, err error) {
	retValue, err := instance.GetProperty("AllowRemoveProvisioningPackage")
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

// SetClearTPMIfNotReady sets the value of ClearTPMIfNotReady for the instance
func (instance *MDM_Policy_Result01_Security02) SetPropertyClearTPMIfNotReady(value int32) (err error) {
	return instance.SetProperty("ClearTPMIfNotReady", (value))
}

// GetClearTPMIfNotReady gets the value of ClearTPMIfNotReady for the instance
func (instance *MDM_Policy_Result01_Security02) GetPropertyClearTPMIfNotReady() (value int32, err error) {
	retValue, err := instance.GetProperty("ClearTPMIfNotReady")
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

// SetConfigureWindowsPasswords sets the value of ConfigureWindowsPasswords for the instance
func (instance *MDM_Policy_Result01_Security02) SetPropertyConfigureWindowsPasswords(value int32) (err error) {
	return instance.SetProperty("ConfigureWindowsPasswords", (value))
}

// GetConfigureWindowsPasswords gets the value of ConfigureWindowsPasswords for the instance
func (instance *MDM_Policy_Result01_Security02) GetPropertyConfigureWindowsPasswords() (value int32, err error) {
	retValue, err := instance.GetProperty("ConfigureWindowsPasswords")
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
func (instance *MDM_Policy_Result01_Security02) SetPropertyInstanceID(value string) (err error) {
	return instance.SetProperty("InstanceID", (value))
}

// GetInstanceID gets the value of InstanceID for the instance
func (instance *MDM_Policy_Result01_Security02) GetPropertyInstanceID() (value string, err error) {
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

// SetParentID sets the value of ParentID for the instance
func (instance *MDM_Policy_Result01_Security02) SetPropertyParentID(value string) (err error) {
	return instance.SetProperty("ParentID", (value))
}

// GetParentID gets the value of ParentID for the instance
func (instance *MDM_Policy_Result01_Security02) GetPropertyParentID() (value string, err error) {
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

// SetPreventAutomaticDeviceEncryptionForAzureADJoinedDevices sets the value of PreventAutomaticDeviceEncryptionForAzureADJoinedDevices for the instance
func (instance *MDM_Policy_Result01_Security02) SetPropertyPreventAutomaticDeviceEncryptionForAzureADJoinedDevices(value int32) (err error) {
	return instance.SetProperty("PreventAutomaticDeviceEncryptionForAzureADJoinedDevices", (value))
}

// GetPreventAutomaticDeviceEncryptionForAzureADJoinedDevices gets the value of PreventAutomaticDeviceEncryptionForAzureADJoinedDevices for the instance
func (instance *MDM_Policy_Result01_Security02) GetPropertyPreventAutomaticDeviceEncryptionForAzureADJoinedDevices() (value int32, err error) {
	retValue, err := instance.GetProperty("PreventAutomaticDeviceEncryptionForAzureADJoinedDevices")
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

// SetRecoveryEnvironmentAuthentication sets the value of RecoveryEnvironmentAuthentication for the instance
func (instance *MDM_Policy_Result01_Security02) SetPropertyRecoveryEnvironmentAuthentication(value int32) (err error) {
	return instance.SetProperty("RecoveryEnvironmentAuthentication", (value))
}

// GetRecoveryEnvironmentAuthentication gets the value of RecoveryEnvironmentAuthentication for the instance
func (instance *MDM_Policy_Result01_Security02) GetPropertyRecoveryEnvironmentAuthentication() (value int32, err error) {
	retValue, err := instance.GetProperty("RecoveryEnvironmentAuthentication")
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

// SetRequireDeviceEncryption sets the value of RequireDeviceEncryption for the instance
func (instance *MDM_Policy_Result01_Security02) SetPropertyRequireDeviceEncryption(value int32) (err error) {
	return instance.SetProperty("RequireDeviceEncryption", (value))
}

// GetRequireDeviceEncryption gets the value of RequireDeviceEncryption for the instance
func (instance *MDM_Policy_Result01_Security02) GetPropertyRequireDeviceEncryption() (value int32, err error) {
	retValue, err := instance.GetProperty("RequireDeviceEncryption")
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

// SetRequireProvisioningPackageSignature sets the value of RequireProvisioningPackageSignature for the instance
func (instance *MDM_Policy_Result01_Security02) SetPropertyRequireProvisioningPackageSignature(value int32) (err error) {
	return instance.SetProperty("RequireProvisioningPackageSignature", (value))
}

// GetRequireProvisioningPackageSignature gets the value of RequireProvisioningPackageSignature for the instance
func (instance *MDM_Policy_Result01_Security02) GetPropertyRequireProvisioningPackageSignature() (value int32, err error) {
	retValue, err := instance.GetProperty("RequireProvisioningPackageSignature")
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

// SetRequireRetrieveHealthCertificateOnBoot sets the value of RequireRetrieveHealthCertificateOnBoot for the instance
func (instance *MDM_Policy_Result01_Security02) SetPropertyRequireRetrieveHealthCertificateOnBoot(value int32) (err error) {
	return instance.SetProperty("RequireRetrieveHealthCertificateOnBoot", (value))
}

// GetRequireRetrieveHealthCertificateOnBoot gets the value of RequireRetrieveHealthCertificateOnBoot for the instance
func (instance *MDM_Policy_Result01_Security02) GetPropertyRequireRetrieveHealthCertificateOnBoot() (value int32, err error) {
	retValue, err := instance.GetProperty("RequireRetrieveHealthCertificateOnBoot")
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
