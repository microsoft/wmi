// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

// 
// Author:
//      Auto Generated on 6/6/2024 using wmigen
//      Source root.WebAdministration
//////////////////////////////////////////////
package webadministration
import (
 "github.com/microsoft/wmi/pkg/base/query"
cim "github.com/microsoft/wmi/pkg/wmiinstance"
 "github.com/microsoft/wmi/pkg/errors"
 "reflect"
)

// PagesSection struct
type PagesSection struct { 
	*ConfigurationSectionWithCollection

	// 
	AsyncTimeout string

	// 
	AutoEventWireup bool

	// 
	Buffer bool

	// 
	ClientIDMode string

	// 
	CompilationMode int32

	// 
	ControlRenderingCompatibilityVersion string

	// 
	Controls TagPrefixInfo

	// 
	EnableEventValidation bool

	// 
	EnableSessionState int32

	// 
	EnableViewState bool

	// 
	EnableViewStateMac bool

	// 
	MaintainScrollPositionOnPostBack bool

	// 
	MasterPageFile string

	// 
	MaxPageStateFieldLength int32

	// 
	Namespaces NamespaceInfo

	// 
	PageBaseType string

	// 
	PageParserFilterType string

	// 
	RenderAllHiddenFieldsAtTopOfForm bool

	// 
	SmartNavigation bool

	// 
	StyleSheetTheme string

	// 
	TagMapping TagMapInfo

	// 
	Theme string

	// 
	UserControlBaseType string

	// 
	ValidateRequest bool

	// 
	ViewStateEncryptionMode int32
}

	func NewPagesSectionEx1(instance *cim.WmiInstance) (newInstance *PagesSection, err error) {tmp, err := NewConfigurationSectionWithCollectionEx1(instance)
		
	if err != nil { return }
	newInstance = &PagesSection {
	ConfigurationSectionWithCollection: tmp,
	}
	return
	}
	

	func NewPagesSectionEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery ) (newInstance *PagesSection, err error) {tmp, err := NewConfigurationSectionWithCollectionEx6(hostName, wmiNamespace, userName, password, domainName, query)
		
	if err != nil { return }
	newInstance = &PagesSection {
	ConfigurationSectionWithCollection: tmp,
	}
	return
	}
	

// SetAsyncTimeout sets the value of AsyncTimeout for the instance
func (instance *PagesSection) SetPropertyAsyncTimeout(value string) (err error) { 
    return instance.SetProperty("AsyncTimeout", (value))
}

// GetAsyncTimeout gets the value of AsyncTimeout for the instance
func (instance *PagesSection) GetPropertyAsyncTimeout()(value string, err error) { 
    retValue, err := instance.GetProperty("AsyncTimeout")
    if err != nil {
        return
    }
    if retValue == nil {
        // Doesn't have any value. Return empty
        return
    }
    
    valuetmp, ok := retValue.(string); 
    if !ok {
        err = errors.Wrapf(errors.InvalidType, " string is Invalid. Expected %s", reflect.TypeOf(retValue))
        return  
    }

    value = string(valuetmp)

    return
}

// SetAutoEventWireup sets the value of AutoEventWireup for the instance
func (instance *PagesSection) SetPropertyAutoEventWireup(value bool) (err error) { 
    return instance.SetProperty("AutoEventWireup", (value))
}

// GetAutoEventWireup gets the value of AutoEventWireup for the instance
func (instance *PagesSection) GetPropertyAutoEventWireup()(value bool, err error) { 
    retValue, err := instance.GetProperty("AutoEventWireup")
    if err != nil {
        return
    }
    if retValue == nil {
        // Doesn't have any value. Return empty
        return
    }
    
    valuetmp, ok := retValue.(bool); 
    if !ok {
        err = errors.Wrapf(errors.InvalidType, " bool is Invalid. Expected %s", reflect.TypeOf(retValue))
        return  
    }

    value = bool(valuetmp)

    return
}

