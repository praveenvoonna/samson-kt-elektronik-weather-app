import React from 'react';
import { BrowserRouter as Router, Switch, Route } from 'react-router-dom';
import Home from './screens/Home';
import Login from './components/Login';
import Registration from './components/Registration';
import Weather from './components/Weather';

function App() {
  return (
    <Router>
      <div>
        <Switch>
          <Route path="/" exact component={Home} />
          <Route path="/login" component={Login} />
          <Route path="/register" component={Registration} />
          <Route path="/weather" component={Weather} />
        </Switch>
      </div>
    </Router>
  );
}

export default App;
