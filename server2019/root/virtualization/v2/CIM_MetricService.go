// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

// 
// Author:
//      Auto Generated on 6/6/2024 using wmigen
//      Source root.virtualization.v2
//////////////////////////////////////////////
package v2
import (
 "github.com/microsoft/wmi/pkg/base/query"
cim "github.com/microsoft/wmi/pkg/wmiinstance"
)

// CIM_MetricService struct
type CIM_MetricService struct { 
	*CIM_Service
}

	func NewCIM_MetricServiceEx1(instance *cim.WmiInstance) (newInstance *CIM_MetricService, err error) {tmp, err := NewCIM_ServiceEx1(instance)
		
	if err != nil { return }
	newInstance = &CIM_MetricService {
	CIM_Service: tmp,
	}
	return
	}
	

	func NewCIM_MetricServiceEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery ) (newInstance *CIM_MetricService, err error) {tmp, err := NewCIM_ServiceEx6(hostName, wmiNamespace, userName, password, domainName, query)
		
	if err != nil { return }
	newInstance = &CIM_MetricService {
	CIM_Service: tmp,
	}
	return
	}
	

// ShowMetrics reports the Metrics available to be collected for a ManagedElement, the ManagedElements for which a metric defined by an instance of CIM_BaseMetricDefinition is available to be collected, and whether or not a particular metric is currently being collected for a ManagedElement. 
///If the Subject parameter is specified and the Definition parameter is NULL, upon successful completion of the method, the DefinitionList[] parameter shall contain a reference to an instance of CIM_BaseMetricDefinition for each instance of CIM_BaseMetricDefinition to which the instance of CIM_ManagedElement identified by the Subject parameter is associated through CIM_MetricDefForME where the CIM_BaseMetricDefinition instance is associated to the CIM_MetricService instance through CIM_ServiceAffectsElement, the MetricCollectionEnabled parameter shall contain a value corresponding to the value of the MetricCollectionEnabled property of the instance of CIM_MetricDefForME that associates the CIM_ManagedElement identified by the Subject parameter with the CIM_BaseMetricDefinition for which a reference is returned in the DefinitionList parameter at the same array index.
///If the Definition parameter is non-NULL and the Subject parameter is NULL, upon successful completion of the method the ManagedElements parameter shall contain a reference to each CIM_ManagedElement instance to which the CIM_BaseMetricDefinition instance referenced by the Definition parameteris associated through CIM_MetricDefForME and the MetricCollectionEnabled parameter shall contain a value corresponding to the value of the MetricCollectionEnabled property of the instance of CIM_MetricDefForME that associates the CIM_BaseMetricDefinition identified by the Definition parameter with the CIM_ManagedElement for which a reference is returned in the ManagedElements parameter at the same array index as the reference to the CIM_ManagedElement.
///If the Subject parameter and Definition parameter are both non-NULL, the method shall return NULL values for the DefinitionList and ManagedElements parameters, and the MetricCollectionEnabled parameter shall contain a single value that corresponds to the value of the MetricCollectionEnabled property of the instance of CIM_MetricDefForME that associates the instance of CIM_ManagedElement identified by the Subject parameter with the instance of CIM_BaseMetricDefinition identified by the Definition parameter. If the Subject and Definition parameters are both NULL, the method shall return 2 "Failed". If the Definition parameter is a reference to an instance of CIM_BaseMetricDefinition that is not associated to the CIM_MetricService through CIM_ServiceAffectsElement, the method shall return 2 "Failed". If the Subject parameter does not identify a single instance the Method shall return 2 "Failed".

// <param name="Definition" type="CIM_BaseMetricDefinition ">The Definition parameter identifies an instance of CIM_BaseMetricDefintion. The method returns references to instances of CIM_ManagedElement for which metrics defined by the instance of CIM_BaseMetricDefinition are available to be collected.</param>
// <param name="Subject" type="CIM_ManagedElement ">The Subject parameter identifies an instance of CIM_ManagedElement for which the method returns references to instances of CIM_BaseMetricDefinition that define metrics that are being captured for the CIM_ManagedElement.</param>

