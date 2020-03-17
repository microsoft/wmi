// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 3/16/2020 using wmigen
//      Source root.CIMV2
//////////////////////////////////////////////
package cimv2

import (
	"github.com/microsoft/wmi/pkg/wmiinstance"
)

// Win32_PingStatus struct
type Win32_PingStatus struct {
	cim.WmiInstance

	//
	Address string

	//
	BufferSize uint32

	//
	NoFragmentation bool

	//
	PrimaryAddressResolutionStatus uint32

	//
	ProtocolAddress string

	//
	ProtocolAddressResolved string

	//
	RecordRoute uint32

	//
	ReplyInconsistency bool

	//
	ReplySize uint32

	//
	ResolveAddressNames bool

	//
	ResponseTime uint32

	//
	ResponseTimeToLive uint32

	//
	RouteRecord []string

	//
	RouteRecordResolved []string

	//
	SourceRoute string

	//
	SourceRouteType uint32

	//
	StatusCode uint32

	//
	Timeout uint32

	//
	TimeStampRecord []uint32

	//
	TimeStampRecordAddress []string

	//
	TimeStampRecordAddressResolved []string

	//
	TimestampRoute uint32

	//
	TimeToLive uint32

	//
	TypeofService uint32
}

// SetAddress sets the value of Address for the instance
func (instance *Win32_PingStatus) SetPropertyAddress(value string) (err error) {
	return instance.SetProperty("Address", value)
}

