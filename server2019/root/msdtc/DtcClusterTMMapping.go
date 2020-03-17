// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 3/16/2020 using wmigen
//      Source root.msdtc
//////////////////////////////////////////////
package msdtc

import (
	"github.com/microsoft/wmi/pkg/wmiinstance"
)

// DtcClusterTMMapping struct
type DtcClusterTMMapping struct {
	cim.WmiInstance

	//
	Application string

	//
	ApplicationType string

	//
	ClusterResourceName string

	//
	Local bool

	//
	Name string
}

// SetApplication sets the value of Application for the instance
func (instance *DtcClusterTMMapping) SetPropertyApplication(value string) (err error) {
	return instance.SetProperty("Application", value)
}

// GetApplication gets the value of Application for the instance
func (instance *DtcClusterTMMapping) GetPropertyApplication() (value string, err error) {
	retValue, err := instance.GetProperty("Application")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetApplicationType sets the value of ApplicationType for the instance
func (instance *DtcClusterTMMapping) SetPropertyApplicationType(value string) (err error) {
	return instance.SetProperty("ApplicationType", value)
}

// GetApplicationType gets the value of ApplicationType for the instance
func (instance *DtcClusterTMMapping) GetPropertyApplicationType() (value string, err error) {
	retValue, err := instance.GetProperty("ApplicationType")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetClusterResourceName sets the value of ClusterResourceName for the instance
func (instance *DtcClusterTMMapping) SetPropertyClusterResourceName(value string) (err error) {
	return instance.SetProperty("ClusterResourceName", value)
}

// GetClusterResourceName gets the value of ClusterResourceName for the instance
func (instance *DtcClusterTMMapping) GetPropertyClusterResourceName() (value string, err error) {
	retValue, err := instance.GetProperty("ClusterResourceName")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetLocal sets the value of Local for the instance
func (instance *DtcClusterTMMapping) SetPropertyLocal(value bool) (err error) {
	return instance.SetProperty("Local", value)
}

// GetLocal gets the value of Local for the instance
func (instance *DtcClusterTMMapping) GetPropertyLocal() (value bool, err error) {
	retValue, err := instance.GetProperty("Local")
	if err != nil {
		return
	}
	value, ok := retValue.(bool)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetName sets the value of Name for the instance
func (instance *DtcClusterTMMapping) SetPropertyName(value string) (err error) {
	return instance.SetProperty("Name", value)
}

// GetName gets the value of Name for the instance
func (instance *DtcClusterTMMapping) GetPropertyName() (value string, err error) {
	retValue, err := instance.GetProperty("Name")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}
