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

// WebServicesSection struct
type WebServicesSection struct { 
	*ConfigurationSectionWithCollection

	// 
	ConformanceWarnings ConformanceWarningSettings

	// 
	Diagnostics DiagnosticsSettings

	// 
	Protocols WebServicesProtocolSettings

	// 
	ServiceDescriptionFormatExtensionTypes ServiceDescriptionFormatSettings

	// 
	SoapEnvelopeProcessing SoapEnvelopeProcessingInfo

	// 
	SoapExtensionImporterTypes SoapExtensionImporterTypesSettings

	// 
	SoapExtensionReflectorTypes SoapExtensionReflectorTypesSettings

	// 
	SoapExtensionTypes SoapExtensionTypesInfo

	// 
	SoapServerProtocolFactory SoapServerProtocolFactory

	// 
	SoapTransportImporterTypes SoapTransportImporterTypesInfo

	// 
	WsdlHelpGenerator WsdlHelpGeneratorInfo
}

	func NewWebServicesSectionEx1(instance *cim.WmiInstance) (newInstance *WebServicesSection, err error) {tmp, err := NewConfigurationSectionWithCollectionEx1(instance)
		
	if err != nil { return }
	newInstance = &WebServicesSection {
	ConfigurationSectionWithCollection: tmp,
	}
	return
	}
	

	func NewWebServicesSectionEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery ) (newInstance *WebServicesSection, err error) {tmp, err := NewConfigurationSectionWithCollectionEx6(hostName, wmiNamespace, userName, password, domainName, query)
		
	if err != nil { return }
	newInstance = &WebServicesSection {
	ConfigurationSectionWithCollection: tmp,
	}
	return
	}
	

// SetConformanceWarnings sets the value of ConformanceWarnings for the instance
func (instance *WebServicesSection) SetPropertyConformanceWarnings(value ConformanceWarningSettings) (err error) { 
    return instance.SetProperty("ConformanceWarnings", (value))
}

// GetConformanceWarnings gets the value of ConformanceWarnings for the instance
func (instance *WebServicesSection) GetPropertyConformanceWarnings()(value ConformanceWarningSettings, err error) { 
    retValue, err := instance.GetProperty("ConformanceWarnings")
    if err != nil {
        return
    }
    if retValue == nil {
        // Doesn't have any value. Return empty
        return
    }
    
    valuetmp, ok := retValue.(ConformanceWarningSettings); 
    if !ok {
        err = errors.Wrapf(errors.InvalidType, " ConformanceWarningSettings is Invalid. Expected %s", reflect.TypeOf(retValue))
        return  
    }

    value = ConformanceWarningSettings(valuetmp)

    return
}

// SetDiagnostics sets the value of Diagnostics for the instance
func (instance *WebServicesSection) SetPropertyDiagnostics(value DiagnosticsSettings) (err error) { 
    return instance.SetProperty("Diagnostics", (value))
}

// GetDiagnostics gets the value of Diagnostics for the instance
func (instance *WebServicesSection) GetPropertyDiagnostics()(value DiagnosticsSettings, err error) { 
    retValue, err := instance.GetProperty("Diagnostics")
    if err != nil {
        return
    }
    if retValue == nil {
        // Doesn't have any value. Return empty
        return
    }
    
    valuetmp, ok := retValue.(DiagnosticsSettings); 
    if !ok {
        err = errors.Wrapf(errors.InvalidType, " DiagnosticsSettings is Invalid. Expected %s", reflect.TypeOf(retValue))
        return  
    }

    value = DiagnosticsSettings(valuetmp)

    return
}

// SetProtocols sets the value of Protocols for the instance
func (instance *WebServicesSection) SetPropertyProtocols(value WebServicesProtocolSettings) (err error) { 
    return instance.SetProperty("Protocols", (value))
}

// GetProtocols gets the value of Protocols for the instance
func (instance *WebServicesSection) GetPropertyProtocols()(value WebServicesProtocolSettings, err error) { 
    retValue, err := instance.GetProperty("Protocols")
    if err != nil {
        return
    }
    if retValue == nil {
        // Doesn't have any value. Return empty
        return
    }
    
    valuetmp, ok := retValue.(WebServicesProtocolSettings); 
    if !ok {
        err = errors.Wrapf(errors.InvalidType, " WebServicesProtocolSettings is Invalid. Expected %s", reflect.TypeOf(retValue))
        return  
    }

    value = WebServicesProtocolSettings(valuetmp)

    return
}

