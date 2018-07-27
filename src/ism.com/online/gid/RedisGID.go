package gid

type RedisChecker struct {
  GidChecker
}

func (gidChecker *RedisChecker) CheckGID(gid string) bool {
  return false
}
