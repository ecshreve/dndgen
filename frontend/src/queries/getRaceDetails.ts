import { gql } from "@apollo/client";

export const GET_RACE_DETAILS = gql`
  fragment RaceDetails on Race {
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
      reference
    }
    startingProficiencyOptions {
      id
      desc
      choose
      proficiencies {
        id
        indx
        reference
      }
    }
  }

  query GetRaceDetails {
    races {
      ...RaceDetails
    }
  }
`;
