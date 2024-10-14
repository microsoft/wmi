// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 6/14/2024 using wmigen
//      Source root.CIMV2
//////////////////////////////////////////////
package cimv2

import (
	"github.com/microsoft/wmi/pkg/base/query"
	"github.com/microsoft/wmi/pkg/errors"
	cim "github.com/microsoft/wmi/pkg/wmiinstance"
	"reflect"
)

// Win32_PerfRawData_Mlx5Provider_MellanoxWinOF2CongestionControl struct
type Win32_PerfRawData_Mlx5Provider_MellanoxWinOF2CongestionControl struct {
	*Win32_PerfRawData

	//
	NotificationPointCNPsSentSuccessfully uint64

	//
	NotificationPointRoCEv2ECNMarkedPackets uint64

	//
	ReactionPointAIFactorValue uint64

	//
	ReactionPointCurrentNumberofFlows uint64

	//
	ReactionPointGFactorValue uint64

	//
	ReactionPointIgnoredCNPPackets uint64

	//
	ReactionPointNumberofAIFactorDecreased uint64

	//
	ReactionPointNumberofAIFactorIncreased uint64

	//
	ReactionPointNumberofGFactorDecreased uint64

	//
	ReactionPointNumberofGFactorIncreased uint64

	//
	ReactionPointNumberofQPsUnderCCDecreased uint64

	//
	ReactionPointNumberofQPsUnderCCIncreased uint64

	//
	ReactionPointRTTValueinMicrosec uint64

	//
	ReactionPointSuccessfullyHandledCNPPackets uint64
}

func NewWin32_PerfRawData_Mlx5Provider_MellanoxWinOF2CongestionControlEx1(instance *cim.WmiInstance) (newInstance *Win32_PerfRawData_Mlx5Provider_MellanoxWinOF2CongestionControl, err error) {
	tmp, err := NewWin32_PerfRawDataEx1(instance)

	if err != nil {
		return
	}
	newInstance = &Win32_PerfRawData_Mlx5Provider_MellanoxWinOF2CongestionControl{
		Win32_PerfRawData: tmp,
	}
	return
}

func NewWin32_PerfRawData_Mlx5Provider_MellanoxWinOF2CongestionControlEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *Win32_PerfRawData_Mlx5Provider_MellanoxWinOF2CongestionControl, err error) {
	tmp, err := NewWin32_PerfRawDataEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &Win32_PerfRawData_Mlx5Provider_MellanoxWinOF2CongestionControl{
		Win32_PerfRawData: tmp,
	}
	return
}

// SetNotificationPointCNPsSentSuccessfully sets the value of NotificationPointCNPsSentSuccessfully for the instance
func (instance *Win32_PerfRawData_Mlx5Provider_MellanoxWinOF2CongestionControl) SetPropertyNotificationPointCNPsSentSuccessfully(value uint64) (err error) {
	return instance.SetProperty("NotificationPointCNPsSentSuccessfully", (value))
}

