import React, { Component } from 'react';
import logo from './logo.svg';
import './App.css';
import Items from './Items/Items';

class App extends Component {
  render() {
    const item = {
      hash: "1675393889",
      icon: "/common/destiny2_content/icons/dc688a83c2420aa0e48403bfa1df54ee.jpg"
    }

    const items = [];
    for (var i = 0; i <50; i++) {
      items.push(item)
    }

    return (
      <div className="App">
          <Items items={items}/>
      </div>
    );
  }
}

export default App;
