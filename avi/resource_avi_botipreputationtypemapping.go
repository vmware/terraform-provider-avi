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

func ResourceBotIPReputationTypeMappingSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"ip_reputation_mappings": {
			Type:     schema.TypeList,
			Optional: true,
			Elem:     ResourceIPReputationTypeMappingSchema(),
		},
		"name": {
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

func resourceAviBotIPReputationTypeMapping() *schema.Resource {
	return &schema.Resource{
		Create: resourceAviBotIPReputationTypeMappingCreate,
		Read:   ResourceAviBotIPReputationTypeMappingRead,
		Update: resourceAviBotIPReputationTypeMappingUpdate,
		Delete: resourceAviBotIPReputationTypeMappingDelete,
		Schema: ResourceBotIPReputationTypeMappingSchema(),
		Importer: &schema.ResourceImporter{
			State: ResourceBotIPReputationTypeMappingImporter,
		},
	}
}

func ResourceBotIPReputationTypeMappingImporter(d *schema.ResourceData, m interface{}) ([]*schema.ResourceData, error) {
	s := ResourceBotIPReputationTypeMappingSchema()
	return ResourceImporter(d, m, "botipreputationtypemapping", s)
}

func ResourceAviBotIPReputationTypeMappingRead(d *schema.ResourceData, meta interface{}) error {
	s := ResourceBotIPReputationTypeMappingSchema()
	err := APIRead(d, meta, "botipreputationtypemapping", s)
	if err != nil {
		log.Printf("[ERROR] in reading object %v\n", err)
	}
	return err
}

func resourceAviBotIPReputationTypeMappingCreate(d *schema.ResourceData, meta interface{}) error {
	s := ResourceBotIPReputationTypeMappingSchema()
	err := APICreateOrUpdate(d, meta, "botipreputationtypemapping", s)
	if err == nil {
		err = ResourceAviBotIPReputationTypeMappingRead(d, meta)
	}
	return err
}

func resourceAviBotIPReputationTypeMappingUpdate(d *schema.ResourceData, meta interface{}) error {
	s := ResourceBotIPReputationTypeMappingSchema()
	var err error
	err = APICreateOrUpdate(d, meta, "botipreputationtypemapping", s)
	if err == nil {
		err = ResourceAviBotIPReputationTypeMappingRead(d, meta)
	}
	return err
}

func resourceAviBotIPReputationTypeMappingDelete(d *schema.ResourceData, meta interface{}) error {
	objType := "botipreputationtypemapping"
	client := meta.(*clients.AviClient)
	if APIDeleteSystemDefaultCheck(d) {
		return nil
	}
	uuid := d.Get("uuid").(string)
	if uuid != "" {
		path := "api/" + objType + "/" + uuid
		err := client.AviSession.Delete(path)
		if err != nil && !(strings.Contains(err.Error(), "404") || strings.Contains(err.Error(), "204") || strings.Contains(err.Error(), "403")) {
			log.Println("[INFO] resourceAviBotIPReputationTypeMappingDelete not found")
			return err
		}
		d.SetId("")
	}
	return nil
}
