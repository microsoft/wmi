// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 6/14/2024 using wmigen
//      Source root.CIMV2.mdm.dmmap
//////////////////////////////////////////////
package dmmap

import (
	"github.com/microsoft/wmi/pkg/base/instance"
	"github.com/microsoft/wmi/pkg/base/query"
	"github.com/microsoft/wmi/pkg/errors"
	cim "github.com/microsoft/wmi/pkg/wmiinstance"
	"reflect"
)

// MDM_Policy_Result01_Privacy02 struct
type MDM_Policy_Result01_Privacy02 struct {
	*cim.WmiInstance

	//
	AllowAutoAcceptPairingAndPrivacyConsentPrompts int32

	//
	AllowCrossDeviceClipboard int32

	//
	AllowInputPersonalization int32

	//
	DisableAdvertisingId int32

	//
	DisablePrivacyExperience int32

	//
	EnableActivityFeed int32

	//
	InstanceID string

	//
	LetAppsAccessAccountInfo int32

	//
	LetAppsAccessAccountInfo_ForceAllowTheseApps string

	//
	LetAppsAccessAccountInfo_ForceDenyTheseApps string

	//
	LetAppsAccessAccountInfo_UserInControlOfTheseApps string

	//
	LetAppsAccessBackgroundSpatialPerception int32

	//
	LetAppsAccessBackgroundSpatialPerception_ForceAllowTheseApps string

	//
	LetAppsAccessBackgroundSpatialPerception_ForceDenyTheseApps string

	//
	LetAppsAccessBackgroundSpatialPerception_UserInControlOfTheseApps string

	//
	LetAppsAccessCalendar int32

	//
	LetAppsAccessCalendar_ForceAllowTheseApps string

	//
	LetAppsAccessCalendar_ForceDenyTheseApps string

	//
	LetAppsAccessCalendar_UserInControlOfTheseApps string

	//
	LetAppsAccessCallHistory int32

	//
	LetAppsAccessCallHistory_ForceAllowTheseApps string

	//
	LetAppsAccessCallHistory_ForceDenyTheseApps string

	//
	LetAppsAccessCallHistory_UserInControlOfTheseApps string

	//
	LetAppsAccessCamera int32

	//
	LetAppsAccessCamera_ForceAllowTheseApps string

	//
	LetAppsAccessCamera_ForceDenyTheseApps string

	//
	LetAppsAccessCamera_UserInControlOfTheseApps string

	//
	LetAppsAccessContacts int32

	//
	LetAppsAccessContacts_ForceAllowTheseApps string

	//
	LetAppsAccessContacts_ForceDenyTheseApps string

	//
	LetAppsAccessContacts_UserInControlOfTheseApps string

	//
	LetAppsAccessEmail int32

	//
	LetAppsAccessEmail_ForceAllowTheseApps string

	//
	LetAppsAccessEmail_ForceDenyTheseApps string

	//
	LetAppsAccessEmail_UserInControlOfTheseApps string

	//
	LetAppsAccessGazeInput int32

	//
	LetAppsAccessGazeInput_ForceAllowTheseApps string

	//
	LetAppsAccessGazeInput_ForceDenyTheseApps string

	//
	LetAppsAccessGazeInput_UserInControlOfTheseApps string

	//
	LetAppsAccessGraphicsCaptureProgrammatic int32

	//
	LetAppsAccessGraphicsCaptureProgrammatic_ForceAllowTheseApps string

	//
	LetAppsAccessGraphicsCaptureProgrammatic_ForceDenyTheseApps string

	//
	LetAppsAccessGraphicsCaptureProgrammatic_UserInControlOfTheseApps string

	//
	LetAppsAccessGraphicsCaptureWithoutBorder int32

	//
	LetAppsAccessGraphicsCaptureWithoutBorder_ForceAllowTheseApps string

	//
	LetAppsAccessGraphicsCaptureWithoutBorder_ForceDenyTheseApps string

	//
	LetAppsAccessGraphicsCaptureWithoutBorder_UserInControlOfTheseApps string

	//
	LetAppsAccessLocation int32

	//
	LetAppsAccessLocation_ForceAllowTheseApps string

	//
	LetAppsAccessLocation_ForceDenyTheseApps string

	//
	LetAppsAccessLocation_UserInControlOfTheseApps string

	//
	LetAppsAccessMessaging int32

	//
	LetAppsAccessMessaging_ForceAllowTheseApps string

	//
	LetAppsAccessMessaging_ForceDenyTheseApps string

	//
	LetAppsAccessMessaging_UserInControlOfTheseApps string

	//
	LetAppsAccessMicrophone int32

	//
	LetAppsAccessMicrophone_ForceAllowTheseApps string

	//
	LetAppsAccessMicrophone_ForceDenyTheseApps string

	//
	LetAppsAccessMicrophone_UserInControlOfTheseApps string

	//
	LetAppsAccessMotion int32

	//
	LetAppsAccessMotion_ForceAllowTheseApps string

	//
	LetAppsAccessMotion_ForceDenyTheseApps string

	//
	LetAppsAccessMotion_UserInControlOfTheseApps string

	//
	LetAppsAccessNotifications int32

	//
	LetAppsAccessNotifications_ForceAllowTheseApps string

	//
	LetAppsAccessNotifications_ForceDenyTheseApps string

	//
	LetAppsAccessNotifications_UserInControlOfTheseApps string

	//
	LetAppsAccessPhone int32

	//
	LetAppsAccessPhone_ForceAllowTheseApps string

	//
	LetAppsAccessPhone_ForceDenyTheseApps string

	//
	LetAppsAccessPhone_UserInControlOfTheseApps string

	//
	LetAppsAccessRadios int32

	//
	LetAppsAccessRadios_ForceAllowTheseApps string

	//
	LetAppsAccessRadios_ForceDenyTheseApps string

	//
	LetAppsAccessRadios_UserInControlOfTheseApps string

	//
	LetAppsAccessTasks int32

	//
	LetAppsAccessTasks_ForceAllowTheseApps string

	//
	LetAppsAccessTasks_ForceDenyTheseApps string

	//
	LetAppsAccessTasks_UserInControlOfTheseApps string

	//
	LetAppsAccessTrustedDevices int32

	//
	LetAppsAccessTrustedDevices_ForceAllowTheseApps string

	//
	LetAppsAccessTrustedDevices_ForceDenyTheseApps string

	//
	LetAppsAccessTrustedDevices_UserInControlOfTheseApps string

	//
	LetAppsActivateWithVoice int32

	//
	LetAppsActivateWithVoiceAboveLock int32

	//
	LetAppsGetDiagnosticInfo int32

	//
	LetAppsGetDiagnosticInfo_ForceAllowTheseApps string

	//
	LetAppsGetDiagnosticInfo_ForceDenyTheseApps string

	//
	LetAppsGetDiagnosticInfo_UserInControlOfTheseApps string

	//
	LetAppsRunInBackground int32

	//
	LetAppsRunInBackground_ForceAllowTheseApps string

	//
	LetAppsRunInBackground_ForceDenyTheseApps string

	//
	LetAppsRunInBackground_UserInControlOfTheseApps string

	//
	LetAppsSyncWithDevices int32

	//
	LetAppsSyncWithDevices_ForceAllowTheseApps string

	//
	LetAppsSyncWithDevices_ForceDenyTheseApps string

	//
	LetAppsSyncWithDevices_UserInControlOfTheseApps string

	//
	ParentID string

	//
	PublishUserActivities int32

	//
	UploadUserActivities int32
}

func NewMDM_Policy_Result01_Privacy02Ex1(instance *cim.WmiInstance) (newInstance *MDM_Policy_Result01_Privacy02, err error) {
	tmp, err := instance, nil

	if err != nil {
		return
	}
	newInstance = &MDM_Policy_Result01_Privacy02{
		WmiInstance: tmp,
	}
	return
}

