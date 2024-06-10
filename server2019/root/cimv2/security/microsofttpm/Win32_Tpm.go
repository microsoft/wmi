// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

// Author:
//
//	Auto Generated on 6/6/2024 using wmigen
//	Source root.CIMV2.Security.MicrosoftTpm
//
// ////////////////////////////////////////////
package microsofttpm

import (
	"github.com/microsoft/wmi/pkg/base/instance"
	"github.com/microsoft/wmi/pkg/base/query"
	"github.com/microsoft/wmi/pkg/errors"
	cim "github.com/microsoft/wmi/pkg/wmiinstance"
	"reflect"
)

// Win32_Tpm struct
type Win32_Tpm struct {
	*cim.WmiInstance

	//
	IsActivated_InitialValue bool

	//
	IsEnabled_InitialValue bool

	//
	IsOwned_InitialValue bool

	//
	ManufacturerId uint32

	//
	ManufacturerIdTxt string

	//
	ManufacturerVersion string

	//
	ManufacturerVersionFull20 string

	//
	ManufacturerVersionInfo string

	//
	PhysicalPresenceVersionInfo string

	//
	SpecVersion string
}

func NewWin32_TpmEx1(instance *cim.WmiInstance) (newInstance *Win32_Tpm, err error) {
	tmp, err := instance, nil

	if err != nil {
		return
	}
	newInstance = &Win32_Tpm{
		WmiInstance: tmp,
	}
	return
}

