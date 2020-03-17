// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 3/16/2020 using wmigen
//      Source root.StandardCimv2
//////////////////////////////////////////////
package standardcimv2

import (
	"github.com/microsoft/wmi/pkg/wmiinstance"
)

// MSFT_NetAdapter_RdmaAdapterInfo struct
type MSFT_NetAdapter_RdmaAdapterInfo struct {
	cim.WmiInstance

	//
	DeviceId uint32

	//
	FRMRPageCount uint32

	//
	InOrderDMA bool

	//
	LargeRequestThreshold uint32

	//
	MajorVersionNumber uint16

	//
	MaxCalleeData uint32

	//
	MaxCallerData uint32

	//
	MaxCompletionQueueDepth uint32

	//
	MaxInboundReadLimit uint32

	//
	MaxInitiatorQueueDepth uint32

	//
	MaxInitiatorRequestSge uint32

	//
	MaxInlineDataSize uint32

	//
	MaxOutboundReadLimit uint32

	//
	MaxReadRequestSge uint32

	//
	MaxReceiveQueueDepth uint32

	//
	MaxReceiveRequestSge uint32

	//
	MaxRegistrationSize uint64

	//
	MaxSharedReceiveQueueDepth uint32

	//
	MaxTransferLength uint32

	//
	MaxWindowSize uint64

	//
	MinorVersionNumber uint16

	//
	RdmaReadSinkFlagNotRequired bool

	//
	RdmaTechnology uint32

	//
	SupportsCompletionQueueInterruptModeration bool

	//
	SupportsCompletionQueueResize bool

	//
	SupportsLoopbackConnections bool

	//
	SupportsMultiEngine bool

	//
	VendorId uint32
}

// SetDeviceId sets the value of DeviceId for the instance
func (instance *MSFT_NetAdapter_RdmaAdapterInfo) SetPropertyDeviceId(value uint32) (err error) {
	return instance.SetProperty("DeviceId", value)
}

