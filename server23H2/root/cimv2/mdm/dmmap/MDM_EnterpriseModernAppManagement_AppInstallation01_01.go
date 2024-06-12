// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 6/12/2024 using wmigen
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

// MDM_EnterpriseModernAppManagement_AppInstallation01_01 struct
type MDM_EnterpriseModernAppManagement_AppInstallation01_01 struct {
	*cim.WmiInstance

	//
	HostedInstall string

	//
	InstanceID string

	//
	LastError int32

	//
	LastErrorDesc string

	//
	ParentID string

	//
	ProgressStatus int32

	//
	Status int32

	//
	StoreInstall string
}

func NewMDM_EnterpriseModernAppManagement_AppInstallation01_01Ex1(instance *cim.WmiInstance) (newInstance *MDM_EnterpriseModernAppManagement_AppInstallation01_01, err error) {
	tmp, err := instance, nil

	if err != nil {
		return
	}
	newInstance = &MDM_EnterpriseModernAppManagement_AppInstallation01_01{
		WmiInstance: tmp,
	}
	return
}

func NewMDM_EnterpriseModernAppManagement_AppInstallation01_01Ex6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *MDM_EnterpriseModernAppManagement_AppInstallation01_01, err error) {
	tmp, err := instance.GetWmiInstance(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &MDM_EnterpriseModernAppManagement_AppInstallation01_01{
		WmiInstance: tmp,
	}
	return
}

// SetHostedInstall sets the value of HostedInstall for the instance
func (instance *MDM_EnterpriseModernAppManagement_AppInstallation01_01) SetPropertyHostedInstall(value string) (err error) {
	return instance.SetProperty("HostedInstall", (value))
}

// GetHostedInstall gets the value of HostedInstall for the instance
func (instance *MDM_EnterpriseModernAppManagement_AppInstallation01_01) GetPropertyHostedInstall() (value string, err error) {
	retValue, err := instance.GetProperty("HostedInstall")
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
func (instance *MDM_EnterpriseModernAppManagement_AppInstallation01_01) SetPropertyInstanceID(value string) (err error) {
	return instance.SetProperty("InstanceID", (value))
}

// GetInstanceID gets the value of InstanceID for the instance
func (instance *MDM_EnterpriseModernAppManagement_AppInstallation01_01) GetPropertyInstanceID() (value string, err error) {
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

// SetLastError sets the value of LastError for the instance
func (instance *MDM_EnterpriseModernAppManagement_AppInstallation01_01) SetPropertyLastError(value int32) (err error) {
	return instance.SetProperty("LastError", (value))
}

// GetLastError gets the value of LastError for the instance
func (instance *MDM_EnterpriseModernAppManagement_AppInstallation01_01) GetPropertyLastError() (value int32, err error) {
	retValue, err := instance.GetProperty("LastError")
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

// SetLastErrorDesc sets the value of LastErrorDesc for the instance
func (instance *MDM_EnterpriseModernAppManagement_AppInstallation01_01) SetPropertyLastErrorDesc(value string) (err error) {
	return instance.SetProperty("LastErrorDesc", (value))
}

// GetLastErrorDesc gets the value of LastErrorDesc for the instance
func (instance *MDM_EnterpriseModernAppManagement_AppInstallation01_01) GetPropertyLastErrorDesc() (value string, err error) {
	retValue, err := instance.GetProperty("LastErrorDesc")
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
func (instance *MDM_EnterpriseModernAppManagement_AppInstallation01_01) SetPropertyParentID(value string) (err error) {
	return instance.SetProperty("ParentID", (value))
}

// GetParentID gets the value of ParentID for the instance
func (instance *MDM_EnterpriseModernAppManagement_AppInstallation01_01) GetPropertyParentID() (value string, err error) {
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

// SetProgressStatus sets the value of ProgressStatus for the instance
func (instance *MDM_EnterpriseModernAppManagement_AppInstallation01_01) SetPropertyProgressStatus(value int32) (err error) {
	return instance.SetProperty("ProgressStatus", (value))
}

// GetProgressStatus gets the value of ProgressStatus for the instance
func (instance *MDM_EnterpriseModernAppManagement_AppInstallation01_01) GetPropertyProgressStatus() (value int32, err error) {
	retValue, err := instance.GetProperty("ProgressStatus")
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

// SetStatus sets the value of Status for the instance
func (instance *MDM_EnterpriseModernAppManagement_AppInstallation01_01) SetPropertyStatus(value int32) (err error) {
	return instance.SetProperty("Status", (value))
}

// GetStatus gets the value of Status for the instance
func (instance *MDM_EnterpriseModernAppManagement_AppInstallation01_01) GetPropertyStatus() (value int32, err error) {
	retValue, err := instance.GetProperty("Status")
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

// SetStoreInstall sets the value of StoreInstall for the instance
func (instance *MDM_EnterpriseModernAppManagement_AppInstallation01_01) SetPropertyStoreInstall(value string) (err error) {
	return instance.SetProperty("StoreInstall", (value))
}

// GetStoreInstall gets the value of StoreInstall for the instance
func (instance *MDM_EnterpriseModernAppManagement_AppInstallation01_01) GetPropertyStoreInstall() (value string, err error) {
	retValue, err := instance.GetProperty("StoreInstall")
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

//

// <param name="param" type="string "></param>

// <param name="ReturnValue" type="uint32 "></param>
func (instance *MDM_EnterpriseModernAppManagement_AppInstallation01_01) StoreInstallMethod( /* IN */ param string) (result uint32, err error) {
	retVal, err := instance.InvokeMethodWithReturn("StoreInstallMethod", param)
	if err != nil {
		return
	}
	result = uint32(retVal)
	return

}

//

// <param name="param" type="string "></param>

// <param name="ReturnValue" type="uint32 "></param>
func (instance *MDM_EnterpriseModernAppManagement_AppInstallation01_01) HostedInstallMethod( /* IN */ param string) (result uint32, err error) {
	retVal, err := instance.InvokeMethodWithReturn("HostedInstallMethod", param)
	if err != nil {
		return
	}
	result = uint32(retVal)
	return

}
