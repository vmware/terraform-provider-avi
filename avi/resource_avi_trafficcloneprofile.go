/*
* Copyright (c) 2017. Avi Networks.
* Author: Gaurav Rastogi (grastogi@avinetworks.com)
*
 */
package avi

import (
	"log"
	"strings"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/vmware/alb-sdk/go/clients"
)

func ResourceTrafficCloneProfileSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"clone_servers": {
			Type:     schema.TypeList,
			Optional: true,
			Elem:     ResourceCloneServerSchema(),
		},
		"cloud_ref": {
			Type:     schema.TypeString,
			Optional: true,
			Computed: true,
		},
		"configpb_attributes": {
			Type:     schema.TypeSet,
			Optional: true,
			Computed: true,
			Elem:     ResourceConfigPbAttributesSchema(),
		},
		"markers": {
			Type:     schema.TypeList,
			Optional: true,
			Elem:     ResourceRoleFilterMatchLabelSchema(),
		},
		"name": {
			Type:     schema.TypeString,
			Required: true,
		},
		"preserve_client_ip": {
			Type:     schema.TypeBool,
			Optional: true,
			Default:  false,
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

func resourceAviTrafficCloneProfile() *schema.Resource {
	return &schema.Resource{
		Create: resourceAviTrafficCloneProfileCreate,
		Read:   ResourceAviTrafficCloneProfileRead,
		Update: resourceAviTrafficCloneProfileUpdate,
		Delete: resourceAviTrafficCloneProfileDelete,
		Schema: ResourceTrafficCloneProfileSchema(),
		Importer: &schema.ResourceImporter{
			State: ResourceTrafficCloneProfileImporter,
		},
	}
}

func ResourceTrafficCloneProfileImporter(d *schema.ResourceData, m interface{}) ([]*schema.ResourceData, error) {
	s := ResourceTrafficCloneProfileSchema()
	return ResourceImporter(d, m, "trafficcloneprofile", s)
}

func ResourceAviTrafficCloneProfileRead(d *schema.ResourceData, meta interface{}) error {
	s := ResourceTrafficCloneProfileSchema()
	err := APIRead(d, meta, "trafficcloneprofile", s)
	if err != nil {
		log.Printf("[ERROR] in reading object %v\n", err)
	}
	return err
}

func resourceAviTrafficCloneProfileCreate(d *schema.ResourceData, meta interface{}) error {
	s := ResourceTrafficCloneProfileSchema()
	err := APICreateOrUpdate(d, meta, "trafficcloneprofile", s)
	if err == nil {
		err = ResourceAviTrafficCloneProfileRead(d, meta)
	}
	return err
}

func resourceAviTrafficCloneProfileUpdate(d *schema.ResourceData, meta interface{}) error {
	s := ResourceTrafficCloneProfileSchema()
	var err error
	err = APICreateOrUpdate(d, meta, "trafficcloneprofile", s)
	if err == nil {
		err = ResourceAviTrafficCloneProfileRead(d, meta)
	}
	return err
}

func resourceAviTrafficCloneProfileDelete(d *schema.ResourceData, meta interface{}) error {
	objType := "trafficcloneprofile"
	client := meta.(*clients.AviClient)
	if APIDeleteSystemDefaultCheck(d) {
		return nil
	}
	uuid := d.Get("uuid").(string)
	if uuid != "" {
		path := "api/" + objType + "/" + uuid
		err := client.AviSession.Delete(path)
		if err != nil && !(strings.Contains(err.Error(), "404") || strings.Contains(err.Error(), "204") || strings.Contains(err.Error(), "403")) {
			log.Println("[INFO] resourceAviTrafficCloneProfileDelete not found")
			return err
		}
		d.SetId("")
	}
	return nil
}
