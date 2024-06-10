// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 9/18/2020 using wmigen
//      Source root.StandardCimv2
//////////////////////////////////////////////
package standardcimv2

import (
	"github.com/microsoft/wmi/pkg/base/query"
	"github.com/microsoft/wmi/pkg/errors"
	cim "github.com/microsoft/wmi/pkg/wmiinstance"
	"reflect"
)

// MSFT_NetQuickModeSA struct
type MSFT_NetQuickModeSA struct {
	*CIM_IPsecSAEndpoint

	//
	EmTargetName string

	//
	ExplicitCredentials uint64

	//
	FirstCipherAlgorithm uint32

	//
	FirstIntegrityAlgorithm uint32

	//
	FirstTransformType uint32

	//
	Flags uint32

	//
	InterfaceAlias string

	//
	IpProtocol uint8

	//
	LifetimePackets uint64

	//
	LocalEndpoint string

	//
	LocalPort uint16

	//
	LocalUdpEncapsulationPort uint16

	//
	MmSaId uint64

	//
	MmTargetName string

	//
	NapContext uint32

	//
	NdAllowClearTimeoutSeconds uint32

	//
	PeerV4PrivateAddress string

	//
	PfsGroupId uint32

	//
	QmSaId uint32

	//
	QuickModeFilterId uint64

	//
	RealIfProfileId uint64

	//
	RemoteEndpoint string

	//
	RemotePort uint16

	//
	RemoteUdpEncapsulationPort uint16

	//
	SecondCipherAlgorithm uint32

	//
	SecondIntegrityAlgorithm uint32

	//
	SecondSPI uint32

	//
	SecondTransformType uint32

	//
	TrafficLuid uint64

	//
	TrafficSelectorId uint64

	//
	TransportLayerFilterName string

	//
	VirtualIfTunnelId uint64
}

func NewMSFT_NetQuickModeSAEx1(instance *cim.WmiInstance) (newInstance *MSFT_NetQuickModeSA, err error) {
	tmp, err := NewCIM_IPsecSAEndpointEx1(instance)

	if err != nil {
		return
	}
	newInstance = &MSFT_NetQuickModeSA{
		CIM_IPsecSAEndpoint: tmp,
	}
	return
}

func NewMSFT_NetQuickModeSAEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *MSFT_NetQuickModeSA, err error) {
	tmp, err := NewCIM_IPsecSAEndpointEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &MSFT_NetQuickModeSA{
		CIM_IPsecSAEndpoint: tmp,
	}
	return
}

// SetEmTargetName sets the value of EmTargetName for the instance
func (instance *MSFT_NetQuickModeSA) SetPropertyEmTargetName(value string) (err error) {
	return instance.SetProperty("EmTargetName", (value))
}

// GetEmTargetName gets the value of EmTargetName for the instance
func (instance *MSFT_NetQuickModeSA) GetPropertyEmTargetName() (value string, err error) {
	retValue, err := instance.GetProperty("EmTargetName")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(string)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " string is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = string(valuetmp)

	return
}

// SetExplicitCredentials sets the value of ExplicitCredentials for the instance
func (instance *MSFT_NetQuickModeSA) SetPropertyExplicitCredentials(value uint64) (err error) {
	return instance.SetProperty("ExplicitCredentials", (value))
}

// GetExplicitCredentials gets the value of ExplicitCredentials for the instance
func (instance *MSFT_NetQuickModeSA) GetPropertyExplicitCredentials() (value uint64, err error) {
	retValue, err := instance.GetProperty("ExplicitCredentials")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(uint64)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " uint64 is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = uint64(valuetmp)

	return
}

// SetFirstCipherAlgorithm sets the value of FirstCipherAlgorithm for the instance
func (instance *MSFT_NetQuickModeSA) SetPropertyFirstCipherAlgorithm(value uint32) (err error) {
	return instance.SetProperty("FirstCipherAlgorithm", (value))
}

// GetFirstCipherAlgorithm gets the value of FirstCipherAlgorithm for the instance
func (instance *MSFT_NetQuickModeSA) GetPropertyFirstCipherAlgorithm() (value uint32, err error) {
	retValue, err := instance.GetProperty("FirstCipherAlgorithm")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(uint32)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " uint32 is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = uint32(valuetmp)

	return
}

