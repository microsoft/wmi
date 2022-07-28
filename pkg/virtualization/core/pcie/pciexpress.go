// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.

package pcie

import (
	wmi "github.com/microsoft/wmi/pkg/wmiinstance"
	v2 "github.com/microsoft/wmi/server2019/root/virtualization/v2"
)

type PciExpress struct {
	*v2.Msvm_PciExpress
}

func NewPciExpress(instance *wmi.WmiInstance) (*PciExpress, error) {
	wmivm, err := v2.NewMsvm_PciExpressEx1(instance)
	if err != nil {
		return nil, err
	}
	return &PciExpress{wmivm}, nil
}