func NewWin32_TpmEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *Win32_Tpm, err error) {
	tmp, err := instance.GetWmiInstance(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &Win32_Tpm{
		WmiInstance: tmp,
	}
	return
}

// SetIsActivated_InitialValue sets the value of IsActivated_InitialValue for the instance
func (instance *Win32_Tpm) SetPropertyIsActivated_InitialValue(value bool) (err error) {
	return instance.SetProperty("IsActivated_InitialValue", (value))
}

// GetIsActivated_InitialValue gets the value of IsActivated_InitialValue for the instance
func (instance *Win32_Tpm) GetPropertyIsActivated_InitialValue() (value bool, err error) {
	retValue, err := instance.GetProperty("IsActivated_InitialValue")
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

// SetIsEnabled_InitialValue sets the value of IsEnabled_InitialValue for the instance
func (instance *Win32_Tpm) SetPropertyIsEnabled_InitialValue(value bool) (err error) {
	return instance.SetProperty("IsEnabled_InitialValue", (value))
}

// GetIsEnabled_InitialValue gets the value of IsEnabled_InitialValue for the instance
func (instance *Win32_Tpm) GetPropertyIsEnabled_InitialValue() (value bool, err error) {
	retValue, err := instance.GetProperty("IsEnabled_InitialValue")
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

// SetIsOwned_InitialValue sets the value of IsOwned_InitialValue for the instance
func (instance *Win32_Tpm) SetPropertyIsOwned_InitialValue(value bool) (err error) {
	return instance.SetProperty("IsOwned_InitialValue", (value))
}

// GetIsOwned_InitialValue gets the value of IsOwned_InitialValue for the instance
func (instance *Win32_Tpm) GetPropertyIsOwned_InitialValue() (value bool, err error) {
	retValue, err := instance.GetProperty("IsOwned_InitialValue")
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

// SetManufacturerId sets the value of ManufacturerId for the instance
func (instance *Win32_Tpm) SetPropertyManufacturerId(value uint32) (err error) {
	return instance.SetProperty("ManufacturerId", (value))
}

// GetManufacturerId gets the value of ManufacturerId for the instance
func (instance *Win32_Tpm) GetPropertyManufacturerId() (value uint32, err error) {
	retValue, err := instance.GetProperty("ManufacturerId")
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

// SetManufacturerIdTxt sets the value of ManufacturerIdTxt for the instance
func (instance *Win32_Tpm) SetPropertyManufacturerIdTxt(value string) (err error) {
	return instance.SetProperty("ManufacturerIdTxt", (value))
}

// GetManufacturerIdTxt gets the value of ManufacturerIdTxt for the instance
func (instance *Win32_Tpm) GetPropertyManufacturerIdTxt() (value string, err error) {
	retValue, err := instance.GetProperty("ManufacturerIdTxt")
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

// SetManufacturerVersion sets the value of ManufacturerVersion for the instance
func (instance *Win32_Tpm) SetPropertyManufacturerVersion(value string) (err error) {
	return instance.SetProperty("ManufacturerVersion", (value))
}

// GetManufacturerVersion gets the value of ManufacturerVersion for the instance
func (instance *Win32_Tpm) GetPropertyManufacturerVersion() (value string, err error) {
	retValue, err := instance.GetProperty("ManufacturerVersion")
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

// SetManufacturerVersionFull20 sets the value of ManufacturerVersionFull20 for the instance
func (instance *Win32_Tpm) SetPropertyManufacturerVersionFull20(value string) (err error) {
	return instance.SetProperty("ManufacturerVersionFull20", (value))
}

// GetManufacturerVersionFull20 gets the value of ManufacturerVersionFull20 for the instance
func (instance *Win32_Tpm) GetPropertyManufacturerVersionFull20() (value string, err error) {
	retValue, err := instance.GetProperty("ManufacturerVersionFull20")
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

// SetManufacturerVersionInfo sets the value of ManufacturerVersionInfo for the instance
func (instance *Win32_Tpm) SetPropertyManufacturerVersionInfo(value string) (err error) {
	return instance.SetProperty("ManufacturerVersionInfo", (value))
}

// GetManufacturerVersionInfo gets the value of ManufacturerVersionInfo for the instance
func (instance *Win32_Tpm) GetPropertyManufacturerVersionInfo() (value string, err error) {
	retValue, err := instance.GetProperty("ManufacturerVersionInfo")
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

// SetPhysicalPresenceVersionInfo sets the value of PhysicalPresenceVersionInfo for the instance
func (instance *Win32_Tpm) SetPropertyPhysicalPresenceVersionInfo(value string) (err error) {
	return instance.SetProperty("PhysicalPresenceVersionInfo", (value))
}

// GetPhysicalPresenceVersionInfo gets the value of PhysicalPresenceVersionInfo for the instance
func (instance *Win32_Tpm) GetPropertyPhysicalPresenceVersionInfo() (value string, err error) {
	retValue, err := instance.GetProperty("PhysicalPresenceVersionInfo")
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

// SetSpecVersion sets the value of SpecVersion for the instance
func (instance *Win32_Tpm) SetPropertySpecVersion(value string) (err error) {
	return instance.SetProperty("SpecVersion", (value))
}

// GetSpecVersion gets the value of SpecVersion for the instance
func (instance *Win32_Tpm) GetPropertySpecVersion() (value string, err error) {
	retValue, err := instance.GetProperty("SpecVersion")
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

// <param name="IsEnabled" type="bool "></param>
// <param name="ReturnValue" type="uint32 "></param>
func (instance *Win32_Tpm) IsEnabled( /* OUT */ IsEnabled bool) (result uint32, err error) {
	retVal, err := instance.InvokeMethod("IsEnabled")
	if err != nil {
		return
	}
	retValue := retVal[0].(int32)
	result = uint32(retValue)
	return

}

//

// <param name="IsOwned" type="bool "></param>
// <param name="ReturnValue" type="uint32 "></param>
func (instance *Win32_Tpm) IsOwned( /* OUT */ IsOwned bool) (result uint32, err error) {
	retVal, err := instance.InvokeMethod("IsOwned")
	if err != nil {
		return
	}
	retValue := retVal[0].(int32)
	result = uint32(retValue)
	return

}

//

// <param name="IsActivated" type="bool "></param>
// <param name="ReturnValue" type="uint32 "></param>
func (instance *Win32_Tpm) IsActivated( /* OUT */ IsActivated bool) (result uint32, err error) {
	retVal, err := instance.InvokeMethod("IsActivated")
	if err != nil {
		return
	}
	retValue := retVal[0].(int32)
	result = uint32(retValue)
	return

}

//

// <param name="IsPhysicalClearDisabled" type="bool "></param>
// <param name="ReturnValue" type="uint32 "></param>
func (instance *Win32_Tpm) IsPhysicalClearDisabled( /* OUT */ IsPhysicalClearDisabled bool) (result uint32, err error) {
	retVal, err := instance.InvokeMethod("IsPhysicalClearDisabled")
	if err != nil {
		return
	}
	retValue := retVal[0].(int32)
	result = uint32(retValue)
	return

}

//

// <param name="IsOwnerClearDisabled" type="bool "></param>
// <param name="ReturnValue" type="uint32 "></param>
func (instance *Win32_Tpm) IsOwnerClearDisabled( /* OUT */ IsOwnerClearDisabled bool) (result uint32, err error) {
	retVal, err := instance.InvokeMethod("IsOwnerClearDisabled")
	if err != nil {
		return
	}
	retValue := retVal[0].(int32)
	result = uint32(retValue)
	return

}

//

// <param name="IsPhysicalPresenceHardwareEnabled" type="bool "></param>
// <param name="ReturnValue" type="uint32 "></param>
func (instance *Win32_Tpm) IsPhysicalPresenceHardwareEnabled( /* OUT */ IsPhysicalPresenceHardwareEnabled bool) (result uint32, err error) {
	retVal, err := instance.InvokeMethod("IsPhysicalPresenceHardwareEnabled")
	if err != nil {
		return
	}
	retValue := retVal[0].(int32)
	result = uint32(retValue)
	return

}

//

// <param name="IsOwnershipAllowed" type="bool "></param>
// <param name="ReturnValue" type="uint32 "></param>
func (instance *Win32_Tpm) IsOwnershipAllowed( /* OUT */ IsOwnershipAllowed bool) (result uint32, err error) {
	retVal, err := instance.InvokeMethod("IsOwnershipAllowed")
	if err != nil {
		return
	}
	retValue := retVal[0].(int32)
	result = uint32(retValue)
	return

}

//

// <param name="CommandOrdinal" type="uint32 "></param>

// <param name="IsCommandPresent" type="bool "></param>
// <param name="ReturnValue" type="uint32 "></param>
func (instance *Win32_Tpm) IsCommandPresent( /* IN */ CommandOrdinal uint32,
	/* OUT */ IsCommandPresent bool) (result uint32, err error) {
	retVal, err := instance.InvokeMethod("IsCommandPresent", CommandOrdinal)
	if err != nil {
		return
	}
	retValue := retVal[0].(int32)
	result = uint32(retValue)
	return

}

//

// <param name="OwnerAuth" type="string "></param>

// <param name="ReturnValue" type="uint32 "></param>
func (instance *Win32_Tpm) Enable( /* OPTIONAL IN */ OwnerAuth string) (result uint32, err error) {
	retVal, err := instance.InvokeMethodWithReturn("Enable", OwnerAuth)
	if err != nil {
		return
	}
	result = uint32(retVal)
	return

}

//

// <param name="OwnerAuth" type="string "></param>

// <param name="ReturnValue" type="uint32 "></param>
func (instance *Win32_Tpm) Disable( /* OPTIONAL IN */ OwnerAuth string) (result uint32, err error) {
	retVal, err := instance.InvokeMethodWithReturn("Disable", OwnerAuth)
	if err != nil {
		return
	}
	result = uint32(retVal)
	return

}

//

// <param name="IsEndorsementKeyPairPresent" type="bool "></param>
// <param name="ReturnValue" type="uint32 "></param>
func (instance *Win32_Tpm) IsEndorsementKeyPairPresent( /* OUT */ IsEndorsementKeyPairPresent bool) (result uint32, err error) {
	retVal, err := instance.InvokeMethod("IsEndorsementKeyPairPresent")
	if err != nil {
		return
	}
	retValue := retVal[0].(int32)
	result = uint32(retValue)
	return

}

//

// <param name="ReturnValue" type="uint32 "></param>
func (instance *Win32_Tpm) CreateEndorsementKeyPair() (result uint32, err error) {
	retVal, err := instance.InvokeMethodWithReturn("CreateEndorsementKeyPair")
	if err != nil {
		return
	}
	result = uint32(retVal)
	return

}

//

// <param name="OwnerAuth" type="string "></param>

// <param name="ReturnValue" type="uint32 "></param>
func (instance *Win32_Tpm) TakeOwnership( /* OPTIONAL IN */ OwnerAuth string) (result uint32, err error) {
	retVal, err := instance.InvokeMethodWithReturn("TakeOwnership", OwnerAuth)
	if err != nil {
		return
	}
	result = uint32(retVal)
	return

}

//

// <param name="OwnerAuth" type="string "></param>

// <param name="ReturnValue" type="uint32 "></param>
func (instance *Win32_Tpm) Clear( /* OPTIONAL IN */ OwnerAuth string) (result uint32, err error) {
	retVal, err := instance.InvokeMethodWithReturn("Clear", OwnerAuth)
	if err != nil {
		return
	}
	result = uint32(retVal)
	return

}

//

// <param name="IsSrkAuthCompatible" type="bool "></param>
// <param name="ReturnValue" type="uint32 "></param>
func (instance *Win32_Tpm) IsSrkAuthCompatible( /* OUT */ IsSrkAuthCompatible bool) (result uint32, err error) {
	retVal, err := instance.InvokeMethod("IsSrkAuthCompatible")
	if err != nil {
		return
	}
	retValue := retVal[0].(int32)
	result = uint32(retValue)
	return

}

//

// <param name="OwnerAuth" type="string "></param>

// <param name="ReturnValue" type="uint32 "></param>
func (instance *Win32_Tpm) ResetSrkAuth( /* OPTIONAL IN */ OwnerAuth string) (result uint32, err error) {
	retVal, err := instance.InvokeMethodWithReturn("ResetSrkAuth", OwnerAuth)
	if err != nil {
		return
	}
	result = uint32(retVal)
	return

}

//

// <param name="NewOwnerAuth" type="string "></param>
// <param name="OldOwnerAuth" type="string "></param>

// <param name="ReturnValue" type="uint32 "></param>
func (instance *Win32_Tpm) ChangeOwnerAuth( /* OPTIONAL IN */ OldOwnerAuth string,
	/* OPTIONAL IN */ NewOwnerAuth string) (result uint32, err error) {
	retVal, err := instance.InvokeMethodWithReturn("ChangeOwnerAuth", OldOwnerAuth, NewOwnerAuth)
	if err != nil {
		return
	}
	result = uint32(retVal)
	return

}

//

// <param name="ReturnValue" type="uint32 "></param>
// <param name="SelfTestResult" type="uint8 []"></param>
func (instance *Win32_Tpm) SelfTest( /* OUT */ SelfTestResult []uint8) (result uint32, err error) {
	retVal, err := instance.InvokeMethod("SelfTest")
	if err != nil {
		return
	}
	retValue := retVal[0].(int32)
	result = uint32(retValue)
	return

}

//

// <param name="OwnerPassPhrase" type="string "></param>

// <param name="OwnerAuth" type="string "></param>
// <param name="ReturnValue" type="uint32 "></param>
func (instance *Win32_Tpm) ConvertToOwnerAuth( /* IN */ OwnerPassPhrase string,
	/* OUT */ OwnerAuth string) (result uint32, err error) {
	retVal, err := instance.InvokeMethod("ConvertToOwnerAuth", OwnerPassPhrase)
	if err != nil {
		return
	}
	retValue := retVal[0].(int32)
	result = uint32(retValue)
	return

}

//

// <param name="Request" type="uint32 "></param>
// <param name="RequestParameter" type="uint32 "></param>

// <param name="ReturnValue" type="uint32 "></param>
func (instance *Win32_Tpm) SetPhysicalPresenceRequest( /* IN */ Request uint32,
	/* OPTIONAL IN */ RequestParameter uint32) (result uint32, err error) {
	retVal, err := instance.InvokeMethodWithReturn("SetPhysicalPresenceRequest", Request, RequestParameter)
	if err != nil {
		return
	}
	result = uint32(retVal)
	return

}

//

// <param name="Request" type="uint32 "></param>
// <param name="ReturnValue" type="uint32 "></param>
func (instance *Win32_Tpm) GetPhysicalPresenceRequest( /* OUT */ Request uint32) (result uint32, err error) {
	retVal, err := instance.InvokeMethod("GetPhysicalPresenceRequest")
	if err != nil {
		return
	}
	retValue := retVal[0].(int32)
	result = uint32(retValue)
	return

}

//

// <param name="ReturnValue" type="uint32 "></param>
// <param name="Transition" type="uint32 "></param>
func (instance *Win32_Tpm) GetPhysicalPresenceTransition( /* OUT */ Transition uint32) (result uint32, err error) {
	retVal, err := instance.InvokeMethod("GetPhysicalPresenceTransition")
	if err != nil {
		return
	}
	retValue := retVal[0].(int32)
	result = uint32(retValue)
	return

}

//

// <param name="Request" type="uint32 "></param>
// <param name="Response" type="uint32 "></param>
// <param name="ReturnValue" type="uint32 "></param>
func (instance *Win32_Tpm) GetPhysicalPresenceResponse( /* OUT */ Request uint32,
	/* OUT */ Response uint32) (result uint32, err error) {
	retVal, err := instance.InvokeMethod("GetPhysicalPresenceResponse")
	if err != nil {
		return
	}
	retValue := retVal[0].(int32)
	result = uint32(retValue)
	return

}

//

// <param name="CommandOrdinal" type="uint32 "></param>

// <param name="ReturnValue" type="uint32 "></param>
func (instance *Win32_Tpm) AddBlockedCommand( /* IN */ CommandOrdinal uint32) (result uint32, err error) {
	retVal, err := instance.InvokeMethodWithReturn("AddBlockedCommand", CommandOrdinal)
	if err != nil {
		return
	}
	result = uint32(retVal)
	return

}

//

// <param name="CommandOrdinal" type="uint32 "></param>

// <param name="ReturnValue" type="uint32 "></param>
func (instance *Win32_Tpm) RemoveBlockedCommand( /* IN */ CommandOrdinal uint32) (result uint32, err error) {
	retVal, err := instance.InvokeMethodWithReturn("RemoveBlockedCommand", CommandOrdinal)
	if err != nil {
		return
	}
	result = uint32(retVal)
	return

}

//

// <param name="CommandOrdinal" type="uint32 "></param>

// <param name="IsCommandBlocked" type="uint32 "></param>
// <param name="ReturnValue" type="uint32 "></param>
func (instance *Win32_Tpm) IsCommandBlocked( /* IN */ CommandOrdinal uint32,
	/* OUT */ IsCommandBlocked uint32) (result uint32, err error) {
	retVal, err := instance.InvokeMethod("IsCommandBlocked", CommandOrdinal)
	if err != nil {
		return
	}
	retValue := retVal[0].(int32)
	result = uint32(retValue)
	return

}

//

// <param name="OwnerAuth" type="string "></param>

// <param name="ReturnValue" type="uint32 "></param>
func (instance *Win32_Tpm) ResetAuthLockOut( /* OPTIONAL IN */ OwnerAuth string) (result uint32, err error) {
	retVal, err := instance.InvokeMethodWithReturn("ResetAuthLockOut", OwnerAuth)
	if err != nil {
		return
	}
	result = uint32(retVal)
	return

}

//

// <param name="IsReady" type="bool "></param>
// <param name="ReturnValue" type="uint32 "></param>
func (instance *Win32_Tpm) IsReady( /* OUT */ IsReady bool) (result uint32, err error) {
	retVal, err := instance.InvokeMethod("IsReady")
	if err != nil {
		return
	}
	retValue := retVal[0].(int32)
	result = uint32(retValue)
	return

}

//

// <param name="Information" type="uint32 "></param>
// <param name="IsReady" type="bool "></param>
// <param name="ReturnValue" type="uint32 "></param>
func (instance *Win32_Tpm) IsReadyInformation( /* OUT */ IsReady bool,
	/* OUT */ Information uint32) (result uint32, err error) {
	retVal, err := instance.InvokeMethod("IsReadyInformation")
	if err != nil {
		return
	}
	retValue := retVal[0].(int32)
	result = uint32(retValue)
	return

}

//

// <param name="IsAutoProvisioningEnabled" type="bool "></param>
// <param name="ReturnValue" type="uint32 "></param>
func (instance *Win32_Tpm) IsAutoProvisioningEnabled( /* OUT */ IsAutoProvisioningEnabled bool) (result uint32, err error) {
	retVal, err := instance.InvokeMethod("IsAutoProvisioningEnabled")
	if err != nil {
		return
	}
	retValue := retVal[0].(int32)
	result = uint32(retValue)
	return

}

//

// <param name="ReturnValue" type="uint32 "></param>
func (instance *Win32_Tpm) EnableAutoProvisioning() (result uint32, err error) {
	retVal, err := instance.InvokeMethodWithReturn("EnableAutoProvisioning")
	if err != nil {
		return
	}
	result = uint32(retVal)
	return

}

//

// <param name="OnlyForNextBoot" type="bool "></param>

// <param name="ReturnValue" type="uint32 "></param>
func (instance *Win32_Tpm) DisableAutoProvisioning( /* OPTIONAL IN */ OnlyForNextBoot bool) (result uint32, err error) {
	retVal, err := instance.InvokeMethodWithReturn("DisableAutoProvisioning", OnlyForNextBoot)
	if err != nil {
		return
	}
	result = uint32(retVal)
	return

}

//

// <param name="OwnerAuth" type="string "></param>
// <param name="ReturnValue" type="uint32 "></param>
func (instance *Win32_Tpm) GetOwnerAuth( /* OUT */ OwnerAuth string) (result uint32, err error) {
	retVal, err := instance.InvokeMethod("GetOwnerAuth")
	if err != nil {
		return
	}
	retValue := retVal[0].(int32)
	result = uint32(retValue)
	return

}

//

// <param name="ForceClear_Allowed" type="bool "></param>
// <param name="PhysicalPresencePrompts_Allowed" type="bool "></param>

// <param name="Information" type="uint32 "></param>
// <param name="ReturnValue" type="uint32 "></param>
func (instance *Win32_Tpm) Provision( /* IN */ ForceClear_Allowed bool,
	/* IN */ PhysicalPresencePrompts_Allowed bool,
	/* OUT */ Information uint32) (result uint32, err error) {
	retVal, err := instance.InvokeMethod("Provision", ForceClear_Allowed, PhysicalPresencePrompts_Allowed)
	if err != nil {
		return
	}
	retValue := retVal[0].(int32)
	result = uint32(retValue)
	return

}

//

// <param name="OwnerAuth" type="string "></param>

// <param name="ReturnValue" type="uint32 "></param>
func (instance *Win32_Tpm) ImportOwnerAuth( /* IN */ OwnerAuth string) (result uint32, err error) {
	retVal, err := instance.InvokeMethodWithReturn("ImportOwnerAuth", OwnerAuth)
	if err != nil {
		return
	}
	result = uint32(retVal)
	return

}

//

// <param name="Operation" type="uint32 "></param>

// <param name="ConfirmationStatus" type="uint32 "></param>
// <param name="ReturnValue" type="uint32 "></param>
func (instance *Win32_Tpm) GetPhysicalPresenceConfirmationStatus( /* IN */ Operation uint32,
	/* OUT */ ConfirmationStatus uint32) (result uint32, err error) {
	retVal, err := instance.InvokeMethod("GetPhysicalPresenceConfirmationStatus", Operation)
	if err != nil {
		return
	}
	retValue := retVal[0].(int32)
	result = uint32(retValue)
	return

}

//

// <param name="ReturnValue" type="uint32 "></param>
// <param name="SrkPublicKeyModulus" type="uint8 []"></param>
func (instance *Win32_Tpm) GetSrkPublicKeyModulus( /* OUT */ SrkPublicKeyModulus []uint8) (result uint32, err error) {
	retVal, err := instance.InvokeMethod("GetSrkPublicKeyModulus")
	if err != nil {
		return
	}
	retValue := retVal[0].(int32)
	result = uint32(retValue)
	return

}

//

// <param name="SrkPublicKeyModulus" type="uint8 []"></param>

// <param name="ReturnValue" type="uint32 "></param>
// <param name="SrkADThumbprint" type="uint8 []"></param>
func (instance *Win32_Tpm) GetSrkADThumbprint( /* IN */ SrkPublicKeyModulus []uint8,
	/* OUT */ SrkADThumbprint []uint8) (result uint32, err error) {
	retVal, err := instance.InvokeMethod("GetSrkADThumbprint", SrkPublicKeyModulus)
	if err != nil {
		return
	}
	retValue := retVal[0].(int32)
	result = uint32(retValue)
	return

}

//

// <param name="ReturnValue" type="uint32 "></param>
// <param name="TcgLog" type="uint8 []"></param>
func (instance *Win32_Tpm) GetTcgLog( /* OUT */ TcgLog []uint8) (result uint32, err error) {
	retVal, err := instance.InvokeMethod("GetTcgLog")
	if err != nil {
		return
	}
	retValue := retVal[0].(int32)
	result = uint32(retValue)
	return

}

//

// <param name="ReturnValue" type="uint32 "></param>
// <param name="TestResult" type="uint32 "></param>
func (instance *Win32_Tpm) IsKeyAttestationCapable( /* OUT */ TestResult uint32) (result uint32, err error) {
	retVal, err := instance.InvokeMethod("IsKeyAttestationCapable")
	if err != nil {
		return
	}
	retValue := retVal[0].(int32)
	result = uint32(retValue)
	return

}

//

// <param name="OwnerAuth" type="string "></param>
// <param name="OwnerAuthStatus" type="uint32 "></param>
// <param name="ReturnValue" type="uint32 "></param>
func (instance *Win32_Tpm) GetOwnerAuthForEscrow( /* OUT */ OwnerAuthStatus uint32,
	/* OUT */ OwnerAuth string) (result uint32, err error) {
	retVal, err := instance.InvokeMethod("GetOwnerAuthForEscrow")
	if err != nil {
		return
	}
	retValue := retVal[0].(int32)
	result = uint32(retValue)
	return

}

//

// <param name="OwnerAuth" type="string "></param>

// <param name="ReturnValue" type="uint32 "></param>
func (instance *Win32_Tpm) OwnerAuthEscrowed( /* IN */ OwnerAuth string) (result uint32, err error) {
	retVal, err := instance.InvokeMethodWithReturn("OwnerAuthEscrowed", OwnerAuth)
	if err != nil {
		return
	}
	result = uint32(retVal)
	return

}

//

// <param name="OwnerAuthStatus" type="uint32 "></param>
// <param name="ReturnValue" type="uint32 "></param>
func (instance *Win32_Tpm) GetOwnerAuthStatus( /* OUT */ OwnerAuthStatus uint32) (result uint32, err error) {
	retVal, err := instance.InvokeMethod("GetOwnerAuthStatus")
	if err != nil {
		return
	}
	retValue := retVal[0].(int32)
	result = uint32(retValue)
	return

}

//

// <param name="IsFIPS" type="bool "></param>
// <param name="ReturnValue" type="uint32 "></param>
func (instance *Win32_Tpm) IsFIPS( /* OUT */ IsFIPS bool) (result uint32, err error) {
	retVal, err := instance.InvokeMethod("IsFIPS")
	if err != nil {
		return
	}
	retValue := retVal[0].(int32)
	result = uint32(retValue)
	return

}

//

// <param name="LockoutRecovery" type="uint32 "></param>
// <param name="MaxTries" type="uint32 "></param>
// <param name="RecoveryTime" type="uint32 "></param>
// <param name="ReturnValue" type="uint32 "></param>
func (instance *Win32_Tpm) GetDictionaryAttackParameters( /* OUT */ MaxTries uint32,
	/* OUT */ RecoveryTime uint32,
	/* OUT */ LockoutRecovery uint32) (result uint32, err error) {
	retVal, err := instance.InvokeMethod("GetDictionaryAttackParameters")
	if err != nil {
		return
	}
	retValue := retVal[0].(int32)
	result = uint32(retValue)
	return

}

//

// <param name="LockoutCounter" type="uint32 "></param>
// <param name="MaxTries" type="uint32 "></param>
// <param name="ReturnValue" type="uint32 "></param>
func (instance *Win32_Tpm) GetCapLockoutInfo( /* OUT */ LockoutCounter uint32,
	/* OUT */ MaxTries uint32) (result uint32, err error) {
	retVal, err := instance.InvokeMethod("GetCapLockoutInfo")
	if err != nil {
		return
	}
	retValue := retVal[0].(int32)
	result = uint32(retValue)
	return

}

//

// <param name="IsLockedOut" type="bool "></param>
// <param name="ReturnValue" type="uint32 "></param>
func (instance *Win32_Tpm) IsLockedOut( /* OUT */ IsLockedOut bool) (result uint32, err error) {
	retVal, err := instance.InvokeMethod("IsLockedOut")
	if err != nil {
		return
	}
	retValue := retVal[0].(int32)
	result = uint32(retValue)
	return

}
