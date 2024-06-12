// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 6/12/2024 using wmigen
//      Source root.Microsoft.Windows.SMB
//////////////////////////////////////////////
package smb

import (
	"github.com/microsoft/wmi/pkg/base/instance"
	"github.com/microsoft/wmi/pkg/base/query"
	"github.com/microsoft/wmi/pkg/errors"
	cim "github.com/microsoft/wmi/pkg/wmiinstance"
	"reflect"
)

// MSFT_SmbGlobalMapping struct
type MSFT_SmbGlobalMapping struct {
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
	Status SmbGlobalMapping_Status

	//
	UseWriteThrough bool
}

func NewMSFT_SmbGlobalMappingEx1(instance *cim.WmiInstance) (newInstance *MSFT_SmbGlobalMapping, err error) {
	tmp, err := instance, nil

	if err != nil {
		return
	}
	newInstance = &MSFT_SmbGlobalMapping{
		WmiInstance: tmp,
	}
	return
}

func NewMSFT_SmbGlobalMappingEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *MSFT_SmbGlobalMapping, err error) {
	tmp, err := instance.GetWmiInstance(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &MSFT_SmbGlobalMapping{
		WmiInstance: tmp,
	}
	return
}

// SetLocalPath sets the value of LocalPath for the instance
func (instance *MSFT_SmbGlobalMapping) SetPropertyLocalPath(value string) (err error) {
	return instance.SetProperty("LocalPath", (value))
}

// GetLocalPath gets the value of LocalPath for the instance
func (instance *MSFT_SmbGlobalMapping) GetPropertyLocalPath() (value string, err error) {
	retValue, err := instance.GetProperty("LocalPath")
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

// SetRemotePath sets the value of RemotePath for the instance
func (instance *MSFT_SmbGlobalMapping) SetPropertyRemotePath(value string) (err error) {
	return instance.SetProperty("RemotePath", (value))
}

// GetRemotePath gets the value of RemotePath for the instance
func (instance *MSFT_SmbGlobalMapping) GetPropertyRemotePath() (value string, err error) {
	retValue, err := instance.GetProperty("RemotePath")
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

// SetRequireIntegrity sets the value of RequireIntegrity for the instance
func (instance *MSFT_SmbGlobalMapping) SetPropertyRequireIntegrity(value bool) (err error) {
	return instance.SetProperty("RequireIntegrity", (value))
}

// GetRequireIntegrity gets the value of RequireIntegrity for the instance
func (instance *MSFT_SmbGlobalMapping) GetPropertyRequireIntegrity() (value bool, err error) {
	retValue, err := instance.GetProperty("RequireIntegrity")
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

// SetRequirePrivacy sets the value of RequirePrivacy for the instance
func (instance *MSFT_SmbGlobalMapping) SetPropertyRequirePrivacy(value bool) (err error) {
	return instance.SetProperty("RequirePrivacy", (value))
}

// GetRequirePrivacy gets the value of RequirePrivacy for the instance
func (instance *MSFT_SmbGlobalMapping) GetPropertyRequirePrivacy() (value bool, err error) {
	retValue, err := instance.GetProperty("RequirePrivacy")
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

// SetStatus sets the value of Status for the instance
func (instance *MSFT_SmbGlobalMapping) SetPropertyStatus(value SmbGlobalMapping_Status) (err error) {
	return instance.SetProperty("Status", (value))
}

// GetStatus gets the value of Status for the instance
func (instance *MSFT_SmbGlobalMapping) GetPropertyStatus() (value SmbGlobalMapping_Status, err error) {
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

	value = SmbGlobalMapping_Status(valuetmp)

	return
}

// SetUseWriteThrough sets the value of UseWriteThrough for the instance
func (instance *MSFT_SmbGlobalMapping) SetPropertyUseWriteThrough(value bool) (err error) {
	return instance.SetProperty("UseWriteThrough", (value))
}

// GetUseWriteThrough gets the value of UseWriteThrough for the instance
func (instance *MSFT_SmbGlobalMapping) GetPropertyUseWriteThrough() (value bool, err error) {
	retValue, err := instance.GetProperty("UseWriteThrough")
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

//

// <param name="Credential" type="string "></param>
// <param name="DenyAccess" type="string []"></param>
// <param name="FullAccess" type="string []"></param>
// <param name="LocalPath" type="string "></param>
// <param name="Persistent" type="bool "></param>
// <param name="RemotePath" type="string "></param>
// <param name="RequireIntegrity" type="bool "></param>
// <param name="RequirePrivacy" type="bool "></param>
// <param name="UseWriteThrough" type="bool "></param>

// <param name="CreatedMapping" type="MSFT_SmbGlobalMapping "></param>
// <param name="ReturnValue" type="uint32 "></param>
func (instance *MSFT_SmbGlobalMapping) Create( /* IN */ LocalPath string,
	/* IN */ RemotePath string,
	/* IN */ Credential string,
	/* IN */ Persistent bool,
	/* OUT */ CreatedMapping MSFT_SmbGlobalMapping,
	/* OPTIONAL IN */ RequireIntegrity bool,
	/* OPTIONAL IN */ RequirePrivacy bool,
	/* OPTIONAL IN */ FullAccess []string,
	/* OPTIONAL IN */ DenyAccess []string,
	/* OPTIONAL IN */ UseWriteThrough bool) (result uint32, err error) {
	retVal, err := instance.InvokeMethod("Create", LocalPath, RemotePath, Credential, Persistent, RequireIntegrity, RequirePrivacy, FullAccess, DenyAccess, UseWriteThrough)
	if err != nil {
		return
	}
	retValue := retVal[0].(int32)
	result = uint32(retValue)
	return

}

//

// <param name="Force" type="bool "></param>

// <param name="ReturnValue" type="uint32 "></param>
func (instance *MSFT_SmbGlobalMapping) Remove( /* IN */ Force bool) (result uint32, err error) {
	retVal, err := instance.InvokeMethodWithReturn("Remove", Force)
	if err != nil {
		return
	}
	result = uint32(retVal)
	return

}
