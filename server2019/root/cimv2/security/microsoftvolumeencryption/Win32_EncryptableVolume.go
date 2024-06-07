// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

// 
// Author:
//      Auto Generated on 6/6/2024 using wmigen
//      Source root.CIMV2.Security.MicrosoftVolumeEncryption
//////////////////////////////////////////////
package microsoftvolumeencryption
import (
 "github.com/microsoft/wmi/pkg/base/query"
cim "github.com/microsoft/wmi/pkg/wmiinstance"
 "github.com/microsoft/wmi/pkg/base/instance"
 "github.com/microsoft/wmi/pkg/errors"
 "reflect"
)

// Win32_EncryptableVolume struct
type Win32_EncryptableVolume struct { 
	*cim.WmiInstance

	// 
	ConversionStatus uint32

	// 
	DeviceID string

	// 
	DriveLetter string

	// 
	EncryptionMethod uint32

	// 
	IsVolumeInitializedForProtection bool

	// 
	PersistentVolumeID string

	// 
	ProtectionStatus uint32

	// 
	VolumeType EncryptableVolume_VolumeType
}

	func NewWin32_EncryptableVolumeEx1(instance *cim.WmiInstance) (newInstance *Win32_EncryptableVolume, err error) {tmp, err := instance, nil
		
	if err != nil { return }
	newInstance = &Win32_EncryptableVolume {
	WmiInstance: tmp,
	}
	return
	}
	

	func NewWin32_EncryptableVolumeEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery ) (newInstance *Win32_EncryptableVolume, err error) {tmp, err := instance.GetWmiInstance(hostName, wmiNamespace, userName, password, domainName, query)
		
	if err != nil { return }
	newInstance = &Win32_EncryptableVolume {
	WmiInstance: tmp,
	}
	return
	}
	

// SetConversionStatus sets the value of ConversionStatus for the instance
func (instance *Win32_EncryptableVolume) SetPropertyConversionStatus(value uint32) (err error) { 
    return instance.SetProperty("ConversionStatus", (value))
}

