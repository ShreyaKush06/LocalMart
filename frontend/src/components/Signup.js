import React, { useState } from 'react';
import axios from 'axios';
import { useNavigate, Link } from 'react-router-dom';

const Signup = () => {
    const [credentials, setCredentials] = useState({ username: '', password: '', confirmPassword: '' });
    const [error, setError] = useState('');
    const navigate = useNavigate();

    const handleSubmit = async (e) => {
        e.preventDefault();
        setError('');
        
        // Validate passwords match
        if (credentials.password !== credentials.confirmPassword) {
            setError("Passwords don't match");
            return;
        }
        
        try {
            await axios.post('/signup', {
                username: credentials.username,
                password: credentials.password
            });
            alert('Signup successful! Please login.');
            navigate('/login');
        } catch (error) {
            if (error.response && error.response.status === 409) {
                setError('Username already exists');
            } else {
                setError('Failed to create account');
            }
        }
    };

    return (
        <div className="login-form">
            <h2>Create Account</h2>
            {error && <div className="error-message">{error}</div>}
            <form onSubmit={handleSubmit}>
                <input
                    type="text"
                    placeholder="Username"
                    value={credentials.username}
                    onChange={(e) => setCredentials({ ...credentials, username: e.target.value })}
                    required
                />
                <input
                    type="password"
                    placeholder="Password"
                    value={credentials.password}
                    onChange={(e) => setCredentials({ ...credentials, password: e.target.value })}
                    required
                />
                <input
                    type="password"
                    placeholder="Confirm Password"
                    value={credentials.confirmPassword}
                    onChange={(e) => setCredentials({ ...credentials, confirmPassword: e.target.value })}
                    required
                />
                <button type="submit">Sign Up</button>
            </form>
            <p className="form-footer">
                Already have an account? <Link to="/login">Login</Link>
            </p>
        </div>
    );
};

export default Signup;