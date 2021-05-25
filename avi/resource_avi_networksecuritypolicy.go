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

func ResourceNetworkSecurityPolicySchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"cloud_config_cksum": {
			Type:     schema.TypeString,
			Optional: true,
			Computed: true,
		},
		"created_by": {
			Type:     schema.TypeString,
			Optional: true,
			Computed: true,
		},
		"description": {
			Type:     schema.TypeString,
			Optional: true,
			Computed: true,
		},
		"ip_reputation_db_ref": {
			Type:     schema.TypeString,
			Optional: true,
			Computed: true,
		},
		"markers": {
			Type:     schema.TypeList,
			Optional: true,
			Elem:     ResourceRoleFilterMatchLabelSchema(),
		},
		"name": {
			Type:     schema.TypeString,
			Optional: true,
			Computed: true,
		},
		"rules": {
			Type:     schema.TypeList,
			Optional: true,
			Elem:     ResourceNetworkSecurityRuleSchema(),
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

func resourceAviNetworkSecurityPolicy() *schema.Resource {
	return &schema.Resource{
		Create: resourceAviNetworkSecurityPolicyCreate,
		Read:   ResourceAviNetworkSecurityPolicyRead,
		Update: resourceAviNetworkSecurityPolicyUpdate,
		Delete: resourceAviNetworkSecurityPolicyDelete,
		Schema: ResourceNetworkSecurityPolicySchema(),
		Importer: &schema.ResourceImporter{
			State: ResourceNetworkSecurityPolicyImporter,
		},
	}
}

func ResourceNetworkSecurityPolicyImporter(d *schema.ResourceData, m interface{}) ([]*schema.ResourceData, error) {
	s := ResourceNetworkSecurityPolicySchema()
	return ResourceImporter(d, m, "networksecuritypolicy", s)
}

func ResourceAviNetworkSecurityPolicyRead(d *schema.ResourceData, meta interface{}) error {
	s := ResourceNetworkSecurityPolicySchema()
	err := APIRead(d, meta, "networksecuritypolicy", s)
	if err != nil {
		log.Printf("[ERROR] in reading object %v\n", err)
	}
	return err
}

func resourceAviNetworkSecurityPolicyCreate(d *schema.ResourceData, meta interface{}) error {
	s := ResourceNetworkSecurityPolicySchema()
	err := APICreateOrUpdate(d, meta, "networksecuritypolicy", s)
	if err == nil {
		err = ResourceAviNetworkSecurityPolicyRead(d, meta)
	}
	return err
}

func resourceAviNetworkSecurityPolicyUpdate(d *schema.ResourceData, meta interface{}) error {
	s := ResourceNetworkSecurityPolicySchema()
	var err error
	err = APICreateOrUpdate(d, meta, "networksecuritypolicy", s)
	if err == nil {
		err = ResourceAviNetworkSecurityPolicyRead(d, meta)
	}
	return err
}

func resourceAviNetworkSecurityPolicyDelete(d *schema.ResourceData, meta interface{}) error {
	objType := "networksecuritypolicy"
	client := meta.(*clients.AviClient)
	if APIDeleteSystemDefaultCheck(d) {
		return nil
	}
	uuid := d.Get("uuid").(string)
	if uuid != "" {
		path := "api/" + objType + "/" + uuid
		err := client.AviSession.Delete(path)
		if err != nil && !(strings.Contains(err.Error(), "404") || strings.Contains(err.Error(), "204") || strings.Contains(err.Error(), "403")) {
			log.Println("[INFO] resourceAviNetworkSecurityPolicyDelete not found")
			return err
		}
		d.SetId("")
	}
	return nil
}
