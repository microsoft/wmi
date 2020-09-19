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

// Msvm_EthernetSwitchHardwareOffloadSettingData struct
type Msvm_EthernetSwitchHardwareOffloadSettingData struct {
	*Msvm_EthernetSwitchFeatureSettingData

	//
	DefaultQueueVmmqEnabled bool

	//
	DefaultQueueVmmqQueuePairs uint32

	//
	DefaultQueueVrssEnabled bool

	//
	DefaultQueueVrssExcludePrimaryProcessor bool

	//
	DefaultQueueVrssIndependentHostSpreading bool

	//
	DefaultQueueVrssMinQueuePairs uint32

	//
	DefaultQueueVrssQueueSchedulingMode uint32

	//
	SoftwareRscEnabled bool
}

func NewMsvm_EthernetSwitchHardwareOffloadSettingDataEx1(instance *cim.WmiInstance) (newInstance *Msvm_EthernetSwitchHardwareOffloadSettingData, err error) {
	tmp, err := NewMsvm_EthernetSwitchFeatureSettingDataEx1(instance)

	if err != nil {
		return
	}
	newInstance = &Msvm_EthernetSwitchHardwareOffloadSettingData{
		Msvm_EthernetSwitchFeatureSettingData: tmp,
	}
	return
}

func NewMsvm_EthernetSwitchHardwareOffloadSettingDataEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *Msvm_EthernetSwitchHardwareOffloadSettingData, err error) {
	tmp, err := NewMsvm_EthernetSwitchFeatureSettingDataEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &Msvm_EthernetSwitchHardwareOffloadSettingData{
		Msvm_EthernetSwitchFeatureSettingData: tmp,
	}
	return
}

// SetDefaultQueueVmmqEnabled sets the value of DefaultQueueVmmqEnabled for the instance
func (instance *Msvm_EthernetSwitchHardwareOffloadSettingData) SetPropertyDefaultQueueVmmqEnabled(value bool) (err error) {
	return instance.SetProperty("DefaultQueueVmmqEnabled", (value))
}

