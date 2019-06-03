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

func ResourceNatPolicySchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"created_by": {
			Type:     schema.TypeString,
			Optional: true,
		},
		"description": {
			Type:     schema.TypeString,
			Optional: true,
		},
		"name": {
			Type:     schema.TypeString,
			Optional: true,
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
	err := ApiRead(d, meta, "natpolicy", s)
	if err != nil {
		log.Printf("[ERROR] in reading object %v\n", err)
	}
	return err
}

func resourceAviNatPolicyCreate(d *schema.ResourceData, meta interface{}) error {
	s := ResourceNatPolicySchema()
	err := ApiCreateOrUpdate(d, meta, "natpolicy", s)
	if err == nil {
		err = ResourceAviNatPolicyRead(d, meta)
	}
	return err
}

func resourceAviNatPolicyUpdate(d *schema.ResourceData, meta interface{}) error {
	s := ResourceNatPolicySchema()
	var err error
	err = ApiCreateOrUpdate(d, meta, "natpolicy", s)
	if err == nil {
		err = ResourceAviNatPolicyRead(d, meta)
	}
	return err
}

func resourceAviNatPolicyDelete(d *schema.ResourceData, meta interface{}) error {
	objType := "natpolicy"
	if ApiDeleteSystemDefaultCheck(d) {
		return nil
	}
	client := meta.(*clients.AviClient)
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