// SetBuffer sets the value of Buffer for the instance
func (instance *PagesSection) SetPropertyBuffer(value bool) (err error) { 
    return instance.SetProperty("Buffer", (value))
}

// GetBuffer gets the value of Buffer for the instance
func (instance *PagesSection) GetPropertyBuffer()(value bool, err error) { 
    retValue, err := instance.GetProperty("Buffer")
    if err != nil {
        return
    }
    if retValue == nil {
        // Doesn't have any value. Return empty
        return
    }
    
    valuetmp, ok := retValue.(bool); 
    if !ok {
        err = errors.Wrapf(errors.InvalidType, " bool is Invalid. Expected %s", reflect.TypeOf(retValue))
        return  
    }

    value = bool(valuetmp)

    return
}

// SetClientIDMode sets the value of ClientIDMode for the instance
func (instance *PagesSection) SetPropertyClientIDMode(value string) (err error) { 
    return instance.SetProperty("ClientIDMode", (value))
}

// GetClientIDMode gets the value of ClientIDMode for the instance
func (instance *PagesSection) GetPropertyClientIDMode()(value string, err error) { 
    retValue, err := instance.GetProperty("ClientIDMode")
    if err != nil {
        return
    }
    if retValue == nil {
        // Doesn't have any value. Return empty
        return
    }
    
    valuetmp, ok := retValue.(string); 
    if !ok {
        err = errors.Wrapf(errors.InvalidType, " string is Invalid. Expected %s", reflect.TypeOf(retValue))
        return  
    }

    value = string(valuetmp)

    return
}

// SetCompilationMode sets the value of CompilationMode for the instance
func (instance *PagesSection) SetPropertyCompilationMode(value int32) (err error) { 
    return instance.SetProperty("CompilationMode", (value))
}

// GetCompilationMode gets the value of CompilationMode for the instance
func (instance *PagesSection) GetPropertyCompilationMode()(value int32, err error) { 
    retValue, err := instance.GetProperty("CompilationMode")
    if err != nil {
        return
    }
    if retValue == nil {
        // Doesn't have any value. Return empty
        return
    }
    
    valuetmp, ok := retValue.(int32); 
    if !ok {
        err = errors.Wrapf(errors.InvalidType, " int32 is Invalid. Expected %s", reflect.TypeOf(retValue))
        return  
    }

    value = int32(valuetmp)

    return
}

// SetControlRenderingCompatibilityVersion sets the value of ControlRenderingCompatibilityVersion for the instance
func (instance *PagesSection) SetPropertyControlRenderingCompatibilityVersion(value string) (err error) { 
    return instance.SetProperty("ControlRenderingCompatibilityVersion", (value))
}

// GetControlRenderingCompatibilityVersion gets the value of ControlRenderingCompatibilityVersion for the instance
func (instance *PagesSection) GetPropertyControlRenderingCompatibilityVersion()(value string, err error) { 
    retValue, err := instance.GetProperty("ControlRenderingCompatibilityVersion")
    if err != nil {
        return
    }
    if retValue == nil {
        // Doesn't have any value. Return empty
        return
    }
    
    valuetmp, ok := retValue.(string); 
    if !ok {
        err = errors.Wrapf(errors.InvalidType, " string is Invalid. Expected %s", reflect.TypeOf(retValue))
        return  
    }

    value = string(valuetmp)

    return
}

// SetControls sets the value of Controls for the instance
func (instance *PagesSection) SetPropertyControls(value TagPrefixInfo) (err error) { 
    return instance.SetProperty("Controls", (value))
}

// GetControls gets the value of Controls for the instance
func (instance *PagesSection) GetPropertyControls()(value TagPrefixInfo, err error) { 
    retValue, err := instance.GetProperty("Controls")
    if err != nil {
        return
    }
    if retValue == nil {
        // Doesn't have any value. Return empty
        return
    }
    
    valuetmp, ok := retValue.(TagPrefixInfo); 
    if !ok {
        err = errors.Wrapf(errors.InvalidType, " TagPrefixInfo is Invalid. Expected %s", reflect.TypeOf(retValue))
        return  
    }

    value = TagPrefixInfo(valuetmp)

    return
}

