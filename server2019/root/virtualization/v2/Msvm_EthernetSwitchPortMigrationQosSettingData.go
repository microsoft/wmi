// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 9/18/2020 using wmigen
//      Source root.virtualization.v2
//////////////////////////////////////////////
package v2

import (
	"github.com/microsoft/wmi/pkg/base/query"
	"github.com/microsoft/wmi/pkg/errors"
	cim "github.com/microsoft/wmi/pkg/wmiinstance"
	"reflect"
)

// Msvm_EthernetSwitchPortMigrationQosSettingData struct
type Msvm_EthernetSwitchPortMigrationQosSettingData struct {
	*Msvm_EthernetSwitchPortFeatureSettingData

	//
	InboundMaximumMbps uint64

	//
	OutboundMaximumMbps uint64

	//
	OutboundReservedValue uint64

	//
	QueueId string

	//
	Switch_DefaultReservation uint64

	//
	Switch_EnableHardwareLimits bool

	//
	Switch_EnableHardwareReservations bool

	//
	Switch_EnableSoftwareReservations bool

	//
	Switch_LinkSpeedPercentage uint8

	//
	Switch_ReservationMode uint8
}

func NewMsvm_EthernetSwitchPortMigrationQosSettingDataEx1(instance *cim.WmiInstance) (newInstance *Msvm_EthernetSwitchPortMigrationQosSettingData, err error) {
	tmp, err := NewMsvm_EthernetSwitchPortFeatureSettingDataEx1(instance)

	if err != nil {
		return
	}
	newInstance = &Msvm_EthernetSwitchPortMigrationQosSettingData{
		Msvm_EthernetSwitchPortFeatureSettingData: tmp,
	}
	return
}

func NewMsvm_EthernetSwitchPortMigrationQosSettingDataEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *Msvm_EthernetSwitchPortMigrationQosSettingData, err error) {
	tmp, err := NewMsvm_EthernetSwitchPortFeatureSettingDataEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &Msvm_EthernetSwitchPortMigrationQosSettingData{
		Msvm_EthernetSwitchPortFeatureSettingData: tmp,
	}
	return
}

// SetInboundMaximumMbps sets the value of InboundMaximumMbps for the instance
func (instance *Msvm_EthernetSwitchPortMigrationQosSettingData) SetPropertyInboundMaximumMbps(value uint64) (err error) {
	return instance.SetProperty("InboundMaximumMbps", (value))
}

