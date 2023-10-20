import React, { useState } from 'react';
import axios from 'axios';

const Weather = () => {
    const [city, setCity] = useState('');
    const [weatherData, setWeatherData] = useState(null);

    const handleSearch = () => {
        // Implement weather search logic here
        axios.get(`https://api.openweathermap.org/data/2.5/weather?q=${city}&appid=YOUR_API_KEY`)
            .then(response => {
                setWeatherData(response.data);
            })
            .catch(error => {
                console.error(error);
            });
    };

    return (
        <div>
            <h1>Current Weather and History</h1>
            <input type="text" placeholder="City Name" value={city} onChange={(e) => setCity(e.target.value)} />
            <button onClick={handleSearch}>Search</button>
            {weatherData && (
                <div>
                    {/* Display weather data in tabular form */}
                </div>
            )}
        </div>
    );
};

export default Weather;
