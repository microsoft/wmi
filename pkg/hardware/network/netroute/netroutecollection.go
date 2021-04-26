// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.

package netroute

import (
	wmi "github.com/microsoft/wmi/pkg/wmiinstance"
)

type NetworkRouteCollection []*NetworkRoute

func NewNetworkRouteCollection(instances []*wmi.WmiInstance) (col NetworkRouteCollection, err error) {
	for _, inst := range instances {
		na, err1 := NewNetworkRoute(inst)
		if err1 != nil {
			err = err1
			return
		}
		col = append(col, na)
	}
	return
}
func (netroutes *NetworkRouteCollection) Close() (err error) {
	for _, value := range *netroutes {
		value.Close()
	}
	return nil
}

func (netroutes *NetworkRouteCollection) String() string {
	return ""
}
