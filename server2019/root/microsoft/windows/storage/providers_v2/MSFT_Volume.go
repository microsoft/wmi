// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

// 
// Author:
//      Auto Generated on 6/6/2024 using wmigen
//      Source root.Microsoft.Windows.Storage.Providers_v2
//////////////////////////////////////////////
package providers_v2
import (
 "github.com/microsoft/wmi/pkg/base/query"
cim "github.com/microsoft/wmi/pkg/wmiinstance"
 "github.com/microsoft/wmi/pkg/errors"
 "reflect"
)

// MSFT_Volume struct
type MSFT_Volume struct { 
	*MSFT_StorageObject

	// The allocation unit size of the volume.
	AllocationUnitSize uint32

	// Indicates the deduplication mode of the volume.
	DedupMode Volume_DedupMode

	// Drive letter assigned to the volume.
	DriveLetter byte

	// Denotes the type of the volume.
	DriveType Volume_DriveType

	// File system on the volume.
	FileSystem string

	// File system label of the volume.
	FileSystemLabel string

	// The underlying file system type of the volume.
	FileSystemType Volume_FileSystemType

	// The health status of the Volume.
	///0 - 'Healthy': The volume is functioning normally.
	///1 - 'Warning': The volume is still functioning, but has detected errors or issues that require administrator intervention.
	///2 - 'Unhealthy': The volume is not functioning, due to errors or failures. The volume needs immediate attention from an administrator.
	HealthStatus Volume_HealthStatus

	// An array of values that denote the current operational status of the volume.
	///0 - 'Unknown': The operational status is unknown.
	///1 - 'Other': A vendor-specific OperationalStatus has been specified by setting the OtherOperationalStatusDescription property.
	///2 - 'OK': The volume is responding to commands and is in a normal operating state.
	///3 - 'Degraded': The volume is responding to commands, but is not running in an optimal operating state.
	///4 - 'Stressed': The volume is functioning, but needs attention. For example, the volume might be overloaded or overheated.
	///5 - 'Predictive Failure': The volume is functioning, but a failure is likely to occur in the near future.
	///6 - 'Error': An error has occurred.
	///7 - 'Non-Recoverable Error': A non-recoverable error has occurred.
	///8 - 'Starting': The volume is in the process of starting.
	///9 - 'Stopping': The volume is in the process of stopping.
	///10 - 'Stopped': The volume was stopped or shut down in a clean and orderly fashion.
	///11 - 'In Service': The volume is being configured, maintained, cleaned, or otherwise administered.
	///12 - 'No Contact': The storage provider has knowledge of the volume, but has never been able to establish communication with it.
	///13 - 'Lost Communication': The storage provider has knowledge of the volume and has contacted it successfully in the past, but the volume is currently unreachable.
	///14 - 'Aborted': Similar to Stopped, except that the volume stopped abruptly and may require configuration or maintenance.
	///15 - 'Dormant': The volume is reachable, but it is inactive.
	///16 - 'Supporting Entity in Error': This status value does not necessarily indicate trouble with the volume, but it does indicate that another device or connection that the volume depends on may need attention.
	///17 - 'Completed': The volume has completed an operation. This status value should be combined with OK, Error, or Degraded, depending on the outcome of the operation.
	///0xD00D - 'Scan Needed': In Windows-based storage subsystems, this indicates a scan is needed but not repair.
	///0xD00E - 'Spot Fix Needed': In Windows-based storage subsystems, this indicates limited repair is needed.
	///0xD00F - 'Full Repair Needed': In Windows-based storage subsystems, this indicates full repair is needed.
	OperationalStatus []Volume_OperationalStatus

	// Guid path of the volume.
	Path string

	// Indicates the ReFS deduplication mode of the volume.
	ReFSDedupMode Volume_ReFSDedupMode

	// Total size of the volume
	Size uint64

	// Available space on the volume
	SizeRemaining uint64
}

	func NewMSFT_VolumeEx1(instance *cim.WmiInstance) (newInstance *MSFT_Volume, err error) {tmp, err := NewMSFT_StorageObjectEx1(instance)
		
	if err != nil { return }
	newInstance = &MSFT_Volume {
	MSFT_StorageObject: tmp,
	}
	return
	}
	

	func NewMSFT_VolumeEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery ) (newInstance *MSFT_Volume, err error) {tmp, err := NewMSFT_StorageObjectEx6(hostName, wmiNamespace, userName, password, domainName, query)
		
	if err != nil { return }
	newInstance = &MSFT_Volume {
	MSFT_StorageObject: tmp,
	}
	return
	}
	

