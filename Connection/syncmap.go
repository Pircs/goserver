package connection

import (
	"net"
	"sync"
)

var nAutoIncrease int
var mpConnection sync.Map

func init() {
	nAutoIncrease = 0
}

// CreateConnection 创建一个新的连接
func CreateConnection(pTCPConn *net.TCPConn) *TConnection {
	nAutoIncrease++
	pConnection := &TConnection{}
	pConnection.nIndex = nAutoIncrease
	pConnection.pTCPConn = pTCPConn
	mpConnection.Store(nAutoIncrease, pConnection)
	return pConnection
}

// FindConnection 查找
func FindConnection(nIndex int) *TConnection {
	v, ok := mpConnection.Load(nIndex)
	if ok {
		return v.(*TConnection)
	}
	return nil
}

func deleteConnection(nIndex int) {
	mpConnection.Delete(nIndex)
}
