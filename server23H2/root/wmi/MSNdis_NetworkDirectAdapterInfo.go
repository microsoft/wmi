// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 6/12/2024 using wmigen
//      Source root.WMI
//////////////////////////////////////////////
package wmi

import (
	"github.com/microsoft/wmi/pkg/base/query"
	"github.com/microsoft/wmi/pkg/errors"
	cim "github.com/microsoft/wmi/pkg/wmiinstance"
	"reflect"
)

// MSNdis_NetworkDirectAdapterInfo struct
type MSNdis_NetworkDirectAdapterInfo struct {
	*MSNdis

	//
	DeviceId uint32

	//
	FRMRPageCount uint32

	//
	InOrderDMA bool

	//
	LargeRequestThreshold uint32

	//
	MaxCalleeData uint32

	//
	MaxCallerData uint32

	//
	MaxCqDepth uint32

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
	MaxSrqDepth uint32

	//
	MaxTransferLength uint32

	//
	MaxWindowSize uint64

	//
	SupportsCQResize bool

	//
	SupportsLoopbackConnections bool

	//
	VendorId uint32

	//
	Version MSNdis_NetworkDirectVersion
}

func NewMSNdis_NetworkDirectAdapterInfoEx1(instance *cim.WmiInstance) (newInstance *MSNdis_NetworkDirectAdapterInfo, err error) {
	tmp, err := NewMSNdisEx1(instance)

	if err != nil {
		return
	}
	newInstance = &MSNdis_NetworkDirectAdapterInfo{
		MSNdis: tmp,
	}
	return
}

func NewMSNdis_NetworkDirectAdapterInfoEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *MSNdis_NetworkDirectAdapterInfo, err error) {
	tmp, err := NewMSNdisEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &MSNdis_NetworkDirectAdapterInfo{
		MSNdis: tmp,
	}
	return
}

// SetDeviceId sets the value of DeviceId for the instance
func (instance *MSNdis_NetworkDirectAdapterInfo) SetPropertyDeviceId(value uint32) (err error) {
	return instance.SetProperty("DeviceId", (value))
}