// SetAllocationUnitSize sets the value of AllocationUnitSize for the instance
func (instance *MSFT_Volume) SetPropertyAllocationUnitSize(value uint32) (err error) { 
    return instance.SetProperty("AllocationUnitSize", (value))
}

// GetAllocationUnitSize gets the value of AllocationUnitSize for the instance
func (instance *MSFT_Volume) GetPropertyAllocationUnitSize()(value uint32, err error) { 
    retValue, err := instance.GetProperty("AllocationUnitSize")
    if err != nil {
        return
    }
    if retValue == nil {
        // Doesn't have any value. Return empty
        return
    }
    
    valuetmp, ok := retValue.(uint32); 
    if !ok {
        err = errors.Wrapf(errors.InvalidType, " uint32 is Invalid. Expected %s", reflect.TypeOf(retValue))
        return  
    }

    value = uint32(valuetmp)

    return
}

// SetDedupMode sets the value of DedupMode for the instance
func (instance *MSFT_Volume) SetPropertyDedupMode(value Volume_DedupMode) (err error) { 
    return instance.SetProperty("DedupMode", (value))
}

// GetDedupMode gets the value of DedupMode for the instance
func (instance *MSFT_Volume) GetPropertyDedupMode()(value Volume_DedupMode, err error) { 
    retValue, err := instance.GetProperty("DedupMode")
    if err != nil {
        return
    }
    if retValue == nil {
        // Doesn't have any value. Return empty
        return
    }
    
    valuetmp, ok := retValue.(int32); 
    if !ok {
        err = errors.Wrapf(errors.InvalidType, " int32 is Invalid. Expected %s", reflect.TypeOf(retValue))
        return  
    }

    value = Volume_DedupMode(valuetmp)

    return
}

// SetDriveLetter sets the value of DriveLetter for the instance
func (instance *MSFT_Volume) SetPropertyDriveLetter(value byte) (err error) { 
    return instance.SetProperty("DriveLetter", (value))
}

// GetDriveLetter gets the value of DriveLetter for the instance
func (instance *MSFT_Volume) GetPropertyDriveLetter()(value byte, err error) { 
    retValue, err := instance.GetProperty("DriveLetter")
    if err != nil {
        return
    }
    if retValue == nil {
        // Doesn't have any value. Return empty
        return
    }
    
    valuetmp, ok := retValue.(byte); 
    if !ok {
        err = errors.Wrapf(errors.InvalidType, " byte is Invalid. Expected %s", reflect.TypeOf(retValue))
        return  
    }

    value = byte(valuetmp)

    return
}

// SetDriveType sets the value of DriveType for the instance
func (instance *MSFT_Volume) SetPropertyDriveType(value Volume_DriveType) (err error) { 
    return instance.SetProperty("DriveType", (value))
}

// GetDriveType gets the value of DriveType for the instance
func (instance *MSFT_Volume) GetPropertyDriveType()(value Volume_DriveType, err error) { 
    retValue, err := instance.GetProperty("DriveType")
    if err != nil {
        return
    }
    if retValue == nil {
        // Doesn't have any value. Return empty
        return
    }
    
    valuetmp, ok := retValue.(int32); 
    if !ok {
        err = errors.Wrapf(errors.InvalidType, " int32 is Invalid. Expected %s", reflect.TypeOf(retValue))
        return  
    }

    value = Volume_DriveType(valuetmp)

    return
}

// SetFileSystem sets the value of FileSystem for the instance
func (instance *MSFT_Volume) SetPropertyFileSystem(value string) (err error) { 
    return instance.SetProperty("FileSystem", (value))
}

