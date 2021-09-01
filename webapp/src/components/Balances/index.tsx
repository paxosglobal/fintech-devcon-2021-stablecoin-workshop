import React from "react";
import { Card, CardContent, Grid } from "@material-ui/core";
import { Balances, currencyFormat } from "../../types/activity";
import Typography from "@material-ui/core/Typography";

interface BalanceProps {
  balances: Balances;
}

export const BalanceCards = (props: BalanceProps) => {
  const { balances } = props;

  if (!balances) return null;

  return (
    <Grid
      container
      direction="row"
      justifyContent="space-between"
      alignItems="center"
    >
      <Card variant="outlined">
        <CardContent>
          <Typography color="textSecondary" gutterBottom>
            Total USDK Minted
          </Typography>
          <Typography color="primary" variant="h5" component="h2">
            {currencyFormat(balances.usdkMinted)}
          </Typography>
        </CardContent>
      </Card>
      <Card variant="outlined">
        <CardContent>
          <Typography color="textSecondary" gutterBottom>
            Total USD On Platform
          </Typography>
          <Typography color="primary" variant="h5" component="h2">
            {currencyFormat(balances.usdOnPlatform)}
          </Typography>
        </CardContent>
      </Card>
      <Card variant="outlined">
        <CardContent>
          <Typography color="textSecondary" gutterBottom>
            Total USDK In Reserve
          </Typography>
          <Typography color="primary" variant="h5" component="h2">
            {currencyFormat(balances.usdInReserve)}
          </Typography>
        </CardContent>
      </Card>
    </Grid>
  );
};
