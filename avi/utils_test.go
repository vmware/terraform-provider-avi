package avi

import (
	"encoding/json"
	"reflect"
	"testing"

	"github.com/avinetworks/sdk/go/models"
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

// Testcase to test conversion of avidata to schema and schema to avidata
// apidata of pool object shoud be equal after conversion of schema to avidata
func TestAPIDataToSchemaViceVersa(t *testing.T) {
	s := ResourcePoolSchema()
	var apidata interface{}
	err := json.Unmarshal([]byte(apipooldata), &apidata)
	if err != nil {
		t.Errorf("ERROR: Error during JSON unmarshal.")
		t.Fail()
	}
	// Converting apidata into schema
	if schemadata, err := APIDataToSchema(apidata, nil, nil); err != nil {
		t.Errorf("ERROR: APIDataToSchema conversion failed with error %v", err)
	} else {
		// Converting schema into avidata
		data, _ := SchemaToAviData(schemadata, s)
		// Return false for not equals
		if !reflect.DeepEqual(apidata, data) {
			t.Fail()
		}
	}
}

// Testcase to test conversion of avidata to schema and schema to avidata on updated avidata
// apidata of pool object shoud not be equal after conversion of schema to avidata
func TestAPIDataToSchemaWithUpdatedApiData(t *testing.T) {
	s := ResourcePoolSchema()
	var apidata models.Pool
	err := json.Unmarshal([]byte(apipooldata), &apidata)
	if err != nil {
		t.Errorf("ERROR: Error during JSON unmarshal.")
		t.Fail()
	}
	// converting apidata into schema
	if schemadata, err := APIDataToSchema(apidata, nil, nil); err != nil {
		t.Errorf("ERROR: APIDataToSchema conversion failed with error %v", err)
	} else {
		// Updating the actual configuration of pool object
		enabled := false
		apidata.Enabled = &enabled
		// Converting schema into avidata.
		data, _ := SchemaToAviData(schemadata, s)
		// Return false for not equal,
		if reflect.DeepEqual(apidata, data) {
			t.Fail()
		}
	}
}