// SetEnableEventValidation sets the value of EnableEventValidation for the instance
func (instance *PagesSection) SetPropertyEnableEventValidation(value bool) (err error) { 
    return instance.SetProperty("EnableEventValidation", (value))
}

// GetEnableEventValidation gets the value of EnableEventValidation for the instance
func (instance *PagesSection) GetPropertyEnableEventValidation()(value bool, err error) { 
    retValue, err := instance.GetProperty("EnableEventValidation")
    if err != nil {
        return
    }
    if retValue == nil {
        // Doesn't have any value. Return empty
        return
    }
    
    valuetmp, ok := retValue.(bool); 
    if !ok {
        err = errors.Wrapf(errors.InvalidType, " bool is Invalid. Expected %s", reflect.TypeOf(retValue))
        return  
    }

    value = bool(valuetmp)

    return
}

// SetEnableSessionState sets the value of EnableSessionState for the instance
func (instance *PagesSection) SetPropertyEnableSessionState(value int32) (err error) { 
    return instance.SetProperty("EnableSessionState", (value))
}

// GetEnableSessionState gets the value of EnableSessionState for the instance
func (instance *PagesSection) GetPropertyEnableSessionState()(value int32, err error) { 
    retValue, err := instance.GetProperty("EnableSessionState")
    if err != nil {
        return
    }
    if retValue == nil {
        // Doesn't have any value. Return empty
        return
    }
    
    valuetmp, ok := retValue.(int32); 
    if !ok {
        err = errors.Wrapf(errors.InvalidType, " int32 is Invalid. Expected %s", reflect.TypeOf(retValue))
        return  
    }

    value = int32(valuetmp)

    return
}

// SetEnableViewState sets the value of EnableViewState for the instance
func (instance *PagesSection) SetPropertyEnableViewState(value bool) (err error) { 
    return instance.SetProperty("EnableViewState", (value))
}

// GetEnableViewState gets the value of EnableViewState for the instance
func (instance *PagesSection) GetPropertyEnableViewState()(value bool, err error) { 
    retValue, err := instance.GetProperty("EnableViewState")
    if err != nil {
        return
    }
    if retValue == nil {
        // Doesn't have any value. Return empty
        return
    }
    
    valuetmp, ok := retValue.(bool); 
    if !ok {
        err = errors.Wrapf(errors.InvalidType, " bool is Invalid. Expected %s", reflect.TypeOf(retValue))
        return  
    }

    value = bool(valuetmp)

    return
}

// SetEnableViewStateMac sets the value of EnableViewStateMac for the instance
func (instance *PagesSection) SetPropertyEnableViewStateMac(value bool) (err error) { 
    return instance.SetProperty("EnableViewStateMac", (value))
}

// GetEnableViewStateMac gets the value of EnableViewStateMac for the instance
func (instance *PagesSection) GetPropertyEnableViewStateMac()(value bool, err error) { 
    retValue, err := instance.GetProperty("EnableViewStateMac")
    if err != nil {
        return
    }
    if retValue == nil {
        // Doesn't have any value. Return empty
        return
    }
    
    valuetmp, ok := retValue.(bool); 
    if !ok {
        err = errors.Wrapf(errors.InvalidType, " bool is Invalid. Expected %s", reflect.TypeOf(retValue))
        return  
    }

    value = bool(valuetmp)

    return
}

// SetMaintainScrollPositionOnPostBack sets the value of MaintainScrollPositionOnPostBack for the instance
func (instance *PagesSection) SetPropertyMaintainScrollPositionOnPostBack(value bool) (err error) { 
    return instance.SetProperty("MaintainScrollPositionOnPostBack", (value))
}

