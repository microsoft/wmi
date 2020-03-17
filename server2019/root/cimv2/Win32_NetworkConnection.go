// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 3/16/2020 using wmigen
//      Source root.CIMV2
//////////////////////////////////////////////
package cimv2

// Win32_NetworkConnection struct
type Win32_NetworkConnection struct {
	CIM_LogicalElement

	//
	AccessMask uint32

	//
	Comment string

	//
	ConnectionState string

	//
	ConnectionType string

	//
	DisplayType string

	//
	LocalName string

	//
	Persistent bool

	//
	ProviderName string

	//
	RemoteName string

	//
	RemotePath string

	//
	ResourceType string

	//
	UserName string
}

// SetAccessMask sets the value of AccessMask for the instance
func (instance *Win32_NetworkConnection) SetPropertyAccessMask(value uint32) (err error) {
	return instance.SetProperty("AccessMask", value)
}

// GetAccessMask gets the value of AccessMask for the instance
func (instance *Win32_NetworkConnection) GetPropertyAccessMask() (value uint32, err error) {
	retValue, err := instance.GetProperty("AccessMask")
	if err != nil {
		return
	}
	value, ok := retValue.(uint32)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetComment sets the value of Comment for the instance
func (instance *Win32_NetworkConnection) SetPropertyComment(value string) (err error) {
	return instance.SetProperty("Comment", value)
}

// GetComment gets the value of Comment for the instance
func (instance *Win32_NetworkConnection) GetPropertyComment() (value string, err error) {
	retValue, err := instance.GetProperty("Comment")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetConnectionState sets the value of ConnectionState for the instance
func (instance *Win32_NetworkConnection) SetPropertyConnectionState(value string) (err error) {
	return instance.SetProperty("ConnectionState", value)
}

// GetConnectionState gets the value of ConnectionState for the instance
func (instance *Win32_NetworkConnection) GetPropertyConnectionState() (value string, err error) {
	retValue, err := instance.GetProperty("ConnectionState")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetConnectionType sets the value of ConnectionType for the instance
func (instance *Win32_NetworkConnection) SetPropertyConnectionType(value string) (err error) {
	return instance.SetProperty("ConnectionType", value)
}

// GetConnectionType gets the value of ConnectionType for the instance
func (instance *Win32_NetworkConnection) GetPropertyConnectionType() (value string, err error) {
	retValue, err := instance.GetProperty("ConnectionType")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetDisplayType sets the value of DisplayType for the instance
func (instance *Win32_NetworkConnection) SetPropertyDisplayType(value string) (err error) {
	return instance.SetProperty("DisplayType", value)
}

// GetDisplayType gets the value of DisplayType for the instance
func (instance *Win32_NetworkConnection) GetPropertyDisplayType() (value string, err error) {
	retValue, err := instance.GetProperty("DisplayType")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetLocalName sets the value of LocalName for the instance
func (instance *Win32_NetworkConnection) SetPropertyLocalName(value string) (err error) {
	return instance.SetProperty("LocalName", value)
}

// GetLocalName gets the value of LocalName for the instance
func (instance *Win32_NetworkConnection) GetPropertyLocalName() (value string, err error) {
	retValue, err := instance.GetProperty("LocalName")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetPersistent sets the value of Persistent for the instance
func (instance *Win32_NetworkConnection) SetPropertyPersistent(value bool) (err error) {
	return instance.SetProperty("Persistent", value)
}

// GetPersistent gets the value of Persistent for the instance
func (instance *Win32_NetworkConnection) GetPropertyPersistent() (value bool, err error) {
	retValue, err := instance.GetProperty("Persistent")
	if err != nil {
		return
	}
	value, ok := retValue.(bool)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetProviderName sets the value of ProviderName for the instance
func (instance *Win32_NetworkConnection) SetPropertyProviderName(value string) (err error) {
	return instance.SetProperty("ProviderName", value)
}

// GetProviderName gets the value of ProviderName for the instance
func (instance *Win32_NetworkConnection) GetPropertyProviderName() (value string, err error) {
	retValue, err := instance.GetProperty("ProviderName")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetRemoteName sets the value of RemoteName for the instance
func (instance *Win32_NetworkConnection) SetPropertyRemoteName(value string) (err error) {
	return instance.SetProperty("RemoteName", value)
}

// GetRemoteName gets the value of RemoteName for the instance
func (instance *Win32_NetworkConnection) GetPropertyRemoteName() (value string, err error) {
	retValue, err := instance.GetProperty("RemoteName")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetRemotePath sets the value of RemotePath for the instance
func (instance *Win32_NetworkConnection) SetPropertyRemotePath(value string) (err error) {
	return instance.SetProperty("RemotePath", value)
}

// GetRemotePath gets the value of RemotePath for the instance
func (instance *Win32_NetworkConnection) GetPropertyRemotePath() (value string, err error) {
	retValue, err := instance.GetProperty("RemotePath")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetResourceType sets the value of ResourceType for the instance
func (instance *Win32_NetworkConnection) SetPropertyResourceType(value string) (err error) {
	return instance.SetProperty("ResourceType", value)
}

// GetResourceType gets the value of ResourceType for the instance
func (instance *Win32_NetworkConnection) GetPropertyResourceType() (value string, err error) {
	retValue, err := instance.GetProperty("ResourceType")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetUserName sets the value of UserName for the instance
func (instance *Win32_NetworkConnection) SetPropertyUserName(value string) (err error) {
	return instance.SetProperty("UserName", value)
}

// GetUserName gets the value of UserName for the instance
func (instance *Win32_NetworkConnection) GetPropertyUserName() (value string, err error) {
	retValue, err := instance.GetProperty("UserName")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}
