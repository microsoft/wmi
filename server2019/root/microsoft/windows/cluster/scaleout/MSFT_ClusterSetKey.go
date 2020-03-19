// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 3/19/2020 using wmigen
//      Source root.Microsoft.Windows.Cluster.Scaleout
//////////////////////////////////////////////
package scaleout

import (
	"github.com/microsoft/wmi/pkg/base/instance"
	"github.com/microsoft/wmi/pkg/base/query"
	cim "github.com/microsoft/wmi/pkg/wmiinstance"
)

// MSFT_ClusterSetKey struct
type MSFT_ClusterSetKey struct {
	*cim.WmiInstance

	//
	Cert []uint8

	//
	key []uint8
}

func NewMSFT_ClusterSetKeyEx1(instance *cim.WmiInstance) (newInstance *MSFT_ClusterSetKey, err error) {
	tmp, err := instance, nil

	if err != nil {
		return
	}
	newInstance = &MSFT_ClusterSetKey{
		WmiInstance: tmp,
	}
	return
}

func NewMSFT_ClusterSetKeyEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *MSFT_ClusterSetKey, err error) {
	tmp, err := instance.GetWmiInstance(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &MSFT_ClusterSetKey{
		WmiInstance: tmp,
	}
	return
}

// SetCert sets the value of Cert for the instance
func (instance *MSFT_ClusterSetKey) SetPropertyCert(value []uint8) (err error) {
	return instance.SetProperty("Cert", value)
}

// GetCert gets the value of Cert for the instance
func (instance *MSFT_ClusterSetKey) GetPropertyCert() (value []uint8, err error) {
	retValue, err := instance.GetProperty("Cert")
	if err != nil {
		return
	}
	value, ok := retValue.([]uint8)
	if !ok {
		// TODO: Set an error
	}
	return
}

// Setkey sets the value of key for the instance
func (instance *MSFT_ClusterSetKey) SetPropertykey(value []uint8) (err error) {
	return instance.SetProperty("key", value)
}

// Getkey gets the value of key for the instance
func (instance *MSFT_ClusterSetKey) GetPropertykey() (value []uint8, err error) {
	retValue, err := instance.GetProperty("key")
	if err != nil {
		return
	}
	value, ok := retValue.([]uint8)
	if !ok {
		// TODO: Set an error
	}
	return
}
