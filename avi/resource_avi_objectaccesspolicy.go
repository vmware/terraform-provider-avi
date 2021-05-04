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

func ResourceObjectAccessPolicySchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"name": {
			Type:     schema.TypeString,
			Required: true,
		},
		"rules": {
			Type:     schema.TypeList,
			Required: true,
			Elem:     ResourceObjectAccessPolicyRuleSchema(),
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

func resourceAviObjectAccessPolicy() *schema.Resource {
	return &schema.Resource{
		Create: resourceAviObjectAccessPolicyCreate,
		Read:   ResourceAviObjectAccessPolicyRead,
		Update: resourceAviObjectAccessPolicyUpdate,
		Delete: resourceAviObjectAccessPolicyDelete,
		Schema: ResourceObjectAccessPolicySchema(),
		Importer: &schema.ResourceImporter{
			State: ResourceObjectAccessPolicyImporter,
		},
	}
}

func ResourceObjectAccessPolicyImporter(d *schema.ResourceData, m interface{}) ([]*schema.ResourceData, error) {
	s := ResourceObjectAccessPolicySchema()
	return ResourceImporter(d, m, "objectaccesspolicy", s)
}

func ResourceAviObjectAccessPolicyRead(d *schema.ResourceData, meta interface{}) error {
	s := ResourceObjectAccessPolicySchema()
	err := APIRead(d, meta, "objectaccesspolicy", s)
	if err != nil {
		log.Printf("[ERROR] in reading object %v\n", err)
	}
	return err
}

func resourceAviObjectAccessPolicyCreate(d *schema.ResourceData, meta interface{}) error {
	s := ResourceObjectAccessPolicySchema()
	err := APICreateOrUpdate(d, meta, "objectaccesspolicy", s)
	if err == nil {
		err = ResourceAviObjectAccessPolicyRead(d, meta)
	}
	return err
}

func resourceAviObjectAccessPolicyUpdate(d *schema.ResourceData, meta interface{}) error {
	s := ResourceObjectAccessPolicySchema()
	var err error
	err = APICreateOrUpdate(d, meta, "objectaccesspolicy", s)
	if err == nil {
		err = ResourceAviObjectAccessPolicyRead(d, meta)
	}
	return err
}

func resourceAviObjectAccessPolicyDelete(d *schema.ResourceData, meta interface{}) error {
	objType := "objectaccesspolicy"
	client := meta.(*clients.AviClient)
	if APIDeleteSystemDefaultCheck(d) {
		return nil
	}
	uuid := d.Get("uuid").(string)
	if uuid != "" {
		path := "api/" + objType + "/" + uuid
		err := client.AviSession.Delete(path)
		if err != nil && !(strings.Contains(err.Error(), "404") || strings.Contains(err.Error(), "204") || strings.Contains(err.Error(), "403")) {
			log.Println("[INFO] resourceAviObjectAccessPolicyDelete not found")
			return err
		}
		d.SetId("")
	}
	return nil
}
