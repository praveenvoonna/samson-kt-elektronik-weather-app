import React, { useState, useEffect } from "react";
import { Button, TextField, Typography, Container, Box } from "@mui/material";
import axios from "axios";
import { useNavigate } from "react-router-dom";
import "./Login.css";

const Login = () => {
  const [username, setUsername] = useState("");
  const [password, setPassword] = useState("");
  const [token, setToken] = useState("");
  const [errMessage, setError] = useState("");
  const navigate = useNavigate();

  useEffect(() => {
    sessionStorage.removeItem("token");
  }, []);

  const navigateToDashboard = () => {
    navigate("/dashboard");
  };

  const navigateToRegistration = () => {
    navigate("/register");
  };

  const handleSubmit = async (e) => {
    e.preventDefault();

    const data = {
      username: username,
      password: password,
    };

    try {
      const response = await axios.post("http://localhost:8080/login", data, {
        headers: {
          "Content-Type": "application/json",
        },
      });

      console.log(response.data);
      setToken(response.data.token);
      sessionStorage.setItem("token", response.data.token);
      navigateToDashboard();
    } catch (error) {
      console.error("error:", error);
      setError("login failed");
    }
  };

  return (
    <Container component="main" maxWidth="xs">
      <div className="login-container">
        <Typography component="h1" variant="h5">
          Login
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
            autoFocus
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
          <Button
            type="submit"
            fullWidth
            variant="contained"
            sx={{ mt: 3, mb: 2 }}
          >
            Login
          </Button>
          <Button
            fullWidth
            variant="contained"
            sx={{ mb: 2 }}
            onClick={navigateToRegistration}
          >
            Registration
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

export default Login;
