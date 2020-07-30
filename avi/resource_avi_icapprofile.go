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

func ResourceIcapProfileSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"buffer_size": {
			Type:     schema.TypeInt,
			Optional: true,
			Default:  51200,
		},
		"buffer_size_exceed_action": {
			Type:     schema.TypeString,
			Optional: true,
			Default:  "ICAP_FAIL_OPEN",
		},
		"cloud_ref": {
			Type:     schema.TypeString,
			Optional: true,
			Computed: true,
		},
		"description": {
			Type:     schema.TypeString,
			Optional: true,
			Computed: true,
		},
		"enable_preview": {
			Type:     schema.TypeBool,
			Optional: true,
			Default:  true,
		},
		"fail_action": {
			Type:     schema.TypeString,
			Optional: true,
			Default:  "ICAP_FAIL_OPEN",
		},
		"name": {
			Type:     schema.TypeString,
			Optional: true,
			Computed: true,
		},
		"pool_group_ref": {
			Type:     schema.TypeString,
			Optional: true,
			Computed: true,
		},
		"preview_size": {
			Type:     schema.TypeInt,
			Optional: true,
			Default:  5000,
		},
		"response_timeout": {
			Type:     schema.TypeInt,
			Optional: true,
			Default:  1000,
		},
		"service_uri": {
			Type:     schema.TypeString,
			Optional: true,
			Computed: true,
		},
		"slow_response_warning_threshold": {
			Type:     schema.TypeInt,
			Optional: true,
			Default:  500,
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
		"vendor": {
			Type:     schema.TypeString,
			Optional: true,
			Default:  "ICAP_VENDOR_OPSWAT",
		},
	}
}

func resourceAviIcapProfile() *schema.Resource {
	return &schema.Resource{
		Create: resourceAviIcapProfileCreate,
		Read:   ResourceAviIcapProfileRead,
		Update: resourceAviIcapProfileUpdate,
		Delete: resourceAviIcapProfileDelete,
		Schema: ResourceIcapProfileSchema(),
		Importer: &schema.ResourceImporter{
			State: ResourceIcapProfileImporter,
		},
	}
}

func ResourceIcapProfileImporter(d *schema.ResourceData, m interface{}) ([]*schema.ResourceData, error) {
	s := ResourceIcapProfileSchema()
	return ResourceImporter(d, m, "icapprofile", s)
}

func ResourceAviIcapProfileRead(d *schema.ResourceData, meta interface{}) error {
	s := ResourceIcapProfileSchema()
	err := ApiRead(d, meta, "icapprofile", s)
	if err != nil {
		log.Printf("[ERROR] in reading object %v\n", err)
	}
	return err
}

func resourceAviIcapProfileCreate(d *schema.ResourceData, meta interface{}) error {
	s := ResourceIcapProfileSchema()
	err := ApiCreateOrUpdate(d, meta, "icapprofile", s)
	if err == nil {
		err = ResourceAviIcapProfileRead(d, meta)
	}
	return err
}

func resourceAviIcapProfileUpdate(d *schema.ResourceData, meta interface{}) error {
	s := ResourceIcapProfileSchema()
	var err error
	err = ApiCreateOrUpdate(d, meta, "icapprofile", s)
	if err == nil {
		err = ResourceAviIcapProfileRead(d, meta)
	}
	return err
}

func resourceAviIcapProfileDelete(d *schema.ResourceData, meta interface{}) error {
	objType := "icapprofile"
	client := meta.(*clients.AviClient)
	if ApiDeleteSystemDefaultCheck(d) {
		return nil
	}
	uuid := d.Get("uuid").(string)
	if uuid != "" {
		path := "api/" + objType + "/" + uuid
		err := client.AviSession.Delete(path)
		if err != nil && !(strings.Contains(err.Error(), "404") || strings.Contains(err.Error(), "204") || strings.Contains(err.Error(), "403")) {
			log.Println("[INFO] resourceAviIcapProfileDelete not found")
			return err
		}
		d.SetId("")
	}
	return nil
}
