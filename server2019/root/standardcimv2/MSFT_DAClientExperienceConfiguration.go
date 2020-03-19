// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 3/19/2020 using wmigen
//      Source root.StandardCimv2
//////////////////////////////////////////////
package standardcimv2

import (
	"github.com/microsoft/wmi/pkg/base/query"
	cim "github.com/microsoft/wmi/pkg/wmiinstance"
)

// MSFT_DAClientExperienceConfiguration struct
type MSFT_DAClientExperienceConfiguration struct {
	*MSFT_NetSettingData

	//
	CorporateResources []string

	//
	CustomCommands []string

	//
	ForceTunneling uint32

	//
	FriendlyName string

	//
	GslbFqdn string

	//
	IPsecTunnelEndpoints []string

	//
	ManualEntryPointSelectionAllowed bool

	//
	PassiveMode bool

	//
	PolicyStore string

	//
	PreferLocalNamesAllowed bool

	//
	SupportEmail string

	//
	UserInterface bool
}

func NewMSFT_DAClientExperienceConfigurationEx1(instance *cim.WmiInstance) (newInstance *MSFT_DAClientExperienceConfiguration, err error) {
	tmp, err := NewMSFT_NetSettingDataEx1(instance)

	if err != nil {
		return
	}
	newInstance = &MSFT_DAClientExperienceConfiguration{
		MSFT_NetSettingData: tmp,
	}
	return
}

func NewMSFT_DAClientExperienceConfigurationEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *MSFT_DAClientExperienceConfiguration, err error) {
	tmp, err := NewMSFT_NetSettingDataEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &MSFT_DAClientExperienceConfiguration{
		MSFT_NetSettingData: tmp,
	}
	return
}

// SetCorporateResources sets the value of CorporateResources for the instance
func (instance *MSFT_DAClientExperienceConfiguration) SetPropertyCorporateResources(value []string) (err error) {
	return instance.SetProperty("CorporateResources", value)
}

