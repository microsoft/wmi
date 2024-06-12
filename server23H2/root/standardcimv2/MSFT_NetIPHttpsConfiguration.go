// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 6/12/2024 using wmigen
//      Source root.StandardCimv2
//////////////////////////////////////////////
package standardcimv2

import (
	"github.com/microsoft/wmi/pkg/base/query"
	"github.com/microsoft/wmi/pkg/errors"
	cim "github.com/microsoft/wmi/pkg/wmiinstance"
	"reflect"
)

// MSFT_NetIPHttpsConfiguration struct
type MSFT_NetIPHttpsConfiguration struct {
	*MSFT_NetSettingData

	//
	AuthMode uint32

	//
	ConfigurationType uint32

	//
	PolicyStore string

	//
	Profile string

	//
	ProfileActivated bool

	//
	ServerURL string

	//
	State uint32

	//
	StrongCRLRequired bool

	//
	Type uint32
}

func NewMSFT_NetIPHttpsConfigurationEx1(instance *cim.WmiInstance) (newInstance *MSFT_NetIPHttpsConfiguration, err error) {
	tmp, err := NewMSFT_NetSettingDataEx1(instance)

	if err != nil {
		return
	}
	newInstance = &MSFT_NetIPHttpsConfiguration{
		MSFT_NetSettingData: tmp,
	}
	return
}

func NewMSFT_NetIPHttpsConfigurationEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *MSFT_NetIPHttpsConfiguration, err error) {
	tmp, err := NewMSFT_NetSettingDataEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &MSFT_NetIPHttpsConfiguration{
		MSFT_NetSettingData: tmp,
	}
	return
}

// SetAuthMode sets the value of AuthMode for the instance
func (instance *MSFT_NetIPHttpsConfiguration) SetPropertyAuthMode(value uint32) (err error) {
	return instance.SetProperty("AuthMode", (value))
}

