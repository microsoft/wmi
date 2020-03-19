// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 3/19/2020 using wmigen
//      Source root.Cli
//////////////////////////////////////////////
package cli

import (
	"github.com/microsoft/wmi/pkg/base/instance"
	"github.com/microsoft/wmi/pkg/base/query"
	cim "github.com/microsoft/wmi/pkg/wmiinstance"
)

// MSFT_CliQualifier struct
type MSFT_CliQualifier struct {
	*cim.WmiInstance

	//
	Name string

	//
	QualifierValue []string
}

func NewMSFT_CliQualifierEx1(instance *cim.WmiInstance) (newInstance *MSFT_CliQualifier, err error) {
	tmp, err := instance, nil

	if err != nil {
		return
	}
	newInstance = &MSFT_CliQualifier{
		WmiInstance: tmp,
	}
	return
}

func NewMSFT_CliQualifierEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *MSFT_CliQualifier, err error) {
	tmp, err := instance.GetWmiInstance(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &MSFT_CliQualifier{
		WmiInstance: tmp,
	}
	return
}

// SetName sets the value of Name for the instance
func (instance *MSFT_CliQualifier) SetPropertyName(value string) (err error) {
	return instance.SetProperty("Name", value)
}

// GetName gets the value of Name for the instance
func (instance *MSFT_CliQualifier) GetPropertyName() (value string, err error) {
	retValue, err := instance.GetProperty("Name")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetQualifierValue sets the value of QualifierValue for the instance
func (instance *MSFT_CliQualifier) SetPropertyQualifierValue(value []string) (err error) {
	return instance.SetProperty("QualifierValue", value)
}

// GetQualifierValue gets the value of QualifierValue for the instance
func (instance *MSFT_CliQualifier) GetPropertyQualifierValue() (value []string, err error) {
	retValue, err := instance.GetProperty("QualifierValue")
	if err != nil {
		return
	}
	value, ok := retValue.([]string)
	if !ok {
		// TODO: Set an error
	}
	return
}