// SetFirstIntegrityAlgorithm sets the value of FirstIntegrityAlgorithm for the instance
func (instance *MSFT_NetQuickModeSA) SetPropertyFirstIntegrityAlgorithm(value uint32) (err error) {
	return instance.SetProperty("FirstIntegrityAlgorithm", (value))
}

// GetFirstIntegrityAlgorithm gets the value of FirstIntegrityAlgorithm for the instance
func (instance *MSFT_NetQuickModeSA) GetPropertyFirstIntegrityAlgorithm() (value uint32, err error) {
	retValue, err := instance.GetProperty("FirstIntegrityAlgorithm")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(uint32)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " uint32 is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = uint32(valuetmp)

	return
}

// SetFirstTransformType sets the value of FirstTransformType for the instance
func (instance *MSFT_NetQuickModeSA) SetPropertyFirstTransformType(value uint32) (err error) {
	return instance.SetProperty("FirstTransformType", (value))
}

// GetFirstTransformType gets the value of FirstTransformType for the instance
func (instance *MSFT_NetQuickModeSA) GetPropertyFirstTransformType() (value uint32, err error) {
	retValue, err := instance.GetProperty("FirstTransformType")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(uint32)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " uint32 is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = uint32(valuetmp)

	return
}

// SetFlags sets the value of Flags for the instance
func (instance *MSFT_NetQuickModeSA) SetPropertyFlags(value uint32) (err error) {
	return instance.SetProperty("Flags", (value))
}

// GetFlags gets the value of Flags for the instance
func (instance *MSFT_NetQuickModeSA) GetPropertyFlags() (value uint32, err error) {
	retValue, err := instance.GetProperty("Flags")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(uint32)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " uint32 is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = uint32(valuetmp)

	return
}

// SetInterfaceAlias sets the value of InterfaceAlias for the instance
func (instance *MSFT_NetQuickModeSA) SetPropertyInterfaceAlias(value string) (err error) {
	return instance.SetProperty("InterfaceAlias", (value))
}

// GetInterfaceAlias gets the value of InterfaceAlias for the instance
func (instance *MSFT_NetQuickModeSA) GetPropertyInterfaceAlias() (value string, err error) {
	retValue, err := instance.GetProperty("InterfaceAlias")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(string)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " string is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = string(valuetmp)

	return
}

// SetIpProtocol sets the value of IpProtocol for the instance
func (instance *MSFT_NetQuickModeSA) SetPropertyIpProtocol(value uint8) (err error) {
	return instance.SetProperty("IpProtocol", (value))
}

// GetIpProtocol gets the value of IpProtocol for the instance
func (instance *MSFT_NetQuickModeSA) GetPropertyIpProtocol() (value uint8, err error) {
	retValue, err := instance.GetProperty("IpProtocol")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(uint8)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " uint8 is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = uint8(valuetmp)

	return
}

// SetLifetimePackets sets the value of LifetimePackets for the instance
func (instance *MSFT_NetQuickModeSA) SetPropertyLifetimePackets(value uint64) (err error) {
	return instance.SetProperty("LifetimePackets", (value))
}

// GetLifetimePackets gets the value of LifetimePackets for the instance
func (instance *MSFT_NetQuickModeSA) GetPropertyLifetimePackets() (value uint64, err error) {
	retValue, err := instance.GetProperty("LifetimePackets")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(uint64)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " uint64 is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = uint64(valuetmp)

	return
}

// SetLocalEndpoint sets the value of LocalEndpoint for the instance
func (instance *MSFT_NetQuickModeSA) SetPropertyLocalEndpoint(value string) (err error) {
	return instance.SetProperty("LocalEndpoint", (value))
}

// GetLocalEndpoint gets the value of LocalEndpoint for the instance
func (instance *MSFT_NetQuickModeSA) GetPropertyLocalEndpoint() (value string, err error) {
	retValue, err := instance.GetProperty("LocalEndpoint")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(string)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " string is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = string(valuetmp)

	return
}

// SetLocalPort sets the value of LocalPort for the instance
func (instance *MSFT_NetQuickModeSA) SetPropertyLocalPort(value uint16) (err error) {
	return instance.SetProperty("LocalPort", (value))
}

