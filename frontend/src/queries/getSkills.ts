import { gql } from "@apollo/client";

export const GET_SKILLS = gql`
  query GetSkills($characterId: ID!) {
    characters(where: {id: $characterId}) {
      edges {
        node {
          id
          proficiencyBonus
          characterSkills {
            id
            proficient
            skill {
              id
              indx
              name
            }
            characterAbilityScore {
              id
              score
              modifier
              abilityScore {
                id
                indx
                name
            }
          }
        }
      }
    }
  }
}
`;
