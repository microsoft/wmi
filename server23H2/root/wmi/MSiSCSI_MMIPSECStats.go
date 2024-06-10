// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

// Author:
//
//	Auto Generated on 6/6/2024 using wmigen
//	Source root.WMI
//
// ////////////////////////////////////////////
package wmi

import (
	"github.com/microsoft/wmi/pkg/base/query"
	"github.com/microsoft/wmi/pkg/errors"
	cim "github.com/microsoft/wmi/pkg/wmiinstance"
	"reflect"
)

// MSiSCSI_MMIPSECStats struct
type MSiSCSI_MMIPSECStats struct {
	*Win32_PerfRawData

	//
	AcquireFailures uint64

	//
	AcquireHeapSize uint64

	//
	Active bool

	//
	ActiveAcquire uint64

	//
	ActiveReceive uint64

	//
	AuthenticationFailures uint64

	//
	ConnectionListSize uint64

	//
	GetSPIFailures uint64

	//
	InstanceName string

	//
	InvalidCookiesReceived uint64

	//
	InvalidPackets uint64

	//
	KeyAdditionFailures uint64

	//
	KeyAdditions uint64

	//
	KeyUpdateFailures uint64

	//
	KeyUpdates uint64

	//
	NegotiationFailures uint64

	//
	OakleyMainMode uint64

	//
	OakleyQuickMode uint64

	//
	ReceiveFailures uint64

	//
	ReceiveHeapSize uint64

	//
	SendFailures uint64

	//
	SoftAssociations uint64

	//
	TotalGetSPI uint64
}

func NewMSiSCSI_MMIPSECStatsEx1(instance *cim.WmiInstance) (newInstance *MSiSCSI_MMIPSECStats, err error) {
	tmp, err := NewWin32_PerfRawDataEx1(instance)

	if err != nil {
		return
	}
	newInstance = &MSiSCSI_MMIPSECStats{
		Win32_PerfRawData: tmp,
	}
	return
}

func NewMSiSCSI_MMIPSECStatsEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *MSiSCSI_MMIPSECStats, err error) {
	tmp, err := NewWin32_PerfRawDataEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &MSiSCSI_MMIPSECStats{
		Win32_PerfRawData: tmp,
	}
	return
}

// SetAcquireFailures sets the value of AcquireFailures for the instance
func (instance *MSiSCSI_MMIPSECStats) SetPropertyAcquireFailures(value uint64) (err error) {
	return instance.SetProperty("AcquireFailures", (value))
}

