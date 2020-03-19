// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 3/19/2020 using wmigen
//      Source root.Microsoft.Windows.DesiredStateConfiguration
//////////////////////////////////////////////
package desiredstateconfiguration

import (
	"github.com/microsoft/wmi/pkg/base/instance"
	"github.com/microsoft/wmi/pkg/base/query"
	cim "github.com/microsoft/wmi/pkg/wmiinstance"
)

// MSFT_DSCConfigurationStatus struct
type MSFT_DSCConfigurationStatus struct {
	*cim.WmiInstance

	//
	DurationInSeconds uint32

	//
	Error string

	//
	HostName string

	//
	IPV4Addresses []string

	//
	IPV6Addresses []string

	//
	JobID string

	//
	LCMVersion string

	//
	Locale string

	//
	MACAddresses []string

	//
	MetaConfiguration MSFT_DSCMetaConfiguration

	//
	MetaData string

	//
	Mode string

	//
	NumberOfResources uint32

	//
	RebootRequested bool

	//
	ResourcesChanged []MSFT_ResourceChanged

	//
	ResourcesInDesiredState []MSFT_ResourceInDesiredState

	//
	ResourcesNotInDesiredState []MSFT_ResourceNotInDesiredState

	//
	StartDate string

	//
	Status string

	//
	Type string
}

func NewMSFT_DSCConfigurationStatusEx1(instance *cim.WmiInstance) (newInstance *MSFT_DSCConfigurationStatus, err error) {
	tmp, err := instance, nil

	if err != nil {
		return
	}
	newInstance = &MSFT_DSCConfigurationStatus{
		WmiInstance: tmp,
	}
	return
}

