// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 3/16/2020 using wmigen
//      Source root.Microsoft.Uev
//////////////////////////////////////////////
package uev

import (
	"github.com/microsoft/wmi/pkg/wmiinstance"
)

// SettingsLocationTemplate struct
type SettingsLocationTemplate struct {
	cim.WmiInstance

	// Flag indicating whether the template defers to MS account.
	DeferToMSAccount bool

	// Flag indicating whether the template is enabled for the current user.
	Enabled bool

	// Enable state of the settings location template.
	EnableStateLocation string

	// Flag indicating whether the template is a suite parent.
	IsSuiteParent bool

	// Flag indicating whether the template represents a template file.
	IsTemplateFile bool

	// ID of the suite parent template.
	SuiteParentId string

	// Description of the settings location template.
	TemplateDescription string

	// Unique ID of the settings location template.
	TemplateId string

	// Friendly name of the settings location template.
	TemplateName string

	// Profile the template is associated with.
	TemplateProfile string

	// Type of the settings location template (OS, Application).
	TemplateType string

	// Version of the settings location template.
	TemplateVersion uint32
}

// SetDeferToMSAccount sets the value of DeferToMSAccount for the instance
func (instance *SettingsLocationTemplate) SetPropertyDeferToMSAccount(value bool) (err error) {
	return instance.SetProperty("DeferToMSAccount", value)
}

