import React, { useState } from 'react';
import axios from 'axios';

const Login = () => {
    const [credentials, setCredentials] = useState({ email: '', password: '' });

    const handleSubmit = async (e) => {
        e.preventDefault();
        try {
            await axios.post('/login', credentials);
            alert('Login successful!');
        } catch (error) {
            alert('Login failed!');
        }
    };

    return (
        <div className="login-form">
            <h2>Student Login</h2>
            <form onSubmit={handleSubmit}>
                <input
                    type="email"
                    placeholder="Email"
                    onChange={(e) => setCredentials({...credentials, email: e.target.value})}
                />
                <input
                    type="password"
                    placeholder="Password"
                    onChange={(e) => setCredentials({...credentials, password: e.target.value})}
                />
                <button type="submit">Login</button>
            </form>
        </div>
    );
};

export default Login;