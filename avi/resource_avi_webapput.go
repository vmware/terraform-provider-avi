// Copyright 2019 VMware, Inc.
// SPDX-License-Identifier: Mozilla Public License 2.0

package avi

import (
	"log"
	"strings"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/vmware/alb-sdk/go/clients"
)

func ResourceWebappUTSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"configpb_attributes": {
			Type:     schema.TypeSet,
			Optional: true,
			Computed: true,
			Elem:     ResourceConfigPbAttributesSchema(),
		},
		"mandatory_test": {
			Type:     schema.TypeSet,
			Optional: true,
			Computed: true,
			Elem:     ResourceL1FMandatoryTestCaseSchema(),
		},
		"mandatory_tests": {
			Type:     schema.TypeList,
			Optional: true,
			Elem:     ResourceL1FMandatoryTestCaseSchema(),
		},
		"name": {
			Type:     schema.TypeString,
			Required: true,
		},
		"sensitive_test": {
			Type:     schema.TypeSet,
			Optional: true,
			Computed: true,
			Elem:     ResourceL1FSensitiveTestCaseSchema(),
		},
		"sensitive_tests": {
			Type:     schema.TypeList,
			Optional: true,
			Elem:     ResourceL1FSensitiveTestCaseSchema(),
		},
		"string_length_test": {
			Type:     schema.TypeSet,
			Optional: true,
			Computed: true,
			Elem:     ResourceL1StringLengthTestCaseSchema(),
		},
		"string_length_tests": {
			Type:     schema.TypeList,
			Optional: true,
			Elem:     ResourceL1StringLengthTestCaseSchema(),
		},
		"tenant_ref": {
			Type:     schema.TypeString,
			Optional: true,
			Computed: true,
		},
		"test_sensitive_string": {
			Type:             schema.TypeString,
			Optional:         true,
			Computed:         true,
			Sensitive:        true,
			DiffSuppressFunc: suppressSensitiveFieldDiffs,
		},
		"test_string": {
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

func resourceAviWebappUT() *schema.Resource {
	return &schema.Resource{
		Create: resourceAviWebappUTCreate,
		Read:   ResourceAviWebappUTRead,
		Update: resourceAviWebappUTUpdate,
		Delete: resourceAviWebappUTDelete,
		Schema: ResourceWebappUTSchema(),
		Importer: &schema.ResourceImporter{
			State: ResourceWebappUTImporter,
		},
	}
}

func ResourceWebappUTImporter(d *schema.ResourceData, m interface{}) ([]*schema.ResourceData, error) {
	s := ResourceWebappUTSchema()
	return ResourceImporter(d, m, "webapput", s)
}

func ResourceAviWebappUTRead(d *schema.ResourceData, meta interface{}) error {
	s := ResourceWebappUTSchema()
	err := APIRead(d, meta, "webapput", s)
	if err != nil {
		log.Printf("[ERROR] in reading object %v\n", err)
	}
	return err
}

func resourceAviWebappUTCreate(d *schema.ResourceData, meta interface{}) error {
	s := ResourceWebappUTSchema()
	err := APICreateOrUpdate(d, meta, "webapput", s)
	if err == nil {
		err = ResourceAviWebappUTRead(d, meta)
	}
	return err
}

func resourceAviWebappUTUpdate(d *schema.ResourceData, meta interface{}) error {
	s := ResourceWebappUTSchema()
	var err error
	err = APICreateOrUpdate(d, meta, "webapput", s)
	if err == nil {
		err = ResourceAviWebappUTRead(d, meta)
	}
	return err
}

func resourceAviWebappUTDelete(d *schema.ResourceData, meta interface{}) error {
	objType := "webapput"
	client := meta.(*clients.AviClient)
	if APIDeleteSystemDefaultCheck(d) {
		return nil
	}
	uuid := d.Get("uuid").(string)
	if uuid != "" {
		path := "api/" + objType + "/" + uuid
		err := client.AviSession.Delete(path)
		if err != nil && !(strings.Contains(err.Error(), "404") || strings.Contains(err.Error(), "204") || strings.Contains(err.Error(), "403")) {
			log.Println("[INFO] resourceAviWebappUTDelete not found")
			return err
		}
		d.SetId("")
	}
	return nil
}
