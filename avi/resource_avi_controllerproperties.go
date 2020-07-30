/*
 * Copyright (c) 2017. Avi Networks.
 * Author: Gaurav Rastogi (grastogi@avinetworks.com)
 *
 */
package avi

import (
	"github.com/avinetworks/sdk/go/clients"
	"github.com/hashicorp/terraform/helper/schema"
	"log"
	"strings"
)

func ResourceControllerPropertiesSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"allow_admin_network_updates": {
			Type:     schema.TypeBool,
			Optional: true,
			Default:  false,
		},
		"allow_ip_forwarding": {
			Type:     schema.TypeBool,
			Optional: true,
			Default:  false,
		},
		"allow_unauthenticated_apis": {
			Type:     schema.TypeBool,
			Optional: true,
			Default:  false,
		},
		"allow_unauthenticated_nodes": {
			Type:     schema.TypeBool,
			Optional: true,
			Default:  false,
		},
		"api_idle_timeout": {
			Type:     schema.TypeInt,
			Optional: true,
			Default:  15,
		},
		"api_perf_logging_threshold": {
			Type:     schema.TypeInt,
			Optional: true,
			Default:  10000,
		},
		"appviewx_compat_mode": {
			Type:     schema.TypeBool,
			Optional: true,
			Default:  false,
		},
		"attach_ip_retry_interval": {
			Type:     schema.TypeInt,
			Optional: true,
			Default:  360,
		},
		"attach_ip_retry_limit": {
			Type:     schema.TypeInt,
			Optional: true,
			Default:  4,
		},
		"bm_use_ansible": {
			Type:     schema.TypeBool,
			Optional: true,
			Default:  true,
		},
		"cleanup_expired_authtoken_timeout_period": {
			Type:     schema.TypeInt,
			Optional: true,
			Default:  60,
		},
		"cleanup_sessions_timeout_period": {
			Type:     schema.TypeInt,
			Optional: true,
			Default:  60,
		},
		"cloud_reconcile": {
			Type:     schema.TypeBool,
			Optional: true,
			Default:  true,
		},
		"cluster_ip_gratuitous_arp_period": {
			Type:     schema.TypeInt,
			Optional: true,
			Default:  60,
		},
		"consistency_check_timeout_period": {
			Type:     schema.TypeInt,
			Optional: true,
			Default:  60,
		},
		"crashed_se_reboot": {
			Type:     schema.TypeInt,
			Optional: true,
			Default:  900,
		},
		"dead_se_detection_timer": {
			Type:     schema.TypeInt,
			Optional: true,
			Default:  360,
		},
		"default_minimum_api_timeout": {
			Type:     schema.TypeInt,
			Optional: true,
			Default:  60,
		},
		"dns_refresh_period": {
			Type:     schema.TypeInt,
			Optional: true,
			Default:  60,
		},
		"dummy": {
			Type:     schema.TypeInt,
			Optional: true,
			Computed: true,
		},
		"edit_system_limits": {
			Type:     schema.TypeBool,
			Optional: true,
			Default:  false,
		},
		"enable_api_sharding": {
			Type:     schema.TypeBool,
			Optional: true,
			Default:  true,
		},
		"enable_memory_balancer": {
			Type:     schema.TypeBool,
			Optional: true,
			Default:  true,
		},
		"fatal_error_lease_time": {
			Type:     schema.TypeInt,
			Optional: true,
			Default:  120,
		},
		"federated_datastore_cleanup_duration": {
			Type:     schema.TypeInt,
			Optional: true,
			Default:  120,
		},
		"file_object_cleanup_period": {
			Type:     schema.TypeInt,
			Optional: true,
			Default:  1440,
		},
		"max_dead_se_in_grp": {
			Type:     schema.TypeInt,
			Optional: true,
			Default:  1,
		},
		"max_pcap_per_tenant": {
			Type:     schema.TypeInt,
			Optional: true,
			Default:  4,
		},
		"max_se_spawn_interval_delay": {
			Type:     schema.TypeInt,
			Optional: true,
			Default:  1800,
		},
		"max_seq_attach_ip_failures": {
			Type:     schema.TypeInt,
			Optional: true,
			Default:  3,
		},
		"max_seq_vnic_failures": {
			Type:     schema.TypeInt,
			Optional: true,
			Default:  3,
		},
		"permission_scoped_shared_admin_networks": {
			Type:     schema.TypeBool,
			Optional: true,
			Default:  false,
		},
		"persistence_key_rotate_period": {
			Type:     schema.TypeInt,
			Optional: true,
			Default:  0,
		},
		"portal_request_burst_limit": {
			Type:     schema.TypeInt,
			Optional: true,
			Default:  0,
		},
		"portal_request_rate_limit": {
			Type:     schema.TypeInt,
			Optional: true,
			Default:  0,
		},
		"portal_token": {
			Type:             schema.TypeString,
			Optional:         true,
			Computed:         true,
			Sensitive:        true,
			DiffSuppressFunc: suppressSensitiveFieldDiffs,
		},
		"process_locked_useraccounts_timeout_period": {
			Type:     schema.TypeInt,
			Optional: true,
			Default:  1,
		},
		"process_pki_profile_timeout_period": {
			Type:     schema.TypeInt,
			Optional: true,
			Default:  1440,
		},
		"query_host_fail": {
			Type:     schema.TypeInt,
			Optional: true,
			Default:  180,
		},
		"safenet_hsm_version": {
			Type:     schema.TypeString,
			Optional: true,
			Computed: true,
		},
		"se_create_timeout": {
			Type:     schema.TypeInt,
			Optional: true,
			Default:  900,
		},
		"se_failover_attempt_interval": {
			Type:     schema.TypeInt,
			Optional: true,
			Default:  300,
		},
		"se_from_marketplace": {
			Type:     schema.TypeString,
			Optional: true,
			Default:  "IMAGE",
		},
		"se_offline_del": {
			Type:     schema.TypeInt,
			Optional: true,
			Default:  172000,
		},
		"se_spawn_retry_interval": {
			Type:     schema.TypeInt,
			Optional: true,
			Default:  300,
		},
		"se_vnic_cooldown": {
			Type:     schema.TypeInt,
			Optional: true,
			Default:  120,
		},
		"secure_channel_cleanup_timeout": {
			Type:     schema.TypeInt,
			Optional: true,
			Default:  60,
		},
		"secure_channel_controller_token_timeout": {
			Type:     schema.TypeInt,
			Optional: true,
			Default:  60,
		},
		"secure_channel_se_token_timeout": {
			Type:     schema.TypeInt,
			Optional: true,
			Default:  60,
		},
		"seupgrade_copy_pool_size": {
			Type:     schema.TypeInt,
			Optional: true,
			Default:  5,
		},
		"seupgrade_fabric_pool_size": {
			Type:     schema.TypeInt,
			Optional: true,
			Default:  20,
		},
		"seupgrade_segroup_min_dead_timeout": {
			Type:     schema.TypeInt,
			Optional: true,
			Default:  360,
		},
		"shared_ssl_certificates": {
			Type:     schema.TypeBool,
			Optional: true,
			Default:  false,
		},
		"ssl_certificate_expiry_warning_days": {
			Type:     schema.TypeList,
			Optional: true,
			Elem:     &schema.Schema{Type: schema.TypeInt},
		},
		"unresponsive_se_reboot": {
			Type:     schema.TypeInt,
			Optional: true,
			Default:  300,
		},
		"upgrade_dns_ttl": {
			Type:     schema.TypeInt,
			Optional: true,
			Default:  5,
		},
		"upgrade_fat_se_lease_time": {
			Type:     schema.TypeInt,
			Optional: true,
			Default:  1200,
		},
		"upgrade_lease_time": {
			Type:     schema.TypeInt,
			Optional: true,
			Default:  600,
		},
		"upgrade_se_per_vs_scale_ops_txn_time": {
			Type:     schema.TypeInt,
			Optional: true,
			Default:  3,
		},
		"uuid": {
			Type:     schema.TypeString,
			Optional: true,
			Computed: true,
		},
		"vnic_op_fail_time": {
			Type:     schema.TypeInt,
			Optional: true,
			Default:  180,
		},
		"vs_apic_scaleout_timeout": {
			Type:     schema.TypeInt,
			Optional: true,
			Default:  360,
		},
		"vs_awaiting_se_timeout": {
			Type:     schema.TypeInt,
			Optional: true,
			Default:  60,
		},
		"vs_key_rotate_period": {
			Type:     schema.TypeInt,
			Optional: true,
			Default:  360,
		},
		"vs_scaleout_ready_check_interval": {
			Type:     schema.TypeInt,
			Optional: true,
			Default:  60,
		},
		"vs_se_attach_ip_fail": {
			Type:     schema.TypeInt,
			Optional: true,
			Default:  600,
		},
		"vs_se_bootup_fail": {
			Type:     schema.TypeInt,
			Optional: true,
			Default:  480,
		},
		"vs_se_create_fail": {
			Type:     schema.TypeInt,
			Optional: true,
			Default:  1500,
		},
		"vs_se_ping_fail": {
			Type:     schema.TypeInt,
			Optional: true,
			Default:  60,
		},
		"vs_se_vnic_fail": {
			Type:     schema.TypeInt,
			Optional: true,
			Default:  300,
		},
		"vs_se_vnic_ip_fail": {
			Type:     schema.TypeInt,
			Optional: true,
			Default:  120,
		},
		"warmstart_se_reconnect_wait_time": {
			Type:     schema.TypeInt,
			Optional: true,
			Default:  480,
		},
		"warmstart_vs_resync_wait_time": {
			Type:     schema.TypeInt,
			Optional: true,
			Default:  300,
		},
	}
}

