// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 3/16/2020 using wmigen
//      Source root.Microsoft.Windows.Storage
//////////////////////////////////////////////
package storage

// MSFT_ReplicaPeer struct
type MSFT_ReplicaPeer struct {
	MSFT_StorageObject

	//
	IsPrimary bool

	//
	PeerObject MSFT_StorageObject

	//
	PeerObjectId string

	//
	PeerObjectName string

	//
	PeerObjectType uint16

	//
	PeerProviderURI string

	//
	PeerSubsystemName string

	//
	PeerUniqueId string
}

// SetIsPrimary sets the value of IsPrimary for the instance
func (instance *MSFT_ReplicaPeer) SetPropertyIsPrimary(value bool) (err error) {
	return instance.SetProperty("IsPrimary", value)
}

// GetIsPrimary gets the value of IsPrimary for the instance
func (instance *MSFT_ReplicaPeer) GetPropertyIsPrimary() (value bool, err error) {
	retValue, err := instance.GetProperty("IsPrimary")
	if err != nil {
		return
	}
	value, ok := retValue.(bool)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetPeerObject sets the value of PeerObject for the instance
func (instance *MSFT_ReplicaPeer) SetPropertyPeerObject(value MSFT_StorageObject) (err error) {
	return instance.SetProperty("PeerObject", value)
}

// GetPeerObject gets the value of PeerObject for the instance
func (instance *MSFT_ReplicaPeer) GetPropertyPeerObject() (value MSFT_StorageObject, err error) {
	retValue, err := instance.GetProperty("PeerObject")
	if err != nil {
		return
	}
	value, ok := retValue.(MSFT_StorageObject)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetPeerObjectId sets the value of PeerObjectId for the instance
func (instance *MSFT_ReplicaPeer) SetPropertyPeerObjectId(value string) (err error) {
	return instance.SetProperty("PeerObjectId", value)
}

// GetPeerObjectId gets the value of PeerObjectId for the instance
func (instance *MSFT_ReplicaPeer) GetPropertyPeerObjectId() (value string, err error) {
	retValue, err := instance.GetProperty("PeerObjectId")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetPeerObjectName sets the value of PeerObjectName for the instance
func (instance *MSFT_ReplicaPeer) SetPropertyPeerObjectName(value string) (err error) {
	return instance.SetProperty("PeerObjectName", value)
}

// GetPeerObjectName gets the value of PeerObjectName for the instance
func (instance *MSFT_ReplicaPeer) GetPropertyPeerObjectName() (value string, err error) {
	retValue, err := instance.GetProperty("PeerObjectName")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetPeerObjectType sets the value of PeerObjectType for the instance
func (instance *MSFT_ReplicaPeer) SetPropertyPeerObjectType(value uint16) (err error) {
	return instance.SetProperty("PeerObjectType", value)
}

// GetPeerObjectType gets the value of PeerObjectType for the instance
func (instance *MSFT_ReplicaPeer) GetPropertyPeerObjectType() (value uint16, err error) {
	retValue, err := instance.GetProperty("PeerObjectType")
	if err != nil {
		return
	}
	value, ok := retValue.(uint16)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetPeerProviderURI sets the value of PeerProviderURI for the instance
func (instance *MSFT_ReplicaPeer) SetPropertyPeerProviderURI(value string) (err error) {
	return instance.SetProperty("PeerProviderURI", value)
}

// GetPeerProviderURI gets the value of PeerProviderURI for the instance
func (instance *MSFT_ReplicaPeer) GetPropertyPeerProviderURI() (value string, err error) {
	retValue, err := instance.GetProperty("PeerProviderURI")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetPeerSubsystemName sets the value of PeerSubsystemName for the instance
func (instance *MSFT_ReplicaPeer) SetPropertyPeerSubsystemName(value string) (err error) {
	return instance.SetProperty("PeerSubsystemName", value)
}

// GetPeerSubsystemName gets the value of PeerSubsystemName for the instance
func (instance *MSFT_ReplicaPeer) GetPropertyPeerSubsystemName() (value string, err error) {
	retValue, err := instance.GetProperty("PeerSubsystemName")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetPeerUniqueId sets the value of PeerUniqueId for the instance
func (instance *MSFT_ReplicaPeer) SetPropertyPeerUniqueId(value string) (err error) {
	return instance.SetProperty("PeerUniqueId", value)
}

// GetPeerUniqueId gets the value of PeerUniqueId for the instance
func (instance *MSFT_ReplicaPeer) GetPropertyPeerUniqueId() (value string, err error) {
	retValue, err := instance.GetProperty("PeerUniqueId")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}
