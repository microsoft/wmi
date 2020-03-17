// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 3/16/2020 using wmigen
//      Source root.Microsoft.Windows.SMB
//////////////////////////////////////////////
package smb

import (
	"github.com/microsoft/wmi/pkg/wmiinstance"
)

// MSFT_SmbShare struct
type MSFT_SmbShare struct {
	cim.WmiInstance

	//
	AvailabilityType SmbShare_AvailabilityType

	//
	CachingMode SmbShare_CachingMode

	//
	CATimeout uint32

	//
	ConcurrentUserLimit uint32

	//
	ContinuouslyAvailable bool

	//
	CurrentUsers uint32

	//
	Description string

	//
	EncryptData bool

	//
	FolderEnumerationMode SmbShare_FolderEnumerationMode

	//
	IdentityRemoting bool

	//
	Infrastructure bool

	//
	LeasingMode SmbShare_LeasingMode

	//
	Name string

	//
	Path string

	//
	Scoped bool

	//
	ScopeName string

	//
	SecurityDescriptor string

	//
	ShadowCopy bool

	//
	ShareState SmbShare_ShareState

	//
	ShareType SmbShare_ShareType

	//
	SmbInstance SmbShare_SmbInstance

	//
	Special bool

	//
	Temporary bool

	//
	Volume string
}

// SetAvailabilityType sets the value of AvailabilityType for the instance
func (instance *MSFT_SmbShare) SetPropertyAvailabilityType(value SmbShare_AvailabilityType) (err error) {
	return instance.SetProperty("AvailabilityType", value)
}

