package rule

const inf_sql = "SELECT INTEGRATIONSERVICEID, ISTWOPHASECOMMIT, MESSAGEVALIDATE, INTEGRATIONTYPE, SERVICETEMPLETEID FROM INTEGRATIONSERVICE WHERE INTEGRATIONSERVICEID = ?"
const rSvc_sql = "select msglog, processlog, asyncretrycount, asyncretrytimeout from realtimeservice where integrationserviceid = ?"
const bSvc_sql = "select fetchcount, rollbacktype, waittimeout, fileseparatorclass, posttaskid, isdirect, pretaskid, checksize, binarymode from batchservice where integrationserviceid = ?"
const dSvc_sql = "select fetchinterval, datamaxseq, fetchcount, dayclosefieldname, dayclosefieldvalue, sequenceinitcondition, posttaskid, rollbacktype, sequencefield, datefield, autogenerate, geninterval, pretaskid, includeendvalue, closehandler from deferredservice where integrationserviceid = ? "

const dStrt_sql = "select * from datastructure where datastructureid = ?"
