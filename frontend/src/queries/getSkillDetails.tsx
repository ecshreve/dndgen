import { gql } from '@apollo/client';

export const GET_SKILL_DETAILS = gql`
  query GetSkillDetails {
    skills {
      id
      indx
      name
      desc
      abilityScore {
        id
        indx
        name
      }
    }
  }
`;