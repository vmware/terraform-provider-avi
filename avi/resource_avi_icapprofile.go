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

func ResourceIcapProfileSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"allow_204": {
			Type:     schema.TypeBool,
			Optional: true,
			Default:  true,
		},
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
		"configpb_attributes": {
			Type:     schema.TypeSet,
			Optional: true,
			Computed: true,
			Elem:     ResourceConfigPbAttributesSchema(),
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
			Required: true,
		},
		"nsx_defender_config": {
			Type:     schema.TypeSet,
			Optional: true,
			Computed: true,
			Elem:     ResourceIcapNsxDefenderConfigSchema(),
		},
		"pool_group_ref": {
			Type:     schema.TypeString,
			Required: true,
		},
		"preview_size": {
			Type:     schema.TypeInt,
			Optional: true,
			Default:  5000,
		},
		"response_timeout": {
			Type:     schema.TypeInt,
			Optional: true,
			Default:  60000,
		},
		"service_uri": {
			Type:     schema.TypeString,
			Required: true,
		},
		"slow_response_warning_threshold": {
			Type:     schema.TypeInt,
			Optional: true,
			Default:  10000,
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
	err := APIRead(d, meta, "icapprofile", s)
	if err != nil {
		log.Printf("[ERROR] in reading object %v\n", err)
	}
	return err
}

func resourceAviIcapProfileCreate(d *schema.ResourceData, meta interface{}) error {
	s := ResourceIcapProfileSchema()
	err := APICreateOrUpdate(d, meta, "icapprofile", s)
	if err == nil {
		err = ResourceAviIcapProfileRead(d, meta)
	}
	return err
}

func resourceAviIcapProfileUpdate(d *schema.ResourceData, meta interface{}) error {
	s := ResourceIcapProfileSchema()
	var err error
	err = APICreateOrUpdate(d, meta, "icapprofile", s)
	if err == nil {
		err = ResourceAviIcapProfileRead(d, meta)
	}
	return err
}

func resourceAviIcapProfileDelete(d *schema.ResourceData, meta interface{}) error {
	objType := "icapprofile"
	client := meta.(*clients.AviClient)
	if APIDeleteSystemDefaultCheck(d) {
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
