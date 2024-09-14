// CharacterCreationPage.tsx

import { useQuery } from '@apollo/client';
import React from 'react';
import { graphql } from '../gql/gql';
import CharacterCreationForm from './CharacterCreationForm';
import { CharacterFormValues } from './types';

const GET_CLASS_OPTIONS_QUERY = graphql(`
  query GetClasses {
    classes {
      id
      name
    }
  }
`);

const GET_RACE_OPTIONS_QUERY = graphql(`
  query GetRaces {
    races {
      id
      name
    }
  }
`);

const GET_ALIGNMENT_OPTIONS_QUERY = graphql(`
  query GetAlignments {
    alignments {
      id
      name
    }
  }
`);

const GET_ABILITY_SCORE_OPTIONS_QUERY = graphql(`
  query GetAbilityScores {
    abilityScores {
      id
      name
    }
  }
`);


const CharacterCreationPage: React.FC = () => {
  const handleSubmit = (values: CharacterFormValues) => {
    // Handle form submission, e.g., send to backend
    console.log('Character created:', values);
  };

  // Get all the options from Graphql
  const { data: classesData } = useQuery(GET_CLASS_OPTIONS_QUERY);
  const { data: racesData } = useQuery(GET_RACE_OPTIONS_QUERY);
  const { data: alignmentsData } = useQuery(GET_ALIGNMENT_OPTIONS_QUERY);
  const { data: abilityScoresData } = useQuery(GET_ABILITY_SCORE_OPTIONS_QUERY);

  const classes = classesData?.classes || [];
  const races = racesData?.races || [];
  const alignments = alignmentsData?.alignments || [];
  const abilityScores = abilityScoresData?.abilityScores || [];

  return (
    <div>
      <h1>Create a New Character</h1>
      <CharacterCreationForm
        classes={classes}
        races={races}
        alignments={alignments}
        onSubmit={handleSubmit}
      />
    </div>
  );
};

export default CharacterCreationPage;
