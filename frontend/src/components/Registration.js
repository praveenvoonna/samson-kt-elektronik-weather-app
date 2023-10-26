import React, { useState } from "react";
import { Form, Button } from "react-bootstrap";
import axios from "axios";
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

  const handleSubmit = async (e) => {
    e.preventDefault();

    const data = {
      username: username,
      password: password,
      date_of_birth: dateOfBirth,
    };

    try {
      const response = await axios.post("http://localhost:8080/register", data, {
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
      setError("Registration failed");
    }
  };

  return (
    <>
      <div className="register-container">
        <h2>Register</h2>
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

          {/* date of birth */}
          <Form.Group controlId="formBasicDate">
            <Form.Label>Date of Birth</Form.Label>
            <Form.Control
              type="date"
              value={dateOfBirth}
              onChange={(e) => setDateOfBirth(e.target.value)}
              placeholder="Date of Birth"
            />
          </Form.Group>

          {/* submit button */}
          <Button variant="primary" type="submit">
            Register
          </Button>
        </Form>
      </div>
    </>
  );
};

export default Register;
