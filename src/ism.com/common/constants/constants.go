package constants

const PARAMTYPE_FIELD = 0
const PARAMTYPE_RULE = 1
const PARAMTYPE_CONSTANT = 2

const ONLINE_INF = 0
const BATCH_INF = 1
const DEFERRED_INF = 2

const ONLINE_APP = 0
const DB_APP = 1
const FILE_APP = 2
const JMS_APP = 3
const EJB_APP = 4
const WEB_APP = 5
const SAP_APP = 6

const ETC_APP = 99

const KEY_NOTKEY = 0
const KEY_NORMALKEY = 1
const KEY_UPDATEKEY = 2

const M_ALL = 0
const M_VALUE = 1
const M_CLASS = 2
const M_PATTERN = 3

const S_SYNC = 0
const S_ASYNC = 1
const S_ASYNC_ACK = 2

const NO_FIELD = -2
const VARIABLE_LENGTH = -1

const INF_PKG = "com.fiss.rule.repository.action."
const INF_ISM_PKG = "com.ism.rule.repository.action."
const INF_INF = INF_ISM_PKG + "InterfaceAction"
const INF_APP = INF_ISM_PKG + "ApplicationAction"
const INF_APS = INF_ISM_PKG + "AppSystemAction"
const INF_DAT = INF_ISM_PKG + "DataStructureAction"
const INF_FLD = INF_ISM_PKG + "FieldAction"
const INF_FLG = INF_ISM_PKG + "FieldGroupAction"
const INF_MAP = INF_ISM_PKG + "DataMappingAction"
const INF_SVR = INF_ISM_PKG + "ServerAction"
const INF_SVC = INF_ISM_PKG + "ServiceAction"
const INF_HDR = INF_ISM_PKG + "HeaderAction"
const INF_HDM = INF_ISM_PKG + "HeaderMappingAction"
const INF_TSK = INF_ISM_PKG + "TaskAction"
const INF_ERRCODE = INF_ISM_PKG + "ErrorCodeAction"
const INF_ERRMSG = INF_ISM_PKG + "ErrorMessageAction"
const INF_COMMON = INF_ISM_PKG + "CommonProcessAction"

const INF_ISVC = INF_ISM_PKG + "IntegrationServiceAction"
const INF_SVCT = INF_ISM_PKG + "ServiceTempleteAction"
const INF_SVCMAP = INF_ISM_PKG + "ServiceMappingRuleAction"

const ALIGN_LEFT = "L"
const ALIGN_RIGHT = "R"

const HEALTH_DOWN = 0
const HEALTH_RUN = 1
const HEALTH_BACKUP = 2
const HEALTH_CHECK = 3
const HEALTH_NO_USE = 4

const LB_ROUNDROBIN = 0
const LB_DEDICATE = 1
const LB_FAILOVER = 2
const LB_MASTER = 3
const LB_BY_RULE = 4

const RETURN_SINGLE = 0
const RETURN_MULTI = 1
const RETURN_NONE = 2

const PRESQL_STR_START = "@PRESQL["
const PRESQL_STR_END = "]PRESQL@"

const CRUD_SKIP = -1
const CRUD_CREATE = 0             // insert "C"
const CRUD_UPDATE = 1             // update "U"
const CRUD_DELETE = 2             // delete "D"
const CRUD_READ = 3               // select "R"
const CRUD_UPDATE_INSERT = 4      // insert if update failed "P"
const CRUD_INSERT_UPDATE = 5      // update if insert failed "N"
const CRUD_USE_RULE = 7           //
const CRUD_INSERT_SKIP_DUP = 8    // Insert(Skip duplicate) "K"
const CRUD_DELETE_IGNORE = 9      // delete(skip count 0) "I"
const CRUD_UPDATE_IGNORE = 10     // update(skip count 0) "E"
const CRUD_USER_SQL = 11          // user sql "Q"
const CRUD_USER_SQL_SKIP_DUP = 12 // user sql(skip duplicate) "L"
const CRUD_DYNAMIC = 13           // dynamic "Y"
const CRUD_DYNAMIC_IGNORE = 14    // dynamic (Skip Error) "Z"

