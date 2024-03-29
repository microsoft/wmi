// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 3/19/2020 using wmigen
//      Source root.SDDC.Management
//////////////////////////////////////////////
package management

import (
	"github.com/microsoft/wmi/pkg/base/instance"
	"github.com/microsoft/wmi/pkg/base/query"
	cim "github.com/microsoft/wmi/pkg/wmiinstance"
)

// SDDC_BitlockerResult struct
type SDDC_BitlockerResult struct {
	*cim.WmiInstance

	//
	ErrorCode uint32

	//
	FailedPhase uint16

	//
	RecoveryPassword string
}

func NewSDDC_BitlockerResultEx1(instance *cim.WmiInstance) (newInstance *SDDC_BitlockerResult, err error) {
	tmp, err := instance, nil

	if err != nil {
		return
	}
	newInstance = &SDDC_BitlockerResult{
		WmiInstance: tmp,
	}
	return
}

func NewSDDC_BitlockerResultEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *SDDC_BitlockerResult, err error) {
	tmp, err := instance.GetWmiInstance(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &SDDC_BitlockerResult{
		WmiInstance: tmp,
	}
	return
}

// SetErrorCode sets the value of ErrorCode for the instance
func (instance *SDDC_BitlockerResult) SetPropertyErrorCode(value uint32) (err error) {
	return instance.SetProperty("ErrorCode", value)
}

// GetErrorCode gets the value of ErrorCode for the instance
func (instance *SDDC_BitlockerResult) GetPropertyErrorCode() (value uint32, err error) {
	retValue, err := instance.GetProperty("ErrorCode")
	if err != nil {
		return
	}
	value, ok := retValue.(uint32)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetFailedPhase sets the value of FailedPhase for the instance
func (instance *SDDC_BitlockerResult) SetPropertyFailedPhase(value uint16) (err error) {
	return instance.SetProperty("FailedPhase", value)
}

// GetFailedPhase gets the value of FailedPhase for the instance
func (instance *SDDC_BitlockerResult) GetPropertyFailedPhase() (value uint16, err error) {
	retValue, err := instance.GetProperty("FailedPhase")
	if err != nil {
		return
	}
	value, ok := retValue.(uint16)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetRecoveryPassword sets the value of RecoveryPassword for the instance
func (instance *SDDC_BitlockerResult) SetPropertyRecoveryPassword(value string) (err error) {
	return instance.SetProperty("RecoveryPassword", value)
}

// GetRecoveryPassword gets the value of RecoveryPassword for the instance
func (instance *SDDC_BitlockerResult) GetPropertyRecoveryPassword() (value string, err error) {
	retValue, err := instance.GetProperty("RecoveryPassword")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}