// GetInboundMaximumMbps gets the value of InboundMaximumMbps for the instance
func (instance *Msvm_EthernetSwitchPortMigrationQosSettingData) GetPropertyInboundMaximumMbps() (value uint64, err error) {
	retValue, err := instance.GetProperty("InboundMaximumMbps")
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

// SetOutboundMaximumMbps sets the value of OutboundMaximumMbps for the instance
func (instance *Msvm_EthernetSwitchPortMigrationQosSettingData) SetPropertyOutboundMaximumMbps(value uint64) (err error) {
	return instance.SetProperty("OutboundMaximumMbps", (value))
}

// GetOutboundMaximumMbps gets the value of OutboundMaximumMbps for the instance
func (instance *Msvm_EthernetSwitchPortMigrationQosSettingData) GetPropertyOutboundMaximumMbps() (value uint64, err error) {
	retValue, err := instance.GetProperty("OutboundMaximumMbps")
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

// SetOutboundReservedValue sets the value of OutboundReservedValue for the instance
func (instance *Msvm_EthernetSwitchPortMigrationQosSettingData) SetPropertyOutboundReservedValue(value uint64) (err error) {
	return instance.SetProperty("OutboundReservedValue", (value))
}

// GetOutboundReservedValue gets the value of OutboundReservedValue for the instance
func (instance *Msvm_EthernetSwitchPortMigrationQosSettingData) GetPropertyOutboundReservedValue() (value uint64, err error) {
	retValue, err := instance.GetProperty("OutboundReservedValue")
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

// SetQueueId sets the value of QueueId for the instance
func (instance *Msvm_EthernetSwitchPortMigrationQosSettingData) SetPropertyQueueId(value string) (err error) {
	return instance.SetProperty("QueueId", (value))
}

// GetQueueId gets the value of QueueId for the instance
func (instance *Msvm_EthernetSwitchPortMigrationQosSettingData) GetPropertyQueueId() (value string, err error) {
	retValue, err := instance.GetProperty("QueueId")
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

// SetSwitch_DefaultReservation sets the value of Switch_DefaultReservation for the instance
func (instance *Msvm_EthernetSwitchPortMigrationQosSettingData) SetPropertySwitch_DefaultReservation(value uint64) (err error) {
	return instance.SetProperty("Switch_DefaultReservation", (value))
}

// GetSwitch_DefaultReservation gets the value of Switch_DefaultReservation for the instance
func (instance *Msvm_EthernetSwitchPortMigrationQosSettingData) GetPropertySwitch_DefaultReservation() (value uint64, err error) {
	retValue, err := instance.GetProperty("Switch_DefaultReservation")
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

// SetSwitch_EnableHardwareLimits sets the value of Switch_EnableHardwareLimits for the instance
func (instance *Msvm_EthernetSwitchPortMigrationQosSettingData) SetPropertySwitch_EnableHardwareLimits(value bool) (err error) {
	return instance.SetProperty("Switch_EnableHardwareLimits", (value))
}

// GetSwitch_EnableHardwareLimits gets the value of Switch_EnableHardwareLimits for the instance
func (instance *Msvm_EthernetSwitchPortMigrationQosSettingData) GetPropertySwitch_EnableHardwareLimits() (value bool, err error) {
	retValue, err := instance.GetProperty("Switch_EnableHardwareLimits")
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

// SetSwitch_EnableHardwareReservations sets the value of Switch_EnableHardwareReservations for the instance
func (instance *Msvm_EthernetSwitchPortMigrationQosSettingData) SetPropertySwitch_EnableHardwareReservations(value bool) (err error) {
	return instance.SetProperty("Switch_EnableHardwareReservations", (value))
}

// GetSwitch_EnableHardwareReservations gets the value of Switch_EnableHardwareReservations for the instance
func (instance *Msvm_EthernetSwitchPortMigrationQosSettingData) GetPropertySwitch_EnableHardwareReservations() (value bool, err error) {
	retValue, err := instance.GetProperty("Switch_EnableHardwareReservations")
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

// SetSwitch_EnableSoftwareReservations sets the value of Switch_EnableSoftwareReservations for the instance
func (instance *Msvm_EthernetSwitchPortMigrationQosSettingData) SetPropertySwitch_EnableSoftwareReservations(value bool) (err error) {
	return instance.SetProperty("Switch_EnableSoftwareReservations", (value))
}

// GetSwitch_EnableSoftwareReservations gets the value of Switch_EnableSoftwareReservations for the instance
func (instance *Msvm_EthernetSwitchPortMigrationQosSettingData) GetPropertySwitch_EnableSoftwareReservations() (value bool, err error) {
	retValue, err := instance.GetProperty("Switch_EnableSoftwareReservations")
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

// SetSwitch_LinkSpeedPercentage sets the value of Switch_LinkSpeedPercentage for the instance
func (instance *Msvm_EthernetSwitchPortMigrationQosSettingData) SetPropertySwitch_LinkSpeedPercentage(value uint8) (err error) {
	return instance.SetProperty("Switch_LinkSpeedPercentage", (value))
}

// GetSwitch_LinkSpeedPercentage gets the value of Switch_LinkSpeedPercentage for the instance
func (instance *Msvm_EthernetSwitchPortMigrationQosSettingData) GetPropertySwitch_LinkSpeedPercentage() (value uint8, err error) {
	retValue, err := instance.GetProperty("Switch_LinkSpeedPercentage")
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

// SetSwitch_ReservationMode sets the value of Switch_ReservationMode for the instance
func (instance *Msvm_EthernetSwitchPortMigrationQosSettingData) SetPropertySwitch_ReservationMode(value uint8) (err error) {
	return instance.SetProperty("Switch_ReservationMode", (value))
}

// GetSwitch_ReservationMode gets the value of Switch_ReservationMode for the instance
func (instance *Msvm_EthernetSwitchPortMigrationQosSettingData) GetPropertySwitch_ReservationMode() (value uint8, err error) {
	retValue, err := instance.GetProperty("Switch_ReservationMode")
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
func (instance *Msvm_EthernetSwitchPortMigrationQosSettingData) GetRelatedEthernetSwitchFeatureCapabilities() (value *cim.WmiInstance, err error) {
	return instance.GetRelated("Msvm_EthernetSwitchFeatureCapabilities")
}