// GetFileSystem gets the value of FileSystem for the instance
func (instance *MSFT_Volume) GetPropertyFileSystem()(value string, err error) { 
    retValue, err := instance.GetProperty("FileSystem")
    if err != nil {
        return
    }
    if retValue == nil {
        // Doesn't have any value. Return empty
        return
    }
    
    valuetmp, ok := retValue.(string); 
    if !ok {
        err = errors.Wrapf(errors.InvalidType, " string is Invalid. Expected %s", reflect.TypeOf(retValue))
        return  
    }

    value = string(valuetmp)

    return
}

// SetFileSystemLabel sets the value of FileSystemLabel for the instance
func (instance *MSFT_Volume) SetPropertyFileSystemLabel(value string) (err error) { 
    return instance.SetProperty("FileSystemLabel", (value))
}

// GetFileSystemLabel gets the value of FileSystemLabel for the instance
func (instance *MSFT_Volume) GetPropertyFileSystemLabel()(value string, err error) { 
    retValue, err := instance.GetProperty("FileSystemLabel")
    if err != nil {
        return
    }
    if retValue == nil {
        // Doesn't have any value. Return empty
        return
    }
    
    valuetmp, ok := retValue.(string); 
    if !ok {
        err = errors.Wrapf(errors.InvalidType, " string is Invalid. Expected %s", reflect.TypeOf(retValue))
        return  
    }

    value = string(valuetmp)

    return
}

// SetFileSystemType sets the value of FileSystemType for the instance
func (instance *MSFT_Volume) SetPropertyFileSystemType(value Volume_FileSystemType) (err error) { 
    return instance.SetProperty("FileSystemType", (value))
}

// GetFileSystemType gets the value of FileSystemType for the instance
func (instance *MSFT_Volume) GetPropertyFileSystemType()(value Volume_FileSystemType, err error) { 
    retValue, err := instance.GetProperty("FileSystemType")
    if err != nil {
        return
    }
    if retValue == nil {
        // Doesn't have any value. Return empty
        return
    }
    
    valuetmp, ok := retValue.(int32); 
    if !ok {
        err = errors.Wrapf(errors.InvalidType, " int32 is Invalid. Expected %s", reflect.TypeOf(retValue))
        return  
    }

    value = Volume_FileSystemType(valuetmp)

    return
}

// SetHealthStatus sets the value of HealthStatus for the instance
func (instance *MSFT_Volume) SetPropertyHealthStatus(value Volume_HealthStatus) (err error) { 
    return instance.SetProperty("HealthStatus", (value))
}

// GetHealthStatus gets the value of HealthStatus for the instance
func (instance *MSFT_Volume) GetPropertyHealthStatus()(value Volume_HealthStatus, err error) { 
    retValue, err := instance.GetProperty("HealthStatus")
    if err != nil {
        return
    }
    if retValue == nil {
        // Doesn't have any value. Return empty
        return
    }
    
    valuetmp, ok := retValue.(int32); 
    if !ok {
        err = errors.Wrapf(errors.InvalidType, " int32 is Invalid. Expected %s", reflect.TypeOf(retValue))
        return  
    }

    value = Volume_HealthStatus(valuetmp)

    return
}

// SetOperationalStatus sets the value of OperationalStatus for the instance
func (instance *MSFT_Volume) SetPropertyOperationalStatus(value []Volume_OperationalStatus) (err error) { 
    return instance.SetProperty("OperationalStatus", (value))
}

// GetOperationalStatus gets the value of OperationalStatus for the instance
func (instance *MSFT_Volume) GetPropertyOperationalStatus()(value []Volume_OperationalStatus, err error) { 
    retValue, err := instance.GetProperty("OperationalStatus")
    if err != nil {
        return
    }
    if retValue == nil {
        // Doesn't have any value. Return empty
        return
    }
    
    for _, interfaceValue := range retValue.([]interface{}) {
        valuetmp, ok := interfaceValue.(int32); 
        if !ok {
            err = errors.Wrapf(errors.InvalidType, " int32 is Invalid. Expected %s", reflect.TypeOf(interfaceValue))
            return  
        }
        value = append(value, Volume_OperationalStatus(valuetmp))
    }

    return
}

// SetPath sets the value of Path for the instance
func (instance *MSFT_Volume) SetPropertyPath(value string) (err error) { 
    return instance.SetProperty("Path", (value))
}

