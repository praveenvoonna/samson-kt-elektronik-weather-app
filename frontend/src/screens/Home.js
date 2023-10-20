import React from 'react';
import { Link } from 'react-router-dom';

const Home = () => {
    return (
        <div>
            <h1>Welcome to the Weather App!</h1>
            <Link to="/login">Login</Link>
            <Link to="/register">Register</Link>
            <Link to="/weather">Check Weather</Link>
        </div>
    );
};

export default Home;
