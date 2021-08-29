import React from "react";
import { Button, Grid, TextField } from "@material-ui/core";
import NumberFormat from "react-number-format";

import { postNewActivity } from "../../api";
import { AssetType, Withdrawal } from "../../types/activity";

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

export const WithdrawalForm = () => {
  const [destinationAddress, setDestinationAddress] = React.useState("");
  const handleAddressChange = (e: React.ChangeEvent<HTMLInputElement>) => {
    setDestinationAddress(e.target.value);
  };

  const [amount, setAmount] = React.useState("0");
  const handleAmountChange = (e: React.ChangeEvent<HTMLInputElement>) => {
    setAmount(e.target.value);
  };

  const handleSubmit = async () => {
    const withdrawal: Withdrawal = {
      asset: AssetType.usd,
      destinationAddress: destinationAddress,
      amount: amount,
    };
    await postNewActivity(withdrawal);
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
          label="Destination Address"
          placeholder="0xaddress"
          fullWidth
          margin="normal"
          InputLabelProps={{
            shrink: true,
          }}
          variant="outlined"
          value={destinationAddress}
          onChange={handleAddressChange}
        />
        <TextField
          id="outlined-full-width"
          label="Amount to Withdrawal"
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
          onChange={handleAmountChange}
        />
        <Button onClick={handleSubmit} variant="contained" color="secondary">
          Submit
        </Button>
      </Grid>
    </form>
  );
};
