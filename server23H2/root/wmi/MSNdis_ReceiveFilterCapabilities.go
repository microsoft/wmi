// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 6/13/2024 using wmigen
//      Source root.WMI
//////////////////////////////////////////////
package wmi

import (
	"github.com/microsoft/wmi/pkg/base/query"
	"github.com/microsoft/wmi/pkg/errors"
	cim "github.com/microsoft/wmi/pkg/wmiinstance"
	"reflect"
)

// MSNdis_ReceiveFilterCapabilities struct
type MSNdis_ReceiveFilterCapabilities struct {
	*MSNdis

	//
	EnabledFilterTypes uint32

	//
	EnabledQueueTypes uint32

	//
	Flags uint32

	//
	Header MSNdis_ObjectHeader

	//
	MaxLookaheadSplitSize uint32

	//
	MaxMacHeaderFilters uint32

	//
	MaxQueueGroups uint32

	//
	MaxQueuesPerQueueGroup uint32

	//
	MinLookaheadSplitSize uint32

	//
	NumQueues uint32

	//
	SupportedFilterTests uint32

	//
	SupportedHeaders uint32

	//
	SupportedMacHeaderFields uint32

	//
	SupportedQueueProperties uint32
}

func NewMSNdis_ReceiveFilterCapabilitiesEx1(instance *cim.WmiInstance) (newInstance *MSNdis_ReceiveFilterCapabilities, err error) {
	tmp, err := NewMSNdisEx1(instance)

	if err != nil {
		return
	}
	newInstance = &MSNdis_ReceiveFilterCapabilities{
		MSNdis: tmp,
	}
	return
}

func NewMSNdis_ReceiveFilterCapabilitiesEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *MSNdis_ReceiveFilterCapabilities, err error) {
	tmp, err := NewMSNdisEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &MSNdis_ReceiveFilterCapabilities{
		MSNdis: tmp,
	}
	return
}

// SetEnabledFilterTypes sets the value of EnabledFilterTypes for the instance
func (instance *MSNdis_ReceiveFilterCapabilities) SetPropertyEnabledFilterTypes(value uint32) (err error) {
	return instance.SetProperty("EnabledFilterTypes", (value))
}

