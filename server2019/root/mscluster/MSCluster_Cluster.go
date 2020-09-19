// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 9/18/2020 using wmigen
//      Source root.MSCluster
//////////////////////////////////////////////
package mscluster

import (
	"github.com/microsoft/wmi/pkg/base/query"
	"github.com/microsoft/wmi/pkg/errors"
	cim "github.com/microsoft/wmi/pkg/wmiinstance"
	"reflect"
)

// MSCluster_Cluster struct
type MSCluster_Cluster struct {
	*CIM_Cluster

	//
	AddEvictDelay uint32

	//
	AdminAccessPoint uint32

	//
	AdminExtensions []string

	//
	AutoAssignNodeSite uint32

	//
	AutoBalancerLevel uint32

	//
	AutoBalancerMode uint32

	//
	BackupInProgress uint32

	//
	BlockCacheSize uint32

	//
	ClusSvcHangTimeout uint32

	//
	ClusSvcRegroupOpeningTimeout uint32

	//
	ClusSvcRegroupPruningTimeout uint32

	//
	ClusSvcRegroupStageTimeout uint32

	//
	ClusSvcRegroupTickInMilliseconds uint32

	//
	ClusterEnforcedAntiAffinity uint32

	//
	ClusterFunctionalLevel uint32

	//
	ClusterGroupWaitDelay uint32

	//
	ClusterLogLevel uint32

	//
	ClusterLogSize uint32

	//
	ClusterUpgradeVersion uint32

	//
	CrossSiteDelay uint32

	//
	CrossSiteThreshold uint32

	//
	CrossSubnetDelay uint32

	//
	CrossSubnetThreshold uint32

	//
	CsvBalancer uint32

	//
	DatabaseReadWriteMode uint32

	//
	DefaultNetworkRole uint32

	//
	DetectedCloudPlatform uint32

	//
	DetectManagedEvents uint32

	//
	DetectManagedEventsThreshold uint32

	//
	DisableGroupPreferredOwnerRandomization uint32

	//
	DrainOnShutdown uint32

	//
	DumpPolicy uint64

	//
	DynamicQuorumEnabled uint32

	//
	EnabledEventLogs []string

	//
	EnableSharedVolumes uint32

	//
	FixQuorum uint32

	//
	Fqdn string

	//
	GracePeriodEnabled uint32

	//
	GracePeriodTimeout uint32

	//
	GroupAdminExtensions []string

	//
	GroupDependencyTimeout uint32

	//
	HangRecoveryAction uint32

	//
	IgnorePersistentStateOnStartup uint32

	//
	LogResourceControls uint32

	//
	LowerQuorumPriorityNodeId uint32

	//
	MaintenanceFile string

	//
	MessageBufferLength uint32

	//
	MinimumNeverPreemptPriority uint32

	//
	MinimumPreemptorPriority uint32

	//
	NetftIPSecEnabled uint32

	//
	NetworkAdminExtensions []string

	//
	NetworkInterfaceAdminExtensions []string

	//
	NetworkPriorities []string

	//
	NodeAdminExtensions []string

	//
	PlacementOptions uint32

	//
	PlumbAllCrossSubnetRoutes uint32

	//
	PreferredSite string

	//
	PreventQuorum uint32

	//
	PrivateProperties MSCluster_Property

	//
	QuarantineDuration uint32

	//
	QuarantineThreshold uint32

	//
	QuorumArbitrationTimeMax uint32

	//
	QuorumArbitrationTimeMin uint32

	//
	QuorumLogFileSize uint32

	//
	QuorumPath string

	//
	QuorumType string

	//
	QuorumTypeValue uint32

	//
	RecentEventsResetTime string

	//
	RequestReplyTimeout uint32

	//
	ResiliencyDefaultPeriod uint32

	//
	ResiliencyLevel uint32

	//
	ResourceAdminExtensions []string

	//
	ResourceDllDeadlockPeriod uint32

	//
	ResourceTypeAdminExtensions []string

	//
	RootMemoryReserved uint32

	//
	RouteHistoryLength uint32

	//
	S2DBusTypes uint32

	//
	S2DCacheBehavior uint64

	//
	S2DCacheDesiredState uint32

	//
	S2DCacheDeviceModel []string

	//
	S2DCacheFlashReservePercent uint32

	//
	S2DCacheMetadataReserveBytes uint64

	//
	S2DCachePageSizeKBytes uint32

	//
	S2DEnabled uint32

	//
	S2DIOLatencyThreshold uint32

	//
	S2DOptimizations uint32

	//
	SameSubnetDelay uint32

	//
	SameSubnetThreshold uint32

	//
	Security []uint8

	//
	Security_Descriptor []uint8

	//
	SecurityLevel uint32

	//
	SecurityLevelForStorage uint32

	//
	SharedVolumeCompatibleFilters []string

	//
	SharedVolumeIncompatibleFilters []string

	//
	SharedVolumeSecurityDescriptor []uint8

	//
	SharedVolumesRoot string

	//
	SharedVolumeVssWriterOperationTimeout uint32

	//
	ShutdownTimeoutInMinutes uint32

	//
	UseClientAccessNetworksForSharedVolumes uint32

	//
	WitnessDatabaseWriteTimeout uint32

	//
	WitnessDynamicWeight uint32

	//
	WitnessRestartInterval uint32
}

func NewMSCluster_ClusterEx1(instance *cim.WmiInstance) (newInstance *MSCluster_Cluster, err error) {
	tmp, err := NewCIM_ClusterEx1(instance)

	if err != nil {
		return
	}
	newInstance = &MSCluster_Cluster{
		CIM_Cluster: tmp,
	}
	return
}

func NewMSCluster_ClusterEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *MSCluster_Cluster, err error) {
	tmp, err := NewCIM_ClusterEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &MSCluster_Cluster{
		CIM_Cluster: tmp,
	}
	return
}

// SetAddEvictDelay sets the value of AddEvictDelay for the instance
func (instance *MSCluster_Cluster) SetPropertyAddEvictDelay(value uint32) (err error) {
	return instance.SetProperty("AddEvictDelay", (value))
}

