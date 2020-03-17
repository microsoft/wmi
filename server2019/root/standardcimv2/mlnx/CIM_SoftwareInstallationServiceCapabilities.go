// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 3/16/2020 using wmigen
//      Source root.StandardCimv2.mlnx
//////////////////////////////////////////////
package mlnx

// CIM_SoftwareInstallationServiceCapabilities struct
type CIM_SoftwareInstallationServiceCapabilities struct {
	CIM_Capabilities

	//
	CanAddToCollection bool

	//
	OtherSupportedExtendedResourceTypeDescriptions []string

	//
	SupportedAsynchronousActions []SoftwareInstallationServiceCapabilities_SupportedAsynchronousActions

	//
	SupportedExtendedResourceTypes []SoftwareInstallationServiceCapabilities_SupportedExtendedResourceTypes

	//
	SupportedExtendedResourceTypesBuildNumbers []uint16

	//
	SupportedExtendedResourceTypesMajorVersions []uint16

	//
	SupportedExtendedResourceTypesMinorVersions []uint16

	//
	SupportedExtendedResourceTypesRevisionNumbers []uint16

	//
	SupportedInstallOptions []SoftwareInstallationServiceCapabilities_SupportedInstallOptions

	//
	SupportedSynchronousActions []SoftwareInstallationServiceCapabilities_SupportedSynchronousActions

	//
	SupportedTargetTypes []string

	//
	SupportedURISchemes []SoftwareInstallationServiceCapabilities_SupportedURISchemes
}

// SetCanAddToCollection sets the value of CanAddToCollection for the instance
func (instance *CIM_SoftwareInstallationServiceCapabilities) SetPropertyCanAddToCollection(value bool) (err error) {
	return instance.SetProperty("CanAddToCollection", value)
}

