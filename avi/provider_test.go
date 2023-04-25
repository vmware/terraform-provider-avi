/***************************************************************************
 * ========================================================================
 * Copyright 2023 VMware, Inc. All rights reserved. VMware Confidential
 * ========================================================================
 */

/*
 * Copyright (c) 2017. Avi Networks.
 * Author: Gaurav Rastogi (grastogi@avinetworks.com)
 *
 */
package avi

import (
	"os"
	"strconv"
	"testing"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
	"github.com/vmware/alb-sdk/go/clients"
	"github.com/vmware/alb-sdk/go/session"
)

var testAccProviders map[string]*schema.Provider
var testAccProvider *schema.Provider

func init() {
	testAccProvider = Provider()
	testAccProviders = map[string]*schema.Provider{
		"avi": testAccProvider,
	}
}

func TestProvider(t *testing.T) {

	if err := Provider().InternalValidate(); err != nil {
		t.Fatalf("err: %s", err)
	}

	// Validating Schema for Provider

	var configs = map[string]interface{}{"avi_username": "",
		"avi_controller": "", "avi_password": "", "avi_tenant": "", "avi_api_timeout": 0}

	diags := Provider().Validate(
		&terraform.ResourceConfig{Config: configs})
	if diags.HasError() {
		t.Fatalf("err: %s", diags[0].Detail)
	}

	// Validating pool resource in avi provider and datasource for pool
	var poolconfigsData = map[string]interface{}{"name": ""}

	diags = Provider().ValidateDataSource("avi_pool",
		&terraform.ResourceConfig{Config: poolconfigsData})
	if diags.HasError() {
		t.Fatalf("err: %s", diags[0].Detail)
	}

	var poolconfigsRes = map[string]interface{}{"name": "", "uuid": "",
		"health_monitor_refs": make([]string, 0), "servers": make([]string, 0),
		"fail_action": make([]string, 0)}

	diags = Provider().ValidateResource("avi_pool",
		&terraform.ResourceConfig{Config: poolconfigsRes})
	if diags.HasError() {
		t.Fatalf("err: %s", diags[0].Detail)
	}

	// Validating pool resource in avi provider and datasource for pool

	var hmconfigsData = map[string]interface{}{"name": ""}

	diags = Provider().ValidateDataSource(
		"avi_healthmonitor", &terraform.ResourceConfig{Config: hmconfigsData})
	if diags.HasError() {
		t.Fatalf("err: %s", diags[0].Detail)
	}

	var hmconfigsRes = map[string]interface{}{"name": "", "uuid": "",
		"type": ""}

	diags = Provider().ValidateResource(
		"avi_healthmonitor", &terraform.ResourceConfig{Config: hmconfigsRes})
	if diags.HasError() {
		t.Fatalf("err: %s", diags[0].Detail)
	}
}

func testAccPreCheck(t *testing.T) {
	var timeout time.Duration
	if tm, err := strconv.Atoi(os.Getenv("AVI_API_TIMEOUT")); err == nil {
		timeout = time.Duration(tm) * time.Second
	}

	config := Credentials{
		Username:   os.Getenv("AVI_USERNAME"),
		Password:   os.Getenv("AVI_PASSWORD"),
		Controller: os.Getenv("AVI_CONTROLLER"),
		Tenant:     os.Getenv("AVI_TENANT"),
		Version:    os.Getenv("AVI_VERSION"),
		AuthToken:  os.Getenv("AVI_AUTHTOKEN"),
		Timeout:    timeout,
	}

	if config.Controller == "" {
		t.Fatalf("AVI_CONTROLLER must be set for acceptance test")
	}

	if config.Password == "" && config.AuthToken == "" {
		t.Fatalf("AVI_PASSWORD or AVI_AUTHTOKEN must be set for acceptance test")
	}

	if config.Username == "" {
		config.Username = "admin"
	}
	if config.Tenant == "" {
		config.Tenant = "admin"
	}
	if config.Version == "" {
		config.Version = "18.2.8"
	}

	errs := os.Setenv("AVI_SUPPRESS_SENSITIVE_FIELDS_DIFF", "true")
	if errs != nil {
		t.Fatalf("Unable to set env variable AVI_SUPPRESS_SENSITIVE_FIELDS_DIFF. Error: %s", errs)
	}

	_, err := clients.NewAviClient(
		config.Controller, config.Username,
		session.SetPassword(config.Password),
		session.SetTenant(config.Tenant),
		session.SetVersion(config.Version),
		session.SetAuthToken(config.AuthToken),
		session.SetInsecure, session.SetTimeout(config.Timeout))

	if err != nil {
		t.Fatalf("%+v", err)
	}
}
