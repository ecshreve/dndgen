import { gql } from "@apollo/client";

export const ProficiencyDetails = gql`
  fragment ProficiencyDetails on Proficiency {
    id
    indx
    reference
  }
`;

export const ProficiencyOptionDetails = gql`
  fragment ProficiencyOptionDetails on ProficiencyChoice {
    id
    desc
    choose
    proficiencies {
      ...ProficiencyDetails
    }
  }
`;

export const EquipmentDetails = gql`
  fragment EquipmentDetails on Equipment {
    id
    indx
    name
    weight
  }
`;

export const StartingEquipmentDetails = gql`
  fragment StartingEquipmentDetails on EquipmentEntry {
    id
    quantity
    equipment {
      id
      indx
      name
      weight
    }
  }
`;
