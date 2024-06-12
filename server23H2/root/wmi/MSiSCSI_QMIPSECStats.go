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

// MSiSCSI_QMIPSECStats struct
type MSiSCSI_QMIPSECStats struct {
	*Win32_PerfRawData

	//
	Active bool

	//
	ActiveSA uint64

	//
	ActiveTunnels uint64

	//
	AuthenticatedBytesReceived uint64

	//
	AuthenticatedBytesSent uint64

	//
	BadSPIPackets uint64

	//
	ConfidentialBytesReceived uint64

	//
	ConfidentialBytesSent uint64

	//
	InstanceName string

	//
	KeyAdditions uint64

	//
	KeyDeletions uint64

	//
	PacketsNotAuthenticated uint64

	//
	PacketsNotDecrypted uint64

	//
	PacketsWithReplayDetection uint64

	//
	PendingKeyOperations uint64

	//
	ReKeys uint64

	//
	TransportBytesReceived uint64

	//
	TransportBytesSent uint64

	//
	TunnelBytesReceived uint64

	//
	TunnelBytesSent uint64
}

func NewMSiSCSI_QMIPSECStatsEx1(instance *cim.WmiInstance) (newInstance *MSiSCSI_QMIPSECStats, err error) {
	tmp, err := NewWin32_PerfRawDataEx1(instance)

	if err != nil {
		return
	}
	newInstance = &MSiSCSI_QMIPSECStats{
		Win32_PerfRawData: tmp,
	}
	return
}

func NewMSiSCSI_QMIPSECStatsEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *MSiSCSI_QMIPSECStats, err error) {
	tmp, err := NewWin32_PerfRawDataEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &MSiSCSI_QMIPSECStats{
		Win32_PerfRawData: tmp,
	}
	return
}

// SetActive sets the value of Active for the instance
func (instance *MSiSCSI_QMIPSECStats) SetPropertyActive(value bool) (err error) {
	return instance.SetProperty("Active", (value))
}

