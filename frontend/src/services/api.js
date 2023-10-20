import axios from 'axios';

const API_KEY = 'YOUR_API_KEY';

const getWeatherData = (city) => {
    return axios.get(`https://api.openweathermap.org/data/2.5/weather?q=${city}&appid=${API_KEY}`);
};

export { getWeatherData };
