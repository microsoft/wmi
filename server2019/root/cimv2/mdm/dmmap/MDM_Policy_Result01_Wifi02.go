// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 3/16/2020 using wmigen
//      Source root.CIMV2.mdm.dmmap
//////////////////////////////////////////////
package dmmap

import (
	"github.com/microsoft/wmi/pkg/wmiinstance"
)

// MDM_Policy_Result01_Wifi02 struct
type MDM_Policy_Result01_Wifi02 struct {
	cim.WmiInstance

	//
	AllowAutoConnectToWiFiSenseHotspots int32

	//
	AllowInternetSharing int32

	//
	AllowManualWiFiConfiguration int32

	//
	AllowWiFi int32

	//
	AllowWiFiDirect int32

	//
	InstanceID string

	//
	ParentID string

	//
	WLANScanMode int32
}

// SetAllowAutoConnectToWiFiSenseHotspots sets the value of AllowAutoConnectToWiFiSenseHotspots for the instance
func (instance *MDM_Policy_Result01_Wifi02) SetPropertyAllowAutoConnectToWiFiSenseHotspots(value int32) (err error) {
	return instance.SetProperty("AllowAutoConnectToWiFiSenseHotspots", value)
}

// GetAllowAutoConnectToWiFiSenseHotspots gets the value of AllowAutoConnectToWiFiSenseHotspots for the instance
func (instance *MDM_Policy_Result01_Wifi02) GetPropertyAllowAutoConnectToWiFiSenseHotspots() (value int32, err error) {
	retValue, err := instance.GetProperty("AllowAutoConnectToWiFiSenseHotspots")
	if err != nil {
		return
	}
	value, ok := retValue.(int32)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetAllowInternetSharing sets the value of AllowInternetSharing for the instance
func (instance *MDM_Policy_Result01_Wifi02) SetPropertyAllowInternetSharing(value int32) (err error) {
	return instance.SetProperty("AllowInternetSharing", value)
}

// GetAllowInternetSharing gets the value of AllowInternetSharing for the instance
func (instance *MDM_Policy_Result01_Wifi02) GetPropertyAllowInternetSharing() (value int32, err error) {
	retValue, err := instance.GetProperty("AllowInternetSharing")
	if err != nil {
		return
	}
	value, ok := retValue.(int32)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetAllowManualWiFiConfiguration sets the value of AllowManualWiFiConfiguration for the instance
func (instance *MDM_Policy_Result01_Wifi02) SetPropertyAllowManualWiFiConfiguration(value int32) (err error) {
	return instance.SetProperty("AllowManualWiFiConfiguration", value)
}

// GetAllowManualWiFiConfiguration gets the value of AllowManualWiFiConfiguration for the instance
func (instance *MDM_Policy_Result01_Wifi02) GetPropertyAllowManualWiFiConfiguration() (value int32, err error) {
	retValue, err := instance.GetProperty("AllowManualWiFiConfiguration")
	if err != nil {
		return
	}
	value, ok := retValue.(int32)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetAllowWiFi sets the value of AllowWiFi for the instance
func (instance *MDM_Policy_Result01_Wifi02) SetPropertyAllowWiFi(value int32) (err error) {
	return instance.SetProperty("AllowWiFi", value)
}

// GetAllowWiFi gets the value of AllowWiFi for the instance
func (instance *MDM_Policy_Result01_Wifi02) GetPropertyAllowWiFi() (value int32, err error) {
	retValue, err := instance.GetProperty("AllowWiFi")
	if err != nil {
		return
	}
	value, ok := retValue.(int32)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetAllowWiFiDirect sets the value of AllowWiFiDirect for the instance
func (instance *MDM_Policy_Result01_Wifi02) SetPropertyAllowWiFiDirect(value int32) (err error) {
	return instance.SetProperty("AllowWiFiDirect", value)
}

// GetAllowWiFiDirect gets the value of AllowWiFiDirect for the instance
func (instance *MDM_Policy_Result01_Wifi02) GetPropertyAllowWiFiDirect() (value int32, err error) {
	retValue, err := instance.GetProperty("AllowWiFiDirect")
	if err != nil {
		return
	}
	value, ok := retValue.(int32)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetInstanceID sets the value of InstanceID for the instance
func (instance *MDM_Policy_Result01_Wifi02) SetPropertyInstanceID(value string) (err error) {
	return instance.SetProperty("InstanceID", value)
}

// GetInstanceID gets the value of InstanceID for the instance
func (instance *MDM_Policy_Result01_Wifi02) GetPropertyInstanceID() (value string, err error) {
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
func (instance *MDM_Policy_Result01_Wifi02) SetPropertyParentID(value string) (err error) {
	return instance.SetProperty("ParentID", value)
}

// GetParentID gets the value of ParentID for the instance
func (instance *MDM_Policy_Result01_Wifi02) GetPropertyParentID() (value string, err error) {
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

// SetWLANScanMode sets the value of WLANScanMode for the instance
func (instance *MDM_Policy_Result01_Wifi02) SetPropertyWLANScanMode(value int32) (err error) {
	return instance.SetProperty("WLANScanMode", value)
}

// GetWLANScanMode gets the value of WLANScanMode for the instance
func (instance *MDM_Policy_Result01_Wifi02) GetPropertyWLANScanMode() (value int32, err error) {
	retValue, err := instance.GetProperty("WLANScanMode")
	if err != nil {
		return
	}
	value, ok := retValue.(int32)
	if !ok {
		// TODO: Set an error
	}
	return
}
