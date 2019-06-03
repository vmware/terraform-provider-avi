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

func ResourcePoolGroupDeploymentPolicySchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"auto_disable_old_prod_pools": {
			Type:     schema.TypeBool,
			Optional: true,
			Default:  true,
		},
		"description": {
			Type:     schema.TypeString,
			Optional: true,
		},
		"evaluation_duration": {
			Type:     schema.TypeInt,
			Optional: true,
			Default:  300,
		},
		"name": {
			Type:     schema.TypeString,
			Required: true,
		},
		"rules": {
			Type:     schema.TypeList,
			Optional: true,
			Elem:     ResourcePGDeploymentRuleSchema(),
		},
		"scheme": {
			Type:     schema.TypeString,
			Optional: true,
			Default:  "BLUE_GREEN",
		},
		"target_test_traffic_ratio": {
			Type:     schema.TypeInt,
			Optional: true,
			Default:  100,
		},
		"tenant_ref": {
			Type:     schema.TypeString,
			Optional: true,
			Computed: true,
		},
		"test_traffic_ratio_rampup": {
			Type:     schema.TypeInt,
			Optional: true,
			Default:  100,
		},
		"uuid": {
			Type:     schema.TypeString,
			Optional: true,
			Computed: true,
		},
		"webhook_ref": {
			Type:     schema.TypeString,
			Optional: true,
			Computed: true,
		},
	}
}

func resourceAviPoolGroupDeploymentPolicy() *schema.Resource {
	return &schema.Resource{
		Create: resourceAviPoolGroupDeploymentPolicyCreate,
		Read:   ResourceAviPoolGroupDeploymentPolicyRead,
		Update: resourceAviPoolGroupDeploymentPolicyUpdate,
		Delete: resourceAviPoolGroupDeploymentPolicyDelete,
		Schema: ResourcePoolGroupDeploymentPolicySchema(),
		Importer: &schema.ResourceImporter{
			State: ResourcePoolGroupDeploymentPolicyImporter,
		},
	}
}

func ResourcePoolGroupDeploymentPolicyImporter(d *schema.ResourceData, m interface{}) ([]*schema.ResourceData, error) {
	s := ResourcePoolGroupDeploymentPolicySchema()
	return ResourceImporter(d, m, "poolgroupdeploymentpolicy", s)
}

func ResourceAviPoolGroupDeploymentPolicyRead(d *schema.ResourceData, meta interface{}) error {
	s := ResourcePoolGroupDeploymentPolicySchema()
	err := ApiRead(d, meta, "poolgroupdeploymentpolicy", s)
	if err != nil {
		log.Printf("[ERROR] in reading object %v\n", err)
	}
	return err
}

func resourceAviPoolGroupDeploymentPolicyCreate(d *schema.ResourceData, meta interface{}) error {
	s := ResourcePoolGroupDeploymentPolicySchema()
	err := ApiCreateOrUpdate(d, meta, "poolgroupdeploymentpolicy", s)
	if err == nil {
		err = ResourceAviPoolGroupDeploymentPolicyRead(d, meta)
	}
	return err
}

func resourceAviPoolGroupDeploymentPolicyUpdate(d *schema.ResourceData, meta interface{}) error {
	s := ResourcePoolGroupDeploymentPolicySchema()
	var err error
	err = ApiCreateOrUpdate(d, meta, "poolgroupdeploymentpolicy", s)
	if err == nil {
		err = ResourceAviPoolGroupDeploymentPolicyRead(d, meta)
	}
	return err
}

func resourceAviPoolGroupDeploymentPolicyDelete(d *schema.ResourceData, meta interface{}) error {
	objType := "poolgroupdeploymentpolicy"
	if ApiDeleteSystemDefaultCheck(d) {
		return nil
	}
	client := meta.(*clients.AviClient)
	uuid := d.Get("uuid").(string)
	if uuid != "" {
		path := "api/" + objType + "/" + uuid
		err := client.AviSession.Delete(path)
		if err != nil && !(strings.Contains(err.Error(), "404") || strings.Contains(err.Error(), "204") || strings.Contains(err.Error(), "403")) {
			log.Println("[INFO] resourceAviPoolGroupDeploymentPolicyDelete not found")
			return err
		}
		d.SetId("")
	}
	return nil
}
