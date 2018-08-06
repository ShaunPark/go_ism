package rule

// used in rule.go
const inf_sql = "SELECT INTEGRATIONSERVICEID, ISTWOPHASECOMMIT, MESSAGEVALIDATE, INTEGRATIONTYPE, SERVICETEMPLETEID FROM INTEGRATIONSERVICE WHERE INTEGRATIONSERVICEID = ?"
const rSvc_sql = "select msglog, processlog, asyncretrycount, asyncretrytimeout from realtimeservice where integrationserviceid = ?"
const bSvc_sql = "select fetchcount, rollbacktype, waittimeout, fileseparatorclass, posttaskid, isdirect, pretaskid, checksize, binarymode from batchservice where integrationserviceid = ?"
const dSvc_sql = "select fetchinterval, datamaxseq, fetchcount, dayclosefieldname, dayclosefieldvalue, sequenceinitcondition, posttaskid, rollbacktype, sequencefield, datefield, autogenerate, geninterval, pretaskid, includeendvalue, closehandler from deferredservice where integrationserviceid = ? "
const sysEntity_sql = "select systemid, parserName, EntityOrder, timeout from SystemEntity where IntegrationServiceId = ? order by EntityOrder"

// used in data.go
const dStrt_sql = "select datastructureid, datastructurename, datastructuretype from datastructure where datastructureid = ?"
const data_sql = "select dataindex, recorddelimeter, fieldgroupid from data where datastructureid = ? order by dataIndex"
const detail_sql = "select fieldgroupid, repeatcount, repeatcountdataindex, repeatcountfieldindex, grouptype from detail where DataStructureID = ? and dataIndex = ? order by detailIndex"
const length_sql = " select lengthDataIndex, lengthDetailIndex, lengthFieldIndex, dataIndex, detailIndex, fieldIndex, diffvalue from LengthFieldInfo where DatastructureId = ? order by dataIndex, detailIndex, fieldIndex"

// used in field.go
const fld_sql = "select fieldid, fieldname, fieldtype, fieldlength, fieldformat, fillchar, aligntype from field where fieldid = ?"
const flg_sql = "select FieldGroupID, FIELDGROUPNAME, FieldDelimeter from fieldgroup where FieldGroupID = ?"
const flgMap_sql = "select fieldindex, fieldid, fieldoffset, diffvalue, iskey, isnull, issqlfunction, LengthFIeldType, inOutType, filterType from fieldgroupmap where FieldGroupid = ? order by fieldIndex "

//used in system.go
const app_sql = "select ApplicationID, ApplicationName, portNumber, AppUserID, AppUserPwd,  applicationType from application where ApplicationID = ?"
const svr_sql = "select ServerId, ServerName, ipAddress from server where ServerId = ?"
const sys_sql = "select systemID, SystemName, LoadBalanceType, useHealth, autoRecovery, checkGid, callLimit from AppSystem where systemID = ?"
const dbApp_sql = "select dbname, jdbctype, dbtype, dbparameter, appsystemname " +
	"from DBApplication where ApplicationID = ?"

// used in service.go
const svc_sql = "select serviceid, servicename, serviceblockflag, inputdatastructureid, outputdatastructureid, " +
	"errordatastructureid, servicetype, interfaceType, SENDRECEIVETYPE from Service where serviceid = ?"
const dbSvc_sql = "select crudtype, tablename, query, querytype, spType, defaultCrudType, filterCheckType, targetFetchCount " +
	"from DBService where serviceid = ?"

// used in svcModel.go
const sModel_sql = "select ServiceTempleteId, returnType, returnIndex, isConcurrent from ServiceTemplete where ServiceTempleteId = ?"
const svcE_sql = "select serviceId, syncType, RoutingMatchMethod, EntityOrder, errorMappingId, servicetype " +
	"from ServiceEntity where ServiceTempleteId = ? order by EntityOrder"
const svcR_sql = "select targetServiceIndex, MappingId, isDefault, routePattern " +
	"from ServiceRoute where ServiceTempleteId = ? and EntityOrder = ? order by RouteOrder"
const rInput_sql = "select sourcedataIndex, sourcedetailIndex, sourcefieldIndex, value, fieldId, path " +
	" from RouteInput where ServiceTempleteId = ? and EntityOrder = ? order by InputIndex"

// used in svcMap.go
const sMap_sql = "select outputServiceId, DataStructureId from ServiceMappingRule where MappingId = ?"
const dMap_sql = "select dataindex, detailindex, columnindex, fieldid, customfunction from DataMapping where mappingid = ? order by dataindex, detailindex, columnindex "
const sCol_sql = "select mappingindex, sourcedataindex, sourcedetailindex, sourcefieldindex, defaultvalue, fieldid, sourcemessageindex , path, isXml " +
	"from SourceColumns where MAPPINGID = ? and DATAINDEX = ? and DETAILINDEX = ? and COLUMNINDEX = ? order by mappingindex"
const inSvc_sql = "select serviceId, dataStructureId from InputServices  where MappingId = ? order by ServiceMappingIndex"
