package types

// msg-gateway 用到的
const (
	WSGetNewestSeq     = 1001
	WSPullMsgBySeqList = 1002
	WSSendMsg          = 1003
	WSSendSignalMsg    = 1004
)

// msg 用到的
const (
	SingleChatType       = 1
	GroupChatType        = 2
	NotificationChatType = 4

	OnlineStatus  = "online"
	OfflineStatus = "offline"
	Registered    = "registered"
	UnRegistered  = "unregistered"

	// 群聊的
	GroupCreatedNotification             = 1501
	GroupInfoSetNotification             = 1502
	JoinGroupApplicationNotification     = 1503
	MemberQuitNotification               = 1504
	GroupApplicationAcceptedNotification = 1505
	GroupApplicationRejectedNotification = 1506
	GroupOwnerTransferredNotification    = 1507
	MemberKickedNotification             = 1508
	MemberInvitedNotification            = 1509
	MemberEnterNotification              = 1510
	GroupDismissedNotification           = 1511
	GroupMemberMutedNotification         = 1512
	GroupMemberCancelMutedNotification   = 1513
	GroupMutedNotification               = 1514
	GroupCancelMutedNotification         = 1515
	GroupMemberInfoSetNotification       = 1516

	///消息类型
	Text           = 101
	Picture        = 102
	Voice          = 103
	Video          = 104
	File           = 105
	AtText         = 106
	Merger         = 107
	Card           = 108
	Location       = 109
	Custom         = 110
	Revoke         = 111
	HasReadReceipt = 112
	Typing         = 113
	Quote          = 114
	Common         = 200
	GroupMsg       = 201
)

const (
	AtAllString = "AtAllTag"
	AtNormal    = 0
	AtMe        = 1
	AtAll       = 2
	AtAllAtMe   = 3
)

// options
const (
	//OptionsKey
	IsHistory                  = "history"
	IsPersistent               = "persistent"
	IsOfflinePush              = "offlinePush"
	IsUnreadCount              = "unreadCount"
	IsConversationUpdate       = "conversationUpdate"
	IsSenderSync               = "senderSync"
	IsNotPrivate               = "notPrivate"
	IsSenderConversationUpdate = "senderConversationUpdate"
)
