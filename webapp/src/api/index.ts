import { Activity, Balances, Deposit, Withdrawal } from "../types/activity";

export const backendUrl = (): string => {
  return process.env.REACT_APP_BACKEND_URL || "http://localhost:8080";
};

const getFullUrlForPath = (path: string): string => {
  return `${backendUrl()}/${path}`;
};

const routes = {
  balances: getFullUrlForPath("balances"),
  activities: getFullUrlForPath("activities"),
  deposits: getFullUrlForPath("deposits"),
  withdrawals: getFullUrlForPath("withdrawals"),
};

export const getBalances = async (): Promise<Balances> => {
  return fetch(routes.balances).then((r) => r.json());
};

export const getActivities = async (): Promise<Array<Activity>> => {
  return fetch(routes.activities).then((r) => r.json());
};

export const postNewActivity = async (
  activity: Deposit | Withdrawal
): Promise<void> => {
  let route = routes.deposits;
  if ("destinationAddress" in activity) {
    route = routes.withdrawals;
  }

  return fetch(route, {
    method: "POST",
    mode: "no-cors", // no-cors, *cors, same-origin
    cache: "no-cache", // *default, no-cache, reload, force-cache, only-if-cached
    headers: {
      "Content-Type": "application/json",
    },
    body: JSON.stringify(activity),
  }).then(console.log);
};
