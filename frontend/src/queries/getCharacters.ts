import { gql } from '@apollo/client';

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
          characterSkills {
            modifier
            proficient
            skill {
              indx
            }
          }
          proficiencies {
            indx
          }
        }
      }
    }
  }
`;
