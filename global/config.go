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
	RAdmin = 1 + iota
	RMember
	RVisitor
)
