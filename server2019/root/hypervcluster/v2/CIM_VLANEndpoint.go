// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

// Author:
//
//	Auto Generated on 6/6/2024 using wmigen
//	Source root.HyperVCluster.v2
//
// ////////////////////////////////////////////
package v2

import (
	"github.com/microsoft/wmi/pkg/base/query"
	"github.com/microsoft/wmi/pkg/errors"
	cim "github.com/microsoft/wmi/pkg/wmiinstance"
	"reflect"
)

// CIM_VLANEndpoint struct
type CIM_VLANEndpoint struct {
	*CIM_ProtocolEndpoint

	// The desired VLAN mode that is requested for use. (Note that the current mode is given by the OperationalEndpointMode property.) The following values are defined:
	///- Access: Puts the endpoint/switch port into permanent nontrunking mode and negotiates to convert the link into a nontrunk link. The endpoint becomes a nontrunk interface.
	///- Dynamic Auto: Makes the endpoint able to convert the link to a trunk link. The endpoint becomes a trunk interface if the neighboring interface is set to trunk or desirable mode.
	///- Dynamic Desirable: Makes the endpoint actively attempt to convert the link to a trunk link. The endpoint becomes a trunk interface if the neighboring interface is set to trunk, desirable, or auto mode. The default switch-port mode for all Ethernet interfaces is 'dynamic desirable.'
	///- Trunk: Puts the endpoint into permanent trunking mode and negotiates to convert the link into a trunk link. The endpoint becomes a trunk interface even if the neighboring interface is not a trunk interface.
	///- Dot1Q Tunnel: Configures the interface as a tunnel (nontrunking) endpoint/port to be connected in an asymmetric link with an 802.1Q trunk port. 802.1Q tunneling is used to maintain customer VLAN integrity across a service provider network.
	DesiredEndpointMode VLANEndpoint_DesiredEndpointMode

	// The type of VLAN encapsulation that is requested for use. (Note that the encapsulation currently in use is given by the OperationalVLANTrunkEncapsulation property.) Note that this property is only applicable when the endpoint is operating in a trunking mode (see the OperationalEndpointMode property for additional details). This property is either 'not applicable' (i.e., the endpoint will never be placed in a trunking mode), a particular type (802.1q or Cisco ISL), or 'negotiate' (i.e., the result of the negotiation between this interface and its neighbor). The value, 'Negotiate' is not allowed if the endpoint does not support negotiation. This capability is hardware and vendor dependent. Refer to the associated VLANEndpointCapabilities.doesTrunkEncapsulationNegotiation property to validate whether a particular endpoint (port) supports encapsulation negotiation.
	DesiredVLANTrunkEncapsulation VLANEndpoint_DesiredVLANTrunkEncapsulation

	// Indicates whether GARP VLAN Registration Protocol (GVRP) is enabled or disabled on the trunk endpoint/port. This property is 'not applicable' unless GVRP is supported by the endpoint. This is indicated in the Capabilities property, VLANEndpointCapabilities.Dot1QTagging. This property is applicable only when the endpoint is operating in trunking mode (determined by examining the SwitchEndpointMode property).
	GVRPStatus VLANEndpoint_GVRPStatus

	// The configuration mode for the VLAN endpoint. The following values are defined: /n - Unknown: If the endpoint is not VLAN aware. /n - Access: Puts the endpoint into permanent nontrunking mode and negotiates to convert the link into a nontrunk link. The endpoint becomes a nontrunk interface.
	///- Dynamic Auto: Makes the endpoint able to convert the link to a trunk link. The endpoint becomes a trunk interface if the neighboring interface is set to trunk or desirable mode.
	///- Dynamic Desirable: Makes the endpoint actively attempt to convert the link to a trunk link. The endpoint becomes a trunk interface if the neighboring interface is set to trunk, desirable, or auto mode. The default switch-port mode for all Ethernet interfaces is 'dynamic desirable.'
	///- Trunk: Puts the endpoint into permanent trunking mode and negotiates to convert the link into a trunk link. The endpoint becomes a trunk interface even if the neighboring interface is not a trunk interface.
	///- Dot1Q Tunnel: Configures the interface as a tunnel (nontrunking) endpoint/port to be connected in an asymmetric link with an 802.1Q trunk port. 802.1Q tunneling is used to maintain customer VLAN integrity across a service provider network.
	OperationalEndpointMode VLANEndpoint_OperationalEndpointMode

	// The type of VLAN encapsulation in use on a trunk endpoint/port. This property is either 'not applicable' (i.e., the endpoint is not operating in trunking mode), a particular type (802.1q or Cisco ISL), 'negotiating' (i.e., the endpoints are negotiating the encapsulation type). Note that this property is only applicable when the endpoint is operating in a trunking mode (see the OperationalEndpointMode property for additional details).
	OperationalVLANTrunkEncapsulation VLANEndpoint_OperationalVLANTrunkEncapsulation

	// A string describing the type of VLAN endpoint model that is supported by this VLANEndpoint, when the value of the mode property is set to 1 (i.e., "Other"). This property should be set to NULL when the mode property is any value other than 1.
	OtherEndpointMode string

	// A string describing the type of VLAN encapsulation that is supported by this VLANEndpoint, when the value of the encapsulation property is set to 1 (i.e., "Other"). This property should be set to NULL when the desired encapsulation property is any value other than 1.
	OtherTrunkEncapsulation string
}

