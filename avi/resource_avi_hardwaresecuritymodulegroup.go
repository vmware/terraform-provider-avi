/*
 * Copyright (c) 2017. Avi Networks.
 * Author: Gaurav Rastogi (grastogi@avinetworks.com)
 *
 */
package avi

import (
	"github.com/avinetworks/sdk/go/clients"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"log"
	"strings"
)

func ResourceHardwareSecurityModuleGroupSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"hsm": {
			Type:     schema.TypeSet,
			Required: true,
			Elem:     ResourceHardwareSecurityModuleSchema(),
		},
		"name": {
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

func resourceAviHardwareSecurityModuleGroup() *schema.Resource {
	return &schema.Resource{
		Create: resourceAviHardwareSecurityModuleGroupCreate,
		Read:   ResourceAviHardwareSecurityModuleGroupRead,
		Update: resourceAviHardwareSecurityModuleGroupUpdate,
		Delete: resourceAviHardwareSecurityModuleGroupDelete,
		Schema: ResourceHardwareSecurityModuleGroupSchema(),
		Importer: &schema.ResourceImporter{
			State: ResourceHardwareSecurityModuleGroupImporter,
		},
	}
}

func ResourceHardwareSecurityModuleGroupImporter(d *schema.ResourceData, m interface{}) ([]*schema.ResourceData, error) {
	s := ResourceHardwareSecurityModuleGroupSchema()
	return ResourceImporter(d, m, "hardwaresecuritymodulegroup", s)
}

func ResourceAviHardwareSecurityModuleGroupRead(d *schema.ResourceData, meta interface{}) error {
	s := ResourceHardwareSecurityModuleGroupSchema()
	err := ApiRead(d, meta, "hardwaresecuritymodulegroup", s)
	if err != nil {
		log.Printf("[ERROR] in reading object %v\n", err)
	}
	return err
}

func resourceAviHardwareSecurityModuleGroupCreate(d *schema.ResourceData, meta interface{}) error {
	s := ResourceHardwareSecurityModuleGroupSchema()
	err := ApiCreateOrUpdate(d, meta, "hardwaresecuritymodulegroup", s)
	if err == nil {
		err = ResourceAviHardwareSecurityModuleGroupRead(d, meta)
	}
	return err
}

func resourceAviHardwareSecurityModuleGroupUpdate(d *schema.ResourceData, meta interface{}) error {
	s := ResourceHardwareSecurityModuleGroupSchema()
	var err error
	err = ApiCreateOrUpdate(d, meta, "hardwaresecuritymodulegroup", s)
	if err == nil {
		err = ResourceAviHardwareSecurityModuleGroupRead(d, meta)
	}
	return err
}

func resourceAviHardwareSecurityModuleGroupDelete(d *schema.ResourceData, meta interface{}) error {
	objType := "hardwaresecuritymodulegroup"
	if ApiDeleteSystemDefaultCheck(d) {
		return nil
	}
	client := meta.(*clients.AviClient)
	uuid := d.Get("uuid").(string)
	if uuid != "" {
		path := "api/" + objType + "/" + uuid
		err := client.AviSession.Delete(path)
		if err != nil && !(strings.Contains(err.Error(), "404") || strings.Contains(err.Error(), "204") || strings.Contains(err.Error(), "403")) {
			log.Println("[INFO] resourceAviHardwareSecurityModuleGroupDelete not found")
			return err
		}
		d.SetId("")
	}
	return nil
}
