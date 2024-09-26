import { gql } from "@apollo/client";

export const GET_RACES = gql`
  query GetRaces {
    races {
      id
      indx
      name
      size
      speed
      size
      sizeDesc
      ageDesc
      alignmentDesc
      abilityBonuses {
        abilityScore {
          indx
        }
        bonus
      }
      languageDesc
      languages {
        name
      }
      startingProficiencies {
        indx
      }
      startingProficiencyOptions {
        desc
        choose
        proficiencies {
          indx
        }
      }
    }
  }
`;
