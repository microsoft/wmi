// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 6/12/2024 using wmigen
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

// MSiSCSI_DiscoveryConfig struct
type MSiSCSI_DiscoveryConfig struct {
	*cim.WmiInstance

	//
	Active bool

	// TRUE if adapter should perform automatic discovery of iSNS server.
	AutomaticiSNSDiscovery bool

	// Default initiator name for registering with iSNS.
	InitiatorName string

	//
	InstanceName string

	// If AutomaticiSNSDiscovery is FALSE then this contains the fixed addresses of iSNS servers
	iSNSServer ISCSI_IP_Address

	// TRUE if adapter should perform target discovery via iSNS.
	PerformiSNSDiscovery bool

	// TRUE if adapter should perform target discovery via SLP.
	PerformSLPDiscovery bool
}

func NewMSiSCSI_DiscoveryConfigEx1(instance *cim.WmiInstance) (newInstance *MSiSCSI_DiscoveryConfig, err error) {
	tmp, err := instance, nil

	if err != nil {
		return
	}
	newInstance = &MSiSCSI_DiscoveryConfig{
		WmiInstance: tmp,
	}
	return
}

func NewMSiSCSI_DiscoveryConfigEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *MSiSCSI_DiscoveryConfig, err error) {
	tmp, err := instance.GetWmiInstance(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &MSiSCSI_DiscoveryConfig{
		WmiInstance: tmp,
	}
	return
}

// SetActive sets the value of Active for the instance
func (instance *MSiSCSI_DiscoveryConfig) SetPropertyActive(value bool) (err error) {
	return instance.SetProperty("Active", (value))
}

// GetActive gets the value of Active for the instance
func (instance *MSiSCSI_DiscoveryConfig) GetPropertyActive() (value bool, err error) {
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

// SetAutomaticiSNSDiscovery sets the value of AutomaticiSNSDiscovery for the instance
func (instance *MSiSCSI_DiscoveryConfig) SetPropertyAutomaticiSNSDiscovery(value bool) (err error) {
	return instance.SetProperty("AutomaticiSNSDiscovery", (value))
}

// GetAutomaticiSNSDiscovery gets the value of AutomaticiSNSDiscovery for the instance
func (instance *MSiSCSI_DiscoveryConfig) GetPropertyAutomaticiSNSDiscovery() (value bool, err error) {
	retValue, err := instance.GetProperty("AutomaticiSNSDiscovery")
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

// SetInitiatorName sets the value of InitiatorName for the instance
func (instance *MSiSCSI_DiscoveryConfig) SetPropertyInitiatorName(value string) (err error) {
	return instance.SetProperty("InitiatorName", (value))
}

// GetInitiatorName gets the value of InitiatorName for the instance
func (instance *MSiSCSI_DiscoveryConfig) GetPropertyInitiatorName() (value string, err error) {
	retValue, err := instance.GetProperty("InitiatorName")
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

// SetInstanceName sets the value of InstanceName for the instance
func (instance *MSiSCSI_DiscoveryConfig) SetPropertyInstanceName(value string) (err error) {
	return instance.SetProperty("InstanceName", (value))
}

// GetInstanceName gets the value of InstanceName for the instance
func (instance *MSiSCSI_DiscoveryConfig) GetPropertyInstanceName() (value string, err error) {
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

// SetiSNSServer sets the value of iSNSServer for the instance
func (instance *MSiSCSI_DiscoveryConfig) SetPropertyiSNSServer(value ISCSI_IP_Address) (err error) {
	return instance.SetProperty("iSNSServer", (value))
}

// GetiSNSServer gets the value of iSNSServer for the instance
func (instance *MSiSCSI_DiscoveryConfig) GetPropertyiSNSServer() (value ISCSI_IP_Address, err error) {
	retValue, err := instance.GetProperty("iSNSServer")
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

// SetPerformiSNSDiscovery sets the value of PerformiSNSDiscovery for the instance
func (instance *MSiSCSI_DiscoveryConfig) SetPropertyPerformiSNSDiscovery(value bool) (err error) {
	return instance.SetProperty("PerformiSNSDiscovery", (value))
}

// GetPerformiSNSDiscovery gets the value of PerformiSNSDiscovery for the instance
func (instance *MSiSCSI_DiscoveryConfig) GetPropertyPerformiSNSDiscovery() (value bool, err error) {
	retValue, err := instance.GetProperty("PerformiSNSDiscovery")
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

// SetPerformSLPDiscovery sets the value of PerformSLPDiscovery for the instance
func (instance *MSiSCSI_DiscoveryConfig) SetPropertyPerformSLPDiscovery(value bool) (err error) {
	return instance.SetProperty("PerformSLPDiscovery", (value))
}

// GetPerformSLPDiscovery gets the value of PerformSLPDiscovery for the instance
func (instance *MSiSCSI_DiscoveryConfig) GetPropertyPerformSLPDiscovery() (value bool, err error) {
	retValue, err := instance.GetProperty("PerformSLPDiscovery")
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
