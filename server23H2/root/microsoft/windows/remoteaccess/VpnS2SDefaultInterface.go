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

// VpnS2SDefaultInterface struct
type VpnS2SDefaultInterface struct {
	*VpnS2SInterface

	//
	EncryptionType string
}

func NewVpnS2SDefaultInterfaceEx1(instance *cim.WmiInstance) (newInstance *VpnS2SDefaultInterface, err error) {
	tmp, err := NewVpnS2SInterfaceEx1(instance)

	if err != nil {
		return
	}
	newInstance = &VpnS2SDefaultInterface{
		VpnS2SInterface: tmp,
	}
	return
}

func NewVpnS2SDefaultInterfaceEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *VpnS2SDefaultInterface, err error) {
	tmp, err := NewVpnS2SInterfaceEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &VpnS2SDefaultInterface{
		VpnS2SInterface: tmp,
	}
	return
}

// SetEncryptionType sets the value of EncryptionType for the instance
func (instance *VpnS2SDefaultInterface) SetPropertyEncryptionType(value string) (err error) {
	return instance.SetProperty("EncryptionType", (value))
}

// GetEncryptionType gets the value of EncryptionType for the instance
func (instance *VpnS2SDefaultInterface) GetPropertyEncryptionType() (value string, err error) {
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
