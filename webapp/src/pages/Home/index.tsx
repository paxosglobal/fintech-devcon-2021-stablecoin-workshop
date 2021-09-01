import React, { useEffect, useState } from "react";
import { Divider, Grid, Paper } from "@material-ui/core";
import { makeStyles } from "@material-ui/core/styles";

import { getBalances } from "../../api";
import { DepositForm } from "../../components/DepositForm";
import { WithdrawalForm } from "../../components/WithdrawalForm";
import Chart from "../../components/Chart";
import { useActivities } from "../../providers/activity-context";
import useInterval from "../../hooks/useInterval";
import { BalanceCards } from "../../components/Balances";
import { Balances } from "../../types/activity";

const useStyles = makeStyles((theme) => ({
  paper: {
    padding: theme.spacing(1),
    textAlign: "center",
    color: theme.palette.text.secondary,
    whiteSpace: "nowrap",
    marginBottom: theme.spacing(1),
  },
  divider: {
    margin: theme.spacing(8, 0),
  },
}));

export default function Home() {
  const classes = useStyles();

  const [balances, setBalances] = useState({} as Balances);
  const {
    state: { activities },
  } = useActivities();
  useInterval(() => {
    async function fetchApi() {
      const balances = await getBalances();
      setBalances(balances);
    }

    fetchApi();
  }, 500);

  useEffect(() => {
    async function fetchApi() {
      const balances = await getBalances();
      setBalances(balances);
    }

    fetchApi();
  }, []);

  return (
    <div>
      <Grid container spacing={8}>
        <Grid item xs={12}>
          <BalanceCards balances={balances} />
        </Grid>
        <Grid item xs={12}>
          <Chart activities={activities} />
        </Grid>
      </Grid>
      <Divider className={classes.divider} />
      <Grid container spacing={2}>
        <Grid item xs={6}>
          <Paper className={classes.paper}>
            <h3>Deposit</h3>
            <DepositForm />
          </Paper>
        </Grid>
        <Grid item xs={6}>
          <Paper className={classes.paper}>
            <h3>Withdraw</h3>
            <WithdrawalForm />
          </Paper>
        </Grid>
      </Grid>
    </div>
  );
}