const DBTYPE_ORACLE = 0
const DBTYPE_DB2 = 1
const DBTYPE_INFOMIX = 2
const DBTYPE_MSSQL = 3
const DBTYPE_SYBASE = 4
const DBTYPE_TERADATA = 5
const DBTYPE_MYSQL = 6
const DBTYPE_DB2400 = 7
const DBTYPE_ALTIBASE = 8
const DBTYPE_TIBERO = 9
const DBTYPE_DB2_UNIVERSAL = 10
const DBTYPE_CUSTOM = 99

const QUERY_LITERAL = 0
const QUERY_PREPARED = 1

const ROLLBACK_ALL = 0
const ROLLBACK_ROW = 1

const SPTYPE_SP = 0
const SPTYPE_NORMAL = 0

const SVC_TYPE_ONLINE = 0
const SVC_TYPE_DB = 1
const SVC_TYPE_FILE = 2
const SVC_TYPE_M3O = 3
const SVC_TYPE_WS = 4
const SVC_TYPE_JMS = 5
const SVC_TYPE_EJB = 6
const SVC_TYPE_SAP = 7
const SVC_TYPE_MQ = 8
const SVC_TYPE_BIZ = 9

// M, A, J, W
const JMS_TYPE_MQ = 1 //Websphere
const JMS_TYPE_ACTIVEMQ = 2
const JMS_TYPE_JBOSSMQ_5 = 3
const JMS_TYPE_WEBLOGICMQ = 4
const JMS_TYPE_JBOSSMQ_7 = 5
const JMS_TYPE_SONICMQ = 6
const JMS_TYPE_WEBSPHEREMQ = 7

//Q, T
const JMS_SEND_QUEUE = 1 //Q
const JMS_SEND_TOPIC = 2 //T

//T, O, B, M
const JMS_MESSAGE_TEXT = 1
const JMS_MESSAGE_OBJECT = 2
const JMS_MESSAGE_BYTEARRAY = 3
const JMS_MESSAGE_MAP = 4

// J = jboss, L = weblogic, W = websphere
const EJB_TYPE_JBOSS = 3
const EJB_TYPE_WEBLOGIC = 4
const EJB_TYPE_WEBSPHERE = 5

/**
 * 2011-11-12 for sap field type. used in fieldgroup
 */
const SAP_TYPE_FIELD = 1       //simple field type parameters
const SAP_TYPE_STRUCT = 2      //struct embedded as a field
const SAP_TYPE_TABLE = 3       //table parameter independent from import/export, can be used for both.
const SAP_TYPE_FIELD_TABLE = 4 //table embedded as a table with repeatable records.

const MQ_MSG_TYPE_OBJECT = 1 //O
const MQ_MSG_TYPE_STRING = 2 //S
const MQ_MSG_TYPE_UTF = 3    //U
const MQ_MSG_TYPE_BYTE = 4   //B

const MQ_SEND_TYPE_MQI = 1 //N
const MQ_SEND_TYPE_JMS = 2 //J

const TYPE_CHAR = "C"
const TYPE_NUMBER = "N"
const TYPE_DATE = "D"
const TYPE_BINARY = "B"
const TYPE_NCHAR = "I"

const FTP_NONE = 0
const FTP_OVERWRITE = 1
const FTP_APPEND = 2
const FTP_OVERWRITE_APPEND = 3

/**
 * target file size check. if size is same, error, different and smaller than local, assume the file is not correct.
 */
const FTP_NONE_CHECK = 4

/**
 * add retrieve
 */
const FTP_RESTART = 5

/**
 * @changed 2011-02-20 ftp protocol added
 */
const FTP_PROTOCOL_FTP = 0
const FTP_PROTOCOL_SFTP = 1
const FTP_PROTOCOL_SCP = 2
const FTP_PROTOCOL_FTPS = 3
const FTP_PROTOCOL_TANDEM = 7
const FTP_PROTOCOL_SFTP_WIN = 8
const FTP_PROTOCOL_CUSTOM = 9

const HTTP_PROTOCOL = 0
const HTTPS_PROTOCOL = 1

// InterfaceInfo Type ------------------------------------------------------
const INTF_TYPE_ONLINE = 0
const INTF_TYPE_BATCH = 1
const INTF_TYPE_DEFERRED = 2