// <param name="DefinitionList" type="CIM_BaseMetricDefinition []">Upon successful completion of the method, the DefinitionList parameter may contain references to instances of CIM_BaseMetricDefinitions that define metrics available for collection for the CIM_ManagedElement identified by the Subject parameter.</param>
// <param name="ManagedElements" type="CIM_ManagedElement []">Upon successful completion of the method, the ManagedElements[] parameter may contain references to CIM_ManagedElements for which the metric identified by Definition parameter is available for collection.</param>
// <param name="MetricCollectionEnabled" type="MetricService_MetricCollectionEnabled []">The MetricCollectionEnabled parameter indicates whether a metric is being collected for a managed element.</param>
// <param name="MetricNames" type="string []">Upon successful completion of the method, each array index of the MetricNames parameter shall contain the value of the Name property for the instance of CIM_BaseMetricDefinition referenced by the corresponding array index of the DefinitionList parameter.</param>
// <param name="ReturnValue" type="uint32 "></param>
func (instance *CIM_MetricService) ShowMetrics( /* IN */ Subject CIM_ManagedElement,
 /* IN */ Definition CIM_BaseMetricDefinition,
 /* OUT */ ManagedElements []CIM_ManagedElement,
 /* OUT */ DefinitionList []CIM_BaseMetricDefinition,
 /* OUT */ MetricNames []string,
 /* OUT */ MetricCollectionEnabled []MetricService_MetricCollectionEnabled) (result uint32, err error) {retVal, err := instance.InvokeMethod("ShowMetrics" , Subject, Definition)
	if err != nil { return }
	retValue := retVal[0].(int32)
	result = uint32(retValue)
	return
	
}


// ShowMetricsByClass reports the Metrics available to be collected for all instances of a CIM class. The CIM classes for which a metric defined by an instance of CIM_BaseMetricDefinition is available to be collected, and whether or not a particular metric is currently being collected for a ManagedElement. 
///If the Subject parameter is specified and the Definition parameter is NULL, upon successful completion of the method, the DefinitionList[] parameter shall contain a reference to an instance of CIM_BaseMetricDefinition for each instance of CIM_BaseMetricDefinition that is associated with all instances of the class identified by the Subject parameter, where the CIM_BaseMetricDefinition instance is associated to the CIM_MetricService instance through CIM_ServiceAffectsElement; the MetricCollectionEnabled parameter shall contain the value 2 "Enabled" if the value of the MetricCollectionEnabled property has the value 2 "Enabled" for every instance of CIM_MetricDefForME that associates the CIM_ManagedElement identified by the Subject parameter with the CIM_BaseMetricDefinition for which a reference CIM_BaseMetricDefinition is returned in the DefinitionList parameter at the same array index as the reference and the MetricCollectionEnabled parameter shall contain the value 3 "Disabled" if the value of the MetricCollectionEnabled property does not have the value 2 "Enabled" for every instance of CIM_MetricDefForME that associates the CIM_ManagedElement identified by the Subject parameter with the CIM_BaseMetricDefinition for which a reference CIM_BaseMetricDefinition is returned in the DefinitionList parameter at the same array index as the reference. 
///If the Subject parameter and Definition parameter are both non-NULL, the method shall return NULL value for the DefinitionList parameter and the MetricCollectionEnabled parameter shall contain a single value that corresponds to the value of the MetricCollectionEnabled property of the instance of CIM_MetricDefForME that associates the instance of CIM_ManagedElement identified by the Subject parameter with the instance of CIM_BaseMetricDefinition identified by the Definition parameter. If the Subject and Definition parameters are both NULL, the method shall return 2 "Failed". If the Definition parameter is a reference to an instance of CIM_BaseMetricDefinition that is not associated to the CIM_MetricService through CIM_ServiceAffectsElement, the method shall return 2 "Failed". If the Subject parameter does not identify a single CIM Class the Method shall return 2 "Failed".

// <param name="Definition" type="CIM_BaseMetricDefinition ">The Definition parameter identifies an instance of CIM_BaseMetricDefinition. The method returns references to instances of CIM_ManagedElement for which metrics defined by the instance of CIM_BaseMetricDefinition are available to be collected.</param>
// <param name="Subject" type="CIM_ManagedElement ">The Subject parameter identifies a CIM class for which the method returns references to instances of CIM_BaseMetricDefinition that define metrics that are available to be captured for all instances of the class.</param>

