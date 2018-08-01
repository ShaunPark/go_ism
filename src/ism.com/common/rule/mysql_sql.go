package rule

const inf_sql = "SELECT INTEGRATIONSERVICEID, ISTWOPHASECOMMIT, MESSAGEVALIDATE, INTEGRATIONTYPE, SERVICETEMPLETEID FROM INTEGRATIONSERVICE WHERE INTEGRATIONSERVICEID = ?"
const rSvc_sql = "select msglog, processlog, asyncretrycount, asyncretrytimeout from realtimeservice where integrationserviceid = ?"
const bSvc_sql = "select fetchcount, rollbacktype, waittimeout, fileseparatorclass, posttaskid, isdirect, pretaskid, checksize, binarymode from batchservice where integrationserviceid = ?"
const dSvc_sql = "select fetchinterval, datamaxseq, fetchcount, dayclosefieldname, dayclosefieldvalue, sequenceinitcondition, posttaskid, rollbacktype, sequencefield, datefield, autogenerate, geninterval, pretaskid, includeendvalue, closehandler from deferredservice where integrationserviceid = ? "
const sysEntity_sql = "select systemid, parserName, EntityOrder, timeout from SystemEntity where IntegrationServiceId = ? order by EntityOrder"

const dStrt_sql = "select datastructureid, datastructurename, datastructuretype from datastructure where datastructureid = ?"
const data_sql = "select dataindex, recorddelimeter, fieldgroupid from data where datastructureid = ? order by dataIndex"
const detail_sql = "select fieldgroupid, repeatcount, repeatcountdataindex, repeatcountfieldindex, grouptype from detail where DataStructureID = ? and dataIndex = ? order by detailIndex"
const length_sql = " select lengthDataIndex, lengthDetailIndex, lengthFieldIndex, dataIndex, detailIndex, fieldIndex, diffvalue from LengthFieldInfo where DatastructureId = ? order by dataIndex, detailIndex, fieldIndex"

const fld_sql = "select fieldid, fieldname, fieldtype, fieldlength, fieldformat, fillchar, aligntype from field where fieldid = ?"
const flg_sql = "select FieldGroupID, FIELDGROUPNAME, FieldDelimeter from fieldgroup where FieldGroupID = ?"
const flgMap_sql = "select fieldindex, fieldid, fieldoffset, diffvalue, iskey, isnull, issqlfunction, LengthFIeldType, inOutType, filterType from fieldgroupmap where FieldGroupid = ? order by fieldIndex "

const app_sql = "select ApplicationID, ApplicationName, portNumber, AppUserID, AppUserPwd,  applicationType from application where ApplicationID = ?"
const svr_sql = "select ServerId, ServerName, ipAddress from server where ServerId = ?"
const sys_sql = "select systemID, SystemName, LoadBalanceType, useHealth, autoRecovery, checkGid, callLimit from AppSystem where systemID = ?"

const dbApp_sql = "select dbname, jdbctype, dbtype, dbparameter, appsystemname from DBApplication where ApplicationID = ?"