func resourceAviControllerProperties() *schema.Resource {
	return &schema.Resource{
		Create: resourceAviControllerPropertiesCreate,
		Read:   ResourceAviControllerPropertiesRead,
		Update: resourceAviControllerPropertiesUpdate,
		Delete: resourceAviControllerPropertiesDelete,
		Schema: ResourceControllerPropertiesSchema(),
		Importer: &schema.ResourceImporter{
			State: ResourceControllerPropertiesImporter,
		},
	}
}

func ResourceControllerPropertiesImporter(d *schema.ResourceData, m interface{}) ([]*schema.ResourceData, error) {
	s := ResourceControllerPropertiesSchema()
	return ResourceImporter(d, m, "controllerproperties", s)
}

func ResourceAviControllerPropertiesRead(d *schema.ResourceData, meta interface{}) error {
	s := ResourceControllerPropertiesSchema()
	err := ApiRead(d, meta, "controllerproperties", s)
	if err != nil {
		log.Printf("[ERROR] in reading object %v\n", err)
	}
	return err
}

func resourceAviControllerPropertiesCreate(d *schema.ResourceData, meta interface{}) error {
	s := ResourceControllerPropertiesSchema()
	err := ApiCreateOrUpdate(d, meta, "controllerproperties", s)
	if err == nil {
		err = ResourceAviControllerPropertiesRead(d, meta)
	}
	return err
}

func resourceAviControllerPropertiesUpdate(d *schema.ResourceData, meta interface{}) error {
	s := ResourceControllerPropertiesSchema()
	var err error
	err = ApiCreateOrUpdate(d, meta, "controllerproperties", s)
	if err == nil {
		err = ResourceAviControllerPropertiesRead(d, meta)
	}
	return err
}

func resourceAviControllerPropertiesDelete(d *schema.ResourceData, meta interface{}) error {
	objType := "controllerproperties"
	client := meta.(*clients.AviClient)
	if ApiDeleteSystemDefaultCheck(d) {
		return nil
	}
	uuid := d.Get("uuid").(string)
	if uuid != "" {
		path := "api/" + objType + "/" + uuid
		err := client.AviSession.Delete(path)
		if err != nil && !(strings.Contains(err.Error(), "404") || strings.Contains(err.Error(), "204") || strings.Contains(err.Error(), "403")) {
			log.Println("[INFO] resourceAviControllerPropertiesDelete not found")
			return err
		}
		d.SetId("")
	}
	return nil
}
