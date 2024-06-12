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

// AcpiPct struct
type AcpiPct struct {
	*cim.WmiInstance

	//
	Control AcpiGenAddr

	//
	Status AcpiGenAddr
}

func NewAcpiPctEx1(instance *cim.WmiInstance) (newInstance *AcpiPct, err error) {
	tmp, err := instance, nil

	if err != nil {
		return
	}
	newInstance = &AcpiPct{
		WmiInstance: tmp,
	}
	return
}

func NewAcpiPctEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *AcpiPct, err error) {
	tmp, err := instance.GetWmiInstance(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &AcpiPct{
		WmiInstance: tmp,
	}
	return
}

// SetControl sets the value of Control for the instance
func (instance *AcpiPct) SetPropertyControl(value AcpiGenAddr) (err error) {
	return instance.SetProperty("Control", (value))
}

// GetControl gets the value of Control for the instance
func (instance *AcpiPct) GetPropertyControl() (value AcpiGenAddr, err error) {
	retValue, err := instance.GetProperty("Control")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(AcpiGenAddr)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " AcpiGenAddr is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = AcpiGenAddr(valuetmp)

	return
}

// SetStatus sets the value of Status for the instance
func (instance *AcpiPct) SetPropertyStatus(value AcpiGenAddr) (err error) {
	return instance.SetProperty("Status", (value))
}

// GetStatus gets the value of Status for the instance
func (instance *AcpiPct) GetPropertyStatus() (value AcpiGenAddr, err error) {
	retValue, err := instance.GetProperty("Status")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(AcpiGenAddr)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " AcpiGenAddr is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = AcpiGenAddr(valuetmp)

	return
}