// GetConversionStatus gets the value of ConversionStatus for the instance
func (instance *Win32_EncryptableVolume) GetPropertyConversionStatus()(value uint32, err error) { 
    retValue, err := instance.GetProperty("ConversionStatus")
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

// SetDeviceID sets the value of DeviceID for the instance
func (instance *Win32_EncryptableVolume) SetPropertyDeviceID(value string) (err error) { 
    return instance.SetProperty("DeviceID", (value))
}

// GetDeviceID gets the value of DeviceID for the instance
func (instance *Win32_EncryptableVolume) GetPropertyDeviceID()(value string, err error) { 
    retValue, err := instance.GetProperty("DeviceID")
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

// SetDriveLetter sets the value of DriveLetter for the instance
func (instance *Win32_EncryptableVolume) SetPropertyDriveLetter(value string) (err error) { 
    return instance.SetProperty("DriveLetter", (value))
}

// GetDriveLetter gets the value of DriveLetter for the instance
func (instance *Win32_EncryptableVolume) GetPropertyDriveLetter()(value string, err error) { 
    retValue, err := instance.GetProperty("DriveLetter")
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

// SetEncryptionMethod sets the value of EncryptionMethod for the instance
func (instance *Win32_EncryptableVolume) SetPropertyEncryptionMethod(value uint32) (err error) { 
    return instance.SetProperty("EncryptionMethod", (value))
}

// GetEncryptionMethod gets the value of EncryptionMethod for the instance
func (instance *Win32_EncryptableVolume) GetPropertyEncryptionMethod()(value uint32, err error) { 
    retValue, err := instance.GetProperty("EncryptionMethod")
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

// SetIsVolumeInitializedForProtection sets the value of IsVolumeInitializedForProtection for the instance
func (instance *Win32_EncryptableVolume) SetPropertyIsVolumeInitializedForProtection(value bool) (err error) { 
    return instance.SetProperty("IsVolumeInitializedForProtection", (value))
}

// GetIsVolumeInitializedForProtection gets the value of IsVolumeInitializedForProtection for the instance
func (instance *Win32_EncryptableVolume) GetPropertyIsVolumeInitializedForProtection()(value bool, err error) { 
    retValue, err := instance.GetProperty("IsVolumeInitializedForProtection")
    if err != nil {
        return
    }
    if retValue == nil {
        // Doesn't have any value. Return empty
        return
    }
    
    valuetmp, ok := retValue.(bool); 
    if !ok {
        err = errors.Wrapf(errors.InvalidType, " bool is Invalid. Expected %s", reflect.TypeOf(retValue))
        return  
    }

    value = bool(valuetmp)

    return
}

// SetPersistentVolumeID sets the value of PersistentVolumeID for the instance
func (instance *Win32_EncryptableVolume) SetPropertyPersistentVolumeID(value string) (err error) { 
    return instance.SetProperty("PersistentVolumeID", (value))
}

// GetPersistentVolumeID gets the value of PersistentVolumeID for the instance
func (instance *Win32_EncryptableVolume) GetPropertyPersistentVolumeID()(value string, err error) { 
    retValue, err := instance.GetProperty("PersistentVolumeID")
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

// SetProtectionStatus sets the value of ProtectionStatus for the instance
func (instance *Win32_EncryptableVolume) SetPropertyProtectionStatus(value uint32) (err error) { 
    return instance.SetProperty("ProtectionStatus", (value))
}

// GetProtectionStatus gets the value of ProtectionStatus for the instance
func (instance *Win32_EncryptableVolume) GetPropertyProtectionStatus()(value uint32, err error) { 
    retValue, err := instance.GetProperty("ProtectionStatus")
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

// SetVolumeType sets the value of VolumeType for the instance
func (instance *Win32_EncryptableVolume) SetPropertyVolumeType(value EncryptableVolume_VolumeType) (err error) { 
    return instance.SetProperty("VolumeType", (value))
}

// GetVolumeType gets the value of VolumeType for the instance
func (instance *Win32_EncryptableVolume) GetPropertyVolumeType()(value EncryptableVolume_VolumeType, err error) { 
    retValue, err := instance.GetProperty("VolumeType")
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

    value = EncryptableVolume_VolumeType(valuetmp)

    return
}

// 

// <param name="PrecisionFactor" type="EncryptableVolume_PrecisionFactor "></param>

// <param name="ConversionStatus" type="EncryptableVolume_ConversionStatus "></param>
// <param name="EncryptionFlags" type="uint32 "></param>
// <param name="EncryptionPercentage" type="uint32 "></param>
// <param name="ReturnValue" type="uint32 "></param>
// <param name="WipingPercentage" type="uint32 "></param>
// <param name="WipingStatus" type="EncryptableVolume_WipingStatus "></param>
func (instance *Win32_EncryptableVolume) GetConversionStatus( /* OUT */ ConversionStatus EncryptableVolume_ConversionStatus,
 /* OUT */ EncryptionPercentage uint32,
 /* OUT */ EncryptionFlags uint32,
 /* OUT */ WipingStatus EncryptableVolume_WipingStatus,
 /* OUT */ WipingPercentage uint32,
 /* OPTIONAL IN */ PrecisionFactor EncryptableVolume_PrecisionFactor) (result uint32, err error) {retVal, err := instance.InvokeMethod("GetConversionStatus" , PrecisionFactor)
	if err != nil { return }
	retValue := retVal[0].(int32)
	result = uint32(retValue)
	return
	
}


// 

// <param name="ReturnValue" type="uint32 "></param>
// <param name="TestError" type="uint32 "></param>
// <param name="TestStatus" type="EncryptableVolume_TestStatus "></param>
func (instance *Win32_EncryptableVolume) GetHardwareTestStatus( /* OUT */ TestStatus EncryptableVolume_TestStatus,
 /* OUT */ TestError uint32) (result uint32, err error) {retVal, err := instance.InvokeMethod("GetHardwareTestStatus" )
	if err != nil { return }
	retValue := retVal[0].(int32)
	result = uint32(retValue)
	return
	
}


// 

// <param name="LockStatus" type="EncryptableVolume_LockStatus "></param>
// <param name="ReturnValue" type="uint32 "></param>
func (instance *Win32_EncryptableVolume) GetLockStatus( /* OUT */ LockStatus EncryptableVolume_LockStatus) (result uint32, err error) {retVal, err := instance.InvokeMethod("GetLockStatus" )
	if err != nil { return }
	retValue := retVal[0].(int32)
	result = uint32(retValue)
	return
	
}


// 

// <param name="ProtectionStatus" type="EncryptableVolume_ProtectionStatus "></param>
// <param name="ReturnValue" type="uint32 "></param>
func (instance *Win32_EncryptableVolume) GetProtectionStatus( /* OUT */ ProtectionStatus EncryptableVolume_ProtectionStatus) (result uint32, err error) {retVal, err := instance.InvokeMethod("GetProtectionStatus" )
	if err != nil { return }
	retValue := retVal[0].(int32)
	result = uint32(retValue)
	return
	
}


// 

// <param name="ReturnValue" type="uint32 "></param>
// <param name="SuspendCount" type="uint32 "></param>
func (instance *Win32_EncryptableVolume) GetSuspendCount( /* OUT */ SuspendCount uint32) (result uint32, err error) {retVal, err := instance.InvokeMethod("GetSuspendCount" )
	if err != nil { return }
	retValue := retVal[0].(int32)
	result = uint32(retValue)
	return
	
}


// 

// <param name="HardwareEncryptionStatus" type="EncryptableVolume_HardwareEncryptionStatus "></param>
// <param name="ReturnValue" type="uint32 "></param>
func (instance *Win32_EncryptableVolume) GetHardwareEncryptionStatus( /* OUT */ HardwareEncryptionStatus EncryptableVolume_HardwareEncryptionStatus) (result uint32, err error) {retVal, err := instance.InvokeMethod("GetHardwareEncryptionStatus" )
	if err != nil { return }
	retValue := retVal[0].(int32)
	result = uint32(retValue)
	return
	
}


// 

// <param name="EncryptionFlags" type="uint32 "></param>
// <param name="EncryptionMethod" type="EncryptableVolume_EncryptionMethod "></param>

// <param name="ReturnValue" type="uint32 "></param>
func (instance *Win32_EncryptableVolume) Encrypt( /* IN */ EncryptionMethod EncryptableVolume_EncryptionMethod,
 /* IN */ EncryptionFlags uint32) (result uint32, err error) {retVal, err := instance.InvokeMethodWithReturn("Encrypt" , EncryptionMethod, EncryptionFlags);
	if err != nil { return }
	result = uint32(retVal)
	return
	
}


// 

// <param name="EncryptionFlags" type="uint32 "></param>
// <param name="EncryptionMethod" type="EncryptableVolume_EncryptionMethod "></param>

// <param name="ReturnValue" type="uint32 "></param>
func (instance *Win32_EncryptableVolume) EncryptAfterHardwareTest( /* IN */ EncryptionMethod EncryptableVolume_EncryptionMethod,
 /* IN */ EncryptionFlags uint32) (result uint32, err error) {retVal, err := instance.InvokeMethodWithReturn("EncryptAfterHardwareTest" , EncryptionMethod, EncryptionFlags);
	if err != nil { return }
	result = uint32(retVal)
	return
	
}


// 

// <param name="EncryptionMethod" type="EncryptableVolume_EncryptionMethod "></param>
// <param name="EncryptionMethodFlags" type="uint32 "></param>
// <param name="ReturnValue" type="uint32 "></param>
// <param name="SelfEncryptionDriveEncryptionMethod" type="string "></param>
func (instance *Win32_EncryptableVolume) GetEncryptionMethod( /* OUT */ EncryptionMethod EncryptableVolume_EncryptionMethod,
 /* OUT */ SelfEncryptionDriveEncryptionMethod string,
 /* OUT */ EncryptionMethodFlags uint32) (result uint32, err error) {retVal, err := instance.InvokeMethod("GetEncryptionMethod" )
	if err != nil { return }
	retValue := retVal[0].(int32)
	result = uint32(retValue)
	return
	
}


// 

// <param name="ReturnValue" type="uint32 "></param>
func (instance *Win32_EncryptableVolume) Decrypt() (result uint32, err error) {retVal, err := instance.InvokeMethodWithReturn("Decrypt" );
	if err != nil { return }
	result = uint32(retVal)
	return
	
}


// 

// <param name="ReturnValue" type="uint32 "></param>
func (instance *Win32_EncryptableVolume) PauseConversion() (result uint32, err error) {retVal, err := instance.InvokeMethodWithReturn("PauseConversion" );
	if err != nil { return }
	result = uint32(retVal)
	return
	
}


// 

// <param name="ReturnValue" type="uint32 "></param>
func (instance *Win32_EncryptableVolume) ResumeConversion() (result uint32, err error) {retVal, err := instance.InvokeMethodWithReturn("ResumeConversion" );
	if err != nil { return }
	result = uint32(retVal)
	return
	
}


// 

// <param name="FriendlyName" type="string "></param>
// <param name="PlatformValidationProfile" type="uint8 []"></param>

// <param name="ReturnValue" type="uint32 "></param>
// <param name="VolumeKeyProtectorID" type="string "></param>
func (instance *Win32_EncryptableVolume) ProtectKeyWithTPM( /* IN */ FriendlyName string,
 /* IN */ PlatformValidationProfile []uint8,
 /* OUT */ VolumeKeyProtectorID string) (result uint32, err error) {retVal, err := instance.InvokeMethod("ProtectKeyWithTPM" , FriendlyName, PlatformValidationProfile)
	if err != nil { return }
	retValue := retVal[0].(int32)
	result = uint32(retValue)
	return
	
}


// 

// <param name="ExternalKey" type="uint8 []"></param>
// <param name="FriendlyName" type="string "></param>

// <param name="ReturnValue" type="uint32 "></param>
// <param name="VolumeKeyProtectorID" type="string "></param>
func (instance *Win32_EncryptableVolume) ProtectKeyWithExternalKey( /* IN */ FriendlyName string,
 /* IN */ ExternalKey []uint8,
 /* OUT */ VolumeKeyProtectorID string) (result uint32, err error) {retVal, err := instance.InvokeMethod("ProtectKeyWithExternalKey" , FriendlyName, ExternalKey)
	if err != nil { return }
	retValue := retVal[0].(int32)
	result = uint32(retValue)
	return
	
}


// 

// <param name="FriendlyName" type="string "></param>
// <param name="NumericalPassword" type="string "></param>

// <param name="ReturnValue" type="uint32 "></param>
// <param name="VolumeKeyProtectorID" type="string "></param>
func (instance *Win32_EncryptableVolume) ProtectKeyWithNumericalPassword( /* IN */ FriendlyName string,
 /* IN */ NumericalPassword string,
 /* OUT */ VolumeKeyProtectorID string) (result uint32, err error) {retVal, err := instance.InvokeMethod("ProtectKeyWithNumericalPassword" , FriendlyName, NumericalPassword)
	if err != nil { return }
	retValue := retVal[0].(int32)
	result = uint32(retValue)
	return
	
}


// 

// <param name="FriendlyName" type="string "></param>
// <param name="PIN" type="string "></param>
// <param name="PlatformValidationProfile" type="uint8 []"></param>

// <param name="ReturnValue" type="uint32 "></param>
// <param name="VolumeKeyProtectorID" type="string "></param>
func (instance *Win32_EncryptableVolume) ProtectKeyWithTPMAndPIN( /* IN */ FriendlyName string,
 /* IN */ PlatformValidationProfile []uint8,
 /* IN */ PIN string,
 /* OUT */ VolumeKeyProtectorID string) (result uint32, err error) {retVal, err := instance.InvokeMethod("ProtectKeyWithTPMAndPIN" , FriendlyName, PlatformValidationProfile, PIN)
	if err != nil { return }
	retValue := retVal[0].(int32)
	result = uint32(retValue)
	return
	
}


// 

// <param name="ExternalKey" type="uint8 []"></param>
// <param name="FriendlyName" type="string "></param>
// <param name="PlatformValidationProfile" type="uint8 []"></param>

// <param name="ReturnValue" type="uint32 "></param>
// <param name="VolumeKeyProtectorID" type="string "></param>
func (instance *Win32_EncryptableVolume) ProtectKeyWithTPMAndStartupKey( /* IN */ FriendlyName string,
 /* IN */ PlatformValidationProfile []uint8,
 /* IN */ ExternalKey []uint8,
 /* OUT */ VolumeKeyProtectorID string) (result uint32, err error) {retVal, err := instance.InvokeMethod("ProtectKeyWithTPMAndStartupKey" , FriendlyName, PlatformValidationProfile, ExternalKey)
	if err != nil { return }
	retValue := retVal[0].(int32)
	result = uint32(retValue)
	return
	
}


// 

// <param name="ExternalKey" type="uint8 []"></param>
// <param name="FriendlyName" type="string "></param>
// <param name="PIN" type="string "></param>
// <param name="PlatformValidationProfile" type="uint8 []"></param>

// <param name="ReturnValue" type="uint32 "></param>
// <param name="VolumeKeyProtectorID" type="string "></param>
func (instance *Win32_EncryptableVolume) ProtectKeyWithTPMAndPINAndStartupKey( /* IN */ FriendlyName string,
 /* IN */ PlatformValidationProfile []uint8,
 /* IN */ PIN string,
 /* IN */ ExternalKey []uint8,
 /* OUT */ VolumeKeyProtectorID string) (result uint32, err error) {retVal, err := instance.InvokeMethod("ProtectKeyWithTPMAndPINAndStartupKey" , FriendlyName, PlatformValidationProfile, PIN, ExternalKey)
	if err != nil { return }
	retValue := retVal[0].(int32)
	result = uint32(retValue)
	return
	
}


// 

// <param name="FriendlyName" type="string "></param>
// <param name="PathWithFileName" type="string "></param>

// <param name="ReturnValue" type="uint32 "></param>
// <param name="VolumeKeyProtectorID" type="string "></param>
func (instance *Win32_EncryptableVolume) ProtectKeyWithCertificateFile( /* IN */ FriendlyName string,
 /* IN */ PathWithFileName string,
 /* OUT */ VolumeKeyProtectorID string) (result uint32, err error) {retVal, err := instance.InvokeMethod("ProtectKeyWithCertificateFile" , FriendlyName, PathWithFileName)
	if err != nil { return }
	retValue := retVal[0].(int32)
	result = uint32(retValue)
	return
	
}


// 

// <param name="CertThumbprint" type="string "></param>
// <param name="FriendlyName" type="string "></param>

// <param name="ReturnValue" type="uint32 "></param>
// <param name="VolumeKeyProtectorID" type="string "></param>
func (instance *Win32_EncryptableVolume) ProtectKeyWithCertificateThumbprint( /* IN */ FriendlyName string,
 /* IN */ CertThumbprint string,
 /* OUT */ VolumeKeyProtectorID string) (result uint32, err error) {retVal, err := instance.InvokeMethod("ProtectKeyWithCertificateThumbprint" , FriendlyName, CertThumbprint)
	if err != nil { return }
	retValue := retVal[0].(int32)
	result = uint32(retValue)
	return
	
}


// 

// <param name="FriendlyName" type="string "></param>
// <param name="PassPhrase" type="string "></param>

// <param name="ReturnValue" type="uint32 "></param>
// <param name="VolumeKeyProtectorID" type="string "></param>
func (instance *Win32_EncryptableVolume) ProtectKeyWithPassPhrase( /* IN */ FriendlyName string,
 /* IN */ PassPhrase string,
 /* OUT */ VolumeKeyProtectorID string) (result uint32, err error) {retVal, err := instance.InvokeMethod("ProtectKeyWithPassPhrase" , FriendlyName, PassPhrase)
	if err != nil { return }
	retValue := retVal[0].(int32)
	result = uint32(retValue)
	return
	
}


// 

// <param name="Flags" type="uint32 "></param>
// <param name="FriendlyName" type="string "></param>
// <param name="SidString" type="string "></param>

// <param name="ReturnValue" type="uint32 "></param>
// <param name="VolumeKeyProtectorID" type="string "></param>
func (instance *Win32_EncryptableVolume) ProtectKeyWithAdSid( /* IN */ FriendlyName string,
 /* IN */ SidString string,
 /* IN */ Flags uint32,
 /* OUT */ VolumeKeyProtectorID string) (result uint32, err error) {retVal, err := instance.InvokeMethod("ProtectKeyWithAdSid" , FriendlyName, SidString, Flags)
	if err != nil { return }
	retValue := retVal[0].(int32)
	result = uint32(retValue)
	return
	
}


// 

// <param name="VolumeKeyProtectorID" type="string "></param>

// <param name="ReturnValue" type="uint32 "></param>
func (instance *Win32_EncryptableVolume) EnableAutoUnlock( /* IN */ VolumeKeyProtectorID string) (result uint32, err error) {retVal, err := instance.InvokeMethodWithReturn("EnableAutoUnlock" , VolumeKeyProtectorID);
	if err != nil { return }
	result = uint32(retVal)
	return
	
}


// 

// <param name="ReturnValue" type="uint32 "></param>
func (instance *Win32_EncryptableVolume) DisableAutoUnlock() (result uint32, err error) {retVal, err := instance.InvokeMethodWithReturn("DisableAutoUnlock" );
	if err != nil { return }
	result = uint32(retVal)
	return
	
}


// 

// <param name="IsAutoUnlockEnabled" type="bool "></param>
// <param name="ReturnValue" type="uint32 "></param>
// <param name="VolumeKeyProtectorID" type="string "></param>
func (instance *Win32_EncryptableVolume) IsAutoUnlockEnabled( /* OUT */ IsAutoUnlockEnabled bool,
 /* OUT */ VolumeKeyProtectorID string) (result uint32, err error) {retVal, err := instance.InvokeMethod("IsAutoUnlockEnabled" )
	if err != nil { return }
	retValue := retVal[0].(int32)
	result = uint32(retValue)
	return
	
}


// 

// <param name="ReturnValue" type="uint32 "></param>
func (instance *Win32_EncryptableVolume) ClearAllAutoUnlockKeys() (result uint32, err error) {retVal, err := instance.InvokeMethodWithReturn("ClearAllAutoUnlockKeys" );
	if err != nil { return }
	result = uint32(retVal)
	return
	
}


// 

// <param name="IsAutoUnlockKeyStored" type="bool "></param>
// <param name="ReturnValue" type="uint32 "></param>
func (instance *Win32_EncryptableVolume) IsAutoUnlockKeyStored( /* OUT */ IsAutoUnlockKeyStored bool) (result uint32, err error) {retVal, err := instance.InvokeMethod("IsAutoUnlockKeyStored" )
	if err != nil { return }
	retValue := retVal[0].(int32)
	result = uint32(retValue)
	return
	
}


// 

// <param name="Path" type="string "></param>
// <param name="VolumeKeyProtectorID" type="string "></param>

// <param name="ReturnValue" type="uint32 "></param>
func (instance *Win32_EncryptableVolume) SaveExternalKeyToFile( /* IN */ VolumeKeyProtectorID string,
 /* IN */ Path string) (result uint32, err error) {retVal, err := instance.InvokeMethodWithReturn("SaveExternalKeyToFile" , VolumeKeyProtectorID, Path);
	if err != nil { return }
	result = uint32(retVal)
	return
	
}


// 

// <param name="VolumeKeyProtectorID" type="string "></param>

// <param name="FileName" type="string "></param>
// <param name="ReturnValue" type="uint32 "></param>
func (instance *Win32_EncryptableVolume) GetExternalKeyFileName( /* IN */ VolumeKeyProtectorID string,
 /* OUT */ FileName string) (result uint32, err error) {retVal, err := instance.InvokeMethod("GetExternalKeyFileName" , VolumeKeyProtectorID)
	if err != nil { return }
	retValue := retVal[0].(int32)
	result = uint32(retValue)
	return
	
}


// 

// <param name="PathWithFileName" type="string "></param>

// <param name="ExternalKey" type="uint8 []"></param>
// <param name="ReturnValue" type="uint32 "></param>
func (instance *Win32_EncryptableVolume) GetExternalKeyFromFile( /* IN */ PathWithFileName string,
 /* OUT */ ExternalKey []uint8) (result uint32, err error) {retVal, err := instance.InvokeMethod("GetExternalKeyFromFile" , PathWithFileName)
	if err != nil { return }
	retValue := retVal[0].(int32)
	result = uint32(retValue)
	return
	
}


// 

// <param name="KeyProtectorType" type="EncryptableVolume_KeyProtectorType "></param>

// <param name="ReturnValue" type="uint32 "></param>
// <param name="VolumeKeyProtectorID" type="string []"></param>
func (instance *Win32_EncryptableVolume) GetKeyProtectors( /* IN */ KeyProtectorType EncryptableVolume_KeyProtectorType,
 /* OUT */ VolumeKeyProtectorID []string) (result uint32, err error) {retVal, err := instance.InvokeMethod("GetKeyProtectors" , KeyProtectorType)
	if err != nil { return }
	retValue := retVal[0].(int32)
	result = uint32(retValue)
	return
	
}


// 

// <param name="DisableCount" type="uint32 "></param>

// <param name="ReturnValue" type="uint32 "></param>
func (instance *Win32_EncryptableVolume) DisableKeyProtectors( /* IN */ DisableCount uint32) (result uint32, err error) {retVal, err := instance.InvokeMethodWithReturn("DisableKeyProtectors" , DisableCount);
	if err != nil { return }
	result = uint32(retVal)
	return
	
}


// 

// <param name="ReturnValue" type="uint32 "></param>
func (instance *Win32_EncryptableVolume) EnableKeyProtectors() (result uint32, err error) {retVal, err := instance.InvokeMethodWithReturn("EnableKeyProtectors" );
	if err != nil { return }
	result = uint32(retVal)
	return
	
}


// 

// <param name="VolumeKeyProtectorID" type="string "></param>

// <param name="ReturnValue" type="uint32 "></param>
func (instance *Win32_EncryptableVolume) DeleteKeyProtector( /* IN */ VolumeKeyProtectorID string) (result uint32, err error) {retVal, err := instance.InvokeMethodWithReturn("DeleteKeyProtector" , VolumeKeyProtectorID);
	if err != nil { return }
	result = uint32(retVal)
	return
	
}


// 

// <param name="ReturnValue" type="uint32 "></param>
func (instance *Win32_EncryptableVolume) DeleteKeyProtectors() (result uint32, err error) {retVal, err := instance.InvokeMethodWithReturn("DeleteKeyProtectors" );
	if err != nil { return }
	result = uint32(retVal)
	return
	
}


// 

// <param name="NumericalPassword" type="string "></param>

// <param name="ReturnValue" type="uint32 "></param>
func (instance *Win32_EncryptableVolume) UnlockWithNumericalPassword( /* IN */ NumericalPassword string) (result uint32, err error) {retVal, err := instance.InvokeMethodWithReturn("UnlockWithNumericalPassword" , NumericalPassword);
	if err != nil { return }
	result = uint32(retVal)
	return
	
}


// 

// <param name="ExternalKey" type="uint8 []"></param>

// <param name="ReturnValue" type="uint32 "></param>
func (instance *Win32_EncryptableVolume) UnlockWithExternalKey( /* IN */ ExternalKey []uint8) (result uint32, err error) {retVal, err := instance.InvokeMethodWithReturn("UnlockWithExternalKey" , ExternalKey);
	if err != nil { return }
	result = uint32(retVal)
	return
	
}


// 

// <param name="PathWithFileName" type="string "></param>
// <param name="Pin" type="string "></param>

// <param name="ReturnValue" type="uint32 "></param>
func (instance *Win32_EncryptableVolume) UnlockWithCertificateFile( /* IN */ PathWithFileName string,
 /* IN */ Pin string) (result uint32, err error) {retVal, err := instance.InvokeMethodWithReturn("UnlockWithCertificateFile" , PathWithFileName, Pin);
	if err != nil { return }
	result = uint32(retVal)
	return
	
}


// 

// <param name="CertThumbprint" type="string "></param>
// <param name="Pin" type="string "></param>

// <param name="ReturnValue" type="uint32 "></param>
func (instance *Win32_EncryptableVolume) UnlockWithCertificateThumbprint( /* IN */ CertThumbprint string,
 /* IN */ Pin string) (result uint32, err error) {retVal, err := instance.InvokeMethodWithReturn("UnlockWithCertificateThumbprint" , CertThumbprint, Pin);
	if err != nil { return }
	result = uint32(retVal)
	return
	
}


// 

// <param name="PassPhrase" type="string "></param>

// <param name="ReturnValue" type="uint32 "></param>
func (instance *Win32_EncryptableVolume) UnlockWithPassPhrase( /* IN */ PassPhrase string) (result uint32, err error) {retVal, err := instance.InvokeMethodWithReturn("UnlockWithPassPhrase" , PassPhrase);
	if err != nil { return }
	result = uint32(retVal)
	return
	
}


// 

// <param name="SidString" type="string "></param>

// <param name="ReturnValue" type="uint32 "></param>
func (instance *Win32_EncryptableVolume) UnlockWithAdSid( /* OPTIONAL IN */ SidString string) (result uint32, err error) {retVal, err := instance.InvokeMethodWithReturn("UnlockWithAdSid" , SidString);
	if err != nil { return }
	result = uint32(retVal)
	return
	
}


// 

// <param name="ForceDismount" type="bool "></param>

// <param name="ReturnValue" type="uint32 "></param>
func (instance *Win32_EncryptableVolume) Lock( /* IN */ ForceDismount bool) (result uint32, err error) {retVal, err := instance.InvokeMethodWithReturn("Lock" , ForceDismount);
	if err != nil { return }
	result = uint32(retVal)
	return
	
}


// 

// <param name="NumericalPassword" type="string "></param>

// <param name="IsNumericalPasswordValid" type="bool "></param>
// <param name="ReturnValue" type="uint32 "></param>
func (instance *Win32_EncryptableVolume) IsNumericalPasswordValid( /* IN */ NumericalPassword string,
 /* OUT */ IsNumericalPasswordValid bool) (result uint32, err error) {retVal, err := instance.InvokeMethod("IsNumericalPasswordValid" , NumericalPassword)
	if err != nil { return }
	retValue := retVal[0].(int32)
	result = uint32(retValue)
	return
	
}


// 

// <param name="KeyProtectorType" type="EncryptableVolume_KeyProtectorType "></param>

// <param name="IsKeyProtectorAvailable" type="bool "></param>
// <param name="ReturnValue" type="uint32 "></param>
func (instance *Win32_EncryptableVolume) IsKeyProtectorAvailable( /* IN */ KeyProtectorType EncryptableVolume_KeyProtectorType,
 /* OUT */ IsKeyProtectorAvailable bool) (result uint32, err error) {retVal, err := instance.InvokeMethod("IsKeyProtectorAvailable" , KeyProtectorType)
	if err != nil { return }
	retValue := retVal[0].(int32)
	result = uint32(retValue)
	return
	
}


// 

// <param name="VolumeKeyProtectorID" type="string "></param>

// <param name="KeyProtectorType" type="EncryptableVolume_KeyProtectorType "></param>
// <param name="ReturnValue" type="uint32 "></param>
func (instance *Win32_EncryptableVolume) GetKeyProtectorType( /* IN */ VolumeKeyProtectorID string,
 /* OUT */ KeyProtectorType EncryptableVolume_KeyProtectorType) (result uint32, err error) {retVal, err := instance.InvokeMethod("GetKeyProtectorType" , VolumeKeyProtectorID)
	if err != nil { return }
	retValue := retVal[0].(int32)
	result = uint32(retValue)
	return
	
}


// 

// <param name="VolumeKeyProtectorID" type="string "></param>

// <param name="FriendlyName" type="string "></param>
// <param name="ReturnValue" type="uint32 "></param>
func (instance *Win32_EncryptableVolume) GetKeyProtectorFriendlyName( /* IN */ VolumeKeyProtectorID string,
 /* OUT */ FriendlyName string) (result uint32, err error) {retVal, err := instance.InvokeMethod("GetKeyProtectorFriendlyName" , VolumeKeyProtectorID)
	if err != nil { return }
	retValue := retVal[0].(int32)
	result = uint32(retValue)
	return
	
}


// 

// <param name="VolumeKeyProtectorID" type="string "></param>

// <param name="ExternalKey" type="uint8 []"></param>
// <param name="ReturnValue" type="uint32 "></param>
func (instance *Win32_EncryptableVolume) GetKeyProtectorExternalKey( /* IN */ VolumeKeyProtectorID string,
 /* OUT */ ExternalKey []uint8) (result uint32, err error) {retVal, err := instance.InvokeMethod("GetKeyProtectorExternalKey" , VolumeKeyProtectorID)
	if err != nil { return }
	retValue := retVal[0].(int32)
	result = uint32(retValue)
	return
	
}


// 

// <param name="VolumeKeyProtectorID" type="string "></param>

// <param name="NumericalPassword" type="string "></param>
// <param name="ReturnValue" type="uint32 "></param>
func (instance *Win32_EncryptableVolume) GetKeyProtectorNumericalPassword( /* IN */ VolumeKeyProtectorID string,
 /* OUT */ NumericalPassword string) (result uint32, err error) {retVal, err := instance.InvokeMethod("GetKeyProtectorNumericalPassword" , VolumeKeyProtectorID)
	if err != nil { return }
	retValue := retVal[0].(int32)
	result = uint32(retValue)
	return
	
}


// 

// <param name="VolumeKeyProtectorID" type="string "></param>

// <param name="PlatformValidationProfile" type="uint8 []"></param>
// <param name="ReturnValue" type="uint32 "></param>
func (instance *Win32_EncryptableVolume) GetKeyProtectorPlatformValidationProfile( /* IN */ VolumeKeyProtectorID string,
 /* OUT */ PlatformValidationProfile []uint8) (result uint32, err error) {retVal, err := instance.InvokeMethod("GetKeyProtectorPlatformValidationProfile" , VolumeKeyProtectorID)
	if err != nil { return }
	retValue := retVal[0].(int32)
	result = uint32(retValue)
	return
	
}


// 

// <param name="VolumeKeyProtectorID" type="string "></param>

// <param name="CertThumbprint" type="string "></param>
// <param name="CertType" type="uint32 "></param>
// <param name="PublicKey" type="uint8 []"></param>
// <param name="ReturnValue" type="uint32 "></param>
func (instance *Win32_EncryptableVolume) GetKeyProtectorCertificate( /* IN */ VolumeKeyProtectorID string,
 /* OUT */ PublicKey []uint8,
 /* OUT */ CertThumbprint string,
 /* OUT */ CertType uint32) (result uint32, err error) {retVal, err := instance.InvokeMethod("GetKeyProtectorCertificate" , VolumeKeyProtectorID)
	if err != nil { return }
	retValue := retVal[0].(int32)
	result = uint32(retValue)
	return
	
}


// 

// <param name="VolumeKeyProtectorID" type="string "></param>

// <param name="Flags" type="uint32 "></param>
// <param name="ReturnValue" type="uint32 "></param>
// <param name="SidString" type="string "></param>
func (instance *Win32_EncryptableVolume) GetKeyProtectorAdSidInformation( /* IN */ VolumeKeyProtectorID string,
 /* OUT */ SidString string,
 /* OUT */ Flags uint32) (result uint32, err error) {retVal, err := instance.InvokeMethod("GetKeyProtectorAdSidInformation" , VolumeKeyProtectorID)
	if err != nil { return }
	retValue := retVal[0].(int32)
	result = uint32(retValue)
	return
	
}


// 

// <param name="VolumeKeyProtectorID" type="string "></param>

// <param name="KeyPackage" type="uint8 []"></param>
// <param name="ReturnValue" type="uint32 "></param>
func (instance *Win32_EncryptableVolume) GetKeyPackage( /* IN */ VolumeKeyProtectorID string,
 /* OUT */ KeyPackage []uint8) (result uint32, err error) {retVal, err := instance.InvokeMethod("GetKeyPackage" , VolumeKeyProtectorID)
	if err != nil { return }
	retValue := retVal[0].(int32)
	result = uint32(retValue)
	return
	
}


// 

// <param name="VolumeKeyProtectorID" type="string "></param>

// <param name="ReturnValue" type="uint32 "></param>
func (instance *Win32_EncryptableVolume) BackupRecoveryInformationToActiveDirectory( /* IN */ VolumeKeyProtectorID string) (result uint32, err error) {retVal, err := instance.InvokeMethodWithReturn("BackupRecoveryInformationToActiveDirectory" , VolumeKeyProtectorID);
	if err != nil { return }
	result = uint32(retVal)
	return
	
}


// 

// <param name="NewPassPhrase" type="string "></param>
// <param name="VolumeKeyProtectorID" type="string "></param>

// <param name="NewProtectorID" type="string "></param>
// <param name="ReturnValue" type="uint32 "></param>
func (instance *Win32_EncryptableVolume) ChangePassPhrase( /* IN */ VolumeKeyProtectorID string,
 /* IN */ NewPassPhrase string,
 /* OUT */ NewProtectorID string) (result uint32, err error) {retVal, err := instance.InvokeMethod("ChangePassPhrase" , VolumeKeyProtectorID, NewPassPhrase)
	if err != nil { return }
	retValue := retVal[0].(int32)
	result = uint32(retValue)
	return
	
}


// 

// <param name="NewPIN" type="string "></param>
// <param name="VolumeKeyProtectorID" type="string "></param>

// <param name="ReturnValue" type="uint32 "></param>
func (instance *Win32_EncryptableVolume) ChangePIN( /* IN */ VolumeKeyProtectorID string,
 /* IN */ NewPIN string) (result uint32, err error) {retVal, err := instance.InvokeMethodWithReturn("ChangePIN" , VolumeKeyProtectorID, NewPIN);
	if err != nil { return }
	result = uint32(retVal)
	return
	
}


// 

// <param name="ExternalKey" type="uint8 []"></param>
// <param name="VolumeKeyProtectorID" type="string "></param>

// <param name="NewVolumeKeyProtectorID" type="string "></param>
// <param name="ReturnValue" type="uint32 "></param>
func (instance *Win32_EncryptableVolume) ChangeExternalKey( /* IN */ VolumeKeyProtectorID string,
 /* IN */ ExternalKey []uint8,
 /* OUT */ NewVolumeKeyProtectorID string) (result uint32, err error) {retVal, err := instance.InvokeMethod("ChangeExternalKey" , VolumeKeyProtectorID, ExternalKey)
	if err != nil { return }
	retValue := retVal[0].(int32)
	result = uint32(retValue)
	return
	
}


// 

// <param name="DiscoveryVolumeType" type="string "></param>
// <param name="ForceEncryptionType" type="EncryptableVolume_ForceEncryptionType "></param>

// <param name="ReturnValue" type="uint32 "></param>
func (instance *Win32_EncryptableVolume) PrepareVolume( /* IN */ DiscoveryVolumeType string,
 /* IN */ ForceEncryptionType EncryptableVolume_ForceEncryptionType) (result uint32, err error) {retVal, err := instance.InvokeMethodWithReturn("PrepareVolume" , DiscoveryVolumeType, ForceEncryptionType);
	if err != nil { return }
	result = uint32(retVal)
	return
	
}


// 

// <param name="DiscoveryVolumeType" type="string "></param>
// <param name="InitializationFlags" type="uint32 "></param>

// <param name="ReturnValue" type="uint32 "></param>
func (instance *Win32_EncryptableVolume) PrepareVolumeEx( /* IN */ DiscoveryVolumeType string,
 /* IN */ InitializationFlags uint32) (result uint32, err error) {retVal, err := instance.InvokeMethodWithReturn("PrepareVolumeEx" , DiscoveryVolumeType, InitializationFlags);
	if err != nil { return }
	result = uint32(retVal)
	return
	
}


// 

// <param name="IdentificationField" type="string "></param>
// <param name="ReturnValue" type="uint32 "></param>
func (instance *Win32_EncryptableVolume) GetIdentificationField( /* OUT */ IdentificationField string) (result uint32, err error) {retVal, err := instance.InvokeMethod("GetIdentificationField" )
	if err != nil { return }
	retValue := retVal[0].(int32)
	result = uint32(retValue)
	return
	
}


// 

// <param name="IdentificationField" type="string "></param>

// <param name="ReturnValue" type="uint32 "></param>
func (instance *Win32_EncryptableVolume) SetIdentificationField( /* IN */ IdentificationField string) (result uint32, err error) {retVal, err := instance.InvokeMethodWithReturn("SetIdentificationField" , IdentificationField);
	if err != nil { return }
	result = uint32(retVal)
	return
	
}


// 

// <param name="ReturnValue" type="uint32 "></param>
// <param name="Version" type="EncryptableVolume_Version "></param>
func (instance *Win32_EncryptableVolume) GetVersion( /* OUT */ Version EncryptableVolume_Version) (result uint32, err error) {retVal, err := instance.InvokeMethod("GetVersion" )
	if err != nil { return }
	retValue := retVal[0].(int32)
	result = uint32(retValue)
	return
	
}


// 

// <param name="ReturnValue" type="uint32 "></param>
func (instance *Win32_EncryptableVolume) UpgradeVolume() (result uint32, err error) {retVal, err := instance.InvokeMethodWithReturn("UpgradeVolume" );
	if err != nil { return }
	result = uint32(retVal)
	return
	
}


// 

// <param name="CertThumbprint" type="string []"></param>
// <param name="ReturnValue" type="uint32 "></param>
func (instance *Win32_EncryptableVolume) FindValidCertificates( /* OUT */ CertThumbprint []string) (result uint32, err error) {retVal, err := instance.InvokeMethod("FindValidCertificates" )
	if err != nil { return }
	retValue := retVal[0].(int32)
	result = uint32(retValue)
	return
	
}


// 

// <param name="BindingState" type="EncryptableVolume_BindingState "></param>
// <param name="ReturnValue" type="uint32 "></param>
func (instance *Win32_EncryptableVolume) GetSecureBootBindingState( /* OUT */ BindingState EncryptableVolume_BindingState) (result uint32, err error) {retVal, err := instance.InvokeMethod("GetSecureBootBindingState" )
	if err != nil { return }
	retValue := retVal[0].(int32)
	result = uint32(retValue)
	return
	
}


// 

// <param name="CertThumbprint" type="string "></param>
// <param name="FriendlyName" type="string "></param>

// <param name="ReturnValue" type="uint32 "></param>
// <param name="VolumeKeyProtectorID" type="string "></param>
func (instance *Win32_EncryptableVolume) ProtectKeyWithNetworkCertificate( /* IN */ FriendlyName string,
 /* IN */ CertThumbprint string,
 /* OUT */ VolumeKeyProtectorID string) (result uint32, err error) {retVal, err := instance.InvokeMethod("ProtectKeyWithNetworkCertificate" , FriendlyName, CertThumbprint)
	if err != nil { return }
	retValue := retVal[0].(int32)
	result = uint32(retValue)
	return
	
}


// 

// <param name="VolumeKeyProtectorID" type="string "></param>

// <param name="ReturnValue" type="uint32 "></param>
func (instance *Win32_EncryptableVolume) BackupRecoveryInformationToCloudDomain( /* IN */ VolumeKeyProtectorID string) (result uint32, err error) {retVal, err := instance.InvokeMethodWithReturn("BackupRecoveryInformationToCloudDomain" , VolumeKeyProtectorID);
	if err != nil { return }
	result = uint32(retVal)
	return
	
}


// 

// <param name="LocalIPAddress" type="string "></param>
// <param name="ServerIPAddresses" type="string []"></param>

// <param name="ReturnValue" type="uint32 "></param>
func (instance *Win32_EncryptableVolume) UnlockWithNetworkServerKey( /* IN */ ServerIPAddresses []string,
 /* IN */ LocalIPAddress string) (result uint32, err error) {retVal, err := instance.InvokeMethodWithReturn("UnlockWithNetworkServerKey" , ServerIPAddresses, LocalIPAddress);
	if err != nil { return }
	result = uint32(retVal)
	return
	
}


// 

// <param name="BackupInfoType" type="uint32 "></param>
// <param name="VolumeKeyProtectorID" type="string "></param>

// <param name="BackupAccounts" type="string "></param>
// <param name="ReturnValue" type="uint32 "></param>
func (instance *Win32_EncryptableVolume) GetNumericalPasswordBackupAccounts( /* IN */ VolumeKeyProtectorID string,
 /* IN */ BackupInfoType uint32,
 /* OUT */ BackupAccounts string) (result uint32, err error) {retVal, err := instance.InvokeMethod("GetNumericalPasswordBackupAccounts" , VolumeKeyProtectorID, BackupInfoType)
	if err != nil { return }
	retValue := retVal[0].(int32)
	result = uint32(retValue)
	return
	
}


// 

// <param name="VolumeKeyProtectorID" type="string "></param>

// <param name="BackupInfoType" type="uint32 "></param>
// <param name="ReturnValue" type="uint32 "></param>
func (instance *Win32_EncryptableVolume) GetNumericalPasswordBackupType( /* IN */ VolumeKeyProtectorID string,
 /* OUT */ BackupInfoType uint32) (result uint32, err error) {retVal, err := instance.InvokeMethod("GetNumericalPasswordBackupType" , VolumeKeyProtectorID)
	if err != nil { return }
	retValue := retVal[0].(int32)
	result = uint32(retValue)
	return
	
}


