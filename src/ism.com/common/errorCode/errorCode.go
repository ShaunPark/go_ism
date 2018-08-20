package errorCode

const _COMMON_BASE = 0
const _RULE_BASE = 100
const _SVCMAP_BASE = 200
const _TRNS_BASE = 300
const _DDR_BASE = 400
const _VALIDATE_BASE = 500

const _BAT_DEF_BASE = 900
const _ONL_BASE = 1000
const _DFD_BASE = 2000
const _BAT_BASE = 3000

const _SOC_BASE = 4000
const _FTP_BASE = 4100
const _DB_BASE = 4200
const _TELNET_BASE = 5300
const _LICENSE_BASE = 6000

const _WAS_BASE = 6300
const _SAP_BASE = 7000
const _WEB_BASE = 7500

const NOTDEFINED = -1

// Common
const FISS_COMMON = _COMMON_BASE

// Loading Config and Error	Message
const CONF_CORRUPTED_FILE = _COMMON_BASE + 1 // "fail to	read %S	file[%S]. File is corrupted"
const CONF_LOADING_FAIL = _COMMON_BASE + 2   // "fail to	read %S	file[%S]."
const CONF_ILLEGAL_TAG = _COMMON_BASE + 3    // "root tag is	not	%S - %S"
const CONF_PROPERTY_FAIL = _COMMON_BASE + 4  // "fail to	load %S	: %S"

// ClassLoader ( CODE :	CL )
const CLASSLOADER_COMMON = _COMMON_BASE + 5    // "fail to	load class"
const CLASSLOADER_NOT_FOUND = _COMMON_BASE + 6 // "class not found	[%S]"
const CLASSLOADER_INIT_FAIL = _COMMON_BASE + 7 // "class init fail	[%S]"
const ROUTING_TARGET_FAIL = _COMMON_BASE + 8   //routing failed.
const INVALID_APP = _COMMON_BASE + 9           //invalid application
const EXCEPTION_HANDLER_FAIL = _COMMON_BASE + 10
const NOT_IMPLEMENTED = _COMMON_BASE + 11
const NOT_SUPPORTED = _COMMON_BASE + 12

// DDR ( CODE :	DD )
const DDR_COMMON = _DDR_BASE
const DDR_INPUT_PARSE_FAIL = _DDR_BASE + 1
const DDR_PARSER_INIT_FAIL = _DDR_BASE + 2
const DDR_ROUTING_INPUT_NOT_DEFINED = _DDR_BASE + 3
const DDR_ROUTING_INVALID_ROUTING_INPUT = _DDR_BASE + 4
const DDR_INVALID_ROUTING_INPUT_DATA_INDEX = _DDR_BASE + 5
const DDR_INVALID_ROUTING_INPUT_DETAIL_INDEX = _DDR_BASE + 6
const DDR_INVALID_ROUTING_INPUT_COLUMN_INDEX = _DDR_BASE + 7
const DDR_NO_MAP_MATCHED = _DDR_BASE + 8
const DDR_INVALID_TARGET_SYSTEM_INDEX = _DDR_BASE + 9
const DDR_INVALID_INPUT_INDEX = _DDR_BASE + 10
const DDR_INVALID_INPUT_DATA = _DDR_BASE + 11
const DDR_GET_RULE_FAIL = _DDR_BASE + 12 //	"fail to get interface info	in service map process : [%S:%S]"

// Rule	( CODE : RU	)
const RULE_CACHE_COMMON = _RULE_BASE + 0
const RULE_CACHE_INIT_FAIL = _RULE_BASE + 1     //	"Rule Cache	Init Failure : %S"
const RULE_CACHE_ACCESS_FAIL = _RULE_BASE + 2   //	"Rule Cache	Access Failure : %S"
const RULE_CACHE_INF_NOT_FOUND = _RULE_BASE + 3 //	"Rule Cache	Interface Class	not	found :	%S"
const RULE_CACHE_REPOSITORY_CONNECT_FAIL = _RULE_BASE + 4
const RULE_NOT_FOUND = _RULE_BASE + 5
const RULE_FIND_FAIL_WITH_DB_ERROR = _RULE_BASE + 6
const RULE_FIND_FAIL_IN_REPOSITORY = _RULE_BASE + 7 //	"리파지토리에서	룰을 찾을 수 없습니다[%S:%S]"
const RULE_CACHE_ILLEGAL_INTF = _RULE_BASE + 8      //	"illegal InterfaceInfo %S[%S] -	%S"
const RULE_CACHE_NO_TARGET_SYSTEM = _RULE_BASE + 9  //	"no	target system -	sequence[%S]"
const RULE_NO_SERVICE_MAP = _RULE_BASE + 10         //	"no	ServiceMap [%S]"
const RULE_NO_MASTER_INFO = _RULE_BASE + 11         //	"no	master info."
const RULE_ERRCODE_INIT_FAIL = _RULE_BASE + 12
const RULE_TYPE_UNMATCH = _RULE_BASE + 13             //unmatched service type. ex) batch request but not batch rule.
const RULE_SVCTYPE_UNMATCH = _RULE_BASE + 14          //unmatched service type. ex) batch direct request. but source service typ is not file type.
const RULE_INVALID_SVC_ENTITY_COUNT = _RULE_BASE + 15 //서비스 엔티티 수가 잘못된 경우.

// Validate	- 600
const VALIDATE_FAIL = _VALIDATE_BASE + 1
const VALIDATE_INPUT_IS_NULL = _VALIDATE_BASE + 2
const VALIDATE_INPUT_FORMAT_FAIL = _VALIDATE_BASE + 3
const VALIDATE_VALUE_IS_NULL = _VALIDATE_BASE + 4
const VALIDATE_UNSUPPORTED_OPERATION = _VALIDATE_BASE + 5
const VALIDATE_INVALID_INPUT_COUNT = _VALIDATE_BASE + 6
const VALIDATE_LENGTH_FAIL = _VALIDATE_BASE + 7

