package miraiGo

const (
	EventReceiveFriendMessage            = "FriendMessage"
	EventReceiveGroupMessage             = "GroupMessage"
	EventReceiveTempMessage              = "TempMessage"
	EventBotOnline                       = "BotOnlineEvent"
	EventBotOfflineActive                = "BotOfflineEventActive"
	EventBotOfflineForce                 = "BotOfflineEventForce"
	EventBotOfflineDropped               = "BotOfflineEventDropped"
	EventBotRelogin                      = "BotReloginEvent"
	EventGroupRecall                     = "GroupRecallEvent"
	EventFriendRecall                    = "FriendRecallEvent"
	EventBotGroupPermissionChange        = "BotGroupPermissionChangeEvent"
	EventBotMute                         = "BotMuteEvent"
	EventBotUnmute                       = "BotUnmuteEvent"
	EventBotJoinGroup                    = "BotJoinGroupEvent"
	EventBotLeaveActive                  = "BotLeaveEventActive"
	EventBotLeaveKick                    = "BotLeaveEventKick"
	EventGroupNameChange                 = "GroupNameChangeEvent"
	EventGroupEntranceAnnouncementChange = "GroupEntranceAnnouncementChangeEvent"
	EventGroupMuteAll                    = "GroupMuteAllEvent"
	EventGroupAllowAnonymousChat         = "GroupAllowAnonymousChatEvent"
	EventGroupAllowConfessTalk           = "GroupAllowConfessTalkEvent"
	EventGroupAllowMemberInvite          = "GroupAllowMemberInviteEvent"
	EventMemberJoin                      = "MemberJoinEvent"
	EventMemberLeaveKick                 = "MemberLeaveEventKick"
	EventMemberLeaveQuit                 = "MemberLeaveEventQuit"
	EventMemberCardChange                = "MemberCardChangeEvent"
	EventMemberSpecialTitleChange        = "MemberSpecialTitleChangeEvent"
	EventMemberPermissionChange          = "MemberPermissionChangeEvent"
	EventMemberMute                      = "MemberMuteEvent"
	EventMemberUnmute                    = "MemberUnmuteEvent"
	EventNewFriendRequest                = "NewFriendRequestEvent"
	EventMemberJoinRequest               = "MemberJoinRequestEvent"
)

type Event struct {
	Type string `json:"type"` //事件类型
	MessageChain []Message `json:"messageChain"` //(ReceiveMessage)消息链
	Sender Sender `json:"sender"` //(ReceiveMessage)发送者信息
	//QQ uint `json:"qq"` //(BotOnline)
}