// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.

package ethernetport

import (
	"github.com/microsoft/wmi/pkg/base/host"
	"github.com/microsoft/wmi/pkg/base/query"
	"github.com/microsoft/wmi/pkg/constant"
	wmi "github.com/microsoft/wmi/pkg/wmiinstance"
	v2 "github.com/microsoft/wmi/server2019/root/virtualization/v2"
)

type ExternalEthernetPort struct {
	*v2.Msvm_ExternalEthernetPort
}

// NewExternalEthernetPort
func NewExternalEthernetPort(instance *wmi.WmiInstance) (*ExternalEthernetPort, error) {
	wmivm, err := v2.NewMsvm_ExternalEthernetPortEx1(instance)
	if err != nil {
		return nil, err
	}
	return &ExternalEthernetPort{wmivm}, nil
}

// GetExternalEthernetPort gets an existing virtual machine
// Make sure to call Close once done using this instance
func GetExternalEthernetPort(whost *host.WmiHost, ethernetName string) (ethernet *ExternalEthernetPort, err error) {
	creds := whost.GetCredential()
	query := query.NewWmiQuery("Msvm_ExternalEthernetPort", "ElementName", ethernetName)
	wmivm, err := v2.NewMsvm_ExternalEthernetPortEx6(whost.HostName, string(constant.Virtualization), creds.UserName, creds.Password, creds.Domain, query)
	if err != nil {
		return
	}
	ethernet = &ExternalEthernetPort{wmivm}
	return
}

// GetExternalEthernetPort gets an existing virtual machine
// Make sure to call Close once done using this instance
func GetExternalEthernetPortByID(whost *host.WmiHost, ethernetID string) (ethernet *ExternalEthernetPort, err error) {
	creds := whost.GetCredential()
	query := query.NewWmiQuery("Msvm_ExternalEthernetPort", "Name", ethernetID)
	wmivm, err := v2.NewMsvm_ExternalEthernetPortEx6(whost.HostName, string(constant.Virtualization), creds.UserName, creds.Password, creds.Domain, query)
	if err != nil {
		return
	}
	ethernet = &ExternalEthernetPort{wmivm}
	return
}