func NewMDM_Policy_Result01_Privacy02Ex6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *MDM_Policy_Result01_Privacy02, err error) {
	tmp, err := instance.GetWmiInstance(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &MDM_Policy_Result01_Privacy02{
		WmiInstance: tmp,
	}
	return
}

// SetAllowAutoAcceptPairingAndPrivacyConsentPrompts sets the value of AllowAutoAcceptPairingAndPrivacyConsentPrompts for the instance
func (instance *MDM_Policy_Result01_Privacy02) SetPropertyAllowAutoAcceptPairingAndPrivacyConsentPrompts(value int32) (err error) {
	return instance.SetProperty("AllowAutoAcceptPairingAndPrivacyConsentPrompts", (value))
}

// GetAllowAutoAcceptPairingAndPrivacyConsentPrompts gets the value of AllowAutoAcceptPairingAndPrivacyConsentPrompts for the instance
func (instance *MDM_Policy_Result01_Privacy02) GetPropertyAllowAutoAcceptPairingAndPrivacyConsentPrompts() (value int32, err error) {
	retValue, err := instance.GetProperty("AllowAutoAcceptPairingAndPrivacyConsentPrompts")
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

// SetAllowCrossDeviceClipboard sets the value of AllowCrossDeviceClipboard for the instance
func (instance *MDM_Policy_Result01_Privacy02) SetPropertyAllowCrossDeviceClipboard(value int32) (err error) {
	return instance.SetProperty("AllowCrossDeviceClipboard", (value))
}

// GetAllowCrossDeviceClipboard gets the value of AllowCrossDeviceClipboard for the instance
func (instance *MDM_Policy_Result01_Privacy02) GetPropertyAllowCrossDeviceClipboard() (value int32, err error) {
	retValue, err := instance.GetProperty("AllowCrossDeviceClipboard")
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

// SetAllowInputPersonalization sets the value of AllowInputPersonalization for the instance
func (instance *MDM_Policy_Result01_Privacy02) SetPropertyAllowInputPersonalization(value int32) (err error) {
	return instance.SetProperty("AllowInputPersonalization", (value))
}

// GetAllowInputPersonalization gets the value of AllowInputPersonalization for the instance
func (instance *MDM_Policy_Result01_Privacy02) GetPropertyAllowInputPersonalization() (value int32, err error) {
	retValue, err := instance.GetProperty("AllowInputPersonalization")
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

// SetDisableAdvertisingId sets the value of DisableAdvertisingId for the instance
func (instance *MDM_Policy_Result01_Privacy02) SetPropertyDisableAdvertisingId(value int32) (err error) {
	return instance.SetProperty("DisableAdvertisingId", (value))
}

// GetDisableAdvertisingId gets the value of DisableAdvertisingId for the instance
func (instance *MDM_Policy_Result01_Privacy02) GetPropertyDisableAdvertisingId() (value int32, err error) {
	retValue, err := instance.GetProperty("DisableAdvertisingId")
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

// SetDisablePrivacyExperience sets the value of DisablePrivacyExperience for the instance
func (instance *MDM_Policy_Result01_Privacy02) SetPropertyDisablePrivacyExperience(value int32) (err error) {
	return instance.SetProperty("DisablePrivacyExperience", (value))
}

// GetDisablePrivacyExperience gets the value of DisablePrivacyExperience for the instance
func (instance *MDM_Policy_Result01_Privacy02) GetPropertyDisablePrivacyExperience() (value int32, err error) {
	retValue, err := instance.GetProperty("DisablePrivacyExperience")
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

// SetEnableActivityFeed sets the value of EnableActivityFeed for the instance
func (instance *MDM_Policy_Result01_Privacy02) SetPropertyEnableActivityFeed(value int32) (err error) {
	return instance.SetProperty("EnableActivityFeed", (value))
}

// GetEnableActivityFeed gets the value of EnableActivityFeed for the instance
func (instance *MDM_Policy_Result01_Privacy02) GetPropertyEnableActivityFeed() (value int32, err error) {
	retValue, err := instance.GetProperty("EnableActivityFeed")
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
func (instance *MDM_Policy_Result01_Privacy02) SetPropertyInstanceID(value string) (err error) {
	return instance.SetProperty("InstanceID", (value))
}

// GetInstanceID gets the value of InstanceID for the instance
func (instance *MDM_Policy_Result01_Privacy02) GetPropertyInstanceID() (value string, err error) {
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

// SetLetAppsAccessAccountInfo sets the value of LetAppsAccessAccountInfo for the instance
func (instance *MDM_Policy_Result01_Privacy02) SetPropertyLetAppsAccessAccountInfo(value int32) (err error) {
	return instance.SetProperty("LetAppsAccessAccountInfo", (value))
}

// GetLetAppsAccessAccountInfo gets the value of LetAppsAccessAccountInfo for the instance
func (instance *MDM_Policy_Result01_Privacy02) GetPropertyLetAppsAccessAccountInfo() (value int32, err error) {
	retValue, err := instance.GetProperty("LetAppsAccessAccountInfo")
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

// SetLetAppsAccessAccountInfo_ForceAllowTheseApps sets the value of LetAppsAccessAccountInfo_ForceAllowTheseApps for the instance
func (instance *MDM_Policy_Result01_Privacy02) SetPropertyLetAppsAccessAccountInfo_ForceAllowTheseApps(value string) (err error) {
	return instance.SetProperty("LetAppsAccessAccountInfo_ForceAllowTheseApps", (value))
}

// GetLetAppsAccessAccountInfo_ForceAllowTheseApps gets the value of LetAppsAccessAccountInfo_ForceAllowTheseApps for the instance
func (instance *MDM_Policy_Result01_Privacy02) GetPropertyLetAppsAccessAccountInfo_ForceAllowTheseApps() (value string, err error) {
	retValue, err := instance.GetProperty("LetAppsAccessAccountInfo_ForceAllowTheseApps")
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

// SetLetAppsAccessAccountInfo_ForceDenyTheseApps sets the value of LetAppsAccessAccountInfo_ForceDenyTheseApps for the instance
func (instance *MDM_Policy_Result01_Privacy02) SetPropertyLetAppsAccessAccountInfo_ForceDenyTheseApps(value string) (err error) {
	return instance.SetProperty("LetAppsAccessAccountInfo_ForceDenyTheseApps", (value))
}

// GetLetAppsAccessAccountInfo_ForceDenyTheseApps gets the value of LetAppsAccessAccountInfo_ForceDenyTheseApps for the instance
func (instance *MDM_Policy_Result01_Privacy02) GetPropertyLetAppsAccessAccountInfo_ForceDenyTheseApps() (value string, err error) {
	retValue, err := instance.GetProperty("LetAppsAccessAccountInfo_ForceDenyTheseApps")
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

// SetLetAppsAccessAccountInfo_UserInControlOfTheseApps sets the value of LetAppsAccessAccountInfo_UserInControlOfTheseApps for the instance
func (instance *MDM_Policy_Result01_Privacy02) SetPropertyLetAppsAccessAccountInfo_UserInControlOfTheseApps(value string) (err error) {
	return instance.SetProperty("LetAppsAccessAccountInfo_UserInControlOfTheseApps", (value))
}

// GetLetAppsAccessAccountInfo_UserInControlOfTheseApps gets the value of LetAppsAccessAccountInfo_UserInControlOfTheseApps for the instance
func (instance *MDM_Policy_Result01_Privacy02) GetPropertyLetAppsAccessAccountInfo_UserInControlOfTheseApps() (value string, err error) {
	retValue, err := instance.GetProperty("LetAppsAccessAccountInfo_UserInControlOfTheseApps")
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

// SetLetAppsAccessBackgroundSpatialPerception sets the value of LetAppsAccessBackgroundSpatialPerception for the instance
func (instance *MDM_Policy_Result01_Privacy02) SetPropertyLetAppsAccessBackgroundSpatialPerception(value int32) (err error) {
	return instance.SetProperty("LetAppsAccessBackgroundSpatialPerception", (value))
}

// GetLetAppsAccessBackgroundSpatialPerception gets the value of LetAppsAccessBackgroundSpatialPerception for the instance
func (instance *MDM_Policy_Result01_Privacy02) GetPropertyLetAppsAccessBackgroundSpatialPerception() (value int32, err error) {
	retValue, err := instance.GetProperty("LetAppsAccessBackgroundSpatialPerception")
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

// SetLetAppsAccessBackgroundSpatialPerception_ForceAllowTheseApps sets the value of LetAppsAccessBackgroundSpatialPerception_ForceAllowTheseApps for the instance
func (instance *MDM_Policy_Result01_Privacy02) SetPropertyLetAppsAccessBackgroundSpatialPerception_ForceAllowTheseApps(value string) (err error) {
	return instance.SetProperty("LetAppsAccessBackgroundSpatialPerception_ForceAllowTheseApps", (value))
}

// GetLetAppsAccessBackgroundSpatialPerception_ForceAllowTheseApps gets the value of LetAppsAccessBackgroundSpatialPerception_ForceAllowTheseApps for the instance
func (instance *MDM_Policy_Result01_Privacy02) GetPropertyLetAppsAccessBackgroundSpatialPerception_ForceAllowTheseApps() (value string, err error) {
	retValue, err := instance.GetProperty("LetAppsAccessBackgroundSpatialPerception_ForceAllowTheseApps")
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

// SetLetAppsAccessBackgroundSpatialPerception_ForceDenyTheseApps sets the value of LetAppsAccessBackgroundSpatialPerception_ForceDenyTheseApps for the instance
func (instance *MDM_Policy_Result01_Privacy02) SetPropertyLetAppsAccessBackgroundSpatialPerception_ForceDenyTheseApps(value string) (err error) {
	return instance.SetProperty("LetAppsAccessBackgroundSpatialPerception_ForceDenyTheseApps", (value))
}

// GetLetAppsAccessBackgroundSpatialPerception_ForceDenyTheseApps gets the value of LetAppsAccessBackgroundSpatialPerception_ForceDenyTheseApps for the instance
func (instance *MDM_Policy_Result01_Privacy02) GetPropertyLetAppsAccessBackgroundSpatialPerception_ForceDenyTheseApps() (value string, err error) {
	retValue, err := instance.GetProperty("LetAppsAccessBackgroundSpatialPerception_ForceDenyTheseApps")
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

// SetLetAppsAccessBackgroundSpatialPerception_UserInControlOfTheseApps sets the value of LetAppsAccessBackgroundSpatialPerception_UserInControlOfTheseApps for the instance
func (instance *MDM_Policy_Result01_Privacy02) SetPropertyLetAppsAccessBackgroundSpatialPerception_UserInControlOfTheseApps(value string) (err error) {
	return instance.SetProperty("LetAppsAccessBackgroundSpatialPerception_UserInControlOfTheseApps", (value))
}

// GetLetAppsAccessBackgroundSpatialPerception_UserInControlOfTheseApps gets the value of LetAppsAccessBackgroundSpatialPerception_UserInControlOfTheseApps for the instance
func (instance *MDM_Policy_Result01_Privacy02) GetPropertyLetAppsAccessBackgroundSpatialPerception_UserInControlOfTheseApps() (value string, err error) {
	retValue, err := instance.GetProperty("LetAppsAccessBackgroundSpatialPerception_UserInControlOfTheseApps")
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

// SetLetAppsAccessCalendar sets the value of LetAppsAccessCalendar for the instance
func (instance *MDM_Policy_Result01_Privacy02) SetPropertyLetAppsAccessCalendar(value int32) (err error) {
	return instance.SetProperty("LetAppsAccessCalendar", (value))
}

// GetLetAppsAccessCalendar gets the value of LetAppsAccessCalendar for the instance
func (instance *MDM_Policy_Result01_Privacy02) GetPropertyLetAppsAccessCalendar() (value int32, err error) {
	retValue, err := instance.GetProperty("LetAppsAccessCalendar")
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

// SetLetAppsAccessCalendar_ForceAllowTheseApps sets the value of LetAppsAccessCalendar_ForceAllowTheseApps for the instance
func (instance *MDM_Policy_Result01_Privacy02) SetPropertyLetAppsAccessCalendar_ForceAllowTheseApps(value string) (err error) {
	return instance.SetProperty("LetAppsAccessCalendar_ForceAllowTheseApps", (value))
}

// GetLetAppsAccessCalendar_ForceAllowTheseApps gets the value of LetAppsAccessCalendar_ForceAllowTheseApps for the instance
func (instance *MDM_Policy_Result01_Privacy02) GetPropertyLetAppsAccessCalendar_ForceAllowTheseApps() (value string, err error) {
	retValue, err := instance.GetProperty("LetAppsAccessCalendar_ForceAllowTheseApps")
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

// SetLetAppsAccessCalendar_ForceDenyTheseApps sets the value of LetAppsAccessCalendar_ForceDenyTheseApps for the instance
func (instance *MDM_Policy_Result01_Privacy02) SetPropertyLetAppsAccessCalendar_ForceDenyTheseApps(value string) (err error) {
	return instance.SetProperty("LetAppsAccessCalendar_ForceDenyTheseApps", (value))
}

// GetLetAppsAccessCalendar_ForceDenyTheseApps gets the value of LetAppsAccessCalendar_ForceDenyTheseApps for the instance
func (instance *MDM_Policy_Result01_Privacy02) GetPropertyLetAppsAccessCalendar_ForceDenyTheseApps() (value string, err error) {
	retValue, err := instance.GetProperty("LetAppsAccessCalendar_ForceDenyTheseApps")
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

// SetLetAppsAccessCalendar_UserInControlOfTheseApps sets the value of LetAppsAccessCalendar_UserInControlOfTheseApps for the instance
func (instance *MDM_Policy_Result01_Privacy02) SetPropertyLetAppsAccessCalendar_UserInControlOfTheseApps(value string) (err error) {
	return instance.SetProperty("LetAppsAccessCalendar_UserInControlOfTheseApps", (value))
}

// GetLetAppsAccessCalendar_UserInControlOfTheseApps gets the value of LetAppsAccessCalendar_UserInControlOfTheseApps for the instance
func (instance *MDM_Policy_Result01_Privacy02) GetPropertyLetAppsAccessCalendar_UserInControlOfTheseApps() (value string, err error) {
	retValue, err := instance.GetProperty("LetAppsAccessCalendar_UserInControlOfTheseApps")
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

// SetLetAppsAccessCallHistory sets the value of LetAppsAccessCallHistory for the instance
func (instance *MDM_Policy_Result01_Privacy02) SetPropertyLetAppsAccessCallHistory(value int32) (err error) {
	return instance.SetProperty("LetAppsAccessCallHistory", (value))
}

// GetLetAppsAccessCallHistory gets the value of LetAppsAccessCallHistory for the instance
func (instance *MDM_Policy_Result01_Privacy02) GetPropertyLetAppsAccessCallHistory() (value int32, err error) {
	retValue, err := instance.GetProperty("LetAppsAccessCallHistory")
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

// SetLetAppsAccessCallHistory_ForceAllowTheseApps sets the value of LetAppsAccessCallHistory_ForceAllowTheseApps for the instance
func (instance *MDM_Policy_Result01_Privacy02) SetPropertyLetAppsAccessCallHistory_ForceAllowTheseApps(value string) (err error) {
	return instance.SetProperty("LetAppsAccessCallHistory_ForceAllowTheseApps", (value))
}

// GetLetAppsAccessCallHistory_ForceAllowTheseApps gets the value of LetAppsAccessCallHistory_ForceAllowTheseApps for the instance
func (instance *MDM_Policy_Result01_Privacy02) GetPropertyLetAppsAccessCallHistory_ForceAllowTheseApps() (value string, err error) {
	retValue, err := instance.GetProperty("LetAppsAccessCallHistory_ForceAllowTheseApps")
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

// SetLetAppsAccessCallHistory_ForceDenyTheseApps sets the value of LetAppsAccessCallHistory_ForceDenyTheseApps for the instance
func (instance *MDM_Policy_Result01_Privacy02) SetPropertyLetAppsAccessCallHistory_ForceDenyTheseApps(value string) (err error) {
	return instance.SetProperty("LetAppsAccessCallHistory_ForceDenyTheseApps", (value))
}

// GetLetAppsAccessCallHistory_ForceDenyTheseApps gets the value of LetAppsAccessCallHistory_ForceDenyTheseApps for the instance
func (instance *MDM_Policy_Result01_Privacy02) GetPropertyLetAppsAccessCallHistory_ForceDenyTheseApps() (value string, err error) {
	retValue, err := instance.GetProperty("LetAppsAccessCallHistory_ForceDenyTheseApps")
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

// SetLetAppsAccessCallHistory_UserInControlOfTheseApps sets the value of LetAppsAccessCallHistory_UserInControlOfTheseApps for the instance
func (instance *MDM_Policy_Result01_Privacy02) SetPropertyLetAppsAccessCallHistory_UserInControlOfTheseApps(value string) (err error) {
	return instance.SetProperty("LetAppsAccessCallHistory_UserInControlOfTheseApps", (value))
}

// GetLetAppsAccessCallHistory_UserInControlOfTheseApps gets the value of LetAppsAccessCallHistory_UserInControlOfTheseApps for the instance
func (instance *MDM_Policy_Result01_Privacy02) GetPropertyLetAppsAccessCallHistory_UserInControlOfTheseApps() (value string, err error) {
	retValue, err := instance.GetProperty("LetAppsAccessCallHistory_UserInControlOfTheseApps")
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

// SetLetAppsAccessCamera sets the value of LetAppsAccessCamera for the instance
func (instance *MDM_Policy_Result01_Privacy02) SetPropertyLetAppsAccessCamera(value int32) (err error) {
	return instance.SetProperty("LetAppsAccessCamera", (value))
}

// GetLetAppsAccessCamera gets the value of LetAppsAccessCamera for the instance
func (instance *MDM_Policy_Result01_Privacy02) GetPropertyLetAppsAccessCamera() (value int32, err error) {
	retValue, err := instance.GetProperty("LetAppsAccessCamera")
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

// SetLetAppsAccessCamera_ForceAllowTheseApps sets the value of LetAppsAccessCamera_ForceAllowTheseApps for the instance
func (instance *MDM_Policy_Result01_Privacy02) SetPropertyLetAppsAccessCamera_ForceAllowTheseApps(value string) (err error) {
	return instance.SetProperty("LetAppsAccessCamera_ForceAllowTheseApps", (value))
}

// GetLetAppsAccessCamera_ForceAllowTheseApps gets the value of LetAppsAccessCamera_ForceAllowTheseApps for the instance
func (instance *MDM_Policy_Result01_Privacy02) GetPropertyLetAppsAccessCamera_ForceAllowTheseApps() (value string, err error) {
	retValue, err := instance.GetProperty("LetAppsAccessCamera_ForceAllowTheseApps")
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

// SetLetAppsAccessCamera_ForceDenyTheseApps sets the value of LetAppsAccessCamera_ForceDenyTheseApps for the instance
func (instance *MDM_Policy_Result01_Privacy02) SetPropertyLetAppsAccessCamera_ForceDenyTheseApps(value string) (err error) {
	return instance.SetProperty("LetAppsAccessCamera_ForceDenyTheseApps", (value))
}

// GetLetAppsAccessCamera_ForceDenyTheseApps gets the value of LetAppsAccessCamera_ForceDenyTheseApps for the instance
func (instance *MDM_Policy_Result01_Privacy02) GetPropertyLetAppsAccessCamera_ForceDenyTheseApps() (value string, err error) {
	retValue, err := instance.GetProperty("LetAppsAccessCamera_ForceDenyTheseApps")
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

// SetLetAppsAccessCamera_UserInControlOfTheseApps sets the value of LetAppsAccessCamera_UserInControlOfTheseApps for the instance
func (instance *MDM_Policy_Result01_Privacy02) SetPropertyLetAppsAccessCamera_UserInControlOfTheseApps(value string) (err error) {
	return instance.SetProperty("LetAppsAccessCamera_UserInControlOfTheseApps", (value))
}

// GetLetAppsAccessCamera_UserInControlOfTheseApps gets the value of LetAppsAccessCamera_UserInControlOfTheseApps for the instance
func (instance *MDM_Policy_Result01_Privacy02) GetPropertyLetAppsAccessCamera_UserInControlOfTheseApps() (value string, err error) {
	retValue, err := instance.GetProperty("LetAppsAccessCamera_UserInControlOfTheseApps")
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

// SetLetAppsAccessContacts sets the value of LetAppsAccessContacts for the instance
func (instance *MDM_Policy_Result01_Privacy02) SetPropertyLetAppsAccessContacts(value int32) (err error) {
	return instance.SetProperty("LetAppsAccessContacts", (value))
}

// GetLetAppsAccessContacts gets the value of LetAppsAccessContacts for the instance
func (instance *MDM_Policy_Result01_Privacy02) GetPropertyLetAppsAccessContacts() (value int32, err error) {
	retValue, err := instance.GetProperty("LetAppsAccessContacts")
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

// SetLetAppsAccessContacts_ForceAllowTheseApps sets the value of LetAppsAccessContacts_ForceAllowTheseApps for the instance
func (instance *MDM_Policy_Result01_Privacy02) SetPropertyLetAppsAccessContacts_ForceAllowTheseApps(value string) (err error) {
	return instance.SetProperty("LetAppsAccessContacts_ForceAllowTheseApps", (value))
}

// GetLetAppsAccessContacts_ForceAllowTheseApps gets the value of LetAppsAccessContacts_ForceAllowTheseApps for the instance
func (instance *MDM_Policy_Result01_Privacy02) GetPropertyLetAppsAccessContacts_ForceAllowTheseApps() (value string, err error) {
	retValue, err := instance.GetProperty("LetAppsAccessContacts_ForceAllowTheseApps")
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

// SetLetAppsAccessContacts_ForceDenyTheseApps sets the value of LetAppsAccessContacts_ForceDenyTheseApps for the instance
func (instance *MDM_Policy_Result01_Privacy02) SetPropertyLetAppsAccessContacts_ForceDenyTheseApps(value string) (err error) {
	return instance.SetProperty("LetAppsAccessContacts_ForceDenyTheseApps", (value))
}

// GetLetAppsAccessContacts_ForceDenyTheseApps gets the value of LetAppsAccessContacts_ForceDenyTheseApps for the instance
func (instance *MDM_Policy_Result01_Privacy02) GetPropertyLetAppsAccessContacts_ForceDenyTheseApps() (value string, err error) {
	retValue, err := instance.GetProperty("LetAppsAccessContacts_ForceDenyTheseApps")
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

// SetLetAppsAccessContacts_UserInControlOfTheseApps sets the value of LetAppsAccessContacts_UserInControlOfTheseApps for the instance
func (instance *MDM_Policy_Result01_Privacy02) SetPropertyLetAppsAccessContacts_UserInControlOfTheseApps(value string) (err error) {
	return instance.SetProperty("LetAppsAccessContacts_UserInControlOfTheseApps", (value))
}

// GetLetAppsAccessContacts_UserInControlOfTheseApps gets the value of LetAppsAccessContacts_UserInControlOfTheseApps for the instance
func (instance *MDM_Policy_Result01_Privacy02) GetPropertyLetAppsAccessContacts_UserInControlOfTheseApps() (value string, err error) {
	retValue, err := instance.GetProperty("LetAppsAccessContacts_UserInControlOfTheseApps")
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

// SetLetAppsAccessEmail sets the value of LetAppsAccessEmail for the instance
func (instance *MDM_Policy_Result01_Privacy02) SetPropertyLetAppsAccessEmail(value int32) (err error) {
	return instance.SetProperty("LetAppsAccessEmail", (value))
}

// GetLetAppsAccessEmail gets the value of LetAppsAccessEmail for the instance
func (instance *MDM_Policy_Result01_Privacy02) GetPropertyLetAppsAccessEmail() (value int32, err error) {
	retValue, err := instance.GetProperty("LetAppsAccessEmail")
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

// SetLetAppsAccessEmail_ForceAllowTheseApps sets the value of LetAppsAccessEmail_ForceAllowTheseApps for the instance
func (instance *MDM_Policy_Result01_Privacy02) SetPropertyLetAppsAccessEmail_ForceAllowTheseApps(value string) (err error) {
	return instance.SetProperty("LetAppsAccessEmail_ForceAllowTheseApps", (value))
}

// GetLetAppsAccessEmail_ForceAllowTheseApps gets the value of LetAppsAccessEmail_ForceAllowTheseApps for the instance
func (instance *MDM_Policy_Result01_Privacy02) GetPropertyLetAppsAccessEmail_ForceAllowTheseApps() (value string, err error) {
	retValue, err := instance.GetProperty("LetAppsAccessEmail_ForceAllowTheseApps")
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

// SetLetAppsAccessEmail_ForceDenyTheseApps sets the value of LetAppsAccessEmail_ForceDenyTheseApps for the instance
func (instance *MDM_Policy_Result01_Privacy02) SetPropertyLetAppsAccessEmail_ForceDenyTheseApps(value string) (err error) {
	return instance.SetProperty("LetAppsAccessEmail_ForceDenyTheseApps", (value))
}

// GetLetAppsAccessEmail_ForceDenyTheseApps gets the value of LetAppsAccessEmail_ForceDenyTheseApps for the instance
func (instance *MDM_Policy_Result01_Privacy02) GetPropertyLetAppsAccessEmail_ForceDenyTheseApps() (value string, err error) {
	retValue, err := instance.GetProperty("LetAppsAccessEmail_ForceDenyTheseApps")
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

// SetLetAppsAccessEmail_UserInControlOfTheseApps sets the value of LetAppsAccessEmail_UserInControlOfTheseApps for the instance
func (instance *MDM_Policy_Result01_Privacy02) SetPropertyLetAppsAccessEmail_UserInControlOfTheseApps(value string) (err error) {
	return instance.SetProperty("LetAppsAccessEmail_UserInControlOfTheseApps", (value))
}

// GetLetAppsAccessEmail_UserInControlOfTheseApps gets the value of LetAppsAccessEmail_UserInControlOfTheseApps for the instance
func (instance *MDM_Policy_Result01_Privacy02) GetPropertyLetAppsAccessEmail_UserInControlOfTheseApps() (value string, err error) {
	retValue, err := instance.GetProperty("LetAppsAccessEmail_UserInControlOfTheseApps")
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

// SetLetAppsAccessGazeInput sets the value of LetAppsAccessGazeInput for the instance
func (instance *MDM_Policy_Result01_Privacy02) SetPropertyLetAppsAccessGazeInput(value int32) (err error) {
	return instance.SetProperty("LetAppsAccessGazeInput", (value))
}

// GetLetAppsAccessGazeInput gets the value of LetAppsAccessGazeInput for the instance
func (instance *MDM_Policy_Result01_Privacy02) GetPropertyLetAppsAccessGazeInput() (value int32, err error) {
	retValue, err := instance.GetProperty("LetAppsAccessGazeInput")
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

// SetLetAppsAccessGazeInput_ForceAllowTheseApps sets the value of LetAppsAccessGazeInput_ForceAllowTheseApps for the instance
func (instance *MDM_Policy_Result01_Privacy02) SetPropertyLetAppsAccessGazeInput_ForceAllowTheseApps(value string) (err error) {
	return instance.SetProperty("LetAppsAccessGazeInput_ForceAllowTheseApps", (value))
}

// GetLetAppsAccessGazeInput_ForceAllowTheseApps gets the value of LetAppsAccessGazeInput_ForceAllowTheseApps for the instance
func (instance *MDM_Policy_Result01_Privacy02) GetPropertyLetAppsAccessGazeInput_ForceAllowTheseApps() (value string, err error) {
	retValue, err := instance.GetProperty("LetAppsAccessGazeInput_ForceAllowTheseApps")
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

// SetLetAppsAccessGazeInput_ForceDenyTheseApps sets the value of LetAppsAccessGazeInput_ForceDenyTheseApps for the instance
func (instance *MDM_Policy_Result01_Privacy02) SetPropertyLetAppsAccessGazeInput_ForceDenyTheseApps(value string) (err error) {
	return instance.SetProperty("LetAppsAccessGazeInput_ForceDenyTheseApps", (value))
}

// GetLetAppsAccessGazeInput_ForceDenyTheseApps gets the value of LetAppsAccessGazeInput_ForceDenyTheseApps for the instance
func (instance *MDM_Policy_Result01_Privacy02) GetPropertyLetAppsAccessGazeInput_ForceDenyTheseApps() (value string, err error) {
	retValue, err := instance.GetProperty("LetAppsAccessGazeInput_ForceDenyTheseApps")
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

// SetLetAppsAccessGazeInput_UserInControlOfTheseApps sets the value of LetAppsAccessGazeInput_UserInControlOfTheseApps for the instance
func (instance *MDM_Policy_Result01_Privacy02) SetPropertyLetAppsAccessGazeInput_UserInControlOfTheseApps(value string) (err error) {
	return instance.SetProperty("LetAppsAccessGazeInput_UserInControlOfTheseApps", (value))
}

// GetLetAppsAccessGazeInput_UserInControlOfTheseApps gets the value of LetAppsAccessGazeInput_UserInControlOfTheseApps for the instance
func (instance *MDM_Policy_Result01_Privacy02) GetPropertyLetAppsAccessGazeInput_UserInControlOfTheseApps() (value string, err error) {
	retValue, err := instance.GetProperty("LetAppsAccessGazeInput_UserInControlOfTheseApps")
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

// SetLetAppsAccessGraphicsCaptureProgrammatic sets the value of LetAppsAccessGraphicsCaptureProgrammatic for the instance
func (instance *MDM_Policy_Result01_Privacy02) SetPropertyLetAppsAccessGraphicsCaptureProgrammatic(value int32) (err error) {
	return instance.SetProperty("LetAppsAccessGraphicsCaptureProgrammatic", (value))
}

// GetLetAppsAccessGraphicsCaptureProgrammatic gets the value of LetAppsAccessGraphicsCaptureProgrammatic for the instance
func (instance *MDM_Policy_Result01_Privacy02) GetPropertyLetAppsAccessGraphicsCaptureProgrammatic() (value int32, err error) {
	retValue, err := instance.GetProperty("LetAppsAccessGraphicsCaptureProgrammatic")
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

// SetLetAppsAccessGraphicsCaptureProgrammatic_ForceAllowTheseApps sets the value of LetAppsAccessGraphicsCaptureProgrammatic_ForceAllowTheseApps for the instance
func (instance *MDM_Policy_Result01_Privacy02) SetPropertyLetAppsAccessGraphicsCaptureProgrammatic_ForceAllowTheseApps(value string) (err error) {
	return instance.SetProperty("LetAppsAccessGraphicsCaptureProgrammatic_ForceAllowTheseApps", (value))
}

// GetLetAppsAccessGraphicsCaptureProgrammatic_ForceAllowTheseApps gets the value of LetAppsAccessGraphicsCaptureProgrammatic_ForceAllowTheseApps for the instance
func (instance *MDM_Policy_Result01_Privacy02) GetPropertyLetAppsAccessGraphicsCaptureProgrammatic_ForceAllowTheseApps() (value string, err error) {
	retValue, err := instance.GetProperty("LetAppsAccessGraphicsCaptureProgrammatic_ForceAllowTheseApps")
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

// SetLetAppsAccessGraphicsCaptureProgrammatic_ForceDenyTheseApps sets the value of LetAppsAccessGraphicsCaptureProgrammatic_ForceDenyTheseApps for the instance
func (instance *MDM_Policy_Result01_Privacy02) SetPropertyLetAppsAccessGraphicsCaptureProgrammatic_ForceDenyTheseApps(value string) (err error) {
	return instance.SetProperty("LetAppsAccessGraphicsCaptureProgrammatic_ForceDenyTheseApps", (value))
}

// GetLetAppsAccessGraphicsCaptureProgrammatic_ForceDenyTheseApps gets the value of LetAppsAccessGraphicsCaptureProgrammatic_ForceDenyTheseApps for the instance
func (instance *MDM_Policy_Result01_Privacy02) GetPropertyLetAppsAccessGraphicsCaptureProgrammatic_ForceDenyTheseApps() (value string, err error) {
	retValue, err := instance.GetProperty("LetAppsAccessGraphicsCaptureProgrammatic_ForceDenyTheseApps")
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

// SetLetAppsAccessGraphicsCaptureProgrammatic_UserInControlOfTheseApps sets the value of LetAppsAccessGraphicsCaptureProgrammatic_UserInControlOfTheseApps for the instance
func (instance *MDM_Policy_Result01_Privacy02) SetPropertyLetAppsAccessGraphicsCaptureProgrammatic_UserInControlOfTheseApps(value string) (err error) {
	return instance.SetProperty("LetAppsAccessGraphicsCaptureProgrammatic_UserInControlOfTheseApps", (value))
}

// GetLetAppsAccessGraphicsCaptureProgrammatic_UserInControlOfTheseApps gets the value of LetAppsAccessGraphicsCaptureProgrammatic_UserInControlOfTheseApps for the instance
func (instance *MDM_Policy_Result01_Privacy02) GetPropertyLetAppsAccessGraphicsCaptureProgrammatic_UserInControlOfTheseApps() (value string, err error) {
	retValue, err := instance.GetProperty("LetAppsAccessGraphicsCaptureProgrammatic_UserInControlOfTheseApps")
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

// SetLetAppsAccessGraphicsCaptureWithoutBorder sets the value of LetAppsAccessGraphicsCaptureWithoutBorder for the instance
func (instance *MDM_Policy_Result01_Privacy02) SetPropertyLetAppsAccessGraphicsCaptureWithoutBorder(value int32) (err error) {
	return instance.SetProperty("LetAppsAccessGraphicsCaptureWithoutBorder", (value))
}

// GetLetAppsAccessGraphicsCaptureWithoutBorder gets the value of LetAppsAccessGraphicsCaptureWithoutBorder for the instance
func (instance *MDM_Policy_Result01_Privacy02) GetPropertyLetAppsAccessGraphicsCaptureWithoutBorder() (value int32, err error) {
	retValue, err := instance.GetProperty("LetAppsAccessGraphicsCaptureWithoutBorder")
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

// SetLetAppsAccessGraphicsCaptureWithoutBorder_ForceAllowTheseApps sets the value of LetAppsAccessGraphicsCaptureWithoutBorder_ForceAllowTheseApps for the instance
func (instance *MDM_Policy_Result01_Privacy02) SetPropertyLetAppsAccessGraphicsCaptureWithoutBorder_ForceAllowTheseApps(value string) (err error) {
	return instance.SetProperty("LetAppsAccessGraphicsCaptureWithoutBorder_ForceAllowTheseApps", (value))
}

// GetLetAppsAccessGraphicsCaptureWithoutBorder_ForceAllowTheseApps gets the value of LetAppsAccessGraphicsCaptureWithoutBorder_ForceAllowTheseApps for the instance
func (instance *MDM_Policy_Result01_Privacy02) GetPropertyLetAppsAccessGraphicsCaptureWithoutBorder_ForceAllowTheseApps() (value string, err error) {
	retValue, err := instance.GetProperty("LetAppsAccessGraphicsCaptureWithoutBorder_ForceAllowTheseApps")
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

// SetLetAppsAccessGraphicsCaptureWithoutBorder_ForceDenyTheseApps sets the value of LetAppsAccessGraphicsCaptureWithoutBorder_ForceDenyTheseApps for the instance
func (instance *MDM_Policy_Result01_Privacy02) SetPropertyLetAppsAccessGraphicsCaptureWithoutBorder_ForceDenyTheseApps(value string) (err error) {
	return instance.SetProperty("LetAppsAccessGraphicsCaptureWithoutBorder_ForceDenyTheseApps", (value))
}

// GetLetAppsAccessGraphicsCaptureWithoutBorder_ForceDenyTheseApps gets the value of LetAppsAccessGraphicsCaptureWithoutBorder_ForceDenyTheseApps for the instance
func (instance *MDM_Policy_Result01_Privacy02) GetPropertyLetAppsAccessGraphicsCaptureWithoutBorder_ForceDenyTheseApps() (value string, err error) {
	retValue, err := instance.GetProperty("LetAppsAccessGraphicsCaptureWithoutBorder_ForceDenyTheseApps")
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

// SetLetAppsAccessGraphicsCaptureWithoutBorder_UserInControlOfTheseApps sets the value of LetAppsAccessGraphicsCaptureWithoutBorder_UserInControlOfTheseApps for the instance
func (instance *MDM_Policy_Result01_Privacy02) SetPropertyLetAppsAccessGraphicsCaptureWithoutBorder_UserInControlOfTheseApps(value string) (err error) {
	return instance.SetProperty("LetAppsAccessGraphicsCaptureWithoutBorder_UserInControlOfTheseApps", (value))
}

// GetLetAppsAccessGraphicsCaptureWithoutBorder_UserInControlOfTheseApps gets the value of LetAppsAccessGraphicsCaptureWithoutBorder_UserInControlOfTheseApps for the instance
func (instance *MDM_Policy_Result01_Privacy02) GetPropertyLetAppsAccessGraphicsCaptureWithoutBorder_UserInControlOfTheseApps() (value string, err error) {
	retValue, err := instance.GetProperty("LetAppsAccessGraphicsCaptureWithoutBorder_UserInControlOfTheseApps")
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

// SetLetAppsAccessLocation sets the value of LetAppsAccessLocation for the instance
func (instance *MDM_Policy_Result01_Privacy02) SetPropertyLetAppsAccessLocation(value int32) (err error) {
	return instance.SetProperty("LetAppsAccessLocation", (value))
}

// GetLetAppsAccessLocation gets the value of LetAppsAccessLocation for the instance
func (instance *MDM_Policy_Result01_Privacy02) GetPropertyLetAppsAccessLocation() (value int32, err error) {
	retValue, err := instance.GetProperty("LetAppsAccessLocation")
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

// SetLetAppsAccessLocation_ForceAllowTheseApps sets the value of LetAppsAccessLocation_ForceAllowTheseApps for the instance
func (instance *MDM_Policy_Result01_Privacy02) SetPropertyLetAppsAccessLocation_ForceAllowTheseApps(value string) (err error) {
	return instance.SetProperty("LetAppsAccessLocation_ForceAllowTheseApps", (value))
}

// GetLetAppsAccessLocation_ForceAllowTheseApps gets the value of LetAppsAccessLocation_ForceAllowTheseApps for the instance
func (instance *MDM_Policy_Result01_Privacy02) GetPropertyLetAppsAccessLocation_ForceAllowTheseApps() (value string, err error) {
	retValue, err := instance.GetProperty("LetAppsAccessLocation_ForceAllowTheseApps")
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

// SetLetAppsAccessLocation_ForceDenyTheseApps sets the value of LetAppsAccessLocation_ForceDenyTheseApps for the instance
func (instance *MDM_Policy_Result01_Privacy02) SetPropertyLetAppsAccessLocation_ForceDenyTheseApps(value string) (err error) {
	return instance.SetProperty("LetAppsAccessLocation_ForceDenyTheseApps", (value))
}

// GetLetAppsAccessLocation_ForceDenyTheseApps gets the value of LetAppsAccessLocation_ForceDenyTheseApps for the instance
func (instance *MDM_Policy_Result01_Privacy02) GetPropertyLetAppsAccessLocation_ForceDenyTheseApps() (value string, err error) {
	retValue, err := instance.GetProperty("LetAppsAccessLocation_ForceDenyTheseApps")
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

// SetLetAppsAccessLocation_UserInControlOfTheseApps sets the value of LetAppsAccessLocation_UserInControlOfTheseApps for the instance
func (instance *MDM_Policy_Result01_Privacy02) SetPropertyLetAppsAccessLocation_UserInControlOfTheseApps(value string) (err error) {
	return instance.SetProperty("LetAppsAccessLocation_UserInControlOfTheseApps", (value))
}

// GetLetAppsAccessLocation_UserInControlOfTheseApps gets the value of LetAppsAccessLocation_UserInControlOfTheseApps for the instance
func (instance *MDM_Policy_Result01_Privacy02) GetPropertyLetAppsAccessLocation_UserInControlOfTheseApps() (value string, err error) {
	retValue, err := instance.GetProperty("LetAppsAccessLocation_UserInControlOfTheseApps")
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

// SetLetAppsAccessMessaging sets the value of LetAppsAccessMessaging for the instance
func (instance *MDM_Policy_Result01_Privacy02) SetPropertyLetAppsAccessMessaging(value int32) (err error) {
	return instance.SetProperty("LetAppsAccessMessaging", (value))
}

// GetLetAppsAccessMessaging gets the value of LetAppsAccessMessaging for the instance
func (instance *MDM_Policy_Result01_Privacy02) GetPropertyLetAppsAccessMessaging() (value int32, err error) {
	retValue, err := instance.GetProperty("LetAppsAccessMessaging")
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

// SetLetAppsAccessMessaging_ForceAllowTheseApps sets the value of LetAppsAccessMessaging_ForceAllowTheseApps for the instance
func (instance *MDM_Policy_Result01_Privacy02) SetPropertyLetAppsAccessMessaging_ForceAllowTheseApps(value string) (err error) {
	return instance.SetProperty("LetAppsAccessMessaging_ForceAllowTheseApps", (value))
}

// GetLetAppsAccessMessaging_ForceAllowTheseApps gets the value of LetAppsAccessMessaging_ForceAllowTheseApps for the instance
func (instance *MDM_Policy_Result01_Privacy02) GetPropertyLetAppsAccessMessaging_ForceAllowTheseApps() (value string, err error) {
	retValue, err := instance.GetProperty("LetAppsAccessMessaging_ForceAllowTheseApps")
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

// SetLetAppsAccessMessaging_ForceDenyTheseApps sets the value of LetAppsAccessMessaging_ForceDenyTheseApps for the instance
func (instance *MDM_Policy_Result01_Privacy02) SetPropertyLetAppsAccessMessaging_ForceDenyTheseApps(value string) (err error) {
	return instance.SetProperty("LetAppsAccessMessaging_ForceDenyTheseApps", (value))
}

// GetLetAppsAccessMessaging_ForceDenyTheseApps gets the value of LetAppsAccessMessaging_ForceDenyTheseApps for the instance
func (instance *MDM_Policy_Result01_Privacy02) GetPropertyLetAppsAccessMessaging_ForceDenyTheseApps() (value string, err error) {
	retValue, err := instance.GetProperty("LetAppsAccessMessaging_ForceDenyTheseApps")
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

// SetLetAppsAccessMessaging_UserInControlOfTheseApps sets the value of LetAppsAccessMessaging_UserInControlOfTheseApps for the instance
func (instance *MDM_Policy_Result01_Privacy02) SetPropertyLetAppsAccessMessaging_UserInControlOfTheseApps(value string) (err error) {
	return instance.SetProperty("LetAppsAccessMessaging_UserInControlOfTheseApps", (value))
}

// GetLetAppsAccessMessaging_UserInControlOfTheseApps gets the value of LetAppsAccessMessaging_UserInControlOfTheseApps for the instance
func (instance *MDM_Policy_Result01_Privacy02) GetPropertyLetAppsAccessMessaging_UserInControlOfTheseApps() (value string, err error) {
	retValue, err := instance.GetProperty("LetAppsAccessMessaging_UserInControlOfTheseApps")
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

// SetLetAppsAccessMicrophone sets the value of LetAppsAccessMicrophone for the instance
func (instance *MDM_Policy_Result01_Privacy02) SetPropertyLetAppsAccessMicrophone(value int32) (err error) {
	return instance.SetProperty("LetAppsAccessMicrophone", (value))
}

// GetLetAppsAccessMicrophone gets the value of LetAppsAccessMicrophone for the instance
func (instance *MDM_Policy_Result01_Privacy02) GetPropertyLetAppsAccessMicrophone() (value int32, err error) {
	retValue, err := instance.GetProperty("LetAppsAccessMicrophone")
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

// SetLetAppsAccessMicrophone_ForceAllowTheseApps sets the value of LetAppsAccessMicrophone_ForceAllowTheseApps for the instance
func (instance *MDM_Policy_Result01_Privacy02) SetPropertyLetAppsAccessMicrophone_ForceAllowTheseApps(value string) (err error) {
	return instance.SetProperty("LetAppsAccessMicrophone_ForceAllowTheseApps", (value))
}

// GetLetAppsAccessMicrophone_ForceAllowTheseApps gets the value of LetAppsAccessMicrophone_ForceAllowTheseApps for the instance
func (instance *MDM_Policy_Result01_Privacy02) GetPropertyLetAppsAccessMicrophone_ForceAllowTheseApps() (value string, err error) {
	retValue, err := instance.GetProperty("LetAppsAccessMicrophone_ForceAllowTheseApps")
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

// SetLetAppsAccessMicrophone_ForceDenyTheseApps sets the value of LetAppsAccessMicrophone_ForceDenyTheseApps for the instance
func (instance *MDM_Policy_Result01_Privacy02) SetPropertyLetAppsAccessMicrophone_ForceDenyTheseApps(value string) (err error) {
	return instance.SetProperty("LetAppsAccessMicrophone_ForceDenyTheseApps", (value))
}

// GetLetAppsAccessMicrophone_ForceDenyTheseApps gets the value of LetAppsAccessMicrophone_ForceDenyTheseApps for the instance
func (instance *MDM_Policy_Result01_Privacy02) GetPropertyLetAppsAccessMicrophone_ForceDenyTheseApps() (value string, err error) {
	retValue, err := instance.GetProperty("LetAppsAccessMicrophone_ForceDenyTheseApps")
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

// SetLetAppsAccessMicrophone_UserInControlOfTheseApps sets the value of LetAppsAccessMicrophone_UserInControlOfTheseApps for the instance
func (instance *MDM_Policy_Result01_Privacy02) SetPropertyLetAppsAccessMicrophone_UserInControlOfTheseApps(value string) (err error) {
	return instance.SetProperty("LetAppsAccessMicrophone_UserInControlOfTheseApps", (value))
}

// GetLetAppsAccessMicrophone_UserInControlOfTheseApps gets the value of LetAppsAccessMicrophone_UserInControlOfTheseApps for the instance
func (instance *MDM_Policy_Result01_Privacy02) GetPropertyLetAppsAccessMicrophone_UserInControlOfTheseApps() (value string, err error) {
	retValue, err := instance.GetProperty("LetAppsAccessMicrophone_UserInControlOfTheseApps")
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

// SetLetAppsAccessMotion sets the value of LetAppsAccessMotion for the instance
func (instance *MDM_Policy_Result01_Privacy02) SetPropertyLetAppsAccessMotion(value int32) (err error) {
	return instance.SetProperty("LetAppsAccessMotion", (value))
}

// GetLetAppsAccessMotion gets the value of LetAppsAccessMotion for the instance
func (instance *MDM_Policy_Result01_Privacy02) GetPropertyLetAppsAccessMotion() (value int32, err error) {
	retValue, err := instance.GetProperty("LetAppsAccessMotion")
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

// SetLetAppsAccessMotion_ForceAllowTheseApps sets the value of LetAppsAccessMotion_ForceAllowTheseApps for the instance
func (instance *MDM_Policy_Result01_Privacy02) SetPropertyLetAppsAccessMotion_ForceAllowTheseApps(value string) (err error) {
	return instance.SetProperty("LetAppsAccessMotion_ForceAllowTheseApps", (value))
}

// GetLetAppsAccessMotion_ForceAllowTheseApps gets the value of LetAppsAccessMotion_ForceAllowTheseApps for the instance
func (instance *MDM_Policy_Result01_Privacy02) GetPropertyLetAppsAccessMotion_ForceAllowTheseApps() (value string, err error) {
	retValue, err := instance.GetProperty("LetAppsAccessMotion_ForceAllowTheseApps")
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

// SetLetAppsAccessMotion_ForceDenyTheseApps sets the value of LetAppsAccessMotion_ForceDenyTheseApps for the instance
func (instance *MDM_Policy_Result01_Privacy02) SetPropertyLetAppsAccessMotion_ForceDenyTheseApps(value string) (err error) {
	return instance.SetProperty("LetAppsAccessMotion_ForceDenyTheseApps", (value))
}

// GetLetAppsAccessMotion_ForceDenyTheseApps gets the value of LetAppsAccessMotion_ForceDenyTheseApps for the instance
func (instance *MDM_Policy_Result01_Privacy02) GetPropertyLetAppsAccessMotion_ForceDenyTheseApps() (value string, err error) {
	retValue, err := instance.GetProperty("LetAppsAccessMotion_ForceDenyTheseApps")
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

// SetLetAppsAccessMotion_UserInControlOfTheseApps sets the value of LetAppsAccessMotion_UserInControlOfTheseApps for the instance
func (instance *MDM_Policy_Result01_Privacy02) SetPropertyLetAppsAccessMotion_UserInControlOfTheseApps(value string) (err error) {
	return instance.SetProperty("LetAppsAccessMotion_UserInControlOfTheseApps", (value))
}

// GetLetAppsAccessMotion_UserInControlOfTheseApps gets the value of LetAppsAccessMotion_UserInControlOfTheseApps for the instance
func (instance *MDM_Policy_Result01_Privacy02) GetPropertyLetAppsAccessMotion_UserInControlOfTheseApps() (value string, err error) {
	retValue, err := instance.GetProperty("LetAppsAccessMotion_UserInControlOfTheseApps")
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

// SetLetAppsAccessNotifications sets the value of LetAppsAccessNotifications for the instance
func (instance *MDM_Policy_Result01_Privacy02) SetPropertyLetAppsAccessNotifications(value int32) (err error) {
	return instance.SetProperty("LetAppsAccessNotifications", (value))
}

// GetLetAppsAccessNotifications gets the value of LetAppsAccessNotifications for the instance
func (instance *MDM_Policy_Result01_Privacy02) GetPropertyLetAppsAccessNotifications() (value int32, err error) {
	retValue, err := instance.GetProperty("LetAppsAccessNotifications")
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

// SetLetAppsAccessNotifications_ForceAllowTheseApps sets the value of LetAppsAccessNotifications_ForceAllowTheseApps for the instance
func (instance *MDM_Policy_Result01_Privacy02) SetPropertyLetAppsAccessNotifications_ForceAllowTheseApps(value string) (err error) {
	return instance.SetProperty("LetAppsAccessNotifications_ForceAllowTheseApps", (value))
}

// GetLetAppsAccessNotifications_ForceAllowTheseApps gets the value of LetAppsAccessNotifications_ForceAllowTheseApps for the instance
func (instance *MDM_Policy_Result01_Privacy02) GetPropertyLetAppsAccessNotifications_ForceAllowTheseApps() (value string, err error) {
	retValue, err := instance.GetProperty("LetAppsAccessNotifications_ForceAllowTheseApps")
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

// SetLetAppsAccessNotifications_ForceDenyTheseApps sets the value of LetAppsAccessNotifications_ForceDenyTheseApps for the instance
func (instance *MDM_Policy_Result01_Privacy02) SetPropertyLetAppsAccessNotifications_ForceDenyTheseApps(value string) (err error) {
	return instance.SetProperty("LetAppsAccessNotifications_ForceDenyTheseApps", (value))
}

// GetLetAppsAccessNotifications_ForceDenyTheseApps gets the value of LetAppsAccessNotifications_ForceDenyTheseApps for the instance
func (instance *MDM_Policy_Result01_Privacy02) GetPropertyLetAppsAccessNotifications_ForceDenyTheseApps() (value string, err error) {
	retValue, err := instance.GetProperty("LetAppsAccessNotifications_ForceDenyTheseApps")
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

// SetLetAppsAccessNotifications_UserInControlOfTheseApps sets the value of LetAppsAccessNotifications_UserInControlOfTheseApps for the instance
func (instance *MDM_Policy_Result01_Privacy02) SetPropertyLetAppsAccessNotifications_UserInControlOfTheseApps(value string) (err error) {
	return instance.SetProperty("LetAppsAccessNotifications_UserInControlOfTheseApps", (value))
}

// GetLetAppsAccessNotifications_UserInControlOfTheseApps gets the value of LetAppsAccessNotifications_UserInControlOfTheseApps for the instance
func (instance *MDM_Policy_Result01_Privacy02) GetPropertyLetAppsAccessNotifications_UserInControlOfTheseApps() (value string, err error) {
	retValue, err := instance.GetProperty("LetAppsAccessNotifications_UserInControlOfTheseApps")
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

// SetLetAppsAccessPhone sets the value of LetAppsAccessPhone for the instance
func (instance *MDM_Policy_Result01_Privacy02) SetPropertyLetAppsAccessPhone(value int32) (err error) {
	return instance.SetProperty("LetAppsAccessPhone", (value))
}

// GetLetAppsAccessPhone gets the value of LetAppsAccessPhone for the instance
func (instance *MDM_Policy_Result01_Privacy02) GetPropertyLetAppsAccessPhone() (value int32, err error) {
	retValue, err := instance.GetProperty("LetAppsAccessPhone")
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

// SetLetAppsAccessPhone_ForceAllowTheseApps sets the value of LetAppsAccessPhone_ForceAllowTheseApps for the instance
func (instance *MDM_Policy_Result01_Privacy02) SetPropertyLetAppsAccessPhone_ForceAllowTheseApps(value string) (err error) {
	return instance.SetProperty("LetAppsAccessPhone_ForceAllowTheseApps", (value))
}

// GetLetAppsAccessPhone_ForceAllowTheseApps gets the value of LetAppsAccessPhone_ForceAllowTheseApps for the instance
func (instance *MDM_Policy_Result01_Privacy02) GetPropertyLetAppsAccessPhone_ForceAllowTheseApps() (value string, err error) {
	retValue, err := instance.GetProperty("LetAppsAccessPhone_ForceAllowTheseApps")
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

// SetLetAppsAccessPhone_ForceDenyTheseApps sets the value of LetAppsAccessPhone_ForceDenyTheseApps for the instance
func (instance *MDM_Policy_Result01_Privacy02) SetPropertyLetAppsAccessPhone_ForceDenyTheseApps(value string) (err error) {
	return instance.SetProperty("LetAppsAccessPhone_ForceDenyTheseApps", (value))
}

// GetLetAppsAccessPhone_ForceDenyTheseApps gets the value of LetAppsAccessPhone_ForceDenyTheseApps for the instance
func (instance *MDM_Policy_Result01_Privacy02) GetPropertyLetAppsAccessPhone_ForceDenyTheseApps() (value string, err error) {
	retValue, err := instance.GetProperty("LetAppsAccessPhone_ForceDenyTheseApps")
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

// SetLetAppsAccessPhone_UserInControlOfTheseApps sets the value of LetAppsAccessPhone_UserInControlOfTheseApps for the instance
func (instance *MDM_Policy_Result01_Privacy02) SetPropertyLetAppsAccessPhone_UserInControlOfTheseApps(value string) (err error) {
	return instance.SetProperty("LetAppsAccessPhone_UserInControlOfTheseApps", (value))
}

// GetLetAppsAccessPhone_UserInControlOfTheseApps gets the value of LetAppsAccessPhone_UserInControlOfTheseApps for the instance
func (instance *MDM_Policy_Result01_Privacy02) GetPropertyLetAppsAccessPhone_UserInControlOfTheseApps() (value string, err error) {
	retValue, err := instance.GetProperty("LetAppsAccessPhone_UserInControlOfTheseApps")
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

// SetLetAppsAccessRadios sets the value of LetAppsAccessRadios for the instance
func (instance *MDM_Policy_Result01_Privacy02) SetPropertyLetAppsAccessRadios(value int32) (err error) {
	return instance.SetProperty("LetAppsAccessRadios", (value))
}

// GetLetAppsAccessRadios gets the value of LetAppsAccessRadios for the instance
func (instance *MDM_Policy_Result01_Privacy02) GetPropertyLetAppsAccessRadios() (value int32, err error) {
	retValue, err := instance.GetProperty("LetAppsAccessRadios")
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

// SetLetAppsAccessRadios_ForceAllowTheseApps sets the value of LetAppsAccessRadios_ForceAllowTheseApps for the instance
func (instance *MDM_Policy_Result01_Privacy02) SetPropertyLetAppsAccessRadios_ForceAllowTheseApps(value string) (err error) {
	return instance.SetProperty("LetAppsAccessRadios_ForceAllowTheseApps", (value))
}

// GetLetAppsAccessRadios_ForceAllowTheseApps gets the value of LetAppsAccessRadios_ForceAllowTheseApps for the instance
func (instance *MDM_Policy_Result01_Privacy02) GetPropertyLetAppsAccessRadios_ForceAllowTheseApps() (value string, err error) {
	retValue, err := instance.GetProperty("LetAppsAccessRadios_ForceAllowTheseApps")
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

// SetLetAppsAccessRadios_ForceDenyTheseApps sets the value of LetAppsAccessRadios_ForceDenyTheseApps for the instance
func (instance *MDM_Policy_Result01_Privacy02) SetPropertyLetAppsAccessRadios_ForceDenyTheseApps(value string) (err error) {
	return instance.SetProperty("LetAppsAccessRadios_ForceDenyTheseApps", (value))
}

// GetLetAppsAccessRadios_ForceDenyTheseApps gets the value of LetAppsAccessRadios_ForceDenyTheseApps for the instance
func (instance *MDM_Policy_Result01_Privacy02) GetPropertyLetAppsAccessRadios_ForceDenyTheseApps() (value string, err error) {
	retValue, err := instance.GetProperty("LetAppsAccessRadios_ForceDenyTheseApps")
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

// SetLetAppsAccessRadios_UserInControlOfTheseApps sets the value of LetAppsAccessRadios_UserInControlOfTheseApps for the instance
func (instance *MDM_Policy_Result01_Privacy02) SetPropertyLetAppsAccessRadios_UserInControlOfTheseApps(value string) (err error) {
	return instance.SetProperty("LetAppsAccessRadios_UserInControlOfTheseApps", (value))
}

// GetLetAppsAccessRadios_UserInControlOfTheseApps gets the value of LetAppsAccessRadios_UserInControlOfTheseApps for the instance
func (instance *MDM_Policy_Result01_Privacy02) GetPropertyLetAppsAccessRadios_UserInControlOfTheseApps() (value string, err error) {
	retValue, err := instance.GetProperty("LetAppsAccessRadios_UserInControlOfTheseApps")
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

// SetLetAppsAccessTasks sets the value of LetAppsAccessTasks for the instance
func (instance *MDM_Policy_Result01_Privacy02) SetPropertyLetAppsAccessTasks(value int32) (err error) {
	return instance.SetProperty("LetAppsAccessTasks", (value))
}

// GetLetAppsAccessTasks gets the value of LetAppsAccessTasks for the instance
func (instance *MDM_Policy_Result01_Privacy02) GetPropertyLetAppsAccessTasks() (value int32, err error) {
	retValue, err := instance.GetProperty("LetAppsAccessTasks")
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

// SetLetAppsAccessTasks_ForceAllowTheseApps sets the value of LetAppsAccessTasks_ForceAllowTheseApps for the instance
func (instance *MDM_Policy_Result01_Privacy02) SetPropertyLetAppsAccessTasks_ForceAllowTheseApps(value string) (err error) {
	return instance.SetProperty("LetAppsAccessTasks_ForceAllowTheseApps", (value))
}

// GetLetAppsAccessTasks_ForceAllowTheseApps gets the value of LetAppsAccessTasks_ForceAllowTheseApps for the instance
func (instance *MDM_Policy_Result01_Privacy02) GetPropertyLetAppsAccessTasks_ForceAllowTheseApps() (value string, err error) {
	retValue, err := instance.GetProperty("LetAppsAccessTasks_ForceAllowTheseApps")
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

// SetLetAppsAccessTasks_ForceDenyTheseApps sets the value of LetAppsAccessTasks_ForceDenyTheseApps for the instance
func (instance *MDM_Policy_Result01_Privacy02) SetPropertyLetAppsAccessTasks_ForceDenyTheseApps(value string) (err error) {
	return instance.SetProperty("LetAppsAccessTasks_ForceDenyTheseApps", (value))
}

// GetLetAppsAccessTasks_ForceDenyTheseApps gets the value of LetAppsAccessTasks_ForceDenyTheseApps for the instance
func (instance *MDM_Policy_Result01_Privacy02) GetPropertyLetAppsAccessTasks_ForceDenyTheseApps() (value string, err error) {
	retValue, err := instance.GetProperty("LetAppsAccessTasks_ForceDenyTheseApps")
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

// SetLetAppsAccessTasks_UserInControlOfTheseApps sets the value of LetAppsAccessTasks_UserInControlOfTheseApps for the instance
func (instance *MDM_Policy_Result01_Privacy02) SetPropertyLetAppsAccessTasks_UserInControlOfTheseApps(value string) (err error) {
	return instance.SetProperty("LetAppsAccessTasks_UserInControlOfTheseApps", (value))
}

// GetLetAppsAccessTasks_UserInControlOfTheseApps gets the value of LetAppsAccessTasks_UserInControlOfTheseApps for the instance
func (instance *MDM_Policy_Result01_Privacy02) GetPropertyLetAppsAccessTasks_UserInControlOfTheseApps() (value string, err error) {
	retValue, err := instance.GetProperty("LetAppsAccessTasks_UserInControlOfTheseApps")
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

// SetLetAppsAccessTrustedDevices sets the value of LetAppsAccessTrustedDevices for the instance
func (instance *MDM_Policy_Result01_Privacy02) SetPropertyLetAppsAccessTrustedDevices(value int32) (err error) {
	return instance.SetProperty("LetAppsAccessTrustedDevices", (value))
}

// GetLetAppsAccessTrustedDevices gets the value of LetAppsAccessTrustedDevices for the instance
func (instance *MDM_Policy_Result01_Privacy02) GetPropertyLetAppsAccessTrustedDevices() (value int32, err error) {
	retValue, err := instance.GetProperty("LetAppsAccessTrustedDevices")
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

// SetLetAppsAccessTrustedDevices_ForceAllowTheseApps sets the value of LetAppsAccessTrustedDevices_ForceAllowTheseApps for the instance
func (instance *MDM_Policy_Result01_Privacy02) SetPropertyLetAppsAccessTrustedDevices_ForceAllowTheseApps(value string) (err error) {
	return instance.SetProperty("LetAppsAccessTrustedDevices_ForceAllowTheseApps", (value))
}

// GetLetAppsAccessTrustedDevices_ForceAllowTheseApps gets the value of LetAppsAccessTrustedDevices_ForceAllowTheseApps for the instance
func (instance *MDM_Policy_Result01_Privacy02) GetPropertyLetAppsAccessTrustedDevices_ForceAllowTheseApps() (value string, err error) {
	retValue, err := instance.GetProperty("LetAppsAccessTrustedDevices_ForceAllowTheseApps")
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

// SetLetAppsAccessTrustedDevices_ForceDenyTheseApps sets the value of LetAppsAccessTrustedDevices_ForceDenyTheseApps for the instance
func (instance *MDM_Policy_Result01_Privacy02) SetPropertyLetAppsAccessTrustedDevices_ForceDenyTheseApps(value string) (err error) {
	return instance.SetProperty("LetAppsAccessTrustedDevices_ForceDenyTheseApps", (value))
}

// GetLetAppsAccessTrustedDevices_ForceDenyTheseApps gets the value of LetAppsAccessTrustedDevices_ForceDenyTheseApps for the instance
func (instance *MDM_Policy_Result01_Privacy02) GetPropertyLetAppsAccessTrustedDevices_ForceDenyTheseApps() (value string, err error) {
	retValue, err := instance.GetProperty("LetAppsAccessTrustedDevices_ForceDenyTheseApps")
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

// SetLetAppsAccessTrustedDevices_UserInControlOfTheseApps sets the value of LetAppsAccessTrustedDevices_UserInControlOfTheseApps for the instance
func (instance *MDM_Policy_Result01_Privacy02) SetPropertyLetAppsAccessTrustedDevices_UserInControlOfTheseApps(value string) (err error) {
	return instance.SetProperty("LetAppsAccessTrustedDevices_UserInControlOfTheseApps", (value))
}

// GetLetAppsAccessTrustedDevices_UserInControlOfTheseApps gets the value of LetAppsAccessTrustedDevices_UserInControlOfTheseApps for the instance
func (instance *MDM_Policy_Result01_Privacy02) GetPropertyLetAppsAccessTrustedDevices_UserInControlOfTheseApps() (value string, err error) {
	retValue, err := instance.GetProperty("LetAppsAccessTrustedDevices_UserInControlOfTheseApps")
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

// SetLetAppsActivateWithVoice sets the value of LetAppsActivateWithVoice for the instance
func (instance *MDM_Policy_Result01_Privacy02) SetPropertyLetAppsActivateWithVoice(value int32) (err error) {
	return instance.SetProperty("LetAppsActivateWithVoice", (value))
}

// GetLetAppsActivateWithVoice gets the value of LetAppsActivateWithVoice for the instance
func (instance *MDM_Policy_Result01_Privacy02) GetPropertyLetAppsActivateWithVoice() (value int32, err error) {
	retValue, err := instance.GetProperty("LetAppsActivateWithVoice")
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

// SetLetAppsActivateWithVoiceAboveLock sets the value of LetAppsActivateWithVoiceAboveLock for the instance
func (instance *MDM_Policy_Result01_Privacy02) SetPropertyLetAppsActivateWithVoiceAboveLock(value int32) (err error) {
	return instance.SetProperty("LetAppsActivateWithVoiceAboveLock", (value))
}

// GetLetAppsActivateWithVoiceAboveLock gets the value of LetAppsActivateWithVoiceAboveLock for the instance
func (instance *MDM_Policy_Result01_Privacy02) GetPropertyLetAppsActivateWithVoiceAboveLock() (value int32, err error) {
	retValue, err := instance.GetProperty("LetAppsActivateWithVoiceAboveLock")
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

// SetLetAppsGetDiagnosticInfo sets the value of LetAppsGetDiagnosticInfo for the instance
func (instance *MDM_Policy_Result01_Privacy02) SetPropertyLetAppsGetDiagnosticInfo(value int32) (err error) {
	return instance.SetProperty("LetAppsGetDiagnosticInfo", (value))
}

// GetLetAppsGetDiagnosticInfo gets the value of LetAppsGetDiagnosticInfo for the instance
func (instance *MDM_Policy_Result01_Privacy02) GetPropertyLetAppsGetDiagnosticInfo() (value int32, err error) {
	retValue, err := instance.GetProperty("LetAppsGetDiagnosticInfo")
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

// SetLetAppsGetDiagnosticInfo_ForceAllowTheseApps sets the value of LetAppsGetDiagnosticInfo_ForceAllowTheseApps for the instance
func (instance *MDM_Policy_Result01_Privacy02) SetPropertyLetAppsGetDiagnosticInfo_ForceAllowTheseApps(value string) (err error) {
	return instance.SetProperty("LetAppsGetDiagnosticInfo_ForceAllowTheseApps", (value))
}

// GetLetAppsGetDiagnosticInfo_ForceAllowTheseApps gets the value of LetAppsGetDiagnosticInfo_ForceAllowTheseApps for the instance
func (instance *MDM_Policy_Result01_Privacy02) GetPropertyLetAppsGetDiagnosticInfo_ForceAllowTheseApps() (value string, err error) {
	retValue, err := instance.GetProperty("LetAppsGetDiagnosticInfo_ForceAllowTheseApps")
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

// SetLetAppsGetDiagnosticInfo_ForceDenyTheseApps sets the value of LetAppsGetDiagnosticInfo_ForceDenyTheseApps for the instance
func (instance *MDM_Policy_Result01_Privacy02) SetPropertyLetAppsGetDiagnosticInfo_ForceDenyTheseApps(value string) (err error) {
	return instance.SetProperty("LetAppsGetDiagnosticInfo_ForceDenyTheseApps", (value))
}

// GetLetAppsGetDiagnosticInfo_ForceDenyTheseApps gets the value of LetAppsGetDiagnosticInfo_ForceDenyTheseApps for the instance
func (instance *MDM_Policy_Result01_Privacy02) GetPropertyLetAppsGetDiagnosticInfo_ForceDenyTheseApps() (value string, err error) {
	retValue, err := instance.GetProperty("LetAppsGetDiagnosticInfo_ForceDenyTheseApps")
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

// SetLetAppsGetDiagnosticInfo_UserInControlOfTheseApps sets the value of LetAppsGetDiagnosticInfo_UserInControlOfTheseApps for the instance
func (instance *MDM_Policy_Result01_Privacy02) SetPropertyLetAppsGetDiagnosticInfo_UserInControlOfTheseApps(value string) (err error) {
	return instance.SetProperty("LetAppsGetDiagnosticInfo_UserInControlOfTheseApps", (value))
}

// GetLetAppsGetDiagnosticInfo_UserInControlOfTheseApps gets the value of LetAppsGetDiagnosticInfo_UserInControlOfTheseApps for the instance
func (instance *MDM_Policy_Result01_Privacy02) GetPropertyLetAppsGetDiagnosticInfo_UserInControlOfTheseApps() (value string, err error) {
	retValue, err := instance.GetProperty("LetAppsGetDiagnosticInfo_UserInControlOfTheseApps")
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

// SetLetAppsRunInBackground sets the value of LetAppsRunInBackground for the instance
func (instance *MDM_Policy_Result01_Privacy02) SetPropertyLetAppsRunInBackground(value int32) (err error) {
	return instance.SetProperty("LetAppsRunInBackground", (value))
}

// GetLetAppsRunInBackground gets the value of LetAppsRunInBackground for the instance
func (instance *MDM_Policy_Result01_Privacy02) GetPropertyLetAppsRunInBackground() (value int32, err error) {
	retValue, err := instance.GetProperty("LetAppsRunInBackground")
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

// SetLetAppsRunInBackground_ForceAllowTheseApps sets the value of LetAppsRunInBackground_ForceAllowTheseApps for the instance
func (instance *MDM_Policy_Result01_Privacy02) SetPropertyLetAppsRunInBackground_ForceAllowTheseApps(value string) (err error) {
	return instance.SetProperty("LetAppsRunInBackground_ForceAllowTheseApps", (value))
}

// GetLetAppsRunInBackground_ForceAllowTheseApps gets the value of LetAppsRunInBackground_ForceAllowTheseApps for the instance
func (instance *MDM_Policy_Result01_Privacy02) GetPropertyLetAppsRunInBackground_ForceAllowTheseApps() (value string, err error) {
	retValue, err := instance.GetProperty("LetAppsRunInBackground_ForceAllowTheseApps")
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

// SetLetAppsRunInBackground_ForceDenyTheseApps sets the value of LetAppsRunInBackground_ForceDenyTheseApps for the instance
func (instance *MDM_Policy_Result01_Privacy02) SetPropertyLetAppsRunInBackground_ForceDenyTheseApps(value string) (err error) {
	return instance.SetProperty("LetAppsRunInBackground_ForceDenyTheseApps", (value))
}

// GetLetAppsRunInBackground_ForceDenyTheseApps gets the value of LetAppsRunInBackground_ForceDenyTheseApps for the instance
func (instance *MDM_Policy_Result01_Privacy02) GetPropertyLetAppsRunInBackground_ForceDenyTheseApps() (value string, err error) {
	retValue, err := instance.GetProperty("LetAppsRunInBackground_ForceDenyTheseApps")
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

// SetLetAppsRunInBackground_UserInControlOfTheseApps sets the value of LetAppsRunInBackground_UserInControlOfTheseApps for the instance
func (instance *MDM_Policy_Result01_Privacy02) SetPropertyLetAppsRunInBackground_UserInControlOfTheseApps(value string) (err error) {
	return instance.SetProperty("LetAppsRunInBackground_UserInControlOfTheseApps", (value))
}

// GetLetAppsRunInBackground_UserInControlOfTheseApps gets the value of LetAppsRunInBackground_UserInControlOfTheseApps for the instance
func (instance *MDM_Policy_Result01_Privacy02) GetPropertyLetAppsRunInBackground_UserInControlOfTheseApps() (value string, err error) {
	retValue, err := instance.GetProperty("LetAppsRunInBackground_UserInControlOfTheseApps")
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

// SetLetAppsSyncWithDevices sets the value of LetAppsSyncWithDevices for the instance
func (instance *MDM_Policy_Result01_Privacy02) SetPropertyLetAppsSyncWithDevices(value int32) (err error) {
	return instance.SetProperty("LetAppsSyncWithDevices", (value))
}

// GetLetAppsSyncWithDevices gets the value of LetAppsSyncWithDevices for the instance
func (instance *MDM_Policy_Result01_Privacy02) GetPropertyLetAppsSyncWithDevices() (value int32, err error) {
	retValue, err := instance.GetProperty("LetAppsSyncWithDevices")
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

// SetLetAppsSyncWithDevices_ForceAllowTheseApps sets the value of LetAppsSyncWithDevices_ForceAllowTheseApps for the instance
func (instance *MDM_Policy_Result01_Privacy02) SetPropertyLetAppsSyncWithDevices_ForceAllowTheseApps(value string) (err error) {
	return instance.SetProperty("LetAppsSyncWithDevices_ForceAllowTheseApps", (value))
}

// GetLetAppsSyncWithDevices_ForceAllowTheseApps gets the value of LetAppsSyncWithDevices_ForceAllowTheseApps for the instance
func (instance *MDM_Policy_Result01_Privacy02) GetPropertyLetAppsSyncWithDevices_ForceAllowTheseApps() (value string, err error) {
	retValue, err := instance.GetProperty("LetAppsSyncWithDevices_ForceAllowTheseApps")
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

// SetLetAppsSyncWithDevices_ForceDenyTheseApps sets the value of LetAppsSyncWithDevices_ForceDenyTheseApps for the instance
func (instance *MDM_Policy_Result01_Privacy02) SetPropertyLetAppsSyncWithDevices_ForceDenyTheseApps(value string) (err error) {
	return instance.SetProperty("LetAppsSyncWithDevices_ForceDenyTheseApps", (value))
}

// GetLetAppsSyncWithDevices_ForceDenyTheseApps gets the value of LetAppsSyncWithDevices_ForceDenyTheseApps for the instance
func (instance *MDM_Policy_Result01_Privacy02) GetPropertyLetAppsSyncWithDevices_ForceDenyTheseApps() (value string, err error) {
	retValue, err := instance.GetProperty("LetAppsSyncWithDevices_ForceDenyTheseApps")
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

// SetLetAppsSyncWithDevices_UserInControlOfTheseApps sets the value of LetAppsSyncWithDevices_UserInControlOfTheseApps for the instance
func (instance *MDM_Policy_Result01_Privacy02) SetPropertyLetAppsSyncWithDevices_UserInControlOfTheseApps(value string) (err error) {
	return instance.SetProperty("LetAppsSyncWithDevices_UserInControlOfTheseApps", (value))
}

// GetLetAppsSyncWithDevices_UserInControlOfTheseApps gets the value of LetAppsSyncWithDevices_UserInControlOfTheseApps for the instance
func (instance *MDM_Policy_Result01_Privacy02) GetPropertyLetAppsSyncWithDevices_UserInControlOfTheseApps() (value string, err error) {
	retValue, err := instance.GetProperty("LetAppsSyncWithDevices_UserInControlOfTheseApps")
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
func (instance *MDM_Policy_Result01_Privacy02) SetPropertyParentID(value string) (err error) {
	return instance.SetProperty("ParentID", (value))
}

// GetParentID gets the value of ParentID for the instance
func (instance *MDM_Policy_Result01_Privacy02) GetPropertyParentID() (value string, err error) {
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

// SetPublishUserActivities sets the value of PublishUserActivities for the instance
func (instance *MDM_Policy_Result01_Privacy02) SetPropertyPublishUserActivities(value int32) (err error) {
	return instance.SetProperty("PublishUserActivities", (value))
}

// GetPublishUserActivities gets the value of PublishUserActivities for the instance
func (instance *MDM_Policy_Result01_Privacy02) GetPropertyPublishUserActivities() (value int32, err error) {
	retValue, err := instance.GetProperty("PublishUserActivities")
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

// SetUploadUserActivities sets the value of UploadUserActivities for the instance
func (instance *MDM_Policy_Result01_Privacy02) SetPropertyUploadUserActivities(value int32) (err error) {
	return instance.SetProperty("UploadUserActivities", (value))
}

// GetUploadUserActivities gets the value of UploadUserActivities for the instance
func (instance *MDM_Policy_Result01_Privacy02) GetPropertyUploadUserActivities() (value int32, err error) {
	retValue, err := instance.GetProperty("UploadUserActivities")
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