// <param name="DefinitionList" type="CIM_BaseMetricDefinition []">Upon successful completion of the method, the DefinitionList parameter may contain references to instances of CIM_BaseMetricDefinitions that define metrics available for collection for the CIM_ManagedElement identified by the Subject parameter.</param>
// <param name="MetricCollectionEnabled" type="MetricService_MetricCollectionEnabled []">The MetricCollectionEnabled parameter indicates whether a metric is being collected for all instances of a class of managed elements.</param>
// <param name="MetricNames" type="string []">Upon successful completion of the method, each array index of the MetricNames parameter shall contain the value of the Name property for the instance of CIM_BaseMetricDefinition referenced by the corresponding array index of the DefinitionList parameter.</param>
// <param name="ReturnValue" type="uint32 "></param>
func (instance *CIM_MetricService) ShowMetricsByClass( /* IN */ Subject CIM_ManagedElement,
 /* IN */ Definition CIM_BaseMetricDefinition,
 /* OUT */ DefinitionList []CIM_BaseMetricDefinition,
 /* OUT */ MetricNames []string,
 /* OUT */ MetricCollectionEnabled []MetricService_MetricCollectionEnabled) (result uint32, err error) {retVal, err := instance.InvokeMethod("ShowMetricsByClass" , Subject, Definition)
	if err != nil { return }
	retValue := retVal[0].(int32)
	result = uint32(retValue)
	return
	
}


// ControlMetrics enables and disables the collection of metrics. It is used to control the collection of each type of metric for a CIM_ManagedElement, the collection of a given type of metric for all ManagedElements, or the collection of a specific metric for a specific ManagedElement. 
///If the Subject parameter is specified and the Definition parameter is NULL and the MetricCollectionEnabled parameter has the value 2 "Enabled" or 3 "Disabled", upon successful completion of the method, the MetricCollectionEnabled property of each instance of CIM_MetricDefForME that references the instance of CIM_ManagedElement identified by the Subject parameter and references an instance of CIM_BaseMetricDefinition that is associated to the CIM_BaseMetricService through the CIM_ServiceAffectsElement association shall have the value of the MetricCollectionEnabled parameter. If the Subject parameter is specified and the Definition parameter is NULL and the value of the MetricCollectionEnabled parameter is 4 "Reset" upon successful completion of the method, the value of the MetricCollectionEnabled of each instance of CIM_MetricDefForME that references the instance of CIM_ManagedElement identified by the Subject parameter and references an instance of CIM_BaseMetricDefinition that is associated to the CIM_BaseMetricService through the CIM_ServiceAffectsElement association shall transition to 3 "Disabled" then to 2 "Enabled". If the Definition parameter is non-NULL and the Subject parameter is NULL, and the MetricCollectionEnabled parameter has the value 2 "Enabled" or 3 "Disabled", upon successful completion of the method the MetricCollectionEnabled property of each instance of CIM_MetricDefForME that references the instance of CIM_BaseMetricDefinition identified by the Definition parameter shall have the value of the MetricCollectionEnabled parameter. If the Definition parameter is non-NULL and the Subject parameter is NULL, and the value of the MetricCollectionEnabled parameter is 4 "Reset" upon successful completion of the method the MetricCollectionEnabled property of each instance of CIM_MetricDefForME that references the instance of CIM_BaseMetricDefinition identified by the Definition parameter shall transition to 3 "Disabled" then to 2 "Enabled". If the Subject parameter and Definition parameter are both non-NULL, and the MetricCollectionEnabled parameter has the value 2 "Enable" or 3 "Disable", upon successful completion of the method, the MetricCollectionEnabled property of the instance of CIM_MetricDefForME that references the instance of CIM_ManagedElement identified by the Subject parameter and references the instance of CIM_BaseMetricDefinition identified by the Definition parameter shall have the value of the MetricCollectionEnabled parameter.
///If the Subject parameter and Definition parameter are both non-NULL and the value of the MetricCollectionEnabled parameter is 4 "Reset" upon successful completion of the method, the MetricCollectionEnabled property of the instance of CIM_MetricDefForME that references the instance of CIM_ManagedElement identified by the Subject parameter and references the instance of CIM_BaseMetricDefinition identified by the Definition parameter shall transition to 3 "Disabled" then to 2 "Enabled". If the Subject parameter and Definition parameter are both non-NULL and there is not an instance of CIM_MetricDefForME that associates the two instances, the method shall return 2 "Failed". If the Subject and Definition parameters are both NULL, the method shall return 2 "Failed". If the Definition parameter is a reference to an instance of CIM_BaseMetricDefinition that is not associated to the CIM_MetricService through CIM_ServiceAffectsElement, the method shall return 2 "Failed". If the Subject parameter does not identify a single instance the Method shall return 2 "Failed".