// GetAddEvictDelay gets the value of AddEvictDelay for the instance
func (instance *MSCluster_Cluster) GetPropertyAddEvictDelay() (value uint32, err error) {
	retValue, err := instance.GetProperty("AddEvictDelay")
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

// SetAdminAccessPoint sets the value of AdminAccessPoint for the instance
func (instance *MSCluster_Cluster) SetPropertyAdminAccessPoint(value uint32) (err error) {
	return instance.SetProperty("AdminAccessPoint", (value))
}

// GetAdminAccessPoint gets the value of AdminAccessPoint for the instance
func (instance *MSCluster_Cluster) GetPropertyAdminAccessPoint() (value uint32, err error) {
	retValue, err := instance.GetProperty("AdminAccessPoint")
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

// SetAdminExtensions sets the value of AdminExtensions for the instance
func (instance *MSCluster_Cluster) SetPropertyAdminExtensions(value []string) (err error) {
	return instance.SetProperty("AdminExtensions", (value))
}

// GetAdminExtensions gets the value of AdminExtensions for the instance
func (instance *MSCluster_Cluster) GetPropertyAdminExtensions() (value []string, err error) {
	retValue, err := instance.GetProperty("AdminExtensions")
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

// SetAutoAssignNodeSite sets the value of AutoAssignNodeSite for the instance
func (instance *MSCluster_Cluster) SetPropertyAutoAssignNodeSite(value uint32) (err error) {
	return instance.SetProperty("AutoAssignNodeSite", (value))
}

// GetAutoAssignNodeSite gets the value of AutoAssignNodeSite for the instance
func (instance *MSCluster_Cluster) GetPropertyAutoAssignNodeSite() (value uint32, err error) {
	retValue, err := instance.GetProperty("AutoAssignNodeSite")
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

// SetAutoBalancerLevel sets the value of AutoBalancerLevel for the instance
func (instance *MSCluster_Cluster) SetPropertyAutoBalancerLevel(value uint32) (err error) {
	return instance.SetProperty("AutoBalancerLevel", (value))
}

// GetAutoBalancerLevel gets the value of AutoBalancerLevel for the instance
func (instance *MSCluster_Cluster) GetPropertyAutoBalancerLevel() (value uint32, err error) {
	retValue, err := instance.GetProperty("AutoBalancerLevel")
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

// SetAutoBalancerMode sets the value of AutoBalancerMode for the instance
func (instance *MSCluster_Cluster) SetPropertyAutoBalancerMode(value uint32) (err error) {
	return instance.SetProperty("AutoBalancerMode", (value))
}

// GetAutoBalancerMode gets the value of AutoBalancerMode for the instance
func (instance *MSCluster_Cluster) GetPropertyAutoBalancerMode() (value uint32, err error) {
	retValue, err := instance.GetProperty("AutoBalancerMode")
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

// SetBackupInProgress sets the value of BackupInProgress for the instance
func (instance *MSCluster_Cluster) SetPropertyBackupInProgress(value uint32) (err error) {
	return instance.SetProperty("BackupInProgress", (value))
}

// GetBackupInProgress gets the value of BackupInProgress for the instance
func (instance *MSCluster_Cluster) GetPropertyBackupInProgress() (value uint32, err error) {
	retValue, err := instance.GetProperty("BackupInProgress")
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

// SetBlockCacheSize sets the value of BlockCacheSize for the instance
func (instance *MSCluster_Cluster) SetPropertyBlockCacheSize(value uint32) (err error) {
	return instance.SetProperty("BlockCacheSize", (value))
}

// GetBlockCacheSize gets the value of BlockCacheSize for the instance
func (instance *MSCluster_Cluster) GetPropertyBlockCacheSize() (value uint32, err error) {
	retValue, err := instance.GetProperty("BlockCacheSize")
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

// SetClusSvcHangTimeout sets the value of ClusSvcHangTimeout for the instance
func (instance *MSCluster_Cluster) SetPropertyClusSvcHangTimeout(value uint32) (err error) {
	return instance.SetProperty("ClusSvcHangTimeout", (value))
}

// GetClusSvcHangTimeout gets the value of ClusSvcHangTimeout for the instance
func (instance *MSCluster_Cluster) GetPropertyClusSvcHangTimeout() (value uint32, err error) {
	retValue, err := instance.GetProperty("ClusSvcHangTimeout")
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

// SetClusSvcRegroupOpeningTimeout sets the value of ClusSvcRegroupOpeningTimeout for the instance
func (instance *MSCluster_Cluster) SetPropertyClusSvcRegroupOpeningTimeout(value uint32) (err error) {
	return instance.SetProperty("ClusSvcRegroupOpeningTimeout", (value))
}

// GetClusSvcRegroupOpeningTimeout gets the value of ClusSvcRegroupOpeningTimeout for the instance
func (instance *MSCluster_Cluster) GetPropertyClusSvcRegroupOpeningTimeout() (value uint32, err error) {
	retValue, err := instance.GetProperty("ClusSvcRegroupOpeningTimeout")
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

// SetClusSvcRegroupPruningTimeout sets the value of ClusSvcRegroupPruningTimeout for the instance
func (instance *MSCluster_Cluster) SetPropertyClusSvcRegroupPruningTimeout(value uint32) (err error) {
	return instance.SetProperty("ClusSvcRegroupPruningTimeout", (value))
}

// GetClusSvcRegroupPruningTimeout gets the value of ClusSvcRegroupPruningTimeout for the instance
func (instance *MSCluster_Cluster) GetPropertyClusSvcRegroupPruningTimeout() (value uint32, err error) {
	retValue, err := instance.GetProperty("ClusSvcRegroupPruningTimeout")
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

// SetClusSvcRegroupStageTimeout sets the value of ClusSvcRegroupStageTimeout for the instance
func (instance *MSCluster_Cluster) SetPropertyClusSvcRegroupStageTimeout(value uint32) (err error) {
	return instance.SetProperty("ClusSvcRegroupStageTimeout", (value))
}

// GetClusSvcRegroupStageTimeout gets the value of ClusSvcRegroupStageTimeout for the instance
func (instance *MSCluster_Cluster) GetPropertyClusSvcRegroupStageTimeout() (value uint32, err error) {
	retValue, err := instance.GetProperty("ClusSvcRegroupStageTimeout")
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

// SetClusSvcRegroupTickInMilliseconds sets the value of ClusSvcRegroupTickInMilliseconds for the instance
func (instance *MSCluster_Cluster) SetPropertyClusSvcRegroupTickInMilliseconds(value uint32) (err error) {
	return instance.SetProperty("ClusSvcRegroupTickInMilliseconds", (value))
}

// GetClusSvcRegroupTickInMilliseconds gets the value of ClusSvcRegroupTickInMilliseconds for the instance
func (instance *MSCluster_Cluster) GetPropertyClusSvcRegroupTickInMilliseconds() (value uint32, err error) {
	retValue, err := instance.GetProperty("ClusSvcRegroupTickInMilliseconds")
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

// SetClusterEnforcedAntiAffinity sets the value of ClusterEnforcedAntiAffinity for the instance
func (instance *MSCluster_Cluster) SetPropertyClusterEnforcedAntiAffinity(value uint32) (err error) {
	return instance.SetProperty("ClusterEnforcedAntiAffinity", (value))
}

// GetClusterEnforcedAntiAffinity gets the value of ClusterEnforcedAntiAffinity for the instance
func (instance *MSCluster_Cluster) GetPropertyClusterEnforcedAntiAffinity() (value uint32, err error) {
	retValue, err := instance.GetProperty("ClusterEnforcedAntiAffinity")
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

// SetClusterFunctionalLevel sets the value of ClusterFunctionalLevel for the instance
func (instance *MSCluster_Cluster) SetPropertyClusterFunctionalLevel(value uint32) (err error) {
	return instance.SetProperty("ClusterFunctionalLevel", (value))
}

// GetClusterFunctionalLevel gets the value of ClusterFunctionalLevel for the instance
func (instance *MSCluster_Cluster) GetPropertyClusterFunctionalLevel() (value uint32, err error) {
	retValue, err := instance.GetProperty("ClusterFunctionalLevel")
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

// SetClusterGroupWaitDelay sets the value of ClusterGroupWaitDelay for the instance
func (instance *MSCluster_Cluster) SetPropertyClusterGroupWaitDelay(value uint32) (err error) {
	return instance.SetProperty("ClusterGroupWaitDelay", (value))
}

// GetClusterGroupWaitDelay gets the value of ClusterGroupWaitDelay for the instance
func (instance *MSCluster_Cluster) GetPropertyClusterGroupWaitDelay() (value uint32, err error) {
	retValue, err := instance.GetProperty("ClusterGroupWaitDelay")
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

// SetClusterLogLevel sets the value of ClusterLogLevel for the instance
func (instance *MSCluster_Cluster) SetPropertyClusterLogLevel(value uint32) (err error) {
	return instance.SetProperty("ClusterLogLevel", (value))
}

// GetClusterLogLevel gets the value of ClusterLogLevel for the instance
func (instance *MSCluster_Cluster) GetPropertyClusterLogLevel() (value uint32, err error) {
	retValue, err := instance.GetProperty("ClusterLogLevel")
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

// SetClusterLogSize sets the value of ClusterLogSize for the instance
func (instance *MSCluster_Cluster) SetPropertyClusterLogSize(value uint32) (err error) {
	return instance.SetProperty("ClusterLogSize", (value))
}

// GetClusterLogSize gets the value of ClusterLogSize for the instance
func (instance *MSCluster_Cluster) GetPropertyClusterLogSize() (value uint32, err error) {
	retValue, err := instance.GetProperty("ClusterLogSize")
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

// SetClusterUpgradeVersion sets the value of ClusterUpgradeVersion for the instance
func (instance *MSCluster_Cluster) SetPropertyClusterUpgradeVersion(value uint32) (err error) {
	return instance.SetProperty("ClusterUpgradeVersion", (value))
}

// GetClusterUpgradeVersion gets the value of ClusterUpgradeVersion for the instance
func (instance *MSCluster_Cluster) GetPropertyClusterUpgradeVersion() (value uint32, err error) {
	retValue, err := instance.GetProperty("ClusterUpgradeVersion")
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

// SetCrossSiteDelay sets the value of CrossSiteDelay for the instance
func (instance *MSCluster_Cluster) SetPropertyCrossSiteDelay(value uint32) (err error) {
	return instance.SetProperty("CrossSiteDelay", (value))
}

// GetCrossSiteDelay gets the value of CrossSiteDelay for the instance
func (instance *MSCluster_Cluster) GetPropertyCrossSiteDelay() (value uint32, err error) {
	retValue, err := instance.GetProperty("CrossSiteDelay")
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

// SetCrossSiteThreshold sets the value of CrossSiteThreshold for the instance
func (instance *MSCluster_Cluster) SetPropertyCrossSiteThreshold(value uint32) (err error) {
	return instance.SetProperty("CrossSiteThreshold", (value))
}

// GetCrossSiteThreshold gets the value of CrossSiteThreshold for the instance
func (instance *MSCluster_Cluster) GetPropertyCrossSiteThreshold() (value uint32, err error) {
	retValue, err := instance.GetProperty("CrossSiteThreshold")
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

// SetCrossSubnetDelay sets the value of CrossSubnetDelay for the instance
func (instance *MSCluster_Cluster) SetPropertyCrossSubnetDelay(value uint32) (err error) {
	return instance.SetProperty("CrossSubnetDelay", (value))
}

// GetCrossSubnetDelay gets the value of CrossSubnetDelay for the instance
func (instance *MSCluster_Cluster) GetPropertyCrossSubnetDelay() (value uint32, err error) {
	retValue, err := instance.GetProperty("CrossSubnetDelay")
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

// SetCrossSubnetThreshold sets the value of CrossSubnetThreshold for the instance
func (instance *MSCluster_Cluster) SetPropertyCrossSubnetThreshold(value uint32) (err error) {
	return instance.SetProperty("CrossSubnetThreshold", (value))
}

// GetCrossSubnetThreshold gets the value of CrossSubnetThreshold for the instance
func (instance *MSCluster_Cluster) GetPropertyCrossSubnetThreshold() (value uint32, err error) {
	retValue, err := instance.GetProperty("CrossSubnetThreshold")
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

// SetCsvBalancer sets the value of CsvBalancer for the instance
func (instance *MSCluster_Cluster) SetPropertyCsvBalancer(value uint32) (err error) {
	return instance.SetProperty("CsvBalancer", (value))
}

// GetCsvBalancer gets the value of CsvBalancer for the instance
func (instance *MSCluster_Cluster) GetPropertyCsvBalancer() (value uint32, err error) {
	retValue, err := instance.GetProperty("CsvBalancer")
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

// SetDatabaseReadWriteMode sets the value of DatabaseReadWriteMode for the instance
func (instance *MSCluster_Cluster) SetPropertyDatabaseReadWriteMode(value uint32) (err error) {
	return instance.SetProperty("DatabaseReadWriteMode", (value))
}

// GetDatabaseReadWriteMode gets the value of DatabaseReadWriteMode for the instance
func (instance *MSCluster_Cluster) GetPropertyDatabaseReadWriteMode() (value uint32, err error) {
	retValue, err := instance.GetProperty("DatabaseReadWriteMode")
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

// SetDefaultNetworkRole sets the value of DefaultNetworkRole for the instance
func (instance *MSCluster_Cluster) SetPropertyDefaultNetworkRole(value uint32) (err error) {
	return instance.SetProperty("DefaultNetworkRole", (value))
}

// GetDefaultNetworkRole gets the value of DefaultNetworkRole for the instance
func (instance *MSCluster_Cluster) GetPropertyDefaultNetworkRole() (value uint32, err error) {
	retValue, err := instance.GetProperty("DefaultNetworkRole")
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

// SetDetectedCloudPlatform sets the value of DetectedCloudPlatform for the instance
func (instance *MSCluster_Cluster) SetPropertyDetectedCloudPlatform(value uint32) (err error) {
	return instance.SetProperty("DetectedCloudPlatform", (value))
}

// GetDetectedCloudPlatform gets the value of DetectedCloudPlatform for the instance
func (instance *MSCluster_Cluster) GetPropertyDetectedCloudPlatform() (value uint32, err error) {
	retValue, err := instance.GetProperty("DetectedCloudPlatform")
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

// SetDetectManagedEvents sets the value of DetectManagedEvents for the instance
func (instance *MSCluster_Cluster) SetPropertyDetectManagedEvents(value uint32) (err error) {
	return instance.SetProperty("DetectManagedEvents", (value))
}

// GetDetectManagedEvents gets the value of DetectManagedEvents for the instance
func (instance *MSCluster_Cluster) GetPropertyDetectManagedEvents() (value uint32, err error) {
	retValue, err := instance.GetProperty("DetectManagedEvents")
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

// SetDetectManagedEventsThreshold sets the value of DetectManagedEventsThreshold for the instance
func (instance *MSCluster_Cluster) SetPropertyDetectManagedEventsThreshold(value uint32) (err error) {
	return instance.SetProperty("DetectManagedEventsThreshold", (value))
}

// GetDetectManagedEventsThreshold gets the value of DetectManagedEventsThreshold for the instance
func (instance *MSCluster_Cluster) GetPropertyDetectManagedEventsThreshold() (value uint32, err error) {
	retValue, err := instance.GetProperty("DetectManagedEventsThreshold")
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

// SetDisableGroupPreferredOwnerRandomization sets the value of DisableGroupPreferredOwnerRandomization for the instance
func (instance *MSCluster_Cluster) SetPropertyDisableGroupPreferredOwnerRandomization(value uint32) (err error) {
	return instance.SetProperty("DisableGroupPreferredOwnerRandomization", (value))
}

// GetDisableGroupPreferredOwnerRandomization gets the value of DisableGroupPreferredOwnerRandomization for the instance
func (instance *MSCluster_Cluster) GetPropertyDisableGroupPreferredOwnerRandomization() (value uint32, err error) {
	retValue, err := instance.GetProperty("DisableGroupPreferredOwnerRandomization")
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

// SetDrainOnShutdown sets the value of DrainOnShutdown for the instance
func (instance *MSCluster_Cluster) SetPropertyDrainOnShutdown(value uint32) (err error) {
	return instance.SetProperty("DrainOnShutdown", (value))
}

// GetDrainOnShutdown gets the value of DrainOnShutdown for the instance
func (instance *MSCluster_Cluster) GetPropertyDrainOnShutdown() (value uint32, err error) {
	retValue, err := instance.GetProperty("DrainOnShutdown")
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

// SetDumpPolicy sets the value of DumpPolicy for the instance
func (instance *MSCluster_Cluster) SetPropertyDumpPolicy(value uint64) (err error) {
	return instance.SetProperty("DumpPolicy", (value))
}

// GetDumpPolicy gets the value of DumpPolicy for the instance
func (instance *MSCluster_Cluster) GetPropertyDumpPolicy() (value uint64, err error) {
	retValue, err := instance.GetProperty("DumpPolicy")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(uint64)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " uint64 is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = uint64(valuetmp)

	return
}

// SetDynamicQuorumEnabled sets the value of DynamicQuorumEnabled for the instance
func (instance *MSCluster_Cluster) SetPropertyDynamicQuorumEnabled(value uint32) (err error) {
	return instance.SetProperty("DynamicQuorumEnabled", (value))
}

// GetDynamicQuorumEnabled gets the value of DynamicQuorumEnabled for the instance
func (instance *MSCluster_Cluster) GetPropertyDynamicQuorumEnabled() (value uint32, err error) {
	retValue, err := instance.GetProperty("DynamicQuorumEnabled")
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

// SetEnabledEventLogs sets the value of EnabledEventLogs for the instance
func (instance *MSCluster_Cluster) SetPropertyEnabledEventLogs(value []string) (err error) {
	return instance.SetProperty("EnabledEventLogs", (value))
}

// GetEnabledEventLogs gets the value of EnabledEventLogs for the instance
func (instance *MSCluster_Cluster) GetPropertyEnabledEventLogs() (value []string, err error) {
	retValue, err := instance.GetProperty("EnabledEventLogs")
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

// SetEnableSharedVolumes sets the value of EnableSharedVolumes for the instance
func (instance *MSCluster_Cluster) SetPropertyEnableSharedVolumes(value uint32) (err error) {
	return instance.SetProperty("EnableSharedVolumes", (value))
}

// GetEnableSharedVolumes gets the value of EnableSharedVolumes for the instance
func (instance *MSCluster_Cluster) GetPropertyEnableSharedVolumes() (value uint32, err error) {
	retValue, err := instance.GetProperty("EnableSharedVolumes")
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

// SetFixQuorum sets the value of FixQuorum for the instance
func (instance *MSCluster_Cluster) SetPropertyFixQuorum(value uint32) (err error) {
	return instance.SetProperty("FixQuorum", (value))
}

// GetFixQuorum gets the value of FixQuorum for the instance
func (instance *MSCluster_Cluster) GetPropertyFixQuorum() (value uint32, err error) {
	retValue, err := instance.GetProperty("FixQuorum")
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

// SetFqdn sets the value of Fqdn for the instance
func (instance *MSCluster_Cluster) SetPropertyFqdn(value string) (err error) {
	return instance.SetProperty("Fqdn", (value))
}

// GetFqdn gets the value of Fqdn for the instance
func (instance *MSCluster_Cluster) GetPropertyFqdn() (value string, err error) {
	retValue, err := instance.GetProperty("Fqdn")
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

// SetGracePeriodEnabled sets the value of GracePeriodEnabled for the instance
func (instance *MSCluster_Cluster) SetPropertyGracePeriodEnabled(value uint32) (err error) {
	return instance.SetProperty("GracePeriodEnabled", (value))
}

// GetGracePeriodEnabled gets the value of GracePeriodEnabled for the instance
func (instance *MSCluster_Cluster) GetPropertyGracePeriodEnabled() (value uint32, err error) {
	retValue, err := instance.GetProperty("GracePeriodEnabled")
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

// SetGracePeriodTimeout sets the value of GracePeriodTimeout for the instance
func (instance *MSCluster_Cluster) SetPropertyGracePeriodTimeout(value uint32) (err error) {
	return instance.SetProperty("GracePeriodTimeout", (value))
}

// GetGracePeriodTimeout gets the value of GracePeriodTimeout for the instance
func (instance *MSCluster_Cluster) GetPropertyGracePeriodTimeout() (value uint32, err error) {
	retValue, err := instance.GetProperty("GracePeriodTimeout")
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

// SetGroupAdminExtensions sets the value of GroupAdminExtensions for the instance
func (instance *MSCluster_Cluster) SetPropertyGroupAdminExtensions(value []string) (err error) {
	return instance.SetProperty("GroupAdminExtensions", (value))
}

// GetGroupAdminExtensions gets the value of GroupAdminExtensions for the instance
func (instance *MSCluster_Cluster) GetPropertyGroupAdminExtensions() (value []string, err error) {
	retValue, err := instance.GetProperty("GroupAdminExtensions")
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

// SetGroupDependencyTimeout sets the value of GroupDependencyTimeout for the instance
func (instance *MSCluster_Cluster) SetPropertyGroupDependencyTimeout(value uint32) (err error) {
	return instance.SetProperty("GroupDependencyTimeout", (value))
}

// GetGroupDependencyTimeout gets the value of GroupDependencyTimeout for the instance
func (instance *MSCluster_Cluster) GetPropertyGroupDependencyTimeout() (value uint32, err error) {
	retValue, err := instance.GetProperty("GroupDependencyTimeout")
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

// SetHangRecoveryAction sets the value of HangRecoveryAction for the instance
func (instance *MSCluster_Cluster) SetPropertyHangRecoveryAction(value uint32) (err error) {
	return instance.SetProperty("HangRecoveryAction", (value))
}

// GetHangRecoveryAction gets the value of HangRecoveryAction for the instance
func (instance *MSCluster_Cluster) GetPropertyHangRecoveryAction() (value uint32, err error) {
	retValue, err := instance.GetProperty("HangRecoveryAction")
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

// SetIgnorePersistentStateOnStartup sets the value of IgnorePersistentStateOnStartup for the instance
func (instance *MSCluster_Cluster) SetPropertyIgnorePersistentStateOnStartup(value uint32) (err error) {
	return instance.SetProperty("IgnorePersistentStateOnStartup", (value))
}

// GetIgnorePersistentStateOnStartup gets the value of IgnorePersistentStateOnStartup for the instance
func (instance *MSCluster_Cluster) GetPropertyIgnorePersistentStateOnStartup() (value uint32, err error) {
	retValue, err := instance.GetProperty("IgnorePersistentStateOnStartup")
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

// SetLogResourceControls sets the value of LogResourceControls for the instance
func (instance *MSCluster_Cluster) SetPropertyLogResourceControls(value uint32) (err error) {
	return instance.SetProperty("LogResourceControls", (value))
}

// GetLogResourceControls gets the value of LogResourceControls for the instance
func (instance *MSCluster_Cluster) GetPropertyLogResourceControls() (value uint32, err error) {
	retValue, err := instance.GetProperty("LogResourceControls")
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

// SetLowerQuorumPriorityNodeId sets the value of LowerQuorumPriorityNodeId for the instance
func (instance *MSCluster_Cluster) SetPropertyLowerQuorumPriorityNodeId(value uint32) (err error) {
	return instance.SetProperty("LowerQuorumPriorityNodeId", (value))
}

// GetLowerQuorumPriorityNodeId gets the value of LowerQuorumPriorityNodeId for the instance
func (instance *MSCluster_Cluster) GetPropertyLowerQuorumPriorityNodeId() (value uint32, err error) {
	retValue, err := instance.GetProperty("LowerQuorumPriorityNodeId")
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

// SetMaintenanceFile sets the value of MaintenanceFile for the instance
func (instance *MSCluster_Cluster) SetPropertyMaintenanceFile(value string) (err error) {
	return instance.SetProperty("MaintenanceFile", (value))
}

// GetMaintenanceFile gets the value of MaintenanceFile for the instance
func (instance *MSCluster_Cluster) GetPropertyMaintenanceFile() (value string, err error) {
	retValue, err := instance.GetProperty("MaintenanceFile")
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

// SetMessageBufferLength sets the value of MessageBufferLength for the instance
func (instance *MSCluster_Cluster) SetPropertyMessageBufferLength(value uint32) (err error) {
	return instance.SetProperty("MessageBufferLength", (value))
}

// GetMessageBufferLength gets the value of MessageBufferLength for the instance
func (instance *MSCluster_Cluster) GetPropertyMessageBufferLength() (value uint32, err error) {
	retValue, err := instance.GetProperty("MessageBufferLength")
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

// SetMinimumNeverPreemptPriority sets the value of MinimumNeverPreemptPriority for the instance
func (instance *MSCluster_Cluster) SetPropertyMinimumNeverPreemptPriority(value uint32) (err error) {
	return instance.SetProperty("MinimumNeverPreemptPriority", (value))
}

// GetMinimumNeverPreemptPriority gets the value of MinimumNeverPreemptPriority for the instance
func (instance *MSCluster_Cluster) GetPropertyMinimumNeverPreemptPriority() (value uint32, err error) {
	retValue, err := instance.GetProperty("MinimumNeverPreemptPriority")
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

// SetMinimumPreemptorPriority sets the value of MinimumPreemptorPriority for the instance
func (instance *MSCluster_Cluster) SetPropertyMinimumPreemptorPriority(value uint32) (err error) {
	return instance.SetProperty("MinimumPreemptorPriority", (value))
}

// GetMinimumPreemptorPriority gets the value of MinimumPreemptorPriority for the instance
func (instance *MSCluster_Cluster) GetPropertyMinimumPreemptorPriority() (value uint32, err error) {
	retValue, err := instance.GetProperty("MinimumPreemptorPriority")
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

// SetNetftIPSecEnabled sets the value of NetftIPSecEnabled for the instance
func (instance *MSCluster_Cluster) SetPropertyNetftIPSecEnabled(value uint32) (err error) {
	return instance.SetProperty("NetftIPSecEnabled", (value))
}

// GetNetftIPSecEnabled gets the value of NetftIPSecEnabled for the instance
func (instance *MSCluster_Cluster) GetPropertyNetftIPSecEnabled() (value uint32, err error) {
	retValue, err := instance.GetProperty("NetftIPSecEnabled")
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

// SetNetworkAdminExtensions sets the value of NetworkAdminExtensions for the instance
func (instance *MSCluster_Cluster) SetPropertyNetworkAdminExtensions(value []string) (err error) {
	return instance.SetProperty("NetworkAdminExtensions", (value))
}

// GetNetworkAdminExtensions gets the value of NetworkAdminExtensions for the instance
func (instance *MSCluster_Cluster) GetPropertyNetworkAdminExtensions() (value []string, err error) {
	retValue, err := instance.GetProperty("NetworkAdminExtensions")
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

// SetNetworkInterfaceAdminExtensions sets the value of NetworkInterfaceAdminExtensions for the instance
func (instance *MSCluster_Cluster) SetPropertyNetworkInterfaceAdminExtensions(value []string) (err error) {
	return instance.SetProperty("NetworkInterfaceAdminExtensions", (value))
}

// GetNetworkInterfaceAdminExtensions gets the value of NetworkInterfaceAdminExtensions for the instance
func (instance *MSCluster_Cluster) GetPropertyNetworkInterfaceAdminExtensions() (value []string, err error) {
	retValue, err := instance.GetProperty("NetworkInterfaceAdminExtensions")
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

// SetNetworkPriorities sets the value of NetworkPriorities for the instance
func (instance *MSCluster_Cluster) SetPropertyNetworkPriorities(value []string) (err error) {
	return instance.SetProperty("NetworkPriorities", (value))
}

// GetNetworkPriorities gets the value of NetworkPriorities for the instance
func (instance *MSCluster_Cluster) GetPropertyNetworkPriorities() (value []string, err error) {
	retValue, err := instance.GetProperty("NetworkPriorities")
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

// SetNodeAdminExtensions sets the value of NodeAdminExtensions for the instance
func (instance *MSCluster_Cluster) SetPropertyNodeAdminExtensions(value []string) (err error) {
	return instance.SetProperty("NodeAdminExtensions", (value))
}

// GetNodeAdminExtensions gets the value of NodeAdminExtensions for the instance
func (instance *MSCluster_Cluster) GetPropertyNodeAdminExtensions() (value []string, err error) {
	retValue, err := instance.GetProperty("NodeAdminExtensions")
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

// SetPlacementOptions sets the value of PlacementOptions for the instance
func (instance *MSCluster_Cluster) SetPropertyPlacementOptions(value uint32) (err error) {
	return instance.SetProperty("PlacementOptions", (value))
}

// GetPlacementOptions gets the value of PlacementOptions for the instance
func (instance *MSCluster_Cluster) GetPropertyPlacementOptions() (value uint32, err error) {
	retValue, err := instance.GetProperty("PlacementOptions")
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

// SetPlumbAllCrossSubnetRoutes sets the value of PlumbAllCrossSubnetRoutes for the instance
func (instance *MSCluster_Cluster) SetPropertyPlumbAllCrossSubnetRoutes(value uint32) (err error) {
	return instance.SetProperty("PlumbAllCrossSubnetRoutes", (value))
}

// GetPlumbAllCrossSubnetRoutes gets the value of PlumbAllCrossSubnetRoutes for the instance
func (instance *MSCluster_Cluster) GetPropertyPlumbAllCrossSubnetRoutes() (value uint32, err error) {
	retValue, err := instance.GetProperty("PlumbAllCrossSubnetRoutes")
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

// SetPreferredSite sets the value of PreferredSite for the instance
func (instance *MSCluster_Cluster) SetPropertyPreferredSite(value string) (err error) {
	return instance.SetProperty("PreferredSite", (value))
}

// GetPreferredSite gets the value of PreferredSite for the instance
func (instance *MSCluster_Cluster) GetPropertyPreferredSite() (value string, err error) {
	retValue, err := instance.GetProperty("PreferredSite")
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

// SetPreventQuorum sets the value of PreventQuorum for the instance
func (instance *MSCluster_Cluster) SetPropertyPreventQuorum(value uint32) (err error) {
	return instance.SetProperty("PreventQuorum", (value))
}

// GetPreventQuorum gets the value of PreventQuorum for the instance
func (instance *MSCluster_Cluster) GetPropertyPreventQuorum() (value uint32, err error) {
	retValue, err := instance.GetProperty("PreventQuorum")
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

// SetPrivateProperties sets the value of PrivateProperties for the instance
func (instance *MSCluster_Cluster) SetPropertyPrivateProperties(value MSCluster_Property) (err error) {
	return instance.SetProperty("PrivateProperties", (value))
}

// GetPrivateProperties gets the value of PrivateProperties for the instance
func (instance *MSCluster_Cluster) GetPropertyPrivateProperties() (value MSCluster_Property, err error) {
	retValue, err := instance.GetProperty("PrivateProperties")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(MSCluster_Property)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " MSCluster_Property is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = MSCluster_Property(valuetmp)

	return
}

// SetQuarantineDuration sets the value of QuarantineDuration for the instance
func (instance *MSCluster_Cluster) SetPropertyQuarantineDuration(value uint32) (err error) {
	return instance.SetProperty("QuarantineDuration", (value))
}

// GetQuarantineDuration gets the value of QuarantineDuration for the instance
func (instance *MSCluster_Cluster) GetPropertyQuarantineDuration() (value uint32, err error) {
	retValue, err := instance.GetProperty("QuarantineDuration")
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

// SetQuarantineThreshold sets the value of QuarantineThreshold for the instance
func (instance *MSCluster_Cluster) SetPropertyQuarantineThreshold(value uint32) (err error) {
	return instance.SetProperty("QuarantineThreshold", (value))
}

// GetQuarantineThreshold gets the value of QuarantineThreshold for the instance
func (instance *MSCluster_Cluster) GetPropertyQuarantineThreshold() (value uint32, err error) {
	retValue, err := instance.GetProperty("QuarantineThreshold")
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

// SetQuorumArbitrationTimeMax sets the value of QuorumArbitrationTimeMax for the instance
func (instance *MSCluster_Cluster) SetPropertyQuorumArbitrationTimeMax(value uint32) (err error) {
	return instance.SetProperty("QuorumArbitrationTimeMax", (value))
}

// GetQuorumArbitrationTimeMax gets the value of QuorumArbitrationTimeMax for the instance
func (instance *MSCluster_Cluster) GetPropertyQuorumArbitrationTimeMax() (value uint32, err error) {
	retValue, err := instance.GetProperty("QuorumArbitrationTimeMax")
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

// SetQuorumArbitrationTimeMin sets the value of QuorumArbitrationTimeMin for the instance
func (instance *MSCluster_Cluster) SetPropertyQuorumArbitrationTimeMin(value uint32) (err error) {
	return instance.SetProperty("QuorumArbitrationTimeMin", (value))
}

// GetQuorumArbitrationTimeMin gets the value of QuorumArbitrationTimeMin for the instance
func (instance *MSCluster_Cluster) GetPropertyQuorumArbitrationTimeMin() (value uint32, err error) {
	retValue, err := instance.GetProperty("QuorumArbitrationTimeMin")
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

// SetQuorumLogFileSize sets the value of QuorumLogFileSize for the instance
func (instance *MSCluster_Cluster) SetPropertyQuorumLogFileSize(value uint32) (err error) {
	return instance.SetProperty("QuorumLogFileSize", (value))
}

// GetQuorumLogFileSize gets the value of QuorumLogFileSize for the instance
func (instance *MSCluster_Cluster) GetPropertyQuorumLogFileSize() (value uint32, err error) {
	retValue, err := instance.GetProperty("QuorumLogFileSize")
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

// SetQuorumPath sets the value of QuorumPath for the instance
func (instance *MSCluster_Cluster) SetPropertyQuorumPath(value string) (err error) {
	return instance.SetProperty("QuorumPath", (value))
}

// GetQuorumPath gets the value of QuorumPath for the instance
func (instance *MSCluster_Cluster) GetPropertyQuorumPath() (value string, err error) {
	retValue, err := instance.GetProperty("QuorumPath")
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

// SetQuorumType sets the value of QuorumType for the instance
func (instance *MSCluster_Cluster) SetPropertyQuorumType(value string) (err error) {
	return instance.SetProperty("QuorumType", (value))
}

// GetQuorumType gets the value of QuorumType for the instance
func (instance *MSCluster_Cluster) GetPropertyQuorumType() (value string, err error) {
	retValue, err := instance.GetProperty("QuorumType")
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

// SetQuorumTypeValue sets the value of QuorumTypeValue for the instance
func (instance *MSCluster_Cluster) SetPropertyQuorumTypeValue(value uint32) (err error) {
	return instance.SetProperty("QuorumTypeValue", (value))
}

// GetQuorumTypeValue gets the value of QuorumTypeValue for the instance
func (instance *MSCluster_Cluster) GetPropertyQuorumTypeValue() (value uint32, err error) {
	retValue, err := instance.GetProperty("QuorumTypeValue")
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

// SetRecentEventsResetTime sets the value of RecentEventsResetTime for the instance
func (instance *MSCluster_Cluster) SetPropertyRecentEventsResetTime(value string) (err error) {
	return instance.SetProperty("RecentEventsResetTime", (value))
}

// GetRecentEventsResetTime gets the value of RecentEventsResetTime for the instance
func (instance *MSCluster_Cluster) GetPropertyRecentEventsResetTime() (value string, err error) {
	retValue, err := instance.GetProperty("RecentEventsResetTime")
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

// SetRequestReplyTimeout sets the value of RequestReplyTimeout for the instance
func (instance *MSCluster_Cluster) SetPropertyRequestReplyTimeout(value uint32) (err error) {
	return instance.SetProperty("RequestReplyTimeout", (value))
}

// GetRequestReplyTimeout gets the value of RequestReplyTimeout for the instance
func (instance *MSCluster_Cluster) GetPropertyRequestReplyTimeout() (value uint32, err error) {
	retValue, err := instance.GetProperty("RequestReplyTimeout")
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

// SetResiliencyDefaultPeriod sets the value of ResiliencyDefaultPeriod for the instance
func (instance *MSCluster_Cluster) SetPropertyResiliencyDefaultPeriod(value uint32) (err error) {
	return instance.SetProperty("ResiliencyDefaultPeriod", (value))
}

// GetResiliencyDefaultPeriod gets the value of ResiliencyDefaultPeriod for the instance
func (instance *MSCluster_Cluster) GetPropertyResiliencyDefaultPeriod() (value uint32, err error) {
	retValue, err := instance.GetProperty("ResiliencyDefaultPeriod")
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

// SetResiliencyLevel sets the value of ResiliencyLevel for the instance
func (instance *MSCluster_Cluster) SetPropertyResiliencyLevel(value uint32) (err error) {
	return instance.SetProperty("ResiliencyLevel", (value))
}

// GetResiliencyLevel gets the value of ResiliencyLevel for the instance
func (instance *MSCluster_Cluster) GetPropertyResiliencyLevel() (value uint32, err error) {
	retValue, err := instance.GetProperty("ResiliencyLevel")
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

// SetResourceAdminExtensions sets the value of ResourceAdminExtensions for the instance
func (instance *MSCluster_Cluster) SetPropertyResourceAdminExtensions(value []string) (err error) {
	return instance.SetProperty("ResourceAdminExtensions", (value))
}

// GetResourceAdminExtensions gets the value of ResourceAdminExtensions for the instance
func (instance *MSCluster_Cluster) GetPropertyResourceAdminExtensions() (value []string, err error) {
	retValue, err := instance.GetProperty("ResourceAdminExtensions")
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

// SetResourceDllDeadlockPeriod sets the value of ResourceDllDeadlockPeriod for the instance
func (instance *MSCluster_Cluster) SetPropertyResourceDllDeadlockPeriod(value uint32) (err error) {
	return instance.SetProperty("ResourceDllDeadlockPeriod", (value))
}

// GetResourceDllDeadlockPeriod gets the value of ResourceDllDeadlockPeriod for the instance
func (instance *MSCluster_Cluster) GetPropertyResourceDllDeadlockPeriod() (value uint32, err error) {
	retValue, err := instance.GetProperty("ResourceDllDeadlockPeriod")
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

// SetResourceTypeAdminExtensions sets the value of ResourceTypeAdminExtensions for the instance
func (instance *MSCluster_Cluster) SetPropertyResourceTypeAdminExtensions(value []string) (err error) {
	return instance.SetProperty("ResourceTypeAdminExtensions", (value))
}

// GetResourceTypeAdminExtensions gets the value of ResourceTypeAdminExtensions for the instance
func (instance *MSCluster_Cluster) GetPropertyResourceTypeAdminExtensions() (value []string, err error) {
	retValue, err := instance.GetProperty("ResourceTypeAdminExtensions")
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

// SetRootMemoryReserved sets the value of RootMemoryReserved for the instance
func (instance *MSCluster_Cluster) SetPropertyRootMemoryReserved(value uint32) (err error) {
	return instance.SetProperty("RootMemoryReserved", (value))
}

// GetRootMemoryReserved gets the value of RootMemoryReserved for the instance
func (instance *MSCluster_Cluster) GetPropertyRootMemoryReserved() (value uint32, err error) {
	retValue, err := instance.GetProperty("RootMemoryReserved")
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

// SetRouteHistoryLength sets the value of RouteHistoryLength for the instance
func (instance *MSCluster_Cluster) SetPropertyRouteHistoryLength(value uint32) (err error) {
	return instance.SetProperty("RouteHistoryLength", (value))
}

// GetRouteHistoryLength gets the value of RouteHistoryLength for the instance
func (instance *MSCluster_Cluster) GetPropertyRouteHistoryLength() (value uint32, err error) {
	retValue, err := instance.GetProperty("RouteHistoryLength")
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

// SetS2DBusTypes sets the value of S2DBusTypes for the instance
func (instance *MSCluster_Cluster) SetPropertyS2DBusTypes(value uint32) (err error) {
	return instance.SetProperty("S2DBusTypes", (value))
}

// GetS2DBusTypes gets the value of S2DBusTypes for the instance
func (instance *MSCluster_Cluster) GetPropertyS2DBusTypes() (value uint32, err error) {
	retValue, err := instance.GetProperty("S2DBusTypes")
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

// SetS2DCacheBehavior sets the value of S2DCacheBehavior for the instance
func (instance *MSCluster_Cluster) SetPropertyS2DCacheBehavior(value uint64) (err error) {
	return instance.SetProperty("S2DCacheBehavior", (value))
}

// GetS2DCacheBehavior gets the value of S2DCacheBehavior for the instance
func (instance *MSCluster_Cluster) GetPropertyS2DCacheBehavior() (value uint64, err error) {
	retValue, err := instance.GetProperty("S2DCacheBehavior")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(uint64)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " uint64 is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = uint64(valuetmp)

	return
}

// SetS2DCacheDesiredState sets the value of S2DCacheDesiredState for the instance
func (instance *MSCluster_Cluster) SetPropertyS2DCacheDesiredState(value uint32) (err error) {
	return instance.SetProperty("S2DCacheDesiredState", (value))
}

// GetS2DCacheDesiredState gets the value of S2DCacheDesiredState for the instance
func (instance *MSCluster_Cluster) GetPropertyS2DCacheDesiredState() (value uint32, err error) {
	retValue, err := instance.GetProperty("S2DCacheDesiredState")
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

// SetS2DCacheDeviceModel sets the value of S2DCacheDeviceModel for the instance
func (instance *MSCluster_Cluster) SetPropertyS2DCacheDeviceModel(value []string) (err error) {
	return instance.SetProperty("S2DCacheDeviceModel", (value))
}

// GetS2DCacheDeviceModel gets the value of S2DCacheDeviceModel for the instance
func (instance *MSCluster_Cluster) GetPropertyS2DCacheDeviceModel() (value []string, err error) {
	retValue, err := instance.GetProperty("S2DCacheDeviceModel")
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

// SetS2DCacheFlashReservePercent sets the value of S2DCacheFlashReservePercent for the instance
func (instance *MSCluster_Cluster) SetPropertyS2DCacheFlashReservePercent(value uint32) (err error) {
	return instance.SetProperty("S2DCacheFlashReservePercent", (value))
}

// GetS2DCacheFlashReservePercent gets the value of S2DCacheFlashReservePercent for the instance
func (instance *MSCluster_Cluster) GetPropertyS2DCacheFlashReservePercent() (value uint32, err error) {
	retValue, err := instance.GetProperty("S2DCacheFlashReservePercent")
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

// SetS2DCacheMetadataReserveBytes sets the value of S2DCacheMetadataReserveBytes for the instance
func (instance *MSCluster_Cluster) SetPropertyS2DCacheMetadataReserveBytes(value uint64) (err error) {
	return instance.SetProperty("S2DCacheMetadataReserveBytes", (value))
}

// GetS2DCacheMetadataReserveBytes gets the value of S2DCacheMetadataReserveBytes for the instance
func (instance *MSCluster_Cluster) GetPropertyS2DCacheMetadataReserveBytes() (value uint64, err error) {
	retValue, err := instance.GetProperty("S2DCacheMetadataReserveBytes")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(uint64)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " uint64 is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = uint64(valuetmp)

	return
}

// SetS2DCachePageSizeKBytes sets the value of S2DCachePageSizeKBytes for the instance
func (instance *MSCluster_Cluster) SetPropertyS2DCachePageSizeKBytes(value uint32) (err error) {
	return instance.SetProperty("S2DCachePageSizeKBytes", (value))
}

// GetS2DCachePageSizeKBytes gets the value of S2DCachePageSizeKBytes for the instance
func (instance *MSCluster_Cluster) GetPropertyS2DCachePageSizeKBytes() (value uint32, err error) {
	retValue, err := instance.GetProperty("S2DCachePageSizeKBytes")
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

// SetS2DEnabled sets the value of S2DEnabled for the instance
func (instance *MSCluster_Cluster) SetPropertyS2DEnabled(value uint32) (err error) {
	return instance.SetProperty("S2DEnabled", (value))
}

// GetS2DEnabled gets the value of S2DEnabled for the instance
func (instance *MSCluster_Cluster) GetPropertyS2DEnabled() (value uint32, err error) {
	retValue, err := instance.GetProperty("S2DEnabled")
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

// SetS2DIOLatencyThreshold sets the value of S2DIOLatencyThreshold for the instance
func (instance *MSCluster_Cluster) SetPropertyS2DIOLatencyThreshold(value uint32) (err error) {
	return instance.SetProperty("S2DIOLatencyThreshold", (value))
}

// GetS2DIOLatencyThreshold gets the value of S2DIOLatencyThreshold for the instance
func (instance *MSCluster_Cluster) GetPropertyS2DIOLatencyThreshold() (value uint32, err error) {
	retValue, err := instance.GetProperty("S2DIOLatencyThreshold")
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

// SetS2DOptimizations sets the value of S2DOptimizations for the instance
func (instance *MSCluster_Cluster) SetPropertyS2DOptimizations(value uint32) (err error) {
	return instance.SetProperty("S2DOptimizations", (value))
}

// GetS2DOptimizations gets the value of S2DOptimizations for the instance
func (instance *MSCluster_Cluster) GetPropertyS2DOptimizations() (value uint32, err error) {
	retValue, err := instance.GetProperty("S2DOptimizations")
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

// SetSameSubnetDelay sets the value of SameSubnetDelay for the instance
func (instance *MSCluster_Cluster) SetPropertySameSubnetDelay(value uint32) (err error) {
	return instance.SetProperty("SameSubnetDelay", (value))
}

// GetSameSubnetDelay gets the value of SameSubnetDelay for the instance
func (instance *MSCluster_Cluster) GetPropertySameSubnetDelay() (value uint32, err error) {
	retValue, err := instance.GetProperty("SameSubnetDelay")
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

// SetSameSubnetThreshold sets the value of SameSubnetThreshold for the instance
func (instance *MSCluster_Cluster) SetPropertySameSubnetThreshold(value uint32) (err error) {
	return instance.SetProperty("SameSubnetThreshold", (value))
}

// GetSameSubnetThreshold gets the value of SameSubnetThreshold for the instance
func (instance *MSCluster_Cluster) GetPropertySameSubnetThreshold() (value uint32, err error) {
	retValue, err := instance.GetProperty("SameSubnetThreshold")
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

// SetSecurity sets the value of Security for the instance
func (instance *MSCluster_Cluster) SetPropertySecurity(value []uint8) (err error) {
	return instance.SetProperty("Security", (value))
}

// GetSecurity gets the value of Security for the instance
func (instance *MSCluster_Cluster) GetPropertySecurity() (value []uint8, err error) {
	retValue, err := instance.GetProperty("Security")
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

// SetSecurity_Descriptor sets the value of Security_Descriptor for the instance
func (instance *MSCluster_Cluster) SetPropertySecurity_Descriptor(value []uint8) (err error) {
	return instance.SetProperty("Security_Descriptor", (value))
}

// GetSecurity_Descriptor gets the value of Security_Descriptor for the instance
func (instance *MSCluster_Cluster) GetPropertySecurity_Descriptor() (value []uint8, err error) {
	retValue, err := instance.GetProperty("Security_Descriptor")
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

// SetSecurityLevel sets the value of SecurityLevel for the instance
func (instance *MSCluster_Cluster) SetPropertySecurityLevel(value uint32) (err error) {
	return instance.SetProperty("SecurityLevel", (value))
}

// GetSecurityLevel gets the value of SecurityLevel for the instance
func (instance *MSCluster_Cluster) GetPropertySecurityLevel() (value uint32, err error) {
	retValue, err := instance.GetProperty("SecurityLevel")
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

// SetSecurityLevelForStorage sets the value of SecurityLevelForStorage for the instance
func (instance *MSCluster_Cluster) SetPropertySecurityLevelForStorage(value uint32) (err error) {
	return instance.SetProperty("SecurityLevelForStorage", (value))
}

// GetSecurityLevelForStorage gets the value of SecurityLevelForStorage for the instance
func (instance *MSCluster_Cluster) GetPropertySecurityLevelForStorage() (value uint32, err error) {
	retValue, err := instance.GetProperty("SecurityLevelForStorage")
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

// SetSharedVolumeCompatibleFilters sets the value of SharedVolumeCompatibleFilters for the instance
func (instance *MSCluster_Cluster) SetPropertySharedVolumeCompatibleFilters(value []string) (err error) {
	return instance.SetProperty("SharedVolumeCompatibleFilters", (value))
}

// GetSharedVolumeCompatibleFilters gets the value of SharedVolumeCompatibleFilters for the instance
func (instance *MSCluster_Cluster) GetPropertySharedVolumeCompatibleFilters() (value []string, err error) {
	retValue, err := instance.GetProperty("SharedVolumeCompatibleFilters")
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

// SetSharedVolumeIncompatibleFilters sets the value of SharedVolumeIncompatibleFilters for the instance
func (instance *MSCluster_Cluster) SetPropertySharedVolumeIncompatibleFilters(value []string) (err error) {
	return instance.SetProperty("SharedVolumeIncompatibleFilters", (value))
}

// GetSharedVolumeIncompatibleFilters gets the value of SharedVolumeIncompatibleFilters for the instance
func (instance *MSCluster_Cluster) GetPropertySharedVolumeIncompatibleFilters() (value []string, err error) {
	retValue, err := instance.GetProperty("SharedVolumeIncompatibleFilters")
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

// SetSharedVolumeSecurityDescriptor sets the value of SharedVolumeSecurityDescriptor for the instance
func (instance *MSCluster_Cluster) SetPropertySharedVolumeSecurityDescriptor(value []uint8) (err error) {
	return instance.SetProperty("SharedVolumeSecurityDescriptor", (value))
}

// GetSharedVolumeSecurityDescriptor gets the value of SharedVolumeSecurityDescriptor for the instance
func (instance *MSCluster_Cluster) GetPropertySharedVolumeSecurityDescriptor() (value []uint8, err error) {
	retValue, err := instance.GetProperty("SharedVolumeSecurityDescriptor")
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

// SetSharedVolumesRoot sets the value of SharedVolumesRoot for the instance
func (instance *MSCluster_Cluster) SetPropertySharedVolumesRoot(value string) (err error) {
	return instance.SetProperty("SharedVolumesRoot", (value))
}

// GetSharedVolumesRoot gets the value of SharedVolumesRoot for the instance
func (instance *MSCluster_Cluster) GetPropertySharedVolumesRoot() (value string, err error) {
	retValue, err := instance.GetProperty("SharedVolumesRoot")
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

// SetSharedVolumeVssWriterOperationTimeout sets the value of SharedVolumeVssWriterOperationTimeout for the instance
func (instance *MSCluster_Cluster) SetPropertySharedVolumeVssWriterOperationTimeout(value uint32) (err error) {
	return instance.SetProperty("SharedVolumeVssWriterOperationTimeout", (value))
}

// GetSharedVolumeVssWriterOperationTimeout gets the value of SharedVolumeVssWriterOperationTimeout for the instance
func (instance *MSCluster_Cluster) GetPropertySharedVolumeVssWriterOperationTimeout() (value uint32, err error) {
	retValue, err := instance.GetProperty("SharedVolumeVssWriterOperationTimeout")
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

// SetShutdownTimeoutInMinutes sets the value of ShutdownTimeoutInMinutes for the instance
func (instance *MSCluster_Cluster) SetPropertyShutdownTimeoutInMinutes(value uint32) (err error) {
	return instance.SetProperty("ShutdownTimeoutInMinutes", (value))
}

// GetShutdownTimeoutInMinutes gets the value of ShutdownTimeoutInMinutes for the instance
func (instance *MSCluster_Cluster) GetPropertyShutdownTimeoutInMinutes() (value uint32, err error) {
	retValue, err := instance.GetProperty("ShutdownTimeoutInMinutes")
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

// SetUseClientAccessNetworksForSharedVolumes sets the value of UseClientAccessNetworksForSharedVolumes for the instance
func (instance *MSCluster_Cluster) SetPropertyUseClientAccessNetworksForSharedVolumes(value uint32) (err error) {
	return instance.SetProperty("UseClientAccessNetworksForSharedVolumes", (value))
}

// GetUseClientAccessNetworksForSharedVolumes gets the value of UseClientAccessNetworksForSharedVolumes for the instance
func (instance *MSCluster_Cluster) GetPropertyUseClientAccessNetworksForSharedVolumes() (value uint32, err error) {
	retValue, err := instance.GetProperty("UseClientAccessNetworksForSharedVolumes")
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

// SetWitnessDatabaseWriteTimeout sets the value of WitnessDatabaseWriteTimeout for the instance
func (instance *MSCluster_Cluster) SetPropertyWitnessDatabaseWriteTimeout(value uint32) (err error) {
	return instance.SetProperty("WitnessDatabaseWriteTimeout", (value))
}

// GetWitnessDatabaseWriteTimeout gets the value of WitnessDatabaseWriteTimeout for the instance
func (instance *MSCluster_Cluster) GetPropertyWitnessDatabaseWriteTimeout() (value uint32, err error) {
	retValue, err := instance.GetProperty("WitnessDatabaseWriteTimeout")
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

// SetWitnessDynamicWeight sets the value of WitnessDynamicWeight for the instance
func (instance *MSCluster_Cluster) SetPropertyWitnessDynamicWeight(value uint32) (err error) {
	return instance.SetProperty("WitnessDynamicWeight", (value))
}

// GetWitnessDynamicWeight gets the value of WitnessDynamicWeight for the instance
func (instance *MSCluster_Cluster) GetPropertyWitnessDynamicWeight() (value uint32, err error) {
	retValue, err := instance.GetProperty("WitnessDynamicWeight")
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

// SetWitnessRestartInterval sets the value of WitnessRestartInterval for the instance
func (instance *MSCluster_Cluster) SetPropertyWitnessRestartInterval(value uint32) (err error) {
	return instance.SetProperty("WitnessRestartInterval", (value))
}

// GetWitnessRestartInterval gets the value of WitnessRestartInterval for the instance
func (instance *MSCluster_Cluster) GetPropertyWitnessRestartInterval() (value uint32, err error) {
	retValue, err := instance.GetProperty("WitnessRestartInterval")
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

// <param name="NewName" type="string "></param>
func (instance *MSCluster_Cluster) Rename( /* IN */ NewName string) (err error) {
	_, err = instance.InvokeMethodWithReturn("Rename", NewName)
	if err != nil {
		return
	}
	return

}

//

// <param name="QuorumPath" type="string "></param>
// <param name="Resource" type="string "></param>
func (instance *MSCluster_Cluster) SetMajorityQuorum( /* IN */ Resource string,
	/* IN */ QuorumPath string) (err error) {
	_, err = instance.InvokeMethodWithReturn("SetMajorityQuorum", Resource, QuorumPath)
	if err != nil {
		return
	}
	return

}

//

// <param name="QuorumPath" type="string "></param>
// <param name="Resource" type="string "></param>
func (instance *MSCluster_Cluster) SetDiskQuorum( /* IN */ Resource string,
	/* IN */ QuorumPath string) (err error) {
	_, err = instance.InvokeMethodWithReturn("SetDiskQuorum", Resource, QuorumPath)
	if err != nil {
		return
	}
	return

}

//
func (instance *MSCluster_Cluster) SetNodeMajorityQuorum() (err error) {
	_, err = instance.InvokeMethodWithReturn("SetNodeMajorityQuorum")
	if err != nil {
		return
	}
	return

}

//

// <param name="ClusterState" type="int32 "></param>
func (instance *MSCluster_Cluster) GetNodeClusterState( /* OUT */ ClusterState int32) (err error) {
	_, err = instance.InvokeMethod("GetNodeClusterState")
	if err != nil {
		return
	}
	return

}

//

// <param name="AdministrativeAccessPoint" type="uint32 "></param>
// <param name="ClusterName" type="string "></param>
// <param name="IPAddresses" type="string []"></param>
// <param name="NodeNames" type="string []"></param>
// <param name="SubnetMasks" type="string []"></param>
func (instance *MSCluster_Cluster) CreateCluster( /* IN */ ClusterName string,
	/* IN */ NodeNames []string,
	/* IN */ IPAddresses []string,
	/* IN */ SubnetMasks []string,
	/* IN */ AdministrativeAccessPoint uint32) (err error) {
	_, err = instance.InvokeMethodWithReturn("CreateCluster", ClusterName, NodeNames, IPAddresses, SubnetMasks, AdministrativeAccessPoint)
	if err != nil {
		return
	}
	return

}

//

// <param name="NodeName" type="string "></param>
func (instance *MSCluster_Cluster) AddNode( /* IN */ NodeName string) (err error) {
	_, err = instance.InvokeMethodWithReturn("AddNode", NodeName)
	if err != nil {
		return
	}
	return

}

//

// <param name="NodeName" type="string "></param>
func (instance *MSCluster_Cluster) EvictNode( /* IN */ NodeName string) (err error) {
	_, err = instance.InvokeMethodWithReturn("EvictNode", NodeName)
	if err != nil {
		return
	}
	return

}

//

// <param name="CleanupAD" type="bool "></param>
func (instance *MSCluster_Cluster) DestroyCluster( /* IN */ CleanupAD bool) (err error) {
	_, err = instance.InvokeMethodWithReturn("DestroyCluster", CleanupAD)
	if err != nil {
		return
	}
	return

}

//

// <param name="NodeName" type="string "></param>
// <param name="Timeout" type="uint32 "></param>
func (instance *MSCluster_Cluster) ForceCleanup( /* IN */ NodeName string,
	/* IN */ Timeout uint32) (err error) {
	_, err = instance.InvokeMethodWithReturn("ForceCleanup", NodeName, Timeout)
	if err != nil {
		return
	}
	return

}

//

// <param name="ControlCode" type="int32 "></param>
// <param name="InputBuffer" type="uint8 []"></param>

// <param name="OutputBuffer" type="uint8 []"></param>
// <param name="OutputBufferSize" type="int32 "></param>
func (instance *MSCluster_Cluster) ExecuteClusterControl( /* IN */ ControlCode int32,
	/* IN */ InputBuffer []uint8,
	/* OUT */ OutputBuffer []uint8,
	/* OUT */ OutputBufferSize int32) (err error) {
	_, err = instance.InvokeMethod("ExecuteClusterControl", ControlCode, InputBuffer)
	if err != nil {
		return
	}
	return

}

//

// <param name="Resource" type="string "></param>
func (instance *MSCluster_Cluster) AddResourceToClusterSharedVolumes( /* IN */ Resource string) (err error) {
	_, err = instance.InvokeMethodWithReturn("AddResourceToClusterSharedVolumes", Resource)
	if err != nil {
		return
	}
	return

}

//

// <param name="Resource" type="string "></param>
func (instance *MSCluster_Cluster) RemoveResourceFromClusterSharedVolumes( /* IN */ Resource string) (err error) {
	_, err = instance.InvokeMethodWithReturn("RemoveResourceFromClusterSharedVolumes", Resource)
	if err != nil {
		return
	}
	return

}

//

// <param name="Status" type="MSCluster_ValidationStatus "></param>
func (instance *MSCluster_Cluster) GenerateValidationStatus( /* OUT */ Status MSCluster_ValidationStatus) (err error) {
	_, err = instance.InvokeMethod("GenerateValidationStatus")
	if err != nil {
		return
	}
	return

}

//

// <param name="VirtualMachine" type="string "></param>
func (instance *MSCluster_Cluster) AddVirtualMachine( /* IN */ VirtualMachine string) (err error) {
	_, err = instance.InvokeMethodWithReturn("AddVirtualMachine", VirtualMachine)
	if err != nil {
		return
	}
	return

}

//

// <param name="Group" type="string "></param>
// <param name="Path" type="string "></param>

// <param name="result" type="uint32 "></param>
func (instance *MSCluster_Cluster) VerifyPath( /* IN */ Path string,
	/* IN */ Group string,
	/* OUT */ result uint32) (err error) {
	_, err = instance.InvokeMethod("VerifyPath", Path, Group)
	if err != nil {
		return
	}
	return

}

//

// <param name="AdministrativeAccessPoint" type="uint32 "></param>
// <param name="DomainName" type="string "></param>
// <param name="Name" type="string "></param>
// <param name="Password" type="string "></param>
// <param name="UserName" type="string "></param>
func (instance *MSCluster_Cluster) AddClusterNameAccount( /* IN */ Name string,
	/* IN */ DomainName string,
	/* IN */ UserName string,
	/* IN */ Password string,
	/* IN */ AdministrativeAccessPoint uint32) (err error) {
	_, err = instance.InvokeMethodWithReturn("AddClusterNameAccount", Name, DomainName, UserName, Password, AdministrativeAccessPoint)
	if err != nil {
		return
	}
	return

}
