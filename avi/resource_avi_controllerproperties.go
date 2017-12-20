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
		"allow_ip_forwarding": &schema.Schema{
			Type:     schema.TypeBool,
			Optional: true,
			Default:  false,
		},
		"allow_unauthenticated_apis": &schema.Schema{
			Type:     schema.TypeBool,
			Optional: true,
			Default:  false,
		},
		"allow_unauthenticated_nodes": &schema.Schema{
			Type:     schema.TypeBool,
			Optional: true,
			Default:  false,
		},
		"api_idle_timeout": &schema.Schema{
			Type:     schema.TypeInt,
			Optional: true,
			Default:  15,
		},
		"appviewx_compat_mode": &schema.Schema{
			Type:     schema.TypeBool,
			Optional: true,
			Default:  false,
		},
		"attach_ip_retry_interval": &schema.Schema{
			Type:     schema.TypeInt,
			Optional: true,
			Default:  360,
		},
		"attach_ip_retry_limit": &schema.Schema{
			Type:     schema.TypeInt,
			Optional: true,
			Default:  4,
		},
		"bm_use_ansible": &schema.Schema{
			Type:     schema.TypeBool,
			Optional: true,
			Default:  true,
		},
		"cluster_ip_gratuitous_arp_period": &schema.Schema{
			Type:     schema.TypeInt,
			Optional: true,
			Default:  60,
		},
		"crashed_se_reboot": &schema.Schema{
			Type:     schema.TypeInt,
			Optional: true,
			Default:  900,
		},
		"dead_se_detection_timer": &schema.Schema{
			Type:     schema.TypeInt,
			Optional: true,
			Default:  360,
		},
		"dns_refresh_period": &schema.Schema{
			Type:     schema.TypeInt,
			Optional: true,
			Default:  60,
		},
		"dummy": &schema.Schema{
			Type:     schema.TypeInt,
			Optional: true,
		},
		"fatal_error_lease_time": &schema.Schema{
			Type:     schema.TypeInt,
			Optional: true,
			Default:  120,
		},
		"max_dead_se_in_grp": &schema.Schema{
			Type:     schema.TypeInt,
			Optional: true,
			Default:  1,
		},
		"max_pcap_per_tenant": &schema.Schema{
			Type:     schema.TypeInt,
			Optional: true,
			Default:  4,
		},
		"max_seq_attach_ip_failures": &schema.Schema{
			Type:     schema.TypeInt,
			Optional: true,
			Default:  3,
		},
		"max_seq_vnic_failures": &schema.Schema{
			Type:     schema.TypeInt,
			Optional: true,
			Default:  3,
		},
		"persistence_key_rotate_period": &schema.Schema{
			Type:     schema.TypeInt,
			Optional: true,
			Default:  60,
		},
		"portal_token": &schema.Schema{
			Type:     schema.TypeString,
			Optional: true,
		},
		"query_host_fail": &schema.Schema{
			Type:     schema.TypeInt,
			Optional: true,
			Default:  180,
		},
		"safenet_hsm_version": &schema.Schema{
			Type:     schema.TypeString,
			Optional: true,
		},
		"se_create_timeout": &schema.Schema{
			Type:     schema.TypeInt,
			Optional: true,
			Default:  900,
		},
		"se_failover_attempt_interval": &schema.Schema{
			Type:     schema.TypeInt,
			Optional: true,
			Default:  300,
		},
		"se_offline_del": &schema.Schema{
			Type:     schema.TypeInt,
			Optional: true,
			Default:  172000,
		},
		"se_vnic_cooldown": &schema.Schema{
			Type:     schema.TypeInt,
			Optional: true,
			Default:  120,
		},
		"secure_channel_cleanup_timeout": &schema.Schema{
			Type:     schema.TypeInt,
			Optional: true,
			Default:  60,
		},
		"secure_channel_controller_token_timeout": &schema.Schema{
			Type:     schema.TypeInt,
			Optional: true,
			Default:  60,
		},
		"secure_channel_se_token_timeout": &schema.Schema{
			Type:     schema.TypeInt,
			Optional: true,
			Default:  60,
		},
		"seupgrade_fabric_pool_size": &schema.Schema{
			Type:     schema.TypeInt,
			Optional: true,
			Default:  20,
		},
		"seupgrade_segroup_min_dead_timeout": &schema.Schema{
			Type:     schema.TypeInt,
			Optional: true,
			Default:  360,
		},
		"ssl_certificate_expiry_warning_days": &schema.Schema{
			Type:     schema.TypeList,
			Optional: true,
			Elem:     &schema.Schema{Type: schema.TypeInt},
		},
		"unresponsive_se_reboot": &schema.Schema{
			Type:     schema.TypeInt,
			Optional: true,
			Default:  300,
		},
		"upgrade_dns_ttl": &schema.Schema{
			Type:     schema.TypeInt,
			Optional: true,
			Default:  5,
		},
		"upgrade_lease_time": &schema.Schema{
			Type:     schema.TypeInt,
			Optional: true,
			Default:  360,
		},
		"uuid": &schema.Schema{
			Type:     schema.TypeString,
			Optional: true,
			Computed: true,
		},
		"vnic_op_fail_time": &schema.Schema{
			Type:     schema.TypeInt,
			Optional: true,
			Default:  180,
		},
		"vs_apic_scaleout_timeout": &schema.Schema{
			Type:     schema.TypeInt,
			Optional: true,
			Default:  360,
		},
		"vs_awaiting_se_timeout": &schema.Schema{
			Type:     schema.TypeInt,
			Optional: true,
			Default:  60,
		},
		"vs_key_rotate_period": &schema.Schema{
			Type:     schema.TypeInt,
			Optional: true,
			Default:  60,
		},
		"vs_se_attach_ip_fail": &schema.Schema{
			Type:     schema.TypeInt,
			Optional: true,
			Default:  3600,
		},
		"vs_se_bootup_fail": &schema.Schema{
			Type:     schema.TypeInt,
			Optional: true,
			Default:  480,
		},
		"vs_se_create_fail": &schema.Schema{
			Type:     schema.TypeInt,
			Optional: true,
			Default:  1500,
		},
		"vs_se_ping_fail": &schema.Schema{
			Type:     schema.TypeInt,
			Optional: true,
			Default:  60,
		},
		"vs_se_vnic_fail": &schema.Schema{
			Type:     schema.TypeInt,
			Optional: true,
			Default:  300,
		},
		"vs_se_vnic_ip_fail": &schema.Schema{
			Type:     schema.TypeInt,
			Optional: true,
			Default:  120,
		},
		"warmstart_se_reconnect_wait_time": &schema.Schema{
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
	}
}

