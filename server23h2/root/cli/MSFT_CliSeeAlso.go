// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 6/14/2024 using wmigen
//      Source root.Cli
//////////////////////////////////////////////
package cli

import (
	"github.com/microsoft/wmi/pkg/base/instance"
	"github.com/microsoft/wmi/pkg/base/query"
	"github.com/microsoft/wmi/pkg/errors"
	cim "github.com/microsoft/wmi/pkg/wmiinstance"
	"reflect"
)

// MSFT_CliSeeAlso struct
type MSFT_CliSeeAlso struct {
	*cim.WmiInstance

	//
	Description string

	//
	Original MSFT_CliAlias

	//
	Related MSFT_CliAlias
}

func NewMSFT_CliSeeAlsoEx1(instance *cim.WmiInstance) (newInstance *MSFT_CliSeeAlso, err error) {
	tmp, err := instance, nil

	if err != nil {
		return
	}
	newInstance = &MSFT_CliSeeAlso{
		WmiInstance: tmp,
	}
	return
}

func NewMSFT_CliSeeAlsoEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *MSFT_CliSeeAlso, err error) {
	tmp, err := instance.GetWmiInstance(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &MSFT_CliSeeAlso{
		WmiInstance: tmp,
	}
	return
}

// SetDescription sets the value of Description for the instance
func (instance *MSFT_CliSeeAlso) SetPropertyDescription(value string) (err error) {
	return instance.SetProperty("Description", (value))
}

// GetDescription gets the value of Description for the instance
func (instance *MSFT_CliSeeAlso) GetPropertyDescription() (value string, err error) {
	retValue, err := instance.GetProperty("Description")
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

// SetOriginal sets the value of Original for the instance
func (instance *MSFT_CliSeeAlso) SetPropertyOriginal(value MSFT_CliAlias) (err error) {
	return instance.SetProperty("Original", (value))
}

// GetOriginal gets the value of Original for the instance
func (instance *MSFT_CliSeeAlso) GetPropertyOriginal() (value MSFT_CliAlias, err error) {
	retValue, err := instance.GetProperty("Original")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(MSFT_CliAlias)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " MSFT_CliAlias is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = MSFT_CliAlias(valuetmp)

	return
}

// SetRelated sets the value of Related for the instance
func (instance *MSFT_CliSeeAlso) SetPropertyRelated(value MSFT_CliAlias) (err error) {
	return instance.SetProperty("Related", (value))
}

// GetRelated gets the value of Related for the instance
func (instance *MSFT_CliSeeAlso) GetPropertyRelated() (value MSFT_CliAlias, err error) {
	retValue, err := instance.GetProperty("Related")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(MSFT_CliAlias)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " MSFT_CliAlias is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = MSFT_CliAlias(valuetmp)

	return
}