// GetPath gets the value of Path for the instance
func (instance *MSFT_Volume) GetPropertyPath()(value string, err error) { 
    retValue, err := instance.GetProperty("Path")
    if err != nil {
        return
    }
    if retValue == nil {
        // Doesn't have any value. Return empty
        return
    }
    
    valuetmp, ok := retValue.(string); 
    if !ok {
        err = errors.Wrapf(errors.InvalidType, " string is Invalid. Expected %s", reflect.TypeOf(retValue))
        return  
    }

    value = string(valuetmp)

    return
}

// SetReFSDedupMode sets the value of ReFSDedupMode for the instance
func (instance *MSFT_Volume) SetPropertyReFSDedupMode(value Volume_ReFSDedupMode) (err error) { 
    return instance.SetProperty("ReFSDedupMode", (value))
}

// GetReFSDedupMode gets the value of ReFSDedupMode for the instance
func (instance *MSFT_Volume) GetPropertyReFSDedupMode()(value Volume_ReFSDedupMode, err error) { 
    retValue, err := instance.GetProperty("ReFSDedupMode")
    if err != nil {
        return
    }
    if retValue == nil {
        // Doesn't have any value. Return empty
        return
    }
    
    valuetmp, ok := retValue.(int32); 
    if !ok {
        err = errors.Wrapf(errors.InvalidType, " int32 is Invalid. Expected %s", reflect.TypeOf(retValue))
        return  
    }

    value = Volume_ReFSDedupMode(valuetmp)

    return
}

// SetSize sets the value of Size for the instance
func (instance *MSFT_Volume) SetPropertySize(value uint64) (err error) { 
    return instance.SetProperty("Size", (value))
}

// GetSize gets the value of Size for the instance
func (instance *MSFT_Volume) GetPropertySize()(value uint64, err error) { 
    retValue, err := instance.GetProperty("Size")
    if err != nil {
        return
    }
    if retValue == nil {
        // Doesn't have any value. Return empty
        return
    }
    
    valuetmp, ok := retValue.(uint64); 
    if !ok {
        err = errors.Wrapf(errors.InvalidType, " uint64 is Invalid. Expected %s", reflect.TypeOf(retValue))
        return  
    }

    value = uint64(valuetmp)

    return
}

// SetSizeRemaining sets the value of SizeRemaining for the instance
func (instance *MSFT_Volume) SetPropertySizeRemaining(value uint64) (err error) { 
    return instance.SetProperty("SizeRemaining", (value))
}

// GetSizeRemaining gets the value of SizeRemaining for the instance
func (instance *MSFT_Volume) GetPropertySizeRemaining()(value uint64, err error) { 
    retValue, err := instance.GetProperty("SizeRemaining")
    if err != nil {
        return
    }
    if retValue == nil {
        // Doesn't have any value. Return empty
        return
    }
    
    valuetmp, ok := retValue.(uint64); 
    if !ok {
        err = errors.Wrapf(errors.InvalidType, " uint64 is Invalid. Expected %s", reflect.TypeOf(retValue))
        return  
    }

    value = uint64(valuetmp)

    return
}

// 

// <param name="CreatedStorageJob" type="MSFT_StorageJob "></param>
// <param name="ExtendedStatus" type="MSFT_StorageExtendedStatus "></param>
// <param name="ReturnValue" type="uint32 "></param>
func (instance *MSFT_Volume) DeleteObject( /* OUT */ CreatedStorageJob MSFT_StorageJob,
 /* OUT */ ExtendedStatus MSFT_StorageExtendedStatus) (result uint32, err error) {retVal, err := instance.InvokeMethod("DeleteObject" )
	if err != nil { return }
	retValue := retVal[0].(int32)
	result = uint32(retValue)
	return
	
}


// 

// <param name="AllocationUnitSize" type="uint32 "></param>
// <param name="Compress" type="bool "></param>
// <param name="DevDrive" type="bool "></param>
// <param name="DisableHeatGathering" type="bool "></param>
// <param name="FileSystem" type="string "></param>
// <param name="FileSystemLabel" type="string "></param>
// <param name="Force" type="bool "></param>
// <param name="Full" type="bool "></param>
// <param name="IsDAX" type="bool "></param>
// <param name="NoTrim" type="bool "></param>
// <param name="SetIntegrityStreams" type="bool "></param>
// <param name="SHA256Checksums" type="bool "></param>
// <param name="ShortFileNameSupport" type="bool "></param>
// <param name="UseLargeFRS" type="bool "></param>

