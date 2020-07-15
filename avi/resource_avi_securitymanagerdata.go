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

func ResourceSecurityManagerDataSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"app_learning_info": {
			Type:     schema.TypeList,
			Optional: true,
			Elem:     ResourceDbAppLearningInfoSchema(),
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

func resourceAviSecurityManagerData() *schema.Resource {
	return &schema.Resource{
		Create: resourceAviSecurityManagerDataCreate,
		Read:   ResourceAviSecurityManagerDataRead,
		Update: resourceAviSecurityManagerDataUpdate,
		Delete: resourceAviSecurityManagerDataDelete,
		Schema: ResourceSecurityManagerDataSchema(),
		Importer: &schema.ResourceImporter{
			State: ResourceSecurityManagerDataImporter,
		},
	}
}

func ResourceSecurityManagerDataImporter(d *schema.ResourceData, m interface{}) ([]*schema.ResourceData, error) {
	s := ResourceSecurityManagerDataSchema()
	return ResourceImporter(d, m, "securitymanagerdata", s)
}

func ResourceAviSecurityManagerDataRead(d *schema.ResourceData, meta interface{}) error {
	s := ResourceSecurityManagerDataSchema()
	err := ApiRead(d, meta, "securitymanagerdata", s)
	if err != nil {
		log.Printf("[ERROR] in reading object %v\n", err)
	}
	return err
}

func resourceAviSecurityManagerDataCreate(d *schema.ResourceData, meta interface{}) error {
	s := ResourceSecurityManagerDataSchema()
	err := ApiCreateOrUpdate(d, meta, "securitymanagerdata", s)
	if err == nil {
		err = ResourceAviSecurityManagerDataRead(d, meta)
	}
	return err
}

func resourceAviSecurityManagerDataUpdate(d *schema.ResourceData, meta interface{}) error {
	s := ResourceSecurityManagerDataSchema()
	var err error
	err = ApiCreateOrUpdate(d, meta, "securitymanagerdata", s)
	if err == nil {
		err = ResourceAviSecurityManagerDataRead(d, meta)
	}
	return err
}

func resourceAviSecurityManagerDataDelete(d *schema.ResourceData, meta interface{}) error {
	objType := "securitymanagerdata"
	client := meta.(*clients.AviClient)
	if ApiDeleteSystemDefaultCheck(d) {
		return nil
	}
	uuid := d.Get("uuid").(string)
	if uuid != "" {
		path := "api/" + objType + "/" + uuid
		err := client.AviSession.Delete(path)
		if err != nil && !(strings.Contains(err.Error(), "404") || strings.Contains(err.Error(), "204") || strings.Contains(err.Error(), "403")) {
			log.Println("[INFO] resourceAviSecurityManagerDataDelete not found")
			return err
		}
		d.SetId("")
	}
	return nil
}
