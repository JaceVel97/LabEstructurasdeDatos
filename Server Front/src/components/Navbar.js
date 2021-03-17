import React from 'react';
import { Link } from 'react-router-dom';

export const Navbar = () => (
    <nav className="navbar navbar-expand-lg " style={{backgroundColor: '#2874a6'}}>
        <Link className="navbar-brand" to="/home">Reactjs</Link>
        <button className="navbar-toggler" type="button" data-toggle="collapse" data-target="#navbarNav" aria-controls="navbarNav" aria-expanded="false" aria-label="Toggle navigation">
          <span className="navbar-toggler-icon" />
        </button>
        <div className="collapse navbar-collapse" id="navbarNav">
          <ul className="navbar-nav">
            <li className="nav-item">
              <Link className="nav-link" to="/exa1">Ejemplo 1</Link>
            </li>
            <li className="nav-item">
              <Link className="nav-link" to="/exa2">Ejemplo 2</Link>
            </li>
          </ul>
        </div>
      </nav>
)