// Connector Socket	( CONSO	)
const SOC_COMMON = _SOC_BASE
const SOC_CONNECT_FAIL = _SOC_BASE + 1
const SOC_CONNECT_TIMEOUT = _SOC_BASE + 2
const SOC_CONNECT_RECEIVE_TIMEOUT = _SOC_BASE + 3

// FTP ( FTP00 )
const FTP_COMMON = _FTP_BASE + 0
const FTP_NO_SERVER_INFO = _FTP_BASE + 1         // "no ftp server	info."
const FTP_PWD_DECRYPT_FAIL = _FTP_BASE + 2       // "fail to decrypt ftp password."
const FTP_CONNECT_FAIL = _FTP_BASE + 3           // failed	to connect ftp server
const FTP_LOGIN_FAIL = _FTP_BASE + 4             // fail to login ftp server [user=<%S>]
const FTP_SET_ACTIVE_MODE_FAIL = _FTP_BASE + 5   // "failed to	set	ftp	active mode"
const FTP_SET_TRANSFER_TYPE_FAIL = _FTP_BASE + 6 // "failed to	set	ftp	transfer type"
const FTP_CHANGE_DIRECTORY_FAIL = _FTP_BASE + 7  // fail to change	directory :<%S>
const FTP_FETCH_COUNT_FAIL = _FTP_BASE + 8       // fail to get file count
const FTP_GET_FAIL = _FTP_BASE + 9               // fail to get file <%S>
const FTP_GET_ERROR = _FTP_BASE + 10             // failed	to get files
const FTP_ONLY_OVERWRITE = _FTP_BASE + 11        // overwrite 만 가능함
const FTP_FILE_ALREADY_EXIST = _FTP_BASE + 12    // file already exists!
const FTP_PUT_FAIL = _FTP_BASE + 13              // fail to put file <%S>
const FTP_SIZE_NOT_SAME = _FTP_BASE + 14         // file size is not same. original[<%S>][<%S>]
const FTP_SET_ACCESS_MODE_FAIL = _FTP_BASE + 15  // fail to change file access mode, System Type [%S]
const FTP_PUT_ERROR = _FTP_BASE + 16             // failed	to put files
const FTP_REMOVE_ERROR = _FTP_BASE + 17          // failed	to remove files
const FTP_CHMOD_FAIL = _FTP_BASE + 18            // Change	mode failed
const FTP_FILE_NOT_FOUND_ERROR = _FTP_BASE + 19  // no	source file
const FTP_FAIL_CREATE_LOCAL_DIR = _FTP_BASE + 20 // "fail to create directory in shared path :	%S"
const FILE_SEP_ERROR = _FTP_BASE + 21            //file separation failed.[%s]
const FTP_RENAME_ERROR = _FTP_BASE + 22          // failed to rename file

const FTP_SET_SOURCE_TRANSFER_TYPE_FAIL = _FTP_BASE + 31
const FTP_SET_SOURCE_TRANSFER_TYPE_ERROR = _FTP_BASE + 32
const FTP_SOURCE_LIST_FILE_ERROR = _FTP_BASE + 33
const FTP_SOURCE_CHANGE_DIRECTORY_FAIL = _FTP_BASE + 34
const FTP_SOURCE_CHANGE_DIRECTORY_ERROR = _FTP_BASE + 35
const FTP_SET_TARGET_TRANSFER_TYPE_FAIL = _FTP_BASE + 51
const FTP_SET_TARGET_TRANSFER_TYPE_ERROR = _FTP_BASE + 52
const FTP_TARGET_LIST_FILE_ERROR = _FTP_BASE + 53
const FTP_TARGET_CHANGE_DIRECTORY_FAIL = _FTP_BASE + 54
const FTP_TARGET_CHANGE_DIRECTORY_ERROR = _FTP_BASE + 55
const FTP_TARGET_TEMPFILE_RENAME_FAIL = _FTP_BASE + 56
const FTP_TARGET_TEMPFILE_RENAME_ERROR = _FTP_BASE + 57
const FTP_TARGET_SET_ACCESS_MODE_FAIL = _FTP_BASE + 58
const FTP_TARGET_SET_ACCESS_MODE_ERROR = _FTP_BASE + 59
const FTP_SOURCE_SENT_FILE_JOB_ERROR = _FTP_BASE + 71

