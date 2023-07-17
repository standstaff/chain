package keeper_test

import (
	"encoding/hex"
	"fmt"
	"time"

	"github.com/bandprotocol/chain/v2/pkg/tss"
	"github.com/bandprotocol/chain/v2/x/tss/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/x/authz"
)

func (s *KeeperTestSuite) TestGRPCQueryGroup() {
	ctx, msgSrvr, q, k := s.ctx, s.msgSrvr, s.querier, s.app.TSSKeeper
	groupID, memberID1, memberID2 := tss.GroupID(1), tss.MemberID(1), tss.MemberID(2)

	members := []string{
		"band18gtd9xgw6z5fma06fxnhet7z2ctrqjm3z4k7ad",
		"band1s743ydr36t6p29jsmrxm064guklgthsn3t90ym",
		"band1p08slm6sv2vqy4j48hddkd6hpj8yp6vlw3pf8p",
		"band1p08slm6sv2vqy4j48hddkd6hpj8yp6vlw3pf8p",
		"band12jf07lcaj67mthsnklngv93qkeuphhmxst9mh8",
	}
	round1Info1 := types.Round1Info{
		MemberID: memberID1,
		CoefficientCommits: []tss.Point{
			[]byte("point1"),
			[]byte("point2"),
			[]byte("point3"),
		},
		OneTimePubKey: []byte("OneTimePubKeySample"),
		A0Sig:         []byte("A0SigSample"),
		OneTimeSig:    []byte("OneTimeSigSample"),
	}
	round1Info2 := types.Round1Info{
		MemberID: memberID2,
		CoefficientCommits: []tss.Point{
			[]byte("point1"),
			[]byte("point2"),
			[]byte("point3"),
		},
		OneTimePubKey: []byte("OneTimePubKeySample"),
		A0Sig:         []byte("A0SigSample"),
		OneTimeSig:    []byte("OneTimeSigSample"),
	}
	round2Info1 := types.Round2Info{
		MemberID: memberID1,
		EncryptedSecretShares: tss.Scalars{
			[]byte("scalar1"),
			[]byte("scalar2"),
		},
	}
	round2Info2 := types.Round2Info{
		MemberID: memberID2,
		EncryptedSecretShares: tss.Scalars{
			[]byte("scalar1"),
			[]byte("scalar2"),
		},
	}
	complaintWithStatus1 := types.ComplaintsWithStatus{
		MemberID: memberID1,
		ComplaintsWithStatus: []types.ComplaintWithStatus{
			{
				Complaint: types.Complaint{
					Complainant: 1,
					Respondent:  2,
					KeySym:      []byte("key_sym"),
					Signature:   []byte("signature"),
				},
				ComplaintStatus: types.COMPLAINT_STATUS_SUCCESS,
			},
		},
	}
	complaintWithStatus2 := types.ComplaintsWithStatus{
		MemberID: memberID2,
		ComplaintsWithStatus: []types.ComplaintWithStatus{
			{
				Complaint: types.Complaint{
					Complainant: 1,
					Respondent:  2,
					KeySym:      []byte("key_sym"),
					Signature:   []byte("signature"),
				},
				ComplaintStatus: types.COMPLAINT_STATUS_SUCCESS,
			},
		},
	}
	confirm1 := types.Confirm{
		MemberID:     memberID1,
		OwnPubKeySig: []byte("own_pub_key_sig"),
	}
	confirm2 := types.Confirm{
		MemberID:     memberID2,
		OwnPubKeySig: []byte("own_pub_key_sig"),
	}

	msgSrvr.CreateGroup(ctx, &types.MsgCreateGroup{
		Members:   members,
		Threshold: 3,
		Sender:    members[0],
	})

	// Set round 1 infos
	k.SetRound1Info(ctx, groupID, round1Info1)
	k.SetRound1Info(ctx, groupID, round1Info2)

	// Set round 2 infos
	k.SetRound2Info(ctx, groupID, round2Info1)
	k.SetRound2Info(ctx, groupID, round2Info2)

	// Set complains
	k.SetComplaintsWithStatus(ctx, groupID, complaintWithStatus1)
	k.SetComplaintsWithStatus(ctx, groupID, complaintWithStatus2)

	// Set confirms
	k.SetConfirm(ctx, groupID, confirm1)
	k.SetConfirm(ctx, groupID, confirm2)

	var req types.QueryGroupRequest
	testCases := []struct {
		msg      string
		malleate func()
		expPass  bool
		postTest func(res *types.QueryGroupResponse)
	}{
		{
			"non existing group",
			func() {
				req = types.QueryGroupRequest{
					GroupId: 2,
				}
			},
			false,
			func(res *types.QueryGroupResponse) {},
		},
		{
			"success",
			func() {
				req = types.QueryGroupRequest{
					GroupId: uint64(groupID),
				}
			},
			true,
			func(res *types.QueryGroupResponse) {
				dkgContextB, _ := hex.DecodeString("6c31fc15422ebad28aaf9089c306702f67540b53c7eea8b7d2941044b027100f")

				s.Require().Equal(&types.QueryGroupResponse{
					Group: types.Group{
						GroupID:   1,
						Size_:     5,
						Threshold: 3,
						PubKey:    nil,
						Status:    types.GROUP_STATUS_ROUND_1,
					},
					DKGContext: dkgContextB,
					Members: []types.Member{
						{
							MemberID:    1,
							Address:     "band18gtd9xgw6z5fma06fxnhet7z2ctrqjm3z4k7ad",
							PubKey:      nil,
							IsMalicious: false,
						},
						{
							MemberID:    2,
							Address:     "band1s743ydr36t6p29jsmrxm064guklgthsn3t90ym",
							PubKey:      nil,
							IsMalicious: false,
						},
						{
							MemberID:    3,
							Address:     "band1p08slm6sv2vqy4j48hddkd6hpj8yp6vlw3pf8p",
							PubKey:      nil,
							IsMalicious: false,
						},
						{
							MemberID:    4,
							Address:     "band1p08slm6sv2vqy4j48hddkd6hpj8yp6vlw3pf8p",
							PubKey:      nil,
							IsMalicious: false,
						},
						{
							MemberID:    5,
							Address:     "band12jf07lcaj67mthsnklngv93qkeuphhmxst9mh8",
							PubKey:      nil,
							IsMalicious: false,
						},
					},
					Round1Infos: []types.Round1Info{
						round1Info1,
						round1Info2,
					},
					Round2Infos: []types.Round2Info{
						round2Info1,
						round2Info2,
					},
					ComplaintsWithStatus: []types.ComplaintsWithStatus{
						complaintWithStatus1,
						complaintWithStatus2,
					},
					Confirms: []types.Confirm{
						confirm1,
						confirm2,
					},
				}, res)
			},
		},
	}

	for _, tc := range testCases {
		s.Run(fmt.Sprintf("Case %s", tc.msg), func() {
			tc.malleate()

			res, err := q.Group(ctx, &req)
			if tc.expPass {
				s.Require().NoError(err)
			} else {
				s.Require().Error(err)
			}

			tc.postTest(res)
		})
	}
}

