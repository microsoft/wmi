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

// MDM_Policy_Config01_TextInput02 struct
type MDM_Policy_Config01_TextInput02 struct {
	*cim.WmiInstance

	//
	AllowHardwareKeyboardTextSuggestions int32

	//
	AllowIMELogging int32

	//
	AllowIMENetworkAccess int32

	//
	AllowInputPanel int32

	//
	AllowJapaneseIMESurrogatePairCharacters int32

	//
	AllowJapaneseIVSCharacters int32

	//
	AllowJapaneseNonPublishingStandardGlyph int32

	//
	AllowJapaneseUserDictionary int32

	//
	AllowKeyboardTextSuggestions int32

	//
	AllowLanguageFeaturesUninstall int32

	//
	AllowLinguisticDataCollection int32

	//
	AllowTextInputSuggestionUpdate int32

	//
	ConfigureJapaneseIMEVersion int32

	//
	ConfigureKoreanIMEVersion int32

	//
	ConfigureSimplifiedChineseIMEVersion int32

	//
	ConfigureTraditionalChineseIMEVersion int32

	//
	EnableTouchKeyboardAutoInvokeInDesktopMode int32

	//
	ExcludeJapaneseIMEExceptJIS0208 int32

	//
	ExcludeJapaneseIMEExceptJIS0208andEUDC int32

	//
	ExcludeJapaneseIMEExceptShiftJIS int32

	//
	ForceTouchKeyboardDockedState int32

	//
	InstanceID string

	//
	ParentID string

	//
	TouchKeyboardDictationButtonAvailability int32

	//
	TouchKeyboardEmojiButtonAvailability int32

	//
	TouchKeyboardFullModeAvailability int32

	//
	TouchKeyboardHandwritingModeAvailability int32

	//
	TouchKeyboardNarrowModeAvailability int32

	//
	TouchKeyboardSplitModeAvailability int32

	//
	TouchKeyboardWideModeAvailability int32
}

func NewMDM_Policy_Config01_TextInput02Ex1(instance *cim.WmiInstance) (newInstance *MDM_Policy_Config01_TextInput02, err error) {
	tmp, err := instance, nil

	if err != nil {
		return
	}
	newInstance = &MDM_Policy_Config01_TextInput02{
		WmiInstance: tmp,
	}
	return
}