// GetDeferToMSAccount gets the value of DeferToMSAccount for the instance
func (instance *SettingsLocationTemplate) GetPropertyDeferToMSAccount() (value bool, err error) {
	retValue, err := instance.GetProperty("DeferToMSAccount")
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
func (instance *SettingsLocationTemplate) SetPropertyEnabled(value bool) (err error) {
	return instance.SetProperty("Enabled", value)
}

// GetEnabled gets the value of Enabled for the instance
func (instance *SettingsLocationTemplate) GetPropertyEnabled() (value bool, err error) {
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

// SetEnableStateLocation sets the value of EnableStateLocation for the instance
func (instance *SettingsLocationTemplate) SetPropertyEnableStateLocation(value string) (err error) {
	return instance.SetProperty("EnableStateLocation", value)
}

// GetEnableStateLocation gets the value of EnableStateLocation for the instance
func (instance *SettingsLocationTemplate) GetPropertyEnableStateLocation() (value string, err error) {
	retValue, err := instance.GetProperty("EnableStateLocation")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetIsSuiteParent sets the value of IsSuiteParent for the instance
func (instance *SettingsLocationTemplate) SetPropertyIsSuiteParent(value bool) (err error) {
	return instance.SetProperty("IsSuiteParent", value)
}

// GetIsSuiteParent gets the value of IsSuiteParent for the instance
func (instance *SettingsLocationTemplate) GetPropertyIsSuiteParent() (value bool, err error) {
	retValue, err := instance.GetProperty("IsSuiteParent")
	if err != nil {
		return
	}
	value, ok := retValue.(bool)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetIsTemplateFile sets the value of IsTemplateFile for the instance
func (instance *SettingsLocationTemplate) SetPropertyIsTemplateFile(value bool) (err error) {
	return instance.SetProperty("IsTemplateFile", value)
}

// GetIsTemplateFile gets the value of IsTemplateFile for the instance
func (instance *SettingsLocationTemplate) GetPropertyIsTemplateFile() (value bool, err error) {
	retValue, err := instance.GetProperty("IsTemplateFile")
	if err != nil {
		return
	}
	value, ok := retValue.(bool)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetSuiteParentId sets the value of SuiteParentId for the instance
func (instance *SettingsLocationTemplate) SetPropertySuiteParentId(value string) (err error) {
	return instance.SetProperty("SuiteParentId", value)
}

// GetSuiteParentId gets the value of SuiteParentId for the instance
func (instance *SettingsLocationTemplate) GetPropertySuiteParentId() (value string, err error) {
	retValue, err := instance.GetProperty("SuiteParentId")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetTemplateDescription sets the value of TemplateDescription for the instance
func (instance *SettingsLocationTemplate) SetPropertyTemplateDescription(value string) (err error) {
	return instance.SetProperty("TemplateDescription", value)
}

// GetTemplateDescription gets the value of TemplateDescription for the instance
func (instance *SettingsLocationTemplate) GetPropertyTemplateDescription() (value string, err error) {
	retValue, err := instance.GetProperty("TemplateDescription")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetTemplateId sets the value of TemplateId for the instance
func (instance *SettingsLocationTemplate) SetPropertyTemplateId(value string) (err error) {
	return instance.SetProperty("TemplateId", value)
}

// GetTemplateId gets the value of TemplateId for the instance
func (instance *SettingsLocationTemplate) GetPropertyTemplateId() (value string, err error) {
	retValue, err := instance.GetProperty("TemplateId")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetTemplateName sets the value of TemplateName for the instance
func (instance *SettingsLocationTemplate) SetPropertyTemplateName(value string) (err error) {
	return instance.SetProperty("TemplateName", value)
}

// GetTemplateName gets the value of TemplateName for the instance
func (instance *SettingsLocationTemplate) GetPropertyTemplateName() (value string, err error) {
	retValue, err := instance.GetProperty("TemplateName")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetTemplateProfile sets the value of TemplateProfile for the instance
func (instance *SettingsLocationTemplate) SetPropertyTemplateProfile(value string) (err error) {
	return instance.SetProperty("TemplateProfile", value)
}

// GetTemplateProfile gets the value of TemplateProfile for the instance
func (instance *SettingsLocationTemplate) GetPropertyTemplateProfile() (value string, err error) {
	retValue, err := instance.GetProperty("TemplateProfile")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetTemplateType sets the value of TemplateType for the instance
func (instance *SettingsLocationTemplate) SetPropertyTemplateType(value string) (err error) {
	return instance.SetProperty("TemplateType", value)
}

// GetTemplateType gets the value of TemplateType for the instance
func (instance *SettingsLocationTemplate) GetPropertyTemplateType() (value string, err error) {
	retValue, err := instance.GetProperty("TemplateType")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetTemplateVersion sets the value of TemplateVersion for the instance
func (instance *SettingsLocationTemplate) SetPropertyTemplateVersion(value uint32) (err error) {
	return instance.SetProperty("TemplateVersion", value)
}

// GetTemplateVersion gets the value of TemplateVersion for the instance
func (instance *SettingsLocationTemplate) GetPropertyTemplateVersion() (value uint32, err error) {
	retValue, err := instance.GetProperty("TemplateVersion")
	if err != nil {
		return
	}
	value, ok := retValue.(uint32)
	if !ok {
		// TODO: Set an error
	}
	return
}

// Get the content of the template file of the current instance.

// <param name="ReturnValue" type="string "></param>
func (instance *SettingsLocationTemplate) GetContent() (result string, err error) {
	retVal, err := instance.InvokeMethodWithReturn("GetContent")
	if err != nil {
		return
	}
	result = string(retVal)
	return

}

// Get the content of a setting location template file by template ID.

// <param name="TemplateId" type="string ">Unique ID of the settings location template</param>

// <param name="ReturnValue" type="string "></param>
func (instance *SettingsLocationTemplate) GetContentByTemplateId( /* IN */ TemplateId string) (result string, err error) {
	retVal, err := instance.InvokeMethodWithReturn("GetContentByTemplateId", TemplateId)
	if err != nil {
		return
	}
	result = string(retVal)
	return

}

// Register a settings location template.

// <param name="AbsolutePathToTemplate" type="string ">Absolute path to the settings location template file</param>
func (instance *SettingsLocationTemplate) Register( /* IN */ AbsolutePathToTemplate string) (err error) {
	_, err = instance.InvokeMethodWithReturn("Register", AbsolutePathToTemplate)
	if err != nil {
		return
	}
	return

}

// Unregister the current settings location template.
func (instance *SettingsLocationTemplate) Unregister() (err error) {
	_, err = instance.InvokeMethodWithReturn("Unregister")
	if err != nil {
		return
	}
	return

}

// Unregister all settings location templates.
func (instance *SettingsLocationTemplate) UnregisterAll() (err error) {
	_, err = instance.InvokeMethodWithReturn("UnregisterAll")
	if err != nil {
		return
	}
	return

}

// Unregister a settings location template by template ID.

// <param name="TemplateId" type="string ">Unique ID of the settings location template</param>
func (instance *SettingsLocationTemplate) UnregisterByTemplateId( /* IN */ TemplateId string) (err error) {
	_, err = instance.InvokeMethodWithReturn("UnregisterByTemplateId", TemplateId)
	if err != nil {
		return
	}
	return

}

// Validate a settings location template.

// <param name="AbsolutePathToTemplate" type="string ">Absolute path to the settings location template file</param>

// <param name="ReturnValue" type="string "></param>
func (instance *SettingsLocationTemplate) Validate( /* IN */ AbsolutePathToTemplate string) (result string, err error) {
	retVal, err := instance.InvokeMethodWithReturn("Validate", AbsolutePathToTemplate)
	if err != nil {
		return
	}
	result = string(retVal)
	return

}

// Update a settings location template by template file.

// <param name="AbsolutePathToTemplate" type="string ">Absolute path to the settings location template file</param>
func (instance *SettingsLocationTemplate) Update( /* IN */ AbsolutePathToTemplate string) (err error) {
	_, err = instance.InvokeMethodWithReturn("Update", AbsolutePathToTemplate)
	if err != nil {
		return
	}
	return

}

// Rebuild the template index from the existing registered template files.
func (instance *SettingsLocationTemplate) RebuildIndex() (err error) {
	_, err = instance.InvokeMethodWithReturn("RebuildIndex")
	if err != nil {
		return
	}
	return

}

// Enable the current settings location template.
func (instance *SettingsLocationTemplate) Enable() (err error) {
	_, err = instance.InvokeMethodWithReturn("Enable")
	if err != nil {
		return
	}
	return

}

// Enable a settings location template.

// <param name="TemplateId" type="string ">Unique ID of the settings location template</param>
func (instance *SettingsLocationTemplate) EnableByTemplateId( /* IN */ TemplateId string) (err error) {
	_, err = instance.InvokeMethodWithReturn("EnableByTemplateId", TemplateId)
	if err != nil {
		return
	}
	return

}

// Disable the current settings location template.
func (instance *SettingsLocationTemplate) Disable() (err error) {
	_, err = instance.InvokeMethodWithReturn("Disable")
	if err != nil {
		return
	}
	return

}

// Disable a settings location template.

// <param name="TemplateId" type="string ">Unique ID of the settings location template</param>
func (instance *SettingsLocationTemplate) DisableByTemplateId( /* IN */ TemplateId string) (err error) {
	_, err = instance.InvokeMethodWithReturn("DisableByTemplateId", TemplateId)
	if err != nil {
		return
	}
	return

}

// Get the information about the processes monitored by the current settings location template.

// <param name="ReturnValue" type="string "></param>
func (instance *SettingsLocationTemplate) GetProcessInfo() (result string, err error) {
	retVal, err := instance.InvokeMethodWithReturn("GetProcessInfo")
	if err != nil {
		return
	}
	result = string(retVal)
	return

}

// Get the information about the processes monitored by the given settings location template.

// <param name="TemplateId" type="string ">Unique ID of the settings location template</param>

// <param name="ReturnValue" type="string "></param>
func (instance *SettingsLocationTemplate) GetProcessInfoByTemplateId( /* IN */ TemplateId string) (result string, err error) {
	retVal, err := instance.InvokeMethodWithReturn("GetProcessInfoByTemplateId", TemplateId)
	if err != nil {
		return
	}
	result = string(retVal)
	return

}

// Get the schema used by the current settings location template.

// <param name="ReturnValue" type="string "></param>
func (instance *SettingsLocationTemplate) GetSchema() (result string, err error) {
	retVal, err := instance.InvokeMethodWithReturn("GetSchema")
	if err != nil {
		return
	}
	result = string(retVal)
	return

}

// Get the schema used by a settings location template.

// <param name="TemplateId" type="string ">Unique ID of the settings location template</param>

// <param name="ReturnValue" type="string "></param>
func (instance *SettingsLocationTemplate) GetSchemaByTemplateId( /* IN */ TemplateId string) (result string, err error) {
	retVal, err := instance.InvokeMethodWithReturn("GetSchemaByTemplateId", TemplateId)
	if err != nil {
		return
	}
	result = string(retVal)
	return

}

// Associate a template with a profile

// <param name="Profile" type="string ">The profile</param>
// <param name="TemplateId" type="string ">The template ID</param>
func (instance *SettingsLocationTemplate) SetTemplateProfileByTemplateId( /* IN */ TemplateId string,
	/* IN */ Profile string) (err error) {
	_, err = instance.InvokeMethodWithReturn("SetTemplateProfileByTemplateId", TemplateId, Profile)
	if err != nil {
		return
	}
	return

}
