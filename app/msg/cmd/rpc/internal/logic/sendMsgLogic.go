package logic

import (
	"context"
	"github.com/golang/protobuf/proto"
	conversationpb "github.com/showurl/Zero-IM-Server/app/conversation/cmd/rpc/pb"
	imuserpb "github.com/showurl/Zero-IM-Server/app/im-user/cmd/rpc/pb"
	chatpb "github.com/showurl/Zero-IM-Server/app/msg/cmd/rpc/pb"
	"github.com/showurl/Zero-IM-Server/common/types"
	"github.com/showurl/Zero-IM-Server/common/utils"
	strUtils "github.com/showurl/Zero-IM-Server/common/utils/str"
	"sync"

	"github.com/showurl/Zero-IM-Server/app/msg/cmd/rpc/internal/svc"
	"github.com/showurl/Zero-IM-Server/app/msg/cmd/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type SendMsgLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewSendMsgLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SendMsgLogic {
	return &SendMsgLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *SendMsgLogic) SendMsg(pb *pb.SendMsgReq) (*pb.SendMsgResp, error) {
	replay := chatpb.SendMsgResp{}
	flag, errCode, errMsg := l.userRelationshipVerification(pb)
	if !flag {
		return returnMsg(&replay, pb, errCode, errMsg, "", 0)
	}
	//if !utils.VerifyToken(pb.Token, pb.SendID) {
	//	return returnMsg(&replay, pb, http.StatusUnauthorized, "token validate err,not authorized", "", 0)
	l.encapsulateMsgData(pb.MsgData)
	logx.WithContext(l.ctx).Info("this is a test MsgData ", pb.MsgData)
	msgToMQSingle := chatpb.MsgDataToMQ{Token: pb.Token, OperationID: pb.OperationID, MsgData: pb.MsgData}
	//options := utils.JsonStringToMap(pbData.Options)
	isHistory := utils.GetSwitchFromOptions(pb.MsgData.Options, types.IsHistory)
	mReq := MsgCallBackReq{
		SendID:      pb.MsgData.SendID,
		RecvID:      pb.MsgData.RecvID,
		Content:     string(pb.MsgData.Content),
		SendTime:    pb.MsgData.SendTime,
		MsgFrom:     pb.MsgData.MsgFrom,
		ContentType: pb.MsgData.ContentType,
		SessionType: pb.MsgData.SessionType,
		PlatformID:  pb.MsgData.SenderPlatformID,
		MsgID:       pb.MsgData.ClientMsgID,
	}
	if !isHistory {
		mReq.IsOnlineOnly = true
	}

	// callback
	canSend, err := l.callbackWordFilter(pb)
	if err != nil {
		logx.WithContext(l.ctx).Error(utils.GetSelfFuncName(), "callbackWordFilter failed", err.Error(), pb.MsgData)
	}
	if !canSend {
		return returnMsg(&replay, pb, 201, "callbackWordFilter result stop rpc and return", "", 0)
	}
	switch pb.MsgData.SessionType {
	case types.SingleChatType:
		// callback
		canSend, err := l.callbackBeforeSendSingleMsg(pb)
		if err != nil {
			logx.WithContext(l.ctx).Error(utils.GetSelfFuncName(), "callbackBeforeSendSingleMsg failed", err.Error())
		}
		if !canSend {
			return returnMsg(&replay, pb, 201, "callbackBeforeSendSingleMsg result stop rpc and return", "", 0)
		}
		isSend := l.modifyMessageByUserMessageReceiveOpt(pb.MsgData.RecvID, pb.MsgData.SendID, types.SingleChatType, pb)
		if isSend {
			msgToMQSingle.MsgData = pb.MsgData
			logx.WithContext(l.ctx).Info(msgToMQSingle.OperationID, msgToMQSingle)
			err1 := l.sendMsgToKafka(&msgToMQSingle, msgToMQSingle.MsgData.RecvID, types.OnlineStatus)
			if err1 != nil {
				logx.WithContext(l.ctx).Error(msgToMQSingle.OperationID, "kafka send msg err:RecvID", msgToMQSingle.MsgData.RecvID, msgToMQSingle.String())
				return returnMsg(&replay, pb, 201, "kafka send msg err", "", 0)
			}
		}
		if msgToMQSingle.MsgData.SendID != msgToMQSingle.MsgData.RecvID { //Filter messages sent to yourself
			err2 := l.sendMsgToKafka(&msgToMQSingle, msgToMQSingle.MsgData.SendID, types.OnlineStatus)
			if err2 != nil {
				logx.WithContext(l.ctx).Error(msgToMQSingle.OperationID, "kafka send msg err:SendID", msgToMQSingle.MsgData.SendID, msgToMQSingle.String())
				return returnMsg(&replay, pb, 201, "kafka send msg err", "", 0)
			}
		}
		// callback
		if err := l.callbackAfterSendSingleMsg(pb); err != nil {
			logx.WithContext(l.ctx).Error(utils.GetSelfFuncName(), "callbackAfterSendSingleMsg failed", err.Error())
		}
		return returnMsg(&replay, pb, 0, "", msgToMQSingle.MsgData.ServerMsgID, msgToMQSingle.MsgData.SendTime)
	case types.GroupChatType:
		// callback
		canSend, err := l.callbackBeforeSendGroupMsg(pb)
		if err != nil {
			logx.WithContext(l.ctx).Error(utils.GetSelfFuncName(), "callbackBeforeSendGroupMsg failed", err.Error())
		}
		if !canSend {
			return returnMsg(&replay, pb, 201, "callbackBeforeSendGroupMsg result stop rpc and return", "", 0)
		}
		getGroupMemberIDListFromCacheReq := &imuserpb.GetGroupMemberIDListFromCacheReq{
			GroupID: pb.MsgData.RecvID,
		}
		memberListResp, err := l.svcCtx.ImUser.GetGroupMemberIDListFromCache(l.ctx, getGroupMemberIDListFromCacheReq)
		if err != nil {
			logx.WithContext(l.ctx).Error("GetGroupMemberIDListFromCache rpc call failed ", err.Error())
			return returnMsg(&replay, pb, 201, "GetGroupMemberIDListFromCache failed", "", 0)
		}
		if memberListResp.CommonResp.ErrCode != 0 {
			logx.WithContext(l.ctx).Error("GetGroupMemberIDListFromCache rpc logic call failed ", memberListResp.String())
			return returnMsg(&replay, pb, 201, "GetGroupMemberIDListFromCache logic failed", "", 0)
		}
		memberUserIDList := memberListResp.MemberIDList
		var addUidList []string
		switch pb.MsgData.ContentType {
		case types.MemberKickedNotification:
			var tips chatpb.TipsComm
			var memberKickedTips chatpb.MemberKickedTips
			err := proto.Unmarshal(pb.MsgData.Content, &tips)
			if err != nil {
				logx.WithContext(l.ctx).Error(pb.OperationID, "Unmarshal err", err.Error())
			}
			err = proto.Unmarshal(tips.Detail, &memberKickedTips)
			if err != nil {
				logx.WithContext(l.ctx).Error(pb.OperationID, "Unmarshal err", err.Error())
			}
			logx.WithContext(l.ctx).Info(pb.OperationID, "data is ", memberKickedTips)
			for _, v := range memberKickedTips.KickedUserList {
				addUidList = append(addUidList, v.UserID)
			}
		case types.MemberQuitNotification:
			addUidList = append(addUidList, pb.MsgData.SendID)

		default:
		}
		onUserIDList, offUserIDList := l.getOnlineAndOfflineUserIDList(memberUserIDList)
		groupID := pb.MsgData.GroupID
		//split  parallel send
		var wg sync.WaitGroup
		var sendTag bool
		var split = 50
		remain := len(onUserIDList) % split
		for i := 0; i < len(onUserIDList)/split; i++ {
			wg.Add(1)
			go l.sendMsgToGroup(onUserIDList[i*split:(i+1)*split], pb, types.OnlineStatus, &sendTag, &wg)
		}
		if remain > 0 {
			wg.Add(1)
			go l.sendMsgToGroup(onUserIDList[split*(len(onUserIDList)/split):], pb, types.OnlineStatus, &sendTag, &wg)
		}
		wg.Wait()
		remain = len(offUserIDList) % split
		for i := 0; i < len(offUserIDList)/split; i++ {
			wg.Add(1)
			go l.sendMsgToGroup(offUserIDList[i*split:(i+1)*split], pb, types.OfflineStatus, &sendTag, &wg)
		}
		if remain > 0 {
			wg.Add(1)
			go l.sendMsgToGroup(offUserIDList[split*(len(offUserIDList)/split):], pb, types.OfflineStatus, &sendTag, &wg)
		}
		wg.Wait()
		logx.WithContext(l.ctx).Info(msgToMQSingle.OperationID, "addUidList", addUidList)
		for _, v := range addUidList {
			pb.MsgData.RecvID = v
			isSend := l.modifyMessageByUserMessageReceiveOpt(v, groupID, types.GroupChatType, pb)
			logx.WithContext(l.ctx).Info(msgToMQSingle.OperationID, "isSend", isSend)
			if isSend {
				msgToMQSingle.MsgData = pb.MsgData
				err := l.sendMsgToKafka(&msgToMQSingle, v, types.OnlineStatus)
				if err != nil {
					logx.WithContext(l.ctx).Error(msgToMQSingle.OperationID, "kafka send msg err:UserId", v, msgToMQSingle.String())
				} else {
					sendTag = true
				}
			}
		}
		// callback
		if err := l.callbackAfterSendGroupMsg(pb); err != nil {
			logx.WithContext(l.ctx).Error(utils.GetSelfFuncName(), "callbackAfterSendGroupMsg failed", err.Error())
		}
		if !sendTag {
			return returnMsg(&replay, pb, 201, "kafka send msg err", "", 0)
		} else {
			if pb.MsgData.ContentType == types.AtText {
				go func() {
					var conversationReq conversationpb.ModifyConversationFieldReq
					var tag bool
					var atUserID []string
					conversation := conversationpb.Conversation{
						OwnerUserID:      pb.MsgData.SendID,
						ConversationID:   types.GetConversationIDBySessionType(pb.MsgData.GroupID, types.GroupChatType),
						ConversationType: types.GroupChatType,
						GroupID:          pb.MsgData.GroupID,
					}
					conversationReq.Conversation = &conversation
					conversationReq.OperationID = pb.OperationID
					conversationReq.FieldType = conversationpb.FieldType_FieldGroupAtType
					tagAll := strUtils.IsContain(types.AtAllString, pb.MsgData.AtUserIDList)
					if tagAll {
						atUserID = strUtils.DifferenceString([]string{types.AtAllString}, pb.MsgData.AtUserIDList)
						if len(atUserID) == 0 { //just @everyone
							conversationReq.UserIDList = memberUserIDList
							conversation.GroupAtType = types.AtAll
						} else { //@Everyone and @other people
							conversationReq.UserIDList = atUserID
							conversation.GroupAtType = types.AtAllAtMe
							tag = true
						}
					} else {
						conversationReq.UserIDList = pb.MsgData.AtUserIDList
						conversation.GroupAtType = types.AtMe
					}
					conversationReply, err := l.svcCtx.Conversation.ModifyConversationField(l.ctx, &conversationReq)
					if err != nil {
						logx.WithContext(l.ctx).Error(conversationReq.OperationID, "ModifyConversationField rpc failed, ", conversationReq.String(), err.Error())
					} else if conversationReply.CommonResp.ErrCode != 0 {
						logx.WithContext(l.ctx).Error(conversationReq.OperationID, "ModifyConversationField rpc failed, ", conversationReq.String(), conversationReply.String())
					}
					if tag {
						conversationReq.UserIDList = strUtils.DifferenceString(atUserID, memberUserIDList)
						conversation.GroupAtType = types.AtAll
						conversationReply, err := l.svcCtx.Conversation.ModifyConversationField(context.Background(), &conversationReq)
						if err != nil {
							logx.WithContext(l.ctx).Error(conversationReq.OperationID, "ModifyConversationField rpc failed, ", conversationReq.String(), err.Error())
						} else if conversationReply.CommonResp.ErrCode != 0 {
							logx.WithContext(l.ctx).Error(conversationReq.OperationID, "ModifyConversationField rpc failed, ", conversationReq.String(), conversationReply.String())
						}
					}
				}()
			}
			return returnMsg(&replay, pb, 0, "", msgToMQSingle.MsgData.ServerMsgID, msgToMQSingle.MsgData.SendTime)

		}
	case types.NotificationChatType:
		msgToMQSingle.MsgData = pb.MsgData
		logx.WithContext(l.ctx).Info(msgToMQSingle.OperationID, msgToMQSingle)
		err1 := l.sendMsgToKafka(&msgToMQSingle, msgToMQSingle.MsgData.RecvID, types.OnlineStatus)
		if err1 != nil {
			logx.WithContext(l.ctx).Error(msgToMQSingle.OperationID, "kafka send msg err:RecvID", msgToMQSingle.MsgData.RecvID, msgToMQSingle.String())
			return returnMsg(&replay, pb, 201, "kafka send msg err", "", 0)
		}

		if msgToMQSingle.MsgData.SendID != msgToMQSingle.MsgData.RecvID { //Filter messages sent to yourself
			err2 := l.sendMsgToKafka(&msgToMQSingle, msgToMQSingle.MsgData.SendID, types.OnlineStatus)
			if err2 != nil {
				logx.WithContext(l.ctx).Error(msgToMQSingle.OperationID, "kafka send msg err:SendID", msgToMQSingle.MsgData.SendID, msgToMQSingle.String())
				return returnMsg(&replay, pb, 201, "kafka send msg err", "", 0)
			}
		}
		return returnMsg(&replay, pb, 0, "", msgToMQSingle.MsgData.ServerMsgID, msgToMQSingle.MsgData.SendTime)
	default:
		return returnMsg(&replay, pb, 203, "unkonwn sessionType", "", 0)
	}
}

