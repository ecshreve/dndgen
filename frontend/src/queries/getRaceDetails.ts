import { gql } from "@apollo/client";
import { ProficiencyDetails, ProficiencyOptionDetails } from "./fragments";

export const GET_RACE_DETAILS = gql`
  ${ProficiencyDetails}
  ${ProficiencyOptionDetails}
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
      ...ProficiencyDetails
    }
    startingProficiencyOptions {
      ...ProficiencyOptionDetails
    }
  }

  query GetRaceDetails {
    races {
      ...RaceDetails
    }
  }
`;
