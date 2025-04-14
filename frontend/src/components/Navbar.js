import React, { useState, useEffect } from 'react';
import { Link } from 'react-router-dom';
import axios from 'axios';
import '../App.css';

const Navbar = () => {
    const [role, setRole] = useState('');

    useEffect(() => {
        const fetchRole = async () => {
            try {
                const res = await axios.get('/api/role'); // Add an endpoint to fetch user role
                setRole(res.data.role);
            } catch (error) {
                console.error('Error fetching role:', error);
            }
        };
        fetchRole();
    }, []);

    return (
        <nav className="navbar">
            <div className="navbar-brand">
                <Link to="/">
                    <h1>Blinket Gap Filler</h1>
                </Link>
            </div>
            <div className="navbar-links">
                <Link to="/" className="nav-link">Products</Link>
                {role === 'admin' && <Link to="/admin" className="nav-link">Admin Panel</Link>}
                <Link to="/login" className="nav-link">Login</Link>
            </div>
        </nav>
    );
};

export default Navbar;