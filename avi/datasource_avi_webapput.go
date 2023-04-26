// Copyright 2019 VMware, Inc.
// SPDX-License-Identifier: Mozilla Public License 2.0

package avi

import "github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

func dataSourceAviWebappUT() *schema.Resource {
	return &schema.Resource{
		Read: ResourceAviWebappUTRead,
		Schema: map[string]*schema.Schema{
			"configpb_attributes": {
				Type:     schema.TypeSet,
				Computed: true,
				Elem:     ResourceConfigPbAttributesSchema(),
			},
			"mandatory_test": {
				Type:     schema.TypeSet,
				Computed: true,
				Elem:     ResourceL1FMandatoryTestCaseSchema(),
			},
			"mandatory_tests": {
				Type:     schema.TypeList,
				Computed: true,
				Elem:     ResourceL1FMandatoryTestCaseSchema(),
			},
			"name": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"sensitive_test": {
				Type:     schema.TypeSet,
				Computed: true,
				Elem:     ResourceL1FSensitiveTestCaseSchema(),
			},
			"sensitive_tests": {
				Type:     schema.TypeList,
				Computed: true,
				Elem:     ResourceL1FSensitiveTestCaseSchema(),
			},
			"string_length_test": {
				Type:     schema.TypeSet,
				Computed: true,
				Elem:     ResourceL1StringLengthTestCaseSchema(),
			},
			"string_length_tests": {
				Type:     schema.TypeList,
				Computed: true,
				Elem:     ResourceL1StringLengthTestCaseSchema(),
			},
			"tenant_ref": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"test_sensitive_string": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"test_string": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"uuid": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
		},
	}
}