// GetAcquireFailures gets the value of AcquireFailures for the instance
func (instance *MSiSCSI_MMIPSECStats) GetPropertyAcquireFailures() (value uint64, err error) {
	retValue, err := instance.GetProperty("AcquireFailures")
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

// SetAcquireHeapSize sets the value of AcquireHeapSize for the instance
func (instance *MSiSCSI_MMIPSECStats) SetPropertyAcquireHeapSize(value uint64) (err error) {
	return instance.SetProperty("AcquireHeapSize", (value))
}

// GetAcquireHeapSize gets the value of AcquireHeapSize for the instance
func (instance *MSiSCSI_MMIPSECStats) GetPropertyAcquireHeapSize() (value uint64, err error) {
	retValue, err := instance.GetProperty("AcquireHeapSize")
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

// SetActive sets the value of Active for the instance
func (instance *MSiSCSI_MMIPSECStats) SetPropertyActive(value bool) (err error) {
	return instance.SetProperty("Active", (value))
}

// GetActive gets the value of Active for the instance
func (instance *MSiSCSI_MMIPSECStats) GetPropertyActive() (value bool, err error) {
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

// SetActiveAcquire sets the value of ActiveAcquire for the instance
func (instance *MSiSCSI_MMIPSECStats) SetPropertyActiveAcquire(value uint64) (err error) {
	return instance.SetProperty("ActiveAcquire", (value))
}

// GetActiveAcquire gets the value of ActiveAcquire for the instance
func (instance *MSiSCSI_MMIPSECStats) GetPropertyActiveAcquire() (value uint64, err error) {
	retValue, err := instance.GetProperty("ActiveAcquire")
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

// SetActiveReceive sets the value of ActiveReceive for the instance
func (instance *MSiSCSI_MMIPSECStats) SetPropertyActiveReceive(value uint64) (err error) {
	return instance.SetProperty("ActiveReceive", (value))
}

// GetActiveReceive gets the value of ActiveReceive for the instance
func (instance *MSiSCSI_MMIPSECStats) GetPropertyActiveReceive() (value uint64, err error) {
	retValue, err := instance.GetProperty("ActiveReceive")
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

// SetAuthenticationFailures sets the value of AuthenticationFailures for the instance
func (instance *MSiSCSI_MMIPSECStats) SetPropertyAuthenticationFailures(value uint64) (err error) {
	return instance.SetProperty("AuthenticationFailures", (value))
}

// GetAuthenticationFailures gets the value of AuthenticationFailures for the instance
func (instance *MSiSCSI_MMIPSECStats) GetPropertyAuthenticationFailures() (value uint64, err error) {
	retValue, err := instance.GetProperty("AuthenticationFailures")
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

// SetConnectionListSize sets the value of ConnectionListSize for the instance
func (instance *MSiSCSI_MMIPSECStats) SetPropertyConnectionListSize(value uint64) (err error) {
	return instance.SetProperty("ConnectionListSize", (value))
}

// GetConnectionListSize gets the value of ConnectionListSize for the instance
func (instance *MSiSCSI_MMIPSECStats) GetPropertyConnectionListSize() (value uint64, err error) {
	retValue, err := instance.GetProperty("ConnectionListSize")
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

// SetGetSPIFailures sets the value of GetSPIFailures for the instance
func (instance *MSiSCSI_MMIPSECStats) SetPropertyGetSPIFailures(value uint64) (err error) {
	return instance.SetProperty("GetSPIFailures", (value))
}

// GetGetSPIFailures gets the value of GetSPIFailures for the instance
func (instance *MSiSCSI_MMIPSECStats) GetPropertyGetSPIFailures() (value uint64, err error) {
	retValue, err := instance.GetProperty("GetSPIFailures")
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
func (instance *MSiSCSI_MMIPSECStats) SetPropertyInstanceName(value string) (err error) {
	return instance.SetProperty("InstanceName", (value))
}

// GetInstanceName gets the value of InstanceName for the instance
func (instance *MSiSCSI_MMIPSECStats) GetPropertyInstanceName() (value string, err error) {
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

// SetInvalidCookiesReceived sets the value of InvalidCookiesReceived for the instance
func (instance *MSiSCSI_MMIPSECStats) SetPropertyInvalidCookiesReceived(value uint64) (err error) {
	return instance.SetProperty("InvalidCookiesReceived", (value))
}

// GetInvalidCookiesReceived gets the value of InvalidCookiesReceived for the instance
func (instance *MSiSCSI_MMIPSECStats) GetPropertyInvalidCookiesReceived() (value uint64, err error) {
	retValue, err := instance.GetProperty("InvalidCookiesReceived")
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

// SetInvalidPackets sets the value of InvalidPackets for the instance
func (instance *MSiSCSI_MMIPSECStats) SetPropertyInvalidPackets(value uint64) (err error) {
	return instance.SetProperty("InvalidPackets", (value))
}

// GetInvalidPackets gets the value of InvalidPackets for the instance
func (instance *MSiSCSI_MMIPSECStats) GetPropertyInvalidPackets() (value uint64, err error) {
	retValue, err := instance.GetProperty("InvalidPackets")
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

// SetKeyAdditionFailures sets the value of KeyAdditionFailures for the instance
func (instance *MSiSCSI_MMIPSECStats) SetPropertyKeyAdditionFailures(value uint64) (err error) {
	return instance.SetProperty("KeyAdditionFailures", (value))
}

// GetKeyAdditionFailures gets the value of KeyAdditionFailures for the instance
func (instance *MSiSCSI_MMIPSECStats) GetPropertyKeyAdditionFailures() (value uint64, err error) {
	retValue, err := instance.GetProperty("KeyAdditionFailures")
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

// SetKeyAdditions sets the value of KeyAdditions for the instance
func (instance *MSiSCSI_MMIPSECStats) SetPropertyKeyAdditions(value uint64) (err error) {
	return instance.SetProperty("KeyAdditions", (value))
}

// GetKeyAdditions gets the value of KeyAdditions for the instance
func (instance *MSiSCSI_MMIPSECStats) GetPropertyKeyAdditions() (value uint64, err error) {
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

// SetKeyUpdateFailures sets the value of KeyUpdateFailures for the instance
func (instance *MSiSCSI_MMIPSECStats) SetPropertyKeyUpdateFailures(value uint64) (err error) {
	return instance.SetProperty("KeyUpdateFailures", (value))
}

// GetKeyUpdateFailures gets the value of KeyUpdateFailures for the instance
func (instance *MSiSCSI_MMIPSECStats) GetPropertyKeyUpdateFailures() (value uint64, err error) {
	retValue, err := instance.GetProperty("KeyUpdateFailures")
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

// SetKeyUpdates sets the value of KeyUpdates for the instance
func (instance *MSiSCSI_MMIPSECStats) SetPropertyKeyUpdates(value uint64) (err error) {
	return instance.SetProperty("KeyUpdates", (value))
}

// GetKeyUpdates gets the value of KeyUpdates for the instance
func (instance *MSiSCSI_MMIPSECStats) GetPropertyKeyUpdates() (value uint64, err error) {
	retValue, err := instance.GetProperty("KeyUpdates")
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

// SetNegotiationFailures sets the value of NegotiationFailures for the instance
func (instance *MSiSCSI_MMIPSECStats) SetPropertyNegotiationFailures(value uint64) (err error) {
	return instance.SetProperty("NegotiationFailures", (value))
}

// GetNegotiationFailures gets the value of NegotiationFailures for the instance
func (instance *MSiSCSI_MMIPSECStats) GetPropertyNegotiationFailures() (value uint64, err error) {
	retValue, err := instance.GetProperty("NegotiationFailures")
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

// SetOakleyMainMode sets the value of OakleyMainMode for the instance
func (instance *MSiSCSI_MMIPSECStats) SetPropertyOakleyMainMode(value uint64) (err error) {
	return instance.SetProperty("OakleyMainMode", (value))
}

// GetOakleyMainMode gets the value of OakleyMainMode for the instance
func (instance *MSiSCSI_MMIPSECStats) GetPropertyOakleyMainMode() (value uint64, err error) {
	retValue, err := instance.GetProperty("OakleyMainMode")
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

// SetOakleyQuickMode sets the value of OakleyQuickMode for the instance
func (instance *MSiSCSI_MMIPSECStats) SetPropertyOakleyQuickMode(value uint64) (err error) {
	return instance.SetProperty("OakleyQuickMode", (value))
}

// GetOakleyQuickMode gets the value of OakleyQuickMode for the instance
func (instance *MSiSCSI_MMIPSECStats) GetPropertyOakleyQuickMode() (value uint64, err error) {
	retValue, err := instance.GetProperty("OakleyQuickMode")
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

// SetReceiveFailures sets the value of ReceiveFailures for the instance
func (instance *MSiSCSI_MMIPSECStats) SetPropertyReceiveFailures(value uint64) (err error) {
	return instance.SetProperty("ReceiveFailures", (value))
}

// GetReceiveFailures gets the value of ReceiveFailures for the instance
func (instance *MSiSCSI_MMIPSECStats) GetPropertyReceiveFailures() (value uint64, err error) {
	retValue, err := instance.GetProperty("ReceiveFailures")
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

// SetReceiveHeapSize sets the value of ReceiveHeapSize for the instance
func (instance *MSiSCSI_MMIPSECStats) SetPropertyReceiveHeapSize(value uint64) (err error) {
	return instance.SetProperty("ReceiveHeapSize", (value))
}

// GetReceiveHeapSize gets the value of ReceiveHeapSize for the instance
func (instance *MSiSCSI_MMIPSECStats) GetPropertyReceiveHeapSize() (value uint64, err error) {
	retValue, err := instance.GetProperty("ReceiveHeapSize")
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

// SetSendFailures sets the value of SendFailures for the instance
func (instance *MSiSCSI_MMIPSECStats) SetPropertySendFailures(value uint64) (err error) {
	return instance.SetProperty("SendFailures", (value))
}

// GetSendFailures gets the value of SendFailures for the instance
func (instance *MSiSCSI_MMIPSECStats) GetPropertySendFailures() (value uint64, err error) {
	retValue, err := instance.GetProperty("SendFailures")
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

// SetSoftAssociations sets the value of SoftAssociations for the instance
func (instance *MSiSCSI_MMIPSECStats) SetPropertySoftAssociations(value uint64) (err error) {
	return instance.SetProperty("SoftAssociations", (value))
}

// GetSoftAssociations gets the value of SoftAssociations for the instance
func (instance *MSiSCSI_MMIPSECStats) GetPropertySoftAssociations() (value uint64, err error) {
	retValue, err := instance.GetProperty("SoftAssociations")
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

// SetTotalGetSPI sets the value of TotalGetSPI for the instance
func (instance *MSiSCSI_MMIPSECStats) SetPropertyTotalGetSPI(value uint64) (err error) {
	return instance.SetProperty("TotalGetSPI", (value))
}

// GetTotalGetSPI gets the value of TotalGetSPI for the instance
func (instance *MSiSCSI_MMIPSECStats) GetPropertyTotalGetSPI() (value uint64, err error) {
	retValue, err := instance.GetProperty("TotalGetSPI")
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
