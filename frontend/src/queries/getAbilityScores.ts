import { gql } from "@apollo/client";

export const GET_ABILITY_SCORES = gql`
  query GetAbilityScores($characterId: ID!) {
    characters(where: {id: $characterId}) {
      edges {
        node {
          id
          characterAbilityScores {
            id
            score
            modifier
            abilityScore {
              id
              indx
              name
            }
            characterSkills {
              id
              proficient
              skill {
                id
                indx
                name
              }
            }
          }
        }
      }
    }
  }
`;
