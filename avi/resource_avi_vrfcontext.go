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

func ResourceVrfContextSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"bgp_profile": &schema.Schema{
			Type:     schema.TypeSet,
			Optional: true,
			Elem:     ResourceBgpProfileSchema(),
			Set: func(v interface{}) int {
				return 0
			},
		},
		"cloud_ref": &schema.Schema{
			Type:     schema.TypeString,
			Optional: true,
			Default:  "/api/cloud?name=Default-Cloud",
		},
		"debugvrfcontext": &schema.Schema{
			Type:     schema.TypeSet,
			Optional: true,
			Elem:     ResourceDebugVrfContextSchema(),
			Set: func(v interface{}) int {
				return 0
			},
		},
		"description": &schema.Schema{
			Type:     schema.TypeString,
			Optional: true,
		},
		"gateway_mon": &schema.Schema{
			Type:     schema.TypeList,
			Optional: true,
			Elem:     ResourceGatewayMonitorSchema(),
		},
		"internal_gateway_monitor": &schema.Schema{
			Type:     schema.TypeSet,
			Optional: true,
			Elem:     ResourceInternalGatewayMonitorSchema(),
			Set: func(v interface{}) int {
				return 0
			},
		},
		"name": &schema.Schema{
			Type:     schema.TypeString,
			Required: true,
		},
		"static_routes": &schema.Schema{
			Type:     schema.TypeList,
			Optional: true,
			Elem:     ResourceStaticRouteSchema(),
		},
		"system_default": &schema.Schema{
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
	}
}

func resourceAviVrfContext() *schema.Resource {
	return &schema.Resource{
		Create: resourceAviVrfContextCreate,
		Read:   ResourceAviVrfContextRead,
		Update: resourceAviVrfContextUpdate,
		Delete: resourceAviVrfContextDelete,
		Schema: ResourceVrfContextSchema(),
		Importer: &schema.ResourceImporter{
			State: ResourceVrfContextImporter,
		},
	}
}

func ResourceVrfContextImporter(d *schema.ResourceData, m interface{}) ([]*schema.ResourceData, error) {
	s := ResourceVrfContextSchema()
	return ResourceImporter(d, m, "vrfcontext", s)
}

func ResourceAviVrfContextRead(d *schema.ResourceData, meta interface{}) error {
	s := ResourceVrfContextSchema()
	err := ApiRead(d, meta, "vrfcontext", s)
	if err != nil {
		log.Printf("[ERROR] in reading object %v\n", err)
	}
	return err
}

func resourceAviVrfContextCreate(d *schema.ResourceData, meta interface{}) error {
	s := ResourceVrfContextSchema()
	err := ApiCreateOrUpdate(d, meta, "vrfcontext", s)
	if err == nil {
		err = ResourceAviVrfContextRead(d, meta)
	}
	return err
}

func resourceAviVrfContextUpdate(d *schema.ResourceData, meta interface{}) error {
	s := ResourceVrfContextSchema()
	err := ApiCreateOrUpdate(d, meta, "vrfcontext", s)
	if err == nil {
		err = ResourceAviVrfContextRead(d, meta)
	}
	return err
}

func resourceAviVrfContextDelete(d *schema.ResourceData, meta interface{}) error {
	objType := "vrfcontext"
	client := meta.(*clients.AviClient)
	uuid := d.Get("uuid").(string)
	if uuid != "" {
		path := "api/" + objType + "/" + uuid
		err := client.AviSession.Delete(path)
		if err != nil && !(strings.Contains(err.Error(), "404") || strings.Contains(err.Error(), "204")) {
			log.Println("[INFO] resourceAviVrfContextDelete not found")
			return err
		}
		d.SetId("")
	}
	return nil
}
