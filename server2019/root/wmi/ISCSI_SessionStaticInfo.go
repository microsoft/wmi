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
	"github.com/microsoft/wmi/pkg/base/instance"
	"github.com/microsoft/wmi/pkg/base/query"
	"github.com/microsoft/wmi/pkg/errors"
	cim "github.com/microsoft/wmi/pkg/wmiinstance"
	"reflect"
)

// ISCSI_SessionStaticInfo struct
type ISCSI_SessionStaticInfo struct {
	*cim.WmiInstance

	// The number of connections that currently belong to this session
	ConnectionCount uint16

	// List of ISCSI_ConnectionStaticInfo.  ConnectionCount specifies the number of elements in the array. NOTE: This is a variable length array.
	ConnectionsList []ISCSI_ConnectionStaticInfo

	// If FALSE indicates that data PDUs within sequences may be in any order. If TRUE indicates that data PDUs within sequences must be at continuously increasing addresses, with no gaps or overlay between PDUs.
	DataPduInOrder bool

	// If FALSE indicates that data PDU Sequences may be transferred in any order.  If TRUE indicates that data PDU sequences must be transferred using continuously increasing offsets, except during error recovery.
	DataSequenceInOrder bool

	// The level of error recovery negotiated between the initiator and the target.
	ErrorRecoveryLevel uint8

	// The maximum length supported for unsolicited data sent within this session
	FirstBurstLength uint32

	// If TRUE indicates whether the initiator and target have agreed to support immediate commands on this session.
	ImmediateData bool

	// If TRUE, the initiator must wait for an R2T before sending data to the target.  If FALSE, the initiator may send data immediately, within limits set by FirstBurstSize and the expected data transfer length of the request.
	InitialR2t bool

	// Initiator node name used to establish the session
	InitiatoriSCSIName string

	// Initiator-defined portion of the iSCSI Session ID
	ISID []uint8

	// The maximum number of bytes which can be sent within a single sequence of Data-In or Data-Out PDUs
	MaxBurstLength uint32

	// The maximum number of connections that will be allowed within this session
	MaxConnections uint32

	// The maximum number of outstanding request-to-transmit (R2T) per task within this session
	MaxOutstandingR2t uint32

	// iSCSI node name of the target
	TargetiSCSIName string

	// Target-defined portion of the iSCSI Session ID
	TSID uint16

	// **typedef** Type of iSCSI session
	Type SessionStaticInfo_Type

	// A uniquely generated session ID, it is the same id returned by the LoginToTarget method.  Do not confuse this with ISID or SSID.
	UniqueSessionId uint64
}

func NewISCSI_SessionStaticInfoEx1(instance *cim.WmiInstance) (newInstance *ISCSI_SessionStaticInfo, err error) {
	tmp, err := instance, nil

	if err != nil {
		return
	}
	newInstance = &ISCSI_SessionStaticInfo{
		WmiInstance: tmp,
	}
	return
}

