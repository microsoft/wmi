// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

// Author:
//
//	Auto Generated on 6/6/2024 using wmigen
//	Source root.virtualization.v2
//
// ////////////////////////////////////////////
package v2

import (
	"github.com/microsoft/wmi/pkg/base/instance"
	"github.com/microsoft/wmi/pkg/base/query"
	"github.com/microsoft/wmi/pkg/errors"
	cim "github.com/microsoft/wmi/pkg/wmiinstance"
	"reflect"
)

// Msvm_NetworkConnectionDiagnosticInformation struct
type Msvm_NetworkConnectionDiagnosticInformation struct {
	*cim.WmiInstance

	// The round trip time for the Ping request.
	RoundTripTime uint32
}

func NewMsvm_NetworkConnectionDiagnosticInformationEx1(instance *cim.WmiInstance) (newInstance *Msvm_NetworkConnectionDiagnosticInformation, err error) {
	tmp, err := instance, nil

	if err != nil {
		return
	}
	newInstance = &Msvm_NetworkConnectionDiagnosticInformation{
		WmiInstance: tmp,
	}
	return
}

func NewMsvm_NetworkConnectionDiagnosticInformationEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *Msvm_NetworkConnectionDiagnosticInformation, err error) {
	tmp, err := instance.GetWmiInstance(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &Msvm_NetworkConnectionDiagnosticInformation{
		WmiInstance: tmp,
	}
	return
}

// SetRoundTripTime sets the value of RoundTripTime for the instance
func (instance *Msvm_NetworkConnectionDiagnosticInformation) SetPropertyRoundTripTime(value uint32) (err error) {
	return instance.SetProperty("RoundTripTime", (value))
}

// GetRoundTripTime gets the value of RoundTripTime for the instance
func (instance *Msvm_NetworkConnectionDiagnosticInformation) GetPropertyRoundTripTime() (value uint32, err error) {
	retValue, err := instance.GetProperty("RoundTripTime")
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
