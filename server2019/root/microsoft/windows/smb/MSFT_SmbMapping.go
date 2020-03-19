// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 3/19/2020 using wmigen
//      Source root.Microsoft.Windows.SMB
//////////////////////////////////////////////
package smb

import (
	"github.com/microsoft/wmi/pkg/base/instance"
	"github.com/microsoft/wmi/pkg/base/query"
	cim "github.com/microsoft/wmi/pkg/wmiinstance"
)

// MSFT_SmbMapping struct
type MSFT_SmbMapping struct {
	*cim.WmiInstance

	//
	LocalPath string

	//
	RemotePath string

	//
	RequireIntegrity bool

	//
	RequirePrivacy bool

	//
	Status SmbMapping_Status

	//
	UseWriteThrough bool
}

func NewMSFT_SmbMappingEx1(instance *cim.WmiInstance) (newInstance *MSFT_SmbMapping, err error) {
	tmp, err := instance, nil

	if err != nil {
		return
	}
	newInstance = &MSFT_SmbMapping{
		WmiInstance: tmp,
	}
	return
}

func NewMSFT_SmbMappingEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *MSFT_SmbMapping, err error) {
	tmp, err := instance.GetWmiInstance(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &MSFT_SmbMapping{
		WmiInstance: tmp,
	}
	return
}

// SetLocalPath sets the value of LocalPath for the instance
func (instance *MSFT_SmbMapping) SetPropertyLocalPath(value string) (err error) {
	return instance.SetProperty("LocalPath", value)
}

// GetLocalPath gets the value of LocalPath for the instance
func (instance *MSFT_SmbMapping) GetPropertyLocalPath() (value string, err error) {
	retValue, err := instance.GetProperty("LocalPath")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetRemotePath sets the value of RemotePath for the instance
func (instance *MSFT_SmbMapping) SetPropertyRemotePath(value string) (err error) {
	return instance.SetProperty("RemotePath", value)
}

// GetRemotePath gets the value of RemotePath for the instance
func (instance *MSFT_SmbMapping) GetPropertyRemotePath() (value string, err error) {
	retValue, err := instance.GetProperty("RemotePath")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetRequireIntegrity sets the value of RequireIntegrity for the instance
func (instance *MSFT_SmbMapping) SetPropertyRequireIntegrity(value bool) (err error) {
	return instance.SetProperty("RequireIntegrity", value)
}

// GetRequireIntegrity gets the value of RequireIntegrity for the instance
func (instance *MSFT_SmbMapping) GetPropertyRequireIntegrity() (value bool, err error) {
	retValue, err := instance.GetProperty("RequireIntegrity")
	if err != nil {
		return
	}
	value, ok := retValue.(bool)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetRequirePrivacy sets the value of RequirePrivacy for the instance
func (instance *MSFT_SmbMapping) SetPropertyRequirePrivacy(value bool) (err error) {
	return instance.SetProperty("RequirePrivacy", value)
}

// GetRequirePrivacy gets the value of RequirePrivacy for the instance
func (instance *MSFT_SmbMapping) GetPropertyRequirePrivacy() (value bool, err error) {
	retValue, err := instance.GetProperty("RequirePrivacy")
	if err != nil {
		return
	}
	value, ok := retValue.(bool)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetStatus sets the value of Status for the instance
func (instance *MSFT_SmbMapping) SetPropertyStatus(value SmbMapping_Status) (err error) {
	return instance.SetProperty("Status", value)
}

// GetStatus gets the value of Status for the instance
func (instance *MSFT_SmbMapping) GetPropertyStatus() (value SmbMapping_Status, err error) {
	retValue, err := instance.GetProperty("Status")
	if err != nil {
		return
	}
	value, ok := retValue.(SmbMapping_Status)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetUseWriteThrough sets the value of UseWriteThrough for the instance
func (instance *MSFT_SmbMapping) SetPropertyUseWriteThrough(value bool) (err error) {
	return instance.SetProperty("UseWriteThrough", value)
}

// GetUseWriteThrough gets the value of UseWriteThrough for the instance
func (instance *MSFT_SmbMapping) GetPropertyUseWriteThrough() (value bool, err error) {
	retValue, err := instance.GetProperty("UseWriteThrough")
	if err != nil {
		return
	}
	value, ok := retValue.(bool)
	if !ok {
		// TODO: Set an error
	}
	return
}

//

// <param name="Force" type="bool "></param>
// <param name="UpdateProfile" type="bool "></param>

// <param name="ReturnValue" type="uint32 "></param>
func (instance *MSFT_SmbMapping) Remove( /* IN */ UpdateProfile bool,
	/* IN */ Force bool) (result uint32, err error) {
	retVal, err := instance.InvokeMethodWithReturn("Remove", UpdateProfile, Force)
	if err != nil {
		return
	}
	result = uint32(retVal)
	return

}

//

// <param name="HomeFolder" type="bool "></param>
// <param name="LocalPath" type="string "></param>
// <param name="Password" type="string "></param>
// <param name="Persistent" type="bool "></param>
// <param name="RemotePath" type="string "></param>
// <param name="RequireIntegrity" type="bool "></param>
// <param name="RequirePrivacy" type="bool "></param>
// <param name="SaveCredentials" type="bool "></param>
// <param name="UserName" type="string "></param>
// <param name="UseWriteThrough" type="bool "></param>

// <param name="CreatedMapping" type="MSFT_SmbMapping "></param>
// <param name="ReturnValue" type="uint32 "></param>
func (instance *MSFT_SmbMapping) Create( /* IN */ LocalPath string,
	/* IN */ RemotePath string,
	/* IN */ UserName string,
	/* IN */ Password string,
	/* IN */ Persistent bool,
	/* IN */ SaveCredentials bool,
	/* IN */ HomeFolder bool,
	/* OUT */ CreatedMapping MSFT_SmbMapping,
	/* OPTIONAL IN */ RequireIntegrity bool,
	/* OPTIONAL IN */ RequirePrivacy bool,
	/* OPTIONAL IN */ UseWriteThrough bool) (result uint32, err error) {
	retVal, err := instance.InvokeMethod("Create", LocalPath, RemotePath, UserName, Password, Persistent, SaveCredentials, HomeFolder, RequireIntegrity, RequirePrivacy, UseWriteThrough)
	if err != nil {
		return
	}
	retValue := retVal[0].(int32)
	result = uint32(retValue)
	return

}