// DB (	DB000 )
const DB_COMMON = _DB_BASE + 0
const DB_NO_SERVER_INFO = _DB_BASE + 1         // "no DB server info."
const DB_NO_DBNAME = _DB_BASE + 2              // "no DB name."
const DB_NO_USERID = _DB_BASE + 3              // "no DB User ID."
const DB_NO_USERPASSWORD = _DB_BASE + 4        // "no DB User Password."
const DB_PASSWORD_DECRYPT_FAIL = _DB_BASE + 5  // "fail to	decrypt	DB User	password."
const DB_PASSWORD_DECRYPT_ERROR = _DB_BASE + 6 // "fail to	decrypt	DB User	password. Decrypt result is	null."
const DB_NO_HOST = _DB_BASE + 7                // "no DB Host."
const DB_NO_PORT = _DB_BASE + 8                // "no DB Port."
const DB_NO_DATABASE_TYPE = _DB_BASE + 9       // 데이터베이스	타입이 설정되지	않았습니다.
const DB_NOT_SUPPORTED_DB_TYPE = _DB_BASE + 10 // 지원되지	않는 DB	타입["+_dbType+"] 입니다.
const DB_PROVIDER_ERROR = _DB_BASE + 11
const DB_NO_PROPERTIES = _DB_BASE + 12         // 데이터베이스	설정 정보가	없습니다.
const DB_CONNECT_FAIL = _DB_BASE + 13          // DB 연결에 실패했습니다.
const DB_CONNECTION_RETRY_FAIL = _DB_BASE + 14 // 전체	반영할 Data	수 조회	중 Exception["+e.getMessage()+"]이 발생. 재시도	횟수["+_retryCount+"]를	초과하여 작업을	중단합니다.")
const DB_FETCH_COUNT_FAIL = _DB_BASE + 15      // 데이터 건수 조회	중 오류	발생.
const DB_NO_QUERY = _DB_BASE + 16              // SQL문이 존재하지	않습니다.
const DB_NO_CONNECTION = _DB_BASE + 17         // DB가	연결되어 있지 않습니다.
const DB_FETCH_DATA_FAIL = _DB_BASE + 18       // "[데이터조회]["+interfaceId+"] 데이터 조회 오류입니다:"+e.toint()
const DB_FETCH_DATA_ERROR = _DB_BASE + 19      // 데이터 조회 중에	오류 발생했습니다.
const DB_DECOMPRESS_FAIL = _DB_BASE + 20       // 압축된 데이터를 복구하는	과정에서 오류가	발생하였습니다.
const DB_NEGATIVE_INDEX = _DB_BASE + 21        // Index 값[" +	index +	"]이 음수입니다.
const DB_INDEX_ERROR = _DB_BASE + 22           // "Index 값[" + index + "]이 Result의 길이["+results.length+"]보다	깁니다."
const DB_NO_RESULT = _DB_BASE + 23
const DB_QUERY_HANDLER = _DB_BASE + 24      //
const DB_FETCH_FILE_TOO_BIG = _DB_BASE + 25 //

// ServiceMap -	700
const SVCMAP_COMMON = _SVCMAP_BASE                 //
const SVCMAP_GET_RULE_FAIL = _SVCMAP_BASE + 1      // "fail to	get	interface info in service map process :	[%S:%S]"
const SVCMAP_NO_SERVICE_DEFINED = _SVCMAP_BASE + 2 //	"fail to get interface info	in service map process : [%S:%S]"
const SVCMAP_MATCH_CLASS_INIT_FAIL = _SVCMAP_BASE + 3

