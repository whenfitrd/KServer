package global

const (
	MyMessage = 1 + iota
)

const (
	MyMsgLen = 9
	MsgInfoLen = 8
)

const (
	MMsgHead = "use json"
)

//游戏组权限
const (
	Admin = 1 + iota
	Member
)

//路由权限
const (
	RVisitor = 1
	RMember = 2
	RAdmin = 4
	RAll = 7
)

//模块
const (
	LoginModule = 1 + iota
)
