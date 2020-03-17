// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 3/16/2020 using wmigen
//      Source root.RSOP.Computer
//////////////////////////////////////////////
package computer

// RSOP_FolderRedirectionPolicySetting struct
type RSOP_FolderRedirectionPolicySetting struct {
	RSOP_PolicySetting

	//
	configurationControl uint32

	//
	grantType bool

	//
	installationType uint32

	//
	moveType bool

	//
	parentFolderId string

	//
	policyRemoval uint32

	//
	primaryComputerEvaluation uint32

	//
	redirectedPaths []string

	//
	redirectingGroup string

	//
	redirectionFlags uint32

	//
	resultantPath string

	//
	securityGroups []string
}

// SetconfigurationControl sets the value of configurationControl for the instance
func (instance *RSOP_FolderRedirectionPolicySetting) SetPropertyconfigurationControl(value uint32) (err error) {
	return instance.SetProperty("configurationControl", value)
}

// GetconfigurationControl gets the value of configurationControl for the instance
func (instance *RSOP_FolderRedirectionPolicySetting) GetPropertyconfigurationControl() (value uint32, err error) {
	retValue, err := instance.GetProperty("configurationControl")
	if err != nil {
		return
	}
	value, ok := retValue.(uint32)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetgrantType sets the value of grantType for the instance
func (instance *RSOP_FolderRedirectionPolicySetting) SetPropertygrantType(value bool) (err error) {
	return instance.SetProperty("grantType", value)
}

// GetgrantType gets the value of grantType for the instance
func (instance *RSOP_FolderRedirectionPolicySetting) GetPropertygrantType() (value bool, err error) {
	retValue, err := instance.GetProperty("grantType")
	if err != nil {
		return
	}
	value, ok := retValue.(bool)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetinstallationType sets the value of installationType for the instance
func (instance *RSOP_FolderRedirectionPolicySetting) SetPropertyinstallationType(value uint32) (err error) {
	return instance.SetProperty("installationType", value)
}

// GetinstallationType gets the value of installationType for the instance
func (instance *RSOP_FolderRedirectionPolicySetting) GetPropertyinstallationType() (value uint32, err error) {
	retValue, err := instance.GetProperty("installationType")
	if err != nil {
		return
	}
	value, ok := retValue.(uint32)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetmoveType sets the value of moveType for the instance
func (instance *RSOP_FolderRedirectionPolicySetting) SetPropertymoveType(value bool) (err error) {
	return instance.SetProperty("moveType", value)
}

// GetmoveType gets the value of moveType for the instance
func (instance *RSOP_FolderRedirectionPolicySetting) GetPropertymoveType() (value bool, err error) {
	retValue, err := instance.GetProperty("moveType")
	if err != nil {
		return
	}
	value, ok := retValue.(bool)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetparentFolderId sets the value of parentFolderId for the instance
func (instance *RSOP_FolderRedirectionPolicySetting) SetPropertyparentFolderId(value string) (err error) {
	return instance.SetProperty("parentFolderId", value)
}

// GetparentFolderId gets the value of parentFolderId for the instance
func (instance *RSOP_FolderRedirectionPolicySetting) GetPropertyparentFolderId() (value string, err error) {
	retValue, err := instance.GetProperty("parentFolderId")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetpolicyRemoval sets the value of policyRemoval for the instance
func (instance *RSOP_FolderRedirectionPolicySetting) SetPropertypolicyRemoval(value uint32) (err error) {
	return instance.SetProperty("policyRemoval", value)
}

// GetpolicyRemoval gets the value of policyRemoval for the instance
func (instance *RSOP_FolderRedirectionPolicySetting) GetPropertypolicyRemoval() (value uint32, err error) {
	retValue, err := instance.GetProperty("policyRemoval")
	if err != nil {
		return
	}
	value, ok := retValue.(uint32)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetprimaryComputerEvaluation sets the value of primaryComputerEvaluation for the instance
func (instance *RSOP_FolderRedirectionPolicySetting) SetPropertyprimaryComputerEvaluation(value uint32) (err error) {
	return instance.SetProperty("primaryComputerEvaluation", value)
}

// GetprimaryComputerEvaluation gets the value of primaryComputerEvaluation for the instance
func (instance *RSOP_FolderRedirectionPolicySetting) GetPropertyprimaryComputerEvaluation() (value uint32, err error) {
	retValue, err := instance.GetProperty("primaryComputerEvaluation")
	if err != nil {
		return
	}
	value, ok := retValue.(uint32)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetredirectedPaths sets the value of redirectedPaths for the instance
func (instance *RSOP_FolderRedirectionPolicySetting) SetPropertyredirectedPaths(value []string) (err error) {
	return instance.SetProperty("redirectedPaths", value)
}

// GetredirectedPaths gets the value of redirectedPaths for the instance
func (instance *RSOP_FolderRedirectionPolicySetting) GetPropertyredirectedPaths() (value []string, err error) {
	retValue, err := instance.GetProperty("redirectedPaths")
	if err != nil {
		return
	}
	value, ok := retValue.([]string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetredirectingGroup sets the value of redirectingGroup for the instance
func (instance *RSOP_FolderRedirectionPolicySetting) SetPropertyredirectingGroup(value string) (err error) {
	return instance.SetProperty("redirectingGroup", value)
}

// GetredirectingGroup gets the value of redirectingGroup for the instance
func (instance *RSOP_FolderRedirectionPolicySetting) GetPropertyredirectingGroup() (value string, err error) {
	retValue, err := instance.GetProperty("redirectingGroup")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetredirectionFlags sets the value of redirectionFlags for the instance
func (instance *RSOP_FolderRedirectionPolicySetting) SetPropertyredirectionFlags(value uint32) (err error) {
	return instance.SetProperty("redirectionFlags", value)
}

// GetredirectionFlags gets the value of redirectionFlags for the instance
func (instance *RSOP_FolderRedirectionPolicySetting) GetPropertyredirectionFlags() (value uint32, err error) {
	retValue, err := instance.GetProperty("redirectionFlags")
	if err != nil {
		return
	}
	value, ok := retValue.(uint32)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetresultantPath sets the value of resultantPath for the instance
func (instance *RSOP_FolderRedirectionPolicySetting) SetPropertyresultantPath(value string) (err error) {
	return instance.SetProperty("resultantPath", value)
}

// GetresultantPath gets the value of resultantPath for the instance
func (instance *RSOP_FolderRedirectionPolicySetting) GetPropertyresultantPath() (value string, err error) {
	retValue, err := instance.GetProperty("resultantPath")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetsecurityGroups sets the value of securityGroups for the instance
func (instance *RSOP_FolderRedirectionPolicySetting) SetPropertysecurityGroups(value []string) (err error) {
	return instance.SetProperty("securityGroups", value)
}

// GetsecurityGroups gets the value of securityGroups for the instance
func (instance *RSOP_FolderRedirectionPolicySetting) GetPropertysecurityGroups() (value []string, err error) {
	retValue, err := instance.GetProperty("securityGroups")
	if err != nil {
		return
	}
	value, ok := retValue.([]string)
	if !ok {
		// TODO: Set an error
	}
	return
}
