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

func ResourceNetworkSecurityPolicySchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"cloud_config_cksum": &schema.Schema{
			Type:     schema.TypeString,
			Optional: true,
		},
		"created_by": &schema.Schema{
			Type:     schema.TypeString,
			Optional: true,
		},
		"description": &schema.Schema{
			Type:     schema.TypeString,
			Optional: true,
		},
		"name": &schema.Schema{
			Type:     schema.TypeString,
			Optional: true,
		},
		"rules": &schema.Schema{
			Type:     schema.TypeList,
			Optional: true,
			Elem:     ResourceNetworkSecurityRuleSchema(),
		},
		"tenant_ref": &schema.Schema{
			Type:     schema.TypeString,
			Optional: true,
			Computed: true,
		},
		"uuid": &schema.Schema{
			Type:     schema.TypeString,
			Optional: true,
			Computed: true,
		},
	}
}

func resourceAviNetworkSecurityPolicy() *schema.Resource {
	return &schema.Resource{
		Create: resourceAviNetworkSecurityPolicyCreate,
		Read:   ResourceAviNetworkSecurityPolicyRead,
		Update: resourceAviNetworkSecurityPolicyUpdate,
		Delete: resourceAviNetworkSecurityPolicyDelete,
		Schema: ResourceNetworkSecurityPolicySchema(),
		Importer: &schema.ResourceImporter{
			State: ResourceNetworkSecurityPolicyImporter,
		},
	}
}

func ResourceNetworkSecurityPolicyImporter(d *schema.ResourceData, m interface{}) ([]*schema.ResourceData, error) {
	s := ResourceNetworkSecurityPolicySchema()
	return ResourceImporter(d, m, "networksecuritypolicy", s)
}

func ResourceAviNetworkSecurityPolicyRead(d *schema.ResourceData, meta interface{}) error {
	s := ResourceNetworkSecurityPolicySchema()
	err := ApiRead(d, meta, "networksecuritypolicy", s)
	if err != nil {
		log.Printf("[ERROR] in reading object %v\n", err)
	}
	return err
}

func resourceAviNetworkSecurityPolicyCreate(d *schema.ResourceData, meta interface{}) error {
	s := ResourceNetworkSecurityPolicySchema()
	err := ApiCreateOrUpdate(d, meta, "networksecuritypolicy", s)
	if err == nil {
		err = ResourceAviNetworkSecurityPolicyRead(d, meta)
	}
	return err
}

func resourceAviNetworkSecurityPolicyUpdate(d *schema.ResourceData, meta interface{}) error {
	s := ResourceNetworkSecurityPolicySchema()
	var err error
	err = ApiCreateOrUpdate(d, meta, "networksecuritypolicy", s)
	if err == nil {
		err = ResourceAviNetworkSecurityPolicyRead(d, meta)
	}
	return err
}

func resourceAviNetworkSecurityPolicyDelete(d *schema.ResourceData, meta interface{}) error {
	objType := "networksecuritypolicy"
	if ApiDeleteSystemDefaultCheck(d) {
		return nil
	}
	client := meta.(*clients.AviClient)
	uuid := d.Get("uuid").(string)
	if uuid != "" {
		path := "api/" + objType + "/" + uuid
		err := client.AviSession.Delete(path)
		if err != nil && !(strings.Contains(err.Error(), "404") || strings.Contains(err.Error(), "204") || strings.Contains(err.Error(), "403")) {
			log.Println("[INFO] resourceAviNetworkSecurityPolicyDelete not found")
			return err
		}
		d.SetId("")
	}
	return nil
}
