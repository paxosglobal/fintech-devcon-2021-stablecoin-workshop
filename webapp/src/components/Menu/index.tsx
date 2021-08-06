import React from 'react';
import DashboardIcon from '@material-ui/icons/Dashboard';
import AccountBalanceIcon from '@material-ui/icons/AccountBalance';
import MonetizationOnIcon from '@material-ui/icons/MonetizationOn';
import SendIcon from '@material-ui/icons/Send';

import { ListItemLink } from '../ListItemLink';

export const menuItems = (
  <div>
    <ListItemLink to="/" primary="Dashboard" icon={<DashboardIcon />} />
    <ListItemLink to="/balances" primary="Balances" icon={<AccountBalanceIcon />} />
    <ListItemLink to="/deposits" primary="Deposits" icon={<MonetizationOnIcon />} />
    <ListItemLink to="/withdrawals" primary="Withdrawals" icon={<SendIcon />} />
  </div>
);