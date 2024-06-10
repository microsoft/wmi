// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

// Author:
//
//	Auto Generated on 6/6/2024 using wmigen
//	Source root.CIMV2.mdm.dmmap
//
// ////////////////////////////////////////////
package dmmap

import (
	"github.com/microsoft/wmi/pkg/base/instance"
	"github.com/microsoft/wmi/pkg/base/query"
	"github.com/microsoft/wmi/pkg/errors"
	cim "github.com/microsoft/wmi/pkg/wmiinstance"
	"reflect"
)

// MDM_Policy_Result01_Search02 struct
type MDM_Policy_Result01_Search02 struct {
	*cim.WmiInstance

	//
	AllowCloudSearch int32

	//
	AllowCortanaInAAD int32

	//
	AllowFindMyFiles int32

	//
	AllowIndexingEncryptedStoresOrItems int32

	//
	AllowSearchToUseLocation int32

	//
	AllowStoringImagesFromVisionSearch int32

	//
	AllowUsingDiacritics int32

	//
	AllowWindowsIndexer int32

	//
	AlwaysUseAutoLangDetection int32

	//
	DisableBackoff int32

	//
	DisableRemovableDriveIndexing int32

	//
	DoNotUseWebResults int32

	//
	InstanceID string

	//
	ParentID string

	//
	PreventIndexingLowDiskSpaceMB int32

	//
	PreventRemoteQueries int32
}

func NewMDM_Policy_Result01_Search02Ex1(instance *cim.WmiInstance) (newInstance *MDM_Policy_Result01_Search02, err error) {
	tmp, err := instance, nil

	if err != nil {
		return
	}
	newInstance = &MDM_Policy_Result01_Search02{
		WmiInstance: tmp,
	}
	return
}