// SetServiceDescriptionFormatExtensionTypes sets the value of ServiceDescriptionFormatExtensionTypes for the instance
func (instance *WebServicesSection) SetPropertyServiceDescriptionFormatExtensionTypes(value ServiceDescriptionFormatSettings) (err error) { 
    return instance.SetProperty("ServiceDescriptionFormatExtensionTypes", (value))
}

// GetServiceDescriptionFormatExtensionTypes gets the value of ServiceDescriptionFormatExtensionTypes for the instance
func (instance *WebServicesSection) GetPropertyServiceDescriptionFormatExtensionTypes()(value ServiceDescriptionFormatSettings, err error) { 
    retValue, err := instance.GetProperty("ServiceDescriptionFormatExtensionTypes")
    if err != nil {
        return
    }
    if retValue == nil {
        // Doesn't have any value. Return empty
        return
    }
    
    valuetmp, ok := retValue.(ServiceDescriptionFormatSettings); 
    if !ok {
        err = errors.Wrapf(errors.InvalidType, " ServiceDescriptionFormatSettings is Invalid. Expected %s", reflect.TypeOf(retValue))
        return  
    }

    value = ServiceDescriptionFormatSettings(valuetmp)

    return
}

// SetSoapEnvelopeProcessing sets the value of SoapEnvelopeProcessing for the instance
func (instance *WebServicesSection) SetPropertySoapEnvelopeProcessing(value SoapEnvelopeProcessingInfo) (err error) { 
    return instance.SetProperty("SoapEnvelopeProcessing", (value))
}

// GetSoapEnvelopeProcessing gets the value of SoapEnvelopeProcessing for the instance
func (instance *WebServicesSection) GetPropertySoapEnvelopeProcessing()(value SoapEnvelopeProcessingInfo, err error) { 
    retValue, err := instance.GetProperty("SoapEnvelopeProcessing")
    if err != nil {
        return
    }
    if retValue == nil {
        // Doesn't have any value. Return empty
        return
    }
    
    valuetmp, ok := retValue.(SoapEnvelopeProcessingInfo); 
    if !ok {
        err = errors.Wrapf(errors.InvalidType, " SoapEnvelopeProcessingInfo is Invalid. Expected %s", reflect.TypeOf(retValue))
        return  
    }

    value = SoapEnvelopeProcessingInfo(valuetmp)

    return
}

// SetSoapExtensionImporterTypes sets the value of SoapExtensionImporterTypes for the instance
func (instance *WebServicesSection) SetPropertySoapExtensionImporterTypes(value SoapExtensionImporterTypesSettings) (err error) { 
    return instance.SetProperty("SoapExtensionImporterTypes", (value))
}

// GetSoapExtensionImporterTypes gets the value of SoapExtensionImporterTypes for the instance
func (instance *WebServicesSection) GetPropertySoapExtensionImporterTypes()(value SoapExtensionImporterTypesSettings, err error) { 
    retValue, err := instance.GetProperty("SoapExtensionImporterTypes")
    if err != nil {
        return
    }
    if retValue == nil {
        // Doesn't have any value. Return empty
        return
    }
    
    valuetmp, ok := retValue.(SoapExtensionImporterTypesSettings); 
    if !ok {
        err = errors.Wrapf(errors.InvalidType, " SoapExtensionImporterTypesSettings is Invalid. Expected %s", reflect.TypeOf(retValue))
        return  
    }

    value = SoapExtensionImporterTypesSettings(valuetmp)

    return
}

// SetSoapExtensionReflectorTypes sets the value of SoapExtensionReflectorTypes for the instance
func (instance *WebServicesSection) SetPropertySoapExtensionReflectorTypes(value SoapExtensionReflectorTypesSettings) (err error) { 
    return instance.SetProperty("SoapExtensionReflectorTypes", (value))
}

// GetSoapExtensionReflectorTypes gets the value of SoapExtensionReflectorTypes for the instance
func (instance *WebServicesSection) GetPropertySoapExtensionReflectorTypes()(value SoapExtensionReflectorTypesSettings, err error) { 
    retValue, err := instance.GetProperty("SoapExtensionReflectorTypes")
    if err != nil {
        return
    }
    if retValue == nil {
        // Doesn't have any value. Return empty
        return
    }
    
    valuetmp, ok := retValue.(SoapExtensionReflectorTypesSettings); 
    if !ok {
        err = errors.Wrapf(errors.InvalidType, " SoapExtensionReflectorTypesSettings is Invalid. Expected %s", reflect.TypeOf(retValue))
        return  
    }

    value = SoapExtensionReflectorTypesSettings(valuetmp)

    return
}

