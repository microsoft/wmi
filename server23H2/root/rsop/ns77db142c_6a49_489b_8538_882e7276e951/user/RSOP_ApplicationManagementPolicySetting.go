// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 6/12/2024 using wmigen
//      Source root.RSOP.NS77DB142C_6A49_489B_8538_882E7276E951.User
//////////////////////////////////////////////
package user

import (
	"github.com/microsoft/wmi/pkg/base/query"
	"github.com/microsoft/wmi/pkg/errors"
	cim "github.com/microsoft/wmi/pkg/wmiinstance"
	"reflect"
)

// RSOP_ApplicationManagementPolicySetting struct
type RSOP_ApplicationManagementPolicySetting struct {
	*RSOP_PolicySetting

	//
	AllowX86OnIA64 bool

	//
	ApplicationId string

	//
	ApplyCause uint32

	//
	AssignmentType uint32

	//
	Categories []string

	//
	DemandInstallable bool

	//
	DeploymentLastModifyTime string

	//
	DeploymentType uint32

	//
	DisplayInARP bool

	//
	Eligibility uint32

	//
	EntryType uint32

	//
	IgnoreLanguage bool

	//
	InstallationUI uint32

	//
	LanguageId uint32

	//
	LanguageMatch uint32

	//
	LossOfScopeAction uint32

	//
	MachineArchitectures []uint32

	//
	OnDemandClsid string

	//
	OnDemandFileExtension string

	//
	OnDemandProgId string

	//
	PackageLocation string

	//
	PackageType uint32

	//
	PrecedenceReason uint32

	//
	ProductId string

	//
	Publisher string

	//
	RedeployCount uint32

	//
	RemovalCause uint32

	//
	RemovalType uint32

	//
	RemovingApplication string

	//
	ReplaceableApplications []string

	//
	ScriptFile string

	//
	SecurityDescriptor []uint8

	//
	SupportURL string

	//
	Transforms []string

	//
	UninstallUnmanaged bool

	//
	UpgradeableApplications []string

	//
	UpgradeSettingsMandatory bool

	//
	VersionNumberHi uint32

	//
	VersionNumberLo uint32
}

func NewRSOP_ApplicationManagementPolicySettingEx1(instance *cim.WmiInstance) (newInstance *RSOP_ApplicationManagementPolicySetting, err error) {
	tmp, err := NewRSOP_PolicySettingEx1(instance)

	if err != nil {
		return
	}
	newInstance = &RSOP_ApplicationManagementPolicySetting{
		RSOP_PolicySetting: tmp,
	}
	return
}

func NewRSOP_ApplicationManagementPolicySettingEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *RSOP_ApplicationManagementPolicySetting, err error) {
	tmp, err := NewRSOP_PolicySettingEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &RSOP_ApplicationManagementPolicySetting{
		RSOP_PolicySetting: tmp,
	}
	return
}

// SetAllowX86OnIA64 sets the value of AllowX86OnIA64 for the instance
func (instance *RSOP_ApplicationManagementPolicySetting) SetPropertyAllowX86OnIA64(value bool) (err error) {
	return instance.SetProperty("AllowX86OnIA64", (value))
}

// GetAllowX86OnIA64 gets the value of AllowX86OnIA64 for the instance
func (instance *RSOP_ApplicationManagementPolicySetting) GetPropertyAllowX86OnIA64() (value bool, err error) {
	retValue, err := instance.GetProperty("AllowX86OnIA64")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(bool)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " bool is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = bool(valuetmp)

	return
}

// SetApplicationId sets the value of ApplicationId for the instance
func (instance *RSOP_ApplicationManagementPolicySetting) SetPropertyApplicationId(value string) (err error) {
	return instance.SetProperty("ApplicationId", (value))
}

