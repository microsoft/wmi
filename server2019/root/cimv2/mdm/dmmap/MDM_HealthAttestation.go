// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 3/19/2020 using wmigen
//      Source root.CIMV2.mdm.dmmap
//////////////////////////////////////////////
package dmmap

import (
	"github.com/microsoft/wmi/pkg/base/instance"
	"github.com/microsoft/wmi/pkg/base/query"
	cim "github.com/microsoft/wmi/pkg/wmiinstance"
)

// MDM_HealthAttestation struct
type MDM_HealthAttestation struct {
	*cim.WmiInstance

	//
	Certificate string

	//
	CorrelationID string

	//
	CurrentProtocolVersion int32

	//
	ForceRetrieve bool

	//
	HASEndpoint string

	//
	InstanceID string

	//
	MaxSupportedProtocolVersion int32

	//
	Nonce string

	//
	ParentID string

	//
	PreferredMaxProtocolVersion int32

	//
	Status int32

	//
	TpmReadyStatus int32
}

func NewMDM_HealthAttestationEx1(instance *cim.WmiInstance) (newInstance *MDM_HealthAttestation, err error) {
	tmp, err := instance, nil

	if err != nil {
		return
	}
	newInstance = &MDM_HealthAttestation{
		WmiInstance: tmp,
	}
	return
}

