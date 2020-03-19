// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 3/19/2020 using wmigen
//      Source root.StandardCimv2.mlnx
//////////////////////////////////////////////
package mlnx

import (
	"github.com/microsoft/wmi/pkg/base/query"
	cim "github.com/microsoft/wmi/pkg/wmiinstance"
)

// MLNX_NetAdapterEcnSettingData struct
type MLNX_NetAdapterEcnSettingData struct {
	*MLNX_NetAdapterSettingData

	//
	EcnAlphaToRateCoeff uint32

	//
	EcnAlphaToRateShift uint32

	//
	EcnBurstSize uint32

	//
	EcnCatchRateLimitEn uint32

	//
	EcnClampTgtRate uint32

	//
	EcnClampTgtRateAfterTimeInc uint32

	//
	EcnCnp802pPrio uint32

	//
	EcnCnpDscp uint32

	//
	EcnCnpRecEn uint32

	//
	EcnCnpTimer uint32

	//
	EcnCoalesceCnpInRp uint32

	//
	EcnCompnRateLimit uint32

	//
	EcnDceTcpG uint32

	//
	EcnDceTcpRtt uint32

	//
	EcnDceTcpRttDelay uint32

	//
	EcnEnable bool

	//
	EcnExpectIpv6 uint32

	//
	EcnExpectVlanTagged uint32

	//
	EcnFastRise uint32

	//
	EcnForceRcTos uint32

	//
	EcnForceUcTos uint32

	//
	EcnInitialAlphaValue uint32

	//
	EcnMarkedRatioMultiplier uint32

	//
	EcnMarkedRatioShift uint32

	//
	EcnMaxByteRise uint32

	//
	EcnMaxTimeRise uint32

	//
	EcnMinLosslessBufferForCatches uint32

	//
	EcnMinLossyBufferForCatches uint32

	//
	EcnNoCongestionCyclesToKeep uint32

	//
	EcnNpRecEn uint32

	//
	EcnNumInjector uint32

	//
	EcnPriorityEnable uint32

	//
	EcnRateToSetOnFirstCnp uint32

	//
	EcnRpgAiRate uint32

	//
	EcnRpgByteReset uint32

	//
	EcnRpgHaiRate uint32

	//
	EcnRpgMaxRate uint32

	//
	EcnRpgMinDecFac uint32

	//
	EcnRpgMinRate uint32

	//
	EcnRpgThreshold uint32

	//
	EcnRpgTimeReset uint32

	//
	EcnSupportIBStandardCnp uint32

	//
	EcnTrapDisablePendingCnpThreshold uint32

	//
	EcnTrapDisablePeriod uint32

	//
	EcnUseIBStandardCnp uint32
}

func NewMLNX_NetAdapterEcnSettingDataEx1(instance *cim.WmiInstance) (newInstance *MLNX_NetAdapterEcnSettingData, err error) {
	tmp, err := NewMLNX_NetAdapterSettingDataEx1(instance)

	if err != nil {
		return
	}
	newInstance = &MLNX_NetAdapterEcnSettingData{
		MLNX_NetAdapterSettingData: tmp,
	}
	return
}

func NewMLNX_NetAdapterEcnSettingDataEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *MLNX_NetAdapterEcnSettingData, err error) {
	tmp, err := NewMLNX_NetAdapterSettingDataEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &MLNX_NetAdapterEcnSettingData{
		MLNX_NetAdapterSettingData: tmp,
	}
	return
}

// SetEcnAlphaToRateCoeff sets the value of EcnAlphaToRateCoeff for the instance
func (instance *MLNX_NetAdapterEcnSettingData) SetPropertyEcnAlphaToRateCoeff(value uint32) (err error) {
	return instance.SetProperty("EcnAlphaToRateCoeff", value)
}

