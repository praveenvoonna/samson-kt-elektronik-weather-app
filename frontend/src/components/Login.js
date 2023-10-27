import React, { useState, useEffect } from "react";
import { Form, Button } from "react-bootstrap";
import axios from "axios";
import { useNavigate } from "react-router-dom";
import "./Login.css";

const Login = () => {
  const [username, setUsername] = useState("");
  const [password, setPassword] = useState("");
  const [token, setToken] = useState("");
  const [errMessage, setError] = useState("");

  useEffect(() =>  {
    sessionStorage.removeItem("token");
  }, []);

  const navigate = useNavigate();
  const navigateToDashboard = () => {
    navigate("/dashboard");
  };
  const navigateToRegistration = () => {
    navigate("/register")
  }
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
      console.error("Error:", error);
      setError("Login failed");
    }
  };

  return (
    <>
      <div className="login-container">
        <h2>Login</h2>
        <Form onSubmit={handleSubmit}>
          {/* username */}
          <Form.Group controlId="formBasicUsername">
            <Form.Label>Username</Form.Label>
            <Form.Control
              type="text"
              name="username"
              value={username}
              onChange={(e) => setUsername(e.target.value)}
              placeholder="Enter username"
            />
          </Form.Group>

          {/* password */}
          <Form.Group controlId="formBasicPassword">
            <Form.Label>Password</Form.Label>
            <Form.Control
              type="password"
              name="password"
              value={password}
              onChange={(e) => setPassword(e.target.value)}
              placeholder="Password"
            />
          </Form.Group>

          {/* submit button */}
          <Button variant="primary" type="submit">
            Login
          </Button>
          <Button variant="primary" onClick={navigateToRegistration}>
            Registration
          </Button>
          {errMessage ? <h1>{errMessage}</h1> : null}
        </Form>
      </div>
    </>
  );
};

export default Login;
