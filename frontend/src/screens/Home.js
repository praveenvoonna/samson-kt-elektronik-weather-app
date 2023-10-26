import React from 'react';
import { useNavigate } from 'react-router-dom';
import './Home.css';

const Home = () => {
    const navigate = useNavigate();
    const navigateToLogin = () => {
        navigate("/login");
    };

    const navigateToRegistration = () => {
        navigate("/register");
    };

    return (
        <div className="home-container">
            <h1 className="home-title">Welcome to the Weather App!</h1>
            <div className="home-buttons">
                <button className="home-button" onClick={navigateToLogin}>Login</button>
                <button className="home-button" onClick={navigateToRegistration}>Register</button>
            </div>
        </div>
    );
};

export default Home;
