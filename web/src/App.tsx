import React, { Component } from 'react';
import logo from './logo.svg';
import './App.css';
import Items from './Items/Items';

class App extends Component {
  render() {
    return (
      <div className="App">
        <header className="App-header">
          <img src={logo} className="App-logo" alt="logo" />
          <p>
            Edit <code>src/App.js</code> and save to reload. Check?
          </p>
          <Items items={[
            {
              hash: "1675393889",
              icon: "/common/destiny2_content/icons/dc688a83c2420aa0e48403bfa1df54ee.jpg"
            },
            {
              hash: "1675393889",
              icon: "/common/destiny2_content/icons/dc688a83c2420aa0e48403bfa1df54ee.jpg"
            },
            {
              hash: "1675393889",
              icon: "/common/destiny2_content/icons/dc688a83c2420aa0e48403bfa1df54ee.jpg"
            }
          ]}/>
          <a
            className="App-link"
            href="https://reactjs.org"
            target="_blank"
            rel="noopener noreferrer"
          >
            Learn React
          </a>
        </header>
      </div>
    );
  }
}

export default App;
