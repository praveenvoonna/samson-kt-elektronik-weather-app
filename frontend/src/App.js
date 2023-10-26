import React, { useState, useEffect } from "react";
import { BrowserRouter as Router, Routes, Route } from "react-router-dom";
import Home from "./screens/Home";
import Login from "./components/Login";
import Registration from "./components/Registration";
import Weather from "./components/Weather";
import Dashboard from "./screens/Dashboard"


function App() {

  return (
    <Router>
      <Routes>
        <Route path="/login" Component={Login} />
        <Route path="/register" Component={Registration} />
        <Route path="/weather" Component={Weather} />
        <Route path="/dashboard" Component={Dashboard} />
        <Route path="/" exact Component={Home} />
      </Routes>
    </Router>
  );
}

export default App;