// GetAuthMode gets the value of AuthMode for the instance
func (instance *MSFT_NetIPHttpsConfiguration) GetPropertyAuthMode() (value uint32, err error) {
	retValue, err := instance.GetProperty("AuthMode")
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

// SetConfigurationType sets the value of ConfigurationType for the instance
func (instance *MSFT_NetIPHttpsConfiguration) SetPropertyConfigurationType(value uint32) (err error) {
	return instance.SetProperty("ConfigurationType", (value))
}

// GetConfigurationType gets the value of ConfigurationType for the instance
func (instance *MSFT_NetIPHttpsConfiguration) GetPropertyConfigurationType() (value uint32, err error) {
	retValue, err := instance.GetProperty("ConfigurationType")
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

// SetPolicyStore sets the value of PolicyStore for the instance
func (instance *MSFT_NetIPHttpsConfiguration) SetPropertyPolicyStore(value string) (err error) {
	return instance.SetProperty("PolicyStore", (value))
}

// GetPolicyStore gets the value of PolicyStore for the instance
func (instance *MSFT_NetIPHttpsConfiguration) GetPropertyPolicyStore() (value string, err error) {
	retValue, err := instance.GetProperty("PolicyStore")
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

// SetProfile sets the value of Profile for the instance
func (instance *MSFT_NetIPHttpsConfiguration) SetPropertyProfile(value string) (err error) {
	return instance.SetProperty("Profile", (value))
}

// GetProfile gets the value of Profile for the instance
func (instance *MSFT_NetIPHttpsConfiguration) GetPropertyProfile() (value string, err error) {
	retValue, err := instance.GetProperty("Profile")
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

// SetProfileActivated sets the value of ProfileActivated for the instance
func (instance *MSFT_NetIPHttpsConfiguration) SetPropertyProfileActivated(value bool) (err error) {
	return instance.SetProperty("ProfileActivated", (value))
}

// GetProfileActivated gets the value of ProfileActivated for the instance
func (instance *MSFT_NetIPHttpsConfiguration) GetPropertyProfileActivated() (value bool, err error) {
	retValue, err := instance.GetProperty("ProfileActivated")
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

// SetServerURL sets the value of ServerURL for the instance
func (instance *MSFT_NetIPHttpsConfiguration) SetPropertyServerURL(value string) (err error) {
	return instance.SetProperty("ServerURL", (value))
}

// GetServerURL gets the value of ServerURL for the instance
func (instance *MSFT_NetIPHttpsConfiguration) GetPropertyServerURL() (value string, err error) {
	retValue, err := instance.GetProperty("ServerURL")
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

// SetState sets the value of State for the instance
func (instance *MSFT_NetIPHttpsConfiguration) SetPropertyState(value uint32) (err error) {
	return instance.SetProperty("State", (value))
}

// GetState gets the value of State for the instance
func (instance *MSFT_NetIPHttpsConfiguration) GetPropertyState() (value uint32, err error) {
	retValue, err := instance.GetProperty("State")
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

// SetStrongCRLRequired sets the value of StrongCRLRequired for the instance
func (instance *MSFT_NetIPHttpsConfiguration) SetPropertyStrongCRLRequired(value bool) (err error) {
	return instance.SetProperty("StrongCRLRequired", (value))
}

// GetStrongCRLRequired gets the value of StrongCRLRequired for the instance
func (instance *MSFT_NetIPHttpsConfiguration) GetPropertyStrongCRLRequired() (value bool, err error) {
	retValue, err := instance.GetProperty("StrongCRLRequired")
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

// SetType sets the value of Type for the instance
func (instance *MSFT_NetIPHttpsConfiguration) SetPropertyType(value uint32) (err error) {
	return instance.SetProperty("Type", (value))
}

// GetType gets the value of Type for the instance
func (instance *MSFT_NetIPHttpsConfiguration) GetPropertyType() (value uint32, err error) {
	retValue, err := instance.GetProperty("Type")
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

//

// <param name="Profile" type="string "></param>

// <param name="ReturnValue" type="uint32 "></param>
func (instance *MSFT_NetIPHttpsConfiguration) EnableProfile( /* IN */ Profile string) (result uint32, err error) {
	retVal, err := instance.InvokeMethodWithReturn("EnableProfile", Profile)
	if err != nil {
		return
	}
	result = uint32(retVal)
	return

}

//

// <param name="ReturnValue" type="uint32 "></param>
func (instance *MSFT_NetIPHttpsConfiguration) DisableProfile() (result uint32, err error) {
	retVal, err := instance.InvokeMethodWithReturn("DisableProfile")
	if err != nil {
		return
	}
	result = uint32(retVal)
	return

}

//

// <param name="ApplicationId" type="string "></param>
// <param name="CertificateHash" type="string "></param>
// <param name="CertificateStoreName" type="string "></param>
// <param name="IpPort" type="string "></param>
// <param name="NullEncryption" type="bool "></param>

// <param name="ReturnValue" type="uint32 "></param>
func (instance *MSFT_NetIPHttpsConfiguration) AddCertBinding( /* IN */ CertificateHash string,
	/* IN */ ApplicationId string,
	/* IN */ IpPort string,
	/* IN */ CertificateStoreName string,
	/* IN */ NullEncryption bool) (result uint32, err error) {
	retVal, err := instance.InvokeMethodWithReturn("AddCertBinding", CertificateHash, ApplicationId, IpPort, CertificateStoreName, NullEncryption)
	if err != nil {
		return
	}
	result = uint32(retVal)
	return

}

//

// <param name="ReturnValue" type="uint32 "></param>
func (instance *MSFT_NetIPHttpsConfiguration) RemoveCertBinding() (result uint32, err error) {
	retVal, err := instance.InvokeMethodWithReturn("RemoveCertBinding")
	if err != nil {
		return
	}
	result = uint32(retVal)
	return

}

//

// <param name="NewName" type="string "></param>
// <param name="PassThru" type="bool "></param>

// <param name="OutputObject" type="MSFT_NetIPHttpsConfiguration "></param>
// <param name="ReturnValue" type="uint32 "></param>
func (instance *MSFT_NetIPHttpsConfiguration) Rename( /* IN */ NewName string,
	/* IN */ PassThru bool,
	/* OUT */ OutputObject MSFT_NetIPHttpsConfiguration) (result uint32, err error) {
	retVal, err := instance.InvokeMethod("Rename", NewName, PassThru)
	if err != nil {
		return
	}
	retValue := retVal[0].(int32)
	result = uint32(retValue)
	return

}

//

// <param name="AuthMode" type="bool "></param>
// <param name="PassThru" type="bool "></param>
// <param name="State" type="bool "></param>
// <param name="StrongCRLRequired" type="bool "></param>

// <param name="OutputObject" type="MSFT_NetIPHttpsConfiguration "></param>
// <param name="ReturnValue" type="uint32 "></param>
func (instance *MSFT_NetIPHttpsConfiguration) Reset( /* IN */ State bool,
	/* IN */ AuthMode bool,
	/* IN */ StrongCRLRequired bool,
	/* IN */ PassThru bool,
	/* OUT */ OutputObject MSFT_NetIPHttpsConfiguration) (result uint32, err error) {
	retVal, err := instance.InvokeMethod("Reset", State, AuthMode, StrongCRLRequired, PassThru)
	if err != nil {
		return
	}
	retValue := retVal[0].(int32)
	result = uint32(retValue)
	return

}
