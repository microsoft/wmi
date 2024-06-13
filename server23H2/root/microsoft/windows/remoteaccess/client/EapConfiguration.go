// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 6/13/2024 using wmigen
//      Source root.Microsoft.Windows.RemoteAccess.Client
//////////////////////////////////////////////
package client

import (
	"github.com/microsoft/wmi/pkg/base/instance"
	"github.com/microsoft/wmi/pkg/base/query"
	"github.com/microsoft/wmi/pkg/errors"
	cim "github.com/microsoft/wmi/pkg/wmiinstance"
	"reflect"
)

// EapConfiguration struct
type EapConfiguration struct {
	*cim.WmiInstance

	//
	EapConfigXmlStream string
}

func NewEapConfigurationEx1(instance *cim.WmiInstance) (newInstance *EapConfiguration, err error) {
	tmp, err := instance, nil

	if err != nil {
		return
	}
	newInstance = &EapConfiguration{
		WmiInstance: tmp,
	}
	return
}

func NewEapConfigurationEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *EapConfiguration, err error) {
	tmp, err := instance.GetWmiInstance(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &EapConfiguration{
		WmiInstance: tmp,
	}
	return
}

// SetEapConfigXmlStream sets the value of EapConfigXmlStream for the instance
func (instance *EapConfiguration) SetPropertyEapConfigXmlStream(value string) (err error) {
	return instance.SetProperty("EapConfigXmlStream", (value))
}

// GetEapConfigXmlStream gets the value of EapConfigXmlStream for the instance
func (instance *EapConfiguration) GetPropertyEapConfigXmlStream() (value string, err error) {
	retValue, err := instance.GetProperty("EapConfigXmlStream")
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
