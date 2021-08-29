enum ActivityType {
  deposit = "DEPOSIT",
  withdrawal = "WITHDRAWAL",
}

enum ActivityStatus {
  created = "CREATED",
  completed = "COMPLETED",
}

export enum AssetType {
  usd = "USD",
}

export interface Balances {
  usdOnPlatform: number;
  usdInReserve: number;
  usdkMinted: number;
}

export interface Deposit {
  asset: AssetType;
  amount: string;
}

export interface Withdrawal {
  asset: AssetType;
  amount: string;
  destinationAddress: string;
}

export interface Activity {
  type: ActivityType;
  asset: AssetType;
  amount: string;
  destinationAddress: string;
  iD: number;
  chainTxId: string;
  status: ActivityStatus;
  createdAt: Date;
  completedAt: Date;
}

export const calcActivityAmount = (activity: Activity): number => {
  const sign = activity.type === ActivityType.deposit ? 1 : -1;
  return parseFloat(activity.amount) * sign;
};

export const currencyFormat = (amount: string | number): string => {
  if (typeof amount === "string") {
    amount = parseFloat(amount);
  }

  if (!amount || isNaN(amount)) return ''

  return amount.toLocaleString("en-US", {
    style: "currency",
    currency: "USD",
  });
};
