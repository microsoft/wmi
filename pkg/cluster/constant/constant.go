// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.

package constant

const (
	STATUS_OK          string = "OK"
	STATUS_ERROR       string = "Error"
	STATUS_DEGRADED    string = "Degraded"
	STATUS_UNKNOWN     string = "Unknown"
	STATUS_PRED_FAIL   string = "Pred Fail"
	STATUS_STARTING    string = "Starting"
	STATUS_STOPPING    string = "Stopping"
	STATUS_SERVICE     string = "Service"
	STATUS_STRESSED    string = "Stressed"
	STATUS_NON_RECOVER string = "NonRecover"
)

const (
	FAULT_STATE_NO_FAULT          int32 = 0
	FAULT_STATE_REDIRECTED_ACCESS int32 = 1
	FAULT_STATE_NO_ACCESS         int32 = 2
	FAULT_STATE_IN_MAINTENANCE    int32 = 4
)

const (
	ClusterStateNotInstalled  int32 = 0
	ClusterStateNotConfigured int32 = 1
	ClusterStateNotRunning    int32 = 3
	ClusterStateRunning       int32 = 19
)

const (
	CLUSTER_NODE_STATE_UNKNOWN int32 = -1
	CLUSTER_NODE_STATE_UP      int32 = 0
	CLUSTER_NODE_STATE_DOWN    int32 = 1
	CLUSTER_NODE_STATE_PAUSED  int32 = 2
	CLUSTER_NODE_STATE_JOINING int32 = 3
)

const (
	CLUSTER_NETWORK_STATE_UNKNOWN int32 = -1
	CLUSTER_NETWORK_STATE_UNAVAILABLE int32 = 0
	CLUSTER_NETWORK_STATE_DOWN int32 = 1
	CLUSTER_NETWORK_STATE_PARTITIONED int32 = 2
	CLUSTER_NETWORK_STATE_UP int32 = 3
)