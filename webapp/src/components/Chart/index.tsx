import React from "react";
import { useTheme } from "@material-ui/core/styles";
import {
  LineChart,
  Line,
  XAxis,
  YAxis,
  Label,
  ResponsiveContainer,
} from "recharts";
import { Title } from "../Title";
import { Activity, calcActivityAmount } from "../../types/activity";

interface ChartProps {
  activities: Array<Activity>;
}

interface ChartPoint {
  count: number;
  amount: number;
}

type ChartData = Array<ChartPoint>;

const activitiesToChartData = (activities: Array<Activity>): ChartData => {
  let previousSum = 0;
  return activities.map((activity, i) => {
    if (activities[i - 1] !== undefined) {
      previousSum = previousSum + calcActivityAmount(activities[i - 1]);
    }

    return {
      count: i,
      amount: previousSum + calcActivityAmount(activity),
    };
  });
};

export default function Chart(props: ChartProps) {
  const theme = useTheme();
  const data = activitiesToChartData(props.activities);

  return (
    <React.Fragment>
      <Title>Balance History</Title>
      <ResponsiveContainer width="100%" height={300}>
        <LineChart
          data={data}
          margin={{
            top: 16,
            right: 16,
            bottom: 0,
            left: 24,
          }}
        >
          <XAxis dataKey="time" stroke={theme.palette.text.secondary} />
          <YAxis stroke={theme.palette.text.secondary}>
            <Label
              angle={270}
              position="left"
              style={{ textAnchor: "middle", fill: theme.palette.text.primary }}
            >
              Activity History ($)
            </Label>
          </YAxis>
          <Line
            type="monotone"
            dataKey="amount"
            stroke={theme.palette.primary.main}
            dot={false}
          />
        </LineChart>
      </ResponsiveContainer>
    </React.Fragment>
  );
}