func NewMSFT_DSCConfigurationStatusEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *MSFT_DSCConfigurationStatus, err error) {
	tmp, err := instance.GetWmiInstance(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &MSFT_DSCConfigurationStatus{
		WmiInstance: tmp,
	}
	return
}

// SetDurationInSeconds sets the value of DurationInSeconds for the instance
func (instance *MSFT_DSCConfigurationStatus) SetPropertyDurationInSeconds(value uint32) (err error) {
	return instance.SetProperty("DurationInSeconds", value)
}

// GetDurationInSeconds gets the value of DurationInSeconds for the instance
func (instance *MSFT_DSCConfigurationStatus) GetPropertyDurationInSeconds() (value uint32, err error) {
	retValue, err := instance.GetProperty("DurationInSeconds")
	if err != nil {
		return
	}
	value, ok := retValue.(uint32)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetError sets the value of Error for the instance
func (instance *MSFT_DSCConfigurationStatus) SetPropertyError(value string) (err error) {
	return instance.SetProperty("Error", value)
}

// GetError gets the value of Error for the instance
func (instance *MSFT_DSCConfigurationStatus) GetPropertyError() (value string, err error) {
	retValue, err := instance.GetProperty("Error")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetHostName sets the value of HostName for the instance
func (instance *MSFT_DSCConfigurationStatus) SetPropertyHostName(value string) (err error) {
	return instance.SetProperty("HostName", value)
}

// GetHostName gets the value of HostName for the instance
func (instance *MSFT_DSCConfigurationStatus) GetPropertyHostName() (value string, err error) {
	retValue, err := instance.GetProperty("HostName")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetIPV4Addresses sets the value of IPV4Addresses for the instance
func (instance *MSFT_DSCConfigurationStatus) SetPropertyIPV4Addresses(value []string) (err error) {
	return instance.SetProperty("IPV4Addresses", value)
}

// GetIPV4Addresses gets the value of IPV4Addresses for the instance
func (instance *MSFT_DSCConfigurationStatus) GetPropertyIPV4Addresses() (value []string, err error) {
	retValue, err := instance.GetProperty("IPV4Addresses")
	if err != nil {
		return
	}
	value, ok := retValue.([]string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetIPV6Addresses sets the value of IPV6Addresses for the instance
func (instance *MSFT_DSCConfigurationStatus) SetPropertyIPV6Addresses(value []string) (err error) {
	return instance.SetProperty("IPV6Addresses", value)
}

// GetIPV6Addresses gets the value of IPV6Addresses for the instance
func (instance *MSFT_DSCConfigurationStatus) GetPropertyIPV6Addresses() (value []string, err error) {
	retValue, err := instance.GetProperty("IPV6Addresses")
	if err != nil {
		return
	}
	value, ok := retValue.([]string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetJobID sets the value of JobID for the instance
func (instance *MSFT_DSCConfigurationStatus) SetPropertyJobID(value string) (err error) {
	return instance.SetProperty("JobID", value)
}

// GetJobID gets the value of JobID for the instance
func (instance *MSFT_DSCConfigurationStatus) GetPropertyJobID() (value string, err error) {
	retValue, err := instance.GetProperty("JobID")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetLCMVersion sets the value of LCMVersion for the instance
func (instance *MSFT_DSCConfigurationStatus) SetPropertyLCMVersion(value string) (err error) {
	return instance.SetProperty("LCMVersion", value)
}

// GetLCMVersion gets the value of LCMVersion for the instance
func (instance *MSFT_DSCConfigurationStatus) GetPropertyLCMVersion() (value string, err error) {
	retValue, err := instance.GetProperty("LCMVersion")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetLocale sets the value of Locale for the instance
func (instance *MSFT_DSCConfigurationStatus) SetPropertyLocale(value string) (err error) {
	return instance.SetProperty("Locale", value)
}

// GetLocale gets the value of Locale for the instance
func (instance *MSFT_DSCConfigurationStatus) GetPropertyLocale() (value string, err error) {
	retValue, err := instance.GetProperty("Locale")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetMACAddresses sets the value of MACAddresses for the instance
func (instance *MSFT_DSCConfigurationStatus) SetPropertyMACAddresses(value []string) (err error) {
	return instance.SetProperty("MACAddresses", value)
}

// GetMACAddresses gets the value of MACAddresses for the instance
func (instance *MSFT_DSCConfigurationStatus) GetPropertyMACAddresses() (value []string, err error) {
	retValue, err := instance.GetProperty("MACAddresses")
	if err != nil {
		return
	}
	value, ok := retValue.([]string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetMetaConfiguration sets the value of MetaConfiguration for the instance
func (instance *MSFT_DSCConfigurationStatus) SetPropertyMetaConfiguration(value MSFT_DSCMetaConfiguration) (err error) {
	return instance.SetProperty("MetaConfiguration", value)
}

// GetMetaConfiguration gets the value of MetaConfiguration for the instance
func (instance *MSFT_DSCConfigurationStatus) GetPropertyMetaConfiguration() (value MSFT_DSCMetaConfiguration, err error) {
	retValue, err := instance.GetProperty("MetaConfiguration")
	if err != nil {
		return
	}
	value, ok := retValue.(MSFT_DSCMetaConfiguration)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetMetaData sets the value of MetaData for the instance
func (instance *MSFT_DSCConfigurationStatus) SetPropertyMetaData(value string) (err error) {
	return instance.SetProperty("MetaData", value)
}

// GetMetaData gets the value of MetaData for the instance
func (instance *MSFT_DSCConfigurationStatus) GetPropertyMetaData() (value string, err error) {
	retValue, err := instance.GetProperty("MetaData")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetMode sets the value of Mode for the instance
func (instance *MSFT_DSCConfigurationStatus) SetPropertyMode(value string) (err error) {
	return instance.SetProperty("Mode", value)
}

// GetMode gets the value of Mode for the instance
func (instance *MSFT_DSCConfigurationStatus) GetPropertyMode() (value string, err error) {
	retValue, err := instance.GetProperty("Mode")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetNumberOfResources sets the value of NumberOfResources for the instance
func (instance *MSFT_DSCConfigurationStatus) SetPropertyNumberOfResources(value uint32) (err error) {
	return instance.SetProperty("NumberOfResources", value)
}

// GetNumberOfResources gets the value of NumberOfResources for the instance
func (instance *MSFT_DSCConfigurationStatus) GetPropertyNumberOfResources() (value uint32, err error) {
	retValue, err := instance.GetProperty("NumberOfResources")
	if err != nil {
		return
	}
	value, ok := retValue.(uint32)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetRebootRequested sets the value of RebootRequested for the instance
func (instance *MSFT_DSCConfigurationStatus) SetPropertyRebootRequested(value bool) (err error) {
	return instance.SetProperty("RebootRequested", value)
}

// GetRebootRequested gets the value of RebootRequested for the instance
func (instance *MSFT_DSCConfigurationStatus) GetPropertyRebootRequested() (value bool, err error) {
	retValue, err := instance.GetProperty("RebootRequested")
	if err != nil {
		return
	}
	value, ok := retValue.(bool)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetResourcesChanged sets the value of ResourcesChanged for the instance
func (instance *MSFT_DSCConfigurationStatus) SetPropertyResourcesChanged(value []MSFT_ResourceChanged) (err error) {
	return instance.SetProperty("ResourcesChanged", value)
}

// GetResourcesChanged gets the value of ResourcesChanged for the instance
func (instance *MSFT_DSCConfigurationStatus) GetPropertyResourcesChanged() (value []MSFT_ResourceChanged, err error) {
	retValue, err := instance.GetProperty("ResourcesChanged")
	if err != nil {
		return
	}
	value, ok := retValue.([]MSFT_ResourceChanged)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetResourcesInDesiredState sets the value of ResourcesInDesiredState for the instance
func (instance *MSFT_DSCConfigurationStatus) SetPropertyResourcesInDesiredState(value []MSFT_ResourceInDesiredState) (err error) {
	return instance.SetProperty("ResourcesInDesiredState", value)
}

// GetResourcesInDesiredState gets the value of ResourcesInDesiredState for the instance
func (instance *MSFT_DSCConfigurationStatus) GetPropertyResourcesInDesiredState() (value []MSFT_ResourceInDesiredState, err error) {
	retValue, err := instance.GetProperty("ResourcesInDesiredState")
	if err != nil {
		return
	}
	value, ok := retValue.([]MSFT_ResourceInDesiredState)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetResourcesNotInDesiredState sets the value of ResourcesNotInDesiredState for the instance
func (instance *MSFT_DSCConfigurationStatus) SetPropertyResourcesNotInDesiredState(value []MSFT_ResourceNotInDesiredState) (err error) {
	return instance.SetProperty("ResourcesNotInDesiredState", value)
}

// GetResourcesNotInDesiredState gets the value of ResourcesNotInDesiredState for the instance
func (instance *MSFT_DSCConfigurationStatus) GetPropertyResourcesNotInDesiredState() (value []MSFT_ResourceNotInDesiredState, err error) {
	retValue, err := instance.GetProperty("ResourcesNotInDesiredState")
	if err != nil {
		return
	}
	value, ok := retValue.([]MSFT_ResourceNotInDesiredState)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetStartDate sets the value of StartDate for the instance
func (instance *MSFT_DSCConfigurationStatus) SetPropertyStartDate(value string) (err error) {
	return instance.SetProperty("StartDate", value)
}

// GetStartDate gets the value of StartDate for the instance
func (instance *MSFT_DSCConfigurationStatus) GetPropertyStartDate() (value string, err error) {
	retValue, err := instance.GetProperty("StartDate")
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
func (instance *MSFT_DSCConfigurationStatus) SetPropertyStatus(value string) (err error) {
	return instance.SetProperty("Status", value)
}

// GetStatus gets the value of Status for the instance
func (instance *MSFT_DSCConfigurationStatus) GetPropertyStatus() (value string, err error) {
	retValue, err := instance.GetProperty("Status")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetType sets the value of Type for the instance
func (instance *MSFT_DSCConfigurationStatus) SetPropertyType(value string) (err error) {
	return instance.SetProperty("Type", value)
}

// GetType gets the value of Type for the instance
func (instance *MSFT_DSCConfigurationStatus) GetPropertyType() (value string, err error) {
	retValue, err := instance.GetProperty("Type")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}
