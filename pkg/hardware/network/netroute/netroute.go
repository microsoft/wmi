// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.

package netroute

import (
	wmi "github.com/microsoft/wmi/pkg/wmiinstance"
	"github.com/microsoft/wmi/server2019/root/standardcimv2"
)

type NetworkRoute struct {
	*standardcimv2.MSFT_NetRoute
}

// NewNetworkRoute
func NewNetworkRoute(instance *wmi.WmiInstance) (*NetworkRoute, error) {
	wminetroute, err := standardcimv2.NewMSFT_NetRouteEx1(instance)
	if err != nil {
		return nil, err
	}
	return &NetworkRoute{wminetroute}, nil
}

func (r *NetworkRoute) GetInterfaceIndex() (value int32, err error) {
	retValue, err := r.GetProperty("InterfaceIndex")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	value, _ = retValue.(int32)
	return
}
