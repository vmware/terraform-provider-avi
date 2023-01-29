data "avi_wafcrs" "wafcrs_2020_3" {
  name = var.csr_version
}

data "avi_wafprofile" "default_wafprofile" {
  name =  var.wafprofile #"System-WAF-Profile"
}

resource "avi_wafpolicy" "custom_waf_policy" {
  mode            = "WAF_MODE_ENFORCEMENT"
  name            = "custom_waf_policy"
  waf_crs_ref     = data.avi_wafcrs.wafcrs_2020_3.id
  waf_profile_ref = data.avi_wafprofile.default_wafprofile.uuid

  # Enabling/disabling CRS rule groups and individual rules is now simple with the
  # object model changes in 20.1.6, via the new "crs_overrides" list.

  crs_overrides {
    enable = true
    name   = "CRS_942_Application_Attack_SQLi"

    exclude_list {
      match_element = "ARGS:password"
      uri_path      = "/exception_path"
    }
  }
  crs_overrides {
    enable = true
    name   = "CRS_933_Application_Attack_PHP"
  }
  crs_overrides {
    enable = true
    name   = "CRS_932_Application_Attack_RCE"

    rule_overrides {
      enable  = true
      rule_id = "932115"
    }
    rule_overrides {
      enable  = true
      rule_id = "932120"
    }
    rule_overrides {
      enable  = true
      rule_id = "932140"
    }
    rule_overrides {
      enable  = true
      rule_id = "932190"
    }
  }

  # Adding custom rules directly through HCL is possible, but note that there is some complexity around escaping for HCL special
  # character sequences etc. There have been various bugs and regressions in Terraform around escape sequences over the years and
  # the syntax of ModSec signatures, especially more complex ones, often include multiple special characters in combination.
  #
  # We are looking at whether a model where an external file can be used to hold signatures would be useful to make this easier to manage.

  pre_crs_groups {
    enable = true
    index  = 0
    name   = "JSON Parsing Rules"

    rules {
      enable       = true
      index        = 0
      is_sensitive = false
      name         = "1 - Maximum JSON Parsing Depth"
      rule         = "SecRule ARGS_NAMES \"@rx (\\.[^\\.]*){2}\" \"id:10001,phase:2,t:none,block,msg:'Exceeded maximum JSON nested structure depth'\""
      rule_id      = "10001"
      tags         = []
    }
    rules {
      enable       = true
      index        = 1
      is_sensitive = false
      name         = "2 - Maximum value length"
      rule         = "SecRule ARGS \"@ge 100\" \"id:10002,phase:2,t:none,t:length,block,msg:'JSON value too long'\""
      rule_id      = "10002"
      tags         = []
    }
    rules {
      enable       = true
      index        = 2
      is_sensitive = false
      name         = "3a - Enable raw request body processing"
      rule         = "SecRule REQUEST_URI \"@unconditionalMatch\" \"id:10003,phase:1,pass,ctl:forceRequestBodyVariable=On\""
      rule_id      = "10003"
      tags         = []
    }
    rules {
      enable       = true
      index        = 3
      is_sensitive = false
      name         = "3b - Maximum request body size"
      rule         = "SecRule REQUEST_BODY_LENGTH \"@ge 150\" \"id:10004,phase:2,t:none,block,msg:'Request mody size too long'\""
      rule_id      = "10004"
      tags         = []
    }

    # Note that in the below example, %{MATCHED_VAR_NAME} requires the % character to be escaped (%%) - this would also apply to $
    # and other special characters/sequences in HCL. Terraform still does not properly generate escaped strings when exporting resource
    # state to HCL with "terraform state show" etc. (although quotes and backslashes are handled).

    rules {
      enable       = true
      index        = 4
      is_sensitive = false
      name         = "4a - Count lengths of JSON arrays"
      rule         = "SecRule ARGS_NAMES \"@unconditionalMatch\" \"id:10005,phase:2,pass,setvar:'TX.count_%%{MATCHED_VAR_NAME}=+1'\""
      rule_id      = "10005"
      tags         = []
    }
    rules {
      enable       = true
      index        = 5
      is_sensitive = false
      name         = "4b - Limit lengths of JSON arrays"
      rule         = "SecRule TX:/count_.*/  \"@gt 2\" \"id:10006,phase:2,t:none,block,msg:'JSON array has too many elements'\""
      rule_id      = "10006"
      tags         = []
    }
    rules {
      enable       = true
      index        = 6
      is_sensitive = false
      name         = "2 - Alternative max value length"
      rule         = "SecRule ARGS \"@rx ^.{101}\" \"id:10007,phase:2,t:none,block,msg:'JSON value too long'\""
      rule_id      = "10007"
      tags         = []
    }
    rules {
      enable       = true
      index        = 7
      is_sensitive = false
      name         = "22060 | Demo vulnerability"
      rule         = "SecRule REQUEST_HEADERS:Host \"192.168.1.100\" \"id:'1011', phase:1,t:none,nolog,pass,ctl:ruleRemoveById=960017\""
      rule_id      = "10008"
      tags         = []
    }
    rules {
      enable       = true
      index        = 8
      is_sensitive = false
      name         = "22061 | Log4j2-Demo 1 vulnerability"
      rule         = <<EOT
      SecRule REQUEST_LINE|ARGS|ARGS_NAMES|REQUEST_COOKIES|REQUEST_COOKIES_NAMES|REQUEST_BODY|REQUEST_HEADERS|XML:/*|XML://@* "@rx \$${(?:jndi|java):" "id:4022060, phase:2, block, t:none, t:lowercase, t:urlDecodeUni, msg:'CVE-2021-44228 log4j2 vulnerability', tag:'language-java', tag:'attack-multi', tag:'attack-rce', tag:'paranoia-level/1', tag:'CAPEC-152', tag:'CAPEC-242', tag:'CRS-group-402', ver:'AVI_CRS/2021_3', severity:'CRITICAL', multiMatch, setvar:'tx.anomaly_score_pl1=+%%{tx.critical_anomaly_score}', setvar:'tx.rce_score=+%%{tx.critical_anomaly_score}'" 
       EOT
      rule_id      = "10009"
      tags         = []

     }
       
   }

}