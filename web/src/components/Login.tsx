import React, { useEffect, useReducer } from 'react';
import { Outlet } from 'react-router-dom';

import Button from '@material-ui/core/Button';
import Card from '@material-ui/core/Card';
import CardActions from '@material-ui/core/CardActions';
import CardContent from '@material-ui/core/CardContent';
import CardHeader from '@material-ui/core/CardHeader';
import { createStyles, makeStyles, Theme } from '@material-ui/core/styles';
import TextField from '@material-ui/core/TextField';

import UserService from '@services/user.service';

// connects to grpcwebproxy which is listening on port 8080
// TODO: configurize eventually..
const user = new UserService();

// Reference: https://github.com/creativesuraj/react-material-ui-login/blob/master/src/components/Login.tsx

// TODO: Learn more about colors and material-ui: https://material-ui.com/customization/color/
const useStyles = makeStyles((theme: Theme) => createStyles({
  container: {
    display: 'flex',
    flexWrap: 'wrap',
    width: 400,
    margin: `${theme.spacing(0)} auto`,
  },
  loginBtn: {
    marginTop: theme.spacing(2),
    flexGrow: 1,
    background: '#339933',
    '&:hover': {
      backgroundColor: '#43a943',
    },
  },
  header: {
    textAlign: 'center',
    background: '#282c34',
    color: '#fff',
  },
  card: {
    marginTop: theme.spacing(10),
  },
}));

type State = {
  username: string
  password: string
  isButtonDisabled: boolean
  helperText: string
  isError: boolean
};

const initialState: State = {
  username: '',
  password: '',
  isButtonDisabled: true,
  helperText: '',
  isError: false,
};

// TODO: 'type' should be an enum
type Action = { type: 'setUsername', payload: string }
  | { type: 'setPassword', payload: string }
  | { type: 'setIsButtonDisabled', payload: boolean }
  | { type: 'loginSuccess', payload: string }
  | { type: 'loginFailed', payload: string }
  | { type: 'setIsError', payload: boolean };

const reducer = (state: State, action: Action): State => {
  switch (action.type) {
    case 'setUsername':
      return { ...state, username: action.payload };
    case 'setPassword':
      return { ...state, password: action.payload };
    case 'setIsButtonDisabled':
      return { ...state, isButtonDisabled: action.payload };
    case 'loginSuccess':
      return { ...state, helperText: action.payload, isError: false };
    case 'loginFailed':
      return { ...state, helperText: action.payload, isError: true };
    case 'setIsError':
      return { ...state, isError: action.payload };
    default:
      throw new Error('login.tsx: unknown reducer action');
  }
};

// Login component
export default function Login() {
  const classes = useStyles();
  const [state, dispatch] = useReducer(reducer, initialState);

  useEffect(() => {
    if (state.username.trim() && state.password.trim()) {
      dispatch({ type: 'setIsButtonDisabled', payload: false });
    } else {
      dispatch({ type: 'setIsButtonDisabled', payload: true });
    }
  }, [state.username, state.password]);
  const handleLogin = async () => {
    const response = await user.getUser('1');
    dispatch({ type: 'loginSuccess', payload: `returned '${response.toString()}' from user-service` });

    // // TODO: handle logins + session management
    // if (state.username === 'abc@email.com' && state.password === 'password') {
    //   dispatch({ type: 'loginSuccess', payload: `returned ${name}` });
    // } else {
    //   dispatch({ type: 'loginFailed', payload: 'Incorrect username or password' });
    // }
  };

  const handleKeyPress = (event: React.KeyboardEvent) => {
    if (event.keyCode === 13 || event.which === 13) {
      state.isButtonDisabled || handleLogin();
    }
  };

  const handleUsernameChange: React.ChangeEventHandler<HTMLInputElement> = (event) => {
    dispatch({ type: 'setUsername', payload: event.target.value });
  };

  const handlePasswordChange: React.ChangeEventHandler<HTMLInputElement> = (event) => {
    dispatch({ type: 'setPassword', payload: event.target.value });
  };
  return (
    <form className={classes.container} noValidate autoComplete="off">
      <h3>SIMPLE TEST: use any user/pw and see grpc response below</h3>
      <Card className={classes.card}>
        <CardHeader className={classes.header} title="Login" />
        <CardContent>
          <div>
            <TextField
              error={state.isError}
              fullWidth
              id="username"
              type="email"
              label="Username"
              placeholder="Username"
              margin="normal"
              onChange={handleUsernameChange}
              onKeyPress={handleKeyPress}
            />
            <TextField
              error={state.isError}
              fullWidth
              id="password"
              type="password"
              label="Password"
              placeholder="Password"
              margin="normal"
              helperText={state.helperText}
              onChange={handlePasswordChange}
              onKeyPress={handleKeyPress}
            />
          </div>
        </CardContent>
        <CardActions>
          <Button
            variant="contained"
            size="large"
            color="primary"
            className={classes.loginBtn}
            onClick={handleLogin}
            disabled={state.isButtonDisabled}>
            Login
          </Button>
        </CardActions>
      </Card>
      <Outlet />
    </form>
  );
}
