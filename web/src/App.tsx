import React from 'react';
import { Link, Outlet } from 'react-router-dom';
import './App.css';
import logo from './logo.svg';

/* TODO: 
  - break apart nav and layouts
  - switch to kebab-cased files
  - deeper folder structure

  - register form call backend to register
  - login for logs in
  - logged in page
*/

// Main app/layout
export default function App() {
  return (
    <div className="App">
      <header className="App-header">
        <img src={logo} className="App-logo" alt="logo" />
        <p>
          Edit <code>src/App.tsx</code> and save to reload.
        </p>
        <nav>
          <ul>
            <li>
              <Link to={`/`} className="App-link">Home</Link>
            </li>
            <li>
              <Link to={`/login`} className="App-link">Login</Link>
            </li>
            <li>
              <Link to={`/register`} className="App-link">Register</Link>
            </li>
          </ul>
        </nav>
      </header>
      <main>
        <Outlet />
      </main>
    </div>
  );
}
