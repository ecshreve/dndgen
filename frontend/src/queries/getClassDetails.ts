import { gql } from "@apollo/client";

import { ProficiencyDetails, ProficiencyOptionDetails, StartingEquipmentDetails } from "./fragments";

export const GET_CLASS_DETAILS = gql`
  ${ProficiencyDetails}
  ${ProficiencyOptionDetails}
  ${StartingEquipmentDetails}
  fragment ClassDetails on Class {
    id
    indx
    name
    savingThrows {
      id
      indx
    }
    proficiencies {
      ...ProficiencyDetails
    }
    proficiencyOptions {
      ...ProficiencyOptionDetails
    }
    startingEquipment {
      ...StartingEquipmentDetails
    }
  }

  query GetClassDetails {
    classes {
      ...ClassDetails
    }
  }
`;