// GetDefaultQueueVmmqEnabled gets the value of DefaultQueueVmmqEnabled for the instance
func (instance *Msvm_EthernetSwitchHardwareOffloadSettingData) GetPropertyDefaultQueueVmmqEnabled() (value bool, err error) {
	retValue, err := instance.GetProperty("DefaultQueueVmmqEnabled")
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

// SetDefaultQueueVmmqQueuePairs sets the value of DefaultQueueVmmqQueuePairs for the instance
func (instance *Msvm_EthernetSwitchHardwareOffloadSettingData) SetPropertyDefaultQueueVmmqQueuePairs(value uint32) (err error) {
	return instance.SetProperty("DefaultQueueVmmqQueuePairs", (value))
}

// GetDefaultQueueVmmqQueuePairs gets the value of DefaultQueueVmmqQueuePairs for the instance
func (instance *Msvm_EthernetSwitchHardwareOffloadSettingData) GetPropertyDefaultQueueVmmqQueuePairs() (value uint32, err error) {
	retValue, err := instance.GetProperty("DefaultQueueVmmqQueuePairs")
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

// SetDefaultQueueVrssEnabled sets the value of DefaultQueueVrssEnabled for the instance
func (instance *Msvm_EthernetSwitchHardwareOffloadSettingData) SetPropertyDefaultQueueVrssEnabled(value bool) (err error) {
	return instance.SetProperty("DefaultQueueVrssEnabled", (value))
}

// GetDefaultQueueVrssEnabled gets the value of DefaultQueueVrssEnabled for the instance
func (instance *Msvm_EthernetSwitchHardwareOffloadSettingData) GetPropertyDefaultQueueVrssEnabled() (value bool, err error) {
	retValue, err := instance.GetProperty("DefaultQueueVrssEnabled")
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

// SetDefaultQueueVrssExcludePrimaryProcessor sets the value of DefaultQueueVrssExcludePrimaryProcessor for the instance
func (instance *Msvm_EthernetSwitchHardwareOffloadSettingData) SetPropertyDefaultQueueVrssExcludePrimaryProcessor(value bool) (err error) {
	return instance.SetProperty("DefaultQueueVrssExcludePrimaryProcessor", (value))
}

// GetDefaultQueueVrssExcludePrimaryProcessor gets the value of DefaultQueueVrssExcludePrimaryProcessor for the instance
func (instance *Msvm_EthernetSwitchHardwareOffloadSettingData) GetPropertyDefaultQueueVrssExcludePrimaryProcessor() (value bool, err error) {
	retValue, err := instance.GetProperty("DefaultQueueVrssExcludePrimaryProcessor")
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

// SetDefaultQueueVrssIndependentHostSpreading sets the value of DefaultQueueVrssIndependentHostSpreading for the instance
func (instance *Msvm_EthernetSwitchHardwareOffloadSettingData) SetPropertyDefaultQueueVrssIndependentHostSpreading(value bool) (err error) {
	return instance.SetProperty("DefaultQueueVrssIndependentHostSpreading", (value))
}

// GetDefaultQueueVrssIndependentHostSpreading gets the value of DefaultQueueVrssIndependentHostSpreading for the instance
func (instance *Msvm_EthernetSwitchHardwareOffloadSettingData) GetPropertyDefaultQueueVrssIndependentHostSpreading() (value bool, err error) {
	retValue, err := instance.GetProperty("DefaultQueueVrssIndependentHostSpreading")
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

// SetDefaultQueueVrssMinQueuePairs sets the value of DefaultQueueVrssMinQueuePairs for the instance
func (instance *Msvm_EthernetSwitchHardwareOffloadSettingData) SetPropertyDefaultQueueVrssMinQueuePairs(value uint32) (err error) {
	return instance.SetProperty("DefaultQueueVrssMinQueuePairs", (value))
}

// GetDefaultQueueVrssMinQueuePairs gets the value of DefaultQueueVrssMinQueuePairs for the instance
func (instance *Msvm_EthernetSwitchHardwareOffloadSettingData) GetPropertyDefaultQueueVrssMinQueuePairs() (value uint32, err error) {
	retValue, err := instance.GetProperty("DefaultQueueVrssMinQueuePairs")
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

// SetDefaultQueueVrssQueueSchedulingMode sets the value of DefaultQueueVrssQueueSchedulingMode for the instance
func (instance *Msvm_EthernetSwitchHardwareOffloadSettingData) SetPropertyDefaultQueueVrssQueueSchedulingMode(value uint32) (err error) {
	return instance.SetProperty("DefaultQueueVrssQueueSchedulingMode", (value))
}

// GetDefaultQueueVrssQueueSchedulingMode gets the value of DefaultQueueVrssQueueSchedulingMode for the instance
func (instance *Msvm_EthernetSwitchHardwareOffloadSettingData) GetPropertyDefaultQueueVrssQueueSchedulingMode() (value uint32, err error) {
	retValue, err := instance.GetProperty("DefaultQueueVrssQueueSchedulingMode")
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

// SetSoftwareRscEnabled sets the value of SoftwareRscEnabled for the instance
func (instance *Msvm_EthernetSwitchHardwareOffloadSettingData) SetPropertySoftwareRscEnabled(value bool) (err error) {
	return instance.SetProperty("SoftwareRscEnabled", (value))
}

// GetSoftwareRscEnabled gets the value of SoftwareRscEnabled for the instance
func (instance *Msvm_EthernetSwitchHardwareOffloadSettingData) GetPropertySoftwareRscEnabled() (value bool, err error) {
	retValue, err := instance.GetProperty("SoftwareRscEnabled")
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
func (instance *Msvm_EthernetSwitchHardwareOffloadSettingData) GetRelatedVirtualEthernetSwitchSettingData() (value *cim.WmiInstance, err error) {
	return instance.GetRelated("Msvm_VirtualEthernetSwitchSettingData")
}
