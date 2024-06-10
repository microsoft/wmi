// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 9/18/2020 using wmigen
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

// MSFT_SmbOpenFile struct
type MSFT_SmbOpenFile struct {
	*cim.WmiInstance

	//
	ClientComputerName string

	//
	ClientUserName string

	//
	ClusterNodeName string

	//
	ContinuouslyAvailable bool

	//
	Encrypted bool

	//
	FileId uint64

	//
	Locks uint32

	//
	Path string

	//
	Permissions uint32

	//
	ScopeName string

	//
	SessionId uint64

	//
	ShareRelativePath string

	//
	Signed bool

	//
	SmbInstance SmbOpenFile_SmbInstance
}

func NewMSFT_SmbOpenFileEx1(instance *cim.WmiInstance) (newInstance *MSFT_SmbOpenFile, err error) {
	tmp, err := instance, nil

	if err != nil {
		return
	}
	newInstance = &MSFT_SmbOpenFile{
		WmiInstance: tmp,
	}
	return
}

func NewMSFT_SmbOpenFileEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *MSFT_SmbOpenFile, err error) {
	tmp, err := instance.GetWmiInstance(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &MSFT_SmbOpenFile{
		WmiInstance: tmp,
	}
	return
}

// SetClientComputerName sets the value of ClientComputerName for the instance
func (instance *MSFT_SmbOpenFile) SetPropertyClientComputerName(value string) (err error) {
	return instance.SetProperty("ClientComputerName", (value))
}

// GetClientComputerName gets the value of ClientComputerName for the instance
func (instance *MSFT_SmbOpenFile) GetPropertyClientComputerName() (value string, err error) {
	retValue, err := instance.GetProperty("ClientComputerName")
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

// SetClientUserName sets the value of ClientUserName for the instance
func (instance *MSFT_SmbOpenFile) SetPropertyClientUserName(value string) (err error) {
	return instance.SetProperty("ClientUserName", (value))
}

// GetClientUserName gets the value of ClientUserName for the instance
func (instance *MSFT_SmbOpenFile) GetPropertyClientUserName() (value string, err error) {
	retValue, err := instance.GetProperty("ClientUserName")
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

// SetClusterNodeName sets the value of ClusterNodeName for the instance
func (instance *MSFT_SmbOpenFile) SetPropertyClusterNodeName(value string) (err error) {
	return instance.SetProperty("ClusterNodeName", (value))
}

// GetClusterNodeName gets the value of ClusterNodeName for the instance
func (instance *MSFT_SmbOpenFile) GetPropertyClusterNodeName() (value string, err error) {
	retValue, err := instance.GetProperty("ClusterNodeName")
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

// SetContinuouslyAvailable sets the value of ContinuouslyAvailable for the instance
func (instance *MSFT_SmbOpenFile) SetPropertyContinuouslyAvailable(value bool) (err error) {
	return instance.SetProperty("ContinuouslyAvailable", (value))
}

// GetContinuouslyAvailable gets the value of ContinuouslyAvailable for the instance
func (instance *MSFT_SmbOpenFile) GetPropertyContinuouslyAvailable() (value bool, err error) {
	retValue, err := instance.GetProperty("ContinuouslyAvailable")
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

// SetEncrypted sets the value of Encrypted for the instance
func (instance *MSFT_SmbOpenFile) SetPropertyEncrypted(value bool) (err error) {
	return instance.SetProperty("Encrypted", (value))
}

// GetEncrypted gets the value of Encrypted for the instance
func (instance *MSFT_SmbOpenFile) GetPropertyEncrypted() (value bool, err error) {
	retValue, err := instance.GetProperty("Encrypted")
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

// SetFileId sets the value of FileId for the instance
func (instance *MSFT_SmbOpenFile) SetPropertyFileId(value uint64) (err error) {
	return instance.SetProperty("FileId", (value))
}

// GetFileId gets the value of FileId for the instance
func (instance *MSFT_SmbOpenFile) GetPropertyFileId() (value uint64, err error) {
	retValue, err := instance.GetProperty("FileId")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(uint64)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " uint64 is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = uint64(valuetmp)

	return
}

// SetLocks sets the value of Locks for the instance
func (instance *MSFT_SmbOpenFile) SetPropertyLocks(value uint32) (err error) {
	return instance.SetProperty("Locks", (value))
}

// GetLocks gets the value of Locks for the instance
func (instance *MSFT_SmbOpenFile) GetPropertyLocks() (value uint32, err error) {
	retValue, err := instance.GetProperty("Locks")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(uint32)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " uint32 is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = uint32(valuetmp)

	return
}

// SetPath sets the value of Path for the instance
func (instance *MSFT_SmbOpenFile) SetPropertyPath(value string) (err error) {
	return instance.SetProperty("Path", (value))
}

// GetPath gets the value of Path for the instance
func (instance *MSFT_SmbOpenFile) GetPropertyPath() (value string, err error) {
	retValue, err := instance.GetProperty("Path")
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

// SetPermissions sets the value of Permissions for the instance
func (instance *MSFT_SmbOpenFile) SetPropertyPermissions(value uint32) (err error) {
	return instance.SetProperty("Permissions", (value))
}

// GetPermissions gets the value of Permissions for the instance
func (instance *MSFT_SmbOpenFile) GetPropertyPermissions() (value uint32, err error) {
	retValue, err := instance.GetProperty("Permissions")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(uint32)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " uint32 is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = uint32(valuetmp)

	return
}

// SetScopeName sets the value of ScopeName for the instance
func (instance *MSFT_SmbOpenFile) SetPropertyScopeName(value string) (err error) {
	return instance.SetProperty("ScopeName", (value))
}

// GetScopeName gets the value of ScopeName for the instance
func (instance *MSFT_SmbOpenFile) GetPropertyScopeName() (value string, err error) {
	retValue, err := instance.GetProperty("ScopeName")
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

// SetSessionId sets the value of SessionId for the instance
func (instance *MSFT_SmbOpenFile) SetPropertySessionId(value uint64) (err error) {
	return instance.SetProperty("SessionId", (value))
}

// GetSessionId gets the value of SessionId for the instance
func (instance *MSFT_SmbOpenFile) GetPropertySessionId() (value uint64, err error) {
	retValue, err := instance.GetProperty("SessionId")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(uint64)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " uint64 is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = uint64(valuetmp)

	return
}

// SetShareRelativePath sets the value of ShareRelativePath for the instance
func (instance *MSFT_SmbOpenFile) SetPropertyShareRelativePath(value string) (err error) {
	return instance.SetProperty("ShareRelativePath", (value))
}

// GetShareRelativePath gets the value of ShareRelativePath for the instance
func (instance *MSFT_SmbOpenFile) GetPropertyShareRelativePath() (value string, err error) {
	retValue, err := instance.GetProperty("ShareRelativePath")
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

// SetSigned sets the value of Signed for the instance
func (instance *MSFT_SmbOpenFile) SetPropertySigned(value bool) (err error) {
	return instance.SetProperty("Signed", (value))
}

// GetSigned gets the value of Signed for the instance
func (instance *MSFT_SmbOpenFile) GetPropertySigned() (value bool, err error) {
	retValue, err := instance.GetProperty("Signed")
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

// SetSmbInstance sets the value of SmbInstance for the instance
func (instance *MSFT_SmbOpenFile) SetPropertySmbInstance(value SmbOpenFile_SmbInstance) (err error) {
	return instance.SetProperty("SmbInstance", (value))
}

// GetSmbInstance gets the value of SmbInstance for the instance
func (instance *MSFT_SmbOpenFile) GetPropertySmbInstance() (value SmbOpenFile_SmbInstance, err error) {
	retValue, err := instance.GetProperty("SmbInstance")
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

	value = SmbOpenFile_SmbInstance(valuetmp)

	return
}

//

// <param name="ReturnValue" type="uint32 "></param>
func (instance *MSFT_SmbOpenFile) ForceClose() (result uint32, err error) {
	retVal, err := instance.InvokeMethodWithReturn("ForceClose")
	if err != nil {
		return
	}
	result = uint32(retVal)
	return

}
