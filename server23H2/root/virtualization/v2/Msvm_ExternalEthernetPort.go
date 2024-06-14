// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 6/14/2024 using wmigen
//      Source root.virtualization.v2
//////////////////////////////////////////////
package v2

import (
	"github.com/microsoft/wmi/pkg/base/query"
	"github.com/microsoft/wmi/pkg/errors"
	cim "github.com/microsoft/wmi/pkg/wmiinstance"
	"reflect"
)

// Msvm_ExternalEthernetPort struct
type Msvm_ExternalEthernetPort struct {
	*CIM_EthernetPort

	//
	IsBound bool
}

func NewMsvm_ExternalEthernetPortEx1(instance *cim.WmiInstance) (newInstance *Msvm_ExternalEthernetPort, err error) {
	tmp, err := NewCIM_EthernetPortEx1(instance)

	if err != nil {
		return
	}
	newInstance = &Msvm_ExternalEthernetPort{
		CIM_EthernetPort: tmp,
	}
	return
}

func NewMsvm_ExternalEthernetPortEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *Msvm_ExternalEthernetPort, err error) {
	tmp, err := NewCIM_EthernetPortEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &Msvm_ExternalEthernetPort{
		CIM_EthernetPort: tmp,
	}
	return
}

// SetIsBound sets the value of IsBound for the instance
func (instance *Msvm_ExternalEthernetPort) SetPropertyIsBound(value bool) (err error) {
	return instance.SetProperty("IsBound", (value))
}

// GetIsBound gets the value of IsBound for the instance
func (instance *Msvm_ExternalEthernetPort) GetPropertyIsBound() (value bool, err error) {
	retValue, err := instance.GetProperty("IsBound")
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
func (instance *Msvm_ExternalEthernetPort) GetRelatedComputerSystem() (value *cim.WmiInstance, err error) {
	return instance.GetRelated("Msvm_ComputerSystem")
}

func (instance *Msvm_ExternalEthernetPort) GetRelatedExternalEthernetPort() (value *cim.WmiInstance, err error) {
	return instance.GetRelated("Msvm_ExternalEthernetPort")
}

func (instance *Msvm_ExternalEthernetPort) GetRelatedExternalEthernetPortCapabilities() (value *cim.WmiInstance, err error) {
	return instance.GetRelated("Msvm_ExternalEthernetPortCapabilities")
}
