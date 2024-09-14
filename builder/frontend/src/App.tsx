import React from 'react';
import './App.css';
import { graphql } from './gql';
import { useQuery } from '@apollo/client';

const GET_CLASSES = graphql(`
  query GetClasses {
    classes {
      id
      name
    }
  }
`);


const App: React.FC = () => {
  const { data } = useQuery(GET_CLASSES);
  return (
    <div className="App">
      <h1>Hello, World!</h1>
      <hr />
      {data && data.classes.map((c) => (
        <div key={c.id}>{c.name}</div>
      ))}
    </div>
  );
};

export default App;