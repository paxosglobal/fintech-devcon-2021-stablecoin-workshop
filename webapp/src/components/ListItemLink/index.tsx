import React from 'react'
import { Link as RouterLink } from 'react-router-dom';

import ListItem from '@material-ui/core/ListItem';
import ListItemIcon from '@material-ui/core/ListItemIcon';
import ListItemText from '@material-ui/core/ListItemText';

interface Props{
    icon?: React.ReactNode;
    primary: string;
    to: string;
}

export const ListItemLink: React.FC<Props> = ({icon, primary, to}) => {
    const renderLink = React.useMemo(
      () => React.forwardRef<HTMLAnchorElement>((itemProps, ref) => <RouterLink to={to} ref={ref} {...itemProps} />),
      [to],
    );
  
    return (
      <li>
        <ListItem button component={renderLink}>
          {icon ? <ListItemIcon>{icon}</ListItemIcon> : null}
          <ListItemText primary={primary} />
        </ListItem>
      </li>
    );
  }