// GetLocalPort gets the value of LocalPort for the instance
func (instance *MSFT_NetQuickModeSA) GetPropertyLocalPort() (value uint16, err error) {
	retValue, err := instance.GetProperty("LocalPort")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(uint16)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " uint16 is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = uint16(valuetmp)

	return
}

// SetLocalUdpEncapsulationPort sets the value of LocalUdpEncapsulationPort for the instance
func (instance *MSFT_NetQuickModeSA) SetPropertyLocalUdpEncapsulationPort(value uint16) (err error) {
	return instance.SetProperty("LocalUdpEncapsulationPort", (value))
}

// GetLocalUdpEncapsulationPort gets the value of LocalUdpEncapsulationPort for the instance
func (instance *MSFT_NetQuickModeSA) GetPropertyLocalUdpEncapsulationPort() (value uint16, err error) {
	retValue, err := instance.GetProperty("LocalUdpEncapsulationPort")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(uint16)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " uint16 is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = uint16(valuetmp)

	return
}

// SetMmSaId sets the value of MmSaId for the instance
func (instance *MSFT_NetQuickModeSA) SetPropertyMmSaId(value uint64) (err error) {
	return instance.SetProperty("MmSaId", (value))
}

// GetMmSaId gets the value of MmSaId for the instance
func (instance *MSFT_NetQuickModeSA) GetPropertyMmSaId() (value uint64, err error) {
	retValue, err := instance.GetProperty("MmSaId")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(uint64)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " uint64 is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = uint64(valuetmp)

	return
}

// SetMmTargetName sets the value of MmTargetName for the instance
func (instance *MSFT_NetQuickModeSA) SetPropertyMmTargetName(value string) (err error) {
	return instance.SetProperty("MmTargetName", (value))
}

// GetMmTargetName gets the value of MmTargetName for the instance
func (instance *MSFT_NetQuickModeSA) GetPropertyMmTargetName() (value string, err error) {
	retValue, err := instance.GetProperty("MmTargetName")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(string)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " string is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = string(valuetmp)

	return
}

// SetNapContext sets the value of NapContext for the instance
func (instance *MSFT_NetQuickModeSA) SetPropertyNapContext(value uint32) (err error) {
	return instance.SetProperty("NapContext", (value))
}

// GetNapContext gets the value of NapContext for the instance
func (instance *MSFT_NetQuickModeSA) GetPropertyNapContext() (value uint32, err error) {
	retValue, err := instance.GetProperty("NapContext")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(uint32)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " uint32 is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = uint32(valuetmp)

	return
}

// SetNdAllowClearTimeoutSeconds sets the value of NdAllowClearTimeoutSeconds for the instance
func (instance *MSFT_NetQuickModeSA) SetPropertyNdAllowClearTimeoutSeconds(value uint32) (err error) {
	return instance.SetProperty("NdAllowClearTimeoutSeconds", (value))
}

// GetNdAllowClearTimeoutSeconds gets the value of NdAllowClearTimeoutSeconds for the instance
func (instance *MSFT_NetQuickModeSA) GetPropertyNdAllowClearTimeoutSeconds() (value uint32, err error) {
	retValue, err := instance.GetProperty("NdAllowClearTimeoutSeconds")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(uint32)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " uint32 is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = uint32(valuetmp)

	return
}

// SetPeerV4PrivateAddress sets the value of PeerV4PrivateAddress for the instance
func (instance *MSFT_NetQuickModeSA) SetPropertyPeerV4PrivateAddress(value string) (err error) {
	return instance.SetProperty("PeerV4PrivateAddress", (value))
}

// GetPeerV4PrivateAddress gets the value of PeerV4PrivateAddress for the instance
func (instance *MSFT_NetQuickModeSA) GetPropertyPeerV4PrivateAddress() (value string, err error) {
	retValue, err := instance.GetProperty("PeerV4PrivateAddress")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(string)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " string is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = string(valuetmp)

	return
}

// SetPfsGroupId sets the value of PfsGroupId for the instance
func (instance *MSFT_NetQuickModeSA) SetPropertyPfsGroupId(value uint32) (err error) {
	return instance.SetProperty("PfsGroupId", (value))
}

// GetPfsGroupId gets the value of PfsGroupId for the instance
func (instance *MSFT_NetQuickModeSA) GetPropertyPfsGroupId() (value uint32, err error) {
	retValue, err := instance.GetProperty("PfsGroupId")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(uint32)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " uint32 is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = uint32(valuetmp)

	return
}

