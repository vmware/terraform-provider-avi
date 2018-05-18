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

func ResourceNetworkSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"cloud_ref": &schema.Schema{
			Type:     schema.TypeString,
			Optional: true,
		},
		"configured_subnets": &schema.Schema{
			Type:     schema.TypeList,
			Optional: true,
			Elem:     ResourceSubnetSchema(),
		},
		"dhcp_enabled": &schema.Schema{
			Type:     schema.TypeBool,
			Optional: true,
			Default:  true,
		},
		"exclude_discovered_subnets": &schema.Schema{
			Type:     schema.TypeBool,
			Optional: true,
			Default:  false,
		},
		"name": &schema.Schema{
			Type:     schema.TypeString,
			Required: true,
		},
		"synced_from_se": &schema.Schema{
			Type:     schema.TypeBool,
			Optional: true,
			Default:  false,
		},
		"tenant_ref": &schema.Schema{
			Type:     schema.TypeString,
			Optional: true,
		},
		"uuid": &schema.Schema{
			Type:     schema.TypeString,
			Optional: true,
			Computed: true,
		},
		"vcenter_dvs": &schema.Schema{
			Type:     schema.TypeBool,
			Optional: true,
			Default:  true,
		},
		"vrf_context_ref": &schema.Schema{
			Type:     schema.TypeString,
			Optional: true,
		},
	}
}

func resourceAviNetwork() *schema.Resource {
	return &schema.Resource{
		Create: resourceAviNetworkCreate,
		Read:   ResourceAviNetworkRead,
		Update: resourceAviNetworkUpdate,
		Delete: resourceAviNetworkDelete,
		Schema: ResourceNetworkSchema(),
		Importer: &schema.ResourceImporter{
			State: ResourceNetworkImporter,
		},
	}
}

func ResourceNetworkImporter(d *schema.ResourceData, m interface{}) ([]*schema.ResourceData, error) {
	s := ResourceNetworkSchema()
	return ResourceImporter(d, m, "network", s)
}

func ResourceAviNetworkRead(d *schema.ResourceData, meta interface{}) error {
	s := ResourceNetworkSchema()
	err := ApiRead(d, meta, "network", s)
	if err != nil {
		log.Printf("[ERROR] in reading object %v\n", err)
	}
	return err
}

func resourceAviNetworkCreate(d *schema.ResourceData, meta interface{}) error {
	s := ResourceNetworkSchema()
	err := ApiCreateOrUpdate(d, meta, "network", s)
	if err == nil {
		err = ResourceAviNetworkRead(d, meta)
	}
	return err
}

func resourceAviNetworkUpdate(d *schema.ResourceData, meta interface{}) error {
	s := ResourceNetworkSchema()
	var err error
	err = ApiCreateOrUpdate(d, meta, "network", s)
	if err == nil {
		err = ResourceAviNetworkRead(d, meta)
	}
	return err
}

func resourceAviNetworkDelete(d *schema.ResourceData, meta interface{}) error {
	objType := "network"
	if ApiDeleteSystemDefaultCheck(d) {
		return nil
	}
	client := meta.(*clients.AviClient)
	uuid := d.Get("uuid").(string)
	if uuid != "" {
		path := "api/" + objType + "/" + uuid
		err := client.AviSession.Delete(path)
		if err != nil && !(strings.Contains(err.Error(), "404") || strings.Contains(err.Error(), "204") || strings.Contains(err.Error(), "403")) {
			log.Println("[INFO] resourceAviNetworkDelete not found")
			return err
		}
		d.SetId("")
	}
	return nil
}
