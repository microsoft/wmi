// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 3/19/2020 using wmigen
//      Source root.StandardCimv2
//////////////////////////////////////////////
package standardcimv2

import (
	"github.com/microsoft/wmi/pkg/base/query"
	cim "github.com/microsoft/wmi/pkg/wmiinstance"
)

// MSFT_NetTCPSetting struct
type MSFT_NetTCPSetting struct {
	*CIM_PolicyAction

	//
	AutomaticUseCustom uint8

	//
	AutoReusePortRangeNumberOfPorts uint16

	//
	AutoReusePortRangeStartPort uint16

	//
	AutoTuningLevelEffective uint8

	//
	AutoTuningLevelGroupPolicy uint8

	//
	AutoTuningLevelLocal uint8

	//
	CongestionProvider uint8

	//
	CwndRestart uint8

	//
	DelayedAckFrequency uint8

	//
	DelayedAckTimeout uint32

	//
	DynamicPortRangeNumberOfPorts uint16

	//
	DynamicPortRangeStartPort uint16

	//
	EcnCapability uint8

	//
	ForceWS uint8

	//
	InitialCongestionWindow uint32

	//
	InitialRto uint32

	//
	MaxSynRetransmissions uint8

	//
	MemoryPressureProtection uint8

	//
	MinRto uint32

	//
	NonSackRttResiliency uint8

	//
	ScalingHeuristics uint8

	//
	SettingName string

	//
	Timestamps uint8
}

func NewMSFT_NetTCPSettingEx1(instance *cim.WmiInstance) (newInstance *MSFT_NetTCPSetting, err error) {
	tmp, err := NewCIM_PolicyActionEx1(instance)

	if err != nil {
		return
	}
	newInstance = &MSFT_NetTCPSetting{
		CIM_PolicyAction: tmp,
	}
	return
}

func NewMSFT_NetTCPSettingEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *MSFT_NetTCPSetting, err error) {
	tmp, err := NewCIM_PolicyActionEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &MSFT_NetTCPSetting{
		CIM_PolicyAction: tmp,
	}
	return
}

// SetAutomaticUseCustom sets the value of AutomaticUseCustom for the instance
func (instance *MSFT_NetTCPSetting) SetPropertyAutomaticUseCustom(value uint8) (err error) {
	return instance.SetProperty("AutomaticUseCustom", value)
}

