import React from 'react';
import { Link } from 'react-router-dom';
import '../App.css';

const Navbar = () => {
  return (
    <nav className="navbar">
      <div className="navbar-brand">
        <Link to="/">
          <h1>Blinket Gap Filler</h1>
        </Link>
      </div>
      <div className="navbar-links">
        <Link to="/" className="nav-link">Products</Link>
        <Link to="/login" className="nav-link">Login</Link>
      </div>
    </nav>
  );
};

export default Navbar;