func (l *SendMsgLogic) getOnlineAndOfflineUserIDList(list []string) (online []string, offline []string) {
	return list, nil
}

func returnMsg(replay *chatpb.SendMsgResp, pb *chatpb.SendMsgReq, errCode int32, errMsg, serverMsgID string, sendTime int64) (*chatpb.SendMsgResp, error) {
	replay.ErrCode = errCode
	replay.ErrMsg = errMsg
	replay.ServerMsgID = serverMsgID
	replay.ClientMsgID = pb.MsgData.ClientMsgID
	replay.SendTime = sendTime
	return replay, nil
}

func (l *SendMsgLogic) userRelationshipVerification(data *chatpb.SendMsgReq) (bool, int32, string) {
	if data.MsgData.SessionType == types.GroupChatType {
		return true, 0, ""
	}
	// 是不是拉黑了
	ifInBlackResp, err := l.svcCtx.ImUser.IfAInBBlacklist(l.ctx, &imuserpb.IfAInBBlacklistReq{
		AUserID: data.MsgData.SendID,
		BUserID: data.MsgData.RecvID,
	})
	if err != nil {
		logx.WithContext(l.ctx).Error(data.OperationID, "GetBlackIDListFromCache rpc call failed ", err.Error())
	} else {
		if ifInBlackResp.CommonResp.ErrCode != 0 {
			logx.WithContext(l.ctx).Error(data.OperationID, "GetBlackIDListFromCache rpc logic call failed ", ifInBlackResp.String())
		} else {
			if ifInBlackResp.IsInBlacklist {
				return false, 600, "in black list"
			}
		}
	}
	if l.svcCtx.Config.MessageVerify.FriendVerify {
		// 是不是好友
		ifInFriendResp, err := l.svcCtx.ImUser.IfAInBFriendList(l.ctx, &imuserpb.IfAInBFriendListReq{
			AUserID: data.MsgData.SendID,
			BUserID: data.MsgData.RecvID,
		})
		if err != nil {
			logx.WithContext(l.ctx).Error(data.OperationID, "GetFriendIDListFromCache rpc call failed ", err.Error())
		} else {
			if ifInFriendResp.CommonResp.ErrCode != 0 {
				logx.WithContext(l.ctx).Error(data.OperationID, "GetFriendIDListFromCache rpc logic call failed ", ifInFriendResp.String())
			} else {
				if !ifInFriendResp.IsInFriendList {
					return false, 601, "not friend"
				}
			}
		}
		return true, 0, ""
	} else {
		return true, 0, ""
	}
}