// GetActive gets the value of Active for the instance
func (instance *MSiSCSI_QMIPSECStats) GetPropertyActive() (value bool, err error) {
	retValue, err := instance.GetProperty("Active")
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

// SetActiveSA sets the value of ActiveSA for the instance
func (instance *MSiSCSI_QMIPSECStats) SetPropertyActiveSA(value uint64) (err error) {
	return instance.SetProperty("ActiveSA", (value))
}

// GetActiveSA gets the value of ActiveSA for the instance
func (instance *MSiSCSI_QMIPSECStats) GetPropertyActiveSA() (value uint64, err error) {
	retValue, err := instance.GetProperty("ActiveSA")
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

// SetActiveTunnels sets the value of ActiveTunnels for the instance
func (instance *MSiSCSI_QMIPSECStats) SetPropertyActiveTunnels(value uint64) (err error) {
	return instance.SetProperty("ActiveTunnels", (value))
}

// GetActiveTunnels gets the value of ActiveTunnels for the instance
func (instance *MSiSCSI_QMIPSECStats) GetPropertyActiveTunnels() (value uint64, err error) {
	retValue, err := instance.GetProperty("ActiveTunnels")
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

// SetAuthenticatedBytesReceived sets the value of AuthenticatedBytesReceived for the instance
func (instance *MSiSCSI_QMIPSECStats) SetPropertyAuthenticatedBytesReceived(value uint64) (err error) {
	return instance.SetProperty("AuthenticatedBytesReceived", (value))
}

// GetAuthenticatedBytesReceived gets the value of AuthenticatedBytesReceived for the instance
func (instance *MSiSCSI_QMIPSECStats) GetPropertyAuthenticatedBytesReceived() (value uint64, err error) {
	retValue, err := instance.GetProperty("AuthenticatedBytesReceived")
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

// SetAuthenticatedBytesSent sets the value of AuthenticatedBytesSent for the instance
func (instance *MSiSCSI_QMIPSECStats) SetPropertyAuthenticatedBytesSent(value uint64) (err error) {
	return instance.SetProperty("AuthenticatedBytesSent", (value))
}

// GetAuthenticatedBytesSent gets the value of AuthenticatedBytesSent for the instance
func (instance *MSiSCSI_QMIPSECStats) GetPropertyAuthenticatedBytesSent() (value uint64, err error) {
	retValue, err := instance.GetProperty("AuthenticatedBytesSent")
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

// SetBadSPIPackets sets the value of BadSPIPackets for the instance
func (instance *MSiSCSI_QMIPSECStats) SetPropertyBadSPIPackets(value uint64) (err error) {
	return instance.SetProperty("BadSPIPackets", (value))
}

// GetBadSPIPackets gets the value of BadSPIPackets for the instance
func (instance *MSiSCSI_QMIPSECStats) GetPropertyBadSPIPackets() (value uint64, err error) {
	retValue, err := instance.GetProperty("BadSPIPackets")
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

// SetConfidentialBytesReceived sets the value of ConfidentialBytesReceived for the instance
func (instance *MSiSCSI_QMIPSECStats) SetPropertyConfidentialBytesReceived(value uint64) (err error) {
	return instance.SetProperty("ConfidentialBytesReceived", (value))
}

// GetConfidentialBytesReceived gets the value of ConfidentialBytesReceived for the instance
func (instance *MSiSCSI_QMIPSECStats) GetPropertyConfidentialBytesReceived() (value uint64, err error) {
	retValue, err := instance.GetProperty("ConfidentialBytesReceived")
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

// SetConfidentialBytesSent sets the value of ConfidentialBytesSent for the instance
func (instance *MSiSCSI_QMIPSECStats) SetPropertyConfidentialBytesSent(value uint64) (err error) {
	return instance.SetProperty("ConfidentialBytesSent", (value))
}

// GetConfidentialBytesSent gets the value of ConfidentialBytesSent for the instance
func (instance *MSiSCSI_QMIPSECStats) GetPropertyConfidentialBytesSent() (value uint64, err error) {
	retValue, err := instance.GetProperty("ConfidentialBytesSent")
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

// SetInstanceName sets the value of InstanceName for the instance
func (instance *MSiSCSI_QMIPSECStats) SetPropertyInstanceName(value string) (err error) {
	return instance.SetProperty("InstanceName", (value))
}

// GetInstanceName gets the value of InstanceName for the instance
func (instance *MSiSCSI_QMIPSECStats) GetPropertyInstanceName() (value string, err error) {
	retValue, err := instance.GetProperty("InstanceName")
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

// SetKeyAdditions sets the value of KeyAdditions for the instance
func (instance *MSiSCSI_QMIPSECStats) SetPropertyKeyAdditions(value uint64) (err error) {
	return instance.SetProperty("KeyAdditions", (value))
}

// GetKeyAdditions gets the value of KeyAdditions for the instance
func (instance *MSiSCSI_QMIPSECStats) GetPropertyKeyAdditions() (value uint64, err error) {
	retValue, err := instance.GetProperty("KeyAdditions")
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

// SetKeyDeletions sets the value of KeyDeletions for the instance
func (instance *MSiSCSI_QMIPSECStats) SetPropertyKeyDeletions(value uint64) (err error) {
	return instance.SetProperty("KeyDeletions", (value))
}

// GetKeyDeletions gets the value of KeyDeletions for the instance
func (instance *MSiSCSI_QMIPSECStats) GetPropertyKeyDeletions() (value uint64, err error) {
	retValue, err := instance.GetProperty("KeyDeletions")
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

// SetPacketsNotAuthenticated sets the value of PacketsNotAuthenticated for the instance
func (instance *MSiSCSI_QMIPSECStats) SetPropertyPacketsNotAuthenticated(value uint64) (err error) {
	return instance.SetProperty("PacketsNotAuthenticated", (value))
}

// GetPacketsNotAuthenticated gets the value of PacketsNotAuthenticated for the instance
func (instance *MSiSCSI_QMIPSECStats) GetPropertyPacketsNotAuthenticated() (value uint64, err error) {
	retValue, err := instance.GetProperty("PacketsNotAuthenticated")
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

// SetPacketsNotDecrypted sets the value of PacketsNotDecrypted for the instance
func (instance *MSiSCSI_QMIPSECStats) SetPropertyPacketsNotDecrypted(value uint64) (err error) {
	return instance.SetProperty("PacketsNotDecrypted", (value))
}

// GetPacketsNotDecrypted gets the value of PacketsNotDecrypted for the instance
func (instance *MSiSCSI_QMIPSECStats) GetPropertyPacketsNotDecrypted() (value uint64, err error) {
	retValue, err := instance.GetProperty("PacketsNotDecrypted")
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

// SetPacketsWithReplayDetection sets the value of PacketsWithReplayDetection for the instance
func (instance *MSiSCSI_QMIPSECStats) SetPropertyPacketsWithReplayDetection(value uint64) (err error) {
	return instance.SetProperty("PacketsWithReplayDetection", (value))
}

// GetPacketsWithReplayDetection gets the value of PacketsWithReplayDetection for the instance
func (instance *MSiSCSI_QMIPSECStats) GetPropertyPacketsWithReplayDetection() (value uint64, err error) {
	retValue, err := instance.GetProperty("PacketsWithReplayDetection")
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

// SetPendingKeyOperations sets the value of PendingKeyOperations for the instance
func (instance *MSiSCSI_QMIPSECStats) SetPropertyPendingKeyOperations(value uint64) (err error) {
	return instance.SetProperty("PendingKeyOperations", (value))
}

// GetPendingKeyOperations gets the value of PendingKeyOperations for the instance
func (instance *MSiSCSI_QMIPSECStats) GetPropertyPendingKeyOperations() (value uint64, err error) {
	retValue, err := instance.GetProperty("PendingKeyOperations")
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

// SetReKeys sets the value of ReKeys for the instance
func (instance *MSiSCSI_QMIPSECStats) SetPropertyReKeys(value uint64) (err error) {
	return instance.SetProperty("ReKeys", (value))
}

// GetReKeys gets the value of ReKeys for the instance
func (instance *MSiSCSI_QMIPSECStats) GetPropertyReKeys() (value uint64, err error) {
	retValue, err := instance.GetProperty("ReKeys")
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

// SetTransportBytesReceived sets the value of TransportBytesReceived for the instance
func (instance *MSiSCSI_QMIPSECStats) SetPropertyTransportBytesReceived(value uint64) (err error) {
	return instance.SetProperty("TransportBytesReceived", (value))
}

// GetTransportBytesReceived gets the value of TransportBytesReceived for the instance
func (instance *MSiSCSI_QMIPSECStats) GetPropertyTransportBytesReceived() (value uint64, err error) {
	retValue, err := instance.GetProperty("TransportBytesReceived")
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

// SetTransportBytesSent sets the value of TransportBytesSent for the instance
func (instance *MSiSCSI_QMIPSECStats) SetPropertyTransportBytesSent(value uint64) (err error) {
	return instance.SetProperty("TransportBytesSent", (value))
}

// GetTransportBytesSent gets the value of TransportBytesSent for the instance
func (instance *MSiSCSI_QMIPSECStats) GetPropertyTransportBytesSent() (value uint64, err error) {
	retValue, err := instance.GetProperty("TransportBytesSent")
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

// SetTunnelBytesReceived sets the value of TunnelBytesReceived for the instance
func (instance *MSiSCSI_QMIPSECStats) SetPropertyTunnelBytesReceived(value uint64) (err error) {
	return instance.SetProperty("TunnelBytesReceived", (value))
}

// GetTunnelBytesReceived gets the value of TunnelBytesReceived for the instance
func (instance *MSiSCSI_QMIPSECStats) GetPropertyTunnelBytesReceived() (value uint64, err error) {
	retValue, err := instance.GetProperty("TunnelBytesReceived")
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

// SetTunnelBytesSent sets the value of TunnelBytesSent for the instance
func (instance *MSiSCSI_QMIPSECStats) SetPropertyTunnelBytesSent(value uint64) (err error) {
	return instance.SetProperty("TunnelBytesSent", (value))
}

// GetTunnelBytesSent gets the value of TunnelBytesSent for the instance
func (instance *MSiSCSI_QMIPSECStats) GetPropertyTunnelBytesSent() (value uint64, err error) {
	retValue, err := instance.GetProperty("TunnelBytesSent")
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
