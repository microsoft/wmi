// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 3/16/2020 using wmigen
//      Source root.CIMV2
//////////////////////////////////////////////
package cimv2

// CIM_NFS struct
type CIM_NFS struct {
	CIM_RemoteFileSystem

	//
	AttributeCaching bool

	//
	AttributeCachingForDirectoriesMax uint16

	//
	AttributeCachingForDirectoriesMin uint16

	//
	AttributeCachingForRegularFilesMax uint16

	//
	AttributeCachingForRegularFilesMin uint16

	//
	ForegroundMount bool

	//
	HardMount bool

	//
	Interrupt bool

	//
	MountFailureRetries uint16

	//
	ReadBufferSize uint64

	//
	RetransmissionAttempts uint16

	//
	RetransmissionTimeout uint32

	//
	ServerCommunicationPort uint32

	//
	WriteBufferSize uint64
}

// SetAttributeCaching sets the value of AttributeCaching for the instance
func (instance *CIM_NFS) SetPropertyAttributeCaching(value bool) (err error) {
	return instance.SetProperty("AttributeCaching", value)
}

// GetAttributeCaching gets the value of AttributeCaching for the instance
func (instance *CIM_NFS) GetPropertyAttributeCaching() (value bool, err error) {
	retValue, err := instance.GetProperty("AttributeCaching")
	if err != nil {
		return
	}
	value, ok := retValue.(bool)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetAttributeCachingForDirectoriesMax sets the value of AttributeCachingForDirectoriesMax for the instance
func (instance *CIM_NFS) SetPropertyAttributeCachingForDirectoriesMax(value uint16) (err error) {
	return instance.SetProperty("AttributeCachingForDirectoriesMax", value)
}

// GetAttributeCachingForDirectoriesMax gets the value of AttributeCachingForDirectoriesMax for the instance
func (instance *CIM_NFS) GetPropertyAttributeCachingForDirectoriesMax() (value uint16, err error) {
	retValue, err := instance.GetProperty("AttributeCachingForDirectoriesMax")
	if err != nil {
		return
	}
	value, ok := retValue.(uint16)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetAttributeCachingForDirectoriesMin sets the value of AttributeCachingForDirectoriesMin for the instance
func (instance *CIM_NFS) SetPropertyAttributeCachingForDirectoriesMin(value uint16) (err error) {
	return instance.SetProperty("AttributeCachingForDirectoriesMin", value)
}

// GetAttributeCachingForDirectoriesMin gets the value of AttributeCachingForDirectoriesMin for the instance
func (instance *CIM_NFS) GetPropertyAttributeCachingForDirectoriesMin() (value uint16, err error) {
	retValue, err := instance.GetProperty("AttributeCachingForDirectoriesMin")
	if err != nil {
		return
	}
	value, ok := retValue.(uint16)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetAttributeCachingForRegularFilesMax sets the value of AttributeCachingForRegularFilesMax for the instance
func (instance *CIM_NFS) SetPropertyAttributeCachingForRegularFilesMax(value uint16) (err error) {
	return instance.SetProperty("AttributeCachingForRegularFilesMax", value)
}

// GetAttributeCachingForRegularFilesMax gets the value of AttributeCachingForRegularFilesMax for the instance
func (instance *CIM_NFS) GetPropertyAttributeCachingForRegularFilesMax() (value uint16, err error) {
	retValue, err := instance.GetProperty("AttributeCachingForRegularFilesMax")
	if err != nil {
		return
	}
	value, ok := retValue.(uint16)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetAttributeCachingForRegularFilesMin sets the value of AttributeCachingForRegularFilesMin for the instance
func (instance *CIM_NFS) SetPropertyAttributeCachingForRegularFilesMin(value uint16) (err error) {
	return instance.SetProperty("AttributeCachingForRegularFilesMin", value)
}

// GetAttributeCachingForRegularFilesMin gets the value of AttributeCachingForRegularFilesMin for the instance
func (instance *CIM_NFS) GetPropertyAttributeCachingForRegularFilesMin() (value uint16, err error) {
	retValue, err := instance.GetProperty("AttributeCachingForRegularFilesMin")
	if err != nil {
		return
	}
	value, ok := retValue.(uint16)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetForegroundMount sets the value of ForegroundMount for the instance
func (instance *CIM_NFS) SetPropertyForegroundMount(value bool) (err error) {
	return instance.SetProperty("ForegroundMount", value)
}

// GetForegroundMount gets the value of ForegroundMount for the instance
func (instance *CIM_NFS) GetPropertyForegroundMount() (value bool, err error) {
	retValue, err := instance.GetProperty("ForegroundMount")
	if err != nil {
		return
	}
	value, ok := retValue.(bool)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetHardMount sets the value of HardMount for the instance
func (instance *CIM_NFS) SetPropertyHardMount(value bool) (err error) {
	return instance.SetProperty("HardMount", value)
}

// GetHardMount gets the value of HardMount for the instance
func (instance *CIM_NFS) GetPropertyHardMount() (value bool, err error) {
	retValue, err := instance.GetProperty("HardMount")
	if err != nil {
		return
	}
	value, ok := retValue.(bool)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetInterrupt sets the value of Interrupt for the instance
func (instance *CIM_NFS) SetPropertyInterrupt(value bool) (err error) {
	return instance.SetProperty("Interrupt", value)
}

// GetInterrupt gets the value of Interrupt for the instance
func (instance *CIM_NFS) GetPropertyInterrupt() (value bool, err error) {
	retValue, err := instance.GetProperty("Interrupt")
	if err != nil {
		return
	}
	value, ok := retValue.(bool)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetMountFailureRetries sets the value of MountFailureRetries for the instance
func (instance *CIM_NFS) SetPropertyMountFailureRetries(value uint16) (err error) {
	return instance.SetProperty("MountFailureRetries", value)
}

// GetMountFailureRetries gets the value of MountFailureRetries for the instance
func (instance *CIM_NFS) GetPropertyMountFailureRetries() (value uint16, err error) {
	retValue, err := instance.GetProperty("MountFailureRetries")
	if err != nil {
		return
	}
	value, ok := retValue.(uint16)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetReadBufferSize sets the value of ReadBufferSize for the instance
func (instance *CIM_NFS) SetPropertyReadBufferSize(value uint64) (err error) {
	return instance.SetProperty("ReadBufferSize", value)
}

// GetReadBufferSize gets the value of ReadBufferSize for the instance
func (instance *CIM_NFS) GetPropertyReadBufferSize() (value uint64, err error) {
	retValue, err := instance.GetProperty("ReadBufferSize")
	if err != nil {
		return
	}
	value, ok := retValue.(uint64)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetRetransmissionAttempts sets the value of RetransmissionAttempts for the instance
func (instance *CIM_NFS) SetPropertyRetransmissionAttempts(value uint16) (err error) {
	return instance.SetProperty("RetransmissionAttempts", value)
}

// GetRetransmissionAttempts gets the value of RetransmissionAttempts for the instance
func (instance *CIM_NFS) GetPropertyRetransmissionAttempts() (value uint16, err error) {
	retValue, err := instance.GetProperty("RetransmissionAttempts")
	if err != nil {
		return
	}
	value, ok := retValue.(uint16)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetRetransmissionTimeout sets the value of RetransmissionTimeout for the instance
func (instance *CIM_NFS) SetPropertyRetransmissionTimeout(value uint32) (err error) {
	return instance.SetProperty("RetransmissionTimeout", value)
}

// GetRetransmissionTimeout gets the value of RetransmissionTimeout for the instance
func (instance *CIM_NFS) GetPropertyRetransmissionTimeout() (value uint32, err error) {
	retValue, err := instance.GetProperty("RetransmissionTimeout")
	if err != nil {
		return
	}
	value, ok := retValue.(uint32)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetServerCommunicationPort sets the value of ServerCommunicationPort for the instance
func (instance *CIM_NFS) SetPropertyServerCommunicationPort(value uint32) (err error) {
	return instance.SetProperty("ServerCommunicationPort", value)
}

// GetServerCommunicationPort gets the value of ServerCommunicationPort for the instance
func (instance *CIM_NFS) GetPropertyServerCommunicationPort() (value uint32, err error) {
	retValue, err := instance.GetProperty("ServerCommunicationPort")
	if err != nil {
		return
	}
	value, ok := retValue.(uint32)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetWriteBufferSize sets the value of WriteBufferSize for the instance
func (instance *CIM_NFS) SetPropertyWriteBufferSize(value uint64) (err error) {
	return instance.SetProperty("WriteBufferSize", value)
}

// GetWriteBufferSize gets the value of WriteBufferSize for the instance
func (instance *CIM_NFS) GetPropertyWriteBufferSize() (value uint64, err error) {
	retValue, err := instance.GetProperty("WriteBufferSize")
	if err != nil {
		return
	}
	value, ok := retValue.(uint64)
	if !ok {
		// TODO: Set an error
	}
	return
}