// GetApplicationId gets the value of ApplicationId for the instance
func (instance *RSOP_ApplicationManagementPolicySetting) GetPropertyApplicationId() (value string, err error) {
	retValue, err := instance.GetProperty("ApplicationId")
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

// SetApplyCause sets the value of ApplyCause for the instance
func (instance *RSOP_ApplicationManagementPolicySetting) SetPropertyApplyCause(value uint32) (err error) {
	return instance.SetProperty("ApplyCause", (value))
}

// GetApplyCause gets the value of ApplyCause for the instance
func (instance *RSOP_ApplicationManagementPolicySetting) GetPropertyApplyCause() (value uint32, err error) {
	retValue, err := instance.GetProperty("ApplyCause")
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

// SetAssignmentType sets the value of AssignmentType for the instance
func (instance *RSOP_ApplicationManagementPolicySetting) SetPropertyAssignmentType(value uint32) (err error) {
	return instance.SetProperty("AssignmentType", (value))
}

// GetAssignmentType gets the value of AssignmentType for the instance
func (instance *RSOP_ApplicationManagementPolicySetting) GetPropertyAssignmentType() (value uint32, err error) {
	retValue, err := instance.GetProperty("AssignmentType")
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

// SetCategories sets the value of Categories for the instance
func (instance *RSOP_ApplicationManagementPolicySetting) SetPropertyCategories(value []string) (err error) {
	return instance.SetProperty("Categories", (value))
}

// GetCategories gets the value of Categories for the instance
func (instance *RSOP_ApplicationManagementPolicySetting) GetPropertyCategories() (value []string, err error) {
	retValue, err := instance.GetProperty("Categories")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	for _, interfaceValue := range retValue.([]interface{}) {
		valuetmp, ok := interfaceValue.(string)
		if !ok {
			err = errors.Wrapf(errors.InvalidType, " string is Invalid. Expected %s", reflect.TypeOf(interfaceValue))
			return
		}
		value = append(value, string(valuetmp))
	}

	return
}

// SetDemandInstallable sets the value of DemandInstallable for the instance
func (instance *RSOP_ApplicationManagementPolicySetting) SetPropertyDemandInstallable(value bool) (err error) {
	return instance.SetProperty("DemandInstallable", (value))
}

// GetDemandInstallable gets the value of DemandInstallable for the instance
func (instance *RSOP_ApplicationManagementPolicySetting) GetPropertyDemandInstallable() (value bool, err error) {
	retValue, err := instance.GetProperty("DemandInstallable")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(bool)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " bool is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = bool(valuetmp)

	return
}

// SetDeploymentLastModifyTime sets the value of DeploymentLastModifyTime for the instance
func (instance *RSOP_ApplicationManagementPolicySetting) SetPropertyDeploymentLastModifyTime(value string) (err error) {
	return instance.SetProperty("DeploymentLastModifyTime", (value))
}

// GetDeploymentLastModifyTime gets the value of DeploymentLastModifyTime for the instance
func (instance *RSOP_ApplicationManagementPolicySetting) GetPropertyDeploymentLastModifyTime() (value string, err error) {
	retValue, err := instance.GetProperty("DeploymentLastModifyTime")
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

// SetDeploymentType sets the value of DeploymentType for the instance
func (instance *RSOP_ApplicationManagementPolicySetting) SetPropertyDeploymentType(value uint32) (err error) {
	return instance.SetProperty("DeploymentType", (value))
}

// GetDeploymentType gets the value of DeploymentType for the instance
func (instance *RSOP_ApplicationManagementPolicySetting) GetPropertyDeploymentType() (value uint32, err error) {
	retValue, err := instance.GetProperty("DeploymentType")
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

// SetDisplayInARP sets the value of DisplayInARP for the instance
func (instance *RSOP_ApplicationManagementPolicySetting) SetPropertyDisplayInARP(value bool) (err error) {
	return instance.SetProperty("DisplayInARP", (value))
}

// GetDisplayInARP gets the value of DisplayInARP for the instance
func (instance *RSOP_ApplicationManagementPolicySetting) GetPropertyDisplayInARP() (value bool, err error) {
	retValue, err := instance.GetProperty("DisplayInARP")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(bool)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " bool is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = bool(valuetmp)

	return
}

// SetEligibility sets the value of Eligibility for the instance
func (instance *RSOP_ApplicationManagementPolicySetting) SetPropertyEligibility(value uint32) (err error) {
	return instance.SetProperty("Eligibility", (value))
}

// GetEligibility gets the value of Eligibility for the instance
func (instance *RSOP_ApplicationManagementPolicySetting) GetPropertyEligibility() (value uint32, err error) {
	retValue, err := instance.GetProperty("Eligibility")
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

// SetEntryType sets the value of EntryType for the instance
func (instance *RSOP_ApplicationManagementPolicySetting) SetPropertyEntryType(value uint32) (err error) {
	return instance.SetProperty("EntryType", (value))
}

// GetEntryType gets the value of EntryType for the instance
func (instance *RSOP_ApplicationManagementPolicySetting) GetPropertyEntryType() (value uint32, err error) {
	retValue, err := instance.GetProperty("EntryType")
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

// SetIgnoreLanguage sets the value of IgnoreLanguage for the instance
func (instance *RSOP_ApplicationManagementPolicySetting) SetPropertyIgnoreLanguage(value bool) (err error) {
	return instance.SetProperty("IgnoreLanguage", (value))
}

// GetIgnoreLanguage gets the value of IgnoreLanguage for the instance
func (instance *RSOP_ApplicationManagementPolicySetting) GetPropertyIgnoreLanguage() (value bool, err error) {
	retValue, err := instance.GetProperty("IgnoreLanguage")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(bool)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " bool is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = bool(valuetmp)

	return
}

// SetInstallationUI sets the value of InstallationUI for the instance
func (instance *RSOP_ApplicationManagementPolicySetting) SetPropertyInstallationUI(value uint32) (err error) {
	return instance.SetProperty("InstallationUI", (value))
}

// GetInstallationUI gets the value of InstallationUI for the instance
func (instance *RSOP_ApplicationManagementPolicySetting) GetPropertyInstallationUI() (value uint32, err error) {
	retValue, err := instance.GetProperty("InstallationUI")
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

// SetLanguageId sets the value of LanguageId for the instance
func (instance *RSOP_ApplicationManagementPolicySetting) SetPropertyLanguageId(value uint32) (err error) {
	return instance.SetProperty("LanguageId", (value))
}

// GetLanguageId gets the value of LanguageId for the instance
func (instance *RSOP_ApplicationManagementPolicySetting) GetPropertyLanguageId() (value uint32, err error) {
	retValue, err := instance.GetProperty("LanguageId")
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

// SetLanguageMatch sets the value of LanguageMatch for the instance
func (instance *RSOP_ApplicationManagementPolicySetting) SetPropertyLanguageMatch(value uint32) (err error) {
	return instance.SetProperty("LanguageMatch", (value))
}

// GetLanguageMatch gets the value of LanguageMatch for the instance
func (instance *RSOP_ApplicationManagementPolicySetting) GetPropertyLanguageMatch() (value uint32, err error) {
	retValue, err := instance.GetProperty("LanguageMatch")
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

// SetLossOfScopeAction sets the value of LossOfScopeAction for the instance
func (instance *RSOP_ApplicationManagementPolicySetting) SetPropertyLossOfScopeAction(value uint32) (err error) {
	return instance.SetProperty("LossOfScopeAction", (value))
}

// GetLossOfScopeAction gets the value of LossOfScopeAction for the instance
func (instance *RSOP_ApplicationManagementPolicySetting) GetPropertyLossOfScopeAction() (value uint32, err error) {
	retValue, err := instance.GetProperty("LossOfScopeAction")
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

// SetMachineArchitectures sets the value of MachineArchitectures for the instance
func (instance *RSOP_ApplicationManagementPolicySetting) SetPropertyMachineArchitectures(value []uint32) (err error) {
	return instance.SetProperty("MachineArchitectures", (value))
}

// GetMachineArchitectures gets the value of MachineArchitectures for the instance
func (instance *RSOP_ApplicationManagementPolicySetting) GetPropertyMachineArchitectures() (value []uint32, err error) {
	retValue, err := instance.GetProperty("MachineArchitectures")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	for _, interfaceValue := range retValue.([]interface{}) {
		valuetmp, ok := interfaceValue.(uint32)
		if !ok {
			err = errors.Wrapf(errors.InvalidType, " uint32 is Invalid. Expected %s", reflect.TypeOf(interfaceValue))
			return
		}
		value = append(value, uint32(valuetmp))
	}

	return
}

// SetOnDemandClsid sets the value of OnDemandClsid for the instance
func (instance *RSOP_ApplicationManagementPolicySetting) SetPropertyOnDemandClsid(value string) (err error) {
	return instance.SetProperty("OnDemandClsid", (value))
}

// GetOnDemandClsid gets the value of OnDemandClsid for the instance
func (instance *RSOP_ApplicationManagementPolicySetting) GetPropertyOnDemandClsid() (value string, err error) {
	retValue, err := instance.GetProperty("OnDemandClsid")
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

// SetOnDemandFileExtension sets the value of OnDemandFileExtension for the instance
func (instance *RSOP_ApplicationManagementPolicySetting) SetPropertyOnDemandFileExtension(value string) (err error) {
	return instance.SetProperty("OnDemandFileExtension", (value))
}

// GetOnDemandFileExtension gets the value of OnDemandFileExtension for the instance
func (instance *RSOP_ApplicationManagementPolicySetting) GetPropertyOnDemandFileExtension() (value string, err error) {
	retValue, err := instance.GetProperty("OnDemandFileExtension")
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

// SetOnDemandProgId sets the value of OnDemandProgId for the instance
func (instance *RSOP_ApplicationManagementPolicySetting) SetPropertyOnDemandProgId(value string) (err error) {
	return instance.SetProperty("OnDemandProgId", (value))
}

// GetOnDemandProgId gets the value of OnDemandProgId for the instance
func (instance *RSOP_ApplicationManagementPolicySetting) GetPropertyOnDemandProgId() (value string, err error) {
	retValue, err := instance.GetProperty("OnDemandProgId")
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

// SetPackageLocation sets the value of PackageLocation for the instance
func (instance *RSOP_ApplicationManagementPolicySetting) SetPropertyPackageLocation(value string) (err error) {
	return instance.SetProperty("PackageLocation", (value))
}

// GetPackageLocation gets the value of PackageLocation for the instance
func (instance *RSOP_ApplicationManagementPolicySetting) GetPropertyPackageLocation() (value string, err error) {
	retValue, err := instance.GetProperty("PackageLocation")
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

// SetPackageType sets the value of PackageType for the instance
func (instance *RSOP_ApplicationManagementPolicySetting) SetPropertyPackageType(value uint32) (err error) {
	return instance.SetProperty("PackageType", (value))
}

// GetPackageType gets the value of PackageType for the instance
func (instance *RSOP_ApplicationManagementPolicySetting) GetPropertyPackageType() (value uint32, err error) {
	retValue, err := instance.GetProperty("PackageType")
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

// SetPrecedenceReason sets the value of PrecedenceReason for the instance
func (instance *RSOP_ApplicationManagementPolicySetting) SetPropertyPrecedenceReason(value uint32) (err error) {
	return instance.SetProperty("PrecedenceReason", (value))
}

// GetPrecedenceReason gets the value of PrecedenceReason for the instance
func (instance *RSOP_ApplicationManagementPolicySetting) GetPropertyPrecedenceReason() (value uint32, err error) {
	retValue, err := instance.GetProperty("PrecedenceReason")
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

// SetProductId sets the value of ProductId for the instance
func (instance *RSOP_ApplicationManagementPolicySetting) SetPropertyProductId(value string) (err error) {
	return instance.SetProperty("ProductId", (value))
}

// GetProductId gets the value of ProductId for the instance
func (instance *RSOP_ApplicationManagementPolicySetting) GetPropertyProductId() (value string, err error) {
	retValue, err := instance.GetProperty("ProductId")
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

// SetPublisher sets the value of Publisher for the instance
func (instance *RSOP_ApplicationManagementPolicySetting) SetPropertyPublisher(value string) (err error) {
	return instance.SetProperty("Publisher", (value))
}

// GetPublisher gets the value of Publisher for the instance
func (instance *RSOP_ApplicationManagementPolicySetting) GetPropertyPublisher() (value string, err error) {
	retValue, err := instance.GetProperty("Publisher")
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

// SetRedeployCount sets the value of RedeployCount for the instance
func (instance *RSOP_ApplicationManagementPolicySetting) SetPropertyRedeployCount(value uint32) (err error) {
	return instance.SetProperty("RedeployCount", (value))
}

// GetRedeployCount gets the value of RedeployCount for the instance
func (instance *RSOP_ApplicationManagementPolicySetting) GetPropertyRedeployCount() (value uint32, err error) {
	retValue, err := instance.GetProperty("RedeployCount")
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

// SetRemovalCause sets the value of RemovalCause for the instance
func (instance *RSOP_ApplicationManagementPolicySetting) SetPropertyRemovalCause(value uint32) (err error) {
	return instance.SetProperty("RemovalCause", (value))
}

// GetRemovalCause gets the value of RemovalCause for the instance
func (instance *RSOP_ApplicationManagementPolicySetting) GetPropertyRemovalCause() (value uint32, err error) {
	retValue, err := instance.GetProperty("RemovalCause")
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

// SetRemovalType sets the value of RemovalType for the instance
func (instance *RSOP_ApplicationManagementPolicySetting) SetPropertyRemovalType(value uint32) (err error) {
	return instance.SetProperty("RemovalType", (value))
}

// GetRemovalType gets the value of RemovalType for the instance
func (instance *RSOP_ApplicationManagementPolicySetting) GetPropertyRemovalType() (value uint32, err error) {
	retValue, err := instance.GetProperty("RemovalType")
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

// SetRemovingApplication sets the value of RemovingApplication for the instance
func (instance *RSOP_ApplicationManagementPolicySetting) SetPropertyRemovingApplication(value string) (err error) {
	return instance.SetProperty("RemovingApplication", (value))
}

// GetRemovingApplication gets the value of RemovingApplication for the instance
func (instance *RSOP_ApplicationManagementPolicySetting) GetPropertyRemovingApplication() (value string, err error) {
	retValue, err := instance.GetProperty("RemovingApplication")
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

// SetReplaceableApplications sets the value of ReplaceableApplications for the instance
func (instance *RSOP_ApplicationManagementPolicySetting) SetPropertyReplaceableApplications(value []string) (err error) {
	return instance.SetProperty("ReplaceableApplications", (value))
}

// GetReplaceableApplications gets the value of ReplaceableApplications for the instance
func (instance *RSOP_ApplicationManagementPolicySetting) GetPropertyReplaceableApplications() (value []string, err error) {
	retValue, err := instance.GetProperty("ReplaceableApplications")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	for _, interfaceValue := range retValue.([]interface{}) {
		valuetmp, ok := interfaceValue.(string)
		if !ok {
			err = errors.Wrapf(errors.InvalidType, " string is Invalid. Expected %s", reflect.TypeOf(interfaceValue))
			return
		}
		value = append(value, string(valuetmp))
	}

	return
}

// SetScriptFile sets the value of ScriptFile for the instance
func (instance *RSOP_ApplicationManagementPolicySetting) SetPropertyScriptFile(value string) (err error) {
	return instance.SetProperty("ScriptFile", (value))
}

// GetScriptFile gets the value of ScriptFile for the instance
func (instance *RSOP_ApplicationManagementPolicySetting) GetPropertyScriptFile() (value string, err error) {
	retValue, err := instance.GetProperty("ScriptFile")
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

// SetSecurityDescriptor sets the value of SecurityDescriptor for the instance
func (instance *RSOP_ApplicationManagementPolicySetting) SetPropertySecurityDescriptor(value []uint8) (err error) {
	return instance.SetProperty("SecurityDescriptor", (value))
}

// GetSecurityDescriptor gets the value of SecurityDescriptor for the instance
func (instance *RSOP_ApplicationManagementPolicySetting) GetPropertySecurityDescriptor() (value []uint8, err error) {
	retValue, err := instance.GetProperty("SecurityDescriptor")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	for _, interfaceValue := range retValue.([]interface{}) {
		valuetmp, ok := interfaceValue.(uint8)
		if !ok {
			err = errors.Wrapf(errors.InvalidType, " uint8 is Invalid. Expected %s", reflect.TypeOf(interfaceValue))
			return
		}
		value = append(value, uint8(valuetmp))
	}

	return
}

// SetSupportURL sets the value of SupportURL for the instance
func (instance *RSOP_ApplicationManagementPolicySetting) SetPropertySupportURL(value string) (err error) {
	return instance.SetProperty("SupportURL", (value))
}

// GetSupportURL gets the value of SupportURL for the instance
func (instance *RSOP_ApplicationManagementPolicySetting) GetPropertySupportURL() (value string, err error) {
	retValue, err := instance.GetProperty("SupportURL")
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

// SetTransforms sets the value of Transforms for the instance
func (instance *RSOP_ApplicationManagementPolicySetting) SetPropertyTransforms(value []string) (err error) {
	return instance.SetProperty("Transforms", (value))
}

// GetTransforms gets the value of Transforms for the instance
func (instance *RSOP_ApplicationManagementPolicySetting) GetPropertyTransforms() (value []string, err error) {
	retValue, err := instance.GetProperty("Transforms")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	for _, interfaceValue := range retValue.([]interface{}) {
		valuetmp, ok := interfaceValue.(string)
		if !ok {
			err = errors.Wrapf(errors.InvalidType, " string is Invalid. Expected %s", reflect.TypeOf(interfaceValue))
			return
		}
		value = append(value, string(valuetmp))
	}

	return
}

// SetUninstallUnmanaged sets the value of UninstallUnmanaged for the instance
func (instance *RSOP_ApplicationManagementPolicySetting) SetPropertyUninstallUnmanaged(value bool) (err error) {
	return instance.SetProperty("UninstallUnmanaged", (value))
}

// GetUninstallUnmanaged gets the value of UninstallUnmanaged for the instance
func (instance *RSOP_ApplicationManagementPolicySetting) GetPropertyUninstallUnmanaged() (value bool, err error) {
	retValue, err := instance.GetProperty("UninstallUnmanaged")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(bool)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " bool is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = bool(valuetmp)

	return
}

// SetUpgradeableApplications sets the value of UpgradeableApplications for the instance
func (instance *RSOP_ApplicationManagementPolicySetting) SetPropertyUpgradeableApplications(value []string) (err error) {
	return instance.SetProperty("UpgradeableApplications", (value))
}

// GetUpgradeableApplications gets the value of UpgradeableApplications for the instance
func (instance *RSOP_ApplicationManagementPolicySetting) GetPropertyUpgradeableApplications() (value []string, err error) {
	retValue, err := instance.GetProperty("UpgradeableApplications")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	for _, interfaceValue := range retValue.([]interface{}) {
		valuetmp, ok := interfaceValue.(string)
		if !ok {
			err = errors.Wrapf(errors.InvalidType, " string is Invalid. Expected %s", reflect.TypeOf(interfaceValue))
			return
		}
		value = append(value, string(valuetmp))
	}

	return
}

// SetUpgradeSettingsMandatory sets the value of UpgradeSettingsMandatory for the instance
func (instance *RSOP_ApplicationManagementPolicySetting) SetPropertyUpgradeSettingsMandatory(value bool) (err error) {
	return instance.SetProperty("UpgradeSettingsMandatory", (value))
}

// GetUpgradeSettingsMandatory gets the value of UpgradeSettingsMandatory for the instance
func (instance *RSOP_ApplicationManagementPolicySetting) GetPropertyUpgradeSettingsMandatory() (value bool, err error) {
	retValue, err := instance.GetProperty("UpgradeSettingsMandatory")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(bool)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " bool is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = bool(valuetmp)

	return
}

// SetVersionNumberHi sets the value of VersionNumberHi for the instance
func (instance *RSOP_ApplicationManagementPolicySetting) SetPropertyVersionNumberHi(value uint32) (err error) {
	return instance.SetProperty("VersionNumberHi", (value))
}

// GetVersionNumberHi gets the value of VersionNumberHi for the instance
func (instance *RSOP_ApplicationManagementPolicySetting) GetPropertyVersionNumberHi() (value uint32, err error) {
	retValue, err := instance.GetProperty("VersionNumberHi")
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

// SetVersionNumberLo sets the value of VersionNumberLo for the instance
func (instance *RSOP_ApplicationManagementPolicySetting) SetPropertyVersionNumberLo(value uint32) (err error) {
	return instance.SetProperty("VersionNumberLo", (value))
}

// GetVersionNumberLo gets the value of VersionNumberLo for the instance
func (instance *RSOP_ApplicationManagementPolicySetting) GetPropertyVersionNumberLo() (value uint32, err error) {
	retValue, err := instance.GetProperty("VersionNumberLo")
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