// SetSoapExtensionTypes sets the value of SoapExtensionTypes for the instance
func (instance *WebServicesSection) SetPropertySoapExtensionTypes(value SoapExtensionTypesInfo) (err error) { 
    return instance.SetProperty("SoapExtensionTypes", (value))
}

// GetSoapExtensionTypes gets the value of SoapExtensionTypes for the instance
func (instance *WebServicesSection) GetPropertySoapExtensionTypes()(value SoapExtensionTypesInfo, err error) { 
    retValue, err := instance.GetProperty("SoapExtensionTypes")
    if err != nil {
        return
    }
    if retValue == nil {
        // Doesn't have any value. Return empty
        return
    }
    
    valuetmp, ok := retValue.(SoapExtensionTypesInfo); 
    if !ok {
        err = errors.Wrapf(errors.InvalidType, " SoapExtensionTypesInfo is Invalid. Expected %s", reflect.TypeOf(retValue))
        return  
    }

    value = SoapExtensionTypesInfo(valuetmp)

    return
}

// SetSoapServerProtocolFactory sets the value of SoapServerProtocolFactory for the instance
func (instance *WebServicesSection) SetPropertySoapServerProtocolFactory(value SoapServerProtocolFactory) (err error) { 
    return instance.SetProperty("SoapServerProtocolFactory", (value))
}

// GetSoapServerProtocolFactory gets the value of SoapServerProtocolFactory for the instance
func (instance *WebServicesSection) GetPropertySoapServerProtocolFactory()(value SoapServerProtocolFactory, err error) { 
    retValue, err := instance.GetProperty("SoapServerProtocolFactory")
    if err != nil {
        return
    }
    if retValue == nil {
        // Doesn't have any value. Return empty
        return
    }
    
    valuetmp, ok := retValue.(SoapServerProtocolFactory); 
    if !ok {
        err = errors.Wrapf(errors.InvalidType, " SoapServerProtocolFactory is Invalid. Expected %s", reflect.TypeOf(retValue))
        return  
    }

    value = SoapServerProtocolFactory(valuetmp)

    return
}

// SetSoapTransportImporterTypes sets the value of SoapTransportImporterTypes for the instance
func (instance *WebServicesSection) SetPropertySoapTransportImporterTypes(value SoapTransportImporterTypesInfo) (err error) { 
    return instance.SetProperty("SoapTransportImporterTypes", (value))
}

// GetSoapTransportImporterTypes gets the value of SoapTransportImporterTypes for the instance
func (instance *WebServicesSection) GetPropertySoapTransportImporterTypes()(value SoapTransportImporterTypesInfo, err error) { 
    retValue, err := instance.GetProperty("SoapTransportImporterTypes")
    if err != nil {
        return
    }
    if retValue == nil {
        // Doesn't have any value. Return empty
        return
    }
    
    valuetmp, ok := retValue.(SoapTransportImporterTypesInfo); 
    if !ok {
        err = errors.Wrapf(errors.InvalidType, " SoapTransportImporterTypesInfo is Invalid. Expected %s", reflect.TypeOf(retValue))
        return  
    }

    value = SoapTransportImporterTypesInfo(valuetmp)

    return
}

// SetWsdlHelpGenerator sets the value of WsdlHelpGenerator for the instance
func (instance *WebServicesSection) SetPropertyWsdlHelpGenerator(value WsdlHelpGeneratorInfo) (err error) { 
    return instance.SetProperty("WsdlHelpGenerator", (value))
}

// GetWsdlHelpGenerator gets the value of WsdlHelpGenerator for the instance
func (instance *WebServicesSection) GetPropertyWsdlHelpGenerator()(value WsdlHelpGeneratorInfo, err error) { 
    retValue, err := instance.GetProperty("WsdlHelpGenerator")
    if err != nil {
        return
    }
    if retValue == nil {
        // Doesn't have any value. Return empty
        return
    }
    
    valuetmp, ok := retValue.(WsdlHelpGeneratorInfo); 
    if !ok {
        err = errors.Wrapf(errors.InvalidType, " WsdlHelpGeneratorInfo is Invalid. Expected %s", reflect.TypeOf(retValue))
        return  
    }

    value = WsdlHelpGeneratorInfo(valuetmp)

    return
}

