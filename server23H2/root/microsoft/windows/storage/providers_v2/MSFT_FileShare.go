// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 6/14/2024 using wmigen
//      Source root.Microsoft.Windows.Storage.Providers_v2
//////////////////////////////////////////////
package providers_v2

import (
	"github.com/microsoft/wmi/pkg/base/query"
	"github.com/microsoft/wmi/pkg/errors"
	cim "github.com/microsoft/wmi/pkg/wmiinstance"
	"reflect"
)

// MSFT_FileShare struct
type MSFT_FileShare struct {
	*MSFT_StorageObject

	// If TRUE the share is continuously available.
	ContinuouslyAvailable bool

	// A user settable description of the file share. This field can be used to store extra free-form information, such as notes or details about the intended usage.
	Description string

	// If TRUE the share data transmission is encrypted.
	EncryptData bool

	// The file sharing protocol used by the share.
	FileSharingProtocol FileShare_FileSharingProtocol

	// Denotes the current health status of the file share.
	/// 0 - 'Healthy': TBD.
	///1 - 'Warning': TBD.
	///2 - 'Unhealthy': TBD.
	HealthStatus FileShare_HealthStatus

	// Name is a semi-unique (scoped to the owning file server), human-readable string used to access and identify a file share.
	Name string

	// An array of values that denote the current operational status of the fileshare.
	///0 - 'Unknown': The operational status is unknown.
	///1 - 'Other': A vendor-specific OperationalStatus has been specified by setting the OtherOperationalStatusDescription property.
	///2 - 'OK': The disk is responding to commands and is in a normal operating state.
	///3 - 'Degraded': The disk is responding to commands, but is not running in an optimal operating state.
	///4 - 'Stressed': The disk is functioning, but needs attention. For example, the disk might be overloaded or overheated.
	///5 - 'Predictive Failure': The disk is functioning, but a failure is likely to occur in the near future.
	///6 - 'Error': An error has occurred.
	///7 - 'Non-Recoverable Error': A non-recoverable error has occurred.
	///8 - 'Starting': The disk is in the process of starting.
	///9 - 'Stopping': The disk is in the process of stopping.
	///10 - 'Stopped': The disk was stopped or shut down in a clean and orderly fashion.
	///11 - 'In Service': The disk is being configured, maintained, cleaned, or otherwise administered.
	///12 - 'No Contact': The storage provider has knowledge of the disk, but has never been able to establish communication with it.
	///13 - 'Lost Communication': The storage provider has knowledge of the disk and has contacted it successfully in the past, but the disk is currently unreachable.
	///14 - 'Aborted': Similar to Stopped, except that the disk stopped abruptly and may require configuration or maintenance.
	///15 - 'Dormant': The disk is reachable, but it is inactive.
	///16 - 'Supporting Entity in Error': This status value does not necessarily indicate trouble with the disk, but it does indicate that another device or connection that the disk depends on may need attention.
	///17 - 'Completed': The disk has completed an operation. This status value should be combined with OK, Error, or Degraded, depending on the outcome of the operation.
	///0xD010 - 'Online': In Windows-based storage subsystems, this indicates that the object is online.
	///0xD011 - 'Not Ready': In Windows-based storage subsystems, this indicates that the object is not ready.
	///0xD012 - 'No Media': In Windows-based storage subsystems, this indicates that the object has no media present.
	///0xD013 - 'Offline': In Windows-based storage subsystems, this indicates that the object is offline.
	///0xD014 - 'Failed': In Windows-based storage subsystems, this indicates that the object is in a failed state.
	OperationalStatus []FileShare_OperationalStatus

	// TODO
	ShareState FileShare_ShareState

	// The volume relative path to the directory that is being shared.
	VolumeRelativePath string
}

func NewMSFT_FileShareEx1(instance *cim.WmiInstance) (newInstance *MSFT_FileShare, err error) {
	tmp, err := NewMSFT_StorageObjectEx1(instance)

	if err != nil {
		return
	}
	newInstance = &MSFT_FileShare{
		MSFT_StorageObject: tmp,
	}
	return
}

func NewMSFT_FileShareEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *MSFT_FileShare, err error) {
	tmp, err := NewMSFT_StorageObjectEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &MSFT_FileShare{
		MSFT_StorageObject: tmp,
	}
	return
}

// SetContinuouslyAvailable sets the value of ContinuouslyAvailable for the instance
func (instance *MSFT_FileShare) SetPropertyContinuouslyAvailable(value bool) (err error) {
	return instance.SetProperty("ContinuouslyAvailable", (value))
}

