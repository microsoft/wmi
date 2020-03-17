// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 3/16/2020 using wmigen
//      Source root.StandardCimv2
//////////////////////////////////////////////
package standardcimv2

// MSFT_NetMainModeSA struct
type MSFT_NetMainModeSA struct {
	CIM_IKESAEndpoint

	//
	ExtendedFilterId uint64

	//
	IkePolicyKey string

	//
	KeyModule uint16

	//
	LocalEndpoint string

	//
	LocalFirstId MSFT_NetIPsecIdentity

	//
	LocalSecondId MSFT_NetIPsecIdentity

	//
	LocalUdpEncapsulationPort uint16

	//
	MaxQMSAs uint32

	//
	OtherGroupId string

	//
	RemoteEndpoint string

	//
	RemoteFirstId MSFT_NetIPsecIdentity

	//
	RemoteSecondId MSFT_NetIPsecIdentity

	//
	RemoteUdpEncapsulationPort uint16

	//
	VirtualIfTunnelId uint64
}

// SetExtendedFilterId sets the value of ExtendedFilterId for the instance
func (instance *MSFT_NetMainModeSA) SetPropertyExtendedFilterId(value uint64) (err error) {
	return instance.SetProperty("ExtendedFilterId", value)
}

// GetExtendedFilterId gets the value of ExtendedFilterId for the instance
func (instance *MSFT_NetMainModeSA) GetPropertyExtendedFilterId() (value uint64, err error) {
	retValue, err := instance.GetProperty("ExtendedFilterId")
	if err != nil {
		return
	}
	value, ok := retValue.(uint64)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetIkePolicyKey sets the value of IkePolicyKey for the instance
func (instance *MSFT_NetMainModeSA) SetPropertyIkePolicyKey(value string) (err error) {
	return instance.SetProperty("IkePolicyKey", value)
}

// GetIkePolicyKey gets the value of IkePolicyKey for the instance
func (instance *MSFT_NetMainModeSA) GetPropertyIkePolicyKey() (value string, err error) {
	retValue, err := instance.GetProperty("IkePolicyKey")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetKeyModule sets the value of KeyModule for the instance
func (instance *MSFT_NetMainModeSA) SetPropertyKeyModule(value uint16) (err error) {
	return instance.SetProperty("KeyModule", value)
}

// GetKeyModule gets the value of KeyModule for the instance
func (instance *MSFT_NetMainModeSA) GetPropertyKeyModule() (value uint16, err error) {
	retValue, err := instance.GetProperty("KeyModule")
	if err != nil {
		return
	}
	value, ok := retValue.(uint16)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetLocalEndpoint sets the value of LocalEndpoint for the instance
func (instance *MSFT_NetMainModeSA) SetPropertyLocalEndpoint(value string) (err error) {
	return instance.SetProperty("LocalEndpoint", value)
}

// GetLocalEndpoint gets the value of LocalEndpoint for the instance
func (instance *MSFT_NetMainModeSA) GetPropertyLocalEndpoint() (value string, err error) {
	retValue, err := instance.GetProperty("LocalEndpoint")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetLocalFirstId sets the value of LocalFirstId for the instance
func (instance *MSFT_NetMainModeSA) SetPropertyLocalFirstId(value MSFT_NetIPsecIdentity) (err error) {
	return instance.SetProperty("LocalFirstId", value)
}

// GetLocalFirstId gets the value of LocalFirstId for the instance
func (instance *MSFT_NetMainModeSA) GetPropertyLocalFirstId() (value MSFT_NetIPsecIdentity, err error) {
	retValue, err := instance.GetProperty("LocalFirstId")
	if err != nil {
		return
	}
	value, ok := retValue.(MSFT_NetIPsecIdentity)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetLocalSecondId sets the value of LocalSecondId for the instance
func (instance *MSFT_NetMainModeSA) SetPropertyLocalSecondId(value MSFT_NetIPsecIdentity) (err error) {
	return instance.SetProperty("LocalSecondId", value)
}

// GetLocalSecondId gets the value of LocalSecondId for the instance
func (instance *MSFT_NetMainModeSA) GetPropertyLocalSecondId() (value MSFT_NetIPsecIdentity, err error) {
	retValue, err := instance.GetProperty("LocalSecondId")
	if err != nil {
		return
	}
	value, ok := retValue.(MSFT_NetIPsecIdentity)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetLocalUdpEncapsulationPort sets the value of LocalUdpEncapsulationPort for the instance
func (instance *MSFT_NetMainModeSA) SetPropertyLocalUdpEncapsulationPort(value uint16) (err error) {
	return instance.SetProperty("LocalUdpEncapsulationPort", value)
}

// GetLocalUdpEncapsulationPort gets the value of LocalUdpEncapsulationPort for the instance
func (instance *MSFT_NetMainModeSA) GetPropertyLocalUdpEncapsulationPort() (value uint16, err error) {
	retValue, err := instance.GetProperty("LocalUdpEncapsulationPort")
	if err != nil {
		return
	}
	value, ok := retValue.(uint16)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetMaxQMSAs sets the value of MaxQMSAs for the instance
func (instance *MSFT_NetMainModeSA) SetPropertyMaxQMSAs(value uint32) (err error) {
	return instance.SetProperty("MaxQMSAs", value)
}

// GetMaxQMSAs gets the value of MaxQMSAs for the instance
func (instance *MSFT_NetMainModeSA) GetPropertyMaxQMSAs() (value uint32, err error) {
	retValue, err := instance.GetProperty("MaxQMSAs")
	if err != nil {
		return
	}
	value, ok := retValue.(uint32)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetOtherGroupId sets the value of OtherGroupId for the instance
func (instance *MSFT_NetMainModeSA) SetPropertyOtherGroupId(value string) (err error) {
	return instance.SetProperty("OtherGroupId", value)
}

// GetOtherGroupId gets the value of OtherGroupId for the instance
func (instance *MSFT_NetMainModeSA) GetPropertyOtherGroupId() (value string, err error) {
	retValue, err := instance.GetProperty("OtherGroupId")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetRemoteEndpoint sets the value of RemoteEndpoint for the instance
func (instance *MSFT_NetMainModeSA) SetPropertyRemoteEndpoint(value string) (err error) {
	return instance.SetProperty("RemoteEndpoint", value)
}

// GetRemoteEndpoint gets the value of RemoteEndpoint for the instance
func (instance *MSFT_NetMainModeSA) GetPropertyRemoteEndpoint() (value string, err error) {
	retValue, err := instance.GetProperty("RemoteEndpoint")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetRemoteFirstId sets the value of RemoteFirstId for the instance
func (instance *MSFT_NetMainModeSA) SetPropertyRemoteFirstId(value MSFT_NetIPsecIdentity) (err error) {
	return instance.SetProperty("RemoteFirstId", value)
}

// GetRemoteFirstId gets the value of RemoteFirstId for the instance
func (instance *MSFT_NetMainModeSA) GetPropertyRemoteFirstId() (value MSFT_NetIPsecIdentity, err error) {
	retValue, err := instance.GetProperty("RemoteFirstId")
	if err != nil {
		return
	}
	value, ok := retValue.(MSFT_NetIPsecIdentity)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetRemoteSecondId sets the value of RemoteSecondId for the instance
func (instance *MSFT_NetMainModeSA) SetPropertyRemoteSecondId(value MSFT_NetIPsecIdentity) (err error) {
	return instance.SetProperty("RemoteSecondId", value)
}

// GetRemoteSecondId gets the value of RemoteSecondId for the instance
func (instance *MSFT_NetMainModeSA) GetPropertyRemoteSecondId() (value MSFT_NetIPsecIdentity, err error) {
	retValue, err := instance.GetProperty("RemoteSecondId")
	if err != nil {
		return
	}
	value, ok := retValue.(MSFT_NetIPsecIdentity)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetRemoteUdpEncapsulationPort sets the value of RemoteUdpEncapsulationPort for the instance
func (instance *MSFT_NetMainModeSA) SetPropertyRemoteUdpEncapsulationPort(value uint16) (err error) {
	return instance.SetProperty("RemoteUdpEncapsulationPort", value)
}

// GetRemoteUdpEncapsulationPort gets the value of RemoteUdpEncapsulationPort for the instance
func (instance *MSFT_NetMainModeSA) GetPropertyRemoteUdpEncapsulationPort() (value uint16, err error) {
	retValue, err := instance.GetProperty("RemoteUdpEncapsulationPort")
	if err != nil {
		return
	}
	value, ok := retValue.(uint16)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetVirtualIfTunnelId sets the value of VirtualIfTunnelId for the instance
func (instance *MSFT_NetMainModeSA) SetPropertyVirtualIfTunnelId(value uint64) (err error) {
	return instance.SetProperty("VirtualIfTunnelId", value)
}

// GetVirtualIfTunnelId gets the value of VirtualIfTunnelId for the instance
func (instance *MSFT_NetMainModeSA) GetPropertyVirtualIfTunnelId() (value uint64, err error) {
	retValue, err := instance.GetProperty("VirtualIfTunnelId")
	if err != nil {
		return
	}
	value, ok := retValue.(uint64)
	if !ok {
		// TODO: Set an error
	}
	return
}
