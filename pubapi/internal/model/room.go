package model

type Room struct {
	Id         uint64 // 房间ID
	Cap        int32  // 房间容量
	Creator    uint64 // 房间创建者ID
	CreateTime uint64 // 创建时间
	UpdateTime uint64 // 更新时间
}

type RoomLog struct {
	Id         uint64 // 日志ID
	UserId     uint64 // 用户ID
	Value      string // 操作
	CreateTime uint64 // 创建时间
}

type RoomOperation int

const (
	RoomOperation_NONE RoomOperation = iota
	RoomOperation_OPEN               // 开启房间
	RoomOperation_JOIN               // 进入房间
	RoomOperation_EXIT               // 退出房间
	RoomOperation_OVER               // 结束房间
)

func NewRoomOperationWithValue(value string) RoomOperation {
	var o RoomOperation
	switch value {
	case "OPEN":
		o = RoomOperation_OPEN
	case "JOIN":
		o = RoomOperation_JOIN
	case "EXIT":
		o = RoomOperation_EXIT
	case "OVER":
		o = RoomOperation_OVER
	}
	return o
}

func NewRoomOperation(enum int) RoomOperation {
	return RoomOperation(enum)
}

func (o RoomOperation) String() (value string) {
	switch o {
	case RoomOperation_OPEN:
		value = "开启房间"
	case RoomOperation_JOIN:
		value = "加入房间"
	case RoomOperation_EXIT:
		value = "退出房间"
	case RoomOperation_OVER:
		value = "解散房间"
	default:
		value = "异常操作"
	}
	return value
}
