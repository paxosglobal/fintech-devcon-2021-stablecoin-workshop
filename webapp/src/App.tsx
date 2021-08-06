import React from 'react';
import { createTheme } from '@material-ui/core/styles';
import { ThemeProvider } from '@material-ui/styles';

import { Dashboard } from './Dashboard'
import Home from './pages/Home';
import Balances from './pages/Balances';
import Deposits from './pages/Deposit';
import Withdrawals from './pages/Withdraw';
import {
  BrowserRouter as Router,
  Switch,
  Route,
} from "react-router-dom";

import './App.css';

const theme = createTheme({
  palette: {
    primary: {
      main: '#005121'
    },
  },
});

export default function App() {
  return (
    <Router>
      <ThemeProvider theme={theme}>
        <Dashboard>
          {/*
            A <Switch> looks through all its children <Route>
            elements and renders the first one whose path
            matches the current URL. Use a <Switch> any time
            you have multiple routes, but you want only one
            of them to render at a time
          */}
          <Switch>
            <Route exact path="/">
              <Home />
            </Route>
            <Route path="/deposits">
              <Deposits />
            </Route>
            <Route path="/withdrawals">
              <Withdrawals />
            </Route>
            <Route path="/balances">
              <Balances />
            </Route>
          </Switch>
        </Dashboard>
      </ThemeProvider>
    </Router>
  );
}