// Transformer ( CODE :	TR )
const TRNS_COMMON = _TRNS_BASE + 0
const TRNS_ILLEGAL_DATA = _TRNS_BASE + 1              //	"data	is null	or illegal index[%S]"
const TRNS_ILLEGAL_COL_INDEX = _TRNS_BASE + 2         //	"Illegal column	index	[%S], Data Length: %S"
const TRNS_LENGTH_PARSING_ERROR = _TRNS_BASE + 3      //	"variable	length field parsing error [COLUMN:%S -	%S]"
const TRNS_LENGTH_PARSING_FAIL = _TRNS_BASE + 4       //	"variable	length field parsing fail [COLUMN:%S, LENGTH:%S]"
const TRNS_NO_MASTER = _TRNS_BASE + 5                 //	"no	master info. DataStructure ID: %S"
const TRNS_NO_FIELD = _TRNS_BASE + 6                  //	"no	field	info. Data ID: %S"
const TRNS_ILLEGAL_ROW_INDEX = _TRNS_BASE + 7         //	"illegal row index %S"
const TRNS_ILLEGAL_DATA_FORMAT = _TRNS_BASE + 8       //	"Illegal data	format"
const TRNS_ILLEGAL_DETAIL_INDEX = _TRNS_BASE + 9      //	"Illegal detail	index	[%S]"
const TRNS_NO_COLUMN_DATA = _TRNS_BASE + 10           // "column[%S] has no	data"
const TRNS_ILLEGAL_DATA_INDEX = _TRNS_BASE + 11       // "Illegal data	index [%S]"
const TRNS_ARRAY_COPY_ERROR = _TRNS_BASE + 12         // ArrayCopy중 Exception	발생 [<%S>:<%S>:<%S>]
const TRNS_CRUD_TYPE_ERROR = _TRNS_BASE + 13          // cannot	accpet	crudTypeCd : <%S>
const TRNS_NO_KEY_FIELD = _TRNS_BASE + 14             // ""there is	no	key	field. so Transformer can't	write where	clause."
const TRNS_CONVERT_DATEFORMAT_ERROR = _TRNS_BASE + 15 // "fail	to convert date	format [%S - %S]"
const TRNS_CONVERT_BINARY_FAIL = _TRNS_BASE + 16      // "Transformer cannot convert binary	type :	%S"
const TRNS_NO_DATASTRUCTURE = _TRNS_BASE + 17         // "no data structure."
const TRNS_NO_DATA = _TRNS_BASE + 18                  // "no data [Data	Strcuture ID: %S]"
const TRNS_NO_HEADER = _TRNS_BASE + 19                //	"no	Header	[System	ID:	%S]"
const TRNS_FAIL_GET_RULE = _TRNS_BASE + 20
const TRNS_NO_ERR_DATASTRUCTURE = _TRNS_BASE + 21 // "no data structure."
const TRNS_INVALID_REPEAT_COUNT_DATA_INDEX = _TRNS_BASE + 22
const TRNS_INVALID_REPEAT_COUNT_FIELD_INDEX = _TRNS_BASE + 23
const TRNS_INVALID_REPEAT_COUNT_FORMAT = _TRNS_BASE + 24
const TRNS_INVALID_REPEAT_COUNT_VALUE = _TRNS_BASE + 25
const TRNS_CUSTOMFUCNTION_CLASS_NOT_FOUND = _TRNS_BASE + 26
const TRNS_CUSTOMFUCNTION_CLASS_ILLEGALACCESS = _TRNS_BASE + 27
const TRNS_CUSTOMFUCNTION_CLASS_INSTANTIATION_FAIL = _TRNS_BASE + 28
const TRNS_CUSTOMFUNCTION_CONVERT_FAIL = _TRNS_BASE + 29
const TRNS_CUSTOMFUNCTION_DIVIDE_BY_ZERO = _TRNS_BASE + 30
const TRNS_CUSTOMFUNCTION_OPERATION_FAIL = _TRNS_BASE + 31
const TRNS_CUSTOMFUNCTION_RULE_PARSE_FAIL = _TRNS_BASE + 32
const TRNS_CUSTOMFUNCTION_RULE_EXECUTE_FAIL = _TRNS_BASE + 33
const TRNS_INVALID_DATA_LENGTH = _TRNS_BASE + 34
const TRNS_HEADER_MAP_RESULT_NULL = _TRNS_BASE + 35       // 2008-05-16 From 35 ~
const TRNS_HEADER_PARSED_RESULT_NULL = _TRNS_BASE + 36    // 2008-05-16
const TRNS_HEADER_PARSED_LENGTH_ERROR = _TRNS_BASE + 37   // 2008-05-16
const TRNS_MASTER_OFFSET_INFO_IS_NULL = _TRNS_BASE + 38   // 2008-05-16
const TRNS_INVALID_MASTER_OFFSET_INDEX = _TRNS_BASE + 39  // 2008-05-16
const TRNS_DETAIL_OFFSET_INFO_IS_NULL = _TRNS_BASE + 40   // 2008-05-16
const TRNS_DETAIL_OFFSET_INFO_IS_NULL_2 = _TRNS_BASE + 41 // 2008-05-16
const TRNS_DETAIL_OFFSET_INFO_IS_NULL_3 = _TRNS_BASE + 42 // 2008-05-16
const TRNS_INVALID_DETAIL_OFFSET_INDEX = _TRNS_BASE + 43  // 2008-05-16
const TRNS_DATA_MAP_RESULT_NULL = _TRNS_BASE + 44         // 2008-05-16
const TRNS_DATA_PARSED_RESULT_NULL = _TRNS_BASE + 45      // 2008-05-16
const TRNS_DATA_PARSED_LENGTH_NULL = _TRNS_BASE + 46      // 2008-05-16
const TRNS_DATA_PARSED_LENGTH_ERROR = _TRNS_BASE + 47     // 2008-05-16
const TRNS_DETAIL_PARSED_LENGTH_NULL = _TRNS_BASE + 48
const TRNS_DETAIL_PARSED_LENGTH_ERROR = _TRNS_BASE + 49
const TRNS_DETAIL_PARSED_LENGTH_NULL_2 = _TRNS_BASE + 50
const TRNS_DETAIL_PARSED_LENGTH_ERROR_2 = _TRNS_BASE + 51
const TRNS_DETAIL_PARSED_LENGTH_NULL_3 = _TRNS_BASE + 52
const TRNS_DETAIL_PARSED_LENGTH_ERROR_3 = _TRNS_BASE + 53
const TRNS_MAPPED_DATA_NULL = _TRNS_BASE + 54
const TRNS_MAPPED_DATA_LENGTH_ERROR = _TRNS_BASE + 55
const TRNS_MAPPED_DETAIL_NULL = _TRNS_BASE + 56
const TRNS_MAPPED_DETAIL_LENGTH_ERROR = _TRNS_BASE + 57
const TRNS_MAPPED_FIELD_LENGTH_ERROR = _TRNS_BASE + 58
const TRNS_VARIABLE_LENGTH_FIELD_VALUE_NULL = _TRNS_BASE + 59
const TRNS_RULE_INIT_FAIL = _TRNS_BASE + 60
const TRNS_ENCODING_FAIL = _TRNS_BASE + 61
const TRNS_RECORDDELIMTER_NOT_DIFINED = _TRNS_BASE + 62

//
//	Online Error Code
//
const ONL_SOURCE_COMMON = _ONL_BASE

// Global ID ( CODE	: GI 	)
const ONL_GID_COMMON = _ONL_BASE + 1
const ONL_GID_NO_CHECKER = _ONL_BASE + 2        // "no Global ID checker"
const ONL_GID_MGR_CONNECT_FAIL = _ONL_BASE + 3  // "fail to	connect	to GID Check Manager[%S]"
const ONL_GID_HANDLER_UNKNOWN = _ONL_BASE + 4   // "unknown	handler	type %S	- %S"
const ONL_GID_HANDLER_LOAD_FAIL = _ONL_BASE + 5 // "fail to	load Handler - %S"
const ONL_GID_HANDLER_INIT_FAIL = _ONL_BASE + 6 // "fail to	initialize Handler"
const ONL_GID_CHECK_FAIL = _ONL_BASE + 7        // "fail to	check GID %S"

// Domain Route	( CODE : DR	)
const ONL_ROUTE_COMMON = _ONL_BASE + 8
const ONL_ROUTE_NO_FILE = _ONL_BASE + 9            // "no %S file %S"
const ONL_ROUTE_LOAD_FAIL = _ONL_BASE + 10         //	"fail to load %S [%S]"
const ONL_ROUTE_NO_DOMAIN = _ONL_BASE + 11         //	"no	domain info."
const ONL_ROUTE_CONTEXT_FAIL = _ONL_BASE + 12      //	fail to	create Context [%S,	%S]
const ONL_ROUTE_ALL_TROUBLE = _ONL_BASE + 13       //	"all domain	in trouble"
const ONL_ROUTE_CLASS_NOT_DEFINED = _ONL_BASE + 14 //	"can not find route	class name in config file"
const ONL_ROUTE_CLASS_NOT_FOUND = _ONL_BASE + 15   //	"can not find route	class"
const ONL_ROUTE_CLASS_INIT_FAIL = _ONL_BASE + 16   //	"route class init failed [%S]"
const ONL_ROUTE_INTERNAL_TIMEOUT = _ONL_BASE + 17  // timeout internally