// GetCorporateResources gets the value of CorporateResources for the instance
func (instance *MSFT_DAClientExperienceConfiguration) GetPropertyCorporateResources() (value []string, err error) {
	retValue, err := instance.GetProperty("CorporateResources")
	if err != nil {
		return
	}
	value, ok := retValue.([]string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetCustomCommands sets the value of CustomCommands for the instance
func (instance *MSFT_DAClientExperienceConfiguration) SetPropertyCustomCommands(value []string) (err error) {
	return instance.SetProperty("CustomCommands", value)
}

// GetCustomCommands gets the value of CustomCommands for the instance
func (instance *MSFT_DAClientExperienceConfiguration) GetPropertyCustomCommands() (value []string, err error) {
	retValue, err := instance.GetProperty("CustomCommands")
	if err != nil {
		return
	}
	value, ok := retValue.([]string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetForceTunneling sets the value of ForceTunneling for the instance
func (instance *MSFT_DAClientExperienceConfiguration) SetPropertyForceTunneling(value uint32) (err error) {
	return instance.SetProperty("ForceTunneling", value)
}

// GetForceTunneling gets the value of ForceTunneling for the instance
func (instance *MSFT_DAClientExperienceConfiguration) GetPropertyForceTunneling() (value uint32, err error) {
	retValue, err := instance.GetProperty("ForceTunneling")
	if err != nil {
		return
	}
	value, ok := retValue.(uint32)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetFriendlyName sets the value of FriendlyName for the instance
func (instance *MSFT_DAClientExperienceConfiguration) SetPropertyFriendlyName(value string) (err error) {
	return instance.SetProperty("FriendlyName", value)
}

// GetFriendlyName gets the value of FriendlyName for the instance
func (instance *MSFT_DAClientExperienceConfiguration) GetPropertyFriendlyName() (value string, err error) {
	retValue, err := instance.GetProperty("FriendlyName")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetGslbFqdn sets the value of GslbFqdn for the instance
func (instance *MSFT_DAClientExperienceConfiguration) SetPropertyGslbFqdn(value string) (err error) {
	return instance.SetProperty("GslbFqdn", value)
}

// GetGslbFqdn gets the value of GslbFqdn for the instance
func (instance *MSFT_DAClientExperienceConfiguration) GetPropertyGslbFqdn() (value string, err error) {
	retValue, err := instance.GetProperty("GslbFqdn")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetIPsecTunnelEndpoints sets the value of IPsecTunnelEndpoints for the instance
func (instance *MSFT_DAClientExperienceConfiguration) SetPropertyIPsecTunnelEndpoints(value []string) (err error) {
	return instance.SetProperty("IPsecTunnelEndpoints", value)
}

// GetIPsecTunnelEndpoints gets the value of IPsecTunnelEndpoints for the instance
func (instance *MSFT_DAClientExperienceConfiguration) GetPropertyIPsecTunnelEndpoints() (value []string, err error) {
	retValue, err := instance.GetProperty("IPsecTunnelEndpoints")
	if err != nil {
		return
	}
	value, ok := retValue.([]string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetManualEntryPointSelectionAllowed sets the value of ManualEntryPointSelectionAllowed for the instance
func (instance *MSFT_DAClientExperienceConfiguration) SetPropertyManualEntryPointSelectionAllowed(value bool) (err error) {
	return instance.SetProperty("ManualEntryPointSelectionAllowed", value)
}

// GetManualEntryPointSelectionAllowed gets the value of ManualEntryPointSelectionAllowed for the instance
func (instance *MSFT_DAClientExperienceConfiguration) GetPropertyManualEntryPointSelectionAllowed() (value bool, err error) {
	retValue, err := instance.GetProperty("ManualEntryPointSelectionAllowed")
	if err != nil {
		return
	}
	value, ok := retValue.(bool)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetPassiveMode sets the value of PassiveMode for the instance
func (instance *MSFT_DAClientExperienceConfiguration) SetPropertyPassiveMode(value bool) (err error) {
	return instance.SetProperty("PassiveMode", value)
}

// GetPassiveMode gets the value of PassiveMode for the instance
func (instance *MSFT_DAClientExperienceConfiguration) GetPropertyPassiveMode() (value bool, err error) {
	retValue, err := instance.GetProperty("PassiveMode")
	if err != nil {
		return
	}
	value, ok := retValue.(bool)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetPolicyStore sets the value of PolicyStore for the instance
func (instance *MSFT_DAClientExperienceConfiguration) SetPropertyPolicyStore(value string) (err error) {
	return instance.SetProperty("PolicyStore", value)
}

// GetPolicyStore gets the value of PolicyStore for the instance
func (instance *MSFT_DAClientExperienceConfiguration) GetPropertyPolicyStore() (value string, err error) {
	retValue, err := instance.GetProperty("PolicyStore")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetPreferLocalNamesAllowed sets the value of PreferLocalNamesAllowed for the instance
func (instance *MSFT_DAClientExperienceConfiguration) SetPropertyPreferLocalNamesAllowed(value bool) (err error) {
	return instance.SetProperty("PreferLocalNamesAllowed", value)
}

// GetPreferLocalNamesAllowed gets the value of PreferLocalNamesAllowed for the instance
func (instance *MSFT_DAClientExperienceConfiguration) GetPropertyPreferLocalNamesAllowed() (value bool, err error) {
	retValue, err := instance.GetProperty("PreferLocalNamesAllowed")
	if err != nil {
		return
	}
	value, ok := retValue.(bool)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetSupportEmail sets the value of SupportEmail for the instance
func (instance *MSFT_DAClientExperienceConfiguration) SetPropertySupportEmail(value string) (err error) {
	return instance.SetProperty("SupportEmail", value)
}

// GetSupportEmail gets the value of SupportEmail for the instance
func (instance *MSFT_DAClientExperienceConfiguration) GetPropertySupportEmail() (value string, err error) {
	retValue, err := instance.GetProperty("SupportEmail")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetUserInterface sets the value of UserInterface for the instance
func (instance *MSFT_DAClientExperienceConfiguration) SetPropertyUserInterface(value bool) (err error) {
	return instance.SetProperty("UserInterface", value)
}

// GetUserInterface gets the value of UserInterface for the instance
func (instance *MSFT_DAClientExperienceConfiguration) GetPropertyUserInterface() (value bool, err error) {
	retValue, err := instance.GetProperty("UserInterface")
	if err != nil {
		return
	}
	value, ok := retValue.(bool)
	if !ok {
		// TODO: Set an error
	}
	return
}

//

// <param name="CorporateResources" type="bool "></param>
// <param name="CustomCommands" type="bool "></param>
// <param name="ForceTunneling" type="bool "></param>
// <param name="FriendlyName" type="bool "></param>
// <param name="GslbFqdn" type="bool "></param>
// <param name="IPsecTunnelEndpoints" type="bool "></param>
// <param name="ManualEntryPointSelectionAllowed" type="bool "></param>
// <param name="PassiveMode" type="bool "></param>
// <param name="PassThru" type="bool "></param>
// <param name="PreferLocalNamesAllowed" type="bool "></param>
// <param name="SupportEmail" type="bool "></param>
// <param name="UserInterface" type="bool "></param>

// <param name="OutputObject" type="MSFT_DAClientExperienceConfiguration "></param>
// <param name="ReturnValue" type="uint32 "></param>
func (instance *MSFT_DAClientExperienceConfiguration) Reset( /* IN */ CorporateResources bool,
	/* IN */ IPsecTunnelEndpoints bool,
	/* IN */ PreferLocalNamesAllowed bool,
	/* IN */ UserInterface bool,
	/* IN */ SupportEmail bool,
	/* IN */ FriendlyName bool,
	/* IN */ PassiveMode bool,
	/* IN */ CustomCommands bool,
	/* IN */ ManualEntryPointSelectionAllowed bool,
	/* IN */ GslbFqdn bool,
	/* IN */ ForceTunneling bool,
	/* IN */ PassThru bool,
	/* OUT */ OutputObject MSFT_DAClientExperienceConfiguration) (result uint32, err error) {
	retVal, err := instance.InvokeMethod("Reset", CorporateResources, IPsecTunnelEndpoints, PreferLocalNamesAllowed, UserInterface, SupportEmail, FriendlyName, PassiveMode, CustomCommands, ManualEntryPointSelectionAllowed, GslbFqdn, ForceTunneling, PassThru)
	if err != nil {
		return
	}
	retValue := retVal[0].(int32)
	result = uint32(retValue)
	return

}