// GetContinuouslyAvailable gets the value of ContinuouslyAvailable for the instance
func (instance *MSFT_FileShare) GetPropertyContinuouslyAvailable() (value bool, err error) {
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

// SetDescription sets the value of Description for the instance
func (instance *MSFT_FileShare) SetPropertyDescription(value string) (err error) {
	return instance.SetProperty("Description", (value))
}

// GetDescription gets the value of Description for the instance
func (instance *MSFT_FileShare) GetPropertyDescription() (value string, err error) {
	retValue, err := instance.GetProperty("Description")
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

// SetEncryptData sets the value of EncryptData for the instance
func (instance *MSFT_FileShare) SetPropertyEncryptData(value bool) (err error) {
	return instance.SetProperty("EncryptData", (value))
}

// GetEncryptData gets the value of EncryptData for the instance
func (instance *MSFT_FileShare) GetPropertyEncryptData() (value bool, err error) {
	retValue, err := instance.GetProperty("EncryptData")
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

// SetFileSharingProtocol sets the value of FileSharingProtocol for the instance
func (instance *MSFT_FileShare) SetPropertyFileSharingProtocol(value FileShare_FileSharingProtocol) (err error) {
	return instance.SetProperty("FileSharingProtocol", (value))
}

// GetFileSharingProtocol gets the value of FileSharingProtocol for the instance
func (instance *MSFT_FileShare) GetPropertyFileSharingProtocol() (value FileShare_FileSharingProtocol, err error) {
	retValue, err := instance.GetProperty("FileSharingProtocol")
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

	value = FileShare_FileSharingProtocol(valuetmp)

	return
}

// SetHealthStatus sets the value of HealthStatus for the instance
func (instance *MSFT_FileShare) SetPropertyHealthStatus(value FileShare_HealthStatus) (err error) {
	return instance.SetProperty("HealthStatus", (value))
}

// GetHealthStatus gets the value of HealthStatus for the instance
func (instance *MSFT_FileShare) GetPropertyHealthStatus() (value FileShare_HealthStatus, err error) {
	retValue, err := instance.GetProperty("HealthStatus")
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

	value = FileShare_HealthStatus(valuetmp)

	return
}

// SetName sets the value of Name for the instance
func (instance *MSFT_FileShare) SetPropertyName(value string) (err error) {
	return instance.SetProperty("Name", (value))
}

// GetName gets the value of Name for the instance
func (instance *MSFT_FileShare) GetPropertyName() (value string, err error) {
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

// SetOperationalStatus sets the value of OperationalStatus for the instance
func (instance *MSFT_FileShare) SetPropertyOperationalStatus(value []FileShare_OperationalStatus) (err error) {
	return instance.SetProperty("OperationalStatus", (value))
}

// GetOperationalStatus gets the value of OperationalStatus for the instance
func (instance *MSFT_FileShare) GetPropertyOperationalStatus() (value []FileShare_OperationalStatus, err error) {
	retValue, err := instance.GetProperty("OperationalStatus")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	for _, interfaceValue := range retValue.([]interface{}) {
		valuetmp, ok := interfaceValue.(int32)
		if !ok {
			err = errors.Wrapf(errors.InvalidType, " int32 is Invalid. Expected %s", reflect.TypeOf(interfaceValue))
			return
		}
		value = append(value, FileShare_OperationalStatus(valuetmp))
	}

	return
}

// SetShareState sets the value of ShareState for the instance
func (instance *MSFT_FileShare) SetPropertyShareState(value FileShare_ShareState) (err error) {
	return instance.SetProperty("ShareState", (value))
}

// GetShareState gets the value of ShareState for the instance
func (instance *MSFT_FileShare) GetPropertyShareState() (value FileShare_ShareState, err error) {
	retValue, err := instance.GetProperty("ShareState")
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

	value = FileShare_ShareState(valuetmp)

	return
}

// SetVolumeRelativePath sets the value of VolumeRelativePath for the instance
func (instance *MSFT_FileShare) SetPropertyVolumeRelativePath(value string) (err error) {
	return instance.SetProperty("VolumeRelativePath", (value))
}

// GetVolumeRelativePath gets the value of VolumeRelativePath for the instance
func (instance *MSFT_FileShare) GetPropertyVolumeRelativePath() (value string, err error) {
	retValue, err := instance.GetProperty("VolumeRelativePath")
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

//

// <param name="CreatedStorageJob" type="MSFT_StorageJob "></param>
// <param name="ExtendedStatus" type="MSFT_StorageExtendedStatus "></param>
// <param name="ReturnValue" type="uint32 "></param>
func (instance *MSFT_FileShare) DeleteObject( /* OUT */ CreatedStorageJob MSFT_StorageJob,
	/* OUT */ ExtendedStatus MSFT_StorageExtendedStatus) (result uint32, err error) {
	retVal, err := instance.InvokeMethod("DeleteObject")
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