// SetQmSaId sets the value of QmSaId for the instance
func (instance *MSFT_NetQuickModeSA) SetPropertyQmSaId(value uint32) (err error) {
	return instance.SetProperty("QmSaId", (value))
}

// GetQmSaId gets the value of QmSaId for the instance
func (instance *MSFT_NetQuickModeSA) GetPropertyQmSaId() (value uint32, err error) {
	retValue, err := instance.GetProperty("QmSaId")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(uint32)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " uint32 is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = uint32(valuetmp)

	return
}

// SetQuickModeFilterId sets the value of QuickModeFilterId for the instance
func (instance *MSFT_NetQuickModeSA) SetPropertyQuickModeFilterId(value uint64) (err error) {
	return instance.SetProperty("QuickModeFilterId", (value))
}

// GetQuickModeFilterId gets the value of QuickModeFilterId for the instance
func (instance *MSFT_NetQuickModeSA) GetPropertyQuickModeFilterId() (value uint64, err error) {
	retValue, err := instance.GetProperty("QuickModeFilterId")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(uint64)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " uint64 is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = uint64(valuetmp)

	return
}

// SetRealIfProfileId sets the value of RealIfProfileId for the instance
func (instance *MSFT_NetQuickModeSA) SetPropertyRealIfProfileId(value uint64) (err error) {
	return instance.SetProperty("RealIfProfileId", (value))
}

// GetRealIfProfileId gets the value of RealIfProfileId for the instance
func (instance *MSFT_NetQuickModeSA) GetPropertyRealIfProfileId() (value uint64, err error) {
	retValue, err := instance.GetProperty("RealIfProfileId")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(uint64)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " uint64 is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = uint64(valuetmp)

	return
}

// SetRemoteEndpoint sets the value of RemoteEndpoint for the instance
func (instance *MSFT_NetQuickModeSA) SetPropertyRemoteEndpoint(value string) (err error) {
	return instance.SetProperty("RemoteEndpoint", (value))
}

// GetRemoteEndpoint gets the value of RemoteEndpoint for the instance
func (instance *MSFT_NetQuickModeSA) GetPropertyRemoteEndpoint() (value string, err error) {
	retValue, err := instance.GetProperty("RemoteEndpoint")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(string)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " string is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = string(valuetmp)

	return
}

// SetRemotePort sets the value of RemotePort for the instance
func (instance *MSFT_NetQuickModeSA) SetPropertyRemotePort(value uint16) (err error) {
	return instance.SetProperty("RemotePort", (value))
}

// GetRemotePort gets the value of RemotePort for the instance
func (instance *MSFT_NetQuickModeSA) GetPropertyRemotePort() (value uint16, err error) {
	retValue, err := instance.GetProperty("RemotePort")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(uint16)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " uint16 is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = uint16(valuetmp)

	return
}

// SetRemoteUdpEncapsulationPort sets the value of RemoteUdpEncapsulationPort for the instance
func (instance *MSFT_NetQuickModeSA) SetPropertyRemoteUdpEncapsulationPort(value uint16) (err error) {
	return instance.SetProperty("RemoteUdpEncapsulationPort", (value))
}

// GetRemoteUdpEncapsulationPort gets the value of RemoteUdpEncapsulationPort for the instance
func (instance *MSFT_NetQuickModeSA) GetPropertyRemoteUdpEncapsulationPort() (value uint16, err error) {
	retValue, err := instance.GetProperty("RemoteUdpEncapsulationPort")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(uint16)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " uint16 is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = uint16(valuetmp)

	return
}

// SetSecondCipherAlgorithm sets the value of SecondCipherAlgorithm for the instance
func (instance *MSFT_NetQuickModeSA) SetPropertySecondCipherAlgorithm(value uint32) (err error) {
	return instance.SetProperty("SecondCipherAlgorithm", (value))
}

// GetSecondCipherAlgorithm gets the value of SecondCipherAlgorithm for the instance
func (instance *MSFT_NetQuickModeSA) GetPropertySecondCipherAlgorithm() (value uint32, err error) {
	retValue, err := instance.GetProperty("SecondCipherAlgorithm")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(uint32)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " uint32 is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = uint32(valuetmp)

	return
}