// GetDeviceId gets the value of DeviceId for the instance
func (instance *MSNdis_NetworkDirectAdapterInfo) GetPropertyDeviceId() (value uint32, err error) {
	retValue, err := instance.GetProperty("DeviceId")
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

// SetFRMRPageCount sets the value of FRMRPageCount for the instance
func (instance *MSNdis_NetworkDirectAdapterInfo) SetPropertyFRMRPageCount(value uint32) (err error) {
	return instance.SetProperty("FRMRPageCount", (value))
}

// GetFRMRPageCount gets the value of FRMRPageCount for the instance
func (instance *MSNdis_NetworkDirectAdapterInfo) GetPropertyFRMRPageCount() (value uint32, err error) {
	retValue, err := instance.GetProperty("FRMRPageCount")
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

// SetInOrderDMA sets the value of InOrderDMA for the instance
func (instance *MSNdis_NetworkDirectAdapterInfo) SetPropertyInOrderDMA(value bool) (err error) {
	return instance.SetProperty("InOrderDMA", (value))
}

// GetInOrderDMA gets the value of InOrderDMA for the instance
func (instance *MSNdis_NetworkDirectAdapterInfo) GetPropertyInOrderDMA() (value bool, err error) {
	retValue, err := instance.GetProperty("InOrderDMA")
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

// SetLargeRequestThreshold sets the value of LargeRequestThreshold for the instance
func (instance *MSNdis_NetworkDirectAdapterInfo) SetPropertyLargeRequestThreshold(value uint32) (err error) {
	return instance.SetProperty("LargeRequestThreshold", (value))
}

// GetLargeRequestThreshold gets the value of LargeRequestThreshold for the instance
func (instance *MSNdis_NetworkDirectAdapterInfo) GetPropertyLargeRequestThreshold() (value uint32, err error) {
	retValue, err := instance.GetProperty("LargeRequestThreshold")
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

// SetMaxCalleeData sets the value of MaxCalleeData for the instance
func (instance *MSNdis_NetworkDirectAdapterInfo) SetPropertyMaxCalleeData(value uint32) (err error) {
	return instance.SetProperty("MaxCalleeData", (value))
}

// GetMaxCalleeData gets the value of MaxCalleeData for the instance
func (instance *MSNdis_NetworkDirectAdapterInfo) GetPropertyMaxCalleeData() (value uint32, err error) {
	retValue, err := instance.GetProperty("MaxCalleeData")
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

// SetMaxCallerData sets the value of MaxCallerData for the instance
func (instance *MSNdis_NetworkDirectAdapterInfo) SetPropertyMaxCallerData(value uint32) (err error) {
	return instance.SetProperty("MaxCallerData", (value))
}

// GetMaxCallerData gets the value of MaxCallerData for the instance
func (instance *MSNdis_NetworkDirectAdapterInfo) GetPropertyMaxCallerData() (value uint32, err error) {
	retValue, err := instance.GetProperty("MaxCallerData")
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

// SetMaxCqDepth sets the value of MaxCqDepth for the instance
func (instance *MSNdis_NetworkDirectAdapterInfo) SetPropertyMaxCqDepth(value uint32) (err error) {
	return instance.SetProperty("MaxCqDepth", (value))
}

// GetMaxCqDepth gets the value of MaxCqDepth for the instance
func (instance *MSNdis_NetworkDirectAdapterInfo) GetPropertyMaxCqDepth() (value uint32, err error) {
	retValue, err := instance.GetProperty("MaxCqDepth")
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

// SetMaxInboundReadLimit sets the value of MaxInboundReadLimit for the instance
func (instance *MSNdis_NetworkDirectAdapterInfo) SetPropertyMaxInboundReadLimit(value uint32) (err error) {
	return instance.SetProperty("MaxInboundReadLimit", (value))
}

// GetMaxInboundReadLimit gets the value of MaxInboundReadLimit for the instance
func (instance *MSNdis_NetworkDirectAdapterInfo) GetPropertyMaxInboundReadLimit() (value uint32, err error) {
	retValue, err := instance.GetProperty("MaxInboundReadLimit")
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

// SetMaxInitiatorQueueDepth sets the value of MaxInitiatorQueueDepth for the instance
func (instance *MSNdis_NetworkDirectAdapterInfo) SetPropertyMaxInitiatorQueueDepth(value uint32) (err error) {
	return instance.SetProperty("MaxInitiatorQueueDepth", (value))
}

// GetMaxInitiatorQueueDepth gets the value of MaxInitiatorQueueDepth for the instance
func (instance *MSNdis_NetworkDirectAdapterInfo) GetPropertyMaxInitiatorQueueDepth() (value uint32, err error) {
	retValue, err := instance.GetProperty("MaxInitiatorQueueDepth")
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

// SetMaxInitiatorRequestSge sets the value of MaxInitiatorRequestSge for the instance
func (instance *MSNdis_NetworkDirectAdapterInfo) SetPropertyMaxInitiatorRequestSge(value uint32) (err error) {
	return instance.SetProperty("MaxInitiatorRequestSge", (value))
}

// GetMaxInitiatorRequestSge gets the value of MaxInitiatorRequestSge for the instance
func (instance *MSNdis_NetworkDirectAdapterInfo) GetPropertyMaxInitiatorRequestSge() (value uint32, err error) {
	retValue, err := instance.GetProperty("MaxInitiatorRequestSge")
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

// SetMaxInlineDataSize sets the value of MaxInlineDataSize for the instance
func (instance *MSNdis_NetworkDirectAdapterInfo) SetPropertyMaxInlineDataSize(value uint32) (err error) {
	return instance.SetProperty("MaxInlineDataSize", (value))
}

// GetMaxInlineDataSize gets the value of MaxInlineDataSize for the instance
func (instance *MSNdis_NetworkDirectAdapterInfo) GetPropertyMaxInlineDataSize() (value uint32, err error) {
	retValue, err := instance.GetProperty("MaxInlineDataSize")
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

// SetMaxOutboundReadLimit sets the value of MaxOutboundReadLimit for the instance
func (instance *MSNdis_NetworkDirectAdapterInfo) SetPropertyMaxOutboundReadLimit(value uint32) (err error) {
	return instance.SetProperty("MaxOutboundReadLimit", (value))
}

// GetMaxOutboundReadLimit gets the value of MaxOutboundReadLimit for the instance
func (instance *MSNdis_NetworkDirectAdapterInfo) GetPropertyMaxOutboundReadLimit() (value uint32, err error) {
	retValue, err := instance.GetProperty("MaxOutboundReadLimit")
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

// SetMaxReadRequestSge sets the value of MaxReadRequestSge for the instance
func (instance *MSNdis_NetworkDirectAdapterInfo) SetPropertyMaxReadRequestSge(value uint32) (err error) {
	return instance.SetProperty("MaxReadRequestSge", (value))
}

// GetMaxReadRequestSge gets the value of MaxReadRequestSge for the instance
func (instance *MSNdis_NetworkDirectAdapterInfo) GetPropertyMaxReadRequestSge() (value uint32, err error) {
	retValue, err := instance.GetProperty("MaxReadRequestSge")
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

// SetMaxReceiveQueueDepth sets the value of MaxReceiveQueueDepth for the instance
func (instance *MSNdis_NetworkDirectAdapterInfo) SetPropertyMaxReceiveQueueDepth(value uint32) (err error) {
	return instance.SetProperty("MaxReceiveQueueDepth", (value))
}

// GetMaxReceiveQueueDepth gets the value of MaxReceiveQueueDepth for the instance
func (instance *MSNdis_NetworkDirectAdapterInfo) GetPropertyMaxReceiveQueueDepth() (value uint32, err error) {
	retValue, err := instance.GetProperty("MaxReceiveQueueDepth")
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

// SetMaxReceiveRequestSge sets the value of MaxReceiveRequestSge for the instance
func (instance *MSNdis_NetworkDirectAdapterInfo) SetPropertyMaxReceiveRequestSge(value uint32) (err error) {
	return instance.SetProperty("MaxReceiveRequestSge", (value))
}

// GetMaxReceiveRequestSge gets the value of MaxReceiveRequestSge for the instance
func (instance *MSNdis_NetworkDirectAdapterInfo) GetPropertyMaxReceiveRequestSge() (value uint32, err error) {
	retValue, err := instance.GetProperty("MaxReceiveRequestSge")
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

// SetMaxRegistrationSize sets the value of MaxRegistrationSize for the instance
func (instance *MSNdis_NetworkDirectAdapterInfo) SetPropertyMaxRegistrationSize(value uint64) (err error) {
	return instance.SetProperty("MaxRegistrationSize", (value))
}

// GetMaxRegistrationSize gets the value of MaxRegistrationSize for the instance
func (instance *MSNdis_NetworkDirectAdapterInfo) GetPropertyMaxRegistrationSize() (value uint64, err error) {
	retValue, err := instance.GetProperty("MaxRegistrationSize")
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

// SetMaxSrqDepth sets the value of MaxSrqDepth for the instance
func (instance *MSNdis_NetworkDirectAdapterInfo) SetPropertyMaxSrqDepth(value uint32) (err error) {
	return instance.SetProperty("MaxSrqDepth", (value))
}

// GetMaxSrqDepth gets the value of MaxSrqDepth for the instance
func (instance *MSNdis_NetworkDirectAdapterInfo) GetPropertyMaxSrqDepth() (value uint32, err error) {
	retValue, err := instance.GetProperty("MaxSrqDepth")
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

// SetMaxTransferLength sets the value of MaxTransferLength for the instance
func (instance *MSNdis_NetworkDirectAdapterInfo) SetPropertyMaxTransferLength(value uint32) (err error) {
	return instance.SetProperty("MaxTransferLength", (value))
}

// GetMaxTransferLength gets the value of MaxTransferLength for the instance
func (instance *MSNdis_NetworkDirectAdapterInfo) GetPropertyMaxTransferLength() (value uint32, err error) {
	retValue, err := instance.GetProperty("MaxTransferLength")
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

// SetMaxWindowSize sets the value of MaxWindowSize for the instance
func (instance *MSNdis_NetworkDirectAdapterInfo) SetPropertyMaxWindowSize(value uint64) (err error) {
	return instance.SetProperty("MaxWindowSize", (value))
}

// GetMaxWindowSize gets the value of MaxWindowSize for the instance
func (instance *MSNdis_NetworkDirectAdapterInfo) GetPropertyMaxWindowSize() (value uint64, err error) {
	retValue, err := instance.GetProperty("MaxWindowSize")
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

// SetSupportsCQResize sets the value of SupportsCQResize for the instance
func (instance *MSNdis_NetworkDirectAdapterInfo) SetPropertySupportsCQResize(value bool) (err error) {
	return instance.SetProperty("SupportsCQResize", (value))
}

// GetSupportsCQResize gets the value of SupportsCQResize for the instance
func (instance *MSNdis_NetworkDirectAdapterInfo) GetPropertySupportsCQResize() (value bool, err error) {
	retValue, err := instance.GetProperty("SupportsCQResize")
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

// SetSupportsLoopbackConnections sets the value of SupportsLoopbackConnections for the instance
func (instance *MSNdis_NetworkDirectAdapterInfo) SetPropertySupportsLoopbackConnections(value bool) (err error) {
	return instance.SetProperty("SupportsLoopbackConnections", (value))
}

// GetSupportsLoopbackConnections gets the value of SupportsLoopbackConnections for the instance
func (instance *MSNdis_NetworkDirectAdapterInfo) GetPropertySupportsLoopbackConnections() (value bool, err error) {
	retValue, err := instance.GetProperty("SupportsLoopbackConnections")
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

// SetVendorId sets the value of VendorId for the instance
func (instance *MSNdis_NetworkDirectAdapterInfo) SetPropertyVendorId(value uint32) (err error) {
	return instance.SetProperty("VendorId", (value))
}

// GetVendorId gets the value of VendorId for the instance
func (instance *MSNdis_NetworkDirectAdapterInfo) GetPropertyVendorId() (value uint32, err error) {
	retValue, err := instance.GetProperty("VendorId")
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

// SetVersion sets the value of Version for the instance
func (instance *MSNdis_NetworkDirectAdapterInfo) SetPropertyVersion(value MSNdis_NetworkDirectVersion) (err error) {
	return instance.SetProperty("Version", (value))
}

// GetVersion gets the value of Version for the instance
func (instance *MSNdis_NetworkDirectAdapterInfo) GetPropertyVersion() (value MSNdis_NetworkDirectVersion, err error) {
	retValue, err := instance.GetProperty("Version")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(MSNdis_NetworkDirectVersion)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " MSNdis_NetworkDirectVersion is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = MSNdis_NetworkDirectVersion(valuetmp)

	return
}
