/* eslint-disable */
import { Writer, Reader } from "protobufjs/minimal";

export const protobufPackage = "DecentralCardGame.cardchain.cardchain";

export enum SellOfferStatus {
  open = 0,
  sold = 1,
  removed = 2,
  UNRECOGNIZED = -1,
}

export function sellOfferStatusFromJSON(object: any): SellOfferStatus {
  switch (object) {
    case 0:
    case "open":
      return SellOfferStatus.open;
    case 1:
    case "sold":
      return SellOfferStatus.sold;
    case 2:
    case "removed":
      return SellOfferStatus.removed;
    case -1:
    case "UNRECOGNIZED":
    default:
      return SellOfferStatus.UNRECOGNIZED;
  }
}

export function sellOfferStatusToJSON(object: SellOfferStatus): string {
  switch (object) {
    case SellOfferStatus.open:
      return "open";
    case SellOfferStatus.sold:
      return "sold";
    case SellOfferStatus.removed:
      return "removed";
    default:
      return "UNKNOWN";
  }
}

export interface SellOffer {
  seller: string;
  buyer: string;
  card: string;
  status: SellOfferStatus;
}

const baseSellOffer: object = { seller: "", buyer: "", card: "", status: 0 };

export const SellOffer = {
  encode(message: SellOffer, writer: Writer = Writer.create()): Writer {
    if (message.seller !== "") {
      writer.uint32(10).string(message.seller);
    }
    if (message.buyer !== "") {
      writer.uint32(18).string(message.buyer);
    }
    if (message.card !== "") {
      writer.uint32(26).string(message.card);
    }
    if (message.status !== 0) {
      writer.uint32(32).int32(message.status);
    }
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): SellOffer {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseSellOffer } as SellOffer;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.seller = reader.string();
          break;
        case 2:
          message.buyer = reader.string();
          break;
        case 3:
          message.card = reader.string();
          break;
        case 4:
          message.status = reader.int32() as any;
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): SellOffer {
    const message = { ...baseSellOffer } as SellOffer;
    if (object.seller !== undefined && object.seller !== null) {
      message.seller = String(object.seller);
    } else {
      message.seller = "";
    }
    if (object.buyer !== undefined && object.buyer !== null) {
      message.buyer = String(object.buyer);
    } else {
      message.buyer = "";
    }
    if (object.card !== undefined && object.card !== null) {
      message.card = String(object.card);
    } else {
      message.card = "";
    }
    if (object.status !== undefined && object.status !== null) {
      message.status = sellOfferStatusFromJSON(object.status);
    } else {
      message.status = 0;
    }
    return message;
  },

  toJSON(message: SellOffer): unknown {
    const obj: any = {};
    message.seller !== undefined && (obj.seller = message.seller);
    message.buyer !== undefined && (obj.buyer = message.buyer);
    message.card !== undefined && (obj.card = message.card);
    message.status !== undefined &&
      (obj.status = sellOfferStatusToJSON(message.status));
    return obj;
  },

  fromPartial(object: DeepPartial<SellOffer>): SellOffer {
    const message = { ...baseSellOffer } as SellOffer;
    if (object.seller !== undefined && object.seller !== null) {
      message.seller = object.seller;
    } else {
      message.seller = "";
    }
    if (object.buyer !== undefined && object.buyer !== null) {
      message.buyer = object.buyer;
    } else {
      message.buyer = "";
    }
    if (object.card !== undefined && object.card !== null) {
      message.card = object.card;
    } else {
      message.card = "";
    }
    if (object.status !== undefined && object.status !== null) {
      message.status = object.status;
    } else {
      message.status = 0;
    }
    return message;
  },
};

type Builtin = Date | Function | Uint8Array | string | number | undefined;
export type DeepPartial<T> = T extends Builtin
  ? T
  : T extends Array<infer U>
  ? Array<DeepPartial<U>>
  : T extends ReadonlyArray<infer U>
  ? ReadonlyArray<DeepPartial<U>>
  : T extends {}
  ? { [K in keyof T]?: DeepPartial<T[K]> }
  : Partial<T>;
