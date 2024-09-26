import { gql } from '@apollo/client';

export const GET_CLASSES = gql`
  query GetClasses {
    classes {
      edges {
        node {
          id
          indx
          name
          savingThrows {
            indx
          }
          proficiencies {
            indx
          }
          proficiencyOptions {
            desc
          }
          startingEquipment {
            quantity
            equipment {
              indx
            }
          }
        }
      }
    }
  }
`;
