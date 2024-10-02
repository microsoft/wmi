// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 6/14/2024 using wmigen
//      Source ManagementOperations_Status
//////////////////////////////////////////////
package wmi

// ManagementOperations_Status
type ManagementOperations_Status int

const (
	// Success enum
	ManagementOperations_Status_Success ManagementOperations_Status = 0
	// Non_Specific_Error enum
	ManagementOperations_Status_Non_Specific_Error ManagementOperations_Status = 1
	// Login_Failed enum
	ManagementOperations_Status_Login_Failed ManagementOperations_Status = 2
	// Connection_Failed enum
	ManagementOperations_Status_Connection_Failed ManagementOperations_Status = 3
	// Initiator_Node_Already_Exists enum
	ManagementOperations_Status_Initiator_Node_Already_Exists ManagementOperations_Status = 4
	// Initiator_Node_Does_Not_Exist enum
	ManagementOperations_Status_Initiator_Node_Does_Not_Exist ManagementOperations_Status = 5
	// Target_Moved_Temporarily enum
	ManagementOperations_Status_Target_Moved_Temporarily ManagementOperations_Status = 6
	// Target_Moved_Permamently enum
	ManagementOperations_Status_Target_Moved_Permamently ManagementOperations_Status = 7
	// Initiator_Error enum
	ManagementOperations_Status_Initiator_Error ManagementOperations_Status = 8
	// Authentication_Failure enum
	ManagementOperations_Status_Authentication_Failure ManagementOperations_Status = 9
	// Authorization_Failure enum
	ManagementOperations_Status_Authorization_Failure ManagementOperations_Status = 10
	// Not_Found enum
	ManagementOperations_Status_Not_Found ManagementOperations_Status = 11
	// Target_Removed enum
	ManagementOperations_Status_Target_Removed ManagementOperations_Status = 12
	// Unsupported_Version enum
	ManagementOperations_Status_Unsupported_Version ManagementOperations_Status = 13
	// Too_many_Connections enum
	ManagementOperations_Status_Too_many_Connections ManagementOperations_Status = 14
	// Missing_Parameter enum
	ManagementOperations_Status_Missing_Parameter ManagementOperations_Status = 15
	// Can_not_include_in_session enum
	ManagementOperations_Status_Can_not_include_in_session ManagementOperations_Status = 16
	// Session_type_not_supported enum
	ManagementOperations_Status_Session_type_not_supported ManagementOperations_Status = 17
	// Target_Error enum
	ManagementOperations_Status_Target_Error ManagementOperations_Status = 18
	// Service_Unavailable enum
	ManagementOperations_Status_Service_Unavailable ManagementOperations_Status = 19
	// Out_of_Resources enum
	ManagementOperations_Status_Out_of_Resources ManagementOperations_Status = 20
	// Connections_already_exist_on_initiator_node enum
	ManagementOperations_Status_Connections_already_exist_on_initiator_node ManagementOperations_Status = 21
	// Session_Already_Exists enum
	ManagementOperations_Status_Session_Already_Exists ManagementOperations_Status = 22
	// Initiator_Instance_Does_Not_Exist enum
	ManagementOperations_Status_Initiator_Instance_Does_Not_Exist ManagementOperations_Status = 23
	// Target_Already_Exists enum
	ManagementOperations_Status_Target_Already_Exists ManagementOperations_Status = 24
	// The_iscsi_driver_implementation_did_not_complete_an_operation_correctly enum
	ManagementOperations_Status_The_iscsi_driver_implementation_did_not_complete_an_operation_correctly ManagementOperations_Status = 25
	// An_invalid_key_text_was_encountered enum
	ManagementOperations_Status_An_invalid_key_text_was_encountered ManagementOperations_Status = 26
	// Invalid_SendTargets_response_text_was_encountered enum
	ManagementOperations_Status_Invalid_SendTargets_response_text_was_encountered ManagementOperations_Status = 27
	// Invalid_Session_Id enum
	ManagementOperations_Status_Invalid_Session_Id ManagementOperations_Status = 28
	// The_scsi_request_failed enum
	ManagementOperations_Status_The_scsi_request_failed ManagementOperations_Status = 29
	// Exceeded_max_sessions_for_this_initiator_ enum
	ManagementOperations_Status_Exceeded_max_sessions_for_this_initiator_ ManagementOperations_Status = 30
	// Session_is_busy_since_a_request_is_already_in_progress_ enum
	ManagementOperations_Status_Session_is_busy_since_a_request_is_already_in_progress_ ManagementOperations_Status = 31
	// The_target_mapping_is_not_available enum
	ManagementOperations_Status_The_target_mapping_is_not_available ManagementOperations_Status = 32
	// The_Target_Address_type_given_is_not_supported enum
	ManagementOperations_Status_The_Target_Address_type_given_is_not_supported ManagementOperations_Status = 33
	// Logon_Failed enum
	ManagementOperations_Status_Logon_Failed ManagementOperations_Status = 34
	// TCP_Send_Failed enum
	ManagementOperations_Status_TCP_Send_Failed ManagementOperations_Status = 35
	// TCP_Transport_Error enum
	ManagementOperations_Status_TCP_Transport_Error ManagementOperations_Status = 36
	// iSCSI_Version_Mismatch enum
	ManagementOperations_Status_iSCSI_Version_Mismatch ManagementOperations_Status = 37
	// The_Target_Mapping_Address_passed_is_out_of_range_for_the_adapter_configuration enum
	ManagementOperations_Status_The_Target_Mapping_Address_passed_is_out_of_range_for_the_adapter_configuration ManagementOperations_Status = 38
	// The_preshared_key_for_the_target_or_IKE_identification_payload_is_not_available_ enum
	ManagementOperations_Status_The_preshared_key_for_the_target_or_IKE_identification_payload_is_not_available_ ManagementOperations_Status = 39
	// The_authentication_information_for_the_target_is_not_available enum
	ManagementOperations_Status_The_authentication_information_for_the_target_is_not_available ManagementOperations_Status = 40
	// The_target_name_is_not_found_or_is_marked_as_hidden_from_login_ enum
	ManagementOperations_Status_The_target_name_is_not_found_or_is_marked_as_hidden_from_login_ ManagementOperations_Status = 41
	// One_or_more_parameters_specified_in_LoginTargetIN_structure_is_invalid_ enum
	ManagementOperations_Status_One_or_more_parameters_specified_in_LoginTargetIN_structure_is_invalid_ ManagementOperations_Status = 42
	// Given_target_mapping_already_exists_ enum
	ManagementOperations_Status_Given_target_mapping_already_exists_ ManagementOperations_Status = 43
	// The_HBA_security_information_cache_is_full_ enum
	ManagementOperations_Status_The_HBA_security_information_cache_is_full_ ManagementOperations_Status = 44
	// The_port_number_passed_is_not_valid_for_the_initiator_ enum
	ManagementOperations_Status_The_port_number_passed_is_not_valid_for_the_initiator_ ManagementOperations_Status = 45
	// Operation_was_not_successful_for_all_initiators_ enum
	ManagementOperations_Status_Operation_was_not_successful_for_all_initiators_ ManagementOperations_Status = 46
	// The_HBA_security_information_cache_is_not_supported_by_this_adapter_ enum
	ManagementOperations_Status_The_HBA_security_information_cache_is_not_supported_by_this_adapter_ ManagementOperations_Status = 47
	// The_IKE_id_payload_type_specified_is_not_supported_ enum
	ManagementOperations_Status_The_IKE_id_payload_type_specified_is_not_supported_ ManagementOperations_Status = 48
	// The_IKE_id_payload_size_specified_is_not_correct_ enum
	ManagementOperations_Status_The_IKE_id_payload_size_specified_is_not_correct_ ManagementOperations_Status = 49
	// Target_Portal_Structure_Already_Exists_ enum
	ManagementOperations_Status_Target_Portal_Structure_Already_Exists_ ManagementOperations_Status = 50
	// Target_Address_Structure_Already_Exists_ enum
	ManagementOperations_Status_Target_Address_Structure_Already_Exists_ ManagementOperations_Status = 51
	// There_is_no_IKE_authentication_information_available_ enum
	ManagementOperations_Status_There_is_no_IKE_authentication_information_available_ ManagementOperations_Status = 52
	// There_is_no_tunnel_mode_outer_address_specified_ enum
	ManagementOperations_Status_There_is_no_tunnel_mode_outer_address_specified_ ManagementOperations_Status = 53
	// Authentication_or_tunnel_address_cache_is_corrupted_ enum
	ManagementOperations_Status_Authentication_or_tunnel_address_cache_is_corrupted_ ManagementOperations_Status = 54
	// The_request_or_operation_is_not_supported_ enum
	ManagementOperations_Status_The_request_or_operation_is_not_supported_ ManagementOperations_Status = 55
	// The_target_does_not_have_enough_resources_to_process_the_given_request_ enum
	ManagementOperations_Status_The_target_does_not_have_enough_resources_to_process_the_given_request_ ManagementOperations_Status = 56
	// The_initiator_service_did_not_respond_to_the_request_sent_by_the_driver_ enum
	ManagementOperations_Status_The_initiator_service_did_not_respond_to_the_request_sent_by_the_driver_ ManagementOperations_Status = 57
	// The_iSNS_server_was_not_found_or_is_unavailable_ enum
	ManagementOperations_Status_The_iSNS_server_was_not_found_or_is_unavailable_ ManagementOperations_Status = 58
	// The_operation_was_successful_but_requires_a_driver_reload_or_reboot_to_become_effective_ enum
	ManagementOperations_Status_The_operation_was_successful_but_requires_a_driver_reload_or_reboot_to_become_effective_ ManagementOperations_Status = 59
	// There_is_no_target_portal_available_to_complete_the_login_ enum
	ManagementOperations_Status_There_is_no_target_portal_available_to_complete_the_login_ ManagementOperations_Status = 60
	// Cannot_remove_the_last_connection_for_a_session_ enum
	ManagementOperations_Status_Cannot_remove_the_last_connection_for_a_session_ ManagementOperations_Status = 61
	// The_Microsoft_iSCSI_Initiator_Service_is_not_running__Please_start_the_service_and_retry_ enum
	ManagementOperations_Status_The_Microsoft_iSCSI_Initiator_Service_is_not_running__Please_start_the_service_and_retry_ ManagementOperations_Status = 62
	// The_target_has_already_been_logged_in_via_an_iSCSI_session_ enum
	ManagementOperations_Status_The_target_has_already_been_logged_in_via_an_iSCSI_session_ ManagementOperations_Status = 63
	// The_session_cannot_be_logged_out_since_a_device_on_that_session_is_currently_being_used_ enum
	ManagementOperations_Status_The_session_cannot_be_logged_out_since_a_device_on_that_session_is_currently_being_used_ ManagementOperations_Status = 64
	// Failed_to_save_persistent_login_information_ enum
	ManagementOperations_Status_Failed_to_save_persistent_login_information_ ManagementOperations_Status = 65
	// Failed_to_remove_persistent_login_information_ enum
	ManagementOperations_Status_Failed_to_remove_persistent_login_information_ ManagementOperations_Status = 66
	// The_specified_initiator_name_was_not_found_ enum
	ManagementOperations_Status_The_specified_initiator_name_was_not_found_ ManagementOperations_Status = 67
	// The_specified_portal_was_not_found_ enum
	ManagementOperations_Status_The_specified_portal_was_not_found_ ManagementOperations_Status = 68
	// The_specified_discovery_mechanism_was_not_found_ enum
	ManagementOperations_Status_The_specified_discovery_mechanism_was_not_found_ ManagementOperations_Status = 69
	// iSCSI_does_not_support_IPSEC_for_this_version_of_the_OS_ enum
	ManagementOperations_Status_iSCSI_does_not_support_IPSEC_for_this_version_of_the_OS_ ManagementOperations_Status = 70
	// The_iSCSI_service_timed_out_waiting_for_all_persistent_logins_to_complete_ enum
	ManagementOperations_Status_The_iSCSI_service_timed_out_waiting_for_all_persistent_logins_to_complete_ ManagementOperations_Status = 71
	// The_specified_CHAP_secret_is_less_than_96_bits_and_will_not_be_usable_for_authenticating_over_non_ipsec_connections_ enum
	ManagementOperations_Status_The_specified_CHAP_secret_is_less_than_96_bits_and_will_not_be_usable_for_authenticating_over_non_ipsec_connections_ ManagementOperations_Status = 72
	// The_evaluation_period_for_the_iSCSI_initiator_service_has_expired_ enum
	ManagementOperations_Status_The_evaluation_period_for_the_iSCSI_initiator_service_has_expired_ ManagementOperations_Status = 73
	// CHAP_secret_given_does_not_conform_to_the_standard__Please_see_system_event_log_for_more_information_ enum
	ManagementOperations_Status_CHAP_secret_given_does_not_conform_to_the_standard__Please_see_system_event_log_for_more_information_ ManagementOperations_Status = 74
	// Target_CHAP_secret_given_is_invalid_ enum
	ManagementOperations_Status_Target_CHAP_secret_given_is_invalid_ ManagementOperations_Status = 75
	// Initiator_CHAP_secret_given_is_invalid_ enum
	ManagementOperations_Status_Initiator_CHAP_secret_given_is_invalid_ ManagementOperations_Status = 76
	// CHAP_Username_given_is_invalid_ enum
	ManagementOperations_Status_CHAP_Username_given_is_invalid_ ManagementOperations_Status = 77
	// Logon_Authentication_type_given_is_invalid_ enum
	ManagementOperations_Status_Logon_Authentication_type_given_is_invalid_ ManagementOperations_Status = 78
	// Target_Mapping_information_given_is_invalid_ enum
	ManagementOperations_Status_Target_Mapping_information_given_is_invalid_ ManagementOperations_Status = 79
	// Target_Id_given_in_Target_Mapping_is_invalid_ enum
	ManagementOperations_Status_Target_Id_given_in_Target_Mapping_is_invalid_ ManagementOperations_Status = 80
	// The_iSCSI_name_specified_contains_invalid_characters_or_is_too_long_ enum
	ManagementOperations_Status_The_iSCSI_name_specified_contains_invalid_characters_or_is_too_long_ ManagementOperations_Status = 81
	// The_iSNS_version_number_returned_from_the_iSNS_server_is_not_compatible_with_this_version_of_the_iSNS_client_ enum
	ManagementOperations_Status_The_iSNS_version_number_returned_from_the_iSNS_server_is_not_compatible_with_this_version_of_the_iSNS_client_ ManagementOperations_Status = 82
	// Initiator_failed_to_configure_IPSec_for_the_given_connection__This_could_be_because_of_low_resources_ enum
	ManagementOperations_Status_Initiator_failed_to_configure_IPSec_for_the_given_connection__This_could_be_because_of_low_resources_ ManagementOperations_Status = 83
	// The_buffer_given_for_processing_the_request_is_too_small_ enum
	ManagementOperations_Status_The_buffer_given_for_processing_the_request_is_too_small_ ManagementOperations_Status = 84
	// The_given_Load_Balance_policy_is_not_recognized_by_iScsi_initiator_ enum
	ManagementOperations_Status_The_given_Load_Balance_policy_is_not_recognized_by_iScsi_initiator_ ManagementOperations_Status = 85
	// One_or_more_paramaters_specified_is_not_valid_ enum
	ManagementOperations_Status_One_or_more_paramaters_specified_is_not_valid_ ManagementOperations_Status = 86
	// Duplicate_PathIds_were_specified_in_the_call_to_set_Load_Balance_Policy_ enum
	ManagementOperations_Status_Duplicate_PathIds_were_specified_in_the_call_to_set_Load_Balance_Policy_ ManagementOperations_Status = 87
	// Number_of_paths_specified_in_Set_Load_Balance_Policy_does_not_match_the_number_of_paths_to_the_target_ enum
	ManagementOperations_Status_Number_of_paths_specified_in_Set_Load_Balance_Policy_does_not_match_the_number_of_paths_to_the_target_ ManagementOperations_Status = 88
	// Path_Id_specified_in_the_call_to_set_Load_Balance_Policy_is_not_valid_ enum
	ManagementOperations_Status_Path_Id_specified_in_the_call_to_set_Load_Balance_Policy_is_not_valid_ ManagementOperations_Status = 89
	// Multiple_primary_paths_specified_when_only_one_primary_path_is_expected_ enum
	ManagementOperations_Status_Multiple_primary_paths_specified_when_only_one_primary_path_is_expected_ ManagementOperations_Status = 90
	// No_primary_path_specified_when_at_least_one_is_expected_ enum
	ManagementOperations_Status_No_primary_path_specified_when_at_least_one_is_expected_ ManagementOperations_Status = 91
	// Volume_is_already_a_persistently_bound_volume_ enum
	ManagementOperations_Status_Volume_is_already_a_persistently_bound_volume_ ManagementOperations_Status = 92
	// Volume_was_not_found_ enum
	ManagementOperations_Status_Volume_was_not_found_ ManagementOperations_Status = 93
	// The_volume_specified_does_not_originate_from_an_iSCSI_disk_ enum
	ManagementOperations_Status_The_volume_specified_does_not_originate_from_an_iSCSI_disk_ ManagementOperations_Status = 94
	// The_DNS_name_specified_was_not_resolved_ enum
	ManagementOperations_Status_The_DNS_name_specified_was_not_resolved_ ManagementOperations_Status = 95
	// There_is_no_connection_available_in_the_iSCSI_session_to_process_the_request_ enum
	ManagementOperations_Status_There_is_no_connection_available_in_the_iSCSI_session_to_process_the_request_ ManagementOperations_Status = 96
	// The_given_Load_Balance_policy_is_not_supported_ enum
	ManagementOperations_Status_The_given_Load_Balance_policy_is_not_supported_ ManagementOperations_Status = 97
	// A_remove_connection_request_is_already_in_progress_for_this_session_ enum
	ManagementOperations_Status_A_remove_connection_request_is_already_in_progress_for_this_session_ ManagementOperations_Status = 98
	// Given_connection_was_not_found_in_the_session_ enum
	ManagementOperations_Status_Given_connection_was_not_found_in_the_session_ ManagementOperations_Status = 99
	// The_leading_connection_in_the_session_cannot_be_removed_ enum
	ManagementOperations_Status_The_leading_connection_in_the_session_cannot_be_removed_ ManagementOperations_Status = 100
)