// GetMaintainScrollPositionOnPostBack gets the value of MaintainScrollPositionOnPostBack for the instance
func (instance *PagesSection) GetPropertyMaintainScrollPositionOnPostBack()(value bool, err error) { 
    retValue, err := instance.GetProperty("MaintainScrollPositionOnPostBack")
    if err != nil {
        return
    }
    if retValue == nil {
        // Doesn't have any value. Return empty
        return
    }
    
    valuetmp, ok := retValue.(bool); 
    if !ok {
        err = errors.Wrapf(errors.InvalidType, " bool is Invalid. Expected %s", reflect.TypeOf(retValue))
        return  
    }

    value = bool(valuetmp)

    return
}

// SetMasterPageFile sets the value of MasterPageFile for the instance
func (instance *PagesSection) SetPropertyMasterPageFile(value string) (err error) { 
    return instance.SetProperty("MasterPageFile", (value))
}

// GetMasterPageFile gets the value of MasterPageFile for the instance
func (instance *PagesSection) GetPropertyMasterPageFile()(value string, err error) { 
    retValue, err := instance.GetProperty("MasterPageFile")
    if err != nil {
        return
    }
    if retValue == nil {
        // Doesn't have any value. Return empty
        return
    }
    
    valuetmp, ok := retValue.(string); 
    if !ok {
        err = errors.Wrapf(errors.InvalidType, " string is Invalid. Expected %s", reflect.TypeOf(retValue))
        return  
    }

    value = string(valuetmp)

    return
}

// SetMaxPageStateFieldLength sets the value of MaxPageStateFieldLength for the instance
func (instance *PagesSection) SetPropertyMaxPageStateFieldLength(value int32) (err error) { 
    return instance.SetProperty("MaxPageStateFieldLength", (value))
}

// GetMaxPageStateFieldLength gets the value of MaxPageStateFieldLength for the instance
func (instance *PagesSection) GetPropertyMaxPageStateFieldLength()(value int32, err error) { 
    retValue, err := instance.GetProperty("MaxPageStateFieldLength")
    if err != nil {
        return
    }
    if retValue == nil {
        // Doesn't have any value. Return empty
        return
    }
    
    valuetmp, ok := retValue.(int32); 
    if !ok {
        err = errors.Wrapf(errors.InvalidType, " int32 is Invalid. Expected %s", reflect.TypeOf(retValue))
        return  
    }

    value = int32(valuetmp)

    return
}

// SetNamespaces sets the value of Namespaces for the instance
func (instance *PagesSection) SetPropertyNamespaces(value NamespaceInfo) (err error) { 
    return instance.SetProperty("Namespaces", (value))
}

// GetNamespaces gets the value of Namespaces for the instance
func (instance *PagesSection) GetPropertyNamespaces()(value NamespaceInfo, err error) { 
    retValue, err := instance.GetProperty("Namespaces")
    if err != nil {
        return
    }
    if retValue == nil {
        // Doesn't have any value. Return empty
        return
    }
    
    valuetmp, ok := retValue.(NamespaceInfo); 
    if !ok {
        err = errors.Wrapf(errors.InvalidType, " NamespaceInfo is Invalid. Expected %s", reflect.TypeOf(retValue))
        return  
    }

    value = NamespaceInfo(valuetmp)

    return
}

// SetPageBaseType sets the value of PageBaseType for the instance
func (instance *PagesSection) SetPropertyPageBaseType(value string) (err error) { 
    return instance.SetProperty("PageBaseType", (value))
}

// GetPageBaseType gets the value of PageBaseType for the instance
func (instance *PagesSection) GetPropertyPageBaseType()(value string, err error) { 
    retValue, err := instance.GetProperty("PageBaseType")
    if err != nil {
        return
    }
    if retValue == nil {
        // Doesn't have any value. Return empty
        return
    }
    
    valuetmp, ok := retValue.(string); 
    if !ok {
        err = errors.Wrapf(errors.InvalidType, " string is Invalid. Expected %s", reflect.TypeOf(retValue))
        return  
    }

    value = string(valuetmp)

    return
}

