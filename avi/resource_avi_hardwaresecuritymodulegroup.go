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
	log.Printf("[INFO] ResourceAviHardwareSecurityModuleGroupRead Avi Client %v\n", d)
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
	// no need to set the ID
	log.Printf("ResourceAviHardwareSecurityModuleGroupRead CURRENT obj %v\n", d)

	log.Printf("ResourceAviHardwareSecurityModuleGroupRead Read API obj %v\n", obj)
	if tObj, err := ApiDataToSchema(obj, d, s); err == nil {
		log.Printf("[INFO] ResourceAviHardwareSecurityModuleGroupRead Converted obj %v\n", tObj)
		//err = d.Set("obj", tObj)
		if err != nil {
			log.Printf("[ERROR] in setting read object %v\n", err)
		}
	}
	log.Printf("[INFO] ResourceAviHardwareSecurityModuleGroupRead Updated %v\n", d)
	return nil
}

func resourceAviHardwareSecurityModuleGroupCreate(d *schema.ResourceData, meta interface{}) error {
	s := ResourceHardwareSecurityModuleGroupSchema()
	err := ApiCreateOrUpdate(d, meta, "hardwaresecuritymodulegroup", s)
	log.Printf("[DEBUG] created object %v: %v", "hardwaresecuritymodulegroup", d)
	if err == nil {
		err = ResourceAviHardwareSecurityModuleGroupRead(d, meta)
	}
	log.Printf("[DEBUG] created object %v: %v", "hardwaresecuritymodulegroup", d)
	return err
}

func resourceAviHardwareSecurityModuleGroupUpdate(d *schema.ResourceData, meta interface{}) error {
	s := ResourceHardwareSecurityModuleGroupSchema()
	err := ApiCreateOrUpdate(d, meta, "hardwaresecuritymodulegroup", s)
	if err == nil {
		err = ResourceAviHardwareSecurityModuleGroupRead(d, meta)
	}
	log.Printf("[DEBUG] updated object %v: %v", "hardwaresecuritymodulegroup", d)
	return err
}

func resourceAviHardwareSecurityModuleGroupDelete(d *schema.ResourceData, meta interface{}) error {
	objType := "hardwaresecuritymodulegroup"
	log.Println("[INFO] ResourceAviHardwareSecurityModuleGroupRead Avi Client")
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
