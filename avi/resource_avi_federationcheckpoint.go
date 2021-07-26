// Copyright 2019 VMware, Inc.
// SPDX-License-Identifier: Mozilla Public License 2.0

package avi

import (
	"log"
	"strings"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/vmware/alb-sdk/go/clients"
)

func ResourceFederationCheckpointSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"configpb_attributes": {
			Type:     schema.TypeSet,
			Optional: true,
			Computed: true,
			Elem:     ResourceConfigPbAttributesSchema(),
		},
		"date": {
			Type:     schema.TypeString,
			Optional: true,
			Computed: true,
		},
		"description": {
			Type:     schema.TypeString,
			Optional: true,
			Computed: true,
		},
		"is_federated": {
			Type:         schema.TypeString,
			Optional:     true,
			Default:      "true",
			ValidateFunc: validateBool,
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

func resourceAviFederationCheckpoint() *schema.Resource {
	return &schema.Resource{
		Create: resourceAviFederationCheckpointCreate,
		Read:   ResourceAviFederationCheckpointRead,
		Update: resourceAviFederationCheckpointUpdate,
		Delete: resourceAviFederationCheckpointDelete,
		Schema: ResourceFederationCheckpointSchema(),
		Importer: &schema.ResourceImporter{
			State: ResourceFederationCheckpointImporter,
		},
	}
}

func ResourceFederationCheckpointImporter(d *schema.ResourceData, m interface{}) ([]*schema.ResourceData, error) {
	s := ResourceFederationCheckpointSchema()
	return ResourceImporter(d, m, "federationcheckpoint", s)
}

func ResourceAviFederationCheckpointRead(d *schema.ResourceData, meta interface{}) error {
	s := ResourceFederationCheckpointSchema()
	err := APIRead(d, meta, "federationcheckpoint", s)
	if err != nil {
		log.Printf("[ERROR] in reading object %v\n", err)
	}
	return err
}

func resourceAviFederationCheckpointCreate(d *schema.ResourceData, meta interface{}) error {
	s := ResourceFederationCheckpointSchema()
	err := APICreateOrUpdate(d, meta, "federationcheckpoint", s)
	if err == nil {
		err = ResourceAviFederationCheckpointRead(d, meta)
	}
	return err
}

func resourceAviFederationCheckpointUpdate(d *schema.ResourceData, meta interface{}) error {
	s := ResourceFederationCheckpointSchema()
	var err error
	err = APICreateOrUpdate(d, meta, "federationcheckpoint", s)
	if err == nil {
		err = ResourceAviFederationCheckpointRead(d, meta)
	}
	return err
}

func resourceAviFederationCheckpointDelete(d *schema.ResourceData, meta interface{}) error {
	objType := "federationcheckpoint"
	client := meta.(*clients.AviClient)
	if APIDeleteSystemDefaultCheck(d) {
		return nil
	}
	uuid := d.Get("uuid").(string)
	if uuid != "" {
		path := "api/" + objType + "/" + uuid
		err := client.AviSession.Delete(path)
		if err != nil && !(strings.Contains(err.Error(), "404") || strings.Contains(err.Error(), "204") || strings.Contains(err.Error(), "403")) {
			log.Println("[INFO] resourceAviFederationCheckpointDelete not found")
			return err
		}
		d.SetId("")
	}
	return nil
}