func (s *KeeperTestSuite) TestGRPCQueryMembers() {
	ctx, q, k := s.ctx, s.querier, s.app.TSSKeeper
	members := []types.Member{
		{
			MemberID:    1,
			Address:     "band1m5lq9u533qaya4q3nfyl6ulzqkpkhge9q8tpzs",
			PubKey:      nil,
			IsMalicious: false,
		},
		{
			MemberID:    2,
			Address:     "band1p40yh3zkmhcv0ecqp3mcazy83sa57rgjp07dun",
			PubKey:      nil,
			IsMalicious: false,
		},
	}

	// Set members
	for _, m := range members {
		k.SetMember(ctx, tss.GroupID(1), m)
	}

	var req types.QueryMembersRequest
	testCases := []struct {
		msg      string
		malleate func()
		expPass  bool
		postTest func(res *types.QueryMembersResponse)
	}{
		{
			"non existing member",
			func() {
				req = types.QueryMembersRequest{
					GroupId: 2,
				}
			},
			false,
			func(res *types.QueryMembersResponse) {},
		},
		{
			"success",
			func() {
				req = types.QueryMembersRequest{
					GroupId: 1,
				}
			},
			true,
			func(res *types.QueryMembersResponse) {
				s.Require().Equal(members, res.Members)
			},
		},
	}

	for _, tc := range testCases {
		s.Run(fmt.Sprintf("Case %s", tc.msg), func() {
			tc.malleate()

			_, err := q.Members(ctx, &req)
			if tc.expPass {
				s.Require().NoError(err)
			} else {
				s.Require().Error(err)
			}
		})
	}
}

