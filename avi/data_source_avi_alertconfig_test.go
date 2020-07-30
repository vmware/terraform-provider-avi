package avi

import (
	"github.com/hashicorp/terraform/helper/resource"
	"testing"
)

func TestAVIDataSourceAlertConfigBasic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccAVIDSAlertConfigConfig,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(
						"avi_alertconfig.testAlertConfig", "name", "test-System-CC-Alert-abc"),
					resource.TestCheckResourceAttr(
						"avi_alertconfig.testAlertConfig", "enabled", "true"),
					resource.TestCheckResourceAttr(
						"avi_alertconfig.testAlertConfig", "expiry_time", "86400"),
					resource.TestCheckResourceAttr(
						"avi_alertconfig.testAlertConfig", "rolling_window", "300"),
					resource.TestCheckResourceAttr(
						"avi_alertconfig.testAlertConfig", "threshold", "1"),
					resource.TestCheckResourceAttr(
						"avi_alertconfig.testAlertConfig", "throttle", "0"),
				),
			},
		},
	})

}

const testAccAVIDSAlertConfigConfig = `
data "avi_tenant" "default_tenant"{
    name= "admin"
}
data "avi_actiongroupconfig" "system_actiongroupconfig" {
    name= "System-Alert-Level-Medium"
}
resource "avi_alertconfig" "testAlertConfig" {
	name = "test-System-CC-Alert-abc"
	enabled = true
	tenant_ref = data.avi_tenant.default_tenant.id
	category = "REALTIME"
	expiry_time = "86400"
	summary = "System-CC-Alert System Alert Triggered"
	rolling_window = "300"
	source = "EVENT_LOGS"
	alert_rule {
		operator = "OPERATOR_OR"
		sys_event_rule {
	event_id = "DISCOVERY_DATACENTER_DEL"
	not_cond = false
}
sys_event_rule {
	event_id = "VCENTER_ADDRESS_ERROR"
	not_cond = false
}
sys_event_rule {
	event_id = "SE_GROUP_MGMT_NW_DEL"
	not_cond = false
}
sys_event_rule {
	event_id = "MGMT_NW_DEL"
	not_cond = false
}
sys_event_rule {
	event_id = "VCENTER_BAD_CREDENTIALS"
	not_cond = false
}
sys_event_rule {
	event_id = "ESX_HOST_UNREACHABLE"
	not_cond = false
}
sys_event_rule {
	event_id = "VINFRA_DISC_FAILURE"
	not_cond = false
}
sys_event_rule {
	event_id = "VCENTER_ACCESS_SLOW"
	not_cond = false
}
sys_event_rule {
	event_id = "OPENSTACK_ACCESS_FAILURE"
	not_cond = false
}
sys_event_rule {
	event_id = "OPENSTACK_IMAGE_UPLOAD_FAILURE"
	not_cond = false
}
sys_event_rule {
	event_id = "AWS_ACCESS_FAILURE"
	not_cond = false
}
sys_event_rule {
	event_id = "AWS_IMAGE_UPLOAD_FAILURE"
	not_cond = false
}
sys_event_rule {
	event_id = "AWS_SNS_ACCESS_FAILURE"
	not_cond = false
}
sys_event_rule {
	event_id = "AWS_SQS_ACCESS_FAILURE"
	not_cond = false
}
sys_event_rule {
	event_id = "AWS_ASG_PUT_NOTIFICATION_CONFIGURATION_FAILURE"
	not_cond = false
}
sys_event_rule {
	event_id = "AWS_ASG_DELETE_NOTIFICATION_CONFIGURATION_FAILURE"
	not_cond = false
}
sys_event_rule {
	event_id = "AWS_ASG_NOTIFICATION_PROCESSING_FAILURE"
	not_cond = false
}
sys_event_rule {
	event_id = "AWS_ASG_ACCESS_FAILURE"
	not_cond = false
}
sys_event_rule {
	event_id = "AWS_ASG_NOTIFICATION_INSTANCE_LAUNCH_ERROR"
	not_cond = false
}
sys_event_rule {
	event_id = "AWS_ASG_NOTIFICATION_INSTANCE_TERMINATE_ERROR"
	not_cond = false
}
sys_event_rule {
	event_id = "CLOUDSTACK_ACCESS_FAILURE"
	not_cond = false
}
sys_event_rule {
	event_id = "CLOUDSTACK_IMAGE_UPLOAD_FAILURE"
	not_cond = false
}
sys_event_rule {
	event_id = "DOCKER_UCP_ACCESS_FAILURE"
	not_cond = false
}
sys_event_rule {
	event_id = "DOCKER_UCP_IMAGE_UPLOAD_FAILURE"
	not_cond = false
}
sys_event_rule {
	event_id = "VCA_ACCESS_FAILURE"
	not_cond = false
}
sys_event_rule {
	event_id = "VCA_IMAGE_UPLOAD_FAILURE"
	not_cond = false
}
sys_event_rule {
	event_id = "LS_ACCESS_FAILURE"
	not_cond = false
}
sys_event_rule {
	event_id = "LS_IMAGE_UPLOAD_FAILURE"
	not_cond = false
}
sys_event_rule {
	event_id = "MESOS_ACCESS_FAILURE"
	not_cond = false
}
sys_event_rule {
	event_id = "MESOS_IMAGE_UPLOAD_FAILURE"
	not_cond = false
}
sys_event_rule {
	event_id = "CC_SE_CREATION_FAILURE"
	not_cond = false
}
sys_event_rule {
	event_id = "CC_VNIC_ADDITION_FAILURE"
	not_cond = false
}
sys_event_rule {
	event_id = "CC_IP_ATTACH_FAILURE"
	not_cond = false
}
sys_event_rule {
	event_id = "CC_SYNC_SERVICES_FAILURE"
	not_cond = false
}
sys_event_rule {
	event_id = "CC_UPDATE_VIP_FAILURE"
	not_cond = false
}
sys_event_rule {
	event_id = "CC_CONFIG_FAILURE"
	not_cond = false
}
sys_event_rule {
	event_id = "CC_DECONFIG_FAILURE"
	not_cond = false
}
sys_event_rule {
	event_id = "CC_CLUSTER_VIP_CONFIG_FAILURE"
	not_cond = false
}
sys_event_rule {
	event_id = "CC_MARATHON_SERVICE_PORT_OUTSIDE_VALID_RANGE"
	not_cond = false
}
sys_event_rule {
	event_id = "CC_MARATHON_SERVICE_PORT_ALREADY_IN_USE"
	not_cond = false
}
sys_event_rule {
	event_id = "CC_VIP_DNS_REGISTER_FAILURE"
	not_cond = false
}
sys_event_rule {
	event_id = "CC_TENANT_INIT_FAILURE"
	not_cond = false
}
sys_event_rule {
	event_id = "CC_HEALTH_FAILURE"
	not_cond = false
}
sys_event_rule {
	event_id = "CC_SE_START_FAILURE"
	not_cond = false
}
sys_event_rule {
	event_id = "CC_SE_STOP_FAILURE"
	not_cond = false
}
sys_event_rule {
	event_id = "CC_VIP_PARK_INTF_FAILURE"
	not_cond = false
}
sys_event_rule {
	event_id = "AWS_ROUTE53_ACCESS_FAILURE"
	not_cond = false
}
sys_event_rule {
	event_id = "APIC_BAD_CREDENTIALS"
	not_cond = false
}
sys_event_rule {
	event_id = "CONTAINER_CLOUD_ACCESS_FAILURE"
	not_cond = false
}
sys_event_rule {
	event_id = "CONTAINER_CLOUD_IMAGE_UPLOAD_FAILURE"
	not_cond = false
}
sys_event_rule {
	event_id = "GCP_API_FAIL"
	not_cond = false
}
sys_event_rule {
	event_id = "GCP_SUBNET_NOT_FOUND"
	not_cond = false
}
sys_event_rule {
	event_id = "GCP_SUBNET_ATTACH_FAIL"
	not_cond = false
}
sys_event_rule {
	event_id = "GCP_ROUTE_ADD_FAIL"
	not_cond = false
}
sys_event_rule {
	event_id = "GCP_ROUTE_DELETE_FAIL"
	not_cond = false
}
sys_event_rule {
	event_id = "CC_HOST_SSH_FAILURE"
	not_cond = false
}
sys_event_rule {
	event_id = "AZURE_ALB_UPDATE_FAILURE"
	not_cond = false
}
sys_event_rule {
	event_id = "AZURE_NIC_UPDATE_FAILURE"
	not_cond = false
}
sys_event_rule {
	event_id = "AZURE_NIC_DELETE_FAILURE"
	not_cond = false
}
sys_event_rule {
	event_id = "AZURE_IMAGE_UPLOAD_FAILURE"
	not_cond = false
}
	}
	threshold = "1"
	throttle = "0"
	action_group_ref = data.avi_actiongroupconfig.system_actiongroupconfig.id
}

data "avi_alertconfig" "testAlertConfig" {
    name= "${avi_alertconfig.testAlertConfig.name}"
}
`
