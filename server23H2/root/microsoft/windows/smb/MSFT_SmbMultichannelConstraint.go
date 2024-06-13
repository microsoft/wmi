// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 6/13/2024 using wmigen
//      Source root.Microsoft.Windows.SMB
//////////////////////////////////////////////
package smb

import (
	"github.com/microsoft/wmi/pkg/base/instance"
	"github.com/microsoft/wmi/pkg/base/query"
	"github.com/microsoft/wmi/pkg/errors"
	cim "github.com/microsoft/wmi/pkg/wmiinstance"
	"reflect"
)

// MSFT_SmbMultichannelConstraint struct
type MSFT_SmbMultichannelConstraint struct {
	*cim.WmiInstance

	//
	InterfaceAlias string

	//
	InterfaceGuid string

	//
	InterfaceIndex uint32

	//
	ServerName string
}

func NewMSFT_SmbMultichannelConstraintEx1(instance *cim.WmiInstance) (newInstance *MSFT_SmbMultichannelConstraint, err error) {
	tmp, err := instance, nil

	if err != nil {
		return
	}
	newInstance = &MSFT_SmbMultichannelConstraint{
		WmiInstance: tmp,
	}
	return
}

func NewMSFT_SmbMultichannelConstraintEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *MSFT_SmbMultichannelConstraint, err error) {
	tmp, err := instance.GetWmiInstance(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &MSFT_SmbMultichannelConstraint{
		WmiInstance: tmp,
	}
	return
}

// SetInterfaceAlias sets the value of InterfaceAlias for the instance
func (instance *MSFT_SmbMultichannelConstraint) SetPropertyInterfaceAlias(value string) (err error) {
	return instance.SetProperty("InterfaceAlias", (value))
}

// GetInterfaceAlias gets the value of InterfaceAlias for the instance
func (instance *MSFT_SmbMultichannelConstraint) GetPropertyInterfaceAlias() (value string, err error) {
	retValue, err := instance.GetProperty("InterfaceAlias")
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

// SetInterfaceGuid sets the value of InterfaceGuid for the instance
func (instance *MSFT_SmbMultichannelConstraint) SetPropertyInterfaceGuid(value string) (err error) {
	return instance.SetProperty("InterfaceGuid", (value))
}

// GetInterfaceGuid gets the value of InterfaceGuid for the instance
func (instance *MSFT_SmbMultichannelConstraint) GetPropertyInterfaceGuid() (value string, err error) {
	retValue, err := instance.GetProperty("InterfaceGuid")
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

// SetInterfaceIndex sets the value of InterfaceIndex for the instance
func (instance *MSFT_SmbMultichannelConstraint) SetPropertyInterfaceIndex(value uint32) (err error) {
	return instance.SetProperty("InterfaceIndex", (value))
}

// GetInterfaceIndex gets the value of InterfaceIndex for the instance
func (instance *MSFT_SmbMultichannelConstraint) GetPropertyInterfaceIndex() (value uint32, err error) {
	retValue, err := instance.GetProperty("InterfaceIndex")
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

// SetServerName sets the value of ServerName for the instance
func (instance *MSFT_SmbMultichannelConstraint) SetPropertyServerName(value string) (err error) {
	return instance.SetProperty("ServerName", (value))
}

// GetServerName gets the value of ServerName for the instance
func (instance *MSFT_SmbMultichannelConstraint) GetPropertyServerName() (value string, err error) {
	retValue, err := instance.GetProperty("ServerName")
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

//

// <param name="InterfaceAlias" type="string []"></param>
// <param name="InterfaceIndex" type="uint32 []"></param>
// <param name="ServerName" type="string "></param>

// <param name="Output" type="MSFT_SmbMultichannelConstraint []"></param>
// <param name="ReturnValue" type="uint32 "></param>
func (instance *MSFT_SmbMultichannelConstraint) CreateConstraint( /* IN */ ServerName string,
	/* IN */ InterfaceIndex []uint32,
	/* IN */ InterfaceAlias []string,
	/* OUT */ Output []MSFT_SmbMultichannelConstraint) (result uint32, err error) {
	retVal, err := instance.InvokeMethod("CreateConstraint", ServerName, InterfaceIndex, InterfaceAlias)
	if err != nil {
		return
	}
	retValue := retVal[0].(int32)
	result = uint32(retValue)
	return

}
