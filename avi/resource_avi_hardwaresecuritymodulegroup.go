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

func ResourceHardwareSecurityModuleGroupSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"hsm": &schema.Schema{
			Type:     schema.TypeSet,
			Required: true, Elem: ResourceHardwareSecurityModuleSchema(),
			Set: func(v interface{}) int {
				return 0
			},
		},
		"name": &schema.Schema{
			Type:     schema.TypeString,
			Required: true,
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

func resourceAviHardwareSecurityModuleGroup() *schema.Resource {
	return &schema.Resource{
		Create: resourceAviHardwareSecurityModuleGroupCreate,
		Read:   ResourceAviHardwareSecurityModuleGroupRead,
		Update: resourceAviHardwareSecurityModuleGroupUpdate,
		Delete: resourceAviHardwareSecurityModuleGroupDelete,
		Schema: ResourceHardwareSecurityModuleGroupSchema(),
	}
}

func ResourceAviHardwareSecurityModuleGroupRead(d *schema.ResourceData, meta interface{}) error {
	s := ResourceHardwareSecurityModuleGroupSchema()
	client := meta.(*clients.AviClient)
	var obj interface{}
	if uuid, ok := d.GetOk("uuid"); ok {
		path := "api/hardwaresecuritymodulegroup/" + uuid.(string)
		err := client.AviSession.Get(path, &obj)
		if err != nil {
			d.SetId("")
			return nil
		}
	} else {
		d.SetId("")
		return nil
	}
	if _, err := ApiDataToSchema(obj, d, s); err == nil {
		if err != nil {
			log.Printf("[ERROR] in setting read object %v\n", err)
		}
	}
	return nil
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
	err := ApiCreateOrUpdate(d, meta, "hardwaresecuritymodulegroup", s)
	if err == nil {
		err = ResourceAviHardwareSecurityModuleGroupRead(d, meta)
	}
	return err
}

func resourceAviHardwareSecurityModuleGroupDelete(d *schema.ResourceData, meta interface{}) error {
	objType := "hardwaresecuritymodulegroup"
	client := meta.(*clients.AviClient)
	uuid := d.Get("uuid").(string)
	if uuid != "" {
		path := "api/" + objType + "/" + uuid
		err := client.AviSession.Delete(path)
		if err != nil && !(strings.Contains(err.Error(), "404") || strings.Contains(err.Error(), "204")) {
			log.Println("[INFO] resourceAviHardwareSecurityModuleGroupDelete not found")
			return err
		}
		d.SetId("")
	}
	return nil
}
