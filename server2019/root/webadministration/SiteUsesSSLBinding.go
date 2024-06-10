// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

// Author:
//
//	Auto Generated on 6/6/2024 using wmigen
//	Source root.WebAdministration
//
// ////////////////////////////////////////////
package webadministration

import (
	"github.com/microsoft/wmi/pkg/base/query"
	cim "github.com/microsoft/wmi/pkg/wmiinstance"
)

// SiteUsesSSLBinding struct
type SiteUsesSSLBinding struct {
	*ObjectContainerAssociation
}

func NewSiteUsesSSLBindingEx1(instance *cim.WmiInstance) (newInstance *SiteUsesSSLBinding, err error) {
	tmp, err := NewObjectContainerAssociationEx1(instance)

	if err != nil {
		return
	}
	newInstance = &SiteUsesSSLBinding{
		ObjectContainerAssociation: tmp,
	}
	return
}

func NewSiteUsesSSLBindingEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *SiteUsesSSLBinding, err error) {
	tmp, err := NewObjectContainerAssociationEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &SiteUsesSSLBinding{
		ObjectContainerAssociation: tmp,
	}
	return
}