func (s *KeeperTestSuite) TestGRPCQueryIsGrantee() {
	ctx, q, authzKeeper := s.ctx, s.querier, s.app.AuthzKeeper
	expTime := time.Unix(0, 0)

	// Init grantee address
	grantee, _ := sdk.AccAddressFromBech32("band1m5lq9u533qaya4q3nfyl6ulzqkpkhge9q8tpzs")

	// Init granter address
	granter, _ := sdk.AccAddressFromBech32("band1p40yh3zkmhcv0ecqp3mcazy83sa57rgjp07dun")

	// Save grant msgs to grantee
	for _, m := range types.GetMsgGrants() {
		authzKeeper.SaveGrant(s.ctx, grantee, granter, authz.NewGenericAuthorization(m), &expTime)
	}

	var req types.QueryIsGranteeRequest
	testCases := []struct {
		msg      string
		malleate func()
		expPass  bool
		postTest func(res *types.QueryIsGranteeResponse)
	}{
		{
			"address is not bech32",
			func() {
				req = types.QueryIsGranteeRequest{
					Granter: "asdsd1m5lq9u533qaya4q3nfyl6ulzqkpkhge9q8tpzs",
					Grantee: "padads40yh3zkmhcv0ecqp3mcazy83sa57rgjp07dun",
				}
			},
			false,
			func(res *types.QueryIsGranteeResponse) {},
		},
		{
			"address is empty",
			func() {
				req = types.QueryIsGranteeRequest{
					Granter: "",
					Grantee: "",
				}
			},
			false,
			func(res *types.QueryIsGranteeResponse) {},
		},
		{
			"grantee address is not grant by granter",
			func() {
				req = types.QueryIsGranteeRequest{
					Granter: "band1m5lq9u533qaya4q3nfyl6ulzqkpkhge9q8tpzs",
					Grantee: "band17eplw6tga7wqgruqdtalw3rky4njkx6vngxjlt",
				}
			},
			true,
			func(res *types.QueryIsGranteeResponse) {
				s.Require().False(res.IsGrantee)
			},
		},
		{
			"success",
			func() {
				req = types.QueryIsGranteeRequest{
					Granter: "band1m5lq9u533qaya4q3nfyl6ulzqkpkhge9q8tpzs",
					Grantee: "band1p40yh3zkmhcv0ecqp3mcazy83sa57rgjp07dun",
				}
			},
			true,
			func(res *types.QueryIsGranteeResponse) {
				s.Require().True(res.IsGrantee)
			},
		},
	}

	for _, tc := range testCases {
		s.Run(fmt.Sprintf("Case %s", tc.msg), func() {
			tc.malleate()

			_, err := q.IsGrantee(ctx, &req)
			if tc.expPass {
				s.Require().NoError(err)
			} else {
				s.Require().Error(err)
			}
		})
	}
}

func (s *KeeperTestSuite) TestGRPCQueryDE() {
	ctx, q := s.ctx, s.querier

	var req types.QueryDERequest
	testCases := []struct {
		msg      string
		malleate func()
		expPass  bool
		postTest func(res *types.QueryDEResponse, err error)
	}{
		{
			"success",
			func() {
				req = types.QueryDERequest{
					Address: "band1m5lq9u533qaya4q3nfyl6ulzqkpkhge9q8tpzs",
				}
			},
			true,
			func(res *types.QueryDEResponse, err error) {
				s.Require().NoError(err)
				s.Require().NotNil(res)
				s.Require().Len(res.DEs, 0)
			},
		},
	}

	for _, tc := range testCases {
		s.Run(fmt.Sprintf("Case %s", tc.msg), func() {
			tc.malleate()

			res, err := q.DE(ctx, &req)
			if tc.expPass {
				s.Require().NoError(err)
			} else {
				s.Require().Error(err)
			}

			tc.postTest(res, err)
		})
	}
}

