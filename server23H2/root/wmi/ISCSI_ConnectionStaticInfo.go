// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 6/13/2024 using wmigen
//      Source root.WMI
//////////////////////////////////////////////
package wmi

import (
	"github.com/microsoft/wmi/pkg/base/instance"
	"github.com/microsoft/wmi/pkg/base/query"
	"github.com/microsoft/wmi/pkg/errors"
	cim "github.com/microsoft/wmi/pkg/wmiinstance"
	"reflect"
)

// ISCSI_ConnectionStaticInfo struct
type ISCSI_ConnectionStaticInfo struct {
	*cim.WmiInstance

	// **typedef** Authentication type used when establishing the connection.
	AuthType ConnectionStaticInfo_AuthType

	// The iSCSI connection ID for this connection instance.
	CID uint16

	// **typedef** The name of the iSCSI data digest scheme in use within this session.
	DataIntegrity ConnectionStaticInfo_DataIntegrity

	// Estimated throughput of the link in bytes per second
	EstimatedThroughput uint64

	// **typedef** The name of the iSCSI header digest scheme in use within this session.
	HeaderIntegrity ConnectionStaticInfo_HeaderIntegrity

	// The local network address used for the connection
	LocalAddr ISCSI_IP_Address

	// The local port used for the connection
	LocalPort uint32

	// Maximum Datagram size supported by the transport in bytes
	MaxDatagramSize uint32

	// The maximum data payload size supported for command or data PDUs within this session.
	MaxRecvDataSegmentLength uint32

	// **typedef** The transport protocol over which this connection instance is running.
	Protocol ConnectionStaticInfo_Protocol

	// The remote network address used for the connection
	RemoteAddr ISCSI_IP_Address

	// The remote port used for the connection
	RemotePort uint32

	// Must be zero
	Reserved uint16

	// **typedef** Indicates the current state of this connection
	State ConnectionStaticInfo_State

	// A uniquely generated connection ID. Do not confuse this with CID.
	UniqueConnectionId uint64
}

func NewISCSI_ConnectionStaticInfoEx1(instance *cim.WmiInstance) (newInstance *ISCSI_ConnectionStaticInfo, err error) {
	tmp, err := instance, nil

	if err != nil {
		return
	}
	newInstance = &ISCSI_ConnectionStaticInfo{
		WmiInstance: tmp,
	}
	return
}