// ServiceBlock	( CODE : SB	)
const ONL_SB_COMMON = _ONL_BASE + 18
const ONL_SB_NO_SERVICE_INFO = _ONL_BASE + 19 //	"Can't find	service	info %S"
const ONL_SB_BLOCKING = _ONL_BASE + 20        //	"%S	is blocking"

// HealthCheck ( CODE :	HC )
const ONL_HC_COMMON = _ONL_BASE + 21
const ONL_HC_FAIL_TO_LOAD_SYSTEMID = _ONL_BASE + 22           //	"fail to load System IDs"
const ONL_HC_FAIL_TO_LOAD_HEALTHINFO = _ONL_BASE + 23         //	"fail to load HealthInfo for Target	System[%S]"
const ONL_HC_FAIL_TO_UPDATE_HEALTH_STATUS = _ONL_BASE + 24    //	"fail to update	Health status for %S[%S] - %S"
const ONL_HC_FAIL_TO_BROADCAST_HEALTH_STATUS = _ONL_BASE + 25 //	"fail to broadcast for %S[%S] -	%S [%S]"
const ONL_HC_CONNECT_FAIL = _ONL_BASE + 26                    //	"fail to connect
const ONL_HC_FAIL_TO_GET_HEALTHINFO = _ONL_BASE + 27          //	"fail to get HealthInfo	for	Target Systemp[%S]"
const ONL_HC_FAIL_TO_SET_HEALTHINFO = _ONL_BASE + 28          //	"fail to set Health	status for %S[%S] -	%S"
const ONL_HC_NO_HEALTHENTITY = _ONL_BASE + 29                 //	"no	HealthCheckEntity [%S]"
const ONL_HC_NO_HEALTHINFO = _ONL_BASE + 30                   //	"no	HealthCheckInfo	[%S]"
const ONL_HC_ALL_TROUBLE = _ONL_BASE + 31                     //	"all HealthCheckEntity in trouble"
const ONL_HC_INDEX_OUT_OF_BOUNDS = _ONL_BASE + 32             //	"request index(%S) is out of bounds(%S)"
const ONL_NO_TARGET_SYSTEM = _ONL_BASE + 33
const ONL_NO_TARGET_APPSYSTEM = _ONL_BASE + 34
const ONL_NO_TARGET_APPINFO = _ONL_BASE + 35
const ONL_TARGET_TIMEOUT = _ONL_BASE + 36
const ONL_TARGET_NULL = _ONL_BASE + 37            // failed	to find	target model
const ONL_TARGET_NAME_ERROR = _ONL_BASE + 38      // naming	exception
const ONL_TARGET_INTERFACE_ERROR = _ONL_BASE + 39 // naming	exception

// Online Target
const ONL_TARGET_COMMON = _ONL_BASE + 40
const ONL_TIMEOUT_INVALID_VALUE = _ONL_BASE + 41

// 송수신 처리
const ONL_SA_WAIT_TIMEOUT = _ONL_BASE + 42 // 동기 비동기 패턴에서 수신응답을 대기하던중 타임아웃이 발생하였습니다.
// 송신 소켓 커넥터 쓰레드 수 초과
const ONL_SOC_SRC_MAX = _ONL_BASE + 43 // EAI서버가 거래를 추가 처리할 수 없는 상태입니다.
/**
 * 업무 오류 코드(공통). not used for error message. only for log
 */
const ONL_COMMON_BIZ_ERRORCODE = _ONL_BASE + 44 // 업무 오류 코드(공통). not used for error message. only for log
const ONL_INVALID_RETURN_INDEX = _ONL_BASE + 45 // 리턴인덱스를 DDR에의해 호출 안한 경우

// TP 서비스 혹은  EJB 가 없는 경우
const ONL_TGT_SVC_NOT_EXIST = _ONL_BASE + 46      // 타겟시스템에  대상 서비스가 없습니다.
const ONL_MESSAGE_TOO_SHORT = _ONL_BASE + 47      // online message is too short
const ONL_TARGET_RETURN_MSG_NULL = _ONL_BASE + 48 // 20100323 shpark naming	exception
const ONL_SA_INVALID_NODE_INDEX = _ONL_BASE + 49  // 20100407 shpark 비동기 응답의 node 인덱스가 잘못된 경우
const ONL_INVALID_MESSAGE = _ONL_BASE + 50        // 20100407 shpark 전문 값 파싱중 값 오류

// Router 오류 코드 추가
const ONL_ROUTE_TARGET_DOWN = _ONL_BASE + 51             // 20100407 shpark 전문 값 파싱중 값 오류
const ONL_ROUTE_TARGET_BUSY = _ONL_BASE + 52             // 20100407 shpark 전문 값 파싱중 값 오류
const ONL_DB_SERVICE_NOT_MATCH = _ONL_BASE + 53          // 20100407 shpark 전문 값 파싱중 값 오류
const ONL_INVALID_GID = _ONL_BASE + 54                   // 20100407 shpark 전문 값 파싱중 값 오류
const ONL_XA_NOT_SUPPORTED_IN_SIMULCALL = _ONL_BASE + 55 // 20100407 shpark 전문 값 파싱중 값 오류
const ONL_INVALID_TGT_PORT_NUMBER = _ONL_BASE + 56

