import React, { useState, useEffect } from "react";
import axios from "axios";
import {
  Button,
  TextField,
  Typography,
  Table,
  TableBody,
  TableCell,
  TableContainer,
  TableHead,
  TableRow,
  Paper,
  Container,
  Box,
} from "@mui/material";
import { useNavigate } from "react-router-dom";
import "./Dashboard.css";

const Dashboard = () => {
  const [city, setCity] = useState("");
  const [weatherData, setWeatherData] = useState({});
  const [historyData, setHistoryData] = useState([]);
  const [weatherErrorMessage, setWeatherError] = useState("");
  const [historyErrorMessage, setHistoryError] = useState("");
  const token = sessionStorage.getItem("token");
  const navigate = useNavigate();
  const navigateToHome = () => {
    navigate("/");
  };

  if (!token) {
    navigateToHome();
  }

  const getWeatherData = async () => {
    try {
      const token = sessionStorage.getItem("token");
      const response = await axios.get(
        `http://localhost:8080/weather?city=${city}`,
        {
          headers: {
            Authorization: `Bearer ${token}`,
          },
        }
      );
      setWeatherData(response.data);
      setWeatherError("");
      getHistoryData();
    } catch (error) {
      console.error("Error:", error);
      if (error.response && error.response.data && error.response.data.error) {
        setHistoryError(error.response.data.error);
      } else {
        setHistoryError("Registration failed. Please try again.");
      }
    }
  };

  const getHistoryData = async () => {
    try {
      const token = sessionStorage.getItem("token");
      const response = await axios.get("http://localhost:8080/history", {
        headers: {
          Authorization: `Bearer ${token}`,
        },
      });
      setHistoryData(response.data);
      setHistoryError("");
    } catch (error) {
      console.error("Error:", error);
      if (error.response && error.response.data && error.response.data.error) {
        setHistoryError(error.response.data.error);
      } else {
        setHistoryError("Registration failed. Please try again.");
      }
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

  const deleteHistoryHandller = async (id) => {
    try {
      const response = await axios.delete(
        `http://localhost:8080/history?id=${id}`,
        {
          headers: {
            Authorization: `Bearer ${token}`,
          },
        }
      );
      if (response.status === 200) {
        getHistoryData();
      }
      setHistoryError("");
    } catch (error) {
      console.error("Error:", error);
      if (error.response && error.response.data && error.response.data.error) {
        setHistoryError(error.response.data.error);
      } else {
        setHistoryError("Registration failed. Please try again.");
      }
    }
  };

  return (
    token && (
      <Container maxWidth="sm">
        <Box
          sx={{
            my: 4,
            display: "flex",
            flexDirection: "column",
            alignItems: "center",
          }}
        >
          <Typography variant="h4" component="h2" gutterBottom>
            Dashboard
          </Typography>
          <Button variant="contained" onClick={handleLogout}>
            Logout
          </Button>
          <Box component="form" sx={{ mt: 3 }} onSubmit={handleSubmit}>
            <TextField
              margin="normal"
              required
              fullWidth
              id="city"
              label="City"
              name="city"
              value={city}
              onChange={(e) => setCity(e.target.value)}
              placeholder="Enter city name"
            />
            <Button type="submit" variant="contained" sx={{ mt: 2, mb: 2 }}>
              Get Weather
            </Button>
            {weatherErrorMessage && (
              <Typography variant="body2" color="error">
                {weatherErrorMessage}
              </Typography>
            )}
          </Box>

          {weatherData && weatherData.weather && (
            <Box sx={{ my: 4 }}>
              <Typography variant="h5" component="h3" gutterBottom>
                Weather Data for {city}
              </Typography>
              <TableContainer component={Paper}>
                <Table sx={{ minWidth: 650 }} aria-label="weather data table">
                  <TableHead>
                    <TableRow>
                      <TableCell>Description</TableCell>
                      <TableCell>Temperature (Celsius)</TableCell>
                      <TableCell>Pressure</TableCell>
                      <TableCell>Humidity</TableCell>
                      <TableCell>Wind Speed</TableCell>
                    </TableRow>
                  </TableHead>
                  <TableBody>
                    <TableRow key={weatherData.id}>
                      <TableCell>
                        {weatherData.weather[0].description}
                      </TableCell>
                      <TableCell>
                        {convertKelvinToCelsius(weatherData.main.temp)}
                      </TableCell>
                      <TableCell>{weatherData.main.pressure}</TableCell>
                      <TableCell>{weatherData.main.humidity}</TableCell>
                      <TableCell>{weatherData.wind.speed}</TableCell>
                    </TableRow>
                  </TableBody>
                </Table>
              </TableContainer>
            </Box>
          )}

          {historyData && (
            <Box sx={{ my: 4 }}>
              <Typography variant="h5" component="h3" gutterBottom>
                Search History
              </Typography>
              <TableContainer component={Paper}>
                <Table sx={{ minWidth: 650 }} aria-label="history data table">
                  <TableHead>
                    <TableRow>
                      <TableCell>ID</TableCell>
                      <TableCell>City Name</TableCell>
                      <TableCell>Search Time</TableCell>
                      <TableCell>Action</TableCell>{" "}
                      {/* New Table Header Column for Actions */}
                    </TableRow>
                  </TableHead>
                  <TableBody>
                    {historyData.map((item) => (
                      <TableRow key={item.id}>
                        <TableCell>{item.id}</TableCell>
                        <TableCell>{item.city_name}</TableCell>
                        <TableCell>
                          {new Date(item.search_time).toLocaleString()}
                        </TableCell>
                        <TableCell>
                          <Button
                            variant="contained"
                            color="error"
                            onClick={() => {
                              deleteHistoryHandller(item.id);
                            }}
                          >
                            Delete
                          </Button>
                        </TableCell>
                      </TableRow>
                    ))}
                  </TableBody>
                </Table>
              </TableContainer>
              {historyErrorMessage && (
                <Typography variant="body2" color="error">
                  {historyErrorMessage}
                </Typography>
              )}
            </Box>
          )}
        </Box>
      </Container>
    )
  );
};

export default Dashboard;
