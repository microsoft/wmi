// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 6/12/2024 using wmigen
//      Source root.CIMV2.mdm
//////////////////////////////////////////////
package mdm

import (
	"github.com/microsoft/wmi/pkg/base/instance"
	"github.com/microsoft/wmi/pkg/base/query"
	"github.com/microsoft/wmi/pkg/errors"
	cim "github.com/microsoft/wmi/pkg/wmiinstance"
	"reflect"
)

// MDM_MgmtAuthority struct
type MDM_MgmtAuthority struct {
	*cim.WmiInstance

	//
	AuthorityName string

	//
	ClientSearchCriteria string

	//
	ProvisionedCertThumbprint string

	//
	RootThumbprint string

	//
	ServerList string
}

func NewMDM_MgmtAuthorityEx1(instance *cim.WmiInstance) (newInstance *MDM_MgmtAuthority, err error) {
	tmp, err := instance, nil

	if err != nil {
		return
	}
	newInstance = &MDM_MgmtAuthority{
		WmiInstance: tmp,
	}
	return
}

func NewMDM_MgmtAuthorityEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *MDM_MgmtAuthority, err error) {
	tmp, err := instance.GetWmiInstance(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &MDM_MgmtAuthority{
		WmiInstance: tmp,
	}
	return
}

// SetAuthorityName sets the value of AuthorityName for the instance
func (instance *MDM_MgmtAuthority) SetPropertyAuthorityName(value string) (err error) {
	return instance.SetProperty("AuthorityName", (value))
}

// GetAuthorityName gets the value of AuthorityName for the instance
func (instance *MDM_MgmtAuthority) GetPropertyAuthorityName() (value string, err error) {
	retValue, err := instance.GetProperty("AuthorityName")
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

// SetClientSearchCriteria sets the value of ClientSearchCriteria for the instance
func (instance *MDM_MgmtAuthority) SetPropertyClientSearchCriteria(value string) (err error) {
	return instance.SetProperty("ClientSearchCriteria", (value))
}

// GetClientSearchCriteria gets the value of ClientSearchCriteria for the instance
func (instance *MDM_MgmtAuthority) GetPropertyClientSearchCriteria() (value string, err error) {
	retValue, err := instance.GetProperty("ClientSearchCriteria")
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

// SetProvisionedCertThumbprint sets the value of ProvisionedCertThumbprint for the instance
func (instance *MDM_MgmtAuthority) SetPropertyProvisionedCertThumbprint(value string) (err error) {
	return instance.SetProperty("ProvisionedCertThumbprint", (value))
}

// GetProvisionedCertThumbprint gets the value of ProvisionedCertThumbprint for the instance
func (instance *MDM_MgmtAuthority) GetPropertyProvisionedCertThumbprint() (value string, err error) {
	retValue, err := instance.GetProperty("ProvisionedCertThumbprint")
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

// SetRootThumbprint sets the value of RootThumbprint for the instance
func (instance *MDM_MgmtAuthority) SetPropertyRootThumbprint(value string) (err error) {
	return instance.SetProperty("RootThumbprint", (value))
}

// GetRootThumbprint gets the value of RootThumbprint for the instance
func (instance *MDM_MgmtAuthority) GetPropertyRootThumbprint() (value string, err error) {
	retValue, err := instance.GetProperty("RootThumbprint")
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

// SetServerList sets the value of ServerList for the instance
func (instance *MDM_MgmtAuthority) SetPropertyServerList(value string) (err error) {
	return instance.SetProperty("ServerList", (value))
}

// GetServerList gets the value of ServerList for the instance
func (instance *MDM_MgmtAuthority) GetPropertyServerList() (value string, err error) {
	retValue, err := instance.GetProperty("ServerList")
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

// <param name="AuthorityName" type="string "></param>

// <param name="ReturnValue" type="uint32 "></param>
func (instance *MDM_MgmtAuthority) CreateNewAuthority( /* IN */ AuthorityName string) (result uint32, err error) {
	retVal, err := instance.InvokeMethodWithReturn("CreateNewAuthority", AuthorityName)
	if err != nil {
		return
	}
	result = uint32(retVal)
	return

}
