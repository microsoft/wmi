// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 9/18/2020 using wmigen
//      Source root.Microsoft.Windows.Storage.Providers_v2
//////////////////////////////////////////////
package providers_v2

import (
	"github.com/microsoft/wmi/pkg/base/query"
	"github.com/microsoft/wmi/pkg/errors"
	cim "github.com/microsoft/wmi/pkg/wmiinstance"
	"reflect"
)

// MSFT_MaskingSet struct
type MSFT_MaskingSet struct {
	*MSFT_StorageObject

	// FriendlyName is a user-friendly name of the masking set. It is specified during the creation of the masking set, and can be changed using the SetFriendlyName method.
	FriendlyName string

	// This field specifies the operating system, version, driver, and other host environment factors that influence the behavior exposed by the storage subsystem.
	HostType MaskingSet_HostType

	// Name is a user-friendly system defined name for the masking set. Name is unique within the scope of the owning storage subsystem.
	Name string
}

func NewMSFT_MaskingSetEx1(instance *cim.WmiInstance) (newInstance *MSFT_MaskingSet, err error) {
	tmp, err := NewMSFT_StorageObjectEx1(instance)

	if err != nil {
		return
	}
	newInstance = &MSFT_MaskingSet{
		MSFT_StorageObject: tmp,
	}
	return
}

func NewMSFT_MaskingSetEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *MSFT_MaskingSet, err error) {
	tmp, err := NewMSFT_StorageObjectEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &MSFT_MaskingSet{
		MSFT_StorageObject: tmp,
	}
	return
}

// SetFriendlyName sets the value of FriendlyName for the instance
func (instance *MSFT_MaskingSet) SetPropertyFriendlyName(value string) (err error) {
	return instance.SetProperty("FriendlyName", (value))
}