// GetAvailabilityType gets the value of AvailabilityType for the instance
func (instance *MSFT_SmbShare) GetPropertyAvailabilityType() (value SmbShare_AvailabilityType, err error) {
	retValue, err := instance.GetProperty("AvailabilityType")
	if err != nil {
		return
	}
	value, ok := retValue.(SmbShare_AvailabilityType)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetCachingMode sets the value of CachingMode for the instance
func (instance *MSFT_SmbShare) SetPropertyCachingMode(value SmbShare_CachingMode) (err error) {
	return instance.SetProperty("CachingMode", value)
}

// GetCachingMode gets the value of CachingMode for the instance
func (instance *MSFT_SmbShare) GetPropertyCachingMode() (value SmbShare_CachingMode, err error) {
	retValue, err := instance.GetProperty("CachingMode")
	if err != nil {
		return
	}
	value, ok := retValue.(SmbShare_CachingMode)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetCATimeout sets the value of CATimeout for the instance
func (instance *MSFT_SmbShare) SetPropertyCATimeout(value uint32) (err error) {
	return instance.SetProperty("CATimeout", value)
}

// GetCATimeout gets the value of CATimeout for the instance
func (instance *MSFT_SmbShare) GetPropertyCATimeout() (value uint32, err error) {
	retValue, err := instance.GetProperty("CATimeout")
	if err != nil {
		return
	}
	value, ok := retValue.(uint32)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetConcurrentUserLimit sets the value of ConcurrentUserLimit for the instance
func (instance *MSFT_SmbShare) SetPropertyConcurrentUserLimit(value uint32) (err error) {
	return instance.SetProperty("ConcurrentUserLimit", value)
}

// GetConcurrentUserLimit gets the value of ConcurrentUserLimit for the instance
func (instance *MSFT_SmbShare) GetPropertyConcurrentUserLimit() (value uint32, err error) {
	retValue, err := instance.GetProperty("ConcurrentUserLimit")
	if err != nil {
		return
	}
	value, ok := retValue.(uint32)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetContinuouslyAvailable sets the value of ContinuouslyAvailable for the instance
func (instance *MSFT_SmbShare) SetPropertyContinuouslyAvailable(value bool) (err error) {
	return instance.SetProperty("ContinuouslyAvailable", value)
}

// GetContinuouslyAvailable gets the value of ContinuouslyAvailable for the instance
func (instance *MSFT_SmbShare) GetPropertyContinuouslyAvailable() (value bool, err error) {
	retValue, err := instance.GetProperty("ContinuouslyAvailable")
	if err != nil {
		return
	}
	value, ok := retValue.(bool)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetCurrentUsers sets the value of CurrentUsers for the instance
func (instance *MSFT_SmbShare) SetPropertyCurrentUsers(value uint32) (err error) {
	return instance.SetProperty("CurrentUsers", value)
}

// GetCurrentUsers gets the value of CurrentUsers for the instance
func (instance *MSFT_SmbShare) GetPropertyCurrentUsers() (value uint32, err error) {
	retValue, err := instance.GetProperty("CurrentUsers")
	if err != nil {
		return
	}
	value, ok := retValue.(uint32)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetDescription sets the value of Description for the instance
func (instance *MSFT_SmbShare) SetPropertyDescription(value string) (err error) {
	return instance.SetProperty("Description", value)
}

// GetDescription gets the value of Description for the instance
func (instance *MSFT_SmbShare) GetPropertyDescription() (value string, err error) {
	retValue, err := instance.GetProperty("Description")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetEncryptData sets the value of EncryptData for the instance
func (instance *MSFT_SmbShare) SetPropertyEncryptData(value bool) (err error) {
	return instance.SetProperty("EncryptData", value)
}

// GetEncryptData gets the value of EncryptData for the instance
func (instance *MSFT_SmbShare) GetPropertyEncryptData() (value bool, err error) {
	retValue, err := instance.GetProperty("EncryptData")
	if err != nil {
		return
	}
	value, ok := retValue.(bool)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetFolderEnumerationMode sets the value of FolderEnumerationMode for the instance
func (instance *MSFT_SmbShare) SetPropertyFolderEnumerationMode(value SmbShare_FolderEnumerationMode) (err error) {
	return instance.SetProperty("FolderEnumerationMode", value)
}

// GetFolderEnumerationMode gets the value of FolderEnumerationMode for the instance
func (instance *MSFT_SmbShare) GetPropertyFolderEnumerationMode() (value SmbShare_FolderEnumerationMode, err error) {
	retValue, err := instance.GetProperty("FolderEnumerationMode")
	if err != nil {
		return
	}
	value, ok := retValue.(SmbShare_FolderEnumerationMode)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetIdentityRemoting sets the value of IdentityRemoting for the instance
func (instance *MSFT_SmbShare) SetPropertyIdentityRemoting(value bool) (err error) {
	return instance.SetProperty("IdentityRemoting", value)
}

// GetIdentityRemoting gets the value of IdentityRemoting for the instance
func (instance *MSFT_SmbShare) GetPropertyIdentityRemoting() (value bool, err error) {
	retValue, err := instance.GetProperty("IdentityRemoting")
	if err != nil {
		return
	}
	value, ok := retValue.(bool)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetInfrastructure sets the value of Infrastructure for the instance
func (instance *MSFT_SmbShare) SetPropertyInfrastructure(value bool) (err error) {
	return instance.SetProperty("Infrastructure", value)
}

// GetInfrastructure gets the value of Infrastructure for the instance
func (instance *MSFT_SmbShare) GetPropertyInfrastructure() (value bool, err error) {
	retValue, err := instance.GetProperty("Infrastructure")
	if err != nil {
		return
	}
	value, ok := retValue.(bool)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetLeasingMode sets the value of LeasingMode for the instance
func (instance *MSFT_SmbShare) SetPropertyLeasingMode(value SmbShare_LeasingMode) (err error) {
	return instance.SetProperty("LeasingMode", value)
}

// GetLeasingMode gets the value of LeasingMode for the instance
func (instance *MSFT_SmbShare) GetPropertyLeasingMode() (value SmbShare_LeasingMode, err error) {
	retValue, err := instance.GetProperty("LeasingMode")
	if err != nil {
		return
	}
	value, ok := retValue.(SmbShare_LeasingMode)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetName sets the value of Name for the instance
func (instance *MSFT_SmbShare) SetPropertyName(value string) (err error) {
	return instance.SetProperty("Name", value)
}

// GetName gets the value of Name for the instance
func (instance *MSFT_SmbShare) GetPropertyName() (value string, err error) {
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

// SetPath sets the value of Path for the instance
func (instance *MSFT_SmbShare) SetPropertyPath(value string) (err error) {
	return instance.SetProperty("Path", value)
}

// GetPath gets the value of Path for the instance
func (instance *MSFT_SmbShare) GetPropertyPath() (value string, err error) {
	retValue, err := instance.GetProperty("Path")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetScoped sets the value of Scoped for the instance
func (instance *MSFT_SmbShare) SetPropertyScoped(value bool) (err error) {
	return instance.SetProperty("Scoped", value)
}

// GetScoped gets the value of Scoped for the instance
func (instance *MSFT_SmbShare) GetPropertyScoped() (value bool, err error) {
	retValue, err := instance.GetProperty("Scoped")
	if err != nil {
		return
	}
	value, ok := retValue.(bool)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetScopeName sets the value of ScopeName for the instance
func (instance *MSFT_SmbShare) SetPropertyScopeName(value string) (err error) {
	return instance.SetProperty("ScopeName", value)
}

// GetScopeName gets the value of ScopeName for the instance
func (instance *MSFT_SmbShare) GetPropertyScopeName() (value string, err error) {
	retValue, err := instance.GetProperty("ScopeName")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetSecurityDescriptor sets the value of SecurityDescriptor for the instance
func (instance *MSFT_SmbShare) SetPropertySecurityDescriptor(value string) (err error) {
	return instance.SetProperty("SecurityDescriptor", value)
}

// GetSecurityDescriptor gets the value of SecurityDescriptor for the instance
func (instance *MSFT_SmbShare) GetPropertySecurityDescriptor() (value string, err error) {
	retValue, err := instance.GetProperty("SecurityDescriptor")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetShadowCopy sets the value of ShadowCopy for the instance
func (instance *MSFT_SmbShare) SetPropertyShadowCopy(value bool) (err error) {
	return instance.SetProperty("ShadowCopy", value)
}

// GetShadowCopy gets the value of ShadowCopy for the instance
func (instance *MSFT_SmbShare) GetPropertyShadowCopy() (value bool, err error) {
	retValue, err := instance.GetProperty("ShadowCopy")
	if err != nil {
		return
	}
	value, ok := retValue.(bool)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetShareState sets the value of ShareState for the instance
func (instance *MSFT_SmbShare) SetPropertyShareState(value SmbShare_ShareState) (err error) {
	return instance.SetProperty("ShareState", value)
}

// GetShareState gets the value of ShareState for the instance
func (instance *MSFT_SmbShare) GetPropertyShareState() (value SmbShare_ShareState, err error) {
	retValue, err := instance.GetProperty("ShareState")
	if err != nil {
		return
	}
	value, ok := retValue.(SmbShare_ShareState)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetShareType sets the value of ShareType for the instance
func (instance *MSFT_SmbShare) SetPropertyShareType(value SmbShare_ShareType) (err error) {
	return instance.SetProperty("ShareType", value)
}

// GetShareType gets the value of ShareType for the instance
func (instance *MSFT_SmbShare) GetPropertyShareType() (value SmbShare_ShareType, err error) {
	retValue, err := instance.GetProperty("ShareType")
	if err != nil {
		return
	}
	value, ok := retValue.(SmbShare_ShareType)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetSmbInstance sets the value of SmbInstance for the instance
func (instance *MSFT_SmbShare) SetPropertySmbInstance(value SmbShare_SmbInstance) (err error) {
	return instance.SetProperty("SmbInstance", value)
}

// GetSmbInstance gets the value of SmbInstance for the instance
func (instance *MSFT_SmbShare) GetPropertySmbInstance() (value SmbShare_SmbInstance, err error) {
	retValue, err := instance.GetProperty("SmbInstance")
	if err != nil {
		return
	}
	value, ok := retValue.(SmbShare_SmbInstance)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetSpecial sets the value of Special for the instance
func (instance *MSFT_SmbShare) SetPropertySpecial(value bool) (err error) {
	return instance.SetProperty("Special", value)
}

// GetSpecial gets the value of Special for the instance
func (instance *MSFT_SmbShare) GetPropertySpecial() (value bool, err error) {
	retValue, err := instance.GetProperty("Special")
	if err != nil {
		return
	}
	value, ok := retValue.(bool)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetTemporary sets the value of Temporary for the instance
func (instance *MSFT_SmbShare) SetPropertyTemporary(value bool) (err error) {
	return instance.SetProperty("Temporary", value)
}

// GetTemporary gets the value of Temporary for the instance
func (instance *MSFT_SmbShare) GetPropertyTemporary() (value bool, err error) {
	retValue, err := instance.GetProperty("Temporary")
	if err != nil {
		return
	}
	value, ok := retValue.(bool)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetVolume sets the value of Volume for the instance
func (instance *MSFT_SmbShare) SetPropertyVolume(value string) (err error) {
	return instance.SetProperty("Volume", value)
}

// GetVolume gets the value of Volume for the instance
func (instance *MSFT_SmbShare) GetPropertyVolume() (value string, err error) {
	retValue, err := instance.GetProperty("Volume")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

//

// <param name="CachingMode" type="uint32 "></param>
// <param name="CATimeout" type="uint32 "></param>
// <param name="ChangeAccess" type="string []"></param>
// <param name="ConcurrentUserLimit" type="uint32 "></param>
// <param name="ContinuouslyAvailable" type="bool "></param>
// <param name="Description" type="string "></param>
// <param name="EncryptData" type="bool "></param>
// <param name="FolderEnumerationMode" type="uint32 "></param>
// <param name="FullAccess" type="string []"></param>
// <param name="LeasingMode" type="uint32 "></param>
// <param name="Name" type="string "></param>
// <param name="NoAccess" type="string []"></param>
// <param name="Path" type="string "></param>
// <param name="ReadAccess" type="string []"></param>
// <param name="ScopeName" type="string "></param>
// <param name="SecurityDescriptor" type="string "></param>
// <param name="Temporary" type="bool "></param>

// <param name="CreatedShare" type="MSFT_SmbShare "></param>
// <param name="ReturnValue" type="uint32 "></param>
func (instance *MSFT_SmbShare) CreateShare( /* IN */ Name string,
	/* IN */ ScopeName string,
	/* IN */ Path string,
	/* IN */ Description string,
	/* IN */ ConcurrentUserLimit uint32,
	/* IN */ FolderEnumerationMode uint32,
	/* IN */ CachingMode uint32,
	/* IN */ Temporary bool,
	/* IN */ ContinuouslyAvailable bool,
	/* IN */ CATimeout uint32,
	/* IN */ EncryptData bool,
	/* IN */ FullAccess []string,
	/* IN */ ChangeAccess []string,
	/* IN */ ReadAccess []string,
	/* IN */ NoAccess []string,
	/* IN */ SecurityDescriptor string,
	/* OUT */ CreatedShare MSFT_SmbShare,
	/* OPTIONAL IN */ LeasingMode uint32) (result uint32, err error) {
	retVal, err := instance.InvokeMethod("CreateShare", Name, ScopeName, Path, Description, ConcurrentUserLimit, FolderEnumerationMode, CachingMode, Temporary, ContinuouslyAvailable, CATimeout, EncryptData, FullAccess, ChangeAccess, ReadAccess, NoAccess, SecurityDescriptor, LeasingMode)
	if err != nil {
		return
	}
	retValue := retVal[0].(int32)
	result = uint32(retValue)
	return

}

//

// <param name="AccessRight" type="uint32 "></param>
// <param name="AccountName" type="string []"></param>

// <param name="Output" type="MSFT_SmbShareAccessControlEntry []"></param>
// <param name="ReturnValue" type="uint32 "></param>
func (instance *MSFT_SmbShare) GrantAccess( /* IN */ AccountName []string,
	/* IN */ AccessRight uint32,
	/* OUT */ Output []MSFT_SmbShareAccessControlEntry) (result uint32, err error) {
	retVal, err := instance.InvokeMethod("GrantAccess", AccountName, AccessRight)
	if err != nil {
		return
	}
	retValue := retVal[0].(int32)
	result = uint32(retValue)
	return

}

//

// <param name="AccountName" type="string []"></param>

// <param name="Output" type="MSFT_SmbShareAccessControlEntry []"></param>
// <param name="ReturnValue" type="uint32 "></param>
func (instance *MSFT_SmbShare) RevokeAccess( /* IN */ AccountName []string,
	/* OUT */ Output []MSFT_SmbShareAccessControlEntry) (result uint32, err error) {
	retVal, err := instance.InvokeMethod("RevokeAccess", AccountName)
	if err != nil {
		return
	}
	retValue := retVal[0].(int32)
	result = uint32(retValue)
	return

}

//

// <param name="AccountName" type="string []"></param>

// <param name="Output" type="MSFT_SmbShareAccessControlEntry []"></param>
// <param name="ReturnValue" type="uint32 "></param>
func (instance *MSFT_SmbShare) BlockAccess( /* IN */ AccountName []string,
	/* OUT */ Output []MSFT_SmbShareAccessControlEntry) (result uint32, err error) {
	retVal, err := instance.InvokeMethod("BlockAccess", AccountName)
	if err != nil {
		return
	}
	retValue := retVal[0].(int32)
	result = uint32(retValue)
	return

}

//

// <param name="AccountName" type="string []"></param>

// <param name="Output" type="MSFT_SmbShareAccessControlEntry []"></param>
// <param name="ReturnValue" type="uint32 "></param>
func (instance *MSFT_SmbShare) UnblockAccess( /* IN */ AccountName []string,
	/* OUT */ Output []MSFT_SmbShareAccessControlEntry) (result uint32, err error) {
	retVal, err := instance.InvokeMethod("UnblockAccess", AccountName)
	if err != nil {
		return
	}
	retValue := retVal[0].(int32)
	result = uint32(retValue)
	return

}

//

// <param name="Output" type="MSFT_SmbShareAccessControlEntry []"></param>
// <param name="ReturnValue" type="uint32 "></param>
func (instance *MSFT_SmbShare) GetAccessControlEntries( /* OUT */ Output []MSFT_SmbShareAccessControlEntry) (result uint32, err error) {
	retVal, err := instance.InvokeMethod("GetAccessControlEntries")
	if err != nil {
		return
	}
	retValue := retVal[0].(int32)
	result = uint32(retValue)
	return

}

//

// <param name="PopulateVolumeProperty" type="bool "></param>
// <param name="ScopeName" type="string "></param>

// <param name="Output" type="MSFT_SmbShare []"></param>
// <param name="ReturnValue" type="uint32 "></param>
func (instance *MSFT_SmbShare) EnumerateShares( /* IN */ ScopeName string,
	/* IN */ PopulateVolumeProperty bool,
	/* OUT */ Output []MSFT_SmbShare) (result uint32, err error) {
	retVal, err := instance.InvokeMethod("EnumerateShares", ScopeName, PopulateVolumeProperty)
	if err != nil {
		return
	}
	retValue := retVal[0].(int32)
	result = uint32(retValue)
	return

}

//

// <param name="GetAclNonAdmin" type="bool "></param>
// <param name="ScopeName" type="string "></param>
// <param name="ShareName" type="string "></param>

// <param name="Output" type="MSFT_SmbShare "></param>
// <param name="ReturnValue" type="uint32 "></param>
func (instance *MSFT_SmbShare) GetShare( /* IN */ ScopeName string,
	/* IN */ ShareName string,
	/* IN */ GetAclNonAdmin bool,
	/* OUT */ Output MSFT_SmbShare) (result uint32, err error) {
	retVal, err := instance.InvokeMethod("GetShare", ScopeName, ShareName, GetAclNonAdmin)
	if err != nil {
		return
	}
	retValue := retVal[0].(int32)
	result = uint32(retValue)
	return

}

//

// <param name="CATimeout" type="uint32 "></param>
// <param name="ConcurrentUserLimit" type="uint32 "></param>
// <param name="EventType" type="uint32 "></param>
// <param name="Flags" type="uint32 "></param>
// <param name="Path" type="string "></param>
// <param name="Remark" type="string "></param>
// <param name="ScopeName" type="string "></param>
// <param name="SecurityDescriptor" type="string "></param>
// <param name="ShareName" type="string "></param>
// <param name="ShareState" type="uint32 "></param>
// <param name="Type" type="uint32 "></param>

// <param name="ReturnValue" type="uint32 "></param>
func (instance *MSFT_SmbShare) FireShareChangeEvent( /* IN */ EventType uint32,
	/* IN */ ScopeName string,
	/* IN */ ShareName string,
	/* IN */ Path string,
	/* IN */ Remark string,
	/* IN */ SecurityDescriptor string,
	/* IN */ ShareState uint32,
	/* IN */ CATimeout uint32,
	/* IN */ Flags uint32,
	/* IN */ Type uint32,
	/* IN */ ConcurrentUserLimit uint32) (result uint32, err error) {
	retVal, err := instance.InvokeMethodWithReturn("FireShareChangeEvent", EventType, ScopeName, ShareName, Path, Remark, SecurityDescriptor, ShareState, CATimeout, Flags, Type, ConcurrentUserLimit)
	if err != nil {
		return
	}
	result = uint32(retVal)
	return

}
