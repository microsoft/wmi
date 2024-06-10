// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

// Author:
//
//	Auto Generated on 6/6/2024 using wmigen
//	Source root.WebAdministration
//
// ////////////////////////////////////////////
package webadministration

import (
	"github.com/microsoft/wmi/pkg/base/query"
	"github.com/microsoft/wmi/pkg/errors"
	cim "github.com/microsoft/wmi/pkg/wmiinstance"
	"reflect"
)

// CompilationSection struct
type CompilationSection struct {
	*ConfigurationSectionWithCollection

	//
	Assemblies AssemblySettings

	//
	AssemblyPostProcessorType string

	//
	Batch bool

	//
	BatchTimeout string

	//
	BuildProviders BuildProviderSettings

	//
	CodeSubDirectories DirectorySettings

	//
	Debug bool

	//
	DefaultLanguage string

	//
	EnablePrefetchOptimization bool

	//
	Explicit bool

	//
	ExpressionBuilders ExpressionBuilderSettings

	//
	FolderLevelBuildProviders FolderLevelBuildProviderSettings

	//
	MaxBatchGeneratedFileSize int32

	//
	MaxBatchSize int32

	//
	NumRecompilesBeforeAppRestart int32

	//
	OptimizeCompilations bool

	//
	ProfileGuidedOptimizations string

	//
	Strict bool

	//
	TargetFramework string

	//
	TempDirectory string

	//
	UrlLinePragmas bool
}

func NewCompilationSectionEx1(instance *cim.WmiInstance) (newInstance *CompilationSection, err error) {
	tmp, err := NewConfigurationSectionWithCollectionEx1(instance)

	if err != nil {
		return
	}
	newInstance = &CompilationSection{
		ConfigurationSectionWithCollection: tmp,
	}
	return
}

func NewCompilationSectionEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *CompilationSection, err error) {
	tmp, err := NewConfigurationSectionWithCollectionEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &CompilationSection{
		ConfigurationSectionWithCollection: tmp,
	}
	return
}

// SetAssemblies sets the value of Assemblies for the instance
func (instance *CompilationSection) SetPropertyAssemblies(value AssemblySettings) (err error) {
	return instance.SetProperty("Assemblies", (value))
}

// GetAssemblies gets the value of Assemblies for the instance
func (instance *CompilationSection) GetPropertyAssemblies() (value AssemblySettings, err error) {
	retValue, err := instance.GetProperty("Assemblies")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(AssemblySettings)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " AssemblySettings is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = AssemblySettings(valuetmp)

	return
}

// SetAssemblyPostProcessorType sets the value of AssemblyPostProcessorType for the instance
func (instance *CompilationSection) SetPropertyAssemblyPostProcessorType(value string) (err error) {
	return instance.SetProperty("AssemblyPostProcessorType", (value))
}

