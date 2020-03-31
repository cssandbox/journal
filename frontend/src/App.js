import React from 'react';
import logo from './logo.svg';
import Journal from './features/journal/Journal';
import JournalList from './features/journal/JournalList'
import './App.css';

function App() {
  return (
    <div className="App">
      <header className="App-header">
        <img src={logo} className="App-logo" alt="logo" />
        <Journal />
        <JournalList />
      </header>
    </div>
  );
}

export default App;
