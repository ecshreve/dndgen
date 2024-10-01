import { gql } from "@apollo/client";

export const GET_CHARACTERS = gql`
  query GetCharacter($characterId: ID!) {
    characters(where: {id: $characterId}) {
      edges {
        node {
          id
          name
          age
          level
          proficiencyBonus
          class {
            id
            indx
            name
          }
          race {
            id
            indx
            name
          }
          alignment {
            id
            indx
            name
          }
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
          characterSkills {
            id
            proficient
            skill {
              id
              indx
              name
            }
            characterAbilityScore {
              id
              score
              modifier
              abilityScore {
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
