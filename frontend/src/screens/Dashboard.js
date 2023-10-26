import React, { useState, useEffect } from "react";
import { Form, Button } from "react-bootstrap";
import axios from "axios";
import Weather from "../components/Weather";
import { useNavigate } from "react-router-dom";
import "./Dashboard.css";

const Dashboard = () => {
  const [city, setCity] = useState("");
  const [weatherData, setWeatherData] = useState(null);
  const [historyData, setHistoryData] = useState([]);
  const token = sessionStorage.getItem("token");
  const navigate = useNavigate();
  const navigateToHome = () => {
    navigate("/");
  };

  if (!token) {
    navigateToHome()
  }

  const getWeatherData = async () => {
    try {
      const token = sessionStorage.getItem("token");
      const response = await axios.get(`http://localhost:8080/weather?city=${city}`, {
        headers: {
          Authorization: `Bearer ${token}`,
        },
      });
      setWeatherData(response.data);
    } catch (error) {
      console.error("Error:", error);
    }
  };

  const getHistoryData = async () => {
    try {
      // const token = sessionStorage.getItem("token");
      const response = await axios.get("http://localhost:8080/history");
      // {
      //   headers: {
      //     Authorization: `Bearer ${token}`,
      //   },
      // }
      setHistoryData(response.data);
    } catch (error) {
      console.error("Error:", error);
    }
  };

  useEffect(() => {
    getHistoryData();
  }, []);

  const handleSubmit = (e) => {
    e.preventDefault();
    getWeatherData();
  };

  const handleLogout = () => {
    sessionStorage.removeItem("token");
    navigateToHome();
  };

  const convertKelvinToCelsius = (temp) => {
    return (temp - 273.15).toFixed(2);
  };

  return (
    token &&
    <div className="dashboard-container">
      <h2>Dashboard</h2>
      <Button variant="secondary" onClick={handleLogout}>
        Logout
      </Button>
      <Form onSubmit={handleSubmit}>
        {/* city input */}
        <Form.Group controlId="formBasicCity">
          <Form.Label>City</Form.Label>
          <Form.Control
            type="text"
            value={city}
            onChange={(e) => setCity(e.target.value)}
            placeholder="Enter city name"
          />
        </Form.Group>

        {/* get weather button */}
        <Button variant="primary" type="submit">
          Get Weather
        </Button>
      </Form>

      {/* display weather data */}
      {weatherData && (
        <div>
          {weatherData && (
            <div>
              <h3>Weather Data for {city}</h3>
              <table className="table">
                <thead>
                  <tr>
                    <th>Description</th>
                    <th>Temperature (Celsius)</th>
                    <th>Pressure</th>
                    <th>Humidity</th>
                    <th>Wind Speed</th>
                  </tr>
                </thead>
                <tbody>
                  <tr>
                    <td>{weatherData.weather[0].description}</td>
                    <td>{convertKelvinToCelsius(weatherData.main.temp)}</td>
                    <td>{weatherData.main.pressure}</td>
                    <td>{weatherData.main.humidity}</td>
                    <td>{weatherData.wind.speed}</td>
                  </tr>
                </tbody>
              </table>
            </div>
          )}
        </div>
      )}

      {/* display history data */}
      <div>
        <h3>Search History</h3>
        {historyData &&
          historyData.map((item, index) => (
            <div key={index}>
              <p>{item}</p>
            </div>
          ))}
      </div>
    </div>
  );
};

export default Dashboard;
