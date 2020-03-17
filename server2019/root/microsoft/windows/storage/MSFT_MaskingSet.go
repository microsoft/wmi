// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 3/16/2020 using wmigen
//      Source root.Microsoft.Windows.Storage
//////////////////////////////////////////////
package storage

// MSFT_MaskingSet struct
type MSFT_MaskingSet struct {
	MSFT_StorageObject

	//
	FriendlyName string

	//
	HostType uint16

	//
	Name string
}

// SetFriendlyName sets the value of FriendlyName for the instance
func (instance *MSFT_MaskingSet) SetPropertyFriendlyName(value string) (err error) {
	return instance.SetProperty("FriendlyName", value)
}

// GetFriendlyName gets the value of FriendlyName for the instance
func (instance *MSFT_MaskingSet) GetPropertyFriendlyName() (value string, err error) {
	retValue, err := instance.GetProperty("FriendlyName")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetHostType sets the value of HostType for the instance
func (instance *MSFT_MaskingSet) SetPropertyHostType(value uint16) (err error) {
	return instance.SetProperty("HostType", value)
}

// GetHostType gets the value of HostType for the instance
func (instance *MSFT_MaskingSet) GetPropertyHostType() (value uint16, err error) {
	retValue, err := instance.GetProperty("HostType")
	if err != nil {
		return
	}
	value, ok := retValue.(uint16)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetName sets the value of Name for the instance
func (instance *MSFT_MaskingSet) SetPropertyName(value string) (err error) {
	return instance.SetProperty("Name", value)
}

// GetName gets the value of Name for the instance
func (instance *MSFT_MaskingSet) GetPropertyName() (value string, err error) {
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

//

// <param name="HostType" type="uint16 "></param>
// <param name="InitiatorIds" type="string []"></param>
// <param name="RunAsJob" type="bool "></param>

// <param name="CreatedStorageJob" type="MSFT_StorageJob "></param>
// <param name="ExtendedStatus" type="MSFT_StorageExtendedStatus "></param>
// <param name="ReturnValue" type="uint32 "></param>
func (instance *MSFT_MaskingSet) AddInitiatorId( /* IN */ InitiatorIds []string,
	/* IN */ HostType uint16,
	/* IN */ RunAsJob bool,
	/* OUT */ CreatedStorageJob MSFT_StorageJob,
	/* OUT */ ExtendedStatus MSFT_StorageExtendedStatus) (result uint32, err error) {
	retVal, err := instance.InvokeMethod("AddInitiatorId", InitiatorIds, HostType, RunAsJob)
	if err != nil {
		return
	}
	retValue := retVal[0].(int32)
	result = uint32(retValue)
	return

}

//

// <param name="InitiatorIds" type="string []"></param>
// <param name="RunAsJob" type="bool "></param>

// <param name="CreatedStorageJob" type="MSFT_StorageJob "></param>
// <param name="ExtendedStatus" type="MSFT_StorageExtendedStatus "></param>
// <param name="ReturnValue" type="uint32 "></param>
func (instance *MSFT_MaskingSet) RemoveInitiatorId( /* IN */ InitiatorIds []string,
	/* IN */ RunAsJob bool,
	/* OUT */ CreatedStorageJob MSFT_StorageJob,
	/* OUT */ ExtendedStatus MSFT_StorageExtendedStatus) (result uint32, err error) {
	retVal, err := instance.InvokeMethod("RemoveInitiatorId", InitiatorIds, RunAsJob)
	if err != nil {
		return
	}
	retValue := retVal[0].(int32)
	result = uint32(retValue)
	return

}

//

// <param name="RunAsJob" type="bool "></param>
// <param name="TargetPortAddresses" type="string []"></param>

// <param name="CreatedStorageJob" type="MSFT_StorageJob "></param>
// <param name="ExtendedStatus" type="MSFT_StorageExtendedStatus "></param>
// <param name="ReturnValue" type="uint32 "></param>
func (instance *MSFT_MaskingSet) AddTargetPort( /* IN */ TargetPortAddresses []string,
	/* IN */ RunAsJob bool,
	/* OUT */ CreatedStorageJob MSFT_StorageJob,
	/* OUT */ ExtendedStatus MSFT_StorageExtendedStatus) (result uint32, err error) {
	retVal, err := instance.InvokeMethod("AddTargetPort", TargetPortAddresses, RunAsJob)
	if err != nil {
		return
	}
	retValue := retVal[0].(int32)
	result = uint32(retValue)
	return

}

//

// <param name="RunAsJob" type="bool "></param>
// <param name="TargetPortAddresses" type="string []"></param>

// <param name="CreatedStorageJob" type="MSFT_StorageJob "></param>
// <param name="ExtendedStatus" type="MSFT_StorageExtendedStatus "></param>
// <param name="ReturnValue" type="uint32 "></param>
func (instance *MSFT_MaskingSet) RemoveTargetPort( /* IN */ TargetPortAddresses []string,
	/* IN */ RunAsJob bool,
	/* OUT */ CreatedStorageJob MSFT_StorageJob,
	/* OUT */ ExtendedStatus MSFT_StorageExtendedStatus) (result uint32, err error) {
	retVal, err := instance.InvokeMethod("RemoveTargetPort", TargetPortAddresses, RunAsJob)
	if err != nil {
		return
	}
	retValue := retVal[0].(int32)
	result = uint32(retValue)
	return

}

//

// <param name="DeviceAccesses" type="uint16 []"></param>
// <param name="DeviceNumbers" type="string []"></param>
// <param name="RunAsJob" type="bool "></param>
// <param name="VirtualDiskNames" type="string []"></param>

// <param name="CreatedStorageJob" type="MSFT_StorageJob "></param>
// <param name="ExtendedStatus" type="MSFT_StorageExtendedStatus "></param>
// <param name="ReturnValue" type="uint32 "></param>
func (instance *MSFT_MaskingSet) AddVirtualDisk( /* IN */ VirtualDiskNames []string,
	/* IN */ DeviceNumbers []string,
	/* IN */ DeviceAccesses []uint16,
	/* IN */ RunAsJob bool,
	/* OUT */ CreatedStorageJob MSFT_StorageJob,
	/* OUT */ ExtendedStatus MSFT_StorageExtendedStatus) (result uint32, err error) {
	retVal, err := instance.InvokeMethod("AddVirtualDisk", VirtualDiskNames, DeviceNumbers, DeviceAccesses, RunAsJob)
	if err != nil {
		return
	}
	retValue := retVal[0].(int32)
	result = uint32(retValue)
	return

}

//

// <param name="RunAsJob" type="bool "></param>
// <param name="VirtualDiskNames" type="string []"></param>

// <param name="CreatedStorageJob" type="MSFT_StorageJob "></param>
// <param name="ExtendedStatus" type="MSFT_StorageExtendedStatus "></param>
// <param name="ReturnValue" type="uint32 "></param>
func (instance *MSFT_MaskingSet) RemoveVirtualDisk( /* IN */ VirtualDiskNames []string,
	/* IN */ RunAsJob bool,
	/* OUT */ CreatedStorageJob MSFT_StorageJob,
	/* OUT */ ExtendedStatus MSFT_StorageExtendedStatus) (result uint32, err error) {
	retVal, err := instance.InvokeMethod("RemoveVirtualDisk", VirtualDiskNames, RunAsJob)
	if err != nil {
		return
	}
	retValue := retVal[0].(int32)
	result = uint32(retValue)
	return

}

//

// <param name="RunAsJob" type="bool "></param>

// <param name="CreatedStorageJob" type="MSFT_StorageJob "></param>
// <param name="ExtendedStatus" type="MSFT_StorageExtendedStatus "></param>
// <param name="ReturnValue" type="uint32 "></param>
func (instance *MSFT_MaskingSet) DeleteObject( /* IN */ RunAsJob bool,
	/* OUT */ CreatedStorageJob MSFT_StorageJob,
	/* OUT */ ExtendedStatus MSFT_StorageExtendedStatus) (result uint32, err error) {
	retVal, err := instance.InvokeMethod("DeleteObject", RunAsJob)
	if err != nil {
		return
	}
	retValue := retVal[0].(int32)
	result = uint32(retValue)
	return

}

//

// <param name="FriendlyName" type="string "></param>

// <param name="ExtendedStatus" type="MSFT_StorageExtendedStatus "></param>
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

//

// <param name="ExtendedStatus" type="MSFT_StorageExtendedStatus "></param>
// <param name="ReturnValue" type="uint32 "></param>
// <param name="SecurityDescriptor" type="string "></param>
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

//

// <param name="SecurityDescriptor" type="string "></param>

// <param name="ExtendedStatus" type="MSFT_StorageExtendedStatus "></param>
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