// GetCanAddToCollection gets the value of CanAddToCollection for the instance
func (instance *CIM_SoftwareInstallationServiceCapabilities) GetPropertyCanAddToCollection() (value bool, err error) {
	retValue, err := instance.GetProperty("CanAddToCollection")
	if err != nil {
		return
	}
	value, ok := retValue.(bool)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetOtherSupportedExtendedResourceTypeDescriptions sets the value of OtherSupportedExtendedResourceTypeDescriptions for the instance
func (instance *CIM_SoftwareInstallationServiceCapabilities) SetPropertyOtherSupportedExtendedResourceTypeDescriptions(value []string) (err error) {
	return instance.SetProperty("OtherSupportedExtendedResourceTypeDescriptions", value)
}

// GetOtherSupportedExtendedResourceTypeDescriptions gets the value of OtherSupportedExtendedResourceTypeDescriptions for the instance
func (instance *CIM_SoftwareInstallationServiceCapabilities) GetPropertyOtherSupportedExtendedResourceTypeDescriptions() (value []string, err error) {
	retValue, err := instance.GetProperty("OtherSupportedExtendedResourceTypeDescriptions")
	if err != nil {
		return
	}
	value, ok := retValue.([]string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetSupportedAsynchronousActions sets the value of SupportedAsynchronousActions for the instance
func (instance *CIM_SoftwareInstallationServiceCapabilities) SetPropertySupportedAsynchronousActions(value []SoftwareInstallationServiceCapabilities_SupportedAsynchronousActions) (err error) {
	return instance.SetProperty("SupportedAsynchronousActions", value)
}

// GetSupportedAsynchronousActions gets the value of SupportedAsynchronousActions for the instance
func (instance *CIM_SoftwareInstallationServiceCapabilities) GetPropertySupportedAsynchronousActions() (value []SoftwareInstallationServiceCapabilities_SupportedAsynchronousActions, err error) {
	retValue, err := instance.GetProperty("SupportedAsynchronousActions")
	if err != nil {
		return
	}
	value, ok := retValue.([]SoftwareInstallationServiceCapabilities_SupportedAsynchronousActions)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetSupportedExtendedResourceTypes sets the value of SupportedExtendedResourceTypes for the instance
func (instance *CIM_SoftwareInstallationServiceCapabilities) SetPropertySupportedExtendedResourceTypes(value []SoftwareInstallationServiceCapabilities_SupportedExtendedResourceTypes) (err error) {
	return instance.SetProperty("SupportedExtendedResourceTypes", value)
}

// GetSupportedExtendedResourceTypes gets the value of SupportedExtendedResourceTypes for the instance
func (instance *CIM_SoftwareInstallationServiceCapabilities) GetPropertySupportedExtendedResourceTypes() (value []SoftwareInstallationServiceCapabilities_SupportedExtendedResourceTypes, err error) {
	retValue, err := instance.GetProperty("SupportedExtendedResourceTypes")
	if err != nil {
		return
	}
	value, ok := retValue.([]SoftwareInstallationServiceCapabilities_SupportedExtendedResourceTypes)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetSupportedExtendedResourceTypesBuildNumbers sets the value of SupportedExtendedResourceTypesBuildNumbers for the instance
func (instance *CIM_SoftwareInstallationServiceCapabilities) SetPropertySupportedExtendedResourceTypesBuildNumbers(value []uint16) (err error) {
	return instance.SetProperty("SupportedExtendedResourceTypesBuildNumbers", value)
}

// GetSupportedExtendedResourceTypesBuildNumbers gets the value of SupportedExtendedResourceTypesBuildNumbers for the instance
func (instance *CIM_SoftwareInstallationServiceCapabilities) GetPropertySupportedExtendedResourceTypesBuildNumbers() (value []uint16, err error) {
	retValue, err := instance.GetProperty("SupportedExtendedResourceTypesBuildNumbers")
	if err != nil {
		return
	}
	value, ok := retValue.([]uint16)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetSupportedExtendedResourceTypesMajorVersions sets the value of SupportedExtendedResourceTypesMajorVersions for the instance
func (instance *CIM_SoftwareInstallationServiceCapabilities) SetPropertySupportedExtendedResourceTypesMajorVersions(value []uint16) (err error) {
	return instance.SetProperty("SupportedExtendedResourceTypesMajorVersions", value)
}

// GetSupportedExtendedResourceTypesMajorVersions gets the value of SupportedExtendedResourceTypesMajorVersions for the instance
func (instance *CIM_SoftwareInstallationServiceCapabilities) GetPropertySupportedExtendedResourceTypesMajorVersions() (value []uint16, err error) {
	retValue, err := instance.GetProperty("SupportedExtendedResourceTypesMajorVersions")
	if err != nil {
		return
	}
	value, ok := retValue.([]uint16)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetSupportedExtendedResourceTypesMinorVersions sets the value of SupportedExtendedResourceTypesMinorVersions for the instance
func (instance *CIM_SoftwareInstallationServiceCapabilities) SetPropertySupportedExtendedResourceTypesMinorVersions(value []uint16) (err error) {
	return instance.SetProperty("SupportedExtendedResourceTypesMinorVersions", value)
}

// GetSupportedExtendedResourceTypesMinorVersions gets the value of SupportedExtendedResourceTypesMinorVersions for the instance
func (instance *CIM_SoftwareInstallationServiceCapabilities) GetPropertySupportedExtendedResourceTypesMinorVersions() (value []uint16, err error) {
	retValue, err := instance.GetProperty("SupportedExtendedResourceTypesMinorVersions")
	if err != nil {
		return
	}
	value, ok := retValue.([]uint16)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetSupportedExtendedResourceTypesRevisionNumbers sets the value of SupportedExtendedResourceTypesRevisionNumbers for the instance
func (instance *CIM_SoftwareInstallationServiceCapabilities) SetPropertySupportedExtendedResourceTypesRevisionNumbers(value []uint16) (err error) {
	return instance.SetProperty("SupportedExtendedResourceTypesRevisionNumbers", value)
}

// GetSupportedExtendedResourceTypesRevisionNumbers gets the value of SupportedExtendedResourceTypesRevisionNumbers for the instance
func (instance *CIM_SoftwareInstallationServiceCapabilities) GetPropertySupportedExtendedResourceTypesRevisionNumbers() (value []uint16, err error) {
	retValue, err := instance.GetProperty("SupportedExtendedResourceTypesRevisionNumbers")
	if err != nil {
		return
	}
	value, ok := retValue.([]uint16)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetSupportedInstallOptions sets the value of SupportedInstallOptions for the instance
func (instance *CIM_SoftwareInstallationServiceCapabilities) SetPropertySupportedInstallOptions(value []SoftwareInstallationServiceCapabilities_SupportedInstallOptions) (err error) {
	return instance.SetProperty("SupportedInstallOptions", value)
}

// GetSupportedInstallOptions gets the value of SupportedInstallOptions for the instance
func (instance *CIM_SoftwareInstallationServiceCapabilities) GetPropertySupportedInstallOptions() (value []SoftwareInstallationServiceCapabilities_SupportedInstallOptions, err error) {
	retValue, err := instance.GetProperty("SupportedInstallOptions")
	if err != nil {
		return
	}
	value, ok := retValue.([]SoftwareInstallationServiceCapabilities_SupportedInstallOptions)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetSupportedSynchronousActions sets the value of SupportedSynchronousActions for the instance
func (instance *CIM_SoftwareInstallationServiceCapabilities) SetPropertySupportedSynchronousActions(value []SoftwareInstallationServiceCapabilities_SupportedSynchronousActions) (err error) {
	return instance.SetProperty("SupportedSynchronousActions", value)
}

// GetSupportedSynchronousActions gets the value of SupportedSynchronousActions for the instance
func (instance *CIM_SoftwareInstallationServiceCapabilities) GetPropertySupportedSynchronousActions() (value []SoftwareInstallationServiceCapabilities_SupportedSynchronousActions, err error) {
	retValue, err := instance.GetProperty("SupportedSynchronousActions")
	if err != nil {
		return
	}
	value, ok := retValue.([]SoftwareInstallationServiceCapabilities_SupportedSynchronousActions)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetSupportedTargetTypes sets the value of SupportedTargetTypes for the instance
func (instance *CIM_SoftwareInstallationServiceCapabilities) SetPropertySupportedTargetTypes(value []string) (err error) {
	return instance.SetProperty("SupportedTargetTypes", value)
}

// GetSupportedTargetTypes gets the value of SupportedTargetTypes for the instance
func (instance *CIM_SoftwareInstallationServiceCapabilities) GetPropertySupportedTargetTypes() (value []string, err error) {
	retValue, err := instance.GetProperty("SupportedTargetTypes")
	if err != nil {
		return
	}
	value, ok := retValue.([]string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetSupportedURISchemes sets the value of SupportedURISchemes for the instance
func (instance *CIM_SoftwareInstallationServiceCapabilities) SetPropertySupportedURISchemes(value []SoftwareInstallationServiceCapabilities_SupportedURISchemes) (err error) {
	return instance.SetProperty("SupportedURISchemes", value)
}

// GetSupportedURISchemes gets the value of SupportedURISchemes for the instance
func (instance *CIM_SoftwareInstallationServiceCapabilities) GetPropertySupportedURISchemes() (value []SoftwareInstallationServiceCapabilities_SupportedURISchemes, err error) {
	retValue, err := instance.GetProperty("SupportedURISchemes")
	if err != nil {
		return
	}
	value, ok := retValue.([]SoftwareInstallationServiceCapabilities_SupportedURISchemes)
	if !ok {
		// TODO: Set an error
	}
	return
}
