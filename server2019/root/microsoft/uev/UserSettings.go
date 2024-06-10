// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 9/18/2020 using wmigen
//      Source root.Microsoft.Uev
//////////////////////////////////////////////
package uev

import (
	"github.com/microsoft/wmi/pkg/base/instance"
	"github.com/microsoft/wmi/pkg/base/query"
	cim "github.com/microsoft/wmi/pkg/wmiinstance"
)

// UserSettings struct
type UserSettings struct {
	*cim.WmiInstance
}

func NewUserSettingsEx1(instance *cim.WmiInstance) (newInstance *UserSettings, err error) {
	tmp, err := instance, nil

	if err != nil {
		return
	}
	newInstance = &UserSettings{
		WmiInstance: tmp,
	}
	return
}

func NewUserSettingsEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *UserSettings, err error) {
	tmp, err := instance.GetWmiInstance(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &UserSettings{
		WmiInstance: tmp,
	}
	return
}

// Restore users' application and Windows settings back to the original settings state.

// <param name="RestoreType" type="uint32 ">Restore Type</param>
// <param name="TemplateId" type="string ">Template ID</param>
func (instance *UserSettings) RestoreByTemplateId( /* IN */ TemplateId string,
	/* IN */ RestoreType uint32) (err error) {
	_, err = instance.InvokeMethodWithReturn("RestoreByTemplateId", TemplateId, RestoreType)
	if err != nil {
		return
	}
	return

}

// Export a settings package file to a XML file listing all the settings in the package.

// <param name="AbsolutePathToPackage" type="string ">Absolute path to the settings package file</param>

// <param name="ReturnValue" type="string "></param>
func (instance *UserSettings) ExportPackage( /* IN */ AbsolutePathToPackage string) (result string, err error) {
	retVal, err := instance.InvokeMethodWithReturn("ExportPackage", AbsolutePathToPackage)
	if err != nil {
		return
	}
	result = string(retVal)
	return

}

// Get the version of a settings package.

// <param name="AbsolutePathToPackage" type="string ">Absolute path to the settings package file</param>

// <param name="ReturnValue" type="uint32 "></param>
func (instance *UserSettings) GetPackageVersion( /* IN */ AbsolutePathToPackage string) (result uint32, err error) {
	retVal, err := instance.InvokeMethodWithReturn("GetPackageVersion", AbsolutePathToPackage)
	if err != nil {
		return
	}
	result = uint32(retVal)
	return

}

// Check if the user live ID is connected.

// <param name="ReturnValue" type="bool "></param>
func (instance *UserSettings) IsUserMsaConnected() (result bool, err error) {
	retVal, err := instance.InvokeMethodWithReturn("IsUserMsaConnected")
	if err != nil {
		return
	}
	result = (retVal > 0)
	return

}

// Restores the backup packages from another machine. Returns false if a restore is already pending, true otherwise

// <param name="ComputerName" type="string ">Fully qualified computer name from which to restore packages</param>

// <param name="ReturnValue" type="bool "></param>
func (instance *UserSettings) RestoreBackup( /* IN */ ComputerName string) (result bool, err error) {
	retVal, err := instance.InvokeMethodWithReturn("RestoreBackup", ComputerName)
	if err != nil {
		return
	}
	result = (retVal > 0)
	return

}
