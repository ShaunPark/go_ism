package gid

type GidChecker struct {
}
type GidCheckerInterface interface {
	CheckGID(gid string) bool
}
