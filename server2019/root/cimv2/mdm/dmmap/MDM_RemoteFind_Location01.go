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

// MDM_RemoteFind_Location01 struct
type MDM_RemoteFind_Location01 struct {
	cim.WmiInstance

	//
	Accuracy int32

	//
	Age string

	//
	Altitude float32

	//
	AltitudeAccuracy int32

	//
	InstanceID string

	//
	Latitude float32

	//
	Longitude float32

	//
	ParentID string
}

// SetAccuracy sets the value of Accuracy for the instance
func (instance *MDM_RemoteFind_Location01) SetPropertyAccuracy(value int32) (err error) {
	return instance.SetProperty("Accuracy", value)
}

// GetAccuracy gets the value of Accuracy for the instance
func (instance *MDM_RemoteFind_Location01) GetPropertyAccuracy() (value int32, err error) {
	retValue, err := instance.GetProperty("Accuracy")
	if err != nil {
		return
	}
	value, ok := retValue.(int32)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetAge sets the value of Age for the instance
func (instance *MDM_RemoteFind_Location01) SetPropertyAge(value string) (err error) {
	return instance.SetProperty("Age", value)
}

// GetAge gets the value of Age for the instance
func (instance *MDM_RemoteFind_Location01) GetPropertyAge() (value string, err error) {
	retValue, err := instance.GetProperty("Age")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetAltitude sets the value of Altitude for the instance
func (instance *MDM_RemoteFind_Location01) SetPropertyAltitude(value float32) (err error) {
	return instance.SetProperty("Altitude", value)
}

// GetAltitude gets the value of Altitude for the instance
func (instance *MDM_RemoteFind_Location01) GetPropertyAltitude() (value float32, err error) {
	retValue, err := instance.GetProperty("Altitude")
	if err != nil {
		return
	}
	value, ok := retValue.(float32)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetAltitudeAccuracy sets the value of AltitudeAccuracy for the instance
func (instance *MDM_RemoteFind_Location01) SetPropertyAltitudeAccuracy(value int32) (err error) {
	return instance.SetProperty("AltitudeAccuracy", value)
}

// GetAltitudeAccuracy gets the value of AltitudeAccuracy for the instance
func (instance *MDM_RemoteFind_Location01) GetPropertyAltitudeAccuracy() (value int32, err error) {
	retValue, err := instance.GetProperty("AltitudeAccuracy")
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
func (instance *MDM_RemoteFind_Location01) SetPropertyInstanceID(value string) (err error) {
	return instance.SetProperty("InstanceID", value)
}

// GetInstanceID gets the value of InstanceID for the instance
func (instance *MDM_RemoteFind_Location01) GetPropertyInstanceID() (value string, err error) {
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

// SetLatitude sets the value of Latitude for the instance
func (instance *MDM_RemoteFind_Location01) SetPropertyLatitude(value float32) (err error) {
	return instance.SetProperty("Latitude", value)
}

// GetLatitude gets the value of Latitude for the instance
func (instance *MDM_RemoteFind_Location01) GetPropertyLatitude() (value float32, err error) {
	retValue, err := instance.GetProperty("Latitude")
	if err != nil {
		return
	}
	value, ok := retValue.(float32)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetLongitude sets the value of Longitude for the instance
func (instance *MDM_RemoteFind_Location01) SetPropertyLongitude(value float32) (err error) {
	return instance.SetProperty("Longitude", value)
}

// GetLongitude gets the value of Longitude for the instance
func (instance *MDM_RemoteFind_Location01) GetPropertyLongitude() (value float32, err error) {
	retValue, err := instance.GetProperty("Longitude")
	if err != nil {
		return
	}
	value, ok := retValue.(float32)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetParentID sets the value of ParentID for the instance
func (instance *MDM_RemoteFind_Location01) SetPropertyParentID(value string) (err error) {
	return instance.SetProperty("ParentID", value)
}

// GetParentID gets the value of ParentID for the instance
func (instance *MDM_RemoteFind_Location01) GetPropertyParentID() (value string, err error) {
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
