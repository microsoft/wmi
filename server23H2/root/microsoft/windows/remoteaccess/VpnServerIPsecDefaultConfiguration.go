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

// VpnServerIPsecDefaultConfiguration struct
type VpnServerIPsecDefaultConfiguration struct {
	*VpnServerIPsecConfiguration

	//
	EncryptionType string
}

func NewVpnServerIPsecDefaultConfigurationEx1(instance *cim.WmiInstance) (newInstance *VpnServerIPsecDefaultConfiguration, err error) {
	tmp, err := NewVpnServerIPsecConfigurationEx1(instance)

	if err != nil {
		return
	}
	newInstance = &VpnServerIPsecDefaultConfiguration{
		VpnServerIPsecConfiguration: tmp,
	}
	return
}

func NewVpnServerIPsecDefaultConfigurationEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *VpnServerIPsecDefaultConfiguration, err error) {
	tmp, err := NewVpnServerIPsecConfigurationEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &VpnServerIPsecDefaultConfiguration{
		VpnServerIPsecConfiguration: tmp,
	}
	return
}

// SetEncryptionType sets the value of EncryptionType for the instance
func (instance *VpnServerIPsecDefaultConfiguration) SetPropertyEncryptionType(value string) (err error) {
	return instance.SetProperty("EncryptionType", (value))
}

// GetEncryptionType gets the value of EncryptionType for the instance
func (instance *VpnServerIPsecDefaultConfiguration) GetPropertyEncryptionType() (value string, err error) {
	retValue, err := instance.GetProperty("EncryptionType")
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
