import React from "react";
import { Button, Grid, TextField } from "@material-ui/core";
import NumberFormat from "react-number-format";

import { postNewActivity } from "../../api";
import { AssetType, Deposit } from "../../types/activity";

function NumberFormatCustom(props: any) {
  const { inputRef, onChange, ...other } = props;

  return (
    <NumberFormat
      {...other}
      getInputRef={inputRef}
      onValueChange={(values) => {
        onChange({
          target: {
            name: props.name,
            value: values.value,
          },
        });
      }}
      thousandSeparator
      isNumericString
      prefix="$"
    />
  );
}

export const DepositForm = () => {
  const [amount, setAmount] = React.useState("0");
  const handleChange = (e: React.ChangeEvent<HTMLInputElement>) => {
    setAmount(e.target.value);
  };

  const handleSubmit = async () => {
    const deposit: Deposit = {
      asset: AssetType.usd,
      amount: amount,
    };
    await postNewActivity(deposit);
  };

  return (
    <form noValidate autoComplete="off">
      <Grid
        container
        direction="row"
        justifyContent="space-around"
        alignItems="center"
      >
        <TextField
          id="outlined-full-width"
          label="Amount to Deposit"
          placeholder="800000"
          fullWidth
          margin="normal"
          InputLabelProps={{
            shrink: true,
          }}
          InputProps={{
            inputComponent: NumberFormatCustom,
          }}
          variant="outlined"
          value={amount}
          onChange={handleChange}
        />
        <Button onClick={handleSubmit} variant="contained" color="secondary">
          Submit
        </Button>
      </Grid>
    </form>
  );
};
