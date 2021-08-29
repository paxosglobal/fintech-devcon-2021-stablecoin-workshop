import React from "react";
import { Grid } from "@material-ui/core";
import {
  DataGrid,
  GridColDef,
  GridValueFormatterParams,
  GridValueGetterParams,
} from "@material-ui/data-grid";

import { useActivities } from "../../providers/activity-context";
import { AssetType, currencyFormat } from "../../types/activity";

// destinationAddress: string;
// chainTxId: string;

const columns: Array<GridColDef> = [
  {
    field: "id",
    type: "number",
    width: 110,
    headerName: "ID",
  },
  {
    field: "type",
    width: 170,
    headerName: "Type",
    valueFormatter: (params) =>
      (params.value as string).charAt(0).toUpperCase() +
      (params.value as string).slice(1).toLowerCase(),
  },
  {
    field: "amount",
    width: 170,
    type: "number",
    headerName: "Amount",
    valueFormatter: (params) => currencyFormat(params.value as string),
  },
  {
    field: "destinationAddress",
    width: 170,
    headerName: "Destination",
    valueFormatter: (params) => params.value || "-",
  },
  {
    type: "string",
    field: "status",
    headerName: "Status",
    width: 160,
    valueFormatter: (params) =>
      (params.value as string).charAt(0).toUpperCase() +
      (params.value as string).slice(1).toLowerCase(),
  },
  {
    field: "createdAt",
    width: 410,
    type: "date",
    headerName: "Created At",
    valueFormatter: (params) =>
      new Date(params.value as string).toLocaleString("en-US"),
  },
];

export default function Deposit() {
  const {
    state: { activities },
  } = useActivities();
  return (
    <div>
      <h2>Activities</h2>

      <div style={{ height: 800, width: "100%" }}>
        <DataGrid disableSelectionOnClick rows={activities} columns={columns} />
      </div>
      {/*<div style={{ maxWidth: "100%;" }}>*/}
      {/*  <pre>*/}
      {/*    <code>{JSON.stringify(activities, null, 2)}</code>*/}
      {/*  </pre>*/}
      {/*</div>*/}
    </div>
  );
}
