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

// MDM_WindowsDefenderApplicationGuard_Settings01 struct
type MDM_WindowsDefenderApplicationGuard_Settings01 struct {
	*cim.WmiInstance

	//
	AllowCameraMicrophoneRedirection int32

	//
	AllowPersistence int32

	//
	AllowVirtualGPU int32

	//
	AllowWindowsDefenderApplicationGuard int32

	//
	BlockNonEnterpriseContent int32

	//
	CertificateThumbprints string

	//
	ClipboardFileType int32

	//
	ClipboardSettings int32

	//
	FileTrustCriteria int32

	//
	FileTrustOriginMarkOfTheWeb int32

	//
	FileTrustOriginNetworkShare int32

	//
	FileTrustOriginRemovableMedia int32

	//
	InstanceID string

	//
	ParentID string

	//
	PrintingSettings int32

	//
	SaveFilesToHost int32
}

func NewMDM_WindowsDefenderApplicationGuard_Settings01Ex1(instance *cim.WmiInstance) (newInstance *MDM_WindowsDefenderApplicationGuard_Settings01, err error) {
	tmp, err := instance, nil

	if err != nil {
		return
	}
	newInstance = &MDM_WindowsDefenderApplicationGuard_Settings01{
		WmiInstance: tmp,
	}
	return
}