// GetEcnAlphaToRateCoeff gets the value of EcnAlphaToRateCoeff for the instance
func (instance *MLNX_NetAdapterEcnSettingData) GetPropertyEcnAlphaToRateCoeff() (value uint32, err error) {
	retValue, err := instance.GetProperty("EcnAlphaToRateCoeff")
	if err != nil {
		return
	}
	value, ok := retValue.(uint32)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetEcnAlphaToRateShift sets the value of EcnAlphaToRateShift for the instance
func (instance *MLNX_NetAdapterEcnSettingData) SetPropertyEcnAlphaToRateShift(value uint32) (err error) {
	return instance.SetProperty("EcnAlphaToRateShift", value)
}

// GetEcnAlphaToRateShift gets the value of EcnAlphaToRateShift for the instance
func (instance *MLNX_NetAdapterEcnSettingData) GetPropertyEcnAlphaToRateShift() (value uint32, err error) {
	retValue, err := instance.GetProperty("EcnAlphaToRateShift")
	if err != nil {
		return
	}
	value, ok := retValue.(uint32)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetEcnBurstSize sets the value of EcnBurstSize for the instance
func (instance *MLNX_NetAdapterEcnSettingData) SetPropertyEcnBurstSize(value uint32) (err error) {
	return instance.SetProperty("EcnBurstSize", value)
}

// GetEcnBurstSize gets the value of EcnBurstSize for the instance
func (instance *MLNX_NetAdapterEcnSettingData) GetPropertyEcnBurstSize() (value uint32, err error) {
	retValue, err := instance.GetProperty("EcnBurstSize")
	if err != nil {
		return
	}
	value, ok := retValue.(uint32)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetEcnCatchRateLimitEn sets the value of EcnCatchRateLimitEn for the instance
func (instance *MLNX_NetAdapterEcnSettingData) SetPropertyEcnCatchRateLimitEn(value uint32) (err error) {
	return instance.SetProperty("EcnCatchRateLimitEn", value)
}

// GetEcnCatchRateLimitEn gets the value of EcnCatchRateLimitEn for the instance
func (instance *MLNX_NetAdapterEcnSettingData) GetPropertyEcnCatchRateLimitEn() (value uint32, err error) {
	retValue, err := instance.GetProperty("EcnCatchRateLimitEn")
	if err != nil {
		return
	}
	value, ok := retValue.(uint32)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetEcnClampTgtRate sets the value of EcnClampTgtRate for the instance
func (instance *MLNX_NetAdapterEcnSettingData) SetPropertyEcnClampTgtRate(value uint32) (err error) {
	return instance.SetProperty("EcnClampTgtRate", value)
}

// GetEcnClampTgtRate gets the value of EcnClampTgtRate for the instance
func (instance *MLNX_NetAdapterEcnSettingData) GetPropertyEcnClampTgtRate() (value uint32, err error) {
	retValue, err := instance.GetProperty("EcnClampTgtRate")
	if err != nil {
		return
	}
	value, ok := retValue.(uint32)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetEcnClampTgtRateAfterTimeInc sets the value of EcnClampTgtRateAfterTimeInc for the instance
func (instance *MLNX_NetAdapterEcnSettingData) SetPropertyEcnClampTgtRateAfterTimeInc(value uint32) (err error) {
	return instance.SetProperty("EcnClampTgtRateAfterTimeInc", value)
}

// GetEcnClampTgtRateAfterTimeInc gets the value of EcnClampTgtRateAfterTimeInc for the instance
func (instance *MLNX_NetAdapterEcnSettingData) GetPropertyEcnClampTgtRateAfterTimeInc() (value uint32, err error) {
	retValue, err := instance.GetProperty("EcnClampTgtRateAfterTimeInc")
	if err != nil {
		return
	}
	value, ok := retValue.(uint32)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetEcnCnp802pPrio sets the value of EcnCnp802pPrio for the instance
func (instance *MLNX_NetAdapterEcnSettingData) SetPropertyEcnCnp802pPrio(value uint32) (err error) {
	return instance.SetProperty("EcnCnp802pPrio", value)
}

// GetEcnCnp802pPrio gets the value of EcnCnp802pPrio for the instance
func (instance *MLNX_NetAdapterEcnSettingData) GetPropertyEcnCnp802pPrio() (value uint32, err error) {
	retValue, err := instance.GetProperty("EcnCnp802pPrio")
	if err != nil {
		return
	}
	value, ok := retValue.(uint32)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetEcnCnpDscp sets the value of EcnCnpDscp for the instance
func (instance *MLNX_NetAdapterEcnSettingData) SetPropertyEcnCnpDscp(value uint32) (err error) {
	return instance.SetProperty("EcnCnpDscp", value)
}

// GetEcnCnpDscp gets the value of EcnCnpDscp for the instance
func (instance *MLNX_NetAdapterEcnSettingData) GetPropertyEcnCnpDscp() (value uint32, err error) {
	retValue, err := instance.GetProperty("EcnCnpDscp")
	if err != nil {
		return
	}
	value, ok := retValue.(uint32)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetEcnCnpRecEn sets the value of EcnCnpRecEn for the instance
func (instance *MLNX_NetAdapterEcnSettingData) SetPropertyEcnCnpRecEn(value uint32) (err error) {
	return instance.SetProperty("EcnCnpRecEn", value)
}

// GetEcnCnpRecEn gets the value of EcnCnpRecEn for the instance
func (instance *MLNX_NetAdapterEcnSettingData) GetPropertyEcnCnpRecEn() (value uint32, err error) {
	retValue, err := instance.GetProperty("EcnCnpRecEn")
	if err != nil {
		return
	}
	value, ok := retValue.(uint32)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetEcnCnpTimer sets the value of EcnCnpTimer for the instance
func (instance *MLNX_NetAdapterEcnSettingData) SetPropertyEcnCnpTimer(value uint32) (err error) {
	return instance.SetProperty("EcnCnpTimer", value)
}

// GetEcnCnpTimer gets the value of EcnCnpTimer for the instance
func (instance *MLNX_NetAdapterEcnSettingData) GetPropertyEcnCnpTimer() (value uint32, err error) {
	retValue, err := instance.GetProperty("EcnCnpTimer")
	if err != nil {
		return
	}
	value, ok := retValue.(uint32)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetEcnCoalesceCnpInRp sets the value of EcnCoalesceCnpInRp for the instance
func (instance *MLNX_NetAdapterEcnSettingData) SetPropertyEcnCoalesceCnpInRp(value uint32) (err error) {
	return instance.SetProperty("EcnCoalesceCnpInRp", value)
}

// GetEcnCoalesceCnpInRp gets the value of EcnCoalesceCnpInRp for the instance
func (instance *MLNX_NetAdapterEcnSettingData) GetPropertyEcnCoalesceCnpInRp() (value uint32, err error) {
	retValue, err := instance.GetProperty("EcnCoalesceCnpInRp")
	if err != nil {
		return
	}
	value, ok := retValue.(uint32)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetEcnCompnRateLimit sets the value of EcnCompnRateLimit for the instance
func (instance *MLNX_NetAdapterEcnSettingData) SetPropertyEcnCompnRateLimit(value uint32) (err error) {
	return instance.SetProperty("EcnCompnRateLimit", value)
}

// GetEcnCompnRateLimit gets the value of EcnCompnRateLimit for the instance
func (instance *MLNX_NetAdapterEcnSettingData) GetPropertyEcnCompnRateLimit() (value uint32, err error) {
	retValue, err := instance.GetProperty("EcnCompnRateLimit")
	if err != nil {
		return
	}
	value, ok := retValue.(uint32)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetEcnDceTcpG sets the value of EcnDceTcpG for the instance
func (instance *MLNX_NetAdapterEcnSettingData) SetPropertyEcnDceTcpG(value uint32) (err error) {
	return instance.SetProperty("EcnDceTcpG", value)
}

// GetEcnDceTcpG gets the value of EcnDceTcpG for the instance
func (instance *MLNX_NetAdapterEcnSettingData) GetPropertyEcnDceTcpG() (value uint32, err error) {
	retValue, err := instance.GetProperty("EcnDceTcpG")
	if err != nil {
		return
	}
	value, ok := retValue.(uint32)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetEcnDceTcpRtt sets the value of EcnDceTcpRtt for the instance
func (instance *MLNX_NetAdapterEcnSettingData) SetPropertyEcnDceTcpRtt(value uint32) (err error) {
	return instance.SetProperty("EcnDceTcpRtt", value)
}

// GetEcnDceTcpRtt gets the value of EcnDceTcpRtt for the instance
func (instance *MLNX_NetAdapterEcnSettingData) GetPropertyEcnDceTcpRtt() (value uint32, err error) {
	retValue, err := instance.GetProperty("EcnDceTcpRtt")
	if err != nil {
		return
	}
	value, ok := retValue.(uint32)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetEcnDceTcpRttDelay sets the value of EcnDceTcpRttDelay for the instance
func (instance *MLNX_NetAdapterEcnSettingData) SetPropertyEcnDceTcpRttDelay(value uint32) (err error) {
	return instance.SetProperty("EcnDceTcpRttDelay", value)
}

// GetEcnDceTcpRttDelay gets the value of EcnDceTcpRttDelay for the instance
func (instance *MLNX_NetAdapterEcnSettingData) GetPropertyEcnDceTcpRttDelay() (value uint32, err error) {
	retValue, err := instance.GetProperty("EcnDceTcpRttDelay")
	if err != nil {
		return
	}
	value, ok := retValue.(uint32)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetEcnEnable sets the value of EcnEnable for the instance
func (instance *MLNX_NetAdapterEcnSettingData) SetPropertyEcnEnable(value bool) (err error) {
	return instance.SetProperty("EcnEnable", value)
}

// GetEcnEnable gets the value of EcnEnable for the instance
func (instance *MLNX_NetAdapterEcnSettingData) GetPropertyEcnEnable() (value bool, err error) {
	retValue, err := instance.GetProperty("EcnEnable")
	if err != nil {
		return
	}
	value, ok := retValue.(bool)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetEcnExpectIpv6 sets the value of EcnExpectIpv6 for the instance
func (instance *MLNX_NetAdapterEcnSettingData) SetPropertyEcnExpectIpv6(value uint32) (err error) {
	return instance.SetProperty("EcnExpectIpv6", value)
}

// GetEcnExpectIpv6 gets the value of EcnExpectIpv6 for the instance
func (instance *MLNX_NetAdapterEcnSettingData) GetPropertyEcnExpectIpv6() (value uint32, err error) {
	retValue, err := instance.GetProperty("EcnExpectIpv6")
	if err != nil {
		return
	}
	value, ok := retValue.(uint32)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetEcnExpectVlanTagged sets the value of EcnExpectVlanTagged for the instance
func (instance *MLNX_NetAdapterEcnSettingData) SetPropertyEcnExpectVlanTagged(value uint32) (err error) {
	return instance.SetProperty("EcnExpectVlanTagged", value)
}

// GetEcnExpectVlanTagged gets the value of EcnExpectVlanTagged for the instance
func (instance *MLNX_NetAdapterEcnSettingData) GetPropertyEcnExpectVlanTagged() (value uint32, err error) {
	retValue, err := instance.GetProperty("EcnExpectVlanTagged")
	if err != nil {
		return
	}
	value, ok := retValue.(uint32)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetEcnFastRise sets the value of EcnFastRise for the instance
func (instance *MLNX_NetAdapterEcnSettingData) SetPropertyEcnFastRise(value uint32) (err error) {
	return instance.SetProperty("EcnFastRise", value)
}

// GetEcnFastRise gets the value of EcnFastRise for the instance
func (instance *MLNX_NetAdapterEcnSettingData) GetPropertyEcnFastRise() (value uint32, err error) {
	retValue, err := instance.GetProperty("EcnFastRise")
	if err != nil {
		return
	}
	value, ok := retValue.(uint32)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetEcnForceRcTos sets the value of EcnForceRcTos for the instance
func (instance *MLNX_NetAdapterEcnSettingData) SetPropertyEcnForceRcTos(value uint32) (err error) {
	return instance.SetProperty("EcnForceRcTos", value)
}

// GetEcnForceRcTos gets the value of EcnForceRcTos for the instance
func (instance *MLNX_NetAdapterEcnSettingData) GetPropertyEcnForceRcTos() (value uint32, err error) {
	retValue, err := instance.GetProperty("EcnForceRcTos")
	if err != nil {
		return
	}
	value, ok := retValue.(uint32)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetEcnForceUcTos sets the value of EcnForceUcTos for the instance
func (instance *MLNX_NetAdapterEcnSettingData) SetPropertyEcnForceUcTos(value uint32) (err error) {
	return instance.SetProperty("EcnForceUcTos", value)
}

// GetEcnForceUcTos gets the value of EcnForceUcTos for the instance
func (instance *MLNX_NetAdapterEcnSettingData) GetPropertyEcnForceUcTos() (value uint32, err error) {
	retValue, err := instance.GetProperty("EcnForceUcTos")
	if err != nil {
		return
	}
	value, ok := retValue.(uint32)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetEcnInitialAlphaValue sets the value of EcnInitialAlphaValue for the instance
func (instance *MLNX_NetAdapterEcnSettingData) SetPropertyEcnInitialAlphaValue(value uint32) (err error) {
	return instance.SetProperty("EcnInitialAlphaValue", value)
}

// GetEcnInitialAlphaValue gets the value of EcnInitialAlphaValue for the instance
func (instance *MLNX_NetAdapterEcnSettingData) GetPropertyEcnInitialAlphaValue() (value uint32, err error) {
	retValue, err := instance.GetProperty("EcnInitialAlphaValue")
	if err != nil {
		return
	}
	value, ok := retValue.(uint32)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetEcnMarkedRatioMultiplier sets the value of EcnMarkedRatioMultiplier for the instance
func (instance *MLNX_NetAdapterEcnSettingData) SetPropertyEcnMarkedRatioMultiplier(value uint32) (err error) {
	return instance.SetProperty("EcnMarkedRatioMultiplier", value)
}

// GetEcnMarkedRatioMultiplier gets the value of EcnMarkedRatioMultiplier for the instance
func (instance *MLNX_NetAdapterEcnSettingData) GetPropertyEcnMarkedRatioMultiplier() (value uint32, err error) {
	retValue, err := instance.GetProperty("EcnMarkedRatioMultiplier")
	if err != nil {
		return
	}
	value, ok := retValue.(uint32)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetEcnMarkedRatioShift sets the value of EcnMarkedRatioShift for the instance
func (instance *MLNX_NetAdapterEcnSettingData) SetPropertyEcnMarkedRatioShift(value uint32) (err error) {
	return instance.SetProperty("EcnMarkedRatioShift", value)
}

// GetEcnMarkedRatioShift gets the value of EcnMarkedRatioShift for the instance
func (instance *MLNX_NetAdapterEcnSettingData) GetPropertyEcnMarkedRatioShift() (value uint32, err error) {
	retValue, err := instance.GetProperty("EcnMarkedRatioShift")
	if err != nil {
		return
	}
	value, ok := retValue.(uint32)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetEcnMaxByteRise sets the value of EcnMaxByteRise for the instance
func (instance *MLNX_NetAdapterEcnSettingData) SetPropertyEcnMaxByteRise(value uint32) (err error) {
	return instance.SetProperty("EcnMaxByteRise", value)
}

// GetEcnMaxByteRise gets the value of EcnMaxByteRise for the instance
func (instance *MLNX_NetAdapterEcnSettingData) GetPropertyEcnMaxByteRise() (value uint32, err error) {
	retValue, err := instance.GetProperty("EcnMaxByteRise")
	if err != nil {
		return
	}
	value, ok := retValue.(uint32)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetEcnMaxTimeRise sets the value of EcnMaxTimeRise for the instance
func (instance *MLNX_NetAdapterEcnSettingData) SetPropertyEcnMaxTimeRise(value uint32) (err error) {
	return instance.SetProperty("EcnMaxTimeRise", value)
}

// GetEcnMaxTimeRise gets the value of EcnMaxTimeRise for the instance
func (instance *MLNX_NetAdapterEcnSettingData) GetPropertyEcnMaxTimeRise() (value uint32, err error) {
	retValue, err := instance.GetProperty("EcnMaxTimeRise")
	if err != nil {
		return
	}
	value, ok := retValue.(uint32)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetEcnMinLosslessBufferForCatches sets the value of EcnMinLosslessBufferForCatches for the instance
func (instance *MLNX_NetAdapterEcnSettingData) SetPropertyEcnMinLosslessBufferForCatches(value uint32) (err error) {
	return instance.SetProperty("EcnMinLosslessBufferForCatches", value)
}

// GetEcnMinLosslessBufferForCatches gets the value of EcnMinLosslessBufferForCatches for the instance
func (instance *MLNX_NetAdapterEcnSettingData) GetPropertyEcnMinLosslessBufferForCatches() (value uint32, err error) {
	retValue, err := instance.GetProperty("EcnMinLosslessBufferForCatches")
	if err != nil {
		return
	}
	value, ok := retValue.(uint32)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetEcnMinLossyBufferForCatches sets the value of EcnMinLossyBufferForCatches for the instance
func (instance *MLNX_NetAdapterEcnSettingData) SetPropertyEcnMinLossyBufferForCatches(value uint32) (err error) {
	return instance.SetProperty("EcnMinLossyBufferForCatches", value)
}

// GetEcnMinLossyBufferForCatches gets the value of EcnMinLossyBufferForCatches for the instance
func (instance *MLNX_NetAdapterEcnSettingData) GetPropertyEcnMinLossyBufferForCatches() (value uint32, err error) {
	retValue, err := instance.GetProperty("EcnMinLossyBufferForCatches")
	if err != nil {
		return
	}
	value, ok := retValue.(uint32)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetEcnNoCongestionCyclesToKeep sets the value of EcnNoCongestionCyclesToKeep for the instance
func (instance *MLNX_NetAdapterEcnSettingData) SetPropertyEcnNoCongestionCyclesToKeep(value uint32) (err error) {
	return instance.SetProperty("EcnNoCongestionCyclesToKeep", value)
}

// GetEcnNoCongestionCyclesToKeep gets the value of EcnNoCongestionCyclesToKeep for the instance
func (instance *MLNX_NetAdapterEcnSettingData) GetPropertyEcnNoCongestionCyclesToKeep() (value uint32, err error) {
	retValue, err := instance.GetProperty("EcnNoCongestionCyclesToKeep")
	if err != nil {
		return
	}
	value, ok := retValue.(uint32)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetEcnNpRecEn sets the value of EcnNpRecEn for the instance
func (instance *MLNX_NetAdapterEcnSettingData) SetPropertyEcnNpRecEn(value uint32) (err error) {
	return instance.SetProperty("EcnNpRecEn", value)
}

// GetEcnNpRecEn gets the value of EcnNpRecEn for the instance
func (instance *MLNX_NetAdapterEcnSettingData) GetPropertyEcnNpRecEn() (value uint32, err error) {
	retValue, err := instance.GetProperty("EcnNpRecEn")
	if err != nil {
		return
	}
	value, ok := retValue.(uint32)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetEcnNumInjector sets the value of EcnNumInjector for the instance
func (instance *MLNX_NetAdapterEcnSettingData) SetPropertyEcnNumInjector(value uint32) (err error) {
	return instance.SetProperty("EcnNumInjector", value)
}

// GetEcnNumInjector gets the value of EcnNumInjector for the instance
func (instance *MLNX_NetAdapterEcnSettingData) GetPropertyEcnNumInjector() (value uint32, err error) {
	retValue, err := instance.GetProperty("EcnNumInjector")
	if err != nil {
		return
	}
	value, ok := retValue.(uint32)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetEcnPriorityEnable sets the value of EcnPriorityEnable for the instance
func (instance *MLNX_NetAdapterEcnSettingData) SetPropertyEcnPriorityEnable(value uint32) (err error) {
	return instance.SetProperty("EcnPriorityEnable", value)
}

// GetEcnPriorityEnable gets the value of EcnPriorityEnable for the instance
func (instance *MLNX_NetAdapterEcnSettingData) GetPropertyEcnPriorityEnable() (value uint32, err error) {
	retValue, err := instance.GetProperty("EcnPriorityEnable")
	if err != nil {
		return
	}
	value, ok := retValue.(uint32)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetEcnRateToSetOnFirstCnp sets the value of EcnRateToSetOnFirstCnp for the instance
func (instance *MLNX_NetAdapterEcnSettingData) SetPropertyEcnRateToSetOnFirstCnp(value uint32) (err error) {
	return instance.SetProperty("EcnRateToSetOnFirstCnp", value)
}

// GetEcnRateToSetOnFirstCnp gets the value of EcnRateToSetOnFirstCnp for the instance
func (instance *MLNX_NetAdapterEcnSettingData) GetPropertyEcnRateToSetOnFirstCnp() (value uint32, err error) {
	retValue, err := instance.GetProperty("EcnRateToSetOnFirstCnp")
	if err != nil {
		return
	}
	value, ok := retValue.(uint32)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetEcnRpgAiRate sets the value of EcnRpgAiRate for the instance
func (instance *MLNX_NetAdapterEcnSettingData) SetPropertyEcnRpgAiRate(value uint32) (err error) {
	return instance.SetProperty("EcnRpgAiRate", value)
}

// GetEcnRpgAiRate gets the value of EcnRpgAiRate for the instance
func (instance *MLNX_NetAdapterEcnSettingData) GetPropertyEcnRpgAiRate() (value uint32, err error) {
	retValue, err := instance.GetProperty("EcnRpgAiRate")
	if err != nil {
		return
	}
	value, ok := retValue.(uint32)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetEcnRpgByteReset sets the value of EcnRpgByteReset for the instance
func (instance *MLNX_NetAdapterEcnSettingData) SetPropertyEcnRpgByteReset(value uint32) (err error) {
	return instance.SetProperty("EcnRpgByteReset", value)
}

// GetEcnRpgByteReset gets the value of EcnRpgByteReset for the instance
func (instance *MLNX_NetAdapterEcnSettingData) GetPropertyEcnRpgByteReset() (value uint32, err error) {
	retValue, err := instance.GetProperty("EcnRpgByteReset")
	if err != nil {
		return
	}
	value, ok := retValue.(uint32)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetEcnRpgHaiRate sets the value of EcnRpgHaiRate for the instance
func (instance *MLNX_NetAdapterEcnSettingData) SetPropertyEcnRpgHaiRate(value uint32) (err error) {
	return instance.SetProperty("EcnRpgHaiRate", value)
}

// GetEcnRpgHaiRate gets the value of EcnRpgHaiRate for the instance
func (instance *MLNX_NetAdapterEcnSettingData) GetPropertyEcnRpgHaiRate() (value uint32, err error) {
	retValue, err := instance.GetProperty("EcnRpgHaiRate")
	if err != nil {
		return
	}
	value, ok := retValue.(uint32)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetEcnRpgMaxRate sets the value of EcnRpgMaxRate for the instance
func (instance *MLNX_NetAdapterEcnSettingData) SetPropertyEcnRpgMaxRate(value uint32) (err error) {
	return instance.SetProperty("EcnRpgMaxRate", value)
}

// GetEcnRpgMaxRate gets the value of EcnRpgMaxRate for the instance
func (instance *MLNX_NetAdapterEcnSettingData) GetPropertyEcnRpgMaxRate() (value uint32, err error) {
	retValue, err := instance.GetProperty("EcnRpgMaxRate")
	if err != nil {
		return
	}
	value, ok := retValue.(uint32)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetEcnRpgMinDecFac sets the value of EcnRpgMinDecFac for the instance
func (instance *MLNX_NetAdapterEcnSettingData) SetPropertyEcnRpgMinDecFac(value uint32) (err error) {
	return instance.SetProperty("EcnRpgMinDecFac", value)
}

// GetEcnRpgMinDecFac gets the value of EcnRpgMinDecFac for the instance
func (instance *MLNX_NetAdapterEcnSettingData) GetPropertyEcnRpgMinDecFac() (value uint32, err error) {
	retValue, err := instance.GetProperty("EcnRpgMinDecFac")
	if err != nil {
		return
	}
	value, ok := retValue.(uint32)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetEcnRpgMinRate sets the value of EcnRpgMinRate for the instance
func (instance *MLNX_NetAdapterEcnSettingData) SetPropertyEcnRpgMinRate(value uint32) (err error) {
	return instance.SetProperty("EcnRpgMinRate", value)
}

// GetEcnRpgMinRate gets the value of EcnRpgMinRate for the instance
func (instance *MLNX_NetAdapterEcnSettingData) GetPropertyEcnRpgMinRate() (value uint32, err error) {
	retValue, err := instance.GetProperty("EcnRpgMinRate")
	if err != nil {
		return
	}
	value, ok := retValue.(uint32)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetEcnRpgThreshold sets the value of EcnRpgThreshold for the instance
func (instance *MLNX_NetAdapterEcnSettingData) SetPropertyEcnRpgThreshold(value uint32) (err error) {
	return instance.SetProperty("EcnRpgThreshold", value)
}

// GetEcnRpgThreshold gets the value of EcnRpgThreshold for the instance
func (instance *MLNX_NetAdapterEcnSettingData) GetPropertyEcnRpgThreshold() (value uint32, err error) {
	retValue, err := instance.GetProperty("EcnRpgThreshold")
	if err != nil {
		return
	}
	value, ok := retValue.(uint32)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetEcnRpgTimeReset sets the value of EcnRpgTimeReset for the instance
func (instance *MLNX_NetAdapterEcnSettingData) SetPropertyEcnRpgTimeReset(value uint32) (err error) {
	return instance.SetProperty("EcnRpgTimeReset", value)
}

// GetEcnRpgTimeReset gets the value of EcnRpgTimeReset for the instance
func (instance *MLNX_NetAdapterEcnSettingData) GetPropertyEcnRpgTimeReset() (value uint32, err error) {
	retValue, err := instance.GetProperty("EcnRpgTimeReset")
	if err != nil {
		return
	}
	value, ok := retValue.(uint32)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetEcnSupportIBStandardCnp sets the value of EcnSupportIBStandardCnp for the instance
func (instance *MLNX_NetAdapterEcnSettingData) SetPropertyEcnSupportIBStandardCnp(value uint32) (err error) {
	return instance.SetProperty("EcnSupportIBStandardCnp", value)
}

// GetEcnSupportIBStandardCnp gets the value of EcnSupportIBStandardCnp for the instance
func (instance *MLNX_NetAdapterEcnSettingData) GetPropertyEcnSupportIBStandardCnp() (value uint32, err error) {
	retValue, err := instance.GetProperty("EcnSupportIBStandardCnp")
	if err != nil {
		return
	}
	value, ok := retValue.(uint32)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetEcnTrapDisablePendingCnpThreshold sets the value of EcnTrapDisablePendingCnpThreshold for the instance
func (instance *MLNX_NetAdapterEcnSettingData) SetPropertyEcnTrapDisablePendingCnpThreshold(value uint32) (err error) {
	return instance.SetProperty("EcnTrapDisablePendingCnpThreshold", value)
}

// GetEcnTrapDisablePendingCnpThreshold gets the value of EcnTrapDisablePendingCnpThreshold for the instance
func (instance *MLNX_NetAdapterEcnSettingData) GetPropertyEcnTrapDisablePendingCnpThreshold() (value uint32, err error) {
	retValue, err := instance.GetProperty("EcnTrapDisablePendingCnpThreshold")
	if err != nil {
		return
	}
	value, ok := retValue.(uint32)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetEcnTrapDisablePeriod sets the value of EcnTrapDisablePeriod for the instance
func (instance *MLNX_NetAdapterEcnSettingData) SetPropertyEcnTrapDisablePeriod(value uint32) (err error) {
	return instance.SetProperty("EcnTrapDisablePeriod", value)
}

// GetEcnTrapDisablePeriod gets the value of EcnTrapDisablePeriod for the instance
func (instance *MLNX_NetAdapterEcnSettingData) GetPropertyEcnTrapDisablePeriod() (value uint32, err error) {
	retValue, err := instance.GetProperty("EcnTrapDisablePeriod")
	if err != nil {
		return
	}
	value, ok := retValue.(uint32)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetEcnUseIBStandardCnp sets the value of EcnUseIBStandardCnp for the instance
func (instance *MLNX_NetAdapterEcnSettingData) SetPropertyEcnUseIBStandardCnp(value uint32) (err error) {
	return instance.SetProperty("EcnUseIBStandardCnp", value)
}

// GetEcnUseIBStandardCnp gets the value of EcnUseIBStandardCnp for the instance
func (instance *MLNX_NetAdapterEcnSettingData) GetPropertyEcnUseIBStandardCnp() (value uint32, err error) {
	retValue, err := instance.GetProperty("EcnUseIBStandardCnp")
	if err != nil {
		return
	}
	value, ok := retValue.(uint32)
	if !ok {
		// TODO: Set an error
	}
	return
}

//

// <param name="EcnAlphaToRateCoeff" type="uint32 "></param>
// <param name="EcnAlphaToRateShift" type="uint32 "></param>
// <param name="EcnBurstSize" type="uint32 "></param>
// <param name="EcnCatchRateLimitEn" type="uint32 "></param>
// <param name="EcnClampTgtRate" type="uint32 "></param>
// <param name="EcnClampTgtRateAfterTimeInc" type="uint32 "></param>
// <param name="EcnCnp802pPrio" type="uint32 "></param>
// <param name="EcnCnpDscp" type="uint32 "></param>
// <param name="EcnCnpRecEn" type="uint32 "></param>
// <param name="EcnCnpTimer" type="uint32 "></param>
// <param name="EcnCoalesceCnpInRp" type="uint32 "></param>
// <param name="EcnCompnRateLimit" type="uint32 "></param>
// <param name="EcnDceTcpG" type="uint32 "></param>
// <param name="EcnDceTcpRtt" type="uint32 "></param>
// <param name="EcnDceTcpRttDelay" type="uint32 "></param>
// <param name="EcnEnable" type="bool "></param>
// <param name="EcnExpectIpv6" type="uint32 "></param>
// <param name="EcnExpectVlanTagged" type="uint32 "></param>
// <param name="EcnFastRise" type="uint32 "></param>
// <param name="EcnForceRcTos" type="uint32 "></param>
// <param name="EcnForceUcTos" type="uint32 "></param>
// <param name="EcnInitialAlphaValue" type="uint32 "></param>
// <param name="EcnMarkedRatioMultiplier" type="uint32 "></param>
// <param name="EcnMarkedRatioShift" type="uint32 "></param>
// <param name="EcnMaxByteRise" type="uint32 "></param>
// <param name="EcnMaxTimeRise" type="uint32 "></param>
// <param name="EcnMinLosslessBufferForCatches" type="uint32 "></param>
// <param name="EcnMinLossyBufferForCatches" type="uint32 "></param>
// <param name="EcnNoCongestionCyclesToKeep" type="uint32 "></param>
// <param name="EcnNpRecEn" type="uint32 "></param>
// <param name="EcnNumInjector" type="uint32 "></param>
// <param name="EcnPriorityEnable" type="uint32 "></param>
// <param name="EcnRateToSetOnFirstCnp" type="uint32 "></param>
// <param name="EcnRpgAiRate" type="uint32 "></param>
// <param name="EcnRpgByteReset" type="uint32 "></param>
// <param name="EcnRpgHaiRate" type="uint32 "></param>
// <param name="EcnRpgMaxRate" type="uint32 "></param>
// <param name="EcnRpgMinDecFac" type="uint32 "></param>
// <param name="EcnRpgMinRate" type="uint32 "></param>
// <param name="EcnRpgThreshold" type="uint32 "></param>
// <param name="EcnRpgTimeReset" type="uint32 "></param>
// <param name="EcnSupportIBStandardCnp" type="uint32 "></param>
// <param name="EcnTrapDisablePendingCnpThreshold" type="uint32 "></param>
// <param name="EcnTrapDisablePeriod" type="uint32 "></param>
// <param name="EcnUseIBStandardCnp" type="uint32 "></param>

// <param name="ReturnValue" type="uint32 "></param>
func (instance *MLNX_NetAdapterEcnSettingData) SetValue( /* IN */ EcnExpectIpv6 uint32,
	/* IN */ EcnExpectVlanTagged uint32,
	/* IN */ EcnCatchRateLimitEn uint32,
	/* IN */ EcnTrapDisablePeriod uint32,
	/* IN */ EcnMinLosslessBufferForCatches uint32,
	/* IN */ EcnMinLossyBufferForCatches uint32,
	/* IN */ EcnTrapDisablePendingCnpThreshold uint32,
	/* IN */ EcnEnable bool,
	/* IN */ EcnForceRcTos uint32,
	/* IN */ EcnForceUcTos uint32,
	/* IN */ EcnCnpRecEn uint32,
	/* IN */ EcnNpRecEn uint32,
	/* IN */ EcnFastRise uint32,
	/* IN */ EcnClampTgtRate uint32,
	/* IN */ EcnSupportIBStandardCnp uint32,
	/* IN */ EcnCoalesceCnpInRp uint32,
	/* IN */ EcnClampTgtRateAfterTimeInc uint32,
	/* IN */ EcnRpgTimeReset uint32,
	/* IN */ EcnRpgByteReset uint32,
	/* IN */ EcnRpgThreshold uint32,
	/* IN */ EcnRpgMaxRate uint32,
	/* IN */ EcnRpgAiRate uint32,
	/* IN */ EcnRpgHaiRate uint32,
	/* IN */ EcnAlphaToRateShift uint32,
	/* IN */ EcnRpgMinDecFac uint32,
	/* IN */ EcnRpgMinRate uint32,
	/* IN */ EcnMaxTimeRise uint32,
	/* IN */ EcnMaxByteRise uint32,
	/* IN */ EcnAlphaToRateCoeff uint32,
	/* IN */ EcnMarkedRatioMultiplier uint32,
	/* IN */ EcnMarkedRatioShift uint32,
	/* IN */ EcnRateToSetOnFirstCnp uint32,
	/* IN */ EcnDceTcpG uint32,
	/* IN */ EcnDceTcpRtt uint32,
	/* IN */ EcnDceTcpRttDelay uint32,
	/* IN */ EcnInitialAlphaValue uint32,
	/* IN */ EcnUseIBStandardCnp uint32,
	/* IN */ EcnCompnRateLimit uint32,
	/* IN */ EcnNumInjector uint32,
	/* IN */ EcnCnpTimer uint32,
	/* IN */ EcnCnpDscp uint32,
	/* IN */ EcnCnp802pPrio uint32,
	/* IN */ EcnNoCongestionCyclesToKeep uint32,
	/* IN */ EcnPriorityEnable uint32,
	/* IN */ EcnBurstSize uint32) (result uint32, err error) {
	retVal, err := instance.InvokeMethodWithReturn("SetValue", EcnExpectIpv6, EcnExpectVlanTagged, EcnCatchRateLimitEn, EcnTrapDisablePeriod, EcnMinLosslessBufferForCatches, EcnMinLossyBufferForCatches, EcnTrapDisablePendingCnpThreshold, EcnEnable, EcnForceRcTos, EcnForceUcTos, EcnCnpRecEn, EcnNpRecEn, EcnFastRise, EcnClampTgtRate, EcnSupportIBStandardCnp, EcnCoalesceCnpInRp, EcnClampTgtRateAfterTimeInc, EcnRpgTimeReset, EcnRpgByteReset, EcnRpgThreshold, EcnRpgMaxRate, EcnRpgAiRate, EcnRpgHaiRate, EcnAlphaToRateShift, EcnRpgMinDecFac, EcnRpgMinRate, EcnMaxTimeRise, EcnMaxByteRise, EcnAlphaToRateCoeff, EcnMarkedRatioMultiplier, EcnMarkedRatioShift, EcnRateToSetOnFirstCnp, EcnDceTcpG, EcnDceTcpRtt, EcnDceTcpRttDelay, EcnInitialAlphaValue, EcnUseIBStandardCnp, EcnCompnRateLimit, EcnNumInjector, EcnCnpTimer, EcnCnpDscp, EcnCnp802pPrio, EcnNoCongestionCyclesToKeep, EcnPriorityEnable, EcnBurstSize)
	if err != nil {
		return
	}
	result = uint32(retVal)
	return

}

//

// <param name="ReturnValue" type="uint8 "></param>
func (instance *MLNX_NetAdapterEcnSettingData) Disable() (result uint8, err error) {
	retVal, err := instance.InvokeMethodWithReturn("Disable")
	if err != nil {
		return
	}
	result = uint8(retVal)
	return

}

//

// <param name="ReturnValue" type="uint8 "></param>
func (instance *MLNX_NetAdapterEcnSettingData) Enable() (result uint8, err error) {
	retVal, err := instance.InvokeMethodWithReturn("Enable")
	if err != nil {
		return
	}
	result = uint8(retVal)
	return

}