// <param name="CreatedStorageJob" type="MSFT_StorageJob "></param>
// <param name="ExtendedStatus" type="MSFT_StorageExtendedStatus "></param>
// <param name="FormattedVolume" type="MSFT_Volume "></param>
// <param name="ReturnValue" type="uint32 "></param>
func (instance *MSFT_Volume) Format( /* IN */ FileSystem string,
 /* IN */ FileSystemLabel string,
 /* IN */ AllocationUnitSize uint32,
 /* IN */ Full bool,
 /* IN */ Force bool,
 /* IN */ Compress bool,
 /* IN */ ShortFileNameSupport bool,
 /* IN */ SetIntegrityStreams bool,
 /* IN */ UseLargeFRS bool,
 /* IN */ DisableHeatGathering bool,
 /* IN */ IsDAX bool,
 /* IN */ NoTrim bool,
 /* IN */ DevDrive bool,
 /* IN */ SHA256Checksums bool,
 /* OUT */ FormattedVolume MSFT_Volume,
 /* OUT */ CreatedStorageJob MSFT_StorageJob,
 /* OUT */ ExtendedStatus MSFT_StorageExtendedStatus) (result uint32, err error) {retVal, err := instance.InvokeMethod("Format" , FileSystem, FileSystemLabel, AllocationUnitSize, Full, Force, Compress, ShortFileNameSupport, SetIntegrityStreams, UseLargeFRS, DisableHeatGathering, IsDAX, NoTrim, DevDrive, SHA256Checksums)
	if err != nil { return }
	retValue := retVal[0].(int32)
	result = uint32(retValue)
	return
	
}


// 

// <param name="DetectLeaks" type="uint32 "></param>
// <param name="DirectoryIds" type="uint64 []"></param>
// <param name="OfflineScanAndFix" type="bool "></param>
// <param name="Salvage" type="uint32 "></param>
// <param name="Scan" type="bool "></param>
// <param name="ScratchDir" type="string "></param>
// <param name="ScratchFile" type="string "></param>
// <param name="SpotFix" type="bool "></param>
// <param name="TargetDir" type="string "></param>
// <param name="TargetFile" type="string "></param>
// <param name="Threads" type="uint32 "></param>
// <param name="Triage" type="bool "></param>

// <param name="CreatedStorageJob" type="MSFT_StorageJob "></param>
// <param name="ExtendedStatus" type="MSFT_StorageExtendedStatus "></param>
// <param name="Output" type="uint32 "></param>
// <param name="ReturnValue" type="uint32 "></param>
func (instance *MSFT_Volume) Repair( /* IN */ OfflineScanAndFix bool,
 /* IN */ Scan bool,
 /* IN */ SpotFix bool,
 /* IN */ DetectLeaks uint32,
 /* IN */ ScratchFile string,
 /* IN */ Threads uint32,
 /* IN */ Triage bool,
 /* IN */ DirectoryIds []uint64,
 /* IN */ Salvage uint32,
 /* IN */ ScratchDir string,
 /* IN */ TargetFile string,
 /* IN */ TargetDir string,
 /* OUT */ Output uint32,
 /* OUT */ CreatedStorageJob MSFT_StorageJob,
 /* OUT */ ExtendedStatus MSFT_StorageExtendedStatus) (result uint32, err error) {retVal, err := instance.InvokeMethod("Repair" , OfflineScanAndFix, Scan, SpotFix, DetectLeaks, ScratchFile, Threads, Triage, DirectoryIds, Salvage, ScratchDir, TargetFile, TargetDir)
	if err != nil { return }
	retValue := retVal[0].(int32)
	result = uint32(retValue)
	return
	
}


// 

// <param name="Analyze" type="bool "></param>
// <param name="Defrag" type="bool "></param>
// <param name="NormalPriority" type="bool "></param>
// <param name="ReTrim" type="bool "></param>
// <param name="SlabConsolidate" type="bool "></param>
// <param name="TierOptimize" type="bool "></param>