const ONL_TMX_TPESVCFAIL = _ONL_BASE + 151 // 20100407 shpark 전문 값 파싱중 값 오류
const ONL_TMX_TPENOENT = _ONL_BASE + 152   // 20100407 shpark 전문 값 파싱중 값 오류
const ONL_TMX_TPESVCERR = _ONL_BASE + 153  // 20100407 shpark 전문 값 파싱중 값 오류
const ONL_TMX_TPETIME = _ONL_BASE + 154    // 20100407 shpark 전문 값 파싱중 값 오류
const ONL_TMX_TPESYSTEM = _ONL_BASE + 155  // 20100407 shpark 전문 값 파싱중 값 오류
const ONL_TMX_TPECLOSE = _ONL_BASE + 156   // 20100407 shpark 전문 값 파싱중 값 오류
const ONL_TMX_TPETRAN = _ONL_BASE + 157    // 20100407 shpark 전문 값 파싱중 값 오류
const ONL_TMX_TPEBLOCK = _ONL_BASE + 158   // 20100407 shpark 전문 값 파싱중 값 오류
const ONL_TMX_TPESVRDOWN = _ONL_BASE + 159 // 20100407 shpark 전문 값 파싱중 값 오류
const ONL_TMX_TPEQFULL = _ONL_BASE + 160   // 20100407 shpark 전문 값 파싱중 값 오류
const ONL_TMX_TPENOREADY = _ONL_BASE + 161 // 20100407 shpark 전문 값 파싱중 값 오류
const ONL_TMX_TPELIMIT = _ONL_BASE + 162   // 20100407 shpark 전문 값 파싱중 값 오류
const ONL_TMX_TPEPROTO = _ONL_BASE + 163   // 20100407 shpark 전문 값 파싱중 값 오류
const ONL_TMX_OTHER = _ONL_BASE + 164      // 20100407 shpark 전문 값 파싱중 값 오류
const ONL_TMX_TPEOS = _ONL_BASE + 165      // 20100407 shpark 전문 값 파싱중 값 오류
const ONL_TMX_TPEQPURGE = _ONL_BASE + 166  // 20100407 shpark 전문 값 파싱중 값 오류

const ONL_TUX_TPESVCFAIL = _ONL_BASE + 251      // 20100407 shpark 전문 값 파싱중 값 오류
const ONL_TUX_TPENOENT = _ONL_BASE + 252        // 20100407 shpark 전문 값 파싱중 값 오류
const ONL_TUX_TPESVCERR = _ONL_BASE + 253       // 20100407 shpark 전문 값 파싱중 값 오류
const ONL_TUX_TPETIME = _ONL_BASE + 254         // 20100407 shpark 전문 값 파싱중 값 오류
const ONL_TUX_TPESYSTEM = _ONL_BASE + 255       // 20100407 shpark 전문 값 파싱중 값 오류
const ONL_TUX_TPECLOSE = _ONL_BASE + 256        // 20100407 shpark 전문 값 파싱중 값 오류
const ONL_TUX_TPETRAN = _ONL_BASE + 257         // 20100407 shpark 전문 값 파싱중 값 오류
const ONL_TUX_TPEBLOCK = _ONL_BASE + 258        // 20100407 shpark 전문 값 파싱중 값 오류
const ONL_TUX_TPESVRDOWN = _ONL_BASE + 259      // 20100407 shpark 전문 값 파싱중 값 오류
const ONL_TUX_TPEQFULL = _ONL_BASE + 260        // 20100407 shpark 전문 값 파싱중 값 오류
const ONL_TUX_TPENOREADY = _ONL_BASE + 261      // 20100407 shpark 전문 값 파싱중 값 오류
const ONL_TUX_TPELIMIT = _ONL_BASE + 262        // 20100407 shpark 전문 값 파싱중 값 오류
const ONL_TUX_TPEPROTO = _ONL_BASE + 263        // 20100407 shpark 전문 값 파싱중 값 오류
const ONL_TUX_OTHER = _ONL_BASE + 264           // 20100407 shpark 전문 값 파싱중 값 오류
const ONL_TUX_TPEOS = _ONL_BASE + 265           // 20100407 shpark 전문 값 파싱중 값 오류
const ONL_TUX_TPEQPURGE = _ONL_BASE + 266       // 20100407 shpark 전문 값 파싱중 값 오류
const ONL_SVC_TYPE_NOT_MATCH = _ONL_BASE + 267  // 20100407
const ONL_TGT_OVER_CALL_LIMIT = _ONL_BASE + 268 // 20100407
const ONL_ASYNC_MAX_LIMIT = _ONL_BASE + 269     // 20130715

///
///	 Deferred Error	Code
///
const DFD_TARGET_BYTE_TRANSFORMER_ERROR = _DFD_BASE + 1 // ByteToByteTransformer 에서 문제가 발생하였습니다.
const DFD_TARGET_FTP_UPLOAD_ERROR = _DFD_BASE + 2       // FTP에 업로드	하는 도중 에러가 발생하였습니다.
//DB 소스
const DFD_SOURCE_APPLICATION_ERROR = _DFD_BASE + 3 // 소스	application	정보를 가져오는데 실패하였습니다.
const DFD_SOURCE_SQL_PARSE_ERROR = _DFD_BASE + 4   // 조회	SQL	문을 파싱하는데	문제가 발생하였습니다.
const DFD_SOURCE_COLLECT_ERROR = _DFD_BASE + 5     // 조회작업을 실행하는 도중에 에러가 발생하였습니다.
/**
 * @changed 2009-06-03
 */
