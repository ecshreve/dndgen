import { gql } from "@apollo/client";

export const UPDATE_CHARACTER = gql`
  mutation UpdateCharacter($updateCharacterId: ID!, $input: UpdateCharacterInput!) {
    updateCharacter(id: $updateCharacterId, input: $input) {
      id
      name
      age
      level
      race {
        id
        indx
        name
      }
      class {
        id
        indx
        name
      }
    }
  }
`;
