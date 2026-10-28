// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 10/28/2026 using wmigen
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

// TeeDriverAmtAdtRecPkiCertInfoWmiInfo struct
type TeeDriverAmtAdtRecPkiCertInfoWmiInfo struct {
	*cim.WmiInstance

	//
	Active bool

	//
	InstanceName string
}

func NewTeeDriverAmtAdtRecPkiCertInfoWmiInfoEx1(instance *cim.WmiInstance) (newInstance *TeeDriverAmtAdtRecPkiCertInfoWmiInfo, err error) {
	tmp, err := instance, nil

	if err != nil {
		return
	}
	newInstance = &TeeDriverAmtAdtRecPkiCertInfoWmiInfo{
		WmiInstance: tmp,
	}
	return
}

func NewTeeDriverAmtAdtRecPkiCertInfoWmiInfoEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *TeeDriverAmtAdtRecPkiCertInfoWmiInfo, err error) {
	tmp, err := instance.GetWmiInstance(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &TeeDriverAmtAdtRecPkiCertInfoWmiInfo{
		WmiInstance: tmp,
	}
	return
}

// SetActive sets the value of Active for the instance
func (instance *TeeDriverAmtAdtRecPkiCertInfoWmiInfo) SetPropertyActive(value bool) (err error) {
	return instance.SetProperty("Active", (value))
}

// GetActive gets the value of Active for the instance
func (instance *TeeDriverAmtAdtRecPkiCertInfoWmiInfo) GetPropertyActive() (value bool, err error) {
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

// SetInstanceName sets the value of InstanceName for the instance
func (instance *TeeDriverAmtAdtRecPkiCertInfoWmiInfo) SetPropertyInstanceName(value string) (err error) {
	return instance.SetProperty("InstanceName", (value))
}

// GetInstanceName gets the value of InstanceName for the instance
func (instance *TeeDriverAmtAdtRecPkiCertInfoWmiInfo) GetPropertyInstanceName() (value string, err error) {
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
