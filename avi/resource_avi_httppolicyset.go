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

func ResourceHTTPPolicySetSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"cloud_config_cksum": &schema.Schema{
			Type:     schema.TypeString,
			Optional: true,
		},
		"created_by": &schema.Schema{
			Type:     schema.TypeString,
			Optional: true,
		},
		"description": &schema.Schema{
			Type:     schema.TypeString,
			Optional: true,
		},
		"http_request_policy": &schema.Schema{
			Type:     schema.TypeSet,
			Optional: true,
			Elem:     ResourceHTTPRequestPolicySchema(),
			Set: func(v interface{}) int {
				return 0
			},
		},
		"http_response_policy": &schema.Schema{
			Type:     schema.TypeSet,
			Optional: true,
			Elem:     ResourceHTTPResponsePolicySchema(),
			Set: func(v interface{}) int {
				return 0
			},
		},
		"http_security_policy": &schema.Schema{
			Type:     schema.TypeSet,
			Optional: true,
			Elem:     ResourceHTTPSecurityPolicySchema(),
			Set: func(v interface{}) int {
				return 0
			},
		},
		"is_internal_policy": &schema.Schema{
			Type:     schema.TypeBool,
			Optional: true,
			Default:  false,
		},
		"name": &schema.Schema{
			Type:     schema.TypeString,
			Required: true,
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

func resourceAviHTTPPolicySet() *schema.Resource {
	return &schema.Resource{
		Create: resourceAviHTTPPolicySetCreate,
		Read:   ResourceAviHTTPPolicySetRead,
		Update: resourceAviHTTPPolicySetUpdate,
		Delete: resourceAviHTTPPolicySetDelete,
		Schema: ResourceHTTPPolicySetSchema(),
		Importer: &schema.ResourceImporter{
			State: ResourceHTTPPolicySetImporter,
		},
	}
}

func ResourceHTTPPolicySetImporter(d *schema.ResourceData, m interface{}) ([]*schema.ResourceData, error) {
	s := ResourceHTTPPolicySetSchema()
	return ResourceImporter(d, m, "httppolicyset", s)
}

func ResourceAviHTTPPolicySetRead(d *schema.ResourceData, meta interface{}) error {
	s := ResourceHTTPPolicySetSchema()
	err := ApiRead(d, meta, "httppolicyset", s)
	if err != nil {
		log.Printf("[ERROR] in reading object %v\n", err)
	}
	return err
}

func resourceAviHTTPPolicySetCreate(d *schema.ResourceData, meta interface{}) error {
	s := ResourceHTTPPolicySetSchema()
	err := ApiCreateOrUpdate(d, meta, "httppolicyset", s)
	if err == nil {
		err = ResourceAviHTTPPolicySetRead(d, meta)
	}
	return err
}

func resourceAviHTTPPolicySetUpdate(d *schema.ResourceData, meta interface{}) error {
	s := ResourceHTTPPolicySetSchema()
	var err error

	err = ApiCreateOrUpdate(d, meta, "httppolicyset", s)
	if err == nil {
		err = ResourceAviHTTPPolicySetRead(d, meta)
	}
	return err
}

func resourceAviHTTPPolicySetDelete(d *schema.ResourceData, meta interface{}) error {
	objType := "httppolicyset"
	client := meta.(*clients.AviClient)
	uuid := d.Get("uuid").(string)
	if uuid != "" {
		path := "api/" + objType + "/" + uuid
		err := client.AviSession.Delete(path)
		if err != nil && !(strings.Contains(err.Error(), "404") || strings.Contains(err.Error(), "204")) {
			log.Println("[INFO] resourceAviHTTPPolicySetDelete not found")
			return err
		}
		d.SetId("")
	}
	return nil
}
