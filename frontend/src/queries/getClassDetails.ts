import { gql } from "@apollo/client";

export const GET_CLASS_DETAILS = gql`
  fragment ProficiencyDetails on Proficiency {
    id
    indx
    reference
  }

  fragment ProficiencyOptionDetails on ProficiencyChoice {
    id
    desc
    choose
    proficiencies {
      ...ProficiencyDetails
    }
  }

  fragment EquipmentDetails on Equipment {
    id
    indx
    name
    weight
  }

  fragment StartingEquipmentDetails on EquipmentEntry {
    id
    quantity
    equipment {
      ...EquipmentDetails
    }
  }
  
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
