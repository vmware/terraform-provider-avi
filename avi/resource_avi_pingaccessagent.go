/*
* Copyright (c) 2017. Avi Networks.
* Author: Gaurav Rastogi (grastogi@avinetworks.com)
*
 */
package avi

import (
	"log"
	"strings"

	"github.com/avinetworks/sdk/go/clients"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func ResourcePingAccessAgentSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"description": {
			Type:     schema.TypeString,
			Optional: true,
			Computed: true,
		},
		"name": {
			Type:     schema.TypeString,
			Required: true,
		},
		"pingaccess_pool_ref": {
			Type:     schema.TypeString,
			Required: true,
		},
		"primary_server": {
			Type:     schema.TypeSet,
			Required: true,
			Elem:     ResourcePoolServerSchema(),
		},
		"properties_file_data": {
			Type:     schema.TypeString,
			Required: true,
		},
		"tenant_ref": {
			Type:     schema.TypeString,
			Optional: true,
			Computed: true,
		},
		"uuid": {
			Type:     schema.TypeString,
			Optional: true,
			Computed: true,
		},
	}
}

func resourceAviPingAccessAgent() *schema.Resource {
	return &schema.Resource{
		Create: resourceAviPingAccessAgentCreate,
		Read:   ResourceAviPingAccessAgentRead,
		Update: resourceAviPingAccessAgentUpdate,
		Delete: resourceAviPingAccessAgentDelete,
		Schema: ResourcePingAccessAgentSchema(),
		Importer: &schema.ResourceImporter{
			State: ResourcePingAccessAgentImporter,
		},
	}
}

func ResourcePingAccessAgentImporter(d *schema.ResourceData, m interface{}) ([]*schema.ResourceData, error) {
	s := ResourcePingAccessAgentSchema()
	return ResourceImporter(d, m, "pingaccessagent", s)
}

func ResourceAviPingAccessAgentRead(d *schema.ResourceData, meta interface{}) error {
	s := ResourcePingAccessAgentSchema()
	err := APIRead(d, meta, "pingaccessagent", s)
	if err != nil {
		log.Printf("[ERROR] in reading object %v\n", err)
	}
	return err
}

func resourceAviPingAccessAgentCreate(d *schema.ResourceData, meta interface{}) error {
	s := ResourcePingAccessAgentSchema()
	err := APICreateOrUpdate(d, meta, "pingaccessagent", s)
	if err == nil {
		err = ResourceAviPingAccessAgentRead(d, meta)
	}
	return err
}

func resourceAviPingAccessAgentUpdate(d *schema.ResourceData, meta interface{}) error {
	s := ResourcePingAccessAgentSchema()
	var err error
	err = APICreateOrUpdate(d, meta, "pingaccessagent", s)
	if err == nil {
		err = ResourceAviPingAccessAgentRead(d, meta)
	}
	return err
}

func resourceAviPingAccessAgentDelete(d *schema.ResourceData, meta interface{}) error {
	objType := "pingaccessagent"
	client := meta.(*clients.AviClient)
	if APIDeleteSystemDefaultCheck(d) {
		return nil
	}
	uuid := d.Get("uuid").(string)
	if uuid != "" {
		path := "api/" + objType + "/" + uuid
		err := client.AviSession.Delete(path)
		if err != nil && !(strings.Contains(err.Error(), "404") || strings.Contains(err.Error(), "204") || strings.Contains(err.Error(), "403")) {
			log.Println("[INFO] resourceAviPingAccessAgentDelete not found")
			return err
		}
		d.SetId("")
	}
	return nil
}