const DFD_SOURCE_COUNT_ERROR = _DFD_BASE + 24 // 데이터 건수 조회 도중에 에러가 발생하였습니다.
const DFD_TGT_TIMEOUT_ERROR = _DFD_BASE + 25  // 멀티 sequential async 거래에서 commit 테이블의 체크 타임아웃
const DFD_WAIT_TIMEOUT_ERROR = _DFD_BASE + 26 // async 거래에서 commit 테이블의 체크 타임아웃
const DFD_RESULT_ERROR = _DFD_BASE + 27       // async 거래에서 commit 테이블의 체크 오류 발생. DFD MGR
// DFD 동작	관련 에러
const DFD_DB_CONNECT_ERROR = _DFD_BASE + 6       // 디퍼드정보DB	를 연결하는데 실패하였습니다.
const DFD_SRC_SAVE_ERROR = _DFD_BASE + 7         // 소스	정보를 디퍼드정보DB	에 저장하는데 실패하였습니다.
const DFD_TGT_SAVE_ERROR = _DFD_BASE + 8         // 타겟	정보를 디퍼드정보DB	에 저장하는데 실패하였습니다.
const DFD_RULE_REMOVE_ERROR = _DFD_BASE + 9      // 디퍼드정보DB에서	룰을 삭제하는데	실패하였습니다.
const DFD_SRC_GET_ERROR = _DFD_BASE + 10         //	소스 정보를	디퍼드정보DB 에서 가져오는데 실패하였습니다.
const DFD_TFT_GET_ERROR = _DFD_BASE + 11         //	타겟 정보를	디퍼드정보DB 에서 가져오는데 실패하였습니다.
const DFD_ROUTING_ERROR = _DFD_BASE + 12         //	요청을 보내는데	실패하였습니다.
const DFD_SRC_FILE_NOT_FOUND = _DFD_BASE + 13    //	리턴결과파일을 찾지	못했습니다.	이 룰의	패턴이 Async 인지 체크하신후 룰의 패턴을 변경해	주세요.
const DFD_TABLENAME_NOT_FOUND = _DFD_BASE + 14   //	TableNameManager  클래스를 생성하는	도중 에러가	발생하였습니다.
const DFD_TABLENAME_CALL_ERROR = _DFD_BASE + 15  //	TableNameManager 의	getTableName 메소드	호출시 에러가 발생하였습니다.
const DFD_ROLLBACK_ERROR = _DFD_BASE + 16        //	롤백 작업중에 에러가 발생하였습니다.
const DFD_TABLENAME_PARSE_ERROR = _DFD_BASE + 17 //	SQL	문에 #TABLE_NAME# 이 존재하지만	실제 TableNameManaager에서 잘못된 값을 리턴하였습니다.
// DB 타겟
const DFD_TARGET_APPLICATION_ERROR = _DFD_BASE + 18           //	타겟 applciation 정보를	가져오는데 실패하였습니다.
const DFD_TARGET_FILE_NOT_FOUND = _DFD_BASE + 19              //	조회된 데이터 파일이 존재하지 않습니다.
const DFD_TARGET_SQL_TRANSFORMER_ERROR = _DFD_BASE + 20       //	SQLTransformer에서 문제가 발생하였습니다.
const DFD_TARGET_COMMON_PROCESS_ERROR = _DFD_BASE + 21        //	데이터 저장	작업시 문제가 발생하였습니다.
const DFD_TARGET_ROLLBACK_ALL_ERROR = _DFD_BASE + 22          //	롤백 타입이	All	인 상태에서	데이터 저장	작업시 문제가 발생하였습니다. 이전에 모든 작업은 Rollback 됩니다.
const DFD_TARGET_COMMON_PROCESS_DETAIL_ERROR = _DFD_BASE + 23 //
const DFD_INVALID_RULE = _DFD_BASE + 24                       //

//
//	Batch Error Code
//
// Batch ( CODE	: BH )
const BAT_COMMON = _BAT_BASE
const BAT_NO_LOG_RESOURCE_KEY = _BAT_BASE + 1 //	"no	batch log resource key."
const BAT_NO_LOG_RESOURCE = _BAT_BASE + 2     //	"no	batch log resource [%S]"
const BAT_UNKNOWN_PROC_TYPE = _BAT_BASE + 3   //	"unknown process type [%S]"

// Batch Source	Action ( CODE :	BS )
const BAT_SOURCE_COMMON = _BAT_BASE + 4
const BAT_SOURCE_NO_FILE = _BAT_BASE + 5         //	"no	source file[%S]	in FTP"
const BAT_SOURCE_NO_QUERY_RESULT = _BAT_BASE + 6 //	"no	query result [%S]"

// Batch Target	Action ( CODE :	BT )
const BAT_TARGET_COMMON = _BAT_BASE + 7
const BAT_NOT_SUPPORTED_CRUD = _BAT_BASE + 8   //	"not supported CRUD	type [%S]"
const BAT_UNKNOWN_CRUD = _BAT_BASE + 9         //	"unknown CRUD type [%S]"
const BAT_TARGET_NO_DATA_FILE = _BAT_BASE + 10 //	"데이터	파일이 존재하지	않습니다.[%S]"

// Batch Task Action ( CODE	: TK )
const BAT_TASK_COMMON = _BAT_BASE + 11
const BAT_TASK_EXECUTE_QUERY_FAIL = _BAT_BASE + 12    //	"fail to execute query [%S]"
const BAT_TASK_CREATE_TEMP_FILE_FAIL = _BAT_BASE + 13 // "fail to create temp file [%S : %S]"

