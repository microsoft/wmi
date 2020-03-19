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

// MDM_VPNv2_DomainNameInformationList02_01 struct
type MDM_VPNv2_DomainNameInformationList02_01 struct {
	*cim.WmiInstance

	//
	AutoTrigger bool

	//
	DnsServers string

	//
	DomainName string

	//
	DomainNameType string

	//
	InstanceID string

	//
	ParentID string

	//
	Persistent bool

	//
	WebProxyServers string
}

func NewMDM_VPNv2_DomainNameInformationList02_01Ex1(instance *cim.WmiInstance) (newInstance *MDM_VPNv2_DomainNameInformationList02_01, err error) {
	tmp, err := instance, nil

	if err != nil {
		return
	}
	newInstance = &MDM_VPNv2_DomainNameInformationList02_01{
		WmiInstance: tmp,
	}
	return
}

func NewMDM_VPNv2_DomainNameInformationList02_01Ex6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *MDM_VPNv2_DomainNameInformationList02_01, err error) {
	tmp, err := instance.GetWmiInstance(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &MDM_VPNv2_DomainNameInformationList02_01{
		WmiInstance: tmp,
	}
	return
}

// SetAutoTrigger sets the value of AutoTrigger for the instance
func (instance *MDM_VPNv2_DomainNameInformationList02_01) SetPropertyAutoTrigger(value bool) (err error) {
	return instance.SetProperty("AutoTrigger", value)
}

// GetAutoTrigger gets the value of AutoTrigger for the instance
func (instance *MDM_VPNv2_DomainNameInformationList02_01) GetPropertyAutoTrigger() (value bool, err error) {
	retValue, err := instance.GetProperty("AutoTrigger")
	if err != nil {
		return
	}
	value, ok := retValue.(bool)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetDnsServers sets the value of DnsServers for the instance
func (instance *MDM_VPNv2_DomainNameInformationList02_01) SetPropertyDnsServers(value string) (err error) {
	return instance.SetProperty("DnsServers", value)
}

// GetDnsServers gets the value of DnsServers for the instance
func (instance *MDM_VPNv2_DomainNameInformationList02_01) GetPropertyDnsServers() (value string, err error) {
	retValue, err := instance.GetProperty("DnsServers")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetDomainName sets the value of DomainName for the instance
func (instance *MDM_VPNv2_DomainNameInformationList02_01) SetPropertyDomainName(value string) (err error) {
	return instance.SetProperty("DomainName", value)
}

// GetDomainName gets the value of DomainName for the instance
func (instance *MDM_VPNv2_DomainNameInformationList02_01) GetPropertyDomainName() (value string, err error) {
	retValue, err := instance.GetProperty("DomainName")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetDomainNameType sets the value of DomainNameType for the instance
func (instance *MDM_VPNv2_DomainNameInformationList02_01) SetPropertyDomainNameType(value string) (err error) {
	return instance.SetProperty("DomainNameType", value)
}

// GetDomainNameType gets the value of DomainNameType for the instance
func (instance *MDM_VPNv2_DomainNameInformationList02_01) GetPropertyDomainNameType() (value string, err error) {
	retValue, err := instance.GetProperty("DomainNameType")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetInstanceID sets the value of InstanceID for the instance
func (instance *MDM_VPNv2_DomainNameInformationList02_01) SetPropertyInstanceID(value string) (err error) {
	return instance.SetProperty("InstanceID", value)
}

// GetInstanceID gets the value of InstanceID for the instance
func (instance *MDM_VPNv2_DomainNameInformationList02_01) GetPropertyInstanceID() (value string, err error) {
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

// SetParentID sets the value of ParentID for the instance
func (instance *MDM_VPNv2_DomainNameInformationList02_01) SetPropertyParentID(value string) (err error) {
	return instance.SetProperty("ParentID", value)
}

// GetParentID gets the value of ParentID for the instance
func (instance *MDM_VPNv2_DomainNameInformationList02_01) GetPropertyParentID() (value string, err error) {
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

// SetPersistent sets the value of Persistent for the instance
func (instance *MDM_VPNv2_DomainNameInformationList02_01) SetPropertyPersistent(value bool) (err error) {
	return instance.SetProperty("Persistent", value)
}

// GetPersistent gets the value of Persistent for the instance
func (instance *MDM_VPNv2_DomainNameInformationList02_01) GetPropertyPersistent() (value bool, err error) {
	retValue, err := instance.GetProperty("Persistent")
	if err != nil {
		return
	}
	value, ok := retValue.(bool)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetWebProxyServers sets the value of WebProxyServers for the instance
func (instance *MDM_VPNv2_DomainNameInformationList02_01) SetPropertyWebProxyServers(value string) (err error) {
	return instance.SetProperty("WebProxyServers", value)
}

// GetWebProxyServers gets the value of WebProxyServers for the instance
func (instance *MDM_VPNv2_DomainNameInformationList02_01) GetPropertyWebProxyServers() (value string, err error) {
	retValue, err := instance.GetProperty("WebProxyServers")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}
