import { gql } from "@apollo/client";

export const GET_CHARACTERS = gql`
  query GetCharacters {
    characters {
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
          characterSkills {
          modifier
          proficient
          skill {
            indx
            abilityScore {
              indx
            }
          }
        }
        characterAbilityScores {
          abilityScore {
            indx
            skills {
              indx
              characterSkills {
                proficient
                modifier
              }
            }
          }
          modifier
          score
          }
        }
      }
    }
  }
`;