// GetAddress gets the value of Address for the instance
func (instance *Win32_PingStatus) GetPropertyAddress() (value string, err error) {
	retValue, err := instance.GetProperty("Address")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetBufferSize sets the value of BufferSize for the instance
func (instance *Win32_PingStatus) SetPropertyBufferSize(value uint32) (err error) {
	return instance.SetProperty("BufferSize", value)
}

// GetBufferSize gets the value of BufferSize for the instance
func (instance *Win32_PingStatus) GetPropertyBufferSize() (value uint32, err error) {
	retValue, err := instance.GetProperty("BufferSize")
	if err != nil {
		return
	}
	value, ok := retValue.(uint32)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetNoFragmentation sets the value of NoFragmentation for the instance
func (instance *Win32_PingStatus) SetPropertyNoFragmentation(value bool) (err error) {
	return instance.SetProperty("NoFragmentation", value)
}

// GetNoFragmentation gets the value of NoFragmentation for the instance
func (instance *Win32_PingStatus) GetPropertyNoFragmentation() (value bool, err error) {
	retValue, err := instance.GetProperty("NoFragmentation")
	if err != nil {
		return
	}
	value, ok := retValue.(bool)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetPrimaryAddressResolutionStatus sets the value of PrimaryAddressResolutionStatus for the instance
func (instance *Win32_PingStatus) SetPropertyPrimaryAddressResolutionStatus(value uint32) (err error) {
	return instance.SetProperty("PrimaryAddressResolutionStatus", value)
}

// GetPrimaryAddressResolutionStatus gets the value of PrimaryAddressResolutionStatus for the instance
func (instance *Win32_PingStatus) GetPropertyPrimaryAddressResolutionStatus() (value uint32, err error) {
	retValue, err := instance.GetProperty("PrimaryAddressResolutionStatus")
	if err != nil {
		return
	}
	value, ok := retValue.(uint32)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetProtocolAddress sets the value of ProtocolAddress for the instance
func (instance *Win32_PingStatus) SetPropertyProtocolAddress(value string) (err error) {
	return instance.SetProperty("ProtocolAddress", value)
}

// GetProtocolAddress gets the value of ProtocolAddress for the instance
func (instance *Win32_PingStatus) GetPropertyProtocolAddress() (value string, err error) {
	retValue, err := instance.GetProperty("ProtocolAddress")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetProtocolAddressResolved sets the value of ProtocolAddressResolved for the instance
func (instance *Win32_PingStatus) SetPropertyProtocolAddressResolved(value string) (err error) {
	return instance.SetProperty("ProtocolAddressResolved", value)
}

// GetProtocolAddressResolved gets the value of ProtocolAddressResolved for the instance
func (instance *Win32_PingStatus) GetPropertyProtocolAddressResolved() (value string, err error) {
	retValue, err := instance.GetProperty("ProtocolAddressResolved")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetRecordRoute sets the value of RecordRoute for the instance
func (instance *Win32_PingStatus) SetPropertyRecordRoute(value uint32) (err error) {
	return instance.SetProperty("RecordRoute", value)
}

// GetRecordRoute gets the value of RecordRoute for the instance
func (instance *Win32_PingStatus) GetPropertyRecordRoute() (value uint32, err error) {
	retValue, err := instance.GetProperty("RecordRoute")
	if err != nil {
		return
	}
	value, ok := retValue.(uint32)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetReplyInconsistency sets the value of ReplyInconsistency for the instance
func (instance *Win32_PingStatus) SetPropertyReplyInconsistency(value bool) (err error) {
	return instance.SetProperty("ReplyInconsistency", value)
}

// GetReplyInconsistency gets the value of ReplyInconsistency for the instance
func (instance *Win32_PingStatus) GetPropertyReplyInconsistency() (value bool, err error) {
	retValue, err := instance.GetProperty("ReplyInconsistency")
	if err != nil {
		return
	}
	value, ok := retValue.(bool)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetReplySize sets the value of ReplySize for the instance
func (instance *Win32_PingStatus) SetPropertyReplySize(value uint32) (err error) {
	return instance.SetProperty("ReplySize", value)
}

// GetReplySize gets the value of ReplySize for the instance
func (instance *Win32_PingStatus) GetPropertyReplySize() (value uint32, err error) {
	retValue, err := instance.GetProperty("ReplySize")
	if err != nil {
		return
	}
	value, ok := retValue.(uint32)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetResolveAddressNames sets the value of ResolveAddressNames for the instance
func (instance *Win32_PingStatus) SetPropertyResolveAddressNames(value bool) (err error) {
	return instance.SetProperty("ResolveAddressNames", value)
}

// GetResolveAddressNames gets the value of ResolveAddressNames for the instance
func (instance *Win32_PingStatus) GetPropertyResolveAddressNames() (value bool, err error) {
	retValue, err := instance.GetProperty("ResolveAddressNames")
	if err != nil {
		return
	}
	value, ok := retValue.(bool)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetResponseTime sets the value of ResponseTime for the instance
func (instance *Win32_PingStatus) SetPropertyResponseTime(value uint32) (err error) {
	return instance.SetProperty("ResponseTime", value)
}

// GetResponseTime gets the value of ResponseTime for the instance
func (instance *Win32_PingStatus) GetPropertyResponseTime() (value uint32, err error) {
	retValue, err := instance.GetProperty("ResponseTime")
	if err != nil {
		return
	}
	value, ok := retValue.(uint32)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetResponseTimeToLive sets the value of ResponseTimeToLive for the instance
func (instance *Win32_PingStatus) SetPropertyResponseTimeToLive(value uint32) (err error) {
	return instance.SetProperty("ResponseTimeToLive", value)
}

// GetResponseTimeToLive gets the value of ResponseTimeToLive for the instance
func (instance *Win32_PingStatus) GetPropertyResponseTimeToLive() (value uint32, err error) {
	retValue, err := instance.GetProperty("ResponseTimeToLive")
	if err != nil {
		return
	}
	value, ok := retValue.(uint32)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetRouteRecord sets the value of RouteRecord for the instance
func (instance *Win32_PingStatus) SetPropertyRouteRecord(value []string) (err error) {
	return instance.SetProperty("RouteRecord", value)
}

// GetRouteRecord gets the value of RouteRecord for the instance
func (instance *Win32_PingStatus) GetPropertyRouteRecord() (value []string, err error) {
	retValue, err := instance.GetProperty("RouteRecord")
	if err != nil {
		return
	}
	value, ok := retValue.([]string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetRouteRecordResolved sets the value of RouteRecordResolved for the instance
func (instance *Win32_PingStatus) SetPropertyRouteRecordResolved(value []string) (err error) {
	return instance.SetProperty("RouteRecordResolved", value)
}

// GetRouteRecordResolved gets the value of RouteRecordResolved for the instance
func (instance *Win32_PingStatus) GetPropertyRouteRecordResolved() (value []string, err error) {
	retValue, err := instance.GetProperty("RouteRecordResolved")
	if err != nil {
		return
	}
	value, ok := retValue.([]string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetSourceRoute sets the value of SourceRoute for the instance
func (instance *Win32_PingStatus) SetPropertySourceRoute(value string) (err error) {
	return instance.SetProperty("SourceRoute", value)
}

// GetSourceRoute gets the value of SourceRoute for the instance
func (instance *Win32_PingStatus) GetPropertySourceRoute() (value string, err error) {
	retValue, err := instance.GetProperty("SourceRoute")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetSourceRouteType sets the value of SourceRouteType for the instance
func (instance *Win32_PingStatus) SetPropertySourceRouteType(value uint32) (err error) {
	return instance.SetProperty("SourceRouteType", value)
}

// GetSourceRouteType gets the value of SourceRouteType for the instance
func (instance *Win32_PingStatus) GetPropertySourceRouteType() (value uint32, err error) {
	retValue, err := instance.GetProperty("SourceRouteType")
	if err != nil {
		return
	}
	value, ok := retValue.(uint32)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetStatusCode sets the value of StatusCode for the instance
func (instance *Win32_PingStatus) SetPropertyStatusCode(value uint32) (err error) {
	return instance.SetProperty("StatusCode", value)
}

// GetStatusCode gets the value of StatusCode for the instance
func (instance *Win32_PingStatus) GetPropertyStatusCode() (value uint32, err error) {
	retValue, err := instance.GetProperty("StatusCode")
	if err != nil {
		return
	}
	value, ok := retValue.(uint32)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetTimeout sets the value of Timeout for the instance
func (instance *Win32_PingStatus) SetPropertyTimeout(value uint32) (err error) {
	return instance.SetProperty("Timeout", value)
}

// GetTimeout gets the value of Timeout for the instance
func (instance *Win32_PingStatus) GetPropertyTimeout() (value uint32, err error) {
	retValue, err := instance.GetProperty("Timeout")
	if err != nil {
		return
	}
	value, ok := retValue.(uint32)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetTimeStampRecord sets the value of TimeStampRecord for the instance
func (instance *Win32_PingStatus) SetPropertyTimeStampRecord(value []uint32) (err error) {
	return instance.SetProperty("TimeStampRecord", value)
}

// GetTimeStampRecord gets the value of TimeStampRecord for the instance
func (instance *Win32_PingStatus) GetPropertyTimeStampRecord() (value []uint32, err error) {
	retValue, err := instance.GetProperty("TimeStampRecord")
	if err != nil {
		return
	}
	value, ok := retValue.([]uint32)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetTimeStampRecordAddress sets the value of TimeStampRecordAddress for the instance
func (instance *Win32_PingStatus) SetPropertyTimeStampRecordAddress(value []string) (err error) {
	return instance.SetProperty("TimeStampRecordAddress", value)
}

// GetTimeStampRecordAddress gets the value of TimeStampRecordAddress for the instance
func (instance *Win32_PingStatus) GetPropertyTimeStampRecordAddress() (value []string, err error) {
	retValue, err := instance.GetProperty("TimeStampRecordAddress")
	if err != nil {
		return
	}
	value, ok := retValue.([]string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetTimeStampRecordAddressResolved sets the value of TimeStampRecordAddressResolved for the instance
func (instance *Win32_PingStatus) SetPropertyTimeStampRecordAddressResolved(value []string) (err error) {
	return instance.SetProperty("TimeStampRecordAddressResolved", value)
}

// GetTimeStampRecordAddressResolved gets the value of TimeStampRecordAddressResolved for the instance
func (instance *Win32_PingStatus) GetPropertyTimeStampRecordAddressResolved() (value []string, err error) {
	retValue, err := instance.GetProperty("TimeStampRecordAddressResolved")
	if err != nil {
		return
	}
	value, ok := retValue.([]string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetTimestampRoute sets the value of TimestampRoute for the instance
func (instance *Win32_PingStatus) SetPropertyTimestampRoute(value uint32) (err error) {
	return instance.SetProperty("TimestampRoute", value)
}

// GetTimestampRoute gets the value of TimestampRoute for the instance
func (instance *Win32_PingStatus) GetPropertyTimestampRoute() (value uint32, err error) {
	retValue, err := instance.GetProperty("TimestampRoute")
	if err != nil {
		return
	}
	value, ok := retValue.(uint32)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetTimeToLive sets the value of TimeToLive for the instance
func (instance *Win32_PingStatus) SetPropertyTimeToLive(value uint32) (err error) {
	return instance.SetProperty("TimeToLive", value)
}

// GetTimeToLive gets the value of TimeToLive for the instance
func (instance *Win32_PingStatus) GetPropertyTimeToLive() (value uint32, err error) {
	retValue, err := instance.GetProperty("TimeToLive")
	if err != nil {
		return
	}
	value, ok := retValue.(uint32)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetTypeofService sets the value of TypeofService for the instance
func (instance *Win32_PingStatus) SetPropertyTypeofService(value uint32) (err error) {
	return instance.SetProperty("TypeofService", value)
}

// GetTypeofService gets the value of TypeofService for the instance
func (instance *Win32_PingStatus) GetPropertyTypeofService() (value uint32, err error) {
	retValue, err := instance.GetProperty("TypeofService")
	if err != nil {
		return
	}
	value, ok := retValue.(uint32)
	if !ok {
		// TODO: Set an error
	}
	return
}