// GetNotificationPointCNPsSentSuccessfully gets the value of NotificationPointCNPsSentSuccessfully for the instance
func (instance *Win32_PerfRawData_Mlx5Provider_MellanoxWinOF2CongestionControl) GetPropertyNotificationPointCNPsSentSuccessfully() (value uint64, err error) {
	retValue, err := instance.GetProperty("NotificationPointCNPsSentSuccessfully")
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

// SetNotificationPointRoCEv2ECNMarkedPackets sets the value of NotificationPointRoCEv2ECNMarkedPackets for the instance
func (instance *Win32_PerfRawData_Mlx5Provider_MellanoxWinOF2CongestionControl) SetPropertyNotificationPointRoCEv2ECNMarkedPackets(value uint64) (err error) {
	return instance.SetProperty("NotificationPointRoCEv2ECNMarkedPackets", (value))
}

// GetNotificationPointRoCEv2ECNMarkedPackets gets the value of NotificationPointRoCEv2ECNMarkedPackets for the instance
func (instance *Win32_PerfRawData_Mlx5Provider_MellanoxWinOF2CongestionControl) GetPropertyNotificationPointRoCEv2ECNMarkedPackets() (value uint64, err error) {
	retValue, err := instance.GetProperty("NotificationPointRoCEv2ECNMarkedPackets")
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

// SetReactionPointAIFactorValue sets the value of ReactionPointAIFactorValue for the instance
func (instance *Win32_PerfRawData_Mlx5Provider_MellanoxWinOF2CongestionControl) SetPropertyReactionPointAIFactorValue(value uint64) (err error) {
	return instance.SetProperty("ReactionPointAIFactorValue", (value))
}

// GetReactionPointAIFactorValue gets the value of ReactionPointAIFactorValue for the instance
func (instance *Win32_PerfRawData_Mlx5Provider_MellanoxWinOF2CongestionControl) GetPropertyReactionPointAIFactorValue() (value uint64, err error) {
	retValue, err := instance.GetProperty("ReactionPointAIFactorValue")
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

// SetReactionPointCurrentNumberofFlows sets the value of ReactionPointCurrentNumberofFlows for the instance
func (instance *Win32_PerfRawData_Mlx5Provider_MellanoxWinOF2CongestionControl) SetPropertyReactionPointCurrentNumberofFlows(value uint64) (err error) {
	return instance.SetProperty("ReactionPointCurrentNumberofFlows", (value))
}

// GetReactionPointCurrentNumberofFlows gets the value of ReactionPointCurrentNumberofFlows for the instance
func (instance *Win32_PerfRawData_Mlx5Provider_MellanoxWinOF2CongestionControl) GetPropertyReactionPointCurrentNumberofFlows() (value uint64, err error) {
	retValue, err := instance.GetProperty("ReactionPointCurrentNumberofFlows")
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

// SetReactionPointGFactorValue sets the value of ReactionPointGFactorValue for the instance
func (instance *Win32_PerfRawData_Mlx5Provider_MellanoxWinOF2CongestionControl) SetPropertyReactionPointGFactorValue(value uint64) (err error) {
	return instance.SetProperty("ReactionPointGFactorValue", (value))
}

// GetReactionPointGFactorValue gets the value of ReactionPointGFactorValue for the instance
func (instance *Win32_PerfRawData_Mlx5Provider_MellanoxWinOF2CongestionControl) GetPropertyReactionPointGFactorValue() (value uint64, err error) {
	retValue, err := instance.GetProperty("ReactionPointGFactorValue")
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

// SetReactionPointIgnoredCNPPackets sets the value of ReactionPointIgnoredCNPPackets for the instance
func (instance *Win32_PerfRawData_Mlx5Provider_MellanoxWinOF2CongestionControl) SetPropertyReactionPointIgnoredCNPPackets(value uint64) (err error) {
	return instance.SetProperty("ReactionPointIgnoredCNPPackets", (value))
}

// GetReactionPointIgnoredCNPPackets gets the value of ReactionPointIgnoredCNPPackets for the instance
func (instance *Win32_PerfRawData_Mlx5Provider_MellanoxWinOF2CongestionControl) GetPropertyReactionPointIgnoredCNPPackets() (value uint64, err error) {
	retValue, err := instance.GetProperty("ReactionPointIgnoredCNPPackets")
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

// SetReactionPointNumberofAIFactorDecreased sets the value of ReactionPointNumberofAIFactorDecreased for the instance
func (instance *Win32_PerfRawData_Mlx5Provider_MellanoxWinOF2CongestionControl) SetPropertyReactionPointNumberofAIFactorDecreased(value uint64) (err error) {
	return instance.SetProperty("ReactionPointNumberofAIFactorDecreased", (value))
}

// GetReactionPointNumberofAIFactorDecreased gets the value of ReactionPointNumberofAIFactorDecreased for the instance
func (instance *Win32_PerfRawData_Mlx5Provider_MellanoxWinOF2CongestionControl) GetPropertyReactionPointNumberofAIFactorDecreased() (value uint64, err error) {
	retValue, err := instance.GetProperty("ReactionPointNumberofAIFactorDecreased")
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

// SetReactionPointNumberofAIFactorIncreased sets the value of ReactionPointNumberofAIFactorIncreased for the instance
func (instance *Win32_PerfRawData_Mlx5Provider_MellanoxWinOF2CongestionControl) SetPropertyReactionPointNumberofAIFactorIncreased(value uint64) (err error) {
	return instance.SetProperty("ReactionPointNumberofAIFactorIncreased", (value))
}

// GetReactionPointNumberofAIFactorIncreased gets the value of ReactionPointNumberofAIFactorIncreased for the instance
func (instance *Win32_PerfRawData_Mlx5Provider_MellanoxWinOF2CongestionControl) GetPropertyReactionPointNumberofAIFactorIncreased() (value uint64, err error) {
	retValue, err := instance.GetProperty("ReactionPointNumberofAIFactorIncreased")
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

// SetReactionPointNumberofGFactorDecreased sets the value of ReactionPointNumberofGFactorDecreased for the instance
func (instance *Win32_PerfRawData_Mlx5Provider_MellanoxWinOF2CongestionControl) SetPropertyReactionPointNumberofGFactorDecreased(value uint64) (err error) {
	return instance.SetProperty("ReactionPointNumberofGFactorDecreased", (value))
}

// GetReactionPointNumberofGFactorDecreased gets the value of ReactionPointNumberofGFactorDecreased for the instance
func (instance *Win32_PerfRawData_Mlx5Provider_MellanoxWinOF2CongestionControl) GetPropertyReactionPointNumberofGFactorDecreased() (value uint64, err error) {
	retValue, err := instance.GetProperty("ReactionPointNumberofGFactorDecreased")
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

// SetReactionPointNumberofGFactorIncreased sets the value of ReactionPointNumberofGFactorIncreased for the instance
func (instance *Win32_PerfRawData_Mlx5Provider_MellanoxWinOF2CongestionControl) SetPropertyReactionPointNumberofGFactorIncreased(value uint64) (err error) {
	return instance.SetProperty("ReactionPointNumberofGFactorIncreased", (value))
}

// GetReactionPointNumberofGFactorIncreased gets the value of ReactionPointNumberofGFactorIncreased for the instance
func (instance *Win32_PerfRawData_Mlx5Provider_MellanoxWinOF2CongestionControl) GetPropertyReactionPointNumberofGFactorIncreased() (value uint64, err error) {
	retValue, err := instance.GetProperty("ReactionPointNumberofGFactorIncreased")
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

// SetReactionPointNumberofQPsUnderCCDecreased sets the value of ReactionPointNumberofQPsUnderCCDecreased for the instance
func (instance *Win32_PerfRawData_Mlx5Provider_MellanoxWinOF2CongestionControl) SetPropertyReactionPointNumberofQPsUnderCCDecreased(value uint64) (err error) {
	return instance.SetProperty("ReactionPointNumberofQPsUnderCCDecreased", (value))
}

// GetReactionPointNumberofQPsUnderCCDecreased gets the value of ReactionPointNumberofQPsUnderCCDecreased for the instance
func (instance *Win32_PerfRawData_Mlx5Provider_MellanoxWinOF2CongestionControl) GetPropertyReactionPointNumberofQPsUnderCCDecreased() (value uint64, err error) {
	retValue, err := instance.GetProperty("ReactionPointNumberofQPsUnderCCDecreased")
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

// SetReactionPointNumberofQPsUnderCCIncreased sets the value of ReactionPointNumberofQPsUnderCCIncreased for the instance
func (instance *Win32_PerfRawData_Mlx5Provider_MellanoxWinOF2CongestionControl) SetPropertyReactionPointNumberofQPsUnderCCIncreased(value uint64) (err error) {
	return instance.SetProperty("ReactionPointNumberofQPsUnderCCIncreased", (value))
}

// GetReactionPointNumberofQPsUnderCCIncreased gets the value of ReactionPointNumberofQPsUnderCCIncreased for the instance
func (instance *Win32_PerfRawData_Mlx5Provider_MellanoxWinOF2CongestionControl) GetPropertyReactionPointNumberofQPsUnderCCIncreased() (value uint64, err error) {
	retValue, err := instance.GetProperty("ReactionPointNumberofQPsUnderCCIncreased")
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

// SetReactionPointRTTValueinMicrosec sets the value of ReactionPointRTTValueinMicrosec for the instance
func (instance *Win32_PerfRawData_Mlx5Provider_MellanoxWinOF2CongestionControl) SetPropertyReactionPointRTTValueinMicrosec(value uint64) (err error) {
	return instance.SetProperty("ReactionPointRTTValueinMicrosec", (value))
}

// GetReactionPointRTTValueinMicrosec gets the value of ReactionPointRTTValueinMicrosec for the instance
func (instance *Win32_PerfRawData_Mlx5Provider_MellanoxWinOF2CongestionControl) GetPropertyReactionPointRTTValueinMicrosec() (value uint64, err error) {
	retValue, err := instance.GetProperty("ReactionPointRTTValueinMicrosec")
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

// SetReactionPointSuccessfullyHandledCNPPackets sets the value of ReactionPointSuccessfullyHandledCNPPackets for the instance
func (instance *Win32_PerfRawData_Mlx5Provider_MellanoxWinOF2CongestionControl) SetPropertyReactionPointSuccessfullyHandledCNPPackets(value uint64) (err error) {
	return instance.SetProperty("ReactionPointSuccessfullyHandledCNPPackets", (value))
}

// GetReactionPointSuccessfullyHandledCNPPackets gets the value of ReactionPointSuccessfullyHandledCNPPackets for the instance
func (instance *Win32_PerfRawData_Mlx5Provider_MellanoxWinOF2CongestionControl) GetPropertyReactionPointSuccessfullyHandledCNPPackets() (value uint64, err error) {
	retValue, err := instance.GetProperty("ReactionPointSuccessfullyHandledCNPPackets")
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