func NewMDM_HealthAttestationEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *MDM_HealthAttestation, err error) {
	tmp, err := instance.GetWmiInstance(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &MDM_HealthAttestation{
		WmiInstance: tmp,
	}
	return
}

// SetCertificate sets the value of Certificate for the instance
func (instance *MDM_HealthAttestation) SetPropertyCertificate(value string) (err error) {
	return instance.SetProperty("Certificate", value)
}

// GetCertificate gets the value of Certificate for the instance
func (instance *MDM_HealthAttestation) GetPropertyCertificate() (value string, err error) {
	retValue, err := instance.GetProperty("Certificate")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetCorrelationID sets the value of CorrelationID for the instance
func (instance *MDM_HealthAttestation) SetPropertyCorrelationID(value string) (err error) {
	return instance.SetProperty("CorrelationID", value)
}

// GetCorrelationID gets the value of CorrelationID for the instance
func (instance *MDM_HealthAttestation) GetPropertyCorrelationID() (value string, err error) {
	retValue, err := instance.GetProperty("CorrelationID")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetCurrentProtocolVersion sets the value of CurrentProtocolVersion for the instance
func (instance *MDM_HealthAttestation) SetPropertyCurrentProtocolVersion(value int32) (err error) {
	return instance.SetProperty("CurrentProtocolVersion", value)
}

// GetCurrentProtocolVersion gets the value of CurrentProtocolVersion for the instance
func (instance *MDM_HealthAttestation) GetPropertyCurrentProtocolVersion() (value int32, err error) {
	retValue, err := instance.GetProperty("CurrentProtocolVersion")
	if err != nil {
		return
	}
	value, ok := retValue.(int32)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetForceRetrieve sets the value of ForceRetrieve for the instance
func (instance *MDM_HealthAttestation) SetPropertyForceRetrieve(value bool) (err error) {
	return instance.SetProperty("ForceRetrieve", value)
}

// GetForceRetrieve gets the value of ForceRetrieve for the instance
func (instance *MDM_HealthAttestation) GetPropertyForceRetrieve() (value bool, err error) {
	retValue, err := instance.GetProperty("ForceRetrieve")
	if err != nil {
		return
	}
	value, ok := retValue.(bool)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetHASEndpoint sets the value of HASEndpoint for the instance
func (instance *MDM_HealthAttestation) SetPropertyHASEndpoint(value string) (err error) {
	return instance.SetProperty("HASEndpoint", value)
}

// GetHASEndpoint gets the value of HASEndpoint for the instance
func (instance *MDM_HealthAttestation) GetPropertyHASEndpoint() (value string, err error) {
	retValue, err := instance.GetProperty("HASEndpoint")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetInstanceID sets the value of InstanceID for the instance
func (instance *MDM_HealthAttestation) SetPropertyInstanceID(value string) (err error) {
	return instance.SetProperty("InstanceID", value)
}

// GetInstanceID gets the value of InstanceID for the instance
func (instance *MDM_HealthAttestation) GetPropertyInstanceID() (value string, err error) {
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

// SetMaxSupportedProtocolVersion sets the value of MaxSupportedProtocolVersion for the instance
func (instance *MDM_HealthAttestation) SetPropertyMaxSupportedProtocolVersion(value int32) (err error) {
	return instance.SetProperty("MaxSupportedProtocolVersion", value)
}

// GetMaxSupportedProtocolVersion gets the value of MaxSupportedProtocolVersion for the instance
func (instance *MDM_HealthAttestation) GetPropertyMaxSupportedProtocolVersion() (value int32, err error) {
	retValue, err := instance.GetProperty("MaxSupportedProtocolVersion")
	if err != nil {
		return
	}
	value, ok := retValue.(int32)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetNonce sets the value of Nonce for the instance
func (instance *MDM_HealthAttestation) SetPropertyNonce(value string) (err error) {
	return instance.SetProperty("Nonce", value)
}

// GetNonce gets the value of Nonce for the instance
func (instance *MDM_HealthAttestation) GetPropertyNonce() (value string, err error) {
	retValue, err := instance.GetProperty("Nonce")
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
func (instance *MDM_HealthAttestation) SetPropertyParentID(value string) (err error) {
	return instance.SetProperty("ParentID", value)
}

// GetParentID gets the value of ParentID for the instance
func (instance *MDM_HealthAttestation) GetPropertyParentID() (value string, err error) {
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

// SetPreferredMaxProtocolVersion sets the value of PreferredMaxProtocolVersion for the instance
func (instance *MDM_HealthAttestation) SetPropertyPreferredMaxProtocolVersion(value int32) (err error) {
	return instance.SetProperty("PreferredMaxProtocolVersion", value)
}

// GetPreferredMaxProtocolVersion gets the value of PreferredMaxProtocolVersion for the instance
func (instance *MDM_HealthAttestation) GetPropertyPreferredMaxProtocolVersion() (value int32, err error) {
	retValue, err := instance.GetProperty("PreferredMaxProtocolVersion")
	if err != nil {
		return
	}
	value, ok := retValue.(int32)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetStatus sets the value of Status for the instance
func (instance *MDM_HealthAttestation) SetPropertyStatus(value int32) (err error) {
	return instance.SetProperty("Status", value)
}

// GetStatus gets the value of Status for the instance
func (instance *MDM_HealthAttestation) GetPropertyStatus() (value int32, err error) {
	retValue, err := instance.GetProperty("Status")
	if err != nil {
		return
	}
	value, ok := retValue.(int32)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetTpmReadyStatus sets the value of TpmReadyStatus for the instance
func (instance *MDM_HealthAttestation) SetPropertyTpmReadyStatus(value int32) (err error) {
	return instance.SetProperty("TpmReadyStatus", value)
}

// GetTpmReadyStatus gets the value of TpmReadyStatus for the instance
func (instance *MDM_HealthAttestation) GetPropertyTpmReadyStatus() (value int32, err error) {
	retValue, err := instance.GetProperty("TpmReadyStatus")
	if err != nil {
		return
	}
	value, ok := retValue.(int32)
	if !ok {
		// TODO: Set an error
	}
	return
}

//

// <param name="ReturnValue" type="uint32 "></param>
func (instance *MDM_HealthAttestation) VerifyHealthMethod() (result uint32, err error) {
	retVal, err := instance.InvokeMethodWithReturn("VerifyHealthMethod")
	if err != nil {
		return
	}
	result = uint32(retVal)
	return

}
