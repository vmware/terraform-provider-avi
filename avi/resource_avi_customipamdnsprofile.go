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

func ResourceCustomIpamDnsProfileSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"name": {
			Type:     schema.TypeString,
			Optional: true,
			Computed: true,
		},
		"script_params": {
			Type:     schema.TypeList,
			Optional: true,
			Elem:     ResourceCustomParamsSchema(),
		},
		"script_uri": {
			Type:     schema.TypeString,
			Optional: true,
			Computed: true,
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

func resourceAviCustomIpamDnsProfile() *schema.Resource {
	return &schema.Resource{
		Create: resourceAviCustomIpamDnsProfileCreate,
		Read:   ResourceAviCustomIpamDnsProfileRead,
		Update: resourceAviCustomIpamDnsProfileUpdate,
		Delete: resourceAviCustomIpamDnsProfileDelete,
		Schema: ResourceCustomIpamDnsProfileSchema(),
		Importer: &schema.ResourceImporter{
			State: ResourceCustomIpamDnsProfileImporter,
		},
	}
}

func ResourceCustomIpamDnsProfileImporter(d *schema.ResourceData, m interface{}) ([]*schema.ResourceData, error) {
	s := ResourceCustomIpamDnsProfileSchema()
	return ResourceImporter(d, m, "customipamdnsprofile", s)
}

func ResourceAviCustomIpamDnsProfileRead(d *schema.ResourceData, meta interface{}) error {
	s := ResourceCustomIpamDnsProfileSchema()
	err := ApiRead(d, meta, "customipamdnsprofile", s)
	if err != nil {
		log.Printf("[ERROR] in reading object %v\n", err)
	}
	return err
}

func resourceAviCustomIpamDnsProfileCreate(d *schema.ResourceData, meta interface{}) error {
	s := ResourceCustomIpamDnsProfileSchema()
	err := ApiCreateOrUpdate(d, meta, "customipamdnsprofile", s)
	if err == nil {
		err = ResourceAviCustomIpamDnsProfileRead(d, meta)
	}
	return err
}

func resourceAviCustomIpamDnsProfileUpdate(d *schema.ResourceData, meta interface{}) error {
	s := ResourceCustomIpamDnsProfileSchema()
	var err error
	err = ApiCreateOrUpdate(d, meta, "customipamdnsprofile", s)
	if err == nil {
		err = ResourceAviCustomIpamDnsProfileRead(d, meta)
	}
	return err
}

func resourceAviCustomIpamDnsProfileDelete(d *schema.ResourceData, meta interface{}) error {
	objType := "customipamdnsprofile"
	client := meta.(*clients.AviClient)
	if ApiDeleteSystemDefaultCheck(d) {
		return nil
	}
	uuid := d.Get("uuid").(string)
	if uuid != "" {
		path := "api/" + objType + "/" + uuid
		err := client.AviSession.Delete(path)
		if err != nil && !(strings.Contains(err.Error(), "404") || strings.Contains(err.Error(), "204") || strings.Contains(err.Error(), "403")) {
			log.Println("[INFO] resourceAviCustomIpamDnsProfileDelete not found")
			return err
		}
		d.SetId("")
	}
	return nil
}