// GetDeviceId gets the value of DeviceId for the instance
func (instance *MSFT_NetAdapter_RdmaAdapterInfo) GetPropertyDeviceId() (value uint32, err error) {
	retValue, err := instance.GetProperty("DeviceId")
	if err != nil {
		return
	}
	value, ok := retValue.(uint32)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetFRMRPageCount sets the value of FRMRPageCount for the instance
func (instance *MSFT_NetAdapter_RdmaAdapterInfo) SetPropertyFRMRPageCount(value uint32) (err error) {
	return instance.SetProperty("FRMRPageCount", value)
}

// GetFRMRPageCount gets the value of FRMRPageCount for the instance
func (instance *MSFT_NetAdapter_RdmaAdapterInfo) GetPropertyFRMRPageCount() (value uint32, err error) {
	retValue, err := instance.GetProperty("FRMRPageCount")
	if err != nil {
		return
	}
	value, ok := retValue.(uint32)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetInOrderDMA sets the value of InOrderDMA for the instance
func (instance *MSFT_NetAdapter_RdmaAdapterInfo) SetPropertyInOrderDMA(value bool) (err error) {
	return instance.SetProperty("InOrderDMA", value)
}

// GetInOrderDMA gets the value of InOrderDMA for the instance
func (instance *MSFT_NetAdapter_RdmaAdapterInfo) GetPropertyInOrderDMA() (value bool, err error) {
	retValue, err := instance.GetProperty("InOrderDMA")
	if err != nil {
		return
	}
	value, ok := retValue.(bool)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetLargeRequestThreshold sets the value of LargeRequestThreshold for the instance
func (instance *MSFT_NetAdapter_RdmaAdapterInfo) SetPropertyLargeRequestThreshold(value uint32) (err error) {
	return instance.SetProperty("LargeRequestThreshold", value)
}

// GetLargeRequestThreshold gets the value of LargeRequestThreshold for the instance
func (instance *MSFT_NetAdapter_RdmaAdapterInfo) GetPropertyLargeRequestThreshold() (value uint32, err error) {
	retValue, err := instance.GetProperty("LargeRequestThreshold")
	if err != nil {
		return
	}
	value, ok := retValue.(uint32)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetMajorVersionNumber sets the value of MajorVersionNumber for the instance
func (instance *MSFT_NetAdapter_RdmaAdapterInfo) SetPropertyMajorVersionNumber(value uint16) (err error) {
	return instance.SetProperty("MajorVersionNumber", value)
}

// GetMajorVersionNumber gets the value of MajorVersionNumber for the instance
func (instance *MSFT_NetAdapter_RdmaAdapterInfo) GetPropertyMajorVersionNumber() (value uint16, err error) {
	retValue, err := instance.GetProperty("MajorVersionNumber")
	if err != nil {
		return
	}
	value, ok := retValue.(uint16)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetMaxCalleeData sets the value of MaxCalleeData for the instance
func (instance *MSFT_NetAdapter_RdmaAdapterInfo) SetPropertyMaxCalleeData(value uint32) (err error) {
	return instance.SetProperty("MaxCalleeData", value)
}

// GetMaxCalleeData gets the value of MaxCalleeData for the instance
func (instance *MSFT_NetAdapter_RdmaAdapterInfo) GetPropertyMaxCalleeData() (value uint32, err error) {
	retValue, err := instance.GetProperty("MaxCalleeData")
	if err != nil {
		return
	}
	value, ok := retValue.(uint32)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetMaxCallerData sets the value of MaxCallerData for the instance
func (instance *MSFT_NetAdapter_RdmaAdapterInfo) SetPropertyMaxCallerData(value uint32) (err error) {
	return instance.SetProperty("MaxCallerData", value)
}

// GetMaxCallerData gets the value of MaxCallerData for the instance
func (instance *MSFT_NetAdapter_RdmaAdapterInfo) GetPropertyMaxCallerData() (value uint32, err error) {
	retValue, err := instance.GetProperty("MaxCallerData")
	if err != nil {
		return
	}
	value, ok := retValue.(uint32)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetMaxCompletionQueueDepth sets the value of MaxCompletionQueueDepth for the instance
func (instance *MSFT_NetAdapter_RdmaAdapterInfo) SetPropertyMaxCompletionQueueDepth(value uint32) (err error) {
	return instance.SetProperty("MaxCompletionQueueDepth", value)
}

// GetMaxCompletionQueueDepth gets the value of MaxCompletionQueueDepth for the instance
func (instance *MSFT_NetAdapter_RdmaAdapterInfo) GetPropertyMaxCompletionQueueDepth() (value uint32, err error) {
	retValue, err := instance.GetProperty("MaxCompletionQueueDepth")
	if err != nil {
		return
	}
	value, ok := retValue.(uint32)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetMaxInboundReadLimit sets the value of MaxInboundReadLimit for the instance
func (instance *MSFT_NetAdapter_RdmaAdapterInfo) SetPropertyMaxInboundReadLimit(value uint32) (err error) {
	return instance.SetProperty("MaxInboundReadLimit", value)
}

// GetMaxInboundReadLimit gets the value of MaxInboundReadLimit for the instance
func (instance *MSFT_NetAdapter_RdmaAdapterInfo) GetPropertyMaxInboundReadLimit() (value uint32, err error) {
	retValue, err := instance.GetProperty("MaxInboundReadLimit")
	if err != nil {
		return
	}
	value, ok := retValue.(uint32)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetMaxInitiatorQueueDepth sets the value of MaxInitiatorQueueDepth for the instance
func (instance *MSFT_NetAdapter_RdmaAdapterInfo) SetPropertyMaxInitiatorQueueDepth(value uint32) (err error) {
	return instance.SetProperty("MaxInitiatorQueueDepth", value)
}

// GetMaxInitiatorQueueDepth gets the value of MaxInitiatorQueueDepth for the instance
func (instance *MSFT_NetAdapter_RdmaAdapterInfo) GetPropertyMaxInitiatorQueueDepth() (value uint32, err error) {
	retValue, err := instance.GetProperty("MaxInitiatorQueueDepth")
	if err != nil {
		return
	}
	value, ok := retValue.(uint32)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetMaxInitiatorRequestSge sets the value of MaxInitiatorRequestSge for the instance
func (instance *MSFT_NetAdapter_RdmaAdapterInfo) SetPropertyMaxInitiatorRequestSge(value uint32) (err error) {
	return instance.SetProperty("MaxInitiatorRequestSge", value)
}

// GetMaxInitiatorRequestSge gets the value of MaxInitiatorRequestSge for the instance
func (instance *MSFT_NetAdapter_RdmaAdapterInfo) GetPropertyMaxInitiatorRequestSge() (value uint32, err error) {
	retValue, err := instance.GetProperty("MaxInitiatorRequestSge")
	if err != nil {
		return
	}
	value, ok := retValue.(uint32)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetMaxInlineDataSize sets the value of MaxInlineDataSize for the instance
func (instance *MSFT_NetAdapter_RdmaAdapterInfo) SetPropertyMaxInlineDataSize(value uint32) (err error) {
	return instance.SetProperty("MaxInlineDataSize", value)
}

// GetMaxInlineDataSize gets the value of MaxInlineDataSize for the instance
func (instance *MSFT_NetAdapter_RdmaAdapterInfo) GetPropertyMaxInlineDataSize() (value uint32, err error) {
	retValue, err := instance.GetProperty("MaxInlineDataSize")
	if err != nil {
		return
	}
	value, ok := retValue.(uint32)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetMaxOutboundReadLimit sets the value of MaxOutboundReadLimit for the instance
func (instance *MSFT_NetAdapter_RdmaAdapterInfo) SetPropertyMaxOutboundReadLimit(value uint32) (err error) {
	return instance.SetProperty("MaxOutboundReadLimit", value)
}

// GetMaxOutboundReadLimit gets the value of MaxOutboundReadLimit for the instance
func (instance *MSFT_NetAdapter_RdmaAdapterInfo) GetPropertyMaxOutboundReadLimit() (value uint32, err error) {
	retValue, err := instance.GetProperty("MaxOutboundReadLimit")
	if err != nil {
		return
	}
	value, ok := retValue.(uint32)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetMaxReadRequestSge sets the value of MaxReadRequestSge for the instance
func (instance *MSFT_NetAdapter_RdmaAdapterInfo) SetPropertyMaxReadRequestSge(value uint32) (err error) {
	return instance.SetProperty("MaxReadRequestSge", value)
}

// GetMaxReadRequestSge gets the value of MaxReadRequestSge for the instance
func (instance *MSFT_NetAdapter_RdmaAdapterInfo) GetPropertyMaxReadRequestSge() (value uint32, err error) {
	retValue, err := instance.GetProperty("MaxReadRequestSge")
	if err != nil {
		return
	}
	value, ok := retValue.(uint32)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetMaxReceiveQueueDepth sets the value of MaxReceiveQueueDepth for the instance
func (instance *MSFT_NetAdapter_RdmaAdapterInfo) SetPropertyMaxReceiveQueueDepth(value uint32) (err error) {
	return instance.SetProperty("MaxReceiveQueueDepth", value)
}

// GetMaxReceiveQueueDepth gets the value of MaxReceiveQueueDepth for the instance
func (instance *MSFT_NetAdapter_RdmaAdapterInfo) GetPropertyMaxReceiveQueueDepth() (value uint32, err error) {
	retValue, err := instance.GetProperty("MaxReceiveQueueDepth")
	if err != nil {
		return
	}
	value, ok := retValue.(uint32)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetMaxReceiveRequestSge sets the value of MaxReceiveRequestSge for the instance
func (instance *MSFT_NetAdapter_RdmaAdapterInfo) SetPropertyMaxReceiveRequestSge(value uint32) (err error) {
	return instance.SetProperty("MaxReceiveRequestSge", value)
}

// GetMaxReceiveRequestSge gets the value of MaxReceiveRequestSge for the instance
func (instance *MSFT_NetAdapter_RdmaAdapterInfo) GetPropertyMaxReceiveRequestSge() (value uint32, err error) {
	retValue, err := instance.GetProperty("MaxReceiveRequestSge")
	if err != nil {
		return
	}
	value, ok := retValue.(uint32)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetMaxRegistrationSize sets the value of MaxRegistrationSize for the instance
func (instance *MSFT_NetAdapter_RdmaAdapterInfo) SetPropertyMaxRegistrationSize(value uint64) (err error) {
	return instance.SetProperty("MaxRegistrationSize", value)
}

// GetMaxRegistrationSize gets the value of MaxRegistrationSize for the instance
func (instance *MSFT_NetAdapter_RdmaAdapterInfo) GetPropertyMaxRegistrationSize() (value uint64, err error) {
	retValue, err := instance.GetProperty("MaxRegistrationSize")
	if err != nil {
		return
	}
	value, ok := retValue.(uint64)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetMaxSharedReceiveQueueDepth sets the value of MaxSharedReceiveQueueDepth for the instance
func (instance *MSFT_NetAdapter_RdmaAdapterInfo) SetPropertyMaxSharedReceiveQueueDepth(value uint32) (err error) {
	return instance.SetProperty("MaxSharedReceiveQueueDepth", value)
}

// GetMaxSharedReceiveQueueDepth gets the value of MaxSharedReceiveQueueDepth for the instance
func (instance *MSFT_NetAdapter_RdmaAdapterInfo) GetPropertyMaxSharedReceiveQueueDepth() (value uint32, err error) {
	retValue, err := instance.GetProperty("MaxSharedReceiveQueueDepth")
	if err != nil {
		return
	}
	value, ok := retValue.(uint32)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetMaxTransferLength sets the value of MaxTransferLength for the instance
func (instance *MSFT_NetAdapter_RdmaAdapterInfo) SetPropertyMaxTransferLength(value uint32) (err error) {
	return instance.SetProperty("MaxTransferLength", value)
}

// GetMaxTransferLength gets the value of MaxTransferLength for the instance
func (instance *MSFT_NetAdapter_RdmaAdapterInfo) GetPropertyMaxTransferLength() (value uint32, err error) {
	retValue, err := instance.GetProperty("MaxTransferLength")
	if err != nil {
		return
	}
	value, ok := retValue.(uint32)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetMaxWindowSize sets the value of MaxWindowSize for the instance
func (instance *MSFT_NetAdapter_RdmaAdapterInfo) SetPropertyMaxWindowSize(value uint64) (err error) {
	return instance.SetProperty("MaxWindowSize", value)
}

// GetMaxWindowSize gets the value of MaxWindowSize for the instance
func (instance *MSFT_NetAdapter_RdmaAdapterInfo) GetPropertyMaxWindowSize() (value uint64, err error) {
	retValue, err := instance.GetProperty("MaxWindowSize")
	if err != nil {
		return
	}
	value, ok := retValue.(uint64)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetMinorVersionNumber sets the value of MinorVersionNumber for the instance
func (instance *MSFT_NetAdapter_RdmaAdapterInfo) SetPropertyMinorVersionNumber(value uint16) (err error) {
	return instance.SetProperty("MinorVersionNumber", value)
}

// GetMinorVersionNumber gets the value of MinorVersionNumber for the instance
func (instance *MSFT_NetAdapter_RdmaAdapterInfo) GetPropertyMinorVersionNumber() (value uint16, err error) {
	retValue, err := instance.GetProperty("MinorVersionNumber")
	if err != nil {
		return
	}
	value, ok := retValue.(uint16)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetRdmaReadSinkFlagNotRequired sets the value of RdmaReadSinkFlagNotRequired for the instance
func (instance *MSFT_NetAdapter_RdmaAdapterInfo) SetPropertyRdmaReadSinkFlagNotRequired(value bool) (err error) {
	return instance.SetProperty("RdmaReadSinkFlagNotRequired", value)
}

// GetRdmaReadSinkFlagNotRequired gets the value of RdmaReadSinkFlagNotRequired for the instance
func (instance *MSFT_NetAdapter_RdmaAdapterInfo) GetPropertyRdmaReadSinkFlagNotRequired() (value bool, err error) {
	retValue, err := instance.GetProperty("RdmaReadSinkFlagNotRequired")
	if err != nil {
		return
	}
	value, ok := retValue.(bool)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetRdmaTechnology sets the value of RdmaTechnology for the instance
func (instance *MSFT_NetAdapter_RdmaAdapterInfo) SetPropertyRdmaTechnology(value uint32) (err error) {
	return instance.SetProperty("RdmaTechnology", value)
}

// GetRdmaTechnology gets the value of RdmaTechnology for the instance
func (instance *MSFT_NetAdapter_RdmaAdapterInfo) GetPropertyRdmaTechnology() (value uint32, err error) {
	retValue, err := instance.GetProperty("RdmaTechnology")
	if err != nil {
		return
	}
	value, ok := retValue.(uint32)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetSupportsCompletionQueueInterruptModeration sets the value of SupportsCompletionQueueInterruptModeration for the instance
func (instance *MSFT_NetAdapter_RdmaAdapterInfo) SetPropertySupportsCompletionQueueInterruptModeration(value bool) (err error) {
	return instance.SetProperty("SupportsCompletionQueueInterruptModeration", value)
}

// GetSupportsCompletionQueueInterruptModeration gets the value of SupportsCompletionQueueInterruptModeration for the instance
func (instance *MSFT_NetAdapter_RdmaAdapterInfo) GetPropertySupportsCompletionQueueInterruptModeration() (value bool, err error) {
	retValue, err := instance.GetProperty("SupportsCompletionQueueInterruptModeration")
	if err != nil {
		return
	}
	value, ok := retValue.(bool)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetSupportsCompletionQueueResize sets the value of SupportsCompletionQueueResize for the instance
func (instance *MSFT_NetAdapter_RdmaAdapterInfo) SetPropertySupportsCompletionQueueResize(value bool) (err error) {
	return instance.SetProperty("SupportsCompletionQueueResize", value)
}

// GetSupportsCompletionQueueResize gets the value of SupportsCompletionQueueResize for the instance
func (instance *MSFT_NetAdapter_RdmaAdapterInfo) GetPropertySupportsCompletionQueueResize() (value bool, err error) {
	retValue, err := instance.GetProperty("SupportsCompletionQueueResize")
	if err != nil {
		return
	}
	value, ok := retValue.(bool)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetSupportsLoopbackConnections sets the value of SupportsLoopbackConnections for the instance
func (instance *MSFT_NetAdapter_RdmaAdapterInfo) SetPropertySupportsLoopbackConnections(value bool) (err error) {
	return instance.SetProperty("SupportsLoopbackConnections", value)
}

// GetSupportsLoopbackConnections gets the value of SupportsLoopbackConnections for the instance
func (instance *MSFT_NetAdapter_RdmaAdapterInfo) GetPropertySupportsLoopbackConnections() (value bool, err error) {
	retValue, err := instance.GetProperty("SupportsLoopbackConnections")
	if err != nil {
		return
	}
	value, ok := retValue.(bool)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetSupportsMultiEngine sets the value of SupportsMultiEngine for the instance
func (instance *MSFT_NetAdapter_RdmaAdapterInfo) SetPropertySupportsMultiEngine(value bool) (err error) {
	return instance.SetProperty("SupportsMultiEngine", value)
}

// GetSupportsMultiEngine gets the value of SupportsMultiEngine for the instance
func (instance *MSFT_NetAdapter_RdmaAdapterInfo) GetPropertySupportsMultiEngine() (value bool, err error) {
	retValue, err := instance.GetProperty("SupportsMultiEngine")
	if err != nil {
		return
	}
	value, ok := retValue.(bool)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetVendorId sets the value of VendorId for the instance
func (instance *MSFT_NetAdapter_RdmaAdapterInfo) SetPropertyVendorId(value uint32) (err error) {
	return instance.SetProperty("VendorId", value)
}

// GetVendorId gets the value of VendorId for the instance
func (instance *MSFT_NetAdapter_RdmaAdapterInfo) GetPropertyVendorId() (value uint32, err error) {
	retValue, err := instance.GetProperty("VendorId")
	if err != nil {
		return
	}
	value, ok := retValue.(uint32)
	if !ok {
		// TODO: Set an error
	}
	return
}
