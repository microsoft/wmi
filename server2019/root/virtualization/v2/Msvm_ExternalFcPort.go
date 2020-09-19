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

// Msvm_ExternalFcPort struct
type Msvm_ExternalFcPort struct {
	*CIM_FCPort

	//
	IsHyperVCapable bool

	//
	WWNN string

	//
	WWPN string
}

func NewMsvm_ExternalFcPortEx1(instance *cim.WmiInstance) (newInstance *Msvm_ExternalFcPort, err error) {
	tmp, err := NewCIM_FCPortEx1(instance)

	if err != nil {
		return
	}
	newInstance = &Msvm_ExternalFcPort{
		CIM_FCPort: tmp,
	}
	return
}

func NewMsvm_ExternalFcPortEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *Msvm_ExternalFcPort, err error) {
	tmp, err := NewCIM_FCPortEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &Msvm_ExternalFcPort{
		CIM_FCPort: tmp,
	}
	return
}

// SetIsHyperVCapable sets the value of IsHyperVCapable for the instance
func (instance *Msvm_ExternalFcPort) SetPropertyIsHyperVCapable(value bool) (err error) {
	return instance.SetProperty("IsHyperVCapable", (value))
}

// GetIsHyperVCapable gets the value of IsHyperVCapable for the instance
func (instance *Msvm_ExternalFcPort) GetPropertyIsHyperVCapable() (value bool, err error) {
	retValue, err := instance.GetProperty("IsHyperVCapable")
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

// SetWWNN sets the value of WWNN for the instance
func (instance *Msvm_ExternalFcPort) SetPropertyWWNN(value string) (err error) {
	return instance.SetProperty("WWNN", (value))
}

// GetWWNN gets the value of WWNN for the instance
func (instance *Msvm_ExternalFcPort) GetPropertyWWNN() (value string, err error) {
	retValue, err := instance.GetProperty("WWNN")
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

// SetWWPN sets the value of WWPN for the instance
func (instance *Msvm_ExternalFcPort) SetPropertyWWPN(value string) (err error) {
	return instance.SetProperty("WWPN", (value))
}

// GetWWPN gets the value of WWPN for the instance
func (instance *Msvm_ExternalFcPort) GetPropertyWWPN() (value string, err error) {
	retValue, err := instance.GetProperty("WWPN")
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
