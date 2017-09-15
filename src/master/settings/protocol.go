package settings

const (
	System_Msg_S2C     = 10000
	System_Heart_C2S   = 10001
	System_Heart_S2C   = 10002
	System_Message_S2C = 10003
	System_Ping        = 10004
	System_Pong        = 10005
)

const (
	Cell_Login_C2S  = 10100
	Cell_Login_S2C  = 10101
	Cell_Logout_C2S = 10102

	Cell_Fake_S2C = 10103
	Cell_Fake_C2S = 10104

	Cell_Stop_S2C = 10105
	Cell_Stop_C2S = 10106
)

const (
	Task_Create_S2C   = 10200
	Task_Create_C2S   = 10201
	Task_Progress_S2C = 10201
	Task_Progress_C2S = 10202
	Task_Report_C2S   = 10203
)
