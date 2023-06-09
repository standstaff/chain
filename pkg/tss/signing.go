package tss

import (
	"errors"

	"github.com/bandprotocol/chain/v2/pkg/tss/internal/lagrange"
	"github.com/bandprotocol/chain/v2/pkg/tss/internal/schnorr"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/decred/dcrd/dcrec/secp256k1/v4"
)

// ComputeLagrangeCoefficient calculates the Lagrange coefficient for a given member ID and total number of members.
// Note: Currently, supports a maximum mid at 20.
func ComputeLagrangeCoefficient(mid MemberID, memberList []MemberID) Scalar {
	var mids []int64
	for _, member := range memberList {
		mids = append(mids, int64(member))
	}

	coeff := lagrange.ComputeCoefficientPreCompute(int64(mid), mids).Bytes()

	scalarValue := new(secp256k1.ModNScalar)
	scalarValue.SetByteSlice(coeff)

	return ParseScalar(scalarValue)
}

// ComputeCommitment calculates the bytes that consists of memberID, public D, and public E.
func ComputeCommitment(mids []MemberID, pubDs PublicKeys, pubEs PublicKeys) ([]byte, error) {
	if len(mids) != len(pubDs) {
		return nil, errors.New("length is not equal")
	}

	if len(mids) != len(pubEs) {
		return nil, errors.New("length is not equal")
	}

	var commitment []byte
	for i, mid := range mids {
		commitment = append(commitment, sdk.Uint64ToBigEndian(uint64(mid))...)
		commitment = append(commitment, pubDs[i]...)
		commitment = append(commitment, pubEs[i]...)
	}

	return commitment, nil
}

// ComputeOwnBindingFactor calculates the own binding factor (Lo) value for a given member ID, data, and commitment.
// bindingFactor = Hash(i, data , B)
// B = <<i,Di,Ei>,...>
func ComputeOwnBindingFactor(mid MemberID, data []byte, commitment []byte) Scalar {
	bz := Hash([]byte("signingLo"), sdk.Uint64ToBigEndian(uint64(mid)), Hash(data), Hash(commitment))

	var bindingFactor secp256k1.ModNScalar
	bindingFactor.SetByteSlice(bz)

	return ParseScalar(&bindingFactor)
}

// ComputeOwnPubNonce calculates the own public nonce for a given public D, public E, and binding factor.
// Formula: D + bindingFactor * E
func ComputeOwnPubNonce(rawPubD PublicKey, rawPubE PublicKey, rawBindingFactor Scalar) (PublicKey, error) {
	bindingFactor, err := rawBindingFactor.Parse()
	if err != nil {
		return nil, err
	}

	pubD, err := rawPubD.Point()
	if err != nil {
		return nil, err
	}

	pubE, err := rawPubE.Point()
	if err != nil {
		return nil, err
	}

	var mulE secp256k1.JacobianPoint
	secp256k1.ScalarMultNonConst(bindingFactor, pubE, &mulE)

	var ownPubNonce secp256k1.JacobianPoint
	secp256k1.AddNonConst(pubD, &mulE, &ownPubNonce)

	return ParsePublicKey(&ownPubNonce), nil
}

// ComputeOwnPrivNonce calculates the own private nonce for a given private d, private e, and binsing factor.
// Formula: d + bindingFactor * e
func ComputeOwnPrivNonce(rawPrivD PrivateKey, rawPrivE PrivateKey, rawBindingFactor Scalar) (PrivateKey, error) {
	bindingFactor, err := rawBindingFactor.Parse()
	if err != nil {
		return nil, err
	}

	privD, err := rawPrivD.Scalar()
	if err != nil {
		return nil, err
	}

	privE, err := rawPrivE.Scalar()
	if err != nil {
		return nil, err
	}

	bindingFactor.Mul(privE)
	privD.Add(bindingFactor)

	return ParsePrivateKey(privD), nil
}

// ComputeGroupPublicNonce calculates the group public nonce for a given slice of own public nonces.
// Formula: Sum(PubNonce1, PubNonce2, ..., PubNonceN)
func ComputeGroupPublicNonce(rawOwnPubNonces ...PublicKey) (PublicKey, error) {
	pubNonces, err := PublicKeys(rawOwnPubNonces).Points()
	if err != nil {
		return nil, err
	}

	return ParsePublicKey(sumPoints(pubNonces...)), nil
}

// CombineSignatures performs combining all signatures by sum up R and sum up S.
func CombineSignatures(rawSigs ...Signature) (Signature, error) {
	var allR []*secp256k1.JacobianPoint
	var allS []*secp256k1.ModNScalar
	for _, rawSig := range rawSigs {
		sig, err := rawSig.Parse()
		if err != nil {
			return nil, err
		}

		allR = append(allR, &sig.R)
		allS = append(allS, &sig.S)
	}

	return ParseSignature(schnorr.NewSignature(sumPoints(allR...), sumScalars(allS...))), nil
}

// SignSigning performs signing using the group public nonce, group public key, data, Lagrange coefficient,
// own private nonce, and own private key.
func SignSigning(
	groupPubNonce PublicKey,
	groupPubKey PublicKey,
	data []byte,
	rawLagrange Scalar,
	ownPrivNonce PrivateKey,
	ownPrivKey PrivateKey,
) (Signature, error) {
	msg := ConcatBytes(groupPubNonce, generateMessageGroupSigning(groupPubKey, data))
	return Sign(ownPrivKey, msg, Scalar(ownPrivNonce), rawLagrange)
}

// VerifySigning verifies the signing using the group public nonce, group public key, data, Lagrange coefficient,
// signature, and own public key.
func VerifySigningSig(
	groupPubNonce PublicKey,
	groupPubKey PublicKey,
	data []byte,
	rawLagrange Scalar,
	sig Signature,
	ownPubKey PublicKey,
) error {
	msg := ConcatBytes(groupPubNonce, generateMessageGroupSigning(groupPubKey, data))
	return Verify(sig.R(), sig.S(), msg, ownPubKey, nil, rawLagrange)
}

// VerifyGroupSigning verifies the group signing using the group public key, data, and signature.
func VerifyGroupSigningSig(
	groupPubKey PublicKey,
	data []byte,
	sig Signature,
) error {
	msg := ConcatBytes(sig.R(), generateMessageGroupSigning(groupPubKey, data))
	return Verify(sig.R(), sig.S(), msg, groupPubKey, nil, nil)
}

// generateMessageGroupSigning generates the message for group signing using the group public key and data.
func generateMessageGroupSigning(rawGroupPubKey PublicKey, data []byte) []byte {
	return ConcatBytes([]byte("signing"), rawGroupPubKey, data)
}
