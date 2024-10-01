import { gql } from "@apollo/client";

export const GET_CLASS_DETAILS = gql`
  query GetClassDetails {
    classes {
      id
      indx
      name
      savingThrows {
        id
        indx
      }
      proficiencies {
        id
        indx
        reference
      }
      proficiencyOptions {
        id
        desc
        choose
        proficiencies {
          id
          indx
          reference
        }
      }
      startingEquipment {
        id
        quantity
        equipment {
          id
          indx
        }
      }
    }
  }
`;