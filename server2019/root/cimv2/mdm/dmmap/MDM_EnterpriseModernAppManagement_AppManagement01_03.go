// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 3/16/2020 using wmigen
//      Source root.CIMV2.mdm.dmmap
//////////////////////////////////////////////
package dmmap

import (
	"github.com/microsoft/wmi/pkg/wmiinstance"
)

// MDM_EnterpriseModernAppManagement_AppManagement01_03 struct
type MDM_EnterpriseModernAppManagement_AppManagement01_03 struct {
	cim.WmiInstance

	//
	Architecture string

	//
	InstallDate string

	//
	InstallLocation string

	//
	InstanceID string

	//
	IsBundle int32

	//
	IsFramework int32

	//
	IsProvisioned int32

	//
	IsStub int32

	//
	Name string

	//
	PackageStatus int32

	//
	ParentID string

	//
	Publisher string

	//
	RequiresReinstall int32

	//
	ResourceID string

	//
	Users string

	//
	Version string
}

// SetArchitecture sets the value of Architecture for the instance
func (instance *MDM_EnterpriseModernAppManagement_AppManagement01_03) SetPropertyArchitecture(value string) (err error) {
	return instance.SetProperty("Architecture", value)
}

// GetArchitecture gets the value of Architecture for the instance
func (instance *MDM_EnterpriseModernAppManagement_AppManagement01_03) GetPropertyArchitecture() (value string, err error) {
	retValue, err := instance.GetProperty("Architecture")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetInstallDate sets the value of InstallDate for the instance
func (instance *MDM_EnterpriseModernAppManagement_AppManagement01_03) SetPropertyInstallDate(value string) (err error) {
	return instance.SetProperty("InstallDate", value)
}

// GetInstallDate gets the value of InstallDate for the instance
func (instance *MDM_EnterpriseModernAppManagement_AppManagement01_03) GetPropertyInstallDate() (value string, err error) {
	retValue, err := instance.GetProperty("InstallDate")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetInstallLocation sets the value of InstallLocation for the instance
func (instance *MDM_EnterpriseModernAppManagement_AppManagement01_03) SetPropertyInstallLocation(value string) (err error) {
	return instance.SetProperty("InstallLocation", value)
}

// GetInstallLocation gets the value of InstallLocation for the instance
func (instance *MDM_EnterpriseModernAppManagement_AppManagement01_03) GetPropertyInstallLocation() (value string, err error) {
	retValue, err := instance.GetProperty("InstallLocation")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetInstanceID sets the value of InstanceID for the instance
func (instance *MDM_EnterpriseModernAppManagement_AppManagement01_03) SetPropertyInstanceID(value string) (err error) {
	return instance.SetProperty("InstanceID", value)
}

// GetInstanceID gets the value of InstanceID for the instance
func (instance *MDM_EnterpriseModernAppManagement_AppManagement01_03) GetPropertyInstanceID() (value string, err error) {
	retValue, err := instance.GetProperty("InstanceID")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetIsBundle sets the value of IsBundle for the instance
func (instance *MDM_EnterpriseModernAppManagement_AppManagement01_03) SetPropertyIsBundle(value int32) (err error) {
	return instance.SetProperty("IsBundle", value)
}

// GetIsBundle gets the value of IsBundle for the instance
func (instance *MDM_EnterpriseModernAppManagement_AppManagement01_03) GetPropertyIsBundle() (value int32, err error) {
	retValue, err := instance.GetProperty("IsBundle")
	if err != nil {
		return
	}
	value, ok := retValue.(int32)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetIsFramework sets the value of IsFramework for the instance
func (instance *MDM_EnterpriseModernAppManagement_AppManagement01_03) SetPropertyIsFramework(value int32) (err error) {
	return instance.SetProperty("IsFramework", value)
}

// GetIsFramework gets the value of IsFramework for the instance
func (instance *MDM_EnterpriseModernAppManagement_AppManagement01_03) GetPropertyIsFramework() (value int32, err error) {
	retValue, err := instance.GetProperty("IsFramework")
	if err != nil {
		return
	}
	value, ok := retValue.(int32)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetIsProvisioned sets the value of IsProvisioned for the instance
func (instance *MDM_EnterpriseModernAppManagement_AppManagement01_03) SetPropertyIsProvisioned(value int32) (err error) {
	return instance.SetProperty("IsProvisioned", value)
}

// GetIsProvisioned gets the value of IsProvisioned for the instance
func (instance *MDM_EnterpriseModernAppManagement_AppManagement01_03) GetPropertyIsProvisioned() (value int32, err error) {
	retValue, err := instance.GetProperty("IsProvisioned")
	if err != nil {
		return
	}
	value, ok := retValue.(int32)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetIsStub sets the value of IsStub for the instance
func (instance *MDM_EnterpriseModernAppManagement_AppManagement01_03) SetPropertyIsStub(value int32) (err error) {
	return instance.SetProperty("IsStub", value)
}

// GetIsStub gets the value of IsStub for the instance
func (instance *MDM_EnterpriseModernAppManagement_AppManagement01_03) GetPropertyIsStub() (value int32, err error) {
	retValue, err := instance.GetProperty("IsStub")
	if err != nil {
		return
	}
	value, ok := retValue.(int32)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetName sets the value of Name for the instance
func (instance *MDM_EnterpriseModernAppManagement_AppManagement01_03) SetPropertyName(value string) (err error) {
	return instance.SetProperty("Name", value)
}

// GetName gets the value of Name for the instance
func (instance *MDM_EnterpriseModernAppManagement_AppManagement01_03) GetPropertyName() (value string, err error) {
	retValue, err := instance.GetProperty("Name")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetPackageStatus sets the value of PackageStatus for the instance
func (instance *MDM_EnterpriseModernAppManagement_AppManagement01_03) SetPropertyPackageStatus(value int32) (err error) {
	return instance.SetProperty("PackageStatus", value)
}

// GetPackageStatus gets the value of PackageStatus for the instance
func (instance *MDM_EnterpriseModernAppManagement_AppManagement01_03) GetPropertyPackageStatus() (value int32, err error) {
	retValue, err := instance.GetProperty("PackageStatus")
	if err != nil {
		return
	}
	value, ok := retValue.(int32)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetParentID sets the value of ParentID for the instance
func (instance *MDM_EnterpriseModernAppManagement_AppManagement01_03) SetPropertyParentID(value string) (err error) {
	return instance.SetProperty("ParentID", value)
}

// GetParentID gets the value of ParentID for the instance
func (instance *MDM_EnterpriseModernAppManagement_AppManagement01_03) GetPropertyParentID() (value string, err error) {
	retValue, err := instance.GetProperty("ParentID")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetPublisher sets the value of Publisher for the instance
func (instance *MDM_EnterpriseModernAppManagement_AppManagement01_03) SetPropertyPublisher(value string) (err error) {
	return instance.SetProperty("Publisher", value)
}

// GetPublisher gets the value of Publisher for the instance
func (instance *MDM_EnterpriseModernAppManagement_AppManagement01_03) GetPropertyPublisher() (value string, err error) {
	retValue, err := instance.GetProperty("Publisher")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetRequiresReinstall sets the value of RequiresReinstall for the instance
func (instance *MDM_EnterpriseModernAppManagement_AppManagement01_03) SetPropertyRequiresReinstall(value int32) (err error) {
	return instance.SetProperty("RequiresReinstall", value)
}

// GetRequiresReinstall gets the value of RequiresReinstall for the instance
func (instance *MDM_EnterpriseModernAppManagement_AppManagement01_03) GetPropertyRequiresReinstall() (value int32, err error) {
	retValue, err := instance.GetProperty("RequiresReinstall")
	if err != nil {
		return
	}
	value, ok := retValue.(int32)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetResourceID sets the value of ResourceID for the instance
func (instance *MDM_EnterpriseModernAppManagement_AppManagement01_03) SetPropertyResourceID(value string) (err error) {
	return instance.SetProperty("ResourceID", value)
}

// GetResourceID gets the value of ResourceID for the instance
func (instance *MDM_EnterpriseModernAppManagement_AppManagement01_03) GetPropertyResourceID() (value string, err error) {
	retValue, err := instance.GetProperty("ResourceID")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetUsers sets the value of Users for the instance
func (instance *MDM_EnterpriseModernAppManagement_AppManagement01_03) SetPropertyUsers(value string) (err error) {
	return instance.SetProperty("Users", value)
}

// GetUsers gets the value of Users for the instance
func (instance *MDM_EnterpriseModernAppManagement_AppManagement01_03) GetPropertyUsers() (value string, err error) {
	retValue, err := instance.GetProperty("Users")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetVersion sets the value of Version for the instance
func (instance *MDM_EnterpriseModernAppManagement_AppManagement01_03) SetPropertyVersion(value string) (err error) {
	return instance.SetProperty("Version", value)
}

// GetVersion gets the value of Version for the instance
func (instance *MDM_EnterpriseModernAppManagement_AppManagement01_03) GetPropertyVersion() (value string, err error) {
	retValue, err := instance.GetProperty("Version")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}
