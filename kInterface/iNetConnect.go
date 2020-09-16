package kInterface

type INetConnect interface {
	Config(name, ip, port string)

	Start()

	Stop()

	AcceptConnect()
}
