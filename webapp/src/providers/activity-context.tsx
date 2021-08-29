import * as React from "react";
import { Activity } from "../types/activity";
import { Context, useEffect } from "react";
import useInterval from "../hooks/useInterval";
import { getActivities } from "../api";

interface ActivityProviderProps {
  children: React.ReactNode;
}

interface ActivityState {
  activities: Array<Activity>;
}

type ActivityAction = {
  type: "updateActivities";
  activities: Array<Activity>;
};

type Dispatch = (action: ActivityAction) => void;

const ActivityContext = React.createContext<
  { state: ActivityState; dispatch: Dispatch } | undefined
>(undefined);

const activityReducer = (state: ActivityState, action: ActivityAction) => {
  switch (action.type) {
    case "updateActivities": {
      return {
        activities: action.activities,
      };
    }
    default: {
      throw new Error("Unhandled action type");
    }
  }
};

export const ActivityProvider = ({ children }: ActivityProviderProps) => {
  const [state, dispatch] = React.useReducer(activityReducer, {
    activities: [],
  });

  useInterval(() => {
    getActivities().then((activities) => {
      dispatch({
        type: "updateActivities",
        activities: activities,
      });
    });
  }, 500);

  const value = { state, dispatch };
  return (
    <ActivityContext.Provider value={value}>
      {children}
    </ActivityContext.Provider>
  );
};

export const useActivities = (): {
  state: ActivityState;
  dispatch: Dispatch;
} => {
  const context = React.useContext(ActivityContext);
  if (context === undefined) {
    throw new Error("useActivities must be used within an ActivityProvider");
  }
  return context;
};
