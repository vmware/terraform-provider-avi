// Copyright 2019 VMware, Inc.
// SPDX-License-Identifier: Mozilla Public License 2.0

package avi

import (
	"log"
	"strings"

	"github.com/avinetworks/sdk/go/clients"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func ResourceAuthProfileSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"description": {
			Type:     schema.TypeString,
			Optional: true,
			Computed: true,
		},
		"http": {
			Type:     schema.TypeSet,
			Optional: true,
			Computed: true,
			Elem:     ResourceAuthProfileHTTPClientParamsSchema(),
		},
		"ldap": {
			Type:     schema.TypeSet,
			Optional: true,
			Computed: true,
			Elem:     ResourceLdapAuthSettingsSchema(),
		},
		"name": {
			Type:     schema.TypeString,
			Required: true,
		},
		"pa_agent_ref": {
			Type:     schema.TypeString,
			Optional: true,
			Computed: true,
		},
		"saml": {
			Type:     schema.TypeSet,
			Optional: true,
			Computed: true,
			Elem:     ResourceSamlSettingsSchema(),
		},
		"tacacs_plus": {
			Type:     schema.TypeSet,
			Optional: true,
			Computed: true,
			Elem:     ResourceTacacsPlusAuthSettingsSchema(),
		},
		"tenant_ref": {
			Type:     schema.TypeString,
			Optional: true,
			Computed: true,
		},
		"type": {
			Type:     schema.TypeString,
			Required: true,
		},
		"uuid": {
			Type:     schema.TypeString,
			Optional: true,
			Computed: true,
		},
	}
}

func resourceAviAuthProfile() *schema.Resource {
	return &schema.Resource{
		Create: resourceAviAuthProfileCreate,
		Read:   ResourceAviAuthProfileRead,
		Update: resourceAviAuthProfileUpdate,
		Delete: resourceAviAuthProfileDelete,
		Schema: ResourceAuthProfileSchema(),
		Importer: &schema.ResourceImporter{
			State: ResourceAuthProfileImporter,
		},
	}
}

func ResourceAuthProfileImporter(d *schema.ResourceData, m interface{}) ([]*schema.ResourceData, error) {
	s := ResourceAuthProfileSchema()
	return ResourceImporter(d, m, "authprofile", s)
}

func ResourceAviAuthProfileRead(d *schema.ResourceData, meta interface{}) error {
	s := ResourceAuthProfileSchema()
	err := APIRead(d, meta, "authprofile", s)
	if err != nil {
		log.Printf("[ERROR] in reading object %v\n", err)
	}
	return err
}

func resourceAviAuthProfileCreate(d *schema.ResourceData, meta interface{}) error {
	s := ResourceAuthProfileSchema()
	err := APICreateOrUpdate(d, meta, "authprofile", s)
	if err == nil {
		err = ResourceAviAuthProfileRead(d, meta)
	}
	return err
}

func resourceAviAuthProfileUpdate(d *schema.ResourceData, meta interface{}) error {
	s := ResourceAuthProfileSchema()
	var err error
	err = APICreateOrUpdate(d, meta, "authprofile", s)
	if err == nil {
		err = ResourceAviAuthProfileRead(d, meta)
	}
	return err
}

func resourceAviAuthProfileDelete(d *schema.ResourceData, meta interface{}) error {
	objType := "authprofile"
	client := meta.(*clients.AviClient)
	if APIDeleteSystemDefaultCheck(d) {
		return nil
	}
	uuid := d.Get("uuid").(string)
	if uuid != "" {
		path := "api/" + objType + "/" + uuid
		err := client.AviSession.Delete(path)
		if err != nil && !(strings.Contains(err.Error(), "404") || strings.Contains(err.Error(), "204") || strings.Contains(err.Error(), "403")) {
			log.Println("[INFO] resourceAviAuthProfileDelete not found")
			return err
		}
		d.SetId("")
	}
	return nil
}
