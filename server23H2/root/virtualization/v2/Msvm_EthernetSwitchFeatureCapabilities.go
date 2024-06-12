// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 6/12/2024 using wmigen
//      Source root.virtualization.v2
//////////////////////////////////////////////
package v2

import (
	"github.com/microsoft/wmi/pkg/base/query"
	"github.com/microsoft/wmi/pkg/errors"
	cim "github.com/microsoft/wmi/pkg/wmiinstance"
	"reflect"
)

// Msvm_EthernetSwitchFeatureCapabilities struct
type Msvm_EthernetSwitchFeatureCapabilities struct {
	*CIM_Capabilities

	//
	Applicability EthernetSwitchFeatureCapabilities_Applicability

	//
	FeatureId string

	//
	Version string
}

func NewMsvm_EthernetSwitchFeatureCapabilitiesEx1(instance *cim.WmiInstance) (newInstance *Msvm_EthernetSwitchFeatureCapabilities, err error) {
	tmp, err := NewCIM_CapabilitiesEx1(instance)

	if err != nil {
		return
	}
	newInstance = &Msvm_EthernetSwitchFeatureCapabilities{
		CIM_Capabilities: tmp,
	}
	return
}

func NewMsvm_EthernetSwitchFeatureCapabilitiesEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *Msvm_EthernetSwitchFeatureCapabilities, err error) {
	tmp, err := NewCIM_CapabilitiesEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &Msvm_EthernetSwitchFeatureCapabilities{
		CIM_Capabilities: tmp,
	}
	return
}

// SetApplicability sets the value of Applicability for the instance
func (instance *Msvm_EthernetSwitchFeatureCapabilities) SetPropertyApplicability(value EthernetSwitchFeatureCapabilities_Applicability) (err error) {
	return instance.SetProperty("Applicability", (value))
}

// GetApplicability gets the value of Applicability for the instance
func (instance *Msvm_EthernetSwitchFeatureCapabilities) GetPropertyApplicability() (value EthernetSwitchFeatureCapabilities_Applicability, err error) {
	retValue, err := instance.GetProperty("Applicability")
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

	value = EthernetSwitchFeatureCapabilities_Applicability(valuetmp)

	return
}

// SetFeatureId sets the value of FeatureId for the instance
func (instance *Msvm_EthernetSwitchFeatureCapabilities) SetPropertyFeatureId(value string) (err error) {
	return instance.SetProperty("FeatureId", (value))
}

// GetFeatureId gets the value of FeatureId for the instance
func (instance *Msvm_EthernetSwitchFeatureCapabilities) GetPropertyFeatureId() (value string, err error) {
	retValue, err := instance.GetProperty("FeatureId")
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

// SetVersion sets the value of Version for the instance
func (instance *Msvm_EthernetSwitchFeatureCapabilities) SetPropertyVersion(value string) (err error) {
	return instance.SetProperty("Version", (value))
}

// GetVersion gets the value of Version for the instance
func (instance *Msvm_EthernetSwitchFeatureCapabilities) GetPropertyVersion() (value string, err error) {
	retValue, err := instance.GetProperty("Version")
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
func (instance *Msvm_EthernetSwitchFeatureCapabilities) GetRelatedEthernetSwitchHardwareOffloadSettingData() (value *cim.WmiInstance, err error) {
	return instance.GetRelated("Msvm_EthernetSwitchHardwareOffloadSettingData")
}

func (instance *Msvm_EthernetSwitchFeatureCapabilities) GetRelatedInstalledEthernetSwitchExtension() (value *cim.WmiInstance, err error) {
	return instance.GetRelated("Msvm_InstalledEthernetSwitchExtension")
}
