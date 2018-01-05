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

func ResourceTrafficCloneProfileSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"clone_servers": &schema.Schema{
			Type:     schema.TypeList,
			Optional: true,
			Elem:     ResourceCloneServerSchema(),
		},
		"cloud_ref": &schema.Schema{
			Type:     schema.TypeString,
			Optional: true,
			Default:  "/api/cloud?name=Default-Cloud",
		},
		"name": &schema.Schema{
			Type:     schema.TypeString,
			Required: true,
		},
		"preserve_client_ip": &schema.Schema{
			Type:     schema.TypeBool,
			Optional: true,
			Default:  false,
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

func resourceAviTrafficCloneProfile() *schema.Resource {
	return &schema.Resource{
		Create: resourceAviTrafficCloneProfileCreate,
		Read:   ResourceAviTrafficCloneProfileRead,
		Update: resourceAviTrafficCloneProfileUpdate,
		Delete: resourceAviTrafficCloneProfileDelete,
		Schema: ResourceTrafficCloneProfileSchema(),
	}
}

func ResourceAviTrafficCloneProfileRead(d *schema.ResourceData, meta interface{}) error {
	s := ResourceTrafficCloneProfileSchema()
	err := ApiRead(d, meta, "trafficcloneprofile", s)
	return err
}

func resourceAviTrafficCloneProfileCreate(d *schema.ResourceData, meta interface{}) error {
	s := ResourceTrafficCloneProfileSchema()
	err := ApiCreateOrUpdate(d, meta, "trafficcloneprofile", s)
	if err == nil {
		err = ResourceAviTrafficCloneProfileRead(d, meta)
	}
	return err
}

func resourceAviTrafficCloneProfileUpdate(d *schema.ResourceData, meta interface{}) error {
	s := ResourceTrafficCloneProfileSchema()
	err := ApiCreateOrUpdate(d, meta, "trafficcloneprofile", s)
	if err == nil {
		err = ResourceAviTrafficCloneProfileRead(d, meta)
	}
	return err
}

func resourceAviTrafficCloneProfileDelete(d *schema.ResourceData, meta interface{}) error {
	objType := "trafficcloneprofile"
	client := meta.(*clients.AviClient)
	uuid := d.Get("uuid").(string)
	if uuid != "" {
		path := "api/" + objType + "/" + uuid
		err := client.AviSession.Delete(path)
		if err != nil && !(strings.Contains(err.Error(), "404") || strings.Contains(err.Error(), "204")) {
			log.Println("[INFO] resourceAviTrafficCloneProfileDelete not found")
			return err
		}
		d.SetId("")
	}
	return nil
}