// GetEnabledFilterTypes gets the value of EnabledFilterTypes for the instance
func (instance *MSNdis_ReceiveFilterCapabilities) GetPropertyEnabledFilterTypes() (value uint32, err error) {
	retValue, err := instance.GetProperty("EnabledFilterTypes")
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

// SetEnabledQueueTypes sets the value of EnabledQueueTypes for the instance
func (instance *MSNdis_ReceiveFilterCapabilities) SetPropertyEnabledQueueTypes(value uint32) (err error) {
	return instance.SetProperty("EnabledQueueTypes", (value))
}

// GetEnabledQueueTypes gets the value of EnabledQueueTypes for the instance
func (instance *MSNdis_ReceiveFilterCapabilities) GetPropertyEnabledQueueTypes() (value uint32, err error) {
	retValue, err := instance.GetProperty("EnabledQueueTypes")
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

// SetFlags sets the value of Flags for the instance
func (instance *MSNdis_ReceiveFilterCapabilities) SetPropertyFlags(value uint32) (err error) {
	return instance.SetProperty("Flags", (value))
}

// GetFlags gets the value of Flags for the instance
func (instance *MSNdis_ReceiveFilterCapabilities) GetPropertyFlags() (value uint32, err error) {
	retValue, err := instance.GetProperty("Flags")
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

// SetHeader sets the value of Header for the instance
func (instance *MSNdis_ReceiveFilterCapabilities) SetPropertyHeader(value MSNdis_ObjectHeader) (err error) {
	return instance.SetProperty("Header", (value))
}

// GetHeader gets the value of Header for the instance
func (instance *MSNdis_ReceiveFilterCapabilities) GetPropertyHeader() (value MSNdis_ObjectHeader, err error) {
	retValue, err := instance.GetProperty("Header")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(MSNdis_ObjectHeader)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " MSNdis_ObjectHeader is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = MSNdis_ObjectHeader(valuetmp)

	return
}

// SetMaxLookaheadSplitSize sets the value of MaxLookaheadSplitSize for the instance
func (instance *MSNdis_ReceiveFilterCapabilities) SetPropertyMaxLookaheadSplitSize(value uint32) (err error) {
	return instance.SetProperty("MaxLookaheadSplitSize", (value))
}

// GetMaxLookaheadSplitSize gets the value of MaxLookaheadSplitSize for the instance
func (instance *MSNdis_ReceiveFilterCapabilities) GetPropertyMaxLookaheadSplitSize() (value uint32, err error) {
	retValue, err := instance.GetProperty("MaxLookaheadSplitSize")
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

// SetMaxMacHeaderFilters sets the value of MaxMacHeaderFilters for the instance
func (instance *MSNdis_ReceiveFilterCapabilities) SetPropertyMaxMacHeaderFilters(value uint32) (err error) {
	return instance.SetProperty("MaxMacHeaderFilters", (value))
}

// GetMaxMacHeaderFilters gets the value of MaxMacHeaderFilters for the instance
func (instance *MSNdis_ReceiveFilterCapabilities) GetPropertyMaxMacHeaderFilters() (value uint32, err error) {
	retValue, err := instance.GetProperty("MaxMacHeaderFilters")
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

// SetMaxQueueGroups sets the value of MaxQueueGroups for the instance
func (instance *MSNdis_ReceiveFilterCapabilities) SetPropertyMaxQueueGroups(value uint32) (err error) {
	return instance.SetProperty("MaxQueueGroups", (value))
}

// GetMaxQueueGroups gets the value of MaxQueueGroups for the instance
func (instance *MSNdis_ReceiveFilterCapabilities) GetPropertyMaxQueueGroups() (value uint32, err error) {
	retValue, err := instance.GetProperty("MaxQueueGroups")
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

// SetMaxQueuesPerQueueGroup sets the value of MaxQueuesPerQueueGroup for the instance
func (instance *MSNdis_ReceiveFilterCapabilities) SetPropertyMaxQueuesPerQueueGroup(value uint32) (err error) {
	return instance.SetProperty("MaxQueuesPerQueueGroup", (value))
}

// GetMaxQueuesPerQueueGroup gets the value of MaxQueuesPerQueueGroup for the instance
func (instance *MSNdis_ReceiveFilterCapabilities) GetPropertyMaxQueuesPerQueueGroup() (value uint32, err error) {
	retValue, err := instance.GetProperty("MaxQueuesPerQueueGroup")
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

// SetMinLookaheadSplitSize sets the value of MinLookaheadSplitSize for the instance
func (instance *MSNdis_ReceiveFilterCapabilities) SetPropertyMinLookaheadSplitSize(value uint32) (err error) {
	return instance.SetProperty("MinLookaheadSplitSize", (value))
}

// GetMinLookaheadSplitSize gets the value of MinLookaheadSplitSize for the instance
func (instance *MSNdis_ReceiveFilterCapabilities) GetPropertyMinLookaheadSplitSize() (value uint32, err error) {
	retValue, err := instance.GetProperty("MinLookaheadSplitSize")
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

// SetNumQueues sets the value of NumQueues for the instance
func (instance *MSNdis_ReceiveFilterCapabilities) SetPropertyNumQueues(value uint32) (err error) {
	return instance.SetProperty("NumQueues", (value))
}

// GetNumQueues gets the value of NumQueues for the instance
func (instance *MSNdis_ReceiveFilterCapabilities) GetPropertyNumQueues() (value uint32, err error) {
	retValue, err := instance.GetProperty("NumQueues")
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

// SetSupportedFilterTests sets the value of SupportedFilterTests for the instance
func (instance *MSNdis_ReceiveFilterCapabilities) SetPropertySupportedFilterTests(value uint32) (err error) {
	return instance.SetProperty("SupportedFilterTests", (value))
}

// GetSupportedFilterTests gets the value of SupportedFilterTests for the instance
func (instance *MSNdis_ReceiveFilterCapabilities) GetPropertySupportedFilterTests() (value uint32, err error) {
	retValue, err := instance.GetProperty("SupportedFilterTests")
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

// SetSupportedHeaders sets the value of SupportedHeaders for the instance
func (instance *MSNdis_ReceiveFilterCapabilities) SetPropertySupportedHeaders(value uint32) (err error) {
	return instance.SetProperty("SupportedHeaders", (value))
}

// GetSupportedHeaders gets the value of SupportedHeaders for the instance
func (instance *MSNdis_ReceiveFilterCapabilities) GetPropertySupportedHeaders() (value uint32, err error) {
	retValue, err := instance.GetProperty("SupportedHeaders")
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

// SetSupportedMacHeaderFields sets the value of SupportedMacHeaderFields for the instance
func (instance *MSNdis_ReceiveFilterCapabilities) SetPropertySupportedMacHeaderFields(value uint32) (err error) {
	return instance.SetProperty("SupportedMacHeaderFields", (value))
}

// GetSupportedMacHeaderFields gets the value of SupportedMacHeaderFields for the instance
func (instance *MSNdis_ReceiveFilterCapabilities) GetPropertySupportedMacHeaderFields() (value uint32, err error) {
	retValue, err := instance.GetProperty("SupportedMacHeaderFields")
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

// SetSupportedQueueProperties sets the value of SupportedQueueProperties for the instance
func (instance *MSNdis_ReceiveFilterCapabilities) SetPropertySupportedQueueProperties(value uint32) (err error) {
	return instance.SetProperty("SupportedQueueProperties", (value))
}

// GetSupportedQueueProperties gets the value of SupportedQueueProperties for the instance
func (instance *MSNdis_ReceiveFilterCapabilities) GetPropertySupportedQueueProperties() (value uint32, err error) {
	retValue, err := instance.GetProperty("SupportedQueueProperties")
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
