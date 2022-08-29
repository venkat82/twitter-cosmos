// THIS FILE IS GENERATED AUTOMATICALLY. DO NOT MODIFY.

import { StdFee } from "@cosmjs/launchpad";
import { SigningStargateClient } from "@cosmjs/stargate";
import { Registry, OfflineSigner, EncodeObject, DirectSecp256k1HdWallet } from "@cosmjs/proto-signing";
import { Api } from "./rest";
import { MsgAddFollower } from "./types/twitter/tx";
import { MsgAddTweet } from "./types/twitter/tx";
import { MsgFetchFeed } from "./types/twitter/tx";


const types = [
  ["/venkat.twitter.twitter.MsgAddFollower", MsgAddFollower],
  ["/venkat.twitter.twitter.MsgAddTweet", MsgAddTweet],
  ["/venkat.twitter.twitter.MsgFetchFeed", MsgFetchFeed],
  
];
export const MissingWalletError = new Error("wallet is required");

export const registry = new Registry(<any>types);

const defaultFee = {
  amount: [],
  gas: "200000",
};

interface TxClientOptions {
  addr: string
}

interface SignAndBroadcastOptions {
  fee: StdFee,
  memo?: string
}

const txClient = async (wallet: OfflineSigner, { addr: addr }: TxClientOptions = { addr: "http://localhost:26657" }) => {
  if (!wallet) throw MissingWalletError;
  let client;
  if (addr) {
    client = await SigningStargateClient.connectWithSigner(addr, wallet, { registry });
  }else{
    client = await SigningStargateClient.offline( wallet, { registry });
  }
  const { address } = (await wallet.getAccounts())[0];

  return {
    signAndBroadcast: (msgs: EncodeObject[], { fee, memo }: SignAndBroadcastOptions = {fee: defaultFee, memo: ""}) => client.signAndBroadcast(address, msgs, fee,memo),
    msgAddFollower: (data: MsgAddFollower): EncodeObject => ({ typeUrl: "/venkat.twitter.twitter.MsgAddFollower", value: MsgAddFollower.fromPartial( data ) }),
    msgAddTweet: (data: MsgAddTweet): EncodeObject => ({ typeUrl: "/venkat.twitter.twitter.MsgAddTweet", value: MsgAddTweet.fromPartial( data ) }),
    msgFetchFeed: (data: MsgFetchFeed): EncodeObject => ({ typeUrl: "/venkat.twitter.twitter.MsgFetchFeed", value: MsgFetchFeed.fromPartial( data ) }),
    
  };
};

interface QueryClientOptions {
  addr: string
}

const queryClient = async ({ addr: addr }: QueryClientOptions = { addr: "http://localhost:1317" }) => {
  return new Api({ baseUrl: addr });
};

export {
  txClient,
  queryClient,
};
