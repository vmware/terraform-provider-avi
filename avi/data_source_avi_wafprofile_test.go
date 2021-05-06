// Copyright 2019 VMware, Inc.
// SPDX-License-Identifier: Mozilla Public License 2.0

package avi

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAVIDataSourceWafProfileBasic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccAVIDSWafProfileConfig,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(
						"avi_wafprofile.testWafProfile", "name", "test-System-WAF-Profile-abc"),
				),
			},
		},
	})

}

const testAccAVIDSWafProfileConfig = `
data "avi_tenant" "default_tenant"{
    name= "admin"
}
resource "avi_wafprofile" "testWafProfile" {
	files {
	data = <<EOF
# Search engine crawlers and other bots

# site ripper
# http://www.softbytelabs.com/en/BlackWidow/
black widow
blackwidow
# crawler
# 2006
prowebwalker
# generic crawler
pymills-spider/
# SEO
# https://ahrefs.com/robot
AhrefsBot
# people database
# https://pipl.com/bot/
PiplBot
# advertising targeting
# https://www.grapeshot.com/crawler/
GrapeshotCrawler/2.0
grapeFX
# SEO
# http://www.searchmetrics.com/searchmetricsbot/
SearchmetricsBot
# SEO
# https://moz.com/help/guides/moz-procedures/what-is-rogerbot
rogerbot
# SEO
# http://www.majestic12.co.uk/projects/dsearch/mj12bot.php
MJ12bot
# news service
Owlin bot
# misbehaving spider
Lingewoud-550-Spyder
EOF
	name = "test-crawlers-user-agents.dat-abca"
}
files {
	data = <<EOF
<b>Warning</b>:
No row with the given identifier
open_basedir restriction in effect
eval()'d code</b> on line <b>
Cannot execute a blank command in
Fatal error</b>:  preg_replace
thrown in <b>
#0 {main}
Stack trace:
</b> on line <b>
() cannot be called statically
- not a Class::Method
::__toString() must not throw an exception
Access to undeclared static property:
An iterator cannot be used with foreach by reference
Array callback has to contain indices 0 and 1
Arrived at end of main loop which shouldn't happen
Attempt to destruct pending exception
Attempt to unset static property
Balloc() allocation exceeds list boundary
Balloc() failed to allocate memory
Base lambda function for closure not found
Call to a member function
Call to private 
Call to protected 
Call to undefined function
Call to undefined method
Can only throw objects
Cannot access empty property
Cannot access parent:: when current class scope has no parent
Cannot access parent:: when no class scope is active
Cannot access property started with '\\0'
Cannot access self:: when no class scope is active
Cannot access static:: when no class scope is active
Cannot access undefined property for object with overloaded property access
Cannot assign by reference to overloaded object
Cannot break/continue %d level%s
Cannot call abstract method 
Cannot call constructor
Cannot call forward_static_call() when no class scope is active
Cannot call non static method
Cannot call overloaded function for non-object
Cannot call private 
Cannot create references to/from string offsets
Cannot create references to/from string offsets nor overloaded objects
Cannot declare self-referencing constant 
Cannot destroy active lambda function
Cannot get arguments for 
Cannot increment/decrement overloaded objects nor string offsets
Cannot instantiate abstract class 
Cannot instantiate interface 
Cannot instantiate trait 
Cannot override final 
Cannot pass parameter 
Cannot redeclare 
Cannot redeclare class 
Cannot register a reverse output handler conflict outside of MINIT
Cannot register an output handler alias outside of MINIT
Cannot register an output handler conflict outside of MINIT
Cannot resume an already running generator
Cannot return string offsets by reference
Cannot set non exception as previous exception
Cannot unset string offsets
Cannot use [] for reading
Cannot use assign-op operators with overloaded objects nor string offsets
Cannot use object as array
Cannot use object of type 
Cannot use string offset as an array
Cannot use string offset as an object
Cannot yield from finally in a force-closed generator
Cannot yield string offsets by reference
Class entry requested for an object without PHP class
Class name must be a valid object or a string
Corrupted fcall_info provided to zend_call_function()
DCOM has been disabled by your administrator [com.allow_dcom=0]
DateFormat class not defined
DateTimeInterface can't be implemented by user classes
EXTREMELY fatal error: jmpbuf unrecoverable; terminating
EXTREMELY fatal error: jmpbuf unrecoverable; terminating.
EXTREMELY fatal error: longjmp returned control; terminating
Encoding: '*' may only be first arraySize value in list
Encoding: Attribute 
Encoding: Can't decode apache map, missing key
Encoding: Can't decode apache map, missing value
Encoding: Can't decode apache map, only Strings or Longs are allowd as keys
Encoding: Cannot find encoding
Encoding: Element 
Encoding: Error calling from_xml callback
Encoding: Error calling to_xml callback
Encoding: External reference 
Encoding: Internal Error
Encoding: Invalid timestamp 
Encoding: SoapVar has no 'enc_type' property
Encoding: Unresolved reference 
Encoding: Violation of encoding rules
Encoding: Violation of id and ref information items 
Encoding: object has no '
Encoding: object has no 'any' property
Encoding: string '
Error installing signal handler for 
Exception thrown without a stack frame
Exceptions must be valid objects derived from the Exception base class
Failed to clone SpoofChecker object
Failed to register IntlDateFormatter class
Failed to register MessageFormatter class
Failed to register NumberFormatter class
Failed to register ResourceBundle class
Field width %d is too long
First array member is not a valid class name or object
Function name must be a string
Illegal length modifier specified 
Illegal offset type
Input string is too long
Invalid RelaxNG Validation Context
Invalid Schema Validation Context
Invalid opcode
Invalid serialization data for DatePeriod object
Invalid serialization data for DateTime object
Invalid serialization data for DateTimeImmutable object
Maximum execution time of
Method name must be a string
Need to supply an object when throwing an exception
Nesting level too deep - recursive dependency?
NumberFormatter class not defined
Object does not support method calls
Only variables can be passed by reference
PDO: driver 
Parsing Schema: <restriction> or <extension> expected in complexContent
Parsing Schema: attribute
Parsing Schema: attribute has both 'ref' and 'type' attributes
Parsing Schema: attribute has both 'ref' attribute and subtype
Parsing Schema: attribute has both 'type' attribute and subtype
Parsing Schema: attribute has no 'name' nor 'ref' attributes
Parsing Schema: attributeGroup 
Parsing Schema: attributeGroup has both 'ref' attribute and subattribute
Parsing Schema: attributeGroup has no 'name' nor 'ref' attributes
Parsing Schema: can't import schema from 
Parsing Schema: complexType has no 'name' attribute
Parsing Schema: element has both 'default' and 'fixed' attributes
Parsing Schema: element has both 'itemType' attribute and subtype
Parsing Schema: element has both 'ref' and 'fixed' attributes
Parsing Schema: element has both 'ref' and 'nillable' attributes
Parsing Schema: element has both 'ref' and 'type' attributes
Parsing Schema: element has both 'ref' attribute and subtype
Parsing Schema: element has both 'type' attribute and subtype
Parsing Schema: element has no 'name' nor 'ref' attributes
Parsing Schema: expected <restriction> or <extension> in simpleContent
Parsing Schema: expected <restriction>, <list> or <union> in simpleType
Parsing Schema: extension has no 'base' attribute
Parsing Schema: group has both 'ref' attribute and subcontent
Parsing Schema: group has no 'name' nor 'ref' attributes
Parsing Schema: include has no 'schemaLocation' attribute
Parsing Schema: missing restriction value
Parsing Schema: redefine has no 'schemaLocation' attribute
Parsing Schema: restriction has no 'base' attribute
Parsing Schema: simpleType has no 'name' attribute
Parsing Schema: unexpected
Parsing Schema: unresolved element 'ref' attribute 
Parsing Schema: unresolved group 'ref' attribute 
Parsing WSDL: <binding> 
Parsing WSDL: <binding> has no name attribute
Parsing WSDL: <fault> with name 
Parsing WSDL: <message> has no name attribute
Parsing WSDL: <portType> 
Parsing WSDL: <portType> has no name attribute
Parsing WSDL: <service> 
Parsing WSDL: <service> has no name attribute
Parsing WSDL: Could not find any usable binding services in WSDL.
Parsing WSDL: Couldn't bind to service
Parsing WSDL: Couldn't find <definitions> in 
Parsing WSDL: Couldn't load from 
Parsing WSDL: Missing 'name' attribute for <binding>
Parsing WSDL: Missing 'name' attribute for <operation>
Parsing WSDL: Missing 'type' attribute for <binding>
Parsing WSDL: Missing <message> with name 
Parsing WSDL: Missing <portType> with name 
Parsing WSDL: Missing <portType>/<operation> with name 
Parsing WSDL: Missing message attribute for <header>
Parsing WSDL: Missing name for <fault> of 
Parsing WSDL: Missing name for <input> of 
Parsing WSDL: Missing name for <output> of 
Parsing WSDL: Missing part '%s' in <message>
Parsing WSDL: Missing part attribute for <header>
Parsing WSDL: No <binding> element with name 
Parsing WSDL: No address associated with <port>
Parsing WSDL: No binding associated with <port>
Parsing WSDL: No location associated with <port>
Parsing WSDL: No name associated with <part> 
Parsing WSDL: The fault message 
Parsing WSDL: Unexpected WSDL element <
Parsing WSDL: Unexpected extensibility element <
Parsing WSDL: Unknown encodingStyle '
Parsing WSDL: Unknown required WSDL extension '
Parsing WSDL: Unspecified encodingStyle
Possible integer overflow in memory allocation (
Second array member is not a valid method
Spoofchecker class not defined
String size overflow
The object to be iterated is in an invalid state: "
Trying to clone an uncloneable object
Trying to clone an uncloneable object of class 
Unable to call 
Unable to cast node to string
Undefined class constant '
Undefined constant '
Undefined offset for object of type 
Unexpected inconsistency in create_function()
Unknown SOAP version
Unknown typehint
Unsupported operand types
Using $this when not in object context
Wrong parameters for ErrorException(
Wrong parameters for Exception(
You MUST load PDO before loading any PDO drivers
[] operator not supported for strings
and must therefore be declared abstract or implement the remaining methods
namespace must not match the enclosing schema 'targetNamespace'
requires PDO API version 
EOF
	name = "test-php-errors.dat-abca"
}
files {
	data = <<EOF
__autoload
addcslashes
addslashes
apache_child_terminate
apache_get_modules
apache_get_version
apache_getenv
apache_lookup_uri
apache_note
apache_request_headers
apache_reset_timeout
apache_response_headers
apache_setenv
array_change_key_case
array_chunk
array_column
array_combine
array_count_values
array_diff
array_diff_assoc
array_diff_key
array_fill
array_fill_keys
array_flip
array_intersect
array_intersect_assoc
array_intersect_key
array_key_exists
array_keys
array_merge
array_merge_recursive
array_multisort
array_pad
array_pop
array_product
array_push
array_rand
array_replace
array_replace_recursive
array_reverse
array_search
array_shift
array_slice
array_splice
array_sum
array_unique
array_unshift
array_values
array_walk
array_walk_recursive
base_convert
bin2hex
bind_textdomain_codeset
bindtextdomain
blenc_encrypt
boolval
bzclose
bzcompress
bzerrno
bzerror
bzerrstr
bzflush
bzread
bzwrite
calcul_hmac
calculhmac
chdb_create
checkdnsrr
chgrp
chunk_split
class_alias
class_exists
class_implements
class_parents
class_uses
clearstatcache
cli_get_process_title
cli_set_process_title
com_create_guid
com_event_sink
com_get_active_object
com_load_typelib
com_message_pump
com_print_typeinfo
config_get_hash
connection_aborted
connection_status
convert_cyr_string
count_chars
crack_check
crack_closedict
crack_getlastmessage
crack_opendict
ctype_alnum
ctype_alpha
ctype_cntrl
ctype_digit
ctype_graph
ctype_lower
ctype_print
ctype_punct
ctype_space
ctype_upper
ctype_xdigit
curl_close
curl_copy_handle
curl_errno
curl_error
curl_escape
curl_getinfo
curl_multi_add_handle
curl_multi_close
curl_multi_exec
curl_multi_getcontent
curl_multi_info_read
curl_multi_init
curl_multi_remove_handle
curl_multi_select
curl_multi_setopt
curl_multi_strerror
curl_pause
curl_reset
curl_setopt
curl_setopt_array
curl_share_close
curl_share_init
curl_share_setopt
curl_strerror
curl_unescape
curl_version
date_add
date_create
date_create_from_format
date_create_immutable
date_create_immutable_from_format
date_date_set
date_default_timezone_get
date_default_timezone_set
date_diff
date_format
date_get_last_errors
date_interval_create_from_date_string
date_interval_format
date_isodate_set
date_modify
date_offset_get
date_parse
date_parse_from_format
date_sub
date_sun_info
date_sunrise
date_sunset
date_time_set
date_timestamp_get
date_timestamp_set
date_timezone_get
date_timezone_set
dcgettext
dcngettext
debug_print_backtrace
debug_zval_dump
decbin
dechex
define_syslog_variables
deg2rad
dgettext
disk_free_space
disk_total_space
dngettext
dns_check_record
dns_get_mx
dns_get_record
dom_import_simplexml
ereg_replace
eregi_replace
error_clear_last
error_get_last
expect_expectl
expect_popen
expm1
extension_loaded
ezmlm_hash
fastcgi_finish_request
fflush
fgetc
fgetcsv
fgetss
filter_has_var
filter_id
filter_input
filter_input_array
filter_list
filter_var
filter_var_array
finfo_close
fnmatch
foreach
forward_static_call
forward_static_call_array
fpassthru
fprintf
fputcsv
frenchtojd
fribidi_log2vis
fscanf
fseek
ftp_ssl_connect
ftruncate
func_get_arg
func_get_args
func_num_args
gc_collect_cycles
gc_disable
gc_enable
gc_enabled
gc_mem_caches
gd_info
get_browser
get_called_class
get_class
get_class_methods
get_class_vars
get_declared_classes
get_declared_interfaces
get_declared_traits
get_defined_constants
get_defined_functions
get_defined_vars
get_extension_funcs
get_headers
get_html_translation_table
get_include_path
get_included_files
get_loaded_extensions
get_magic_quotes_gpc
get_magic_quotes_runtime
get_object_vars
get_parent_class
get_required_files
get_resource_type
get_resources
getallheaders
gethostbyaddr
gethostbyname
gethostbynamel
gethostname
getimagesizefromstring
getmxrr
getopt
getprotobyname
getprotobynumber
getrandmax
getrusage
getservbyname
getservbyport
gettimeofday
gmmktime
gmstrftime
gopher_parsedir
gregoriantojd
gzclose
gzeof
gzgetc
gzgets
gzgetss
gzpassthru
gzputs
gzrewind
gzseek
gztell
hash_algos
hash_copy
hash_equals
hash_final
hash_hmac
hash_init
hash_pbkdf2
hash_update
hash_update_stream
header_remove
hebrevc
hexdec
highlight_string
http_build_query
http_response_code
iconv_get_encoding
iconv_mime_decode
iconv_mime_decode_headers
iconv_mime_encode
iconv_set_encoding
iconv_strlen
iconv_strpos
iconv_strrpos
iconv_substr
idn_to_ascii
idn_to_utf8
ignore_user_abort
image_type_to_extension
image_type_to_mime_type
import_request_variables
in_array
inet_ntop
inet_pton
ini_alter
ini_restore
interface_exists
intl_error_name
intl_get_error_code
intl_get_error_message
intl_is_failure
ip2long
iptcembed
iptcparse
is_subclass_of
is_uploaded_file
iterator_apply
iterator_count
iterator_to_array
jddayofweek
jdmonthname
jdtofrench
jdtogregorian
jdtojewish
jdtojulian
jdtounix
jewishtojd
jpeg2wbmp
json_last_error
json_last_error_msg
judy_type
judy_version
juliantojd
key_exists
krsort
lcg_value
lchgrp
lchown
libxml_clear_errors
libxml_disable_entity_loader
libxml_get_errors
libxml_get_last_error
libxml_set_external_entity_loader
libxml_set_streams_context
libxml_use_internal_errors
localeconv
long2ip
lzf_compress
lzf_decompress
lzf_optimized_for
magic_quotes_runtime
mb_check_encoding
mb_convert_case
mb_convert_encoding
mb_convert_kana
mb_convert_variables
mb_decode_mimeheader
mb_decode_numericentity
mb_detect_encoding
mb_detect_order
mb_encode_mimeheader
mb_encode_numericentity
mb_encoding_aliases
mb_ereg_search
mb_ereg_search_getpos
mb_ereg_search_getregs
mb_ereg_search_init
mb_ereg_search_pos
mb_ereg_search_regs
mb_ereg_search_setpos
mb_get_info
mb_http_input
mb_http_output
mb_internal_encoding
mb_language
mb_list_encodings
mb_output_handler
mb_preferred_mime_name
mb_regex_encoding
mb_regex_set_options
mb_send_mail
mb_split
mb_strcut
mb_strimwidth
mb_stripos
mb_stristr
mb_strlen
mb_strpos
mb_strrchr
mb_strrichr
mb_strripos
mb_strrpos
mb_strstr
mb_strtolower
mb_strtoupper
mb_strwidth
mb_substitute_character
mb_substr
mb_substr_count
mbereg_match
mbereg_replace
mbereg_search
mbereg_search_getpos
mbereg_search_getregs
mbereg_search_init
mbereg_search_pos
mbereg_search_regs
mbereg_search_setpos
mberegi
mberegi_replace
mbregex_encoding
mcrypt_cbc
mcrypt_cfb
mcrypt_create_iv
mcrypt_decrypt
mcrypt_ecb
mcrypt_enc_get_algorithms_name
mcrypt_enc_get_block_size
mcrypt_enc_get_iv_size
mcrypt_enc_get_key_size
mcrypt_enc_get_modes_name
mcrypt_enc_get_supported_key_sizes
mcrypt_enc_is_block_algorithm
mcrypt_enc_is_block_algorithm_mode
mcrypt_enc_is_block_mode
mcrypt_enc_self_test
mcrypt_encrypt
mcrypt_generic
mcrypt_generic_deinit
mcrypt_generic_end
mcrypt_generic_init
mcrypt_get_block_size
mcrypt_get_cipher_name
mcrypt_get_iv_size
mcrypt_get_key_size
mcrypt_list_algorithms
mcrypt_list_modes
mcrypt_module_close
mcrypt_module_get_algo_block_size
mcrypt_module_get_algo_key_size
mcrypt_module_get_supported_key_sizes
mcrypt_module_is_block_algorithm
mcrypt_module_is_block_algorithm_mode
mcrypt_module_is_block_mode
mcrypt_module_open
mcrypt_module_self_test
mcrypt_ofb
mdecrypt_generic
memcache_debug
memory_get_peak_usage
memory_get_usage
mhash_count
mhash_get_block_size
mhash_get_hash_name
mhash_keygen_s2k
mime_content_type
mktime
money_format
msg_get_queue
msg_queue_exists
msg_receive
msg_remove_queue
msg_send
msg_set_queue
msg_stat_queue
mssql_bind
mssql_close
mssql_connect
mssql_data_seek
mssql_execute
mssql_fetch_array
mssql_fetch_assoc
mssql_fetch_batch
mssql_fetch_field
mssql_fetch_object
mssql_fetch_row
mssql_field_length
mssql_field_name
mssql_field_seek
mssql_field_type
mssql_free_result
mssql_free_statement
mssql_get_last_message
mssql_guid_string
mssql_init
mssql_min_error_severity
mssql_min_message_severity
mssql_next_result
mssql_num_fields
mssql_num_rows
mssql_pconnect
mssql_query
mssql_result
mssql_rows_affected
mssql_select_db
mt_getrandmax
mt_rand
mt_srand
mysql_affected_rows
mysql_client_encoding
mysql_close
mysql_connect
mysql_create_db
mysql_createdb
mysql_data_seek
mysql_db_name
mysql_db_query
mysql_dbname
mysql_drop_db
mysql_dropdb
mysql_errno
mysql_error
mysql_escape_string
mysql_fetch_array
mysql_fetch_assoc
mysql_fetch_field
mysql_fetch_lengths
mysql_fetch_object
mysql_fetch_row
mysql_field_flags
mysql_field_len
mysql_field_name
mysql_field_seek
mysql_field_table
mysql_field_type
mysql_fieldflags
mysql_fieldlen
mysql_fieldname
mysql_fieldtable
mysql_fieldtype
mysql_free_result
mysql_freeresult
mysql_get_client_info
mysql_get_host_info
mysql_get_proto_info
mysql_get_server_info
mysql_info
mysql_insert_id
mysql_list_dbs
mysql_list_fields
mysql_list_processes
mysql_list_tables
mysql_listdbs
mysql_listfields
mysql_listtables
mysql_num_fields
mysql_num_rows
mysql_numfields
mysql_numrows
mysql_pconnect
mysql_ping
mysql_real_escape_string
mysql_result
mysql_select_db
mysql_selectdb
mysql_set_charset
mysql_stat
mysql_table_name
mysql_tablename
mysql_thread_id
mysql_unbuffered_query
mysqli_bind_param
mysqli_bind_result
mysqli_client_encoding
mysqli_connect
mysqli_disable_rpl_parse
mysqli_enable_reads_from_master
mysqli_enable_rpl_parse
mysqli_escape_string
mysqli_execute
mysqli_fetch
mysqli_get_cache_stats
mysqli_get_client_stats
mysqli_get_client_version
mysqli_get_links_stats
mysqli_get_metadata
mysqli_master_query
mysqli_param_count
mysqli_report
mysqli_rpl_parse_enabled
mysqli_rpl_probe
mysqli_send_long_data
mysqli_slave_query
mysqlnd_memcache_get_config
mysqlnd_memcache_set
mysqlnd_ms_dump_servers
mysqlnd_ms_fabric_select_global
mysqlnd_ms_fabric_select_shard
mysqlnd_ms_get_last_gtid
mysqlnd_ms_get_last_used_connection
mysqlnd_ms_get_stats
mysqlnd_ms_match_wild
mysqlnd_ms_query_is_select
mysqlnd_ms_set_qos
mysqlnd_ms_set_user_pick_server
mysqlnd_ms_xa_begin
mysqlnd_ms_xa_commit
mysqlnd_ms_xa_gc
mysqlnd_ms_xa_rollback
mysqlnd_qc_clear_cache
mysqlnd_qc_get_available_handlers
mysqlnd_qc_get_cache_info
mysqlnd_qc_get_core_stats
mysqlnd_qc_get_normalized_query_trace_log
mysqlnd_qc_get_query_trace_log
mysqlnd_qc_set_cache_condition
mysqlnd_qc_set_is_select
mysqlnd_qc_set_storage_handler
mysqlnd_qc_set_user_handlers
mysqlnd_uh_convert_to_mysqlnd
mysqlnd_uh_set_connection_proxy
mysqlnd_uh_set_statement_proxy
natcasesort
ngettext
nl2br
nl_langinfo
nsapi_request_headers
nsapi_response_headers
nsapi_virtual
nthmac
number_format
oauth_get_sbs
oauth_urlencode
ob_get_length
ob_get_level
ob_get_status
ob_gzhandler
ob_iconv_handler
ob_implicit_flush
ob_list_handlers
ob_tidyhandler
odbc_autocommit
odbc_binmode
odbc_close
odbc_close_all
odbc_columnprivileges
odbc_columns
odbc_commit
odbc_cursor
odbc_data_source
odbc_do
odbc_error
odbc_errormsg
odbc_fetch_array
odbc_fetch_into
odbc_fetch_object
odbc_fetch_row
odbc_field_len
odbc_field_name
odbc_field_num
odbc_field_precision
odbc_field_scale
odbc_field_type
odbc_foreignkeys
odbc_free_result
odbc_gettypeinfo
odbc_longreadlen
odbc_next_result
odbc_num_fields
odbc_num_rows
odbc_pconnect
odbc_prepare
odbc_primarykeys
odbc_procedurecolumns
odbc_procedures
odbc_rollback
odbc_setoption
odbc_specialcolumns
odbc_statistics
odbc_tableprivileges
odbc_tables
opcache_compile_file
opcache_get_configuration
opcache_get_status
opcache_invalidate
opcache_is_script_cached
opcache_reset
openssl_cipher_iv_length
openssl_csr_export
openssl_csr_export_to_file
openssl_csr_get_public_key
openssl_csr_get_subject
openssl_csr_new
openssl_csr_sign
openssl_decrypt
openssl_dh_compute_key
openssl_digest
openssl_encrypt
openssl_error_string
openssl_free_key
openssl_get_cert_locations
openssl_get_cipher_methods
openssl_get_md_methods
openssl_get_privatekey
openssl_get_publickey
openssl_open
openssl_pbkdf2
openssl_pkcs12_export
openssl_pkcs12_export_to_file
openssl_pkcs12_read
openssl_pkcs7_decrypt
openssl_pkcs7_encrypt
openssl_pkcs7_sign
openssl_pkcs7_verify
openssl_pkey_export
openssl_pkey_export_to_file
openssl_pkey_free
openssl_pkey_get_details
openssl_pkey_get_private
openssl_pkey_get_public
openssl_pkey_new
openssl_private_decrypt
openssl_private_encrypt
openssl_public_decrypt
openssl_public_encrypt
openssl_random_pseudo_bytes
openssl_seal
openssl_sign
openssl_spki_export
openssl_spki_export_challenge
openssl_spki_new
openssl_spki_verify
openssl_verify
openssl_x509_check_private_key
openssl_x509_checkpurpose
openssl_x509_export
openssl_x509_export_to_file
openssl_x509_fingerprint
openssl_x509_free
openssl_x509_parse
openssl_x509_read
output_add_rewrite_var
output_reset_rewrite_vars
override_function
parse_ini_string
parse_url
parsekit_compile_file
parsekit_compile_string
parsekit_func_arginfo
password_get_info
password_hash
password_needs_rehash
password_verify
pcntl_alarm
pcntl_errno
pcntl_get_last_error
pcntl_getpriority
pcntl_setpriority
pcntl_signal
pcntl_signal_dispatch
pcntl_sigprocmask
pcntl_sigtimedwait
pcntl_sigwaitinfo
pcntl_strerror
pcntl_wait
pcntl_waitpid
pcntl_wexitstatus
pcntl_wifexited
pcntl_wifsignaled
pcntl_wifstopped
pcntl_wstopsig
pcntl_wtermsig
pg_affected_rows
pg_cancel_query
pg_client_encoding
pg_close
pg_connect_poll
pg_connection_busy
pg_connection_reset
pg_connection_status
pg_consume_input
pg_convert
pg_copy_from
pg_copy_to
pg_dbname
pg_delete
pg_end_copy
pg_escape_bytea
pg_escape_identifier
pg_escape_literal
pg_escape_string
pg_fetch_all
pg_fetch_all_columns
pg_fetch_array
pg_fetch_assoc
pg_fetch_object
pg_fetch_result
pg_fetch_row
pg_field_is_null
pg_field_name
pg_field_num
pg_field_prtlen
pg_field_size
pg_field_table
pg_field_type
pg_field_type_oid
pg_flush
pg_free_result
pg_get_notify
pg_get_pid
pg_get_result
pg_host
pg_insert
pg_last_error
pg_last_notice
pg_last_oid
pg_lo_close
pg_lo_create
pg_lo_export
pg_lo_import
pg_lo_open
pg_lo_read
pg_lo_read_all
pg_lo_seek
pg_lo_tell
pg_lo_truncate
pg_lo_unlink
pg_lo_write
pg_meta_data
pg_num_fields
pg_num_rows
pg_options
pg_parameter_status
pg_pconnect
pg_ping
pg_port
pg_put_line
pg_query_params
pg_result_error
pg_result_error_field
pg_result_seek
pg_result_status
pg_select
pg_send_execute
pg_send_prepare
pg_send_query
pg_send_query_params
pg_set_client_encoding
pg_set_error_verbosity
pg_socket
pg_trace
pg_transaction_status
pg_tty
pg_unescape_bytea
pg_untrace
pg_update
pg_version
php_check_syntax
php_ini_loaded_file
php_ini_scanned_files
php_logo_guid
php_sapi_name
phpcredits
png2wbmp
posix_access
posix_ctermid
posix_errno
posix_get_last_error
posix_getgrgid
posix_getgrnam
posix_getgroups
posix_getpgid
posix_getpgrp
posix_getpid
posix_getppid
posix_getrlimit
posix_getsid
posix_initgroups
posix_isatty
posix_setegid
posix_seteuid
posix_setgid
posix_setpgid
posix_setrlimit
posix_setsid
posix_setuid
posix_strerror
posix_times
preg_filter
preg_grep
preg_last_error
preg_quote
property_exists
quoted_printable_decode
quoted_printable_encode
rad2deg
random_bytes
random_int
rar_wrapper_cache_stats
readline_add_history
readline_callback_handler_install
readline_callback_handler_remove
readline_callback_read_char
readline_clear_history
readline_completion_function
readline_info
readline_list_history
readline_on_new_line
readline_read_history
readline_redisplay
readline_write_history
realpath
realpath_cache_get
realpath_cache_size
recode_file
recode_string
restore_error_handler
restore_exception_handler
restore_include_path
rewinddir
rmdir
rpm_close
rpm_get_tag
rpm_is_valid
rpm_open
rpm_version
rrdc_disconnect
runkit_class_adopt
runkit_class_emancipate
runkit_constant_remove
runkit_function_remove
runkit_import
runkit_lint
runkit_lint_file
runkit_method_remove
runkit_return_value_used
runkit_sandbox_output_handler
runkit_superglobals
sem_acquire
sem_get
sem_release
sem_remove
session_abort
session_cache_expire
session_cache_limiter
session_commit
session_decode
session_destroy
session_encode
session_get_cookie_params
session_id
session_is_registered
session_module_name
session_name
session_pgsql_add_error
session_pgsql_get_error
session_pgsql_get_field
session_pgsql_reset
session_pgsql_set_field
session_pgsql_status
session_regenerate_id
session_register
session_register_shutdown
session_reset
session_save_path
session_set_cookie_params
session_status
session_unregister
session_unset
session_write_close
set_file_buffer
set_socket_blocking
set_time_limit
setcookie
setlocale
setproctitle
setrawcookie
setthreadtitle
shm_attach
shm_detach
shm_get_var
shm_has_var
shm_put_var
shm_remove
shm_remove_var
shmop_close
shmop_delete
shmop_open
shmop_read
shmop_size
shmop_write
simplexml_import_dom
socket_accept
socket_bind
socket_clear_error
socket_close
socket_cmsg_space
socket_create_listen
socket_create_pair
socket_get_option
socket_get_status
socket_getopt
socket_getpeername
socket_getsockname
socket_import_stream
socket_last_error
socket_listen
socket_read
socket_recv
socket_recvfrom
socket_recvmsg
socket_select
socket_send
socket_sendmsg
socket_sendto
socket_set_block
socket_set_blocking
socket_set_nonblock
socket_set_option
socket_set_timeout
socket_setopt
socket_shutdown
socket_strerror
socket_write
solr_get_version
spl_autoload
spl_autoload_call
spl_autoload_extensions
spl_autoload_functions
spl_autoload_register
spl_autoload_unregister
spl_classes
spl_object_hash
sql_regcase
sqlite_busy_timeout
sqlite_changes
sqlite_close
sqlite_column
sqlite_current
sqlite_error_string
sqlite_escape_string
sqlite_factory
sqlite_fetch_all
sqlite_fetch_array
sqlite_fetch_column_types
sqlite_fetch_object
sqlite_fetch_single
sqlite_fetch_string
sqlite_field_name
sqlite_has_more
sqlite_has_prev
sqlite_key
sqlite_last_error
sqlite_last_insert_rowid
sqlite_libencoding
sqlite_libversion
sqlite_next
sqlite_num_fields
sqlite_num_rows
sqlite_prev
sqlite_rewind
sqlite_seek
sqlite_udf_decode_binary
sqlite_udf_encode_binary
sqlite_valid
sqlsrv_begin_transaction
sqlsrv_cancel
sqlsrv_client_info
sqlsrv_close
sqlsrv_commit
sqlsrv_configure
sqlsrv_connect
sqlsrv_errors
sqlsrv_execute
sqlsrv_fetch
sqlsrv_fetch_array
sqlsrv_fetch_object
sqlsrv_field_metadata
sqlsrv_free_stmt
sqlsrv_get_config
sqlsrv_get_field
sqlsrv_has_rows
sqlsrv_next_result
sqlsrv_num_fields
sqlsrv_num_rows
sqlsrv_prepare
sqlsrv_query
sqlsrv_rollback
sqlsrv_rows_affected
sqlsrv_send_stream_data
sqlsrv_server_info
sscanf
ssdeep_fuzzy_compare
ssdeep_fuzzy_hash
ssdeep_fuzzy_hash_filename
stomp_connect_error
stomp_version
str_getcsv
str_ireplace
str_pad
str_repeat
str_replace
str_shuffle
str_split
str_word_count
strcasecmp
strchr
strcmp
strcspn
stream_bucket_append
stream_bucket_make_writeable
stream_bucket_new
stream_bucket_prepend
stream_context_get_default
stream_context_get_options
stream_context_get_params
stream_context_set_default
stream_context_set_option
stream_context_set_params
stream_copy_to_stream
stream_encoding
stream_filter_append
stream_filter_prepend
stream_filter_register
stream_filter_remove
stream_get_contents
stream_get_filters
stream_get_line
stream_get_meta_data
stream_get_transports
stream_get_wrappers
stream_is_local
stream_notification_callback
stream_register_wrapper
stream_resolve_include_path
stream_select
stream_set_blocking
stream_set_chunk_size
stream_set_read_buffer
stream_set_timeout
stream_set_write_buffer
stream_socket_accept
stream_socket_enable_crypto
stream_socket_get_name
stream_socket_pair
stream_socket_recvfrom
stream_socket_sendto
stream_socket_server
stream_socket_shutdown
stream_supports_lock
stream_wrapper_register
stream_wrapper_restore
stream_wrapper_unregister
strftime
strip_tags
stripos
stristr
strnatcasecmp
strnatcmp
strncasecmp
strncmp
strpbrk
strpos
strptime
strrchr
strripos
strrpos
strstr
strtok
strtolower
strtotime
strtoupper
strtr
strval
substr_compare
substr_count
substr_replace
sys_getloadavg
tcpwrap_check
time_nanosleep
time_sleep_until
timezone_abbreviations_list
timezone_identifiers_list
timezone_location_get
timezone_name_from_abbr
timezone_name_get
timezone_offset_get
timezone_open
timezone_transitions_get
timezone_version_get
token_get_all
token_name
trait_exists
trigger_error
ucwords
unixtojd
unregister_tick_function
use_soap_error_handler
user_error
utf8_decode
utf8_encode
var_export
version_compare
vfprintf
vprintf
vsprintf
win32_continue_service
win32_create_service
win32_delete_service
win32_get_last_control_message
win32_pause_service
win32_ps_list_procs
win32_ps_stat_mem
win32_ps_stat_proc
win32_query_service_status
win32_set_service_status
win32_start_service
win32_start_service_ctrl_dispatcher
win32_stop_service
xattr_get
xattr_list
xattr_remove
xattr_set
xattr_supported
xml_error_string
xml_get_current_byte_index
xml_get_current_column_number
xml_get_current_line_number
xml_get_error_code
xml_parse
xml_parse_into_struct
xml_parser_create
xml_parser_create_ns
xml_parser_free
xml_parser_get_option
xml_parser_set_option
xml_set_character_data_handler
xml_set_default_handler
xml_set_element_handler
xml_set_end_namespace_decl_handler
xml_set_external_entity_ref_handler
xml_set_notation_decl_handler
xml_set_object
xml_set_processing_instruction_handler
xml_set_start_namespace_decl_handler
xml_set_unparsed_entity_decl_handler
xmlrpc_decode
xmlrpc_decode_request
xmlrpc_encode
xmlrpc_encode_request
xmlrpc_get_type
xmlrpc_is_fault
xmlrpc_parse_method_descriptions
xmlrpc_server_add_introspection_data
xmlrpc_server_call_method
xmlrpc_server_create
xmlrpc_server_destroy
xmlrpc_server_register_introspection_callback
xmlrpc_server_register_method
xmlrpc_set_type
yaml_emit
yaml_emit_file
yaml_parse
yaml_parse_file
yaml_parse_url
zend_logo_guid
zend_thread_id
zend_version
zip_close
zip_entry_close
zip_entry_compressedsize
zip_entry_compressionmethod
zip_entry_filesize
zip_entry_name
zip_entry_open
zip_entry_read
zip_open
zip_read
zlib_encode
zlib_get_coding_type
EOF
	name = "test-php-function-names-933151.dat-abca"
}
files {
	data = <<EOF
# Apache
# (no slash; also guards against old.htaccess, old.htpasswd, etc.)
.htaccess
.htdigest
.htpasswd
# Version control
/.git/
/.gitignore
/.hg/
/.hgignore
/.svn/
# Wordpress
wp-config.php
wp-config.bak
wp-config.old
wp-config.temp
wp-config.tmp
wp-config.txt
# Symfony
/config/config.yml
/config/config_dev.yml
/config/config_prod.yml
/config/config_test.yml
/config/parameters.yml
/config/routing.yml
/config/security.yml
/config/services.yml
# Drupal
/sites/default/default.settings.php
/sites/default/settings.php
# Magento
/app/etc/local.xml
# Sublime Text
/sftp-config.json
# ASP.NET
/Web.config
# Node
/gruntfile.js
/npm-debug.log
# Composer
/composer.json
/composer.lock
/packages.json
# dotenv
/.env
EOF
	name = "test-restricted-files.dat-abca"
}
files {
	data = <<EOF
<jsp:
javax.servlet
.addheader
.createtextfile
.getfile
.loadfromfile
response.binarywrite
response.write
scripting.filesystemobject
server.createobject
server.execute
server.htmlencode
server.mappath
server.urlencode
vbscript.encode
wscript.network
wscript.shell
EOF
	name = "test-java-code-leakages.dat-abca"
}
files {
	data = <<EOF
powershell.exe
Add-BitsFile
Add-Computer
Add-Content
Add-History
Add-Member
Add-PSSnapin
Add-Type
Checkpoint-Computer
Clear-Content
Clear-EventLog
Clear-History
Clear-Item
Clear-ItemProperty
Clear-Variable
Compare-Object
Complete-BitsTransfer
Complete-Transaction
Connect-WSMan
ConvertFrom-CSV
ConvertFrom-SecureString
ConvertFrom-StringData
Convert-Path
ConvertTo-CSV
ConvertTo-Html
ConvertTo-SecureString
ConvertTo-XML
Copy-Item
Copy-ItemProperty
Debug-Process
Disable-ComputerRestore
Disable-PSBreakpoint
Disable-PSSessionConfiguration
Disable-WSManCredSSP
Disconnect-WSMan
Enable-ComputerRestore
Enable-PSBreakpoint
Enable-PSRemoting
Enable-PSSessionConfiguration
Enable-WSManCredSSP
Enter-PSSession
Exit-PSSession
Export-Alias
Export-Clixml
Export-Console
Export-Counter
Export-CSV
Export-FormatData
Export-ModuleMember
Export-PSSession
ForEach-Object
Format-Custom
Format-List
Format-Table
Format-Wide
Get-Acl
Get-Alias
Get-AppLockerFileInformation
Get-AppLockerPolicy
Get-AuthenticodeSignature
Get-BitsTransfer
Get-ChildItem
Get-Command
Get-ComputerRestorePoint
Get-Content
Get-Counter
Get-Credential
Get-Culture
Get-Event
Get-EventLog
Get-EventSubscriber
Get-ExecutionPolicy
Get-FormatData
Get-History
Get-Host
Get-HotFix
Get-Item
Get-ItemProperty
Get-Job
Get-Location
Get-Member
Get-Module
Get-PfxCertificate
Get-Process
Get-PSBreakpoint
Get-PSCallStack
Get-PSDrive
Get-PSProvider
Get-PSSession
Get-PSSessionConfiguration
Get-PSSnapin
Get-Random
Get-Service
Get-TraceSource
Get-Transaction
Get-TroubleshootingPack
Get-UICulture
Get-Unique
Get-Variable
Get-WinEvent
Get-WmiObject
Get-WSManCredSSP
Get-WSManInstance
Group-Object
Import-Alias
Import-Clixml
Import-Counter
Import-CSV
Import-LocalizedData
Import-Module
Import-PSSession
Invoke-Command
Invoke-Expression
Invoke-History
Invoke-Item
Invoke-TroubleshootingPack
Invoke-WmiMethod
Invoke-WSManAction
Join-Path
Limit-EventLog
Measure-Command
Measure-Object
Move-Item
Move-ItemProperty
New-Alias
New-AppLockerPolicy
New-Event
New-EventLog
New-Item
New-ItemProperty
New-Module
New-ModuleManifest
New-Object
New-PSDrive
New-PSSession
New-PSSessionOption
New-Service
New-TimeSpan
New-Variable
New-WebServiceProxy
New-WSManInstance
New-WSManSessionOption
Out-Default
Out-File
Out-GridView
Out-Host
Out-Null
Out-Printer
Out-String
Pop-Location
Push-Location
Read-Host
Receive-Job
Register-EngineEvent
Register-ObjectEvent
Register-PSSessionConfiguration
Register-WmiEvent
Remove-BitsTransfer
Remove-Computer
Remove-Event
Remove-EventLog
Remove-Item
Remove-ItemProperty
Remove-Job
Remove-Module
Remove-PSBreakpoint
Remove-PSDrive
Remove-PSSession
Remove-PSSnapin
Remove-Variable
Remove-WmiObject
Remove-WSManInstance
Rename-Item
Rename-ItemProperty
Reset-ComputerMachinePassword
Resolve-Path
Restart-Computer
Restart-Service
Restore-Computer
Resume-BitsTransfer
Resume-Service
Select-Object
Select-String
Select-XML
Send-MailMessage
Set-Acl
Set-Alias
Set-AppLockerPolicy
Set-AuthenticodeSignature
Set-BitsTransfer
Set-Content
Set-Date
Set-ExecutionPolicy
Set-Item
Set-ItemProperty
Set-Location
Set-PSBreakpoint
Set-PSDebug
Set-PSSessionConfiguration
Set-Service
Set-StrictMode
Set-TraceSource
Set-Variable
Set-WmiInstance
Set-WSManInstance
Set-WSManQuickConfig
Show-EventLog
Sort-Object
Split-Path
Start-BitsTransfer
Start-Job
Start-Process
Start-Service
Start-Sleep
Start-Transaction
Start-Transcript
Stop-Computer
Stop-Job
Stop-Process
Stop-Service
Stop-Transcript
Suspend-BitsTransfer
Suspend-Service
Tee-Object
Test-AppLockerPolicy
Test-ComputerSecureChannel
Test-Connection
Test-ModuleManifest
Test-Path
Test-WSMan
Trace-Command
Undo-Transaction
Unregister-Event
Unregister-PSSessionConfiguration
Update-FormatData
Update-List
Update-TypeData
Use-Transaction
Wait-Event
Wait-Job
Wait-Process
Where-Object
Write-Debug
Write-Error
Write-EventLog
Write-Host
Write-Output
Write-Progress
Write-Verbose
Write-Warning
-EncodedCommand
-ExecutionPolicy
-PSConsoleFile
EOF
	name = "test-windows-powershell-commands.dat-abca"
}
files {
	data = <<EOF
[java.lang.
class java.lang.
java.lang.NullPointerException
java.rmi.ServerException
at java.lang.
onclick="toggle('full exception chain stacktrace')"
at org.apache.catalina
at org.apache.coyote.
at org.apache.tomcat.
at org.apache.jasper.
EOF
	name = "test-java-errors.dat-abca"
}
files {
	data = <<EOF
MySqlClient.
Server message
SQL error
Oracle error
JET Database Engine
Procedure or function 
SQLite.Exception
[IBM][CLI Driver][DB2/6000]
the used select statements have different number of columns
org.postgresql.util.PSQLException
Access Database Engine
Incorrect syntax near
Syntax error in string in query expression
SQLiteException
' doesn't exist
CLI Driver
on MySQL result index
sybase
com.informix.jdbc
[MySQL][ODBC
Error
has occurred in the vicinity of:
Sintaxis incorrecta cerca de
MySQL server version for the right syntax to use
com.mysql.jdbc.exceptions
You have an error in your SQL syntax near
You have an error in your SQL syntax;
An illegal character has been found in the statement
pg_query() [:
supplied argument is not a valid MySQL
mssql_query()
mysql_fetch_array()
Exception
java.sql.SQLException
Column count doesn't match value count at row
Sybase message
 SQL Server
PostgreSQL query failed:
Dynamic SQL Error
System.Data.SQLite.SQLiteException
SQLite/JDBCDriver
Unclosed quotation mark before the character string
System.Data.SqlClient.
Unclosed quotation mark after the character string
System.Data.OleDb.OleDbException
[DM_QUERY_E_SYNTAX]
[SqlException
Unexpected end of command in statement
valid PostgreSQL result
pg_exec() [:
SQL Server
[SQLITE_ERROR]
Microsoft OLE DB Provider for ODBC Drivers
PostgreSQL
org.hsqldb.jdbc
ADODB.Field (0x800A0BCD)
SQL syntax
Exception 
System.Data.SqlClient.SqlException
Data type mismatch in criteria expression.
Driver
DB2 SQL error
Sybase message:
ORA-
[Microsoft][ODBC SQL Server Driver]
'80040e14'
Microsoft OLE DB Provider for SQL Server
 in query expression
Npgsql.
valid MySQL result
supplied argument is not a valid PostgreSQL result
db2_
Ingres SQLSTATE
Column count doesn't match
Warning
[Microsoft][ODBC Microsoft Access Driver]
[Macromedia][SQLServer JDBC Driver]
<b>Warning</b>: ibase_
Roadhouse.Cms.
DB2 SQL error:
EOF
	name = "test-sql-errors.dat-abca"
}
files {
	data = <<EOF
allow_call_time_pass_reference
allow_url_fopen
allow_url_include
always_populate_raw_post_data
arg_separator.input
arg_separator.output
asp_tags
assert.active
assert.bail
assert.callback
assert.quiet_eval
assert.warning
auto_append_file
auto_detect_line_endings
auto_globals_jit
auto_prepend_file
bcmath.scale
birdstep.max_links
browscap
cgi.fix_pathinfo
cgi.force_redirect
cgi.nph
cgi.redirect_status_env
cgi.rfc2616_headers
com.allow_dcom
com.autoregister_casesensitive
com.autoregister_typelib
com.autoregister_verbose
com.code_page
com.typelib_file
date.default_latitude
date.default_longitude
date.sunrise_zenith
date.sunset_zenith
date.timezone
dba.default_handler
default_charset
default_mimetype
default_socket_timeout
define_syslog_variables
disable_classes
disable_functions
display_errors
display_startup_errors
doc_root
docref_ext
docref_root
enable_dl
error_append_string
error_log
error_prepend_string
error_reporting
exif.decode_jis_intel
exif.decode_jis_motorola
exif.decode_unicode_intel
exif.decode_unicode_motorola
exif.encode_jis
exif.encode_unicode
expose_php
extension_dir
fastcgi.impersonate
fastcgi.logging
file_uploads
filter.default
filter.default_flags
gd.jpeg_ignore_warning
highlight.bg
highlight.comment
highlight.default
highlight.html
highlight.keyword
highlight.string
html_errors
ibase.allow_persistent
ibase.dateformat
ibase.default_charset
ibase.default_db
ibase.default_password
ibase.default_user
ibase.max_links
ibase.max_persistent
ibase.timeformat
ibase.timestampformat
iconv.input_encoding
iconv.internal_encoding
iconv.output_encoding
ignore_repeated_errors
ignore_repeated_source
ignore_user_abort
implicit_flush
include_path
intl.default_locale
intl.error_level
ldap.max_links
log_errors
log_errors_max_len
magic_quotes_gpc
magic_quotes_runtime
magic_quotes_sybase
mail.add_x_header
mail.force_extra_parameters
mail.log
max_execution_time
max_file_uploads
max_input_nesting_level
max_input_time
mbstring.detect_order
mbstring.encoding_translation
mbstring.func_overload
mbstring.http_input
mbstring.http_output
mbstring.http_output_conv_mimetype
mbstring.internal_encoding
mbstring.language
mbstring.script_encoding
mbstring.strict_detection
mbstring.substitute_character
mcrypt.algorithms_dir
mcrypt.modes_dir
memory_limit
mssql.allow_persistent
mssql.batchsize
mssql.charset
mssql.compatability_mode
mssql.connect_timeout
mssql.datetimeconvert
mssql.max_links
mssql.max_persistent
mssql.max_procs
mssql.min_error_severity
mssql.min_message_severity
mssql.secure_connection
mssql.textlimit
mssql.textsize
mssql.timeout
mysql.allow_local_infile
mysql.allow_persistent
mysql.cache_size
mysql.connect_timeout
mysql.default_host
mysql.default_password
mysql.default_port
mysql.default_socket
mysql.default_user
mysql.max_links
mysql.max_persistent
mysql.trace_mode
mysqli.allow_local_infile
mysqli.allow_persistent
mysqli.cache_size
mysqli.default_host
mysqli.default_port
mysqli.default_pw
mysqli.default_socket
mysqli.default_user
mysqli.max_links
mysqli.max_persistent
mysqli.reconnect
mysqlnd.collect_memory_statistics
mysqlnd.collect_statistics
mysqlnd.net_cmd_buffer_size
mysqlnd.net_read_buffer_size
oci8.connection_class
oci8.default_prefetch
oci8.events
oci8.max_persistent
oci8.old_oci_close_semantics
oci8.persistent_timeout
oci8.ping_interval
oci8.privileged_connect
oci8.statement_cache_size
odbc.allow_persistent
odbc.check_persistent
odbc.default_db
odbc.default_pw
odbc.default_user
odbc.defaultbinmode
odbc.defaultlrl
odbc.max_links
odbc.max_persistent
open_basedir
output_buffering
output_handler
pcre.backtrack_limit
pcre.recursion_limit
pdo_mysql.cache_size
pdo_mysql.default_socket
pdo_odbc.connection_pooling
pgsql.allow_persistent
pgsql.auto_reset_persistent
pgsql.ignore_notice
pgsql.log_notice
pgsql.max_links
pgsql.max_persistent
phar.cache_list
phar.readonly
phar.require_hash
post_max_size
realpath_cache_size
realpath_cache_ttl
register_argc_argv
register_globals
register_long_arrays
report_memleaks
report_zend_debug
request_order
safe_mode
safe_mode_allowed_env_vars
safe_mode_exec_dir
safe_mode_gid
safe_mode_include_dir
safe_mode_protected_env_vars
sendmail_from
sendmail_path
serialize_precision
session.auto_start
session.bug_compat_42
session.bug_compat_warn
session.cache_expire
session.cache_limiter
session.cookie_domain
session.cookie_httponly
session.cookie_lifetime
session.cookie_path
session.cookie_secure
session.entropy_file
session.entropy_length
session.gc_divisor
session.gc_maxlifetime
session.gc_probability
session.hash_bits_per_character
session.hash_function
session.name
session.referer_check
session.save_handler
session.save_path
session.serialize_handler
session.use_cookies
session.use_only_cookies
session.use_trans_sid
short_open_tag
soap.wsdl_cache_dir
soap.wsdl_cache_enabled
soap.wsdl_cache_limit
soap.wsdl_cache_ttl
sql.safe_mode
sqlite.assoc_case
sqlite3.extension_dir
sybct.allow_persistent
sybct.deadlock_retry_count
sybct.hostname
sybct.login_timeout
sybct.max_links
sybct.max_persistent
sybct.min_client_severity
sybct.min_server_severity
sybct.timeout
sysvshm.init_mem
tidy.clean_output
tidy.default_config
track_errors
unserialize_callback_func
upload_max_filesize
upload_tmp_dir
url_rewriter.tags
user_agent
user_dir
user_ini.cache_ttl
user_ini.filename
variables_order
xmlrpc_error_number
xmlrpc_errors
y2k_compliance
zlib.output_compression
zlib.output_compression_level
zlib.output_handler
EOF
	name = "test-php-config-directives.dat-abca"
}
files {
	data = <<EOF
$GLOBALS
$HTTP_COOKIE_VARS
$HTTP_ENV_VARS
$HTTP_GET_VARS
$HTTP_POST_FILES
$HTTP_POST_VARS
$HTTP_RAW_POST_DATA
$HTTP_REQUEST_VARS
$HTTP_SERVER_VARS
$_COOKIE
$_ENV
$_FILES
$_GET
$_POST
$_REQUEST
$_SERVER
$_SESSION
$argc
$argv
EOF
	name = "test-php-variables.dat-abca"
}
files {
	data = <<EOF
.aptitude/config
.bash_config
.bash_history
.bash_logout
.bashrc
.cache/notify-osd.log
.config/odesk/odesk team.conf
.cshrc
.gitconfig
.hplip/hplip.conf
.ksh_history
.my.cnf
.mysql_history
.nano_history
.profile
.psql_history
.sqlite_history
.ssh/authorized_keys
.ssh/config
.ssh/id_dsa
.ssh/id_dsa.pub
.ssh/identity
.ssh/identity.pub
.ssh/id_rsa
.ssh/id_rsa.pub
.ssh/known_hosts
.subversion/servers
.tconn/tconn.conf
.vidalia/vidalia.conf
.xauthority
.zshrc
etc/php.ini
bin/php.ini
etc/httpd/php.ini
usr/lib/php.ini
usr/lib/php/php.ini
usr/local/etc/php.ini
usr/local/lib/php.ini
usr/local/php/lib/php.ini
usr/local/php4/lib/php.ini
usr/local/php5/lib/php.ini
usr/local/apache/conf/php.ini
etc/php4.4/fcgi/php.ini
etc/php4/apache/php.ini
etc/php4/apache2/php.ini
etc/php5/apache/php.ini
etc/php5/apache2/php.ini
etc/php/php.ini
etc/php/php4/php.ini
etc/php/apache/php.ini
etc/php/apache2/php.ini
web/conf/php.ini
usr/local/zend/etc/php.ini
opt/xampp/etc/php.ini
var/local/www/conf/php.ini
etc/php/cgi/php.ini
etc/php4/cgi/php.ini
etc/php5/cgi/php.ini
home2/bin/stable/apache/php.ini
home/bin/stable/apache/php.ini
etc/httpd/conf.d/php.conf
php5/php.ini
php4/php.ini
php/php.ini
windows/php.ini
winnt/php.ini
apache/php/php.ini
xampp/apache/bin/php.ini
netserver/bin/stable/apache/php.ini
volumes/macintosh_hd1/usr/local/php/lib/php.ini
etc/mono/1.0/machine.config
etc/mono/2.0/machine.config
etc/mono/2.0/web.config
etc/mono/config
usr/local/cpanel/logs/stats_log
usr/local/cpanel/logs/access_log
usr/local/cpanel/logs/error_log
usr/local/cpanel/logs/license_log
usr/local/cpanel/logs/login_log
var/cpanel/cpanel.config
var/log/sw-cp-server/error_log
usr/local/psa/admin/logs/httpsd_access_log
usr/local/psa/admin/logs/panel.log
var/log/sso/sso.log
usr/local/psa/admin/conf/php.ini
etc/sw-cp-server/applications.d/plesk.conf
usr/local/psa/admin/conf/site_isolation_settings.ini
usr/local/sb/config
etc/sw-cp-server/applications.d/00-sso-cpserver.conf
etc/sso/sso_config.ini
etc/mysql/conf.d/old_passwords.cnf
var/log/mysql/mysql-bin.log
var/log/mysql/mysql-bin.index
var/log/mysql/data/mysql-bin.index
var/log/mysql.log
var/log/mysql.err
var/log/mysqlderror.log
var/log/mysql/mysql.log
var/log/mysql/mysql-slow.log
var/log/mysql-bin.index
var/log/data/mysql-bin.index
var/mysql.log
var/mysql-bin.index
var/data/mysql-bin.index
program files/mysql/mysql server 5.0/data/{host}.err
program files/mysql/mysql server 5.0/data/mysql.log
program files/mysql/mysql server 5.0/data/mysql.err
program files/mysql/mysql server 5.0/data/mysql-bin.log
program files/mysql/mysql server 5.0/data/mysql-bin.index
program files/mysql/data/{host}.err
program files/mysql/data/mysql.log
program files/mysql/data/mysql.err
program files/mysql/data/mysql-bin.log
program files/mysql/data/mysql-bin.index
mysql/data/{host}.err
mysql/data/mysql.log
mysql/data/mysql.err
mysql/data/mysql-bin.log
mysql/data/mysql-bin.index
usr/local/mysql/data/mysql.log
usr/local/mysql/data/mysql.err
usr/local/mysql/data/mysql-bin.log
usr/local/mysql/data/mysql-slow.log
usr/local/mysql/data/mysqlderror.log
usr/local/mysql/data/{host}.err
usr/local/mysql/data/mysql-bin.index
var/lib/mysql/my.cnf
etc/mysql/my.cnf
etc/my.cnf
program files/mysql/mysql server 5.0/my.ini
program files/mysql/mysql server 5.0/my.cnf
program files/mysql/my.ini
program files/mysql/my.cnf
mysql/my.ini
mysql/my.cnf
mysql/bin/my.ini
var/postgresql/log/postgresql.log
var/log/postgresql/postgresql.log
var/log/postgres/pg_backup.log
var/log/postgres/postgres.log
var/log/postgresql.log
var/log/pgsql/pgsql.log
var/log/postgresql/postgresql-8.1-main.log
var/log/postgresql/postgresql-8.3-main.log
var/log/postgresql/postgresql-8.4-main.log
var/log/postgresql/postgresql-9.0-main.log
var/log/postgresql/postgresql-9.1-main.log
var/log/pgsql8.log
var/log/postgresql/postgres.log
var/log/pgsql_log
var/log/postgresql/main.log
var/log/cron/var/log/postgres.log
usr/internet/pgsql/data/postmaster.log
usr/local/pgsql/data/postgresql.log
usr/local/pgsql/data/pg_log
postgresql/log/pgadmin.log
var/lib/pgsql/data/postgresql.conf
var/postgresql/db/postgresql.conf
var/nm2/postgresql.conf
usr/local/pgsql/data/postgresql.conf
usr/local/pgsql/data/pg_hba.conf
usr/internet/pgsql/data/pg_hba.conf
usr/local/pgsql/data/passwd
usr/local/pgsql/bin/pg_passwd
etc/postgresql/postgresql.conf
etc/postgresql/pg_hba.conf
home/postgres/data/postgresql.conf
home/postgres/data/pg_version
home/postgres/data/pg_ident.conf
home/postgres/data/pg_hba.conf
program files/postgresql/8.3/data/pg_hba.conf
program files/postgresql/8.3/data/pg_ident.conf
program files/postgresql/8.3/data/postgresql.conf
program files/postgresql/8.4/data/pg_hba.conf
program files/postgresql/8.4/data/pg_ident.conf
program files/postgresql/8.4/data/postgresql.conf
program files/postgresql/9.0/data/pg_hba.conf
program files/postgresql/9.0/data/pg_ident.conf
program files/postgresql/9.0/data/postgresql.conf
program files/postgresql/9.1/data/pg_hba.conf
program files/postgresql/9.1/data/pg_ident.conf
program files/postgresql/9.1/data/postgresql.conf
wamp/logs/access.log
wamp/logs/apache_error.log
wamp/logs/genquery.log
wamp/logs/mysql.log
wamp/logs/slowquery.log
wamp/bin/apache/apache2.2.22/logs/access.log
wamp/bin/apache/apache2.2.22/logs/error.log
wamp/bin/apache/apache2.2.21/logs/access.log
wamp/bin/apache/apache2.2.21/logs/error.log
wamp/bin/mysql/mysql5.5.24/data/mysql-bin.index
wamp/bin/mysql/mysql5.5.16/data/mysql-bin.index
wamp/bin/apache/apache2.2.21/conf/httpd.conf
wamp/bin/apache/apache2.2.22/conf/httpd.conf
wamp/bin/apache/apache2.2.21/wampserver.conf
wamp/bin/apache/apache2.2.22/wampserver.conf
wamp/bin/apache/apache2.2.22/conf/wampserver.conf
wamp/bin/mysql/mysql5.5.24/my.ini
wamp/bin/mysql/mysql5.5.24/wampserver.conf
wamp/bin/mysql/mysql5.5.16/my.ini
wamp/bin/mysql/mysql5.5.16/wampserver.conf
wamp/bin/php/php5.3.8/php.ini
wamp/bin/php/php5.4.3/php.ini
xampp/apache/logs/access.log
xampp/apache/logs/error.log
xampp/mysql/data/mysql-bin.index
xampp/mysql/data/mysql.err
xampp/mysql/data/{host}.err
xampp/sendmail/sendmail.log
xampp/apache/conf/httpd.conf
xampp/filezillaftp/filezilla server.xml
xampp/mercurymail/mercury.ini
xampp/php/php.ini
xampp/phpmyadmin/config.inc.php
xampp/sendmail/sendmail.ini
xampp/webalizer/webalizer.conf
opt/lampp/etc/httpd.conf
xampp/htdocs/aca.txt
xampp/htdocs/admin.php
xampp/htdocs/leer.txt
usr/local/apache/logs/audit_log
usr/local/apache2/logs/audit_log
logs/security_debug_log
logs/security_log
usr/local/apache/conf/modsec.conf
usr/local/apache2/conf/modsec.conf
winnt/system32/logfiles/msftpsvc
winnt/system32/logfiles/msftpsvc1
winnt/system32/logfiles/msftpsvc2
windows/system32/logfiles/msftpsvc
windows/system32/logfiles/msftpsvc1
windows/system32/logfiles/msftpsvc2
etc/logrotate.d/proftpd
www/logs/proftpd.system.log
var/log/proftpd
var/log/proftpd/xferlog.legacy
var/log/proftpd.access_log
var/log/proftpd.xferlog
etc/pam.d/proftpd
etc/proftp.conf
etc/protpd/proftpd.conf
etc/vhcs2/proftpd/proftpd.conf
etc/proftpd/modules.conf
var/log/vsftpd.log
etc/vsftpd.chroot_list
etc/logrotate.d/vsftpd.log
etc/vsftpd/vsftpd.conf
etc/vsftpd.conf
etc/chrootusers
var/log/xferlog
var/adm/log/xferlog
etc/wu-ftpd/ftpaccess
etc/wu-ftpd/ftphosts
etc/wu-ftpd/ftpusers
var/log/pure-ftpd/pure-ftpd.log
logs/pure-ftpd.log
var/log/pureftpd.log
usr/sbin/pure-config.pl
usr/etc/pure-ftpd.conf
etc/pure-ftpd/pure-ftpd.conf
usr/local/etc/pure-ftpd.conf
usr/local/etc/pureftpd.pdb
usr/local/pureftpd/etc/pureftpd.pdb
usr/local/pureftpd/sbin/pure-config.pl
usr/local/pureftpd/etc/pure-ftpd.conf
etc/pure-ftpd.conf
etc/pure-ftpd/pure-ftpd.pdb
etc/pureftpd.pdb
etc/pureftpd.passwd
etc/pure-ftpd/pureftpd.pdb
usr/ports/ftp/pure-ftpd/pure-ftpd.conf
usr/ports/ftp/pure-ftpd/pureftpd.pdb
usr/ports/ftp/pure-ftpd/pureftpd.passwd
usr/ports/net/pure-ftpd/pure-ftpd.conf
usr/ports/net/pure-ftpd/pureftpd.pdb
usr/ports/net/pure-ftpd/pureftpd.passwd
usr/pkgsrc/net/pureftpd/pure-ftpd.conf
usr/pkgsrc/net/pureftpd/pureftpd.pdb
usr/pkgsrc/net/pureftpd/pureftpd.passwd
usr/ports/contrib/pure-ftpd/pure-ftpd.conf
usr/ports/contrib/pure-ftpd/pureftpd.pdb
usr/ports/contrib/pure-ftpd/pureftpd.passwd
var/log/muddleftpd
usr/sbin/mudlogd
etc/muddleftpd/mudlog
etc/muddleftpd.com
etc/muddleftpd/mudlogd.conf
etc/muddleftpd/muddleftpd.conf
var/log/muddleftpd.conf
usr/sbin/mudpasswd
etc/muddleftpd/muddleftpd.passwd
etc/muddleftpd/passwd
var/log/ftp-proxy/ftp-proxy.log
var/log/ftp-proxy
var/log/ftplog
etc/logrotate.d/ftp
etc/ftpchroot
etc/ftphosts
etc/ftpusers
var/log/exim_mainlog
var/log/exim/mainlog
var/log/maillog
var/log/exim_paniclog
var/log/exim/paniclog
var/log/exim/rejectlog
var/log/exim_rejectlog
winnt/system32/logfiles/smtpsvc
winnt/system32/logfiles/smtpsvc1
winnt/system32/logfiles/smtpsvc2
winnt/system32/logfiles/smtpsvc3
winnt/system32/logfiles/smtpsvc4
winnt/system32/logfiles/smtpsvc5
windows/system32/logfiles/smtpsvc
windows/system32/logfiles/smtpsvc1
windows/system32/logfiles/smtpsvc2
windows/system32/logfiles/smtpsvc3
windows/system32/logfiles/smtpsvc4
windows/system32/logfiles/smtpsvc5
etc/osxhttpd/osxhttpd.conf
system/library/webobjects/adaptors/apache2.2/apache.conf
library/webserver/documents/.htaccess
etc/apache2/sites-available/default
etc/apache2/sites-available/default-ssl
etc/apache2/sites-enabled/000-default
etc/apache2/sites-enabled/default
etc/apache2/apache2.conf
etc/apache2/ports.conf
usr/local/etc/apache/httpd.conf
usr/pkg/etc/httpd/httpd.conf
usr/pkg/etc/httpd/httpd-default.conf
usr/pkg/etc/httpd/httpd-vhosts.conf
etc/httpd/mod_php.conf
etc/httpd/extra/httpd-ssl.conf
etc/rc.d/rc.httpd
usr/local/apache/conf/httpd.conf.default
usr/local/apache/conf/access.conf
usr/local/apache22/conf/httpd.conf
usr/local/apache22/httpd.conf
usr/local/etc/apache22/conf/httpd.conf
usr/local/apps/apache22/conf/httpd.conf
etc/apache22/conf/httpd.conf
etc/apache22/httpd.conf
opt/apache22/conf/httpd.conf
usr/local/etc/apache2/vhosts.conf
usr/local/apache/conf/vhosts.conf
usr/local/apache2/conf/vhosts.conf
usr/local/apache/conf/vhosts-custom.conf
usr/local/apache2/conf/vhosts-custom.conf
etc/apache/default-server.conf
etc/apache2/default-server.conf
usr/local/apache2/conf/extra/httpd-ssl.conf
usr/local/apache2/conf/ssl.conf
etc/httpd/conf.d
usr/local/etc/apache22/httpd.conf
usr/local/etc/apache2/httpd.conf
etc/apache2/httpd2.conf
etc/apache2/ssl-global.conf
etc/apache2/vhosts.d/00_default_vhost.conf
apache/conf/httpd.conf
etc/apache/httpd.conf
etc/httpd/conf
http/httpd.conf
usr/local/apache1.3/conf/httpd.conf
usr/local/etc/httpd/conf
var/apache/conf/httpd.conf
var/www/conf
www/apache/conf/httpd.conf
www/conf/httpd.conf
etc/init.d
etc/apache/access.conf
etc/rc.conf
www/logs/freebsddiary-error.log
www/logs/freebsddiary-access_log
library/webserver/documents/index.html
library/webserver/documents/index.htm
library/webserver/documents/default.html
library/webserver/documents/default.htm
library/webserver/documents/index.php
library/webserver/documents/default.php
var/log/webmin/miniserv.log
usr/local/etc/webmin/miniserv.conf
etc/webmin/miniserv.conf
usr/local/etc/webmin/miniserv.users
etc/webmin/miniserv.users
winnt/system32/logfiles/w3svc/inetsvn1.log
winnt/system32/logfiles/w3svc1/inetsvn1.log
winnt/system32/logfiles/w3svc2/inetsvn1.log
winnt/system32/logfiles/w3svc3/inetsvn1.log
windows/system32/logfiles/w3svc/inetsvn1.log
windows/system32/logfiles/w3svc1/inetsvn1.log
windows/system32/logfiles/w3svc2/inetsvn1.log
windows/system32/logfiles/w3svc3/inetsvn1.log
var/log/httpd/access_log
var/log/httpd/error_log
apache/logs/error.log
apache/logs/access.log
apache2/logs/error.log
apache2/logs/access.log
logs/error.log
logs/access.log
etc/httpd/logs/access_log
etc/httpd/logs/access.log
etc/httpd/logs/error_log
etc/httpd/logs/error.log
usr/local/apache/logs/access_log
usr/local/apache/logs/access.log
usr/local/apache/logs/error_log
usr/local/apache/logs/error.log
usr/local/apache2/logs/access_log
usr/local/apache2/logs/access.log
usr/local/apache2/logs/error_log
usr/local/apache2/logs/error.log
var/www/logs/access_log
var/www/logs/access.log
var/www/logs/error_log
var/www/logs/error.log
var/log/httpd/access.log
var/log/httpd/error.log
var/log/apache/access_log
var/log/apache/access.log
var/log/apache/error_log
var/log/apache/error.log
var/log/apache2/access_log
var/log/apache2/access.log
var/log/apache2/error_log
var/log/apache2/error.log
var/log/access_log
var/log/access.log
var/log/error_log
var/log/error.log
opt/lampp/logs/access_log
opt/lampp/logs/error_log
opt/xampp/logs/access_log
opt/xampp/logs/error_log
opt/lampp/logs/access.log
opt/lampp/logs/error.log
opt/xampp/logs/access.log
opt/xampp/logs/error.log
program files/apache group/apache/logs/access.log
program files/apache group/apache/logs/error.log
program files/apache software foundation/apache2.2/logs/error.log
program files/apache software foundation/apache2.2/logs/access.log
opt/apache/apache.conf
opt/apache/conf/apache.conf
opt/apache2/apache.conf
opt/apache2/conf/apache.conf
opt/httpd/apache.conf
opt/httpd/conf/apache.conf
etc/httpd/apache.conf
etc/apache2/apache.conf
etc/httpd/conf/apache.conf
usr/local/apache/apache.conf
usr/local/apache/conf/apache.conf
usr/local/apache2/apache.conf
usr/local/apache2/conf/apache.conf
usr/local/php/apache.conf.php
usr/local/php4/apache.conf.php
usr/local/php5/apache.conf.php
usr/local/php/apache.conf
usr/local/php4/apache.conf
usr/local/php5/apache.conf
private/etc/httpd/apache.conf
opt/apache/apache2.conf
opt/apache/conf/apache2.conf
opt/apache2/apache2.conf
opt/apache2/conf/apache2.conf
opt/httpd/apache2.conf
opt/httpd/conf/apache2.conf
etc/httpd/apache2.conf
etc/httpd/conf/apache2.conf
usr/local/apache/apache2.conf
usr/local/apache/conf/apache2.conf
usr/local/apache2/apache2.conf
usr/local/apache2/conf/apache2.conf
usr/local/php/apache2.conf.php
usr/local/php4/apache2.conf.php
usr/local/php5/apache2.conf.php
usr/local/php/apache2.conf
usr/local/php4/apache2.conf
usr/local/php5/apache2.conf
private/etc/httpd/apache2.conf
usr/local/apache/conf/httpd.conf
usr/local/apache2/conf/httpd.conf
etc/httpd/conf/httpd.conf
etc/apache/apache.conf
etc/apache/conf/httpd.conf
etc/apache2/httpd.conf
usr/apache2/conf/httpd.conf
usr/apache/conf/httpd.conf
usr/local/etc/apache/conf/httpd.conf
usr/local/apache/httpd.conf
usr/local/apache2/httpd.conf
usr/local/httpd/conf/httpd.conf
usr/local/etc/apache2/conf/httpd.conf
usr/local/etc/httpd/conf/httpd.conf
usr/local/apps/apache2/conf/httpd.conf
usr/local/apps/apache/conf/httpd.conf
usr/local/php/httpd.conf.php
usr/local/php4/httpd.conf.php
usr/local/php5/httpd.conf.php
usr/local/php/httpd.conf
usr/local/php4/httpd.conf
usr/local/php5/httpd.conf
etc/apache2/conf/httpd.conf
etc/http/conf/httpd.conf
etc/httpd/httpd.conf
etc/http/httpd.conf
etc/httpd.conf
opt/apache/conf/httpd.conf
opt/apache2/conf/httpd.conf
var/www/conf/httpd.conf
private/etc/httpd/httpd.conf
private/etc/httpd/httpd.conf.default
etc/apache2/vhosts.d/default_vhost.include
etc/apache2/conf.d/charset
etc/apache2/conf.d/security
etc/apache2/envvars
etc/apache2/mods-available/autoindex.conf
etc/apache2/mods-available/deflate.conf
etc/apache2/mods-available/dir.conf
etc/apache2/mods-available/mem_cache.conf
etc/apache2/mods-available/mime.conf
etc/apache2/mods-available/proxy.conf
etc/apache2/mods-available/setenvif.conf
etc/apache2/mods-available/ssl.conf
etc/apache2/mods-enabled/alias.conf
etc/apache2/mods-enabled/deflate.conf
etc/apache2/mods-enabled/dir.conf
etc/apache2/mods-enabled/mime.conf
etc/apache2/mods-enabled/negotiation.conf
etc/apache2/mods-enabled/php5.conf
etc/apache2/mods-enabled/status.conf
program files/apache group/apache/conf/httpd.conf
program files/apache group/apache2/conf/httpd.conf
program files/xampp/apache/conf/apache.conf
program files/xampp/apache/conf/apache2.conf
program files/xampp/apache/conf/httpd.conf
program files/apache group/apache/apache.conf
program files/apache group/apache/conf/apache.conf
program files/apache group/apache2/conf/apache.conf
program files/apache group/apache/apache2.conf
program files/apache group/apache/conf/apache2.conf
program files/apache group/apache2/conf/apache2.conf
program files/apache software foundation/apache2.2/conf/httpd.conf
volumes/macintosh_hd1/opt/httpd/conf/httpd.conf
volumes/macintosh_hd1/opt/apache/conf/httpd.conf
volumes/macintosh_hd1/opt/apache2/conf/httpd.conf
volumes/macintosh_hd1/usr/local/php/httpd.conf.php
volumes/macintosh_hd1/usr/local/php4/httpd.conf.php
volumes/macintosh_hd1/usr/local/php5/httpd.conf.php
volumes/webbackup/opt/apache2/conf/httpd.conf
volumes/webbackup/private/etc/httpd/httpd.conf
volumes/webbackup/private/etc/httpd/httpd.conf.default
usr/local/etc/apache/vhosts.conf
usr/local/jakarta/tomcat/conf/jakarta.conf
usr/local/jakarta/tomcat/conf/server.xml
usr/local/jakarta/tomcat/conf/context.xml
usr/local/jakarta/tomcat/conf/workers.properties
usr/local/jakarta/tomcat/conf/logging.properties
usr/local/jakarta/dist/tomcat/conf/jakarta.conf
usr/local/jakarta/dist/tomcat/conf/server.xml
usr/local/jakarta/dist/tomcat/conf/context.xml
usr/local/jakarta/dist/tomcat/conf/workers.properties
usr/local/jakarta/dist/tomcat/conf/logging.properties
usr/share/tomcat6/conf/server.xml
usr/share/tomcat6/conf/context.xml
usr/share/tomcat6/conf/workers.properties
usr/share/tomcat6/conf/logging.properties
var/log/tomcat6/catalina.out
var/cpanel/tomcat.options
usr/local/jakarta/tomcat/logs/catalina.out
usr/local/jakarta/tomcat/logs/catalina.err
opt/tomcat/logs/catalina.out
opt/tomcat/logs/catalina.err
usr/share/logs/catalina.out
usr/share/logs/catalina.err
usr/share/tomcat/logs/catalina.out
usr/share/tomcat/logs/catalina.err
usr/share/tomcat6/logs/catalina.out
usr/share/tomcat6/logs/catalina.err
usr/local/apache/logs/mod_jk.log
usr/local/jakarta/tomcat/logs/mod_jk.log
usr/local/jakarta/dist/tomcat/logs/mod_jk.log
opt/[jboss]/server/default/conf/jboss-minimal.xml
opt/[jboss]/server/default/conf/jboss-service.xml
opt/[jboss]/server/default/conf/jndi.properties
opt/[jboss]/server/default/conf/log4j.xml
opt/[jboss]/server/default/conf/login-config.xml
opt/[jboss]/server/default/conf/standardjaws.xml
opt/[jboss]/server/default/conf/standardjboss.xml
opt/[jboss]/server/default/conf/server.log.properties
opt/[jboss]/server/default/deploy/jboss-logging.xml
usr/local/[jboss]/server/default/conf/jboss-minimal.xml
usr/local/[jboss]/server/default/conf/jboss-service.xml
usr/local/[jboss]/server/default/conf/jndi.properties
usr/local/[jboss]/server/default/conf/log4j.xml
usr/local/[jboss]/server/default/conf/login-config.xml
usr/local/[jboss]/server/default/conf/standardjaws.xml
usr/local/[jboss]/server/default/conf/standardjboss.xml
usr/local/[jboss]/server/default/conf/server.log.properties
usr/local/[jboss]/server/default/deploy/jboss-logging.xml
private/tmp/[jboss]/server/default/conf/jboss-minimal.xml
private/tmp/[jboss]/server/default/conf/jboss-service.xml
private/tmp/[jboss]/server/default/conf/jndi.properties
private/tmp/[jboss]/server/default/conf/log4j.xml
private/tmp/[jboss]/server/default/conf/login-config.xml
private/tmp/[jboss]/server/default/conf/standardjaws.xml
private/tmp/[jboss]/server/default/conf/standardjboss.xml
private/tmp/[jboss]/server/default/conf/server.log.properties
private/tmp/[jboss]/server/default/deploy/jboss-logging.xml
tmp/[jboss]/server/default/conf/jboss-minimal.xml
tmp/[jboss]/server/default/conf/jboss-service.xml
tmp/[jboss]/server/default/conf/jndi.properties
tmp/[jboss]/server/default/conf/log4j.xml
tmp/[jboss]/server/default/conf/login-config.xml
tmp/[jboss]/server/default/conf/standardjaws.xml
tmp/[jboss]/server/default/conf/standardjboss.xml
tmp/[jboss]/server/default/conf/server.log.properties
tmp/[jboss]/server/default/deploy/jboss-logging.xml
program files/[jboss]/server/default/conf/jboss-minimal.xml
program files/[jboss]/server/default/conf/jboss-service.xml
program files/[jboss]/server/default/conf/jndi.properties
program files/[jboss]/server/default/conf/log4j.xml
program files/[jboss]/server/default/conf/login-config.xml
program files/[jboss]/server/default/conf/standardjaws.xml
program files/[jboss]/server/default/conf/standardjboss.xml
program files/[jboss]/server/default/conf/server.log.properties
program files/[jboss]/server/default/deploy/jboss-logging.xml
[jboss]/server/default/conf/jboss-minimal.xml
[jboss]/server/default/conf/jboss-service.xml
[jboss]/server/default/conf/jndi.properties
[jboss]/server/default/conf/log4j.xml
[jboss]/server/default/conf/login-config.xml
[jboss]/server/default/conf/standardjaws.xml
[jboss]/server/default/conf/standardjboss.xml
[jboss]/server/default/conf/server.log.properties
[jboss]/server/default/deploy/jboss-logging.xml
opt/[jboss]/server/default/log/server.log
opt/[jboss]/server/default/log/boot.log
usr/local/[jboss]/server/default/log/server.log
usr/local/[jboss]/server/default/log/boot.log
private/tmp/[jboss]/server/default/log/server.log
private/tmp/[jboss]/server/default/log/boot.log
tmp/[jboss]/server/default/log/server.log
tmp/[jboss]/server/default/log/boot.log
program files/[jboss]/server/default/log/server.log
program files/[jboss]/server/default/log/boot.log
[jboss]/server/default/log/server.log
[jboss]/server/default/log/boot.log
var/log/lighttpd.error.log
var/log/lighttpd.access.log
var/lighttpd.log
var/logs/access.log
var/log/lighttpd/
var/log/lighttpd/error.log
var/log/lighttpd/access.www.log
var/log/lighttpd/error.www.log
var/log/lighttpd/access.log
usr/local/apache2/logs/lighttpd.error.log
usr/local/apache2/logs/lighttpd.log
usr/local/apache/logs/lighttpd.error.log
usr/local/apache/logs/lighttpd.log
usr/local/lighttpd/log/lighttpd.error.log
usr/local/lighttpd/log/access.log
var/log/lighttpd/{domain}/access.log
var/log/lighttpd/{domain}/error.log
usr/home/user/var/log/lighttpd.error.log
usr/home/user/var/log/apache.log
home/user/lighttpd/lighttpd.conf
usr/home/user/lighttpd/lighttpd.conf
etc/lighttpd/lighthttpd.conf
usr/local/etc/lighttpd.conf
usr/local/lighttpd/conf/lighttpd.conf
usr/local/etc/lighttpd.conf.new
var/www/.lighttpdpassword
var/log/nginx/access_log
var/log/nginx/error_log
var/log/nginx/access.log
var/log/nginx/error.log
var/log/nginx.access_log
var/log/nginx.error_log
logs/access_log
logs/error_log
etc/nginx/nginx.conf
usr/local/etc/nginx/nginx.conf
usr/local/nginx/conf/nginx.conf
usr/local/zeus/web/global.cfg
usr/local/zeus/web/log/errors
opt/lsws/conf/httpd_conf.xml
usr/local/lsws/conf/httpd_conf.xml
opt/lsws/logs/error.log
opt/lsws/logs/access.log
usr/local/lsws/logs/error.log
usr/local/logs/access.log
usr/local/samba/lib/log.user
usr/local/logs/samba.log
var/log/samba/log.smbd
var/log/samba/log.nmbd
var/log/samba.log
var/log/samba.log1
var/log/samba.log2
var/log/log.smb
etc/samba/netlogon
etc/smbpasswd
etc/smb.conf
etc/samba/dhcp.conf
etc/samba/smb.conf
etc/samba/samba.conf
etc/samba/smb.conf.user
etc/samba/smbpasswd
etc/samba/smbusers
etc/samba/private/smbpasswd
usr/local/etc/smb.conf
usr/local/samba/lib/smb.conf.user
etc/dhcp3/dhclient.conf
etc/dhcp3/dhcpd.conf
etc/dhcp/dhclient.conf
program files/vidalia bundle/polipo/polipo.conf
etc/tor/tor-tsocks.conf
etc/stunnel/stunnel.conf
etc/tsocks.conf
etc/tinyproxy/tinyproxy.conf
etc/miredo-server.conf
etc/miredo.conf
etc/miredo/miredo-server.conf
etc/miredo/miredo.conf
etc/wicd/dhclient.conf.template.default
etc/wicd/manager-settings.conf
etc/wicd/wired-settings.conf
etc/wicd/wireless-settings.conf
var/log/ipfw.log
var/log/ipfw
var/log/ipfw/ipfw.log
var/log/ipfw.today
etc/ipfw.rules
etc/ipfw.conf
etc/firewall.rules
winnt/system32/logfiles/firewall/pfirewall.log
winnt/system32/logfiles/firewall/pfirewall.log.old
windows/system32/logfiles/firewall/pfirewall.log
windows/system32/logfiles/firewall/pfirewall.log.old
etc/clamav/clamd.conf
etc/clamav/freshclam.conf
etc/x11/xorg.conf
etc/x11/xorg.conf-vesa
etc/x11/xorg.conf-vmware
etc/x11/xorg.conf.beforevmwaretoolsinstall
etc/x11/xorg.conf.orig
etc/bluetooth/input.conf
etc/bluetooth/main.conf
etc/bluetooth/network.conf
etc/bluetooth/rfcomm.conf
proc/self/environ
proc/self/mounts
proc/self/stat
proc/self/status
proc/self/cmdline
proc/self/fd/0
proc/self/fd/1
proc/self/fd/2
proc/self/fd/3
proc/self/fd/4
proc/self/fd/5
proc/self/fd/6
proc/self/fd/7
proc/self/fd/8
proc/self/fd/9
proc/self/fd/10
proc/self/fd/11
proc/self/fd/12
proc/self/fd/13
proc/self/fd/14
proc/self/fd/15
proc/version
proc/devices
proc/cpuinfo
proc/meminfo
proc/net/tcp
proc/net/udp
etc/bash_completion.d/debconf
root/.bash_logout
root/.bash_history
root/.bash_config
root/.bashrc
etc/bash.bashrc
var/adm/syslog
var/adm/sulog
var/adm/utmp
var/adm/utmpx
var/adm/wtmp
var/adm/wtmpx
var/adm/lastlog/username
usr/spool/lp/log
var/adm/lp/lpd-errs
usr/lib/cron/log
var/adm/loginlog
var/adm/pacct
var/adm/dtmp
var/adm/acct/sum/loginlog
var/adm/x0msgs
var/adm/crash/vmcore
var/adm/crash/unix
etc/newsyslog.conf
var/adm/qacct
var/adm/ras/errlog
var/adm/ras/bootlog
var/adm/cron/log
etc/utmp
etc/security/lastlog
etc/security/failedlogin
usr/spool/mqueue/syslog
var/adm/messages
var/adm/aculogs
var/adm/aculog
var/adm/vold.log
var/adm/log/asppp.log
var/log/poplog
var/log/authlog
var/lp/logs/lpsched
var/lp/logs/lpnet
var/lp/logs/requests
var/cron/log
var/saf/_log
var/saf/port/log
var/log/news.all
var/log/news/news.all
var/log/news/news.crit
var/log/news/news.err
var/log/news/news.notice
var/log/news/suck.err
var/log/news/suck.notice
var/log/messages
var/log/messages.1
var/log/user.log
var/log/user.log.1
var/log/auth.log
var/log/pm-powersave.log
var/log/xorg.0.log
var/log/daemon.log
var/log/daemon.log.1
var/log/kern.log
var/log/kern.log.1
var/log/mail.err
var/log/mail.info
var/log/mail.warn
var/log/ufw.log
var/log/boot.log
var/log/syslog
var/log/syslog.1
tmp/access.log
etc/sensors.conf
etc/sensors3.conf
etc/host.conf
etc/pam.conf
etc/resolv.conf
etc/apt/apt.conf
etc/inetd.conf
etc/syslog.conf
etc/sysctl.conf
etc/sysctl.d/10-console-messages.conf
etc/sysctl.d/10-network-security.conf
etc/sysctl.d/10-process-security.conf
etc/sysctl.d/wine.sysctl.conf
etc/security/access.conf
etc/security/group.conf
etc/security/limits.conf
etc/security/namespace.conf
etc/security/pam_env.conf
etc/security/sepermit.conf
etc/security/time.conf
etc/ssh/sshd_config
etc/adduser.conf
etc/deluser.conf
etc/avahi/avahi-daemon.conf
etc/ca-certificates.conf
etc/ca-certificates.conf.dpkg-old
etc/casper.conf
etc/chkrootkit.conf
etc/debconf.conf
etc/dns2tcpd.conf
etc/e2fsck.conf
etc/esound/esd.conf
etc/etter.conf
etc/fuse.conf
etc/foremost.conf
etc/hdparm.conf
etc/kernel-img.conf
etc/kernel-pkg.conf
etc/ld.so.conf
etc/ltrace.conf
etc/mail/sendmail.conf
etc/manpath.config
etc/kbd/config
etc/ldap/ldap.conf
etc/logrotate.conf
etc/mtools.conf
etc/smi.conf
etc/updatedb.conf
etc/pulse/client.conf
usr/share/adduser/adduser.conf
etc/hostname
etc/networks
etc/timezone
etc/modules
etc/passwd
etc/passwd~
etc/passwd-
etc/shadow
etc/shadow~
etc/shadow-
etc/fstab
etc/motd
etc/hosts
etc/group
etc/group-
etc/alias
etc/crontab
etc/crypttab
etc/exports
etc/mtab
etc/hosts.allow
etc/hosts.deny
etc/os-release
etc/password.master
etc/profile
etc/default/grub
etc/resolvconf/update-libc.d/sendmail
etc/inittab
etc/issue
etc/issue.net
etc/login.defs
etc/sudoers
etc/sysconfig/network-scripts/ifcfg-eth0
etc/redhat-release
etc/debian_version
etc/fedora-release
etc/mandrake-release
etc/slackware-release
etc/suse-release
etc/security/group
etc/security/passwd
etc/security/user
etc/security/environ
etc/security/limits
etc/security/opasswd
boot/grub/grub.cfg
boot/grub/menu.lst
root/.ksh_history
root/.xauthority
usr/lib/security/mkuser.default
var/log/squirrelmail.log
var/log/apache2/squirrelmail.log
var/log/apache2/squirrelmail.err.log
var/lib/squirrelmail/prefs/squirrelmail.log
var/log/mail.log
etc/squirrelmail/apache.conf
etc/squirrelmail/config_local.php
etc/squirrelmail/default_pref
etc/squirrelmail/index.php
etc/squirrelmail/config_default.php
etc/squirrelmail/config.php
etc/squirrelmail/filters_setup.php
etc/squirrelmail/sqspell_config.php
etc/squirrelmail/config/config.php
etc/httpd/conf.d/squirrelmail.conf
usr/share/squirrelmail/config/config.php
private/etc/squirrelmail/config/config.php
srv/www/htdos/squirrelmail/config/config.php
var/www/squirrelmail/config/config.php
var/www/html/squirrelmail/config/config.php
var/www/html/squirrelmail-1.2.9/config/config.php
usr/share/squirrelmail/plugins/squirrel_logger/setup.php
usr/local/squirrelmail/www/readme
windows/system32/drivers/etc/hosts
windows/system32/drivers/etc/lmhosts.sam
windows/system32/drivers/etc/networks
windows/system32/drivers/etc/protocol
windows/system32/drivers/etc/services
/boot.ini
windows/debug/netsetup.log
windows/comsetup.log
windows/repair/setup.log
windows/setupact.log
windows/setupapi.log
windows/setuperr.log
windows/updspapi.log
windows/wmsetup.log
windows/windowsupdate.log
windows/odbc.ini
usr/local/psa/admin/htdocs/domains/databases/phpmyadmin/libraries/config.default.php
etc/apache2/conf.d/phpmyadmin.conf
etc/phpmyadmin/config.inc.php
etc/openldap/ldap.conf
etc/cups/acroread.conf
etc/cups/cupsd.conf
etc/cups/cupsd.conf.default
etc/cups/pdftops.conf
etc/cups/printers.conf
windows/system32/macromed/flash/flashinstall.log
windows/system32/macromed/flash/install.log
etc/cvs-cron.conf
etc/cvs-pserver.conf
etc/subversion/config
etc/modprobe.d/vmware-tools.conf
etc/updatedb.conf.beforevmwaretoolsinstall
etc/vmware-tools/config
etc/vmware-tools/tpvmlp.conf
etc/vmware-tools/vmware-tools-libraries.conf
var/log/vmware/hostd.log
var/log/vmware/hostd-1.log
# Wordpress
wp-config.php
# Symfony
config.yml
config_dev.yml
config_prod.yml
config_test.yml
parameters.yml
routing.yml
security.yml
services.yml
# Drupal
sites/default/default.settings.php
sites/default/settings.php
# Magento
app/etc/local.xml
# Sublime Text
sftp-config.json
# ASP.NET
Web.config
# vBulletin
includes/config.php
# OSCommerce
includes/configure.php
# phpMyAdmin
config.inc.php
# MediaWiki
LocalSettings.php
# MyBB
inc/config.php
# TYPO3
typo3conf/localconf.php
# Laravel
# Note: these entries might be benign in REQUEST_FILENAME
config/app.php
config/custom.php
config/database.php
# Joomla
# Note: this string might be benign in REQUEST_FILENAME
/configuration.php
# phpBB
# Note: this string might be benign in REQUEST_FILENAME
/config.php
EOF
	name = "test-lfi-os-files.dat-abca"
}
files {
	data = <<EOF
# Generic HTTP clients (popular libraries)

# http library
# http://search.cpan.org/~opera/HTTP-DAV/DAV.pm
dav.pm/v
# http library
# http://search.cpan.org/dist/libwww-perl/lib/LWP.pm
libwww-perl
# generic
mozilla/4.0 (compatible)
mozilla/4.0 (compatible; msie 6.0; win32)
mozilla/5.0 sf/
mozilla/5.0 sf//
# http library
# https://pypi.python.org/pypi/httplib2
python-httplib2
# http library
# http://docs.python-requests.org/en/master/
python-requests
# http library
# https://docs.python.org/2/library/urllib.html
Python-urllib
# http library
# https://github.com/typhoeus/typhoeus
typhoeus
# http library
# https://msdn.microsoft.com/en-us/library/windows/desktop/aa382925%28v=vs.85%29.aspx
winhttp.winhttprequest
EOF
	name = "test-scripting-user-agents.dat-abca"
}
files {
	data = <<EOF
__halt_compiler
apache_child_terminate
base64_decode
bzdecompress
call_user_func
call_user_func_array
call_user_method
call_user_method_array
convert_uudecode
file_get_contents
file_put_contents
fsockopen
gzdecode
gzinflate
gzuncompress
include_once
invokeargs
pcntl_exec
pcntl_fork
pfsockopen
posix_getcwd
posix_getpwuid
posix_getuid
posix_uname
ReflectionFunction
require_once
shell_exec
str_rot13
sys_get_temp_dir
wp_remote_fopen
wp_remote_get
wp_remote_head
wp_remote_post
wp_remote_request
wp_safe_remote_get
wp_safe_remote_head
wp_safe_remote_post
wp_safe_remote_request
zlib_decode
EOF
	name = "test-php-function-names-933150.dat-abca"
}
files {
	data = <<EOF
bin/bash
bin/csh
bin/dash
bin/du
bin/echo
bin/less
bin/ls
bin/more
bin/nc
bin/ps
bin/rbash
bin/sh
bin/sleep
bin/su
bin/tcsh
bin/uname
dev/fd/
dev/null
dev/stderr
dev/stdin
dev/stdout
dev/tcp/
dev/udp/
dev/zero
etc/group
etc/master.passwd
etc/passwd
etc/pwd.db
etc/shadow
etc/shells
etc/spwd.db
proc/self/
usr/bin/cc
usr/bin/clang
usr/bin/clang++
usr/bin/curl
usr/bin/env
usr/bin/fetch
usr/bin/file
usr/bin/find
usr/bin/ftp
usr/bin/gcc
usr/bin/head
usr/bin/id
usr/bin/less
usr/bin/more
usr/bin/nc
usr/bin/nice
usr/bin/nmap
usr/bin/perl
usr/bin/php
usr/bin/php5
usr/bin/php7
usr/bin/python
usr/bin/python2
usr/bin/python3
usr/bin/ruby
usr/bin/tail
usr/bin/top
usr/bin/uname
usr/bin/wget
usr/bin/who
usr/bin/whoami
usr/bin/xargs
usr/local/bin/bash
usr/local/bin/curl
usr/local/bin/nmap
usr/local/bin/perl
usr/local/bin/php
usr/local/bin/python
usr/local/bin/python2
usr/local/bin/python3
usr/local/bin/rbash
usr/local/bin/ruby
usr/local/bin/wget
EOF
	name = "test-unix-shell.dat-abca"
}
files {
	data = <<EOF
/.adSensepostnottherenonobook
/<invalid>hello.html
/actSensepostnottherenonotive
/acunetix-wvs-test-for-some-inexistent-file
/antidisestablishmentarianism
/appscan_fingerprint/mac_address
/arachni-
/cybercop
/nessus_is_probing_you_
/nessustest
/netsparker-
/rfiinc.txt
/thereisnowaythat-you-canbethere
/w3af/remotefileinclude.html
appscan_fingerprint
w00tw00t.at.ISC.SANS.DFind
w00tw00t.at.blackhats.romanian.anti-sec
EOF
	name = "test-scanners-urls.dat-abca"
}
files {
	data = <<EOF
abs
acos
adddate
addtime
aes_decrypt
aes_encrypt
ascii
asciistr
asin
atan
atan2
avg
benchmark
bin
bin_to_num
bit_and
bit_count
bit_length
bit_or
bit_xor
cast
ciel
cieling
char_length
char
character_length
charset
chr
coalesce
coercibility
collation
compress
concat_ws
concat
connection_id
conv
convert_tz
convert
cos
cot
count
dcount
cr32
curdate
current_date
current_time
current_timestamp
current_user
curtime
database
date
date_add
date_format
date_sub
datediff
day
dayname
dayofmonth
dayofweek
dayofyear
decode
default
degrees
des_decrypt
des_encrypt
dump
elt
encode
encrypt
exp
export_set
extract
extractvalue
field
field_in_set
find_in_set
floor
format
found_rows
from_base64
from_days
from_unixtime
get_format
get_lock
greatest
group_concat
hex
hextoraw
rawtohex
hour
if
ifnull
in
inet6_aton
inet6_ntoa
inet_aton
inet_ntoa
insert
instr
interval
isnull
is_free_lock
is_ipv4_compat
is_ipv4_mapped
is_ipv4
is_ipv6
is_not_null
is_not
is_null
is_used_lock
last
last_day
last_inser_id
lcase
least
left
length
ln
load_file
local
localtimestamp
locate
log
log2
log10
lower
lpad
ltrim
make_set
makedate
master_pos_wait
max
md5
microsecond
mid
min
minute
mod
month
monthname
name_const
not_in
now
nullif
oct
octet_length
old_password
ord
password
period_add
period_diff
pi
position
pow
power
procedure_analyse
quarter
quote
radians
rand
release_lock
repeat
replace
reverse
right
round
row_count
rpad
rtrim
schema
sec_to_time
second
session_user
sha
sha1
sha2
sign
sin
pg_sleep
sleep
soundex
space
sqrt
std
stddev_pop
stddev_samp
str_to_date
strcmp
subdate
substring
substring_index
substr
subtime
sum
sysdate
system_user
tan
time
timestamp
timestampadd
timestampdiff
timediff
time_format
time_to_sec
to_base64
todays
toseconds
tochar
tonchar
trim
truncate
ucase
uncompress
uncompressed_length
unhex
unix_timestamp
updatexml
upper
user
utc_date
utc_time
utc_timestamp
uuid
uuid_short
values
var_pop
var_samp
variance
version
week
weekday
weekofyear
weight_string
year
yearweek
xmltype
EOF
	name = "test-sql-function-names.dat-abca"
}
files {
	data = <<EOF
acunetix-product
(acunetix web vulnerability scanner
acunetix-scanning-agreement
acunetix-user-agreement
myvar=1234
x-ratproxy-loop
bytes=0-,5-0,5-1,5-2,5-3,5-4,5-5,5-6,5-7,5-8,5-9,5-10,5-11,5-12,5-13,5-14
x-scanner
EOF
	name = "test-scanners-headers.dat-abca"
}
files {
	data = <<EOF
# Vulnerability scanners, bruteforce password crackers and exploitation tools

# password cracker
# http://sectools.org/tool/hydra/
(hydra)
# vuln scanner
# http://virtualblueness.net/nasl.html
.nasl
# sql injection
# https://sourceforge.net/projects/absinthe/
absinthe
# email harvesting
# dead? 2004
advanced email extractor
# vuln scanner
# http://www.arachni-scanner.com/
arachni/
#
autogetcontent
# nessus frontend
# http://www.crossley-nilsen.com/Linux/Bilbo_-_Nessus_WEB/bilbo_-_nessus_web.html
# dead? 2003
bilbo
# Backup File Artifacts Checker
# https://github.com/mazen160/bfac
BFAC
# password cracker
# http://sectools.org/tool/brutus/
brutus
brutus/aet
# sql injection
# https://www.notsosecure.com/bsqlbf-v2-blind-sql-injection-brute-forcer/
bsqlbf
# vuln scanner
# http://freecode.com/projects/cgichk dead? 2001
cgichk
# vuln scanner
# https://sourceforge.net/projects/cisco-torch/
cisco-torch
# vuln scanner
# https://github.com/stasinopoulos/commix
commix
# MS FrontPage vuln scanner?
core-project/1.0
# vuln scanner?
crimscanner/
# vuln scanner
datacha0s
# hidden page scanner
# https://www.owasp.org/index.php/Category:OWASP_DirBuster_Project
dirbuster
# vuln scanner
# https://sourceforge.net/projects/dominohunter/
domino hunter
# vuln scanner - directory traversal fuzzer
# https://github.com/wireghoul/dotdotpwn
dotdotpwn
#
#
email extractor
# vuln scanner
fhscan core 1.
#
floodgate
#
get-minimal
# vuln scanner
gootkit auto-rooter scanner
#
grabber
# vuln scanner
# https://sourceforge.net/projects/grendel/
grendel-scan
# sql injection
havij
# vuln scanner - path disclosure finder
# http://seclists.org/fulldisclosure/2010/Sep/375
inspath
#
internet ninja
# vuln scanner
jaascois
# vuln scanner
zmeu
# port scanner
# https://github.com/robertdavidgraham/masscan
masscan
# vuln scanner
# http://www.severus.org/sacha/metis/
metis
# vuln scanner
morfeus fucking scanner
# sql injection
# https://github.com/dtrip/mysqloit
mysqloit
# vuln scanner
# http://www.nstalker.com/
n-stealth
# vuln scanner
# http://www.tenable.com/products/nessus-vulnerability-scanner
nessus
# vuln scanner
# https://www.netsparker.com/web-vulnerability-scanner/
netsparker
# vuln scanner
# https://cirt.net/Nikto2
nikto
# vuln scanner
nmap nse
nmap scripting engine
nmap-nse
# vuln scanner
# http://www.nsauditor.com/
nsauditor
# vuln scanner
# http://www.openvas.org/
openvas
# sql injection
# http://www.vealtel.com/software/nosec/pangolin/
pangolin
# web proxy & vuln scanner
# https://sourceforge.net/projects/paros/
paros
# phpmyadmin vuln scanner
# dead 2005?
pmafind
#
prog.customcrawler
# vuln scanner
# https://www.qualys.com/suite/web-application-scanning/
qualys was
# 
s.t.a.l.k.e.r.
#
security scan
# vuln scanner
# https://sourceforge.net/projects/springenwerk/
springenwerk
# sql injection
# http://www.sqlpowerinjector.com/
sql power injector
# sql injection
# http://sqlmap.org/
sqlmap
# sql injection
# http://sqlninja.sourceforge.net/
sqlninja
# password cracker
# http://foofus.net/goons/jmk/medusa/medusa.html
the forest lobster
# 
this is an exploit
# vuln scanner?
toata dragostea
toata dragostea mea pentru diavola
# SQL bot
# http://tools.cisco.com/security/center/viewIpsSignature.x?signatureId=22142&signatureSubId=0
uil2pn
# badly scripted UAs (e.g. User-Agent: User-Agent: foo)
user-agent:
# vuln scannr
# https://subgraph.com/vega/
vega/
# vuln scanner
# dead?
voideye
# vuln scanner
# http://w3af.org/
w3af.sf.net
w3af.sourceforge.net
w3af.org
# site scanner (legacy)
# http://www.robotstxt.org/db/webbandit.html
webbandit
# vuln scanner
# http://www8.hp.com/us/en/software-solutions/webinspect-dynamic-analysis-dast/
webinspect
# site scanner
# http://www.scrt.ch/en/attack/downloads/webshag
webshag
# vuln scanner
# dead?
webtrends security analyzer
# vuln scanner
# https://github.com/hhucn/webvulnscan
webvulnscan
# web technology scanner
# https://www.morningstarsecurity.com/research/whatweb
whatweb
# vuln scanner
whcc/
# exploit poc
wordpress hash grabber
# exploit
xmlrpc exploit
# wordpress vuln scanner
# https://wpscan.org/
WPScan
EOF
	name = "test-scanners-user-agents.dat-abca"
}
files {
	data = <<EOF
<h2 style="font:8pt/11pt verdana; color:000000">HTTP 403.6 - Forbidden: IP address rejected<br>
<TITLE>500 Internal Server Error</TITLE>
Microsoft VBScript runtime (0x8
error '800
Application uses a value of the wrong type for the current operation
Microsoft VBScript compilation (0x8
Microsoft VBScript compilation error
Microsoft .NET Framework Version:
A trappable error occurred in an external object. The script cannot continue running
Microsoft VBScript runtime Error
>Syntax error in string in query expression
ADODB.Command
Object required: '
EOF
	name = "test-iis-errors.dat-abca"
}
	config {
		regex_match_limit = "1500"
		allowed_request_content_types = ["application/x-www-form-urlencoded","multipart/form-data","text/xml","application/xml","application/x-amf","application/json"]
		allowed_methods = ["HTTP_METHOD_GET","HTTP_METHOD_HEAD","HTTP_METHOD_POST","HTTP_METHOD_OPTIONS"]
		response_hdr_default_action = "phase:3,deny,status:403,log,auditlog"
		cookie_format_version = "0"
		request_body_default_action = "phase:2,deny,status:403,log,auditlog"
		response_body_default_action = "phase:4,deny,status:403,log,auditlog"
		request_hdr_default_action = "phase:1,deny,status:403,log,auditlog"
		static_extensions = [".gif",".jpg",".jpeg",".png",".js",".css",".ico",".svg",".webp"]
		restricted_extensions = [".asa",".asax",".ascx",".axd",".backup",".bak",".bat",".cdx",".cer",".cfg",".cmd",".com",".config",".conf",".cs",".csproj",".csr",".dat",".db",".dbf",".dll",".dos",".htr",".htw",".ida",".idc",".idq",".inc",".ini",".key",".licx",".lnk",".log",".mdb",".old",".pass",".pdb",".pol",".printer",".pwd",".resources",".resx",".sql",".sys",".vb",".vbs",".vbproj",".vsdisco",".webinfo",".xsd",".xsx"]
		allowed_http_versions = ["ONE_ZERO","ONE_ONE"]
		restricted_headers = ["Proxy-Connection","Lock-Token","Content-Range","Translate","via","if"]
		argument_separator = "&"
	}
	tenant_ref = data.avi_tenant.default_tenant.id
	name = "test-System-WAF-Profile-abc"
}

data "avi_wafprofile" "testWafProfile" {
    name= "${avi_wafprofile.testWafProfile.name}"
}
`
