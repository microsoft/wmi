// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 3/19/2020 using wmigen
//      Source root.CIMV2.mdm.dmmap
//////////////////////////////////////////////
package dmmap

import (
	"github.com/microsoft/wmi/pkg/base/instance"
	"github.com/microsoft/wmi/pkg/base/query"
	cim "github.com/microsoft/wmi/pkg/wmiinstance"
)

// MDM_VPNv2_Sso03 struct
type MDM_VPNv2_Sso03 struct {
	*cim.WmiInstance

	//
	Eku string

	//
	Enabled bool

	//
	InstanceID string

	//
	IssuerHash string

	//
	ParentID string
}

func NewMDM_VPNv2_Sso03Ex1(instance *cim.WmiInstance) (newInstance *MDM_VPNv2_Sso03, err error) {
	tmp, err := instance, nil

	if err != nil {
		return
	}
	newInstance = &MDM_VPNv2_Sso03{
		WmiInstance: tmp,
	}
	return
}

func NewMDM_VPNv2_Sso03Ex6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *MDM_VPNv2_Sso03, err error) {
	tmp, err := instance.GetWmiInstance(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &MDM_VPNv2_Sso03{
		WmiInstance: tmp,
	}
	return
}

// SetEku sets the value of Eku for the instance
func (instance *MDM_VPNv2_Sso03) SetPropertyEku(value string) (err error) {
	return instance.SetProperty("Eku", value)
}

// GetEku gets the value of Eku for the instance
func (instance *MDM_VPNv2_Sso03) GetPropertyEku() (value string, err error) {
	retValue, err := instance.GetProperty("Eku")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetEnabled sets the value of Enabled for the instance
func (instance *MDM_VPNv2_Sso03) SetPropertyEnabled(value bool) (err error) {
	return instance.SetProperty("Enabled", value)
}

// GetEnabled gets the value of Enabled for the instance
func (instance *MDM_VPNv2_Sso03) GetPropertyEnabled() (value bool, err error) {
	retValue, err := instance.GetProperty("Enabled")
	if err != nil {
		return
	}
	value, ok := retValue.(bool)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetInstanceID sets the value of InstanceID for the instance
func (instance *MDM_VPNv2_Sso03) SetPropertyInstanceID(value string) (err error) {
	return instance.SetProperty("InstanceID", value)
}

// GetInstanceID gets the value of InstanceID for the instance
func (instance *MDM_VPNv2_Sso03) GetPropertyInstanceID() (value string, err error) {
	retValue, err := instance.GetProperty("InstanceID")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetIssuerHash sets the value of IssuerHash for the instance
func (instance *MDM_VPNv2_Sso03) SetPropertyIssuerHash(value string) (err error) {
	return instance.SetProperty("IssuerHash", value)
}

// GetIssuerHash gets the value of IssuerHash for the instance
func (instance *MDM_VPNv2_Sso03) GetPropertyIssuerHash() (value string, err error) {
	retValue, err := instance.GetProperty("IssuerHash")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetParentID sets the value of ParentID for the instance
func (instance *MDM_VPNv2_Sso03) SetPropertyParentID(value string) (err error) {
	return instance.SetProperty("ParentID", value)
}

// GetParentID gets the value of ParentID for the instance
func (instance *MDM_VPNv2_Sso03) GetPropertyParentID() (value string, err error) {
	retValue, err := instance.GetProperty("ParentID")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}