func NewISCSI_SessionStaticInfoEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *ISCSI_SessionStaticInfo, err error) {
	tmp, err := instance.GetWmiInstance(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &ISCSI_SessionStaticInfo{
		WmiInstance: tmp,
	}
	return
}

// SetConnectionCount sets the value of ConnectionCount for the instance
func (instance *ISCSI_SessionStaticInfo) SetPropertyConnectionCount(value uint16) (err error) {
	return instance.SetProperty("ConnectionCount", (value))
}

// GetConnectionCount gets the value of ConnectionCount for the instance
func (instance *ISCSI_SessionStaticInfo) GetPropertyConnectionCount() (value uint16, err error) {
	retValue, err := instance.GetProperty("ConnectionCount")
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

// SetConnectionsList sets the value of ConnectionsList for the instance
func (instance *ISCSI_SessionStaticInfo) SetPropertyConnectionsList(value []ISCSI_ConnectionStaticInfo) (err error) {
	return instance.SetProperty("ConnectionsList", (value))
}

// GetConnectionsList gets the value of ConnectionsList for the instance
func (instance *ISCSI_SessionStaticInfo) GetPropertyConnectionsList() (value []ISCSI_ConnectionStaticInfo, err error) {
	retValue, err := instance.GetProperty("ConnectionsList")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	for _, interfaceValue := range retValue.([]interface{}) {
		valuetmp, ok := interfaceValue.(ISCSI_ConnectionStaticInfo)
		if !ok {
			err = errors.Wrapf(errors.InvalidType, " ISCSI_ConnectionStaticInfo is Invalid. Expected %s", reflect.TypeOf(interfaceValue))
			return
		}
		value = append(value, ISCSI_ConnectionStaticInfo(valuetmp))
	}

	return
}

// SetDataPduInOrder sets the value of DataPduInOrder for the instance
func (instance *ISCSI_SessionStaticInfo) SetPropertyDataPduInOrder(value bool) (err error) {
	return instance.SetProperty("DataPduInOrder", (value))
}

// GetDataPduInOrder gets the value of DataPduInOrder for the instance
func (instance *ISCSI_SessionStaticInfo) GetPropertyDataPduInOrder() (value bool, err error) {
	retValue, err := instance.GetProperty("DataPduInOrder")
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

// SetDataSequenceInOrder sets the value of DataSequenceInOrder for the instance
func (instance *ISCSI_SessionStaticInfo) SetPropertyDataSequenceInOrder(value bool) (err error) {
	return instance.SetProperty("DataSequenceInOrder", (value))
}

// GetDataSequenceInOrder gets the value of DataSequenceInOrder for the instance
func (instance *ISCSI_SessionStaticInfo) GetPropertyDataSequenceInOrder() (value bool, err error) {
	retValue, err := instance.GetProperty("DataSequenceInOrder")
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

// SetErrorRecoveryLevel sets the value of ErrorRecoveryLevel for the instance
func (instance *ISCSI_SessionStaticInfo) SetPropertyErrorRecoveryLevel(value uint8) (err error) {
	return instance.SetProperty("ErrorRecoveryLevel", (value))
}

// GetErrorRecoveryLevel gets the value of ErrorRecoveryLevel for the instance
func (instance *ISCSI_SessionStaticInfo) GetPropertyErrorRecoveryLevel() (value uint8, err error) {
	retValue, err := instance.GetProperty("ErrorRecoveryLevel")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(uint8)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " uint8 is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = uint8(valuetmp)

	return
}

// SetFirstBurstLength sets the value of FirstBurstLength for the instance
func (instance *ISCSI_SessionStaticInfo) SetPropertyFirstBurstLength(value uint32) (err error) {
	return instance.SetProperty("FirstBurstLength", (value))
}

// GetFirstBurstLength gets the value of FirstBurstLength for the instance
func (instance *ISCSI_SessionStaticInfo) GetPropertyFirstBurstLength() (value uint32, err error) {
	retValue, err := instance.GetProperty("FirstBurstLength")
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

// SetImmediateData sets the value of ImmediateData for the instance
func (instance *ISCSI_SessionStaticInfo) SetPropertyImmediateData(value bool) (err error) {
	return instance.SetProperty("ImmediateData", (value))
}

// GetImmediateData gets the value of ImmediateData for the instance
func (instance *ISCSI_SessionStaticInfo) GetPropertyImmediateData() (value bool, err error) {
	retValue, err := instance.GetProperty("ImmediateData")
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

// SetInitialR2t sets the value of InitialR2t for the instance
func (instance *ISCSI_SessionStaticInfo) SetPropertyInitialR2t(value bool) (err error) {
	return instance.SetProperty("InitialR2t", (value))
}

// GetInitialR2t gets the value of InitialR2t for the instance
func (instance *ISCSI_SessionStaticInfo) GetPropertyInitialR2t() (value bool, err error) {
	retValue, err := instance.GetProperty("InitialR2t")
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

// SetInitiatoriSCSIName sets the value of InitiatoriSCSIName for the instance
func (instance *ISCSI_SessionStaticInfo) SetPropertyInitiatoriSCSIName(value string) (err error) {
	return instance.SetProperty("InitiatoriSCSIName", (value))
}

// GetInitiatoriSCSIName gets the value of InitiatoriSCSIName for the instance
func (instance *ISCSI_SessionStaticInfo) GetPropertyInitiatoriSCSIName() (value string, err error) {
	retValue, err := instance.GetProperty("InitiatoriSCSIName")
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

// SetISID sets the value of ISID for the instance
func (instance *ISCSI_SessionStaticInfo) SetPropertyISID(value []uint8) (err error) {
	return instance.SetProperty("ISID", (value))
}

// GetISID gets the value of ISID for the instance
func (instance *ISCSI_SessionStaticInfo) GetPropertyISID() (value []uint8, err error) {
	retValue, err := instance.GetProperty("ISID")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	for _, interfaceValue := range retValue.([]interface{}) {
		valuetmp, ok := interfaceValue.(uint8)
		if !ok {
			err = errors.Wrapf(errors.InvalidType, " uint8 is Invalid. Expected %s", reflect.TypeOf(interfaceValue))
			return
		}
		value = append(value, uint8(valuetmp))
	}

	return
}

// SetMaxBurstLength sets the value of MaxBurstLength for the instance
func (instance *ISCSI_SessionStaticInfo) SetPropertyMaxBurstLength(value uint32) (err error) {
	return instance.SetProperty("MaxBurstLength", (value))
}

// GetMaxBurstLength gets the value of MaxBurstLength for the instance
func (instance *ISCSI_SessionStaticInfo) GetPropertyMaxBurstLength() (value uint32, err error) {
	retValue, err := instance.GetProperty("MaxBurstLength")
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

// SetMaxConnections sets the value of MaxConnections for the instance
func (instance *ISCSI_SessionStaticInfo) SetPropertyMaxConnections(value uint32) (err error) {
	return instance.SetProperty("MaxConnections", (value))
}

// GetMaxConnections gets the value of MaxConnections for the instance
func (instance *ISCSI_SessionStaticInfo) GetPropertyMaxConnections() (value uint32, err error) {
	retValue, err := instance.GetProperty("MaxConnections")
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

// SetMaxOutstandingR2t sets the value of MaxOutstandingR2t for the instance
func (instance *ISCSI_SessionStaticInfo) SetPropertyMaxOutstandingR2t(value uint32) (err error) {
	return instance.SetProperty("MaxOutstandingR2t", (value))
}

// GetMaxOutstandingR2t gets the value of MaxOutstandingR2t for the instance
func (instance *ISCSI_SessionStaticInfo) GetPropertyMaxOutstandingR2t() (value uint32, err error) {
	retValue, err := instance.GetProperty("MaxOutstandingR2t")
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

// SetTargetiSCSIName sets the value of TargetiSCSIName for the instance
func (instance *ISCSI_SessionStaticInfo) SetPropertyTargetiSCSIName(value string) (err error) {
	return instance.SetProperty("TargetiSCSIName", (value))
}

// GetTargetiSCSIName gets the value of TargetiSCSIName for the instance
func (instance *ISCSI_SessionStaticInfo) GetPropertyTargetiSCSIName() (value string, err error) {
	retValue, err := instance.GetProperty("TargetiSCSIName")
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

// SetTSID sets the value of TSID for the instance
func (instance *ISCSI_SessionStaticInfo) SetPropertyTSID(value uint16) (err error) {
	return instance.SetProperty("TSID", (value))
}

// GetTSID gets the value of TSID for the instance
func (instance *ISCSI_SessionStaticInfo) GetPropertyTSID() (value uint16, err error) {
	retValue, err := instance.GetProperty("TSID")
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

// SetType sets the value of Type for the instance
func (instance *ISCSI_SessionStaticInfo) SetPropertyType(value SessionStaticInfo_Type) (err error) {
	return instance.SetProperty("Type", (value))
}

// GetType gets the value of Type for the instance
func (instance *ISCSI_SessionStaticInfo) GetPropertyType() (value SessionStaticInfo_Type, err error) {
	retValue, err := instance.GetProperty("Type")
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

	value = SessionStaticInfo_Type(valuetmp)

	return
}

// SetUniqueSessionId sets the value of UniqueSessionId for the instance
func (instance *ISCSI_SessionStaticInfo) SetPropertyUniqueSessionId(value uint64) (err error) {
	return instance.SetProperty("UniqueSessionId", (value))
}

// GetUniqueSessionId gets the value of UniqueSessionId for the instance
func (instance *ISCSI_SessionStaticInfo) GetPropertyUniqueSessionId() (value uint64, err error) {
	retValue, err := instance.GetProperty("UniqueSessionId")
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