func NewCIM_VLANEndpointEx1(instance *cim.WmiInstance) (newInstance *CIM_VLANEndpoint, err error) {
	tmp, err := NewCIM_ProtocolEndpointEx1(instance)

	if err != nil {
		return
	}
	newInstance = &CIM_VLANEndpoint{
		CIM_ProtocolEndpoint: tmp,
	}
	return
}

func NewCIM_VLANEndpointEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *CIM_VLANEndpoint, err error) {
	tmp, err := NewCIM_ProtocolEndpointEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &CIM_VLANEndpoint{
		CIM_ProtocolEndpoint: tmp,
	}
	return
}

// SetDesiredEndpointMode sets the value of DesiredEndpointMode for the instance
func (instance *CIM_VLANEndpoint) SetPropertyDesiredEndpointMode(value VLANEndpoint_DesiredEndpointMode) (err error) {
	return instance.SetProperty("DesiredEndpointMode", (value))
}

// GetDesiredEndpointMode gets the value of DesiredEndpointMode for the instance
func (instance *CIM_VLANEndpoint) GetPropertyDesiredEndpointMode() (value VLANEndpoint_DesiredEndpointMode, err error) {
	retValue, err := instance.GetProperty("DesiredEndpointMode")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(int32)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " int32 is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = VLANEndpoint_DesiredEndpointMode(valuetmp)

	return
}

// SetDesiredVLANTrunkEncapsulation sets the value of DesiredVLANTrunkEncapsulation for the instance
func (instance *CIM_VLANEndpoint) SetPropertyDesiredVLANTrunkEncapsulation(value VLANEndpoint_DesiredVLANTrunkEncapsulation) (err error) {
	return instance.SetProperty("DesiredVLANTrunkEncapsulation", (value))
}

// GetDesiredVLANTrunkEncapsulation gets the value of DesiredVLANTrunkEncapsulation for the instance
func (instance *CIM_VLANEndpoint) GetPropertyDesiredVLANTrunkEncapsulation() (value VLANEndpoint_DesiredVLANTrunkEncapsulation, err error) {
	retValue, err := instance.GetProperty("DesiredVLANTrunkEncapsulation")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(int32)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " int32 is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = VLANEndpoint_DesiredVLANTrunkEncapsulation(valuetmp)

	return
}

// SetGVRPStatus sets the value of GVRPStatus for the instance
func (instance *CIM_VLANEndpoint) SetPropertyGVRPStatus(value VLANEndpoint_GVRPStatus) (err error) {
	return instance.SetProperty("GVRPStatus", (value))
}

// GetGVRPStatus gets the value of GVRPStatus for the instance
func (instance *CIM_VLANEndpoint) GetPropertyGVRPStatus() (value VLANEndpoint_GVRPStatus, err error) {
	retValue, err := instance.GetProperty("GVRPStatus")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(int32)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " int32 is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = VLANEndpoint_GVRPStatus(valuetmp)

	return
}

// SetOperationalEndpointMode sets the value of OperationalEndpointMode for the instance
func (instance *CIM_VLANEndpoint) SetPropertyOperationalEndpointMode(value VLANEndpoint_OperationalEndpointMode) (err error) {
	return instance.SetProperty("OperationalEndpointMode", (value))
}

// GetOperationalEndpointMode gets the value of OperationalEndpointMode for the instance
func (instance *CIM_VLANEndpoint) GetPropertyOperationalEndpointMode() (value VLANEndpoint_OperationalEndpointMode, err error) {
	retValue, err := instance.GetProperty("OperationalEndpointMode")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(int32)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " int32 is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = VLANEndpoint_OperationalEndpointMode(valuetmp)

	return
}

// SetOperationalVLANTrunkEncapsulation sets the value of OperationalVLANTrunkEncapsulation for the instance
func (instance *CIM_VLANEndpoint) SetPropertyOperationalVLANTrunkEncapsulation(value VLANEndpoint_OperationalVLANTrunkEncapsulation) (err error) {
	return instance.SetProperty("OperationalVLANTrunkEncapsulation", (value))
}

// GetOperationalVLANTrunkEncapsulation gets the value of OperationalVLANTrunkEncapsulation for the instance
func (instance *CIM_VLANEndpoint) GetPropertyOperationalVLANTrunkEncapsulation() (value VLANEndpoint_OperationalVLANTrunkEncapsulation, err error) {
	retValue, err := instance.GetProperty("OperationalVLANTrunkEncapsulation")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(int32)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " int32 is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = VLANEndpoint_OperationalVLANTrunkEncapsulation(valuetmp)

	return
}

// SetOtherEndpointMode sets the value of OtherEndpointMode for the instance
func (instance *CIM_VLANEndpoint) SetPropertyOtherEndpointMode(value string) (err error) {
	return instance.SetProperty("OtherEndpointMode", (value))
}

// GetOtherEndpointMode gets the value of OtherEndpointMode for the instance
func (instance *CIM_VLANEndpoint) GetPropertyOtherEndpointMode() (value string, err error) {
	retValue, err := instance.GetProperty("OtherEndpointMode")
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

// SetOtherTrunkEncapsulation sets the value of OtherTrunkEncapsulation for the instance
func (instance *CIM_VLANEndpoint) SetPropertyOtherTrunkEncapsulation(value string) (err error) {
	return instance.SetProperty("OtherTrunkEncapsulation", (value))
}

// GetOtherTrunkEncapsulation gets the value of OtherTrunkEncapsulation for the instance
func (instance *CIM_VLANEndpoint) GetPropertyOtherTrunkEncapsulation() (value string, err error) {
	retValue, err := instance.GetProperty("OtherTrunkEncapsulation")
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
