/*
* Copyright (c) 2017. Avi Networks.
* Author: Gaurav Rastogi (grastogi@avinetworks.com)
*
 */
package avi

import (
	"log"
	"strings"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/vmware/alb-sdk/go/clients"
)

//nolint
func ResourceDnsPolicySchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"configpb_attributes": {
			Type:     schema.TypeSet,
			Optional: true,
			Computed: true,
			Elem:     ResourceConfigPbAttributesSchema(),
		},
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
		"internal": {
			Type:     schema.TypeBool,
			Optional: true,
			Computed: true,
		},
		"markers": {
			Type:     schema.TypeList,
			Optional: true,
			Elem:     ResourceRoleFilterMatchLabelSchema(),
		},
		"name": {
			Type:     schema.TypeString,
			Required: true,
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

//nolint
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

//nolint
func ResourceDnsPolicyImporter(d *schema.ResourceData, m interface{}) ([]*schema.ResourceData, error) {
	s := ResourceDnsPolicySchema()
	return ResourceImporter(d, m, "dnspolicy", s)
}

//nolint
func ResourceAviDnsPolicyRead(d *schema.ResourceData, meta interface{}) error {
	s := ResourceDnsPolicySchema()
	err := APIRead(d, meta, "dnspolicy", s)
	if err != nil {
		log.Printf("[ERROR] in reading object %v\n", err)
	}
	return err
}

//nolint
func resourceAviDnsPolicyCreate(d *schema.ResourceData, meta interface{}) error {
	s := ResourceDnsPolicySchema()
	err := APICreateOrUpdate(d, meta, "dnspolicy", s)
	if err == nil {
		err = ResourceAviDnsPolicyRead(d, meta)
	}
	return err
}

//nolint
func resourceAviDnsPolicyUpdate(d *schema.ResourceData, meta interface{}) error {
	s := ResourceDnsPolicySchema()
	var err error
	err = APICreateOrUpdate(d, meta, "dnspolicy", s)
	if err == nil {
		err = ResourceAviDnsPolicyRead(d, meta)
	}
	return err
}

//nolint
func resourceAviDnsPolicyDelete(d *schema.ResourceData, meta interface{}) error {
	objType := "dnspolicy"
	client := meta.(*clients.AviClient)
	if APIDeleteSystemDefaultCheck(d) {
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