func NewISCSI_ConnectionStaticInfoEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *ISCSI_ConnectionStaticInfo, err error) {
	tmp, err := instance.GetWmiInstance(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &ISCSI_ConnectionStaticInfo{
		WmiInstance: tmp,
	}
	return
}

// SetAuthType sets the value of AuthType for the instance
func (instance *ISCSI_ConnectionStaticInfo) SetPropertyAuthType(value ConnectionStaticInfo_AuthType) (err error) {
	return instance.SetProperty("AuthType", (value))
}

// GetAuthType gets the value of AuthType for the instance
func (instance *ISCSI_ConnectionStaticInfo) GetPropertyAuthType() (value ConnectionStaticInfo_AuthType, err error) {
	retValue, err := instance.GetProperty("AuthType")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(int32)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " int32 is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = ConnectionStaticInfo_AuthType(valuetmp)

	return
}

// SetCID sets the value of CID for the instance
func (instance *ISCSI_ConnectionStaticInfo) SetPropertyCID(value uint16) (err error) {
	return instance.SetProperty("CID", (value))
}

// GetCID gets the value of CID for the instance
func (instance *ISCSI_ConnectionStaticInfo) GetPropertyCID() (value uint16, err error) {
	retValue, err := instance.GetProperty("CID")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(uint16)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " uint16 is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = uint16(valuetmp)

	return
}

// SetDataIntegrity sets the value of DataIntegrity for the instance
func (instance *ISCSI_ConnectionStaticInfo) SetPropertyDataIntegrity(value ConnectionStaticInfo_DataIntegrity) (err error) {
	return instance.SetProperty("DataIntegrity", (value))
}

// GetDataIntegrity gets the value of DataIntegrity for the instance
func (instance *ISCSI_ConnectionStaticInfo) GetPropertyDataIntegrity() (value ConnectionStaticInfo_DataIntegrity, err error) {
	retValue, err := instance.GetProperty("DataIntegrity")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(int32)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " int32 is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = ConnectionStaticInfo_DataIntegrity(valuetmp)

	return
}

// SetEstimatedThroughput sets the value of EstimatedThroughput for the instance
func (instance *ISCSI_ConnectionStaticInfo) SetPropertyEstimatedThroughput(value uint64) (err error) {
	return instance.SetProperty("EstimatedThroughput", (value))
}

// GetEstimatedThroughput gets the value of EstimatedThroughput for the instance
func (instance *ISCSI_ConnectionStaticInfo) GetPropertyEstimatedThroughput() (value uint64, err error) {
	retValue, err := instance.GetProperty("EstimatedThroughput")
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

// SetHeaderIntegrity sets the value of HeaderIntegrity for the instance
func (instance *ISCSI_ConnectionStaticInfo) SetPropertyHeaderIntegrity(value ConnectionStaticInfo_HeaderIntegrity) (err error) {
	return instance.SetProperty("HeaderIntegrity", (value))
}

// GetHeaderIntegrity gets the value of HeaderIntegrity for the instance
func (instance *ISCSI_ConnectionStaticInfo) GetPropertyHeaderIntegrity() (value ConnectionStaticInfo_HeaderIntegrity, err error) {
	retValue, err := instance.GetProperty("HeaderIntegrity")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(int32)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " int32 is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = ConnectionStaticInfo_HeaderIntegrity(valuetmp)

	return
}

// SetLocalAddr sets the value of LocalAddr for the instance
func (instance *ISCSI_ConnectionStaticInfo) SetPropertyLocalAddr(value ISCSI_IP_Address) (err error) {
	return instance.SetProperty("LocalAddr", (value))
}

// GetLocalAddr gets the value of LocalAddr for the instance
func (instance *ISCSI_ConnectionStaticInfo) GetPropertyLocalAddr() (value ISCSI_IP_Address, err error) {
	retValue, err := instance.GetProperty("LocalAddr")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(ISCSI_IP_Address)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " ISCSI_IP_Address is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = ISCSI_IP_Address(valuetmp)

	return
}

// SetLocalPort sets the value of LocalPort for the instance
func (instance *ISCSI_ConnectionStaticInfo) SetPropertyLocalPort(value uint32) (err error) {
	return instance.SetProperty("LocalPort", (value))
}

// GetLocalPort gets the value of LocalPort for the instance
func (instance *ISCSI_ConnectionStaticInfo) GetPropertyLocalPort() (value uint32, err error) {
	retValue, err := instance.GetProperty("LocalPort")
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

// SetMaxDatagramSize sets the value of MaxDatagramSize for the instance
func (instance *ISCSI_ConnectionStaticInfo) SetPropertyMaxDatagramSize(value uint32) (err error) {
	return instance.SetProperty("MaxDatagramSize", (value))
}

// GetMaxDatagramSize gets the value of MaxDatagramSize for the instance
func (instance *ISCSI_ConnectionStaticInfo) GetPropertyMaxDatagramSize() (value uint32, err error) {
	retValue, err := instance.GetProperty("MaxDatagramSize")
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

// SetMaxRecvDataSegmentLength sets the value of MaxRecvDataSegmentLength for the instance
func (instance *ISCSI_ConnectionStaticInfo) SetPropertyMaxRecvDataSegmentLength(value uint32) (err error) {
	return instance.SetProperty("MaxRecvDataSegmentLength", (value))
}

// GetMaxRecvDataSegmentLength gets the value of MaxRecvDataSegmentLength for the instance
func (instance *ISCSI_ConnectionStaticInfo) GetPropertyMaxRecvDataSegmentLength() (value uint32, err error) {
	retValue, err := instance.GetProperty("MaxRecvDataSegmentLength")
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

// SetProtocol sets the value of Protocol for the instance
func (instance *ISCSI_ConnectionStaticInfo) SetPropertyProtocol(value ConnectionStaticInfo_Protocol) (err error) {
	return instance.SetProperty("Protocol", (value))
}

// GetProtocol gets the value of Protocol for the instance
func (instance *ISCSI_ConnectionStaticInfo) GetPropertyProtocol() (value ConnectionStaticInfo_Protocol, err error) {
	retValue, err := instance.GetProperty("Protocol")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(int32)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " int32 is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = ConnectionStaticInfo_Protocol(valuetmp)

	return
}

// SetRemoteAddr sets the value of RemoteAddr for the instance
func (instance *ISCSI_ConnectionStaticInfo) SetPropertyRemoteAddr(value ISCSI_IP_Address) (err error) {
	return instance.SetProperty("RemoteAddr", (value))
}

// GetRemoteAddr gets the value of RemoteAddr for the instance
func (instance *ISCSI_ConnectionStaticInfo) GetPropertyRemoteAddr() (value ISCSI_IP_Address, err error) {
	retValue, err := instance.GetProperty("RemoteAddr")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(ISCSI_IP_Address)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " ISCSI_IP_Address is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = ISCSI_IP_Address(valuetmp)

	return
}

// SetRemotePort sets the value of RemotePort for the instance
func (instance *ISCSI_ConnectionStaticInfo) SetPropertyRemotePort(value uint32) (err error) {
	return instance.SetProperty("RemotePort", (value))
}

// GetRemotePort gets the value of RemotePort for the instance
func (instance *ISCSI_ConnectionStaticInfo) GetPropertyRemotePort() (value uint32, err error) {
	retValue, err := instance.GetProperty("RemotePort")
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

// SetReserved sets the value of Reserved for the instance
func (instance *ISCSI_ConnectionStaticInfo) SetPropertyReserved(value uint16) (err error) {
	return instance.SetProperty("Reserved", (value))
}

// GetReserved gets the value of Reserved for the instance
func (instance *ISCSI_ConnectionStaticInfo) GetPropertyReserved() (value uint16, err error) {
	retValue, err := instance.GetProperty("Reserved")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(uint16)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " uint16 is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = uint16(valuetmp)

	return
}

// SetState sets the value of State for the instance
func (instance *ISCSI_ConnectionStaticInfo) SetPropertyState(value ConnectionStaticInfo_State) (err error) {
	return instance.SetProperty("State", (value))
}

// GetState gets the value of State for the instance
func (instance *ISCSI_ConnectionStaticInfo) GetPropertyState() (value ConnectionStaticInfo_State, err error) {
	retValue, err := instance.GetProperty("State")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(int32)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " int32 is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = ConnectionStaticInfo_State(valuetmp)

	return
}

// SetUniqueConnectionId sets the value of UniqueConnectionId for the instance
func (instance *ISCSI_ConnectionStaticInfo) SetPropertyUniqueConnectionId(value uint64) (err error) {
	return instance.SetProperty("UniqueConnectionId", (value))
}

// GetUniqueConnectionId gets the value of UniqueConnectionId for the instance
func (instance *ISCSI_ConnectionStaticInfo) GetPropertyUniqueConnectionId() (value uint64, err error) {
	retValue, err := instance.GetProperty("UniqueConnectionId")
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
