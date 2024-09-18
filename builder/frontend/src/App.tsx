import React from 'react';
import './App.css';
import CharacterCreationPage from './components/CharacterCreationPage';


const App: React.FC = () => {
  return (
    <div className="App">
      <h1>Hello, World!</h1>
      <hr />
      <CharacterCreationPage />
    </div>
  );
};

export default App;