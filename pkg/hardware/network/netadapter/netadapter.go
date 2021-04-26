// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.

package netadapter

import (
	wmi "github.com/microsoft/wmi/pkg/wmiinstance"
	"github.com/microsoft/wmi/server2019/root/standardcimv2"
)

type NetworkAdapter struct {
	*standardcimv2.MSFT_NetAdapter
}

// NewNetworkAdapter
func NewNetworkAdapter(instance *wmi.WmiInstance) (*NetworkAdapter, error) {
	wmiadapter, err := standardcimv2.NewMSFT_NetAdapterEx1(instance)
	if err != nil {
		return nil, err
	}
	return &NetworkAdapter{wmiadapter}, nil
}

func (na *NetworkAdapter) CloneEx1() (*NetworkAdapter, error) {
	tmp, err := na.Clone()
	if err != nil {
		return nil, err
	}
	return NewNetworkAdapter(tmp)
}

func (r *NetworkAdapter) GetInterfaceIndex() (value int32, err error) {
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