// <param name="Definition" type="CIM_BaseMetricDefinition ">The Definition parameter identifies a CIM_BaseMetricDefinition for which metrics will be controlled.</param>
// <param name="MetricCollectionEnabled" type="MetricService_MetricCollectionEnabled ">The MetricCollectionEnabled parameter indicates the desired operation to perform on the metrics.</param>
// <param name="Subject" type="CIM_ManagedElement ">The Subject parameter identifies managed element(s) for which metrics will be controlled.</param>

// <param name="ReturnValue" type="uint32 "></param>
func (instance *CIM_MetricService) ControlMetrics( /* IN */ Subject CIM_ManagedElement,
 /* IN */ Definition CIM_BaseMetricDefinition,
 /* IN */ MetricCollectionEnabled MetricService_MetricCollectionEnabled) (result uint32, err error) {retVal, err := instance.InvokeMethodWithReturn("ControlMetrics" , Subject, Definition, MetricCollectionEnabled);
	if err != nil { return }
	result = uint32(retVal)
	return
	
}


// ControlMetricsByClass enables and disables the collection of metrics. It is used to control the collection of each type of metric for all instances of a class or the collection of a specific metric for all instances of a class. 
///If the Definition parameter is NULL, and the MetricCollectionEnabled parameter has the value 2 "Enabled" or 3 "Disabled", upon successful completion of the method, the MetricCollectionEnabled property of each instance of CIM_MetricDefForME that references an instance of the class identified by the Subject parameter and references an instance of CIM_BaseMetricDefinition that is associated to the CIM_BaseMetricService through the CIM_ServiceAffectsElement association shall have the value of the MetricCollectionEnabled parameter. If the Definition parameter is NULL, and the MetricCollectionEnabled parameter has the value 4 "Reset", upon successful completion of the method, the MetricCollectionEnabled property of each instance of CIM_MetricDefForME that references an instance of the class identified by the Subject parameter and references an instance of CIM_BaseMetricDefinition that is associated to the CIM_BaseMetricService through the CIM_ServiceAffectsElement association shall transition to 3 "Disabled" then to 2 "Enabled". If the Definition parameter is non-NULL, and the MetricCollectionEnabled parameter has the value 2 "Enable" or 3 "Disable", upon successful completion of the method the MetricCollectionEnabled property of each instance of CIM_MetricDefForME that references the instance of CIM_BaseMetricDefinition identified by the Definition parameter and references an instance of the class identified by the Subject parameter shall have the value of the MetricCollectionEnabled parameter. If the Definition parameter is non-NULL, and the MetricCollectionEnabled parameter has the value 4 "Reset", upon successful completion of the method the MetricCollectionEnabled property of each instance of CIM_MetricDefForME that references the instance of CIM_BaseMetricDefinition identified by the Definition parameter and references an instance of the class identified by the Subject parameter shall transition to 3 "Disabled" then to 2 "Enabled". For each instance of the class identified by the Subject parameter, if there is not an instance of CIM_MetricDefForME that associates the CIM_BaseMetricDefinition instance identified by the Definition parameter to the instance, the method shall return 2 "Failed". If the Subject and Definition parameters are both NULL, the method shall return 2 "Failed". If the Definition parameter is a reference to an instance of CIM_BaseMetricDefinition that is not associated to the CIM_MetricService through CIM_ServiceAffectsElement, the method shall return 2 "Failed". If the Subject parameter does not identify a single CIM Class the Method shall return 2 "Failed".