// GetFriendlyName gets the value of FriendlyName for the instance
func (instance *MSFT_MaskingSet) GetPropertyFriendlyName() (value string, err error) {
	retValue, err := instance.GetProperty("FriendlyName")
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

// SetHostType sets the value of HostType for the instance
func (instance *MSFT_MaskingSet) SetPropertyHostType(value MaskingSet_HostType) (err error) {
	return instance.SetProperty("HostType", (value))
}

// GetHostType gets the value of HostType for the instance
func (instance *MSFT_MaskingSet) GetPropertyHostType() (value MaskingSet_HostType, err error) {
	retValue, err := instance.GetProperty("HostType")
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

	value = MaskingSet_HostType(valuetmp)

	return
}

// SetName sets the value of Name for the instance
func (instance *MSFT_MaskingSet) SetPropertyName(value string) (err error) {
	return instance.SetProperty("Name", (value))
}

// GetName gets the value of Name for the instance
func (instance *MSFT_MaskingSet) GetPropertyName() (value string, err error) {
	retValue, err := instance.GetProperty("Name")
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

// This method adds an initiator to the masking set. All virtual disks in the masking set will be accessible (shown) to these initiators.

// <param name="HostType" type="MaskingSet_HostType "></param>
// <param name="InitiatorIds" type="string []">This parameter is an array of initiator addresses. For each address contained in this array, a corresponding initiator ID instance should be created and then associated with this masking set.</param>

// <param name="CreatedStorageJob" type="MSFT_StorageJob ">If RunAsJob is set to TRUE and this method takes a while to execute, this parameter returns a reference to the storage job used to track the long running operation.</param>
// <param name="ExtendedStatus" type="MSFT_StorageExtendedStatus ">ExtendedStatus allows the storage provider to return extended (implementation specific) error information.</param>
// <param name="ReturnValue" type="uint32 "></param>
func (instance *MSFT_MaskingSet) AddInitiatorId( /* IN */ InitiatorIds []string,
	/* IN */ HostType MaskingSet_HostType,
	/* OUT */ CreatedStorageJob MSFT_StorageJob,
	/* OUT */ ExtendedStatus MSFT_StorageExtendedStatus) (result uint32, err error) {
	retVal, err := instance.InvokeMethod("AddInitiatorId", InitiatorIds, HostType)
	if err != nil {
		return
	}
	retValue := retVal[0].(int32)
	result = uint32(retValue)
	return

}

// This method removes one or more initiator ids from the masking set. Note that the initiator id instances themselves should not be deleted from the system.

// <param name="InitiatorIds" type="string []"></param>

// <param name="CreatedStorageJob" type="MSFT_StorageJob ">If RunAsJob is set to TRUE and this method takes a while to execute, this parameter returns a reference to the storage job used to track the long running operation.</param>
// <param name="ExtendedStatus" type="MSFT_StorageExtendedStatus ">ExtendedStatus allows the storage provider to return extended (implementation specific) error information.</param>
// <param name="ReturnValue" type="uint32 "></param>
func (instance *MSFT_MaskingSet) RemoveInitiatorId( /* IN */ InitiatorIds []string,
	/* OUT */ CreatedStorageJob MSFT_StorageJob,
	/* OUT */ ExtendedStatus MSFT_StorageExtendedStatus) (result uint32, err error) {
	retVal, err := instance.InvokeMethod("RemoveInitiatorId", InitiatorIds)
	if err != nil {
		return
	}
	retValue := retVal[0].(int32)
	result = uint32(retValue)
	return

}

// This method adds one or more target ports to the masking set.

// <param name="TargetPortAddresses" type="string []"></param>

// <param name="CreatedStorageJob" type="MSFT_StorageJob ">If RunAsJob is set to TRUE and this method takes a while to execute, this parameter returns a reference to the storage job used to track the long running operation.</param>
// <param name="ExtendedStatus" type="MSFT_StorageExtendedStatus ">ExtendedStatus allows the storage provider to return extended (implementation specific) error information.</param>
// <param name="ReturnValue" type="uint32 "></param>
func (instance *MSFT_MaskingSet) AddTargetPort( /* IN */ TargetPortAddresses []string,
	/* OUT */ CreatedStorageJob MSFT_StorageJob,
	/* OUT */ ExtendedStatus MSFT_StorageExtendedStatus) (result uint32, err error) {
	retVal, err := instance.InvokeMethod("AddTargetPort", TargetPortAddresses)
	if err != nil {
		return
	}
	retValue := retVal[0].(int32)
	result = uint32(retValue)
	return

}

// This method removes one or more target ports from the masking set.

// <param name="TargetPortAddresses" type="string []"></param>

// <param name="CreatedStorageJob" type="MSFT_StorageJob ">If RunAsJob is set to TRUE and this method takes a while to execute, this parameter returns a reference to the storage job used to track the long running operation.</param>
// <param name="ExtendedStatus" type="MSFT_StorageExtendedStatus ">ExtendedStatus allows the storage provider to return extended (implementation specific) error information.</param>
// <param name="ReturnValue" type="uint32 "></param>
func (instance *MSFT_MaskingSet) RemoveTargetPort( /* IN */ TargetPortAddresses []string,
	/* OUT */ CreatedStorageJob MSFT_StorageJob,
	/* OUT */ ExtendedStatus MSFT_StorageExtendedStatus) (result uint32, err error) {
	retVal, err := instance.InvokeMethod("RemoveTargetPort", TargetPortAddresses)
	if err != nil {
		return
	}
	retValue := retVal[0].(int32)
	result = uint32(retValue)
	return

}

// This method adds a virtual disk to the masking set, allowing it to be shown to the initiators contained in the set.

// <param name="DeviceAccesses" type="MaskingSet_DeviceAccesses []"></param>
// <param name="DeviceNumbers" type="string []"></param>
// <param name="VirtualDiskNames" type="string []"></param>

// <param name="CreatedStorageJob" type="MSFT_StorageJob ">If RunAsJob is set to TRUE and this method takes a while to execute, this parameter returns a reference to the storage job used to track the long running operation.</param>
// <param name="ExtendedStatus" type="MSFT_StorageExtendedStatus ">ExtendedStatus allows the storage provider to return extended (implementation specific) error information.</param>
// <param name="ReturnValue" type="uint32 "></param>
func (instance *MSFT_MaskingSet) AddVirtualDisk( /* IN */ VirtualDiskNames []string,
	/* IN */ DeviceNumbers []string,
	/* IN */ DeviceAccesses []MaskingSet_DeviceAccesses,
	/* OUT */ CreatedStorageJob MSFT_StorageJob,
	/* OUT */ ExtendedStatus MSFT_StorageExtendedStatus) (result uint32, err error) {
	retVal, err := instance.InvokeMethod("AddVirtualDisk", VirtualDiskNames, DeviceNumbers, DeviceAccesses)
	if err != nil {
		return
	}
	retValue := retVal[0].(int32)
	result = uint32(retValue)
	return

}

// This method removes a virtual disk from the masking set. Once removed, this virtual disk will no longer be shown to the initiators contained in this masking set.

// <param name="VirtualDiskNames" type="string []"></param>

// <param name="CreatedStorageJob" type="MSFT_StorageJob ">If RunAsJob is set to TRUE and this method takes a while to execute, this parameter returns a reference to the storage job used to track the long running operation.</param>
// <param name="ExtendedStatus" type="MSFT_StorageExtendedStatus ">ExtendedStatus allows the storage provider to return extended (implementation specific) error information.</param>
// <param name="ReturnValue" type="uint32 "></param>
func (instance *MSFT_MaskingSet) RemoveVirtualDisk( /* IN */ VirtualDiskNames []string,
	/* OUT */ CreatedStorageJob MSFT_StorageJob,
	/* OUT */ ExtendedStatus MSFT_StorageExtendedStatus) (result uint32, err error) {
	retVal, err := instance.InvokeMethod("RemoveVirtualDisk", VirtualDiskNames)
	if err != nil {
		return
	}
	retValue := retVal[0].(int32)
	result = uint32(retValue)
	return

}

// This method deletes the masking set instance.

// <param name="CreatedStorageJob" type="MSFT_StorageJob ">If RunAsJob is set to TRUE and this method takes a while to execute, this parameter returns a reference to the storage job used to track the long running operation.</param>
// <param name="ExtendedStatus" type="MSFT_StorageExtendedStatus ">ExtendedStatus allows the storage provider to return extended (implementation specific) error information.</param>
// <param name="ReturnValue" type="uint32 "></param>
func (instance *MSFT_MaskingSet) DeleteObject( /* OUT */ CreatedStorageJob MSFT_StorageJob,
	/* OUT */ ExtendedStatus MSFT_StorageExtendedStatus) (result uint32, err error) {
	retVal, err := instance.InvokeMethod("DeleteObject")
	if err != nil {
		return
	}
	retValue := retVal[0].(int32)
	result = uint32(retValue)
	return

}

// This method allows the FriendlyName to be set.

// <param name="FriendlyName" type="string "></param>

// <param name="ExtendedStatus" type="MSFT_StorageExtendedStatus ">ExtendedStatus allows the storage provider to return extended (implementation specific) error information.</param>
// <param name="ReturnValue" type="uint32 "></param>
func (instance *MSFT_MaskingSet) SetFriendlyName( /* IN */ FriendlyName string,
	/* OUT */ ExtendedStatus MSFT_StorageExtendedStatus) (result uint32, err error) {
	retVal, err := instance.InvokeMethod("SetFriendlyName", FriendlyName)
	if err != nil {
		return
	}
	retValue := retVal[0].(int32)
	result = uint32(retValue)
	return

}

// This method returns the security descriptor that controls access to this specific object instance.

// <param name="ExtendedStatus" type="MSFT_StorageExtendedStatus ">ExtendedStatus allows the storage provider to return extended (implementation specific) error information.</param>
// <param name="ReturnValue" type="uint32 "></param>
// <param name="SecurityDescriptor" type="string ">A Security Descriptor Definition Language (SDDL) formed string describing the access control list of the object.</param>
func (instance *MSFT_MaskingSet) GetSecurityDescriptor( /* OUT */ SecurityDescriptor string,
	/* OUT */ ExtendedStatus MSFT_StorageExtendedStatus) (result uint32, err error) {
	retVal, err := instance.InvokeMethod("GetSecurityDescriptor")
	if err != nil {
		return
	}
	retValue := retVal[0].(int32)
	result = uint32(retValue)
	return

}

// This method allows a user with sufficient privileges to set the security descriptor that control access to this specific object instance. If the call is not made in the context of a user specified in the security descriptor's access control list, this method will fail with 40001 - 'Access Denied'. If an empty security descriptor is passed to this function, the behavior is left to the specific implementation so long as there is some user context (typically domain administrators) that can access and administer the object.

// <param name="SecurityDescriptor" type="string ">A Security Descriptor Definition Language (SDDL) formed string describing the desired access control list for this object.</param>

// <param name="ExtendedStatus" type="MSFT_StorageExtendedStatus ">ExtendedStatus allows the storage provider to return extended (implementation specific) error information.</param>
// <param name="ReturnValue" type="uint32 "></param>
func (instance *MSFT_MaskingSet) SetSecurityDescriptor( /* IN */ SecurityDescriptor string,
	/* OUT */ ExtendedStatus MSFT_StorageExtendedStatus) (result uint32, err error) {
	retVal, err := instance.InvokeMethod("SetSecurityDescriptor", SecurityDescriptor)
	if err != nil {
		return
	}
	retValue := retVal[0].(int32)
	result = uint32(retValue)
	return

}