// <param name="CreatedStorageJob" type="MSFT_StorageJob "></param>
// <param name="ExtendedStatus" type="MSFT_StorageExtendedStatus "></param>
// <param name="ReturnValue" type="uint32 "></param>
func (instance *MSFT_Volume) Optimize( /* IN */ ReTrim bool,
 /* IN */ Analyze bool,
 /* IN */ Defrag bool,
 /* IN */ SlabConsolidate bool,
 /* IN */ TierOptimize bool,
 /* IN */ NormalPriority bool,
 /* OUT */ CreatedStorageJob MSFT_StorageJob,
 /* OUT */ ExtendedStatus MSFT_StorageExtendedStatus) (result uint32, err error) {retVal, err := instance.InvokeMethod("Optimize" , ReTrim, Analyze, Defrag, SlabConsolidate, TierOptimize, NormalPriority)
	if err != nil { return }
	retValue := retVal[0].(int32)
	result = uint32(retValue)
	return
	
}


// 

// <param name="FileSystemLabel" type="string "></param>

// <param name="ExtendedStatus" type="MSFT_StorageExtendedStatus "></param>
// <param name="ReturnValue" type="uint32 "></param>
func (instance *MSFT_Volume) SetFileSystemLabel( /* IN */ FileSystemLabel string,
 /* OUT */ ExtendedStatus MSFT_StorageExtendedStatus) (result uint32, err error) {retVal, err := instance.InvokeMethod("SetFileSystemLabel" , FileSystemLabel)
	if err != nil { return }
	retValue := retVal[0].(int32)
	result = uint32(retValue)
	return
	
}


// 

// <param name="ExtendedStatus" type="MSFT_StorageExtendedStatus "></param>
// <param name="ReturnValue" type="uint32 "></param>
// <param name="SupportedFileSystems" type="string []"></param>
func (instance *MSFT_Volume) GetSupportedFileSystems( /* OUT */ SupportedFileSystems []string,
 /* OUT */ ExtendedStatus MSFT_StorageExtendedStatus) (result uint32, err error) {retVal, err := instance.InvokeMethod("GetSupportedFileSystems" )
	if err != nil { return }
	retValue := retVal[0].(int32)
	result = uint32(retValue)
	return
	
}


// 

// <param name="FileSystem" type="string "></param>

// <param name="ExtendedStatus" type="MSFT_StorageExtendedStatus "></param>
// <param name="ReturnValue" type="uint32 "></param>
// <param name="SupportedClusterSizes" type="uint32 []"></param>
func (instance *MSFT_Volume) GetSupportedClusterSizes( /* IN */ FileSystem string,
 /* OUT */ SupportedClusterSizes []uint32,
 /* OUT */ ExtendedStatus MSFT_StorageExtendedStatus) (result uint32, err error) {retVal, err := instance.InvokeMethod("GetSupportedClusterSizes" , FileSystem)
	if err != nil { return }
	retValue := retVal[0].(int32)
	result = uint32(retValue)
	return
	
}


// 

// <param name="CorruptionCount" type="uint32 "></param>
// <param name="ExtendedStatus" type="MSFT_StorageExtendedStatus "></param>
// <param name="ReturnValue" type="uint32 "></param>
func (instance *MSFT_Volume) GetCorruptionCount( /* OUT */ CorruptionCount uint32,
 /* OUT */ ExtendedStatus MSFT_StorageExtendedStatus) (result uint32, err error) {retVal, err := instance.InvokeMethod("GetCorruptionCount" )
	if err != nil { return }
	retValue := retVal[0].(int32)
	result = uint32(retValue)
	return
	
}


// 

// <param name="ReturnValue" type="uint32 "></param>
// <param name="VolumeScrubEnabled" type="bool "></param>
func (instance *MSFT_Volume) GetAttributes( /* OUT */ VolumeScrubEnabled bool) (result uint32, err error) {retVal, err := instance.InvokeMethod("GetAttributes" )
	if err != nil { return }
	retValue := retVal[0].(int32)
	result = uint32(retValue)
	return
	
}