func NewMDM_WindowsDefenderApplicationGuard_Settings01Ex6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *MDM_WindowsDefenderApplicationGuard_Settings01, err error) {
	tmp, err := instance.GetWmiInstance(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &MDM_WindowsDefenderApplicationGuard_Settings01{
		WmiInstance: tmp,
	}
	return
}

// SetAllowCameraMicrophoneRedirection sets the value of AllowCameraMicrophoneRedirection for the instance
func (instance *MDM_WindowsDefenderApplicationGuard_Settings01) SetPropertyAllowCameraMicrophoneRedirection(value int32) (err error) {
	return instance.SetProperty("AllowCameraMicrophoneRedirection", (value))
}

// GetAllowCameraMicrophoneRedirection gets the value of AllowCameraMicrophoneRedirection for the instance
func (instance *MDM_WindowsDefenderApplicationGuard_Settings01) GetPropertyAllowCameraMicrophoneRedirection() (value int32, err error) {
	retValue, err := instance.GetProperty("AllowCameraMicrophoneRedirection")
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

// SetAllowPersistence sets the value of AllowPersistence for the instance
func (instance *MDM_WindowsDefenderApplicationGuard_Settings01) SetPropertyAllowPersistence(value int32) (err error) {
	return instance.SetProperty("AllowPersistence", (value))
}

// GetAllowPersistence gets the value of AllowPersistence for the instance
func (instance *MDM_WindowsDefenderApplicationGuard_Settings01) GetPropertyAllowPersistence() (value int32, err error) {
	retValue, err := instance.GetProperty("AllowPersistence")
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

// SetAllowVirtualGPU sets the value of AllowVirtualGPU for the instance
func (instance *MDM_WindowsDefenderApplicationGuard_Settings01) SetPropertyAllowVirtualGPU(value int32) (err error) {
	return instance.SetProperty("AllowVirtualGPU", (value))
}

// GetAllowVirtualGPU gets the value of AllowVirtualGPU for the instance
func (instance *MDM_WindowsDefenderApplicationGuard_Settings01) GetPropertyAllowVirtualGPU() (value int32, err error) {
	retValue, err := instance.GetProperty("AllowVirtualGPU")
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

// SetAllowWindowsDefenderApplicationGuard sets the value of AllowWindowsDefenderApplicationGuard for the instance
func (instance *MDM_WindowsDefenderApplicationGuard_Settings01) SetPropertyAllowWindowsDefenderApplicationGuard(value int32) (err error) {
	return instance.SetProperty("AllowWindowsDefenderApplicationGuard", (value))
}

// GetAllowWindowsDefenderApplicationGuard gets the value of AllowWindowsDefenderApplicationGuard for the instance
func (instance *MDM_WindowsDefenderApplicationGuard_Settings01) GetPropertyAllowWindowsDefenderApplicationGuard() (value int32, err error) {
	retValue, err := instance.GetProperty("AllowWindowsDefenderApplicationGuard")
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

// SetBlockNonEnterpriseContent sets the value of BlockNonEnterpriseContent for the instance
func (instance *MDM_WindowsDefenderApplicationGuard_Settings01) SetPropertyBlockNonEnterpriseContent(value int32) (err error) {
	return instance.SetProperty("BlockNonEnterpriseContent", (value))
}

// GetBlockNonEnterpriseContent gets the value of BlockNonEnterpriseContent for the instance
func (instance *MDM_WindowsDefenderApplicationGuard_Settings01) GetPropertyBlockNonEnterpriseContent() (value int32, err error) {
	retValue, err := instance.GetProperty("BlockNonEnterpriseContent")
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

// SetCertificateThumbprints sets the value of CertificateThumbprints for the instance
func (instance *MDM_WindowsDefenderApplicationGuard_Settings01) SetPropertyCertificateThumbprints(value string) (err error) {
	return instance.SetProperty("CertificateThumbprints", (value))
}

// GetCertificateThumbprints gets the value of CertificateThumbprints for the instance
func (instance *MDM_WindowsDefenderApplicationGuard_Settings01) GetPropertyCertificateThumbprints() (value string, err error) {
	retValue, err := instance.GetProperty("CertificateThumbprints")
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

// SetClipboardFileType sets the value of ClipboardFileType for the instance
func (instance *MDM_WindowsDefenderApplicationGuard_Settings01) SetPropertyClipboardFileType(value int32) (err error) {
	return instance.SetProperty("ClipboardFileType", (value))
}

// GetClipboardFileType gets the value of ClipboardFileType for the instance
func (instance *MDM_WindowsDefenderApplicationGuard_Settings01) GetPropertyClipboardFileType() (value int32, err error) {
	retValue, err := instance.GetProperty("ClipboardFileType")
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

// SetClipboardSettings sets the value of ClipboardSettings for the instance
func (instance *MDM_WindowsDefenderApplicationGuard_Settings01) SetPropertyClipboardSettings(value int32) (err error) {
	return instance.SetProperty("ClipboardSettings", (value))
}

// GetClipboardSettings gets the value of ClipboardSettings for the instance
func (instance *MDM_WindowsDefenderApplicationGuard_Settings01) GetPropertyClipboardSettings() (value int32, err error) {
	retValue, err := instance.GetProperty("ClipboardSettings")
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

// SetFileTrustCriteria sets the value of FileTrustCriteria for the instance
func (instance *MDM_WindowsDefenderApplicationGuard_Settings01) SetPropertyFileTrustCriteria(value int32) (err error) {
	return instance.SetProperty("FileTrustCriteria", (value))
}

// GetFileTrustCriteria gets the value of FileTrustCriteria for the instance
func (instance *MDM_WindowsDefenderApplicationGuard_Settings01) GetPropertyFileTrustCriteria() (value int32, err error) {
	retValue, err := instance.GetProperty("FileTrustCriteria")
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

// SetFileTrustOriginMarkOfTheWeb sets the value of FileTrustOriginMarkOfTheWeb for the instance
func (instance *MDM_WindowsDefenderApplicationGuard_Settings01) SetPropertyFileTrustOriginMarkOfTheWeb(value int32) (err error) {
	return instance.SetProperty("FileTrustOriginMarkOfTheWeb", (value))
}

// GetFileTrustOriginMarkOfTheWeb gets the value of FileTrustOriginMarkOfTheWeb for the instance
func (instance *MDM_WindowsDefenderApplicationGuard_Settings01) GetPropertyFileTrustOriginMarkOfTheWeb() (value int32, err error) {
	retValue, err := instance.GetProperty("FileTrustOriginMarkOfTheWeb")
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

// SetFileTrustOriginNetworkShare sets the value of FileTrustOriginNetworkShare for the instance
func (instance *MDM_WindowsDefenderApplicationGuard_Settings01) SetPropertyFileTrustOriginNetworkShare(value int32) (err error) {
	return instance.SetProperty("FileTrustOriginNetworkShare", (value))
}

// GetFileTrustOriginNetworkShare gets the value of FileTrustOriginNetworkShare for the instance
func (instance *MDM_WindowsDefenderApplicationGuard_Settings01) GetPropertyFileTrustOriginNetworkShare() (value int32, err error) {
	retValue, err := instance.GetProperty("FileTrustOriginNetworkShare")
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

// SetFileTrustOriginRemovableMedia sets the value of FileTrustOriginRemovableMedia for the instance
func (instance *MDM_WindowsDefenderApplicationGuard_Settings01) SetPropertyFileTrustOriginRemovableMedia(value int32) (err error) {
	return instance.SetProperty("FileTrustOriginRemovableMedia", (value))
}

// GetFileTrustOriginRemovableMedia gets the value of FileTrustOriginRemovableMedia for the instance
func (instance *MDM_WindowsDefenderApplicationGuard_Settings01) GetPropertyFileTrustOriginRemovableMedia() (value int32, err error) {
	retValue, err := instance.GetProperty("FileTrustOriginRemovableMedia")
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
func (instance *MDM_WindowsDefenderApplicationGuard_Settings01) SetPropertyInstanceID(value string) (err error) {
	return instance.SetProperty("InstanceID", (value))
}

// GetInstanceID gets the value of InstanceID for the instance
func (instance *MDM_WindowsDefenderApplicationGuard_Settings01) GetPropertyInstanceID() (value string, err error) {
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
func (instance *MDM_WindowsDefenderApplicationGuard_Settings01) SetPropertyParentID(value string) (err error) {
	return instance.SetProperty("ParentID", (value))
}

// GetParentID gets the value of ParentID for the instance
func (instance *MDM_WindowsDefenderApplicationGuard_Settings01) GetPropertyParentID() (value string, err error) {
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

// SetPrintingSettings sets the value of PrintingSettings for the instance
func (instance *MDM_WindowsDefenderApplicationGuard_Settings01) SetPropertyPrintingSettings(value int32) (err error) {
	return instance.SetProperty("PrintingSettings", (value))
}

// GetPrintingSettings gets the value of PrintingSettings for the instance
func (instance *MDM_WindowsDefenderApplicationGuard_Settings01) GetPropertyPrintingSettings() (value int32, err error) {
	retValue, err := instance.GetProperty("PrintingSettings")
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

// SetSaveFilesToHost sets the value of SaveFilesToHost for the instance
func (instance *MDM_WindowsDefenderApplicationGuard_Settings01) SetPropertySaveFilesToHost(value int32) (err error) {
	return instance.SetProperty("SaveFilesToHost", (value))
}

// GetSaveFilesToHost gets the value of SaveFilesToHost for the instance
func (instance *MDM_WindowsDefenderApplicationGuard_Settings01) GetPropertySaveFilesToHost() (value int32, err error) {
	retValue, err := instance.GetProperty("SaveFilesToHost")
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
