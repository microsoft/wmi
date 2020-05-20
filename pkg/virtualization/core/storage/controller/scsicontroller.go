// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.

package controller

import (
	wmi "github.com/microsoft/wmi/pkg/wmiinstance"
	"github.com/microsoft/wmi/server2019/root/virtualization/v2"
)

type SCSIController struct {
	*v2.Msvm_SCSIProtocolController
}

// NewSCSIController
func NewSCSIController(instance *wmi.WmiInstance) (*SCSIController, error) {
	wmivm, err := v2.NewMsvm_SCSIProtocolControllerEx1(instance)
	if err != nil {
		return nil, err
	}
	return &SCSIController{wmivm}, nil
}
