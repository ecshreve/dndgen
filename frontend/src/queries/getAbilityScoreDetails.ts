import { gql } from '@apollo/client';

export const GET_ABILITY_SCORE_DETAILS = gql`
  query GetAbilityScoreDetails {
    abilityScores {
      indx
      name
      fullName
      desc
    }
  }
`;