// SetPageParserFilterType sets the value of PageParserFilterType for the instance
func (instance *PagesSection) SetPropertyPageParserFilterType(value string) (err error) { 
    return instance.SetProperty("PageParserFilterType", (value))
}

// GetPageParserFilterType gets the value of PageParserFilterType for the instance
func (instance *PagesSection) GetPropertyPageParserFilterType()(value string, err error) { 
    retValue, err := instance.GetProperty("PageParserFilterType")
    if err != nil {
        return
    }
    if retValue == nil {
        // Doesn't have any value. Return empty
        return
    }
    
    valuetmp, ok := retValue.(string); 
    if !ok {
        err = errors.Wrapf(errors.InvalidType, " string is Invalid. Expected %s", reflect.TypeOf(retValue))
        return  
    }

    value = string(valuetmp)

    return
}

// SetRenderAllHiddenFieldsAtTopOfForm sets the value of RenderAllHiddenFieldsAtTopOfForm for the instance
func (instance *PagesSection) SetPropertyRenderAllHiddenFieldsAtTopOfForm(value bool) (err error) { 
    return instance.SetProperty("RenderAllHiddenFieldsAtTopOfForm", (value))
}

// GetRenderAllHiddenFieldsAtTopOfForm gets the value of RenderAllHiddenFieldsAtTopOfForm for the instance
func (instance *PagesSection) GetPropertyRenderAllHiddenFieldsAtTopOfForm()(value bool, err error) { 
    retValue, err := instance.GetProperty("RenderAllHiddenFieldsAtTopOfForm")
    if err != nil {
        return
    }
    if retValue == nil {
        // Doesn't have any value. Return empty
        return
    }
    
    valuetmp, ok := retValue.(bool); 
    if !ok {
        err = errors.Wrapf(errors.InvalidType, " bool is Invalid. Expected %s", reflect.TypeOf(retValue))
        return  
    }

    value = bool(valuetmp)

    return
}

// SetSmartNavigation sets the value of SmartNavigation for the instance
func (instance *PagesSection) SetPropertySmartNavigation(value bool) (err error) { 
    return instance.SetProperty("SmartNavigation", (value))
}

// GetSmartNavigation gets the value of SmartNavigation for the instance
func (instance *PagesSection) GetPropertySmartNavigation()(value bool, err error) { 
    retValue, err := instance.GetProperty("SmartNavigation")
    if err != nil {
        return
    }
    if retValue == nil {
        // Doesn't have any value. Return empty
        return
    }
    
    valuetmp, ok := retValue.(bool); 
    if !ok {
        err = errors.Wrapf(errors.InvalidType, " bool is Invalid. Expected %s", reflect.TypeOf(retValue))
        return  
    }

    value = bool(valuetmp)

    return
}

// SetStyleSheetTheme sets the value of StyleSheetTheme for the instance
func (instance *PagesSection) SetPropertyStyleSheetTheme(value string) (err error) { 
    return instance.SetProperty("StyleSheetTheme", (value))
}

// GetStyleSheetTheme gets the value of StyleSheetTheme for the instance
func (instance *PagesSection) GetPropertyStyleSheetTheme()(value string, err error) { 
    retValue, err := instance.GetProperty("StyleSheetTheme")
    if err != nil {
        return
    }
    if retValue == nil {
        // Doesn't have any value. Return empty
        return
    }
    
    valuetmp, ok := retValue.(string); 
    if !ok {
        err = errors.Wrapf(errors.InvalidType, " string is Invalid. Expected %s", reflect.TypeOf(retValue))
        return  
    }

    value = string(valuetmp)

    return
}

// SetTagMapping sets the value of TagMapping for the instance
func (instance *PagesSection) SetPropertyTagMapping(value TagMapInfo) (err error) { 
    return instance.SetProperty("TagMapping", (value))
}

