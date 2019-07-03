package avi

import (
	"encoding/json"
	"github.com/avinetworks/sdk/go/models"
	"reflect"
	"testing"
)

var apipooldata = `{
        "name": "Pool-abc",
        "tenant_ref": "/api/tenant/?name=admin",
        "ignore_servers": false,
        "servers": [{
            "ratio": 1,
            "ip": {
                "type": "V4",
                "addr": "10.90.64.66"
            },
            "hostname": "10.90.64.66",
            "port": 8080,
            "enabled": true
        }],
        "fail_action": {
            "type": "FAIL_ACTION_CLOSE_CONN"
        }
    }`

func TestApiDataToSchemaViceVersa(t *testing.T) {
	s := ResourcePoolSchema()
	var apidata models.Pool
	err := json.Unmarshal([]byte(apipooldata), &apidata)
	if err != nil {
		t.Errorf("ERROR: Error during JSON unmarshal.")
		t.Fail()
	}
	if schemadata, err := ApiDataToSchema(apidata, nil, nil); err != nil {
		t.Errorf("ERROR: ApiDataToSchema conversion failed with error %v", err)
	} else {
		data, _ := SchemaToAviData(schemadata, s)
		if !reflect.DeepEqual(apidata, data) {
			t.Fail()
		}
	}
}

func TestApiDataToSchemaWithUpdatedApiData(t *testing.T) {
	s := ResourcePoolSchema()
	var apidata models.Pool
	err := json.Unmarshal([]byte(apipooldata), &apidata)
	if err != nil {
		t.Errorf("ERROR: Error during JSON unmarshal.")
		t.Fail()
	}
	if schemadata, err := ApiDataToSchema(apidata, nil, nil); err != nil {
		t.Errorf("ERROR: ApiDataToSchema conversion failed with error %v", err)
	} else {
		enabled := false
		apidata.Enabled = &enabled
		data, _ := SchemaToAviData(schemadata, s)
		if reflect.DeepEqual(apidata, data) {
			t.Fail()
		}
	}
}