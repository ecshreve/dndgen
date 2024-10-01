import { gql } from "@apollo/client";

export const GET_RACE_DETAILS = gql`
  query GetRaceDetails {
    races {
      id
      indx
      name
      ageDesc
      alignmentDesc
      size
      sizeDesc
      speed
      languageDesc
      startingProficiencies {
        id
        indx
        name
      }
      startingProficiencyOptions {
        id
        desc
        choose
        proficiencies {
          id
          indx
          name
        }
      }
    }
  }
`;
