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

func ResourceSnmpTrapProfileSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"name": {
			Type:     schema.TypeString,
			Required: true,
		},
		"tenant_ref": {
			Type:     schema.TypeString,
			Optional: true,
			Computed: true,
		},
		"trap_servers": {
			Type:     schema.TypeList,
			Optional: true,
			Elem:     ResourceSnmpTrapServerSchema(),
		},
		"uuid": {
			Type:     schema.TypeString,
			Optional: true,
			Computed: true,
		},
	}
}

func resourceAviSnmpTrapProfile() *schema.Resource {
	return &schema.Resource{
		Create: resourceAviSnmpTrapProfileCreate,
		Read:   ResourceAviSnmpTrapProfileRead,
		Update: resourceAviSnmpTrapProfileUpdate,
		Delete: resourceAviSnmpTrapProfileDelete,
		Schema: ResourceSnmpTrapProfileSchema(),
		Importer: &schema.ResourceImporter{
			State: ResourceSnmpTrapProfileImporter,
		},
	}
}

func ResourceSnmpTrapProfileImporter(d *schema.ResourceData, m interface{}) ([]*schema.ResourceData, error) {
	s := ResourceSnmpTrapProfileSchema()
	return ResourceImporter(d, m, "snmptrapprofile", s)
}

func ResourceAviSnmpTrapProfileRead(d *schema.ResourceData, meta interface{}) error {
	s := ResourceSnmpTrapProfileSchema()
	err := ApiRead(d, meta, "snmptrapprofile", s)
	if err != nil {
		log.Printf("[ERROR] in reading object %v\n", err)
	}
	return err
}

func resourceAviSnmpTrapProfileCreate(d *schema.ResourceData, meta interface{}) error {
	s := ResourceSnmpTrapProfileSchema()
	err := ApiCreateOrUpdate(d, meta, "snmptrapprofile", s)
	if err == nil {
		err = ResourceAviSnmpTrapProfileRead(d, meta)
	}
	return err
}

func resourceAviSnmpTrapProfileUpdate(d *schema.ResourceData, meta interface{}) error {
	s := ResourceSnmpTrapProfileSchema()
	var err error
	err = ApiCreateOrUpdate(d, meta, "snmptrapprofile", s)
	if err == nil {
		err = ResourceAviSnmpTrapProfileRead(d, meta)
	}
	return err
}

func resourceAviSnmpTrapProfileDelete(d *schema.ResourceData, meta interface{}) error {
	objType := "snmptrapprofile"
	if ApiDeleteSystemDefaultCheck(d) {
		return nil
	}
	client := meta.(*clients.AviClient)
	uuid := d.Get("uuid").(string)
	if uuid != "" {
		path := "api/" + objType + "/" + uuid
		err := client.AviSession.Delete(path)
		if err != nil && !(strings.Contains(err.Error(), "404") || strings.Contains(err.Error(), "204") || strings.Contains(err.Error(), "403")) {
			log.Println("[INFO] resourceAviSnmpTrapProfileDelete not found")
			return err
		}
		d.SetId("")
	}
	return nil
}
