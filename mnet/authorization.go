package mnet

//使用2进制int来表示权限
func CheckAuth(a , auth int) bool {
	if a & auth == 0 {
		return false
	} else {
		return true
	}
}
