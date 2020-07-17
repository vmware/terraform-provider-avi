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

func ResourceDnsPolicySchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"created_by": {
			Type:     schema.TypeString,
			Optional: true,
			Computed: true,
		},
		"description": {
			Type:     schema.TypeString,
			Optional: true,
			Computed: true,
		},
		"name": {
			Type:     schema.TypeString,
			Optional: true,
			Computed: true,
		},
		"rule": {
			Type:     schema.TypeList,
			Optional: true,
			Elem:     ResourceDnsRuleSchema(),
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

func resourceAviDnsPolicy() *schema.Resource {
	return &schema.Resource{
		Create: resourceAviDnsPolicyCreate,
		Read:   ResourceAviDnsPolicyRead,
		Update: resourceAviDnsPolicyUpdate,
		Delete: resourceAviDnsPolicyDelete,
		Schema: ResourceDnsPolicySchema(),
		Importer: &schema.ResourceImporter{
			State: ResourceDnsPolicyImporter,
		},
	}
}

func ResourceDnsPolicyImporter(d *schema.ResourceData, m interface{}) ([]*schema.ResourceData, error) {
	s := ResourceDnsPolicySchema()
	return ResourceImporter(d, m, "dnspolicy", s)
}

func ResourceAviDnsPolicyRead(d *schema.ResourceData, meta interface{}) error {
	s := ResourceDnsPolicySchema()
	err := ApiRead(d, meta, "dnspolicy", s)
	if err != nil {
		log.Printf("[ERROR] in reading object %v\n", err)
	}
	return err
}

func resourceAviDnsPolicyCreate(d *schema.ResourceData, meta interface{}) error {
	s := ResourceDnsPolicySchema()
	err := ApiCreateOrUpdate(d, meta, "dnspolicy", s)
	if err == nil {
		err = ResourceAviDnsPolicyRead(d, meta)
	}
	return err
}

func resourceAviDnsPolicyUpdate(d *schema.ResourceData, meta interface{}) error {
	s := ResourceDnsPolicySchema()
	var err error
	err = ApiCreateOrUpdate(d, meta, "dnspolicy", s)
	if err == nil {
		err = ResourceAviDnsPolicyRead(d, meta)
	}
	return err
}

func resourceAviDnsPolicyDelete(d *schema.ResourceData, meta interface{}) error {
	objType := "dnspolicy"
	client := meta.(*clients.AviClient)
	if ApiDeleteSystemDefaultCheck(d) {
		return nil
	}
	uuid := d.Get("uuid").(string)
	if uuid != "" {
		path := "api/" + objType + "/" + uuid
		err := client.AviSession.Delete(path)
		if err != nil && !(strings.Contains(err.Error(), "404") || strings.Contains(err.Error(), "204") || strings.Contains(err.Error(), "403")) {
			log.Println("[INFO] resourceAviDnsPolicyDelete not found")
			return err
		}
		d.SetId("")
	}
	return nil
}