func NewMDM_Policy_Config01_TextInput02Ex6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *MDM_Policy_Config01_TextInput02, err error) {
	tmp, err := instance.GetWmiInstance(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &MDM_Policy_Config01_TextInput02{
		WmiInstance: tmp,
	}
	return
}

// SetAllowHardwareKeyboardTextSuggestions sets the value of AllowHardwareKeyboardTextSuggestions for the instance
func (instance *MDM_Policy_Config01_TextInput02) SetPropertyAllowHardwareKeyboardTextSuggestions(value int32) (err error) {
	return instance.SetProperty("AllowHardwareKeyboardTextSuggestions", (value))
}

// GetAllowHardwareKeyboardTextSuggestions gets the value of AllowHardwareKeyboardTextSuggestions for the instance
func (instance *MDM_Policy_Config01_TextInput02) GetPropertyAllowHardwareKeyboardTextSuggestions() (value int32, err error) {
	retValue, err := instance.GetProperty("AllowHardwareKeyboardTextSuggestions")
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

// SetAllowIMELogging sets the value of AllowIMELogging for the instance
func (instance *MDM_Policy_Config01_TextInput02) SetPropertyAllowIMELogging(value int32) (err error) {
	return instance.SetProperty("AllowIMELogging", (value))
}

// GetAllowIMELogging gets the value of AllowIMELogging for the instance
func (instance *MDM_Policy_Config01_TextInput02) GetPropertyAllowIMELogging() (value int32, err error) {
	retValue, err := instance.GetProperty("AllowIMELogging")
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

// SetAllowIMENetworkAccess sets the value of AllowIMENetworkAccess for the instance
func (instance *MDM_Policy_Config01_TextInput02) SetPropertyAllowIMENetworkAccess(value int32) (err error) {
	return instance.SetProperty("AllowIMENetworkAccess", (value))
}

// GetAllowIMENetworkAccess gets the value of AllowIMENetworkAccess for the instance
func (instance *MDM_Policy_Config01_TextInput02) GetPropertyAllowIMENetworkAccess() (value int32, err error) {
	retValue, err := instance.GetProperty("AllowIMENetworkAccess")
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

// SetAllowInputPanel sets the value of AllowInputPanel for the instance
func (instance *MDM_Policy_Config01_TextInput02) SetPropertyAllowInputPanel(value int32) (err error) {
	return instance.SetProperty("AllowInputPanel", (value))
}

// GetAllowInputPanel gets the value of AllowInputPanel for the instance
func (instance *MDM_Policy_Config01_TextInput02) GetPropertyAllowInputPanel() (value int32, err error) {
	retValue, err := instance.GetProperty("AllowInputPanel")
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

// SetAllowJapaneseIMESurrogatePairCharacters sets the value of AllowJapaneseIMESurrogatePairCharacters for the instance
func (instance *MDM_Policy_Config01_TextInput02) SetPropertyAllowJapaneseIMESurrogatePairCharacters(value int32) (err error) {
	return instance.SetProperty("AllowJapaneseIMESurrogatePairCharacters", (value))
}

// GetAllowJapaneseIMESurrogatePairCharacters gets the value of AllowJapaneseIMESurrogatePairCharacters for the instance
func (instance *MDM_Policy_Config01_TextInput02) GetPropertyAllowJapaneseIMESurrogatePairCharacters() (value int32, err error) {
	retValue, err := instance.GetProperty("AllowJapaneseIMESurrogatePairCharacters")
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

// SetAllowJapaneseIVSCharacters sets the value of AllowJapaneseIVSCharacters for the instance
func (instance *MDM_Policy_Config01_TextInput02) SetPropertyAllowJapaneseIVSCharacters(value int32) (err error) {
	return instance.SetProperty("AllowJapaneseIVSCharacters", (value))
}

// GetAllowJapaneseIVSCharacters gets the value of AllowJapaneseIVSCharacters for the instance
func (instance *MDM_Policy_Config01_TextInput02) GetPropertyAllowJapaneseIVSCharacters() (value int32, err error) {
	retValue, err := instance.GetProperty("AllowJapaneseIVSCharacters")
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

// SetAllowJapaneseNonPublishingStandardGlyph sets the value of AllowJapaneseNonPublishingStandardGlyph for the instance
func (instance *MDM_Policy_Config01_TextInput02) SetPropertyAllowJapaneseNonPublishingStandardGlyph(value int32) (err error) {
	return instance.SetProperty("AllowJapaneseNonPublishingStandardGlyph", (value))
}

// GetAllowJapaneseNonPublishingStandardGlyph gets the value of AllowJapaneseNonPublishingStandardGlyph for the instance
func (instance *MDM_Policy_Config01_TextInput02) GetPropertyAllowJapaneseNonPublishingStandardGlyph() (value int32, err error) {
	retValue, err := instance.GetProperty("AllowJapaneseNonPublishingStandardGlyph")
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

// SetAllowJapaneseUserDictionary sets the value of AllowJapaneseUserDictionary for the instance
func (instance *MDM_Policy_Config01_TextInput02) SetPropertyAllowJapaneseUserDictionary(value int32) (err error) {
	return instance.SetProperty("AllowJapaneseUserDictionary", (value))
}

// GetAllowJapaneseUserDictionary gets the value of AllowJapaneseUserDictionary for the instance
func (instance *MDM_Policy_Config01_TextInput02) GetPropertyAllowJapaneseUserDictionary() (value int32, err error) {
	retValue, err := instance.GetProperty("AllowJapaneseUserDictionary")
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

// SetAllowKeyboardTextSuggestions sets the value of AllowKeyboardTextSuggestions for the instance
func (instance *MDM_Policy_Config01_TextInput02) SetPropertyAllowKeyboardTextSuggestions(value int32) (err error) {
	return instance.SetProperty("AllowKeyboardTextSuggestions", (value))
}

// GetAllowKeyboardTextSuggestions gets the value of AllowKeyboardTextSuggestions for the instance
func (instance *MDM_Policy_Config01_TextInput02) GetPropertyAllowKeyboardTextSuggestions() (value int32, err error) {
	retValue, err := instance.GetProperty("AllowKeyboardTextSuggestions")
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

// SetAllowLanguageFeaturesUninstall sets the value of AllowLanguageFeaturesUninstall for the instance
func (instance *MDM_Policy_Config01_TextInput02) SetPropertyAllowLanguageFeaturesUninstall(value int32) (err error) {
	return instance.SetProperty("AllowLanguageFeaturesUninstall", (value))
}

// GetAllowLanguageFeaturesUninstall gets the value of AllowLanguageFeaturesUninstall for the instance
func (instance *MDM_Policy_Config01_TextInput02) GetPropertyAllowLanguageFeaturesUninstall() (value int32, err error) {
	retValue, err := instance.GetProperty("AllowLanguageFeaturesUninstall")
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

// SetAllowLinguisticDataCollection sets the value of AllowLinguisticDataCollection for the instance
func (instance *MDM_Policy_Config01_TextInput02) SetPropertyAllowLinguisticDataCollection(value int32) (err error) {
	return instance.SetProperty("AllowLinguisticDataCollection", (value))
}

// GetAllowLinguisticDataCollection gets the value of AllowLinguisticDataCollection for the instance
func (instance *MDM_Policy_Config01_TextInput02) GetPropertyAllowLinguisticDataCollection() (value int32, err error) {
	retValue, err := instance.GetProperty("AllowLinguisticDataCollection")
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

// SetAllowTextInputSuggestionUpdate sets the value of AllowTextInputSuggestionUpdate for the instance
func (instance *MDM_Policy_Config01_TextInput02) SetPropertyAllowTextInputSuggestionUpdate(value int32) (err error) {
	return instance.SetProperty("AllowTextInputSuggestionUpdate", (value))
}

// GetAllowTextInputSuggestionUpdate gets the value of AllowTextInputSuggestionUpdate for the instance
func (instance *MDM_Policy_Config01_TextInput02) GetPropertyAllowTextInputSuggestionUpdate() (value int32, err error) {
	retValue, err := instance.GetProperty("AllowTextInputSuggestionUpdate")
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

// SetConfigureJapaneseIMEVersion sets the value of ConfigureJapaneseIMEVersion for the instance
func (instance *MDM_Policy_Config01_TextInput02) SetPropertyConfigureJapaneseIMEVersion(value int32) (err error) {
	return instance.SetProperty("ConfigureJapaneseIMEVersion", (value))
}

// GetConfigureJapaneseIMEVersion gets the value of ConfigureJapaneseIMEVersion for the instance
func (instance *MDM_Policy_Config01_TextInput02) GetPropertyConfigureJapaneseIMEVersion() (value int32, err error) {
	retValue, err := instance.GetProperty("ConfigureJapaneseIMEVersion")
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

// SetConfigureKoreanIMEVersion sets the value of ConfigureKoreanIMEVersion for the instance
func (instance *MDM_Policy_Config01_TextInput02) SetPropertyConfigureKoreanIMEVersion(value int32) (err error) {
	return instance.SetProperty("ConfigureKoreanIMEVersion", (value))
}

// GetConfigureKoreanIMEVersion gets the value of ConfigureKoreanIMEVersion for the instance
func (instance *MDM_Policy_Config01_TextInput02) GetPropertyConfigureKoreanIMEVersion() (value int32, err error) {
	retValue, err := instance.GetProperty("ConfigureKoreanIMEVersion")
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

// SetConfigureSimplifiedChineseIMEVersion sets the value of ConfigureSimplifiedChineseIMEVersion for the instance
func (instance *MDM_Policy_Config01_TextInput02) SetPropertyConfigureSimplifiedChineseIMEVersion(value int32) (err error) {
	return instance.SetProperty("ConfigureSimplifiedChineseIMEVersion", (value))
}

// GetConfigureSimplifiedChineseIMEVersion gets the value of ConfigureSimplifiedChineseIMEVersion for the instance
func (instance *MDM_Policy_Config01_TextInput02) GetPropertyConfigureSimplifiedChineseIMEVersion() (value int32, err error) {
	retValue, err := instance.GetProperty("ConfigureSimplifiedChineseIMEVersion")
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

// SetConfigureTraditionalChineseIMEVersion sets the value of ConfigureTraditionalChineseIMEVersion for the instance
func (instance *MDM_Policy_Config01_TextInput02) SetPropertyConfigureTraditionalChineseIMEVersion(value int32) (err error) {
	return instance.SetProperty("ConfigureTraditionalChineseIMEVersion", (value))
}

// GetConfigureTraditionalChineseIMEVersion gets the value of ConfigureTraditionalChineseIMEVersion for the instance
func (instance *MDM_Policy_Config01_TextInput02) GetPropertyConfigureTraditionalChineseIMEVersion() (value int32, err error) {
	retValue, err := instance.GetProperty("ConfigureTraditionalChineseIMEVersion")
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

// SetEnableTouchKeyboardAutoInvokeInDesktopMode sets the value of EnableTouchKeyboardAutoInvokeInDesktopMode for the instance
func (instance *MDM_Policy_Config01_TextInput02) SetPropertyEnableTouchKeyboardAutoInvokeInDesktopMode(value int32) (err error) {
	return instance.SetProperty("EnableTouchKeyboardAutoInvokeInDesktopMode", (value))
}

// GetEnableTouchKeyboardAutoInvokeInDesktopMode gets the value of EnableTouchKeyboardAutoInvokeInDesktopMode for the instance
func (instance *MDM_Policy_Config01_TextInput02) GetPropertyEnableTouchKeyboardAutoInvokeInDesktopMode() (value int32, err error) {
	retValue, err := instance.GetProperty("EnableTouchKeyboardAutoInvokeInDesktopMode")
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

// SetExcludeJapaneseIMEExceptJIS0208 sets the value of ExcludeJapaneseIMEExceptJIS0208 for the instance
func (instance *MDM_Policy_Config01_TextInput02) SetPropertyExcludeJapaneseIMEExceptJIS0208(value int32) (err error) {
	return instance.SetProperty("ExcludeJapaneseIMEExceptJIS0208", (value))
}

// GetExcludeJapaneseIMEExceptJIS0208 gets the value of ExcludeJapaneseIMEExceptJIS0208 for the instance
func (instance *MDM_Policy_Config01_TextInput02) GetPropertyExcludeJapaneseIMEExceptJIS0208() (value int32, err error) {
	retValue, err := instance.GetProperty("ExcludeJapaneseIMEExceptJIS0208")
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

// SetExcludeJapaneseIMEExceptJIS0208andEUDC sets the value of ExcludeJapaneseIMEExceptJIS0208andEUDC for the instance
func (instance *MDM_Policy_Config01_TextInput02) SetPropertyExcludeJapaneseIMEExceptJIS0208andEUDC(value int32) (err error) {
	return instance.SetProperty("ExcludeJapaneseIMEExceptJIS0208andEUDC", (value))
}

// GetExcludeJapaneseIMEExceptJIS0208andEUDC gets the value of ExcludeJapaneseIMEExceptJIS0208andEUDC for the instance
func (instance *MDM_Policy_Config01_TextInput02) GetPropertyExcludeJapaneseIMEExceptJIS0208andEUDC() (value int32, err error) {
	retValue, err := instance.GetProperty("ExcludeJapaneseIMEExceptJIS0208andEUDC")
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

// SetExcludeJapaneseIMEExceptShiftJIS sets the value of ExcludeJapaneseIMEExceptShiftJIS for the instance
func (instance *MDM_Policy_Config01_TextInput02) SetPropertyExcludeJapaneseIMEExceptShiftJIS(value int32) (err error) {
	return instance.SetProperty("ExcludeJapaneseIMEExceptShiftJIS", (value))
}

// GetExcludeJapaneseIMEExceptShiftJIS gets the value of ExcludeJapaneseIMEExceptShiftJIS for the instance
func (instance *MDM_Policy_Config01_TextInput02) GetPropertyExcludeJapaneseIMEExceptShiftJIS() (value int32, err error) {
	retValue, err := instance.GetProperty("ExcludeJapaneseIMEExceptShiftJIS")
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

// SetForceTouchKeyboardDockedState sets the value of ForceTouchKeyboardDockedState for the instance
func (instance *MDM_Policy_Config01_TextInput02) SetPropertyForceTouchKeyboardDockedState(value int32) (err error) {
	return instance.SetProperty("ForceTouchKeyboardDockedState", (value))
}

// GetForceTouchKeyboardDockedState gets the value of ForceTouchKeyboardDockedState for the instance
func (instance *MDM_Policy_Config01_TextInput02) GetPropertyForceTouchKeyboardDockedState() (value int32, err error) {
	retValue, err := instance.GetProperty("ForceTouchKeyboardDockedState")
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
func (instance *MDM_Policy_Config01_TextInput02) SetPropertyInstanceID(value string) (err error) {
	return instance.SetProperty("InstanceID", (value))
}

// GetInstanceID gets the value of InstanceID for the instance
func (instance *MDM_Policy_Config01_TextInput02) GetPropertyInstanceID() (value string, err error) {
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
func (instance *MDM_Policy_Config01_TextInput02) SetPropertyParentID(value string) (err error) {
	return instance.SetProperty("ParentID", (value))
}

// GetParentID gets the value of ParentID for the instance
func (instance *MDM_Policy_Config01_TextInput02) GetPropertyParentID() (value string, err error) {
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

// SetTouchKeyboardDictationButtonAvailability sets the value of TouchKeyboardDictationButtonAvailability for the instance
func (instance *MDM_Policy_Config01_TextInput02) SetPropertyTouchKeyboardDictationButtonAvailability(value int32) (err error) {
	return instance.SetProperty("TouchKeyboardDictationButtonAvailability", (value))
}

// GetTouchKeyboardDictationButtonAvailability gets the value of TouchKeyboardDictationButtonAvailability for the instance
func (instance *MDM_Policy_Config01_TextInput02) GetPropertyTouchKeyboardDictationButtonAvailability() (value int32, err error) {
	retValue, err := instance.GetProperty("TouchKeyboardDictationButtonAvailability")
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

// SetTouchKeyboardEmojiButtonAvailability sets the value of TouchKeyboardEmojiButtonAvailability for the instance
func (instance *MDM_Policy_Config01_TextInput02) SetPropertyTouchKeyboardEmojiButtonAvailability(value int32) (err error) {
	return instance.SetProperty("TouchKeyboardEmojiButtonAvailability", (value))
}

// GetTouchKeyboardEmojiButtonAvailability gets the value of TouchKeyboardEmojiButtonAvailability for the instance
func (instance *MDM_Policy_Config01_TextInput02) GetPropertyTouchKeyboardEmojiButtonAvailability() (value int32, err error) {
	retValue, err := instance.GetProperty("TouchKeyboardEmojiButtonAvailability")
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

// SetTouchKeyboardFullModeAvailability sets the value of TouchKeyboardFullModeAvailability for the instance
func (instance *MDM_Policy_Config01_TextInput02) SetPropertyTouchKeyboardFullModeAvailability(value int32) (err error) {
	return instance.SetProperty("TouchKeyboardFullModeAvailability", (value))
}

// GetTouchKeyboardFullModeAvailability gets the value of TouchKeyboardFullModeAvailability for the instance
func (instance *MDM_Policy_Config01_TextInput02) GetPropertyTouchKeyboardFullModeAvailability() (value int32, err error) {
	retValue, err := instance.GetProperty("TouchKeyboardFullModeAvailability")
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

// SetTouchKeyboardHandwritingModeAvailability sets the value of TouchKeyboardHandwritingModeAvailability for the instance
func (instance *MDM_Policy_Config01_TextInput02) SetPropertyTouchKeyboardHandwritingModeAvailability(value int32) (err error) {
	return instance.SetProperty("TouchKeyboardHandwritingModeAvailability", (value))
}

// GetTouchKeyboardHandwritingModeAvailability gets the value of TouchKeyboardHandwritingModeAvailability for the instance
func (instance *MDM_Policy_Config01_TextInput02) GetPropertyTouchKeyboardHandwritingModeAvailability() (value int32, err error) {
	retValue, err := instance.GetProperty("TouchKeyboardHandwritingModeAvailability")
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

// SetTouchKeyboardNarrowModeAvailability sets the value of TouchKeyboardNarrowModeAvailability for the instance
func (instance *MDM_Policy_Config01_TextInput02) SetPropertyTouchKeyboardNarrowModeAvailability(value int32) (err error) {
	return instance.SetProperty("TouchKeyboardNarrowModeAvailability", (value))
}

// GetTouchKeyboardNarrowModeAvailability gets the value of TouchKeyboardNarrowModeAvailability for the instance
func (instance *MDM_Policy_Config01_TextInput02) GetPropertyTouchKeyboardNarrowModeAvailability() (value int32, err error) {
	retValue, err := instance.GetProperty("TouchKeyboardNarrowModeAvailability")
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

// SetTouchKeyboardSplitModeAvailability sets the value of TouchKeyboardSplitModeAvailability for the instance
func (instance *MDM_Policy_Config01_TextInput02) SetPropertyTouchKeyboardSplitModeAvailability(value int32) (err error) {
	return instance.SetProperty("TouchKeyboardSplitModeAvailability", (value))
}

// GetTouchKeyboardSplitModeAvailability gets the value of TouchKeyboardSplitModeAvailability for the instance
func (instance *MDM_Policy_Config01_TextInput02) GetPropertyTouchKeyboardSplitModeAvailability() (value int32, err error) {
	retValue, err := instance.GetProperty("TouchKeyboardSplitModeAvailability")
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

// SetTouchKeyboardWideModeAvailability sets the value of TouchKeyboardWideModeAvailability for the instance
func (instance *MDM_Policy_Config01_TextInput02) SetPropertyTouchKeyboardWideModeAvailability(value int32) (err error) {
	return instance.SetProperty("TouchKeyboardWideModeAvailability", (value))
}

// GetTouchKeyboardWideModeAvailability gets the value of TouchKeyboardWideModeAvailability for the instance
func (instance *MDM_Policy_Config01_TextInput02) GetPropertyTouchKeyboardWideModeAvailability() (value int32, err error) {
	retValue, err := instance.GetProperty("TouchKeyboardWideModeAvailability")
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
