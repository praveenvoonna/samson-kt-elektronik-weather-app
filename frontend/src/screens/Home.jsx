import React from "react";
import { Button, Typography, Container, Box } from "@mui/material";
import { useNavigate } from "react-router-dom";

const Home = () => {
  const navigate = useNavigate();

  const navigateToLogin = () => {
    navigate("/login");
  };

  const navigateToRegistration = () => {
    navigate("/register");
  };

  return (
    <Container component="main" maxWidth="sm">
      <Box
        sx={{
          display: "flex",
          flexDirection: "column",
          alignItems: "center",
          justifyContent: "center",
          height: "100vh",
          backgroundColor: "#f5f5f5",
          padding: 4,
        }}
      >
        <Typography
          variant="h3"
          component="h1"
          gutterBottom
          sx={{ fontSize: "2rem", marginBottom: "2rem" }}
        >
          Welcome to the Weather App!
        </Typography>
        <Box sx={{ display: "flex", marginBottom: 4 }}>
          <Button
            variant="contained"
            onClick={navigateToLogin}
            sx={{
              padding: "10px 20px",
              margin: "0 10px",
              fontSize: "1rem",
              backgroundColor: "#007bff",
              color: "white",
              border: "none",
              borderRadius: "4px",
              cursor: "pointer",
              transition: "background-color 0.3s",
              "&:hover": {
                backgroundColor: "#0056b3",
              },
            }}
          >
            Login
          </Button>
          <Button
            variant="contained"
            onClick={navigateToRegistration}
            sx={{
              padding: "10px 20px",
              margin: "0 10px",
              fontSize: "1rem",
              backgroundColor: "#007bff",
              color: "white",
              border: "none",
              borderRadius: "4px",
              cursor: "pointer",
              transition: "background-color 0.3s",
              "&:hover": {
                backgroundColor: "#0056b3",
              },
            }}
          >
            Register
          </Button>
        </Box>
      </Box>
    </Container>
  );
};

export default Home;
