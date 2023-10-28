import React, { useState, useEffect } from "react";
import axios from "axios";
import { Button, TextField, Typography, Container, Box } from "@mui/material";
import { useNavigate } from "react-router-dom";
import "./Registration.css";

const Register = () => {
  const [username, setUsername] = useState("");
  const [password, setPassword] = useState("");
  const [dateOfBirth, setDateOfBirth] = useState("");
  const [token, setToken] = useState("");
  const [errMessage, setError] = useState("");

  const navigate = useNavigate();

  const navigateToDashboard = () => {
    navigate("/dashboard");
  };

  const navigateToLogin = () => {
    navigate("/login");
  };

  useEffect(() => {
    sessionStorage.removeItem("token");
  }, []);

  const handleSubmit = async (e) => {
    e.preventDefault();

    const data = {
      username: username,
      password: password,
      date_of_birth: dateOfBirth,
    };

    try {
      const response = await axios.post(
        "http://localhost:8080/register",
        data,
        {
          headers: {
            "Content-Type": "application/json",
          },
        }
      );

      console.log(response.data);
      setToken(response.data.token);
      sessionStorage.setItem("token", response.data.token);
      navigateToDashboard();
    } catch (error) {
      console.error("error:", error);
      setError("registration failed");
    }
  };

  return (
    <Container component="main" maxWidth="xs">
      <div className="register-container">
        <Typography component="h2" variant="h5" gutterBottom>
          Register
        </Typography>
        <Box component="form" onSubmit={handleSubmit} sx={{ mt: 3 }}>
          <TextField
            margin="normal"
            required
            fullWidth
            id="username"
            label="Username"
            name="username"
            autoComplete="username"
            value={username}
            onChange={(e) => setUsername(e.target.value)}
          />
          <TextField
            margin="normal"
            required
            fullWidth
            name="password"
            label="Password"
            type="password"
            id="password"
            autoComplete="current-password"
            value={password}
            onChange={(e) => setPassword(e.target.value)}
          />
          <TextField
            margin="normal"
            required
            fullWidth
            // label="Date of Birth"
            type="date"
            value={dateOfBirth}
            onChange={(e) => setDateOfBirth(e.target.value)}
          />
          <Button
            type="submit"
            fullWidth
            variant="contained"
            sx={{ mt: 3, mb: 2 }}
          >
            Register
          </Button>
          <Button
            fullWidth
            variant="contained"
            sx={{ mb: 2 }}
            onClick={navigateToLogin}
          >
            Login
          </Button>
          {errMessage && (
            <Typography variant="body2" color="error">
              {errMessage}
            </Typography>
          )}
        </Box>
      </div>
    </Container>
  );
};

export default Register;
