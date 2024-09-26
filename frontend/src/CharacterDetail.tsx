import React from 'react';

import { useQuery } from '@apollo/client';

import { gql } from './__generated__/gql';

const GET_CHARACTER = gql(/* GraphQL */ `
  query GetCharacter($characterId: ID!) {
    node(id: $characterId) {
      ... on Character {
        id
        name
        race {
          id
          indx
        }
        class {
          id
          indx
        }
      }
    }
  }
`);

const CharacterDetail = () => {
  // our query's result, data, is typed!
  const { loading, data } = useQuery(
    GET_CHARACTER,
    { variables: { characterId: '21474836481' } }
  );
  return (
    <div>
      <h3>Character Detail</h3>
      {loading ? (
        <p>Loading ...</p>
      ) : (
        <div>
          {data && data.node && data.node.__typename === 'Character' && (
            <div>
              <p>id: {data.node.id}</p>
              <p>name: {data.node.name}</p>
              <p>race: {data.node.race.indx}</p>
              <p>class: {data.node.class.indx}</p>
            </div>
          )}
        </div>
      )}
    </div>
  );
};

export default CharacterDetail;