func NewMDM_Policy_Result01_Search02Ex6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *MDM_Policy_Result01_Search02, err error) {
	tmp, err := instance.GetWmiInstance(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &MDM_Policy_Result01_Search02{
		WmiInstance: tmp,
	}
	return
}

// SetAllowCloudSearch sets the value of AllowCloudSearch for the instance
func (instance *MDM_Policy_Result01_Search02) SetPropertyAllowCloudSearch(value int32) (err error) {
	return instance.SetProperty("AllowCloudSearch", (value))
}

// GetAllowCloudSearch gets the value of AllowCloudSearch for the instance
func (instance *MDM_Policy_Result01_Search02) GetPropertyAllowCloudSearch() (value int32, err error) {
	retValue, err := instance.GetProperty("AllowCloudSearch")
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

	value = int32(valuetmp)

	return
}

// SetAllowCortanaInAAD sets the value of AllowCortanaInAAD for the instance
func (instance *MDM_Policy_Result01_Search02) SetPropertyAllowCortanaInAAD(value int32) (err error) {
	return instance.SetProperty("AllowCortanaInAAD", (value))
}

// GetAllowCortanaInAAD gets the value of AllowCortanaInAAD for the instance
func (instance *MDM_Policy_Result01_Search02) GetPropertyAllowCortanaInAAD() (value int32, err error) {
	retValue, err := instance.GetProperty("AllowCortanaInAAD")
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

	value = int32(valuetmp)

	return
}

// SetAllowFindMyFiles sets the value of AllowFindMyFiles for the instance
func (instance *MDM_Policy_Result01_Search02) SetPropertyAllowFindMyFiles(value int32) (err error) {
	return instance.SetProperty("AllowFindMyFiles", (value))
}

// GetAllowFindMyFiles gets the value of AllowFindMyFiles for the instance
func (instance *MDM_Policy_Result01_Search02) GetPropertyAllowFindMyFiles() (value int32, err error) {
	retValue, err := instance.GetProperty("AllowFindMyFiles")
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

	value = int32(valuetmp)

	return
}

// SetAllowIndexingEncryptedStoresOrItems sets the value of AllowIndexingEncryptedStoresOrItems for the instance
func (instance *MDM_Policy_Result01_Search02) SetPropertyAllowIndexingEncryptedStoresOrItems(value int32) (err error) {
	return instance.SetProperty("AllowIndexingEncryptedStoresOrItems", (value))
}

// GetAllowIndexingEncryptedStoresOrItems gets the value of AllowIndexingEncryptedStoresOrItems for the instance
func (instance *MDM_Policy_Result01_Search02) GetPropertyAllowIndexingEncryptedStoresOrItems() (value int32, err error) {
	retValue, err := instance.GetProperty("AllowIndexingEncryptedStoresOrItems")
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

	value = int32(valuetmp)

	return
}

// SetAllowSearchToUseLocation sets the value of AllowSearchToUseLocation for the instance
func (instance *MDM_Policy_Result01_Search02) SetPropertyAllowSearchToUseLocation(value int32) (err error) {
	return instance.SetProperty("AllowSearchToUseLocation", (value))
}

// GetAllowSearchToUseLocation gets the value of AllowSearchToUseLocation for the instance
func (instance *MDM_Policy_Result01_Search02) GetPropertyAllowSearchToUseLocation() (value int32, err error) {
	retValue, err := instance.GetProperty("AllowSearchToUseLocation")
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

	value = int32(valuetmp)

	return
}

// SetAllowStoringImagesFromVisionSearch sets the value of AllowStoringImagesFromVisionSearch for the instance
func (instance *MDM_Policy_Result01_Search02) SetPropertyAllowStoringImagesFromVisionSearch(value int32) (err error) {
	return instance.SetProperty("AllowStoringImagesFromVisionSearch", (value))
}

// GetAllowStoringImagesFromVisionSearch gets the value of AllowStoringImagesFromVisionSearch for the instance
func (instance *MDM_Policy_Result01_Search02) GetPropertyAllowStoringImagesFromVisionSearch() (value int32, err error) {
	retValue, err := instance.GetProperty("AllowStoringImagesFromVisionSearch")
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

	value = int32(valuetmp)

	return
}

// SetAllowUsingDiacritics sets the value of AllowUsingDiacritics for the instance
func (instance *MDM_Policy_Result01_Search02) SetPropertyAllowUsingDiacritics(value int32) (err error) {
	return instance.SetProperty("AllowUsingDiacritics", (value))
}

// GetAllowUsingDiacritics gets the value of AllowUsingDiacritics for the instance
func (instance *MDM_Policy_Result01_Search02) GetPropertyAllowUsingDiacritics() (value int32, err error) {
	retValue, err := instance.GetProperty("AllowUsingDiacritics")
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

	value = int32(valuetmp)

	return
}

// SetAllowWindowsIndexer sets the value of AllowWindowsIndexer for the instance
func (instance *MDM_Policy_Result01_Search02) SetPropertyAllowWindowsIndexer(value int32) (err error) {
	return instance.SetProperty("AllowWindowsIndexer", (value))
}

// GetAllowWindowsIndexer gets the value of AllowWindowsIndexer for the instance
func (instance *MDM_Policy_Result01_Search02) GetPropertyAllowWindowsIndexer() (value int32, err error) {
	retValue, err := instance.GetProperty("AllowWindowsIndexer")
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

	value = int32(valuetmp)

	return
}

// SetAlwaysUseAutoLangDetection sets the value of AlwaysUseAutoLangDetection for the instance
func (instance *MDM_Policy_Result01_Search02) SetPropertyAlwaysUseAutoLangDetection(value int32) (err error) {
	return instance.SetProperty("AlwaysUseAutoLangDetection", (value))
}

// GetAlwaysUseAutoLangDetection gets the value of AlwaysUseAutoLangDetection for the instance
func (instance *MDM_Policy_Result01_Search02) GetPropertyAlwaysUseAutoLangDetection() (value int32, err error) {
	retValue, err := instance.GetProperty("AlwaysUseAutoLangDetection")
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

	value = int32(valuetmp)

	return
}

// SetDisableBackoff sets the value of DisableBackoff for the instance
func (instance *MDM_Policy_Result01_Search02) SetPropertyDisableBackoff(value int32) (err error) {
	return instance.SetProperty("DisableBackoff", (value))
}

// GetDisableBackoff gets the value of DisableBackoff for the instance
func (instance *MDM_Policy_Result01_Search02) GetPropertyDisableBackoff() (value int32, err error) {
	retValue, err := instance.GetProperty("DisableBackoff")
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

	value = int32(valuetmp)

	return
}

// SetDisableRemovableDriveIndexing sets the value of DisableRemovableDriveIndexing for the instance
func (instance *MDM_Policy_Result01_Search02) SetPropertyDisableRemovableDriveIndexing(value int32) (err error) {
	return instance.SetProperty("DisableRemovableDriveIndexing", (value))
}

// GetDisableRemovableDriveIndexing gets the value of DisableRemovableDriveIndexing for the instance
func (instance *MDM_Policy_Result01_Search02) GetPropertyDisableRemovableDriveIndexing() (value int32, err error) {
	retValue, err := instance.GetProperty("DisableRemovableDriveIndexing")
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

	value = int32(valuetmp)

	return
}

// SetDoNotUseWebResults sets the value of DoNotUseWebResults for the instance
func (instance *MDM_Policy_Result01_Search02) SetPropertyDoNotUseWebResults(value int32) (err error) {
	return instance.SetProperty("DoNotUseWebResults", (value))
}

// GetDoNotUseWebResults gets the value of DoNotUseWebResults for the instance
func (instance *MDM_Policy_Result01_Search02) GetPropertyDoNotUseWebResults() (value int32, err error) {
	retValue, err := instance.GetProperty("DoNotUseWebResults")
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

	value = int32(valuetmp)

	return
}

// SetInstanceID sets the value of InstanceID for the instance
func (instance *MDM_Policy_Result01_Search02) SetPropertyInstanceID(value string) (err error) {
	return instance.SetProperty("InstanceID", (value))
}

// GetInstanceID gets the value of InstanceID for the instance
func (instance *MDM_Policy_Result01_Search02) GetPropertyInstanceID() (value string, err error) {
	retValue, err := instance.GetProperty("InstanceID")
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

// SetParentID sets the value of ParentID for the instance
func (instance *MDM_Policy_Result01_Search02) SetPropertyParentID(value string) (err error) {
	return instance.SetProperty("ParentID", (value))
}

// GetParentID gets the value of ParentID for the instance
func (instance *MDM_Policy_Result01_Search02) GetPropertyParentID() (value string, err error) {
	retValue, err := instance.GetProperty("ParentID")
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

// SetPreventIndexingLowDiskSpaceMB sets the value of PreventIndexingLowDiskSpaceMB for the instance
func (instance *MDM_Policy_Result01_Search02) SetPropertyPreventIndexingLowDiskSpaceMB(value int32) (err error) {
	return instance.SetProperty("PreventIndexingLowDiskSpaceMB", (value))
}

// GetPreventIndexingLowDiskSpaceMB gets the value of PreventIndexingLowDiskSpaceMB for the instance
func (instance *MDM_Policy_Result01_Search02) GetPropertyPreventIndexingLowDiskSpaceMB() (value int32, err error) {
	retValue, err := instance.GetProperty("PreventIndexingLowDiskSpaceMB")
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

	value = int32(valuetmp)

	return
}

// SetPreventRemoteQueries sets the value of PreventRemoteQueries for the instance
func (instance *MDM_Policy_Result01_Search02) SetPropertyPreventRemoteQueries(value int32) (err error) {
	return instance.SetProperty("PreventRemoteQueries", (value))
}

// GetPreventRemoteQueries gets the value of PreventRemoteQueries for the instance
func (instance *MDM_Policy_Result01_Search02) GetPropertyPreventRemoteQueries() (value int32, err error) {
	retValue, err := instance.GetProperty("PreventRemoteQueries")
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

	value = int32(valuetmp)

	return
}
