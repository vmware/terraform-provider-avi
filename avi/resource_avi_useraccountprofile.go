// Copyright 2019 VMware, Inc.
// SPDX-License-Identifier: Mozilla Public License 2.0

package avi

import (
	"log"
	"strings"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/vmware/alb-sdk/go/clients"
)

func ResourceUserAccountProfileSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"account_lock_timeout": {
			Type:     schema.TypeInt,
			Optional: true,
			Default:  30,
		},
		"configpb_attributes": {
			Type:     schema.TypeSet,
			Optional: true,
			Computed: true,
			Elem:     ResourceConfigPbAttributesSchema(),
		},
		"credentials_timeout_threshold": {
			Type:     schema.TypeInt,
			Optional: true,
			Default:  180,
		},
		"max_concurrent_sessions": {
			Type:     schema.TypeInt,
			Optional: true,
			Default:  0,
		},
		"max_login_failure_count": {
			Type:     schema.TypeInt,
			Optional: true,
			Default:  3,
		},
		"max_password_history_count": {
			Type:     schema.TypeInt,
			Optional: true,
			Default:  4,
		},
		"name": {
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

func resourceAviUserAccountProfile() *schema.Resource {
	return &schema.Resource{
		Create: resourceAviUserAccountProfileCreate,
		Read:   ResourceAviUserAccountProfileRead,
		Update: resourceAviUserAccountProfileUpdate,
		Delete: resourceAviUserAccountProfileDelete,
		Schema: ResourceUserAccountProfileSchema(),
		Importer: &schema.ResourceImporter{
			State: ResourceUserAccountProfileImporter,
		},
	}
}

func ResourceUserAccountProfileImporter(d *schema.ResourceData, m interface{}) ([]*schema.ResourceData, error) {
	s := ResourceUserAccountProfileSchema()
	return ResourceImporter(d, m, "useraccountprofile", s)
}

func ResourceAviUserAccountProfileRead(d *schema.ResourceData, meta interface{}) error {
	s := ResourceUserAccountProfileSchema()
	err := APIRead(d, meta, "useraccountprofile", s)
	if err != nil {
		log.Printf("[ERROR] in reading object %v\n", err)
	}
	return err
}

func resourceAviUserAccountProfileCreate(d *schema.ResourceData, meta interface{}) error {
	s := ResourceUserAccountProfileSchema()
	err := APICreateOrUpdate(d, meta, "useraccountprofile", s)
	if err == nil {
		err = ResourceAviUserAccountProfileRead(d, meta)
	}
	return err
}

func resourceAviUserAccountProfileUpdate(d *schema.ResourceData, meta interface{}) error {
	s := ResourceUserAccountProfileSchema()
	var err error
	err = APICreateOrUpdate(d, meta, "useraccountprofile", s)
	if err == nil {
		err = ResourceAviUserAccountProfileRead(d, meta)
	}
	return err
}

func resourceAviUserAccountProfileDelete(d *schema.ResourceData, meta interface{}) error {
	objType := "useraccountprofile"
	client := meta.(*clients.AviClient)
	if APIDeleteSystemDefaultCheck(d) {
		return nil
	}
	uuid := d.Get("uuid").(string)
	if uuid != "" {
		path := "api/" + objType + "/" + uuid
		err := client.AviSession.Delete(path)
		if err != nil && !(strings.Contains(err.Error(), "404") || strings.Contains(err.Error(), "204") || strings.Contains(err.Error(), "403")) {
			log.Println("[INFO] resourceAviUserAccountProfileDelete not found")
			return err
		}
		d.SetId("")
	}
	return nil
}
