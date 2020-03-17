// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 3/16/2020 using wmigen
//      Source root.Microsoft.Windows.Storage
//////////////////////////////////////////////
package storage

// MSFT_FileShare struct
type MSFT_FileShare struct {
	MSFT_StorageObject

	//
	ContinuouslyAvailable bool

	//
	Description string

	//
	EncryptData bool

	//
	FileSharingProtocol uint16

	//
	HealthStatus uint16

	//
	Name string

	//
	OperationalStatus []uint16

	//
	ShareState uint16

	//
	VolumeRelativePath string
}

// SetContinuouslyAvailable sets the value of ContinuouslyAvailable for the instance
func (instance *MSFT_FileShare) SetPropertyContinuouslyAvailable(value bool) (err error) {
	return instance.SetProperty("ContinuouslyAvailable", value)
}

// GetContinuouslyAvailable gets the value of ContinuouslyAvailable for the instance
func (instance *MSFT_FileShare) GetPropertyContinuouslyAvailable() (value bool, err error) {
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

// SetDescription sets the value of Description for the instance
func (instance *MSFT_FileShare) SetPropertyDescription(value string) (err error) {
	return instance.SetProperty("Description", value)
}

// GetDescription gets the value of Description for the instance
func (instance *MSFT_FileShare) GetPropertyDescription() (value string, err error) {
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
func (instance *MSFT_FileShare) SetPropertyEncryptData(value bool) (err error) {
	return instance.SetProperty("EncryptData", value)
}

// GetEncryptData gets the value of EncryptData for the instance
func (instance *MSFT_FileShare) GetPropertyEncryptData() (value bool, err error) {
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

// SetFileSharingProtocol sets the value of FileSharingProtocol for the instance
func (instance *MSFT_FileShare) SetPropertyFileSharingProtocol(value uint16) (err error) {
	return instance.SetProperty("FileSharingProtocol", value)
}

// GetFileSharingProtocol gets the value of FileSharingProtocol for the instance
func (instance *MSFT_FileShare) GetPropertyFileSharingProtocol() (value uint16, err error) {
	retValue, err := instance.GetProperty("FileSharingProtocol")
	if err != nil {
		return
	}
	value, ok := retValue.(uint16)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetHealthStatus sets the value of HealthStatus for the instance
func (instance *MSFT_FileShare) SetPropertyHealthStatus(value uint16) (err error) {
	return instance.SetProperty("HealthStatus", value)
}

// GetHealthStatus gets the value of HealthStatus for the instance
func (instance *MSFT_FileShare) GetPropertyHealthStatus() (value uint16, err error) {
	retValue, err := instance.GetProperty("HealthStatus")
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
func (instance *MSFT_FileShare) SetPropertyName(value string) (err error) {
	return instance.SetProperty("Name", value)
}

// GetName gets the value of Name for the instance
func (instance *MSFT_FileShare) GetPropertyName() (value string, err error) {
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

// SetOperationalStatus sets the value of OperationalStatus for the instance
func (instance *MSFT_FileShare) SetPropertyOperationalStatus(value []uint16) (err error) {
	return instance.SetProperty("OperationalStatus", value)
}

// GetOperationalStatus gets the value of OperationalStatus for the instance
func (instance *MSFT_FileShare) GetPropertyOperationalStatus() (value []uint16, err error) {
	retValue, err := instance.GetProperty("OperationalStatus")
	if err != nil {
		return
	}
	value, ok := retValue.([]uint16)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetShareState sets the value of ShareState for the instance
func (instance *MSFT_FileShare) SetPropertyShareState(value uint16) (err error) {
	return instance.SetProperty("ShareState", value)
}

// GetShareState gets the value of ShareState for the instance
func (instance *MSFT_FileShare) GetPropertyShareState() (value uint16, err error) {
	retValue, err := instance.GetProperty("ShareState")
	if err != nil {
		return
	}
	value, ok := retValue.(uint16)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetVolumeRelativePath sets the value of VolumeRelativePath for the instance
func (instance *MSFT_FileShare) SetPropertyVolumeRelativePath(value string) (err error) {
	return instance.SetProperty("VolumeRelativePath", value)
}

// GetVolumeRelativePath gets the value of VolumeRelativePath for the instance
func (instance *MSFT_FileShare) GetPropertyVolumeRelativePath() (value string, err error) {
	retValue, err := instance.GetProperty("VolumeRelativePath")
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

// <param name="RunAsJob" type="bool "></param>

// <param name="CreatedStorageJob" type="MSFT_StorageJob "></param>
// <param name="ExtendedStatus" type="MSFT_StorageExtendedStatus "></param>
// <param name="ReturnValue" type="uint32 "></param>
func (instance *MSFT_FileShare) DeleteObject( /* IN */ RunAsJob bool,
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

// <param name="Description" type="string "></param>

// <param name="ExtendedStatus" type="MSFT_StorageExtendedStatus "></param>
// <param name="ReturnValue" type="uint32 "></param>
func (instance *MSFT_FileShare) SetDescription( /* IN */ Description string,
	/* OUT */ ExtendedStatus MSFT_StorageExtendedStatus) (result uint32, err error) {
	retVal, err := instance.InvokeMethod("SetDescription", Description)
	if err != nil {
		return
	}
	retValue := retVal[0].(int32)
	result = uint32(retValue)
	return

}

//

// <param name="EncryptData" type="bool "></param>

// <param name="ExtendedStatus" type="MSFT_StorageExtendedStatus "></param>
// <param name="ReturnValue" type="uint32 "></param>
func (instance *MSFT_FileShare) SetAttributes( /* IN */ EncryptData bool,
	/* OUT */ ExtendedStatus MSFT_StorageExtendedStatus) (result uint32, err error) {
	retVal, err := instance.InvokeMethod("SetAttributes", EncryptData)
	if err != nil {
		return
	}
	retValue := retVal[0].(int32)
	result = uint32(retValue)
	return

}

//

// <param name="AccessControlEntries" type="MSFT_FileShareAccessControlEntry []"></param>
// <param name="ExtendedStatus" type="MSFT_StorageExtendedStatus "></param>
// <param name="ReturnValue" type="uint32 "></param>
func (instance *MSFT_FileShare) GetAccessControlEntries( /* OUT */ AccessControlEntries []MSFT_FileShareAccessControlEntry,
	/* OUT */ ExtendedStatus MSFT_StorageExtendedStatus) (result uint32, err error) {
	retVal, err := instance.InvokeMethod("GetAccessControlEntries")
	if err != nil {
		return
	}
	retValue := retVal[0].(int32)
	result = uint32(retValue)
	return

}

//

// <param name="AccessRight" type="uint32 "></param>
// <param name="AccountNames" type="string []"></param>

// <param name="ExtendedStatus" type="MSFT_StorageExtendedStatus "></param>
// <param name="ReturnValue" type="uint32 "></param>
func (instance *MSFT_FileShare) GrantAccess( /* IN */ AccountNames []string,
	/* IN */ AccessRight uint32,
	/* OUT */ ExtendedStatus MSFT_StorageExtendedStatus) (result uint32, err error) {
	retVal, err := instance.InvokeMethod("GrantAccess", AccountNames, AccessRight)
	if err != nil {
		return
	}
	retValue := retVal[0].(int32)
	result = uint32(retValue)
	return

}

//

// <param name="AccountNames" type="string []"></param>

// <param name="ExtendedStatus" type="MSFT_StorageExtendedStatus "></param>
// <param name="ReturnValue" type="uint32 "></param>
func (instance *MSFT_FileShare) RevokeAccess( /* IN */ AccountNames []string,
	/* OUT */ ExtendedStatus MSFT_StorageExtendedStatus) (result uint32, err error) {
	retVal, err := instance.InvokeMethod("RevokeAccess", AccountNames)
	if err != nil {
		return
	}
	retValue := retVal[0].(int32)
	result = uint32(retValue)
	return

}

//

// <param name="AccountNames" type="string []"></param>

// <param name="ExtendedStatus" type="MSFT_StorageExtendedStatus "></param>
// <param name="ReturnValue" type="uint32 "></param>
func (instance *MSFT_FileShare) BlockAccess( /* IN */ AccountNames []string,
	/* OUT */ ExtendedStatus MSFT_StorageExtendedStatus) (result uint32, err error) {
	retVal, err := instance.InvokeMethod("BlockAccess", AccountNames)
	if err != nil {
		return
	}
	retValue := retVal[0].(int32)
	result = uint32(retValue)
	return

}

//

// <param name="AccountNames" type="string []"></param>

// <param name="ExtendedStatus" type="MSFT_StorageExtendedStatus "></param>
// <param name="ReturnValue" type="uint32 "></param>
func (instance *MSFT_FileShare) UnblockAccess( /* IN */ AccountNames []string,
	/* OUT */ ExtendedStatus MSFT_StorageExtendedStatus) (result uint32, err error) {
	retVal, err := instance.InvokeMethod("UnblockAccess", AccountNames)
	if err != nil {
		return
	}
	retValue := retVal[0].(int32)
	result = uint32(retValue)
	return

}

//

// <param name="DiagnoseResults" type="MSFT_StorageDiagnoseResult []"></param>
// <param name="ExtendedStatus" type="MSFT_StorageExtendedStatus "></param>
// <param name="ReturnValue" type="uint32 "></param>
func (instance *MSFT_FileShare) Diagnose( /* OUT */ DiagnoseResults []MSFT_StorageDiagnoseResult,
	/* OUT */ ExtendedStatus MSFT_StorageExtendedStatus) (result uint32, err error) {
	retVal, err := instance.InvokeMethod("Diagnose")
	if err != nil {
		return
	}
	retValue := retVal[0].(int32)
	result = uint32(retValue)
	return

}

//

// <param name="ActionResults" type="MSFT_HealthAction []"></param>
// <param name="ExtendedStatus" type="MSFT_StorageExtendedStatus "></param>
// <param name="ReturnValue" type="uint32 "></param>
func (instance *MSFT_FileShare) GetActions( /* OUT */ ActionResults []MSFT_HealthAction,
	/* OUT */ ExtendedStatus MSFT_StorageExtendedStatus) (result uint32, err error) {
	retVal, err := instance.InvokeMethod("GetActions")
	if err != nil {
		return
	}
	retValue := retVal[0].(int32)
	result = uint32(retValue)
	return

}
