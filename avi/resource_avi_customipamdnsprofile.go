// Copyright 2019 VMware, Inc.
// SPDX-License-Identifier: Mozilla Public License 2.0

package avi

import (
	"log"
	"strings"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/vmware/alb-sdk/go/clients"
)

//nolint
func ResourceCustomIpamDnsProfileSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"configpb_attributes": {
			Type:     schema.TypeSet,
			Optional: true,
			Computed: true,
			Elem:     ResourceConfigPbAttributesSchema(),
		},
		"name": {
			Type:     schema.TypeString,
			Required: true,
		},
		"script_params": {
			Type:     schema.TypeList,
			Optional: true,
			Elem:     ResourceCustomParamsSchema(),
		},
		"script_uri": {
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

//nolint
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

//nolint
func ResourceCustomIpamDnsProfileImporter(d *schema.ResourceData, m interface{}) ([]*schema.ResourceData, error) {
	s := ResourceCustomIpamDnsProfileSchema()
	return ResourceImporter(d, m, "customipamdnsprofile", s)
}

//nolint
func ResourceAviCustomIpamDnsProfileRead(d *schema.ResourceData, meta interface{}) error {
	s := ResourceCustomIpamDnsProfileSchema()
	err := APIRead(d, meta, "customipamdnsprofile", s)
	if err != nil {
		log.Printf("[ERROR] in reading object %v\n", err)
	}
	return err
}

//nolint
func resourceAviCustomIpamDnsProfileCreate(d *schema.ResourceData, meta interface{}) error {
	s := ResourceCustomIpamDnsProfileSchema()
	err := APICreateOrUpdate(d, meta, "customipamdnsprofile", s)
	if err == nil {
		err = ResourceAviCustomIpamDnsProfileRead(d, meta)
	}
	return err
}

//nolint
func resourceAviCustomIpamDnsProfileUpdate(d *schema.ResourceData, meta interface{}) error {
	s := ResourceCustomIpamDnsProfileSchema()
	var err error
	err = APICreateOrUpdate(d, meta, "customipamdnsprofile", s)
	if err == nil {
		err = ResourceAviCustomIpamDnsProfileRead(d, meta)
	}
	return err
}

//nolint
func resourceAviCustomIpamDnsProfileDelete(d *schema.ResourceData, meta interface{}) error {
	objType := "customipamdnsprofile"
	client := meta.(*clients.AviClient)
	if APIDeleteSystemDefaultCheck(d) {
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