func (s *KeeperTestSuite) TestGRPCQueryPendingSigns() {
	ctx, q := s.ctx, s.querier

	var req types.QueryPendingSigningsRequest
	testCases := []struct {
		msg      string
		malleate func()
		expPass  bool
		postTest func(res *types.QueryPendingSigningsResponse, err error)
	}{
		{
			"success",
			func() {
				req = types.QueryPendingSigningsRequest{
					Address: "band1m5lq9u533qaya4q3nfyl6ulzqkpkhge9q8tpzs",
				}
			},
			true,
			func(res *types.QueryPendingSigningsResponse, err error) {
				s.Require().NoError(err)
				s.Require().NotNil(res)
				s.Require().Len(res.PendingSignings, 0)
			},
		},
	}

	for _, tc := range testCases {
		s.Run(fmt.Sprintf("Case %s", tc.msg), func() {
			tc.malleate()

			res, err := q.PendingSignings(ctx, &req)
			if tc.expPass {
				s.Require().NoError(err)
			} else {
				s.Require().Error(err)
			}

			tc.postTest(res, err)
		})
	}
}

func (s *KeeperTestSuite) TestGRPCQuerySigning() {
	ctx, q, k := s.ctx, s.querier, s.app.TSSKeeper
	signingID, memberID, groupID := tss.SigningID(1), tss.MemberID(1), tss.GroupID(1)
	signing := types.Signing{
		SigningID: signingID,
		GroupID:   groupID,
		AssignedMembers: []types.AssignedMember{
			{
				MemberID:      memberID,
				Member:        "band1m5lq9u533qaya4q3nfyl6ulzqkpkhge9q8tpzs",
				PubD:          []byte("D"),
				PubE:          []byte("E"),
				BindingFactor: []byte("binding_factor"),
				PubNonce:      []byte("public_nonce"),
			},
		},
		Message:       []byte("message"),
		GroupPubNonce: []byte("group_pub_nonce"),
		Signature:     []byte("signature"),
	}
	sig := []byte("signatures")

	// Set partial signature
	k.SetPartialSig(ctx, signingID, memberID, []byte("signatures"))

	// Add signing
	k.AddSigning(ctx, signing)

	var req types.QuerySigningRequest
	testCases := []struct {
		msg      string
		malleate func()
		expPass  bool
		postTest func(res *types.QuerySigningResponse, err error)
	}{
		{
			"invalid signing id",
			func() {
				req = types.QuerySigningRequest{
					Id: 999,
				}
			},
			false,
			func(res *types.QuerySigningResponse, err error) {
				s.Require().Error(err)
				s.Require().Nil(res)
			},
		},
		{
			"success",
			func() {
				req = types.QuerySigningRequest{
					Id: 1,
				}
			},
			true,
			func(res *types.QuerySigningResponse, err error) {
				s.Require().NoError(err)
				s.Require().Equal(signing, res.Signing)
				s.Require().
					Equal([]types.PartialSignature{{MemberID: memberID, Signature: sig}}, res.ReceivedPartialSignatures)
			},
		},
	}

	for _, tc := range testCases {
		s.Run(fmt.Sprintf("Case %s", tc.msg), func() {
			tc.malleate()

			res, err := q.Signing(ctx, &req)
			if tc.expPass {
				s.Require().NoError(err)
			} else {
				s.Require().Error(err)
			}

			tc.postTest(res, err)
		})
	}
}

func (s *KeeperTestSuite) TestGRPCQueryStatuses() {
	ctx, q := s.ctx, s.querier

	var req types.QueryStatusesRequest
	testCases := []struct {
		msg      string
		malleate func()
		expPass  bool
		postTest func(res *types.QueryStatusesResponse, err error)
	}{
		{
			"success",
			func() {
				req = types.QueryStatusesRequest{
					Address: "band1m5lq9u533qaya4q3nfyl6ulzqkpkhge9q8tpzs",
				}
			},
			true,
			func(res *types.QueryStatusesResponse, err error) {
				s.Require().NoError(err)
				s.Require().NotNil(res)
				s.Require().Len(res.Statuses, 0)
			},
		},
	}

	for _, tc := range testCases {
		s.Run(fmt.Sprintf("Case %s", tc.msg), func() {
			tc.malleate()

			res, err := q.Statuses(ctx, &req)
			if tc.expPass {
				s.Require().NoError(err)
			} else {
				s.Require().Error(err)
			}

			tc.postTest(res, err)
		})
	}
}
