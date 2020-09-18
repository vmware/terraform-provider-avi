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

func ResourceSSOPolicySchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"authentication_policy": {
			Type:     schema.TypeSet,
			Optional: true,
			Computed: true,
			Elem:     ResourceAuthenticationPolicySchema(),
		},
		"authorization_policy": {
			Type:     schema.TypeSet,
			Optional: true,
			Computed: true,
			Elem:     ResourceAuthorizationPolicySchema(),
		},
		"labels": {
			Type:     schema.TypeList,
			Optional: true,
			Elem:     ResourceKeyValueSchema(),
		},
		"name": {
			Type:     schema.TypeString,
			Optional: true,
			Computed: true,
		},
		"tenant_ref": {
			Type:     schema.TypeString,
			Optional: true,
			Computed: true,
		},
		"type": {
			Type:     schema.TypeString,
			Optional: true,
			Default:  "SSO_TYPE_SAML",
		},
		"uuid": {
			Type:     schema.TypeString,
			Optional: true,
			Computed: true,
		},
	}
}

func resourceAviSSOPolicy() *schema.Resource {
	return &schema.Resource{
		Create: resourceAviSSOPolicyCreate,
		Read:   ResourceAviSSOPolicyRead,
		Update: resourceAviSSOPolicyUpdate,
		Delete: resourceAviSSOPolicyDelete,
		Schema: ResourceSSOPolicySchema(),
		Importer: &schema.ResourceImporter{
			State: ResourceSSOPolicyImporter,
		},
	}
}

func ResourceSSOPolicyImporter(d *schema.ResourceData, m interface{}) ([]*schema.ResourceData, error) {
	s := ResourceSSOPolicySchema()
	return ResourceImporter(d, m, "ssopolicy", s)
}

func ResourceAviSSOPolicyRead(d *schema.ResourceData, meta interface{}) error {
	s := ResourceSSOPolicySchema()
	err := ApiRead(d, meta, "ssopolicy", s)
	if err != nil {
		log.Printf("[ERROR] in reading object %v\n", err)
	}
	return err
}

func resourceAviSSOPolicyCreate(d *schema.ResourceData, meta interface{}) error {
	s := ResourceSSOPolicySchema()
	err := ApiCreateOrUpdate(d, meta, "ssopolicy", s)
	if err == nil {
		err = ResourceAviSSOPolicyRead(d, meta)
	}
	return err
}

func resourceAviSSOPolicyUpdate(d *schema.ResourceData, meta interface{}) error {
	s := ResourceSSOPolicySchema()
	var err error
	err = ApiCreateOrUpdate(d, meta, "ssopolicy", s)
	if err == nil {
		err = ResourceAviSSOPolicyRead(d, meta)
	}
	return err
}

func resourceAviSSOPolicyDelete(d *schema.ResourceData, meta interface{}) error {
	objType := "ssopolicy"
	client := meta.(*clients.AviClient)
	if ApiDeleteSystemDefaultCheck(d) {
		return nil
	}
	uuid := d.Get("uuid").(string)
	if uuid != "" {
		path := "api/" + objType + "/" + uuid
		err := client.AviSession.Delete(path)
		if err != nil && !(strings.Contains(err.Error(), "404") || strings.Contains(err.Error(), "204") || strings.Contains(err.Error(), "403")) {
			log.Println("[INFO] resourceAviSSOPolicyDelete not found")
			return err
		}
		d.SetId("")
	}
	return nil
}