// <param name="Definition" type="CIM_BaseMetricDefinition ">The Definition parameter identifies a CIM_BaseMetricDefinition for which metrics will be controlled.</param>
// <param name="MetricCollectionEnabled" type="MetricService_MetricCollectionEnabled ">The MetricCollectionEnabled parameter indicates the desired operation to perform on the metrics.</param>
// <param name="Subject" type="CIM_ManagedElement ">The Subject parameter identifies the CIM class for which metrics will be controlled.</param>

// <param name="ReturnValue" type="uint32 "></param>
func (instance *CIM_MetricService) ControlMetricsByClass( /* IN */ Subject CIM_ManagedElement,
 /* IN */ Definition CIM_BaseMetricDefinition,
 /* IN */ MetricCollectionEnabled MetricService_MetricCollectionEnabled) (result uint32, err error) {retVal, err := instance.InvokeMethodWithReturn("ControlMetricsByClass" , Subject, Definition, MetricCollectionEnabled);
	if err != nil { return }
	result = uint32(retVal)
	return
	
}


// GetMetricValues provides the ability to return a filtered list of CIM_BaseMetricValue instances. 
///If the Definition parameter is NULL, the method shall return 2 "Failed". If the Definition parameter is a reference to an instance of CIM_BaseMetricDefinition with which the CIM_MetricService is not associated through CIM_ServiceAffectsElement, the method shall return 2 "Failed".

// <param name="Count" type="uint16 ">The Count parameter identifies the maximum number of instances to to be returned by the method.</param>
// <param name="Definition" type="CIM_BaseMetricDefinition ">The Definition parameter identifies a CIM_BaseMetricDefinition for which metrics will be returned.</param>
// <param name="Range" type="MetricService_Range ">The Range parameter identifies how the instances are selected. The algorithm for ordering value instances is metric definition specific.</param>

// <param name="ReturnValue" type="uint32 "></param>
// <param name="Values" type="CIM_BaseMetricValue []">Upon successful completion of the method, the Values parameter contains references to instances of CIM_BaseMetricValue, filteredaccording to the values of the input parameters.</param>
func (instance *CIM_MetricService) GetMetricValues( /* IN */ Definition CIM_BaseMetricDefinition,
 /* IN */ Range MetricService_Range,
 /* IN */ Count uint16,
 /* OUT */ Values []CIM_BaseMetricValue) (result uint32, err error) {retVal, err := instance.InvokeMethod("GetMetricValues" , Definition, Range, Count)
	if err != nil { return }
	retValue := retVal[0].(int32)
	result = uint32(retValue)
	return
	
}


// Method used to allow specification of the point in time metric gathering is to be started and to specify the preferred sample interval time for periodic data gathering. 
///Whenever sampling for additional metrics is started, the settings specified by this method may be used.

// <param name="PreferredSampleInterval" type="string ">Preferred sample interval time. In order to get correlatable metrics, it is recommended that the sample interval be chosen in a way that 3600 modulo sample interval time in seconds is equal to 0. 
///It is the responsibility of the CIM metric service implementation to decide whether the requested sample interval time is honored. 
///The CIM client can check whether or not the metric providers are honoring the requested sample interval time by retrieving related BaseMetricDefinition instances and checking the contents of the "CIM_BaseMetricDefinition.SampleInterval" property.</param>
// <param name="RestartGathering" type="bool ">Boolean that when set to TRUE requests that gathering of all metrics associated to the metric service is re-started with this method call.</param>
// <param name="StartSampleTime" type="string ">Point in time when sampling for the metrics is to be started. 
///A value of 99990101000000.000000+000 shall indicate that sampling should start at the next time it is synchronized to the full hour. Sampling is synchronized to the full hour if seconds since midnight modulo sample interval in seconds is equal to 0.</param>

// <param name="ReturnValue" type="uint32 "></param>
func (instance *CIM_MetricService) ControlSampleTimes( /* IN */ StartSampleTime string,
 /* IN */ PreferredSampleInterval string,
 /* IN */ RestartGathering bool) (result uint32, err error) {retVal, err := instance.InvokeMethodWithReturn("ControlSampleTimes" , StartSampleTime, PreferredSampleInterval, RestartGathering);
	if err != nil { return }
	result = uint32(retVal)
	return
	
}


