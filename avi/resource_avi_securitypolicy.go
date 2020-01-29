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

func ResourceSecurityPolicySchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"description": {
			Type:     schema.TypeString,
			Optional: true,
			Computed: true,
		},
		"dns_attacks": {
			Type:     schema.TypeSet,
			Optional: true,
			Computed: true,
			Elem:     ResourceDnsAttacksSchema(),
		},
		"dns_policy_index": {
			Type:     schema.TypeInt,
			Optional: true,
			Default:  0,
		},
		"name": {
			Type:     schema.TypeString,
			Required: true,
		},
		"network_security_policy_index": {
			Type:     schema.TypeInt,
			Optional: true,
			Default:  0,
		},
		"oper_mode": {
			Type:     schema.TypeString,
			Optional: true,
			Default:  "DETECTION",
		},
		"tcp_attacks": {
			Type:     schema.TypeSet,
			Optional: true,
			Computed: true,
			Elem:     ResourceTcpAttacksSchema(),
		},
		"tenant_ref": {
			Type:     schema.TypeString,
			Optional: true,
			Computed: true,
		},
		"udp_attacks": {
			Type:     schema.TypeSet,
			Optional: true,
			Computed: true,
			Elem:     ResourceUdpAttacksSchema(),
		},
		"uuid": {
			Type:     schema.TypeString,
			Optional: true,
			Computed: true,
		},
	}
}

func resourceAviSecurityPolicy() *schema.Resource {
	return &schema.Resource{
		Create: resourceAviSecurityPolicyCreate,
		Read:   ResourceAviSecurityPolicyRead,
		Update: resourceAviSecurityPolicyUpdate,
		Delete: resourceAviSecurityPolicyDelete,
		Schema: ResourceSecurityPolicySchema(),
		Importer: &schema.ResourceImporter{
			State: ResourceSecurityPolicyImporter,
		},
	}
}

func ResourceSecurityPolicyImporter(d *schema.ResourceData, m interface{}) ([]*schema.ResourceData, error) {
	s := ResourceSecurityPolicySchema()
	return ResourceImporter(d, m, "securitypolicy", s)
}

func ResourceAviSecurityPolicyRead(d *schema.ResourceData, meta interface{}) error {
	s := ResourceSecurityPolicySchema()
	err := ApiRead(d, meta, "securitypolicy", s)
	if err != nil {
		log.Printf("[ERROR] in reading object %v\n", err)
	}
	return err
}

func resourceAviSecurityPolicyCreate(d *schema.ResourceData, meta interface{}) error {
	s := ResourceSecurityPolicySchema()
	err := ApiCreateOrUpdate(d, meta, "securitypolicy", s)
	if err == nil {
		err = ResourceAviSecurityPolicyRead(d, meta)
	}
	return err
}

func resourceAviSecurityPolicyUpdate(d *schema.ResourceData, meta interface{}) error {
	s := ResourceSecurityPolicySchema()
	var err error
	err = ApiCreateOrUpdate(d, meta, "securitypolicy", s)
	if err == nil {
		err = ResourceAviSecurityPolicyRead(d, meta)
	}
	return err
}

func resourceAviSecurityPolicyDelete(d *schema.ResourceData, meta interface{}) error {
	objType := "securitypolicy"
	if ApiDeleteSystemDefaultCheck(d) {
		return nil
	}
	client := meta.(*clients.AviClient)
	uuid := d.Get("uuid").(string)
	if uuid != "" {
		path := "api/" + objType + "/" + uuid
		err := client.AviSession.Delete(path)
		if err != nil && !(strings.Contains(err.Error(), "404") || strings.Contains(err.Error(), "204") || strings.Contains(err.Error(), "403")) {
			log.Println("[INFO] resourceAviSecurityPolicyDelete not found")
			return err
		}
		d.SetId("")
	}
	return nil
}
