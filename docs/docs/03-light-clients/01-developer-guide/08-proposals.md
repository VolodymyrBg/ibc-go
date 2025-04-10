---
title: Handling Proposals
sidebar_label: Handling Proposals
sidebar_position: 8
slug: /ibc/light-clients/proposals
---


# Handling proposals

It is possible to update the client with the state of the substitute client through a governance proposal. This type of governance proposal is typically used to recover an expired or frozen client, as it can recover the entire state and therefore all existing channels built on top of the client. `RecoverClient` should be implemented to handle the proposal.

## Implementing `RecoverClient`

In the [`LightClientModule` interface](https://github.com/cosmos/ibc-go/blob/501a8462345da099144efe91d495bfcfa18d760d/modules/core/exported/client.go#L51), we find:

```go
// RecoverClient must verify that the provided substitute 
// may be used to update the subject client. The light client
// must set the updated client and consensus states within
// the clientStore for the subject client.
RecoverClient(
  ctx sdk.Context,
  clientID,
  substituteClientID string,
) error
```

Prior to updating, this function must verify that:

- the substitute client is the same type as the subject client. For a reference implementation, please see the [Tendermint light client](https://github.com/cosmos/ibc-go/blob/47162061bcbfe74df791161059715a635e31c604/modules/light-clients/07-tendermint/proposal_handle.go#L34).
- the provided substitute may be used to update the subject client. This may mean that certain parameters must remain unaltered. For example, a [valid substitute Tendermint light client](https://github.com/cosmos/ibc-go/blob/47162061bcbfe74df791161059715a635e31c604/modules/light-clients/07-tendermint/proposal_handle.go#L86) must NOT change the chain ID, trust level, max clock drift, unbonding period, proof specs or upgrade path. Please note that `AllowUpdateAfterMisbehaviour` and `AllowUpdateAfterExpiry` have been deprecated (see ADR 026 for more information).

After these checks are performed, the function must [set the updated client and consensus states](https://github.com/cosmos/ibc-go/blob/v7.0.0/modules/light-clients/07-tendermint/proposal_handle.go#L77) within the client store for the subject client.

Please refer to the [Tendermint light client implementation](https://github.com/cosmos/ibc-go/blob/47162061bcbfe74df791161059715a635e31c604/modules/light-clients/07-tendermint/proposal_handle.go#L79) for reference.
