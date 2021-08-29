import React from "react";
import DashboardIcon from "@material-ui/icons/Dashboard";
import AccountBalanceIcon from "@material-ui/icons/AccountBalance";

import { ListItemLink } from "../ListItemLink";

export const menuItems = (
  <div>
    <ListItemLink to="/" primary="Dashboard" icon={<DashboardIcon />} />
    <ListItemLink
      to="/activities"
      primary="Activities"
      icon={<AccountBalanceIcon />}
    />
  </div>
);
