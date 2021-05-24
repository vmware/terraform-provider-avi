// Copyright 2019 VMware, Inc.
// SPDX-License-Identifier: Mozilla Public License 2.0

package avi

import (
	"log"
	"strings"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/vmware/alb-sdk/go/clients"
)

func ResourceApplicationPersistenceProfileSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"app_cookie_persistence_profile": {
			Type:     schema.TypeSet,
			Optional: true,
			Computed: true,
			Elem:     ResourceAppCookiePersistenceProfileSchema(),
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
		"hdr_persistence_profile": {
			Type:     schema.TypeSet,
			Optional: true,
			Computed: true,
			Elem:     ResourceHdrPersistenceProfileSchema(),
		},
		"http_cookie_persistence_profile": {
			Type:     schema.TypeSet,
			Optional: true,
			Computed: true,
			Elem:     ResourceHttpCookiePersistenceProfileSchema(),
		},
		"ip_persistence_profile": {
			Type:     schema.TypeSet,
			Optional: true,
			Computed: true,
			Elem:     ResourceIPPersistenceProfileSchema(),
		},
		"is_federated": {
			Type:     schema.TypeBool,
			Optional: true,
			Default:  false,
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
		"persistence_type": {
			Type:     schema.TypeString,
			Required: true,
		},
		"server_hm_down_recovery": {
			Type:     schema.TypeString,
			Optional: true,
			Default:  "HM_DOWN_PICK_NEW_SERVER",
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

func resourceAviApplicationPersistenceProfile() *schema.Resource {
	return &schema.Resource{
		Create: resourceAviApplicationPersistenceProfileCreate,
		Read:   ResourceAviApplicationPersistenceProfileRead,
		Update: resourceAviApplicationPersistenceProfileUpdate,
		Delete: resourceAviApplicationPersistenceProfileDelete,
		Schema: ResourceApplicationPersistenceProfileSchema(),
		Importer: &schema.ResourceImporter{
			State: ResourceApplicationPersistenceProfileImporter,
		},
	}
}

func ResourceApplicationPersistenceProfileImporter(d *schema.ResourceData, m interface{}) ([]*schema.ResourceData, error) {
	s := ResourceApplicationPersistenceProfileSchema()
	return ResourceImporter(d, m, "applicationpersistenceprofile", s)
}

func ResourceAviApplicationPersistenceProfileRead(d *schema.ResourceData, meta interface{}) error {
	s := ResourceApplicationPersistenceProfileSchema()
	err := APIRead(d, meta, "applicationpersistenceprofile", s)
	if err != nil {
		log.Printf("[ERROR] in reading object %v\n", err)
	}
	return err
}

func resourceAviApplicationPersistenceProfileCreate(d *schema.ResourceData, meta interface{}) error {
	s := ResourceApplicationPersistenceProfileSchema()
	err := APICreateOrUpdate(d, meta, "applicationpersistenceprofile", s)
	if err == nil {
		err = ResourceAviApplicationPersistenceProfileRead(d, meta)
	}
	return err
}

func resourceAviApplicationPersistenceProfileUpdate(d *schema.ResourceData, meta interface{}) error {
	s := ResourceApplicationPersistenceProfileSchema()
	var err error
	err = APICreateOrUpdate(d, meta, "applicationpersistenceprofile", s)
	if err == nil {
		err = ResourceAviApplicationPersistenceProfileRead(d, meta)
	}
	return err
}

func resourceAviApplicationPersistenceProfileDelete(d *schema.ResourceData, meta interface{}) error {
	objType := "applicationpersistenceprofile"
	client := meta.(*clients.AviClient)
	if APIDeleteSystemDefaultCheck(d) {
		return nil
	}
	uuid := d.Get("uuid").(string)
	if uuid != "" {
		path := "api/" + objType + "/" + uuid
		err := client.AviSession.Delete(path)
		if err != nil && !(strings.Contains(err.Error(), "404") || strings.Contains(err.Error(), "204") || strings.Contains(err.Error(), "403")) {
			log.Println("[INFO] resourceAviApplicationPersistenceProfileDelete not found")
			return err
		}
		d.SetId("")
	}
	return nil
}