// Batch Log ( CODE	: BL )
const BLOG_COMMON = _BAT_BASE + 14
const BLOG_FIND_JOB_INFO_FAIL = _BAT_BASE + 15 // "fail to find Job Info	[BatchAgentID=%S]"
const BLOG_GET_JOB_INFO_FAIL = _BAT_BASE + 16  // "fail to find Job Info	[JobID=%S]"

// Job Request ( CODE :	JR )
const JOB_REQ_COMMON = _BAT_BASE + 17
const JOB_REQ_UNKNOWN_METHOD = _BAT_BASE + 18 // "unknown request method [%S]"
const JOB_REQ_ALREADY_DONE = _BAT_BASE + 19   // "ID[%S] is	already	done"
const JOB_REQ_REGIST_FAIL = _BAT_BASE + 20    // "fail to regist new Job Info. ID=%S"
const JOB_REQ_NO_JOB_ID = _BAT_BASE + 21      // "no Job ID"
const JOB_REQ_NO_JOB_INFO = _BAT_BASE + 22    // "no info. Job ID [%S]"

// Job Processing (	CODE : JP )
const JOB_PROC_COMMON = _BAT_BASE + 23
const JOB_PROC_UNKNOWN_METHOD = _BAT_BASE + 24    // "unknown processor	method [%S]"
const JOB_SEND_FAIL = _BAT_BASE + 25              // "fail to send job to PROCESSOR"
const JOB_PROC_COUNT_FAIL = _BAT_BASE + 26        // "fail to count	data [%S]"
const JOB_PROC_COLLECT_FAIL = _BAT_BASE + 27      // "fail to collect data [%S]"
const JOB_PROC_NO_SOURCE_FILE = _BAT_BASE + 28    // "doesn't exist	file [%S]"
const JOB_PROC_UPDATE_FAIL = _BAT_BASE + 29       // "fail to update data [%S]"
const JOB_PROC_TASK_FAIL = _BAT_BASE + 30         // "fail to task data	[%S]"
const JOB_PROC_FAIL = _BAT_BASE + 31              // "fail to process job[%S]."
const JOB_PROC_TASK_SERVICE_FAIL = _BAT_BASE + 32 // "fail to process SERVICE TASK [TASK:%S-%S (%S)]"
const JOB_COUNT_DATA_ERROR = _BAT_BASE + 33       // "fail to count	data[%S	!= %S]"
const JOB_GET_PARAM_FAIL = _BAT_BASE + 34         // "fail to get job parameters"
const JOB_PUT_PARAM_FAIL = _BAT_BASE + 35         // "fail to put job parameters"
const JOB_PRE_TASK_FAIL = _BAT_BASE + 36          // "fail to task data	[%S]"
const JOB_POST_TASK_FAIL = _BAT_BASE + 37         // "fail to task data	[%S]"
const JOB_COLLECT_FAIL = _BAT_BASE + 38           // "fail to task data	[%S]"
const JOB_CANCELED = _BAT_BASE + 39               // "fail to task data	[%S]"
const JOB_DIRECT_NOT_ALLOWED = _BAT_BASE + 40     // "direct is allowed only for file source type"
const BAT_SOURCE_FILE_CORRUPTED = _BAT_BASE + 41  //	"no	source file[%S]	in FTP"
///////////////////////////////////////////////////////////////
const BAT_SQL_INPUT_PARAM_NOT_MATCH = _BAT_BASE + 42 //	"input param count doesn't match query params"
const BAT_INVALID_FILE_PATH = _BAT_BASE + 43         //	"invalid file path - not allowed or wrong path"
const BAT_ROUTING_ERROR = _BAT_BASE + 44             //    요청을 보내는데    실패하였습니다.
const BAT_GET_COUNT_LIMIT = _BAT_BASE + 45           //    데이터 조회 결과가 임계값을 초과하였습니다.

const _DEPRECATED_BASE = 10000
const INTERNAL_TIMEOUT = _DEPRECATED_BASE + 1 // timeout internally
const ERRCODE_INIT_FAIL = _DEPRECATED_BASE + 2

const FETCH_FILE_PATH_ERROR = _BAT_DEF_BASE + 1

const TELNET_USER_ERROR = _TELNET_BASE + 1
const TELNET_PASSWORD_ERROR = _TELNET_BASE + 2
const TELNET_CONNECT_ERROR = _TELNET_BASE + 3
const TELNET_HOST_ERROR = _TELNET_BASE + 4
const TELNET_PROMPT_ERROR = _TELNET_BASE + 5

const LICENSE_SERVICE_MAX = _LICENSE_BASE + 1
const LICENSE_VALIDATION_FAIL = _LICENSE_BASE + 2

const SAP_IN_PARAM_COUNT_NOT_MATCH = _SAP_BASE + 1
const SAP_OUT_PARAM_COUNT_NOT_MATCH = _SAP_BASE + 2
const SAP_FUNCTION_NOT_EXIST = _SAP_BASE + 3
const SAP_GW_CONNECTION_FAIL = _SAP_BASE + 4
const SAP_FUNCTION_EXECUTION_FAIL = _SAP_BASE + 5

const WEB_INVALID_URL = _WEB_BASE + 1
const WEB_WSDL_VERSION_NOT_1_1 = _WEB_BASE + 2
const WEB_WSDL_VERSION_NOT_2_0 = _WEB_BASE + 3

const WAS_JNDI_NOT_FOUND = _WAS_BASE + 1
const WAS_TYPE_INVALID = _WAS_BASE + 2
const WAS_PROVIDER_NOT_SPECIFIED = _WAS_BASE + 3
const WAS_PROPERTIES_NOT_SPECIFIED = _WAS_BASE + 4

const WAS_CONNECITON_FAILED = _WAS_BASE + 51