// GetAutomaticUseCustom gets the value of AutomaticUseCustom for the instance
func (instance *MSFT_NetTCPSetting) GetPropertyAutomaticUseCustom() (value uint8, err error) {
	retValue, err := instance.GetProperty("AutomaticUseCustom")
	if err != nil {
		return
	}
	value, ok := retValue.(uint8)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetAutoReusePortRangeNumberOfPorts sets the value of AutoReusePortRangeNumberOfPorts for the instance
func (instance *MSFT_NetTCPSetting) SetPropertyAutoReusePortRangeNumberOfPorts(value uint16) (err error) {
	return instance.SetProperty("AutoReusePortRangeNumberOfPorts", value)
}

// GetAutoReusePortRangeNumberOfPorts gets the value of AutoReusePortRangeNumberOfPorts for the instance
func (instance *MSFT_NetTCPSetting) GetPropertyAutoReusePortRangeNumberOfPorts() (value uint16, err error) {
	retValue, err := instance.GetProperty("AutoReusePortRangeNumberOfPorts")
	if err != nil {
		return
	}
	value, ok := retValue.(uint16)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetAutoReusePortRangeStartPort sets the value of AutoReusePortRangeStartPort for the instance
func (instance *MSFT_NetTCPSetting) SetPropertyAutoReusePortRangeStartPort(value uint16) (err error) {
	return instance.SetProperty("AutoReusePortRangeStartPort", value)
}

// GetAutoReusePortRangeStartPort gets the value of AutoReusePortRangeStartPort for the instance
func (instance *MSFT_NetTCPSetting) GetPropertyAutoReusePortRangeStartPort() (value uint16, err error) {
	retValue, err := instance.GetProperty("AutoReusePortRangeStartPort")
	if err != nil {
		return
	}
	value, ok := retValue.(uint16)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetAutoTuningLevelEffective sets the value of AutoTuningLevelEffective for the instance
func (instance *MSFT_NetTCPSetting) SetPropertyAutoTuningLevelEffective(value uint8) (err error) {
	return instance.SetProperty("AutoTuningLevelEffective", value)
}

// GetAutoTuningLevelEffective gets the value of AutoTuningLevelEffective for the instance
func (instance *MSFT_NetTCPSetting) GetPropertyAutoTuningLevelEffective() (value uint8, err error) {
	retValue, err := instance.GetProperty("AutoTuningLevelEffective")
	if err != nil {
		return
	}
	value, ok := retValue.(uint8)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetAutoTuningLevelGroupPolicy sets the value of AutoTuningLevelGroupPolicy for the instance
func (instance *MSFT_NetTCPSetting) SetPropertyAutoTuningLevelGroupPolicy(value uint8) (err error) {
	return instance.SetProperty("AutoTuningLevelGroupPolicy", value)
}

// GetAutoTuningLevelGroupPolicy gets the value of AutoTuningLevelGroupPolicy for the instance
func (instance *MSFT_NetTCPSetting) GetPropertyAutoTuningLevelGroupPolicy() (value uint8, err error) {
	retValue, err := instance.GetProperty("AutoTuningLevelGroupPolicy")
	if err != nil {
		return
	}
	value, ok := retValue.(uint8)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetAutoTuningLevelLocal sets the value of AutoTuningLevelLocal for the instance
func (instance *MSFT_NetTCPSetting) SetPropertyAutoTuningLevelLocal(value uint8) (err error) {
	return instance.SetProperty("AutoTuningLevelLocal", value)
}

// GetAutoTuningLevelLocal gets the value of AutoTuningLevelLocal for the instance
func (instance *MSFT_NetTCPSetting) GetPropertyAutoTuningLevelLocal() (value uint8, err error) {
	retValue, err := instance.GetProperty("AutoTuningLevelLocal")
	if err != nil {
		return
	}
	value, ok := retValue.(uint8)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetCongestionProvider sets the value of CongestionProvider for the instance
func (instance *MSFT_NetTCPSetting) SetPropertyCongestionProvider(value uint8) (err error) {
	return instance.SetProperty("CongestionProvider", value)
}

// GetCongestionProvider gets the value of CongestionProvider for the instance
func (instance *MSFT_NetTCPSetting) GetPropertyCongestionProvider() (value uint8, err error) {
	retValue, err := instance.GetProperty("CongestionProvider")
	if err != nil {
		return
	}
	value, ok := retValue.(uint8)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetCwndRestart sets the value of CwndRestart for the instance
func (instance *MSFT_NetTCPSetting) SetPropertyCwndRestart(value uint8) (err error) {
	return instance.SetProperty("CwndRestart", value)
}

// GetCwndRestart gets the value of CwndRestart for the instance
func (instance *MSFT_NetTCPSetting) GetPropertyCwndRestart() (value uint8, err error) {
	retValue, err := instance.GetProperty("CwndRestart")
	if err != nil {
		return
	}
	value, ok := retValue.(uint8)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetDelayedAckFrequency sets the value of DelayedAckFrequency for the instance
func (instance *MSFT_NetTCPSetting) SetPropertyDelayedAckFrequency(value uint8) (err error) {
	return instance.SetProperty("DelayedAckFrequency", value)
}

// GetDelayedAckFrequency gets the value of DelayedAckFrequency for the instance
func (instance *MSFT_NetTCPSetting) GetPropertyDelayedAckFrequency() (value uint8, err error) {
	retValue, err := instance.GetProperty("DelayedAckFrequency")
	if err != nil {
		return
	}
	value, ok := retValue.(uint8)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetDelayedAckTimeout sets the value of DelayedAckTimeout for the instance
func (instance *MSFT_NetTCPSetting) SetPropertyDelayedAckTimeout(value uint32) (err error) {
	return instance.SetProperty("DelayedAckTimeout", value)
}

// GetDelayedAckTimeout gets the value of DelayedAckTimeout for the instance
func (instance *MSFT_NetTCPSetting) GetPropertyDelayedAckTimeout() (value uint32, err error) {
	retValue, err := instance.GetProperty("DelayedAckTimeout")
	if err != nil {
		return
	}
	value, ok := retValue.(uint32)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetDynamicPortRangeNumberOfPorts sets the value of DynamicPortRangeNumberOfPorts for the instance
func (instance *MSFT_NetTCPSetting) SetPropertyDynamicPortRangeNumberOfPorts(value uint16) (err error) {
	return instance.SetProperty("DynamicPortRangeNumberOfPorts", value)
}

// GetDynamicPortRangeNumberOfPorts gets the value of DynamicPortRangeNumberOfPorts for the instance
func (instance *MSFT_NetTCPSetting) GetPropertyDynamicPortRangeNumberOfPorts() (value uint16, err error) {
	retValue, err := instance.GetProperty("DynamicPortRangeNumberOfPorts")
	if err != nil {
		return
	}
	value, ok := retValue.(uint16)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetDynamicPortRangeStartPort sets the value of DynamicPortRangeStartPort for the instance
func (instance *MSFT_NetTCPSetting) SetPropertyDynamicPortRangeStartPort(value uint16) (err error) {
	return instance.SetProperty("DynamicPortRangeStartPort", value)
}

// GetDynamicPortRangeStartPort gets the value of DynamicPortRangeStartPort for the instance
func (instance *MSFT_NetTCPSetting) GetPropertyDynamicPortRangeStartPort() (value uint16, err error) {
	retValue, err := instance.GetProperty("DynamicPortRangeStartPort")
	if err != nil {
		return
	}
	value, ok := retValue.(uint16)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetEcnCapability sets the value of EcnCapability for the instance
func (instance *MSFT_NetTCPSetting) SetPropertyEcnCapability(value uint8) (err error) {
	return instance.SetProperty("EcnCapability", value)
}

// GetEcnCapability gets the value of EcnCapability for the instance
func (instance *MSFT_NetTCPSetting) GetPropertyEcnCapability() (value uint8, err error) {
	retValue, err := instance.GetProperty("EcnCapability")
	if err != nil {
		return
	}
	value, ok := retValue.(uint8)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetForceWS sets the value of ForceWS for the instance
func (instance *MSFT_NetTCPSetting) SetPropertyForceWS(value uint8) (err error) {
	return instance.SetProperty("ForceWS", value)
}

// GetForceWS gets the value of ForceWS for the instance
func (instance *MSFT_NetTCPSetting) GetPropertyForceWS() (value uint8, err error) {
	retValue, err := instance.GetProperty("ForceWS")
	if err != nil {
		return
	}
	value, ok := retValue.(uint8)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetInitialCongestionWindow sets the value of InitialCongestionWindow for the instance
func (instance *MSFT_NetTCPSetting) SetPropertyInitialCongestionWindow(value uint32) (err error) {
	return instance.SetProperty("InitialCongestionWindow", value)
}

// GetInitialCongestionWindow gets the value of InitialCongestionWindow for the instance
func (instance *MSFT_NetTCPSetting) GetPropertyInitialCongestionWindow() (value uint32, err error) {
	retValue, err := instance.GetProperty("InitialCongestionWindow")
	if err != nil {
		return
	}
	value, ok := retValue.(uint32)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetInitialRto sets the value of InitialRto for the instance
func (instance *MSFT_NetTCPSetting) SetPropertyInitialRto(value uint32) (err error) {
	return instance.SetProperty("InitialRto", value)
}

// GetInitialRto gets the value of InitialRto for the instance
func (instance *MSFT_NetTCPSetting) GetPropertyInitialRto() (value uint32, err error) {
	retValue, err := instance.GetProperty("InitialRto")
	if err != nil {
		return
	}
	value, ok := retValue.(uint32)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetMaxSynRetransmissions sets the value of MaxSynRetransmissions for the instance
func (instance *MSFT_NetTCPSetting) SetPropertyMaxSynRetransmissions(value uint8) (err error) {
	return instance.SetProperty("MaxSynRetransmissions", value)
}

// GetMaxSynRetransmissions gets the value of MaxSynRetransmissions for the instance
func (instance *MSFT_NetTCPSetting) GetPropertyMaxSynRetransmissions() (value uint8, err error) {
	retValue, err := instance.GetProperty("MaxSynRetransmissions")
	if err != nil {
		return
	}
	value, ok := retValue.(uint8)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetMemoryPressureProtection sets the value of MemoryPressureProtection for the instance
func (instance *MSFT_NetTCPSetting) SetPropertyMemoryPressureProtection(value uint8) (err error) {
	return instance.SetProperty("MemoryPressureProtection", value)
}

// GetMemoryPressureProtection gets the value of MemoryPressureProtection for the instance
func (instance *MSFT_NetTCPSetting) GetPropertyMemoryPressureProtection() (value uint8, err error) {
	retValue, err := instance.GetProperty("MemoryPressureProtection")
	if err != nil {
		return
	}
	value, ok := retValue.(uint8)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetMinRto sets the value of MinRto for the instance
func (instance *MSFT_NetTCPSetting) SetPropertyMinRto(value uint32) (err error) {
	return instance.SetProperty("MinRto", value)
}

// GetMinRto gets the value of MinRto for the instance
func (instance *MSFT_NetTCPSetting) GetPropertyMinRto() (value uint32, err error) {
	retValue, err := instance.GetProperty("MinRto")
	if err != nil {
		return
	}
	value, ok := retValue.(uint32)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetNonSackRttResiliency sets the value of NonSackRttResiliency for the instance
func (instance *MSFT_NetTCPSetting) SetPropertyNonSackRttResiliency(value uint8) (err error) {
	return instance.SetProperty("NonSackRttResiliency", value)
}

// GetNonSackRttResiliency gets the value of NonSackRttResiliency for the instance
func (instance *MSFT_NetTCPSetting) GetPropertyNonSackRttResiliency() (value uint8, err error) {
	retValue, err := instance.GetProperty("NonSackRttResiliency")
	if err != nil {
		return
	}
	value, ok := retValue.(uint8)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetScalingHeuristics sets the value of ScalingHeuristics for the instance
func (instance *MSFT_NetTCPSetting) SetPropertyScalingHeuristics(value uint8) (err error) {
	return instance.SetProperty("ScalingHeuristics", value)
}

// GetScalingHeuristics gets the value of ScalingHeuristics for the instance
func (instance *MSFT_NetTCPSetting) GetPropertyScalingHeuristics() (value uint8, err error) {
	retValue, err := instance.GetProperty("ScalingHeuristics")
	if err != nil {
		return
	}
	value, ok := retValue.(uint8)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetSettingName sets the value of SettingName for the instance
func (instance *MSFT_NetTCPSetting) SetPropertySettingName(value string) (err error) {
	return instance.SetProperty("SettingName", value)
}

// GetSettingName gets the value of SettingName for the instance
func (instance *MSFT_NetTCPSetting) GetPropertySettingName() (value string, err error) {
	retValue, err := instance.GetProperty("SettingName")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetTimestamps sets the value of Timestamps for the instance
func (instance *MSFT_NetTCPSetting) SetPropertyTimestamps(value uint8) (err error) {
	return instance.SetProperty("Timestamps", value)
}

// GetTimestamps gets the value of Timestamps for the instance
func (instance *MSFT_NetTCPSetting) GetPropertyTimestamps() (value uint8, err error) {
	retValue, err := instance.GetProperty("Timestamps")
	if err != nil {
		return
	}
	value, ok := retValue.(uint8)
	if !ok {
		// TODO: Set an error
	}
	return
}