// GetAssemblyPostProcessorType gets the value of AssemblyPostProcessorType for the instance
func (instance *CompilationSection) GetPropertyAssemblyPostProcessorType() (value string, err error) {
	retValue, err := instance.GetProperty("AssemblyPostProcessorType")
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

// SetBatch sets the value of Batch for the instance
func (instance *CompilationSection) SetPropertyBatch(value bool) (err error) {
	return instance.SetProperty("Batch", (value))
}

// GetBatch gets the value of Batch for the instance
func (instance *CompilationSection) GetPropertyBatch() (value bool, err error) {
	retValue, err := instance.GetProperty("Batch")
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

// SetBatchTimeout sets the value of BatchTimeout for the instance
func (instance *CompilationSection) SetPropertyBatchTimeout(value string) (err error) {
	return instance.SetProperty("BatchTimeout", (value))
}

// GetBatchTimeout gets the value of BatchTimeout for the instance
func (instance *CompilationSection) GetPropertyBatchTimeout() (value string, err error) {
	retValue, err := instance.GetProperty("BatchTimeout")
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

// SetBuildProviders sets the value of BuildProviders for the instance
func (instance *CompilationSection) SetPropertyBuildProviders(value BuildProviderSettings) (err error) {
	return instance.SetProperty("BuildProviders", (value))
}

// GetBuildProviders gets the value of BuildProviders for the instance
func (instance *CompilationSection) GetPropertyBuildProviders() (value BuildProviderSettings, err error) {
	retValue, err := instance.GetProperty("BuildProviders")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(BuildProviderSettings)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " BuildProviderSettings is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = BuildProviderSettings(valuetmp)

	return
}

// SetCodeSubDirectories sets the value of CodeSubDirectories for the instance
func (instance *CompilationSection) SetPropertyCodeSubDirectories(value DirectorySettings) (err error) {
	return instance.SetProperty("CodeSubDirectories", (value))
}

// GetCodeSubDirectories gets the value of CodeSubDirectories for the instance
func (instance *CompilationSection) GetPropertyCodeSubDirectories() (value DirectorySettings, err error) {
	retValue, err := instance.GetProperty("CodeSubDirectories")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(DirectorySettings)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " DirectorySettings is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = DirectorySettings(valuetmp)

	return
}

// SetDebug sets the value of Debug for the instance
func (instance *CompilationSection) SetPropertyDebug(value bool) (err error) {
	return instance.SetProperty("Debug", (value))
}

// GetDebug gets the value of Debug for the instance
func (instance *CompilationSection) GetPropertyDebug() (value bool, err error) {
	retValue, err := instance.GetProperty("Debug")
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

// SetDefaultLanguage sets the value of DefaultLanguage for the instance
func (instance *CompilationSection) SetPropertyDefaultLanguage(value string) (err error) {
	return instance.SetProperty("DefaultLanguage", (value))
}

// GetDefaultLanguage gets the value of DefaultLanguage for the instance
func (instance *CompilationSection) GetPropertyDefaultLanguage() (value string, err error) {
	retValue, err := instance.GetProperty("DefaultLanguage")
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

// SetEnablePrefetchOptimization sets the value of EnablePrefetchOptimization for the instance
func (instance *CompilationSection) SetPropertyEnablePrefetchOptimization(value bool) (err error) {
	return instance.SetProperty("EnablePrefetchOptimization", (value))
}

// GetEnablePrefetchOptimization gets the value of EnablePrefetchOptimization for the instance
func (instance *CompilationSection) GetPropertyEnablePrefetchOptimization() (value bool, err error) {
	retValue, err := instance.GetProperty("EnablePrefetchOptimization")
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

// SetExplicit sets the value of Explicit for the instance
func (instance *CompilationSection) SetPropertyExplicit(value bool) (err error) {
	return instance.SetProperty("Explicit", (value))
}

// GetExplicit gets the value of Explicit for the instance
func (instance *CompilationSection) GetPropertyExplicit() (value bool, err error) {
	retValue, err := instance.GetProperty("Explicit")
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

// SetExpressionBuilders sets the value of ExpressionBuilders for the instance
func (instance *CompilationSection) SetPropertyExpressionBuilders(value ExpressionBuilderSettings) (err error) {
	return instance.SetProperty("ExpressionBuilders", (value))
}

// GetExpressionBuilders gets the value of ExpressionBuilders for the instance
func (instance *CompilationSection) GetPropertyExpressionBuilders() (value ExpressionBuilderSettings, err error) {
	retValue, err := instance.GetProperty("ExpressionBuilders")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(ExpressionBuilderSettings)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " ExpressionBuilderSettings is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = ExpressionBuilderSettings(valuetmp)

	return
}

// SetFolderLevelBuildProviders sets the value of FolderLevelBuildProviders for the instance
func (instance *CompilationSection) SetPropertyFolderLevelBuildProviders(value FolderLevelBuildProviderSettings) (err error) {
	return instance.SetProperty("FolderLevelBuildProviders", (value))
}

// GetFolderLevelBuildProviders gets the value of FolderLevelBuildProviders for the instance
func (instance *CompilationSection) GetPropertyFolderLevelBuildProviders() (value FolderLevelBuildProviderSettings, err error) {
	retValue, err := instance.GetProperty("FolderLevelBuildProviders")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(FolderLevelBuildProviderSettings)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " FolderLevelBuildProviderSettings is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = FolderLevelBuildProviderSettings(valuetmp)

	return
}

// SetMaxBatchGeneratedFileSize sets the value of MaxBatchGeneratedFileSize for the instance
func (instance *CompilationSection) SetPropertyMaxBatchGeneratedFileSize(value int32) (err error) {
	return instance.SetProperty("MaxBatchGeneratedFileSize", (value))
}

// GetMaxBatchGeneratedFileSize gets the value of MaxBatchGeneratedFileSize for the instance
func (instance *CompilationSection) GetPropertyMaxBatchGeneratedFileSize() (value int32, err error) {
	retValue, err := instance.GetProperty("MaxBatchGeneratedFileSize")
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

// SetMaxBatchSize sets the value of MaxBatchSize for the instance
func (instance *CompilationSection) SetPropertyMaxBatchSize(value int32) (err error) {
	return instance.SetProperty("MaxBatchSize", (value))
}

// GetMaxBatchSize gets the value of MaxBatchSize for the instance
func (instance *CompilationSection) GetPropertyMaxBatchSize() (value int32, err error) {
	retValue, err := instance.GetProperty("MaxBatchSize")
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

// SetNumRecompilesBeforeAppRestart sets the value of NumRecompilesBeforeAppRestart for the instance
func (instance *CompilationSection) SetPropertyNumRecompilesBeforeAppRestart(value int32) (err error) {
	return instance.SetProperty("NumRecompilesBeforeAppRestart", (value))
}

// GetNumRecompilesBeforeAppRestart gets the value of NumRecompilesBeforeAppRestart for the instance
func (instance *CompilationSection) GetPropertyNumRecompilesBeforeAppRestart() (value int32, err error) {
	retValue, err := instance.GetProperty("NumRecompilesBeforeAppRestart")
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

// SetOptimizeCompilations sets the value of OptimizeCompilations for the instance
func (instance *CompilationSection) SetPropertyOptimizeCompilations(value bool) (err error) {
	return instance.SetProperty("OptimizeCompilations", (value))
}

// GetOptimizeCompilations gets the value of OptimizeCompilations for the instance
func (instance *CompilationSection) GetPropertyOptimizeCompilations() (value bool, err error) {
	retValue, err := instance.GetProperty("OptimizeCompilations")
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

// SetProfileGuidedOptimizations sets the value of ProfileGuidedOptimizations for the instance
func (instance *CompilationSection) SetPropertyProfileGuidedOptimizations(value string) (err error) {
	return instance.SetProperty("ProfileGuidedOptimizations", (value))
}

// GetProfileGuidedOptimizations gets the value of ProfileGuidedOptimizations for the instance
func (instance *CompilationSection) GetPropertyProfileGuidedOptimizations() (value string, err error) {
	retValue, err := instance.GetProperty("ProfileGuidedOptimizations")
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

// SetStrict sets the value of Strict for the instance
func (instance *CompilationSection) SetPropertyStrict(value bool) (err error) {
	return instance.SetProperty("Strict", (value))
}

// GetStrict gets the value of Strict for the instance
func (instance *CompilationSection) GetPropertyStrict() (value bool, err error) {
	retValue, err := instance.GetProperty("Strict")
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

// SetTargetFramework sets the value of TargetFramework for the instance
func (instance *CompilationSection) SetPropertyTargetFramework(value string) (err error) {
	return instance.SetProperty("TargetFramework", (value))
}

// GetTargetFramework gets the value of TargetFramework for the instance
func (instance *CompilationSection) GetPropertyTargetFramework() (value string, err error) {
	retValue, err := instance.GetProperty("TargetFramework")
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

// SetTempDirectory sets the value of TempDirectory for the instance
func (instance *CompilationSection) SetPropertyTempDirectory(value string) (err error) {
	return instance.SetProperty("TempDirectory", (value))
}

// GetTempDirectory gets the value of TempDirectory for the instance
func (instance *CompilationSection) GetPropertyTempDirectory() (value string, err error) {
	retValue, err := instance.GetProperty("TempDirectory")
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

// SetUrlLinePragmas sets the value of UrlLinePragmas for the instance
func (instance *CompilationSection) SetPropertyUrlLinePragmas(value bool) (err error) {
	return instance.SetProperty("UrlLinePragmas", (value))
}

// GetUrlLinePragmas gets the value of UrlLinePragmas for the instance
func (instance *CompilationSection) GetPropertyUrlLinePragmas() (value bool, err error) {
	retValue, err := instance.GetProperty("UrlLinePragmas")
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