// 

// <param name="EnableVolumeScrub" type="bool "></param>

// <param name="ExtendedStatus" type="MSFT_StorageExtendedStatus "></param>
// <param name="ReturnValue" type="uint32 "></param>
func (instance *MSFT_Volume) SetAttributes( /* IN */ EnableVolumeScrub bool,
 /* OUT */ ExtendedStatus MSFT_StorageExtendedStatus) (result uint32, err error) {retVal, err := instance.InvokeMethod("SetAttributes" , EnableVolumeScrub)
	if err != nil { return }
	retValue := retVal[0].(int32)
	result = uint32(retValue)
	return
	
}


// 

// <param name="ReturnValue" type="uint32 "></param>
func (instance *MSFT_Volume) Flush() (result uint32, err error) {retVal, err := instance.InvokeMethodWithReturn("Flush" );
	if err != nil { return }
	result = uint32(retVal)
	return
	
}


// 

// <param name="Size" type="uint64 "></param>

// <param name="CreatedStorageJob" type="MSFT_StorageJob "></param>
// <param name="ExtendedStatus" type="MSFT_StorageExtendedStatus "></param>
// <param name="ReturnValue" type="uint32 "></param>
func (instance *MSFT_Volume) Resize( /* IN */ Size uint64,
 /* OUT */ CreatedStorageJob MSFT_StorageJob,
 /* OUT */ ExtendedStatus MSFT_StorageExtendedStatus) (result uint32, err error) {retVal, err := instance.InvokeMethod("Resize" , Size)
	if err != nil { return }
	retValue := retVal[0].(int32)
	result = uint32(retValue)
	return
	
}


// 

// <param name="DiagnoseResults" type="MSFT_StorageDiagnoseResult []"></param>
// <param name="ExtendedStatus" type="MSFT_StorageExtendedStatus "></param>
// <param name="ReturnValue" type="uint32 "></param>
func (instance *MSFT_Volume) Diagnose( /* OUT */ DiagnoseResults []MSFT_StorageDiagnoseResult,
 /* OUT */ ExtendedStatus MSFT_StorageExtendedStatus) (result uint32, err error) {retVal, err := instance.InvokeMethod("Diagnose" )
	if err != nil { return }
	retValue := retVal[0].(int32)
	result = uint32(retValue)
	return
	
}


// 

// <param name="DedupMode" type="uint32 "></param>

// <param name="ExtendedStatus" type="MSFT_StorageExtendedStatus "></param>
// <param name="ReturnValue" type="uint32 "></param>
func (instance *MSFT_Volume) SetDedupMode( /* IN */ DedupMode uint32,
 /* OUT */ ExtendedStatus MSFT_StorageExtendedStatus) (result uint32, err error) {retVal, err := instance.InvokeMethod("SetDedupMode" , DedupMode)
	if err != nil { return }
	retValue := retVal[0].(int32)
	result = uint32(retValue)
	return
	
}


// 

// <param name="DedupProperties" type="MSFT_DedupProperties "></param>
// <param name="ExtendedStatus" type="MSFT_StorageExtendedStatus "></param>
// <param name="ReturnValue" type="uint32 "></param>
func (instance *MSFT_Volume) GetDedupProperties( /* OUT */ DedupProperties MSFT_DedupProperties,
 /* OUT */ ExtendedStatus MSFT_StorageExtendedStatus) (result uint32, err error) {retVal, err := instance.InvokeMethod("GetDedupProperties" )
	if err != nil { return }
	retValue := retVal[0].(int32)
	result = uint32(retValue)
	return
	
}


// 

// <param name="ActionResults" type="MSFT_HealthAction []"></param>
// <param name="ExtendedStatus" type="MSFT_StorageExtendedStatus "></param>
// <param name="ReturnValue" type="uint32 "></param>
func (instance *MSFT_Volume) GetActions( /* OUT */ ActionResults []MSFT_HealthAction,
 /* OUT */ ExtendedStatus MSFT_StorageExtendedStatus) (result uint32, err error) {retVal, err := instance.InvokeMethod("GetActions" )
	if err != nil { return }
	retValue := retVal[0].(int32)
	result = uint32(retValue)
	return
	
}


