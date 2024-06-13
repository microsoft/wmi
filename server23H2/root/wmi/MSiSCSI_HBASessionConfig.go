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

// MSiSCSI_HBASessionConfig struct
type MSiSCSI_HBASessionConfig struct {
	*cim.WmiInstance

	//
	Active bool

	// maximum amount in bytes of unsolicited data an iSCSI initiator may send to the target, during the execution of a single SCSI command. This covers the immediate data (if any) and the sequence of unsolicited Data-Out PDUs (if any) that follow the command.
	FirstBurstLength uint32

	// The initiator and target negotiate support for immediate data. To turn immediate data off, the initiator or target must state its desire to do so.  ImmediateData can be turned on if both the initiator and target have ImmediateData=Yes.
	ImmediateData bool

	// The InitialR2T key is used to turn off the default use of R2T, thus allowing an initiator to start sending data to a target as if it has received an initial R2T with Buffer Offset=0 and Desired Data Transfer Length=min (FirstBurstSize, Expected Data Transfer Length).
	InitialR2T bool

	//
	InstanceName string

	// Maximum SCSI data payload in bytes in an Data-In or a solicited Data-Out iSCSI sequence.
	MaxBurstLength uint32

	// Initiator and target negotiate the maximum number of outstanding R2Ts per task, excluding any implied initial R2T that might be part of that task.  An R2T is considered outstanding until the last data PDU (with the F bit set to 1) is transferred, or a sequence reception timeout (section 6.12.1) is encountered for that data sequence.
	MaxOutstandingR2T uint32

	// Maximum data segment length in bytes they can receive in an iSCSI PDU.
	MaxRecvDataSegmentLength uint32
}

func NewMSiSCSI_HBASessionConfigEx1(instance *cim.WmiInstance) (newInstance *MSiSCSI_HBASessionConfig, err error) {
	tmp, err := instance, nil

	if err != nil {
		return
	}
	newInstance = &MSiSCSI_HBASessionConfig{
		WmiInstance: tmp,
	}
	return
}

func NewMSiSCSI_HBASessionConfigEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *MSiSCSI_HBASessionConfig, err error) {
	tmp, err := instance.GetWmiInstance(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &MSiSCSI_HBASessionConfig{
		WmiInstance: tmp,
	}
	return
}

// SetActive sets the value of Active for the instance
func (instance *MSiSCSI_HBASessionConfig) SetPropertyActive(value bool) (err error) {
	return instance.SetProperty("Active", (value))
}

// GetActive gets the value of Active for the instance
func (instance *MSiSCSI_HBASessionConfig) GetPropertyActive() (value bool, err error) {
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

// SetFirstBurstLength sets the value of FirstBurstLength for the instance
func (instance *MSiSCSI_HBASessionConfig) SetPropertyFirstBurstLength(value uint32) (err error) {
	return instance.SetProperty("FirstBurstLength", (value))
}

// GetFirstBurstLength gets the value of FirstBurstLength for the instance
func (instance *MSiSCSI_HBASessionConfig) GetPropertyFirstBurstLength() (value uint32, err error) {
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
func (instance *MSiSCSI_HBASessionConfig) SetPropertyImmediateData(value bool) (err error) {
	return instance.SetProperty("ImmediateData", (value))
}

// GetImmediateData gets the value of ImmediateData for the instance
func (instance *MSiSCSI_HBASessionConfig) GetPropertyImmediateData() (value bool, err error) {
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

// SetInitialR2T sets the value of InitialR2T for the instance
func (instance *MSiSCSI_HBASessionConfig) SetPropertyInitialR2T(value bool) (err error) {
	return instance.SetProperty("InitialR2T", (value))
}

// GetInitialR2T gets the value of InitialR2T for the instance
func (instance *MSiSCSI_HBASessionConfig) GetPropertyInitialR2T() (value bool, err error) {
	retValue, err := instance.GetProperty("InitialR2T")
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

// SetInstanceName sets the value of InstanceName for the instance
func (instance *MSiSCSI_HBASessionConfig) SetPropertyInstanceName(value string) (err error) {
	return instance.SetProperty("InstanceName", (value))
}

// GetInstanceName gets the value of InstanceName for the instance
func (instance *MSiSCSI_HBASessionConfig) GetPropertyInstanceName() (value string, err error) {
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

// SetMaxBurstLength sets the value of MaxBurstLength for the instance
func (instance *MSiSCSI_HBASessionConfig) SetPropertyMaxBurstLength(value uint32) (err error) {
	return instance.SetProperty("MaxBurstLength", (value))
}

// GetMaxBurstLength gets the value of MaxBurstLength for the instance
func (instance *MSiSCSI_HBASessionConfig) GetPropertyMaxBurstLength() (value uint32, err error) {
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

// SetMaxOutstandingR2T sets the value of MaxOutstandingR2T for the instance
func (instance *MSiSCSI_HBASessionConfig) SetPropertyMaxOutstandingR2T(value uint32) (err error) {
	return instance.SetProperty("MaxOutstandingR2T", (value))
}

// GetMaxOutstandingR2T gets the value of MaxOutstandingR2T for the instance
func (instance *MSiSCSI_HBASessionConfig) GetPropertyMaxOutstandingR2T() (value uint32, err error) {
	retValue, err := instance.GetProperty("MaxOutstandingR2T")
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
func (instance *MSiSCSI_HBASessionConfig) SetPropertyMaxRecvDataSegmentLength(value uint32) (err error) {
	return instance.SetProperty("MaxRecvDataSegmentLength", (value))
}

// GetMaxRecvDataSegmentLength gets the value of MaxRecvDataSegmentLength for the instance
func (instance *MSiSCSI_HBASessionConfig) GetPropertyMaxRecvDataSegmentLength() (value uint32, err error) {
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