const INTF_TYPE_ONLINE_S = "O"
const INTF_TYPE_BATCH_S = "B"
const INTF_TYPE_DEFERRED_S = "D"

// WAS Context Type --------------------------------------------------------
const CONTEXT_TYPE_DEFAULT = "default"
const CONTEXT_TYPE_BAT = "batch"
const CONTEXT_TYPE_DFD = "deferred"
const CONTEXT_TYPE_ONL = "realtime"
const CONTEXT_TYPE_CONSOLE = "webconsole"
const CONTEXT_TYPE_LOG_ONL = "realtime-log"
const CONTEXT_TYPE_DFD_SCHE = "deferred-scheduler"

// Task Type ---------------------------------------------------------------
const PRE_TASK = 0
const POST_TASK = 1

const TASK_TYPE_TABLE = 0
const TASK_TYPE_STORED = 1
const TASK_TYPE_FILE = 2
const TASK_TYPE_SERVICE = 3

/**
 * @changed 2008-02-13 telnet, custom task
 */
const TASK_TYPE_TELNET = 4
const TASK_TYPE_CUSTOM = 5

const R_NOT = 0
const R_AND = 1
const R_OR = 2

const TOTAL_LENGTH = "T"
const PARTIAL_LENGTH = "P"
const LENGTH_NOT = "N"
const FIELDGROUP_LENGTH = "F"
const DETAIL_REPEAT_COUNT = "R"

const PRT_STRING = 0
const PRT_INTEGER = 1
const PRT_BOOLEAN = 2

const INOUTTYPE_IN = 0
const INOUTTYPE_OUT = 1
const INOUTTYPE_INOUT = 2

/**
 * @changed 2008-02-28 jdbc type
 *
 */
const JDBC_TYPE_1 = 1
const JDBC_TYPE_2 = 2
const JDBC_TYPE_3 = 3
const JDBC_TYPE_4 = 4
const JDBC_TYPE_FAILOVER = 5

// Log Level
const LOG_ERROR = 0
const LOG_WARNING = 1
const LOG_NORMAL = 2
const LOG_VERBOSE = 3
const LOG_TRIVIA = 4

const JDBM_APP = "application"
const JDBM_APPSYS = "appsystem"
const JDBM_DATAMAP = "map"
const JDBM_DATASTRT = "data"
const JDBM_ERRCODE = "errcode"
const JDBM_ERRMSG = "errmsg"
const JDBM_FLG = "flg"
const JDBM_FLD = "fld"
const JDBM_HDR = "hdr"
const JDBM_HDRMAP = "hdrmap"
const JDBM_INF = "interface"
const JDBM_SVC = "svc"
const JDBM_SVR = "svr"
const JDBM_TASK = "task"
const JDBM_COMMON = "common"

/////// ISM 2.2
const JDBM_ISVC = "integrationService"
const JDBM_SVCT = "serviceTemplete"
const JDBM_SVCMAP = "serviceMapping"

const COM_SRC_TYPE = 0
const COM_TGT_TYPE = 1
const COM_BOTH_TYPE = 2
const COM_INF_DB = 0
const COM_INF_FILE = 1
const COM_INF_SOCKET = 2
const COM_INF_WEB = 3
const COM_INF_OTHER = 4

const NOT_FILTER = 0
const IN_FILTER = 1
const EX_FILTER = 2
const REQUEST = "1"
const SEND = "0"

const LONG_MESSAGE_NOT = 1
const LONG_MESSAGE_START = 2
const LONG_MESSAGE_BODY = 3
const LONG_MESSAGE_END = 4

const BAT_FILE_DIRECT_AGENT = "A"
const BAT_FILE_DIRECT_EAI = "D"
const BAT_FILE_INDIRECT = "E"

const FILE_SENTFILE_NOJOB = 0
const FILE_SENTFILE_MOVE = 1
const FILE_SENTFILE_DELETE = 2

const BAT_FILE_DIRECT_AGENT_INT = 2
const BAT_FILE_DIRECT_EAI_INT = 1
const BAT_FILE_INDIRECT_INT = 0

const KEY_ENCODED_STR = "<!ENCODED!>"

const MESSAGE_SAP = 3
const MESSAGE_XML = 2
const MESSAGE_ISO = 1
const MESSAGE_ISM = 0
