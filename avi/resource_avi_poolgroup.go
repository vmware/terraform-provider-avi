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

func ResourcePoolGroupSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"cloud_config_cksum": {
			Type:     schema.TypeString,
			Optional: true,
			Computed: true,
		},
		"cloud_ref": {
			Type:     schema.TypeString,
			Optional: true,
			Computed: true,
		},
		"created_by": {
			Type:     schema.TypeString,
			Optional: true,
			Computed: true,
		},
		"deployment_policy_ref": {
			Type:     schema.TypeString,
			Optional: true,
			Computed: true,
		},
		"description": {
			Type:     schema.TypeString,
			Optional: true,
			Computed: true,
		},
		"enable_http2": {
			Type:     schema.TypeBool,
			Optional: true,
			Default:  false,
		},
		"fail_action": {
			Type:     schema.TypeSet,
			Optional: true,
			Computed: true,
			Elem:     ResourceFailActionSchema(),
		},
		"implicit_priority_labels": {
			Type:     schema.TypeBool,
			Optional: true,
			Default:  false,
		},
		"markers": {
			Type:     schema.TypeList,
			Optional: true,
			Elem:     ResourceRoleFilterMatchLabelSchema(),
		},
		"members": {
			Type:     schema.TypeList,
			Optional: true,
			Elem:     ResourcePoolGroupMemberSchema(),
		},
		"min_servers": {
			Type:     schema.TypeInt,
			Optional: true,
			Default:  0,
		},
		"name": {
			Type:     schema.TypeString,
			Required: true,
		},
		"priority_labels_ref": {
			Type:     schema.TypeString,
			Optional: true,
			Computed: true,
		},
		"service_metadata": {
			Type:     schema.TypeString,
			Optional: true,
			Computed: true,
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

func resourceAviPoolGroup() *schema.Resource {
	return &schema.Resource{
		Create: resourceAviPoolGroupCreate,
		Read:   ResourceAviPoolGroupRead,
		Update: resourceAviPoolGroupUpdate,
		Delete: resourceAviPoolGroupDelete,
		Schema: ResourcePoolGroupSchema(),
		Importer: &schema.ResourceImporter{
			State: ResourcePoolGroupImporter,
		},
	}
}

func ResourcePoolGroupImporter(d *schema.ResourceData, m interface{}) ([]*schema.ResourceData, error) {
	s := ResourcePoolGroupSchema()
	return ResourceImporter(d, m, "poolgroup", s)
}

func ResourceAviPoolGroupRead(d *schema.ResourceData, meta interface{}) error {
	s := ResourcePoolGroupSchema()
	err := APIRead(d, meta, "poolgroup", s)
	if err != nil {
		log.Printf("[ERROR] in reading object %v\n", err)
	}
	return err
}

func resourceAviPoolGroupCreate(d *schema.ResourceData, meta interface{}) error {
	s := ResourcePoolGroupSchema()
	err := APICreateOrUpdate(d, meta, "poolgroup", s)
	if err == nil {
		err = ResourceAviPoolGroupRead(d, meta)
	}
	return err
}

func resourceAviPoolGroupUpdate(d *schema.ResourceData, meta interface{}) error {
	s := ResourcePoolGroupSchema()
	var err error
	err = APICreateOrUpdate(d, meta, "poolgroup", s)
	if err == nil {
		err = ResourceAviPoolGroupRead(d, meta)
	}
	return err
}

func resourceAviPoolGroupDelete(d *schema.ResourceData, meta interface{}) error {
	objType := "poolgroup"
	client := meta.(*clients.AviClient)
	if APIDeleteSystemDefaultCheck(d) {
		return nil
	}
	uuid := d.Get("uuid").(string)
	if uuid != "" {
		path := "api/" + objType + "/" + uuid
		err := client.AviSession.Delete(path)
		if err != nil && !(strings.Contains(err.Error(), "404") || strings.Contains(err.Error(), "204") || strings.Contains(err.Error(), "403")) {
			log.Println("[INFO] resourceAviPoolGroupDelete not found")
			return err
		}
		d.SetId("")
	}
	return nil
}