// SetSecondIntegrityAlgorithm sets the value of SecondIntegrityAlgorithm for the instance
func (instance *MSFT_NetQuickModeSA) SetPropertySecondIntegrityAlgorithm(value uint32) (err error) {
	return instance.SetProperty("SecondIntegrityAlgorithm", (value))
}

// GetSecondIntegrityAlgorithm gets the value of SecondIntegrityAlgorithm for the instance
func (instance *MSFT_NetQuickModeSA) GetPropertySecondIntegrityAlgorithm() (value uint32, err error) {
	retValue, err := instance.GetProperty("SecondIntegrityAlgorithm")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(uint32)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " uint32 is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = uint32(valuetmp)

	return
}

// SetSecondSPI sets the value of SecondSPI for the instance
func (instance *MSFT_NetQuickModeSA) SetPropertySecondSPI(value uint32) (err error) {
	return instance.SetProperty("SecondSPI", (value))
}

// GetSecondSPI gets the value of SecondSPI for the instance
func (instance *MSFT_NetQuickModeSA) GetPropertySecondSPI() (value uint32, err error) {
	retValue, err := instance.GetProperty("SecondSPI")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(uint32)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " uint32 is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = uint32(valuetmp)

	return
}

// SetSecondTransformType sets the value of SecondTransformType for the instance
func (instance *MSFT_NetQuickModeSA) SetPropertySecondTransformType(value uint32) (err error) {
	return instance.SetProperty("SecondTransformType", (value))
}

// GetSecondTransformType gets the value of SecondTransformType for the instance
func (instance *MSFT_NetQuickModeSA) GetPropertySecondTransformType() (value uint32, err error) {
	retValue, err := instance.GetProperty("SecondTransformType")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(uint32)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " uint32 is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = uint32(valuetmp)

	return
}

// SetTrafficLuid sets the value of TrafficLuid for the instance
func (instance *MSFT_NetQuickModeSA) SetPropertyTrafficLuid(value uint64) (err error) {
	return instance.SetProperty("TrafficLuid", (value))
}

// GetTrafficLuid gets the value of TrafficLuid for the instance
func (instance *MSFT_NetQuickModeSA) GetPropertyTrafficLuid() (value uint64, err error) {
	retValue, err := instance.GetProperty("TrafficLuid")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(uint64)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " uint64 is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = uint64(valuetmp)

	return
}

// SetTrafficSelectorId sets the value of TrafficSelectorId for the instance
func (instance *MSFT_NetQuickModeSA) SetPropertyTrafficSelectorId(value uint64) (err error) {
	return instance.SetProperty("TrafficSelectorId", (value))
}

// GetTrafficSelectorId gets the value of TrafficSelectorId for the instance
func (instance *MSFT_NetQuickModeSA) GetPropertyTrafficSelectorId() (value uint64, err error) {
	retValue, err := instance.GetProperty("TrafficSelectorId")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(uint64)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " uint64 is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = uint64(valuetmp)

	return
}

// SetTransportLayerFilterName sets the value of TransportLayerFilterName for the instance
func (instance *MSFT_NetQuickModeSA) SetPropertyTransportLayerFilterName(value string) (err error) {
	return instance.SetProperty("TransportLayerFilterName", (value))
}

// GetTransportLayerFilterName gets the value of TransportLayerFilterName for the instance
func (instance *MSFT_NetQuickModeSA) GetPropertyTransportLayerFilterName() (value string, err error) {
	retValue, err := instance.GetProperty("TransportLayerFilterName")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(string)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " string is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = string(valuetmp)

	return
}

// SetVirtualIfTunnelId sets the value of VirtualIfTunnelId for the instance
func (instance *MSFT_NetQuickModeSA) SetPropertyVirtualIfTunnelId(value uint64) (err error) {
	return instance.SetProperty("VirtualIfTunnelId", (value))
}

// GetVirtualIfTunnelId gets the value of VirtualIfTunnelId for the instance
func (instance *MSFT_NetQuickModeSA) GetPropertyVirtualIfTunnelId() (value uint64, err error) {
	retValue, err := instance.GetProperty("VirtualIfTunnelId")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(uint64)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " uint64 is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = uint64(valuetmp)

	return
}