// GetTagMapping gets the value of TagMapping for the instance
func (instance *PagesSection) GetPropertyTagMapping()(value TagMapInfo, err error) { 
    retValue, err := instance.GetProperty("TagMapping")
    if err != nil {
        return
    }
    if retValue == nil {
        // Doesn't have any value. Return empty
        return
    }
    
    valuetmp, ok := retValue.(TagMapInfo); 
    if !ok {
        err = errors.Wrapf(errors.InvalidType, " TagMapInfo is Invalid. Expected %s", reflect.TypeOf(retValue))
        return  
    }

    value = TagMapInfo(valuetmp)

    return
}

// SetTheme sets the value of Theme for the instance
func (instance *PagesSection) SetPropertyTheme(value string) (err error) { 
    return instance.SetProperty("Theme", (value))
}

// GetTheme gets the value of Theme for the instance
func (instance *PagesSection) GetPropertyTheme()(value string, err error) { 
    retValue, err := instance.GetProperty("Theme")
    if err != nil {
        return
    }
    if retValue == nil {
        // Doesn't have any value. Return empty
        return
    }
    
    valuetmp, ok := retValue.(string); 
    if !ok {
        err = errors.Wrapf(errors.InvalidType, " string is Invalid. Expected %s", reflect.TypeOf(retValue))
        return  
    }

    value = string(valuetmp)

    return
}

// SetUserControlBaseType sets the value of UserControlBaseType for the instance
func (instance *PagesSection) SetPropertyUserControlBaseType(value string) (err error) { 
    return instance.SetProperty("UserControlBaseType", (value))
}

// GetUserControlBaseType gets the value of UserControlBaseType for the instance
func (instance *PagesSection) GetPropertyUserControlBaseType()(value string, err error) { 
    retValue, err := instance.GetProperty("UserControlBaseType")
    if err != nil {
        return
    }
    if retValue == nil {
        // Doesn't have any value. Return empty
        return
    }
    
    valuetmp, ok := retValue.(string); 
    if !ok {
        err = errors.Wrapf(errors.InvalidType, " string is Invalid. Expected %s", reflect.TypeOf(retValue))
        return  
    }

    value = string(valuetmp)

    return
}

// SetValidateRequest sets the value of ValidateRequest for the instance
func (instance *PagesSection) SetPropertyValidateRequest(value bool) (err error) { 
    return instance.SetProperty("ValidateRequest", (value))
}

// GetValidateRequest gets the value of ValidateRequest for the instance
func (instance *PagesSection) GetPropertyValidateRequest()(value bool, err error) { 
    retValue, err := instance.GetProperty("ValidateRequest")
    if err != nil {
        return
    }
    if retValue == nil {
        // Doesn't have any value. Return empty
        return
    }
    
    valuetmp, ok := retValue.(bool); 
    if !ok {
        err = errors.Wrapf(errors.InvalidType, " bool is Invalid. Expected %s", reflect.TypeOf(retValue))
        return  
    }

    value = bool(valuetmp)

    return
}

// SetViewStateEncryptionMode sets the value of ViewStateEncryptionMode for the instance
func (instance *PagesSection) SetPropertyViewStateEncryptionMode(value int32) (err error) { 
    return instance.SetProperty("ViewStateEncryptionMode", (value))
}

// GetViewStateEncryptionMode gets the value of ViewStateEncryptionMode for the instance
func (instance *PagesSection) GetPropertyViewStateEncryptionMode()(value int32, err error) { 
    retValue, err := instance.GetProperty("ViewStateEncryptionMode")
    if err != nil {
        return
    }
    if retValue == nil {
        // Doesn't have any value. Return empty
        return
    }
    
    valuetmp, ok := retValue.(int32); 
    if !ok {
        err = errors.Wrapf(errors.InvalidType, " int32 is Invalid. Expected %s", reflect.TypeOf(retValue))
        return  
    }

    value = int32(valuetmp)

    return
}

