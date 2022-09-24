import React from 'react';
import ReactDOM from 'react-dom/client';
import { createBrowserRouter, RouterProvider } from 'react-router-dom';
import App from './App';
import './App.css';
import Login from './components/Login';
import Register from './components/Register';
import './index.css';
import * as serviceWorker from './serviceWorker';

const router = createBrowserRouter([{
  element: <App />,
  children: [
    {
      path: '/',
      element: <h1>Welcome, register and login</h1>,
    },
    {
      path: "/login",
      element: <Login />,
      // loader: rootLoader,
    },
    {
      path: "/register",
      element: <Register />,
      // loader: rootLoader,
    }],
}]);

ReactDOM.createRoot(document.getElementById('root') as HTMLElement).render(
  <React.StrictMode>
    <RouterProvider router={router} />
  </React.StrictMode>
)

// If you want your app to work offline and load faster, you can change
// unregister() to register() below. Note this comes with some pitfalls.
// Learn more about service workers: https://bit.ly/CRA-PWA
serviceWorker.unregister();

