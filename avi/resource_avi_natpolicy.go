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

func ResourceNatPolicySchema() map[string]*schema.Schema {
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
		"markers": {
			Type:     schema.TypeList,
			Optional: true,
			Elem:     ResourceRoleFilterMatchLabelSchema(),
		},
		"name": {
			Type:     schema.TypeString,
			Optional: true,
			Computed: true,
		},
		"rules": {
			Type:     schema.TypeList,
			Optional: true,
			Elem:     ResourceNatRuleSchema(),
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

func resourceAviNatPolicy() *schema.Resource {
	return &schema.Resource{
		Create: resourceAviNatPolicyCreate,
		Read:   ResourceAviNatPolicyRead,
		Update: resourceAviNatPolicyUpdate,
		Delete: resourceAviNatPolicyDelete,
		Schema: ResourceNatPolicySchema(),
		Importer: &schema.ResourceImporter{
			State: ResourceNatPolicyImporter,
		},
	}
}

func ResourceNatPolicyImporter(d *schema.ResourceData, m interface{}) ([]*schema.ResourceData, error) {
	s := ResourceNatPolicySchema()
	return ResourceImporter(d, m, "natpolicy", s)
}

func ResourceAviNatPolicyRead(d *schema.ResourceData, meta interface{}) error {
	s := ResourceNatPolicySchema()
	err := APIRead(d, meta, "natpolicy", s)
	if err != nil {
		log.Printf("[ERROR] in reading object %v\n", err)
	}
	return err
}

func resourceAviNatPolicyCreate(d *schema.ResourceData, meta interface{}) error {
	s := ResourceNatPolicySchema()
	err := APICreateOrUpdate(d, meta, "natpolicy", s)
	if err == nil {
		err = ResourceAviNatPolicyRead(d, meta)
	}
	return err
}

func resourceAviNatPolicyUpdate(d *schema.ResourceData, meta interface{}) error {
	s := ResourceNatPolicySchema()
	var err error
	err = APICreateOrUpdate(d, meta, "natpolicy", s)
	if err == nil {
		err = ResourceAviNatPolicyRead(d, meta)
	}
	return err
}

func resourceAviNatPolicyDelete(d *schema.ResourceData, meta interface{}) error {
	objType := "natpolicy"
	client := meta.(*clients.AviClient)
	if APIDeleteSystemDefaultCheck(d) {
		return nil
	}
	uuid := d.Get("uuid").(string)
	if uuid != "" {
		path := "api/" + objType + "/" + uuid
		err := client.AviSession.Delete(path)
		if err != nil && !(strings.Contains(err.Error(), "404") || strings.Contains(err.Error(), "204") || strings.Contains(err.Error(), "403")) {
			log.Println("[INFO] resourceAviNatPolicyDelete not found")
			return err
		}
		d.SetId("")
	}
	return nil
}
