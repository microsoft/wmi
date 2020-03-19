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

// MDM_Firewall_FirewallRules02_01 struct
type MDM_Firewall_FirewallRules02_01 struct {
	*cim.WmiInstance

	//
	Description string

	//
	Direction string

	//
	EdgeTraversal bool

	//
	Enabled bool

	//
	FriendlyName bool

	//
	IcmpTypesAndCodes int32

	//
	InstanceID string

	//
	InterfaceTypes string

	//
	LocalAddressRanges string

	//
	LocalPortRanges string

	//
	LocalUserAuthorizationList string

	//
	Name string

	//
	ParentID string

	//
	Profiles int32

	//
	Protocol int32

	//
	RemoteAddressRanges string

	//
	RemotePortRanges string

	//
	Status int32
}

func NewMDM_Firewall_FirewallRules02_01Ex1(instance *cim.WmiInstance) (newInstance *MDM_Firewall_FirewallRules02_01, err error) {
	tmp, err := instance, nil

	if err != nil {
		return
	}
	newInstance = &MDM_Firewall_FirewallRules02_01{
		WmiInstance: tmp,
	}
	return
}

func NewMDM_Firewall_FirewallRules02_01Ex6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *MDM_Firewall_FirewallRules02_01, err error) {
	tmp, err := instance.GetWmiInstance(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &MDM_Firewall_FirewallRules02_01{
		WmiInstance: tmp,
	}
	return
}

// SetDescription sets the value of Description for the instance
func (instance *MDM_Firewall_FirewallRules02_01) SetPropertyDescription(value string) (err error) {
	return instance.SetProperty("Description", value)
}

// GetDescription gets the value of Description for the instance
func (instance *MDM_Firewall_FirewallRules02_01) GetPropertyDescription() (value string, err error) {
	retValue, err := instance.GetProperty("Description")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetDirection sets the value of Direction for the instance
func (instance *MDM_Firewall_FirewallRules02_01) SetPropertyDirection(value string) (err error) {
	return instance.SetProperty("Direction", value)
}

// GetDirection gets the value of Direction for the instance
func (instance *MDM_Firewall_FirewallRules02_01) GetPropertyDirection() (value string, err error) {
	retValue, err := instance.GetProperty("Direction")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetEdgeTraversal sets the value of EdgeTraversal for the instance
func (instance *MDM_Firewall_FirewallRules02_01) SetPropertyEdgeTraversal(value bool) (err error) {
	return instance.SetProperty("EdgeTraversal", value)
}

// GetEdgeTraversal gets the value of EdgeTraversal for the instance
func (instance *MDM_Firewall_FirewallRules02_01) GetPropertyEdgeTraversal() (value bool, err error) {
	retValue, err := instance.GetProperty("EdgeTraversal")
	if err != nil {
		return
	}
	value, ok := retValue.(bool)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetEnabled sets the value of Enabled for the instance
func (instance *MDM_Firewall_FirewallRules02_01) SetPropertyEnabled(value bool) (err error) {
	return instance.SetProperty("Enabled", value)
}

// GetEnabled gets the value of Enabled for the instance
func (instance *MDM_Firewall_FirewallRules02_01) GetPropertyEnabled() (value bool, err error) {
	retValue, err := instance.GetProperty("Enabled")
	if err != nil {
		return
	}
	value, ok := retValue.(bool)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetFriendlyName sets the value of FriendlyName for the instance
func (instance *MDM_Firewall_FirewallRules02_01) SetPropertyFriendlyName(value bool) (err error) {
	return instance.SetProperty("FriendlyName", value)
}

// GetFriendlyName gets the value of FriendlyName for the instance
func (instance *MDM_Firewall_FirewallRules02_01) GetPropertyFriendlyName() (value bool, err error) {
	retValue, err := instance.GetProperty("FriendlyName")
	if err != nil {
		return
	}
	value, ok := retValue.(bool)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetIcmpTypesAndCodes sets the value of IcmpTypesAndCodes for the instance
func (instance *MDM_Firewall_FirewallRules02_01) SetPropertyIcmpTypesAndCodes(value int32) (err error) {
	return instance.SetProperty("IcmpTypesAndCodes", value)
}

// GetIcmpTypesAndCodes gets the value of IcmpTypesAndCodes for the instance
func (instance *MDM_Firewall_FirewallRules02_01) GetPropertyIcmpTypesAndCodes() (value int32, err error) {
	retValue, err := instance.GetProperty("IcmpTypesAndCodes")
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
func (instance *MDM_Firewall_FirewallRules02_01) SetPropertyInstanceID(value string) (err error) {
	return instance.SetProperty("InstanceID", value)
}

// GetInstanceID gets the value of InstanceID for the instance
func (instance *MDM_Firewall_FirewallRules02_01) GetPropertyInstanceID() (value string, err error) {
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

// SetInterfaceTypes sets the value of InterfaceTypes for the instance
func (instance *MDM_Firewall_FirewallRules02_01) SetPropertyInterfaceTypes(value string) (err error) {
	return instance.SetProperty("InterfaceTypes", value)
}

// GetInterfaceTypes gets the value of InterfaceTypes for the instance
func (instance *MDM_Firewall_FirewallRules02_01) GetPropertyInterfaceTypes() (value string, err error) {
	retValue, err := instance.GetProperty("InterfaceTypes")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetLocalAddressRanges sets the value of LocalAddressRanges for the instance
func (instance *MDM_Firewall_FirewallRules02_01) SetPropertyLocalAddressRanges(value string) (err error) {
	return instance.SetProperty("LocalAddressRanges", value)
}

// GetLocalAddressRanges gets the value of LocalAddressRanges for the instance
func (instance *MDM_Firewall_FirewallRules02_01) GetPropertyLocalAddressRanges() (value string, err error) {
	retValue, err := instance.GetProperty("LocalAddressRanges")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetLocalPortRanges sets the value of LocalPortRanges for the instance
func (instance *MDM_Firewall_FirewallRules02_01) SetPropertyLocalPortRanges(value string) (err error) {
	return instance.SetProperty("LocalPortRanges", value)
}

// GetLocalPortRanges gets the value of LocalPortRanges for the instance
func (instance *MDM_Firewall_FirewallRules02_01) GetPropertyLocalPortRanges() (value string, err error) {
	retValue, err := instance.GetProperty("LocalPortRanges")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetLocalUserAuthorizationList sets the value of LocalUserAuthorizationList for the instance
func (instance *MDM_Firewall_FirewallRules02_01) SetPropertyLocalUserAuthorizationList(value string) (err error) {
	return instance.SetProperty("LocalUserAuthorizationList", value)
}

// GetLocalUserAuthorizationList gets the value of LocalUserAuthorizationList for the instance
func (instance *MDM_Firewall_FirewallRules02_01) GetPropertyLocalUserAuthorizationList() (value string, err error) {
	retValue, err := instance.GetProperty("LocalUserAuthorizationList")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetName sets the value of Name for the instance
func (instance *MDM_Firewall_FirewallRules02_01) SetPropertyName(value string) (err error) {
	return instance.SetProperty("Name", value)
}

// GetName gets the value of Name for the instance
func (instance *MDM_Firewall_FirewallRules02_01) GetPropertyName() (value string, err error) {
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

// SetParentID sets the value of ParentID for the instance
func (instance *MDM_Firewall_FirewallRules02_01) SetPropertyParentID(value string) (err error) {
	return instance.SetProperty("ParentID", value)
}

// GetParentID gets the value of ParentID for the instance
func (instance *MDM_Firewall_FirewallRules02_01) GetPropertyParentID() (value string, err error) {
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

// SetProfiles sets the value of Profiles for the instance
func (instance *MDM_Firewall_FirewallRules02_01) SetPropertyProfiles(value int32) (err error) {
	return instance.SetProperty("Profiles", value)
}

// GetProfiles gets the value of Profiles for the instance
func (instance *MDM_Firewall_FirewallRules02_01) GetPropertyProfiles() (value int32, err error) {
	retValue, err := instance.GetProperty("Profiles")
	if err != nil {
		return
	}
	value, ok := retValue.(int32)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetProtocol sets the value of Protocol for the instance
func (instance *MDM_Firewall_FirewallRules02_01) SetPropertyProtocol(value int32) (err error) {
	return instance.SetProperty("Protocol", value)
}

// GetProtocol gets the value of Protocol for the instance
func (instance *MDM_Firewall_FirewallRules02_01) GetPropertyProtocol() (value int32, err error) {
	retValue, err := instance.GetProperty("Protocol")
	if err != nil {
		return
	}
	value, ok := retValue.(int32)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetRemoteAddressRanges sets the value of RemoteAddressRanges for the instance
func (instance *MDM_Firewall_FirewallRules02_01) SetPropertyRemoteAddressRanges(value string) (err error) {
	return instance.SetProperty("RemoteAddressRanges", value)
}

// GetRemoteAddressRanges gets the value of RemoteAddressRanges for the instance
func (instance *MDM_Firewall_FirewallRules02_01) GetPropertyRemoteAddressRanges() (value string, err error) {
	retValue, err := instance.GetProperty("RemoteAddressRanges")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetRemotePortRanges sets the value of RemotePortRanges for the instance
func (instance *MDM_Firewall_FirewallRules02_01) SetPropertyRemotePortRanges(value string) (err error) {
	return instance.SetProperty("RemotePortRanges", value)
}

// GetRemotePortRanges gets the value of RemotePortRanges for the instance
func (instance *MDM_Firewall_FirewallRules02_01) GetPropertyRemotePortRanges() (value string, err error) {
	retValue, err := instance.GetProperty("RemotePortRanges")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetStatus sets the value of Status for the instance
func (instance *MDM_Firewall_FirewallRules02_01) SetPropertyStatus(value int32) (err error) {
	return instance.SetProperty("Status", value)
}

// GetStatus gets the value of Status for the instance
func (instance *MDM_Firewall_FirewallRules02_01) GetPropertyStatus() (value int32, err error) {
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
