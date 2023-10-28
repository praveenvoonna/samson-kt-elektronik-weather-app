import React from "react";

const Weather = ({ weatherData }) => {
  return (
    <div>
      <h3>Weather Information for {weatherData.name}</h3>
      <p>
        Coordinates: {weatherData.coord.lat} (Latitude), {weatherData.coord.lon}{" "}
        (Longitude)
      </p>
      <p>
        Weather: {weatherData.weather[0].main} -{" "}
        {weatherData.weather[0].description}
      </p>
      <p>Temperature: {weatherData.main.temp} Kelvin</p>
      <p>Feels Like: {weatherData.main.feels_like} Kelvin</p>
      <p>Minimum Temperature: {weatherData.main.temp_min} Kelvin</p>
      <p>Maximum Temperature: {weatherData.main.temp_max} Kelvin</p>
      <p>Pressure: {weatherData.main.pressure} hPa</p>
      <p>Humidity: {weatherData.main.humidity}%</p>
      <p>Wind Speed: {weatherData.wind.speed} m/s</p>
      <p>Cloudiness: {weatherData.clouds.all}%</p>
      <p>
        Sunrise Time:{" "}
        {new Date(weatherData.sys.sunrise * 1000).toLocaleTimeString()}
      </p>
      <p>
        Sunset Time:{" "}
        {new Date(weatherData.sys.sunset * 1000).toLocaleTimeString()}
      </p>
    </div>
  );
};

export default Weather;
