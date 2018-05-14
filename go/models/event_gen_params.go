package models

// This file is auto-generated.
// Please contact avi-sdk@avinetworks.com for any change requests.

// EventGenParams event gen params
// swagger:model EventGenParams
type EventGenParams struct {

	// cluster of EventGenParams.
	Cluster string `json:"cluster,omitempty"`

	//  Enum options - VINFRA_DISC_DC, VINFRA_DISC_HOST, VINFRA_DISC_CLUSTER, VINFRA_DISC_VM, VINFRA_DISC_NW, MGMT_NW_NAME_CHANGED, DISCOVERY_DATACENTER_DEL, VM_ADDED, VM_REMOVED, VINFRA_DISC_COMPLETE, VCENTER_ADDRESS_ERROR, SE_GROUP_CLUSTER_DEL, SE_GROUP_MGMT_NW_DEL, MGMT_NW_DEL, VCENTER_BAD_CREDENTIALS, ESX_HOST_UNREACHABLE, SERVER_DELETED, SE_GROUP_HOST_DEL, VINFRA_DISC_FAILURE, ESX_HOST_POWERED_DOWN, VCENTER_VERSION_NOT_SUPPORTED, VCENTER_CONNECTIVITY_FAIL, VCENTER_CONNECTIVITY_SUCCESS, VCENTER_ACCESS_SLOW, SE_FATAL_ERROR, SE_HEARTBEAT_FAILURE, SE_MARKED_DOWN, SE_VM_DELETED, SE_VM_PURGED, SE_UP, SE_POWERED_DOWN, SE_REBOOTED, SE_HEALTH_CHECK_FAIL, SE_EXTERNAL_HM_RESTART, SE_DOWN, SE_VERSION_CHECK_FAILED, SE_UPGRADING, SE_ENABLE, SE_MIGRATE, CREATING_SE, CREATED_SE, CREATE_SE_FAIL, CREATE_SE_TIMEOUT, DELETING_SE, DELETED_SE, DELETE_SE_FAIL, ADD_NW_SE, DEL_NW_SE, VS_ADD_SE_INT, VS_REMOVED_SE_INT, VS_ADD_SE, VS_REMOVED_SE, ADD_NW_FAIL, RM_DEL_NETWORK_FAIL, REBOOT_SE, MODIFY_NW, MODIFY_NW_FAIL, VS_SE_BOOTUP_FAIL, VS_SE_IP_FAIL, NO_HOST_AVAIL, VS_SWITCHOVER, VS_SWITCHOVER_FAIL, ADD_VIP_VNIC, DEL_VIP_VNIC, VS_FSM_INACTIVE, VS_FSM_AWAITING_SE_ASSIGNMENT, VS_FSM_ACTIVE, VS_FSM_ACTIVE_AWAITING_SE_TRANSITION, VS_FSM_DISABLED, NEW_PROBABLE_SRVR, VS_SCALEOUT_DONE, VS_SCALEOUT_DONE_AWAITING_MORE_SE, VS_SCALEOUT_ERR, VS_SCALEIN_DONE, VS_SCALEIN_DONE_AWAITING_MORE_SE, VS_SCALEIN_ERR, VS_MIGRATE_SCALEOUT_DONE, VS_MIGRATE_SCALEOUT_ERROR, VS_MIGRATE_SCALEIN_DONE, VS_MIGRATE_SCALEIN_ERROR, VS_MIGRATE_DONE, VS_FSM_UNEXPECTED_EVENT, VS_RPC_TO_RESMGR_FAILED_EVENT, VS_RPC_TO_SE_FAILED_EVENT, VS_RPC_FAILED_EVENT, VS_SCALEOUT_COMPLETE, VS_SCALEIN_COMPLETE, VS_MIGRATE_STARTED, VS_MIGRATE_COMPLETE, VS_SCALEOUT_FAILED, VS_SCALEIN_FAILED, VS_MIGRATE_FAILED, VS_AWAITING_SE, VS_INITIAL_PLACEMENT_FAILED, VS_FSM_ACTIVE_AWAITING_SCALEOUT_READY, UPGRADE_ALL_SE_START, UPGRADE_ALL_SE_DONE, UPGRADE_ALL_SE_NOT_NEEDED, UPGRADE_SE_START, UPGRADE_SE_DONE, UPGRADE_SE_NOT_NEEDED, UPGRADE_SE_SUSPENDED, UPGRADE_SE_VS_SCALEOUT, UPGRADE_SE_VS_SCALEIN, UPGRADE_SE_VS_MIGRATE, UPGRADE_SE_VS_DISRUPTED, REBALANCE_VS_SCALEOUT, REBALANCE_VS_SCALEIN, REBALANCE_VS_MIGRATE, DISABLE_SE_VS_MIGRATE, ROLLBACK_ALL_SE_START, ROLLBACK_ALL_SE_DONE, MIGRATE_SE_STARTED, MIGRATE_SE_RESTARTED, MIGRATE_SE_FINISHED, MIGRATE_SE_FAILED, MIGRATE_SE_VS_MIGRATE_STARTED, MIGRATE_SE_VS_MIGRATE_FINISHED, MIGRATE_SE_VS_MIGRATE_FAILED, SE_HM_EVENT_SHM_DOWN, SE_HM_EVENT_SHM_UP, SERVER_DOWN, SERVER_UP, POOL_DOWN, POOL_UP, VS_DOWN, VS_UP, SE_SERVER_DELETED, SE_SERVER_DISABLED, SE_POOL_DELETED, SE_SERVER_APP_CHANGED, VS_CONN_LIMIT, VS_THROUGHPUT_LIMIT, CONN_DROP_MAX_SYN_TBL, CONN_DROP_MAX_FLOW_TBL, CONN_DROP_MAX_PERSIST_TBL, CONN_DROP_POOL_LB_FAILURE, CONN_DROP_NO_CONN_MEM, CONN_DROP_NO_PKT_BUFF, PKT_DROP_NO_PKT_BUFF, PKT_BUFF_ALLOC_FAIL, CACHE_OBJ_ALLOC_FAIL, SE_CPU_HIGH, SE_MEM_HIGH, SE_PKT_BUFF_HIGH, SE_PERSIST_TBL_HIGH, SE_CONN_MEM_HIGH, SE_DISK_HIGH, SE_FLOW_TBL_HIGH, SE_SYN_TBL_HIGH, SE_DP_HB_FAILED, SE_VNIC_DHCP_IP_ALLOC_FAILURE, SE_VNIC_DUPLICATE_IP, SE_SYN_CACHE_USAGE_HIGH, VS_SE_HA_ACTIVE, VS_SE_HA_COMPROMISED, POOL_SE_HA_ACTIVE, POOL_SE_HA_COMPROMISED, SERVER_DOWN_HA_COMPROMISED, SERVER_UP_HA_ACTIVE, SE_VNIC_IP_ADDED, SE_VNIC_IP_REMOVED, GS_MEMBER_DOWN, GS_MEMBER_UP, GS_GROUP_DOWN, GS_GROUP_UP, GS_DOWN, GS_UP, VIP_DOWN, VIP_UP, SE_GEO_DB_FAILURE, VS_GEO_DB_FAILURE, SE_GEO_DB_SUCCESS, VS_GEO_DB_SUCCESS, SE_EV_SERVER_DOWN, SE_EV_SERVER_UP, SE_EV_POOL_DOWN, SE_EV_POOL_UP, SE_EV_VS_DOWN, SE_EV_VS_UP, SE_HM_EVENT_GHM_DOWN, SE_HM_EVENT_GHM_UP, SE_EV_GS_GROUP_DELETED, SE_EV_GS_MEMBER_DOWN, SE_EV_GS_MEMBER_UP, SE_EV_GS_GROUP_DOWN, SE_EV_GS_GROUP_UP, SE_EV_GS_DOWN, SE_EV_GS_UP, CONFIG_CREATE, CONFIG_UPDATE, CONFIG_DELETE, USER_LOGIN, USER_LOGOUT, CONFIG_ACTION, CONFIG_INTERNAL_CREATE, CONFIG_INTERNAL_UPDATE, USER_PASSWORD_CHANGE_REQUEST, USER_AUTHORIZED_BY_RULE, USER_NOT_AUTHORIZED_BY_ANY_RULE, CONFIG_SE_GRP_FLAVOR_UPDATE, SSL_CERT_EXPIRE, SSL_KEY_EXPORTED, SSL_CERT_RENEW, SSL_CERT_RENEW_FAILED, CONTROLLER_NODE_JOINED, CONTROLLER_NODE_LEFT, CONTROLLER_SERVICE_FAILURE, CONTROLLER_LEADER_FAILOVER, CONTROLLER_WARM_REBOOT, CONTROLLER_SERVICE_RESTORED, CONTROLLER_SERVICE_CRITICAL_FAILURE, CONTROLLER_NODE_SHUTDOWN, CONTROLLER_NODE_STARTED, CLUSTER_CONFIG_FAILED, SYSTEM_UPGRADE_STARTED, SYSTEM_UPGRADE_COMPLETE, SYSTEM_UPGRADE_ABORTED, SYSTEM_ROLLBACK_STARTED, SYSTEM_ROLLBACK_COMPLETE, SYSTEM_ROLLBACK_ABORTED, CONTROLLER_NODE_DB_REPLICATION_FAILED, CONTROLLER_PROCESS_STOPPED, CONTROLLER_MEMORY_BALANCER_DISABLED, METRIC_THRESHOLD_UP_VIOLATION, LICENSE_EXPIRY, ANOMALY, LICENSE_ADDITION_NOTIF, LICENSE_REMOVAL_NOTIF, METRICS_DB_DISK_FULL, OPENSTACK_ACCESS_FAILURE, OPENSTACK_ACCESS_SUCCESS, OPENSTACK_IMAGE_UPLOAD_FAILURE, OPENSTACK_IMAGE_UPLOAD_SUCCESS, OPENSTACK_SE_VM_CREATED, OPENSTACK_SE_VM_DELETED, OPENSTACK_SE_VM_DELETION_DETECTED, OPENSTACK_VNIC_ADDED, OPENSTACK_VNIC_REMOVED, OPENSTACK_IP_DETACHED, OPENSTACK_IP_ATTACHED, OPENSTACK_SE_CREATION_FAILURE, OPENSTACK_SE_DELETION_FAILURE, OPENSTACK_VNIC_ADDITION_FAILURE, OPENSTACK_VNIC_DELETION_FAILURE, OPENSTACK_IP_DETACH_FAILURE, OPENSTACK_IP_ATTACH_FAILURE, OPENSTACK_LBPROV_AUDIT_FAILURE, OPENSTACK_LBPROV_AUDIT_SUCCESS, OPENSTACK_LBPLUGIN_OP_FAILURE, OPENSTACK_LBPLUGIN_OP_SUCCESS, OPENSTACK_SYNC_SERVICES_SUCCESS, OPENSTACK_SYNC_SERVICES_FAILURE, OPENSTACK_TENANTS_DELETED, AWS_ACCESS_FAILURE, AWS_ACCESS_SUCCESS, AWS_IMAGE_UPLOAD_FAILURE, AWS_IMAGE_UPLOAD_SUCCESS, AWS_SNS_ACCESS_FAILURE, AWS_SNS_ACCESS_SUCCESS, AWS_SQS_ACCESS_FAILURE, AWS_SQS_ACCESS_SUCCESS, AWS_ASG_PUT_NOTIFICATION_CONFIGURATION_FAILURE, AWS_ASG_PUT_NOTIFICATION_CONFIGURATION_SUCCESS, AWS_ASG_DELETE_NOTIFICATION_CONFIGURATION_FAILURE, AWS_ASG_DELETE_NOTIFICATION_CONFIGURATION_SUCCESS, AWS_ASG_NOTIFICATION_PROCESSING_FAILURE, AWS_ASG_NOTIFICATION_PROCESSING_SUCCESS, AWS_ASG_NOTIFICATION_INSTANCE_ADDED, AWS_ASG_NOTIFICATION_INSTANCE_REMOVED, AWS_ASG_ACCESS_FAILURE, AWS_ASG_ACCESS_SUCCESS, AWS_ASG_NOTIFICATION_INSTANCE_LAUNCH_ERROR, AWS_ASG_NOTIFICATION_INSTANCE_TERMINATE_ERROR, AWS_ASG_NOTIFICATION_AUTOSCALE_GROUP_DELETED, CLOUDSTACK_ACCESS_FAILURE, CLOUDSTACK_ACCESS_SUCCESS, CLOUDSTACK_IMAGE_UPLOAD_FAILURE, CLOUDSTACK_IMAGE_UPLOAD_SUCCESS, DOCKER_UCP_ACCESS_SUCCESS, DOCKER_UCP_ACCESS_FAILURE, DOCKER_UCP_IMAGE_UPLOAD_FAILURE, DOCKER_UCP_IMAGE_UPLOAD_SUCCESS, DOCKER_UCP_IMAGE_UPLOAD_IN_PROGRESS, VCA_ACCESS_FAILURE, VCA_ACCESS_SUCCESS, VCA_IMAGE_UPLOAD_FAILURE, VCA_IMAGE_UPLOAD_SUCCESS, LS_ACCESS_FAILURE, LS_ACCESS_SUCCESS, LS_IMAGE_UPLOAD_FAILURE, LS_IMAGE_UPLOAD_SUCCESS, MESOS_ACCESS_SUCCESS, MESOS_ACCESS_FAILURE, MESOS_IMAGE_UPLOAD_FAILURE, MESOS_IMAGE_UPLOAD_SUCCESS, MESOS_IMAGE_UPLOAD_IN_PROGRESS, MESOS_CREATED_SE, MESOS_CREATE_SE_FAIL, MESOS_DELETED_SE, MESOS_DELETE_SE_FAIL, MESOS_STOPPED_SE, MESOS_STOP_SE_FAIL, MESOS_STARTED_SE, MESOS_START_SE_FAIL, MESOS_UPDATED_HOSTS, CC_SE_CREATED, CC_SE_CREATION_FAILURE, CC_SE_DELETED, CC_SE_DELETION_FAILURE, CC_SE_DELETION_DETECTED, CC_VNIC_ADDED, CC_VNIC_ADDITION_FAILURE, CC_VNIC_DELETED, CC_VNIC_DELETION_FAILURE, CC_IP_ATTACHED, CC_IP_ATTACH_FAILURE, CC_IP_DETACHED, CC_IP_DETACH_FAILURE, CC_SYNC_SERVICES_SUCCESS, CC_SYNC_SERVICES_FAILURE, CC_UPDATE_VIP_FAILURE, CC_DELETE_VIP_FAILURE, CC_CONFIG_FAILURE, CC_DECONFIG_FAILURE, CC_GENERIC_FAILURE, CC_CLUSTER_VIP_CONFIG_SUCCESS, CC_CLUSTER_VIP_CONFIG_FAILURE, CC_CLUSTER_VIP_DECONFIG_SUCCESS, CC_CLUSTER_VIP_DECONFIG_FAILURE, CC_MARATHON_SERVICE_PORT_OUTSIDE_VALID_RANGE, CC_MARATHON_SERVICE_PORT_ALREADY_IN_USE, CC_VIP_DNS_REGISTER_FAILURE, CC_TENANT_INIT_FAILURE, CC_HEALTH_FAILURE, CC_HEALTH_OK, CC_SE_STARTED, CC_SE_START_FAILURE, CC_SE_STOPPED, CC_SE_STOP_FAILURE, CC_VIP_PARK_INTF_SUCCESS, CC_VIP_PARK_INTF_FAILURE, CC_VIP_DNS_DEREGISTER_FAILURE, CC_VIP_DNS_VALIDATION_FAILURE, CC_VIP_DNS_REGISTER_SUCCESS, CC_VIP_DNS_DEREGISTER_SUCCESS, AWS_ROUTE53_ACCESS_FAILURE, AWS_ROUTE53_ACCESS_SUCCESS, VS_HEALTH_CHANGE, SE_HEALTH_CHANGE, POOL_HEALTH_CHANGE, SERVER_HEALTH_CHANGE, VS_HEALTH_DEGRADED, SE_HEALTH_DEGRADED, POOL_HEALTH_DEGRADED, SERVER_HEALTH_DEGRADED, DUPLICATE_SUBNETS, SUMMARIZED_SUBNETS, IP_POOL_ALMOST_EXHAUSTED, IP_POOL_EXHAUSTED, LICENSE_LIMIT_SERVERS, LICENSE_LIMIT_SE_VCPUS, LICENSE_LIMIT_THROUGHPUT, LICENSE_LIMIT_VS, LICENSE_LIMIT_HOSTS, LICENSE_LIMIT_SE_SOCKETS, LICENSE_EXPIRED, BURST_RESOURCE_CONSUMED, BURST_RESOURCE_EXPIRY_ALERT, APIC_BAD_CREDENTIALS, APIC_CREATE_LIFS, APIC_DELETE_LIFS, APIC_CREATE_LIF_CONTEXTS, APIC_DELETE_LIF_CONTEXTS, APIC_CREATE_CDEV, APIC_DELETE_CDEV, APIC_ATTACH_CIF_TO_LIF, APIC_DETACH_CIF_FROM_LIF, APIC_VS_PLACEMENT, APIC_BIND_VNIC_TO_NETWORK, APIC_CREATE_TENANT, APIC_DELETE_TENANT, APIC_CREATE_NETWORK, APIC_DELETE_NETWORK, APIC_NETWORK_VRF_CHANGED, APIC_VS_NETWORK_RESOLVE_ERROR, CONTAINER_CLOUD_ACCESS_SUCCESS, CONTAINER_CLOUD_ACCESS_FAILURE, CONTAINER_CLOUD_IMAGE_UPLOAD_FAILURE, CONTAINER_CLOUD_IMAGE_UPLOAD_SUCCESS, CONTAINER_CLOUD_IMAGE_UPLOAD_IN_PROGRESS, CONTAINER_CLOUD_CREATED_SE, CONTAINER_CLOUD_CREATE_SE_FAIL, CONTAINER_CLOUD_DELETED_SE, CONTAINER_CLOUD_DELETE_SE_FAIL, CONTAINER_CLOUD_STOPPED_SE, CONTAINER_CLOUD_STOP_SE_FAIL, CONTAINER_CLOUD_STARTED_SE, CONTAINER_CLOUD_START_SE_FAIL, CONTAINER_CLOUD_UPDATED_HOSTS, CONTAINER_CLOUD_SERVICE_SUCCESS, CONTAINER_CLOUD_SERVICE_FAILURE, CONTAINER_CLOUD_SERVICE_INCOMPLETE, AVG_UPTIME_CHANGE, DOS_ATTACK, SE_DOS_ATTACK, SERVER_AUTOSCALE_OUT, SERVER_AUTOSCALE_IN, SERVER_AUTOSCALE_OUT_COMPLETE, SERVER_AUTOSCALE_IN_COMPLETE, SERVER_AUTOSCALE_FAILED, SERVER_AUTOSCALE_IN_FAILED, SERVER_AUTOSCALE_OUT_FAILED, SERVER_AUTOSCALE_OUT_TRIGGERED, SERVER_AUTOSCALE_IN_TRIGGERED, SE_GATEWAY_HEARTBEAT_FAILED, SE_GATEWAY_HEARTBEAT_SUCCESS, SE_VNIC_DOWN_EVENT, SE_VNIC_TX_QUEUE_STALL, SE_BGP_PEER_STATE_CHANGE, SE_LICENSED_BANDWIDTH_EXCEEDED, POOL_AUTO_DEPLOYMENT_FAILED, POOL_AUTO_DEPLOYMENT_SUCCESS, SE_VNIC_UP_EVENT, GSLB_SITE_OPER_STATUS, GSLB_DNS_STATUS, GSLB_SITE_EXCEPTION_STATUS, GSLB_GS_STATUS, SCHEDULER_ACTION_SUCCESS, SCHEDULER_ACTION_FAILURE, GCP_ACCESS_SUCCESS, GCP_ACCESS_FAIL, GCP_SE_DETECTED, GCP_API_FAIL, GCP_SUBNET_NOT_FOUND, GCP_SUBNET_ATTACH_FAIL, GCP_ROUTE_ADD_SUCCESS, GCP_ROUTE_DELETE_SUCCESS, GCP_ROUTE_ADD_FAIL, GCP_ROUTE_DELETE_FAIL, VIP_DNS_REGISTER_SUCCESS, VIP_DNS_REGISTER_FAILURE, VIP_DNS_DEREGISTER_SUCCESS, VIP_DNS_DEREGISTER_FAILURE, SYNC_DNS_RECORDS_SUCCESS, SYNC_DNS_RECORDS_FAILURE, FLUSH_DNS_RECORDS_SUCCESS, FLUSH_DNS_RECORDS_FAILURE, CC_HOST_SSH_FAILURE, CC_HOST_SSH_SUCCESS, AZURE_ACCESS_SUCCESS, AZURE_ACCESS_FAILURE, AZURE_ALB_UPDATE_FAILURE, AZURE_NIC_UPDATE_FAILURE, AZURE_ALB_UPDATE_SUCCESS, AZURE_NIC_UPDATE_SUCCESS, AZURE_NIC_DELETE_SUCCESS, AZURE_NIC_DELETE_FAILURE, AZURE_IMAGE_UPLOAD_FAILURE, AZURE_IMAGE_UPLOAD_SUCCESS.
	Events []string `json:"events,omitempty"`

	// pool of EventGenParams.
	Pool string `json:"pool,omitempty"`

	// sslkeyandcertificate of EventGenParams.
	Sslkeyandcertificate string `json:"sslkeyandcertificate,omitempty"`

	// virtualservice of EventGenParams.
	Virtualservice string `json:"virtualservice,omitempty"`
}
