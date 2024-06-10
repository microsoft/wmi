// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

// Author:
//
//	Auto Generated on 6/6/2024 using wmigen
//	Source root.Microsoft.Windows.RemoteAccess
//
// ////////////////////////////////////////////
package remoteaccess

import (
	"github.com/microsoft/wmi/pkg/base/query"
	"github.com/microsoft/wmi/pkg/errors"
	cim "github.com/microsoft/wmi/pkg/wmiinstance"
	"reflect"
)

// RemoteAccessDA struct
type RemoteAccessDA struct {
	*RemoteAccessCommon

	//
	DAConfiguration DirectAccessConfiguration
}

func NewRemoteAccessDAEx1(instance *cim.WmiInstance) (newInstance *RemoteAccessDA, err error) {
	tmp, err := NewRemoteAccessCommonEx1(instance)

	if err != nil {
		return
	}
	newInstance = &RemoteAccessDA{
		RemoteAccessCommon: tmp,
	}
	return
}

func NewRemoteAccessDAEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *RemoteAccessDA, err error) {
	tmp, err := NewRemoteAccessCommonEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &RemoteAccessDA{
		RemoteAccessCommon: tmp,
	}
	return
}

// SetDAConfiguration sets the value of DAConfiguration for the instance
func (instance *RemoteAccessDA) SetPropertyDAConfiguration(value DirectAccessConfiguration) (err error) {
	return instance.SetProperty("DAConfiguration", (value))
}

// GetDAConfiguration gets the value of DAConfiguration for the instance
func (instance *RemoteAccessDA) GetPropertyDAConfiguration() (value DirectAccessConfiguration, err error) {
	retValue, err := instance.GetProperty("DAConfiguration")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(DirectAccessConfiguration)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " DirectAccessConfiguration is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = DirectAccessConfiguration(valuetmp)

	return
}