func ResourceAviControllerPropertiesRead(d *schema.ResourceData, meta interface{}) error {
	s := ResourceControllerPropertiesSchema()
	log.Printf("[INFO] ResourceAviControllerPropertiesRead Avi Client %v\n", d)
	client := meta.(*clients.AviClient)
	var obj interface{}
	if uuid, ok := d.GetOk("uuid"); ok {
		path := "api/controllerproperties/" + uuid.(string)
		err := client.AviSession.Get(path, &obj)
		if err != nil {
			d.SetId("")
			return nil
		}
	} else {
		d.SetId("")
		return nil
	}
	// no need to set the ID
	log.Printf("ResourceAviControllerPropertiesRead CURRENT obj %v\n", d)

	log.Printf("ResourceAviControllerPropertiesRead Read API obj %v\n", obj)
	if tObj, err := ApiDataToSchema(obj, d, s); err == nil {
		log.Printf("[INFO] ResourceAviControllerPropertiesRead Converted obj %v\n", tObj)
		//err = d.Set("obj", tObj)
		if err != nil {
			log.Printf("[ERROR] in setting read object %v\n", err)
		}
	}
	log.Printf("[INFO] ResourceAviControllerPropertiesRead Updated %v\n", d)
	return nil
}

func resourceAviControllerPropertiesCreate(d *schema.ResourceData, meta interface{}) error {
	s := ResourceControllerPropertiesSchema()
	err := ApiCreateOrUpdate(d, meta, "controllerproperties", s)
	log.Printf("[DEBUG] created object %v: %v", "controllerproperties", d)
	if err == nil {
		err = ResourceAviControllerPropertiesRead(d, meta)
	}
	log.Printf("[DEBUG] created object %v: %v", "controllerproperties", d)
	return err
}

func resourceAviControllerPropertiesUpdate(d *schema.ResourceData, meta interface{}) error {
	s := ResourceControllerPropertiesSchema()
	err := ApiCreateOrUpdate(d, meta, "controllerproperties", s)
	if err == nil {
		err = ResourceAviControllerPropertiesRead(d, meta)
	}
	log.Printf("[DEBUG] updated object %v: %v", "controllerproperties", d)
	return err
}

func resourceAviControllerPropertiesDelete(d *schema.ResourceData, meta interface{}) error {
	objType := "controllerproperties"
	log.Println("[INFO] ResourceAviControllerPropertiesRead Avi Client")
	client := meta.(*clients.AviClient)
	uuid := d.Get("uuid").(string)
	if uuid != "" {
		path := "api/" + objType + "/" + uuid
		err := client.AviSession.Delete(path)
		if err != nil && !(strings.Contains(err.Error(), "404") || strings.Contains(err.Error(), "204")) {
			log.Println("[INFO] resourceAviControllerPropertiesDelete not found")
			return err
		}
		d.SetId("")
	}
	return nil
}
