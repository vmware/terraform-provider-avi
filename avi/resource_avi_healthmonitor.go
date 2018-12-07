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

func ResourceHealthMonitorSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"description": &schema.Schema{
			Type:     schema.TypeString,
			Optional: true,
		},
		"dns_monitor": &schema.Schema{
			Type:     schema.TypeSet,
			Optional: true,
			Elem:     ResourceHealthMonitorDNSSchema(),
			Set: func(v interface{}) int {
				return 0
			},
		},
		"external_monitor": &schema.Schema{
			Type:     schema.TypeSet,
			Optional: true,
			Elem:     ResourceHealthMonitorExternalSchema(),
			Set: func(v interface{}) int {
				return 0
			},
		},
		"failed_checks": &schema.Schema{
			Type:     schema.TypeInt,
			Optional: true,
			Default:  2,
		},
		"http_monitor": &schema.Schema{
			Type:     schema.TypeSet,
			Optional: true,
			Elem:     ResourceHealthMonitorHttpSchema(),
			Set: func(v interface{}) int {
				return 0
			},
		},
		"https_monitor": &schema.Schema{
			Type:     schema.TypeSet,
			Optional: true,
			Elem:     ResourceHealthMonitorHttpSchema(),
			Set: func(v interface{}) int {
				return 0
			},
		},
		"is_federated": &schema.Schema{
			Type:     schema.TypeBool,
			Optional: true,
			Default:  false,
		},
		"monitor_port": &schema.Schema{
			Type:     schema.TypeInt,
			Optional: true,
		},
		"name": &schema.Schema{
			Type:     schema.TypeString,
			Required: true,
		},
		"receive_timeout": &schema.Schema{
			Type:     schema.TypeInt,
			Optional: true,
			Default:  4,
		},
		"send_interval": &schema.Schema{
			Type:     schema.TypeInt,
			Optional: true,
			Default:  10,
		},
		"sip_monitor": &schema.Schema{
			Type:     schema.TypeSet,
			Optional: true,
			Elem:     ResourceHealthMonitorSIPSchema(),
			Set: func(v interface{}) int {
				return 0
			},
		},
		"successful_checks": &schema.Schema{
			Type:     schema.TypeInt,
			Optional: true,
			Default:  2,
		},
		"tcp_monitor": &schema.Schema{
			Type:     schema.TypeSet,
			Optional: true,
			Elem:     ResourceHealthMonitorTcpSchema(),
			Set: func(v interface{}) int {
				return 0
			},
		},
		"tenant_ref": &schema.Schema{
			Type:     schema.TypeString,
			Optional: true,
		},
		"type": &schema.Schema{
			Type:     schema.TypeString,
			Required: true,
		},
		"udp_monitor": &schema.Schema{
			Type:     schema.TypeSet,
			Optional: true,
			Elem:     ResourceHealthMonitorUdpSchema(),
			Set: func(v interface{}) int {
				return 0
			},
		},
		"uuid": &schema.Schema{
			Type:     schema.TypeString,
			Optional: true,
			Computed: true,
		},
	}
}

func resourceAviHealthMonitor() *schema.Resource {
	return &schema.Resource{
		Create: resourceAviHealthMonitorCreate,
		Read:   ResourceAviHealthMonitorRead,
		Update: resourceAviHealthMonitorUpdate,
		Delete: resourceAviHealthMonitorDelete,
		Schema: ResourceHealthMonitorSchema(),
		Importer: &schema.ResourceImporter{
			State: ResourceHealthMonitorImporter,
		},
	}
}

func ResourceHealthMonitorImporter(d *schema.ResourceData, m interface{}) ([]*schema.ResourceData, error) {
	s := ResourceHealthMonitorSchema()
	return ResourceImporter(d, m, "healthmonitor", s)
}

func ResourceAviHealthMonitorRead(d *schema.ResourceData, meta interface{}) error {
	s := ResourceHealthMonitorSchema()
	err := ApiRead(d, meta, "healthmonitor", s)
	if err != nil {
		log.Printf("[ERROR] in reading object %v\n", err)
	}
	return err
}

func resourceAviHealthMonitorCreate(d *schema.ResourceData, meta interface{}) error {
	s := ResourceHealthMonitorSchema()
	err := ApiCreateOrUpdate(d, meta, "healthmonitor", s)
	if err == nil {
		err = ResourceAviHealthMonitorRead(d, meta)
	}
	return err
}

func resourceAviHealthMonitorUpdate(d *schema.ResourceData, meta interface{}) error {
	s := ResourceHealthMonitorSchema()
	var err error
	err = ApiCreateOrUpdate(d, meta, "healthmonitor", s)
	if err == nil {
		err = ResourceAviHealthMonitorRead(d, meta)
	}
	return err
}

func resourceAviHealthMonitorDelete(d *schema.ResourceData, meta interface{}) error {
	objType := "healthmonitor"
	if ApiDeleteSystemDefaultCheck(d) {
		return nil
	}
	client := meta.(*clients.AviClient)
	uuid := d.Get("uuid").(string)
	if uuid != "" {
		path := "api/" + objType + "/" + uuid
		err := client.AviSession.Delete(path)
		if err != nil && !(strings.Contains(err.Error(), "404") || strings.Contains(err.Error(), "204") || strings.Contains(err.Error(), "403")) {
			log.Println("[INFO] resourceAviHealthMonitorDelete not found")
			return err
		}
		d.SetId("")
	}
	return nil
}
