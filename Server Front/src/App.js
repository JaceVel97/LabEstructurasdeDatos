import './App.css';
import { BrowserRouter as Router, Switch, Route } from "react-router-dom";
import React from 'react';

import 'bootswatch/dist/slate/bootstrap.min.css'

import Principal from "./components/Principal";
import Ejemplo1 from "./components/Ejemplo1";
import Ejemplo2 from "./components/Ejemplo2";

import { Navbar } from './components/Navbar';
function App() {
  return (
    <div>
      <Router>
        <div>
          <Navbar/>
          <Switch>
            <div>
              <Route path="/home" component={Principal}/>
              <Route path="/exa1" component={Ejemplo1}/>
              <Route path="/exa2" component={Ejemplo2}/>
            </div>
          </Switch>
        </div>
      </Router>
    </div>
  );
}

export default App;
